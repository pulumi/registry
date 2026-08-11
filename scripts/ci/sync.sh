#!/bin/bash

set -o errexit -o pipefail

# This script takes the built Hugo site and:
#   - creates a new S3 bucket named according to whether the action is a push or pull_request.
#   - creates a list of all Hugo-generated client-side ("meta-refresh") redirects that
#   - we'll use to produce proper 301s later.
#   - pushes the content of the website into the new S3 bucket.
#   - tests the built website, first for broken links, then with Cypress to ensure pages
#   - render and behave properly.
#   - emits a metadata file containing information about the commit and bucket, which
#   - Pulumi will use to process its update.
#   - writes a record to AWS Parameter Store relating the generated bucket to the commit
#   - responsible for producing it.
#   - Posts a PR comment back to GitHub, if applicable.

source ./scripts/ci/common.sh

# The docroot of the built website.
build_dir="public"

# The text file we'll write as an output result.
metadata_file="$(origin_bucket_metadata_filepath)"

# Verify we have at least 1000 index.html files in total across the site.
if [ ! "$(find $build_dir -type f | grep index.html | wc -l)" -ge 1000 ]; then
    echo "Page-count check failed. Exiting."
    exit 1
fi

# Move the sitemap from root to the registry directory.
mv "${build_dir}/sitemap.xml" "${build_dir}/registry/sitemap.xml"

# For previews, name the destination bucket with the PR number, to reduce the number of
# buckets we create and to facilitate shorter sync times.
destination_bucket=$(echo "$(origin_bucket_prefix)-$(build_identifier)" | tr '_' '-')
destination_bucket_uri="s3://${destination_bucket}"

# Make the bucket. If this fails, there are two explanations, given the way we're naming
# our buckets: either a previous run failed at some point after creating the bucket, in
# which case we should simply proceed (to repopulate it), or the bucket was somehow
# created in another account, in which case subsequent operations on the bucket will also
# fail, causing this script to exit nonzero. In either case, it's okay to continue.
aws s3 mb $destination_bucket_uri --region "$(aws_region)" || true
aws s3api put-public-access-block --bucket $destination_bucket --public-access-block-configuration BlockPublicAcls=false
aws s3api put-bucket-ownership-controls --bucket $destination_bucket --ownership-controls="Rules=[{ObjectOwnership=ObjectWriter}]"
aws s3api put-bucket-acl --bucket $destination_bucket --acl bucket-owner-full-control --acl public-read

aws s3api put-bucket-tagging --bucket $destination_bucket --tagging "TagSet=[{$(aws_owner_tag)}]" --region "$(aws_region)"

# Make the bucket an S3 website.
aws s3 website $destination_bucket_uri --index-document index.html --error-document 404.html --region "$(aws_region)"

# Sync the local build directory to the bucket using s5cmd for massively parallel uploads.
# s5cmd uses hundreds of concurrent goroutines vs aws cli's ~10-16 concurrent requests,
# resulting in 10-50x faster uploads for large file counts.
# The --delete flag removes destination objects not present locally, keeping the bucket clean
# for PR preview buckets that get reused across commits.
log "Synchronizing to $destination_bucket_uri..."
s5cmd --log error sync --delete --acl public-read \
    "$build_dir/" "$destination_bucket_uri/"
log "Sync complete."

# Sync LLM docs separately. These are generated outside the Hugo build tree to avoid
# Hugo processing static files. They're uploaded directly to the same URL paths
# they'd occupy if they were in the Hugo static directory.
#
# LLM docs bundles are uploaded gzip-compressed with Content-Encoding: gzip so they
# pass through both CloudFront layers (registry CDN + www.pulumi.com CDN) without
# relying on CloudFront's automatic compression -- which is disabled on /registry/*
# because the outer CDN forwards Accept-Encoding via AllViewerExceptHostHeader, and
# wouldn't help anyway for bundles over CloudFront's 10 MB auto-compress ceiling
# (the aws bundle is >100 MB uncompressed). Go's http.Client decompresses
# Content-Encoding: gzip transparently, so the pulumi docs CLI requires no changes.
#
# The llm-docs-out/ tree contains only llm-docs.json files, so it is safe to apply
# --content-encoding/--content-type globally to every object in the sync.
llm_docs_dir="llm-docs-out"
if [[ -d "$llm_docs_dir" ]]; then
    log "Pre-gzipping LLM docs..."
    find "$llm_docs_dir" -type f -name 'llm-docs.json' -print0 \
        | xargs -0 -P"$(nproc)" -I{} sh -c 'gzip -9 < "$1" > "$1.gz.tmp" && mv "$1.gz.tmp" "$1"' _ {}

    log "Synchronizing LLM docs to $destination_bucket_uri..."
    s5cmd --log error sync --acl public-read \
        --content-encoding gzip \
        --content-type application/json \
        "$llm_docs_dir/" "$destination_bucket_uri/"
    log "LLM docs sync complete."
