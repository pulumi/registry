#!/usr/bin/env python3
"""Unit tests for publish_to_registry.py"""

import json
import tempfile
import unittest
from pathlib import Path
from unittest.mock import MagicMock, patch

from publish_to_registry import (
    Config,
    SpecResult,
    build_package_spec,
    build_specs,
    get_changed_packages,
    publish_with_retry,
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


class TestPublishWithRetry(unittest.TestCase):
    @patch("publish_to_registry.ensure_tools_installed")
    @patch("publish_to_registry.publish_specs")
    def test_succeeds_on_first_attempt(self, mock_publish, mock_tools):
        mock_tools.return_value = (Path("/bin/discover"), Path("/bin/publish"))
        mock_publish.return_value = True
        config = Config(repo_root=Path("/repo"))

        result = publish_with_retry(["pulumi/pulumi/aws@6.50.0"], config)

        self.assertTrue(result)
        self.assertEqual(mock_publish.call_count, 1)

    @patch("publish_to_registry.ensure_tools_installed")
    @patch("publish_to_registry.publish_specs")
    @patch("publish_to_registry.time.sleep")
    def test_retries_on_failure(self, mock_sleep, mock_publish, mock_tools):
        mock_tools.return_value = (Path("/bin/discover"), Path("/bin/publish"))
        mock_publish.side_effect = [False, False, True]
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
        mock_publish.return_value = False
        config = Config(repo_root=Path("/repo"), max_retries=3, initial_backoff=1)

        result = publish_with_retry(["pulumi/pulumi/aws@6.50.0"], config)

        self.assertFalse(result)
        self.assertEqual(mock_publish.call_count, 3)

    @patch("publish_to_registry.ensure_tools_installed")
    @patch("publish_to_registry.publish_specs")
    @patch("publish_to_registry.time.sleep")
    def test_exponential_backoff(self, mock_sleep, mock_publish, mock_tools):
        mock_tools.return_value = (Path("/bin/discover"), Path("/bin/publish"))
        mock_publish.side_effect = [False, False, False]
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
        mock_publish.side_effect = [False, False, False, False]
        config = Config(
            repo_root=Path("/repo"),
            max_retries=4,
            initial_backoff=10,
            max_backoff=30,
        )

        publish_with_retry(["pulumi/pulumi/aws@6.50.0"], config)

        calls = [call[0][0] for call in mock_sleep.call_args_list]
        self.assertEqual(calls, [10, 20, 30])


if __name__ == "__main__":
    unittest.main()
