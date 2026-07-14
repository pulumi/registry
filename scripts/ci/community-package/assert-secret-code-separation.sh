#!/usr/bin/env bash
# Fail CI if any community-package workflow both holds a secret and runs a contributor's
# code. That combination would let a malicious package read the secret, so the two are kept
# in separate files, which forces separate jobs and runners. Runs on any pipeline change.
set -euo pipefail

# Patterns that run contributor code. Only `pipeline.py check` installs SDKs; `docs` and
# `publish` do not, so only `check` is listed.
EXEC='pipeline\.py check|npm (install|ci)|pip (install|download)|go get|go mod download|pulumi plugin install'
# Patterns that pull a secret into a job.
SECRET='secrets\.|ESC_ACTION|esc-action|ANTHROPIC_API_KEY|PULUMI_BOT_TOKEN|id-token: *write'

fail=0
shopt -s nullglob
for wf in .github/workflows/community-package-*.yml; do
  has_exec=0;   grep -Eq "$EXEC"   "$wf" && has_exec=1
  has_secret=0; grep -Eq "$SECRET" "$wf" && has_secret=1

  if (( has_exec && has_secret )); then
    echo "❌ ${wf}: holds a secret and runs contributor code; split them into separate files."
    grep -nE "$EXEC|$SECRET" "$wf" | sed 's/^/      /'
    fail=1
  fi
  if grep -q 'pull_request_target' "$wf"; then
    echo "❌ ${wf}: uses pull_request_target, forbidden in the community-package pipeline."
    fail=1
  fi
  # A secret-bearing file must not check out the PR's head ref.
  if (( has_secret )) && grep -Eq 'ref:.*\.head\.' "$wf"; then
    echo "❌ ${wf}: a secret-bearing file checks out the PR's head ref."
    fail=1
  fi
done

(( fail == 0 )) && echo "✅ no community-package workflow mixes secrets with contributor code."
exit $fail
