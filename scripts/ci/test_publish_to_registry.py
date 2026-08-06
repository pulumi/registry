#!/usr/bin/env python3
"""Unit tests for publish_to_registry.py"""

import json
import tempfile
import unittest
from pathlib import Path
from unittest.mock import MagicMock, patch

from publish_to_registry import (
    Config,
    PackageFailure,
    PublishOutcome,
    SpecResult,
    build_package_spec,
    build_specs,
    get_changed_packages,
    parse_publish_output,
    publish_with_retry,
    report_permanent_failures,
)


class TestGetChangedPackages(unittest.TestCase):
    @patch("publish_to_registry.subprocess.run")
    def test_returns_changed_files(self, mock_run):
        mock_run.return_value = MagicMock(
            stdout="themes/default/data/registry/packages/aws.yaml\n"
                   "themes/default/data/registry/packages/gcp.yaml\n",
            returncode=0,
        )

        result = get_changed_packages(Path("/repo"))

        self.assertEqual(result, [
            "themes/default/data/registry/packages/aws.yaml",
            "themes/default/data/registry/packages/gcp.yaml",
        ])

    @patch("publish_to_registry.subprocess.run")
    def test_returns_empty_on_no_changes(self, mock_run):
        mock_run.return_value = MagicMock(stdout="", returncode=0)

        result = get_changed_packages(Path("/repo"))

        self.assertEqual(result, [])

    @patch("publish_to_registry.subprocess.run")
    def test_returns_empty_on_git_error(self, mock_run):
        from subprocess import CalledProcessError
        mock_run.side_effect = CalledProcessError(1, "git")

        result = get_changed_packages(Path("/repo"))

        self.assertEqual(result, [])


class TestBuildPackageSpec(unittest.TestCase):
    def setUp(self):
        self.temp_dir = tempfile.TemporaryDirectory()
        self.repo_root = Path(self.temp_dir.name)

        publishers_dir = self.repo_root / "tools/resourcedocsgen/pkg/publishers"
        publishers_dir.mkdir(parents=True)
        with open(publishers_dir / "publisher-names.json", "w") as f:
            json.dump({"Pulumi": "pulumi", "Pulumiverse": "pulumiverse"}, f)

        self.packages_dir = self.repo_root / "themes/default/data/registry/packages"
        self.packages_dir.mkdir(parents=True)

        self.publishers = {"Pulumi": "pulumi", "Pulumiverse": "pulumiverse"}

    def tearDown(self):
        self.temp_dir.cleanup()

    def _write_yaml(self, name: str, content: str):
        with open(self.packages_dir / f"{name}.yaml", "w") as f:
            f.write(content)

    def test_builds_pulumi_spec(self):
        self._write_yaml("aws", """
version: "6.50.0"
publisher: Pulumi
schema_file_url: "https://raw.githubusercontent.com/pulumi/pulumi-aws/..."
""")

        result = build_package_spec(
            "themes/default/data/registry/packages/aws.yaml",
            self.repo_root,
            self.publishers,
        )

        self.assertEqual(result.spec, "pulumi/pulumi/aws@6.50.0")
        self.assertIsNone(result.error)
        self.assertFalse(result.skipped)

    def test_builds_opentofu_spec(self):
        self._write_yaml("elasticstack", """
version: "v0.14.3"
publisher: Pulumi
schema_file_url: "https://example.com/registry.opentofu.org/schema.json"
""")

        result = build_package_spec(
            "themes/default/data/registry/packages/elasticstack.yaml",
            self.repo_root,
            self.publishers,
        )

        self.assertEqual(result.spec, "opentofu/pulumi/elasticstack@0.14.3")

    def test_strips_v_prefix_from_version(self):
        self._write_yaml("random", """
version: "v4.16.0"
publisher: Pulumi
""")

        result = build_package_spec(
            "themes/default/data/registry/packages/random.yaml",
            self.repo_root,
            self.publishers,
        )

        self.assertEqual(result.spec, "pulumi/pulumi/random@4.16.0")

    def test_skips_deprecated_publisher(self):
        self._write_yaml("old", """
version: "1.0.0"
publisher: DEPRECATED
""")

        result = build_package_spec(
            "themes/default/data/registry/packages/old.yaml",
            self.repo_root,
            self.publishers,
        )

        self.assertIsNone(result.spec)
        self.assertIsNone(result.error)
        self.assertTrue(result.skipped)

    def test_errors_on_missing_version(self):
        self._write_yaml("broken", """
publisher: Pulumi
""")

        result = build_package_spec(
            "themes/default/data/registry/packages/broken.yaml",
            self.repo_root,
            self.publishers,
        )

        self.assertIsNone(result.spec)
        self.assertIsNotNone(result.error)
        self.assertIn("no version field", result.error)

    def test_skips_legacy_azure_native(self):
        self._write_yaml("azure-native-v1", """
version: "1.0.0"
publisher: Pulumi
""")

        result = build_package_spec(
            "themes/default/data/registry/packages/azure-native-v1.yaml",
            self.repo_root,
            self.publishers,
        )

        self.assertIsNone(result.spec)
        self.assertTrue(result.skipped)

    def test_skips_legacy_aws(self):
        self._write_yaml("aws-v5", """
version: "5.0.0"
publisher: Pulumi
""")

        result = build_package_spec(
            "themes/default/data/registry/packages/aws-v5.yaml",
            self.repo_root,
            self.publishers,
        )

        self.assertIsNone(result.spec)
        self.assertTrue(result.skipped)

    def test_errors_on_missing_file(self):
        result = build_package_spec(
            "themes/default/data/registry/packages/nonexistent.yaml",
            self.repo_root,
            self.publishers,
        )

        self.assertIsNone(result.spec)
        self.assertIsNotNone(result.error)
        self.assertIn("does not exist", result.error)

    def test_errors_on_invalid_yaml(self):
        self._write_yaml("invalid", "{{{{not valid yaml")

        result = build_package_spec(
            "themes/default/data/registry/packages/invalid.yaml",
            self.repo_root,
            self.publishers,
        )

        self.assertIsNone(result.spec)
        self.assertIsNotNone(result.error)
        self.assertIn("Failed to parse", result.error)


