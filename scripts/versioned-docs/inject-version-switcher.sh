#!/usr/bin/env bash
#
# inject-version-switcher.sh — insert archive tags into a snapshot's HTML before
# it is published to the permanent archive bucket:
#
#   1. <meta name="robots" content="noindex">   (archives never compete in search)
#   2. <link rel="canonical" href="<live-equivalent page>">  (SEO flows to latest)
#   3. the evergreen selector loader <script src="/registry/js/versioned-docs.js">
#      (permanent contract URL served by the live site, so archive chrome stays
#      current without republishing)
#
# Any canonical/robots tags emitted by the original Hugo build are stripped
# first — head.html always emits a canonical link, and leaving it would collide
# with the archive's own.
#
# Live pages don't use this script: partials/registry/package/package-card-top-of-page.html
# emits the loader tag at build time.
#
# Idempotent: files already carrying `data-vdocs-pkg` are skipped.
#
# Usage:
#   inject-version-switcher.sh --package aiven --version v6.30.0 --slug 6.x \
#       --src ./snapshot [--site https://www.pulumi.com]
#
set -euo pipefail

PACKAGE=""; VERSION=""; SLUG=""; SRC=""; SITE="https://www.pulumi.com"
while [[ $# -gt 0 ]]; do
  case "$1" in
    --package) PACKAGE="$2"; shift 2;;
    --version) VERSION="$2"; shift 2;;
    --slug)    SLUG="$2"; shift 2;;
    --src)     SRC="$2"; shift 2;;
    --site)    SITE="$2"; shift 2;;
    *) echo "inject-version-switcher: unknown arg: $1" >&2; exit 2;;
  esac
done

[[ -n "$PACKAGE" && -n "$VERSION" && -n "$SLUG" && -n "$SRC" ]] \
  || { echo "inject-version-switcher: --package, --version, --slug, --src are required" >&2; exit 2;}
[[ -d "$SRC" ]] || { echo "inject-version-switcher: src dir not found: $SRC" >&2; exit 2; }
[[ "$VERSION" == v* ]] || VERSION="v$VERSION"
SITE="${SITE%/}"

LIVE_ROOT="/registry/packages/${PACKAGE}/"

processed=0; skipped=0
while IFS= read -r -d '' f; do
  # Idempotency marker: specifically the archive-mode tag. (The build-time
  # live-mode tag is stripped by build-snapshot.sh, but be defensive — a
  # leftover live tag must not prevent the noindex/canonical injection.)
  if grep -q 'data-vdocs-mode="archive"' "$f"; then
    skipped=$((skipped + 1)); continue
  fi

  # rel = page path relative to the snapshot root. Directory-index pages
  # (foo/index.html) canonicalize to the directory URL.
  rel="${f#"$SRC"/}"
  page_rel="$rel"
  case "$page_rel" in
    index.html)    page_rel="" ;;
    */index.html)  page_rel="${page_rel%index.html}" ;;
  esac

  loader="<script async src=\"/registry/js/versioned-docs.js\" data-vdocs-pkg=\"$PACKAGE\" data-vdocs-mode=\"archive\" data-vdocs-version=\"$VERSION\" data-vdocs-slug=\"$SLUG\" data-vdocs-live-root=\"$LIVE_ROOT\" data-vdocs-rel=\"$page_rel\"></script>"
  canonical="${SITE}${LIVE_ROOT}${page_rel}"
  snippet="<meta name=\"robots\" content=\"noindex\"><link rel=\"canonical\" href=\"$canonical\">$loader"

  # Strip pre-existing canonical/robots tags, then insert the snippet before
  # the first </head> (falling back to </body>). Slurp mode (-0777) is robust
  # to minified single-line heads; the quotes are optional because
  # hugo --minify drops them where it can. \x22/\x27 are " and '.
  SNIPPET="$snippet" perl -0777 -i -pe '
    s{<link\b[^>]*\brel=[\x22\x27]?canonical\b[^>]*>}{}gi;
    s{<meta\b[^>]*\bname=[\x22\x27]?robots\b[^>]*>}{}gi;
    my $s = $ENV{SNIPPET};
    s{</head>}{$s</head>}i or s{</body>}{$s</body>}i;
  ' "$f"
  processed=$((processed + 1))
done < <(find "$SRC" -type f -name '*.html' -print0)

echo "inject-version-switcher: package=$PACKAGE version=$VERSION processed=$processed skipped=$skipped"
