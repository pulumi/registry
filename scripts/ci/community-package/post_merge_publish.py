from __future__ import annotations

import json
from pathlib import Path

import github_api
import resourcedocsgen
from models import Entry
from verify_entry import provider_name, sha256

LOCKS = Path("community-packages/locks")


def pinned_facts(entry: Entry) -> tuple[str, str, str, str]:
    tag = github_api.latest_release_tag(entry.repoSlug)
    sha = github_api.commit_sha_for_tag(entry.repoSlug, tag)
    schema_bytes = github_api.raw_file(entry.repoSlug, sha, entry.schemaFile)
    if schema_bytes is None:
        raise FileNotFoundError(f"{entry.repoSlug}@{sha}/{entry.schemaFile}")
    return tag, sha, provider_name(entry.repoSlug, json.loads(schema_bytes)), sha256(schema_bytes)


def publish_entries(entries: list[Entry]) -> list[tuple[str, str]]:
    LOCKS.mkdir(parents=True, exist_ok=True)
    published: list[tuple[str, str]] = []
    for entry in entries:
        tag, sha, name, schema_sha = pinned_facts(entry)
        lock = LOCKS / f"{name}.json"
        if lock.exists():
            previous = json.loads(lock.read_text())
            if previous["version"]["tag"] == tag and previous["content"]["schemaSha256"] != schema_sha:
                raise SystemExit(f"REFUSING: {name}@{tag} content hash changed since it was locked (tag swap).")
        lock.write_text(json.dumps({
            "repoSlug": entry.repoSlug, "schemaFile": entry.schemaFile, "providerName": name,
            "version": {"tag": tag, "commitSha": sha}, "content": {"schemaSha256": schema_sha},
        }, indent=2))
        if not resourcedocsgen.generate_metadata(entry.repoSlug, entry.schemaFile, tag):
            raise SystemExit(f"resourcedocsgen failed for {name}@{tag}")
        published.append((name, tag))
    return published
