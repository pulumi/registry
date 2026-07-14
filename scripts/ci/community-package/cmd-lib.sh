#!/usr/bin/env bash
# Shared helpers for the /check and /review comment commands. Requires REPO, PR, COMMENT_ID,
# COMMENTER, ASSOC, COOLDOWN_MINUTES in the environment.
set -euo pipefail

react() { gh api -X POST "/repos/${REPO}/issues/comments/${COMMENT_ID}/reactions" -f "content=$1" >/dev/null 2>&1 || true; }
say()   { gh pr comment "$PR" -R "$REPO" --body "$1" >/dev/null; }

load_pr() { read -r PR_AUTHOR SHA < <(gh pr view "$PR" -R "$REPO" --json author,headRefOid -q '.author.login + " " + .headRefOid'); }

authorize() {
  if [[ "$1" == "author-or-maintainer" && "$COMMENTER" == "$PR_AUTHOR" ]]; then return 0; fi
  case "$ASSOC" in OWNER|MEMBER|COLLABORATOR) return 0 ;; esac
  react "-1"; echo "not authorized: ${COMMENTER} (assoc ${ASSOC})"; return 1
}

cooldown_guard() {
  local re="$1" last elapsed
  last="$(gh api "/repos/${REPO}/commits/${SHA}/check-runs" \
    -q "[.check_runs[] | select(.name|test(\"${re}\"))] | sort_by(.started_at) | last | .started_at // empty" 2>/dev/null || true)"
  [[ -n "$last" ]] || return 0
  elapsed=$(( ( $(date -u +%s) - $(date -u -d "$last" +%s) ) / 60 ))
  if (( elapsed < COOLDOWN_MINUTES )); then
    react "eyes"
    say "⏳ Rate limited. Last ran ${elapsed} min ago; try again in $(( COOLDOWN_MINUTES - elapsed )) min."
    return 1
  fi
  return 0
}

latest_pr_run_id() { gh api "/repos/${REPO}/actions/workflows/$1/runs?event=pull_request&head_sha=${SHA}" -q '.workflow_runs[0].id // empty'; }

FACTSHEET_MARKER="<!-- community-package-fact-sheet -->"
factsheet_url() {
  gh api "/repos/${REPO}/issues/${PR}/comments" \
    -q "[.[] | select(.body | contains(\"${FACTSHEET_MARKER}\"))][0].html_url // empty" 2>/dev/null || true
}
