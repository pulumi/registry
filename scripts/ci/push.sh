#!/bin/bash

set -o errexit -o pipefail

./scripts/ci/login.sh

./scripts/ci/build.sh
./scripts/ci/sync.sh update

./scripts/ci/run-pulumi.sh update
