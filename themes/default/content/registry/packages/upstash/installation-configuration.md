---
# WARNING: this file was fetched from https://raw.githubusercontent.com/upstash/pulumi-upstash/v0.4.0/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Upstash Installation & Configuration
meta_desc: Information on how to install the Upstash provider.
layout: package
---

## Installation

The Pulumi Upstash provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@upstash/pulumi`](https://www.npmjs.com/package/@upstash/pulumi)
* Go: [`github.com/upstash/pulumi-upstash/sdk/go/upstash`](https://pkg.go.dev/github.com/upstash/pulumi-upstash/sdk)
* Python: [`upstash_pulumi`](https://pypi.org/project/upstash-pulumi/)

### Provider Binary

The Upstash provider can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource upstash --server https://github.com/upstash/pulumi-upstash/releases/download/v0.2.0
```

## Setup

To provision resources with the Pulumi Upstash provider, you need to have Upstash credentials. Upstash maintains documentation on how to create API keys [here](https://docs.upstash.com/redis/account/developerapi).

### Set environment variables

Once you have provisioned these credentials, you can set environment variables to provision resources in Upstash:

{{< chooser os "linux,macos,windows" >}}
{{% choosable os linux %}}

```bash
$ export UPSTASH_EMAIL=<UPSTASH_EMAIL>
$ export UPSTASH_API_KEY=<UPSTASH_API_KEY>
```

{{% /choosable %}}

{{% choosable os macos %}}

```bash
$ export UPSTASH_EMAIL=<UPSTASH_EMAIL>
$ export UPSTASH_API_KEY=<UPSTASH_API_KEY>
```

{{% /choosable %}}

{{% choosable os windows %}}

```powershell
> $env:UPSTASH_EMAIL = "<UPSTASH_EMAIL>"
> $env:UPSTASH_API_KEY = "<UPSTASH_API_KEY>"
```

{{% /choosable %}}
{{< /chooser >}}

## Configuration Options

Use `pulumi config set upstash:<option>` or pass options to the [constructor of `new upstash.Provider`](https://www.pulumi.com/registry/packages/upstash/api-docs/provider).

| Option | Required/Optional | Description |
|-----|------|----|
| `email`| Required | Upstash user email
| `apiKey`| Required | [Upstash Api key](https://console.upstash.com/account/api) |
