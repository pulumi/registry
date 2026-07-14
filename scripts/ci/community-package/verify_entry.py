from __future__ import annotations

import hashlib
import json
import tempfile
from pathlib import Path
from typing import Any

import doc_lint
import github_api
import sdk_install_probe
import package_list
import resourcedocsgen
from models import Content, DocFile, Entry, Gate, Manifest, Owner, Version


def sha256(data: bytes) -> str:
    return hashlib.sha256(data).hexdigest()


def provider_name(slug: str, schema: dict[str, Any]) -> str:
    named = schema.get("name")
    if named:
        return str(named)
    repo = slug.split("/")[-1]
    return repo.split("-", 1)[-1] if "-" in repo else repo


def _doc_file_from_bytes(path: str, data: bytes | None) -> DocFile | None:
    if data is None:
        return None
    text = data.decode("utf-8", "replace")
    return DocFile(path, text.count("\n") + 1, text)


def _doc_file(slug: str, sha: str, path: str) -> DocFile | None:
    return _doc_file_from_bytes(path, github_api.raw_file(slug, sha, path))


def verify(entry: Entry) -> Manifest:
    tag = github_api.latest_release_tag(entry.repoSlug)
    sha = github_api.commit_sha_for_tag(entry.repoSlug, tag)

    schema_bytes = github_api.raw_file(entry.repoSlug, sha, entry.schemaFile)
    if schema_bytes is None:
        raise SystemExit(f"schema not found: {entry.repoSlug}@{sha}/{entry.schemaFile}")
    schema = json.loads(schema_bytes)
    name = provider_name(entry.repoSlug, schema)

    index_bytes = github_api.raw_file(entry.repoSlug, sha, "docs/_index.md")
    index = _doc_file_from_bytes("docs/_index.md", index_bytes)
    install_doc = _doc_file(entry.repoSlug, sha, "docs/installation-configuration.md")
    docs = [d for d in (index, install_doc) if d is not None]

    with tempfile.TemporaryDirectory() as scratch:
        generated = resourcedocsgen.generate_metadata(entry.repoSlug, entry.schemaFile, tag, into=Path(scratch))

    installs = sdk_install_probe.probe_installs(name, tag, schema)
    findings = doc_lint.find_issues(index.content if index else "")

    blocking_failures = sum(1 for r in installs if r.blocking and r.result != "pass")
    green = generated and blocking_failures == 0 and index is not None

    owner = entry.repoSlug.split("/")[0]
    return Manifest(
        repoSlug=entry.repoSlug,
        schemaFile=entry.schemaFile,
        providerName=name,
        version=Version(tag, sha),
        content=Content(sha256(schema_bytes), sha256(index_bytes) if index_bytes else ""),
        owner=Owner(owner, owner in package_list.listed_owners()),
        toolchain={"resourcedocsgen": resourcedocsgen.go_version()},
        installMatrix=installs,
        docLint={"findings": findings},
        gate=Gate(green, blocking_failures),
        docs=docs,
        generation=generated,
    )
