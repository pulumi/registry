# Previewing Registry Changes

Registry pages are generated from provider schemas and from this repository's templates, so a change to either isn't visible until it's published. You can render the affected pages locally first, or have a preview site built from a pull request.

Pick the section that matches what you changed:

- [Local preview from a schema file](#local-preview-from-a-schema-file) — a provider schema you have on disk that isn't released yet.
- [Local preview of a published package](#local-preview-of-a-published-package) — a package already in the registry: a docs template, a Hugo layout, the theme, or a package's YAML metadata.
- [Local preview from a branch schema](#local-preview-from-a-branch-schema) — a provider schema you can put at a public URL, such as a branch on GitHub.
- [Pull request preview](#pull-request-preview) — anything you're ready to push to `pulumi/registry`.

## Prerequisites

The local sections below assume you've set the repository up once:

```bash
mise trust && mise install
make ensure
make build-assets
```

See [the README](../README.md#using-this-repository) for details.

## Local preview from a schema file

Use this when you're changing a provider and want to see how its API docs will render before you cut a release.

1. Build the docs generator:

    ```bash
    make bin/resourcedocsgen
    ```

1. Generate your provider's schema in the provider repo. For a bridged provider that's usually `make schema`, which writes `provider/cmd/pulumi-resource-<name>/schema.json`.

1. From the root of this repository, run `resourcedocsgen docs` against that file, writing into the two locations Hugo serves from:

    ```bash
    ./bin/resourcedocsgen docs \
        --schemaFile ../pulumi-aws/provider/cmd/pulumi-resource-aws/schema.json \
        --version v9.9.9 \
        --docsOutDir ./content/registry/packages/aws/api-docs \
        --packageTreeJSONOutDir ./static/registry/packages/navs
    ```

    `--version` is required and must be valid semver. Use a dummy version higher than anything published so it's obvious in the rendered page that you're looking at a local build.

    Both output directories are git-ignored, so nothing you generate here can end up in a commit.

1. The package's landing pages are committed under `themes/default/content/registry/packages/<package>/` and are used as-is. Only `_index.md` is required; `installation-configuration.md` is an optional split for packages whose install and config content outgrows the overview — see [The Overview page](./overview-page.md) and [Publishing packages](https://www.pulumi.com/docs/iac/guides/building-extending/packages/publishing-packages/#overview-installation--configuration). If you're previewing a package that isn't in the registry yet, create that directory and copy `_index.md` into it from your provider repo's `docs/` folder. `resourcedocsgen docs` doesn't write these pages — `resourcedocsgen metadata from-urls` fetches them from the provider repo, and the publish workflow is what runs it.

1. Serve the site:

    ```bash
    make serve
    ```

    Your pages are at `http://localhost:1313/registry/packages/<package>/api-docs/`.

Re-run step 3 after each schema change; the running Hugo server picks up the new files. If you're also changing CSS or JavaScript under `themes/default/theme`, use `make serve-all` instead so assets rebuild too.

## Local preview of a published package

Use this when the schema is already published and you're changing something on this side — a docs template, a Hugo layout, the theme, or a package's YAML metadata.

```bash
make SKIP_VERSIONED_DOCS=1 api-docs/aws
make serve
```

`make api-docs/<package>` reads `themes/default/data/registry/packages/<package>.yaml`, fetches the corresponding schema, and writes to `content/registry/packages/<package>/` at the repository root, which is git-ignored. `make serve` then serves the whole site as usual, so the package lands at `http://localhost:1313/registry/packages/<package>/` — see [Troubleshooting](#troubleshooting) if that port doesn't answer.

`SKIP_VERSIONED_DOCS=1` skips generating the older-major-version snapshots, which requires a Pulumi-internal tool that most contributors can't install. Leave it set unless you are specifically working on versioned docs.

### Forcing a rebuild

`resourcedocsgen` caches on the package YAML plus its own build identity, recorded in a `.generated` file next to the output. A schema that changed behind an unchanged `schema_file_url` therefore looks fresh to it, and it logs `Skipping (output is fresh)` instead of regenerating.

`make -B` does not help here: the Make target's sentinel is an intermediate file that Make deletes after each run, so the recipe already re-runs every time and the skip happens inside the generator. Delete the sentinel and re-run:

```bash
rm content/registry/packages/aws/api-docs/.generated
make SKIP_VERSIONED_DOCS=1 api-docs/aws
```

## Local preview from a branch schema

Use this when you want the full registry pipeline — metadata, nav tree, published schema file, and all — but against a schema that only exists on a branch.

1. Publish the schema to a public URL. For most providers the schema is committed, so pushing your branch is enough; the raw URL looks like `https://raw.githubusercontent.com/<org>/<repo>/<branch>/provider/cmd/pulumi-resource-<name>/schema.json`. For providers whose schema is too large to commit (Azure Native, for example), upload it to S3 or any other public host.

1. Edit `themes/default/data/registry/packages/<package>.yaml`:

    - Set `schema_file_url` to the URL from step 1.
    - Set `version` to a semver version that has **not** been published to the Pulumi Registry service — a bumped dummy version such as `v9.9.9`.

    Both edits are required. `resourcedocsgen` asks `api.pulumi.com` for the package at `version` first and only falls back to `schema_file_url` when that lookup 404s. If you leave `version` at a published value you'll silently get the published schema and none of your changes.

1. Generate and serve:

    ```bash
    make SKIP_VERSIONED_DOCS=1 api-docs/<package>
    make serve
    ```

Revert the YAML edit before committing.

## Pull request preview

A pull request whose branch lives in `pulumi/registry` gets a full site build published to a per-commit S3 bucket. CI maintains a single pinned comment on the PR containing:

- the preview URL for the current commit, and
- a **Changed pages** list linking directly to the pages your PR affects.

The comment is updated in place on each build rather than added per commit, and the preview buckets are deleted when the PR closes.

**Pull requests from forks don't get this automatically.** The preview job only runs for branches pushed to this repository, so if you're contributing from a fork, use the local sections above. A Pulumi maintainer can build you one on demand by commenting `/preview` on the pull request.

If your change is in a provider repo rather than here, you can still get a preview by opening a PR against this repository with the [branch schema](#local-preview-from-a-branch-schema) edits applied — the same YAML change works in CI. Don't merge that PR; it exists to produce the preview.

## Troubleshooting

**`registry-mirror-discover ... Repository not found`** — `make api-docs/<package>` tried to build the versioned-docs tool, which lives in a Pulumi-internal repository. Re-run with `SKIP_VERSIONED_DOCS=1`.

**`Skipping (output is fresh)` and your changes don't appear** — see [Forcing a rebuild](#forcing-a-rebuild).

**Pages 404 in the local server** — check that the nav tree JSON was written to `static/registry/packages/navs/<package>.json` and that `themes/default/content/registry/packages/<package>/_index.md` exists. The API docs pages hang off that landing page.

**Hugo isn't on port 1313** — `make serve` doesn't pass `--port`, so Hugo binds a random free port when 1313 is already taken. Read the port off the server's own startup output rather than assuming it.
