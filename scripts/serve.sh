#!/bin/bash

set -o errexit -o pipefail

source ./scripts/common.sh

# Run an initial asset build if one hasn't been already.
pushd "${LOCAL_THEME_PATH}"
    if ! test -f "${CSS_BUNDLE}" || ! test -f "${JS_BUNDLE}"; then
        echo "Running an initial build..."
        yarn run ensure
        yarn run build
    fi
popd

# Just run Hugo.
hugo serve --buildDrafts --buildFuture --ignoreVendorPaths="github.com/pulumi/pulumi-hugo-internal/**/*" | grep -v -e 'WARN .* REF_NOT_FOUND'
