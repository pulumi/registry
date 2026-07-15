from __future__ import annotations

import json
import subprocess
from pathlib import Path
from typing import Any

from models import Entry

PATH = Path("community-packages/package-list.json")
EMPTY = '{"include":[]}'


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