class TestBuildSpecs(unittest.TestCase):
    @patch("publish_to_registry.load_publishers")
    @patch("publish_to_registry.build_package_spec")
    def test_builds_specs_and_collects_errors(self, mock_build, mock_publishers):
        mock_publishers.return_value = {"Pulumi": "pulumi"}
        mock_build.side_effect = [
            SpecResult("pulumi/pulumi/aws@6.50.0"),
            SpecResult("pulumi/pulumi/gcp@7.0.0"),
            SpecResult(None, skipped=True),
            SpecResult(None, error="missing version"),
        ]

        specs, errors = build_specs(
            [
                "themes/default/data/registry/packages/aws.yaml",
                "themes/default/data/registry/packages/gcp.yaml",
                "themes/default/data/registry/packages/deprecated.yaml",
                "themes/default/data/registry/packages/broken.yaml",
            ],
            Path("/repo"),
        )

        self.assertEqual(specs, [
            "pulumi/pulumi/aws@6.50.0",
            "pulumi/pulumi/gcp@7.0.0",
        ])
        self.assertEqual(errors, ["missing version"])


def _outcome(succeeded, failures=None, parsed=False):
    return PublishOutcome(succeeded, failures or [], parsed)


class TestPublishWithRetry(unittest.TestCase):
    @patch("publish_to_registry.ensure_tools_installed")
    @patch("publish_to_registry.publish_specs")
    def test_succeeds_on_first_attempt(self, mock_publish, mock_tools):
        mock_tools.return_value = (Path("/bin/discover"), Path("/bin/publish"))
        mock_publish.return_value = _outcome(True)
        config = Config(repo_root=Path("/repo"))

        result = publish_with_retry(["pulumi/pulumi/aws@6.50.0"], config)

        self.assertTrue(result)
        self.assertEqual(mock_publish.call_count, 1)

    @patch("publish_to_registry.ensure_tools_installed")
    @patch("publish_to_registry.publish_specs")
    @patch("publish_to_registry.time.sleep")
    def test_retries_on_failure(self, mock_sleep, mock_publish, mock_tools):
        mock_tools.return_value = (Path("/bin/discover"), Path("/bin/publish"))
        mock_publish.side_effect = [_outcome(False), _outcome(False), _outcome(True)]
        config = Config(repo_root=Path("/repo"), max_retries=3, initial_backoff=1)

        result = publish_with_retry(["pulumi/pulumi/aws@6.50.0"], config)

        self.assertTrue(result)
        self.assertEqual(mock_publish.call_count, 3)
        self.assertEqual(mock_sleep.call_count, 2)

    @patch("publish_to_registry.ensure_tools_installed")
    @patch("publish_to_registry.publish_specs")
    @patch("publish_to_registry.time.sleep")
    def test_fails_after_max_retries(self, mock_sleep, mock_publish, mock_tools):
        mock_tools.return_value = (Path("/bin/discover"), Path("/bin/publish"))
        mock_publish.return_value = _outcome(False)
        config = Config(repo_root=Path("/repo"), max_retries=3, initial_backoff=1)

        result = publish_with_retry(["pulumi/pulumi/aws@6.50.0"], config)

        self.assertFalse(result)
        self.assertEqual(mock_publish.call_count, 3)

    @patch("publish_to_registry.ensure_tools_installed")
    @patch("publish_to_registry.publish_specs")
    @patch("publish_to_registry.time.sleep")
    def test_exponential_backoff(self, mock_sleep, mock_publish, mock_tools):
        mock_tools.return_value = (Path("/bin/discover"), Path("/bin/publish"))
        mock_publish.side_effect = [_outcome(False), _outcome(False), _outcome(False)]
        config = Config(
            repo_root=Path("/repo"),
            max_retries=3,
            initial_backoff=10,
            max_backoff=30,
        )

        publish_with_retry(["pulumi/pulumi/aws@6.50.0"], config)

        mock_sleep.assert_any_call(10)
        mock_sleep.assert_any_call(20)

    @patch("publish_to_registry.ensure_tools_installed")
    @patch("publish_to_registry.publish_specs")
    @patch("publish_to_registry.time.sleep")
    def test_backoff_capped_at_max(self, mock_sleep, mock_publish, mock_tools):
        mock_tools.return_value = (Path("/bin/discover"), Path("/bin/publish"))
        mock_publish.side_effect = [_outcome(False)] * 4
        config = Config(
            repo_root=Path("/repo"),
            max_retries=4,
            initial_backoff=10,
            max_backoff=30,
        )

        publish_with_retry(["pulumi/pulumi/aws@6.50.0"], config)

        calls = [call[0][0] for call in mock_sleep.call_args_list]
        self.assertEqual(calls, [10, 20, 30])


