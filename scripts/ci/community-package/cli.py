#!/usr/bin/env python3
from __future__ import annotations

import argparse
import json
import os
import subprocess
import sys
import traceback
from pathlib import Path

import comment_commands
import fact_sheet
import github_api
import models
import package_list
import resourcedocsgen
import verify_entry


def _write_step_summary(text: str) -> None:
    summary = os.environ.get("GITHUB_STEP_SUMMARY")
    if summary:
        with open(summary, "a") as fh:
            fh.write(text + "\n")


def _rejection_sheet(offending: list[str]) -> str:
    lines = ["## ❌ This PR changes files outside the community package allowlist", "",
             "**Not ready.** A community package PR may only touch the package list "
             f"(`{package_list.PATH}`) and, when a new publisher needs onboarding, the publisher "
             f"allowlist (`{package_list.PUBLISHER_NAMES_PATH}`). The following files are generated and "
             "committed automatically after merge, so remove them and push again:", ""]
    lines += [f"- `{f}`" for f in offending]
    return "\n".join(lines)


def _nothing_to_check_sheet() -> str:
    return (f"## ℹ️ Nothing to check\n\nNo entries were added to `{package_list.PATH}` in this PR. "
            "A community package PR adds one entry to that file.")


def verdict_for(manifests: list[models.Manifest]) -> str:
    if any(not m.green for m in manifests):
        return "fail"
    return "warn" if any(m.warnings for m in manifests) else "pass"


def _write_verdict(out: Path, verdict: str) -> None:
    out.mkdir(parents=True, exist_ok=True)
    (out / comment_commands.VERDICT_FILE).write_text(verdict)


def _write_sheet(out: Path, name: str, sheet: str) -> None:
    out.mkdir(parents=True, exist_ok=True)
    (out / name).write_text(sheet)
    _write_step_summary(sheet)
    print(sheet)


def _crash_sheet(failure: BaseException) -> str:
    return ("## ❌ The check could not run\n\n**Not ready.** The check itself failed before it "
            "could reach a verdict, so this package is unverified. The fault is in the check, "
            f"not in the package.\n\n```\n{failure}\n```")


def run_check(args: argparse.Namespace) -> int:
    try:
        return _check(args)
    except Exception as failure:
        _write_sheet(Path(args.out), "000.factsheet.md", _crash_sheet(failure))
        _write_verdict(Path(args.out), "broken")
        traceback.print_exc()
        return 2


def _check(args: argparse.Namespace) -> int:
    out = Path(args.out)
    offending = package_list.files_outside_allowlist(
        Path(args.changed_files).read_text().splitlines())
    if offending:
        _write_sheet(out, "000.factsheet.md", _rejection_sheet(offending))
        _write_verdict(out, "fail")
        return 1
    entries = package_list.added_entries(package_list.at_ref(args.diff), package_list.current())
    if not entries:
        _write_sheet(out, "000.factsheet.md", _nothing_to_check_sheet())
        _write_verdict(out, "nothing")
        return 0
    resourcedocsgen.ensure_built()
    manifests = []
    for index, entry in enumerate(entries):
        manifest = verify_entry.verify(entry)
        _write_sheet(out, f"{index:03d}.factsheet.md", fact_sheet.render(manifest))
        manifests.append(manifest)
    _write_verdict(out, verdict_for(manifests))
    return 1 if any(not m.green for m in manifests) else 0


def run_preview(args: argparse.Namespace) -> int:
    pull = github_api.pull_request(args.pr)
    base = github_api.file_content_at(github_api.repo(), str(package_list.PATH), pull["base"]["sha"])
    head = github_api.file_content_at(github_api.repo(), str(package_list.PATH), pull["head"]["sha"])
    package_list.PATH.write_text(head)
    entries = package_list.added_entries(base, head)
    if not entries:
        print("no added entries to preview")
        return 0
    resourcedocsgen.ensure_built()
    for entry in entries:
        tag = github_api.latest_release_tag(entry.repoSlug)
        if tag is None:
            print(f"no published release for {entry.repoSlug}; skipping its preview metadata")
            continue
        resourcedocsgen.generate_metadata(entry.repoSlug, entry.schemaFile, tag,
                                          _schema_name(entry, tag))
    return 0


