#!/usr/bin/env bash
set -euo pipefail

for cmd in jq yq curl; do
    command -v "$cmd" >/dev/null || { echo "Missing dependency: $cmd" >&2; exit 1; }
done

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"

# Use pre-built resourcedocsgen binary if available, otherwise build it.
RESOURCEDOCSGEN="${RESOURCEDOCSGEN:-$REPO_ROOT/bin/resourcedocsgen}"
if [[ ! -x "$RESOURCEDOCSGEN" ]]; then
    echo "Building resourcedocsgen..."
    make -C "$REPO_ROOT" bin/resourcedocsgen
fi

# Install registry-mirror-discover (pinned commit for stability).
# Cache the built binary to avoid recompiling on every build.
REGISTRY_MIRROR_TOOLS_COMMIT="becca918a99a2a0809ad7d35a551b8a4a8018639"
DISCOVER_BIN="$REPO_ROOT/bin/registry-mirror-discover"
if [[ ! -x "$DISCOVER_BIN" ]]; then
    echo "Building registry-mirror-discover..."
    GOBIN="$REPO_ROOT/bin" GOPRIVATE=github.com/pulumi go install "github.com/pulumi/registry-mirror-tools/cmd/registry-mirror-discover@${REGISTRY_MIRROR_TOOLS_COMMIT}"
fi

# Allow overriding output directories (for local vs CI paths)
CONTENT_DIR="${CONTENT_DIR:-$REPO_ROOT/themes/default/content/registry/packages}"
STATIC_DIR="${STATIC_DIR:-$REPO_ROOT/themes/default/static/registry/packages}"

# Versioned metadata goes to a separate directory so Hugo sees it as
# $.Site.Data.registry.package_versions (not mixed with packages/)
PACKAGE_VERSIONS_DIR="$REPO_ROOT/themes/default/data/registry/package_versions"

MAJOR_COUNT=3
MINOR_COUNT=1

# Maximum number of packages to process in parallel.
MAX_PARALLEL="${MAX_PARALLEL:-8}"

# Blessed packages are first-party providers that get versioned docs.
# Fetch dynamically from ci-mgmt (pinned commit for stability):
CI_MGMT_COMMIT="b3cede4113152996ec67c47a2ca0cef3c5aeb626"
PROVIDERS_JSON_URL="https://raw.githubusercontent.com/pulumi/ci-mgmt/${CI_MGMT_COMMIT}/provider-ci/providers.json"
PROVIDERS_JSON=$(curl -sfL "$PROVIDERS_JSON_URL") || { echo "Failed to fetch providers.json from ci-mgmt" >&2; exit 1; }
mapfile -t BLESSED_PACKAGES < <(echo "$PROVIDERS_JSON" | jq -r '.[]')

is_blessed() {
    local pkg="$1"
    for blessed in "${BLESSED_PACKAGES[@]}"; do
        if [[ "$pkg" == "$blessed" ]]; then
            return 0
        fi
    done
    return 1
}

# If a package is specified, process just that one; otherwise process all blessed packages
if [[ -n "${1:-}" ]]; then
    PACKAGES_TO_PROCESS=("$1")
else
    PACKAGES_TO_PROCESS=("${BLESSED_PACKAGES[@]}")
fi

# Derive index URL from schema URL.
#
# The _index.md URL is not stored in package metadata. For some packages (third-party providers
# using the `from-urls` workflow), it's only passed during the CI publish workflow and discarded.
# For versioned docs, we derive it from the schema URL pattern:
#
# 1. GitHub-hosted packages (Pulumi providers):
#    Schema: https://raw.githubusercontent.com/pulumi/pulumi-aws/v7.20.0/provider/cmd/pulumi-resource-aws/schema.json
#    Index:  https://raw.githubusercontent.com/pulumi/pulumi-aws/v7.20.0/docs/_index.md
#
# 2. CDN-hosted packages (bridged OpenTofu providers):
#    Schema: https://djoiyj6oj2oxz.cloudfront.net/schemas/registry.opentofu.org/elastic/elasticstack/0.14.3/schema.json
#    Index:  https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/elastic/elasticstack/0.14.3/index.md
#
# If derivation fails or the URL returns 404, we fall back to copying the latest version's _index.md.
derive_index_url() {
    local schema_url="$1"

    if [[ "$schema_url" == *"raw.githubusercontent.com"* ]]; then
        # GitHub URL: extract base and replace path
        # Example: https://raw.githubusercontent.com/pulumi/pulumi-aws/v7.20.0/provider/cmd/pulumi-resource-aws/schema.json
        # Result: https://raw.githubusercontent.com/pulumi/pulumi-aws/v7.20.0/docs/_index.md
        echo "$schema_url" | sed -E 's|^(https://raw\.githubusercontent\.com/[^/]+/[^/]+/[^/]+)/.*$|\1/docs/_index.md|'
    elif [[ "$schema_url" == *"cloudfront.net"* ]]; then
        # CDN URL: replace /schemas/ with /docs/ and schema.json with index.md
        # Example: https://djoiyj6oj2oxz.cloudfront.net/schemas/registry.opentofu.org/elastic/elasticstack/0.14.3/schema.json
        # Result: https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/elastic/elasticstack/0.14.3/index.md
        echo "$schema_url" | sed 's|/schemas/|/docs/|; s|schema\.json$|index.md|'
    else
        echo ""
    fi
}

