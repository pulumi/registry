from __future__ import annotations

import os
import re
import subprocess
import sys
import tempfile
from typing import Any, Callable

from models import InstallResult

# The provider supplies the package name, version, and plugin URL. We validate each against a
# strict allowlist and pass it as an argv argument, never through a shell, so a crafted name
# cannot inject a command.
SAFE_NAME = re.compile(r"^[A-Za-z0-9._@/-]+$")
SAFE_VERSION = re.compile(r"^[A-Za-z0-9._+-]+$")
SAFE_URL = re.compile(r"^https://[A-Za-z0-9._~:/?#@!$&()*+,;=-]+$")


def _run(cmd: list[str], cwd: str | None = None, env: dict[str, str] | None = None) -> tuple[bool, str]:
    try:
        subprocess.run(cmd, check=True, capture_output=True, text=True, timeout=300, cwd=cwd, env=env)
        return True, ""
    except subprocess.CalledProcessError as e:
        return False, ((e.stderr or "") + (e.stdout or "")).strip()
    except subprocess.TimeoutExpired:
        return False, f"timed out after 300s: {' '.join(cmd)}"
    except FileNotFoundError:
        return False, f"command not found: {cmd[0]}"


def _go_module_resolves(import_path: str, tag: str) -> tuple[bool, str]:
    # An SDK is a library, so `go install` won't work. Resolve it inside a throwaway module
    # with `go get`, which downloads and records it without building a main package.
    with tempfile.TemporaryDirectory() as td:
        ok, err = _run(["go", "mod", "init", "probe"], cwd=td)
        if not ok:
            return False, err
        return _run(["go", "get", "--", f"{import_path}@{tag}"], cwd=td,
                    env={**os.environ, "GOFLAGS": "-mod=mod"})


def probe_installs(name: str, tag: str, schema: dict[str, Any]) -> list[InstallResult]:
    version = tag[1:] if tag.startswith("v") else tag
    languages = schema.get("language", {})
    results: list[InstallResult] = []

    def record(language: str, command: str, ok: bool, error: str, blocking: bool = False) -> None:
        if not ok and error:  # echo to the Actions log so the run link holds the evidence
            print(f"::group::install {language} FAILED, {command}", file=sys.stderr)
            print(error, file=sys.stderr)
            print("::endgroup::", file=sys.stderr)
        results.append(InstallResult(language, command, "pass" if ok else "fail",
                                     error="" if ok else error[-600:], blocking=blocking))

    if SAFE_NAME.match(name) and SAFE_VERSION.match(tag):
        command = ["pulumi", "plugin", "install", "resource", name, tag]
        url = schema.get("pluginDownloadURL", "")
        if url and SAFE_URL.match(url):
            command += ["--server", url]
        ok, error = _run(command)
        record("plugin", f"pulumi plugin install resource {name} {tag}", ok, error, blocking=True)
    else:
        results.append(InstallResult("plugin", "(rejected: unsafe identifier)", "rejected", blocking=True))

    def probe(language: str, package: str | None, runner: Callable[[], tuple[bool, str]], command: str) -> None:
        if not package:
            return
        if not SAFE_NAME.match(package):
            results.append(InstallResult(language, "(rejected: unsafe name)", "rejected"))
            return
        ok, error = runner()
        record(language, command, ok, error)

    npm_package = languages.get("nodejs", {}).get("packageName")
    probe("nodejs", npm_package,
          lambda: _run(["npm", "install", "--ignore-scripts", "--no-audit", "--no-fund",
                        "--prefix", "/tmp/nn", "--", f"{npm_package}@{version}"]),
          f"npm install {npm_package}@{version}")

    pypi_package = languages.get("python", {}).get("packageName") or f"pulumi_{schema.get('name', '')}"
    probe("python", pypi_package,
          lambda: _run([sys.executable, "-m", "pip", "download", "--no-deps", "--only-binary", ":all:",
                        "--dest", "/tmp/py", "--", f"{pypi_package}=={version}"]),
          f"pip download {pypi_package}=={version}")

    go_import = languages.get("go", {}).get("importBasePath")
    probe("go", go_import, lambda: _go_module_resolves(go_import, tag), f"go get {go_import}@{tag}")

    return results
