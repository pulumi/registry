#!/bin/bash

# Unit tests for the URL mapping behind the pinned preview comment posted by
# scripts/ci/sync.sh. These cover the parts that decide *which* links land in the comment --
# path-to-URL translation, de-duplication, and the existence gate -- without touching the
# network, S3, or a Hugo build.
#
# Run directly (./scripts/ci/test-preview-comment.sh) or via `make test-preview-comment`.

set -o errexit -o pipefail

cd "$(dirname "$0")/../.."

source ./scripts/ci/common.sh

failures=0

# Usage: expect <description> <expected> <actual>
expect() {
    local description=$1
    local expected=$2
    local actual=$3

    if [[ "$expected" == "$actual" ]]; then
        echo "ok   - ${description}"
    else
        echo "FAIL - ${description}"
        echo "         expected: ${expected}"
        echo "         actual:   ${actual}"
        failures=$(( failures + 1 ))
    fi
}

echo "== content_path_to_url =="

expect "package landing page" \
    "/registry/packages/aws/" \
    "$(content_path_to_url themes/default/content/registry/packages/aws/_index.md)"

expect "installation and configuration page" \
    "/registry/packages/aws/installation-configuration/" \
    "$(content_path_to_url themes/default/content/registry/packages/aws/installation-configuration.md)"

expect "how-to guide" \
    "/registry/packages/aws/how-to-guides/7-0-migration/" \
    "$(content_path_to_url themes/default/content/registry/packages/aws/how-to-guides/7-0-migration.md)"

expect "registry section landing page" \
    "/registry/" \
    "$(content_path_to_url themes/default/content/registry/_index.md)"

expect "site root" \
    "/" \
    "$(content_path_to_url themes/default/content/_index.md)"

expect "leaf bundle collapses to its directory" \
    "/registry/packages/aws/how-to-guides/some-guide/" \
    "$(content_path_to_url themes/default/content/registry/packages/aws/how-to-guides/some-guide/index.md)"

echo
echo "== changed_paths_to_urls =="

expect "package YAML maps to the package page" \
    "/registry/packages/aws/" \
    "$(echo themes/default/data/registry/packages/aws.yaml | changed_paths_to_urls)"

expect "a package YAML and its landing page de-duplicate" \
    "/registry/packages/aws/" \
    "$(printf '%s\n' \
        themes/default/data/registry/packages/aws.yaml \
        themes/default/content/registry/packages/aws/_index.md \
        | changed_paths_to_urls)"

expect "non-page sources are dropped" \
    "" \
    "$(printf '%s\n' \
        themes/default/layouts/partials/registry/package/icon.html \
        community-packages/package-list.json \
        scripts/ci/sync.sh \
        themes/default/data/registry/external_logo_treatment.yaml \
        | changed_paths_to_urls)"

expect "first-seen order is preserved across both rules" \
    "$(printf '%s\n' /registry/packages/gcp/ /registry/packages/aws/ /registry/packages/aws/installation-configuration/)" \
    "$(printf '%s\n' \
        themes/default/data/registry/packages/gcp.yaml \
        themes/default/content/registry/packages/aws/_index.md \
        themes/default/data/registry/packages/aws.yaml \
        themes/default/content/registry/packages/aws/installation-configuration.md \
        | changed_paths_to_urls)"

expect "blank lines are ignored" \
    "/registry/packages/aws/" \
    "$(printf '%s\n' "" themes/default/data/registry/packages/aws.yaml "" | changed_paths_to_urls)"

echo
echo "== existence gate =="

# sync.sh only links a URL when the build actually rendered it, so that url:/alias overrides
# and files that don't produce a page are dropped rather than linked as dead URLs. Stand up a
# fake build tree and assert that filter directly.
build_dir=$(mktemp -d)
trap 'rm -rf "$build_dir"' EXIT
mkdir -p "${build_dir}/registry/packages/aws"
touch "${build_dir}/registry/packages/aws/index.html"

rendered=()
while IFS= read -r url; do
    [[ -f "${build_dir}${url}index.html" ]] && rendered+=("$url")
done < <(printf '%s\n' \
    themes/default/data/registry/packages/aws.yaml \
    themes/default/data/registry/packages/never-rendered.yaml \
    | changed_paths_to_urls)

expect "only rendered pages survive" \
    "/registry/packages/aws/" \
    "$(printf '%s\n' "${rendered[@]}")"

echo
echo "== changed_pages_section =="

# The files API is paged at 100. changed_paths_to_urls only de-duplicates within a single
# invocation, so the section has to collect every page before mapping -- otherwise a package's
# YAML on page one and its landing page on page two each produce a link to the same URL.
# Stub two full pages that straddle exactly that split.
page_one='[{"status": "modified", "filename": "themes/default/data/registry/packages/aws.yaml"}'
for i in $(seq 1 99); do
    page_one+=",{\"status\": \"modified\", \"filename\": \"docs/filler-${i}.md\"}"
done
page_one+=']'
page_two='[{"status": "modified", "filename": "themes/default/content/registry/packages/aws/_index.md"}]'

