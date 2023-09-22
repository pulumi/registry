---
title: Cisco IOS XE Installation & Configuration
meta_desc: Information on how to install the Cisco IOS XE.
layout: package
---

## Installation

The Pulumi Grafana provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@lbrlabs/pulumi-nxos`](https://www.npmjs.com/package/@lbrlabs/pulumi-nxos)
* Python: [`lbrlabs_pulumi_nxos`](https://pypi.org/project/lbrlabs-pulumi-nxos/)
* Go: [`github.com/lbrlabs/pulumi-nxos/sdk/go/nxos`](https://pkg.go.dev/github.com/lbrlabs/pulumi-nxos/sdk)
* .NET: [`Lbrlabs.PulumiPackage.nxos`](https://www.nuget.org/packages/Lbrlabs.PulumiPackage.nxos)

### Provider Binary

The Grafana provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource nxos <version> --server github://api.github.com/lbrlabs
```

Replace the version string with your desired version.

## Setup

To provision resources with the Pulumi Cisco IOS XE provider, you need to have credentials. 

### Set environment variables

Once you have provisioned these credentials, you can set environment variables to provision resources in Grafana:

{{< chooser os "linux,macos,windows" >}}
{{% choosable os linux %}}

```bash
$ export nxos_PASSWORD=<PASSWORD>
$ export nxos_URL=<URL>
$ export nxos_USERNAME=<USERNAME>
```

{{% /choosable %}}

{{% choosable os macos %}}

```bash
$ export nxos_PASSWORD=<PASSWORD>
$ export nxos_URL=<URL>
$ export nxos_USERNAME=<USERNAME>
```

{{% /choosable %}}

{{% choosable os windows %}}

```powershell
> $env:nxos_PASSWORD = "<PASSWORD>"
> $env:nxos_URL = "<URL>"
> $env:nxos_USERNAME = "<USERNAME>"
```

{{% /choosable %}}
{{< /chooser >}}

## Configuration Options

Use `pulumi config set nxos:<option>` or pass options to the [constructor of `new nxos.Provider`]({{< relref "/registry/packages/nxos/api-docs/provider" >}}).
