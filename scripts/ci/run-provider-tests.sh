#!/bin/bash

set -o errexit -o pipefail

# runs on script exit to stop hugo server (whether successful or not)
trap 'pkill -f "make serve" || true' EXIT

bin/resourcedocsgen docs --schemaFile scripts/tests/schema.json \
    --version 1.0.0 \
    --docsOutDir themes/default/content/registry/packages/test-provider/api-docs \
    --packageTreeJSONOutDir "themes/default/static/registry/packages/navs"

# The API nav fetches navs/<package directory>.json, but the tree is written under
# the schema's package name, and here the two differ. Without this the API nav
# never renders in these tests.
mv themes/default/static/registry/packages/navs/testprovider.json \
    themes/default/static/registry/packages/navs/test-provider.json

make serve &

sleep 10

./scripts/run-browser-tests.sh "http://localhost:1313" "test-provider-api-docs.cy.js"
