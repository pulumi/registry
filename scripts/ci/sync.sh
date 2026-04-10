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

# Extract schema.json files out of the Hugo build tree and gzip them into a sibling
# schema-out/ tree. They get uploaded separately (below, after the main sync) with
# Content-Encoding: gzip set on the object, so both CloudFront layers pass them through
# verbatim without relying on automatic compression -- which is disabled on /registry/*
# at the outer CDN (Accept-Encoding forwarded via AllViewerExceptHostHeader) and
# wouldn't help anyway for schemas over the 10 MB auto-compress ceiling
# (aws/schema.json is >50 MB).
#
# This pull-aside has to happen BEFORE the main `s5cmd sync --delete` runs: removing
# the raw schema.json files from the source tree first means that (a) the main sync
# won't upload them with text/html or application/json content-type and no encoding
# header, and (b) --delete will clean up any raw schema.json objects left in the
# destination bucket by previous commits on the same PR preview bucket.
schema_out_dir="schema-out"
rm -rf "$schema_out_dir"
if find "$build_dir/registry/packages" -name 'schema.json' -type f -print -quit 2>/dev/null | grep -q .; then
    log "Extracting and pre-gzipping schema.json files..."
    find "$build_dir/registry/packages" -name 'schema.json' -type f -print0 \
        | xargs -0 -P"$(nproc)" -I{} sh -c '
            src="$1"
            build_dir="$2"
            schema_out_dir="$3"
            rel="${src#"$build_dir/"}"
            dest="$schema_out_dir/$rel"
            mkdir -p "$(dirname "$dest")"
            gzip -9 < "$src" > "$dest"
            rm "$src"
        ' _ {} "$build_dir" "$schema_out_dir"
    log "Schema extraction complete."
else
    log "No schema.json files found under $build_dir/registry/packages; skipping schema extraction."
fi

# Sync the local build directory to the bucket using s5cmd for massively parallel uploads.
# s5cmd uses hundreds of concurrent goroutines vs aws cli's ~10-16 concurrent requests,
# resulting in 10-50x faster uploads for large file counts.
# The --delete flag removes destination objects not present locally, keeping the bucket clean
# for PR preview buckets that get reused across commits.
log "Synchronizing to $destination_bucket_uri..."
s5cmd --log error sync --delete --acl public-read \
    "$build_dir/" "$destination_bucket_uri/"
log "Sync complete."

# Sync CLI docs separately. These are generated outside the Hugo build tree to avoid
# Hugo processing static files. They're uploaded directly to the same URL paths
# they'd occupy if they were in the Hugo static directory.
#
# CLI docs bundles are uploaded gzip-compressed with Content-Encoding: gzip so they
# pass through both CloudFront layers (registry CDN + www.pulumi.com CDN) without
# relying on CloudFront's automatic compression -- which is disabled on /registry/*
# because the outer CDN forwards Accept-Encoding via AllViewerExceptHostHeader, and
# wouldn't help anyway for bundles over CloudFront's 10 MB auto-compress ceiling
# (the aws bundle is >100 MB uncompressed). Go's http.Client decompresses
# Content-Encoding: gzip transparently, so the pulumi docs CLI requires no changes.
#
# The cli-docs-out/ tree contains only cli-docs.json files, so it is safe to apply
# --content-encoding/--content-type globally to every object in the sync.
cli_docs_dir="cli-docs-out"
if [[ -d "$cli_docs_dir" ]]; then
    log "Pre-gzipping CLI docs..."
    find "$cli_docs_dir" -type f -name 'cli-docs.json' -print0 \
        | xargs -0 -P"$(nproc)" -I{} sh -c 'gzip -9 < "$1" > "$1.gz.tmp" && mv "$1.gz.tmp" "$1"' _ {}

    log "Synchronizing CLI docs to $destination_bucket_uri..."
    s5cmd --log error sync --acl public-read \
        --content-encoding gzip \
        --content-type application/json \
        "$cli_docs_dir/" "$destination_bucket_uri/"
    log "CLI docs sync complete."
fi

# Sync the extracted + gzipped schema.json tree separately, with Content-Encoding: gzip
# and Content-Type: application/json set on every object. No --delete here: the main
# sync above already reconciled the destination against the (schema-free) public/ tree,
# so this pass only needs to add the correctly-tagged gzipped schemas back.
if [[ -d "$schema_out_dir" ]]; then
    log "Synchronizing schemas to $destination_bucket_uri..."
    s5cmd --log error sync --acl public-read \
        --content-encoding gzip \
        --content-type application/json \
        "$schema_out_dir/" "$destination_bucket_uri/"
    log "Schema sync complete."
fi

s3_website_url="http://${destination_bucket}.s3-website.$(aws_region).amazonaws.com"
echo "$s3_website_url"

# Set the content-type of latest-version explicitly. (Otherwise, it'll be set as binary/octet-stream.)
aws s3 cp "$build_dir/latest-version" "${destination_bucket_uri}/latest-version" \
    --content-type "text/plain" --acl public-read --region "$(aws_region)" --metadata-directive REPLACE

# Smoke test origin-gzipped assets at the S3 website layer (before either CloudFront
# layer has a chance to mask a regression). Uses the random package because it has
# small, reliably-present cli-docs.json and schema.json bundles. Catches the regression
# where we accidentally upload either file without Content-Encoding: gzip, which would
# otherwise only surface after deploy when consumers start paying full wire cost again.
#
# Uses curl --dump-header to validate the body (gunzip -t) and capture the headers in
# a single round trip, rather than issuing separate HEAD and GET requests.
check_gzip_asset() {
    local label="$1"
    local url="$2"
    log "Smoke-testing $label compression at $url..."
    local headers_file
    headers_file=$(mktemp)
    if ! curl -fsS --dump-header "$headers_file" "$url" | gunzip -t; then
        echo "ERROR: $url body did not decompress as gzip (or request failed)." >&2
        rm -f "$headers_file"
        exit 1
    fi
    local headers
    headers=$(cat "$headers_file")
    rm -f "$headers_file"
    if ! grep -qi '^content-encoding: gzip' <<< "$headers"; then
        echo "ERROR: $url is missing 'Content-Encoding: gzip'. Response headers:" >&2
        echo "$headers" >&2
        exit 1
    fi
    log "$label compression smoke test passed."
}

if [[ -d "$cli_docs_dir" && -f "$cli_docs_dir/registry/packages/random/api-docs/cli-docs.json" ]]; then
    check_gzip_asset "CLI docs" "${s3_website_url}/registry/packages/random/api-docs/cli-docs.json"
else
    log "Skipping CLI docs compression smoke test: $cli_docs_dir/registry/packages/random/api-docs/cli-docs.json not found."
fi

if [[ -d "$schema_out_dir" && -f "$schema_out_dir/registry/packages/random/schema.json" ]]; then
    check_gzip_asset "Schema" "${s3_website_url}/registry/packages/random/schema.json"
else
    log "Skipping schema compression smoke test: $schema_out_dir/registry/packages/random/schema.json not found."
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

# Finally, post a comment to the PR that directs the user to the resulting bucket URL.
if [[ "$1" == "preview" ]]; then
    pr_comment_api_url="$(cat "$GITHUB_EVENT_PATH" | jq -r ".pull_request._links.comments.href")"
    post_github_pr_comment \
        "Your site preview for commit $(git_sha_short) is ready! :tada:\n\n${s3_website_url}/registry." \
        $pr_comment_api_url
fi

log "Done! The bucket website is now built and available at ${s3_website_url}."
