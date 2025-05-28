#!/bin/bash

set -o errexit -o pipefail

# runs on script exit to stop hugo server (whether successful or not)
trap 'pkill -f "make serve" || true' EXIT

bin/resourcedocsgen docs --schemaFile scripts/tests/schema.json \
    --version 1.0.0 \
    --docsOutDir themes/default/content/registry/packages/test-provider/api-docs \
    --packageTreeJSONOutDir "themes/default/static/registry/packages/navs"

make serve &

sleep 10

./scripts/run-browser-tests.sh "http://localhost:1313" "test-provider-api-docs.cy.js"
