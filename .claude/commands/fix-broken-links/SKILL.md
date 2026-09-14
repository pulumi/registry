---
name: fix-broken-links
description: Fix broken links reported by the weekly registry link checker. Reads .broken-links.json, triages each entry to a fix strategy (S3 redirect, source edit, exclusion, upstream issue, or third-party/not actioned), and opens an auditable PR. Invoked by the check-links workflow; not user-invocable.
user-invocable: false
---

# Fix Broken Links

You are fixing the broken links found by the weekly link checker (`scripts/link-checker/check-links.js`, run against `https://www.pulumi.com/registry`). Work through every reported entry, pick the right strategy for each, open one auditable PR, and file upstream issues only where Pulumi owns the source.

The registry is different from `pulumi/docs` in one way that drives most of the triage below: **almost every page a reader sees is generated from a provider repository, not authored here.** A package's overview (`_index.md`) and `installation-configuration.md` are fetched verbatim from `docs/` in the upstream provider repo at publish time, and its API docs are generated from the provider schema. A broken link on one of those pages can't be fixed in this repo. Whether anyone can fix it depends on who maintains the upstream repo.

## Input

Read `.broken-links.json` from the repo root. Shape:

```json
{
  "generated": "2026-09-14T15:00:00.000Z",
  "internal": [{ "source": "...", "destination": "...", "reason": "HTTP_404" }],
  "external": [{ "source": "...", "destination": "...", "reason": "HTTP_404" }]
}
```

- `source` — the live page the entry was found **on**.
- `destination` — the URL that failed.
- `reason` — BLC reason code (`HTTP_404`, `HTTP_410`, `HTTP_400`, `ERRNO_ENOTFOUND`, `HTML could not be retrieved`, …).
- `internal` — `destination` is on `pulumi.com`; `external` — third-party.

