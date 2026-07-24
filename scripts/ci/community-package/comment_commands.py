from __future__ import annotations

import os
from pathlib import Path

import github_api


def _authorized(mode: str, commenter: str, author: str, association: str) -> bool:
    if mode == "author-or-maintainer" and commenter == author:
        return True
    return association in ("OWNER", "MEMBER", "COLLABORATOR")


def _within_cooldown(pr: int, comment_id: str, sha: str, check_name: str, cooldown: int) -> bool:
    elapsed = github_api.minutes_since_check_run(sha, check_name)
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
    author, sha = github_api.pull_request_head(pr)

    if not _authorized("author-or-maintainer", commenter, author, association):
        github_api.add_reaction(comment_id, "-1")
        return 0
    if _within_cooldown(pr, comment_id, sha, r"^check$", cooldown):
        return 0

    run_id = github_api.latest_check_workflow_run("community-package-check.yml", sha)
    if run_id is None:
        github_api.add_reaction(comment_id, "confused")
        github_api.post_comment(pr, f"No check run found for `{sha[:12]}`. Push any commit to trigger one.")
        return 0

    github_api.rerun_workflow(run_id)
    github_api.add_reaction(comment_id, "+1")
    sticky = github_api.fact_sheet_comment(pr)
    if sticky:
        github_api.post_comment(pr, f"🔁 Re-checking `{sha[:12]}` against current upstream. "
                                    f"The [fact-sheet]({sticky['html_url']}) updates in place when it finishes.")
    else:
        github_api.post_comment(pr, f"🔁 Checking `{sha[:12]}`. A fact-sheet comment appears when it finishes.")
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


def _target_pr() -> int | None:
    recorded = Path("pr-number.txt")
    if recorded.exists() and recorded.read_text().strip():
        return int(recorded.read_text().strip())
    # Fork branch names collide (many first-time PRs share "master"/"patch-1"), so match the
    # head repo owner too, not the ref alone.
    head_ref = os.environ.get("PR_HEAD", "")
    head_owner = os.environ.get("PR_HEAD_OWNER", "")
    for pull in github_api.open_pull_requests():
        head = pull["head"]
        owner = (head.get("repo") or {}).get("owner", {}).get("login", "")
        if head["ref"] == head_ref and (not head_owner or owner == head_owner):
            return int(pull["number"])
    return None


def report() -> int:
    pr = _target_pr()
    if pr is None:
        print("no target PR for this run")
        return 0
    sheets = [f.read_text() for f in sorted(Path(".").glob("*.factsheet.md"))]
    body = github_api.FACT_SHEET_MARKER + "\n\n" + "\n\n".join(sheets) + "\n"
    existing = github_api.fact_sheet_comment(pr)
    if existing:
        github_api.edit_comment(int(existing["id"]), body)
    else:
        github_api.post_comment(pr, body)
    print(f"posted fact-sheet to PR #{pr}")
    return 0
