#!/usr/bin/env bash
# Prepares a `/review`: authorizes the commenter, rate-limits, and pre-fetches the provider
# docs and schema into ./review-input/ for the AI step to read. Never checks out or runs the
# PR's code. Sets the `proceed` output.
set -euo pipefail
source "$(dirname "$0")/cmd-lib.sh"
: "${REPO:?}" "${PR:?}" "${ASSOC:?}" "${COMMENT_ID:?}" "${COMMENTER:?}"
COOLDOWN_MINUTES="${COOLDOWN_MINUTES:-15}"

deny() { echo "proceed=false" >> "${GITHUB_OUTPUT:-/dev/stdout}"; exit 0; }

load_pr
authorize maintainer-only || deny
cooldown_guard review || deny

cp community-packages/package-list.json /tmp/base.json 2>/dev/null || echo '{"include":[]}' > /tmp/base.json
gh api "/repos/${REPO}/contents/community-packages/package-list.json?ref=${SHA}" -q '.content' | base64 -d > /tmp/head.json

python3 scripts/ci/community-package/pipeline.py docs --base /tmp/base.json --head /tmp/head.json --out review-input

react "+1"
say "🔎 Running the AI documentation review. Findings post as a comment when it finishes."
echo "proceed=true" >> "${GITHUB_OUTPUT:-/dev/stdout}"
