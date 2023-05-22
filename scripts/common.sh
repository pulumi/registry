#!/bin/bash

BUILD_IDENTIFIER="$(git rev-parse HEAD)"
export REPO_THEME_PATH="themes/default/"


list_pull_requests() {
    curl -L   -H "Accept: application/vnd.github+json" \
    -H "Authorization: Bearer ${GITHUB_TOKEN}" \
    -H "X-GitHub-Api-Version: 2022-11-28"   https://api.github.com/repos/pulumi/registry/pulls \
    | jq -c '[.[] | .title]'
}
