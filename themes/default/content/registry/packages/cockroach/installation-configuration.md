---
title: CockroachDB Cloud Installation & Configuration
meta_desc: Provides an overview on how to configure credentials for the CockroachDB provider.
layout: package
---

## Installation

The Pulumi Cockroach provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumiverse/cockroach`](https://www.npmjs.com/package/@pulumiverse/cockroach)
* Python: [`pulumiverse_cockroach`](https://pypi.org/project/pulumiverse-cockroach/)
* Go: [`github.com/pulumiverse/pulumi-cockroach/sdk/go/cockroach`](https://github.com/pulumiverse/pulumi-cockroach)
* .NET: [`Pulumiverse.Cockroach`](https://www.nuget.org/packages/Pulumiverse.Cockroach)

### Provider Binary

The Cockroach provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource cockroach <version> --server github://api.github.com/pulumiverse
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
| JavaScript/TypeScript | `@lbrlabs/pulumi-cockroach` | `@pulumiverse/cockroach` |
| Python | `lbrlabs_pulumi_cockroach` | `pulumiverse_cockroach` |
| Go | `github.com/lbrlabs/pulumi-cockroach/sdk/go/cockroach` | `github.com/pulumiverse/pulumi-cockroach/sdk/go/cockroach` |
| .NET | `Lbrlabs.PulumiPackage.Cockroach` | `Pulumiverse.Cockroach` |

## Setup

To provision resources with the Pulumi Cockroach provider, you need to have Cockroach credentials. 

### Set environment variables

Once you have provisioned these credentials, you can set environment variables to provision resources in Cockroach:

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

Use `pulumi config set cockroach:<option>` or pass options to the [constructor of `new cockroach.Provider`](/registry/packages/cockroach/api-docs/provider/).
