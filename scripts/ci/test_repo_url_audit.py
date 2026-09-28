#!/usr/bin/env python3
"""Unit tests for repo_url_audit.py"""

import json
import tempfile
import unittest
from pathlib import Path
from unittest import mock

from repo_url_audit import MARKER, check, load_listed_slugs, load_packages, render, slug_of, sync_issue


def package(name="doppler", repo_url="https://github.com/pulumiverse/pulumi-doppler", bridged=False):
    return {"name": name, "repo_url": repo_url, "bridged": bridged}


def lookup_from(repos):
    """A lookup over a {requested slug: repo object} map; anything absent is a 404."""
    return lambda slug: repos.get(slug)


class TestSlugOf(unittest.TestCase):
    def test_parses_owner_and_repo(self):
        self.assertEqual(slug_of("https://github.com/pulumi/pulumi-aws"), "pulumi/pulumi-aws")

    def test_tolerates_trailing_slash_and_git_suffix(self):
        self.assertEqual(slug_of("https://github.com/pulumi/pulumi-aws/"), "pulumi/pulumi-aws")
        self.assertEqual(slug_of("https://github.com/pulumi/pulumi-aws.git"), "pulumi/pulumi-aws")

    def test_rejects_non_repo_urls(self):
        self.assertIsNone(slug_of("https://gitlab.com/pulumi/pulumi-aws"))
        self.assertIsNone(slug_of("https://github.com/pulumi"))
        self.assertIsNone(slug_of("https://github.com/pulumi/pulumi-aws/tree/main"))
        self.assertIsNone(slug_of(""))


class TestCheck(unittest.TestCase):
    def test_current_repo_is_not_a_finding(self):
        lookup = lookup_from({"pulumiverse/pulumi-doppler": {"full_name": "pulumiverse/pulumi-doppler"}})
        self.assertIsNone(check(package(), lookup, set()))

    def test_case_only_difference_is_not_a_finding(self):
        lookup = lookup_from({"pulumiverse/pulumi-doppler": {"full_name": "Pulumiverse/Pulumi-Doppler"}})
        self.assertIsNone(check(package(), lookup, set()))

    def test_transferred_repo_is_moved(self):
        lookup = lookup_from({"pulumiverse/pulumi-doppler": {"full_name": "DopplerHQ/pulumi-doppler"}})
        finding = check(package(), lookup, {"pulumiverse/pulumi-doppler"})
        self.assertEqual(finding["kind"], "moved")
        self.assertEqual(finding["moved_to"], "DopplerHQ/pulumi-doppler")
        self.assertTrue(finding["in_package_list"])

    def test_package_list_match_ignores_case(self):
        lookup = lookup_from({"Pulumiverse/Pulumi-Doppler": {"full_name": "DopplerHQ/pulumi-doppler"}})
        mixed_case = package(repo_url="https://github.com/Pulumiverse/Pulumi-Doppler")
        # load_listed_slugs lowercases the list, so check has to lowercase the package's slug.
        self.assertTrue(check(mixed_case, lookup, {"pulumiverse/pulumi-doppler"})["in_package_list"])
        self.assertFalse(check(mixed_case, lookup, set())["in_package_list"])

    def test_archived_repo_wins_over_moved(self):
        lookup = lookup_from({"rancher/terraform-provider-rke": {
            "full_name": "rancher-archives/terraform-provider-rke", "archived": True}})
        finding = check(package("rke", "https://github.com/rancher/terraform-provider-rke", bridged=True),
                        lookup, set())
        self.assertEqual(finding["kind"], "archived")
        self.assertEqual(finding["moved_to"], "rancher-archives/terraform-provider-rke")

    def test_archived_in_place_has_no_new_location(self):
        lookup = lookup_from({"pulumiverse/pulumi-doppler": {
            "full_name": "pulumiverse/pulumi-doppler", "archived": True}})
        finding = check(package(), lookup, set())
        self.assertEqual(finding["kind"], "archived")
        self.assertIsNone(finding["moved_to"])

    def test_404_is_missing(self):
        self.assertEqual(check(package(), lookup_from({}), set())["kind"], "missing")

    def test_unparseable_url_is_invalid_and_never_looked_up(self):
        lookup = mock.Mock()
        self.assertEqual(check(package(repo_url="https://example.com/x"), lookup, set())["kind"], "invalid")
        lookup.assert_not_called()


