from __future__ import annotations

import json
from pathlib import Path

import github_api
import resourcedocsgen
from models import Entry, PackagePin
from verify_entry import provider_name, sha256

LOCKS = Path("community-packages/locks")


def pin_entry(entry: Entry) -> PackagePin:
    tag = github_api.latest_release_tag(entry.repoSlug)
    sha = github_api.commit_sha_for_tag(entry.repoSlug, tag)
    schema_bytes = github_api.raw_file(entry.repoSlug, sha, entry.schemaFile)
    if schema_bytes is None:
        raise FileNotFoundError(f"{entry.repoSlug}@{sha}/{entry.schemaFile}")
    return PackagePin(tag, sha, provider_name(entry.repoSlug, json.loads(schema_bytes)), sha256(schema_bytes))


def _is_tag_swap(lock: Path, pin: PackagePin) -> bool:
    if not lock.exists():
        return False
    previous = json.loads(lock.read_text())
    same_tag = previous["version"]["tag"] == pin.tag
    changed_content = previous["content"]["schemaSha256"] != pin.schemaSha256
    return bool(same_tag and changed_content)


def publish_entries(entries: list[Entry]) -> list[tuple[str, str]]:
    LOCKS.mkdir(parents=True, exist_ok=True)
    published: list[tuple[str, str]] = []
    for entry in entries:
        pin = pin_entry(entry)
        lock = LOCKS / f"{pin.providerName}.json"
        if _is_tag_swap(lock, pin):
            raise SystemExit(f"REFUSING: {pin.providerName}@{pin.tag} content changed since it was "
                             f"locked (the tag was moved).")
        lock.write_text(json.dumps({
            "repoSlug": entry.repoSlug, "schemaFile": entry.schemaFile, "providerName": pin.providerName,
            "version": {"tag": pin.tag, "commitSha": pin.commitSha},
            "content": {"schemaSha256": pin.schemaSha256},
        }, indent=2))
        if not resourcedocsgen.generate_metadata(entry.repoSlug, entry.schemaFile, pin.tag):
            raise SystemExit(f"resourcedocsgen failed for {pin.providerName}@{pin.tag}")
        published.append((pin.providerName, pin.tag))
    return published
