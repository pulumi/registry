from __future__ import annotations

import base64
import json
import os
import re
import urllib.error
import urllib.request
from datetime import datetime, timezone
from typing import Any

FACT_SHEET_MARKER = "<!-- community-package-fact-sheet -->"


def repo() -> str:
    return os.environ.get("REPO") or os.environ["GITHUB_REPOSITORY"]


def request(path: str, method: str = "GET", body: dict[str, Any] | None = None) -> Any:
    encoded = json.dumps(body).encode() if body is not None else None
    req = urllib.request.Request(f"https://api.github.com{path}", data=encoded, method=method)
    token = os.environ.get("GITHUB_TOKEN")
    if token:
        req.add_header("Authorization", f"Bearer {token}")
    req.add_header("Accept", "application/vnd.github+json")
    if encoded is not None:
        req.add_header("Content-Type", "application/json")
    with urllib.request.urlopen(req, timeout=30) as resp:
        payload = resp.read()
    return json.loads(payload) if payload else None


def raw_file(slug: str, ref: str, path: str) -> bytes | None:
    url = f"https://raw.githubusercontent.com/{slug}/{ref}/{path}"
    try:
        with urllib.request.urlopen(urllib.request.Request(url), timeout=30) as resp:
            return bytes(resp.read())
    except urllib.error.HTTPError:
        return None


def latest_release_tag(slug: str) -> str | None:
    try:
        return str(request(f"/repos/{slug}/releases/latest")["tag_name"])
    except urllib.error.HTTPError as e:
        if e.code == 404:  # tags-only repo, or only prereleases: no "latest" release exists
            return None
        raise


def commit_sha_for_tag(slug: str, tag: str) -> str:
    ref = request(f"/repos/{slug}/git/refs/tags/{tag}")["object"]
    if ref["type"] == "tag":  # annotated tag, one more hop to reach the commit
        return str(request(f"/repos/{slug}/git/tags/{ref['sha']}")["object"]["sha"])
    return str(ref["sha"])


def file_content_at(slug: str, path: str, ref: str) -> str:
    encoded = request(f"/repos/{slug}/contents/{path}?ref={ref}")["content"]
    return base64.b64decode(encoded).decode()


def pull_request(pr: int) -> dict[str, Any]:
    return dict(request(f"/repos/{repo()}/pulls/{pr}"))


def pull_request_head(pr: int) -> tuple[str, str]:
    pull = pull_request(pr)
    return str(pull["user"]["login"]), str(pull["head"]["sha"])


def pull_request_is_first_party(pr: int) -> bool:
    pull = pull_request(pr)
    head_repo = (pull["head"].get("repo") or {}).get("full_name")
    return bool(head_repo) and head_repo == pull["base"]["repo"]["full_name"]


def open_pull_requests() -> list[dict[str, Any]]:
    return list(request(f"/repos/{repo()}/pulls?state=open&per_page=100"))


def add_reaction(comment_id: str, content: str) -> None:
    try:
        request(f"/repos/{repo()}/issues/comments/{comment_id}/reactions", "POST", {"content": content})
    except Exception:
        pass


def post_comment(pr: int, body: str) -> None:
    request(f"/repos/{repo()}/issues/{pr}/comments", "POST", {"body": body})


def edit_comment(comment_id: int, body: str) -> None:
    request(f"/repos/{repo()}/issues/comments/{comment_id}", "PATCH", {"body": body})


def fact_sheet_comment(pr: int) -> dict[str, Any] | None:
    for comment in request(f"/repos/{repo()}/issues/{pr}/comments"):
        if FACT_SHEET_MARKER in comment["body"]:
            return dict(comment)
    return None


def _minutes_since(timestamp: str) -> int:
    moment = datetime.strptime(timestamp, "%Y-%m-%dT%H:%M:%SZ").replace(tzinfo=timezone.utc)
    return int((datetime.now(timezone.utc) - moment).total_seconds() // 60)


def minutes_since_check_run(sha: str, name_pattern: str) -> int | None:
    runs = [r for r in request(f"/repos/{repo()}/commits/{sha}/check-runs")["check_runs"]
            if re.search(name_pattern, r["name"])]
    if not runs:
        return None
    return _minutes_since(str(max(runs, key=lambda r: str(r["started_at"]))["started_at"]))


def pull_request_files(pr: int) -> list[str]:
    # The allowlist gate reads this list, so a truncated first page would let unlisted files
    # through unseen. Walk every page.
    names: list[str] = []
    page = 1
    while True:
        batch = request(f"/repos/{repo()}/pulls/{pr}/files?per_page=100&page={page}")
        names += [str(f["filename"]) for f in batch]
        if len(batch) < 100:
            return names
        page += 1


def _short(sha: str) -> str:
    return sha[:12]


def dispatch_run_label(pr: int, sha: str) -> str:
    return f"Community package check · PR #{pr} · {_short(sha)}"


def dispatch_check(workflow_file: str, pr: int, sha: str) -> None:
    """Dispatch a check named exactly `dispatch_run_label`, which is the sweep's dedupe key.

    The workflow interpolates the `head` input into its `run-name`, so the input has to carry
    the same short sha the label does; otherwise no run ever matches and the sweep dispatches
    the same check again on every pass.
    """
    dispatch_workflow(workflow_file, {"pr": str(pr), "head": _short(sha)})


def _dispatched_runs(workflow_file: str) -> list[dict[str, Any]]:
    return list(request(f"/repos/{repo()}/actions/workflows/{workflow_file}/runs"
                        f"?event=workflow_dispatch&per_page=100")["workflow_runs"])


def _run_title(run: dict[str, Any]) -> str:
    # `name` is the workflow's own name; an evaluated `run-name:` arrives as `display_title`.
    return str(run.get("display_title") or run.get("name") or "")


def dispatch_exists(workflow_file: str, label: str) -> bool:
    return any(_run_title(run) == label for run in _dispatched_runs(workflow_file))


def minutes_since_dispatch(workflow_file: str, pr: int) -> int | None:
    mine = [r for r in _dispatched_runs(workflow_file) if f"PR #{pr} ·" in _run_title(r)]
    if not mine:
        return None
    return _minutes_since(str(max(mine, key=lambda r: str(r["created_at"]))["created_at"]))


def pull_request_run_status(workflow_file: str, sha: str) -> str | None:
    """The run's live status, or its conclusion once it has one.

    An approval-gated run surfaces either as `status: action_required` or, once GitHub closes
    it out, as `status: completed` with `conclusion: action_required`. The sweep has to see
    both, or it skips forever exactly the PRs it exists to unblock.
    """
    runs = request(f"/repos/{repo()}/actions/workflows/{workflow_file}/runs"
                   f"?event=pull_request&head_sha={sha}")["workflow_runs"]
    if not runs:
        return None
    return str(runs[0].get("conclusion") or runs[0]["status"])


def dispatch_workflow(workflow_file: str, inputs: dict[str, str]) -> None:
    ref = os.environ.get("DEFAULT_BRANCH") or "master"
    request(f"/repos/{repo()}/actions/workflows/{workflow_file}/dispatches", "POST",
            {"ref": ref, "inputs": inputs})
