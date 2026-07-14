#!/usr/bin/env python3
"""Data-processing core of the community package pipeline.

The JSON-assembly steps (verify, install-probe, doc-lint, manifest) live here in typed
Python instead of bash + jq, which repeatedly mis-assembled JSON. The thin CLI glue
(the comment commands and the secret-vs-code separation check) stays in bash.

Subcommands: check (verify added entries, on the PR), publish (post-merge on master),
docs (pre-fetch a provider's docs for /review). stdlib only; no third-party deps.
"""
from __future__ import annotations

import argparse
import hashlib
import json
import os
import re
import subprocess
import sys
import tempfile
import urllib.error
import urllib.request
from dataclasses import asdict, dataclass, field
from pathlib import Path
from typing import Any, Callable, Literal

PACKAGE_LIST = Path("community-packages/package-list.json")
RESOURCEDOCSGEN = Path(os.environ.get("RESOURCEDOCSGEN") or "tools/resourcedocsgen/resourcedocsgen")

Result = Literal["pass", "fail", "absent", "rejected"]

VALID_NAME = re.compile(r"^[A-Za-z0-9._@/-]+$")
VALID_VER = re.compile(r"^[A-Za-z0-9._+-]+$")
VALID_URL = re.compile(r"^https://[A-Za-z0-9._~:/?#@!$&()*+,;=-]+$")


# ── models ──────────────────────────────────────────────────────────────────
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
class Owner:
    handle: str
    alreadyListed: bool


@dataclass
class InstallResult:
    language: str
    command: str
    result: Result
    hash: str = ""
    blocking: bool = False
    error: str = ""


@dataclass
class DocFinding:
    line: int
    kind: str
    text: str


@dataclass
class DocRef:
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
    docs: list[DocRef] = field(default_factory=list)
    generation: bool = True


# ── github / raw ────────────────────────────────────────────────────────────
def _fetch(url: str, token: str | None) -> bytes:
    req = urllib.request.Request(url)
    if token:
        req.add_header("Authorization", f"Bearer {token}")
    with urllib.request.urlopen(req, timeout=30) as resp:
        return bytes(resp.read())


def gh_api(path: str) -> Any:
    return json.loads(_fetch(f"https://api.github.com{path}", os.environ.get("GITHUB_TOKEN")))


def raw(slug: str, ref: str, path: str) -> bytes | None:
    try:
        return _fetch(f"https://raw.githubusercontent.com/{slug}/{ref}/{path}", None)
    except urllib.error.HTTPError:
        return None


def latest_tag(slug: str) -> str:
    return str(gh_api(f"/repos/{slug}/releases/latest")["tag_name"])


def resolve_commit_sha(slug: str, tag: str) -> str:
    ref = gh_api(f"/repos/{slug}/git/refs/tags/{tag}")["object"]
    if ref["type"] == "tag":  # annotated tag → dereference to the commit
        return str(gh_api(f"/repos/{slug}/git/tags/{ref['sha']}")["object"]["sha"])
    return str(ref["sha"])


# ── entries ─────────────────────────────────────────────────────────────────
def _items(doc: Any) -> list[dict[str, str]]:
    if isinstance(doc, dict):
        return doc.get("include") or doc.get("packages") or []
    return list(doc)


def diff_entries(before: str, after: str) -> list[Entry]:
    b = _items(json.loads(before))
    a = _items(json.loads(after))
    return [Entry(e["repoSlug"], e["schemaFile"]) for e in a if e not in b]


def listed_owners() -> set[str]:
    if not PACKAGE_LIST.exists():
        return set()
    return {e["repoSlug"].split("/")[0] for e in _items(json.loads(PACKAGE_LIST.read_text()))}


# ── doc lint ────────────────────────────────────────────────────────────────
_LINT = [
    (re.compile(r'<img[^>]+src="(\.|/|[^h])'), "relative-image"),
    (re.compile(r"<img[^>]+src='(\.|/|[^h])"), "relative-image"),
    (re.compile(r"!\[[^\]]*\]\((\.|/)"), "relative-image"),
    (re.compile(r'<a[^>]+href="(\.|/)'), "raw-relative-link"),
]


