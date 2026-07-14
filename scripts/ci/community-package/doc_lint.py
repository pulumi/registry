from __future__ import annotations

import re

from models import DocFinding

# Patterns that break on the registry render surfaces. See the review skill for why each one
# matters; the citations there point at the PRs where they bit.
_PATTERNS = [
    (re.compile(r'<img[^>]+src="(\.|/|[^h])'), "relative-image"),
    (re.compile(r"<img[^>]+src='(\.|/|[^h])"), "relative-image"),
    (re.compile(r"!\[[^\]]*\]\((\.|/)"), "relative-image"),
    (re.compile(r'<a[^>]+href="(\.|/)'), "raw-relative-link"),
    (re.compile(r"\]\((?!https?://)[^)]*\.md[)#]"), "md-suffixed-link"),
]


def find_issues(text: str) -> list[DocFinding]:
    issues: list[DocFinding] = []
    for number, line in enumerate(text.splitlines(), 1):
        for pattern, kind in _PATTERNS:
            if pattern.search(line):
                issues.append(DocFinding(number, kind, line[:160]))
                break
    return issues
