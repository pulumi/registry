from __future__ import annotations

import sys
import unittest
from pathlib import Path

sys.path.insert(0, str(Path(__file__).parent))
import fact_sheet  # noqa: E402
from models import DocFile, DocFinding, InstallResult, Manifest, Version  # noqa: E402


def _manifest(findings: list[DocFinding], installs: list[InstallResult]) -> Manifest:
    return Manifest(
        "x/pulumi-demo", "s.json", "demo", Version("v1.0.0", "0" * 40),
        "x", installs, findings, green=True,
    )


class RenderTests(unittest.TestCase):
    def test_renders_install_failure_details(self) -> None:
        out = fact_sheet.render(_manifest([], [
            InstallResult("python", "pip download x==1", "fail", error="No matching distribution found for x==1"),
        ]))
        self.assertIn("pip download x==1", out)
        self.assertIn("No matching distribution found", out)
        self.assertIn("<details>", out)

    def test_renders_lint_findings_with_message(self) -> None:
        out = fact_sheet.render(_manifest([DocFinding(12, "relative-image", "![a](./b.png)")], []))
        self.assertIn("Doc-lint: 1 finding(s)", out)
        self.assertIn("docs/_index.md:12", out)
        self.assertIn("![a](./b.png)", out)

    def test_clean_lint_line_when_no_findings(self) -> None:
        out = fact_sheet.render(_manifest([], [InstallResult("plugin", "pulumi …", "pass", blocking=True)]))
        self.assertIn("No relative images", out)

    def test_red_gate_shows_not_ready_verdict(self) -> None:
        manifest = _manifest([], [InstallResult("plugin", "pulumi …", "fail", blocking=True)])
        manifest.green = False
        out = fact_sheet.render(manifest)
        self.assertIn("❌", out.splitlines()[0])
        self.assertIn("Not ready", out)
        self.assertIn("/check", out)

    def test_embeds_doc_source_inline(self) -> None:
        manifest = _manifest([], [])
        manifest.docs = [DocFile("docs/_index.md", 3, "# Title\n```go\ncode\n```\n")]
        out = fact_sheet.render(manifest)
        self.assertIn("Provider docs", out)
        self.assertIn("open full file", out)
        self.assertIn("/blob/" + "0" * 40 + "/docs/_index.md", out)
        self.assertIn("# Title", out)
        self.assertIn("<details>", out)

    def test_fence_outgrows_embedded_backticks(self) -> None:
        self.assertEqual(fact_sheet._fence_longer_than_any_run_in("no ticks"), "```")
        self.assertEqual(fact_sheet._fence_longer_than_any_run_in("```go\nx\n```"), "````")


if __name__ == "__main__":
    unittest.main()
