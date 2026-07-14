#!/usr/bin/env bash
# Handles a `/check` PR comment by re-running the check workflow against current upstream.
# PR author or maintainer, with a per-PR cooldown. Never runs the PR's code; it only reads
# metadata and triggers the rerun.
set -euo pipefail
source "$(dirname "$0")/cmd-lib.sh"
: "${REPO:?}" "${PR:?}" "${ASSOC:?}" "${COMMENT_ID:?}" "${COMMENTER:?}"
COOLDOWN_MINUTES="${COOLDOWN_MINUTES:-10}"

load_pr
authorize author-or-maintainer || exit 0
cooldown_guard check || exit 0

run_id="$(latest_pr_run_id community-package-check.yml)"
if [[ -z "$run_id" ]]; then
  react "confused"
  say "No check run found for \`${SHA:0:12}\`. Push any commit to trigger one."
  exit 0
fi

gh run rerun "$run_id" -R "$REPO"
react "+1"
url="$(factsheet_url)"
if [[ -n "$url" ]]; then
  say "🔁 Re-checking \`${SHA:0:12}\` against current upstream. The [fact-sheet](${url}) updates in place when it finishes."
else
  say "🔁 Checking \`${SHA:0:12}\`. A fact-sheet comment appears when it finishes."
fi