def lint_docs(text: str) -> list[DocFinding]:
    out: list[DocFinding] = []
    for i, line in enumerate(text.splitlines(), 1):
        for rx, kind in _LINT:
            if rx.search(line):
                out.append(DocFinding(i, kind, line[:160]))
                break
    return out


# ── install probe ───────────────────────────────────────────────────────────
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


def _go_resolves(import_path: str, tag: str) -> tuple[bool, str]:
    # An SDK is a library, so `go install` won't work. Resolve it inside a throwaway
    # module with `go get`, which downloads and records it without building a main package.
    with tempfile.TemporaryDirectory() as td:
        ok, err = _run(["go", "mod", "init", "probe"], cwd=td)
        if not ok:
            return False, err
        return _run(["go", "get", "--", f"{import_path}@{tag}"], cwd=td,
                    env={**os.environ, "GOFLAGS": "-mod=mod"})


def install_verify(name: str, tag: str, schema: dict[str, Any]) -> list[InstallResult]:
    version = tag[1:] if tag.startswith("v") else tag
    lang = schema.get("language", {})
    out: list[InstallResult] = []

    def record(language: str, display: str, ok: bool, err: str, blocking: bool = False) -> None:
        if not ok and err:  # echo to the Actions log so the run link contains the evidence
            print(f"::group::install {language} FAILED, {display}", file=sys.stderr)
            print(err, file=sys.stderr)
            print("::endgroup::", file=sys.stderr)
        out.append(InstallResult(language, display, "pass" if ok else "fail",
                                 error="" if ok else err[-600:], blocking=blocking))

    if VALID_NAME.match(name) and VALID_VER.match(tag):
        cmd = ["pulumi", "plugin", "install", "resource", name, tag]
        url = schema.get("pluginDownloadURL", "")
        if url and VALID_URL.match(url):
            cmd += ["--server", url]
        ok, err = _run(cmd)
        record("plugin", f"pulumi plugin install resource {name} {tag}", ok, err, blocking=True)
    else:
        out.append(InstallResult("plugin", "(rejected: unsafe identifier)", "rejected", blocking=True))

    def probe(language: str, pkg: str | None, runner: Callable[[], tuple[bool, str]], display: str) -> None:
        if not pkg:
            return
        if not VALID_NAME.match(pkg):
            out.append(InstallResult(language, "(rejected: unsafe name)", "rejected"))
            return
        ok, err = runner()
        record(language, display, ok, err)

    npm = lang.get("nodejs", {}).get("packageName")
    probe("nodejs", npm,
          lambda: _run(["npm", "install", "--ignore-scripts", "--no-audit", "--no-fund",
                        "--prefix", "/tmp/nn", "--", f"{npm}@{version}"]),
          f"npm install {npm}@{version}")

    py = lang.get("python", {}).get("packageName") or f"pulumi_{schema.get('name', '')}"
    probe("python", py,
          lambda: _run([sys.executable, "-m", "pip", "download", "--no-deps", "--only-binary", ":all:",
                        "--dest", "/tmp/py", "--", f"{py}=={version}"]),
          f"pip download {py}=={version}")

    go = lang.get("go", {}).get("importBasePath")
    probe("go", go, lambda: _go_resolves(go, tag), f"go get {go}@{tag}")

    return out


# ── verify one entry ────────────────────────────────────────────────────────
def _sha256(b: bytes) -> str:
    return hashlib.sha256(b).hexdigest()


def ensure_resourcedocsgen() -> None:
    if not RESOURCEDOCSGEN.exists():
        subprocess.run(["go", "build", "-C", "tools/resourcedocsgen"], check=True)


def provider_name(slug: str, schema: dict[str, Any]) -> str:
    name = schema.get("name")
    if name:
        return str(name)
    repo = slug.split("/")[-1]
    return repo.split("-", 1)[-1] if "-" in repo else repo


