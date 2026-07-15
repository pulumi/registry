from __future__ import annotations

from dataclasses import dataclass, field
from typing import Any, Literal

InstallOutcome = Literal["pass", "fail", "absent", "rejected"]


def provider_name(slug: str, schema: dict[str, Any]) -> str:
    named = schema.get("name")
    if named:
        return str(named)
    repo = slug.split("/")[-1]
    return repo.split("-", 1)[-1] if "-" in repo else repo


@dataclass(frozen=True)
class Entry:
    repoSlug: str
    schemaFile: str


@dataclass
class Version:
    tag: str
    commitSha: str


@dataclass
class InstallResult:
    language: str
    command: str
    result: InstallOutcome
    blocking: bool = False
    error: str = ""


@dataclass
class DocFinding:
    line: int
    kind: str
    text: str


@dataclass
class DocFile:
    path: str
    lines: int
    content: str = ""


@dataclass
class Manifest:
    repoSlug: str
    schemaFile: str
    providerName: str
    version: Version
    owner: str
    installMatrix: list[InstallResult]
    docLint: list[DocFinding]
    green: bool
    warnings: bool = False
    generation: bool = True
    docs: list[DocFile] = field(default_factory=list)
    error: str = ""
