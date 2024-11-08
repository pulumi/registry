# Registry

[Pulumi Registry](https://pulumi.com/registry) is the global index of everything you can do with Pulumi. The home of pulumi.com/registry.

## Authoring a Pulumi Package

We’re always eager to expand that index with new [Pulumi Packages](https://www.pulumi.com/docs/guides/pulumi-packages/). Whether you want to author a new native provider, bridge a provider from the Terraform ecosystem, or create a cloud component with best practices and sensible defaults built in, we’d like to work with you to list it on Pulumi Registry.
To get started, use our [guide on authoring a Pulumi Package](https://www.pulumi.com/docs/guides/pulumi-packages/how-to-author/) of your own.

### Publishing a Community Package on the Registry

#### Community Steps

To publish a community-maintained package on the Pulumi Registry as a community member:

1. Ensure your provider repo has the following files:
    * `docs/_index.md`, which should contain a summary of the provider's purpose (required) along with a code sample (preferred). This file will render as the index page for your provider ([example](https://www.pulumi.com/registry/packages/aiven/)).
    * `docs/installation-configuration.md`, which should contain links to SDK packages in each language along with instructions for configuring the provider (e.g., required credentials and/or environment variables). This file will render as the Installation & Configuration page for your provider ([example](https://www.pulumi.com/registry/packages/aiven/installation-configuration/)).
1. Create a release of your provider in GitHub.
1. Add your package to the [community packages list](./community-packages/package-list.json) via pull request to this repository.

For assistance, please reach out on the [Pulumi community slack](https://slack.pulumi.com/) or get in touch with us via this [contact form](https://pulumi.com/contact/?form=registry).

#### Pulumi Steps

Once the community member has submitted the PR to add the provider to the registry, a Pulumi staff member should perform the following steps:

1. Review the PR. Ensure that the PR has been rebased if necessary before merging. If ok, merge.
1. Once the PR is merged, a [scheduled task](https://github.com/pulumi/registry/actions/workflows/generate-package-metadata.yml) will pick up the changes and create a PR to add the package metadata to the registry. A correct metadata PR ([example](https://github.com/pulumi/registry/pull/1606/files)) will include the following files, at a minimum:
   * `data/registry/${PROVIDER}.yaml` which includes structured metadata about the provider. This file is always included with every PR that gets generated for a new release.
   * `themes/default/content/registry/packages/${PROVIDER}/installation-configuration.md`, as described above. This file *must* be included in the first release, but will only be included in subsequent PRs if the content has changed.
   * `themes/default/content/registry/packages/exoscale/_index.md`, as described above. This file *must* be included in the first release, but will only be included in subsequent PRs if the content has changed.

   Optionally, the PR may include additional content files like How-To Guides if they are present in the upstream repo.

1. Merge the PR if it looks ok.
1. In pulumi/docs, a [scheduled task](https://github.com/pulumi/docs/actions/workflows/update-theme.yml) runs hourly and will pick up any changes in this repo, generate files from the provider schema and `data/registry/${PROVIDER}.yaml`, and publish to pulumi.com.

  This scheduled task currently lacks adequate monitoring, and **should be watched to ensure that it runs correctly to completion**. (If it fails, it will block all updates to pulumi.com, including marketing and manually maintained docs pages.)

## About this repository

This repository is a [Hugo module](https://gohugo.io/hugo-modules/) that doubles as a development server to make it easier to work on the pages that make up Pulumi Registry. It contains all of the Hugo `content` and `layouts` files, JavaScript, CSS, and web components. comprising what you see at https://pulumi.com/registry

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

This repository does not contain the content of the API docs packages. We generate these pages at deployemnt time. In order to render the API docs for a package locally you will need to generate the API Docs pages for it. The API docs for packages can be generated on-demand using the [`resourcedocsgen` tool](tools/resourcedocsgen/README.md).

```
cd tools/resourcedocsgen
go build -o "${GOPATH}/bin/resourcedocsgen" .
```

As an example, you can generate the API docs for a specific package by running the resourcedocsgen tool and passing it the name of the package as follows:

```
resourcedocsgen docs registry <package_name> --baseDocsOutDir "./themes/default/content/registry/packages"
```

Run `resourcedocsgen --help` for help regarding its use or [see the `resourcedocsgen` README](tools/resourcedocsgen/README.md).

## Submitting, merging and releasing

Before submitting a pull request, run the linter locally:

```bash
make lint
```

When you're ready to submit a pull request, make sure you've removed anything that doesn't seem to belong (`go.mod`/`go.sum` changes, etc.) and submit the PR in the usual way.


> [!NOTE]
> It currently requires a machine with a minimum of 32 GB of memory (64 GB preferred) to build the registry in its entirety including _all_ pacakges.

Once your PR is approved and merged into the default branch of this repository, it will be deployed to the registry site (https://pulumi.com/registry).
