#!/bin/bash

set -o errexit -o pipefail

source ./scripts/common.sh

PORT="${PORT:-1313}"

# Just run Hugo.
hugo serve --port $PORT --buildDrafts --buildFuture --ignoreVendorPaths="github.com/pulumi/registry/**/*" | grep -v -e 'WARN .* REF_NOT_FOUND'
