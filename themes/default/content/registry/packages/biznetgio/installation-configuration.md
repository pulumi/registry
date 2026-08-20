---
# WARNING: this file was fetched from https://raw.githubusercontent.com/shirasakaren/pulumi-biznetgio/v0.1.7/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/shirasakaren/pulumi-biznetgio/blob/v0.1.7/docs/installation-configuration.md
title: BiznetGIO Installation & Configuration
meta_desc: Information on how to install and configure the BiznetGIO provider.
layout: installation
---

## Installation

The BiznetGIO provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@shirasakaren/biznetgio`](https://www.npmjs.com/package/@shirasakaren/biznetgio)
* Python: [`pulumi-biznetgio`](https://pypi.org/project/pulumi-biznetgio/) (import as `pulumi_biznetgio`)
* Go: [`github.com/shirasakaren/pulumi-biznetgio/sdk/go/pulumi-biznetgio`](https://pkg.go.dev/github.com/shirasakaren/pulumi-biznetgio/sdk/go/pulumi-biznetgio)
* .NET: [`Shirasakaren.Biznetgio`](https://www.nuget.org/packages/Shirasakaren.Biznetgio)
* Java: [`ren.shirasaka:biznetgio`](https://central.sonatype.com/artifact/ren.shirasaka/biznetgio)

The provider plugin binary is downloaded automatically from GitHub Releases. To install it directly:

```sh
pulumi plugin install resource biznetgio <version> --server github://api.github.com/shirasakaren/pulumi-biznetgio
```

## Setup

Get an API token from the [BiznetGIO portal](https://portal.biznetgio.com), then set it as a secret config value:

```sh
pulumi config set --secret biznetgio:apiToken <token>
```

Or export it as an environment variable instead: `BIZNETGIO_API_KEY`.

## Configuration Options

Use `pulumi config set biznetgio:<option>`.

| Option | Required/Optional | Description |
|---|---|---|
| `apiToken` | Required | BiznetGIO API token, sent as the `x-token` header. Environment: `BIZNETGIO_API_KEY`. |
| `baseUrl` | Optional | API base URL. Defaults to `https://api.portal.biznetgio.com/v1`. Environment: `BIZNETGIO_BASE_URL`. |

> **Billing note**: `payWithCreditCard` defaults to `true` on order/upgrade calls, so the first `pulumi up` charges
> the credit card on file. Set it to `false` to leave the order `Pending` for manual payment in the portal instead.
