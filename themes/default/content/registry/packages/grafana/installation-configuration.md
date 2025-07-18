---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumiverse/pulumi-grafana/v0.19.2/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Grafana Cloud Installation & Configuration
meta_desc: Information on how to install the Grafana provider.
layout: package
---

## Installation

The Pulumi Grafana provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumiverse/grafana`](https://www.npmjs.com/package/@pulumiverse/grafana)
* Python: [`pulumiverse_grafana`](https://pypi.org/project/pulumiverse-grafana/)
* Go: [`github.com/pulumiverse/pulumi-grafana/sdk/go/grafana`](https://pkg.go.dev/github.com/pulumiverse/pulumi-grafana/sdk)
* .NET: [`Pulumiverse.Grafana`](https://www.nuget.org/packages/Pulumiverse.Grafana)

### Provider Binary

The Grafana provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource grafana <version>  --server github://api.github.com/pulumiverse
```

Replace the version string with your desired version.

### Migrating from the LbrLabs package

The maintenance of this provider has been transferred from LbrLabs to Pulumiverse.
LbrLabs published up to v0.2.0, where Pulumiverse picks up with an initial v0.2.1
containing the renamed packages.

If you were using the LbrLabs edition, please update your dependencies to the
Pulumiverse edition:

| Programming Language | LbrLabs name | Pulumiverse name |
| -- | -- | -- |
| JavaScript/TypeScript | `@lbrlabs/pulumi-grafana` | `@pulumiverse/grafana` |
| Python | `lbrlabs_pulumi_grafana` | `pulumiverse_grafana` |
| Go | `github.com/lbrlabs/pulumi-grafana/sdk/go/grafana` | `github.com/pulumiverse/pulumi-grafana/sdk/go/grafana` |
| .NET | `Lbrlabs.PulumiPackage.Grafana` | `Pulumiverse.Grafana` |

## Setup

To provision resources with the Pulumi Grafana provider, you need to have Grafana credentials. 

### Set environment variables

Once you have provisioned these credentials, you can set environment variables to provision resources in Grafana:

{{< chooser os "linux,macos,windows" >}}
{{% choosable os linux %}}

```bash
$ export GRAFANA_CLOUD_ACCESS_POLICY_TOKEN=<GRAFANA_CLOUD_ACCESS_POLICY_TOKEN>

$ export GRAFANA_URL=<GRAFANA_URL>
$ export GRAFANA_AUTH=<GRAFANA_AUTH>
```

{{% /choosable %}}

{{% choosable os macos %}}

```bash
$ export GRAFANA_CLOUD_ACCESS_POLICY_TOKEN=<GRAFANA_CLOUD_ACCESS_POLICY_TOKEN>

$ export GRAFANA_URL=<GRAFANA_URL>
$ export GRAFANA_AUTH=<GRAFANA_AUTH>
```

{{% /choosable %}}

{{% choosable os windows %}}

```powershell
> $env:GRAFANA_CLOUD_ACCESS_POLICY_TOKEN = "<GRAFANA_CLOUD_ACCESS_POLICY_TOKEN>"

> $env:GRAFANA_URL = "<GRAFANA_URL>"
> $env:GRAFANA_AUTH = "<GRAFANA_AUTH>"
```

{{% /choosable %}}
{{< /chooser >}}

## Configuration Options

Use `pulumi config set grafana:<option>` or pass options to the [constructor of `new grafana.Provider`]({{< relref "/registry/packages/grafana/api-docs/provider" >}}).
