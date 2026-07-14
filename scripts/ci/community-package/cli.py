#!/usr/bin/env python3
"""Command-line entry point for the community package tooling. Each subcommand is one CI step:

    check           verify the added entry and write a fact-sheet (runs on the PR)
    publish         rebuild and lock the entry after merge (runs on master)
    report          post the fact-sheet as a sticky PR comment
    check-command   handle a /check comment
    review-command  gate a /review comment

The logic lives in the sibling modules (verify_entry, post_merge_publish, fact_sheet, comment_commands, ...); stdlib
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
import package_list
import post_merge_publish
import resourcedocsgen
import verify_entry
from models import Entry, Manifest


def _write_step_summary(text: str) -> None:
    summary = os.environ.get("GITHUB_STEP_SUMMARY")
    if summary:
        with open(summary, "a") as fh:
            fh.write(text + "\n")


def _reject_non_entry_changes(base_ref: str) -> None:
    changed = subprocess.run(["git", "diff", "--name-only", f"{base_ref}...HEAD"],
                             capture_output=True, text=True).stdout
    offending = [f for f in changed.splitlines() if f and f != str(package_list.PATH)]
    if offending:
        lines = [f"## ❌ This PR must change only `{package_list.PATH}`", "",
                 "The following files do not belong in a community package PR. They are",
                 "generated and committed automatically after merge:", ""]
        lines += [f"- `{f}`" for f in offending]
        _write_step_summary("\n".join(lines))
        raise SystemExit("PR changes files other than the package list")


def _emit_fact_sheet(manifest: Manifest, out: Path, index: int) -> None:
    out.mkdir(parents=True, exist_ok=True)
    sheet = fact_sheet.render(manifest)
    (out / f"{index}.factsheet.md").write_text(sheet)
    _write_step_summary(sheet)
    print(sheet)


def run_check(args: argparse.Namespace) -> int:
    if args.entry:
        entries = [Entry(args.entry[0], args.entry[1])]
    else:
        _reject_non_entry_changes(args.diff)
        entries = package_list.added_entries(package_list.at_ref(args.diff), package_list.current())
    if not entries:
        print("no changed entries")
        return 0
    resourcedocsgen.ensure_built()
    failed = False
    for index, entry in enumerate(entries):
        manifest = verify_entry.verify(entry)
        _emit_fact_sheet(manifest, Path(args.out), index)
        if not manifest.gate.verificationGreen and not args.no_fail:
            failed = True
    return 1 if failed else 0


def run_publish(args: argparse.Namespace) -> int:
    resourcedocsgen.ensure_built()
    entries = package_list.added_entries(package_list.at_ref(args.diff), package_list.current())
    output = os.environ.get("GITHUB_OUTPUT")
    for name, tag in post_merge_publish.publish_entries(entries):
        print(f"name={name}\nversion={tag}")
        if output:
            with open(output, "a") as fh:
                fh.write(f"name={name}\nversion={tag}\n")
    return 0


def main(argv: list[str]) -> int:
    parser = argparse.ArgumentParser(prog="community-package")
    sub = parser.add_subparsers(dest="command", required=True)

    check = sub.add_parser("check")
    source = check.add_mutually_exclusive_group(required=True)
    source.add_argument("--entry", nargs=2, metavar=("SLUG", "SCHEMAFILE"))
    source.add_argument("--diff", metavar="BASEREF")
    check.add_argument("--out", default=".")
    check.add_argument("--no-fail", action="store_true", help="do not exit non-zero on a red gate")
    check.set_defaults(run=run_check)

    published = sub.add_parser("publish")
    published.add_argument("--diff", metavar="BASEREF", required=True)
    published.set_defaults(run=run_publish)

    sub.add_parser("report").set_defaults(run=lambda _: comment_commands.report())
    sub.add_parser("check-command").set_defaults(run=lambda _: comment_commands.check_command())
    sub.add_parser("review-command").set_defaults(run=lambda _: comment_commands.review_command())

    args = parser.parse_args(argv)
    return int(args.run(args))


if __name__ == "__main__":
    sys.exit(main(sys.argv[1:]))
