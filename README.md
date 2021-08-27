# pulumi-hugo-internal

A private repo for collaborating on content before publishing on pulumi.com. See below for details and usage guidelines.

## Installing tools and prerequisites

Run `make ensure` to check for the appropriate tools and versions and install any dependencies. The script will let you know if you're missing anything important.

```
make ensure
```

## Running Hugo locally

Once you've run `make ensure` successfully, you're ready to run the development server. If you're only planning on writing Markdown or working with Hugo layouts, this command should be all you need:

```
make serve
```

You can browse the development server at http://localhost:1313, and any changes you make to content or layouts should be reloaded automatically.

## Tidying up

To clear away any module caches or build artifacts, run:

```
make clean
```

## About this repository

This repo is a [Hugo module](https://gohugo.io/hugo-modules/) that acts as a supplement of https://github.com/pulumi/pulumi-hugo. The module itself has no dependencies, but the Hugo website defined at the root of this repository depends on both pulumi-hugo and on the module defined at `./themes/default`.

When Hugo builds a website, it resolves all modules and collapses them into a single union file system, according to the mappings and precedence rules defined in its local configuration. In this repo, when you run `make serve`, Hugo fetches the pulumi-hugo module and merges its files with those at `./themes/default`, giving preference to local content where it exists.

So if you were to browse to the home page of this repo's dev server, what you'd see by default would be based on pulumi-hugo's `content/_index.md` file (and whatever `layouts` and other [Hugo components](https://gohugo.io/hugo-modules/configuration/#module-config-mounts) were involved in the rendering of it) &mdash; unless there happened to be a file located at `./themes/default/content/_index.md`, in which case you'd see the content of that file instead. Basically, the content and layouts in this repo mask those of pulumi-hugo.

## Pull requests and previews

When a pull request is submitted on this repository, a preview is generated at a **publicly accessible URL**, just like with pulumi-hugo. So please be careful sharing these URLs, as the content in this repo should generally be considered private. Merging or closing a pull request destroys any associated previews automatically.

## Merging and releasing

**This repository is not yet set up to publish automatically to pulumi.com.** If your pull request is for content you'd like to release on pulumi.com, you'll need to close it &mdash; without merging &mdash; and submit it as a new pull request on https://github.com/pulumi/pulumi-hugo. From there, it'll be merged and released in the normal way.

At some point we'll build in some automation that handles this step for you automatically. But for now, please continue using this repo as a convenient way to preview and collaborate in private, and use pulumi-hugo when you're ready to publish.

## Blogging

See the [blogging README](BLOGGING.md) for details.

## Shortcodes and web components

See the [components README](themes/default/components).

## Administrative miscellany

### Updating go.mod

The aim of this repo is to track with the latest version of pulumi-hugo, so it updates its `go.mod` and `go.sum` files accordingly. But because these two files are usually irrelevant to users of this repo, they're skipped via [`git update-index --skip-worktree`](https://git-scm.com/docs/git-update-index#_skip_worktree_bit). This helps keep the working tree clean and uncluttered with files you probably don't care about.

If you do need to update `go.mod`, however, just reconfigure your local index to track changes to `go.mod`:

```
git update-index --no-skip-worktree go.mod go.sum
```

.. make your changes, commit (making sure not to pin any modules that shouldn't be pinned), and then reapply `skip-worktree` to ignore future changes:

```
git update-index --skip-worktree go.mod go.sum
```
