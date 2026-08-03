#!/usr/bin/env python3
"""Generate `-on-dark.svg` variants for package logos that vanish on a dark page.

The registry ships ~135 local package marks under
`themes/default/assets/fingerprinted/logos/pkg/`. Many are black or near-black
wordmarks that become invisible once the page background flips to service-black.
This script produces a sibling `<name>-on-dark.svg` for each affected mark by
lightening *only* the dark paints and leaving every colored paint alone — so, for
example, the AWS mark keeps its orange smile and gets a light wordmark, which is
how AWS's own dark-mode lockup is drawn.

Rules:

  * A paint counts as "dark" when its relative luminance is below DARK_THRESHOLD.
  * A dark paint is remapped by inverting HSL lightness (clamped into a legible
    band) and halving saturation, so a near-black navy becomes a light gray with a
    faint blue cast rather than a garish pastel.
  * SVGs whose shapes carry no `fill` at all render black by inheritance. Those
    get a light `fill` on the root <svg> element, which every unpainted descendant
    inherits. Roots that already declare a fill are left alone.
  * Marks with no dark paints (full-color logos, light logos) are skipped — they
    read fine as-is and a variant would only be one more file to keep in sync.

`layouts/partials/registry/package/icon.html` picks the variant up automatically;
there is nothing to register. Re-run after adding or replacing a local logo:

    python3 scripts/generate-dark-logos.py            # write variants
    python3 scripts/generate-dark-logos.py --check    # CI-style, no writes
"""

import argparse
import colorsys
import os
import re
import sys

LOGO_DIR = os.path.join(
    os.path.dirname(os.path.dirname(os.path.abspath(__file__))),
    "themes", "default", "assets", "fingerprinted", "logos", "pkg",
)

SUFFIX = "-on-dark.svg"

# What counts as "too dark for the dark page", in HSL. Relative luminance is the
# wrong tool here: blue contributes so little to it that a perfectly legible
# saturated brand blue (Kubernetes #326de6, luminance 0.17) scores darker than a
# mid gray. Lightness plus saturation separates the two cases properly.
#
#   * lightness < VERY_DARK           → always lighten (blacks, near-blacks)
#   * lightness < DARKISH and low sat → lighten (navy/charcoal wordmarks)
#   * anything else                   → leave alone (saturated brand colors)
VERY_DARK = 0.22
DARKISH = 0.35
LOW_SATURATION = 0.25

# The band a remapped paint's lightness is clamped into. The floor keeps a very
# dark source from landing mid-gray; the ceiling stops pure black from becoming
# pure white, which glares against service-black.
LIGHT_MIN, LIGHT_MAX = 0.85, 0.96

HEX_RE = re.compile(r"#(?:[0-9a-fA-F]{8}|[0-9a-fA-F]{6}|[0-9a-fA-F]{3,4})\b")

# Paint keywords worth handling; anything else (currentColor, none, url(#…),
# inherit) either follows the theme already or isn't a color we can reason about.
KEYWORDS = {"black": "#000000"}

PAINTABLE = ("path", "circle", "rect", "polygon", "polyline", "ellipse", "line", "text")


def expand(hex_str):
    """Normalize a #rgb / #rgba / #rrggbb / #rrggbbaa token to (rrggbb, alpha_suffix)."""
    body = hex_str.lstrip("#")
    if len(body) in (3, 4):
        rgb = "".join(c * 2 for c in body[:3])
        alpha = body[3] * 2 if len(body) == 4 else ""
    else:
        rgb = body[:6]
        alpha = body[6:]
    return rgb.lower(), alpha.lower()


def hls(rgb):
    r, g, b = (int(rgb[i:i + 2], 16) / 255 for i in (0, 2, 4))
    return colorsys.rgb_to_hls(r, g, b)


def is_too_dark(rgb):
    _h, l, s = hls(rgb)
    return l < VERY_DARK or (l < DARKISH and s < LOW_SATURATION)


def lighten(rgb):
    """Invert lightness and soften saturation, preserving hue."""
    h, l, s = hls(rgb)
    new_l = min(LIGHT_MAX, max(LIGHT_MIN, 1.0 - l))
    new_s = s * 0.5
    nr, ng, nb = colorsys.hls_to_rgb(h, new_l, new_s)
    return "%02x%02x%02x" % tuple(round(c * 255) for c in (nr, ng, nb))


