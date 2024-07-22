#!/bin/bash


# Get currently website bucket

bucket="$(curl -s https://www.pulumi.com/registry/metadata.json | jq -r '.bucket' || echo '')"

# bucket="registry-production-origin-push-c4c4803f"

echo "Downloading current site to public dir...."
aws s3 cp "s3://${bucket}/registry" public/registry --recursive --quiet


# pkgs=("aws" "gcp" "azure")

# Iterate over the array
# for str in "${pkgs[@]}"; do
#     # Invoke the command with the string as the first argument
#     npm run test-api-docs $str
# done

npm run test-api-docs