def verify(entry: Entry) -> Manifest:
    tag = latest_tag(entry.repoSlug)
    sha = resolve_commit_sha(entry.repoSlug, tag)

    schema_bytes = raw(entry.repoSlug, sha, entry.schemaFile)
    if schema_bytes is None:
        raise SystemExit(f"schema not found: {entry.repoSlug}@{sha}/{entry.schemaFile}")
    schema = json.loads(schema_bytes)
    name = provider_name(entry.repoSlug, schema)

    index_bytes = raw(entry.repoSlug, sha, "docs/_index.md") or b""
    index_text = index_bytes.decode("utf-8", "replace")

    docs: list[DocRef] = []
    if index_bytes:
        docs.append(DocRef("docs/_index.md", index_text.count("\n") + 1, index_text))
    inst_bytes = raw(entry.repoSlug, sha, "docs/installation-configuration.md")
    if inst_bytes is not None:
        inst_text = inst_bytes.decode("utf-8", "replace")
        docs.append(DocRef("docs/installation-configuration.md", inst_text.count("\n") + 1, inst_text))

    with tempfile.TemporaryDirectory() as td:
        gen = subprocess.run(
            [str(RESOURCEDOCSGEN), "metadata", "from-github",
             "--repoSlug", entry.repoSlug, "--schemaFile", entry.schemaFile, "--version", tag,
             "--metadataDir", f"{td}/data", "--packageDocsDir", f"{td}/content"],
            capture_output=True, text=True)
        generation_ok = gen.returncode == 0

    matrix = install_verify(name, tag, schema)
    findings = lint_docs(index_text)

    blocking_fails = sum(1 for m in matrix if m.blocking and m.result != "pass")
    green = generation_ok and blocking_fails == 0 and index_bytes != b""

    return Manifest(
        repoSlug=entry.repoSlug,
        schemaFile=entry.schemaFile,
        providerName=name,
        version=Version(tag, sha),
        content=Content(_sha256(schema_bytes), _sha256(index_bytes) if index_bytes else ""),
        owner=Owner(entry.repoSlug.split("/")[0], entry.repoSlug.split("/")[0] in listed_owners()),
        toolchain={"resourcedocsgen": _go_version()},
        installMatrix=matrix,
        docLint={"findings": findings},
        gate=Gate(green, blocking_fails),
        docs=docs,
        generation=generation_ok,
    )


def _go_version() -> str:
    try:
        return subprocess.run(["go", "version"], capture_output=True, text=True).stdout.split()[2]
    except Exception:
        return "unknown"


# ── after merge: drift + publish (no installs or doc generation) ──────────
LOCKS = Path("community-packages/locks")


def content_hashes(entry: Entry) -> tuple[str, str, str, str]:
    """(tag, sha, providerName, schemaSha256), the cheap facts, no installs or docs gen."""
    tag = latest_tag(entry.repoSlug)
    sha = resolve_commit_sha(entry.repoSlug, tag)
    schema_bytes = raw(entry.repoSlug, sha, entry.schemaFile)
    if schema_bytes is None:
        raise FileNotFoundError(f"{entry.repoSlug}@{sha}/{entry.schemaFile}")
    return tag, sha, provider_name(entry.repoSlug, json.loads(schema_bytes)), _sha256(schema_bytes)


def publish(entries: list[Entry]) -> list[tuple[str, str]]:
    LOCKS.mkdir(parents=True, exist_ok=True)
    published: list[tuple[str, str]] = []
    for e in entries:
        tag, sha, name, schema_sha = content_hashes(e)
        lock = LOCKS / f"{name}.json"
        if lock.exists():
            prev = json.loads(lock.read_text())
            if prev["version"]["tag"] == tag and prev["content"]["schemaSha256"] != schema_sha:
                raise SystemExit(f"REFUSING: {name}@{tag} content hash changed since it was locked (tag swap).")
        lock.write_text(json.dumps({
            "repoSlug": e.repoSlug, "schemaFile": e.schemaFile, "providerName": name,
            "version": {"tag": tag, "commitSha": sha}, "content": {"schemaSha256": schema_sha},
        }, indent=2))
        subprocess.run([str(RESOURCEDOCSGEN), "metadata", "from-github",
                        "--repoSlug", e.repoSlug, "--schemaFile", e.schemaFile, "--version", tag], check=True)
        published.append((name, tag))
    return published


