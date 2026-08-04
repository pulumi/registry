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
| Check the generated dark logo variants | `make lint-dark-logos` |
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
| `config/` | Hugo configuration (includes custom output formats for CLI docs) |
| `public/` | Build output (git-ignored) |

### Provider YAML Files

Every package in the registry is described by a YAML file at `themes/default/data/registry/packages/<name>.yaml`. These files drive:

- API doc generation (`resourcedocsgen`)
- Registry publication (`scripts/ci/push-registry.py`)
- The nightly community package update workflow

### Go Tools

| Tool | Location | Purpose |
|---|---|---|
| `resourcedocsgen` | `tools/resourcedocsgen/` | Generates provider API reference docs and LLM-friendly JSON bundles (`llm-docs.json`) from Pulumi provider schemas |
| `mktutorial` | `tools/mktutorial/` | Generates how-to guides from `pulumi/examples` |

Both are compiled to `bin/` by the Makefile. `mktutorial` is CI-only; it does not run in local builds.

### Content Generation: Local vs. CI

`make api-docs/<pkg>` writes output to **`content/registry/packages/`** (repo root).

The CI build script (`scripts/ci/build.sh`) writes to **`themes/default/content/registry/packages/`**.

This difference is intentional — local API doc generation stays out of the Hugo theme tree. Within `themes/default/content/registry/packages/<pkg>/`, the `api-docs/` subdirectory is regenerated on every build and is git-ignored, so never commit it. The `_index.md` and `installation-configuration.md` landing pages are different: they are committed and maintained by the `generate-package-metadata.yml` publish workflow, and are bundled when onboarding a package so it renders before the next nightly run. Do not hand-edit any of these files; regenerate them with `resourcedocsgen`.

Additionally, `resourcedocsgen` writes **LLM docs** to **`llm-docs-out/registry/packages/`** (repo root, git-ignored). These are terminal-friendly markdown bundles (`llm-docs.json`) uploaded to S3 separately from the Hugo site. The LLM docs format is specified in `docs/llm-markdown-spec.md`.

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

## Dark Mode

The registry has a light/dark/system theme toggle, pinned to the bottom of the left sidebar. Dark is **light-first**: light is the baseline and dark is a pure override. The whole system lives in `themes/default/theme/src/scss/docs/_docs-theme.scss` (read its header comment first) and is driven by semantic `--docs-*` tokens defined on `body.section-registry` and re-pointed under `html[data-theme="dark"]`.

The preference is stored under the **same `localStorage` key as `pulumi.com/docs`** (`pulumi-docs-theme`), so it is one setting across pulumi.com. Picking dark on `/docs` flips the registry too.

| File | Role |
|---|---|
| `theme/src/scss/docs/_docs-theme.scss` | Token layer + every dark override + the toggle's own styles |
| `theme/src/scss/docs/_code-light.scss` | Min Light code theme, gated on `html:not([data-theme="dark"])` |
| `theme/src/ts/docs-theme.ts` | Toggle behavior, `system` tracking |
| `layouts/partials/head.html` | Pre-paint inline script (writes `data-theme` on `<html>`; prevents FOUC) |
| `layouts/partials/docs/theme-toggle.html` | The three-button control |
| `scripts/generate-dark-logos.py` | Generates `-on-dark.svg` variants of local package logos |
| `scripts/classify-external-logos.py` | Decides which external `logo_url` marks need a light chip |

**You must check both modes whenever you add or restyle a visible element** — new partials, cards, callouts, buttons, icons, or any markup that introduces its own colors, backgrounds, borders, or images. Pure content changes (prose, YAML metadata, code samples) don't need a dark-mode pass.

When something needs dark-mode work, prefer the existing levers over hand-written one-off colors:

