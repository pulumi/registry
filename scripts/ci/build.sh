#!/bin/bash

set -o errexit -o pipefail

source ./scripts/ci/common.sh

# URL to the Pulumi conversion service.
export PULUMI_CONVERT_URL="${PULUMI_CONVERT_URL:-$(pulumi stack output --stack pulumi/tf2pulumi-service/production url)}"

PR_COMMIT_HASH=$(build_identifier)

echo "Generating API docs for all packages for PR preview..."
echo ""
go install github.com/pulumi/docs/tools/resourcedocsgen@master
resourcedocsgen docs registry \
    --commitSha "${PR_COMMIT_HASH}" \
    --baseDocsOutDir "themes/default/content/registry/packages" \
    --basePackageTreeJSONOutDir "themes/default/static/registry/packages/navs" \
    --logtostderr

printf "Running Hugo...\n\n"
export REPO_THEME_PATH="themes/default/"
export HUGO_BASEURL="http://$(origin_bucket_prefix)-${PR_COMMIT_HASH}.s3-website.$(aws_region).amazonaws.com"
hugo --minify --templateMetrics --buildDrafts --buildFuture -e "preview" | grep -v -e 'WARN .* REF_NOT_FOUND'

printf "Done!\n\n"
