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

To provision resources with the Pulumi Service provider, you need Pulumi Cloud credentials. The provider looks for an [access token](/docs/administration/concepts/access-tokens/) in this order:

1. The `pulumiservice:accessToken` configuration option.
2. The `PULUMI_ACCESS_TOKEN` environment variable.
3. The account you're logged in to with `pulumi login`.

If you've already run `pulumi login`, you don't need to configure anything else.

### Set environment variables

To override the token from your `pulumi login`, or to run without logging in (for example, in CI), set these environment variables. The Pulumi Service provider uses the same environment variables as the Pulumi CLI.

{{< chooser os "linux,macos,windows" >}}
{{% choosable os linux %}}

```bash
$ export PULUMI_ACCESS_TOKEN=<PULUMI_ACCESS_TOKEN>
# For self-hosted customers. Defaults to https://api.pulumi.com
$ export PULUMI_BACKEND_URL=<PULUMI_BACKEND_URL>
```

{{% /choosable %}}

{{% choosable os macos %}}

```bash
$ export PULUMI_ACCESS_TOKEN=<PULUMI_ACCESS_TOKEN>
# For self-hosted customers. Defaults to https://api.pulumi.com
$ export PULUMI_BACKEND_URL=<PULUMI_BACKEND_URL>
```

{{% /choosable %}}

{{% choosable os windows %}}

```powershell
> $env:PULUMI_ACCESS_TOKEN = "<PULUMI_ACCESS_TOKEN>"
# For self-hosted customers. Defaults to https://api.pulumi.com
> $env:PULUMI_BACKEND_URL = "<PULUMI_BACKEND_URL>"
```

{{% /choosable %}}
{{< /chooser >}}

## Configuration Options

Use `pulumi config set pulumiservice:<option>` or pass options to the [constructor of `new pulumiservice.Provider`](/registry/packages/pulumiservice/api-docs/provider).

| Option | Required? | Description |
|-----|------|----|
| `accessToken`| Optional | [Pulumi Cloud access token](/docs/administration/concepts/access-tokens/). Defaults to `PULUMI_ACCESS_TOKEN`, then to your `pulumi login` credentials. |
| `apiUrl`| Optional | Allows overriding default [Pulumi Cloud API URL](/docs/pulumi-cloud/cloud-rest-api/) for [self-hosted customers](/docs/pulumi-cloud/self-hosted/). Defaults to `PULUMI_BACKEND_URL`, then `PULUMI_API`, then `https://api.pulumi.com`. |