If both lists are empty, do nothing (the workflow won't invoke you in that case, but be defensive).

## Verify each link before fixing it

**The checker produces false positives — never fix a link without confirming it's actually broken.** The crawler hits live URLs with a bot user agent, so it routinely mis-flags links that are fine in a browser: bot-protected sites, transient `5xx`/timeouts, rate-limiting, auth walls, and pages that block automated clients. `excludeAcceptable()` already filters the common transient reasons, but plenty slip through — especially in the `external` list.

For **every** reported link, confirm it's genuinely broken before touching anything:

1. **Re-check the destination** — `WebFetch` (or `curl -sIL`) the URL. If it resolves, it's a false positive.
2. **Weigh the reason code.** Internal `HTTP_404`/`HTTP_410` on a `pulumi.com` path is almost always real. External failures, and codes like `HTTP_400`, `HTTP_403`, `HTTP_401`, `HTTP_429`, `HTTP_503`, or connection/timeout errors, are frequently bot protection or transient — treat as suspect until confirmed.
3. **If it's a false positive, do not edit, redirect, or exclude it.** Leave the link as-is and record it in the PR audit trail under a **False positives / not actioned** section, with the URL, the reason code, and one line on why it's fine (e.g. "resolves in browser; 403 is Cloudflare bot protection"). Optionally suggest it as a future exclusion-list candidate, but don't add it unless it's a recurring offender.

Only links you've confirmed broken proceed to triage below.

## Skip links already being handled (deduplication)

Broken links persist for weeks, so the same ones surface in consecutive runs while a fix is already in flight. **Before actioning a confirmed-broken link, check whether it's already covered** — never open a second PR for a fix that's pending, or file a duplicate issue.

For each confirmed-broken link, search for prior work:

1. **Open PRs in `pulumi/registry`.** List candidates with `gh pr list --state open --search "broken link" --json number,title,url,headRefName` (prior runs use the `fix/broken-links-*` branch pattern). For a likely match, confirm with `gh pr view <n>` / `gh pr diff <n>` that it actually covers this link (same destination URL, redirect, or exclusion entry).
2. **Existing issues, before filing a new one.** For an upstream item that would become an issue in a provider repo, search that repo first — `gh issue list --repo pulumi/<repo> --state open --search "<url-or-path>"` (or `gh search issues`). If an open issue already tracks it, treat it as a duplicate and reuse that issue's link; do **not** file another.

Classify each confirmed-broken link as **actionable** (no existing PR/issue covers it) or **duplicate** (one does — capture the PR/issue URL). Then:

- **All duplicates** (every confirmed-broken link is already covered): do **not** create a branch, PR, or any issue. Write a Slack summary to `.broken-links-pr.txt` saying so and listing each link with its existing PR/issue link, e.g. `:link: 5 broken links — all already tracked (no new PR): <url> → #12345, …`. Stop here.
- **Some actionable, some duplicates:** action only the non-duplicates (open the PR, file new issues). **Omit duplicates from the fixes**, but list them in a short **Already tracked** section in the PR description (one line each, linking the existing PR/issue). Don't re-file or re-fix them.

## Mapping a live URL back to its source

`source` and `destination` are live `https://www.pulumi.com/registry/...` URLs. Before you can pick a strategy you need to know **who owns the page the link lives on**:

1. Strip the origin to get the path, e.g. `/registry/packages/aws/installation-configuration/`.
2. The package name is the segment after `/registry/packages/`. Its metadata is `themes/default/data/registry/packages/<name>.yaml`; read `publisher` and `repo_url` from it.
3. Find the source of the **page** (`source`):
   - `/registry/packages/<name>/` → `themes/default/content/registry/packages/<name>/_index.md`
   - `/registry/packages/<name>/installation-configuration/` → `themes/default/content/registry/packages/<name>/installation-configuration.md`
   - Both files carry a `# WARNING: this file was fetched from <raw URL>` header and an `edit_url:` front-matter key pointing at the upstream file. **That upstream file is the source; the copy here is generated and must never be hand-edited** (the publish workflow overwrites it on the next release).
   - `/registry/packages/<name>/api-docs/...` → generated from the provider schema by `resourcedocsgen`; there is no editable source in either repo for the prose, only the schema in the provider.
   - `/registry/packages/<name>/how-to-guides/...` → these pages were retired in favor of `/dev/examples/`; anything still linking to one is a redirect job (see `scripts/redirects/how-to-guides-to-dev-redirects.txt`).
   - `/registry/` index and any other non-package page → hand-maintained here: Grep under `themes/default/layouts/`, `themes/default/content/`, and `themes/default/data/` for the link text.
   - `https://github.com/pulumi/pulumi` → the checker crawls this page directly as an additional route; links on it are GitHub's own chrome, not ours. Anything here is an exclusion-list entry at most.
4. Classify the **owner** of the page:
   - **This repo** — templates, data files, hand-written content.
   - **First-party upstream** — `repo_url` is under `https://github.com/pulumi/`. Pulumi maintains it; a link fix there is a normal PR/issue for the provider team.
   - **Third-party upstream** — any other `repo_url` (community providers: `pulumiverse/*`, vendor orgs, individuals). Pulumi does not control that content.

## Per-link triage

Apply the first row that matches.

| Situation | Strategy | Mechanism |
|---|---|---|
| Internal `destination` is a retired registry path that has a clear successor (renamed API-docs module, `how-to-guides` page that moved to `/dev/examples/`, deleted sub-page with an obvious parent) | **S3 redirect** | Add a line to the topic-appropriate file in `scripts/redirects/` |
| Internal `destination` points at a `/docs/...` page that moved | **Out of scope (docs)** | File an issue in `pulumi/docs` asking for an alias; don't add a registry redirect for a docs path |
| Broken link in content **this repo** owns (layouts, data YAML such as `repo_url`/`logo_url`, hand-written pages) | **Edit at source** | Fix it in place |
| Broken link on a generated page whose upstream is **first-party** (`repo_url` under `github.com/pulumi/`) | **Upstream issue** | File an issue in the provider repo pointing at the exact `docs/` file and line (via `edit_url`); never edit the generated copy here |
| Broken link on a generated page whose upstream is **third-party** | **Third-party / not actioned** | Record it in the PR description's **Third-party packages** section with the package, `repo_url`, and the broken URL. Do not file issues in repos outside the `pulumi` org and do not edit the generated copy |
| Dead / transient / bot-protected **external** link that recurs | **Exclusion list** | Add the URL to `getDefaultExcludedKeywords()` in `check-links.js` with an inline `//` comment naming the reason + package |
| Redirect **loop** or a redirect line pointing at a dead target | **Repair** | Fix or remove the stale line in `scripts/redirects/` |
| Anything else unfixable here (a package whose `repo_url` is gone, a schema-level link) | **Out of scope** | File an issue in `pulumi/registry` describing it and link it from the PR; a vanished upstream repo is a delisting candidate, so say so in the issue |

### Mechanics

**S3 redirect.** Append one line to the topic-appropriate file in `scripts/redirects/` (`how-to-guides-to-dev-redirects.txt` for retired how-to guides, `aws-v6-redirects.txt` for AWS module renames, `removed-examples-redirects.txt` for deleted examples; create a new `<topic>-redirects.txt` when nothing fits and name it by what it redirects, not by date). Format is `source-path|destination-url`, one per line. The source is the path **without** a leading slash and ends in `/index.html`; the destination is a full URL:

```
registry/packages/aws/how-to-guides/old-guide/index.html|https://www.pulumi.com/dev/examples/old-guide/
```

Redirects are applied to the S3 bucket by `scripts/ci/make-s3-redirects.sh` on the next production deploy. A link to `/registry/packages/<pkg>/api-docs/...` is excluded from the crawl, but a redirect for one is still valid when the destination moved.

**Edit at source.** Only for files this repo owns. Internal links use full root-relative paths (`/registry/packages/aws/`); never `../`.

**Upstream issue (first-party).** `gh issue create --repo pulumi/<provider-repo>` with a title like `Broken link in docs/installation-configuration.md` and a body that gives the registry page URL, the broken `destination`, the reason code, the `edit_url` of the upstream file, and one line on what the link should probably point at. One issue per upstream file, listing every broken link on it, not one per link.

**Exclusion list.** Add the URL to the array returned by `getDefaultExcludedKeywords()` in `scripts/link-checker/check-links.js`, grouped with similar entries, with an inline comment naming the reason and the package, matching the existing style:

```js
"https://example.com/gone", // <package>: 404, no replacement (#<issue>)
```

Exclusions are substring matches first and globs second, and a glob `*` must match at least one character — so `https://example.com/thing*` does **not** match `https://example.com/thing`. Drop the trailing `*` unless you specifically want only sub-paths.

## Editing guardrails

- **Never hand-edit** anything under `themes/default/content/registry/packages/` or `themes/default/data/registry/packages/` to fix a link. Those are generated; the fix lives upstream or in a redirect.
- Do not file issues, open PRs, or comment in repositories outside the `pulumi` GitHub org. Third-party breakage is recorded, not chased.
- `yarn run lint` (eslint + prettier over `scripts/`) and `make lint-markdown` must both pass before you open the PR. Do not run `make build`: it needs every provider and 32 GB+ of RAM.
- Keep the PR to link fixes. Don't refactor the checker, the redirect files, or anything else you pass through.

## Output

If **every** confirmed-broken link is a duplicate (see deduplication above), skip all of the below: open no branch, PR, or issue, write the "all already tracked" Slack summary to `.broken-links-pr.txt`, and stop. Otherwise, for the actionable (non-duplicate) entries:

1. Create a branch `fix/broken-links-<date>` (date from the workflow, e.g. `fix/broken-links-2026-09-14`).
2. Make the fixes, grouping related changes into clear commits.
3. File a GitHub issue (`gh issue create`) for every **new** first-party upstream or out-of-scope item (reuse the existing issue for duplicates).
4. Run `yarn run lint` and `make lint-markdown`; fix anything they surface.
5. Open a **ready** (non-draft) PR to `master`.
6. Write the final PR URL plus a one-line summary to `.broken-links-pr.txt` for the workflow's Slack step, e.g.: `:link: Fixed 4 broken links — <PR URL> (2 redirects, 1 exclusion, 1 upstream issue filed; 3 third-party, 5 false positives)`

If nothing was fixable in this repo but you filed upstream issues, still open the PR only when it contains a change (a redirect, an exclusion, a source edit). Otherwise skip the PR and write a summary to `.broken-links-pr.txt` that links the issues you filed and counts the third-party and false-positive entries, so the Slack post still tells the reader what happened.

## PR description contract (auditability)

The reviewer must be able to audit every decision without re-deriving it. Include:

- **A table or list of every broken link** → the strategy applied → one line of non-obvious reasoning (why a redirect vs. an upstream issue, why excluded, etc.).
- A **Verification** section: confirm `yarn run lint` and `make lint-markdown` passed, and note that each link was re-checked before fixing.
- A **False positives / not actioned** section listing every reported link you confirmed was actually fine, with its reason code and why (so the reviewer knows it was checked, not missed).
- A **Third-party packages** section listing every broken link on a community-maintained package page, with the package name and `repo_url`, so a human can decide whether to nudge the maintainer.
- An **Already tracked** section, when any link was skipped as a duplicate: one succinct line per link linking the existing PR or issue that covers it.
- An **Upstream / filed issues** section linking every issue you created in a first-party provider repo, `pulumi/docs`, or `pulumi/registry`.
