from __future__ import annotations

import json
import tempfile
from pathlib import Path

import yaml

import doc_lint
import github_api
import sdk_install_probe
import resourcedocsgen
from models import DocFile, Entry, Manifest, Version, provider_name

PUBLISHER_NAMES_PATH = Path("tools/resourcedocsgen/pkg/publishers/publisher-names.json")

DELISTED_PUBLISHER = "DEPRECATED"


def _load_publisher_names() -> dict[str, str]:
    try:
        return dict(json.loads(PUBLISHER_NAMES_PATH.read_text()))
    except OSError:
        return {}


def _publisher_known(publisher: str, names: dict[str, str]) -> bool:
    return not publisher or publisher in names


def _doc_file(slug: str, sha: str, path: str) -> DocFile | None:
    data = github_api.raw_file(slug, sha, path)
    if data is None:
        return None
    text = data.decode("utf-8", "replace")
    return DocFile(path, text.count("\n") + 1, text)


def _unverifiable(entry: Entry, reason: str) -> Manifest:
    return Manifest(
        repoSlug=entry.repoSlug,
        schemaFile=entry.schemaFile,
        providerName=provider_name(entry.repoSlug, {}),
        version=Version("(none)", ""),
        owner=entry.repoSlug.split("/")[0],
        installMatrix=[],
        docLint=[],
        green=False,
        generation=False,
        docs=[],
        error=reason,
    )


def verify(entry: Entry) -> Manifest:
    tag = github_api.latest_release_tag(entry.repoSlug)
    if tag is None:
        return _unverifiable(entry, f"`{entry.repoSlug}` has no published GitHub release. The check "
                                    "reads the package at its latest release; publish a release (not "
                                    "just a git tag or a prerelease) and re-run `/check`.")
    sha = github_api.commit_sha_for_tag(entry.repoSlug, tag)

    schema_bytes = github_api.raw_file(entry.repoSlug, sha, entry.schemaFile)
    if schema_bytes is None:
        return _unverifiable(entry, f"Schema not found at `{entry.schemaFile}` in `{entry.repoSlug}` "
                                    f"at release `{tag}`. Check the `schemaFile` path in the package list.")
    schema = json.loads(schema_bytes)
    name = provider_name(entry.repoSlug, schema)
    publisher = str(schema.get("publisher", "")).strip()
    publisher_known = _publisher_known(publisher, _load_publisher_names())

    index = _doc_file(entry.repoSlug, sha, "docs/_index.md")
    install_doc = _doc_file(entry.repoSlug, sha, "docs/installation-configuration.md")
    docs = [d for d in (index, install_doc) if d is not None]

    with tempfile.TemporaryDirectory() as scratch:
        generated = resourcedocsgen.generate_metadata(entry.repoSlug, entry.schemaFile, tag, into=Path(scratch))

    installs = sdk_install_probe.probe_installs(name, tag, schema)
    findings = doc_lint.find_issues(index.content if index else "")

    has_blocking_failure = any(r.blocking and r.result != "pass" for r in installs)
    green = bool(generated and not has_blocking_failure and index is not None)
    advisory_failure = any(not r.blocking and r.result not in ("pass", "absent") for r in installs)
    warnings = green and (advisory_failure or bool(findings) or (bool(publisher) and not publisher_known))

    return Manifest(
        repoSlug=entry.repoSlug,
        schemaFile=entry.schemaFile,
        providerName=name,
        version=Version(tag, sha),
        owner=entry.repoSlug.split("/")[0],
        installMatrix=installs,
        docLint=findings,
        green=green,
        warnings=warnings,
        generation=generated,
        docs=docs,
        publisher=publisher,
        publisherKnown=publisher_known,
    )


def _unverifiable_package(name: str, yaml_path: str, reason: str) -> Manifest:
    return Manifest(
        repoSlug="",
        schemaFile=yaml_path,
        providerName=name,
        version=Version("(none)", ""),
        owner="",
        installMatrix=[],
        docLint=[],
        green=False,
        generation=False,
        docs=[],
        error=reason,
    )


def _delisted(name: str, yaml_path: str, version: str) -> Manifest:
    return Manifest(
        repoSlug="",
        schemaFile=yaml_path,
        providerName=name,
        version=Version(version, ""),
        owner="",
        installMatrix=[],
        docLint=[],
        green=True,
        generation=True,
        docs=[],
        publisher=DELISTED_PUBLISHER,
        delisted=True,
    )


def verify_package_yaml(yaml_path: Path, repo_root: Path) -> Manifest:
    """Verify a package from the registry PR that publishes it.

    Every publish path ends in a PR that writes the package YAML and its `_index.md`,
    so both files are in the tree and this runs without a single network call.

    The install probes are deliberately absent. `pulumi package add terraform-provider`
    runs the upstream provider binary to read its schema, and this job runs in a
    workflow that holds PULUMI_BOT_TOKEN, so probing here would put third-party code
    next to a write token.
    """
    name = yaml_path.stem
    try:
        data = yaml.safe_load(yaml_path.read_text()) or {}
    except (OSError, yaml.YAMLError) as exc:
        return _unverifiable_package(name, str(yaml_path), f"`{yaml_path}` could not be read: {exc}.")

    name = str(data.get("name") or name).strip()
    version = str(data.get("version", "")).strip()
    if not version:
        return _unverifiable_package(name, str(yaml_path), f"`{yaml_path}` declares no `version`.")

    publisher = str(data.get("publisher", "")).strip()
    if publisher == DELISTED_PUBLISHER:
        return _delisted(name, str(yaml_path), version)

    publisher_known = _publisher_known(publisher, _load_publisher_names())

    index_path = repo_root / "themes/default/content/registry/packages" / name / "_index.md"
    try:
        text = index_path.read_text()
    except OSError:
        return _unverifiable_package(name, str(yaml_path),
                                     f"No landing page at `{index_path.relative_to(repo_root)}`. It is the "
                                     "package's front page in the registry, so it is required.")
    index = DocFile("_index.md", text.count("\n") + 1, text)

    findings = doc_lint.find_issues(index.content)

    green = bool(publisher_known)
    warnings = green and bool(findings)

    return Manifest(
        repoSlug="",
        schemaFile=str(data.get("schema_file_url") or yaml_path),
        providerName=name,
        version=Version(version, ""),
        owner=publisher,
        installMatrix=[],
        docLint=findings,
        green=green,
        warnings=warnings,
        generation=True,
        docs=[index],
        publisher=publisher,
        publisherKnown=publisher_known,
        docSourceURL="",
    )
