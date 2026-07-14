#!/usr/bin/env bash
#
# build-snapshot.sh — produce a self-contained, publishable HTML snapshot of one
# (package, version) pair's API docs.
#
# Mechanics: generate the versioned markdown/nav/schema for exactly one package
# version (via the shared lib.sh core), run a targeted Hugo build (fast: only
# the committed landing pages plus this one package's api-docs render), extract
# the versioned page tree from public/, and vendor every deploy-fingerprinted
# asset it references into _vassets/ so the snapshot never rots when per-deploy
# origin buckets are deleted:
#
#   - /registry/js/* and /registry/css/*  (content-hashed theme + Stencil
#     bundles; the JS is required — the left nav and language chooser are web
#     components)
#   - /icons/*.svg                        (fingerprinted icon sprite)
#   - /fingerprinted/*                    (package logos etc.)
#
# Stable-path assets (/fonts/*, /images/*, /logos/*) are left absolute: they are
# committed to static/ and exist at the same URL in every deploy.
#
# The snapshot is NOT yet injected with noindex/canonical/loader tags — that
# happens in publish-version.sh, on a copy, so a local --out-dir result can be
# previewed as-built.
#
# Output layout (under --out-dir):
#   snapshot/   the page tree for /registry/packages/<pkg>@<slug>/ (incl. schema.json)
#   nav.json    the left-nav tree for /registry/packages/navs/<pkg>@<slug>.json
#
# Usage:
#   build-snapshot.sh --package aiven --version v6.30.0 \
#       --schema-url https://raw.githubusercontent.com/pulumi/pulumi-aiven/v6.30.0/provider/cmd/pulumi-resource-aiven/schema.json \
#       --out-dir /tmp/aiven-snapshot
#
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "$SCRIPT_DIR/../.." && pwd)"

for cmd in jq yq curl hugo node; do
    command -v "$cmd" >/dev/null || { echo "Missing dependency: $cmd" >&2; exit 1; }
done

PACKAGE=""; VERSION=""; SCHEMA_URL=""; OUT_DIR=""
while [[ $# -gt 0 ]]; do
    case "$1" in
        --package)    PACKAGE="$2"; shift 2;;
        --version)    VERSION="$2"; shift 2;;
        --schema-url) SCHEMA_URL="$2"; shift 2;;
        --out-dir)    OUT_DIR="$2"; shift 2;;
        *) echo "build-snapshot: unknown arg: $1" >&2; exit 2;;
    esac
done

[[ -n "$PACKAGE" && -n "$VERSION" && -n "$SCHEMA_URL" && -n "$OUT_DIR" ]] \
    || { echo "build-snapshot: --package, --version, --schema-url, --out-dir are required" >&2; exit 2; }

RESOURCEDOCSGEN="${RESOURCEDOCSGEN:-$REPO_ROOT/bin/resourcedocsgen}"
if [[ ! -x "$RESOURCEDOCSGEN" ]]; then
    make -C "$REPO_ROOT" bin/resourcedocsgen
fi

# shellcheck source=scripts/versioned-docs/lib.sh
source "$SCRIPT_DIR/lib.sh"

VERSION="${VERSION#v}"
SLUG=$(version_slug "$VERSION")
VERSIONED_NAME="${PACKAGE}@${SLUG}"

CONTENT_DIR="$REPO_ROOT/themes/default/content/registry/packages"
STATIC_DIR="$REPO_ROOT/themes/default/static/registry/packages"
PACKAGE_VERSIONS_DIR="$REPO_ROOT/themes/default/data/registry/package_versions"

echo "build-snapshot: $VERSIONED_NAME (v$VERSION) from $SCHEMA_URL"

# 1. Theme assets (CSS/JS bundles + asset manifest). Skippable when the caller
#    has already built them (e.g. a workflow step).
if [[ "${SKIP_BUILD_ASSETS:-}" != "1" ]]; then
    make -C "$REPO_ROOT" build-assets
fi

# 2. Generate this one version's markdown, nav, schema, and data entry.
generate_version_content "$PACKAGE" "$VERSION" "$SCHEMA_URL" \
    "$CONTENT_DIR" "$STATIC_DIR" "$PACKAGE_VERSIONS_DIR" \
    || { echo "build-snapshot: generation failed for $VERSIONED_NAME" >&2; exit 1; }

