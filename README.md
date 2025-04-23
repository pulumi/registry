# Registry

[Pulumi Registry](https://pulumi.com/registry) is the public index of Pulumi extensions and integrations.

## Adding a Package

If you want to see a new package to the Pulumi Registry, open an issue requesting the addition using the ["New Package"](https://github.com/pulumi/registry/issues/new?template=new-package.yml) issue template. You can request either a [Pulumi Package](https://www.pulumi.com/docs/guides/pulumi-packages/) (Native, Bridged or Component) that you have written or any Terraform/OpenTofu provider present in the [opentofu package registry](https://search.opentofu.org).

For assistance, please reach out on the [Pulumi community slack](https://slack.pulumi.com/) or get in touch with us via this [contact form](https://pulumi.com/contact/?form=registry).

### Adding a Pulumi Package that you have authored

We’re always eager to expand that index with new [Pulumi Packages](https://www.pulumi.com/docs/guides/pulumi-packages/). To get started, use our [guide on authoring a Pulumi Package](https://www.pulumi.com/docs/guides/pulumi-packages/how-to-author/) of your own.

To publish a community-maintained package on the Pulumi Registry as a community member:

1. Ensure your provider repo has the following files:
    * `docs/_index.md`, which should contain a summary of the provider's purpose (required) along with a code sample in each language (required). This file will render as the index page for your provider ([example](https://www.pulumi.com/registry/packages/aiven/)).
    * `docs/installation-configuration.md`, which should contain links to published SDKs in (Go, Typescript, C#, Python; required) along with instructions for configuring the provider (e.g., required credentials and/or environment variables). This file will render as the Installation & Configuration page for your provider ([example](https://www.pulumi.com/registry/packages/aiven/installation-configuration/)).
1. Create a release of your provider in GitHub with a "v" + [Semver 2.0](https://semver.org) compliant tag (`vX.Y.Z`).

## Internal Process

Pulumi maintainers should follow the [adding a new package guide](./docs/adding-a-new-package.md).

## About this repository

This repository is a [Hugo module](https://gohugo.io/hugo-modules/) that doubles as a development server to make it easier to work on the pages that make up Pulumi Registry. It contains all of the Hugo `content` and `layouts` files, JavaScript, CSS, and web components. comprising what you see at <https://pulumi.com/registry>

We build the JavaScript and CSS bundles that power the Pulumi Registry here, under the `themes/default/theme` directory. If you are making styling changes along-side content changes, use `make serve-all` to enable hot reloading of both the pages and CSS/JS assets.

## Using this repository

### Prerequisites

We build the Pulumi website statically with Hugo, manage our Node.js dependencies with Yarn, and write most of our documentation in Markdown. Below is a list of the tools you'll need to run the website locally:

* [Go](https://golang.org/) (>= 1.23)
* [Hugo](https://gohugo.io) (>= 0.135.0)
* [Node.js](https://nodejs.org/en/) (>= 18)
* [Yarn](https://classic.yarnpkg.com/en/) (1.x)

### Installing dependencies

The prerequisites listed above need to be installed on your machine in order to serve the site.

1. Run `make ensure` to check for the appropriate tools and versions, and install any dependencies. The script will let you know if you're missing anything important.
  
 ```
 make ensure
 ```

1. Once that succeeds, run `make build_assets` to build the assets the site depends on. This needs to be done before the first time you serve this repo so the assets exist on your local machine.

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
> It currently requires a machine with a minimum of 32 GB of memory (64 GB preferred) to build the registry in its entirety including *all* pacakges.

Once your PR is approved and merged into the default branch of this repository, it will be deployed to the registry site (<https://pulumi.com/registry>).
