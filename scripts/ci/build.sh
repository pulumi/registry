#!/bin/bash

set -o errexit -o pipefail

source ./scripts/ci/common.sh

# URL to the Pulumi conversion service.
export PULUMI_CONVERT_URL="${PULUMI_CONVERT_URL:-$(pulumi stack output --stack pulumi/tf2pulumi-service/production url)}"

go install github.com/pulumi/docs/tools/resourcedocsgen@master

PKGS=(
    "aiven"
    "aws"
)

echo "Generating API docs for ${PKGS[*]}..."
echo ""

for PKG in "${PKGS[@]}" ; do \
    resourcedocsgen docs registry "${PKG}" \
        --commitSha "$(git_sha_short)" \
        --baseDocsOutDir "themes/default/content/registry/packages" \
        --basePackageTreeJSONOutDir "themes/default/static/registry/packages/navs" \
        --logtostderr
done

printf "Running Hugo...\n\n"
export REPO_THEME_PATH="themes/default/"
export HUGO_BASEURL="http://$(origin_bucket_prefix)-$(build_identifier).s3-website.$(aws_region).amazonaws.com"
GOGC=5 hugo --minify --templateMetrics --buildDrafts --buildFuture -e "preview" | grep -v -e 'WARN .* REF_NOT_FOUND'

printf "Done!\n\n"
