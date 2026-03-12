---
user-invocable: false
description: Shared review criteria for registry content and code
---

# Registry Review Criteria

Shared criteria used by `/registry-review`, `/pr-review`, and other review skills.

---

## Go Code (`tools/resourcedocsgen/`, `tools/mktutorial/`)

- **Idiomatic Go**: Follow Go conventions (naming, error handling, package structure)
- **Error handling**: All errors must be checked and handled; no silently ignored errors
- **Test coverage**: New functionality should include tests; modifications should not break existing tests
- **Lint compliance**: Must pass `golangci-lint run` (run via `make lint-go`)
- **Module boundaries**: `tools/resourcedocsgen/` and `tools/mktutorial/` are separate Go modules — changes should respect module boundaries

## CI Workflows (`.github/workflows/`)

- **YAML validity**: Well-formed YAML with correct indentation
- **Action version pinning**: Use SHA-pinned actions (e.g., `actions/checkout@v4` with full SHA) for security
- **Secret handling**: No hardcoded secrets; use GitHub secrets or Pulumi ESC references
- **OIDC config**: Verify `id-token: write` permission when using OIDC authentication
- **Trigger conditions**: Verify `on:` triggers, branch filters, and path filters are correct
- **Concurrency**: Check for appropriate `concurrency` groups to prevent duplicate runs

## Build System (`Makefile`, `scripts/`)

- **Backward compatibility**: Changes to Makefile targets should not break existing CI workflows
- **`make lint` passability**: All changes must pass `make lint`
- **Script safety**: Shell scripts should use `set -euo pipefail`, quote variables, handle edge cases
- **Dependency management**: Changes to `package.json` should use Yarn (not npm or pnpm)

## Provider YAML (`themes/default/data/registry/packages/*.yaml`)

### Required Fields

Every package YAML file must include:

| Field | Type | Description |
|---|---|---|
| `name` | string | Package short name (lowercase, alphanumeric + hyphens) |
| `title` | string | Display name |
| `description` | string | Brief package description |
| `publisher` | string | Publisher name (e.g., "Pulumi", "pulumiverse") |
| `repo_url` | string | GitHub repository URL |
| `version` | string | Semver version (e.g., `v1.2.3`) |
| `schema_file_url` | string | URL to the provider's schema JSON |
| `category` | string | One of the valid categories (see below) |
| `component` | boolean | Whether this is a component provider |
| `native` | boolean | Whether this is a native provider |
| `package_status` | string | `ga` or `public_preview` |

### Valid Categories

`Cloud`, `Database`, `Infrastructure`, `Monitoring`, `Network`, `Utility`, `Version Control System`

### Field Validation

- `version` must match semver pattern: `v\d+\.\d+\.\d+` (with optional pre-release suffix)
- `repo_url` must be a valid GitHub URL: `https://github.com/{owner}/{repo}`
- `schema_file_url` must be a valid URL pointing to a JSON file
- `component` and `native` must be boolean (`true` or `false`)
- `package_status` must be `ga` or `public_preview`

## Infrastructure (`infrastructure/`)

- **Pulumi IaC review**: Verify resource configurations, stack references, and outputs
- **CloudFront/S3 impact**: Flag changes that could affect CDN distribution or origin bucket configuration
- **`origin-bucket-metadata.json`**: Changes to this file directly affect production deployment — flag for careful review

## BUILD-AND-DEPLOY.md Currency

If any of the following files change, flag whether `BUILD-AND-DEPLOY.md` needs updating:

- `Makefile`
- `.github/workflows/*.yml`
- `scripts/ci/*.sh`
- `infrastructure/`
- `config/`

## Community Package List (`community-packages/package-list.json`)

- **JSON validity**: Must be well-formed JSON
- **Required fields**: Each entry must have `repoSlug` (string) and `schemaFile` (string)
- **Repo slug format**: Must be `{owner}/{repo}` format (valid GitHub repository slug)
- **No duplicates**: No duplicate `repoSlug` entries

## Hugo Templates and Theme (`themes/default/layouts/`, `themes/default/theme/`)

- **Template syntax**: Valid Hugo template syntax with proper `{{ }}` delimiters
- **Asset compilation**: Theme JS/CSS changes should build cleanly with `make build-assets`
- **Static files**: Changes to `themes/default/static/` should not conflict with generated content
- **Image hosting**: Images referenced in templates (banner ads, promotional graphics, etc.) must use absolute URLs pointing to `https://www.pulumi.com/` (e.g., `https://www.pulumi.com/images/docs/ad/example.png`). The registry site is served from a separate S3 bucket/CloudFront distribution than the main docs site — relative paths like `/images/docs/ad/example.png` will 404 because those images live in `pulumi/docs`, not `pulumi/registry`. **Never add image files to the registry repo's static directory to work around this** — always reference preexisting images on `pulumi.com` with absolute URLs

## General

- **Spelling and grammar**: Check non-generated content for typos and grammatical errors
- **Link validation**: Verify URLs are well-formed; flag broken or placeholder links
- **Newline endings**: Files should end with a trailing newline
- **No generated content commits**: Never commit files from `themes/default/content/registry/packages/` (these are generated)
- **No secrets**: Flag any file that appears to contain API keys, tokens, or credentials
