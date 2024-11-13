---
# WARNING: this file was fetched from https://raw.githubusercontent.com/lbrlabs/pulumi-nxos/v0.0.2/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Cisco NX OS Installation & Configuration
meta_desc: Information on how to install the Cisco NX OS.
layout: package
---

## Installation

The Pulumi NX OS provider is available as a package in all Pulumi languages:

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

To provision resources with the Pulumi Cisco nxos XE provider, you need to have credentials. 

### Set environment variables

Once you have provisioned these credentials, you can set environment variables to provision resources in Grafana:

{{< chooser os "linux,macos,windows" >}}
{{% choosable os linux %}}

```bash
$ export NXOS_PASSWORD=<PASSWORD>
$ export NXOS_URL=<URL>
$ export NXOS_USERNAME=<USERNAME>
```

{{% /choosable %}}

{{% choosable os macos %}}

```bash
$ export NXOS_PASSWORD=<PASSWORD>
$ export NXOS_URL=<URL>
$ export NXOS_USERNAME=<USERNAME>
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