def _schema_name(entry: models.Entry, tag: str) -> str:
    schema = github_api.raw_file(entry.repoSlug, tag, entry.schemaFile)
    return models.provider_name(entry.repoSlug, json.loads(schema) if schema else {})


def run_fetch_pr(args: argparse.Namespace) -> int:
    out = Path(args.out)
    changed = github_api.pull_request_files(args.pr)
    Path(args.changed_files).write_text("\n".join(changed) + "\n")
    offending = package_list.files_outside_allowlist(changed)
    if offending:
        _write_sheet(out, "000.factsheet.md", _rejection_sheet(offending))
        _write_verdict(out, "fail")
        return 1
    head = str(github_api.pull_request(args.pr)["head"]["sha"])
    for path in package_list.ALLOWED_PATHS:
        path.write_text(github_api.file_content_at(github_api.repo(), str(path), head))
    return 0


PACKAGE_YAML_DIR = Path("themes/default/data/registry/packages")


def _changed_package_yamls(base_ref: str) -> list[Path]:
    changed = subprocess.run(["git", "diff", "--name-only", f"{base_ref}...HEAD"],
                             capture_output=True, text=True).stdout
    return [Path(f) for f in changed.splitlines()
            if f.endswith(".yaml") and Path(f).parent == PACKAGE_YAML_DIR]


def run_check_publish(args: argparse.Namespace) -> int:
    out = Path(args.out)
    changed = _changed_package_yamls(args.diff)
    if not changed:
        print("no package YAML changed in this PR")
        return 0

    repo_root = Path.cwd()
    failed = False
    for index, yaml_path in enumerate(sorted(changed)):
        manifest = verify_entry.verify_package_yaml(yaml_path, repo_root)
        if manifest.delisted:
            print(f"{manifest.providerName} is delisted; nothing to check")
            continue
        _write_sheet(out, f"{index:03d}.factsheet.md", fact_sheet.render(manifest))
        if not manifest.green:
            failed = True
    return 1 if failed else 0


def main(argv: list[str]) -> int:
    parser = argparse.ArgumentParser(prog="community-package")
    sub = parser.add_subparsers(dest="command", required=True)

    check = sub.add_parser("check")
    check.add_argument("--diff", metavar="BASEREF", required=True)
    check.add_argument("--changed-files", metavar="PATH", required=True)
    check.add_argument("--out", default=".")
    check.set_defaults(run=run_check)

    fetch_pr = sub.add_parser("fetch-pr")
    fetch_pr.add_argument("--pr", type=int, required=True)
    fetch_pr.add_argument("--changed-files", metavar="PATH", required=True)
    fetch_pr.add_argument("--out", default=".")
    fetch_pr.set_defaults(run=run_fetch_pr)

    check_publish = sub.add_parser("check-publish")
    check_publish.add_argument("--diff", metavar="BASEREF", required=True)
    check_publish.add_argument("--out", default=".")
    check_publish.set_defaults(run=run_check_publish)

    preview = sub.add_parser("preview")
    preview.add_argument("--pr", type=int, required=True)
    preview.set_defaults(run=run_preview)

    sub.add_parser("report").set_defaults(run=lambda _: comment_commands.report())
    sub.add_parser("sweep").set_defaults(run=lambda _: comment_commands.sweep())
    sub.add_parser("check-command").set_defaults(run=lambda _: comment_commands.check_command())
    sub.add_parser("preview-command").set_defaults(run=lambda _: comment_commands.preview_command())
    sub.add_parser("sweep-failed").set_defaults(run=lambda _: comment_commands.sweep_failed())
    sub.add_parser("preview-failed").set_defaults(run=lambda _: comment_commands.preview_failed())

    args = parser.parse_args(argv)
    return int(args.run(args))


if __name__ == "__main__":
    sys.exit(main(sys.argv[1:]))
