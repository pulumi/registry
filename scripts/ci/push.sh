#!/bin/bash

set -o errexit -o pipefail

./scripts/ci/login.sh

./scripts/ci/build.sh update
./scripts/ci/sync.sh update
./scripts/generate-search-index.sh
node ./scripts/await-in-progress.js
./scripts/ci/run-pulumi.sh update
./scripts/ci/make-s3-redirects.sh
