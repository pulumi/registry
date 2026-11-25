#!/usr/bin/env python3
"""
Publish changed packages to the Pulumi registry using registry-mirror-tools.

Runs registry-mirror-discover and registry-mirror-publish with retry logic.
Falls back to push-registry.py on failure.

Usage:
    uv run --with pyyaml scripts/ci/publish_to_registry.py
    uv run --with pyyaml scripts/ci/publish_to_registry.py --dry-run
    uv run --with pyyaml scripts/ci/publish_to_registry.py --reconcile --dry-run
"""

import argparse
import json
import os
import subprocess
import sys
import time
import urllib.error
import urllib.parse
import urllib.request
from dataclasses import dataclass
from pathlib import Path

import yaml

REGISTRY_MIRROR_TOOLS_COMMIT = "dda1dfd85d540fab45a0d19060e963564d0b36aa"
PUBLISHER_NAMES_RELPATH = "tools/resourcedocsgen/pkg/publishers/publisher-names.json"
PACKAGES_RELDIR = "themes/default/data/registry/packages"
SKIP_LIST_RELPATH = "scripts/ci/publish-skip-list.json"
# Versions that landed before this commit are taken as they stand in the registry.
RECONCILE_FLOOR = "157f57b29deb3d1ba460dc71055acaf2d0a463d5"

@dataclass
class Config:
    repo_root: Path
    max_retries: int = 3
    initial_backoff: int = 10
    max_backoff: int = 30
    dry_run: bool = False


def commit_exists(repo_root: Path, rev: str) -> bool:
    result = subprocess.run(
        ["git", "cat-file", "-e", f"{rev}^{{commit}}"],
        cwd=repo_root,
        capture_output=True,
    )
    return result.returncode == 0


def get_changed_packages(repo_root: Path, base: str = "HEAD~1") -> list[str]:
    """Get list of changed package YAML files from git diff."""
    try:
        result = subprocess.run(
            ["git", "diff", "--name-only", base, "HEAD", "--",
             "themes/default/data/registry/packages/*.yaml"],
            cwd=repo_root,
            capture_output=True,
            text=True,
            check=True,
        )
        return [f for f in result.stdout.strip().split("\n") if f]
    except subprocess.CalledProcessError:
        return []


def load_publishers(repo_root: Path) -> dict[str, str]:
    """Load publisher display name to slug mapping."""
    publishers_path = repo_root / PUBLISHER_NAMES_RELPATH
    with open(publishers_path) as f:
        return json.load(f)


def load_yaml_file(yaml_path: Path) -> dict | None:
    """Load and parse a YAML file, returning None on error."""
    try:
        with open(yaml_path) as f:
            return yaml.safe_load(f)
    except (OSError, yaml.YAMLError) as e:
        print(f"Error loading {yaml_path}: {e}")
        return None


class SpecResult:
    """Result of building a package spec."""

    def __init__(self, spec: str | None, error: str | None = None, skipped: bool = False):
        self.spec = spec
        self.error = error
        self.skipped = skipped


def build_package_spec(
    yaml_file: str,
    repo_root: Path,
    publishers: dict[str, str],
) -> SpecResult:
    """Build a package spec string from a YAML file path.

    Returns a SpecResult with:
    - spec set if successful
    - skipped=True for expected skips (legacy, deprecated)
    - error set for unexpected failures that should cause the workflow to fail
    """
    yaml_path = repo_root / yaml_file
    if not yaml_path.exists():
        return SpecResult(None, error=f"{yaml_file} does not exist")

    data = load_yaml_file(yaml_path)
    if data is None:
        return SpecResult(None, error=f"Failed to parse {yaml_file}")

    return spec_from_package_data(yaml_file, data, publishers)


def spec_from_package_data(yaml_file: str, data: dict, publishers: dict[str, str]) -> SpecResult:
    pkg = Path(yaml_file).stem
    if pkg.startswith("azure-native-v") or pkg.startswith("aws-v"):
        return SpecResult(None, skipped=True)

    version = data.get("version")
    publisher_display = data.get("publisher")
    schema_url = data.get("schema_file_url")

    # Older delistings marked the package with a "DEPRECATED" publisher instead of `deprecated: true`.
    if data.get("deprecated") or publisher_display == "DEPRECATED":
        return SpecResult(None, skipped=True)

    if not version:
        return SpecResult(None, error=f"{yaml_file} has no version field")

    if not publisher_display:
        return SpecResult(None, error=f"{yaml_file} has no publisher field")

    if publisher_display not in publishers:
        return SpecResult(None, error=(
            f'{yaml_file} has publisher "{publisher_display}", which is not in '
            f"{PUBLISHER_NAMES_RELPATH}"
        ))

    publisher = publishers[publisher_display]
    source = "opentofu" if schema_url and "opentofu" in schema_url else "pulumi"
    version = str(version).lstrip("v")

    return SpecResult(f"{source}/{publisher}/{pkg}@{version}")


