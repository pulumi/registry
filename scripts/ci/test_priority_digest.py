#!/usr/bin/env python3
"""Unit tests for priority_digest.py"""

import unittest
from datetime import datetime, timezone

from priority_digest import describe, escape, render, repository_of


NOW = datetime(2026, 8, 14, 12, 0, 0, tzinfo=timezone.utc)


def issue(number=1, title="Something broke", repository="registry", created="2026-08-04T12:00:00Z",
          labels=("p1", "kind/bug"), assignees=()):
    return {
        "number": number,
        "title": title,
        "html_url": f"https://github.com/pulumi/{repository}/issues/{number}",
        "repository_url": f"https://api.github.com/repos/pulumi/{repository}",
        "created_at": created,
        "labels": [{"name": name} for name in labels],
        "assignees": [{"login": login} for login in assignees],
    }


class TestEscape(unittest.TestCase):
    def test_escapes_slack_control_characters(self):
        self.assertEqual(escape("<https://evil.example|Click>"), "&lt;https://evil.example|Click&gt;")

    def test_escapes_ampersand_first(self):
        self.assertEqual(escape("a & <b>"), "a &amp; &lt;b&gt;")

    def test_leaves_ordinary_titles_alone(self):
        self.assertEqual(escape("Pipeline failed: intersight did not update"),
                         "Pipeline failed: intersight did not update")


class TestRepositoryOf(unittest.TestCase):
    def test_reads_owner_and_name(self):
        self.assertEqual(repository_of(issue(repository="registry")), "pulumi/registry")


class TestDescribe(unittest.TestCase):
    def test_lists_priority_kind_assignee_and_age(self):
        self.assertEqual(describe(issue(), NOW), "p1 · kind/bug · unassigned · 10d old")

    def test_names_assignees(self):
        self.assertEqual(describe(issue(assignees=("fnune",)), NOW), "p1 · kind/bug · fnune · 10d old")

    def test_omits_labels_it_does_not_recognise(self):
        self.assertEqual(describe(issue(labels=("needs-triage",)), NOW), "unassigned · 10d old")


class TestRender(unittest.TestCase):
    def test_reports_nothing_open(self):
        self.assertEqual(render([]), ":white_check_mark: *Registry P0 and P1*: nothing open.")

    def test_groups_by_repository(self):
        message = render([
            issue(number=1, repository="registry"),
            issue(number=2, repository="terraform-to-pulumi-registry-pipeline"),
        ], NOW)
        self.assertIn("*pulumi/registry* (1)", message)
        self.assertIn("*pulumi/terraform-to-pulumi-registry-pipeline* (1)", message)

    def test_links_each_issue(self):
        message = render([issue(number=12079)], NOW)
        self.assertIn("• <https://github.com/pulumi/registry/issues/12079|#12079>", message)

    def test_escapes_titles(self):
        message = render([issue(title="<https://evil.example|Click here>")], NOW)
        self.assertNotIn("<https://evil.example|Click here>", message)
        self.assertIn("&lt;https://evil.example|Click here&gt;", message)

    def test_truncates_long_lists(self):
        message = render([issue(number=n) for n in range(25)], NOW)
        self.assertIn("• …and 5 more", message)

    def test_counts_every_issue_in_the_header(self):
        message = render([issue(number=n) for n in range(25)], NOW)
        self.assertIn("|25 open>", message)


if __name__ == "__main__":
    unittest.main()