# 3. Targeted Hugo build. ASSET_BUNDLE_ID is deliberately NOT set: without it
#    the templates reference the content-hashed theme bundles directly (see
#    partials/assets.html), which exist in public/ and are what we vendor.
#    scripts/ci/build.sh's minify-css/critical-css post-processing is skipped
#    for the same reason — it only rewrites the ASSET_BUNDLE_ID bundles.
cd "$REPO_ROOT"
node scripts/apply-fixes.js
echo "build-snapshot: running Hugo..."
# With pipefail, this still fails when hugo fails; the || true only shields
# grep's "no lines matched" status.
hugo --minify -e production 2>&1 | { grep -v -e 'WARN .* REF_NOT_FOUND' || true; }
[[ -d "public/registry/packages/$VERSIONED_NAME" ]] \
    || { echo "build-snapshot: Hugo did not produce public/registry/packages/$VERSIONED_NAME" >&2; exit 1; }

# 4. Extract the versioned page tree (HTML + index.md siblings for Accept:
#    text/markdown negotiation + schema.json from static/) and the nav JSON.
rm -rf "$OUT_DIR"
mkdir -p "$OUT_DIR"
SNAP="$OUT_DIR/snapshot"
cp -a "public/registry/packages/$VERSIONED_NAME" "$SNAP"
cp -a "public/registry/packages/navs/${VERSIONED_NAME}.json" "$OUT_DIR/nav.json"

# The .generated sentinel is build cache state, not content.
find "$SNAP" -name '.generated' -delete

# Strip the build-time (live-mode) selector loader tag: archive pages get their
# own archive-mode tag injected at publish time, and a leftover live tag would
# both confuse the loader and trip the injector's idempotency check.
find "$SNAP" -type f -name '*.html' -print0 | xargs -0 perl -0777 -i -pe \
    's{<script\b[^>]*\bdata-vdocs-pkg=[^>]*></script>}{}gi'

# 5. Vendor fingerprinted assets and rewrite references.
BASE="/registry/packages/$VERSIONED_NAME"
mkdir -p "$SNAP/_vassets"
cp -a public/registry/js "$SNAP/_vassets/js"
cp -a public/registry/css "$SNAP/_vassets/css"

# Copy every fingerprinted sprite/logo reference found in the snapshot's HTML.
# (The sprite lives under /registry/icons/, package logos under /fingerprinted/.
# The terminator set covers both quoted and minifier-unquoted attributes, and
# stops before a #fragment.)
mapfile -t FP_ASSETS < <(grep -rhoE "(/registry/icons|/icons|/fingerprinted)/[^\"'#>[:space:]]+" \
    --include='*.html' "$SNAP" | sort -u)
for asset in "${FP_ASSETS[@]}"; do
    [[ -f "public$asset" ]] || { echo "build-snapshot: WARNING referenced asset not in build output: $asset" >&2; continue; }
    mkdir -p "$SNAP/_vassets/$(dirname "${asset#/}")"
    cp -a "public$asset" "$SNAP/_vassets/${asset#/}"
done

# Rewrite the references. hugo --minify drops attribute quotes where it can,
# so the quote is optional in these patterns.
find "$SNAP" -type f -name '*.html' -print0 | xargs -0 sed -i -E \
    -e "s,(src|href)=(\"?)/registry/js/,\1=\2$BASE/_vassets/js/,g" \
    -e "s,(src|href)=(\"?)/registry/css/,\1=\2$BASE/_vassets/css/,g" \
    -e "s,(src|href)=(\"?)/registry/icons/,\1=\2$BASE/_vassets/registry/icons/,g" \
    -e "s,(src|href)=(\"?)/icons/,\1=\2$BASE/_vassets/icons/,g" \
    -e "s,(src|href)=(\"?)/fingerprinted/,\1=\2$BASE/_vassets/fingerprinted/,g"

PAGES=$(find "$SNAP" -name '*.html' | wc -l)
SIZE=$(du -sh "$SNAP" | cut -f1)
echo "build-snapshot: DONE — $PAGES pages, $SIZE, at $OUT_DIR"
echo "build-snapshot: preview locally with: python3 -m http.server -d $SNAP"
