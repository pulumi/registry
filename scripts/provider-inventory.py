#!/usr/bin/env python3
"""Build an automated inventory of all registry provider pages.

Scans every provider YAML under themes/default/data/registry/packages/,
cross-references the committed content pages under
themes/default/content/registry/packages/ and the community package list in
community-packages/package-list.json, and writes a single JSON inventory.

Usage:
    python3 scripts/provider-inventory.py [-o OUTPUT.json]
"""

import argparse
import json
import sys
from datetime import datetime, timezone
from pathlib import Path

import yaml

REPO_ROOT = Path(__file__).resolve().parent.parent
PACKAGES_DIR = REPO_ROOT / "themes/default/data/registry/packages"
CONTENT_DIR = REPO_ROOT / "themes/default/content/registry/packages"
COMMUNITY_LIST = REPO_ROOT / "community-packages/package-list.json"


SECTION_CONCEPTS = [
    ("example", r"example|usage"),
    ("installation", r"install"),
    ("auth", r"credential|auth|token|api.key|service.account|logging in"),
    ("configuration", r"configur|setup|settings"),
    ("prerequisites", r"prerequisit|getting started|quick.?start|before you begin"),
    ("migration", r"migrat|from v\d|upgrad"),
    ("reference", r"reference|input propert|outputs"),
    ("development", r"develop|contribut|building"),
]

CHOOSER_LANGS = ["typescript", "javascript", "python", "go", "csharp", "java", "yaml"]


def analyze_page(path):
    """Fence-aware structural analysis of a content markdown page."""
    import re
    if not path.exists():
        return None
    text = path.read_text(errors="replace")
    frontmatter = {}
    body = text
    m_fm = re.match(r"\A(?:<!--.*?-->\s*)?---\n(.*?)\n---\n(.*)\Z", text, re.S)
    if m_fm:
        try:
            frontmatter = yaml.safe_load(m_fm.group(1)) or {}
        except yaml.YAMLError:
            frontmatter = {}
        body = m_fm.group(2)
    headings, fence = [], False
    for line in body.splitlines():
        if line.startswith("```"):
            fence = not fence
            continue
        m = re.match(r"^(#{1,4})\s+(.*)", line) if not fence else None
        if m:
            headings.append((len(m.group(1)), m.group(2).strip()))
    joined = " | ".join(h[1].lower() for h in headings)
    concepts = [c for c, pat in SECTION_CONCEPTS if re.search(pat, joined)]
    langs = sorted({l for l in CHOOSER_LANGS
                    if re.search(r"choosable\s+language\s+%s\b" % l, body)})
    words = len(re.sub(r"```.*?```", " ", body, flags=re.S).split())
    return {
        "layout": frontmatter.get("layout", ""),
        "words": words,
        "headings": [h[1] for h in headings],
        "concepts": concepts,
        "has_code": "```" in body,
        "chooser_langs": langs,
        "vendor_fetched": "this file was fetched" in text[:400],
    }


def analyze_extras(content_dir):
    extras = []
    if content_dir.is_dir():
        for child in sorted(content_dir.iterdir()):
            if child.name in ("_index.md", "installation-configuration.md"):
                continue
            if child.is_dir():
                extras.append(child.name + "/ (%d files)" % len(list(child.rglob("*.md"))))
            else:
                extras.append(child.name)
    return extras


def load_community_slugs():
    slugs = set()
    if COMMUNITY_LIST.exists():
        data = json.loads(COMMUNITY_LIST.read_text())
        for entry in data.get("include", []):
            repo = entry.get("repoSlug", "")
            name = repo.split("/")[-1]
            if name.startswith("pulumi-"):
                name = name[len("pulumi-"):]
            slugs.add(name)
    return slugs


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("-o", "--output", default="provider-inventory.json")
    args = parser.parse_args()

    community_slugs = load_community_slugs()
    now = datetime.now(timezone.utc)
    packages = []
    all_keys = set()

    for path in sorted(PACKAGES_DIR.glob("*.yaml")):
        data = yaml.safe_load(path.read_text())
        if not isinstance(data, dict):
            print(f"WARN: {path.name} did not parse to a mapping", file=sys.stderr)
            continue
        all_keys.update(data.keys())
        name = data.get("name", path.stem)

        content = CONTENT_DIR / path.stem
        updated_on = data.get("updated_on")
        age_days = None
        if isinstance(updated_on, int):
            age_days = round((now - datetime.fromtimestamp(updated_on, tz=timezone.utc)).days)

        repo_url = data.get("repo_url") or ""
        org = ""
        if "github.com/" in repo_url:
            org = repo_url.split("github.com/")[1].split("/")[0]

        packages.append({
            "name": name,
            "slug": path.stem,
            "title": data.get("title", name),
            "category": data.get("category", ""),
            "publisher": data.get("publisher", ""),
            "component": bool(data.get("component", False)),
            "native": bool(data.get("native", False)),
            "package_status": data.get("package_status", ""),
            "featured": bool(data.get("featured", False)),
            "version": data.get("version", ""),
            "repo_url": repo_url,
            "repo_org": org,
            "schema_file_url": data.get("schema_file_url", ""),
            "description": data.get("description", ""),
            "updated_on": updated_on,
            "age_days": age_days,
            "in_community_list": path.stem in community_slugs,
            "has_index_page": (content / "_index.md").exists(),
            "has_install_config_page": (content / "installation-configuration.md").exists(),
            "overview_page": analyze_page(content / "_index.md"),
            "install_page": analyze_page(content / "installation-configuration.md"),
            "extras": analyze_extras(content),
            "registry_url": f"https://www.pulumi.com/registry/packages/{path.stem}/",
        })

    inventory = {
        "generated_at": now.isoformat(timespec="seconds"),
        "package_count": len(packages),
        "yaml_keys_seen": sorted(all_keys),
        "packages": packages,
    }
    Path(args.output).write_text(json.dumps(inventory, indent=2) + "\n")
    print(f"Wrote {len(packages)} packages to {args.output}")


if __name__ == "__main__":
    main()
