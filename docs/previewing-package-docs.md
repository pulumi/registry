# Previewing your package's docs

Your package's section of the Pulumi Registry is assembled from the following sources:

- **Your provider's schema** becomes the API docs — a page per resource and function, plus the nav tree beside them.
- **`docs/_index.md` in your provider repo** becomes the Overview page at `/registry/packages/<package>/`. See [The Overview page](./overview-page.md) for what belongs in it.
- **`docs/installation-configuration.md` in your provider repo**, if you have one, becomes the Install & config page. It's optional: most packages keep this material in the Overview instead.
- **Guides committed to this repository**, under `themes/default/content/registry/packages/<package>/how-to-guides/`, become the How-to guides section. You can add as many as you like, and this is where anything longer-form goes — the AWS [6.0](https://www.pulumi.com/registry/packages/aws/how-to-guides/6-0-migration/) and [7.0](https://www.pulumi.com/registry/packages/aws/how-to-guides/7-0-migration/) migration guides are two of these.
- **The package's metadata**, at `themes/default/data/registry/packages/<package>.yaml`, records the version to render, the URL to fetch the schema from, and the title, publisher, logo, and category the site displays.

The sidebar links a fixed set of paths, so where a page lands decides whether readers can find it. It shows Overview, then Install & config if `installation-configuration.md` exists, then How-to guides if the `how-to-guides/` directory exists — with a count of the guides in it — then API Docs. (Four more paths are hardcoded for packages that predate this arrangement: `version-guide`, `from-classic`, `from-v1-to-v2`, and `from-v2-to-v3`, all used by Azure Native.) A Markdown file anywhere else under the package still renders at its own URL, but nothing links to it: `gcp/service-account.md` is one of those, reachable only from links in the GCP Install & config page.

Guides are the one part of a package's docs contributed directly to `pulumi/registry`, and you don't have to work at Pulumi or own the provider to add one. Open a pull request adding your page under `themes/default/content/registry/packages/<package>/how-to-guides/` — many of the guides there were written by people outside Pulumi. Nothing in the publish pipeline touches that directory, since `resourcedocsgen` only ever rewrites `_index.md` and `installation-configuration.md`, so a merged guide survives your next release. Working from a fork changes one thing: CI won't build a preview for you, so preview the guide locally and ask a maintainer to comment `/preview` on the pull request if you want a URL to share.

None of this renders when you tag a release. Your version is picked up later — by a `repository_dispatch` from your release workflow, or one of the twice-daily polls for community packages — which opens a pull request here updating your package's metadata, and the site rebuilds when that merges. So by default, the first time you see your docs rendered is after they're live for everyone. The rest of this document is how to see them sooner. If you're changing the registry itself rather than publishing a package, see [Previewing registry changes](./previewing-registry-changes.md) instead.

## Prerequisites

You need a working checkout of this repository, which you set up once by following [Using this repository](../README.md#using-this-repository) in the README:

```bash
mise trust && mise install
make ensure
make build-assets
```

## Preview a release locally

A release may change more than the schema. For example, a major version typically comes with a rewritten Overview, updated provider configuration, and a new migration guide, and you want to read them as one package rather than one file at a time.

You can assemble all of it locally because Hugo serves this repository's root `content/` and `static/` directories alongside the theme's, and a file in the root `content/` wins over the committed copy of the same page. That directory is git-ignored, so you can put your whole release there — schema output and hand-written pages together — without touching a tracked file or risking a stray commit.

Set three variables first, so the commands in this section work for any provider:

```bash
PKG=cloudamqp                          # your package's name in the registry
PROVIDER_REPO=~/src/pulumi-cloudamqp   # wherever your provider repo is checked out
VERSION=v9.9.9                         # the version to render as; any valid semver
```

1. Build the docs generator. It compiles to `./bin/resourcedocsgen`:

    ```bash
    make bin/resourcedocsgen
    ```

1. Get your provider's schema, and point `SCHEMA` at it. If your provider bridges a Terraform provider — anything built from [`pulumi-tf-provider-boilerplate`](https://github.com/pulumi/pulumi-tf-provider-boilerplate) — the `tfgen` target writes it into the provider repo. (Newer provider repos alias the same target as `make schema`.)

    ```bash
    make -C "$PROVIDER_REPO" tfgen
    SCHEMA="$PROVIDER_REPO/provider/cmd/pulumi-resource-$PKG/schema.json"
    ```

    For any other provider, including a component provider, build the plugin binary — `make provider` in most repos — and ask it for its schema:

    ```bash
    pulumi package get-schema "$PROVIDER_REPO/bin/pulumi-resource-$PKG" > "/tmp/$PKG-schema.json"
    SCHEMA="/tmp/$PKG-schema.json"
    ```

1. From the root of this repository, generate the API docs from that schema:

    ```bash
    ./bin/resourcedocsgen docs \
        --schemaFile "$SCHEMA" \
        --version "$VERSION" \
        --docsOutDir "./content/registry/packages/$PKG/api-docs" \
        --packageTreeJSONOutDir ./static/registry/packages/navs
    ```

    `--docsOutDir` takes the generated Markdown — one page per resource and function, plus an index — and the site expects it under `content/registry/packages/<package>/api-docs`.

    `--packageTreeJSONOutDir` takes a single `<package>.json` describing the package's resource tree. The nav on the left of every API docs page fetches it in the browser from `/registry/packages/navs/<package>.json`, so it has to land in `static/registry/packages/navs`.

    `--version` is required and must be valid semver. It's recorded as the package's version and doesn't affect which schema you render — `--schemaFile` already decided that — so use the version you plan to release.

1. Copy in the pages you've written, into the paths the sidebar looks for:

    ```bash
    mkdir -p "content/registry/packages/$PKG"
    cp "$PROVIDER_REPO/docs/_index.md" "content/registry/packages/$PKG/"
    cp "$PROVIDER_REPO/docs/installation-configuration.md" "content/registry/packages/$PKG/"   # if you have one
    ```

    A migration guide or any other long-form page goes under `how-to-guides/`. That directory needs its own `_index.md`: without one, your guide still renders at its URL, but the section page doesn't exist and nothing appears in the sidebar. Create it if the package has no guides yet:

    ```bash
    mkdir -p "content/registry/packages/$PKG/how-to-guides"
    cat > "content/registry/packages/$PKG/how-to-guides/_index.md" <<EOF
    ---
    title: $PKG How-to Guides
    meta_desc: Guides for the Pulumi $PKG package.
    layout: package
    ---
    EOF
    cp ~/my-migration-guide.md "content/registry/packages/$PKG/how-to-guides/"
    ```

    Guides need front matter with `title`, `meta_desc`, and `layout: package`; copy the shape from [an existing one](https://github.com/pulumi/registry/blob/master/themes/default/content/registry/packages/aws/how-to-guides/7-0-migration.md). These copies are local stand-ins. When the release is real, the Overview and Install & config pages are fetched from your release tag automatically, and guides are contributed to `pulumi/registry` by pull request.

1. Start the site locally:

    ```bash
    make serve
    ```

Your package is at `http://localhost:1313/registry/packages/$PKG/`, with the sidebar linking the Overview, Install & config, How-to guides, and API docs pages you just assembled. Re-run the generator after each schema change, and re-copy after editing a page; the running Hugo server picks up both.

To remove the generated package preview, delete `content/registry/packages/$PKG/`.

## Render the page the way the registry will build it

The previous section assembles a preview out of local files, but it bypasses the package metadata: the version, title, publisher, logo, category, and the install snippets that depend on them all come from `themes/default/data/registry/packages/<package>.yaml`, which none of those commands read.

To exercise the path that CI runs, set `schema_file_url` in that `<package>.yaml` to a schema on your branch and let `make` do the whole job. Do this when you want to confirm the metadata-driven parts of the page, and when you want [a preview you can share](#share-a-preview), since the same edit is what makes CI build one. It needs a package that's already in the registry; if yours isn't, stay with the previous section and see [Adding a new package](./adding-a-new-package.md).

1. Put the schema somewhere public. Most providers commit their `schema.json`, so pushing your branch is enough: the raw URL is `https://raw.githubusercontent.com/<org>/<repo>/<branch>/provider/cmd/pulumi-resource-<name>/schema.json`. Providers whose schema is too large to commit, such as Azure Native, need it uploaded to S3 or another public host.

1. In `themes/default/data/registry/packages/<package>.yaml`, set `schema_file_url` to that URL, and set `version` to a semver version that has never been published to the Pulumi Registry service, such as `v9.9.9`.

    Both edits are required, because `resourcedocsgen` prefers the published schema: it asks `api.pulumi.com` for the package at `version` first, and falls back to `schema_file_url` only when that lookup 404s. Leave `version` at a published value and you'll get the published schema and none of your changes, with nothing in the output to say so.

1. Generate the docs and start the site:

    ```bash
    PKG=cloudamqp   # your package's name in the registry
    make SKIP_VERSIONED_DOCS=1 "api-docs/$PKG"
    make serve
    ```

    `SKIP_VERSIONED_DOCS=1` skips the older-major-version snapshots, which need a Pulumi-internal tool that most contributors can't install. Leave it set unless you're specifically working on versioned docs.

This workflow edits a tracked file, unlike the previous one. Revert it before committing anything else, unless you're opening the throwaway pull request described next:

```bash
git checkout -- "themes/default/data/registry/packages/$PKG.yaml"
```

### Picking up a new commit on your branch

`resourcedocsgen` won't regenerate output it believes is current. It records what it generated from in a `.generated` file next to the pages — the contents of the package YAML, plus the generator's own build — and skips the package when neither has changed, logging `Skipping (output is fresh)`. Push a new commit to your branch and the schema behind `schema_file_url` changes while the YAML doesn't, so it skips and you keep reading the old docs.

`make -B` doesn't fix this. The target's sentinel is an intermediate file that Make deletes after each run, so the recipe already re-runs every time; the skip happens inside the generator. Delete the `.generated` file instead:

```bash
rm "content/registry/packages/$PKG/api-docs/.generated"
make SKIP_VERSIONED_DOCS=1 "api-docs/$PKG"
```

## Share a preview

To send someone a rendered preview of unreleased docs, open a pull request against this repository containing only the `schema_file_url` and `version` edit from the previous section. CI builds the full site from your branch's schema and posts a pinned comment holding the preview URL and a list of the changed pages.

Close that pull request when you're done rather than merging it — it exists to produce a URL, not to change the registry. Your real version lands through the [normal publish pipeline](./architecture.md#registry-package-updates) when you tag a release.

A pull request that adds a how-to guide is the opposite case: that one you do merge, and its preview shows the guide in place.

Neither kind of pull request gets a preview from a fork, because the preview job only runs for branches pushed to this repository. Ask a Pulumi maintainer to build you one by commenting `/preview` on the pull request.

## Troubleshooting

**Your Overview page is missing, though the API docs render.** Nothing generates the Overview page from your schema. Copy `docs/_index.md` into `content/registry/packages/<package>/`, as described in [Preview a release locally](#preview-a-release-locally).

**A page you added doesn't appear in the sidebar.** The sidebar only links a fixed set of paths. Long-form pages belong in `how-to-guides/`, and that directory needs an `_index.md` of its own before the sidebar shows the section at all. The paths the sidebar links are listed at the top of this document.

**The API docs render, but the nav on the left is empty.** That nav is fetched in the browser from `/registry/packages/navs/<package>.json`, and it fails quietly when the file is missing. Confirm `static/registry/packages/navs/<package>.json` exists — that's what `--packageTreeJSONOutDir` writes.

**`make api-docs/<package>` fails with `registry-mirror-discover ... Repository not found`.** The build tried to generate versioned docs, which uses a tool that lives in a Pulumi-internal repository. Re-run with `SKIP_VERSIONED_DOCS=1`.

**The generator logs `Skipping (output is fresh)` and your changes don't appear.** Its output cache thinks the pages are current. See [Picking up a new commit on your branch](#picking-up-a-new-commit-on-your-branch).

**Hugo isn't on port 1313.** `make serve` doesn't pass `--port`, so Hugo binds a random free port when 1313 is already taken. Read the port off the server's own startup output rather than assuming it.

## Learn more

- [The Overview page](./overview-page.md) — what belongs in the `docs/_index.md` you author in your provider repo.
- [Adding a new package](./adding-a-new-package.md) — getting a package into the registry in the first place.
- [Architecture](./architecture.md#registry-package-updates) — how a tagged release becomes a published page, and what rebuilds the site.
- [Previewing registry changes](./previewing-registry-changes.md) — the other direction: changing this repository and checking it against published packages.
