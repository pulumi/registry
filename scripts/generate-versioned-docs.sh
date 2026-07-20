#!/usr/bin/env bash
#
# Generates the in-Hugo versioned API docs (last MAJOR_COUNT majors, latest
# minor of each) for blessed first-party providers. Runs after `make api-docs`
# has generated the latest version of every package.
#
# Versions older than this window — and all versions of non-blessed packages —
# are served from the write-once archive instead; see
# scripts/versioned-docs/README.md.
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

# Shared per-(package, version) generation core.
# shellcheck source=scripts/versioned-docs/lib.sh
source "$SCRIPT_DIR/versioned-docs/lib.sh"

# Install registry-mirror-discover (pinned commit for stability).
# Cache the built binary to avoid recompiling on every build.
REGISTRY_MIRROR_TOOLS_COMMIT="dda1dfd85d540fab45a0d19060e963564d0b36aa"
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
# The list comes from ci-mgmt (pinned commit defined in versioned-docs/lib.sh).
mapfile -t BLESSED_PACKAGES < <(fetch_blessed_packages)
(( ${#BLESSED_PACKAGES[@]} > 0 )) || { echo "Failed to fetch providers.json from ci-mgmt" >&2; exit 1; }

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

mkdir -p "$CONTENT_DIR" "$STATIC_DIR/navs" "$PACKAGE_VERSIONS_DIR"

# Persistent cache for versioned docs output. In CI, this directory is preserved between
# runs by actions/cache. On cache hit, the sentinel-based skip logic avoids regenerating
# unchanged packages entirely. The cache stores only versioned (@-suffixed) content.
VERSIONED_DOCS_CACHE="$REPO_ROOT/.cache/versioned-docs"
mkdir -p "$VERSIONED_DOCS_CACHE/content" "$VERSIONED_DOCS_CACHE/navs" "$VERSIONED_DOCS_CACHE/metadata" "$VERSIONED_DOCS_CACHE/schemas"

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
    for cached_schema_dir in "$VERSIONED_DOCS_CACHE/schemas/"*@*; do
        [[ -d "$cached_schema_dir" ]] || continue
        local name
        name=$(basename "$cached_schema_dir")
        [[ -d "$STATIC_DIR/$name" ]] || cp -a "$cached_schema_dir" "$STATIC_DIR/$name"
    done
    if (( count > 0 )); then
        echo "Restored $count versioned doc sets from cache"
    fi
}

# Save versioned docs output back to the cache for next run.
save_cache() {
    # Clear stale cache entries and repopulate from current output.
    rm -rf "$VERSIONED_DOCS_CACHE/content/"*@* "$VERSIONED_DOCS_CACHE/navs/"*@* "$VERSIONED_DOCS_CACHE/metadata/"*@* "$VERSIONED_DOCS_CACHE/schemas/"*@*
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
    for schema_dir in "$STATIC_DIR/"*@*; do
        [[ -d "$schema_dir" ]] || continue
        cp -a "$schema_dir" "$VERSIONED_DOCS_CACHE/schemas/"
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

        local rc=0
        generate_version_content "$PACKAGE" "$VERSION" "$SCHEMA_URL" \
            "$CONTENT_DIR" "$STATIC_DIR" "$PACKAGE_VERSIONS_DIR" || rc=$?
        if (( rc != 0 )); then
            echo "[$PACKAGE] Warning: Skipping $VERSIONED_NAME (generation failed with status $rc)"
            continue
        fi

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
            rm -rf "$STATIC_DIR/${existing_name}"
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
