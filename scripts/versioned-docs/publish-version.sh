#!/usr/bin/env bash
#
# publish-version.sh — publish one immutable snapshot to the permanent archive
# bucket and update the per-package manifest.
#
# Steps (full publish):
#   1. Refuse if registry/packages/<pkg>@<slug>/ already has objects (unless --force).
#   2. Inject archive tags (noindex + canonical + loader) onto a COPY of --src.
#   3. Sync the copy to the bucket with immutable Cache-Control (1 year), plus
#      the nav JSON to registry/packages/navs/<pkg>@<slug>.json and a small
#      .vdocs.json metadata object (consumed by rebuild-manifest.sh).
#   4. Read-modify-write registry/versioned/<pkg>/versions.json: insert the
#      archive entry (deduped by slug; in-Hugo entries win on overlap), replace
#      the hugo-window entries when --hugo-window is given, semver-sort
#      newest-first, upload with Cache-Control max-age=300.
#   5. Invalidate the manifest (and, with --force, the version prefix) in CloudFront.
#
# Manifest-only mode (--manifest-only, requires --hugo-window): skips steps
# 1-3 and only refreshes the hugo-window entries — used by reconcile-package.sh
# on version bumps that don't archive anything (e.g. a minor bump within the
# same major, which still changes the latest version string).
#
# Manifest entry shape (see README.md):
#   {version: "v6.30.0", slug: "6.x", path: "/registry/packages/aiven@6.x/",
#    date: "...", source: "hugo"|"archive", latest?: true}
#
# Requires AWS credentials with write access to the bucket (the
# registry-versioned-docs-publisher role in CI) plus jq and the AWS CLI.
#
# Usage:
#   publish-version.sh --package aiven --version v6.30.0 --src ./out/snapshot \
#       --nav-file ./out/nav.json --bucket pulumi-registry-versioned-testing \
#       [--hugo-window '<json-array>'] [--distribution-id E123ABC] \
#       [--site https://www.pulumi.com] [--date 2026-07-14T00:00:00Z] [--force]
#   publish-version.sh --package aiven --manifest-only --hugo-window '<json-array>' \
#       --bucket pulumi-registry-versioned-testing [--distribution-id E123ABC]
#
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

PACKAGE=""; VERSION=""; SRC=""; NAV_FILE=""; BUCKET=""; HUGO_WINDOW=""
DISTRIBUTION_ID=""; SITE="https://www.pulumi.com"; DATE=""
FORCE="false"; MANIFEST_ONLY="false"

while [[ $# -gt 0 ]]; do
  case "$1" in
    --package)         PACKAGE="$2"; shift 2;;
    --version)         VERSION="$2"; shift 2;;
    --src)             SRC="$2"; shift 2;;
    --nav-file)        NAV_FILE="$2"; shift 2;;
    --bucket)          BUCKET="$2"; shift 2;;
    --hugo-window)     HUGO_WINDOW="$2"; shift 2;;
    --distribution-id) DISTRIBUTION_ID="$2"; shift 2;;
    --site)            SITE="$2"; shift 2;;
    --date)            DATE="$2"; shift 2;;
    --force)           FORCE="true"; shift;;
    --manifest-only)   MANIFEST_ONLY="true"; shift;;
    *) echo "publish-version: unknown arg: $1" >&2; exit 2;;
  esac
done

[[ -n "$PACKAGE" && -n "$BUCKET" ]] || { echo "publish-version: --package and --bucket are required" >&2; exit 2; }
if [[ "$MANIFEST_ONLY" == "true" ]]; then
  [[ -n "$HUGO_WINDOW" ]] || { echo "publish-version: --manifest-only requires --hugo-window" >&2; exit 2; }
else
  [[ -n "$VERSION" && -n "$SRC" && -n "$NAV_FILE" ]] \
    || { echo "publish-version: --version, --src, --nav-file are required (or use --manifest-only)" >&2; exit 2; }
  [[ -d "$SRC" ]] || { echo "publish-version: src dir not found: $SRC" >&2; exit 2; }
  [[ -f "$NAV_FILE" ]] || { echo "publish-version: nav file not found: $NAV_FILE" >&2; exit 2; }
fi

[[ -n "$DATE" ]] || DATE="$(date -u +%Y-%m-%dT%H:%M:%SZ)"
LIVE_ROOT="/registry/packages/${PACKAGE}/"
MANIFEST_KEY="registry/versioned/${PACKAGE}/versions.json"
IMMUTABLE_CC="public, max-age=31536000, immutable"
MANIFEST_CC="public, max-age=300"

SLUG=""; VERSIONED_NAME=""; PREFIX=""; VERSION_PATH=""
if [[ -n "$VERSION" ]]; then
  [[ "$VERSION" == v* ]] || VERSION="v$VERSION"
  MAJOR="${VERSION#v}"; MAJOR="${MAJOR%%.*}"
  SLUG="${MAJOR}.x"
  VERSIONED_NAME="${PACKAGE}@${SLUG}"
  PREFIX="registry/packages/${VERSIONED_NAME}/"       # S3 key prefix (mirrors the URL, no leading slash)
  VERSION_PATH="/registry/packages/${VERSIONED_NAME}/" # URL path
fi

