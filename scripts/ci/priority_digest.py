#!/usr/bin/env python3

import argparse
import os
import sys
from datetime import datetime, timezone

import requests

REPOSITORIES = ["pulumi/registry", "pulumi/terraform-to-pulumi-registry-pipeline"]
PRIORITY_LABELS = ["p0", "p1"]
LISTED_AT_MOST = 20

QUERY = " ".join(
    [f"repo:{repository}" for repository in REPOSITORIES]
    + ["is:issue", "is:open", "label:" + ",".join(PRIORITY_LABELS), "sort:created-asc"]
)


def require_env(name):
    value = os.getenv(name)
    if not value:
        sys.exit(f"{name} is not set")
    return value


def search_issues():
    token = require_env("GITHUB_TOKEN")
    issues = []
    page = 1
    while True:
        response = requests.get(
            "https://api.github.com/search/issues",
            params={"q": QUERY, "per_page": 100, "page": page},
            headers={
                "Accept": "application/vnd.github+json",
                "Authorization": f"Bearer {token}",
            },
            timeout=30,
        )
        response.raise_for_status()
        body = response.json()
        issues += body["items"]
        if len(issues) >= body["total_count"] or not body["items"]:
            return issues
        page += 1


def escape(text):
    return text.replace("&", "&amp;").replace("<", "&lt;").replace(">", "&gt;")


def age_in_days(issue, now=None):
    created = datetime.fromisoformat(issue["created_at"].replace("Z", "+00:00"))
    return ((now or datetime.now(timezone.utc)) - created).days


def repository_of(issue):
    return "/".join(issue["repository_url"].split("/")[-2:])


def describe(issue, now=None):
    labels = [label["name"] for label in issue["labels"]]
    priority = ", ".join(label for label in labels if label in PRIORITY_LABELS)
    kind = ", ".join(label for label in labels if label.startswith("kind/"))
    assignees = ", ".join(assignee["login"] for assignee in issue["assignees"]) or "unassigned"
    return " · ".join(part for part in [priority, kind, assignees, f"{age_in_days(issue, now)}d old"] if part)


def render(issues, now=None):
    if not issues:
        return ":white_check_mark: *Registry P0 and P1*: nothing open."

    search_url = "https://github.com/search?type=issues&q=" + requests.utils.quote(QUERY)
    lines = [f":bell: *Registry P0 and P1*: <{search_url}|{len(issues)} open>, oldest first."]

    for repository in REPOSITORIES:
        listed = [issue for issue in issues if repository_of(issue) == repository]
        if not listed:
            continue
        lines.append("")
        lines.append(f"*{repository}* ({len(listed)})")
        for issue in listed[:LISTED_AT_MOST]:
            lines.append(
                f"• <{issue['html_url']}|#{issue['number']}> {escape(issue['title'])} — {describe(issue, now)}"
            )
        if len(listed) > LISTED_AT_MOST:
            lines.append(f"• …and {len(listed) - LISTED_AT_MOST} more")

    return "\n".join(lines)


def log_counts(issues):
    for repository in REPOSITORIES:
        found = sum(1 for issue in issues if repository_of(issue) == repository)
        print(f"{repository}: {found} open", file=sys.stderr)


def post(channel, message):
    response = requests.post(
        require_env("SLACK_WEBHOOK_URL"),
        json={
            "channel": channel,
            "text": message,
            "username": "registrybot",
            "icon_url": "https://www.pulumi.com/logos/brand/avatar-on-white.png",
            "mrkdwn": True,
            "unfurl_links": False,
        },
        timeout=30,
    )
    if response.status_code != 200 or response.text.strip() != "ok":
        sys.exit(f"Slack rejected the message: {response.status_code} {response.text.strip()}")


def main():
    parser = argparse.ArgumentParser(
        prog="priority-digest",
        description="""Post the open P0 and P1 registry issues to Slack, oldest first

        GITHUB_TOKEN must be set to read issues.
        SLACK_WEBHOOK_URL must be set to post.
        SLACK_CHANNEL may be set. Defaults to #team-iac-cloud.
        """,
        epilog="This is a Pulumi internal tool - it is not intended for external use")
    parser.add_argument("--dry-run", action="store_true", help="Print the message instead of posting it")
    parser.add_argument("--channel", default=os.getenv("SLACK_CHANNEL", "#team-iac-cloud"),
                        help="Slack channel to post to, as #name.")

    args = parser.parse_args()

    issues = search_issues()
    log_counts(issues)
    message = render(issues)

    if args.dry_run:
        print(message)
    else:
        post(args.channel, message)


if __name__ == "__main__":
    main()
