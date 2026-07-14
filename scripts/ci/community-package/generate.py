from __future__ import annotations

import json

import github_api
import resourcedocsgen
from models import Entry
from verify_entry import provider_name


def metadata(entries: list[Entry]) -> list[tuple[str, str]]:
    published: list[tuple[str, str]] = []
    for entry in entries:
        tag = github_api.latest_release_tag(entry.repoSlug)
        schema_bytes = github_api.raw_file(entry.repoSlug, tag, entry.schemaFile)
        if schema_bytes is None:
            raise FileNotFoundError(f"{entry.repoSlug}@{tag}/{entry.schemaFile}")
        name = provider_name(entry.repoSlug, json.loads(schema_bytes))
        if not resourcedocsgen.generate_metadata(entry.repoSlug, entry.schemaFile, tag):
            raise SystemExit(f"resourcedocsgen failed for {name}@{tag}")
        published.append((name, tag))
    return published