mkdir -p "$CONTENT_DIR" "$STATIC_DIR/navs" "$PACKAGE_VERSIONS_DIR"

# Persistent cache for versioned docs output. In CI, this directory is preserved between
# runs by actions/cache. On cache hit, the sentinel-based skip logic avoids regenerating
# unchanged packages entirely. The cache stores only versioned (@-suffixed) content.
VERSIONED_DOCS_CACHE="$REPO_ROOT/.cache/versioned-docs"
mkdir -p "$VERSIONED_DOCS_CACHE/content" "$VERSIONED_DOCS_CACHE/navs" "$VERSIONED_DOCS_CACHE/metadata"

# Restore cached versioned docs into the output directories.
restore_cache() {
    local count=0
    for cached_dir in "$VERSIONED_DOCS_CACHE/content/"*@*; do
        [[ -d "$cached_dir" ]] || continue
        local name
        name=$(basename "$cached_dir")
        if [[ ! -d "$CONTENT_DIR/$name" ]]; then
            cp -a "$cached_dir" "$CONTENT_DIR/$name"
            ((count++)) || true
        fi
    done
    for cached_nav in "$VERSIONED_DOCS_CACHE/navs/"*@*.json; do
        [[ -f "$cached_nav" ]] || continue
        local name
        name=$(basename "$cached_nav")
        [[ -f "$STATIC_DIR/navs/$name" ]] || cp -a "$cached_nav" "$STATIC_DIR/navs/$name"
    done
    for cached_meta in "$VERSIONED_DOCS_CACHE/metadata/"*@*.yaml; do
        [[ -f "$cached_meta" ]] || continue
        local name
        name=$(basename "$cached_meta")
        [[ -f "$PACKAGE_VERSIONS_DIR/$name" ]] || cp -a "$cached_meta" "$PACKAGE_VERSIONS_DIR/$name"
    done
    if (( count > 0 )); then
        echo "Restored $count versioned doc sets from cache"
    fi
}

# Save versioned docs output back to the cache for next run.
save_cache() {
    # Clear stale cache entries and repopulate from current output.
    rm -rf "$VERSIONED_DOCS_CACHE/content/"*@* "$VERSIONED_DOCS_CACHE/navs/"*@* "$VERSIONED_DOCS_CACHE/metadata/"*@*
    for content_dir in "$CONTENT_DIR/"*@*; do
        [[ -d "$content_dir" ]] || continue
        cp -a "$content_dir" "$VERSIONED_DOCS_CACHE/content/"
    done
    for nav_file in "$STATIC_DIR/navs/"*@*.json; do
        [[ -f "$nav_file" ]] || continue
        cp -a "$nav_file" "$VERSIONED_DOCS_CACHE/navs/"
    done
    for meta_file in "$PACKAGE_VERSIONS_DIR/"*@*.yaml; do
        [[ -f "$meta_file" ]] || continue
        cp -a "$meta_file" "$VERSIONED_DOCS_CACHE/metadata/"
    done
    echo "Saved versioned docs to cache"
}

restore_cache

