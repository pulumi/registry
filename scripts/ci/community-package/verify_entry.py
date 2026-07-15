from __future__ import annotations

import json
import tempfile
from pathlib import Path

import doc_lint
import github_api
import sdk_install_probe
import resourcedocsgen
from models import DocFile, Entry, Manifest, Version, provider_name


def _doc_file(slug: str, sha: str, path: str) -> DocFile | None:
    data = github_api.raw_file(slug, sha, path)
    if data is None:
        return None
    text = data.decode("utf-8", "replace")
    return DocFile(path, text.count("\n") + 1, text)


def verify(entry: Entry) -> Manifest:
    tag = github_api.latest_release_tag(entry.repoSlug)
    sha = github_api.commit_sha_for_tag(entry.repoSlug, tag)

    schema_bytes = github_api.raw_file(entry.repoSlug, sha, entry.schemaFile)
    if schema_bytes is None:
        raise SystemExit(f"schema not found: {entry.repoSlug}@{sha}/{entry.schemaFile}")
    schema = json.loads(schema_bytes)
    name = provider_name(entry.repoSlug, schema)

    index = _doc_file(entry.repoSlug, sha, "docs/_index.md")
    install_doc = _doc_file(entry.repoSlug, sha, "docs/installation-configuration.md")
    docs = [d for d in (index, install_doc) if d is not None]

    with tempfile.TemporaryDirectory() as scratch:
        generated = resourcedocsgen.generate_metadata(entry.repoSlug, entry.schemaFile, tag, into=Path(scratch))

    installs = sdk_install_probe.probe_installs(name, tag, schema)
    findings = doc_lint.find_issues(index.content if index else "")

    has_blocking_failure = any(r.blocking and r.result != "pass" for r in installs)
    green = generated and not has_blocking_failure and index is not None

    return Manifest(
        repoSlug=entry.repoSlug,
        schemaFile=entry.schemaFile,
        providerName=name,
        version=Version(tag, sha),
        owner=entry.repoSlug.split("/")[0],
        installMatrix=installs,
        docLint=findings,
        green=green,
        generation=generated,
        docs=docs,
    )
