#!/usr/bin/env bash
#
# discover-versions.sh — enumerate the historical versions (and version-pinned
# schema URLs) of a registry package.
#
# Routing:
#   - Blessed first-party providers use the private registry-mirror-discover
#     tool (authoritative for bridged/CDN-hosted schemas), same as
#     scripts/generate-versioned-docs.sh. If it can't be built (e.g. no access
#     to the private repo), discovery degrades to the public path with a
#     warning.
#   - Everything else uses `resourcedocsgen versions`, which works from GitHub
#     tags + the package's schema_file_url, falling back to the Pulumi Registry
#     API. Unresolvable versions are reported in `skipped` (best-effort).
#
# Output (stdout): {"package": ..., "versions": [{version, schema_url}], "skipped": [{version, reason}]}
# Versions are newest-first and carry no leading "v".
#
# Usage: discover-versions.sh <package> [--major-count N] [--minor-count N]
#
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "$SCRIPT_DIR/../.." && pwd)"

for cmd in jq curl; do
    command -v "$cmd" >/dev/null || { echo "Missing dependency: $cmd" >&2; exit 1; }
done

# shellcheck source=scripts/versioned-docs/lib.sh
source "$SCRIPT_DIR/lib.sh"

PACKAGE="${1:-}"
[[ -n "$PACKAGE" ]] || { echo "usage: discover-versions.sh <package> [--major-count N] [--minor-count N]" >&2; exit 2; }
shift

MAJOR_COUNT=3
MINOR_COUNT=1
while [[ $# -gt 0 ]]; do
    case "$1" in
        --major-count) MAJOR_COUNT="$2"; shift 2;;
        --minor-count) MINOR_COUNT="$2"; shift 2;;
        *) echo "discover-versions: unknown arg: $1" >&2; exit 2;;
    esac
done

if [[ ! -x "$RESOURCEDOCSGEN" ]]; then
    make -C "$REPO_ROOT" bin/resourcedocsgen >&2
fi

is_blessed() {
    fetch_blessed_packages | grep -qx "$1"
}

# Same pinned build as scripts/generate-versioned-docs.sh. Prints the binary
# path on success; fails silently (caller falls back) when the private repo is
# unreachable.
ensure_discover_bin() {
    local commit="dda1dfd85d540fab45a0d19060e963564d0b36aa"
    local bin="$REPO_ROOT/bin/registry-mirror-discover"
    if [[ ! -x "$bin" ]]; then
        GOBIN="$REPO_ROOT/bin" GOPRIVATE=github.com/pulumi \
            go install "github.com/pulumi/registry-mirror-tools/cmd/registry-mirror-discover@${commit}" >&2 || return 1
    fi
    echo "$bin"
}

if [[ "${FORCE_PUBLIC_DISCOVERY:-}" != "1" ]] && is_blessed "$PACKAGE"; then
    if DISCOVER_BIN=$(ensure_discover_bin); then
        VERSIONS_JSON=$("$DISCOVER_BIN" "$REPO_ROOT" \
            --packages "$PACKAGE" \
            --major-count "$MAJOR_COUNT" \
            --minor-count "$MINOR_COUNT" \
            --format json-array < /dev/null)
        jq -n --arg pkg "$PACKAGE" --argjson versions "$VERSIONS_JSON" \
            '{package: $pkg, versions: $versions, skipped: []}'
        exit 0
    fi
    echo "discover-versions: registry-mirror-discover unavailable; falling back to public discovery for blessed package $PACKAGE" >&2
fi

"$RESOURCEDOCSGEN" versions \
    --registryDir "$REPO_ROOT" \
    --package "$PACKAGE" \
    --major-count "$MAJOR_COUNT" \
    --minor-count "$MINOR_COUNT" \
    --format json
