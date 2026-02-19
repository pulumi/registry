# BUILD-AND-DEPLOY.md — pulumi/registry

This document describes the build, test, and deployment system for the `pulumi/registry` repository.

---

## Table of Contents

1. [Overview](#overview)
2. [Architecture Overview](#architecture-overview)
3. [Local Development](#local-development)
4. [Build System](#build-system)
   - [4.1 Makefile Targets](#41-makefile-targets)
   - [4.2 Hugo Build](#42-hugo-build)
   - [4.3 resourcedocsgen Tool](#43-resourcedocsgen-tool)
   - [4.4 mktutorial Tool](#44-mktutorial-tool)
   - [4.5 CI Build Script](#45-ci-build-script-scriptscibuilds)
5. [GitHub Actions Workflows](#github-actions-workflows)
   - [5.1 pull-request.yml — PR Validation + Preview Deploy](#51-pull-requestyml--pr-validation--preview-deploy)
   - [5.2 push.yml — Production Build + Deploy](#52-pushyml--production-build--deploy)
   - [5.3 Other Workflows](#53-other-workflows)
6. [Deployment Infrastructure](#deployment-infrastructure)
   - [6.1 Pulumi IaC](#61-pulumi-iac)
   - [6.2 AWS Resources](#62-aws-resources)
   - [6.3 S3 Preview Bucket Lifecycle](#63-s3-preview-bucket-lifecycle)
   - [6.4 Custom Redirects](#64-custom-redirects)
7. [Registry Publication (push-registry.py)](#registry-publication-push-registrypy)
8. [Testing Strategy](#testing-strategy)
   - [8.1 Go Unit Tests](#81-go-unit-tests)
   - [8.2 Go Linting](#82-go-linting)
   - [8.3 Markdown Linting](#83-markdown-linting)
   - [8.4 Script Linting](#84-script-linting)
   - [8.5 Provider API Docs Tests](#85-provider-api-docs-tests)
   - [8.6 Browser Tests (Cypress)](#86-browser-tests-cypress)
   - [8.7 Link Checking](#87-link-checking)
9. [Environment & Secret Management](#environment--secret-management)
   - [9.1 Pulumi ESC](#91-pulumi-esc)
   - [9.2 Environment Variables Reference](#92-environment-variables-reference)
   - [9.3 Tool Version Matrix](#93-tool-version-matrix)
10. [Scheduled & Maintenance Tasks](#scheduled--maintenance-tasks)
11. [Troubleshooting](#troubleshooting)
12. [Pulumi Cloud Service Integration](#pulumi-cloud-service-integration)
   - [12.1 Two Parallel Systems](#121-two-parallel-systems)
   - [12.2 Pulumi Cloud Registry API](#122-pulumi-cloud-registry-api)
   - [12.3 How push-registry.py Bridges Them](#123-how-push-registrypy-bridges-them)
   - [12.4 Cross-Stack References Used at Build Time](#124-cross-stack-references-used-at-build-time)
   - [12.5 Pulumi Cloud Roles Summary](#125-pulumi-cloud-roles-summary)

---

## Overview

### Quick Reference: Common Commands

| Task | Command |
|---|---|
| Install all dependencies | `make ensure` |
| Build theme assets (CSS/JS) | `make build-assets` |
| Local development server | `make serve` |
| Full local dev (Hugo + asset watch) | `make serve-all` |
| Generate API docs for one provider | `make api-docs/<package>` |
| Run Go tests | `make test` |
| Run all linters | `make lint` |
| Run Markdown linter | `make lint-markdown` |
| Run Go linter | `make lint-go` |
| Run provider API docs tests | `make test_provider_api_docs` |
| Run browser tests | `make run-browser-tests` |
| Check links | `make check_links` |
| Full CI push (production) | `make ci_push` |
| Full CI pull request | `make ci-pull-request` |
| Bucket cleanup | `make ci_bucket_cleanup` |
| Clean build artifacts | `make clean` |

### Key Concepts

- **Hugo**: Static site generator that produces the final HTML from templates and content files.
- **Pulumi ESC (Environments, Secrets, and Configuration)**: Used for OIDC-based secret exchange; no long-lived secrets are stored directly in GitHub.
- **resourcedocsgen**: A Go tool (in `tools/resourcedocsgen/`) that generates provider API documentation from Pulumi provider schemas.
- **S3 origin bucket model**: Each build produces its own uniquely-named S3 bucket. A Pulumi IaC program reads a metadata file to determine which bucket to point CloudFront at. This means each PR commit gets its own preview URL, and production deploys atomically swap the CloudFront origin.
- **mktutorial**: A Go tool (in `tools/mktutorial/`) that generates how-to guide content from the `pulumi/examples` repository.

### Prerequisites

The canonical tool versions are tracked in `mise.toml`:

| Tool | Version |
|---|---|
| Go | 1.23 |
| Node.js | 20 |
| Yarn | 1.22.22 |
| Hugo | 0.135 (extended) |
| golangci-lint | 2.1.6 |
| yq | latest |

Install and manage all tools via [mise](https://mise.jdx.dev/):

```bash
mise trust && mise install
```

---

## Architecture Overview

```
                   ┌─────────────────────────────────────┐
                   │        pulumi/registry repo          │
                   │                                      │
                   │  themes/default/data/registry/       │
                   │    packages/*.yaml  ─────────────┐   │
                   │                                  ▼   │
                   │  tools/resourcedocsgen/ ──► generates API docs
                   │                                  │   │
                   │  themes/default/content/         │   │
                   │    registry/packages/ ◄──────────┘   │
                   │                                  │   │
                   │  Hugo build ◄────────────────────┘   │
                   │       │                              │
                   │       ▼                              │
                   │  public/  (built site)               │
                   └─────────────────────────────────────┘
                              │
                              │  scripts/ci/sync.sh
                              ▼
              ┌─────────────────────────────┐
              │  AWS S3 Origin Bucket        │
              │  registry-{env}-origin-      │
              │  {build-id}                  │
              └──────────────┬──────────────┘
                             │
                             │  Pulumi IaC (infrastructure/)
                             │  reads origin-bucket-metadata.json
                             ▼
              ┌─────────────────────────────┐
              │  AWS CloudFront Distribution │
              │  (per environment)           │
              └──────────────┬──────────────┘
                             │
                             ▼
                   pulumi.com/registry
```

### Two Environments

| Environment | AWS Account ID | Pulumi Stack | Purpose |
|---|---|---|---|
| testing | 571684982431 | `pulumi/registry/testing` | PR previews and test deployments |
| production | 388588623842 | `pulumi/registry/production` | Live site at pulumi.com/registry |

### Component Roles

| Component | Role |
|---|---|
| Hugo 0.135 (extended) | Static site generation from templates and content |
| resourcedocsgen | Generates provider API reference docs from Pulumi schemas |
| mktutorial | Generates how-to guides from `pulumi/examples` |
| Pulumi IaC | Manages AWS resources; reads metadata file to update CloudFront origin |
| Algolia | Search index; updated as part of production deploys via `scripts/search/main.js` |
| AWS S3 | Hosts built site content as a static website origin |
| AWS CloudFront | CDN serving the site; origin pointed at S3 bucket by Pulumi IaC |
| AWS SSM Parameter Store | Maps commit SHAs to their corresponding S3 origin buckets |

### Multi-Repo Touchpoints

- **`pulumi/pulumi-*` provider repos**: Trigger `publish-provider-update.yml` via `repository_dispatch` when a new provider version is released.
- **`pulumi/examples`**: Source of how-to guide content; pulled nightly by `update-tutorials.yml`.

---

## Local Development

### Tool Installation

```bash
# Trust mise config and install all declared tools
mise trust && mise install
```

### First-Time Setup

```bash
# 1. Install Node.js dependencies (yarn packages, etc.)
make ensure

# 2. Build theme CSS and JavaScript assets
make build-assets

# 3. Start local dev server
make serve
```

`make serve` runs:

```
hugo serve --buildDrafts --buildFuture --ignoreVendorPaths="github.com/pulumi/registry/**/*"
```

`REF_NOT_FOUND` Hugo warnings are suppressed automatically.

### Hot-Reload Development

To watch both Hugo content changes and theme asset changes simultaneously:

```bash
make serve-all
```

This uses `concurrently` to run the Hugo server and the Yarn asset watcher (`yarn --cwd ./themes/default/theme run start`) in parallel. If either process exits, both are killed.

To watch only theme assets (without serving Hugo):

```bash
make serve-assets
```

### Generating API Docs for One Package

```bash
make api-docs/<package>
```

For example, to generate docs for the `aws` provider:

```bash
make api-docs/aws
```

This runs `bin/resourcedocsgen docs registry` for the named package, reading its YAML from `themes/default/data/registry/packages/aws.yaml` and writing generated content to `content/registry/packages/aws/` (at the repository root — note this differs from the path used by `scripts/ci/build.sh` in CI, which writes to `themes/default/content/registry/packages/`).

The Makefile uses `.SECONDEXPANSION` and a `.make/` sentinel directory so that `make api-docs/aws` is a no-op if neither the YAML file nor the existing content has changed since the last run.

### Building a Single Provider Locally vs. Full Build

For a single provider:

```bash
make api-docs/<package>
```

For a full local build (all providers, all content):

```bash
make build
```

`make build` runs `make build-assets`, then `scripts/apply-fixes.js`, then Hugo.

> **Memory note**: A full build requires significant RAM. The CI build sets
> `NODE_OPTIONS=--max_old_space_size=8192` to give Node 8 GB of heap. For full
> local builds, consider setting this in your shell:
>
> ```bash
> export NODE_OPTIONS="--max_old_space_size=8192"
> ```
>
> 32 GB+ of RAM is recommended for a complete build.

---

## Build System

### 4.1 Makefile Targets

All targets are defined in the repository root `Makefile`.

| Target | Command / Description |
|---|---|
| `clean` | Runs `scripts/clean.sh` |
| `ensure` | Runs `scripts/ensure.sh` — installs Node/Go/Hugo deps |
| `build` | `build-assets` → `apply-fixes.js` → `hugo` |
| `build-assets` | `ensure` → `yarn --cwd ./themes/default/theme run build` |
| `serve` | `hugo serve --buildDrafts --buildFuture` (suppress REF_NOT_FOUND) |
| `serve-assets` | `yarn --cwd ./themes/default/theme run start` (asset watch) |
| `serve-all` | Concurrent Hugo serve + asset watch |
| `api-docs/<pkg>` | Build API docs for a single package via `resourcedocsgen` |
| `bin/resourcedocsgen` | Compile resourcedocsgen from `tools/resourcedocsgen/` |
| `bin/mktutorial` | Compile mktutorial from `tools/mktutorial/` |
| `lint` | `lint-go` + `lint-markdown` + `yarn run lint` |
| `lint-go` | `lint-resourcedocsgen` + `lint-mktutorial` |
| `lint-resourcedocsgen` | `golangci-lint run` in `tools/resourcedocsgen/` |
| `lint-mktutorial` | `golangci-lint run` in `tools/mktutorial/` |
| `lint-markdown` | `scripts/lint/lint-markdown.js` |
| `test` | `go test ./...` in `tools/resourcedocsgen/` |
| `test_provider_api_docs` | `ensure` + `build-assets` + `bin/resourcedocsgen` → `scripts/ci/run-provider-tests.sh` |
| `run-browser-tests` | `ensure` → `scripts/run-api-docs-tests.sh` |
| `check_links` | `ensure` → `yarn run check-links` |
| `ci-pull-request` | `ensure` → `scripts/ci/pull-request.sh` |
| `ci-pull-request-closed` | `scripts/ci/pull-request-closed.sh` |
| `ci-scheduled` | `scripts/ci/scheduled.sh` |
| `ci_push` | `ensure` → `scripts/ci/push.sh` |
| `ci_bucket_cleanup` | `scripts/ci/bucket-cleanup.sh` |

### 4.2 Hugo Build

Hugo reads from the following key locations:

| Path | Role |
|---|---|
| `config/` | Hugo configuration (base URL, params, etc.) |
| `themes/default/` | The default Hugo theme |
| `themes/default/content/` | Content files (Markdown) |
| `themes/default/data/` | Data files (YAML) |
| `themes/default/static/` | Static assets |
| `public/` | Hugo output directory |

**Build order for a full CI build:**

1. `make build-assets` — compiles theme JS and CSS via `yarn --cwd ./themes/default/theme run build`
2. `resourcedocsgen docs registry` — generates provider API docs into `themes/default/content/registry/packages/`
3. `node scripts/apply-fixes.js` — post-processes Hugo content (see script for details)
4. `hugo --minify --buildFuture --templateMetrics -e <environment>` — generates `public/`
5. `yarn run minify-css` (i.e., `node scripts/minify-css.js`) — purges and minifies CSS in `public/`

The Hugo base URL differs between build modes:

- **preview**: `http://<origin-bucket-prefix>-<build-id>.s3-website.us-west-2.amazonaws.com`
- **update** (production): uses production base URL; Hugo env is set to `production`

### 4.3 resourcedocsgen Tool

**Location**: `tools/resourcedocsgen/`

**Building**:

```bash
# Via Makefile (recommended — uses helpmakego for dependency tracking)
make bin/resourcedocsgen

# Manual
go build -C tools/resourcedocsgen -o ../../bin ./...
```

The compiled binary is placed at `bin/resourcedocsgen`.

**Subcommands**:

| Subcommand | Purpose |
|---|---|
| `docs registry` | Generate API docs for all (or one) provider packages |
| `pkgversion` | Check the latest published version of a community provider on GitHub |
| `metadata from-github` | Generate package YAML metadata from a GitHub-hosted provider repo |
| `metadata from-urls` | Generate package YAML metadata from explicit schema/index URLs |

**Key flags for `docs registry`**:

```bash
bin/resourcedocsgen docs registry \
    --baseDocsOutDir themes/default/content/registry/packages \
    --basePackageTreeJSONOutDir themes/default/static/registry/packages/navs \
    --baseSchemasOutDir themes/default/static/registry/packages \
    [<package-name>]
```

> **Note**: The paths above reflect `scripts/ci/build.sh` (CI runs). The Makefile's
> `make api-docs/<pkg>` target uses `./content/registry/packages` (repo root, without the
> `themes/default/` prefix) for `--baseDocsOutDir`, and similarly for the other `--base*` flags.

When `<package-name>` is omitted, all packages listed in `themes/default/data/registry/packages/` are processed.

**Input**: YAML files at `themes/default/data/registry/packages/*.yaml` — each file describes one provider (name, version, repo URL, schema file path, publisher, etc.)

**Output**:

- Generated docs at `themes/default/content/registry/packages/<pkg>/api-docs/`
- Package navigation JSON at `themes/default/static/registry/packages/navs/<pkg>.json`
- Schema JSON at `themes/default/static/registry/packages/<pkg>.json`

### 4.4 mktutorial Tool

**Location**: `tools/mktutorial/`

**Building**:

```bash
make bin/mktutorial
```

**Purpose**: Generates how-to guide (tutorial) content from the `pulumi/examples` repository for the following clouds: `aws-apigateway`, `aws`, `classic-azure` (mapped to `azure`), `azure` (mapped to `azure-native`), `gcp`, and `kubernetes`.

**Output**: Content written to `themes/default/content/registry/packages/<cloud>/how-to-guides/`.

Versioned packages (`aws-v6`, `azure-native-v2`) are also cleaned of stale tutorials.

**Used only in CI** via `scripts/ci/mktutorial.sh`. Not used in local builds.

### 4.5 CI Build Script (`scripts/ci/build.sh`)

This is the master build script for CI runs. It accepts one argument: `preview` or `update`.

**Steps executed by `build.sh`**:

1. Calls `make build-assets` (theme CSS/JS compilation).
2. Fetches the Pulumi conversion service URL from the `pulumi/tf2pulumi-service/production` stack output (`PULUMI_CONVERT_URL`).
3. Fetches the Pulumi AI WebSocket URL from `pulumi/pulumigpt-api/corp` stack output (`PULUMI_AI_WS_URL`).
4. Reads Copilot (Atlas) API URL from Pulumi IaC config (`atlasUrl`).
5. Computes a `build_identifier` (for preview: `pr-<number>-<sha8>`; for push: `push-<sha8>`).
6. Sets asset bundle paths:
   - `CSS_BUNDLE=static/css/styles.<id>.css`
   - `JS_BUNDLE=static/js/bundle.min.<id>.js`
7. Compiles `resourcedocsgen`: `go build -o $GOPATH/bin/resourcedocsgen .`
8. Runs `resourcedocsgen docs registry --logtostderr` to generate all provider API docs.
9. Runs `node ./scripts/apply-fixes.js`.
10. Runs Hugo with `--minify --buildFuture --templateMetrics`:
    - `preview` mode: sets `HUGO_BASEURL` to the S3 website URL and uses `-e preview`
    - `update` mode: uses `-e production`
11. Runs `yarn run minify-css` to purge and minify CSS.

---

## GitHub Actions Workflows

All workflow files live in `.github/workflows/`.

### Workflow Summary

| File | Name | Trigger |
|---|---|---|
| `pull-request.yml` | Pull request | PR to `master` |
| `push.yml` | Build and deploy | Push to `master` |
| `testing-deploy.yml` | Build and deploy - test environment | `workflow_dispatch` |
| `pull-request-closed.yml` | Close pull request | PR closed |
| `check-go.yml` | Check Go | `workflow_call` (reusable) |
| `check-links.yml` | Scheduled jobs: Check links | Every Monday 3:00 PM UTC |
| `run-browser-tests.yml` | Scheduled jobs: Run browser tests | Daily 2:00 PM UTC |
| `generate-package-metadata.yml` | Check for Community Package Updates | Daily 5:30 AM + 5:30 PM UTC |
| `publish-provider-update.yml` | provider docs build | `repository_dispatch` |
| `bucket-cleanup.yml` | Scheduled jobs: Bucket cleanup | Daily 3:00 PM UTC |
| `update-tutorials.yml` | Scheduled jobs: Update How To Guides | Daily 3:00 PM UTC |
| `export-repo-secrets.yml` | Export secrets to ESC | `workflow_dispatch` |
| `add-triage-label.yml` | Add triage label to new issues | Issue opened / reopened |
| `add-to-project.yml` | Add issues to project | Issue opened / reopened |

### 5.1 `pull-request.yml` — PR Validation + Preview Deploy

**Trigger**: Pull request to `master`

**ESC environment**: `github-secrets/pulumi-registry` (OIDC, no long-lived secrets exported automatically)

**Flow**:

```
PR opened / committed
        │
        ├── resourcedocsgen (calls check-go.yml)
        │       ├── lint (golangci-lint)
        │       └── test (go test ./...)
        │
        ├── mktutorial (calls check-go.yml)
        │       ├── lint
        │       └── test
        │
        ├── lint-markdown
        │       └── yarn install → make lint-markdown
        │
        ├── lint-scripts
        │       └── yarn install → yarn run lint
        │
        ├── test-live-publish
        │       └── uv run push-registry.py --dry-run
        │
        ├── test-provider-api-docs
        │       └── make ensure build-assets → make test_provider_api_docs
        │
        ├── preview  (skipped for fork PRs; skipped for automation/tfgen-provider-docs label)
        │       ├── Fetch ESC secrets
        │       ├── Install Node 22, Go 1.23, Hugo 0.135
        │       ├── Validate community-packages/package-list.json
        │       ├── Configure AWS credentials → assume testing account role
        │       └── make ci-pull-request
        │               ├── scripts/ci/validate-packages.sh
        │               ├── scripts/ci/build.sh preview
        │               └── scripts/ci/sync.sh preview
        │                       ├── Create / reuse S3 bucket
        │                       ├── aws s3 sync public/ → bucket
        │                       ├── Run browser tests (Cypress smoke test)
        │                       ├── Write origin-bucket-metadata.json
        │                       └── Post PR comment with preview URL
        │
        └── sentinel (depends on all jobs above)
                └── Writes "Sentinel" GitHub status check = success
```

**Skipping preview for fork PRs**: Fork PRs are excluded at the workflow level. The `preview` job has an `if:` condition that only allows it to run when `github.event.pull_request.head.repo.full_name == github.repository`, so the job is never scheduled for PRs from forks. `pull-request.sh` also contains a defensive credential check (for `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`, and `PULUMI_ACCESS_TOKEN`) as a fallback, but this code path is not expected to be reached in practice.

**The `sentinel` job**: The sentinel job creates a GitHub status check called `"Sentinel"` only after all required jobs pass. This single required check simplifies branch protection rules. The sentinel job runs for non-fork PRs and for `repository_dispatch` events (condition: `github.event_name == 'repository_dispatch' || github.event.pull_request.head.repo.full_name == github.repository`).

**Key environment variables in `preview` job**:

| Variable | Source |
|---|---|
| `PULUMI_ACCESS_TOKEN` | ESC output |
| `GITHUB_TOKEN` | GitHub Actions default |
| `PULUMI_STACK_NAME` | GitHub Actions variable |
| `NODE_OPTIONS` | Hardcoded: `--max_old_space_size=8192` |
| `ALGOLIA_APP_ID` | GitHub Actions variable |
| `ALGOLIA_APP_SEARCH_KEY` | GitHub Actions variable |

**Runner**: `pulumi-service-ubuntu-24.04-16core` (large runner required for the full build)

### 5.2 `push.yml` — Production Build + Deploy

**Trigger**: Push to `master`

**ESC environment**: `github-secrets/pulumi-registry` (OIDC via `id-token: write`)

**Flow**:

```
Push to master
        │
        └── build job
                ├── Fetch ESC secrets (OIDC)
                ├── Install Node 22, Go 1.23, Hugo 0.135
                ├── Checkout (using PULUMI_BOT_TOKEN for private module access)
                ├── Configure AWS credentials
                │       └── Assume arn:aws:iam::388588623842:role/ContinuousDelivery
                ├── Install Pulumi CLI
                ├── make ci_push
                │       ├── scripts/ci/login.sh
                │       │       └── pulumi login && pulumi -C infrastructure stack select $PULUMI_STACK_NAME
                │       ├── scripts/ci/build.sh update
                │       ├── scripts/ci/sync.sh update
                │       │       ├── Create / reuse S3 bucket
                │       │       ├── aws s3 sync public/ → bucket
                │       │       ├── Run browser tests (Cypress smoke test)
                │       │       └── Write origin-bucket-metadata.json
                │       ├── scripts/generate-search-index.sh
                │       │       └── Upload search-index.json to S3
                │       ├── node scripts/await-in-progress.js
                │       ├── scripts/ci/run-pulumi.sh update
                │       │       └── pulumi -C infrastructure update --yes
                │       │           (reads origin-bucket-metadata.json, updates CloudFront)
                │       └── scripts/ci/make-s3-redirects.sh
                │               └── Apply 301 redirects from scripts/redirects/
                ├── Archive origin-bucket-metadata.json as artifact
                └── uv run push-registry.py
                        └── Publish new provider versions to registry service
```

**Runner**: `pulumi-service-ubuntu-24.04-16core`

**Key environment variables**:

| Variable | Source |
|---|---|
| `PULUMI_ACCESS_TOKEN` | ESC output |
| `GITHUB_TOKEN` | GitHub Actions default |
| `ALGOLIA_APP_ID` | GitHub Actions variable |
| `ALGOLIA_APP_SEARCH_KEY` | GitHub Actions variable |
| `ALGOLIA_APP_ADMIN_KEY` | ESC output |
| `PULUMI_STACK_NAME` | GitHub Actions variable |
| `PULUMI_DOCS_STACK_NAME` | GitHub Actions variable |
| `DEPLOYMENT_ENVIRONMENT` | GitHub Actions variable |
| `NODE_OPTIONS` | Hardcoded: `--max_old_space_size=8192` |

### 5.3 Other Workflows

#### `testing-deploy.yml` — Manual Test Environment Deploy

Identical to `push.yml` but triggered manually via `workflow_dispatch` and deploys to the testing environment (account `571684982431`, role `arn:aws:iam::571684982431:role/ContinuousDelivery`). Uses Go 1.21.x (note: older than production).

#### `pull-request-closed.yml` — PR Preview Cleanup

**Trigger**: PR closed (any PR from this repo, not forks)

Runs `make ci-pull-request-closed` which calls `scripts/ci/pull-request-closed.sh`:

1. Fetches all commits associated with the closed PR from GitHub API.
2. For each commit, looks up the associated S3 bucket in SSM Parameter Store.
3. If a bucket is found and accessible, deletes it (`aws s3 rb ... --force`).
4. Posts a PR comment noting that previews have been removed.

Requires: ESC secrets + AWS credentials (testing account role via `AWS_CI_ROLE_ARN` from ESC).

#### `check-go.yml` — Reusable Go Lint + Test

A reusable `workflow_call` workflow. Called by `pull-request.yml` for both `tools/resourcedocsgen/` and `tools/mktutorial/`.

Jobs:

- **lint**: Sparse checkout → golangci-lint v2.1.6 with `--config .golangci.yml`
- **test**: Sparse checkout → `go test ./... -v`

#### `check-links.yml` — Link Validation

**Trigger**: Every Monday at 3:00 PM UTC (cron: `0 15 * * MON`); also `workflow_dispatch`

Runs `make check_links` which calls `yarn run check-links`, which runs `node scripts/link-checker/check-links.js "https://www.pulumi.com/registry" 2` (2 retries on failure). Broken links are reported to the `#registry-ops` Slack channel.

Node version: 22.x; Hugo 0.135.0 installed but not explicitly used.

#### `run-browser-tests.yml` — Scheduled Browser Tests

**Trigger**: Daily at 2:00 PM UTC; also `workflow_dispatch`

Runs `make run-browser-tests` on a `pulumi-ubuntu-8core` runner. Assumes the production AWS role (`388588623842:role/ContinuousDelivery`) to be able to reach the live site.

Node version: 22.x; Hugo 0.135.0 installed.

#### `generate-package-metadata.yml` — Nightly Community Package Check

**Trigger**: Daily at 5:30 AM UTC and 5:30 PM UTC; also `workflow_dispatch`

**Flow**:

1. `generate-packages-list` job: Runs `python generate_package_list.py` in `community-packages/` to build a matrix of community provider repos to check.
2. `check-for-package-update` job (matrix, max-parallel: 1): For each provider, runs `resourcedocsgen pkgversion` to check if a new version is available. If so, runs `resourcedocsgen metadata from-github` to generate updated metadata and opens a PR via `.github/actions/new-provider-version-pr`.
3. PRs are skipped if an open PR already exists for that provider (deduplication check via `list_pull_requests` in `scripts/common.sh`).

#### `publish-provider-update.yml` — Provider Doc Update via Repository Dispatch

**Trigger**: `repository_dispatch` with event types `resource-provider` or `push-provider-update`

Used by first-party Pulumi provider repos to trigger documentation regeneration when a new provider version is released.

| Event type | Use case | Required inputs |
|---|---|---|
| `resource-provider` | GitHub-hosted provider (Pulumi repo) | `project-shortname`, `ref` (version tag) |
| `push-provider-update` | Opaque provider (no assumed GitHub structure) | `project-shortname`, `schema-url`, `index-url` |

For `resource-provider`: Calls `resourcedocsgen metadata from-github` → creates a PR.
For `push-provider-update`: Downloads schema from `schema-url`, extracts version from schema, calls `resourcedocsgen metadata from-urls` → creates a PR.

#### `bucket-cleanup.yml` — Remove Stale S3 Preview Buckets

**Trigger**: Daily at 3:00 PM UTC; also `workflow_dispatch`

Runs `make ci_bucket_cleanup` which calls `scripts/ci/bucket-cleanup.sh`, which in turn calls `scripts/ci/remove-buckets.sh push` and `scripts/ci/remove-buckets.sh pr`.

For each deletable bucket (associated with a closed PR):

1. Applies a lifecycle policy: all objects expire after 1 day.
2. Adds a `CleanupStarted` tag with a timestamp.
3. If the bucket has been in cleanup state for 48+ hours, attempts `aws s3 rb`.
4. Gives up (with an error) if cleanup has been stalled for 7+ days.

Runs in the production environment (`388588623842:role/ContinuousDelivery`). Node 18.x / Go 1.20.x (older versions pinned in this workflow).

#### `update-tutorials.yml` — Regenerate Tutorials from Examples

**Trigger**: Daily at 3:00 PM UTC; also `workflow_dispatch`

1. Checks out both `pulumi/registry` and `pulumi/examples` (into `examples/`).
2. Runs `scripts/ci/mktutorial.sh $GITHUB_WORKSPACE/examples`.
3. Opens or updates a PR on branch `tutorials/refresh` via `peter-evans/create-pull-request@v7`.

Auto-merge is currently disabled (commented out in the workflow).

#### `export-repo-secrets.yml` — Sync GitHub Secrets → ESC

**Trigger**: `workflow_dispatch` only

Exports all GitHub repository secrets (except `EXPORT_SECRETS_PRIVATE_KEY`) to the `github-secrets/pulumi-registry` ESC environment using the `pulumi/esc-export-secrets-action`.

#### `add-triage-label.yml` — Auto-label New Issues

**Trigger**: Issue opened or reopened

Adds the `needs-triage` label to all new issues automatically.

#### `add-to-project.yml` — Add Issues to GitHub Project

**Trigger**: Issue opened or reopened

Adds new issues to the Pulumi Docs GitHub project (project #79) using `PULUMI_BOT_GHA_MARKETING` token from ESC.

---

## Deployment Infrastructure

### 6.1 Pulumi IaC

**Location**: `infrastructure/`

**Runtime**: Node.js (`runtime: nodejs`, per `infrastructure/Pulumi.yaml`)

**Purpose**: Manages AWS resources. After a successful build and sync, the CI pipeline runs `pulumi -C infrastructure update --yes`, which reads `origin-bucket-metadata.json` to determine the newly-built S3 bucket and updates the CloudFront origin accordingly.

**Stacks**:

| Stack file | Stack name | AWS account | Purpose |
|---|---|---|---|
| `Pulumi.yaml` | (base config) | — | Project definition |
| `Pulumi.testing.yaml` | testing | 571684982431 | Test environment |
| `Pulumi.production.yaml` | production | 388588623842 | Production environment |

**Key config values** (both stacks):

| Key | Purpose |
|---|---|
| `registry:pathToOriginBucketMetadata` | `../origin-bucket-metadata.json` — where Pulumi reads the newly built bucket |
| `registry:websiteLogsBucketName` | S3 bucket for CloudFront access logs |
| `registry:e2eTestsBucketName` | S3 bucket for Cypress test results |
| `registry:copilotUrl` | Copilot feature URL |
| `registry:atlasUrl` | Atlas AI feature URL |

### 6.2 AWS Resources

**IAM roles**: Both accounts have an `arn:aws:iam::<account>:role/ContinuousDelivery` role that CI assumes via OIDC or static key → role assumption.

**S3 origin buckets**: Public static website buckets named by the `origin_bucket_prefix()` + `build_identifier()` functions in `scripts/ci/common.sh`:

- Prefix: `registry-<deployment-env>-origin` (e.g., `registry-testing-origin`)
- Identifier for PRs: `pr-<number>-<sha8>` (e.g., `registry-testing-origin-pr-42-a1b2c3d4`)
- Identifier for pushes: `push-<sha8>`

Each bucket is created as an S3 static website with `index.html` / `404.html` and public-read ACL.

**AWS SSM Parameter Store**: Each commit's bucket name is stored at:

```
/registry/commits/<full-sha>/bucket
```

This mapping is used by `pull-request-closed.sh` to find and delete preview buckets when a PR closes.

**CloudFront distributions**: Managed by the Pulumi IaC program. After each production push, Pulumi reads `origin-bucket-metadata.json` and updates the CloudFront origin to point to the new S3 bucket.

### 6.3 S3 Preview Bucket Lifecycle

```
PR commit pushed
       │
       ▼
scripts/ci/sync.sh preview
  1. aws s3 mb registry-testing-origin-pr-<N>-<sha8>
  2. Enable static website hosting
  3. aws s3 sync public/ → bucket (--delete)
  4. Run Cypress smoke tests
  5. Write origin-bucket-metadata.json
  6. aws ssm put-parameter /registry/commits/<sha>/bucket = <bucket-name>
  7. Post PR comment: "Your preview is ready at http://..."

PR merged or closed
       │
       ▼
scripts/ci/pull-request-closed.sh
  1. List all commits for the PR (GitHub API)
  2. For each commit: aws ssm get-parameter /registry/commits/<sha>/bucket
  3. If bucket exists: aws s3 rb s3://<bucket> --force
  4. aws ssm delete-parameter /registry/commits/<sha>/bucket
  5. Post PR comment: "Site previews have been removed."

Daily bucket-cleanup.yml (3:00 PM UTC)
  - Catches any buckets not cleaned up by pull-request-closed.sh
  - Applies 1-day lifecycle expiration policy
  - Deletes bucket after 48+ hours in cleanup state
```

### 6.4 Custom Redirects

**Source files**: `scripts/redirects/` — pipe-delimited text files with `key | location` entries.

**Applied by**: `scripts/ci/make-s3-redirects.sh` as part of every production push (`make ci_push`).

**Mechanism**: For each redirect entry, creates an S3 object key with a `WebsiteRedirectLocation` header, causing S3 to return a proper 301 HTTP redirect rather than an HTML meta-refresh. This improves SEO and supports URL anchors.

---

## Registry Publication (`push-registry.py`)

**Location**: `scripts/ci/push-registry.py`

**Runtime**: Python 3 (via `uv run --with requests,pyyaml`)

**Invoked**:

- On every push to `master` (in `push.yml`, after the build+deploy completes)
- In dry-run mode on every PR (in `pull-request.yml`, `test-live-publish` job)

**What it does**:

1. Reads all YAML files from `themes/default/data/registry/packages/*.yaml`.
2. For each package:
   - Skips packages where `publisher == "DEPRECATED"`.
   - Skips packages whose name matches `azure-native-v*` (except `azure-native` itself) — these are aliases.
   - Skips packages whose name matches `aws-v<N>` — these are legacy versioned packages.
   - Calls the Pulumi registry API (`https://api.pulumi.com/api/registry/packages/{source}/{publisher}/{name}/versions/{version}`) to check if this version already exists.
   - If it does not exist (404): downloads the schema from the provider repo or `schema_file_url`, corrects the version field if needed, and calls `pulumi package publish`.
   - If `--installation-configuration` exists (`_installation-configuration.md`), passes it to `pulumi package publish`.
3. In `--dry-run` mode: prints the `pulumi package publish` command instead of running it.

**Required environment variable**: `PULUMI_ACCESS_TOKEN`

**Optional**: `PULUMI_BACKEND_URL` (defaults to `https://api.pulumi.com/api`)

**Package source**: Determined by `schema_url`:

- Contains `registry.opentofu.org` → `source = "opentofu"`
- Otherwise → `source = "pulumi"`

**Publisher lookup**: `tools/resourcedocsgen/pkg/publishers/publisher-names.json` maps display names (as in YAML) to canonical publisher IDs.

---

## Testing Strategy

### 8.1 Go Unit Tests

```bash
make test
# Runs: cd tools/resourcedocsgen && go test ./...
```

Also run in CI via `check-go.yml` for both `tools/resourcedocsgen/` and `tools/mktutorial/`.

### 8.2 Go Linting

```bash
make lint-go
# Runs golangci-lint in both:
#   tools/resourcedocsgen/
#   tools/mktutorial/
```

Config: `.golangci.yml` (repo root).

golangci-lint version: **v2.1.6** (specified in both `mise.toml` and `check-go.yml`).

### 8.3 Markdown Linting

```bash
make lint-markdown
# Runs: scripts/lint/lint-markdown.js
```

Node version used in CI: 23.x (in `pull-request.yml` `lint-markdown` job).

### 8.4 Script Linting

```bash
yarn run lint
# Runs: eslint scripts && prettier scripts --check
```

Lints all files under `scripts/` with ESLint and Prettier. Config defined in the root ESLint and Prettier config files.

Node version used in CI: 23.x.

To auto-fix formatting:

```bash
yarn run format
# Runs: prettier scripts --write
```

### 8.5 Provider API Docs Tests

```bash
make test_provider_api_docs
# Runs: scripts/ci/run-provider-tests.sh
```

Requires: `ensure`, `build-assets`, and `bin/resourcedocsgen` to have been built. Runs in PR CI as the `test-provider-api-docs` job.

Node version in CI: 23.x. Go version: stable (latest).

### 8.6 Browser Tests (Cypress)

```bash
make run-browser-tests
# Runs: scripts/run-api-docs-tests.sh
```

**Config**: `cypress.config.js` — default base URL is `http://localhost:1313`.

**Test location**: `cypress/` directory.

**Reporters**: `cypress-multi-reporters` (config in `reporter-config.json`).

Browser tests are run:

1. **As a smoke test inside `scripts/ci/sync.sh`** after each S3 deploy (both preview and production), using the deployed S3 website URL.
2. **Daily at 2:00 PM UTC** via `run-browser-tests.yml` against the live production site.

### 8.7 Link Checking

```bash
make check_links
# Runs: yarn run check-links
# Which runs: node scripts/link-checker/check-links.js "https://www.pulumi.com/registry" 2
```

- Fetches the sitemap from `https://www.pulumi.com/registry/sitemap.xml`.
- Excludes API docs pages, SDK reference pages, and install/versions pages.
- Uses `broken-link-checker` with `filterLevel: 1` and GET requests.
- Up to 2 retries if broken links are found.
- Reports broken links to the `#registry-ops` Slack channel via `SLACK_ACCESS_TOKEN`.
- Many known-flaky domains are excluded (LinkedIn, YouTube, Twitter, etc.).

Runs every Monday at 3:00 PM UTC in CI.

---

## Environment & Secret Management

### 9.1 Pulumi ESC

**Organization**: `pulumi`
**Environment**: `github-secrets/pulumi-registry`

All workflows use OIDC token exchange to authenticate with Pulumi ESC — no long-lived secrets are stored in GitHub Actions secrets directly (except for `AWS_ACCESS_KEY_ID` / `AWS_SECRET_ACCESS_KEY` used in the testing environment for PR preview deploys).

The ESC action is configured via workflow-level environment variables:

```yaml
ESC_ACTION_OIDC_AUTH: true
ESC_ACTION_OIDC_ORGANIZATION: pulumi
ESC_ACTION_OIDC_REQUESTED_TOKEN_TYPE: urn:pulumi:token-type:access_token:organization
ESC_ACTION_ENVIRONMENT: github-secrets/pulumi-registry
ESC_ACTION_EXPORT_ENVIRONMENT_VARIABLES: false  # (or a specific mapping)
```

The `export-repo-secrets.yml` workflow provides a manual escape hatch to sync GitHub repository secrets into ESC.

### 9.2 Environment Variables Reference

| Variable | Source | Used in | Purpose |
|---|---|---|---|
| `PULUMI_ACCESS_TOKEN` | ESC | build, preview, cleanup | Authenticate with Pulumi Cloud |
| `AWS_ACCESS_KEY_ID` | GitHub secret | preview (testing) | Initial AWS auth for testing environment |
| `AWS_SECRET_ACCESS_KEY` | GitHub secret | preview (testing) | Initial AWS auth for testing environment |
| `AWS_CI_ROLE_ARN` | ESC | preview, cleanup | IAM role to assume in testing account |
| `GITHUB_TOKEN` | GitHub Actions | build, preview, cleanup, metadata | GitHub API access |
| `PULUMI_BOT_TOKEN` | ESC | push checkout, tutorials, metadata PRs | Bot token for PR creation |
| `ALGOLIA_APP_ID` | GitHub var | build, preview | Algolia application ID |
| `ALGOLIA_APP_SEARCH_KEY` | GitHub var | build, preview | Algolia public search key |
| `ALGOLIA_APP_ADMIN_KEY` | ESC | production build | Algolia admin key (index writes) |
| `PULUMI_STACK_NAME` | GitHub var | build, cleanup | Pulumi stack to select (e.g., `pulumi/registry/production`) |
| `PULUMI_DOCS_STACK_NAME` | GitHub var | build | Pulumi docs stack reference |
| `DEPLOYMENT_ENVIRONMENT` | GitHub var | build, cleanup | e.g., `testing` or `production` |
| `NODE_OPTIONS` | Hardcoded | build, browser tests | `--max_old_space_size=8192` |
| `SLACK_ACCESS_TOKEN` | ESC | link check, cleanup | Slack Web API token for posting messages |
| `SLACK_WEBHOOK_URL` | ESC | notify jobs | Slack incoming webhook for failure alerts |
| `PULUMI_CONVERT_URL` | Pulumi stack output | `build.sh` | URL for Pulumi conversion service |
| `PULUMI_AI_WS_URL` | Pulumi stack output | `build.sh` | URL for Pulumi AI WebSocket |
| `PULUMI_ATLAS_URL` | Pulumi IaC config | `build.sh` | Copilot (Atlas) API URL |
| `ASSET_BUNDLE_ID` | `build.sh` (computed) | Hugo templates | Unique suffix for CSS/JS cache-busting |
| `CSS_BUNDLE` / `JS_BUNDLE` | `build.sh` (computed) | Hugo templates | Paths to versioned asset bundles |

### 9.3 Tool Version Matrix

| Tool | `mise.toml` | `pull-request.yml` (preview) | `push.yml` (production) | `testing-deploy.yml` |
|---|---|---|---|---|
| Node.js | 20 | 22.x | 22.x | 22.x |
| Go | 1.23 | 1.23.x | 1.23.x | 1.21.x |
| Hugo | 0.135 | 0.135.0 | 0.135.0 | 0.135.0 |
| golangci-lint | 2.1.6 | v2.1.6 (check-go.yml) | — | — |

Note: `mise.toml` specifies Node 20 for local development, while CI workflows use Node 22. The `lint-markdown` and `lint-scripts` jobs in `pull-request.yml` use Node 23.x. The `bucket-cleanup.yml` workflow uses Node 18.x.

---

## Scheduled & Maintenance Tasks

| Task | Schedule (UTC) | Workflow | Key Command |
|---|---|---|---|
| Community package metadata check | 5:30 AM + 5:30 PM daily | `generate-package-metadata.yml` | `python generate_package_list.py` → `resourcedocsgen pkgversion` / `metadata from-github` |
| Tutorial refresh from examples | 3:00 PM daily | `update-tutorials.yml` | `scripts/ci/mktutorial.sh` → PR on branch `tutorials/refresh` |
| Link check | 3:00 PM every Monday | `check-links.yml` | `make check_links` |
| Browser tests (scheduled) | 2:00 PM daily | `run-browser-tests.yml` | `make run-browser-tests` |
| Stale bucket cleanup | 3:00 PM daily | `bucket-cleanup.yml` | `make ci_bucket_cleanup` |

---

## Troubleshooting

### Hugo version mismatch

**Symptom**: Build fails with template errors or unexpected output.

**Fix**: Ensure you are using Hugo 0.135 extended. Check with `hugo version`. If using mise: `mise install` will install the correct version. Ensure you have the **extended** variant (required for SCSS processing).

### Node out-of-memory error

**Symptom**: Node process killed during build with `FATAL ERROR: Reached heap limit`.

**Fix**: Set the Node memory limit before running the build:

```bash
export NODE_OPTIONS="--max_old_space_size=8192"
make build
```

This is set automatically in CI. If local builds still fail, consider increasing the value or ensuring adequate system RAM (32 GB+ recommended for full builds).

### PR preview not posting

**Symptom**: PR has no preview comment; the `preview` job shows "Missing secret tokens, possibly due to a forked PR."

**Cause**: PRs from forks do not have access to repository secrets. `scripts/ci/pull-request.sh` detects the absence of `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`, or `PULUMI_ACCESS_TOKEN` and skips the S3 sync. The build still runs; only the deployment is skipped.

**Fix**: This is expected behavior for fork PRs. If you need a full preview, merge the fork into a branch on the upstream repo.

### `resourcedocsgen` schema fetch failure

**Symptom**: `resourcedocsgen docs registry` exits with a non-zero status; error mentions a specific provider.

**Cause**: The provider's schema URL (in its YAML file under `themes/default/data/registry/packages/`) may be unreachable, or the `version` field may reference a tag that does not exist in the provider's GitHub repo.

**Fix**: Check the YAML file for the failing provider. Verify that the version tag exists on the provider's GitHub repo and that the schema file path is correct.

### S3 bucket cleanup failures

**Symptom**: `bucket-cleanup.yml` fails or leaves stale buckets.

**Cause**: The lifecycle policy approach means buckets are not immediately deleted. The 2-step process (apply lifecycle → wait 48h → delete) is intentional. Buckets that have been in cleanup state for 7+ days will cause the cleanup script to exit non-zero.

**Fix**: Manually inspect and delete the affected bucket via AWS console or CLI. Remove the associated SSM parameter.

### Pulumi stack selection failure

**Symptom**: `scripts/ci/login.sh` fails with "stack not found" or similar.

**Cause**: `PULUMI_STACK_NAME` environment variable is not set or does not match an available stack.

**Fix**: Ensure `PULUMI_STACK_NAME` is set correctly in the GitHub Actions environment variables for the relevant environment (testing or production). The stack must exist in the Pulumi Cloud org. Run `pulumi -C infrastructure stack ls` locally (with correct credentials) to confirm available stacks.

### `push-registry.py` publisher not found

**Symptom**: `push-registry.py` raises `Exception: Missing publisher entry for "<publisher-name>"`.

**Cause**: A package YAML file references a `publisher` display name that is not listed in `tools/resourcedocsgen/pkg/publishers/publisher-names.json`.

**Fix**: Add the publisher to `publisher-names.json` with the correct canonical identifier, or update the package YAML to use an already-registered publisher name.

---

## Pulumi Cloud Service Integration

### 12.1 Two Parallel Systems

The Pulumi Registry is actually **two separate but tightly coupled systems** that must remain in sync:

| System | URL | Primary Consumer | Data Source |
|---|---|---|---|
| **Static Hugo site** | `pulumi.com/registry` | Humans (browser) | YAML files in `themes/default/data/registry/packages/` |
| **Pulumi Cloud Registry API** | `api.pulumi.com/api/registry` | Pulumi CLI (`pulumi up`, `pulumi package add`) | Pulumi Cloud database (populated via `pulumi package publish`) |

These two systems are **not the same thing and do not share a data store**. The static site is rebuilt from YAML files on every push to `master`; it does not query the Pulumi Cloud API at runtime. The Pulumi Cloud API is a live service that stores package metadata independently.

The bridge between them is `scripts/ci/push-registry.py`, which runs after every production build and publishes any new package versions to the Pulumi Cloud API.

```
YAML files in repo                            Pulumi Cloud Registry API
(source of truth for                          (source of truth for CLI
 the static Hugo site)                         package resolution)
        │                                               │
        │  scripts/ci/push-registry.py                 │
        │  (runs on every production push)              │
        └──────────────────────────────────────────────►│
                  pulumi package publish                │
```

**Consequence**: If `push-registry.py` fails silently on a particular package, the Hugo site will show the package correctly but the Pulumi CLI will not be able to resolve it. The two systems can drift.

**Consequence**: The static site does not support version browsing (no "select a version" dropdown) because Hugo generates a fixed set of pages from a fixed set of YAML files. Versioned snapshots (e.g., `aws-v6`) are implemented as entirely separate YAML files, separate Hugo pages, and separate API publication entries — not as a first-class versioned concept.

---

### 12.2 Pulumi Cloud Registry API

**Base URL**: `https://api.pulumi.com/api/registry`

**Authentication**: Bearer token via `PULUMI_ACCESS_TOKEN` for write operations. Read operations (listing packages, checking versions) are publicly accessible.

**Key endpoints**:

| Method | Path | Purpose |
|---|---|---|
| `GET` | `/packages` | List all packages (paginated via `continuationToken`) |
| `GET` | `/packages/{source}/{publisher}/{name}/versions/{version}` | Check if a specific version exists (returns 200 or 404) |
| *(via CLI)* | — | `pulumi package publish <schema.json>` writes a new version |

**Package record fields** stored in the API:

| Field | Type | Notes |
|---|---|---|
| `name` | string | Provider identifier (e.g., `aws`, `kubernetes`) |
| `publisher` | string | Normalized publisher ID (e.g., `pulumi`, `aws`) |
| `publisherDisplayName` | string | Human-readable publisher name |
| `source` | string | `pulumi` or `opentofu` |
| `version` | string | Semver without `v` prefix (e.g., `7.19.0`) |
| `title` | string | Display name |
| `description` | string | Package description |
| `category` | string | `cloud`, `infrastructure`, `utility`, `network` |
| `packageTypes` | array | `["native"]`, `["bridged"]`, or `["component"]` |
| `packageStatus` | string | `ga` or `public_preview` |
| `visibility` | string | `public` |
| `isFeatured` | bool | Whether package is featured in the registry |
| `repoUrl` | string | GitHub repository URL |
| `readmeURL` | string | URL to the package README / overview doc |
| `schemaURL` | string | URL to the Pulumi JSON schema |
| `logoUrl` | string | Package logo URL |
| `pluginDownloadURL` | string | Custom plugin download location (optional) |
| `createdAt` | string | ISO 8601 timestamp |
| `parameterization` | object | For dynamically-bridged providers: base provider and encoded parameters |

**What the Pulumi CLI uses this API for**:

- Resolving package dependencies when you run `pulumi up` with a package reference.
- `pulumi package add <name> <version>` — looks up the schema URL to generate local SDKs.
- Plugin download resolution — finds the correct `pluginDownloadURL` for `pulumi plugin install`.

**What the API does NOT provide for the static site**: The Hugo site does not read from this API at build time or at request time. All package data displayed on `pulumi.com/registry` comes entirely from the YAML files in `themes/default/data/registry/packages/`, and all documentation HTML is generated by `resourcedocsgen` reading schemas from GitHub.

---

### 12.3 How push-registry.py Bridges Them

`scripts/ci/push-registry.py` is the synchronization mechanism. On every push to `master`, after the Hugo site is built and deployed, this script:

1. Reads every YAML file from `themes/default/data/registry/packages/`.
2. For each package, queries `GET /api/registry/packages/{source}/{publisher}/{name}/versions/{version}` to check whether this exact version already exists in the API.
3. If it does **not** exist (404):
   - Downloads the provider schema JSON from the URL in the YAML file.
   - Corrects the `version` field in the schema if it is absent or inconsistent with the YAML.
   - Runs `pulumi package publish <schema.json> --readme <_index.md> --source <source> --publisher <publisher>` to register the version.
4. If it **does** exist (200): no-op.
5. Skips deprecated packages, `azure-native-v*` aliases, and `aws-v*` legacy versioned packages.

**Important**: `push-registry.py` is also run in **dry-run mode** on every PR (`--dry-run` flag) as the `test-live-publish` CI job. This validates that all YAML files are parseable, all publishers are known, and the `pulumi package publish` invocation would be valid — without actually touching the production API.

**Required credential**: `PULUMI_ACCESS_TOKEN` must be set. This is sourced from Pulumi ESC in CI.

---

### 12.4 Cross-Stack References Used at Build Time

During the CI build (`scripts/ci/build.sh`), the build script queries several other Pulumi stacks in the `pulumi` organization to retrieve runtime service URLs. These are injected as environment variables and baked into the Hugo-generated HTML at build time.

| Stack | Output key | Variable | Purpose |
|---|---|---|---|
| `pulumi/tf2pulumi-service/production` | *(Pulumi stack output)* | `PULUMI_CONVERT_URL` | URL of the Terraform-to-Pulumi conversion service, used for interactive code conversion on docs pages |
| `pulumi/pulumigpt-api/corp` | *(Pulumi stack output)* | `PULUMI_AI_WS_URL` | WebSocket URL for Pulumi AI integration embedded in the Registry UI |
| `pulumi/registry/production` (IaC config) | `atlasUrl` | `PULUMI_ATLAS_URL` | URL of the Copilot (Atlas/Neo) AI assistant |

Additionally, the `infrastructure/index.ts` IaC program reads cross-stack references at **deploy time** to grant S3 bucket access to data warehouse workloads:

| Stack | Output key | Purpose |
|---|---|---|
| `pulumi/dwh-workflows-orchestrate-airflow/production` | `airflowTaskRoleArn` | Grants Airflow task role access to the e2e test results S3 bucket |
| `pulumi/dwh-workflows-loader-prodbuckets/production` | `dwhBucketReaderRole` | Grants DWH bucket reader role access to the e2e test results S3 bucket |

The e2e test results bucket (`pulumi-api-docs-e2e-test-results-production`) stores Cypress test output, which is subsequently consumed by the Pulumi data warehouse for quality tracking.

---

### 12.5 Pulumi Cloud Roles Summary

Pulumi Cloud plays **four distinct roles** in this system:

| Role | Mechanism | Purpose |
|---|---|---|
| **Registry API** | `api.pulumi.com/api/registry` | Stores published package versions; queried by the Pulumi CLI for package and plugin resolution |
| **Secrets / ESC** | `github-secrets/pulumi-registry` environment | Provides CI secrets (tokens, keys) via OIDC; eliminates long-lived secrets in GitHub |
| **IaC State Backend** | `pulumi/registry/testing` and `pulumi/registry/production` stacks | Stores Terraform-like state for the AWS infrastructure (CloudFront, S3, policies) managed by `infrastructure/index.ts` |
| **Stack References** | `pulumi/tf2pulumi-service`, `pulumi/pulumigpt-api`, `pulumi/dwh-workflows-*` stacks | Exposes runtime service URLs and IAM role ARNs as stack outputs, consumed by the build and deploy pipeline |

