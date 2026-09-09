from __future__ import annotations

import os
import re
import subprocess
import sys
import tempfile
import urllib.error
import urllib.request
from typing import Any, Callable

from models import InstallResult

SAFE_NAME = re.compile(r"^[A-Za-z0-9._@/-]+$")
SAFE_VERSION = re.compile(r"^[A-Za-z0-9._+-]+$")
SAFE_URL = re.compile(r"^(https|github|gitlab)://[A-Za-z0-9._~:/?#@!$&()*+,;=-]+$")


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


def _pypi_version_exists(package: str, version: str) -> bool:
    try:
        with urllib.request.urlopen(f"https://pypi.org/pypi/{package}/{version}/json", timeout=30) as resp:
            return bool(resp.status == 200)
    except urllib.error.URLError:
        return False


NEVER_BUILD_FROM_SOURCE = ["--only-binary", ":all:"]


def _python_resolves(package: str, version: str) -> tuple[bool, str]:
    ok, err = _run([sys.executable, "-m", "pip", "download", "--no-deps", *NEVER_BUILD_FROM_SOURCE,
                    "--dest", "/tmp/py", "--", f"{package}=={version}"])
    if ok or _pypi_version_exists(package, version):
        return True, ""
    return False, err


SAFE_COORDINATE = re.compile(r"^[A-Za-z0-9._-]+$")
MAVEN_CENTRAL = "https://repo1.maven.org/maven2"


def _maven_artifact_exists(base_package: str, artifact: str, version: str) -> tuple[bool, str]:
    group = base_package.replace(".", "/")
    url = f"{MAVEN_CENTRAL}/{group}/{artifact}/{version}/"
    try:
        with urllib.request.urlopen(url, timeout=30) as resp:
            if resp.status == 200:
                return True, ""
    except urllib.error.HTTPError as e:
        return False, (f"{url} returned {e.code}. The schema advertises a Java SDK at "
                       f"{base_package}:{artifact}, but Maven Central does not carry that version. "
                       "Publish it there, or drop the `java` block from the schema.")
    except urllib.error.URLError as e:
        return False, f"could not reach Maven Central: {e.reason}"
    return False, f"{url} did not return 200"


def _go_module_resolves(import_path: str, tag: str) -> tuple[bool, str]:
    with tempfile.TemporaryDirectory() as td:
        ok, err = _run(["go", "mod", "init", "probe"], cwd=td)
        if not ok:
            return False, err
        return _run(["go", "get", "--", f"{import_path}@{tag}"], cwd=td,
                    env={**os.environ, "GOFLAGS": "-mod=mod"})


def _echo_to_the_run_log(heading: str, body: str) -> None:
    print(f"::group::{heading}", file=sys.stderr)
    print(body, file=sys.stderr)
    print("::endgroup::", file=sys.stderr)


def _rejected(language: str, kind: str, value: str, allowed: str,
              blocking: bool = False) -> InstallResult:
    shown = repr(value[:80]).replace("`", "'")
    cell = shown.replace("|", r"\|")
    return InstallResult(
        language, f"(not run: {kind} is {cell})", "rejected", blocking=blocking,
        error=f"The schema gives {kind} as `{shown}`. The check builds install commands from "
              f"strings in your schema, so it accepts only {allowed}, and skips the probe instead "
              f"of passing anything else to a shell. Fix the value in your schema, or say so on "
              f"this PR if you believe it is correct.")


def probe_installs(name: str, tag: str, schema: dict[str, Any]) -> list[InstallResult]:
    if not SAFE_VERSION.match(tag):
        return [_rejected("plugin", "the release tag", tag,
                          "letters, digits, and . _ + -", blocking=True)]
    version = tag[1:] if tag.startswith("v") else tag
    languages = schema.get("language", {})
    results: list[InstallResult] = []

    def record(language: str, command: str, ok: bool, error: str, blocking: bool = False) -> None:
        if not ok and error:
            _echo_to_the_run_log(f"install {language} FAILED, {command}", error)
        results.append(InstallResult(language, command, "pass" if ok else "fail",
                                     error="" if ok else error[-600:], blocking=blocking))

    if SAFE_NAME.match(name):
        command = ["pulumi", "plugin", "install", "resource", name, tag]
        url = schema.get("pluginDownloadURL", "")
        if url and SAFE_URL.match(url):
            command += ["--server", url]
        ok, error = _run(command)
        record("plugin", f"pulumi plugin install resource {name} {tag}", ok, error, blocking=True)
    else:
        results.append(_rejected("plugin", "the provider name", name,
                                 "letters, digits, and . _ @ / -", blocking=True))

    def probe(language: str, kind: str, package: str | None,
              runner: Callable[[], tuple[bool, str]], command: str) -> None:
        if not package:
            return
        if not SAFE_NAME.match(package):
            results.append(_rejected(language, kind, package, "letters, digits, and . _ @ / -"))
            return
        ok, error = runner()
        record(language, command, ok, error)

    npm_package = languages.get("nodejs", {}).get("packageName")
    probe("nodejs", "the nodejs packageName", npm_package,
          lambda: _run(["npm", "install", "--ignore-scripts", "--no-audit", "--no-fund",
                        "--prefix", "/tmp/nn", "--", f"{npm_package}@{version}"]),
          f"npm install {npm_package}@{version}")

    python = languages.get("python")
    pypi_package = (python.get("packageName") or f"pulumi_{schema.get('name', '')}") if python is not None else None
    python_kind = ("the python packageName" if (python or {}).get("packageName")
                   else "the python package name, taken from the schema name because python "
                        "advertises no packageName")
    probe("python", python_kind, pypi_package,
          lambda: _python_resolves(pypi_package or "", version),
          f"pip download {pypi_package}=={version}")

    go_import = languages.get("go", {}).get("importBasePath")
    probe("go", "the go importBasePath", go_import,
          lambda: _go_module_resolves(go_import, tag), f"go get {go_import}@{tag}")

    base_package = (languages.get("java") or {}).get("basePackage") or ""
    if base_package and SAFE_COORDINATE.match(base_package) and SAFE_COORDINATE.match(name):
        ok, error = _maven_artifact_exists(base_package, name, version)
        record("java", f"{base_package}:{name}:{version} on Maven Central", ok, error)
    elif base_package:
        results.append(InstallResult("java", "(rejected: unsafe coordinate)", "rejected"))

    if "csharp" in languages:
        results.append(InstallResult("dotnet", "(no package id in the schema to check)", "absent"))

    return results
