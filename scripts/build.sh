#!/bin/bash

set -o errexit -o pipefail

source ./scripts/common.sh

# Apply fixes. See script for details.
node scripts/apply-fixes.js

hugo | grep -v -e 'WARN .* REF_NOT_FOUND'
