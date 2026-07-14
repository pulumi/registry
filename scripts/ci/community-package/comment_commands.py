from __future__ import annotations

import os
from pathlib import Path

import github_api
import package_list


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


def _prefetch_provider_docs(before: str, after: str, out: Path) -> None:
    for entry in package_list.added_entries(before, after):
        tag = github_api.latest_release_tag(entry.repoSlug)
        sha = github_api.commit_sha_for_tag(entry.repoSlug, tag)
        destination = out / entry.repoSlug.split("/")[-1]
        destination.mkdir(parents=True, exist_ok=True)
        for path in ("docs/_index.md", "docs/installation-configuration.md", "README.md", entry.schemaFile):
            data = github_api.raw_file(entry.repoSlug, sha, path)
            if data is not None:
                (destination / Path(path).name).write_bytes(data)


def check_command() -> int:
    pr, comment_id = int(os.environ["PR"]), os.environ["COMMENT_ID"]
    commenter, association = os.environ["COMMENTER"], os.environ["ASSOC"]
    cooldown = int(os.environ.get("COOLDOWN_MINUTES", "10"))
    author, sha = github_api.pull_request_head(pr)

    if not _authorized("author-or-maintainer", commenter, author, association):
        github_api.add_reaction(comment_id, "-1")
        return 0
    if _within_cooldown(pr, comment_id, sha, "check", cooldown):
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


def review_command() -> int:
    pr, comment_id = int(os.environ["PR"]), os.environ["COMMENT_ID"]
    commenter, association = os.environ["COMMENTER"], os.environ["ASSOC"]
    cooldown = int(os.environ.get("COOLDOWN_MINUTES", "15"))
    author, sha = github_api.pull_request_head(pr)

    proceed = _run_review_gate(pr, comment_id, commenter, author, association, sha, cooldown)
    output = os.environ.get("GITHUB_OUTPUT")
    if output:
        with open(output, "a") as fh:
            fh.write(f"proceed={'true' if proceed else 'false'}\n")
    return 0


def _run_review_gate(pr: int, comment_id: str, commenter: str, author: str,
                     association: str, sha: str, cooldown: int) -> bool:
    if not _authorized("maintainer-only", commenter, author, association):
        github_api.add_reaction(comment_id, "-1")
        return False
    if _within_cooldown(pr, comment_id, sha, "review", cooldown):
        return False
    after = github_api.file_content_at(github_api.repo(), str(package_list.PATH), sha)
    _prefetch_provider_docs(package_list.current(), after, Path("review-input"))
    github_api.add_reaction(comment_id, "+1")
    github_api.post_comment(pr, "🔎 Running the AI documentation review. "
                                "Findings post as a comment when it finishes.")
    return True


def report() -> int:
    head_ref = os.environ["PR_HEAD"]
    pull = next((p for p in github_api.open_pull_requests() if p["head"]["ref"] == head_ref), None)
    if pull is None:
        print(f"no open PR for {head_ref}")
        return 0
    pr = int(pull["number"])
    sheets = [f.read_text() for f in sorted(Path(".").glob("*.factsheet.md"))]
    body = github_api.FACT_SHEET_MARKER + "\n\n" + "\n\n".join(sheets) + "\n"
    existing = github_api.fact_sheet_comment(pr)
    if existing:
        github_api.edit_comment(int(existing["id"]), body)
    else:
        github_api.post_comment(pr, body)
    print(f"posted fact-sheet to PR #{pr}")
    return 0
