#!/usr/bin/env python3
"""Decide which external package logos need a light chip behind them in dark mode.

About 90 packages point `logo_url` at an image on a host we don't control, so —
unlike the local marks in `themes/default/assets/fingerprinted/logos/pkg/`, which
get generated `-on-dark.svg` variants (scripts/generate-dark-logos.py) — we can't
recolor them. The dark page needs a per-logo decision instead:

  * Ink that has too little contrast with the dark page — a black wordmark, but
    equally a deep brand navy — disappears, and the mark needs a near-white chip
    behind it.
  * A light mark (white text on transparent) or a bright full-color one already
    reads fine, and a chip would either hide it outright or add a loud white box
    for no reason.

This script downloads every external logo, measures how much of it is lost against
the dark background (by pixel, or for SVG by declared paint), and writes the marks
that fail to

    themes/default/data/registry/external_logo_treatment.yaml

which `layouts/partials/registry/package/icon.html` reads to decide whether to emit
the `data-logo-chip` attribute the chip rule keys off.

The output is a snapshot: if a vendor swaps their logo the classification can go
stale, which is what `--check` is for. Re-run after adding packages with a
`logo_url`, and periodically:

    python3 scripts/classify-external-logos.py            # refresh the data file
    python3 scripts/classify-external-logos.py --check    # fail if it's stale

Non-PNG rasters are normalized with `sips`, which is macOS-only. Anything that can't
be decoded — including every non-PNG raster on a Linux box — is reported and left
out of the list, i.e. it gets no chip, the safer default. `--check` refuses to
compare at all when anything was undecidable (exit 2), so a machine without `sips`,
or one flaky download, reports "unknown" rather than falsely calling the committed
file stale.
"""

import argparse
import glob
import os
import re
import struct
import subprocess
import sys
import tempfile
import urllib.parse
import zlib

ROOT = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
PACKAGE_DIR = os.path.join(ROOT, "themes", "default", "data", "registry", "packages")
OUT_PATH = os.path.join(
    ROOT, "themes", "default", "data", "registry", "external_logo_treatment.yaml"
)

# The question isn't "is this logo dark?" but "would this ink disappear against the
# dark page?", so each pixel is scored by its actual contrast with the dark-mode
# page background rather than by hue. Bucketing on lightness/saturation gets it wrong
# in both directions — a deep brand navy (1Password's badge, Airbyte's wordmark) is
# highly saturated and reads as "a color", but it's still invisible on near-black.
DARK_PAGE_BG = "#1f1b21"  # --docs-bg in dark mode (service-black)

# WCAG contrast ratio below which ink is treated as lost against the background.
# 2.5 is deliberately lenient: at that point the ink is legible only as a smudge.
INVISIBLE_CONTRAST_MAX = 2.5

# Share of the mark that has to be lost before it earns a chip. Under half, because
# a logo is usually a mark plus a wordmark and losing either one is enough. Note
# that anything invisible on the dark page is by definition visible on the chip, so
# the failure mode of over-chipping is cosmetic while under-chipping hides the logo.
INVISIBLE_SHARE_MIN = 0.4

ALPHA_MIN = 32          # ignore near-transparent pixels
SAMPLE_TARGET = 20000   # cap on pixels inspected per image

HEX_RE = re.compile(r"#(?:[0-9a-fA-F]{8}|[0-9a-fA-F]{6}|[0-9a-fA-F]{3,4})\b")
LOGO_URL_RE = re.compile(r"^logo_url:\s*(.+)$", re.M)


# --- PNG decoding ----------------------------------------------------------
# Just enough of the spec to sample pixels: no interlacing, 8- or 16-bit depth.
# Pillow isn't a dependency of this repo and this keeps the script standalone.

