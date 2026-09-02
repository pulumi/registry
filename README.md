# Registry

[Pulumi Registry](https://pulumi.com/registry) is the public index of Pulumi extensions and integrations.

## Adding a Package

Adding a community package is one pull request that adds a single entry to [`community-packages/package-list.json`](https://github.com/pulumi/registry/blob/master/community-packages/package-list.json). Your provider repo supplies a `docs/_index.md` and a `v`-prefixed release; the docs and metadata are generated and published for you after merge, so you never commit generated files here. Automated checks post a fact-sheet on the PR, and a Pulumi maintainer reviews it. You do not need to file an issue first.

**[Adding a new package](./docs/adding-a-new-package.md) has the complete instructions** — what the page must contain, the User-Agent your provider should set, how to re-run the checks, and the checklist a maintainer works through before merging. Read it before opening a PR.

One exception: a **dynamically bridged** Terraform provider, consumed with `pulumi package add terraform-provider <name>` and having no provider repo or committed schema, cannot be added by pull request. Open a ["New Package"](https://github.com/pulumi/registry/issues/new?template=new-package.yml) issue instead.

For assistance, please reach out on the [Pulumi community Slack](https://slack.pulumi.com/) or get in touch with us via this [contact form](https://pulumi.com/contact/?form=registry).

## About this repository

This repository is a [Hugo module](https://gohugo.io/hugo-modules/) that doubles as a local development server, making it easier to work on the pages of the Pulumi Registry. It contains everything behind what you see at <https://pulumi.com/registry>: the Hugo `content` and `layouts` files, along with the JavaScript, CSS, and web components.

We build the JavaScript and CSS bundles that power the Pulumi Registry here, under the `themes/default/theme` directory. If you are making styling changes alongside content changes, use `make serve-all` to enable hot reloading of both the pages and CSS/JS assets.

## Using this repository

### Prerequisites

We use [Mise](https://mise.jdx.dev/getting-started.html#installing-mise-cli) to manage the development environment. Install Mise then run:

```sh
mise trust && mise install
```

This should install all the dependencies you need to develop. Check out `mise.toml` to see what dependencies are used.

### Installing dependencies

The prerequisites listed above need to be installed on your machine in order to serve the site.

1. Run `make ensure` to check for the appropriate tools and versions, and install any dependencies. The script will let you know if you're missing anything important.
  
 ```
 make ensure
 ```

1. Once that succeeds, run `make build-assets` to build the assets the site depends on. This needs to be done before the first time you serve this repo so the assets exist on your local machine.

 ```
 make build-assets
 ```

### Running Hugo locally

Once you've run the above successfully, you're ready to run the development server:

```
make serve
```

Optionally, use `make serve-all` to enable hot reloading of both the pages and CSS/JS assets.

#### Generating API docs for packages

This repository does not contain the content of the API docs packages. We generate these pages at deployment time. In order to render the API docs for a package locally you will need to generate the API Docs pages for it. The API docs for packages can be generated on-demand using the [`resourcedocsgen` tool](tools/resourcedocsgen/README.md).

To build the API docs for a single package, run:

``` bash
make .make/content/registry/packages/<package_name>/api-docs
```

Run `make bin/resourcedocsgen && ./bin/resourcedocsgen --help` for help regarding its use or [see the `resourcedocsgen` README](tools/resourcedocsgen/README.md).

## Submitting, merging and releasing

Before submitting a pull request, run the linter locally:

```bash
make lint
```

When you're ready to submit a pull request, make sure you've removed anything that doesn't seem to belong (`go.mod`/`go.sum` changes, etc.) and submit the PR in the usual way.

> [!NOTE]
> It currently requires a machine with a minimum of 32 GB of memory (64 GB preferred) to build the registry in its entirety including *all* packages.

Once your PR is approved and merged into the default branch of this repository, it will be deployed to the registry site (<https://pulumi.com/registry>).
