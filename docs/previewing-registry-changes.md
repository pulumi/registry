# Previewing Registry Changes

This guide covers how to see Pulumi Registry pages — including provider API docs — before they are published, whether the change lives in a provider's schema, in this repository's templates, or in a pull request.

Pick a route based on what you changed:

| What you changed | Route |
|---|---|
| A provider schema you have locally, and it isn't released yet | [Local preview from a schema file](#local-preview-from-a-schema-file) |
| A package that is already published in the registry (templates, layouts, CSS, metadata) | [Local preview of a published package](#local-preview-of-a-published-package) |
| A provider schema you can publish to a public URL (e.g. a branch on GitHub) | [Local preview from a branch schema](#local-preview-from-a-branch-schema) |
| Anything you're ready to push to `pulumi/registry` | [Pull request preview](#pull-request-preview) |

## Prerequisites

All local routes assume you've set the repository up once:

```bash
mise trust && mise install
make ensure
make build-assets
```

See [the README](../README.md#using-this-repository) for details.

Generated docs land in git-ignored directories (`content/` and `static/registry/` at the repository root), so nothing from these routes should ever end up in a commit.

## Local preview from a schema file

Use this when you're changing a provider and want to see how its API docs will render before you cut a release. It reads a `schema.json` straight off your disk, so nothing needs to be pushed or published.

1. Build the docs generator:

    ```bash
    make bin/resourcedocsgen
    ```

1. Generate your provider's schema in the provider repo. For a bridged provider that's usually `make generate_schema`, which writes `provider/cmd/pulumi-resource-<name>/schema.json`.

1. From the root of this repository, run `resourcedocsgen docs` against that file, writing into the two locations Hugo serves from:

    ```bash
    ./bin/resourcedocsgen docs \
        --schemaFile ../pulumi-aws/provider/cmd/pulumi-resource-aws/schema.json \
        --version v9.9.9 \
        --docsOutDir ./content/registry/packages/aws/api-docs \
        --packageTreeJSONOutDir ./static/registry/packages/navs
    ```

    `--version` is required and must be valid semver. Use a dummy version higher than anything published so it's obvious in the rendered page that you're looking at a local build.

1. The package's landing pages (`_index.md` and `installation-configuration.md`) are committed under `themes/default/content/registry/packages/<package>/` and are used as-is. If you're previewing a package that isn't in the registry yet, create that directory and add the two files by hand, copying them from your provider repo's `docs/` folder.

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

`make api-docs/<package>` reads `themes/default/data/registry/packages/<package>.yaml`, fetches the corresponding schema, and writes to `content/registry/packages/<package>/` at the repository root.

`SKIP_VERSIONED_DOCS=1` skips generating the older-major-version snapshots, which requires a Pulumi-internal tool that most contributors can't install. Leave it set unless you are specifically working on versioned docs.

To force a rebuild after a change the generator can't see (it caches on the package YAML and its own build), delete the sentinel file:

```bash
rm content/registry/packages/aws/api-docs/.generated
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

Every pull request against `pulumi/registry` gets a full site build published to a per-commit S3 bucket. CI maintains a single pinned comment on the PR containing:

- the preview URL for the current commit, and
- a **Changed pages** list linking directly to the pages your PR affects.

The comment is updated in place on each build rather than added per commit, and the preview buckets are deleted when the PR closes.

If your change is in a provider repo rather than here, you can still get a preview by opening a PR against this repository with the [branch schema](#local-preview-from-a-branch-schema) edits applied — the same YAML change works in CI. Don't merge that PR; it exists to produce the preview.

For a community package pull request, a Pulumi maintainer can comment `/preview` to build a live preview of the package's pages.

## Troubleshooting

**`registry-mirror-discover ... Repository not found`** — `make api-docs/<package>` tried to build the versioned-docs tool, which lives in a Pulumi-internal repository. Re-run with `SKIP_VERSIONED_DOCS=1`.

**`Skipping (output is fresh)` and your changes don't appear** — the generator caches on the package YAML plus its own build identity, so a change to a remote schema alone won't invalidate it. Delete `content/registry/packages/<package>/api-docs/.generated` and re-run.

**Pages 404 in the local server** — check that the nav tree JSON was written to `static/registry/packages/navs/<package>.json` and that `themes/default/content/registry/packages/<package>/_index.md` exists. The API docs pages hang off that landing page.

**Hugo isn't on port 1313** — `make serve` doesn't pass `--port`, so Hugo binds a random free port when 1313 is already taken. Read the port off the server's own startup output rather than assuming it.
