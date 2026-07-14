from __future__ import annotations

from dataclasses import dataclass, field
from typing import Literal

InstallOutcome = Literal["pass", "fail", "absent", "rejected"]


@dataclass(frozen=True)
class Entry:
    repoSlug: str
    schemaFile: str


@dataclass
class Version:
    tag: str
    commitSha: str


@dataclass
class Content:
    schemaSha256: str
    docsSha256: str


@dataclass
class PackagePin:
    tag: str
    commitSha: str
    providerName: str
    schemaSha256: str


@dataclass
class Owner:
    handle: str
    alreadyListed: bool


@dataclass
class InstallResult:
    language: str
    command: str
    result: InstallOutcome
    hash: str = ""
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
class Gate:
    verificationGreen: bool
    blockingFailures: int
    requiresMaintainerApproval: bool = True


@dataclass
class Manifest:
    repoSlug: str
    schemaFile: str
    providerName: str
    version: Version
    content: Content
    owner: Owner
    toolchain: dict[str, str]
    installMatrix: list[InstallResult]
    docLint: dict[str, list[DocFinding]]
    gate: Gate
    docs: list[DocFile] = field(default_factory=list)
    generation: bool = True