fi

s3_website_url="http://${destination_bucket}.s3-website.$(aws_region).amazonaws.com"
echo "$s3_website_url"

# Set the content-type of latest-version explicitly. (Otherwise, it'll be set as binary/octet-stream.)
aws s3 cp "$build_dir/latest-version" "${destination_bucket_uri}/latest-version" \
    --content-type "text/plain" --acl public-read --region "$(aws_region)" --metadata-directive REPLACE

# Smoke test LLM docs compression at the S3 website layer (before either CloudFront
# layer has a chance to mask a regression). Uses the random package because it has a
# small, reliably-present llm-docs.json bundle. Catches the regression where we
# accidentally upload llm-docs.json files without Content-Encoding: gzip, which would
# otherwise only surface after deploy when the CLI starts paying full wire cost again.
#
# Uses curl --dump-header to validate the body (gunzip -t) and capture the headers in
# a single round trip, rather than issuing separate HEAD and GET requests.
if [[ -d "$llm_docs_dir" && -f "$llm_docs_dir/registry/packages/random/api-docs/llm-docs.json" ]]; then
    log "Smoke-testing LLM docs compression..."
    llm_docs_test_url="${s3_website_url}/registry/packages/random/api-docs/llm-docs.json"
    llm_docs_headers_file=$(mktemp)
    if ! curl -fsS --dump-header "$llm_docs_headers_file" "$llm_docs_test_url" | gunzip -t; then
        echo "ERROR: $llm_docs_test_url body did not decompress as gzip (or request failed)." >&2
        rm -f "$llm_docs_headers_file"
        exit 1
    fi
    llm_docs_headers=$(cat "$llm_docs_headers_file")
    rm -f "$llm_docs_headers_file"
    if ! grep -qi '^content-encoding: gzip' <<< "$llm_docs_headers"; then
        echo "ERROR: $llm_docs_test_url is missing 'Content-Encoding: gzip'. Response headers:" >&2
        echo "$llm_docs_headers" >&2
        exit 1
    fi
    log "LLM docs compression smoke test passed."
else
    log "Skipping LLM docs compression smoke test: $llm_docs_dir/registry/packages/random/api-docs/llm-docs.json not found."
fi

# Smoke test the deployed website.
log "Running browser tests on $s3_website_url..."
./scripts/run-browser-tests.sh "$s3_website_url"

# At this point, we have a bucket that's suitable for deployment. As a result of this run,
# we leave a file in the project root indicating the name of the bucket that was generated
# and the associated commit SHA, and then we upload that file into the bucket as well, for
# reference. The Pulumi program will expect this file to exist, and will use the bucket
# name to set the CloudFront origin on the next Pulumi run.
#
# Why use a local file and not `pulumi config`, or some other persistence store? Because
# we need ensure that every CI job deploys only what it was responsible for building.
# Coupled with the locking we get from the Pulumi Cloud, using a local file is a safe
# way to ensure we're deploying what we just finished building and testing.
log "Writing result metadata."
metadata='{
    "timestamp": %s,
    "commit": "%s",
    "bucket": "%s",
    "url": "%s"
}'
printf "$metadata" "$(current_time_in_ms)" "$(git_sha)" "$destination_bucket" "$s3_website_url" > "$metadata_file"

# Copy the file to the destination bucket, for future reference.
aws s3 cp "$metadata_file" "${destination_bucket_uri}/registry/metadata.json" --region "$(aws_region)" --acl public-read

# Persist an association between the current commit and the bucket we just deployed to.
set_bucket_for_commit "$(git_sha)" "$destination_bucket" "$(aws_region)"

# The marker that makes the preview comment pinned: every build finds the comment carrying it
# and rewrites that one, instead of leaving a fresh comment per commit.
PREVIEW_COMMENT_MARKER="<!-- registry-preview-link -->"

