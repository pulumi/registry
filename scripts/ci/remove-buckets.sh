#!/bin/bash

set -o errexit -o pipefail

# This script uses the list-buckets script to find buckets that can be safely
# deleted, then deletes them. See that script for details, but in general, a bucket is
# deletable if and only if:
#
# * It's associated with a PR that's been closed.

echo "Finding deletable $1 buckets..."

source ./scripts/ci/common.sh

echo "origin bucket prefix: $(origin_bucket_prefix)"

buckets_to_remove="$(./scripts/ci/list-buckets.sh "$1" --only-deletables)"

if [ -z "$buckets_to_remove" ]; then
    echo "None found."
    exit
fi

echo
echo "Buckets to remove:"
echo "$buckets_to_remove"
echo


# This is a lifecycle policy that deletes all objects and delete markers after 1 day (minimum).
# After ~48 hours, the bucket will be empty and can be deleted.
lifecycle_file=$(mktemp)
trap 'rm -f "$lifecycle_file"' EXIT

printf '{
  "Rules": [
    {
      "Expiration": {
        "Days": 1
      },
      "ID": "FullDelete",
      "Filter": {
        "Prefix": ""
      },
      "Status": "Enabled",
      "NoncurrentVersionExpiration": {
        "NoncurrentDays": 1
      },
      "AbortIncompleteMultipartUpload": {
        "DaysAfterInitiation": 1
      }
    },
    {
      "Expiration": {
        "ExpiredObjectDeleteMarker": true
      },
      "ID": "DeleteMarkers",
      "Filter": {
        "Prefix": ""
      },
      "Status": "Enabled"
    }
  ]
}' > "$lifecycle_file"

for bucket in $buckets_to_remove; do
    echo "Removing ${bucket}..."
    echo "Upserting lifecycle policy for ${bucket}..."
    aws s3api put-bucket-lifecycle-configuration --bucket "${bucket}" --lifecycle-configuration "file://${lifecycle_file}"

    # Check if bucket already has cleanup timestamp
    existing_tags=$(aws s3api get-bucket-tagging --bucket "${bucket}" 2>/dev/null || echo '{"TagSet":[]}')
    cleanup_started=$(echo "$existing_tags" | jq -r '.TagSet[] | select(.Key=="CleanupStarted") | .Value // empty')
    
    if [ -z "$cleanup_started" ]; then
        # No cleanup tag exists, add it while preserving existing tags
        timestamp=$(date -u +"%Y-%m-%dT%H:%M:%SZ")
        # Create new tag set by combining existing tags with the new cleanup tag
        new_tags=$(echo "$existing_tags" | jq --arg ts "$timestamp" '.TagSet += [{"Key": "CleanupStarted", "Value": $ts}]')
        aws s3api put-bucket-tagging --bucket "${bucket}" --tagging "$new_tags"
        echo "Added cleanup timestamp tag to ${bucket}. It should be safe to delete within ~48 hours."
    else
        # Calculate time difference
        current_time=$(date -u +%s)
        # Parse the cleanup timestamp to Unix time, handling both Linux (-d) and macOS (-j) date syntax
        cleanup_time=$(date -u -d "$cleanup_started" +%s 2>/dev/null || date -u -j -f "%Y-%m-%dT%H:%M:%SZ" "$cleanup_started" +%s)
        # Calculate seconds elapsed since cleanup started
        time_diff=$((current_time - cleanup_time))

        if [ $time_diff -gt 172800 ]; then  # 48 hours = 172800 seconds
            echo "Attempting to delete bucket ${bucket} after 48+ hours of cleanup..."
            if ! aws s3 rb "s3://${bucket}"; then
                if [ $time_diff -gt 604800 ]; then  # 7 days = 604800 seconds
                    echo "Error: Bucket ${bucket} has been in cleanup state for more than 7 days. Please resolve the issue and try again."
                    exit 1
                fi
                echo "Warning: Failed to delete bucket ${bucket}, cleanup might still be in progress"
            else
                remove_param_for_commit "$(git_sha)" "$(aws_region)"
            fi
        else
            echo "Bucket ${bucket} is still in cleanup idle period (${time_diff} seconds elapsed)"
        fi
    fi

    echo
done

echo "Done!"
