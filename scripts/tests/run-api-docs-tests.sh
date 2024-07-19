#!/bin/bash


# Get currently website bucket

# bucket="$(curl -s https://www.pulumi.com/registry/metadata.json | jq -r '.bucket' || echo '')"

bucket="registry-production-origin-push-c530c8d7"

echo "Downloading current site to public dir...."
aws s3 cp "s3://${bucket}/registry" public --recursive --quiet

npm run test-api-docs