class TestLoading(unittest.TestCase):
    def test_skips_deprecated_and_flags_bridged(self):
        with tempfile.TemporaryDirectory() as tmp:
            d = Path(tmp)
            (d / "a.yaml").write_text("publisher: Acme\nrepo_url: https://github.com/acme/a\n"
                                      "schema_file_url: https://x.cloudfront.net/schemas/registry.opentofu.org/acme/a/1/schema.json\n")
            (d / "b.yaml").write_text("publisher: DEPRECATED\nrepo_url: https://github.com/acme/b\n")
            (d / "c.yaml").write_text('publisher: Acme\nrepo_url: "https://github.com/acme/c"\n')
            packages = load_packages(d)
        self.assertEqual([p["name"] for p in packages], ["a", "c"])
        self.assertTrue(packages[0]["bridged"])
        self.assertFalse(packages[1]["bridged"])
        self.assertEqual(packages[1]["repo_url"], "https://github.com/acme/c")

    def test_listed_slugs_are_lowercased(self):
        with tempfile.TemporaryDirectory() as tmp:
            path = Path(tmp) / "package-list.json"
            path.write_text(json.dumps({"include": [{"repoSlug": "EventStore/pulumi-eventstorecloud"}]}))
            self.assertEqual(load_listed_slugs(path), {"eventstore/pulumi-eventstorecloud"})


class TestRender(unittest.TestCase):
    def finding(self, kind, **overrides):
        return {"name": "doppler", "slug": "pulumiverse/pulumi-doppler", "repo_url": "", "bridged": False,
                "in_package_list": False, "moved_to": None, "kind": kind, **overrides}

    def test_starts_with_marker_and_includes_only_populated_sections(self):
        body = render([self.finding("moved", moved_to="DopplerHQ/pulumi-doppler")])
        self.assertTrue(body.startswith(MARKER))
        self.assertIn("## Moved", body)
        self.assertNotIn("## Archived upstream", body)
        self.assertIn("[`DopplerHQ/pulumi-doppler`](https://github.com/DopplerHQ/pulumi-doppler)", body)

    def test_says_where_each_fix_goes(self):
        body = render([
            self.finding("moved", name="a", moved_to="x/a"),
            self.finding("moved", name="b", moved_to="x/b", in_package_list=True),
            self.finding("archived", name="c", bridged=True, in_package_list=True),
        ])
        self.assertIn("| a | `pulumiverse/pulumi-doppler` | [`x/a`](https://github.com/x/a) | YAML |", body)
        self.assertIn("| YAML + `package-list.json` |", body)
        # Bridged wins: the pipeline regenerates the YAML regardless of package-list.json.
        self.assertIn("| c | `pulumiverse/pulumi-doppler` | unchanged | `pulumi/terraform-to-pulumi-registry-pipeline`"
                      " (dynamically bridged) |", body)

    def test_archived_guidance_sends_bridged_delistings_through_the_pipeline_first(self):
        body = render([self.finding("archived", bridged=True)])
        self.assertIn("remove it from `watched-providers` before delisting it here", body)

    def test_is_deterministic(self):
        findings = [self.finding("missing"), self.finding("invalid", repo_url="")]
        self.assertEqual(render(findings), render(findings))
        self.assertIn("`(empty)`", render(findings))


def response(payload=None):
    r = mock.Mock()
    r.json.return_value = payload
    return r


class TestSyncIssue(unittest.TestCase):
    def session(self, open_issues):
        s = mock.Mock()
        s.get.return_value = response(open_issues)
        s.post.return_value = response({"number": 99})
        s.patch.return_value = response({})
        return s

    def test_opens_an_issue_when_none_is_open(self):
        s = self.session([])
        sync_issue(s, MARKER + "\nbody")
        s.post.assert_called_once()
        self.assertEqual(s.post.call_args.kwargs["json"]["body"], MARKER + "\nbody")
        s.patch.assert_not_called()

    def test_rewrites_the_open_issue_in_place(self):
        s = self.session([{"number": 5, "body": MARKER + "\nold"}])
        sync_issue(s, MARKER + "\nnew")
        s.patch.assert_called_once()
        self.assertEqual(s.patch.call_args.kwargs["json"], {"body": MARKER + "\nnew"})
        s.post.assert_not_called()

    def test_leaves_a_current_issue_alone(self):
        s = self.session([{"number": 5, "body": MARKER + "\nsame"}])
        sync_issue(s, MARKER + "\nsame")
        s.patch.assert_not_called()
        s.post.assert_not_called()

    def test_ignores_issues_and_prs_without_the_marker(self):
        s = self.session([{"number": 1, "body": "unrelated"},
                          {"number": 2, "body": MARKER, "pull_request": {}}])
        sync_issue(s, MARKER + "\nbody")
        s.post.assert_called_once()

    def test_closes_the_open_issue_once_nothing_is_left(self):
        s = self.session([{"number": 5, "body": MARKER + "\nold"}])
        sync_issue(s, None)
        s.post.assert_called_once()  # the closing comment
        self.assertEqual(s.patch.call_args.kwargs["json"], {"state": "closed", "state_reason": "completed"})

    def test_does_nothing_when_clean_and_no_issue_is_open(self):
        s = self.session([])
        sync_issue(s, None)
        s.post.assert_not_called()
        s.patch.assert_not_called()


if __name__ == "__main__":
    unittest.main()
