---
title: Hetzner Cloud Installation & Configuration
meta_desc: Provides an overview on how to configure credentials for the Pulumi Hetzner Cloud Provider.
layout: installation
---

The Pulumi Hetzner Cloud provider uses the Hetzner Cloud SDK to manage and provision resources.

## Installation

The Hetzner Cloud provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/hcloud`](https://www.npmjs.com/package/@pulumi/hcloud)
* Python: [`pulumi-hcloud`](https://pypi.org/project/pulumi-hcloud/)
* Go: [`github.com/pulumi/pulumi-hcloud/sdk/go/hcloud`](https://github.com/pulumi/pulumi-hcloud)
* .NET: [`Pulumi.HCloud`](https://www.nuget.org/packages/Pulumi.HCloud)

## Configuring Credentials

Pulumi relies on the Hetzner Cloud SDK to authenticate requests from your computer to Hetzner Cloud. Your credentials are never sent
to pulumi.com.
The Pulumi Hetzner Cloud Provider needs to be configured with Hetzner Cloud credentials
before it can be used to create resources.

Once the credentials are obtained, there are two ways to communicate your authorization tokens to Pulumi:

1. Set the environment variable `HCLOUD_TOKEN`:

    ```bash
    $ export HCLOUD_TOKEN=XXXXXXXXXXXXXX
    ```

2. Set them using configuration, if you prefer that they be stored alongside your Pulumi stack for easy multi-user access:

    ```bash
    $ pulumi config set hcloud:token XXXXXXXXXXXXXX --secret
    ```

Remember to pass `--secret` when setting `hcloud:token` so that it is properly encrypted. The complete list of
configuration parameters is in the
[Hetzner Cloud provider README](https://github.com/pulumi/pulumi-hcloud/blob/master/README.md).
