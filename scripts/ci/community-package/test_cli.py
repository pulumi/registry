from __future__ import annotations

import sys
import unittest
from pathlib import Path

sys.path.insert(0, str(Path(__file__).parent))
import cli  # noqa: E402


class PublishPrTitleAndBranchTests(unittest.TestCase):
    def test_single_entry_keeps_the_per_provider_title_and_branch(self) -> None:
        title, branch = cli.publish_pr_title_and_branch([("thoth", "v0.1.1")])
        self.assertEqual(title, "Publish Package Metadata thoth@v0.1.1")
        self.assertEqual(branch, "thoth/publish-metadata")

    def test_multiple_entries_name_every_provider(self) -> None:
        title, branch = cli.publish_pr_title_and_branch(
            [("defang-aws", "v1.0.0"), ("defang-gcp", "v1.0.0"), ("defang-azure", "v1.0.0")])
        self.assertEqual(title, "Publish package metadata for defang-aws, defang-gcp, defang-azure")
        self.assertEqual(branch, "publish-metadata/defang-aws-defang-gcp-defang-azure")


if __name__ == "__main__":
    unittest.main()
