---
title: Prodvana Installation & Configuration
meta_desc: Provides an overview of how to install and configure the Prodvana provider.
layout: package
---

## Installation

The Prodvana provider is available as a package in the following Pulumi languages:

* JavaScript/TypeScript: [`@prodvana/pulumi-prodvana`](https://www.npmjs.com/package/@prodvana/pulumi-prodvana)
* Python: [`pulumi_prodvana`](https://pypi.org/project/pulumi-prodvana/)
* Go: [`github.com/prodvana/pulumi-prodvana/sdk/go/prodvana`](https://pkg.go.dev/github.com/prodvana/pulumi-prodvana/sdk)

### Provider Binary

The provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource prodvana <version> --server github://api.github.com/prodvana
```

Replace the version string with your desired version.

## Setup

To provision resources with the Pulumi Prodvana provider, you need to pass in your Prodvana organization's `orgSlug` and a valid `apiToken`. See [these instructions](https://docs.prodvana.io/docs/api-tokens-1) for creating an API Token.

## Configuration Options

Use `pulumi config set prodvana:<option> (--secret)`.

| Option | Environment Variable | Required/Optional | Description |
|-----|------|------|----|
| `orgSlug`| `PVN_ORG_SLUG` | Required | Prodvana Org Slug, found in your Prodvana url `<org_slug>.prodvana.io` |
| `apiToken`| `PVN_API_TOKEN` | Required (Secret) | API Token with permissions to the Prodvana Organization |
