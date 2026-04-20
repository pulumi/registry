#!/bin/bash

# Invalidate cached registry pages on the docs CloudFront distribution.
#
# Each registry deploy swaps the S3 origin bucket. The docs CloudFront
# caches registry HTML for up to 30 minutes, so stale HTML can reference
# CSS/JS bundles from the previous bucket that no longer exist on the new
# one — causing broken styles until the cache expires. Invalidating
# /registry/* after the origin swap forces CloudFront to fetch fresh HTML
# whose <link> tags match the new bucket's CSS bundles.
#
# Requires PULUMI_DOCS_STACK_NAME to be set (e.g. "pulumi/www.pulumi.com/www-production").

set -o errexit -o pipefail

source ./scripts/ci/common.sh

if [[ -z "${PULUMI_DOCS_STACK_NAME:-}" ]]; then
    log "PULUMI_DOCS_STACK_NAME not set, skipping docs CDN invalidation."
    exit 0
fi

DOCS_CF_DIST_ID=$(pulumi stack output cloudFrontDistributionId --stack "$PULUMI_DOCS_STACK_NAME" 2>/dev/null || true)

if [[ -z "$DOCS_CF_DIST_ID" ]]; then
    log "Could not read cloudFrontDistributionId from $PULUMI_DOCS_STACK_NAME, skipping invalidation."
    exit 0
fi

log "Invalidating /registry/* on docs CloudFront distribution $DOCS_CF_DIST_ID..."
aws cloudfront create-invalidation \
    --distribution-id "$DOCS_CF_DIST_ID" \
    --paths "/registry/*" \
    --region us-east-1 > /dev/null

log "Invalidation created."
