from __future__ import annotations

import os
import subprocess
from pathlib import Path

BINARY = Path(os.environ.get("RESOURCEDOCSGEN") or "tools/resourcedocsgen/resourcedocsgen")


def ensure_built() -> None:
    if not BINARY.exists():
        subprocess.run(["go", "build", "-C", "tools/resourcedocsgen"], check=True)


def generate_metadata(slug: str, schema_file: str, tag: str, into: Path | None = None) -> bool:
    args = [str(BINARY), "metadata", "from-github",
            "--repoSlug", slug, "--schemaFile", schema_file, "--version", tag]
    if into is not None:
        args += ["--metadataDir", str(into / "data"), "--packageDocsDir", str(into / "content")]
    return subprocess.run(args, capture_output=True, text=True).returncode == 0
