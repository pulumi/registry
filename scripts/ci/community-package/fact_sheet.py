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
    if not manifest.green:
        return "❌", ("**Not ready.** A required check failed (see below). Fix upstream, then comment "
                     "`/check` to re-run.")
    if manifest.warnings:
        return "🟡", ("**Passes the required checks, but review the warnings.** The plugin installs and "
                     "docs generate, but one or more advisory checks did not pass — see the SDK installs "
                     "and doc-lint below before approving.")
    return "✅", ("**Ready for approval.** All checks passed; a `@pulumi/iac-cloud` review is still "
                 "required to merge.")


def _stamp_lines() -> list[str]:
    stamp = []
    head = (os.environ.get("HEAD_SHA") or os.environ.get("GITHUB_SHA") or "")[:12]
    if head:
        stamp.append(f"PR head `{head}`")
    run = _run_link()
    if run:
        stamp.append(f"[full run & command output]({run})")
    return ["", "_Checked " + " · ".join(stamp) + "_"] if stamp else []


def render(manifest: Manifest) -> str:
    if manifest.error:
        return "\n".join([
            f"## ❌ `{manifest.providerName}` · community package check",
            "",
            f"`{manifest.repoSlug or manifest.schemaFile}`",
            "",
            f"**Not ready.** {manifest.error}",
        ] + _stamp_lines())

    icon, verdict = _verdict(manifest)

    if manifest.version.commitSha:
        commit_url = f"https://github.com/{manifest.repoSlug}/commit/{manifest.version.commitSha}"
        subtitle = (f"`{manifest.repoSlug}` · `{manifest.version.tag}` · commit "
                    f"[`{manifest.version.commitSha[:12]}`]({commit_url})")
    else:
        subtitle = f"`{manifest.repoSlug or manifest.schemaFile}` · `{manifest.version.tag}`"

    lines = [
        f"## {icon} `{manifest.providerName}` · community package check",
        "",
        subtitle,
        "",
        verdict,
        "",
        "| check | | command |",
        "|---|:-:|---|",
        f"| docs generate | {'✅' if manifest.generation else '❌'} | `resourcedocsgen metadata "
        f"{'from-github' if manifest.version.commitSha else 'from-urls'}` |",
    ]
    for result in manifest.installMatrix:
        if result.language == "plugin":
            label = "plugin install"
        elif result.language == "package add":
            label = "package add"
        else:
            label = f"{result.language} SDK"
        command = f"`{result.command}`" if result.command else "_not advertised_"
        lines.append(f"| {label} | {_RESULT_ICON.get(result.result, result.result)} | {command} |")
    if manifest.publisher:
        if manifest.publisherKnown:
            listed = "✅"
        else:
            listed = "⚠️" if manifest.green else "❌"
        lines.append(f"| publisher listed | {listed} | `{manifest.publisher}` is a key in publisher-names.json |")
    if manifest.schemaVersion:
        lines.append(f"| schema version | {'✅' if manifest.schemaVersionMatches else '❌'} "
                     f"| `{manifest.schemaVersion}` in `{manifest.schemaFile}` |")
    lines += ["", f"Owner `{manifest.owner}`"]

    if manifest.publisher and not manifest.publisherKnown:
        lines += ["", f"**Publisher** ❌ no entry for `{manifest.publisher}`. `publisher-names.json` maps the "
                      "`publisher` string in a provider's schema (the key) to that publisher's slug in the "
                      "registry backend (the value, which goes into an API path). Add "
                      f"`\"{manifest.publisher}\": \"<slug>\"` to "
                      "`tools/resourcedocsgen/pkg/publishers/publisher-names.json` in this PR; the key and the "
                      "value are usually the same. If this publisher already ships under a slug, use that one "
                      "so the two do not split. Without the entry the schema fetch from the registry backend "
                      "fails and falls back to VCS."]

    if not manifest.schemaVersionMatches:
        lines += ["", f"**Schema version** ❌ `{manifest.schemaFile}` declares version "
                      f"`{manifest.schemaVersion}`, but the release is `{manifest.version.tag}`. The registry "
                      "publishes the schema straight from the release tag and rejects a version that does not "
                      "match. Either omit the `version` key entirely, which is what bridged providers do so the "
                      "registry takes the version from the publish request, or stamp the real version at "
                      "release time."]

    for result in [r for r in manifest.installMatrix if r.result == "rejected" and r.error]:
        lines += ["", f"**{result.language}** 🚫 {result.error}"]

    findings = manifest.docLint
    lines.append("")
    if not manifest.indexPresent:
        lines.append("**Doc-lint** ❌ the provider repo has no `docs/_index.md` at the reviewed commit. "
                     "It is the package's front page in the registry, so it is required, and "
                     "`resourcedocsgen` cannot generate metadata without it. Add it upstream, cut a "
                     "release, then comment `/check`.")
    elif findings:
        lines.append(f"**Doc-lint: {len(findings)} finding(s)** (advisory; these break the registry render surfaces):")
        lines += [f"- `docs/_index.md:{f.line}` {f.kind}: `{f.text}`" for f in findings]
    else:
        lines.append("**Doc-lint** ✅ clean. No relative images or raw HTML.")

    if manifest.docs:
        source = "source at the reviewed commit" if manifest.version.commitSha else "source as published"
        lines += ["", f"**Provider docs** ({source}):"]
        for doc in manifest.docs:
            url = (f"https://github.com/{manifest.repoSlug}/blob/{manifest.version.commitSha}/{doc.path}"
                   if manifest.version.commitSha else manifest.docSourceURL or manifest.schemaFile)
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

    if not manifest.generation and manifest.generationError:
        lines += [
            "",
            "<details><summary>❌ docs generate failed</summary>",
            "",
            "`resourcedocsgen metadata from-github`",
            "",
            "```",
            manifest.generationError.replace("```", "ˋˋˋ"),
            "```",
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

    lines += _stamp_lines()
    return "\n".join(lines)
