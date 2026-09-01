from __future__ import annotations

import json
import subprocess
from pathlib import Path
from typing import Any

from models import Entry

PATH = Path("community-packages/package-list.json")
PUBLISHER_NAMES_PATH = Path("tools/resourcedocsgen/pkg/publishers/publisher-names.json")
ALLOWED_PATHS = (PATH, PUBLISHER_NAMES_PATH)
EMPTY = '{"include":[]}'


def files_outside_allowlist(changed: list[str]) -> list[str]:
    allowed = {str(path) for path in ALLOWED_PATHS}
    return [f for f in changed if f and f not in allowed]


def _entries(document: Any) -> list[dict[str, str]]:
    if isinstance(document, dict):
        return document.get("include") or document.get("packages") or []
    return list(document)


def added_entries(before: str, after: str) -> list[Entry]:
    old = _entries(json.loads(before))
    new = _entries(json.loads(after))
    return [Entry(e["repoSlug"], e["schemaFile"]) for e in new if e not in old]


def at_ref(ref: str) -> str:
    shown = subprocess.run(["git", "show", f"{ref}:{PATH}"], capture_output=True, text=True).stdout
    return shown or EMPTY


def current() -> str:
    return PATH.read_text() if PATH.exists() else EMPTY
