#!/usr/bin/env bash
# Posts the check workflow's fact-sheet as a single sticky PR comment. Has a write token but
# only posts markdown the check already rendered into its artifact, and runs none of the PR's
# code, so the write token is safe.
set -euo pipefail
: "${REPO:?}" "${PR_HEAD:?}"

MARKER="<!-- community-package-fact-sheet -->"

pr="$(gh pr list -R "$REPO" --head "$PR_HEAD" --json number -q '.[0].number' 2>/dev/null || true)"
[[ -n "$pr" ]] || { echo "no open PR for ${PR_HEAD}"; exit 0; }

body="$MARKER"$'\n'
shopt -s nullglob
for f in *.factsheet.md; do body+=$'\n'"$(cat "$f")"$'\n'; done

cid="$(gh api "/repos/${REPO}/issues/${pr}/comments" \
  -q "[.[] | select(.body | contains(\"${MARKER}\"))][0].id // empty" 2>/dev/null || true)"

if [[ -n "$cid" ]]; then
  gh api -X PATCH "/repos/${REPO}/issues/comments/${cid}" -f body="$body" >/dev/null
else
  gh api -X POST "/repos/${REPO}/issues/${pr}/comments" -f body="$body" >/dev/null
fi
echo "posted fact-sheet to PR #${pr}"
