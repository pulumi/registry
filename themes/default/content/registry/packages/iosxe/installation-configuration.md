---
title: Cisco IOS XE Installation & Configuration
meta_desc: Information on how to install the Cisco IOS XE.
layout: package
---

## Installation

The Pulumi Grafana provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@lbrlabs/pulumi-iosxe`](https://www.npmjs.com/package/@lbrlabs/pulumi-iosxe)
* Python: [`lbrlabs_pulumi_iosxe`](https://pypi.org/project/lbrlabs-pulumi-iosxe/)
* Go: [`github.com/lbrlabs/pulumi-iosxe/sdk/go/iosxe`](https://pkg.go.dev/github.com/lbrlabs/pulumi-iosxe/sdk)
* .NET: [`Lbrlabs.PulumiPackage.Iosxe`](https://www.nuget.org/packages/Lbrlabs.PulumiPackage.Iosxe)

### Provider Binary

The Grafana provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource iosxe <version> --server github://api.github.com/lbrlabs
```

Replace the version string with your desired version.

## Setup

To provision resources with the Pulumi Cisco IOS XE provider, you need to have credentials. 

### Set environment variables

Once you have provisioned these credentials, you can set environment variables to provision resources in Grafana:

{{< chooser os "linux,macos,windows" >}}
{{% choosable os linux %}}

```bash
$ export IOSXE_PASSWORD=<PASSWORD>
$ export IOSXE_URL=<URL>
$ export IOSXE_USERNAME=<USERNAME>
```

{{% /choosable %}}

{{% choosable os macos %}}

```bash
$ export IOSXE_PASSWORD=<PASSWORD>
$ export IOSXE_URL=<URL>
$ export IOSXE_USERNAME=<USERNAME>
```

{{% /choosable %}}

{{% choosable os windows %}}

```powershell
> $env:IOSXE_PASSWORD = "<PASSWORD>"
> $env:IOSXE_URL = "<URL>"
> $env:IOSXE_USERNAME = "<USERNAME>"
```

{{% /choosable %}}
{{< /chooser >}}

## Configuration Options

Use `pulumi config set iosxe:<option>` or pass options to the [constructor of `new iosxe.Provider`]({{< relref "/registry/packages/iosxe/api-docs/provider" >}}).
