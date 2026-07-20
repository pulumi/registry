#!/usr/bin/env bash
#
# Shared generation core for versioned provider API docs.
#
# This library is sourced by two callers:
#   - scripts/generate-versioned-docs.sh: regenerates the in-Hugo versioned docs
#     for blessed first-party providers on every site build.
#   - scripts/versioned-docs/build-snapshot.sh: generates a single (package,
#     version) pair for a write-once archive snapshot.
#
# Callers must set (or accept the defaults for):
#   REPO_ROOT         - repository root
#   RESOURCEDOCSGEN   - path to the resourcedocsgen binary
#   SCHEMA_CACHE_DIR  - where downloaded schemas are cached (50MB+ each)

: "${REPO_ROOT:?REPO_ROOT must be set before sourcing lib.sh}"
: "${RESOURCEDOCSGEN:=$REPO_ROOT/bin/resourcedocsgen}"
: "${SCHEMA_CACHE_DIR:=$REPO_ROOT/.cache/schemas}"

# Blessed packages are first-party providers that get in-Hugo versioned docs on
# every site build. Everything else (and blessed versions older than the
# in-Hugo window) is served from the write-once archive. The list comes from
# ci-mgmt (pinned commit for stability).
: "${CI_MGMT_COMMIT:=b3cede4113152996ec67c47a2ca0cef3c5aeb626}"

# Prints the blessed package names, one per line.
fetch_blessed_packages() {
    local url="https://raw.githubusercontent.com/pulumi/ci-mgmt/${CI_MGMT_COMMIT}/provider-ci/providers.json"
    curl -sfL "$url" | jq -r '.[]'
}

# Derive the version slug used in URLs and directory names: v6.83.0 -> 6.x
version_slug() {
    local version="${1#v}"
    echo "${version%%.*}.x"
}

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
# If derivation fails or the URL returns 404, the caller falls back to copying
# the latest version's _index.md.
derive_index_url() {
    local schema_url="$1"

    if [[ "$schema_url" == *"raw.githubusercontent.com"* ]]; then
        echo "$schema_url" | sed -E 's|^(https://raw\.githubusercontent\.com/[^/]+/[^/]+/[^/]+)/.*$|\1/docs/_index.md|'
    elif [[ "$schema_url" == *"cloudfront.net"* ]]; then
        echo "$schema_url" | sed 's|/schemas/|/docs/|; s|schema\.json$|index.md|'
    else
        echo ""
    fi
}

# Download a schema to the cache (converting YAML schemas to JSON) and echo the
# cached path. Returns non-zero if the schema can't be fetched or is empty.
download_schema() {
    local pkg="$1" version="$2" schema_url="$3"
    local cached_schema="$SCHEMA_CACHE_DIR/${pkg}-v${version#v}.json"

    mkdir -p "$SCHEMA_CACHE_DIR"

    if [[ ! -f "$cached_schema" ]]; then
        echo "[$pkg] Downloading schema from $schema_url..." >&2
        if [[ "$schema_url" == *.yaml || "$schema_url" == *.yml ]]; then
            if ! curl -sfL "$schema_url" | yq -o json > "$cached_schema"; then
                rm -f "$cached_schema"
                return 1
            fi
        else
            if ! curl -sfL "$schema_url" -o "$cached_schema"; then
                rm -f "$cached_schema"
                return 1
            fi
        fi
        if [[ ! -s "$cached_schema" ]]; then
            rm -f "$cached_schema"
            return 1
        fi
    else
        echo "[$pkg] Using cached schema: $cached_schema" >&2
    fi

    echo "$cached_schema"
}

