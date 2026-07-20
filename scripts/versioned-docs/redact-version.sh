#!/usr/bin/env bash
#
# redact-version.sh — remove a published archive snapshot from the bucket and
# manifest.
#
# An auditable, deliberate operation. Deletes every S3 object version under the
# registry/packages/<pkg>@<slug>/ prefix (bucket versioning is on) plus the
# archived nav JSON, drops the entry from versions.json, and invalidates the
# prefix + manifest in CloudFront.
#
# Requires ADMIN AWS credentials (s3:DeleteObjectVersion + s3:ListBucketVersions
# on the bucket). Redaction is a rare admin operation, so immutable archives
# can't be wiped by an automated or compromised CI run.
#
# Usage:
#   redact-version.sh --package P --slug 4.x --bucket B [--distribution-id D] [--yes]
#
set -euo pipefail

PACKAGE=""; SLUG=""; BUCKET=""; DISTRIBUTION_ID=""; ASSUME_YES="false"
while [[ $# -gt 0 ]]; do
  case "$1" in
    --package)         PACKAGE="$2"; shift 2;;
    --slug)            SLUG="$2"; shift 2;;
    --bucket)          BUCKET="$2"; shift 2;;
    --distribution-id) DISTRIBUTION_ID="$2"; shift 2;;
    --yes)             ASSUME_YES="true"; shift;;
    *) echo "redact-version: unknown arg: $1" >&2; exit 2;;
  esac
done

[[ -n "$PACKAGE" && -n "$SLUG" && -n "$BUCKET" ]] \
  || { echo "redact-version: --package, --slug, --bucket are required" >&2; exit 2; }

VERSIONED_NAME="${PACKAGE}@${SLUG}"
PREFIX="registry/packages/${VERSIONED_NAME}/"
NAV_KEY="registry/packages/navs/${VERSIONED_NAME}.json"
MANIFEST_KEY="registry/versioned/${PACKAGE}/versions.json"

echo "redact-version: package=$PACKAGE slug=$SLUG bucket=$BUCKET prefix=$PREFIX"
if [[ "$ASSUME_YES" != "true" ]]; then
  read -r -p "Permanently delete all objects under $PREFIX and remove it from the manifest? [y/N] " ans
  [[ "$ans" == "y" || "$ans" == "Y" ]] || { echo "redact-version: aborted"; exit 1; }
fi

# 1. Delete every object version + delete marker under the prefix (and the nav
#    JSON), paginating until clean.
delete_all_versions() {
  local prefix="$1"
  while :; do
    page="$(aws s3api list-object-versions --bucket "$BUCKET" --prefix "$prefix" --max-items 1000 --output json)"
    del="$(printf '%s' "$page" | jq -c '{Objects: ([ (.Versions // [])[], (.DeleteMarkers // [])[] ] | map({Key:.Key, VersionId:.VersionId})), Quiet: true}')"
    n="$(printf '%s' "$del" | jq '.Objects | length')"
    [[ "$n" -eq 0 ]] && break
    printf '%s' "$del" | aws s3api delete-objects --bucket "$BUCKET" --delete file:///dev/stdin >/dev/null
    echo "redact-version: deleted $n object version(s) under $prefix"
  done
}
delete_all_versions "$PREFIX"
delete_all_versions "$NAV_KEY"

# 2. Drop the entry from the manifest.
existing_manifest="$(aws s3 cp "s3://${BUCKET}/${MANIFEST_KEY}" - 2>/dev/null || true)"
if [[ -n "$existing_manifest" ]]; then
  updated="$(printf '%s' "$existing_manifest" | jq --arg slug "$SLUG" '
    .versions = ( (.versions // []) | map(select((.slug != $slug) or (.source != "archive"))) )
  ')"
  printf '%s' "$updated" | jq -e . >/dev/null
  printf '%s' "$updated" | aws s3 cp - "s3://${BUCKET}/${MANIFEST_KEY}" \
    --content-type "application/json" --cache-control "public, max-age=300"
  echo "redact-version: manifest updated"
else
  echo "redact-version: no manifest found; nothing to update"
fi

# 3. Invalidate the prefix + manifest.
if [[ -n "$DISTRIBUTION_ID" ]]; then
  aws cloudfront create-invalidation --distribution-id "$DISTRIBUTION_ID" \
    --paths "/${PREFIX}*" "/${NAV_KEY}" "/${MANIFEST_KEY}" --query 'Invalidation.Id' --output text \
    | sed 's/^/redact-version: invalidation /'
else
  echo "redact-version: no --distribution-id; skipping invalidation"
fi

echo "redact-version: DONE — $VERSIONED_NAME redacted"
