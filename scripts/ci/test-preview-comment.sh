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
if [[ "$failures" -gt 0 ]]; then
    echo "${failures} test(s) failed."
    exit 1
fi

echo "All tests passed."