def collect_paints(svg):
    """Every color token in the file, as {original_token: (rgb, alpha_suffix)}."""
    found = {}
    for match in HEX_RE.finditer(svg):
        token = match.group(0)
        found[token] = expand(token)
    for keyword, hex_str in KEYWORDS.items():
        if re.search(r'(fill|stroke|stop-color)\s*[=:]\s*["\']?%s\b' % keyword, svg):
            found[keyword] = expand(hex_str)
    return found


def has_unpainted_shapes(svg):
    """True when some shape relies on the inherited (black) default fill.

    Only a heuristic — it ignores CSS class-based fills — but the classes it can't
    see are picked up by the hex scan instead, so the two together cover the file.
    """
    for match in re.finditer(r"<(%s)\b([^>]*)>" % "|".join(PAINTABLE), svg):
        attrs = match.group(2)
        if "fill" not in attrs and "class" not in attrs and "style" not in attrs:
            return True
    return False


def root_declares_fill(svg):
    match = re.search(r"<svg\b[^>]*>", svg)
    return bool(match) and "fill" in match.group(0)


def add_root_fill(svg, color):
    """Set `fill` on the root <svg>; unpainted descendants inherit it."""
    def repl(match):
        tag = match.group(0)
        return tag[:4] + ' fill="%s"' % color + tag[4:]

    return re.sub(r"<svg\b[^>]*>", repl, svg, count=1)


def convert(svg):
    """Return the dark-variant source, or None when the mark needs no variant."""
    paints = collect_paints(svg)
    replacements = {
        token: lighten(rgb)
        for token, (rgb, _alpha) in paints.items()
        if is_too_dark(rgb)
    }

    unpainted = has_unpainted_shapes(svg) and not root_declares_fill(svg)
    if not replacements and not unpainted:
        return None

    out = svg
    # Longest tokens first so #abcdef isn't clipped by a #abc rule.
    for token in sorted(replacements, key=len, reverse=True):
        rgb, alpha = paints[token]
        new_token = "#" + replacements[token] + alpha
        if token in KEYWORDS:
            out = re.sub(r"\b%s\b" % token, new_token, out)
        else:
            out = re.sub(re.escape(token) + r"\b", new_token, out)

    if unpainted:
        out = add_root_fill(out, "#" + lighten("000000"))

    return out


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument(
        "--check",
        action="store_true",
        help="report what would change and exit non-zero if anything is stale",
    )
    args = parser.parse_args()

    sources = sorted(
        f for f in os.listdir(LOGO_DIR)
        if f.endswith(".svg") and not f.endswith(SUFFIX)
    )

    written, removed, skipped, stale = [], [], [], []

    for name in sources:
        base = name[:-len(".svg")]
        src_path = os.path.join(LOGO_DIR, name)
        dst_path = os.path.join(LOGO_DIR, base + SUFFIX)

        with open(src_path, encoding="utf-8") as fh:
            svg = fh.read()

        variant = convert(svg)

        if variant is None:
            skipped.append(name)
            if os.path.exists(dst_path):
                if args.check:
                    stale.append(base + SUFFIX + " (no longer needed)")
                else:
                    os.remove(dst_path)
                    removed.append(base + SUFFIX)
            continue

        current = None
        if os.path.exists(dst_path):
            with open(dst_path, encoding="utf-8") as fh:
                current = fh.read()

        if current == variant:
            continue

        if args.check:
            stale.append(base + SUFFIX)
        else:
            with open(dst_path, "w", encoding="utf-8") as fh:
                fh.write(variant)
            written.append(base + SUFFIX)

    if args.check:
        if stale:
            print("Out of date (%d):" % len(stale))
            for name in stale:
                print("  " + name)
            print("\nRun: python3 scripts/generate-dark-logos.py")
            return 1
        print("All dark logo variants are up to date (%d sources)." % len(sources))
        return 0

    print("sources:  %d" % len(sources))
    print("written:  %d" % len(written))
    print("removed:  %d" % len(removed))
    print("skipped:  %d (no dark paints)" % len(skipped))
    return 0


if __name__ == "__main__":
    sys.exit(main())