# Builds the "Changed pages" section: direct links to the pages this PR changed, so a reviewer
# lands on them instead of hunting through the site.
#
# The base branch isn't reliably available locally (the /preview command path checks out
# differently from the pull_request path), so we ask the GitHub API which files changed rather
# than diffing. Each changed path is mapped to the URL it renders, then kept only if it
# actually rendered to a page in this build -- that drops removed files, non-page sources, and
# url:/alias overrides rather than linking them as dead URLs.
changed_pages_section() {
    local repo_api_url=$1
    local pr_number=$2

    local max_listed_pages=50   # Cap the rendered list so huge PRs stay readable.
    local max_files_pages=10    # Cap API pagination (10 * 100 = up to 1000 files).
    local changed_pages=()
    local page files_json page_count url

    for page in $(seq 1 "$max_files_pages"); do
        files_json=$(curl -s \
            -H "Authorization: token ${GITHUB_TOKEN}" \
            "${repo_api_url}/pulls/${pr_number}/files?per_page=100&page=${page}")

        while IFS= read -r url; do
            [[ -z "$url" ]] && continue
            if [[ -f "${build_dir}${url}index.html" ]]; then
                changed_pages+=("- [${url}](${s3_website_url}${url})")
            fi
        done < <(echo "$files_json" \
            | jq -r '.[] | select(.status != "removed") | .filename' 2>/dev/null \
            | changed_paths_to_urls)

        # Stop on the last (short) page, or if the response wasn't an array (e.g. a transient
        # API error), which yields 0 and breaks cleanly.
        page_count="$(echo "$files_json" | jq -r 'if type == "array" then length else 0 end' 2>/dev/null)"
        if [[ "${page_count:-0}" -lt 100 ]]; then
            break
        fi
    done

    if [[ "${#changed_pages[@]}" -eq 0 ]]; then
        return 0
    fi

    local listed
    listed=$(printf '%s\n' "${changed_pages[@]:0:$max_listed_pages}")
    printf '\n\n**Changed pages:**\n%s' "$listed"
    if [[ "${#changed_pages[@]}" -gt "$max_listed_pages" ]]; then
        printf '\n- …and %s more' "$(( ${#changed_pages[@]} - max_listed_pages ))"
    fi
}

# Posts (or updates) the pinned preview comment on the PR.
post_preview_comment() {
    if [[ -z "${GITHUB_EVENT_PATH:-}" || ! -f "${GITHUB_EVENT_PATH}" || -z "${GITHUB_TOKEN:-}" ]]; then
        log "No GitHub event or token available; skipping the preview comment."
        return 0
    fi

    local event pr_comment_api_url repo_api_url pr_number
    event="$(cat "$GITHUB_EVENT_PATH")"
    pr_comment_api_url="$(echo "$event" | jq -r '.pull_request._links.comments.href // empty')"
    repo_api_url="$(echo "$event" | jq -r '.pull_request.base.repo.url // empty')"
    pr_number="$(echo "$event" | jq -r '.number // empty')"

    if [[ -z "$pr_comment_api_url" || -z "$repo_api_url" || -z "$pr_number" ]]; then
        log "Event payload is missing pull request details; skipping the preview comment."
        return 0
    fi

    local body
    body="${PREVIEW_COMMENT_MARKER}
Your site preview for commit $(git_sha_short) is ready! :tada:

${preview_url}$(changed_pages_section "$repo_api_url" "$pr_number")"

    upsert_github_pr_comment \
        "$PREVIEW_COMMENT_MARKER" \
        "$body" \
        "$pr_comment_api_url" \
        "$repo_api_url"
}

# Finally, for previews, point the PR at the resulting bucket URL. Everything below is
# reporting, not deployment, so it must never fail the build -- hence the `|| log`.
if [[ "$1" == "preview" ]]; then
    preview_url="${s3_website_url}/registry/"

    if [[ -n "${GITHUB_STEP_SUMMARY:-}" ]]; then
        echo "Site preview for commit $(git_sha_short): ${preview_url}" >> "$GITHUB_STEP_SUMMARY" || true
    fi

    post_preview_comment || log "Failed to post the preview comment; continuing."
fi

log "Done! The bucket website is now built and available at ${s3_website_url}."
