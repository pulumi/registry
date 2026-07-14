from __future__ import annotations

import sys
import unittest
from pathlib import Path

sys.path.insert(0, str(Path(__file__).parent))
import doc_lint  # noqa: E402


class FindIssuesTests(unittest.TestCase):
    def test_flags_relative_refs_and_ignores_absolute(self) -> None:
        doc = (
            "![logo](./logo.png)\n"
            "<img src='/x.svg'>\n"
            "<a href=\"./page\">y</a>\n"
            "[ok](https://example.com)\n"
            "![ok](https://example.com/y.png)\n"
        )
        issues = doc_lint.find_issues(doc)
        self.assertEqual({i.line for i in issues}, {1, 2, 3})
        self.assertIn("relative-image", {i.kind for i in issues})
        self.assertIn("raw-relative-link", {i.kind for i in issues})

    def test_clean_doc(self) -> None:
        self.assertEqual(doc_lint.find_issues("all good\n[x](https://y)\n![a](https://y/z.png)\n"), [])

    def test_flags_md_suffixed_cross_link(self) -> None:
        doc = (
            "see [config](./installation-configuration.md).\n"
            "ok [config](./installation-configuration)\n"
            "abs [readme](https://github.com/a/b/blob/main/README.md)\n"
        )
        issues = doc_lint.find_issues(doc)
        self.assertEqual({i.line for i in issues}, {1})
        self.assertEqual(issues[0].kind, "md-suffixed-link")


if __name__ == "__main__":
    unittest.main()
