#!/bin/bash

set -o errexit -o pipefail

missing=()
for name in AWS_ACCESS_KEY_ID AWS_SECRET_ACCESS_KEY PULUMI_ACCESS_TOKEN; do
    [ -n "${!name:-}" ] || missing+=("${name}")
done
if [ ${#missing[@]} -gt 0 ]; then
    echo "Cannot build a preview: ${missing[*]} not set."
    echo "The calling workflow needs id-token: write, a pulumi/auth-actions step that feeds"
    echo "pulumi/esc-action, ESC_ACTION_ENVIRONMENT=pulumi/github-secrets/pulumi-registry,"
    echo "and an OIDC role for AWS."
    exit 1
fi

aws sts get-caller-identity

./scripts/ci/validate-packages.sh
./scripts/ci/build.sh preview
./scripts/ci/sync.sh preview
