#!/usr/bin/env python3
"""Check every listed package's repo_url against GitHub and track drift in one issue.

GitHub keeps serving a renamed or transferred repo from its old URL, so a stale
repo_url still works and nobody notices until the redirect breaks or the old
owner name is reused. This asks the API about each repo (which follows the
redirect) and reports the packages whose repo has moved, been archived, or
disappeared. The report lives in a single open issue that is rewritten on each
run and closed once nothing is left to fix.
"""

import argparse
import json
import os
import re
import sys
from pathlib import Path

import requests
import yaml
from requests.adapters import HTTPAdapter
from urllib3.util.retry import Retry

REPO = "pulumi/registry"
PACKAGES_DIR = Path("themes/default/data/registry/packages")
PACKAGE_LIST = Path("community-packages/package-list.json")
MARKER = "<!-- repo-url-audit -->"
LABEL = "kind/chore"
TITLE = "Package repo URLs have drifted from GitHub"
PIPELINE_REPO = "pulumi/terraform-to-pulumi-registry-pipeline"

GITHUB_URL = re.compile(r"^https://github\.com/([^/\s]+)/([^/\s]+?)(?:\.git)?/?$")


def require_env(name):
    value = os.getenv(name)
    if not value:
        sys.exit(f"{name} is not set")
    return value


def session():
    s = requests.Session()
    retry = Retry(total=4, backoff_factor=2, status_forcelist=[500, 502, 503, 504],
                  allowed_methods=["GET"])
    s.mount("https://", HTTPAdapter(max_retries=retry))
    s.headers.update({
        "Accept": "application/vnd.github+json",
        "Authorization": f"Bearer {require_env('GITHUB_TOKEN')}",
    })
    return s


def load_packages(packages_dir=PACKAGES_DIR):
    """Listed packages only: a DEPRECATED package is already off the registry."""
    packages = []
    for path in sorted(packages_dir.glob("*.yaml")):
        data = yaml.safe_load(path.read_text()) or {}
        if data.get("publisher") == "DEPRECATED":
            continue
        packages.append({
            "name": path.stem,
            "repo_url": str(data.get("repo_url") or ""),
            "bridged": "registry.opentofu.org" in str(data.get("schema_file_url") or ""),
        })
    return packages


def load_listed_slugs(package_list=PACKAGE_LIST):
    entries = json.loads(package_list.read_text())["include"]
    return {entry["repoSlug"].lower() for entry in entries}


def slug_of(repo_url):
    match = GITHUB_URL.match(repo_url)
    return f"{match.group(1)}/{match.group(2)}" if match else None


def check(package, lookup, listed_slugs):
    """Return a finding for the package, or None when its repo_url is current.

    lookup(slug) returns the API's repo object, or None for a 404.
    """
    slug = slug_of(package["repo_url"])
    base = {**package, "slug": slug, "in_package_list": bool(slug) and slug.lower() in listed_slugs}
    if not slug:
        return {**base, "kind": "invalid"}
    repo = lookup(slug)
    if repo is None:
        return {**base, "kind": "missing"}
    current = repo["full_name"]
    # GitHub slugs are case-insensitive, so a case-only difference still resolves.
    moved_to = current if current.lower() != slug.lower() else None
    if repo.get("archived"):
        return {**base, "kind": "archived", "moved_to": moved_to}
    if moved_to:
        return {**base, "kind": "moved", "moved_to": moved_to}
    return None


def github_lookup(s):
    def lookup(slug):
        response = s.get(f"https://api.github.com/repos/{slug}", timeout=30)
        if response.status_code == 404:
            return None
        response.raise_for_status()
        return response.json()
    return lookup


def _where(finding):
    if finding["bridged"]:
        return f"`{PIPELINE_REPO}` (dynamically bridged)"
    if finding["in_package_list"]:
        return "YAML + `package-list.json`"
    return "YAML"


