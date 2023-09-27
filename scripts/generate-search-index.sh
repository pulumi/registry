#!/bin/bash

set -o errexit -o pipefail

source ./scripts/ci/common.sh

# Run the script that updates the Algolia search index. The value passed into this script denotes
# the name of the index to be updated (e.g., 'preview' or 'production'). The index will be written
# to a file named 'search-index.json'.
node ./scripts/search/main.js "$1"

registry_bucket="$(cat "$(origin_bucket_metadata_filepath)" | jq -r ".bucket")"

# Upload the `search-index.json` file to S3 where it can be accessed by the update search index cron.
aws s3 cp "search-index.json" "s3://${registry_bucket}/registry/search-index.json" --acl public-read --region "$(aws_region)"