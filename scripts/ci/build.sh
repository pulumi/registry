#!/bin/bash

set -o errexit -o pipefail

source ./scripts/ci/common.sh

echo "Building CSS and JS assets....."
make build-assets

# URL to the Pulumi conversion service.
export PULUMI_CONVERT_URL="${PULUMI_CONVERT_URL:-$(pulumi stack output --stack pulumi/tf2pulumi-service/production url)}"
export PULUMI_AI_WS_URL=${PULUMI_AI_WS_URL:-$(pulumi stack output --stack pulumi/pulumigpt-api/corp websocketUri)}

printf "Compiling theme JavaScript and CSS...\n\n"
export ASSET_BUNDLE_ID="$(build_identifier)"

# Paths to the CSS and JS bundles we'll generate below. Note that environment variables
# are read by some templates during the Hugo build process.
export CSS_BUNDLE="static/css/styles.${ASSET_BUNDLE_ID}.css"
export JS_BUNDLE="static/js/bundle.min.${ASSET_BUNDLE_ID}.js"

# Relative paths to those same files, read by Hugo templates.
export REL_CSS_BUNDLE="/css/styles.${ASSET_BUNDLE_ID}.css"
export REL_JS_BUNDLE="/js/bundle.min.${ASSET_BUNDLE_ID}.js"
export REPO_THEME_PATH="themes/default/"

pushd tools/resourcedocsgen
go build -o "${GOPATH}/bin/resourcedocsgen" .
popd

REGISTRY_COMMIT="$(git_sha_short)"
printf "Generating API docs from registry commit %s...\n\n" "${REGISTRY_COMMIT}"
resourcedocsgen docs registry --commitSha "${REGISTRY_COMMIT}" \
    --baseDocsOutDir "themes/default/content/registry/packages" \
    --basePackageTreeJSONOutDir "themes/default/static/registry/packages/navs" \
    --logtostderr

printf "Running Hugo...\n\n"

case ${1} in
    preview)
        export HUGO_BASEURL="http://$(origin_bucket_prefix)-$(build_identifier).s3-website.$(aws_region).amazonaws.com"
        GOGC=3 hugo --minify --buildFuture --templateMetrics -e preview
        ;;
    update)
        GOGC=3 hugo --minify --buildFuture --templateMetrics -e production
        ;;
    *)
        echo "Unknown mode, '${1}' must be one of 'preview' or 'update'"
        exit 1
        ;;
esac

# Purge unused CSS.
yarn run minify-css

printf "Done!\n\n"
