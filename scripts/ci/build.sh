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

REGISTRY_COMMIT="$(git_sha_short)"
printf "Generating API docs from registry commit %s...\n\n" "${REGISTRY_COMMIT}"

# Restore cached API docs output so resourcedocsgen can skip unchanged packages.
# The .generated sentinel files inside each api-docs/ directory record the cache key;
# if the key matches, the Go tool skips schema fetching and doc generation entirely.
API_DOCS_CACHE=".cache/api-docs"
CONTENT_DIR="themes/default/content/registry/packages"
NAVS_DIR="themes/default/static/registry/packages/navs"
SCHEMAS_DIR="themes/default/static/registry/packages"
LLM_DOCS_DIR="llm-docs-out/registry/packages"

if [[ -d "$API_DOCS_CACHE/content" ]]; then
    log "Restoring cached API docs..."
    mkdir -p "$CONTENT_DIR" "$NAVS_DIR"
    count=0
    for pkg_dir in "$API_DOCS_CACHE/content"/*/; do
        [[ -d "$pkg_dir/api-docs" ]] || continue
        pkg=$(basename "$pkg_dir")
        if [[ ! -d "$CONTENT_DIR/$pkg/api-docs" ]]; then
            mkdir -p "$CONTENT_DIR/$pkg"
            cp -a "$pkg_dir/api-docs" "$CONTENT_DIR/$pkg/api-docs"
            ((count++)) || true
        fi
    done
    for nav in "$API_DOCS_CACHE/navs"/*.json; do
        [[ -f "$nav" ]] || continue
        name=$(basename "$nav")
        [[ -f "$NAVS_DIR/$name" ]] || cp -a "$nav" "$NAVS_DIR/$name"
    done
    for schema_dir in "$API_DOCS_CACHE/schemas"/*/; do
        [[ -d "$schema_dir" ]] || continue
        pkg=$(basename "$schema_dir")
        [[ -d "$SCHEMAS_DIR/$pkg" ]] || cp -a "$schema_dir" "$SCHEMAS_DIR/$pkg"
    done
    if [[ -d "$API_DOCS_CACHE/llm-docs" ]]; then
        mkdir -p "$LLM_DOCS_DIR"
        for pkg_dir in "$API_DOCS_CACHE/llm-docs"/*/; do
            [[ -d "$pkg_dir/api-docs" ]] || continue
            pkg=$(basename "$pkg_dir")
            if [[ ! -d "$LLM_DOCS_DIR/$pkg/api-docs" ]]; then
                mkdir -p "$LLM_DOCS_DIR/$pkg"
                cp -a "$pkg_dir/api-docs" "$LLM_DOCS_DIR/$pkg/api-docs"
            fi
        done
    fi
    log "Restored $count cached package doc sets"
fi


make api-docs

# Save API docs output to cache for next run.
log "Saving API docs to cache..."
rm -rf "$API_DOCS_CACHE"
mkdir -p "$API_DOCS_CACHE/content" "$API_DOCS_CACHE/navs" "$API_DOCS_CACHE/schemas" "$API_DOCS_CACHE/llm-docs"
for pkg_dir in "$CONTENT_DIR"/*/; do
    [[ -d "$pkg_dir/api-docs" ]] || continue
    pkg=$(basename "$pkg_dir")
    [[ "$pkg" == *@* ]] && continue  # versioned packages have their own cache
    mkdir -p "$API_DOCS_CACHE/content/$pkg"
    cp -a "$pkg_dir/api-docs" "$API_DOCS_CACHE/content/$pkg/api-docs"
done
for nav in "$NAVS_DIR"/*.json; do
    [[ -f "$nav" ]] || continue
    name=$(basename "$nav")
    [[ "$name" == *@* ]] && continue
    cp -a "$nav" "$API_DOCS_CACHE/navs/"
done
for schema_dir in "$SCHEMAS_DIR"/*/; do
    [[ -d "$schema_dir" && -f "$schema_dir/schema.json" ]] || continue
    pkg=$(basename "$schema_dir")
    [[ "$pkg" == *@* ]] && continue
    # Only cache schema.json, not the entire directory tree (which may contain
    # stale LLM doc files from older builds).
    mkdir -p "$API_DOCS_CACHE/schemas/$pkg"
    cp -a "$schema_dir/schema.json" "$API_DOCS_CACHE/schemas/$pkg/schema.json"
done
for pkg_dir in "$LLM_DOCS_DIR"/*/; do
    [[ -d "$pkg_dir/api-docs" ]] || continue
    pkg=$(basename "$pkg_dir")
    [[ "$pkg" == *@* ]] && continue
    mkdir -p "$API_DOCS_CACHE/llm-docs/$pkg"
    cp -a "$pkg_dir/api-docs" "$API_DOCS_CACHE/llm-docs/$pkg/api-docs"
done
log "API docs cache saved"

# Apply fixes. See script for details.
node ./scripts/apply-fixes.js

log "Running Hugo..."
echo
echo

case ${1} in
    preview)
        export HUGO_BASEURL="http://$(origin_bucket_prefix)-$(build_identifier).s3-website.$(aws_region).amazonaws.com"
        hugo --minify --buildFuture --templateMetrics -e preview
        ;;
    update)
        hugo --minify --buildFuture --templateMetrics -e production
        ;;
    *)
        echo "Unknown mode, '${1}' must be one of 'preview' or 'update'"
        exit 1
        ;;
esac

# Split the sitemap into multiple files if it exceeds the per-file URL limit.
yarn run split-sitemap

# Purge unused CSS and write commit-hash-named bundles that match the <link>
# hrefs emitted by assets.html, avoiding stale CloudFront caches (#10466).
yarn run minify-css

# Inline critical CSS for key pages.
node scripts/inline-critical-css.js

log "Done!"
echo
echo
