#!/usr/bin/env bash
# Fail if a PR touches anything other than the package list. Generated content is produced
# after merge, never hand-carried in the PR.
set -euo pipefail

PACKAGE_LIST="community-packages/package-list.json"
base_ref="${1:-origin/master}"
offending="$(git diff --name-only "${base_ref}...HEAD" | grep -v "^${PACKAGE_LIST}\$" || true)"

if [[ -n "$offending" ]]; then
  {
    echo "## ❌ This PR must change only \`${PACKAGE_LIST}\`"
    echo
    echo "The following files do not belong in a community package PR. They are"
    echo "generated and committed automatically after merge:"
    echo
    printf '%s\n' "$offending" | sed 's/^/- `/; s/$/`/'
  } >> "${GITHUB_STEP_SUMMARY:-/dev/stdout}"
  exit 1
fi

echo "Diff is package-list-only."
