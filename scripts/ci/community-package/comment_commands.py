from __future__ import annotations

import os
from pathlib import Path

import github_api
import package_list


def _authorized(mode: str, commenter: str, author: str, association: str) -> bool:
    if mode == "author-or-maintainer" and commenter == author:
        return True
    return association in ("OWNER", "MEMBER", "COLLABORATOR")


CHECK_WORKFLOW = "community-package-check.yml"
SWEEP_FAILURE_MARKER = "<!-- community-package-sweep-failure -->"
SWEEP_FAILURE_LABEL = "p1"
PACKAGE_LIST = str(package_list.PATH)


def _invokes(body: str, command: str) -> bool:
    fenced = False
    for raw in body.splitlines():
        line = raw.strip()
        if line.startswith("```"):
            fenced = not fenced
        elif not fenced and not line.startswith(">"):
            if line == command or line.startswith(command + " "):
                return True
    return False


def _within_cooldown(pr: int, comment_id: str, cooldown: int) -> bool:
    elapsed = github_api.minutes_since_dispatch(CHECK_WORKFLOW, pr)
    if elapsed is not None and elapsed < cooldown:
        github_api.add_reaction(comment_id, "eyes")
        github_api.post_comment(pr, f"⏳ Rate limited. Last ran {elapsed} min ago; "
                                    f"try again in {cooldown - elapsed} min.")
        return True
    return False


def check_command() -> int:
    pr, comment_id = int(os.environ["PR"]), os.environ["COMMENT_ID"]
    commenter, association = os.environ["COMMENTER"], os.environ["ASSOC"]
    cooldown = int(os.environ.get("COOLDOWN_MINUTES", "10"))
    if not _invokes(os.environ.get("COMMENT_BODY", ""), "/check"):
        return 0
    author, sha = github_api.pull_request_head(pr)

    if not _authorized("author-or-maintainer", commenter, author, association):
        github_api.add_reaction(comment_id, "-1")
        return 0
    if _within_cooldown(pr, comment_id, cooldown):
        return 0

    github_api.dispatch_check(CHECK_WORKFLOW, pr, sha)
    github_api.add_reaction(comment_id, "+1")
    sticky = github_api.fact_sheet_comment(pr)
    if sticky:
        github_api.post_comment(pr, f"🔁 Re-checking `{sha[:12]}` against current upstream. "
                                    f"The [fact-sheet]({sticky['html_url']}) updates in place when it finishes.")
    else:
        github_api.post_comment(pr, f"🔁 Checking `{sha[:12]}`. A fact-sheet comment appears when it finishes.")
    return 0


def sweep() -> int:
    for pull in github_api.open_pull_requests():
        pr, sha = int(pull["number"]), str(pull["head"]["sha"])
        if PACKAGE_LIST not in github_api.pull_request_files(pr):
            continue
        if github_api.dispatch_exists(CHECK_WORKFLOW, github_api.dispatch_run_label(pr, sha)):
            continue
        github_api.dispatch_check(CHECK_WORKFLOW, pr, sha)
        print(f"dispatched a check for PR #{pr} at {sha[:12]}")
    return 0


def _set_output(name: str, value: str) -> None:
    path = os.environ.get("GITHUB_OUTPUT")
    if path:
        with open(path, "a") as fh:
            fh.write(f"{name}={value}\n")


def _preview_reply(first_party: bool) -> tuple[str, str, str]:
    if first_party:
        return ("confused",
                "ℹ️ This PR already builds a preview automatically, because its branch lives in this "
                "repo. So `/preview` is not needed here; watch for the URL from the `Build and deploy "
                "preview` check. The command is for fork PRs, whose automatic preview is skipped for "
                "lack of secrets.",
                "false")
    return ("+1", "🔨 Building a preview. The URL will post here when it is ready.", "true")


def preview_command() -> int:
    pr, comment_id = int(os.environ["PR"]), os.environ["COMMENT_ID"]
    reaction, comment, should_build = _preview_reply(github_api.pull_request_is_first_party(pr))
    github_api.add_reaction(comment_id, reaction)
    github_api.post_comment(pr, comment)
    _set_output("should_build", should_build)
    return 0


def _run_url() -> str:
    server = os.environ.get("GITHUB_SERVER_URL", "https://github.com")
    repo = os.environ.get("GITHUB_REPOSITORY") or os.environ.get("REPO", "")
    return f"{server}/{repo}/actions/runs/{os.environ.get('GITHUB_RUN_ID', '')}"


def sweep_failed() -> int:
    open_already = github_api.open_issue_with_marker(SWEEP_FAILURE_MARKER, SWEEP_FAILURE_LABEL)
    if open_already:
        print(f"issue #{open_already['number']} is already open")
        return 0
    github_api.create_issue(
        "The community package sweep is failing",
        f"{SWEEP_FAILURE_MARKER}\n\nNo community package pull request is being checked while "
        f"this is broken, and a contributor sees no fact-sheet. See the [run log]({_run_url()}).\n\n"
        "Close this once a sweep run succeeds.",
        [SWEEP_FAILURE_LABEL])
    print("opened a failure issue")
    return 0


def preview_failed() -> int:
    pr = int(os.environ["PR"])
    github_api.post_comment(pr, f"❌ The preview build failed. See the [run log]({_run_url()}) for details.")
    comment_id = os.environ.get("COMMENT_ID")
    if comment_id:
        github_api.add_reaction(comment_id, "-1")
    return 0


def _unfinished_sheet() -> str:
    return ("## ❌ The check did not finish\n\n"
            "It wrote no fact-sheet, so this package is still unverified. The fault is in the "
            f"check, not in the package. A maintainer has to read the [run log]({_run_url()}).")


def fact_sheet_body(sheets: list[str]) -> str:
    return github_api.FACT_SHEET_MARKER + "\n\n" + "\n\n".join(sheets or [_unfinished_sheet()]) + "\n"


def report() -> int:
    pr = int(os.environ["PR"])
    sheets = [f.read_text() for f in sorted(Path(".").glob("*.factsheet.md"))]
    body = fact_sheet_body(sheets)
    existing = github_api.fact_sheet_comment(pr)
    if existing:
        github_api.edit_comment(int(existing["id"]), body)
    else:
        github_api.post_comment(pr, body)
    print(f"posted fact-sheet to PR #{pr}")
    return 0
