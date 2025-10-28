---
title: Pulumi Service Provider Installation & Configuration
meta_desc: Information on how to install the Pulumi Service provider.
layout: package
---

## Installation

The Pulumi Service provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/pulumiservice`](https://www.npmjs.com/package/@pulumi/pulumiservice)
* Python: [`pulumi_pulumiservice`](https://pypi.org/project/pulumi-pulumiservice/)
* Go: [`github.com/pulumi/pulumi-pulumiservice/sdk/go/pulumiservice`](https://pkg.go.dev/github.com/pulumi/pulumi-pulumiservice/sdk/go/pulumiservice)
* .NET: [`Pulumi.PulumiService`](https://www.nuget.org/packages/Pulumi.PulumiService)
* Java: [`com.pulumi/pulumiservice`](https://search.maven.org/artifact/com.pulumi/pulumiservice)

## Setup

To provision resources with the Pulumi Service provider, you must have Pulumi Cloud credentials. Pulumi Cloud maintains documentation on creating [access tokens](https://www.pulumi.com/docs/intro/pulumi-cloud/accounts/#access-tokens).
{{% notes "info" %}}
While you can use this provider to provision access tokens, you'll still need to have an access token available to generate an access token with the provider.
{{% /notes %}}

### Set environment variables

Once you have an access token, its easy to set the environment variables. The Pulumi Service Provider uses the same environment variables as Pulumi does.

{{< chooser os "linux,macos,windows" >}}
{{% choosable os linux %}}

```bash
$ export PULUMI_ACCESS_TOKEN=<PULUMI_ACCESS_TOKEN>
# For self hosted customers. defaults to https://api.pulumi.com
$ export PULUMI_BACKEND_URL=<PULUMI_BACKEND_URL>
```

{{% /choosable %}}

{{% choosable os macos %}}

```bash
$ export PULUMI_ACCESS_TOKEN=<PULUMI_ACCESS_TOKEN>
# For self hosted customers. defaults to https://api.pulumi.com
$ export PULUMI_BACKEND_URL=<PULUMI_BACKEND_URL>
```

{{% /choosable %}}

{{% choosable os windows %}}

```powershell
> $env:PULUMI_ACCESS_TOKEN = "<PULUMI_ACCESS_TOKEN>"
# For self hosted customers. defaults to https://api.pulumi.com
> $env:PULUMI_BACKEND_URL = "<PULUMI_BACKEND_URL>"
```

{{% /choosable %}}
{{< /chooser >}}

## Configuration Options

Use `pulumi config set pulumiservice:<option>` or pass options to the [constructor of `new pulumiservice.Provider`](/registry/packages/pulumiservice/api-docs/provider).

| Option | Required? | Description |
|-----|------|----|
| `accessToken`| Required | [Pulumi Cloud Access Tokens](/docs/pulumi-cloud/accounts/#personal-access-tokens) |
| `apiUrl`| Optional | Allows overriding default [Pulumi Cloud API URL](/docs/pulumi-cloud/cloud-rest-api/) for [self hosted customers](/docs/pulumi-cloud/self-hosted/).
