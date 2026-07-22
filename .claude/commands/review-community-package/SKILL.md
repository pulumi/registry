---
description: Deep review of a community package contribution PR (one that touches community-packages/package-list.json). Runs schema URL + naming + SDK presence + relative-asset checks against the upstream provider repo and surfaces a structured report. Optionally regenerates metadata locally and verifies plugin install in a sandboxed container. Use when a maintainer wants to thoroughly vet a community-package add before approving — including from a list of open PRs, not necessarily one already open in the browser.
---

# Review community package

**Use this when:** A PR adds (or removes) an entry in `community-packages/package-list.json`, and you want the full deep review before deciding to approve, request changes, or hand off.

Companion to `/pr-review` (which is generic). This one knows the community-package contract.

**CRITICAL — confirmation required for every user-visible action.** Any `gh` command that mutates state (`gh pr review`, `gh pr comment`, `gh pr close`, `gh pr merge`, `gh pr create`), any `git push`, any creation of a new PR, and any closing of the contributor's PR MUST be preceded by an explicit `AskUserQuestion` confirmation that shows the full command and the full message body to be posted. This rule applies even in `--dangerously-skip-permissions` / auto mode — the skill speaks publicly on the user's behalf and to outside contributors, so the bar is "user must see exactly what's about to be said before it's said." Read-only operations (`gh pr view`, `curl`, `gh api`, `gh release list`, `git status`, etc.) and local-only operations (writing the cached artifact, regenerating metadata into the worktree) do not require this confirmation.

References:

