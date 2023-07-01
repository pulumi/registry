---
title: Cockroach Installation & Configuration
meta_desc: Information on how to install the Cockrach provider.
layout: installation
---

## Installation

The Pulumi Cockroach provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@lbrlabs/pulumi-cockroach`](https://www.npmjs.com/package/@lbrlabs/pulumi-cockroach)
* Python: [`lbrlabs_pulumi_cockroach`](https://pypi.org/project/lbrlabs-pulumi-cockroach/)
* Go: [`github.com/lbrlabs/pulumi-cockroach/sdk/go/cockroach`](https://pkg.go.dev/github.com/lbrlabs/pulumi-cockroach/sdk)
* .NET: [`Lbrlabs.PulumiPackage.Cockroach`](https://www.nuget.org/packages/Lbrlabs.PulumiPackage.Cockroach)

### Provider Binary

The Cockroach provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource cockroach <version> --server github://api.github.com/lbrlabs
```

Replace the version string with your desired version.

## Setup

To provision resources with the Pulumi Cockroach provider, you need to have Cockorach credentials. 

### Set environment variables

Once you have provisioned these credentials, you can set environment variables to provision resources in Cockorach:

{{< chooser os "linux,macos,windows" >}}
{{% choosable os linux %}}

```bash
$ export COCKROACH_API_KEY=<COCKROACH_API_KEY>
```

{{% /choosable %}}

{{% choosable os macos %}}

```bash
$ export COCKROACH_API_KEY=<COCKROACH_API_KEY>
```

{{% /choosable %}}

{{% choosable os windows %}}

```powershell
> $env:COCKROACH_API_KEY = "<COCKROACH_API_KEY>"
```

{{% /choosable %}}
{{< /chooser >}}

## Configuration Options

Use `pulumi config set cockroach:<option>` or pass options to the [constructor of `new cockroach.Provider`]({{< relref "/registry/packages/cockroach/api-docs/provider" >}}).
