---
# WARNING: this file was fetched from https://raw.githubusercontent.com/genesiscloud/pulumi-genesiscloud/v0.0.35/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Genesis Cloud Installation & Configuration
meta_desc: Information on how to install the Genesis Cloud provider.
layout: package
---

## Installation

The Pulumi Genesis Cloud provider is available as a package in the following languages:

- JavaScript/TypeScript: [`@genesiscloud/pulumi-genesiscloud`](https://www.npmjs.com/package/@genesiscloud/pulumi-genesiscloud)
- Python: [`pulumi-genesiscloud`](https://pypi.org/project/pulumi-genesiscloud/)
- Go: [`github.com/genesiscloud/pulumi-genesiscloud/sdk/go/genesiscloud`](https://pkg.go.dev/github.com/genesiscloud/pulumi-genesiscloud/sdk/go)
- Dotnet: [GenesisCloud.PackagePulumi.Genesiscloud](https://www.nuget.org/packages/GenesisCloud.PulumiPackage.Genesiscloud)

### Provider Binary

The Genesis Cloud provider binary is a third party binary. It can be installed using the pulumi plugin command.

```bash
pulumi plugin install resource genesiscloud <version> --server github://api.github.com/genesiscloud
```

Replace the version string with your desired version.

## Setup

To provision resources with the Genesis Cloud provider, you need to have a Genesis Cloud token.

## Configuration Options

Use `pulumi config set genesiscloud:<option>`.

| Option     | Environment Variables | Required/Optional | Description                                                          |
| ---------- | --------------------- | ----------------- | -------------------------------------------------------------------- |
| `token`    | GENESISCLOUD_TOKEN    | Required          | Genesis Cloud access token                                           |
| `endpoint` | GENESISCLOUD_ENDPOINT | Optional          | The endpoint use in the provider. Defaults to `api.genesiscloud.com` |
