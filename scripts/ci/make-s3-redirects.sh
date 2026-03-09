#!/bin/bash

set -o errexit -o pipefail

source ./scripts/ci/common.sh

# Create an S3 object for each of the items in the redirect list so it returns a 301
# redirect (instead of serving the HTML with a meta-redirect). This ensures the right HTTP
# response code is returned for search engines and enables better support for URL anchors.

destination_bucket="$(cat "$(origin_bucket_metadata_filepath)" | jq -r ".bucket")"
region="$(aws_region)"

IFS="|"

# Collect all redirects from all files, then apply them in parallel.
pids=()
for redirect_file in ./scripts/redirects/*.txt; do
    [[ -f "$redirect_file" ]] || continue
    while read key location; do
        [[ -z "$key" ]] && continue
        echo "Redirecting $key to $location"
        aws s3api put-object \
            --key "$key" \
            --website-redirect-location "$location" \
            --bucket "$destination_bucket" \
            --acl public-read \
            --region "$region" > /dev/null &
        pids+=($!)
    done < "$redirect_file"
done

# Wait for all uploads and fail if any failed.
failed=0
for pid in "${pids[@]}"; do
    if ! wait "$pid"; then
        ((failed++)) || true
    fi
done

if (( failed > 0 )); then
    echo "ERROR: $failed redirect(s) failed to apply"
    exit 1
fi
