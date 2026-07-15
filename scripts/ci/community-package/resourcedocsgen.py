from __future__ import annotations

import json
import os
import subprocess
from pathlib import Path

import github_api
from models import Entry, provider_name

BINARY = Path(os.environ.get("RESOURCEDOCSGEN") or "tools/resourcedocsgen/resourcedocsgen")


def ensure_built() -> None:
    if not BINARY.exists():
        subprocess.run(["go", "build", "-C", "tools/resourcedocsgen"], check=True)


def generate_metadata(slug: str, schema_file: str, tag: str, into: Path | None = None) -> bool:
    args = [str(BINARY), "metadata", "from-github",
            "--repoSlug", slug, "--schemaFile", schema_file, "--version", tag]
    if into is not None:
        args += ["--metadataDir", str(into / "data"), "--packageDocsDir", str(into / "content")]
    return subprocess.run(args, capture_output=True, text=True).returncode == 0


def publish(entries: list[Entry]) -> list[tuple[str, str]]:
    published: list[tuple[str, str]] = []
    for entry in entries:
        tag = github_api.latest_release_tag(entry.repoSlug)
        schema_bytes = github_api.raw_file(entry.repoSlug, tag, entry.schemaFile)
        if schema_bytes is None:
            raise FileNotFoundError(f"{entry.repoSlug}@{tag}/{entry.schemaFile}")
        name = provider_name(entry.repoSlug, json.loads(schema_bytes))
        if not generate_metadata(entry.repoSlug, entry.schemaFile, tag):
            raise SystemExit(f"resourcedocsgen failed for {name}@{tag}")
        published.append((name, tag))
    return published
