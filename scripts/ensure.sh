#!/bin/bash

set -o errexit -o pipefail

source ./scripts/common.sh

# Check for Go, Hugo, Node, and Yarn.
if [[ -z "$(which go)" ]]; then
    echo "This project requires Go."
    echo "See the README for the complete list of prerequisities and "
    echo "https://golang.org/doc/install for help installing Go."
    exit 1
fi

if [[ -z "$(which hugo)" ]]; then
    echo "This project requires Hugo."
    echo "See the README for the complete list of prerequisities and "
    echo "https://gohugo.io/getting-started/quick-start for help installing Hugo."
    exit 1
fi

if [[ -z "$(which node)" ]]; then
    echo "This project requires Node.js."
    echo "See the README for the complete list of prerequisities and "
    echo "https://nodejs.org/en/download for help installing Node.js."
    exit 1
fi

if [[ -z "$(which yarn)" ]]; then
    echo "This project requires the Yarn package manager."
    echo "See the README for the complete list of prerequisities and "
    echo "https://yarnpkg.com/getting-started/install for help installing Yarn."
    exit 1
fi

# Yarn retries dropped connections and timeouts on its own, but treats an HTTP status from
# the registry as final, so a 502 on any one tarball ends the whole install.
retry_yarn() {
    local max_attempts=3 attempt
    for attempt in $(seq "$max_attempts"); do
        if yarn "$@" --network-timeout 300000; then
            return 0
        fi
        if [[ "$attempt" -lt "$max_attempts" ]]; then
            sleep $((attempt * 10))
        fi
    done
    echo "yarn $* failed after $max_attempts attempts."
    return 1
}

echo "Installing Node.js modules..."
retry_yarn install
retry_yarn --cwd infrastructure install
retry_yarn --cwd ./themes/default/theme install
retry_yarn --cwd ./themes/default/theme/stencil install

