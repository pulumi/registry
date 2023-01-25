---
title: ZPA Setup
meta_desc: Information on how to install the ZPA provider.
layout: installation
---

## Installation

The Pulumi ZPA provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@zscaler/pulumi-zpa`](https://www.npmjs.com/package/@zscaler/pulumi-zpa)
* Python: [`zscaler-pulumi-zpa`](https://pypi.org/project/zscaler-pulumi-zpa/)
* Go: [`github.com/zscaler/pulumi-zpa/sdk/go/zpa`](https://pkg.go.dev/github.com/zscaler/pulumi-zpa/sdk)
* .NET: [`zscaler.PulumiPackage.Zpa`](https://www.nuget.org/packages/zscaler.PulumiPackage.Zpa)

### Provider Binary

The ZPA provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource zpa <version> --server github://api.github.com/zscaler
```

Replace the version string with your desired version.

## Setup

To provision resources with the Pulumi ZPA provider, you need to have ZPA credentials. Zscaler maintains documentation on how to create API keys [here](https://help.zscaler.com/zpa/about-api-keys)

### Set environment variables

Once you have provisioned these credentials, you can set environment variables to provision resources in ZPA:

{{< chooser os "linux,macos,windows" >}}
{{% choosable os linux %}}

```bash
$ export ZPA_CLIENT_ID=<ZPA_CLIENT_ID>
$ export ZPA_CLIENT_SECRET=<ZPA_CLIENT_SECRET>
$ export ZPA_CUSTOMER_ID=<ZPA_CUSTOMER_ID>
```

{{% /choosable %}}

{{% choosable os macos %}}

```bash
$ export ZPA_CLIENT_ID=<ZPA_CLIENT_ID>
$ export ZPA_CLIENT_SECRET=<ZPA_CLIENT_SECRET>
$ export ZPA_CUSTOMER_ID=<ZPA_CUSTOMER_ID>
```

{{% /choosable %}}

{{% choosable os windows %}}

```powershell
> $env:ZPA_CLIENT_ID = "<ZPA_CLIENT_ID>"
> $env:ZPA_CLIENT_SECRET = "<ZPA_CLIENT_SECRET>"
> $env:ZPA_CUSTOMER_ID = "<ZPA_CUSTOMER_ID>"
```

{{% /choosable %}}
{{< /chooser >}}

## Configuration Options

Use `pulumi config set zpa:<option>` or pass options to the [constructor of `new zpa.Provider`]({{< relref "/registry/packages/zpa/api-docs/provider" >}}).

| Option | Required/Optional | Description |
|-----|------|----|
| `zpa_client_id`| Required | [ZPA Client ID](https://help.zscaler.com/zpa/about-api-keys) |
| `zpa_client_secret`| Required | [ZPA Client Secret](https://help.zscaler.com/zpa/about-api-keys) |
| `zpa_customer_id` | Required | [ZPA Customer ID](https://help.zscaler.com/zpa/about-api-keys) |
| `zpa_cloud` | Optional | [ZPA Cloud Name](https://registry.terraform.io/providers/zscaler/zpa/latest/docs) |