# Generate the full content set for one (package, version) pair:
#   <content_dir>/<pkg>@<slug>/            _index.md + api-docs/ markdown
#   <static_dir>/navs/<pkg>@<slug>.json    left-nav tree
#   <static_dir>/<pkg>@<slug>/schema.json  schema served by the "Schema (JSON)" link
#   <versions_dir>/<pkg>@<slug>.yaml       Hugo data for the version selector/card
#
# A .generated sentinel containing the schema URL enables skip-if-unchanged.
#
# Returns: 0 = generated (or already fresh), 2 = schema unavailable (soft skip),
#          1 = generation failed.
generate_version_content() {
    local pkg="$1" version="$2" schema_url="$3"
    local content_dir="$4" static_dir="$5" versions_dir="$6"

    version="${version#v}"
    local slug versioned_name
    slug=$(version_slug "$version")
    versioned_name="${pkg}@${slug}"

    local docs_out_dir="$content_dir/$versioned_name"
    local sentinel="$docs_out_dir/.generated"

    # Skip unchanged versions: docs were already generated from this exact
    # schema URL and every output artifact is present.
    if [[ -f "$sentinel" ]] && grep -qF "$schema_url" "$sentinel" 2>/dev/null \
        && [[ -f "$static_dir/navs/${versioned_name}.json" ]] \
        && [[ -f "$static_dir/$versioned_name/schema.json" ]] \
        && [[ -f "$versions_dir/${versioned_name}.yaml" ]]; then
        echo "[$pkg] Skipping $versioned_name (already generated, inputs unchanged)"
        return 0
    fi

    echo "[$pkg] === Processing v$version ($slug) ==="

    rm -rf "$docs_out_dir"
    mkdir -p "$docs_out_dir" "$static_dir/navs" "$static_dir/$versioned_name" "$versions_dir"

    local cached_schema
    if ! cached_schema=$(download_schema "$pkg" "$version" "$schema_url"); then
        echo "[$pkg] Warning: Failed to download schema from $schema_url" >&2
        rm -rf "$docs_out_dir" "$static_dir/$versioned_name"
        return 2
    fi

    # Fetch _index.md from derived URL, fall back to latest version's copy.
    local index_url
    index_url=$(derive_index_url "$schema_url")
    if [[ -n "$index_url" ]]; then
        echo "[$pkg] Fetching _index.md from $index_url..."
        if ! curl -sfL "$index_url" -o "$docs_out_dir/_index.md" 2>/dev/null; then
            echo "[$pkg] Could not fetch _index.md from $index_url"
            rm -f "$docs_out_dir/_index.md"
        fi
    fi

    if [[ ! -f "$docs_out_dir/_index.md" ]]; then
        local latest_index="$REPO_ROOT/themes/default/content/registry/packages/$pkg/_index.md"
        if [[ -f "$latest_index" ]]; then
            echo "[$pkg] Copying _index.md from latest version..."
            cp "$latest_index" "$docs_out_dir/_index.md"
        else
            echo "[$pkg] Warning: No _index.md available for $versioned_name"
        fi
    fi

    # Generate API docs (nav JSON goes to a temp dir so it can be renamed to
    # the versioned filename).
    local nav_temp_dir
    nav_temp_dir=$(mktemp -d)

    echo "[$pkg] Generating API docs..."
    if ! "$RESOURCEDOCSGEN" docs \
        --schemaFile "$cached_schema" \
        --version "v$version" \
        --docsOutDir "$docs_out_dir/api-docs" \
        --packageTreeJSONOutDir "$nav_temp_dir"; then
        echo "[$pkg] Warning: Failed to generate docs for $versioned_name" >&2
        rm -rf "$docs_out_dir" "$static_dir/$versioned_name" "$nav_temp_dir"
        return 1
    fi

    mv "$nav_temp_dir/${pkg}.json" "$static_dir/navs/${versioned_name}.json"
    rm -rf "$nav_temp_dir"

    # Copy the schema so the "Schema (JSON)" link on versioned pages resolves.
    # (Latest packages get theirs from resourcedocsgen's --baseSchemasOutDir.)
    cp "$cached_schema" "$static_dir/$versioned_name/schema.json"

    # Generate versioned metadata YAML from the schema, falling back to the
    # latest package YAML for fields the schema doesn't carry.
    local package_yaml="$REPO_ROOT/themes/default/data/registry/packages/${pkg}.yaml"
    local schema_title package_title package_publisher package_repo_url
    schema_title=$(jq -r '.displayName // empty' "$cached_schema")
    if [[ -n "$schema_title" ]]; then
        package_title="$schema_title"
    else
        package_title=$(grep '^title:' "$package_yaml" | head -1 | sed 's/^title: *//')
    fi
    package_publisher=$(jq -r '.publisher // "Pulumi"' "$cached_schema")
    package_repo_url=$(jq -r '.repository // empty' "$cached_schema")
    if [[ -z "$package_repo_url" ]]; then
        package_repo_url=$(grep '^repo_url:' "$package_yaml" | head -1 | sed 's/^repo_url: *//')
    fi

    cat > "$versions_dir/${versioned_name}.yaml" <<YAML
name: $pkg
title: $package_title
version: v$version
version_slug: "$slug"
schema_file_url: $schema_url
repo_url: $package_repo_url
publisher: $package_publisher
updated_on: $(date +%s)
YAML

    # Write sentinel file to enable skip-if-unchanged on subsequent builds.
    echo "$schema_url" > "$sentinel"

    echo "[$pkg] Done with $versioned_name"
    return 0
}