def build_specs(changed_files: list[str], repo_root: Path) -> tuple[list[str], list[str]]:
    """Build package specs from changed YAML files.

    Returns a tuple of (specs, errors) where errors are messages for packages
    that failed to parse unexpectedly.
    """
    publishers = load_publishers(repo_root)
    specs = []
    errors = []
    for yaml_file in changed_files:
        result = build_package_spec(yaml_file, repo_root, publishers)
        if result.spec:
            specs.append(result.spec)
        elif result.error:
            errors.append(result.error)
    return specs, errors


def ensure_tools_installed(repo_root: Path) -> tuple[Path, Path]:
    """Install registry-mirror-tools if not present."""
    bin_dir = repo_root / "bin"
    bin_dir.mkdir(exist_ok=True)

    discover_bin = bin_dir / "registry-mirror-discover"
    publish_bin = bin_dir / "registry-mirror-publish"

    env = os.environ.copy()
    env["GOBIN"] = str(bin_dir)
    env["GOPRIVATE"] = "github.com/pulumi"

    for tool, bin_path in [("registry-mirror-discover", discover_bin),
                           ("registry-mirror-publish", publish_bin)]:
        if not bin_path.exists():
            print(f"Installing {tool}...")
            subprocess.run(
                ["go", "install",
                 f"github.com/pulumi/registry-mirror-tools/cmd/{tool}@{REGISTRY_MIRROR_TOOLS_COMMIT}"],
                env=env,
                check=True,
            )

    return discover_bin, publish_bin


def publish_specs(
    specs: list[str],
    repo_root: Path,
    discover_bin: Path,
    publish_bin: Path,
    dry_run: bool = False,
) -> bool:
    """Run discover and publish pipeline for the given specs."""
    specs_input = "\n".join(specs)

    publish_cmd = [str(publish_bin), str(repo_root)]
    if dry_run:
        publish_cmd.append("--dry-run")

    discover = subprocess.Popen(
        [str(discover_bin), str(repo_root)],
        stdin=subprocess.PIPE,
        stdout=subprocess.PIPE,
        text=True,
    )

    publish = subprocess.Popen(
        publish_cmd,
        stdin=discover.stdout,
        text=True,
    )

    discover.stdin.write(specs_input)
    discover.stdin.close()
    discover.stdout.close()

    publish.wait()
    discover.wait()

    return discover.returncode == 0 and publish.returncode == 0


def publish_with_retry(specs: list[str], config: Config) -> bool:
    """Publish specs with exponential backoff retry.

    Assumes specs is non-empty. Returns True on success, False on failure.
    """
    discover_bin, publish_bin = ensure_tools_installed(config.repo_root)

    print("Packages to publish:")
    for spec in specs:
        print(f"  {spec}")
    print()

    backoff = config.initial_backoff
    for attempt in range(1, config.max_retries + 1):
        print(f"Attempt {attempt} of {config.max_retries}...")

        if publish_specs(specs, config.repo_root, discover_bin, publish_bin, config.dry_run):
            print("Successfully published packages")
            return True

        if attempt < config.max_retries:
            print(f"Publish failed, retrying in {backoff}s...")
            time.sleep(backoff)
            backoff = min(backoff * 2, config.max_backoff)

    print(f"Failed to publish after {config.max_retries} attempts")
    return False


def all_package_yamls(repo_root: Path) -> list[str]:
    package_dir = repo_root / PACKAGES_RELDIR
    return sorted(str(p.relative_to(repo_root)) for p in package_dir.glob("*.yaml"))


def validate_all(repo_root: Path) -> int:
    yaml_files = all_package_yamls(repo_root)
    specs, errors = build_specs(yaml_files, repo_root)

    print(f"{len(yaml_files)} package YAML files, {len(specs)} publishable specs")
    if not errors:
        return 0

    print(f"\n{len(errors)} package(s) could not be turned into a publish spec:")
    for error in errors:
        print(f"  - {error}")
    return 1


