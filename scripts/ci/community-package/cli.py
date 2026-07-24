#!/usr/bin/env python3
"""Command-line entry point for the community package tooling. Each subcommand is one CI step:

    check           verify the added entry and write a fact-sheet (runs on the PR)
    preview         materialize a fork PR's entry so its site preview can be built
    report          post the fact-sheet as a sticky PR comment
    check-command   handle a /check comment
    preview-command handle a /preview comment
    preview-failed  post a build-failure comment for a /preview run

The logic lives in the sibling modules (verify_entry, resourcedocsgen, fact_sheet, comment_commands, ...); stdlib
only, no third-party dependencies.
"""
from __future__ import annotations

import argparse
import os
import subprocess
import sys
from pathlib import Path

import comment_commands
import fact_sheet
import github_api
import package_list
import resourcedocsgen
import verify_entry


def _write_step_summary(text: str) -> None:
    summary = os.environ.get("GITHUB_STEP_SUMMARY")
    if summary:
        with open(summary, "a") as fh:
            fh.write(text + "\n")


# The publisher allowlist is hand-maintained Go source embedded into resourcedocsgen, not
# generated after merge, so a new publisher's entry has to ride along with the onboarding PR.
PUBLISHER_NAMES_PATH = Path("tools/resourcedocsgen/pkg/publishers/publisher-names.json")
ALLOWED_PATHS = {str(package_list.PATH), str(PUBLISHER_NAMES_PATH)}


def _files_outside_allowlist(changed: list[str]) -> list[str]:
    return [f for f in changed if f and f not in ALLOWED_PATHS]


def _offending_files(base_ref: str) -> list[str]:
    changed = subprocess.run(["git", "diff", "--name-only", f"{base_ref}...HEAD"],
                             capture_output=True, text=True).stdout
    return _files_outside_allowlist(changed.splitlines())


def _rejection_sheet(offending: list[str]) -> str:
    lines = ["## ❌ This PR changes files outside the community package allowlist", "",
             "**Not ready.** A community package PR may only touch the package list "
             f"(`{package_list.PATH}`) and, when a new publisher needs onboarding, the publisher "
             f"allowlist (`{PUBLISHER_NAMES_PATH}`). The following files are generated and committed "
             "automatically after merge, so remove them and push again:", ""]
    lines += [f"- `{f}`" for f in offending]
    return "\n".join(lines)


def _nothing_to_check_sheet() -> str:
    return (f"## ℹ️ Nothing to check\n\nNo entries were added to `{package_list.PATH}` in this PR. "
            "A community package PR adds one entry to that file.")


def _write_sheet(out: Path, name: str, sheet: str) -> None:
    out.mkdir(parents=True, exist_ok=True)
    (out / name).write_text(sheet)
    _write_step_summary(sheet)
    print(sheet)


def run_check(args: argparse.Namespace) -> int:
    out = Path(args.out)
    offending = _offending_files(args.diff)
    if offending:
        _write_sheet(out, "000.factsheet.md", _rejection_sheet(offending))
        return 1
    entries = package_list.added_entries(package_list.at_ref(args.diff), package_list.current())
    if not entries:
        _write_sheet(out, "000.factsheet.md", _nothing_to_check_sheet())
        return 0
    resourcedocsgen.ensure_built()
    failed = False
    for index, entry in enumerate(entries):
        manifest = verify_entry.verify(entry)
        _write_sheet(out, f"{index:03d}.factsheet.md", fact_sheet.render(manifest))
        if not manifest.green:
            failed = True
    return 1 if failed else 0


def run_preview(args: argparse.Namespace) -> int:
    # Diff against the PR base by SHA, not the on-disk file: a native pull_request run is checked
    # out at the head, where the on-disk list already contains the new entry.
    pull = github_api.pull_request(args.pr)
    base = github_api.file_content_at(github_api.repo(), str(package_list.PATH), pull["base"]["sha"])
    head = github_api.file_content_at(github_api.repo(), str(package_list.PATH), pull["head"]["sha"])
    package_list.PATH.write_text(head)  # bring the entry into the tree so the site build sees it
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
        resourcedocsgen.generate_metadata(entry.repoSlug, entry.schemaFile, tag)
    return 0


def main(argv: list[str]) -> int:
    parser = argparse.ArgumentParser(prog="community-package")
    sub = parser.add_subparsers(dest="command", required=True)

    check = sub.add_parser("check")
    check.add_argument("--diff", metavar="BASEREF", required=True)
    check.add_argument("--out", default=".")
    check.set_defaults(run=run_check)

    preview = sub.add_parser("preview")
    preview.add_argument("--pr", type=int, required=True)
    preview.set_defaults(run=run_preview)

    sub.add_parser("report").set_defaults(run=lambda _: comment_commands.report())
    sub.add_parser("check-command").set_defaults(run=lambda _: comment_commands.check_command())
    sub.add_parser("preview-command").set_defaults(run=lambda _: comment_commands.preview_command())
    sub.add_parser("preview-failed").set_defaults(run=lambda _: comment_commands.preview_failed())

    args = parser.parse_args(argv)
    return int(args.run(args))


if __name__ == "__main__":
    sys.exit(main(sys.argv[1:]))