def _read_png(path):
    with open(path, "rb") as handle:
        data = handle.read()
    if data[:8] != b"\x89PNG\r\n\x1a\n":
        raise ValueError("not a PNG")
    pos, idat, plte, trns = 8, b"", None, None
    width = height = depth = ctype = None
    while pos < len(data):
        length, kind = struct.unpack(">I4s", data[pos:pos + 8])
        pos += 8
        chunk = data[pos:pos + length]
        pos += length + 4
        if kind == b"IHDR":
            width, height, depth, ctype, _comp, _filt, interlace = struct.unpack(
                ">IIBBBBB", chunk
            )
            if interlace:
                raise ValueError("interlaced PNG")
        elif kind == b"IDAT":
            idat += chunk
        elif kind == b"PLTE":
            plte = chunk
        elif kind == b"tRNS":
            trns = chunk
        elif kind == b"IEND":
            break
    if depth not in (8, 16):
        raise ValueError("unsupported bit depth %s" % depth)

    samples = {0: 1, 2: 3, 3: 1, 4: 2, 6: 4}[ctype]
    stride_bytes = width * samples * (depth // 8)
    step = depth // 8

    raw = zlib.decompress(idat)
    out = bytearray()
    prev = bytearray(stride_bytes)
    pos = 0
    for _ in range(height):
        filt = raw[pos]
        pos += 1
        line = bytearray(raw[pos:pos + stride_bytes])
        pos += stride_bytes
        bpp = samples * step
        if filt == 1:
            for i in range(bpp, stride_bytes):
                line[i] = (line[i] + line[i - bpp]) & 255
        elif filt == 2:
            for i in range(stride_bytes):
                line[i] = (line[i] + prev[i]) & 255
        elif filt == 3:
            for i in range(stride_bytes):
                left = line[i - bpp] if i >= bpp else 0
                line[i] = (line[i] + ((left + prev[i]) >> 1)) & 255
        elif filt == 4:
            for i in range(stride_bytes):
                a = line[i - bpp] if i >= bpp else 0
                b = prev[i]
                c = prev[i - bpp] if i >= bpp else 0
                p = a + b - c
                pa, pb, pc = abs(p - a), abs(p - b), abs(p - c)
                pred = a if (pa <= pb and pa <= pc) else (b if pb <= pc else c)
                line[i] = (line[i] + pred) & 255
        out += line
        prev = line
    return width, height, ctype, samples, step, bytes(out), plte, trns


def _png_pixels(path):
    width, height, ctype, samples, step, buf, plte, trns = _read_png(path)
    total = width * height
    stride = max(1, total // SAMPLE_TARGET)

    def sample(i, channel):
        # 16-bit samples are truncated to their high byte — precision we don't need.
        return buf[(i * samples + channel) * step]

    for i in range(0, total, stride):
        if ctype == 6:
            yield sample(i, 0), sample(i, 1), sample(i, 2), sample(i, 3)
        elif ctype == 2:
            yield sample(i, 0), sample(i, 1), sample(i, 2), 255
        elif ctype == 0:
            v = sample(i, 0)
            yield v, v, v, 255
        elif ctype == 4:
            v = sample(i, 0)
            yield v, v, v, sample(i, 1)
        elif ctype == 3:
            idx = sample(i, 0)
            r, g, b = plte[idx * 3], plte[idx * 3 + 1], plte[idx * 3 + 2]
            alpha = trns[idx] if trns and idx < len(trns) else 255
            yield r, g, b, alpha


def _relative_luminance(r, g, b):
    def channel(value):
        value /= 255
        return value / 12.92 if value <= 0.03928 else ((value + 0.055) / 1.055) ** 2.4

    return 0.2126 * channel(r) + 0.7152 * channel(g) + 0.0722 * channel(b)


_BG_LUMINANCE = _relative_luminance(
    *(int(DARK_PAGE_BG.lstrip("#")[i:i + 2], 16) for i in (0, 2, 4))
)


def _is_lost_on_dark(r, g, b):
    """True when this ink has too little contrast with the dark page to be seen."""
    lighter, darker = sorted(
        (_relative_luminance(r, g, b), _BG_LUMINANCE), reverse=True
    )
    return (lighter + 0.05) / (darker + 0.05) < INVISIBLE_CONTRAST_MAX


def _share(lost, total):
    return None if not total else {"invisible": lost / total}


def raster_shares(path):
    lost = total = 0
    for r, g, b, a in _png_pixels(path):
        if a < ALPHA_MIN:
            continue
        total += 1
        lost += _is_lost_on_dark(r, g, b)
    return _share(lost, total)


def svg_shares(path):
    """SVGs are judged by their declared paints rather than by rasterizing.

    Each distinct paint counts once, so this measures "how much of the palette is
    lost" rather than "how much of the area" — close enough for flat vector marks,
    which is what these are.
    """
    with open(path, encoding="utf-8", errors="replace") as handle:
        source = handle.read()
    colors = {m.group(0).lower() for m in HEX_RE.finditer(source)}
    if re.search(r'(fill|stroke|stop-color)\s*[=:]\s*["\']?black\b', source):
        colors.add("#000000")

    lost = total = 0
    for token in colors:
        body = token.lstrip("#")
        if len(body) in (3, 4):
            body = "".join(c * 2 for c in body[:3])
        body = body[:6]
        if len(body) != 6:
            continue
        r, g, b = (int(body[i:i + 2], 16) for i in (0, 2, 4))
        total += 1
        lost += _is_lost_on_dark(r, g, b)

    # Shapes with no fill inherit the initial black, which is always lost.
    unpainted = re.search(
        r"<(path|circle|rect|polygon|ellipse)\b(?![^>]*fill)", source
    ) and not re.search(r"<svg\b[^>]*fill", source)
    if unpainted:
        total += 1
        lost += 1
    return _share(lost, total)


def needs_chip(shares):
    return shares["invisible"] >= INVISIBLE_SHARE_MIN


# --- driver ----------------------------------------------------------------

def external_logo_urls():
    urls = {}
    for path in sorted(glob.glob(os.path.join(PACKAGE_DIR, "*.yaml"))):
        with open(path, encoding="utf-8") as handle:
            match = LOGO_URL_RE.search(handle.read())
        if not match:
            continue
        url = match.group(1).strip().strip('"').strip("'")
        if url:
            urls[os.path.basename(path)[:-len(".yaml")]] = url
    return urls


def fetch(url, dest):
    # --fail matters: without it a 404 writes the host's HTML error page to `dest`
    # and reports success, and the classifier then measures the error page.
    result = subprocess.run(
        ["curl", "-sL", "--fail", "--max-time", "30", "-o", dest, url],
        capture_output=True,
    )
    return result.returncode == 0 and os.path.exists(dest) and os.path.getsize(dest) > 0


def is_svg(path):
    r"""True for an SVG, whether or not the URL carried a .svg extension.

    Sniffing the first tag rather than the first five bytes: real files start
    `<svg\n  xmlns=…>`, `<!DOCTYPE svg …>`, or with a UTF-8 BOM, and any of those
    would otherwise fall through to the raster path and fail to decode.
    """
    if path.lower().endswith(".svg"):
        return True
    with open(path, "rb") as handle:
        head = handle.read(400).lstrip(b"\xef\xbb\xbf")
    return bool(re.match(rb"\s*(<\?xml|<!DOCTYPE\s+svg|<svg\b)", head))


def to_png(path, workdir, name):
    """Return a decodable PNG path, converting with sips when needed."""
    with open(path, "rb") as handle:
        head = handle.read(8)
    if head.startswith(b"\x89PNG\r\n\x1a\n"):
        return path
    converted = os.path.join(workdir, name + ".converted.png")
    try:
        result = subprocess.run(
            ["sips", "-s", "format", "png", path, "--out", converted],
            capture_output=True,
        )
    except FileNotFoundError:
        raise ValueError(
            "cannot decode: `sips` is macOS-only and isn't installed"
        ) from None
    if result.returncode != 0 or not os.path.exists(converted):
        raise ValueError("could not convert to PNG")
    return converted


def render(chips):
    lines = [
        "# Generated by scripts/classify-external-logos.py — do not edit by hand.",
        "#",
        "# Packages whose external `logo_url` points at a dark, near-monochrome mark",
        "# that would disappear against the dark-mode page background. The registry",
        "# can't recolor a third-party image, so icon.html renders these on a light",
        "# chip instead. Every other external logo (light or full-color) reads fine",
        "# in both modes and is deliberately absent from this list.",
        "#",
        "# Re-run the script after adding a package with a `logo_url`, or when a",
        "# vendor changes their logo.",
        "chip:",
    ]
    lines += ["  - %s" % name for name in chips]
    return "\n".join(lines) + "\n"


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument(
        "--check",
        action="store_true",
        help="compare against the committed file and exit non-zero if it's stale",
    )
    parser.add_argument(
        "--verbose", action="store_true", help="print the per-logo measurements"
    )
    args = parser.parse_args()

    urls = external_logo_urls()
    chips, skipped = [], []

    with tempfile.TemporaryDirectory() as workdir:
        for name, url in sorted(urls.items()):
            suffix = os.path.splitext(urllib.parse.urlparse(url).path)[1].lower()
            raw = os.path.join(workdir, name + (suffix or ".bin"))
            if not fetch(url, raw):
                skipped.append((name, "download failed"))
                continue
            try:
                if is_svg(raw):
                    shares = svg_shares(raw)
                else:
                    shares = raster_shares(to_png(raw, workdir, name))
            except Exception as exc:  # noqa: BLE001 — report and move on
                skipped.append((name, str(exc)[:60]))
                continue
            if shares is None:
                skipped.append((name, "no opaque pixels"))
                continue
            if args.verbose:
                print("  %-24s invisible-on-dark=%.2f  %s" % (
                    name, shares["invisible"],
                    "CHIP" if needs_chip(shares) else ""))
            if needs_chip(shares):
                chips.append(name)

    rendered = render(chips)

    if args.check:
        # A logo we couldn't measure isn't evidence the committed file is wrong —
        # it's a missing measurement. One flaky curl (or a Linux runner with no
        # `sips`) would otherwise drop that package from `chips`, fail the
        # comparison, and tell the operator to re-run a generator that would commit
        # the regression. Report and bail out as "unknown" instead of "stale".
        if skipped:
            print("Could not classify %d logo(s); not comparing:" % len(skipped))
            for name, why in skipped:
                print("  %-24s %s" % (name, why))
            return 2

        current = ""
        if os.path.exists(OUT_PATH):
            with open(OUT_PATH, encoding="utf-8") as handle:
                current = handle.read()
        if current != rendered:
            print("external_logo_treatment.yaml is out of date.")
            print("Run: python3 scripts/classify-external-logos.py")
            return 1
        print("external_logo_treatment.yaml is up to date (%d of %d logos chipped)."
              % (len(chips), len(urls)))
        return 0

    with open(OUT_PATH, "w", encoding="utf-8") as handle:
        handle.write(rendered)

    print("external logos: %d" % len(urls))
    print("need a chip:    %d" % len(chips))
    print("read fine:      %d" % (len(urls) - len(chips) - len(skipped)))
    if skipped:
        print("undecidable:    %d (left un-chipped)" % len(skipped))
        for name, why in skipped:
            print("  %-24s %s" % (name, why))
    return 0


if __name__ == "__main__":
    sys.exit(main())
