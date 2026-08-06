#!/usr/bin/env python3
"""
Validate that a package's schema declares the version the registry publishes.

The registry publishes each package's schema straight from its release tag.
The registry service rejects a schema whose own version field disagrees with
the version being published, so a placeholder like 0.0.0 fails every release.
An absent version is valid: bridged providers omit the field and the service
takes the version from the publish request.

Usage:
    uv run --with pyyaml scripts/ci/validate_schema_version.py [yaml...]
    uv run --with pyyaml scripts/ci/validate_schema_version.py --all
"""

from __future__ import annotations

import argparse
import json
import re
import sys
import urllib.error
import urllib.request
from concurrent.futures import ThreadPoolExecutor
from dataclasses import dataclass
from pathlib import Path

import yaml

PACKAGES = Path("themes/default/data/registry/packages")
HEAD_BYTES = 512 * 1024
ALLOWED_MISMATCHES: dict[str, str] = {
    "aem": "https://github.com/wttech/pulumi-aem — schema keeps an 0.1.2-alpha build string",
    "genesiscloud": "https://github.com/genesiscloud/pulumi-genesiscloud — schema keeps the v prefix",
    "kafka-connect": "https://github.com/azaurus1/pulumi-kafka-connect — schema keeps an 0.0.8+dirty build string",
    "powerplatform": "https://github.com/rpothin/pulumi-powerplatform/issues/82",
    "terraform-provider": "https://github.com/pulumi/pulumi-terraform-provider/pull/113",
}


@dataclass
class Result:
    package: str
    status: str
    declared: str = ""
    expected: str = ""
    detail: str = ""

    @property
    def blocking(self) -> bool:
        return self.status == "mismatch" and self.package not in ALLOWED_MISMATCHES


def schema_url(data: dict) -> str | None:
    url = data.get("schema_file_url")
    if url:
        return str(url)
    repo, path = data.get("repo_url"), data.get("schema_file_path")
    if not repo or not path:
        return None
    slug = str(repo).removeprefix("https://github.com/")
    return f"https://raw.githubusercontent.com/{slug}/refs/tags/{data.get('version')}/{path}"


def top_level_version(text: str) -> tuple[str | None, bool]:
    depth, i, n = 0, 0, len(text)
    while i < n:
        char = text[i]
        if char == '"':
            end = i + 1
            while end < n:
                if text[end] == "\\":
                    end += 2
                    continue
                if text[end] == '"':
                    break
                end += 1
            if end >= n:
                return None, True
            if depth == 1 and text[i + 1:end] == "version":
                match = re.match(r'\s*:\s*"((?:[^"\\]|\\.)*)"', text[end + 1:])
                if match:
                    return match.group(1), False
            i = end + 1
            continue
        if char == "{":
            depth += 1
        elif char == "}":
            depth -= 1
            if depth == 0:
                return None, False
        i += 1
    return None, True


def fetch(url: str, head_only: bool) -> bytes:
    request = urllib.request.Request(url)
    if head_only:
        request.add_header("Range", f"bytes=0-{HEAD_BYTES - 1}")
    with urllib.request.urlopen(request, timeout=120) as response:
        return response.read()


def declared_version(url: str) -> str:
    if url.endswith((".yaml", ".yml")):
        return str(yaml.safe_load(fetch(url, head_only=False)).get("version", "")).strip()
    version, truncated = top_level_version(fetch(url, head_only=True).decode("utf-8", "replace"))
    if version is None and truncated:
        schema = json.loads(fetch(url, head_only=False).decode("utf-8", "replace"))
        return str(schema.get("version", "")).strip()
    return (version or "").strip()


def check_url(package: str, url: str, version: str) -> Result:
    expected = version.lstrip("v")
    try:
        declared = declared_version(url)
    except (urllib.error.URLError, OSError, ValueError, yaml.YAMLError) as err:
        return Result(package, "unreachable", expected=expected, detail=f"{type(err).__name__}: {err}")
    if not declared:
        return Result(package, "absent", expected=expected)
    if declared != expected:
        return Result(package, "mismatch", declared=declared, expected=expected)
    return Result(package, "ok", declared=declared, expected=expected)


def check(path: Path) -> Result | None:
    data = yaml.safe_load(path.read_text())
    if not data or data.get("publisher") == "DEPRECATED":
        return None
    package = str(data.get("name") or path.stem)
    expected = str(data.get("version") or "").lstrip("v")
    url = schema_url(data)
    if not url or not expected:
        return Result(package, "unresolvable", detail="no schema URL or version in the package YAML")
    return check_url(package, url, expected)


def report(results: list[Result]) -> int:
    for result in sorted(results, key=lambda r: r.package):
        if result.status == "mismatch":
            waived = "" if result.blocking else f"  (waived: {ALLOWED_MISMATCHES[result.package]})"
            print(f"MISMATCH  {result.package}: schema declares {result.declared!r}, "
                  f"registry publishes {result.expected!r}{waived}")
        elif result.status in ("unreachable", "unresolvable"):
            print(f"SKIP      {result.package}: {result.detail}")

    blocking = [r for r in results if r.blocking]
    counts = {status: sum(1 for r in results if r.status == status)
              for status in ("ok", "absent", "mismatch", "unreachable", "unresolvable")}
    print(f"\n{len(results)} packages checked: {counts['ok']} matching, {counts['absent']} without a "
          f"version field, {counts['mismatch']} mismatched, "
          f"{counts['unreachable'] + counts['unresolvable']} not checkable.")

    if not blocking:
        return 0
    print(f"\n{len(blocking)} package(s) declare a version their release does not match. The registry "
          "service rejects these:\n\n  400 Bad Request: Schema version (X) must match query parameter "
          "version (Y)\n\nFix upstream by omitting the schema's version field, which is what bridged "
          "providers do, or by stamping the real version at release time.")
    return 1


def main() -> int:
    parser = argparse.ArgumentParser(description="Validate package schema versions")
    parser.add_argument("files", nargs="*", type=Path, help="package YAML files to check")
    parser.add_argument("--all", action="store_true", help="check every package in the registry")
    parser.add_argument("--schema-url", help="check a schema by URL instead of a package YAML")
    parser.add_argument("--version", help="the version the schema is expected to declare")
    parser.add_argument("--package", default="", help="package name to report against --schema-url")
    args = parser.parse_args()

    if args.schema_url:
        if not args.version:
            parser.error("--schema-url requires --version")
        return report([check_url(args.package or args.schema_url, args.schema_url, args.version)])

    paths = sorted(PACKAGES.glob("*.yaml")) if args.all else [p for p in args.files if p.exists()]
    if not paths:
        print("No package YAML files to check.")
        return 0

    with ThreadPoolExecutor(max_workers=16) as pool:
        results = [r for r in pool.map(check, paths) if r is not None]
    return report(results)


if __name__ == "__main__":
    sys.exit(main())
