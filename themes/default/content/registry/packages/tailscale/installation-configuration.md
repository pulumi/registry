---
title: Tailscale Setup
meta_desc: Provides an overview on how to configure the Pulumi Tailscale Provider.
layout: installation
---

The Pulumi Tailscale provider uses the Tailscale SDK to manage resources.

## Installation

The Tailscale provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/tailscale`](https://www.npmjs.com/package/@pulumi/tailscale)
* Python: [`pulumi-tailscale`](https://pypi.org/project/pulumi-tailscale/)
* Go: [`github.com/pulumi/pulumi-tailscale/sdk/go/tailscale`](https://github.com/pulumi/pulumi-tailscale)
* .NET: [`Pulumi.Tailscale`](https://www.nuget.org/packages/Pulumi.Tailscale)

### Configuring The Provider

Pulumi relies on the Tailscale SDK to authenticate requests from your computer to Tailscale. Your credentials are never sent
to pulumi.com. Once the credentials are obtained, there are two ways to communicate your configuration parameters to Pulumi:

1. Set the environment variable `TAILSCALE_API_KEY` and `TAILSCALE_TAILNET`:

    ```bash
    $ export TAILSCALE_API_KEY=XXXXXX
    $ export TAILSCALE_TAILNET=YYYYYY
    ```

2. If you prefer that they be stored alongside your Pulumi stack for easy multi-user access:

    ```bash
    $ pulumi config set tailscale:apiKey XXXXXX --secret
    $ pulumi config set tailscale:tailnet YYYYYY
    ```

 The complete list of
configuration parameters is in the [Tailscale provider README](https://github.com/pulumi/pulumi-tailscale/blob/master/README.md).