curl() {
    local url=""
    while [[ $# -gt 0 ]]; do
        case "$1" in
            -H|--max-time) shift 2;;
            -s) shift;;
            *) url="$1"; shift;;
        esac
    done
    if [[ "$url" == *"page=1"* ]]; then echo "$page_one"; else echo "$page_two"; fi
}

expect "a page-straddling package lists once, not twice" \
    "$(printf '\n\n**Changed pages:**\n- [/registry/packages/aws/](http://preview.example/registry/packages/aws/)')" \
    "$(changed_pages_section "https://api.github.com/repos/pulumi/registry" 42 "$build_dir" "http://preview.example")"

curl() { echo '{"message": "Bad credentials", "status": "401"}'; }

expect "an API error payload yields no section rather than a broken one" \
    "" \
    "$(changed_pages_section "https://api.github.com/repos/pulumi/registry" 42 "$build_dir" "http://preview.example")"

unset -f curl

echo
echo "== upsert_github_pr_comment =="

# upsert_github_pr_comment talks to nothing but curl, so shadowing curl drives the whole
# find-or-create decision offline. $stub_comments is the canned comments-list response;
# $stub_calls records what the function asked for.
comments_url="https://api.github.com/repos/pulumi/registry/issues/42/comments"
repo_url="https://api.github.com/repos/pulumi/registry"
marker="<!-- registry-preview-link -->"
export GITHUB_TOKEN=stub-token

curl() {
    local url="" method="GET" write_out=""
    while [[ $# -gt 0 ]]; do
        case "$1" in
            -X) method="$2"; shift 2;;
            -w) write_out="$2"; shift 2;;
            -d|-H|-o|--max-time) shift 2;;
            -s) shift;;
            *) url="$1"; shift;;
        esac
    done

    if [[ "$method" != "GET" ]]; then
        echo "${method} ${url}" >> "$stub_calls"
        # Mimic -w '%{http_code}' so the status check in upsert_github_pr_comment sees a 2xx.
        [[ -n "$write_out" ]] && echo "201"
        return 0
    fi

    # Only page one carries anything; later pages come back empty so the loop terminates.
    if [[ "$url" == *"page=1"* ]]; then
        echo "$stub_comments"
    else
        echo '[]'
    fi
}

# Usage: run_upsert <comments-list-json>
run_upsert() {
    stub_comments=$1
    stub_calls=$(mktemp)
    upsert_github_pr_comment "$marker" "a body" "$comments_url" "$repo_url" || true
    cat "$stub_calls"
    rm -f "$stub_calls"
}

expect "no existing comment creates one" \
    "POST ${comments_url}" \
    "$(run_upsert '[]')"

expect "an existing pinned comment is updated in place" \
    "PATCH ${repo_url}/issues/comments/222" \
    "$(run_upsert '[
        {"id": 111, "user": {"login": "github-actions[bot]"}, "body": "unrelated"},
        {"id": 222, "user": {"login": "github-actions[bot]"}, "body": "<!-- registry-preview-link -->\nold"}
     ]')"

# The marker alone must not be enough to redirect the pinned comment, or any contributor could
# take it over by quoting it.
expect "a marker from another author is ignored" \
    "POST ${comments_url}" \
    "$(run_upsert '[
        {"id": 333, "user": {"login": "some-contributor"}, "body": "<!-- registry-preview-link --> hijack"}
     ]')"

expect "the bot comment wins even when a decoy sorts after it" \
    "PATCH ${repo_url}/issues/comments/444" \
    "$(run_upsert '[
        {"id": 444, "user": {"login": "pulumi-bot"}, "body": "<!-- registry-preview-link -->\nold"},
        {"id": 555, "user": {"login": "some-contributor"}, "body": "<!-- registry-preview-link --> hijack"}
     ]')"

# jq's .[] iterates an object's values as readily as an array's elements, so an error payload
# would blow up the marker filter unless the type is checked first. Falling back to a fresh
# comment is the right failure mode; erroring out and posting nothing is not.
expect "an API error payload falls back to creating a comment" \
    "POST ${comments_url}" \
    "$(run_upsert '{"message": "Bad credentials", "status": "401"}')"

unset -f curl

echo "== preview identity =="

expect "an explicit preview PR names the bucket after it" \
    "pr-12144-abcdef12" \
    "$(PREVIEW_PR=12144 PREVIEW_HEAD_SHA=abcdef1234567890 build_identifier)"

expect "an explicit head sha wins over the checked-out commit" \
    "abcdef1234567890" \
    "$(PREVIEW_HEAD_SHA=abcdef1234567890 git_sha)"

pr_event_file="$(mktemp)"
echo '{"pull_request": {"head": {"sha": "0123456789abcdef"}}}' > "$pr_event_file"

expect "without one, a pull_request event still decides the sha" \
    "0123456789abcdef" \
    "$(GITHUB_EVENT_NAME=pull_request GITHUB_EVENT_PATH="$pr_event_file" git_sha)"

rm -f "$pr_event_file"

echo
if [[ "$failures" -gt 0 ]]; then
    echo "${failures} test(s) failed."
    exit 1
fi

echo "All tests passed."
