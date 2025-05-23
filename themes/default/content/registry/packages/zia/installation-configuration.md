---
# WARNING: this file was fetched from https://raw.githubusercontent.com/zscaler/pulumi-zia/v0.0.7/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
title: ZIA Setup
meta_desc: Information on how to install the ZIA provider.
layout: installation
---

## Installation

The Pulumi ZIA provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@bdzscaler/pulumi-zia`](https://www.npmjs.com/package/@zscaler/pulumi-zia)
* Python: [`zscaler_pulumi_zia`](https://pypi.org/project/zscaler-pulumi-zia/)
* Go: [`github.com/zscaler/pulumi-zia/sdk/go/zia`](https://pkg.go.dev/github.com/zscaler/pulumi-zia/sdk)
* .NET: [`zscaler.PulumiPackage.Zia`](https://www.nuget.org/packages/zscaler.PulumiPackage.Zia)

### Provider Binary

The ZIA provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource zia <version> --server github://api.github.com/zscaler
```

Replace the version string with your desired version.

## Setup

To provision resources with the Pulumi ZIA provider, you need to have ZIA credentials. Zscaler maintains documentation on how to create API keys [here](https://help.zscaler.com/zia/getting-started-zia-api)

### Set environment variables

Once you have provisioned these credentials, you can set environment variables to provision resources in ZIA:

{{< chooser os "linux,macos,windows" >}}
{{% choosable os linux %}}

```bash
$ export ZIA_USERNAME=<ZIA_USERNAME>
$ export ZIA_PASSWORD=<ZIA_PASSWORD>
$ export ZIA_API_KEY=<ZIA_API_KEY>
$ export ZIA_CLOUD=<ZIA_CLOUD>
```

{{% /choosable %}}

{{% choosable os macos %}}

```bash
$ export ZIA_USERNAME=<ZIA_USERNAME>
$ export ZIA_PASSWORD=<ZIA_PASSWORD>
$ export ZIA_API_KEY=<ZIA_API_KEY>
$ export ZIA_CLOUD=<ZIA_CLOUD>
```

{{% /choosable %}}

{{% choosable os windows %}}

```powershell
> $env:ZIA_USERNAME = "<ZIA_USERNAME>"
> $env:ZIA_PASSWORD = "<ZIA_PASSWORD>"
> $env:ZIA_API_KEY = "<ZIA_API_KEY>"
> $env:ZIA_CLOUD = "<ZIA_CLOUD>"
```

{{% /choosable %}}
{{< /chooser >}}

## Configuration Options

Use `pulumi config set zia:<option>` or pass options to the [constructor of `new zia.Provider`]({{< relref "/registry/packages/zia/api-docs/provider" >}}).

| Option | Required/Optional | Description |
|-----|------|----|
| `username`| Required | [ZIA API Username](https://help.zscaler.com/zia/getting-started-zia-api) |
| `password`| Required | [ZIA API Password](https://help.zscaler.com/zia/getting-started-zia-api) |
| `api_key` | Required | [ZIA API Key](https://help.zscaler.com/zia/getting-started-zia-api) |
| `zia_cloud` | Optional | [ZIA Cloud](https://registry.terraform.io/providers/zscaler/zia/latest/docs) |

### Zscaler Sandbox Authentication

The ZIA provider requires both the `ZIA_CLOUD` and `ZIA_SANDBOX_TOKEN` in order to authenticate to the Zscaler Cloud Sandbox environment. For details on how obtain the API Token visit the Zscaler help portal [About Sandbox API Token](https://help.zscaler.com/zia/about-sandbox-api-token)