if [[ "$MANIFEST_ONLY" != "true" ]]; then
  echo "publish-version: package=$PACKAGE version=$VERSION bucket=$BUCKET prefix=$PREFIX force=$FORCE"

  # 1. Refuse to overwrite an existing immutable snapshot unless --force.
  existing_count="$(aws s3api list-objects-v2 --bucket "$BUCKET" --prefix "$PREFIX" --max-items 1 \
    --query 'length(Contents || `[]`)' --output text 2>/dev/null || echo 0)"
  if [[ "$existing_count" != "0" && "$existing_count" != "None" ]]; then
    if [[ "$FORCE" != "true" ]]; then
      echo "publish-version: REFUSING — objects already exist under $PREFIX (pass --force to overwrite)" >&2
      exit 1
    fi
    echo "publish-version: --force set; overwriting existing $PREFIX"
  fi

  # 2. Inject archive tags onto a copy (never mutate the caller's --src).
  WORKDIR="$(mktemp -d)"
  trap 'rm -rf "$WORKDIR"' EXIT
  cp -a "$SRC/." "$WORKDIR/"

  "$SCRIPT_DIR/inject-version-switcher.sh" --package "$PACKAGE" --version "$VERSION" \
    --slug "$SLUG" --src "$WORKDIR" --site "$SITE"

  # Snapshot metadata: lets rebuild-manifest.sh recover the full version number
  # for each archived slug from the bucket alone.
  jq -n --arg pkg "$PACKAGE" --arg version "$VERSION" --arg slug "$SLUG" --arg date "$DATE" \
    '{package: $pkg, version: $version, slug: $slug, date: $date}' > "$WORKDIR/.vdocs.json"

  # 3. Sync to the bucket with immutable Cache-Control.
  sync_args=(--no-progress --cache-control "$IMMUTABLE_CC")
  [[ "$FORCE" == "true" ]] && sync_args+=(--delete)
  aws s3 sync "$WORKDIR/" "s3://${BUCKET}/${PREFIX}" "${sync_args[@]}"
  aws s3 cp "$NAV_FILE" "s3://${BUCKET}/registry/packages/navs/${VERSIONED_NAME}.json" \
    --content-type "application/json" --cache-control "$IMMUTABLE_CC" --no-progress
else
  echo "publish-version: package=$PACKAGE bucket=$BUCKET manifest-only"
fi

# 4. Manifest read-modify-write.
existing_manifest="$(aws s3 cp "s3://${BUCKET}/${MANIFEST_KEY}" - 2>/dev/null || true)"
[[ -n "$existing_manifest" ]] || existing_manifest='{}'

updated_manifest="$(printf '%s' "$existing_manifest" | jq \
  --arg pkg "$PACKAGE" --arg liveRoot "$LIVE_ROOT" \
  --arg version "${VERSION:-}" --arg slug "${SLUG:-}" --arg path "${VERSION_PATH:-}" \
  --arg date "$DATE" --argjson hugoWindow "${HUGO_WINDOW:-null}" '
  ( . // {} ) as $m
  | ( $m.versions // [] ) as $existing
  # The hugo-window entries: replaced wholesale when the caller passes fresh
  # ones (reconcile knows the current window), preserved otherwise.
  | ( if ($hugoWindow != null) then $hugoWindow
      else ($existing | map(select(.source == "hugo"))) end ) as $hugo
  | ( $hugo | map(.slug) ) as $hugoSlugs
  # Existing archive entries, minus any that the fresh hugo window now covers
  # and minus the slug being (re)published.
  | ( $existing
      | map(select(.source == "archive"))
      | map(select(.slug as $s | ($hugoSlugs | index($s)) | not))
      | map(select(.slug != $slug or $slug == "")) ) as $archives
  # The new archive entry — skipped in manifest-only mode ($slug empty) or when
  # the slug is being served in-Hugo (hugo entries always win).
  | ( if ($slug != "" and (($hugoSlugs | index($slug)) | not))
      then [ { version: $version, slug: $slug, path: $path, date: $date, source: "archive" } ]
      else [] end ) as $newEntry
  | {
      package: $pkg,
      liveRoot: $liveRoot,
      versions: (
        ( $hugo + $archives + $newEntry )
        # Order newest-first by SEMVER, not by date — a back-catalog version
        # published today must still sort below newer versions. Numeric
        # component compare so 3.10.0 > 3.9.0; odd versions fall back to the
        # raw string.
        | sort_by(.version | ltrimstr("v") | (try (split(".") | map(tonumber)) catch [.])) | reverse
      )
    }
')"

printf '%s' "$updated_manifest" | jq -e . >/dev/null  # validate JSON before upload
printf '%s' "$updated_manifest" | aws s3 cp - "s3://${BUCKET}/${MANIFEST_KEY}" \
  --content-type "application/json" --cache-control "$MANIFEST_CC"

# 5. CloudFront invalidation (belt-and-braces; the 300s manifest TTL already bounds staleness).
if [[ -n "$DISTRIBUTION_ID" ]]; then
  paths=("/${MANIFEST_KEY}")
  [[ "$FORCE" == "true" && -n "$VERSION_PATH" ]] && paths+=("${VERSION_PATH}*")
  aws cloudfront create-invalidation --distribution-id "$DISTRIBUTION_ID" \
    --paths "${paths[@]}" --query 'Invalidation.Id' --output text \
    | sed 's/^/publish-version: invalidation /'
else
  echo "publish-version: no --distribution-id; skipping invalidation (manifest TTL is 300s)"
fi

if [[ "$MANIFEST_ONLY" == "true" ]]; then
  echo "publish-version: DONE — manifest refreshed for $PACKAGE"
else
  echo "publish-version: DONE — $VERSION available at https://www.pulumi.com${VERSION_PATH}"
fi
