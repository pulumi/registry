## Description

<!-- A brief description of the PR here. -->

## Adding a new package?

If this pull request adds a new package:

- [ ] The package's schema URL is correct.
- [ ] The package metadata file, if present, contains: 
  - [ ] a supported category (one of `Cloud`, `Infrastructure`, `Network`, `Database`, `Monitoring`, or `Utility`).
  - [ ] a description that explains what the package does.
  - [ ] a valid logo URL that points to a PNG whose dimensions conform to the others in this repo (e.g., 100x100).
  - [ ] a version number prefixed with `v` that corresponds with a valid GitHub release.
- [ ] The package repo contains an Overview doc (`/docs/_index.md`) that includes:
  - [ ] a brief explanation of what the package is and what it does.
  - [ ] at least one representative example in all supported languages.
  - [ ] a front-matter property for the `layout` set to `package`.
- [ ] The package repo contains an Installation and Configuration doc (`/docs/installation-configuration.md`) that includes:
  - [ ] links to SDKs in all supported languages.
  - [ ] a copyable command for installing the resource plugin if necessary.
  - [ ] an example of configuring the provider with `pulumi config set`.
  - [ ] an example of configuring the provider with environment variables.
- [ ] Someone from the @pulumi/ai team has reviewed the PR.
- [ ] Someone from the @pulumi/docs team has reviewed all documentation.
