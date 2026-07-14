from __future__ import annotations

import re
import unittest
from pathlib import Path

_WORKFLOWS = Path(__file__).resolve().parents[3] / ".github" / "workflows"

# The check subcommand is what install-probes the advertised SDKs; the rest of the patterns
# are the install commands themselves. Any of them means a job runs contributor-controlled code.
RUNS_CONTRIBUTOR_CODE = re.compile(
    r"cli\.py check(?![-\w])|npm (install|ci)|pip (install|download)|go get|go mod download|pulumi plugin install")
HOLDS_A_SECRET = re.compile(
    r"secrets\.|ESC_ACTION|esc-action|ANTHROPIC_API_KEY|PULUMI_BOT_TOKEN|id-token: *write")


class SecretCodeSeparationTests(unittest.TestCase):
    """No community-package workflow may both hold a secret and run a contributor's code;
    if one did, a malicious package could read the secret."""

    def _workflows(self) -> list[Path]:
        found = sorted(_WORKFLOWS.glob("community-package-*.yml"))
        self.assertTrue(found, "no community-package workflows found")
        return found

    def test_no_secret_shares_a_job_with_contributor_code(self) -> None:
        for workflow in self._workflows():
            text = workflow.read_text()
            mixes = bool(RUNS_CONTRIBUTOR_CODE.search(text)) and bool(HOLDS_A_SECRET.search(text))
            self.assertFalse(mixes, f"{workflow.name} mixes secrets with contributor-code execution")

    def test_no_pull_request_target(self) -> None:
        for workflow in self._workflows():
            self.assertNotIn("pull_request_target", workflow.read_text(), f"{workflow.name} uses pull_request_target")

    def test_secret_workflow_never_checks_out_the_pr_head(self) -> None:
        for workflow in self._workflows():
            text = workflow.read_text()
            if HOLDS_A_SECRET.search(text):
                self.assertNotRegex(text, r"ref:.*\.head\.",
                                    f"{workflow.name} checks out the PR head with a secret present")


if __name__ == "__main__":
    unittest.main()