def package_yaml_revisions(repo_root: Path, floor: str) -> list[tuple[str, str]]:
    """Return (commit, path) for every package YAML changed on HEAD's first-parent line since floor."""
    result = subprocess.run(
        ["git", "log", "--first-parent", "--format=commit %H", "--name-only",
         f"{floor}..HEAD", "--", f"{PACKAGES_RELDIR}/*.yaml"],
        cwd=repo_root,
        capture_output=True,
        text=True,
        check=True,
    )
    revisions = []
    commit = None
    for line in result.stdout.splitlines():
        if line.startswith("commit "):
            commit = line.removeprefix("commit ")
        elif line:
            revisions.append((commit, line))
    return revisions


def read_blobs(repo_root: Path, revs: list[str]) -> dict[str, bytes | None]:
    """Read `<commit>:<path>` blobs in one git process; None marks a path deleted in that commit."""
    proc = subprocess.run(
        ["git", "cat-file", "--batch"],
        cwd=repo_root,
        input="".join(f"{rev}\n" for rev in revs).encode(),
        capture_output=True,
        check=True,
    )
    out = proc.stdout
    blobs = {}
    pos = 0
    for rev in revs:
        header_end = out.index(b"\n", pos)
        header = out[pos:header_end].split()
        pos = header_end + 1
        if header[-1] == b"missing":
            blobs[rev] = None
            continue
        size = int(header[2])
        blobs[rev] = out[pos:pos + size]
        pos += size + 1
    return blobs


def historical_specs(
    repo_root: Path,
    floor: str,
    publishers: dict[str, str],
    live_paths: set[str],
    skipped_revisions: set[str],
) -> tuple[set[str], list[str]]:
    """Build a spec from every revision since floor of a package YAML that is still publishable at HEAD.

    Packages delisted or removed since are left out, so reconciling never republishes them.
    """
    revisions = [
        f"{commit}:{path}"
        for commit, path in package_yaml_revisions(repo_root, floor)
        if path in live_paths and f"{commit}:{path}" not in skipped_revisions
    ]
    specs = set()
    errors = []
    for rev, blob in read_blobs(repo_root, revisions).items():
        if blob is None:
            continue
        path = rev.split(":", 1)[1]
        try:
            data = yaml.safe_load(blob)
        except yaml.YAMLError as e:
            errors.append(f"{rev}: {e}")
            continue
        if not isinstance(data, dict):
            errors.append(f"{rev}: not a YAML mapping")
            continue
        result = spec_from_package_data(path, data, publishers)
        if result.spec:
            specs.add(result.spec)
        elif result.error:
            errors.append(f"{rev}: {result.error}")
    return specs, errors


def split_spec(spec: str) -> tuple[str, str]:
    package, version = spec.rsplit("@", 1)
    return package, version


def published_versions(package: str) -> set[str]:
    """List every published version of `<source>/<publisher>/<name>`; empty if the package doesn't exist."""
    base_url = f"{os.environ['PULUMI_API_URL'].rstrip('/')}/api/registry/packages/{package}/versions"
    headers = {"Accept": "application/json"}
    token = os.environ.get("PULUMI_ACCESS_TOKEN")
    if token:
        headers["Authorization"] = f"token {token}"

    versions = set()
    continuation_token = None
    while True:
        query = {"limit": "500"}
        if continuation_token:
            query["continuationToken"] = continuation_token
        request = urllib.request.Request(f"{base_url}?{urllib.parse.urlencode(query)}", headers=headers)
        try:
            with urllib.request.urlopen(request, timeout=30) as response:
                body = json.load(response)
        except urllib.error.HTTPError as e:
            if e.code == 404:
                return set()
            raise
        versions.update(p["version"] for p in body.get("packages", []))
        continuation_token = body.get("continuationToken")
        if not continuation_token:
            return versions


def missing_specs(desired: set[str]) -> list[str]:
    by_package: dict[str, set[str]] = {}
    for spec in desired:
        package, version = split_spec(spec)
        by_package.setdefault(package, set()).add(version)

    missing = []
    for package, versions in sorted(by_package.items()):
        missing.extend(f"{package}@{v}" for v in sorted(versions - published_versions(package)))
    return missing


def load_skip_list(repo_root: Path) -> dict[str, str]:
    """Map each spec or `<commit>:<path>` revision that reconciling must ignore to the reason why."""
    with open(repo_root / SKIP_LIST_RELPATH) as f:
        return json.load(f)


