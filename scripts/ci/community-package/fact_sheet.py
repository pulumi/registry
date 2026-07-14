from __future__ import annotations

import os

from models import Manifest

_RESULT_ICON = {"pass": "✅", "fail": "❌", "absent": "➖", "rejected": "🚫"}


def _run_link() -> str | None:
    run_id = os.environ.get("GITHUB_RUN_ID")
    if not run_id:
        return None
    server = os.environ.get("GITHUB_SERVER_URL", "https://github.com")
    repo = os.environ.get("GITHUB_REPOSITORY", "")
    return f"{server}/{repo}/actions/runs/{run_id}"


def _fence_longer_than_any_run_in(text: str) -> str:
    longest = current = 0
    for char in text:
        current = current + 1 if char == "`" else 0
        longest = max(longest, current)
    return "`" * max(3, longest + 1)


def _verdict(manifest: Manifest) -> tuple[str, str]:
    if manifest.gate.verificationGreen:
        return "✅", ("**Ready for approval.** All blocking checks passed; a `@pulumi/iac-cloud` "
                     "review is still required to merge.")
    return "❌", ("**Not ready.** A blocking check failed (see below). Fix upstream, then comment "
                 "`/check` to re-run.")


def render(manifest: Manifest) -> str:
    icon, verdict = _verdict(manifest)
    commit_url = f"https://github.com/{manifest.repoSlug}/commit/{manifest.version.commitSha}"
    owner_note = "already listed" if manifest.owner.alreadyListed else "🆕 first-time contributor"

    lines = [
        f"## {icon} `{manifest.providerName}` · community package check",
        "",
        f"`{manifest.repoSlug}` · `{manifest.version.tag}` · commit [`{manifest.version.commitSha[:12]}`]({commit_url})",
        "",
        verdict,
        "",
        "| check | | command |",
        "|---|:-:|---|",
        f"| docs generate | {'✅' if manifest.generation else '❌'} | `resourcedocsgen metadata from-github` |",
    ]
    for result in manifest.installMatrix:
        label = "plugin install" if result.language == "plugin" else f"{result.language} SDK"
        command = f"`{result.command}`" if result.command else "_not advertised_"
        lines.append(f"| {label} | {_RESULT_ICON.get(result.result, result.result)} | {command} |")
    lines += ["", f"Owner `{manifest.owner.handle}` · {owner_note} · "
                  f"schema `sha256:{manifest.content.schemaSha256[:12]}…`"]

    findings = manifest.docLint["findings"]
    lines.append("")
    if findings:
        lines.append(f"**Doc-lint: {len(findings)} finding(s)** (advisory; these break the registry render surfaces):")
        lines += [f"- `docs/_index.md:{f.line}` {f.kind}: `{f.text}`" for f in findings]
    else:
        lines.append("**Doc-lint** ✅ clean. No relative images or raw HTML.")

    if manifest.docs:
        lines += ["", "**Provider docs** (source at the reviewed commit):"]
        for doc in manifest.docs:
            url = f"https://github.com/{manifest.repoSlug}/blob/{manifest.version.commitSha}/{doc.path}"
            body, note = doc.content, ""
            if doc.lines > 400:
                body = "\n".join(doc.content.splitlines()[:400])
                note = f"\n… truncated to 400 of {doc.lines} lines"
            fence = _fence_longer_than_any_run_in(body)
            lines += [
                f'<details><summary><code>{doc.path}</code> · {doc.lines} lines · <a href="{url}">open full file</a></summary>',
                "",
                f"{fence}markdown",
                body + note,
                fence,
                "</details>",
            ]

    for result in [r for r in manifest.installMatrix if r.result == "fail" and r.error]:
        escaped = result.error.replace("```", "ˋˋˋ")
        lines += [
            "",
            f"<details><summary>❌ {result.language} install failed</summary>",
            "",
            f"`{result.command}`",
            "",
            "```",
            escaped,
            "```",
            "</details>",
        ]

    stamp = []
    head = (os.environ.get("HEAD_SHA") or os.environ.get("GITHUB_SHA") or "")[:12]
    if head:
        stamp.append(f"PR head `{head}`")
    run = _run_link()
    if run:
        stamp.append(f"[full run & command output]({run})")
    if stamp:
        lines += ["", "_Checked " + " · ".join(stamp) + "_"]

    return "\n".join(lines)