- `_common:review-criteria` — shared review rubric
- [`docs/adding-a-new-package.md`](../../docs/adding-a-new-package.md) — documented checklist
- See [`~/Documents/Context/community-package-reviews/`](file:///home/fausto/Documents/Context/community-package-reviews/) for the history doc this skill is derived from

---

## Workflow

Six phases, each tier costlier than the last. A failure in an earlier tier short-circuits the heavy ones; reviewer attention is the scarcest resource.

| # | Phase | What it covers | Cost | Always runs? |
|---|---|---|---|---|
| 1 | Triage | Pick the PR (named or from queue). | trivial | yes |
| 2 | Mechanical checks | Schema URL resolves, naming convention, semver tag, SDKs published, docs scannable. Deterministic. | cheap (gh + curl) | yes |
| 3 | Judgement | Is this worth onboarding? README, resource surface, maintainer profile, typo-squat, license, source smell test. | model judgement | yes |
| 4 | Heavy verification | Regenerate metadata locally; sandbox-install per the contributor's instructions. | minutes | only if 2 + 3 clean |
| 5 | Report + persist | Structured markdown to terminal + identical content cached to disk with YAML frontmatter. | trivial | yes |
| 6 | Act | Approve / request-changes / comment / continue-in-new-PR / merge. Update cached artifact in place. | reviewer choice | yes |

Phases 1–3 always run. Phase 4 auto-runs when 2 and 3 are clean and skips silently when either has Fail items. Phase 6 surfaces the action menu shaped by what the previous phases found.

---

## Step 1: Pick the PR

If invoked with a PR number, use it. Otherwise, list open PRs touching the package list and ask:

```bash
gh search prs --repo pulumi/registry --state open \
  --json number,title,author,createdAt,url \
  "in:title community" --limit 20
```

Also surface PRs whose diff includes `community-packages/package-list.json`:

```bash
gh pr list --repo pulumi/registry --state open --search "community-packages" \
  --json number,title,author,createdAt
```

Use `AskUserQuestion` to let the user pick. Single-select. Show 3–4 most recent; include "Other" for free-text PR number.

---

## Step 2: Parse the diff

```bash
gh pr view <num> --repo pulumi/registry --json title,author,baseRefName,headRefName,files,body
gh pr diff <num> --repo pulumi/registry
```

Extract added/removed entries from `community-packages/package-list.json`. For each added entry, capture `repoSlug` and `schemaFile`. Skip the rest of the steps for removals; just summarise the removal context.

---

## Step 3: Static checks (cheap, no auth beyond gh)

For each new entry, in parallel where possible:

1. **Repo exists, not archived, recent activity.**
   ```bash
   gh api repos/<slug> --jq '{archived, pushed_at, default_branch, html_url}'
   ```
2. **`v`-prefixed semver release exists.**
   ```bash
   gh release list -R <slug> -L 5
   ```
3. **Schema URL resolves.** Construct `https://raw.githubusercontent.com/<slug>/<default_branch>/<schemaFile>` and `curl -fIs` it. If 404, this is a hard fail — the nightly pipeline will silently produce nothing.
4. **Schema name matches repo convention.** Download the schema, parse JSON, assert `schema.name == repo_short.removeprefix("pulumi-")`. Mismatches cause the publish pipeline to mis-name everything (#9784 dex class of bug).

---

## Step 4: Docs + SDK presence (network, still cheap)

**Docs.** Fetch both at the latest release tag AND on the default branch (catch the "exists in main but not in the tagged release" case and vice versa):

```
https://raw.githubusercontent.com/<slug>/<tag>/docs/{_index,installation-configuration}.md
https://raw.githubusercontent.com/<slug>/<default_branch>/docs/{_index,installation-configuration}.md
```

If either file is absent in both locations: **Fail** (required by [the checklist](https://github.com/pulumi/registry/blob/master/docs/adding-a-new-package.md)).

For files that exist, check:

- **Code example per language the contributor claims** (TS/Python/Go/C#). Grep for language fences.
- **Relative-asset / raw-HTML scan** (the [#9934 stackit class](https://github.com/pulumi/registry/pull/9934#discussion_r2827346034)). High-confidence Fail patterns: `<img src="(?!https?://)`, markdown `!\[.*?\]\((?!https?://|#)`, raw HTML `<a href="(?!https?://|#|mailto:)`. Lower confidence Warn (rewritten on IDP by [pulumi-service#38008](https://github.com/pulumi/pulumi-service/pull/38008), not on the public registry): plain markdown `]\((?!https?://|#|mailto:)` with `./` or `/` prefix.

**SDK presence probe.** First read the schema's `language` map to learn which languages the provider declares it supports — the keys are the truth, not "all four." Then probe only the declared languages. Classify findings:

- **Fail** when a language is declared in the schema but no SDK can be found at any conventional name.
- **Fail** when a declared language has a published SDK but the version lags the provider release by more than one minor version (the schema promises an API surface that the SDK doesn't expose).
- **Warn** for obvious packaging bugs in published names (e.g. doubled tokens like `Foo.Pulumi.Bar.Bar` — a `tfgen` config bug that affects every future release).
- Languages NOT in the schema's `language` map are not probed and not reported (the contributor chose not to ship them, fine).

Naming variants to try, in parallel:

- npm: `pulumi-<name>`, `@pulumi/<name>`, `@<owner>/<name>`, `@<owner>/pulumi-<name>`
- PyPI: `pulumi-<name>`, `pulumi_<name>`, `pulumi-<owner>-<name>`, `pulumi_<owner>_<name>`
- pkg.go.dev: `github.com/<owner>/<repo>` (module root) AND `github.com/<owner>/<repo>/sdk/go/<name>`
- NuGet: `curl -s 'https://azuresearch-usnc.nuget.org/query?q=<owner>&prerelease=false' | jq '.data[].id'` and grep for `<name>` (case-insensitive).

---

## Step 5: Provider sanity & light code review (cheap, runs by default)

The PR is a four-line addition, but the _provider_ being onboarded is arbitrary third-party code that will be advertised on `pulumi.com/registry`. Reviewer needs a basic "is this worth registering and is it not malicious" pass. Not a security audit — a smell test.

For the upstream repo at the release tag:

1. **Fetch the README** (`README.md` at repo root). Summarise in one sentence: what does this provider actually let users do? Then ask the broader question: **is this useful to have in the public registry?** Failure modes to flag (non-exhaustive):
    - Deployment of one specific branded product/site, not a general-purpose service abstraction — counter-example: [PR #10246 `pdfzus-merge-site`](https://github.com/pulumi/registry/pull/10246), closed and redirected to the [private registry](https://www.pulumi.com/docs/idp/concepts/private-registry/).
    - Wraps an API that already has a well-maintained Pulumi provider (duplicate / vanity).
    - Wraps a deprecated, dead, or barely-used upstream service (low audience).
    - Toy / demo / tutorial code masquerading as a provider.
    - Empty or near-empty provider — schema with one trivial resource and no real surface area.
    - Internal-looking abstractions that should be a private registry entry, not a public one.

    Suggest the [private registry](https://www.pulumi.com/docs/idp/concepts/private-registry/) when the value is real but the audience is one org.

2. **List resource types from the schema** (already downloaded in step 3). One-sentence summary of what's provisioned. Flag if there's exactly one resource and it's named after a product, or if resource types do something obviously off-mission (e.g. a "Provider" resource that takes arbitrary shell commands).

3. **Maintainer sanity.** `gh api users/<owner>` → account creation date, public repo count, company/org if any. Flag accounts created in the last 30 days or with no other repos — not blocking, just worth a second look.

4. **Typo-squat check.** Compare the provider name (`schema.name`) to common high-traffic providers (`aws`, `azure`, `gcp`, `kubernetes`, `docker`, etc.) using simple edit distance. Flag names within edit distance ≤ 2 of a major provider name.

5. **Quick scan of the provider source** for obvious red flags. Identify the main provider entrypoint (Go: `provider/cmd/pulumi-resource-<name>/main.go`, TS: similar). Read it. Look for: `curl … | sh` in install/build scripts, base64-decode-and-exec, unexplained HTTP calls to non-provider domains, eval-of-remote-content. Skip deep security analysis — surface only obvious smells.

6. **License file present?** `gh api repos/<slug>/license` — flag if 404.

For small providers (README + main entrypoint < ~10 kB total), do this inline. For larger ones, use the Agent tool (`subagent_type: "Explore"`) so the dumps don't bloat the parent context. Return only the verdict + one-line-per-finding either way.

---

## Step 6: regenerate metadata locally (auto-runs when Steps 2, 3, and 4 are clean)

Runs the doc-generation step the nightly bot would run, so you can preview the rendered output and so the generated files are available for "Close and continue" in Step 9.

Skipped automatically if Step 2, 3, or 4 produced any `Fail`, or if Step 3 / 5 produced a safety Warn (typo-squat, suspicious source, missing license, etc.). In those cases the next action is to triage the findings, not materialise files for a PR that may not happen.

```bash
make bin/resourcedocsgen
./bin/resourcedocsgen metadata from-github \
  --repoSlug '<slug>' \
  --schemaFile '<schemaFile>' \
  --version '<tag>'
```

Output goes under `themes/default/content/registry/packages/<name>/` and `themes/default/data/registry/packages/<name>.yaml`. Don't commit. Offer to `make serve` so the user can click through the package's preview page locally.

**Publisher allowlist (manual, easy to miss).** `resourcedocsgen` does NOT register the publisher — it only writes the `publisher:` value into the generated YAML. Read that value and confirm it exists as a key in [`tools/resourcedocsgen/pkg/publishers/publisher-names.json`](../../../tools/resourcedocsgen/pkg/publishers/publisher-names.json). If it's absent, add `"<Publisher>": "<slug>"` — single-username publishers map to themselves (`"AdrianSilaghi": "AdrianSilaghi"`), org publishers use a lowercase slug (`"OVHcloud": "ovhcloud"`). A publisher missing from this file fails the `Test Live Registry Publish` check with `publisher "<X>" not found, please add it to publisher-names.json`, so this edit MUST be bundled into the Close-and-continue commit alongside the regenerated metadata.

---

## Step 7: sandboxed install (auto-runs when Steps 2, 3, 4, and 5 are clean; container-only)

**CRITICAL: The install command MUST run inside a throwaway container. Never invoke it on the host, never with a host filesystem mount, never with the host network namespace. Downloaded provider binaries are untrusted code from an arbitrary GitHub user.** If `podman` and `docker` are both unavailable, skip this step and record it as Skipped — do not fall back to running on the host under any circumstance.

**Gating:** runs automatically only if Steps 2, 3, 4, and 5 all came back clean (zero `Fail`; no `Warn` in the "is this useful / is this safe" category — typo-squat, brand-new maintainer account with no other signal, suspicious source patterns, missing license, or "shipping one company's stack"-style flags). On any such finding, skip and record as Skipped with a one-line reason; ask the user to triage the findings first.

Verifies that the install instructions the contributor wrote in `installation-configuration.md` actually work. Don't hardcode a command — different provider types use different installation paths:

- **Native** typically uses `pulumi plugin install resource <name> <ver> --server <url>`
- **Bridged** may use `pulumi plugin install` against a different URL, or instruct installing only the language SDK
- **Component** often uses `pulumi package add github.com/<owner>/<repo>@<tag>`
- **Some** publish to package managers only and have no `pulumi plugin install` step at all

Process:

1. Read the contributor's `installation-configuration.md` fetched in step 4.
2. Extract shell code blocks (` ```bash`, ` ```sh`, ` ```shell`, or unlabelled if it starts with `pulumi`/`go`/`npm`/`pip`/`dotnet`).
3. Identify the install-side commands (skip configuration commands like `pulumi config set`, skip language-SDK package-manager commands unless they're the only install path — those install to a project, not globally, and need a temp project to run in).
4. Show the user the extracted commands and ask which to run. Default to selecting the `pulumi plugin install` / `pulumi package add` one if present.
5. Run the selected command(s) inside a throwaway container — no host filesystem mount, no exec of any downloaded binary by us.

Container invocation (adapt the inner `sh -c` to the actual commands):

```bash
podman run --rm \
  --network=bridge \
  --read-only \
  --tmpfs /tmp \
  --tmpfs /root \
  -e PULUMI_HOME=/tmp/pulumi \
  -e PULUMI_ACCESS_TOKEN=dummy \
  pulumi/pulumi-base \
  sh -c '<extracted install command> \
         && find /tmp/pulumi/plugins -maxdepth 2 -type f -exec file {} \;'
```

For `pulumi package add` (component flow), you'll need a throwaway project. Initialise inside the container:

```bash
podman run --rm \
  --network=bridge \
  --tmpfs /tmp --tmpfs /root \
  -e PULUMI_HOME=/tmp/pulumi \
  pulumi/pulumi-base \
  sh -c 'mkdir /tmp/proj && cd /tmp/proj \
         && pulumi new typescript --name probe --description probe --stack dev --yes --offline \
         && <extracted pulumi package add command> \
         && ls -la sdks/'
```

What this proves: the commands the contributor told users to run actually work. What it does NOT prove: that the resulting provider runs correctly — we deliberately don't exec the downloaded binary.

If `podman` is unavailable, fall back to `docker` with the same args. If neither is available, skip with a note in the report. If the contributor's docs have no extractable install commands at all, flag that as a Warn — install instructions are required by the documented checklist.

### Manual reproduction (always include in the report)

The container exits and disappears on `--rm`. Even when the skill's verdict is "install succeeded," a human reviewer may want to poke around — different terminal output, different download path, "did it actually fetch v1.2.3 or the latest tag," etc. The Step 8 report and cached artifact MUST include, in addition to the captured stdout/stderr:

1. **Verbatim re-run command** (one-shot, identical to what the skill ran):

    ```bash
    podman run --rm --network=bridge --read-only \
      --tmpfs /tmp --tmpfs /root \
      -e PULUMI_HOME=/tmp/pulumi \
      pulumi/pulumi-base \
      sh -c '<the exact extracted command>'
    ```

2. **Interactive inspect variant** (named container, no `--rm`, drops to a shell after the install runs so the reviewer can `ls`, `file`, `cat`, etc.):

    ```bash
    podman run -it --name registry-review-pr<num> --network=bridge \
      --tmpfs /tmp --tmpfs /root \
      -e PULUMI_HOME=/tmp/pulumi \
      pulumi/pulumi-base \
      sh -c '<the exact extracted command>; exec sh'
    ```

    Followed by cleanup hint: `podman rm -f registry-review-pr<num>`.

3. **Already-ran-and-want-to-look-again variant** (if the reviewer opts at runtime to skip `--rm` — see below):

    ```bash
    podman exec -it registry-review-pr<num> sh
    # or to copy something out:
    podman cp registry-review-pr<num>:/tmp/pulumi/plugins ./inspect-plugins/
    ```

The skill's default invocation uses `--rm` (clean exit). Offer a `--keep-container` flag for when the reviewer wants to keep the post-install state around for poking; in that mode, replace `--rm` with `--name registry-review-pr<num>` and print the inspect commands at the end of Step 7.

---

## Step 8: Report

One block, **emitted to chat AND written verbatim to the cached artifact** (the cache adds YAML frontmatter; the body is identical). Grouped by status. **Every PR number, repo slug, package name, file path, line citation, and external reference MUST be a clickable markdown link** (terminals and editors render `[text](url)`). Build URLs explicitly: `https://github.com/pulumi/registry/pull/<num>`, `https://github.com/<slug>/blob/<tag>/<path>#L<line>`, `https://www.npmjs.com/package/<name>`.

**Outcome name canon** (use these exact strings; map prose ↔ machine in one place):

| Prose | Frontmatter `recommended_outcome` / `action_taken` |
|---|---|
| Close and continue | `continue` |
| Ask for submitter changes | `request-changes` |
| Close with reasoning | `close-unsuitable` |

**Summary paragraph** must call out any structural findings that constrain the outcome mechanic but don't change the verdict — most commonly "fork CI did not run, so the merge path must be close-and-continue." Put these in Summary, not buried in Warn, so the Recommended-outcome rationale lands.

**Skipped lines** should cite the gating step that triggered the skip: `Step 7 (sandboxed install): Step 4 fail (no installation-configuration.md to extract install commands from)`.

Shape:

```markdown
# Review: [PR #<num>](url) — <title>

Provider: [<slug>](url) @ [<tag>](url)
Author: [@<author>](url) · Tracking issue: [@<reporter>](url) on [#<issue>](url)
Head ref: `<ref>` on <fork-owner or "pulumi/registry"> · sha [`<sha7>`](url)

## Summary
<2–3 sentence human read: what the provider does, key blocker if any>

## Checks
### Pass
- <link-rich bullets>
### Warn
- <link-rich bullets>
### Fail
- <link-rich bullets>
### Skipped
- Step <N>: <why>

## Recommended outcome
**<Close and continue | Ask for submitter changes | Close with reasoning>** — <one-sentence rationale>

## Details
<one collapsed `<details>` per step that produced verbose output: schema excerpt, SDK probe table, maintainer profile, container output>
```

Write this body verbatim to the cached artifact (Cached artifact format section below), prepending the YAML frontmatter. Update `pr-<num>/latest.md` symlink.

---

## Step 9: Action menu

Three outcomes. Recommended option is determined by the report state:

| State | Recommended outcome |
|---|---|
| All Pass + non-safety Warns | **Close and continue** |
| Mechanical (Step 2) or doc (Step 4) Fail the submitter must fix | **Ask for submitter changes** |
| Judgement / safety (Step 3 / 5) Fail | **Close with reasoning** |

Use `AskUserQuestion` to surface all three with the recommended one first. After the action runs, update the cached artifact's frontmatter `action_taken` field in place.

Use the tone templates from [`pr-review:references:message-templates`](../pr-review/references/message-templates.md) for the actual comment bodies — warm for external contributors, technical for bot-authored PRs.

If the skill detects an internal-contributor PR (author has write+ access to `pulumi/registry`, e.g. rgharris/seanyeh/jasonwzm), bail out at the top of Step 9 with a one-line "use /pr-review instead — this skill is for community contributions" and exit. Internal PRs have valid direct-merge paths this skill shouldn't disable.

### Comment style (applies to every PR comment, review body, and close message the skill generates)

- **Target under 150 words.** A reviewer or contributor should be able to skim it without scrolling. Long rationales belong in the cached artifact, not the PR comment.
- **No em-dashes.** Use a period, comma, parenthesis, or colon instead. Never `—`.
- **No opening preamble.** Skip "Hi @x, thanks for the submission, looks like a real provider, we'd like to onboard it." Open with the first ask or the operative sentence. A contributor opening a PR knows we're considering it.
- **No closing flourish.** Skip "Happy to help if anything is unclear." The maintainer can add that by hand if they want.
- **One line per ask.** Format: `Asked thing (one-clause justification or link)`. Don't explain why each ask matters in 3 sentences. If the link to the relevant doc/spec carries the why, that's enough.
- **Group only when there are multiple asks.** For a single ask, no `Required:` / `Nice to have:` headers; just write the line.
- **Links over prose.** A clickable link to the spec/file/PR is shorter and more verifiable than a paraphrase.

Apply this to all three outcomes' generated bodies. If the message is over the word target after a first draft, cut explanation prose before cutting asks.

---

### Outcome A: Close and continue

The "we want this, we'll take it over" path. Always used when onboarding (matches boris's `#9743 → #10010`, `#9393 → #9934`, `#9455 → #9784` historical pattern): it's the only way to get CI on fork-blocked PRs, it bundles any Step 6 files so the package ships immediately rather than waiting for the next nightly, and it puts a maintainer on the merging commit.

Steps:

1. Confirm the user wants to proceed (one-line preview: new branch name, files to be committed).
2. Locally:
    - `git checkout master && git pull --ff-only`
    - `git checkout -b <gh-username>/community-package-<provider-name>-continued`
    - Re-apply the `package-list.json` entry from the contributor's PR.
    - If Step 6 produced files (`themes/default/content/registry/packages/<name>/*`, `themes/default/data/registry/packages/<name>.yaml`), `git add` them.
    - If the publisher was missing from `tools/resourcedocsgen/pkg/publishers/publisher-names.json` and you added it in Step 6, `git add` that too. This is a manual edit (`resourcedocsgen` never writes it), and omitting it fails the `Test Live Registry Publish` check on the continuation PR.
    - Commit. Message: `Continue #<orig-num>: add <provider-name> to community package list`.
    - Push: `git push -u origin HEAD`.
3. Open the new PR: `gh pr create --base master --title "Add <provider-name> to community package list" --body "<body>"`. Body should include:
    - "Continues [#<orig-num>](https://github.com/pulumi/registry/pull/<orig-num>) from @<contributor>."
    - Link back to the original tracking issue if present.
    - One-line note on what's bundled (just the entry, or entry + regenerated metadata).
4. On the contributor's PR:
    - `gh pr comment <orig-num> --body "Thanks @<contributor> — continuing in #<new-num> so we can land the metadata together: https://github.com/pulumi/registry/pull/<new-num>"`
    - `gh pr close <orig-num>`
5. Update cached artifact: `action_taken: continue`, add `continued_in: <new-num>` to frontmatter.

Do not push to the contributor's fork even when `maintainer_can_modify` is set on their PR — the historical convention is the new-PR pattern, and it keeps the contributor's branch as their own record.

---

### Outcome B: Ask for submitter changes

The PR stays open. Used when there's work only the submitter can do — fixes in their upstream provider repo (broken Go importBasePath, missing SDK, relative image paths in their `docs/_index.md`), fixes in their PR itself (wrong `schemaFile` path, typo in `repoSlug`), or release-side work (cut a new `v`-tagged release).

Steps:

1. Draft a `--request-changes` body from the cached report's Fail and actionable-Warn items. One section per ask, each with: what's wrong, where (clickable link to the offending file:line in their repo), what to do.
2. `gh pr review <num> --request-changes --body "<body>"`.
3. Update cached artifact: `action_taken: request-changes`.

Don't include the full report — only the items that require submitter action. Pass items and non-blocking Warns aren't useful to them here; they're in the cached artifact if you need to recall.

---

### Outcome C: Close with reasoning

The "this isn't a fit for the public registry" path. Used for #10246-class submissions (deploys one branded product, duplicates an existing well-maintained provider, wraps a dead upstream, malware-ish smell from Step 5).

Steps:

1. Draft a closing comment from the report's Step 3 / 5 Fail items. Warm tone, explain why, offer alternatives (most commonly: [private registry](https://www.pulumi.com/docs/idp/concepts/private-registry/)).
2. `gh pr comment <num> --body "<body>"`, then `gh pr close <num>`.
3. Update cached artifact: `action_taken: close-unsuitable`.

The closing comment is the contributor's only signal that the work didn't land — make it human, specific, and constructive about what would change the answer (if anything would).

---

## Cached artifact format

Location: `${XDG_CACHE_HOME:-$HOME/.cache}/pulumi-registry-reviews/pr-<num>/<head-sha>.md` (portable on Linux and macOS — `mkdir -p` handles both). Maintain `pr-<num>/latest.md` as a symlink to the most recent sha file. Body is the Step 8 report verbatim; only the frontmatter is cache-only.

```yaml
---
pr: <int>
title: <string>
repo: pulumi/registry
author: <gh-login>
base_ref: master
head_ref: <branch>
head_repo_owner: <gh-login>      # = "pulumi" if branch is in this repo
head_sha: <full sha>
reviewed_at: <ISO 8601 UTC>
skill_version: 1
fork_ci_ran: <bool>              # false if checks rollup conclusion is null
package_entries:
  - repo_slug: <owner>/<repo>
    schema_file: <path>
    action: add                  # or "remove"
verdict: <pass | pass-with-warnings | fail | incomplete>
recommended_outcome: <continue | request-changes | close-unsuitable>
action_taken: null               # set on Step 9 execution: continue | request-changes | close-unsuitable
continued_in: null               # set to <int> on "continue" outcome
---
```

When the user picks an action in Step 9, update only `action_taken`, `reviewed_at`, and (for continue) `continued_in` in place. Don't append new sections; the per-sha file is the audit trail.

---

## Notes

- For first-party (`pulumi/pulumi-*`) providers, prefer `/pr-review` — they use `repository_dispatch`, not the community package list.
- The container image `pulumi/pulumi-base` is ~200 MB. First run pulls it.
- Step 7's container bounds binary blast radius but doesn't make a suspect package worth onboarding — Steps 3 and 5 still must be clean.