def reconcile(config: Config, floor: str) -> int:
    """Publish every package version on master since floor that the registry doesn't have."""
    publishers = load_publishers(config.repo_root)
    skip_list = load_skip_list(config.repo_root)

    head_specs = set()
    live_paths = set()
    errors = []
    for yaml_file in all_package_yamls(config.repo_root):
        result = build_package_spec(yaml_file, config.repo_root, publishers)
        if result.spec:
            head_specs.add(result.spec)
            live_paths.add(yaml_file)
        elif result.error:
            errors.append(result.error)

    history, history_errors = historical_specs(
        config.repo_root, floor, publishers, live_paths, set(skip_list))
    errors.extend(history_errors)
    desired = (head_specs | history) - skip_list.keys()

    print(f"{len(desired)} package versions expected in the registry "
          f"({len(head_specs)} at HEAD, {len(history)} since {floor[:12]}, {len(skip_list)} on the skip list)")

    missing = missing_specs(desired)
    if not missing:
        print("The registry has every expected version")
    elif not publish_with_retry(missing, config):
        return 1
    elif not config.dry_run:
        missing = missing_specs(set(missing))
        if missing:
            print(f"{len(missing)} versions are still missing from the registry:")
            for spec in missing:
                print(f"  {spec}")
            print(f"Publish them, or add them to {SKIP_LIST_RELPATH} with a reason.")
            return 1

    if errors:
        print("Some package versions could not be checked, because no spec could be built for them:")
        for error in errors:
            print(f"  - {error}")
        print(f"Fix the YAML, or add the spec or `<commit>:<path>` revision to {SKIP_LIST_RELPATH} with a reason.")
        return 1

    return 0


def main() -> int:
    parser = argparse.ArgumentParser(
        description="Publish changed packages to the registry"
    )
    parser.add_argument(
        "--dry-run",
        action="store_true",
        help="Show what would be published without actually publishing",
    )
    parser.add_argument(
        "--repo-root",
        type=Path,
        default=Path(__file__).parent.parent.parent,
        help="Path to the registry repository root",
    )
    parser.add_argument(
        "--max-retries",
        type=int,
        default=3,
        help="Maximum number of retry attempts (default: 3)",
    )
    parser.add_argument(
        "--specs",
        nargs="*",
        help="Package specs to publish (default: detect from git diff)",
    )
    parser.add_argument(
        "--base",
        default="HEAD~1",
        help="Commit to diff HEAD against when detecting changed packages (default: HEAD~1)",
    )
    parser.add_argument(
        "--validate-all",
        action="store_true",
        help="Build a spec for every package YAML and report errors without publishing",
    )
    parser.add_argument(
        "--reconcile",
        action="store_true",
        help="Publish every package version on master since --floor that the registry is missing",
    )
    parser.add_argument(
        "--floor",
        default=RECONCILE_FLOOR,
        help=f"Oldest commit whose package versions --reconcile checks (default: {RECONCILE_FLOOR[:12]})",
    )
    args = parser.parse_args()

    if args.validate_all:
        return validate_all(args.repo_root.resolve())

    if args.reconcile and args.specs:
        print("Error: --reconcile and --specs are mutually exclusive")
        return 1

    if not os.environ.get("PULUMI_API_URL"):
        print("Error: PULUMI_API_URL environment variable is required")
        return 1

    config = Config(
        repo_root=args.repo_root.resolve(),
        max_retries=args.max_retries,
        dry_run=args.dry_run,
    )

    if args.reconcile:
        return reconcile(config, args.floor)

    errors = []
    if args.specs:
        specs = args.specs
    else:
        if not commit_exists(config.repo_root, args.base):
            # Diffing a narrower range could report success while skipping packages, so fail and let the operator name the specs.
            print(
                f"Error: base commit {args.base} is not in the clone (force push?), "
                f"so the packages changed in {args.base}..HEAD can't be detected. "
                "Re-run publish.yml via workflow_dispatch with the changed packages as specs."
            )
            return 1
        changed_files = get_changed_packages(config.repo_root, args.base)
        if not changed_files:
            print("No package YAML files changed")
            return 0
        print(f"Changed package files: {changed_files}")
        specs, errors = build_specs(changed_files, config.repo_root)

    if errors:
        print("Errors building package specs:")
        for error in errors:
            print(f"  - {error}")

    if specs:
        if not publish_with_retry(specs, config):
            print("registry-mirror-publish failed, push-registry.py will attempt fallback")
            return 1
    else:
        print("No packages to publish")

    if errors:
        print("Failing due to errors building some package specs")
        return 1

    return 0


if __name__ == "__main__":
    sys.exit(main())
