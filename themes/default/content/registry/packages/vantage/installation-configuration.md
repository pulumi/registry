---
title: Vantage Installation & Configuration
meta_desc: Information on how to install the Vantage provider.
layout: installation
---

## Installation

The Pulumi Vantage provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@lbrlabs/pulumi-vantage`](https://www.npmjs.com/package/@lbrlabs/pulumi-vantage)
* Python: [`lbrlabs_pulumi_vantage`](https://pypi.org/project/lbrlabs-pulumi-vantage/)
* Go: [`github.com/lbrlabs/pulumi-vantage/sdk/go/vantage`](https://pkg.go.dev/github.com/lbrlabs/pulumi-vantage/sdk)
* .NET: [`Lbrlabs.PulumiPackage.vantage`](https://www.nuget.org/packages/Lbrlabs.PulumiPackage.vantage)

### Provider Binary

The Vantage provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource vantage <version> --server github://api.github.com/lbrlabs
```

Replace the version string with your desired version.

## Setup

To provision resources with the Pulumi Vantage provider, you need to have Vrafana credentials. 

### Set environment variables

Once you have provisioned these credentials, you can set environment variables to provision resources in Vantage:

{{< chooser os "linux,macos,windows" >}}
{{% choosable os linux %}}

```bash
$ export export VANTAGE_API_TOKEN="vntg_tkn_xxxxx"
$ export VANTAGE_API_HOST="https://api.vantage.sh"
```

{{% /choosable %}}

{{% choosable os macos %}}

```bash
$ export export VANTAGE_API_TOKEN="vntg_tkn_xxxxx"
$ export VANTAGE_API_HOST="https://api.vantage.sh"
```

{{% /choosable %}}

{{% choosable os windows %}}

```powershell
> $env:VANTAGE_API_TOKEN = "vntg_tkn_xxxxx"
> $env:VANTAGE_API_HOST = "https://api.vantage.sh"
```

{{% /choosable %}}
{{< /chooser >}}

## Configuration Options

Use `pulumi config set vantage:<option>` or pass options to the [constructor of `new vantage.Provider`]({{< relref "/registry/packages/vantage/api-docs/provider" >}}).
