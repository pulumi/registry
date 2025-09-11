---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-harness/v0.8.4/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Harness Installation & Configuration
meta_desc: Information on how to install the Harness provider.
layout: package
---

## Installation

The Pulumi Harness provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/harness`](https://www.npmjs.com/package/@pulumi/harness)
* Python: [`pulumi_harness`](https://pypi.org/project/pulumi-harness/)
* Go: [`github.com/pulumi/pulumi-harness/sdk/go/harness`](https://pkg.go.dev/github.com/pulumi/pulumi-harness/sdk)
* .NET: [`Pulumi.Harness`](https://www.nuget.org/packages/Pulumi.Harness)

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

