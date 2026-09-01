from __future__ import annotations

import os
import subprocess
import sys
from pathlib import Path

BINARY = Path(os.environ.get("RESOURCEDOCSGEN") or "tools/resourcedocsgen/resourcedocsgen")


def ensure_built() -> None:
    if not BINARY.exists():
        subprocess.run(["go", "build", "-C", "tools/resourcedocsgen"], check=True)


def generate_metadata(slug: str, schema_file: str, tag: str, provider_name: str,
                      into: Path | None = None) -> tuple[bool, str]:
    args = [str(BINARY), "metadata", "from-github",
            "--repoSlug", slug, "--schemaFile", schema_file, "--version", tag,
            "--providerName", provider_name]
    if into is not None:
        args += ["--metadataDir", str(into / "data"), "--packageDocsDir", str(into / "content")]
    run = subprocess.run(args, capture_output=True, text=True)
    if run.returncode == 0:
        return True, ""
    output = ((run.stderr or "") + (run.stdout or "")).strip()
    print(f"::group::docs generate FAILED, resourcedocsgen metadata from-github {slug}@{tag}", file=sys.stderr)
    print(output, file=sys.stderr)
    print("::endgroup::", file=sys.stderr)
    return False, output[:600]
