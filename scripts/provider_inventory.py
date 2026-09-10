#!/usr/bin/env python3
"""Build an automated inventory of all registry provider pages.

Scans every provider YAML under themes/default/data/registry/packages/,
cross-references the committed content pages under
themes/default/content/registry/packages/ and the community package list in
community-packages/package-list.json, and writes a single JSON inventory.

Can be run without preinstalled dependencies with

	uv run --with pyyaml scripts/provider_inventory.py

Usage:
    python3 scripts/provider_inventory.py [-o OUTPUT.json]
"""

import argparse
import json
import re
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

# Hugo shortcode: {{% choosable language go %}} or {{% choosable language "a,b" %}}
CHOOSER_RE = r'choosable\s+language\s+"?([a-z0-9,\s.-]+?)"?\s*%'


def analyze_page(path):
    """Fence-aware structural analysis of a content markdown page."""
    if not path.exists():
        return None
    text = path.read_text(errors="replace")
    frontmatter = {}
    body = text
    m_fm = re.match(r"\A(?:<!--.*?-->\s*)?---\n(.*?)\n---\n(.*)\Z", text, re.S)
    if m_fm:
        try:
            frontmatter = yaml.safe_load(m_fm.group(1))
        except yaml.YAMLError:
            frontmatter = {}
        if not isinstance(frontmatter, dict):
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
    # Extract the languages actually offered rather than probing a fixed list,
    # so quoted forms, comma-separated groups, and languages beyond the Pulumi
    # SDK set (hcl, nodejs, dotnet) are all counted.
    langs = sorted({part.strip() for m in re.finditer(CHOOSER_RE, body)
                    for part in m.group(1).split(",") if part.strip()})
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
    # How-to guides are no longer generated; the only ones in the tree are the
    # hand-committed migration guides under a few packages' how-to-guides/ dirs.
    extras = {"howto_guides": 0, "migration_dirs": [], "other_files": []}
    if content_dir.is_dir():
        for child in sorted(content_dir.iterdir()):
            if child.name in ("_index.md", "installation-configuration.md"):
                continue
            if child.is_dir():
                # api-docs/ is git-ignored build output present only on a built
                # tree; skipping it keeps the inventory identical across checkouts.
                if child.name == "api-docs":
                    continue
                count = len(list(child.rglob("*.md")))
                if child.name == "how-to-guides":
                    extras["howto_guides"] = count
                else:
                    extras["migration_dirs"].append(f"{child.name} ({count})")
            else:
                extras["other_files"].append(child.name)
    return extras


def load_community_slugs():
    """Return (package slugs, repo slugs) derived from the community package list.

    Repo name and package name diverge often enough that stripping the
    "pulumi-" prefix alone is wrong: runpod/pulumi-runpod-native publishes
    "runpod", and DefangLabs/pulumi-defang publishes three separate packages.
    The schemaFile path carries the canonical package name; repo slugs cover
    the rest.
    """
    slugs, repos = set(), set()
    if COMMUNITY_LIST.exists():
        data = json.loads(COMMUNITY_LIST.read_text())
        for entry in data.get("include", []):
            repo = entry.get("repoSlug", "")
            if repo:
                repos.add(repo.lower())
            m = re.search(r"pulumi-resource-([^/]+)/", entry.get("schemaFile", ""))
            if m:
                slugs.add(m.group(1))
                continue
            name = repo.split("/")[-1]
            if name.startswith("pulumi-"):
                name = name[len("pulumi-"):]
            slugs.add(name)
    return slugs, repos


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("-o", "--output", default="provider-inventory.json")
    args = parser.parse_args()

    community_slugs, community_repos = load_community_slugs()
    now = datetime.now(timezone.utc)
    packages = []
    all_keys = set()

    for path in sorted(PACKAGES_DIR.glob("*.yaml")):
        try:
            data = yaml.safe_load(path.read_text())
        except yaml.YAMLError as err:
            print(f"WARN: {path.name} could not be parsed: {err}", file=sys.stderr)
            continue
        if not isinstance(data, dict):
            print(f"WARN: {path.name} did not parse to a mapping", file=sys.stderr)
            continue
        all_keys.update(data.keys())
        name = data.get("name", path.stem)

        content = CONTENT_DIR / path.stem
        updated_on = data.get("updated_on")
        age_days = None
        if isinstance(updated_on, int):
            age_days = (now - datetime.fromtimestamp(updated_on, tz=timezone.utc)).days

        repo_url = data.get("repo_url") or ""
        org = ""
        if "github.com/" in repo_url:
            org = repo_url.split("github.com/")[1].split("/")[0]

        repo_slug = repo_url.rstrip("/").split("github.com/")[-1].lower()

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
            "in_community_list": (path.stem in community_slugs
                                  or repo_slug in community_repos),
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
