---
title: Cloudflare Setup
meta_desc: Provides an overview on how to configure the Pulumi Cloudflare Provider.
layout: installation
---

The Pulumi Cloudflare provider uses the Cloudflare SDK to manage resources.

## Installation

The Cloudflare provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/cloudflare`](https://www.npmjs.com/package/@pulumi/cloudflare)
* Python: [`pulumi-cloudflare`](https://pypi.org/project/pulumi-cloudflare/)
* Go: [`github.com/pulumi/pulumi-cloudflare/sdk/v3/go/cloudflare`](https://github.com/pulumi/pulumi-cloudflare)
* .NET: [`Pulumi.Cloudflare`](https://www.nuget.org/packages/Pulumi.Cloudflare)

### Configuring The Provider

Pulumi relies on the Cloudflare SDK to authenticate requests from your computer to Cloudflare. Your credentials are never sent
to pulumi.com. Once the credentials are obtained, there are two ways to communicate your configuration parameters to Pulumi:

1. Set the environment variable `CLOUDFLARE_API_TOKEN` (or the legacy `CLOUDFLARE_EMAIL` and `CLOUDFLARE_API_KEY`):

    ```bash
    $ export CLOUDFLARE_API_TOKEN=YYYYYY
    # Legacy
    $ export CLOUDFLARE_EMAIL=XXXXXX
    $ export CLOUDFLARE_API_KEY=YYYYYY
    ```

{{% notes type="info" %}}
You can only use one of: `apiToken` _or_ `apiKey`/`email`. We recommend using `apiToken`.
{{% /notes %}}

2. If you prefer that they be stored alongside your Pulumi stack for easy multi-user access:

    ```bash
    $ pulumi config set cloudflare:apiToken
    # Legacy
    $ pulumi config set cloudflare:email XXXXXX
    $ pulumi config set cloudflare:apiKey YYYYYY --secret
    ```

 The complete list of
configuration parameters is in the [Cloudflare provider README](https://github.com/pulumi/pulumi-cloudflare/blob/master/README.md).
