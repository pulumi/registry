from __future__ import annotations

import sys
import unittest
from pathlib import Path

sys.path.insert(0, str(Path(__file__).parent))
import package_list  # noqa: E402


class AddedEntriesTests(unittest.TestCase):
    def test_detects_added(self) -> None:
        before = '{"include":[{"repoSlug":"a/b","schemaFile":"s"}]}'
        after = '{"include":[{"repoSlug":"a/b","schemaFile":"s"},{"repoSlug":"c/d","schemaFile":"t"}]}'
        self.assertEqual([e.repoSlug for e in package_list.added_entries(before, after)], ["c/d"])

    def test_no_change(self) -> None:
        same = '{"include":[{"repoSlug":"a/b","schemaFile":"s"}]}'
        self.assertEqual(package_list.added_entries(same, same), [])

    def test_accepts_packages_key_or_bare_array(self) -> None:
        self.assertEqual(
            [e.repoSlug for e in package_list.added_entries('{"packages":[]}',
                                                            '{"packages":[{"repoSlug":"a/b","schemaFile":"s"}]}')],
            ["a/b"])
        self.assertEqual(
            [e.repoSlug for e in package_list.added_entries("[]", '[{"repoSlug":"a/b","schemaFile":"s"}]')],
            ["a/b"])


if __name__ == "__main__":
    unittest.main()
