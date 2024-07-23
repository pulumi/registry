#!/bin/bash


# Get current website bucket.
bucket="$(curl -s https://www.pulumi.com/registry/metadata.json | jq -r '.bucket' || echo '')"

echo "Downloading current site to public dir...."
aws s3 cp "s3://${bucket}/registry" public/registry --recursive --quiet

node ./scripts/run-browser-tests.js
