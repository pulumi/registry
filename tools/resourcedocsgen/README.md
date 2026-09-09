# Resource Docs generator driver

This tool calls the docs generator tool in `pulumi/pulumi` which uses the Pulumi schema for a package to generate API (resource) docs.

## Installation

You can install the `resourcedocsgen` tool just like any other Go-based CLI tool:

```
go install github.com/pulumi/registry/tools/resourcedocsgen@master
```

To build and install from source:

```
cd tools/resourcedocsgen
go build -o "${GOPATH}/bin/resourcedocsgen" .
```

## Usage

Then you can run any of the available commands using `resourcedocsgen <command> <flags>`. Run `resourcedocsgen --help` to see the available commands.

As of this writing, the tool supports three main purposes:

* Generate the registry metadata
* Generate API docs and the package nav tree
* Generate Overview page snippets a package author pastes into their own `docs/_index.md`

### Generating package metadata

Package metadata is used by the [Pulumi Registry](https://github.com/pulumi/registry) to generate the listing shown at https://pulumi.com/registry.
The metadata file contains information sourced from the package's own Pulumi schema.

### Generating API docs and the package nav tree

* The `docs` command. This command expects the package repo to be cloned alongside the `docs` repo. To help with invoking this command
with the right values, the `gen_resource_docs.sh` script can be used. Run `. ./scripts/gen_resource_docs.sh` from the root of this repo
with the arguments that the script accepts.
* The `registry` command. This command is a child of the above `docs` command. It uses the information in the registry repo at the commit
hash that the `docs` Hugo module depends on. Specifically, it uses the metadata files from the `data` folder of the registry repo's default
Hugo module. The metadata files serve as a snapshot of what packages exist in the registry and therefore, the packages for which API docs
need to be built. If you are running this command from anywhere but the `tools/resourcedocsgen` folder, you should override the values
for `baseDocsOutDir` and `basePackageTreeJSONOutDir`. For example, you might want to run this command to generate API docs and nav tree files
for a certain package (or all packages) within the registry repo for development purposes.

For both of the above commands, the default location for generating the API docs is `content/registry/packages/<package name>/api-docs`
and for the nav tree it is `static/registry/packages/navs/<package name>.json`.

### Generating Overview page sections

**These commands are experimental.** They are new, no CI depends on them, and their output shape and flags may change as we learn what package authors actually need. Nothing in the build calls them.

Two sections of a package's Overview page — `## Installation` and the configuration parameters reference under `## Configuration`, both specified in [`docs/overview-page.md`](../../docs/overview-page.md) — are derivable from the package's schema. `gen-install` and `gen-config` derive them and print markdown to stdout.

These commands are **advisory**. The Overview page is authored in the provider's own repository and fetched from its release tag by `metadata from-github`, so these generate a snippet to paste and edit, not a page. Point them at a local schema, or at the `schema_file_url` from a package's YAML in `themes/default/data/registry/packages/`:

```bash
resourcedocsgen gen-install --schemaFile provider/cmd/pulumi-resource-example/schema.json --version v1.2.3
resourcedocsgen gen-config  --schemaFile provider/cmd/pulumi-resource-example/schema.json --style table
```

A schema cannot say which SDKs a package actually publishes: the `language` blob is written by the code generator, not by the release pipeline. `pulumi-vault` declares no `java` block yet publishes `com.pulumi/vault`, and the roughly hundred bridged providers that publish no SDKs at all still declare four or five language blocks. So `gen-install` guesses from the schema and takes `--languages` as the correction, including `--languages none` for a package whose installation section is a single `pulumi package add` command.

Every one of the seven chooser languages gets a tab either way. A language with a published SDK gets its package-manager command; one without gets `pulumi package add <name>`, which generates an SDK locally, taking its language from the `runtime` in the reader's `Pulumi.yaml`. That is why the command reads identically for C#, Java and YAML on a package that publishes only TypeScript, Python and Go SDKs.

HCL is the exception and never uses `pulumi package add`: its tab emits a `required_providers` block plus `pulumi install`. A native Pulumi provider is sourced as `pulumi/<name>` and pinned to an exact semver version; a package parameterized over a Terraform provider reuses that provider's upstream source and version instead. Because HCL always differs, the chooser is emitted even for a package with no SDKs at all.

A parameterized provider is the one case the schema answers on its own — it is consumed as a local package, so the other tabs default to `pulumi package add terraform-provider <namespace>/<name>`.

`gen-config` reads the schema's provider `config` block and emits a bullet list by default, matching what most existing Overview pages use; `--style table` emits a GFM table instead. Two things the standard requires cannot come from a schema — environment-variable fallbacks (usually read by the vendor SDK a layer beneath it) and mutually exclusive options — so its output ends with a reminder to add them by hand. Requiredness and secrecy come straight from the schema and are only as accurate as the schema is: `pulumi-vault` marks neither its required parameters nor its `token` as secret.

### Updating the API docs templates

This tool depends on the `pulumi/pulumi` repo, namely the `pkg/codegen/docs` generator.
The docs generator uses Go-based [templates](https://github.com/pulumi/pulumi/tree/master/pkg/codegen/docs/templates) to render the markdown files in-memory which this tool then writes to the filesystem.

To make changes to the templates, make sure you have `pulumi/pulumi` cloned locally and override the `github.com/pulumi/pulumi/pkg/v3` dependency to point to
your local repo for testing out the changes to the templates. Once done, you must submit a PR to the `pulumi/pulumi` repo.

Once your `pulumi/pulumi` PR is merged, you should update the pseudo-version that this tool uses by running:

```
go get -u github.com/pulumi/pulumi/pkg/v3@<commit hash>
go get -u github.com/pulumi/pulumi/sdk/v3@<commit hash>
```

To update to latest pulumi/pulumi use

```
go get -u github.com/pulumi/pulumi/pkg/v3
go get -u github.com/pulumi/pulumi/sdk/v3
go mod tidy
```
