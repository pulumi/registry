# AGENTS.md — Canonical Instructions for This Repository

This file defines guidance for AI agents like Claude Code and GitHub Copilot when working in this repository.

---

## Repository Context

This is **pulumi/registry** — the Hugo-based static site for [Pulumi Registry](https://pulumi.com/registry), the public index of Pulumi packages. The deployed site lives at `pulumi.com/registry`.

For comprehensive details on the build system, CI/CD workflows, and deployment infrastructure, see `BUILD-AND-DEPLOY.md`. **IMPORTANT**: `BUILD-AND-DEPLOY.md` is large; to conserve tokens, read only as needed and limit your focus to the sections relevant to your task.

---

## Build / Test / Lint Workflow

Install tools via [mise](https://mise.jdx.dev/) before anything else:

```bash
mise trust && mise install
```

Common commands (all defined in the root `Makefile`):

| Task | Command |
|---|---|
| Install dependencies | `make ensure` |
| Build theme CSS/JS | `make build-assets` |
| Local dev server | `make serve` |
| Dev server + asset hot reload | `make serve-all` |
| Generate API docs for one provider | `make api-docs/<package>` |
| Full local build (all providers) | `make build` |
| Run Go unit tests | `make test` |
| Run all linters | `make lint` |
| Run Go linter only | `make lint-go` |
| Run Markdown linter only | `make lint-markdown` |
| Run provider API docs tests | `make test_provider_api_docs` |
| Run browser tests | `make run-browser-tests` |

`make lint` must pass before merging. `make build` requires 32 GB+ RAM for a complete build; for single-provider work, use `make api-docs/<package>` instead.

---

## Architecture

### Hugo Structure

Hugo reads from `themes/default/`:

| Path | Contents |
|---|---|
| `themes/default/content/` | Markdown pages (generated at build time — do not edit by hand) |
| `themes/default/data/registry/packages/` | One YAML file per provider — the source of truth for the registry |
| `themes/default/layouts/` | Hugo templates |
| `themes/default/static/` | Static assets, generated nav JSON, generated schema JSON |
| `themes/default/theme/` | JS/CSS source (compiled by `yarn --cwd ./themes/default/theme run build`) |
| `config/` | Hugo configuration |
| `public/` | Build output (git-ignored) |

### Provider YAML Files

Every package in the registry is described by a YAML file at `themes/default/data/registry/packages/<name>.yaml`. These files drive:

- API doc generation (`resourcedocsgen`)
- Registry publication (`scripts/ci/push-registry.py`)
- The nightly community package update workflow

### Go Tools

| Tool | Location | Purpose |
|---|---|---|
| `resourcedocsgen` | `tools/resourcedocsgen/` | Generates provider API reference docs from Pulumi provider schemas |
| `mktutorial` | `tools/mktutorial/` | Generates how-to guides from `pulumi/examples` |

Both are compiled to `bin/` by the Makefile. `mktutorial` is CI-only; it does not run in local builds.

### Content Generation: Local vs. CI

`make api-docs/<pkg>` writes output to **`content/registry/packages/`** (repo root).

The CI build script (`scripts/ci/build.sh`) writes to **`themes/default/content/registry/packages/`**.

This difference is intentional — local API doc generation stays out of the Hugo theme tree. Never commit generated content from `themes/default/content/registry/packages/`.

### Deployment Model

Each CI build syncs to a uniquely named S3 bucket. A Pulumi IaC program in `infrastructure/` reads `origin-bucket-metadata.json` to update the CloudFront origin. This means:

- Every PR commit gets its own preview URL (in the `testing` AWS account)
- Production deploys atomically swap the CloudFront origin (in the `production` AWS account)
- Preview buckets are deleted when a PR closes

---

## Adding or Updating a Provider

To add or update a community provider package:

1. Edit (or create) its YAML file in `themes/default/data/registry/packages/`.
2. The nightly `generate-package-metadata.yml` workflow handles version bumps automatically for community packages tracked in `community-packages/package-list.json`.
3. First-party Pulumi provider repos trigger `publish-provider-update.yml` via `repository_dispatch`.

The `push-registry.py` script publishes packages to the live Pulumi registry service on every push to `master`.

---

## Conventions

- **Package manager**: Yarn only. Do not use npm or pnpm.
- **Go modules**: `tools/resourcedocsgen/` and `tools/mktutorial/` are separate Go modules. Run `go test ./...` and `golangci-lint run` from within those directories (or use the Makefile targets).
- **Generated content**: `themes/default/content/registry/packages/` is generated — do not hand-edit files there.
- **Branch naming**: Use `<GitHub Username>/<descriptive-name>` for branches in this repository.
