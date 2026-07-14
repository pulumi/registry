from __future__ import annotations

import sys
import unittest
from pathlib import Path

sys.path.insert(0, str(Path(__file__).parent))
import verify_entry  # noqa: E402


class ProviderNameTests(unittest.TestCase):
    def test_prefers_schema_name_then_repo_suffix(self) -> None:
        self.assertEqual(verify_entry.provider_name("x/pulumi-thoth", {"name": "thoth"}), "thoth")
        self.assertEqual(verify_entry.provider_name("x/pulumi-thoth", {}), "thoth")
        self.assertEqual(verify_entry.provider_name("x/pulumi-provider-dex", {}), "provider-dex")


if __name__ == "__main__":
    unittest.main()
