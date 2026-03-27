#!/usr/bin/env python3
"""
Publish changed packages to the Pulumi registry using registry-mirror-tools.

Runs registry-mirror-discover and registry-mirror-publish with retry logic.
Falls back to push-registry.py on failure.

Usage:
    uv run --with pyyaml scripts/ci/publish_to_registry.py
    uv run --with pyyaml scripts/ci/publish_to_registry.py --dry-run
"""

import argparse
import json
import os
import subprocess
import sys
import time
from dataclasses import dataclass
from pathlib import Path

import yaml

REGISTRY_MIRROR_TOOLS_COMMIT = "dda1dfd85d540fab45a0d19060e963564d0b36aa"

@dataclass
class Config:
    repo_root: Path
    max_retries: int = 3
    initial_backoff: int = 10
    max_backoff: int = 30
    dry_run: bool = False


def get_changed_packages(repo_root: Path) -> list[str]:
    """Get list of changed package YAML files from git diff."""
    try:
        result = subprocess.run(
            ["git", "diff", "--name-only", "HEAD~1", "--",
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
    publishers_path = repo_root / "tools/resourcedocsgen/pkg/publishers/publisher-names.json"
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

    pkg = yaml_path.stem

    if pkg.startswith("azure-native-v") or pkg.startswith("aws-v"):
        return SpecResult(None, skipped=True)

    data = load_yaml_file(yaml_path)
    if data is None:
        return SpecResult(None, error=f"Failed to parse {yaml_file}")

    version = data.get("version")
    publisher_display = data.get("publisher")
    schema_url = data.get("schema_file_url")

    if publisher_display == "DEPRECATED":
        return SpecResult(None, skipped=True)

    if not version:
        return SpecResult(None, error=f"{yaml_file} has no version field")

    publisher = publishers.get(publisher_display, "pulumi")
    source = "opentofu" if schema_url and "opentofu" in schema_url else "pulumi"
    version = version.lstrip("v")

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
    args = parser.parse_args()

    if not os.environ.get("PULUMI_API_URL"):
        print("Error: PULUMI_API_URL environment variable is required")
        return 1

    config = Config(
        repo_root=args.repo_root.resolve(),
        max_retries=args.max_retries,
        dry_run=args.dry_run,
    )

    errors = []
    if args.specs:
        specs = args.specs
    else:
        changed_files = get_changed_packages(config.repo_root)
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
