---
# WARNING: this file was fetched from https://raw.githubusercontent.com/UpCloudLtd/pulumi-upcloud/v0.9.0/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/UpCloudLtd/pulumi-upcloud/blob/v0.9.0/docs/installation-configuration.md
title: UpCloud provider installation & configuration
meta_desc: Instructions on how to install and configure the UpCloud provider.
layout: installation
---

## Installation

The UpCloud Pulumi provider SDK is available as a package in:

* JavaScript/TypeScript: [`@upcloud/pulumi-upcloud`](https://www.npmjs.com/package/@upcloud/pulumi-upcloud)
* Python: [`pulumi-upcloud`](https://pypi.org/project/pulumi-upcloud/)
* Go: [`github.com/UpCloudLtd/pulumi-upcloud/sdk/go/upcloud`](https://pkg.go.dev/github.com/UpCloudLtd/pulumi-upcloud/sdk)
* .NET: [`UpCloud.Pulumi.UpCloud`](https://www.nuget.org/packages/UpCloud.Pulumi.UpCloud/)

### Provider Binary

The UpCloud provider is a third party plugin. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource upcloud --server github://api.github.com/UpCloudLtd/pulumi-upcloud
```

Add version after upcloud if you need to install specific version, e.g., `pulumi plugin install resource upcloud v0.0.5 ...`.

## Configuration

Configure your UpCloud credentials to `UPCLOUD_USERNAME` and `UPCLOUD_PASSWORD` environment variables or your API token to `UPCLOUD_TOKEN` environment variable.
