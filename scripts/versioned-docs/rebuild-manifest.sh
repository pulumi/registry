#!/usr/bin/env bash
#
# rebuild-manifest.sh — regenerate a package's versions.json from what's
# actually in the archive bucket, instead of the incremental read-modify-write
# that publish-version.sh does.
#
# Why: a parallel backfill has many publish jobs racing on the same
# versions.json (last-write-wins drops entries). Letting each job publish its
# archive and then reconciling the manifest ONCE from the bucket listing
# sidesteps the race entirely.
#
# Archive entries are recovered from each snapshot's .vdocs.json metadata
# object (written by publish-version.sh). The hugo-window entries come from
# reconcile-package.sh:
#
#   ./reconcile-package.sh aiven | jq -c '.hugo_window' \
#     | xargs -0 -I{} ./rebuild-manifest.sh --package aiven --bucket B --hugo-window {}
#
# Usage:
#   rebuild-manifest.sh --package aiven --bucket pulumi-registry-versioned-testing \
#       --hugo-window '<json-array>' [--distribution-id E123ABC]
#
set -euo pipefail

PACKAGE=""; BUCKET=""; HUGO_WINDOW=""; DISTRIBUTION_ID=""
while [[ $# -gt 0 ]]; do
  case "$1" in
    --package)         PACKAGE="$2"; shift 2;;
    --bucket)          BUCKET="$2"; shift 2;;
    --hugo-window)     HUGO_WINDOW="$2"; shift 2;;
    --distribution-id) DISTRIBUTION_ID="$2"; shift 2;;
    *) echo "rebuild-manifest: unknown arg: $1" >&2; exit 2;;
  esac
done

[[ -n "$PACKAGE" && -n "$BUCKET" && -n "$HUGO_WINDOW" ]] \
  || { echo "rebuild-manifest: --package, --bucket, --hugo-window are required" >&2; exit 2; }

LIVE_ROOT="/registry/packages/${PACKAGE}/"
MANIFEST_KEY="registry/versioned/${PACKAGE}/versions.json"

# Archived slugs: the "PRE <pkg>@<slug>/" sub-prefixes under registry/packages/.
mapfile -t SLUG_DIRS < <(aws s3api list-objects-v2 --bucket "$BUCKET" \
  --prefix "registry/packages/${PACKAGE}@" --delimiter / \
  --query 'CommonPrefixes[].Prefix' --output text 2>/dev/null | tr '\t' '\n' | grep -v '^None$' || true)

echo "rebuild-manifest: package=$PACKAGE archives=${#SLUG_DIRS[@]}"

ARCHIVE_ENTRIES="[]"
for prefix in "${SLUG_DIRS[@]}"; do
  [[ -n "$prefix" ]] || continue
  meta="$(aws s3 cp "s3://${BUCKET}/${prefix}.vdocs.json" - 2>/dev/null || true)"
  if [[ -z "$meta" ]]; then
    echo "rebuild-manifest: WARNING no .vdocs.json under $prefix; skipping (republish it with --force to fix)" >&2
    continue
  fi
  ARCHIVE_ENTRIES="$(printf '%s' "$ARCHIVE_ENTRIES" | jq \
    --argjson meta "$meta" --arg path "/${prefix}" '
    . + [ { version: $meta.version, slug: $meta.slug, path: $path, date: ($meta.date // ""), source: "archive" } ]
  ')"
done

manifest="$(jq -n \
  --arg pkg "$PACKAGE" --arg liveRoot "$LIVE_ROOT" \
  --argjson hugo "$HUGO_WINDOW" --argjson archives "$ARCHIVE_ENTRIES" '
  ( $hugo | map(.slug) ) as $hugoSlugs
  | {
      package: $pkg,
      liveRoot: $liveRoot,
      versions: (
        ( $hugo + ( $archives | map(select(.slug as $s | ($hugoSlugs | index($s)) | not)) ) )
        | sort_by(.version | ltrimstr("v") | (try (split(".") | map(tonumber)) catch [.])) | reverse
      )
    }
')"

printf '%s' "$manifest" | jq -e . >/dev/null   # validate before upload
echo "rebuild-manifest: writing $(printf '%s' "$manifest" | jq '.versions | length') versions"
printf '%s' "$manifest" | jq -r '.versions[] | "  \(.version) [\(.source)]\(if .latest then "  (latest)" else "" end)"'

printf '%s' "$manifest" | aws s3 cp - "s3://${BUCKET}/${MANIFEST_KEY}" \
  --content-type "application/json" --cache-control "public, max-age=300"

if [[ -n "$DISTRIBUTION_ID" ]]; then
  aws cloudfront create-invalidation --distribution-id "$DISTRIBUTION_ID" \
    --paths "/${MANIFEST_KEY}" --query 'Invalidation.Id' --output text \
    | sed 's/^/rebuild-manifest: invalidation /'
fi

echo "rebuild-manifest: DONE — $PACKAGE"