def render(findings):
    by_kind = {kind: [f for f in findings if f["kind"] == kind]
               for kind in ("moved", "archived", "missing", "invalid")}
    lines = [
        MARKER,
        "",
        "The weekly repo URL audit (`scripts/ci/repo_url_audit.py`) found listed packages whose `repo_url` "
        "no longer matches GitHub. This issue is rewritten on every run and closes itself once the list is empty.",
        "",
        "Where to fix it:",
        "- **YAML**: `repo_url` in `themes/default/data/registry/packages/<name>.yaml`.",
        "- **YAML + `package-list.json`**: also update `repoSlug` in `community-packages/package-list.json`. "
        "`resourcedocsgen metadata from-github` rebuilds `repo_url` from it, so otherwise the next nightly run "
        "puts the old URL back.",
        f"- **`{PIPELINE_REPO}`**: the pipeline regenerates this YAML on every release. File the change there, "
        "and re-point the provider's `fqtp` in `watched-providers` if its Terraform registry namespace moved too.",
    ]
    if by_kind["moved"]:
        lines += ["", "## Moved", "", "| Package | Current `repo_url` | Now at | Fix in |", "|---|---|---|---|"]
        lines += [f"| {f['name']} | `{f['slug']}` | [`{f['moved_to']}`](https://github.com/{f['moved_to']}) "
                  f"| {_where(f)} |" for f in by_kind["moved"]]
    if by_kind["archived"]:
        lines += ["", "## Archived upstream", "",
                  "Still listed, but the repo is archived. Re-point it to a maintained successor, or delist it "
                  "with `/delist-package <name>`. For a dynamically bridged package, start in "
                  f"`{PIPELINE_REPO}`: re-point its `fqtp`, or remove it from `watched-providers` before delisting "
                  "it here, since the pipeline's next push relists a package that is only delisted here.", "",
                  "| Package | Current `repo_url` | Now at | Fix in |", "|---|---|---|---|"]
        lines += [f"| {f['name']} | `{f['slug']}` | {f'`{m}`' if (m := f['moved_to']) else 'unchanged'} "
                  f"| {_where(f)} |" for f in by_kind["archived"]]
    if by_kind["missing"]:
        lines += ["", "## Not found", "", "GitHub returns 404: the repo was deleted or made private.", "",
                  "| Package | Current `repo_url` | Fix in |", "|---|---|---|"]
        lines += [f"| {f['name']} | `{f['slug']}` | {_where(f)} |" for f in by_kind["missing"]]
    if by_kind["invalid"]:
        lines += ["", "## Not a GitHub repo URL", "", "| Package | `repo_url` |", "|---|---|"]
        lines += [f"| {f['name']} | `{f['repo_url'] or '(empty)'}` |" for f in by_kind["invalid"]]
    return "\n".join(lines) + "\n"


def find_open_issue(s):
    response = s.get(f"https://api.github.com/repos/{REPO}/issues",
                     params={"state": "open", "labels": LABEL, "per_page": 100}, timeout=30)
    response.raise_for_status()
    for issue in response.json():
        if "pull_request" not in issue and MARKER in (issue.get("body") or ""):
            return issue
    return None


def sync_issue(s, body):
    """Create, rewrite, or close the tracking issue so it matches body (None = nothing to fix)."""
    issue = find_open_issue(s)
    api = f"https://api.github.com/repos/{REPO}/issues"
    if body is None:
        if issue:
            s.post(f"{api}/{issue['number']}/comments", timeout=30,
                   json={"body": "Every listed package's `repo_url` matches GitHub now. Closing."}
                   ).raise_for_status()
            s.patch(f"{api}/{issue['number']}", json={"state": "closed", "state_reason": "completed"},
                    timeout=30).raise_for_status()
            print(f"nothing to fix; closed #{issue['number']}")
        else:
            print("nothing to fix")
        return
    if issue:
        if issue.get("body") != body:
            s.patch(f"{api}/{issue['number']}", json={"body": body}, timeout=30).raise_for_status()
            print(f"updated #{issue['number']}")
        else:
            print(f"#{issue['number']} is already current")
        return
    created = s.post(api, json={"title": TITLE, "body": body, "labels": [LABEL]}, timeout=30)
    created.raise_for_status()
    print(f"opened #{created.json()['number']}")


def main():
    parser = argparse.ArgumentParser(description=__doc__.splitlines()[0])
    parser.add_argument("--dry-run", action="store_true", help="Print the report instead of filing it")
    args = parser.parse_args()

    s = session()
    lookup = github_lookup(s)
    listed_slugs = load_listed_slugs()
    findings = [f for p in load_packages() if (f := check(p, lookup, listed_slugs))]
    body = render(findings) if findings else None

    if args.dry_run:
        print(body or "nothing to fix")
        return
    sync_issue(s, body)


if __name__ == "__main__":
    main()