# Process all older major versions for a single package.
# Designed to run as a background job alongside other packages.
process_package() {
    local PACKAGE="$1"

    if ! is_blessed "$PACKAGE"; then
        return 0
    fi

    local PACKAGE_YAML="$REPO_ROOT/themes/default/data/registry/packages/${PACKAGE}.yaml"
    if [[ ! -f "$PACKAGE_YAML" ]]; then
        echo "[$PACKAGE] Package YAML not found: $PACKAGE_YAML, skipping" >&2
        return 0
    fi

    # Per-invocation temp dir for resourcedocsgen nav output
    local NAV_TEMP_DIR
    NAV_TEMP_DIR=$(mktemp -d)
    trap 'rm -rf "$NAV_TEMP_DIR"' EXIT

    echo ""
    echo "========================================"
    echo "Processing $PACKAGE"
    echo "========================================"

    echo "[$PACKAGE] Discovering versions..."
    local VERSIONS_JSON
    VERSIONS_JSON=$("$DISCOVER_BIN" \
        "$REPO_ROOT" \
        --packages "$PACKAGE" \
        --major-count "$MAJOR_COUNT" \
        --minor-count "$MINOR_COUNT" \
        --format json-array < /dev/null)

    local VERSION_COUNT
    VERSION_COUNT=$(echo "$VERSIONS_JSON" | jq length)
    echo "[$PACKAGE] Found $VERSION_COUNT versions"

    local LATEST_MAJOR
    LATEST_MAJOR=$(echo "$VERSIONS_JSON" | jq -r '.[0].version' | cut -d. -f1)
    echo "[$PACKAGE] Latest major version is $LATEST_MAJOR.x (already generated above)"

    local GENERATED_VERSIONS=()

    while read -r VERSION_INFO; do
        local VERSION SCHEMA_URL MAJOR VERSION_SLUG VERSIONED_NAME
        VERSION=$(echo "$VERSION_INFO" | jq -r '.version')
        SCHEMA_URL=$(echo "$VERSION_INFO" | jq -r '.schema_url')

        MAJOR=$(echo "$VERSION" | cut -d. -f1)
        VERSION_SLUG="${MAJOR}.x"
        VERSIONED_NAME="${PACKAGE}@${VERSION_SLUG}"

        if [[ "$MAJOR" == "$LATEST_MAJOR" ]]; then
            continue
        fi

        if [[ -z "$SCHEMA_URL" || "$SCHEMA_URL" == "null" ]]; then
            echo "[$PACKAGE] Warning: No schema URL for $VERSIONED_NAME, skipping"
            continue
        fi

        # Skip unchanged packages: if docs were already generated from this exact schema URL
        # with this exact version of resourcedocsgen, reuse them.
        local DOCS_OUT_DIR="$CONTENT_DIR/$VERSIONED_NAME"
        local SENTINEL="$DOCS_OUT_DIR/.generated"
        if [[ -f "$SENTINEL" ]] && grep -qF "$SCHEMA_URL" "$SENTINEL" 2>/dev/null; then
            echo "[$PACKAGE] Skipping $VERSIONED_NAME (already generated, inputs unchanged)"
            GENERATED_VERSIONS+=("$VERSIONED_NAME")
            continue
        fi

        echo ""
        echo "[$PACKAGE] === Processing v$VERSION ($VERSION_SLUG) ==="

        rm -rf "$DOCS_OUT_DIR"
        mkdir -p "$DOCS_OUT_DIR"

        # Download schema once and cache it (schemas can be 50MB+)
        local SCHEMA_CACHE_DIR="$REPO_ROOT/.cache/schemas"
        mkdir -p "$SCHEMA_CACHE_DIR"
        local CACHED_SCHEMA="$SCHEMA_CACHE_DIR/${PACKAGE}-v${VERSION}.json"

        if [[ ! -f "$CACHED_SCHEMA" ]]; then
            echo "[$PACKAGE] Downloading schema from $SCHEMA_URL..."
            if [[ "$SCHEMA_URL" == *.yaml || "$SCHEMA_URL" == *.yml ]]; then
                curl -sfL "$SCHEMA_URL" | yq -o json > "$CACHED_SCHEMA"
            else
                curl -sfL "$SCHEMA_URL" -o "$CACHED_SCHEMA"
            fi
        else
            echo "[$PACKAGE] Using cached schema: $CACHED_SCHEMA"
        fi

        # Fetch _index.md from derived URL, fall back to latest version's copy
        local INDEX_URL
        INDEX_URL=$(derive_index_url "$SCHEMA_URL")
        if [[ -n "$INDEX_URL" ]]; then
            echo "[$PACKAGE] Fetching _index.md from $INDEX_URL..."
            if ! curl -sfL "$INDEX_URL" -o "$DOCS_OUT_DIR/_index.md" 2>/dev/null; then
                echo "[$PACKAGE] Could not fetch _index.md from $INDEX_URL"
                rm -f "$DOCS_OUT_DIR/_index.md"
            fi
        fi

        if [[ ! -f "$DOCS_OUT_DIR/_index.md" ]]; then
            local LATEST_INDEX="$REPO_ROOT/themes/default/content/registry/packages/$PACKAGE/_index.md"
            if [[ -f "$LATEST_INDEX" ]]; then
                echo "[$PACKAGE] Copying _index.md from latest version..."
                cp "$LATEST_INDEX" "$DOCS_OUT_DIR/_index.md"
            else
                echo "[$PACKAGE] Warning: No _index.md available for $VERSIONED_NAME"
            fi
        fi

        # Generate API docs using pre-built binary (output nav to per-invocation temp dir)
        echo "[$PACKAGE] Generating API docs..."
        if ! "$RESOURCEDOCSGEN" docs \
            --schemaFile "$CACHED_SCHEMA" \
            --version "v$VERSION" \
            --docsOutDir "$DOCS_OUT_DIR/api-docs" \
            --packageTreeJSONOutDir "$NAV_TEMP_DIR"; then
            echo "[$PACKAGE] Warning: Failed to generate docs for $VERSIONED_NAME, skipping"
            rm -rf "$DOCS_OUT_DIR"
            continue
        fi

        # Move nav JSON to final versioned name
        mv "$NAV_TEMP_DIR/${PACKAGE}.json" "$STATIC_DIR/navs/${VERSIONED_NAME}.json"

        # Generate versioned metadata YAML from cached schema
        # Use schema's displayName if available, otherwise fall back to latest package's title
        local SCHEMA_TITLE PACKAGE_TITLE PACKAGE_PUBLISHER PACKAGE_REPO_URL
        SCHEMA_TITLE=$(jq -r '.displayName // empty' "$CACHED_SCHEMA")
        if [[ -n "$SCHEMA_TITLE" ]]; then
            PACKAGE_TITLE="$SCHEMA_TITLE"
        else
            PACKAGE_TITLE=$(grep '^title:' "$PACKAGE_YAML" | head -1 | sed 's/^title: *//')
        fi
        PACKAGE_PUBLISHER=$(jq -r '.publisher // "Pulumi"' "$CACHED_SCHEMA")
        PACKAGE_REPO_URL=$(jq -r '.repository // empty' "$CACHED_SCHEMA")
        if [[ -z "$PACKAGE_REPO_URL" ]]; then
            PACKAGE_REPO_URL=$(grep '^repo_url:' "$PACKAGE_YAML" | head -1 | sed 's/^repo_url: *//')
        fi

        cat > "$PACKAGE_VERSIONS_DIR/${VERSIONED_NAME}.yaml" <<YAML
name: $PACKAGE
title: $PACKAGE_TITLE
version: v$VERSION
version_slug: "$VERSION_SLUG"
schema_file_url: $SCHEMA_URL
repo_url: $PACKAGE_REPO_URL
publisher: $PACKAGE_PUBLISHER
updated_on: $(date +%s)
YAML

        # Write sentinel file to enable skip-if-unchanged on subsequent builds.
        echo "$SCHEMA_URL" > "$SENTINEL"

        echo "[$PACKAGE] Done with $VERSIONED_NAME"
        GENERATED_VERSIONS+=("$VERSIONED_NAME")
    done < <(echo "$VERSIONS_JSON" | jq -c '.[]')

    echo ""
    echo "[$PACKAGE] === Cleaning up old versions ==="

    for existing in "$CONTENT_DIR/${PACKAGE}@"*; do
        [[ -e "$existing" ]] || continue
        local existing_name
        existing_name=$(basename "$existing")
        local is_current=false
        for generated in "${GENERATED_VERSIONS[@]}"; do
            if [[ "$existing_name" == "$generated" ]]; then
                is_current=true
                break
            fi
        done
        if [[ "$is_current" == "false" ]]; then
            echo "[$PACKAGE] Removing old version: $existing_name"
            rm -rf "$existing"
            rm -f "$STATIC_DIR/navs/${existing_name}.json"
            rm -f "$PACKAGE_VERSIONS_DIR/${existing_name}.yaml"
        fi
    done

    echo ""
    echo "[$PACKAGE] Done. Generated ${#GENERATED_VERSIONS[@]} older versions."
}

# Run all packages in parallel, up to MAX_PARALLEL at a time.
declare -a bg_pids=()
declare -a bg_pkgs=()

for PACKAGE in "${PACKAGES_TO_PROCESS[@]}"; do
    if ! is_blessed "$PACKAGE"; then
        continue
    fi

    process_package "$PACKAGE" &
    bg_pids+=($!)
    bg_pkgs+=("$PACKAGE")

    # Throttle to MAX_PARALLEL concurrent jobs
    while (( $(jobs -rp | wc -l) >= MAX_PARALLEL )); do
        wait -n 2>/dev/null || true
    done
done

# Wait for all remaining jobs and collect failures
declare -a failed_pkgs=()
for i in "${!bg_pids[@]}"; do
    if ! wait "${bg_pids[$i]}"; then
        failed_pkgs+=("${bg_pkgs[$i]}")
    fi
done

save_cache

echo ""
echo "========================================"
echo "Versioned docs generation complete"
echo "========================================"

if (( ${#failed_pkgs[@]} > 0 )); then
    echo "ERROR: The following packages failed: ${failed_pkgs[*]}" >&2
    exit 1
fi