# ── fact-sheet rendering ────────────────────────────────────────────────────
def _run_link() -> str | None:
    rid = os.environ.get("GITHUB_RUN_ID")
    if not rid:
        return None
    server = os.environ.get("GITHUB_SERVER_URL", "https://github.com")
    repo = os.environ.get("GITHUB_REPOSITORY", "")
    return f"{server}/{repo}/actions/runs/{rid}"


def _fence_for(text: str) -> str:
    longest = cur = 0
    for ch in text:
        cur = cur + 1 if ch == "`" else 0
        longest = max(longest, cur)
    return "`" * max(3, longest + 1)


_RESULT_ICON = {"pass": "✅", "fail": "❌", "absent": "➖", "rejected": "🚫"}


def _verdict(m: Manifest) -> tuple[str, str]:
    if m.gate.verificationGreen:
        return "✅", ("**Ready for approval.** All blocking checks passed; a `@pulumi/iac-cloud` "
                     "review is still required to merge.")
    return "❌", ("**Not ready.** A blocking check failed (see below). Fix upstream, then comment "
                 "`/check` to re-run.")


def factsheet(m: Manifest) -> str:
    icon, verdict = _verdict(m)
    commit_url = f"https://github.com/{m.repoSlug}/commit/{m.version.commitSha}"
    owner_note = "already listed" if m.owner.alreadyListed else "🆕 first-time contributor"

    lines = [
        f"## {icon} `{m.providerName}` · community package check",
        "",
        f"`{m.repoSlug}` · `{m.version.tag}` · commit [`{m.version.commitSha[:12]}`]({commit_url})",
        "",
        verdict,
        "",
        "| check | | command |",
        "|---|:-:|---|",
        f"| docs generate | {'✅' if m.generation else '❌'} | `resourcedocsgen metadata from-github` |",
    ]
    for r in m.installMatrix:
        label = "plugin install" if r.language == "plugin" else f"{r.language} SDK"
        cmd = f"`{r.command}`" if r.command else "_not advertised_"
        lines.append(f"| {label} | {_RESULT_ICON.get(r.result, r.result)} | {cmd} |")
    lines += ["", f"Owner `{m.owner.handle}` · {owner_note} · schema `sha256:{m.content.schemaSha256[:12]}…`"]

    findings = m.docLint["findings"]
    lines.append("")
    if findings:
        lines.append(f"**Doc-lint: {len(findings)} finding(s)** (advisory; these break the registry render surfaces):")
        lines += [f"- `docs/_index.md:{f.line}` {f.kind}: `{f.text}`" for f in findings]
    else:
        lines.append("**Doc-lint** ✅ clean. No relative images or raw HTML.")

    if m.docs:
        # Source embedded, not linked: GitHub's permalink snippet preview only renders for
        # same-repo links in a PR body, not for cross-repo issue comments, so we paste it.
        lines += ["", "**Provider docs** (source at the reviewed commit):"]
        for d in m.docs:
            url = f"https://github.com/{m.repoSlug}/blob/{m.version.commitSha}/{d.path}"
            body, note = d.content, ""
            if d.lines > 400:
                body = "\n".join(d.content.splitlines()[:400])
                note = f"\n… truncated to 400 of {d.lines} lines"
            fence = _fence_for(body)
            lines += [
                f'<details><summary><code>{d.path}</code> · {d.lines} lines · <a href="{url}">open full file</a></summary>',
                "",
                f"{fence}markdown",
                body + note,
                fence,
                "</details>",
            ]

    fails = [r for r in m.installMatrix if r.result == "fail" and r.error]
    for r in fails:
        safe = r.error.replace("```", "ˋˋˋ")
        lines += [
            "",
            f"<details><summary>❌ {r.language} install failed</summary>",
            "",
            f"`{r.command}`",
            "",
            "```",
            safe,
            "```",
            "</details>",
        ]

    stamp_bits = []
    head = (os.environ.get("HEAD_SHA") or os.environ.get("GITHUB_SHA") or "")[:12]
    if head:
        stamp_bits.append(f"PR head `{head}`")
    run = _run_link()
    if run:
        stamp_bits.append(f"[full run & command output]({run})")
    if stamp_bits:
        lines += ["", "_Checked " + " · ".join(stamp_bits) + "_"]

    return "\n".join(lines)


