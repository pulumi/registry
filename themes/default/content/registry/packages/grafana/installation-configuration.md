---
title: Grafana Installation & Configuration
meta_desc: Information on how to install the Grafana provider.
layout: installation
---

## Installation

The Pulumi Grafana provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@lbrlabs/pulumi-grafana`](https://www.npmjs.com/package/@lbrlabs/pulumi-grafana)
* Python: [`lbrlabs_pulumi_grafana`](https://pypi.org/project/lbrlabs-pulumi-grafana/)
* Go: [`github.com/lbrlabs/pulumi-grafana/sdk/go/grafana`](https://pkg.go.dev/github.com/lbrlabs/pulumi-grafana/sdk)
* .NET: [`Lbrlabs.PulumiPackage.Grafana`](https://www.nuget.org/packages/Lbrlabs.PulumiPackage.Grafana)

### Provider Binary

The Grafana provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource grafana <version>
```

Replace the version string with your desired version.

## Setup

To provision resources with the Pulumi Grafana provider, you need to have Grafana credentials. 

### Set environment variables

Once you have provisioned these credentials, you can set environment variables to provision resources in Grafana:

{{< chooser os "linux,macos,windows" >}}
{{% choosable os linux %}}

```bash
$ export GRAFANA_URL=<GRAFANA_URL>
$ export GRAFANA_AUTH=<GRAFANA_AUTH>
```

{{% /choosable %}}

{{% choosable os macos %}}

```bash
$ export GRAFANA_URL=<GRAFANA_URL>
$ export GRAFANA_AUTH=<GRAFANA_AUTH>
```

{{% /choosable %}}

{{% choosable os windows %}}

```powershell
> $env:GRAFANA_URL = "<GRAFANA_URL>"
> $env:GRAFANA_AUTH = "<GRAFANA_AUTH>"
```

{{% /choosable %}}
{{< /chooser >}}

## Configuration Options

Use `pulumi config set grafana:<option>` or pass options to the [constructor of `new grafana.Provider`]({{< relref "/registry/packages/grafana/api-docs/provider" >}}).
