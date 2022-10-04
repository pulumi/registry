---
title: Harness Setup
meta_desc: Information on how to install the Harness provider.
layout: installation
---

## Installation

The Pulumi Harness provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@lbrlabs/pulumi-harness`](https://www.npmjs.com/package/@lbrlabs/pulumi-harness)
* Python: [`lbrlabs_pulumi_harness`](https://pypi.org/project/lbrlabs-pulumi-harness/)
* Go: [`github.com/lbrlabs/pulumi-harness/sdk/go/harness`](https://pkg.go.dev/github.com/lbrlabs/pulumi-harness/sdk)
* .NET: [`Lbrlabs.PulumiPackage.Harness`](https://www.nuget.org/packages/Lbrlabs.PulumiPackage.Harness)

### Provider Binary

The Harness provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource harness <version>
```

Replace the version string with your desired version.

## Setup

To provision resources with the Pulumi Harness provider, you need to have Harness credentials. Harness maintains documentation on how to create API keys [here](https://docs.harness.io/article/smloyragsm-api-keys)

### Set environment variables

Once you have provisioned these credentials, you can set environment variables to provision resources in Harness:

{{< chooser os "linux,macos,windows" >}}
{{% choosable os linux %}}

```bash
$ export HARNESS_ACCOUNT_ID=<HARNESS_ACCOUNT_ID>
$ export HARNESS_API_KEY=<HARNESS_API_KEY>
$ export HARNESS_PLATFORM_API_KEY=<HARNESS_PLATFORM_API_KEY>
```

{{% /choosable %}}

{{% choosable os macos %}}

```bash
$ export HARNESS_ACCOUNT_ID=<HARNESS_ACCOUNT_ID>
$ export HARNESS_API_KEY=<HARNESS_API_KEY>
$ export HARNESS_PLATFORM_API_KEY=<HARNESS_PLATFORM_API_KEY>
```

{{% /choosable %}}

{{% choosable os windows %}}

```powershell
> $env:HARNESS_ACCOUNT_ID = "<HARNESS_ACCOUNT_ID>"
> $env:HARNESS_API_KEY = "<HARNESS_API_KEY>"
> $env:HARNESS_PLATFORM_API_KEY = "<HARNESS_PLATFORM_API_KEY>"
```

{{% /choosable %}}
{{< /chooser >}}

## Configuration Options

Use `pulumi config set harness:<option>` or pass options to the [constructor of `new harness.Provider`]({{< relref "/registry/packages/harness/api-docs/provider" >}}).

