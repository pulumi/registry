# Adding a new package

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

  - [ ] Validate that you can install the new provider with the instructions found in `/docs/installation-configuration.md`.

  Maintainers should run the `pulumi plugin install resource <name> <version> --server <pluginDownloadURL>` command specified in the `/docs/installation-configuration.md` and see a provider be downloaded.

  - [ ] `/docs/installation-configuration.md` contains a link to the published SDK in each language (i.e. TypeScript, Python, Go and C#).

  - [ ] `/docs/_index.md` contains a minimal example in every supported language.
  - [ ] `/docs/_index.md` contains a brief explanation of what the package does.

- [ ] There is a published version:
  - [ ] The repository has a version tag prefixed with `v` that corresponds with a valid GitHub release
  - [ ] Each published SDK has a matching release

- [ ] A CODEOWNER has approved the PR.
