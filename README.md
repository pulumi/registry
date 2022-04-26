# Registry

[Pulumi Registry](https://pulumi.com/registry) is the global index of everything you can do with Pulumi.

## Authoring a Pulumi Package
We’re always eager to expand that index with new [Pulumi Packages](https://www.pulumi.com/docs/guides/pulumi-packages/). Whether you want to author a new native provider, bridge a provider from the Terraform ecosystem, or create a cloud component with best practices and sensible defaults built in, we’d like to work with you to list it on Pulumi Registry.

To get started, use our [guide on authoring a Pulumi Package](https://www.pulumi.com/docs/guides/pulumi-packages/how-to-author/) of your own.
When ready, add your package to the [community packages list](./community-packages/package-list.json) via pull request to this repository.
For assistance, please reach out on the [Pulumi community slack](https://slack.pulumi.com/) or get in touch with us via this [contact form](https://pulumi.com/contact/?form=registry). 

## About this repository

This repository is a [Hugo module](https://gohugo.io/hugo-modules/) that doubles as a development server to make it easier to work on the pages that make up Pulumi Registry. It contains most of the Hugo `content` and `layouts` files comprising what you see at https://pulumi.com/registry -- but not everything. A few things you won't find in this repository include:

* Package-level how-to guides. These files are still built and checked into the [pulumi/docs](https://github.com/pulumi/docs) repository. (We're [working on bringing them into this repository](https://github.com/pulumi/registry/issues/237), though.)

* JavaScript, CSS, and web components. We build the JavaScript and CSS bundles that power the Pulumi website (and therefore the Registry) in the [pulumi/theme](https://github.com/pulumi/theme) repository.

* Layouts and content for pulumi.com marketing pages, CLI docs, the blog, etc., all of which are managed in the [pulumi/pulumi-hugo](https://github.com/pulumi/pulumi-hugo) repository.

    <img src="https://user-images.githubusercontent.com/274700/139131567-b8e3c43d-6407-4638-ae4e-4ad3f3794d89.png" width="60%">

You can, however, develop locally with this repository using content from these other Hugo-module repositories, either by loading their content remotely or pointing your Hugo development server to local clones of them. More on this below.

## Using this repository

### Prerequisites

We build the Pulumi website statically with Hugo, manage our Node.js dependencies with Yarn, and write most of our documentation in Markdown. Below is a list of the tools you'll need to run the website locally:

* [Go](https://golang.org/) (>= 1.15)
* [Hugo](https://gohugo.io) (>= 0.92)
* [Node.js](https://nodejs.org/en/) (>= 1.14)
* [Yarn](https://classic.yarnpkg.com/en/) (1.x)

### Installing dependencies

First, run `make ensure` to check for the appropriate tools, versions, and install any dependencies. The script will let you know if you're missing anything important.

```
make ensure
```

### Running Hugo locally

Once you've run `make ensure` successfully, you're ready to run the development server:

```
make serve
```

When you do this, Hugo will load the latest versions of:

* The [pulumi/pulumi-hugo](https://github.com/pulumi/pulumi-hugo) module, which contains our marketing pages, some docs content, the blog, etc.
* The [pulumi/theme](https://github.com/pulumi/theme) module, which contains our CSS and JavaScript bundles (web components, styles, etc.).

... and then start a development server at http://localhost:1313. Any changes you make to the content, layouts, or other [Hugo component folders](https://gohugo.io/getting-started/directory-structure/) should be reloaded automatically.

### Developing alongside another Hugo module

If you want to develop another module alongside this one -- e.g., add a new web component to use in the Registry, or to make changes to Registry-specific CSS -- you can point your development server to a local clone of [pulumi/theme](https://github.com/pulumi/theme). To do that, first clone the repository, then add a `replace` line to the `go.mod` file at the root of _this_ repository to override the existing reference to `pulumi/theme` temporarily. For instance:

```
module github.com/pulumi/pulumi-hugo

go 1.16

require (
	github.com/pulumi/theme v0.0.0-20211015193555-271ef1f67093 // indirect
)

// Add this line to tell Hugo to use your local clone of pulumi/theme.
replace github.com/pulumi/theme => ../theme
```

**Tip:** If you run `make serve` from the root of pulumi/theme (in another terminal tab) while also running `make serve` in this one, the changes you make to the CSS and JavaScript source files in pulumi/theme will be recompiled and reloaded in the browser automatically.

Be sure to remove the `replace` line before you commit.

### Linking to pages that don't exist in this repository

If you want to link to a page that exists on https://pulumi.com but not in this repository, you can still use a [Hugo `relref`](https://gohugo.io/content-management/shortcodes/#ref-and-relref) in the usual way -- you'll just need to make sure the path you're linking to does exist (or will exist by the time you merge your change). For example, to add a link pointing to the [Pulumi CLI documentation](https://www.pulumi.com/docs/reference/cli/) (which does not exist in this repository), you'd use:

```
{{< relref /docs/reference/cli >}}
```

... and just be aware while the link won't work for you locally, it will work once your change is merged and picked up by our website automation (see [Merging and releasing](#merging-and-releasing)] below for details). This works because we've suppressed Hugo's built-in `relref` validation locally to keep the module-development workflow as lightweight as possible.

### Temporarily pulling in content from pulumi/docs

If the change you're working on requires content from pulumi/docs -- e.g., the aforementioned how-to guides -- you may
want to be able to see some of that content as you develop. To do that, just copy the files you need from the
[`content` folder of pulumi/docs](https://github.com/pulumi/docs/tree/master/content) into the `content` folder of this
repository, remembering to remove those files before you commit. For example:

```
# Copy the AWS how-to guides from a local/sibling clone of pulumi/docs.
cp ../docs/content/registry/packages/aws/how-to-guides ./themes/default/content/registry/packages/aws/
```

#### Generating API docs for packages

The API docs for packages can be generated on-demand using the [`registrygen` tool](https://github.com/pulumi/registrygen).

```
go install github.com/pulumi/registrygen@master
```

Run `registrygen --help` for help regarding its use or [see the `registrygen` README](https://github.com/pulumi/registrygen#readme).

## Submitting, merging and releasing

When you're ready to submit a pull request, make sure you've removed anything that doesn't seem to belong (`go.mod`/`go.sum` changes, content from pulumi/docs, etc.) and submit the PR in the usual way.

If you're doing work in another repository that's associated with the changes in your PR, you can "pin" your PR branch to another module repository branch by pointing Hugo to that branch. To do that, use `hugo mod get` and pass a reference to the target branch:

```
hugo mod get github.com/pulumi/theme@my-special-branch
```

This will modify `go.mod` and `go.sum` accordingly and result in a PR preview that incorporates your changes from the other branch. Just be sure to remove these changes before you're ready to merge.

Once your PR is approved and merged into the default branch of this repository, an automated process that runs on [pulumi/docs](https://github.com/pulumi/docs) will detect your changes and produce a PR and integration build containing content from all other Hugo modules. Once that PR build completes and is approved and merged into pulumi/docs, your changes will be deployed to https://pulumi.com.