class TestPublishSpecsIntegration(unittest.TestCase):
    """Integration tests that run the actual discover/publish binaries."""

    @classmethod
    def setUpClass(cls):
        cls.repo_root = Path(__file__).parent.parent.parent
        cls.discover_bin = cls.repo_root / "bin" / "registry-mirror-discover"
        cls.publish_bin = cls.repo_root / "bin" / "registry-mirror-publish"

        if not cls.discover_bin.exists() or not cls.publish_bin.exists():
            raise unittest.SkipTest(
                "registry-mirror-tools binaries not found in bin/. "
                "Run 'make ensure' or install manually to run integration tests."
            )

    def test_dry_run_succeeds_with_valid_spec(self):
        """Verify the discover/publish pipeline works with --dry-run.

        This test catches issues like missing CLI arguments that mocks would miss.
        """
        import os
        from publish_to_registry import publish_specs

        env_backup = os.environ.copy()
        try:
            os.environ["PULUMI_API_URL"] = "https://api.pulumi.com"
            os.environ["PULUMI_ACCESS_TOKEN"] = "test-token"

            result = publish_specs(
                specs=["pulumi/pulumi/random@4.16.0"],
                repo_root=self.repo_root,
                discover_bin=self.discover_bin,
                publish_bin=self.publish_bin,
                dry_run=True,
            )

            self.assertTrue(result, "publish_specs with --dry-run should succeed")
        finally:
            os.environ.clear()
            os.environ.update(env_backup)


if __name__ == "__main__":
    unittest.main()


REJECTION = (
    'publish_failed: complete publish failed: 400: {"code":400,"message":'
    '"Bad Request: Schema version (0.0.0-dev) must match query parameter version (1.3.0)"}'
)


def _output_line(status, reason=""):
    return "2026-08-04T17:13:16Z [output] " + json.dumps({
        "source": "pulumi", "publisher": "pulumi", "package": "terraform-provider",
        "version": "1.3.0", "status": status, "reason": reason,
    })


class TestParsePublishOutput(unittest.TestCase):
    def test_reads_failures_and_their_status_code(self):
        failures, parsed = parse_publish_output(_output_line("failed", REJECTION))
        self.assertTrue(parsed)
        self.assertEqual(len(failures), 1)
        self.assertEqual(failures[0].spec, "pulumi/pulumi/terraform-provider@1.3.0")
        self.assertEqual(failures[0].status_code, 400)
        self.assertTrue(failures[0].permanent)

    def test_published_packages_are_not_failures(self):
        failures, parsed = parse_publish_output(_output_line("published"))
        self.assertTrue(parsed)
        self.assertEqual(failures, [])

    def test_output_without_json_is_not_parsed(self):
        failures, parsed = parse_publish_output("Done. Published: 0, Skipped: 0, Failed: 1")
        self.assertFalse(parsed)
        self.assertEqual(failures, [])

    def test_malformed_json_is_skipped(self):
        failures, parsed = parse_publish_output("[output] {not json")
        self.assertFalse(parsed)
        self.assertEqual(failures, [])


