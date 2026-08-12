#!/usr/bin/env python3
"""Unit tests for provider_inventory.py"""

import json
import tempfile
import unittest
from pathlib import Path
from unittest.mock import patch

import provider_inventory
from provider_inventory import analyze_extras, analyze_page, load_community_slugs


class TestLoadCommunitySlugs(unittest.TestCase):
    def _load(self, include):
        with tempfile.TemporaryDirectory() as tmp:
            listing = Path(tmp) / "package-list.json"
            listing.write_text(json.dumps({"include": include}))
            with patch.object(provider_inventory, "COMMUNITY_LIST", listing):
                return load_community_slugs()

    def test_prefers_package_name_from_schema_path(self):
        """One repo publishing several packages yields a slug per package."""
        slugs, _ = self._load([
            {"repoSlug": "DefangLabs/pulumi-defang",
             "schemaFile": "provider/cmd/pulumi-resource-defang-aws/schema.json"},
            {"repoSlug": "DefangLabs/pulumi-defang",
             "schemaFile": "provider/cmd/pulumi-resource-defang-gcp/schema.json"},
        ])
        self.assertEqual(slugs, {"defang-aws", "defang-gcp"})

    def test_falls_back_to_repo_name_without_schema_path(self):
        slugs, _ = self._load([
            {"repoSlug": "checkly/pulumi-checkly", "schemaFile": "provider/schema.json"},
        ])
        self.assertEqual(slugs, {"checkly"})

    def test_collects_repo_slugs_lowercased(self):
        """Repo slugs cover packages whose name diverges from the repo name."""
        _, repos = self._load([
            {"repoSlug": "runpod/pulumi-runpod-native", "schemaFile": "provider/schema.json"},
        ])
        self.assertIn("runpod/pulumi-runpod-native", repos)

    def test_missing_list_returns_empty(self):
        with tempfile.TemporaryDirectory() as tmp:
            with patch.object(provider_inventory, "COMMUNITY_LIST", Path(tmp) / "nope.json"):
                self.assertEqual(load_community_slugs(), (set(), set()))


class TestAnalyzePage(unittest.TestCase):
    def _page(self, text):
        with tempfile.TemporaryDirectory() as tmp:
            path = Path(tmp) / "_index.md"
            path.write_text(text)
            return analyze_page(path)

    def test_missing_file_returns_none(self):
        self.assertIsNone(analyze_page(Path("/nonexistent/_index.md")))

    def test_ignores_headings_inside_code_fences(self):
        page = self._page("---\ntitle: X\n---\n"
                          "## Example\n```yaml\n# Not A Heading\n```\n## Installation\n")
        self.assertEqual(page["headings"], ["Example", "Installation"])

    def test_extracts_quoted_and_comma_separated_languages(self):
        page = self._page('---\ntitle: X\n---\n'
                          '{{% choosable language "javascript,typescript" %}}\n'
                          '{{% choosable language go %}}\n'
                          '{{% choosable language "yaml" %}}\n')
        self.assertEqual(page["chooser_langs"], ["go", "javascript", "typescript", "yaml"])

    def test_counts_languages_outside_the_pulumi_sdk_set(self):
        page = self._page("---\ntitle: X\n---\n"
                          "{{% choosable language hcl %}}\n"
                          "{{% choosable language nodejs %}}\n")
        self.assertEqual(page["chooser_langs"], ["hcl", "nodejs"])

    def test_scalar_frontmatter_does_not_raise(self):
        page = self._page("---\njust a string\n---\n## Example\n")
        self.assertEqual(page["layout"], "")

    def test_detects_vendor_fetched_pages(self):
        page = self._page("---\n# WARNING: this file was fetched from https://example\n"
                          "title: X\n---\n## Example\n")
        self.assertTrue(page["vendor_fetched"])


class TestAnalyzeExtras(unittest.TestCase):
    def test_skips_gitignored_api_docs_build_output(self):
        with tempfile.TemporaryDirectory() as tmp:
            pkg = Path(tmp) / "aws"
            (pkg / "api-docs" / "s3").mkdir(parents=True)
            (pkg / "api-docs" / "s3" / "bucket.md").write_text("generated")
            (pkg / "how-to-guides").mkdir()
            (pkg / "how-to-guides" / "guide.md").write_text("guide")
            extras = analyze_extras(pkg, "aws")
        self.assertEqual(extras["migration_dirs"], [])
        self.assertEqual(extras["howto_guides"], 1)
        self.assertTrue(extras["howto_ci_refreshed"])

    def test_marks_hand_committed_guides_as_not_refreshed(self):
        with tempfile.TemporaryDirectory() as tmp:
            pkg = Path(tmp) / "eks"
            (pkg / "how-to-guides").mkdir(parents=True)
            (pkg / "how-to-guides" / "guide.md").write_text("guide")
            extras = analyze_extras(pkg, "eks")
        self.assertFalse(extras["howto_ci_refreshed"])

    def test_missing_directory_is_empty(self):
        extras = analyze_extras(Path("/nonexistent/pkg"), "pkg")
        self.assertEqual(extras["howto_guides"], 0)
        self.assertEqual(extras["other_files"], [])


if __name__ == "__main__":
    unittest.main()
