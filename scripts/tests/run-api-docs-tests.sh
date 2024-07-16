#!/bin/bash


# Get currently website bucket

# bucket="$(curl -s https://www.pulumi.com/registry/metadata.json | jq -r '.bucket' || echo '')"

bucket="registry-production-origin-push-c9412e72"

aws s3 cp "s3://${bucket}/registry" public --recursive

npm run api-docs-tests