class TestFailureClassification(unittest.TestCase):
    def test_server_errors_are_retried(self):
        outcome = PublishOutcome(False, [PackageFailure("a@1", "503: upstream", 503)], parsed=True)
        self.assertTrue(outcome.worth_retrying)
        self.assertEqual(outcome.permanent_failures, [])

    def test_client_errors_are_not_retried(self):
        outcome = PublishOutcome(False, [PackageFailure("a@1", REJECTION, 400)], parsed=True)
        self.assertFalse(outcome.worth_retrying)
        self.assertEqual(len(outcome.permanent_failures), 1)

    def test_a_single_transient_failure_still_retries(self):
        outcome = PublishOutcome(False, [
            PackageFailure("a@1", REJECTION, 400),
            PackageFailure("b@2", "502: bad gateway", 502),
        ], parsed=True)
        self.assertTrue(outcome.worth_retrying)

    def test_unparsed_output_falls_back_to_retrying(self):
        outcome = PublishOutcome(False, [], parsed=False)
        self.assertTrue(outcome.worth_retrying)

    def test_failure_without_a_status_code_is_retried(self):
        outcome = PublishOutcome(False, [PackageFailure("a@1", "connection reset", None)], parsed=True)
        self.assertTrue(outcome.worth_retrying)


class TestRetrySkipsPermanentFailures(unittest.TestCase):
    @patch("publish_to_registry.report_permanent_failures")
    @patch("publish_to_registry.ensure_tools_installed")
    @patch("publish_to_registry.publish_specs")
    @patch("publish_to_registry.time.sleep")
    def test_permanent_rejection_stops_after_one_attempt(self, mock_sleep, mock_publish,
                                                        mock_tools, mock_report):
        mock_tools.return_value = (Path("/bin/discover"), Path("/bin/publish"))
        failure = PackageFailure("pulumi/pulumi/terraform-provider@1.3.0", REJECTION, 400)
        mock_publish.return_value = PublishOutcome(False, [failure], parsed=True)

        result = publish_with_retry(["pulumi/pulumi/terraform-provider@1.3.0"],
                                    Config(repo_root=Path("/repo"), max_retries=3))

        self.assertFalse(result)
        self.assertEqual(mock_publish.call_count, 1)
        self.assertEqual(mock_sleep.call_count, 0)
        mock_report.assert_called_once_with([failure])

    @patch("publish_to_registry.report_permanent_failures")
    @patch("publish_to_registry.ensure_tools_installed")
    @patch("publish_to_registry.publish_specs")
    @patch("publish_to_registry.time.sleep")
    def test_dry_run_does_not_file_issues(self, mock_sleep, mock_publish, mock_tools, mock_report):
        mock_tools.return_value = (Path("/bin/discover"), Path("/bin/publish"))
        mock_publish.return_value = PublishOutcome(
            False, [PackageFailure("a@1", REJECTION, 400)], parsed=True)

        publish_with_retry(["a@1"], Config(repo_root=Path("/repo"), dry_run=True))

        mock_report.assert_not_called()


class TestReportPermanentFailures(unittest.TestCase):
    def setUp(self):
        self.env = {"GITHUB_TOKEN": "t", "GITHUB_REPOSITORY": "pulumi/registry"}

    @patch("publish_to_registry.github_api")
    def test_opens_an_issue_when_none_is_open(self, mock_api):
        mock_api.side_effect = [{"items": []}, {"number": 7}]
        with patch.dict("os.environ", self.env, clear=True):
            report_permanent_failures([PackageFailure(
                "pulumi/pulumi/terraform-provider@1.3.0", REJECTION, 400)])

        method, path = mock_api.call_args_list[1][0][0], mock_api.call_args_list[1][0][1]
        self.assertEqual((method, path), ("POST", "/repos/pulumi/registry/issues"))
        self.assertIn("terraform-provider", mock_api.call_args_list[1][0][3]["title"])

    @patch("publish_to_registry.github_api")
    def test_comments_instead_of_opening_a_duplicate(self, mock_api):
        mock_api.side_effect = [{"items": [{"number": 42}]}, {}]
        with patch.dict("os.environ", self.env, clear=True):
            report_permanent_failures([PackageFailure(
                "pulumi/pulumi/terraform-provider@1.3.0", REJECTION, 400)])

        method, path = mock_api.call_args_list[1][0][0], mock_api.call_args_list[1][0][1]
        self.assertEqual((method, path), ("POST", "/repos/pulumi/registry/issues/42/comments"))

    @patch("publish_to_registry.github_api")
    def test_without_a_token_nothing_is_filed(self, mock_api):
        with patch.dict("os.environ", {}, clear=True):
            report_permanent_failures([PackageFailure("a@1", REJECTION, 400)])
        mock_api.assert_not_called()
