# Adding a new package

This is the complete process for getting a package into the Pulumi Registry. It covers both sides: what a package author submits, and what a Pulumi maintainer verifies before merging.

## Which path applies

Most packages are added by pull request. That is the path for a provider published as a [Pulumi Package](https://www.pulumi.com/docs/guides/pulumi-packages/) with a committed `schema.json` — Native, Component, or a statically bridged Terraform provider you have built into your own provider repo. You do not need to file an issue first.

A **dynamically bridged** Terraform provider is different. If you consume it with `pulumi package add terraform-provider <name>` and there is no provider repo or committed schema, it cannot be added by pull request, because those are listed through a separate Pulumi pipeline. Open a ["New Package"](https://github.com/pulumi/registry/issues/new?template=new-package.yml) issue to request one instead. Use the same issue to request a package you do not maintain, or to discuss before opening a PR.

For help with either path, reach out on the [Pulumi community Slack](https://slack.pulumi.com/) or use the [contact form](https://pulumi.com/contact/?form=registry).

## For package authors

New to authoring a package? Start with the [guide on authoring a Pulumi Package](https://www.pulumi.com/docs/guides/pulumi-packages/how-to-author/).

To publish a community-maintained package on the Pulumi Registry:

1. Ensure your provider repo has `docs/_index.md`, the only documentation page we require. It renders as the index page for your provider. It must begin with YAML front matter, before any other text — doc generation rejects a file that does not, and your package will not publish. See [the Overview page guidelines](./overview-page.md) for the front-matter keys and the sections the page needs.
1. Set a User-Agent header on your provider's API client that identifies the provider and its version, e.g. `pulumi-your-package/1.2.3`. Vendors decide how much to invest in an integration based on the traffic they can attribute to it, and a provider sending its SDK's default user agent doesn't show up in that count. Bridged providers are the worst case: one that falls through to the vendored HTTP client identifies as OpenTofu or Terraform instead of Pulumi. Where you set this depends on the SDK you wrap — look for the client constructor's user-agent or "application name" option. Pulumi has no canonical mechanism for this and there's no `tfbridge.ProviderInfo` field that does it for you, so if you're bridging a Terraform provider and can't find a hook, ask on the [Pulumi community Slack](https://slack.pulumi.com/).
1. Create a release of your provider in GitHub with a "v" + [Semver 2.0](https://semver.org) compliant tag (`vX.Y.Z`).
1. Open a pull request that adds one entry to [`community-packages/package-list.json`](https://github.com/pulumi/registry/blob/master/community-packages/package-list.json). That single entry is the whole registration: your docs and metadata are generated and published for you after merge, so do not commit generated files here. The one exception is a brand-new publisher, whose display name must be added to [`publisher-names.json`](https://github.com/pulumi/registry/blob/master/tools/resourcedocsgen/pkg/publishers/publisher-names.json) in the same PR, because publishing fails without it.

Automated checks then post a fact-sheet on your PR. If anything is flagged, fix it in your provider repo and comment `/check` to re-run — the checks read your live upstream, so you do not push here to re-validate. A maintainer can comment `/preview` to build a live preview of your package's pages, then reviews the fact-sheet and approves. Nothing merges automatically.

## For maintainers

To keep quality in the Pulumi Registry high, we have a check-list before merging a new provider into the registry. Please post a copy of this checklist in the PR under review and check off each item as verified.

- [ ] Pulumi has appropriate contact information from the provider maintainer

  If the provider is community maintained (maintained by a person, not a company), then a GitHub handle is sufficient

  If the provider is maintained by a company, Pulumi needs a contact person at the maintaining company.

- [ ] The package will generate accurate documentation:

  1. Check out the PR under review and run:

  ```sh
  $ make bin/resourcedocsgen
  $ ./bin/resourcedocsgen metadata from-github \
          --repoSlug '<repoSlug>' \
          --schemaFile '<schemaFile>' \
          --version '<version>'
  ```

  Here `<repoSlug>` and `<schemaFile>` should match exactly the values added to `/community-packages/package-list.json`.

  This will generate metadata for the provider locally.

  1. Push the metadata files into a PR (either back to the PR under review or a new PR).

  After pushing the provider to CI and waiting for a preview site:

  - [ ] Confirm that that CI passes **for the PR with the metadata files**.

  - [ ] Click through the site preview and confirm that the docs (for the new provider) render as expected.

  - [ ] The registry renders a valid logo for the new provider.

- [ ] Hand-written docs are complete and accurate:

  - [ ] `/docs/_index.md` — the only documentation page we require — meets [the Overview page guidelines](./overview-page.md).
  - [ ] Validate that you can actually install the provider by following its Installation section. Run the command given — `pulumi plugin install resource <name> <version> --server <pluginDownloadURL>` for a package shipping its own plugin binary, or `pulumi package add <...>` — and see a provider be downloaded.
  - [ ] The provider sets a user agent identifying itself to the vendor API.

- [ ] There is a published version:
  - [ ] The repository has a version tag prefixed with `v` that corresponds with a valid GitHub release
  - [ ] Each published SDK has a matching release

- [ ] A CODEOWNER has approved the PR.

## Adding the package

Most submissions arrive as a pull request the author opened themselves, following the steps above; work that PR through the checklist and merge it. When a community member requests a package they cannot submit — a dynamically bridged provider, or one they do not maintain — a Pulumi staff member opens the `community-packages/package-list.json` PR on their behalf, then validates and merges it the same way.

In pulumi/docs, a [scheduled task](https://github.com/pulumi/docs/actions/workflows/update-theme.yml) runs hourly and will pick up any changes in this repo, generate files from the provider schema and `data/registry/${PROVIDER}.yaml`, and publish to pulumi.com.

  This scheduled task currently lacks adequate monitoring, and **should be watched to ensure that it runs correctly to completion**. (If it fails, it will block all updates to pulumi.com, including marketing and manually maintained docs pages.)

## Search keywords

The filter box on the registry index page matches a package by its title, its name, and its keywords. The keywords come from the `keywords` field of the provider schema. `resourcedocsgen` copies them into `themes/default/data/registry/packages/<name>.yaml`, without the `pulumi` keyword and without tags such as `category/network` or `kind/native`.

To make a package findable by a new term, add the term to the `keywords` field in the provider schema and release a new version. The next metadata publish picks it up. Do not edit the package YAML file by hand.