- **Use the semantic tokens.** Paint with `var(--docs-fg)`, `--docs-fg-muted`, `--docs-bg`, `--docs-bg-alt`, `--docs-surface`, `--docs-border`, `--docs-card`, `--docs-card-border`, `--docs-link`, `--docs-ring` rather than raw `--color-*` scales or hex literals — they flip automatically. Two caveats when you're *retrofitting* a token onto an existing light rule rather than writing a new one:
  - *The fallback form doesn't protect light mode.* The tokens are declared on `body.section-registry`, so anywhere inside the registry they always resolve — `var(--docs-TOKEN, ORIGINAL)` never falls back, and the element silently picks up the token's light value instead of the hex it had. `--docs-fg-muted` in light is gray-700 (`#6a6675`), not whatever grey the rule used before. The fallback only earns its keep on markup that also renders *outside* the section (the 404 page, standalone components). Compare the two values before you swap.
  - *A layered file often can't opt in at all.* Almost everything under `theme/src/scss/` is inside `@layer components`, and `_docs-theme.scss` is unlayered, so its blanket dark rules — `p { color: var(--docs-fg) }`, the per-surface overrides — win regardless of specificity. Pointing a layered rule at a token therefore changes light and does nothing in dark. Declare the dark value in `_docs-theme.scss` and leave the layered file's light value alone.
- **Use Tailwind `dark:` variants.** The `dark:` variant is wired to `data-theme` (`@custom-variant dark` in `theme/src/scss/main.scss`), so `dark:bg-gray-900`, `dark:text-white`, etc. work directly in templates. `bg-docs-bg` / `text-docs-fg` / `border-docs-border` utilities are generated too.
- **Lean on the automatic flips.** Brand violet (`--color-violet-primary` and every `text-violet-primary` / `bg-violet-primary` / `border-violet-primary` utility) and the default border color are already remapped in the dark block, so markup authored with those gets dark mode for free. Surfaces styled via `@apply` (the `.btn-*` variants, for example) carry no literal utility class and need their own override in `_docs-theme.scss`.
- **Package logos** are handled in two tiers by `layouts/partials/registry/package/icon.html`, and neither tier is hand-maintained:
  - *Local marks.* The partial looks for `themes/default/assets/fingerprinted/logos/pkg/<name>-on-dark.svg` and, when it finds one, emits both images with the `docs-logo-light` / `docs-logo-dark` classes the theme swaps. Run `python3 scripts/generate-dark-logos.py` after adding or replacing a local mark; it lightens only the paints that are too dark for the dark page, skips full-color marks entirely, and skips any package whose YAML declares a `logo_url` (`icon.html` short-circuits on that before it ever looks for a local variant). `make lint-dark-logos` runs `--check`, which reports staleness without writing; it's a PR-CI job, so a stale variant fails the build.
  - *External `logo_url` marks.* These can't be recolored, so the ones that would disappear get a near-white chip instead. Which ones those are is measured by `python3 scripts/classify-external-logos.py`, which downloads each logo, samples it, and writes `themes/default/data/registry/external_logo_treatment.yaml`. Re-run it when you add a package with a `logo_url` or when a vendor changes their logo (`--check` in CI-style mode). It needs the network — and macOS `sips` for non-PNG rasters — so it is deliberately *not* in CI; `--check` exits 2 rather than reporting staleness if any logo couldn't be measured. Don't apply the chip by hand or across the board — it hides light marks and boxes in full-color ones.
- **Masked icons** in `docs/_icons.scss` tint with `background: var(--docs-fg-muted)`; `background: url()` colored marks do not and need a per-asset variant.

Note that `_docs-theme.scss` is imported **unlayered**, after everything else, so its rules beat anything in `@layer components` or `@layer utilities` regardless of specificity. That is deliberate (much of the chrome is painted with Tailwind utilities), but it means a dark rule there can't be walked back by a more specific layered rule.

The one deliberate exception is the dark default border color, which sits in `@layer base` inside that file. Unlayered it would outrank `@layer utilities` too and flatten every colored `border-{color}-{step}` a template authored down to a gray hairline; in `@layer base` it still replaces Tailwind's gray-200 default while any layered rule with its own border color keeps it. If you write another blanket `*` rule, put it in `@layer base` for the same reason.

---

## Conventions

- **Package manager**: Yarn only. Do not use npm or pnpm.
- **Go modules**: `tools/resourcedocsgen/` and `tools/mktutorial/` are separate Go modules. Run `go test ./...` and `golangci-lint run` from within those directories (or use the Makefile targets).
- **Generated content**: Files under `themes/default/content/registry/packages/` are generated — regenerate with `resourcedocsgen`, never hand-edit. The `api-docs/` subdirectories are git-ignored build output (never committed); the `_index.md` and `installation-configuration.md` pages are committed metadata maintained by the publish workflow.
- **Branch naming**: Use `<GitHub Username>/<descriptive-name>` for branches in this repository.
