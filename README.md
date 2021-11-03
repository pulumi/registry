# registry

[Pulumi Registry](https://pulumi.com/registry) is the global index of everything you can do with Pulumi.

We’re always eager to expand that index with new [Pulumi Packages](https://www.pulumi.com/docs/guides/pulumi-packages/). Whether you want to author a new native provider, bridge a provider from the Terraform ecosystem, or create a cloud component with best practices and sensible defaults built in, we’d like to work with you to list it on Pulumi Registry.

[Contact us to start publishing to Pulumi Registry](https://pulumi.com/contact/?form=registry). If you’d like, you can also get a head start by [authoring a Pulumi Package](https://pulumi.com/docs/guides/pulumi-packages/).

## About this repository

This repository is consumed by https://github.com/pulumi/docs to produce the website you see at https://pulumi.com. If you're interested in contributing some documentation for a Pulumi package, you've come to the right place! 🙌

This repository is a [Hugo module](https://gohugo.io/hugo-modules/) that doubles as a development server to make it easier to work on the pages that make up Pulumi Registry. It contains most of the Hugo `content` and `layouts` files comprising what you see at https://pulumi.com/registry. **It does not contain every page of the Pulumi website**, because many of those pages (e.g., those comprising our CLI and SDK docs) are generated from source code, so they aren't meant to be edited by humans directly.

Because of this, many of the links you follow when browsing around on the development server (to paths underneath `/docs/reference` for example) will fail to resolve because their content files are checked into a different repository. When we build the Pulumi website, we merge this module along with any others into a single build artifact, but when you're working within an individual module like this one, you may find you're unable to reach certain pages or verify the links you may want to make to them.

See below to learn more about contributing, submitting pull requests, merging, and releasing.

### What's in this repo

* Hand-authored, Registry-specific content and documentation, including package overview pages and install guides.
* Registry-specific Hugo module components (layouts, partials, shortcodes, data).

### What's not in this repo

* Package-level API navigation, documentation, and how-to guides. These files are still built and checked into the [pulumi/docs](https://github.com/pulumi/docs) repository. (We're [working on bringing them into this repository](https://github.com/pulumi/registry/issues/237), though.)
* Generated documentation for the Pulumi CLI and SDK. You'll find this at https://github.com/pulumi/docs.
* CSS and JavaScript source code. You'll find these at https://github.com/pulumi/theme.
* Most content and layouts for pulumi.com. You'll find these at https://github.com/pulumi/pulumi-hugo.
* Templates for generating package documentation. You'll find these at https://github.com/pulumi/pulumi.

## Contributing

Before you get started, be sure to read the [contributing guide](CONTRIBUTING.md) and review our [code of conduct](CODE_OF_CONDUCT.md).

## Toolchain

We build the Pulumi website statically with Hugo, manage our Node.js dependencies with Yarn, and write most of our documentation in Markdown. Below is a list of the tools you'll need to run the website locally:

* [Go](https://golang.org/)
* [Hugo](https://gohugo.io)
* [Node.js](https://nodejs.org/en/)
* [Yarn](https://classic.yarnpkg.com/en/)

### CSS and JavaScript Tools

We also use a handful of tools for compiling and managing our CSS and JavaScript assets, including:

* [Sass](https://sass-lang.com/)
* [TailwindCSS](https://tailwindcss.com/)
* [Stencil.js](https://stenciljs.com/)
* [TypeScript](https://www.typescriptlang.org/)

You don't need to install these tools individually or globally; the scripts below will handle everything for you. But if you'd like to contribute any CSS or JavaScript, you'll probably want to understand how to work with each of these tools as well.

## Installing prerequisites

Run `make ensure` to check for the appropriate tools and versions and install any dependencies. The script will let you know if you're missing anything important.

```
make ensure
```

## Running the development server

Once you've run `make ensure` successfully, you're ready to run the development server. If you're only planning on writing Markdown or working with Hugo layouts, this command should be all you need:

```
make serve
```

When you do this, Hugo will load the latest versions of:

* The [pulumi/pulumi-hugo](https://github.com/pulumi/pulumi-hugo) module, which contains our marketing pages, some docs content, the blog, etc.
* The [pulumi/theme](https://github.com/pulumi/theme) module, which contains our CSS and JavaScript bundles (web components, styles, etc.).

... and then start a development server at http://localhost:1313. Any changes you make to the content, layouts, or other [Hugo component folders](https://gohugo.io/getting-started/directory-structure/) should be reloaded automatically.

### Developing alongside another Hugo module

If you want to develop another module alongside this one -- e.g., to make changes to Registry-specific CSS -- you can point your development server to a local clone of that module. To do so, first clone the repository you want to use, then add a `replace` line to the `go.mod` file at the root of _this_ repository to override the existing reference to the module temporarily. For instance, to make changes to [pulumi/theme](https://github.com/pulumi/theme) as you develop in this repo:

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

### Linking to pages that aren't in this repository

If you want to link to a page that exists on https://pulumi.com but not in this repository, you can still use a [Hugo `relref`](https://gohugo.io/content-management/shortcodes/#ref-and-relref) in the usual way -- you'll just need to make sure the path you're linking to does exist (or will exist by the time you merge your change). For example, to add a link pointing to the [Pulumi CLI documentation](https://www.pulumi.com/docs/reference/cli/) (which does not exist in this repository), you'd use:

```
{{< relref /docs/reference/cli >}}
```

... and just be aware while the link won't work for you locally, it will work once your change is merged and picked up by our website automation (see [Merging and releasing](#merging-and-releasing)] below for details). This works because we've suppressed Hugo's built-in `relref` validation locally to keep the module-development workflow as lightweight as possible.

### Temporarily pulling in content from pulumi/docs

If the change you're working on requires content from pulumi/docs -- e.g., the aforementioned how-to guides or package documentation -- you may want to be able to see some of that content as you develop. To do that, just copy the files you need from the [`content` folder of pulumi/docs](https://github.com/pulumi/docs/tree/master/content) into the `content` folder of this repository (and/or the [`static` folder](https://github.com/pulumi/docs/tree/master/content), if you're working on API Docs changes), if you're looking for the JSON files that power the API Docs navigation), remembering to remove those files before you commit. For example:

```
# Copy the AWS API docs and navigation from a local/sibling clone of pulumi/docs.
cp -R ../docs/content/registry/packages/aws/api-docs ./themes/default/content/registry/packages/aws/
cp ../docs/static/registry/packages/navs/aws.json ./themes/default/static/registry/packages/navs/
```

Again, just be sure to remove these files before you commit:

```
# Remove the files we just added.
rm -rf ./themes/default/content/registry/packages/aws/api-docs/
rm -f ./themes/default/static/registry/packages/navs/aws.json
```

## Linting and testing

To check your code and your Markdown files for problems before submitting, run:

```
make lint test
```

## Tidying up

To clear away any module caches or build artifacts, run:

```
make clean
```

## Submitting and merging pull requests

When you're ready to submit a pull request, make sure you've removed anything that doesn't seem to belong (`go.mod`/`go.sum` changes, content from pulumi/docs, etc.) and [submit the PR](pulls) in the usual way.

If you're doing work in another repository that's associated with the changes in your PR, you can "pin" your PR branch to another module repository branch by pointing Hugo to that branch. To do that, use `hugo mod get` and pass a reference to the target branch:

```
hugo mod get github.com/pulumi/theme@my-special-branch
```

This will modify `go.mod` and `go.sum` accordingly and result in a PR preview that incorporates your changes from the other branch. Just be sure to remove these changes before you're ready to merge.

Once your PR is approved and merged into the default branch of this repository, an automated process that runs on [pulumi/docs](https://github.com/pulumi/docs) will detect your changes and produce a PR and integration build containing content from all other Hugo modules. Once that PR build completes and is approved and merged into pulumi/docs, your changes will be deployed to https://pulumi.com.

<p align="center">
    <img src="https://user-images.githubusercontent.com/274700/139970838-36931fe4-451c-4cd9-ae72-1178a6ef1537.png" width="75%">
</p>
