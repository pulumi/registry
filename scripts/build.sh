#!/bin/bash

set -o errexit -o pipefail

source ./scripts/common.sh

# Apply fixes. These are things we do to patch known-malformed
node scripts/apply-fixes.js

hugo | grep -v -e 'WARN .* REF_NOT_FOUND'