def _emit(m: Manifest, out_dir: Path, idx: int) -> None:
    out_dir.mkdir(parents=True, exist_ok=True)
    (out_dir / f"build-manifest-{idx}.json").write_text(json.dumps(asdict(m), indent=2))
    fs = factsheet(m)
    (out_dir / f"build-manifest-{idx}.factsheet.md").write_text(fs)
    summary = os.environ.get("GITHUB_STEP_SUMMARY")
    if summary:
        with open(summary, "a") as fh:
            fh.write(fs + "\n")
    print(fs)


# ── cli ─────────────────────────────────────────────────────────────────────
def _cmd_check(args: argparse.Namespace) -> int:
    ensure_resourcedocsgen()
    out_dir = Path(args.out)
    if args.entry:
        entries = [Entry(args.entry[0], args.entry[1])]
    else:
        base = subprocess.run(["git", "show", f"{args.diff}:{PACKAGE_LIST}"],
                              capture_output=True, text=True).stdout or '{"include":[]}'
        entries = diff_entries(base, PACKAGE_LIST.read_text())
    if not entries:
        print("no changed entries")
        return 0
    rc = 0
    for i, entry in enumerate(entries):
        m = verify(entry)
        _emit(m, out_dir, i)
        if not m.gate.verificationGreen and not args.no_fail:
            rc = 1
    return rc


def _changed(base_ref: str) -> list[Entry]:
    base = subprocess.run(["git", "show", f"{base_ref}:{PACKAGE_LIST}"],
                          capture_output=True, text=True).stdout or '{"include":[]}'
    return diff_entries(base, PACKAGE_LIST.read_text())


def _cmd_publish(args: argparse.Namespace) -> int:
    ensure_resourcedocsgen()
    published = publish(_changed(args.diff))
    out = os.environ.get("GITHUB_OUTPUT")
    for name, tag in published:
        line = f"name={name}\nversion={tag}"
        print(line)
        if out:
            with open(out, "a") as fh:
                fh.write(line + "\n")
    return 0


def _cmd_docs(args: argparse.Namespace) -> int:
    # Pre-fetch provider docs + schema for the /review AI step to Read (no execution).
    for entry in diff_entries(Path(args.base).read_text(), Path(args.head).read_text()):
        tag = latest_tag(entry.repoSlug)
        sha = resolve_commit_sha(entry.repoSlug, tag)
        dest = Path(args.out) / entry.repoSlug.split("/")[-1]
        dest.mkdir(parents=True, exist_ok=True)
        for path in ("docs/_index.md", "docs/installation-configuration.md", "README.md", entry.schemaFile):
            data = raw(entry.repoSlug, sha, path)
            if data is not None:
                (dest / Path(path).name).write_bytes(data)
    return 0


def main(argv: list[str]) -> int:
    p = argparse.ArgumentParser(prog="pipeline")
    sub = p.add_subparsers(dest="cmd", required=True)

    c = sub.add_parser("check")
    g = c.add_mutually_exclusive_group(required=True)
    g.add_argument("--entry", nargs=2, metavar=("SLUG", "SCHEMAFILE"))
    g.add_argument("--diff", metavar="BASEREF")
    c.add_argument("--out", default=".")
    c.add_argument("--no-fail", action="store_true", help="do not exit non-zero on a red gate")
    c.set_defaults(func=_cmd_check)

    pub = sub.add_parser("publish")
    pub.add_argument("--diff", metavar="BASEREF", required=True)
    pub.set_defaults(func=_cmd_publish)

    dc = sub.add_parser("docs")
    dc.add_argument("--base", required=True)
    dc.add_argument("--head", required=True)
    dc.add_argument("--out", default="review-input")
    dc.set_defaults(func=_cmd_docs)

    args = p.parse_args(argv)
    return int(args.func(args))


if __name__ == "__main__":
    sys.exit(main(sys.argv[1:]))
