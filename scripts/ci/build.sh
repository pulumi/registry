#!/bin/bash

set -o errexit -o pipefail

source ./scripts/ci/common.sh

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

printf "Generating API docs from registry commit %s...\n\n" "${REGISTRY_COMMIT}"
pushd tools/resourcedocsgen
go build -o "${GOPATH}/bin/resourcedocsgen" .
popd

case ${1} in
    preview)
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
        export HUGO_BASEURL="http://$(origin_bucket_prefix)-$(build_identifier).s3-website.$(aws_region).amazonaws.com"
        GOGC=3 hugo --minify --buildFuture --templateMetrics -e preview
        ;;
    update)
        REGISTRY_COMMIT="$(go mod graph | grep pulumi/registry/themes/default | sed 's/.*-//')"
        resourcedocsgen docs registry --commitSha "${REGISTRY_COMMIT}" --logtostderr
        
        printf "Running Hugo...\n\n"
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
