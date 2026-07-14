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
    generation: bool = True
    docs: list[DocFile] = field(default_factory=list)
