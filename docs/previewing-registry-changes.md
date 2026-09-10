# Previewing registry changes

Every package page under `pulumi.com/registry` is assembled at build time from these sources:

- **Provider schemas**, downloaded during the build, which become the API docs — a page per resource and function, plus the nav tree beside them.
- **`docs/_index.md` from each provider's repo**, fetched at publish time and committed here under `themes/default/content/registry/packages/<package>/`, which becomes the package's Overview page. An `installation-configuration.md` beside it, where a package has one, becomes the Install & config page.
- **Guides written directly in this repository**, under that same directory's `how-to-guides/`, which become the How-to guides section — the AWS migration guides, the Kubernetes FAQ, and so on.
- **Per-package metadata**, at `themes/default/data/registry/packages/<package>.yaml`, which records each package's version, schema URL, title, publisher, logo, and category.
- **This repository's layouts, templates, and theme**, which render every one of those into the pages readers see.

A build of this repository produces the whole registry, not one package: `make api-docs` regenerates every package's API docs from its current schema, and Hugo renders all of them into a single site. CI caches per-package output, so packages whose metadata hasn't changed are restored from the previous build rather than regenerated — but the site it publishes always contains every package. That's why a change to a shared layout or theme file lands on all 300-odd packages at once.

Your change reaches the live site when it merges to `master`: the push workflow rebuilds the site and swaps the CloudFront origin to the new build — see [Publish](./architecture.md#publish).

Until then, `make serve` on its own won't show you much. The Overview pages and guides are committed, but every API docs page is generated during the build and isn't in your checkout, so the packages you most want to look at have nothing under them. This document covers how to generate one locally and how to get a shareable preview from a pull request.

If you're a package author who wants to see how your next release will render, see [Previewing your package's docs](./previewing-package-docs.md) instead.

## Prerequisites

You need a working checkout, which you set up once by following [Using this repository](../README.md#using-this-repository) in the README:

```bash
mise trust && mise install
make ensure
make build-assets
```

## Render a package locally

Pick a package that exercises what you changed, generate its docs, and serve the site:

```bash
make SKIP_VERSIONED_DOCS=1 api-docs/aws
make serve
```

The `api-docs/<package>` target reads `themes/default/data/registry/packages/<package>.yaml`, downloads the schema that file points at, and writes the pages to `content/registry/packages/<package>/api-docs` and the nav tree to `static/registry/packages/navs/<package>.json`. Hugo serves this repository's root `content/` and `static/` directories alongside the theme's, which is why generated pages appear on the local site; both are git-ignored, so nothing you generate can end up in a commit. Your package is then at `http://localhost:1313/registry/packages/<package>/`.

`SKIP_VERSIONED_DOCS=1` skips generating the older-major-version snapshots, which needs a Pulumi-internal tool that most contributors can't install. Leave it set unless you're specifically working on versioned docs.

Generation takes a while for large packages, so generate one small package plus one large one rather than the whole registry — `make build` rebuilds everything, but it needs 32 GB of RAM and a lot of patience. `random` and `aws` are a reasonable pair.

If you're changing CSS or JavaScript under `themes/default/theme`, use `make serve-all` in place of `make serve` so assets rebuild as you edit.

### Forcing a rebuild

`resourcedocsgen` won't regenerate output it believes is current. It records what it generated from in a `.generated` file next to the pages — the contents of the package YAML, plus the generator's own build — and skips the package when neither has changed, logging `Skipping (output is fresh)`. Editing the generator invalidates that, but editing a schema behind an unchanged URL doesn't.

`make -B` doesn't fix this. The target's sentinel is an intermediate file that Make deletes after each run, so the recipe already re-runs every time; the skip happens inside the generator. Delete the `.generated` file instead:

```bash
rm content/registry/packages/aws/api-docs/.generated
make SKIP_VERSIONED_DOCS=1 api-docs/aws
```

Changes to layouts, templates, and theme assets don't need any of this. Hugo re-renders those from the pages you've already generated.

## Preview a pull request

A pull request whose branch lives in `pulumi/registry` gets a full site build — every package, generated the same way `master` generates them — published to a per-commit S3 bucket. CI keeps a single pinned comment on the PR holding the preview URL for the current commit and a **Changed pages** list linking straight to the pages your PR affects. That comment is updated in place on each build rather than added per commit, and the bucket is deleted when the PR closes.

This is the only practical way to check a change against the whole registry rather than the handful of packages you generated locally, so it's worth pushing a draft PR early for anything touching shared layouts or the theme.

Pull requests from forks don't get a preview automatically, because the preview job only runs for branches pushed to this repository. If you're contributing from a fork, preview locally, or ask a Pulumi maintainer to build you one by commenting `/preview` on the pull request.

## Troubleshooting

**`make api-docs/<package>` fails with `registry-mirror-discover ... Repository not found`.** The build tried to generate versioned docs, which uses a tool that lives in a Pulumi-internal repository. Re-run with `SKIP_VERSIONED_DOCS=1`.

**The generator logs `Skipping (output is fresh)` and your changes don't appear.** Its output cache thinks the pages are current. See [Forcing a rebuild](#forcing-a-rebuild).

**The API docs render, but the nav on the left is empty.** That nav is fetched in the browser from `/registry/packages/navs/<package>.json`, and it fails quietly when the file is missing. Confirm `static/registry/packages/navs/<package>.json` exists.

**A package page is empty apart from its Overview.** Its API docs haven't been generated in this checkout. Run `make SKIP_VERSIONED_DOCS=1 api-docs/<package>`.

**Hugo isn't on port 1313.** `make serve` doesn't pass `--port`, so Hugo binds a random free port when 1313 is already taken. Read the port off the server's own startup output rather than assuming it.

## Learn more

- [Architecture](./architecture.md) — where each source of registry content comes from, and how a build reaches production.
- [Using this repository](../README.md#using-this-repository) — one-time setup and the everyday build commands.
- [`resourcedocsgen`](../tools/resourcedocsgen/README.md) — every flag the generator accepts.
- [Previewing your package's docs](./previewing-package-docs.md) — the other direction: publishing a package and checking it before release.
