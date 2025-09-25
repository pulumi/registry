#!/bin/bash

set -euo pipefail

cd ./tools/resourcedocsgen && go test ./...
