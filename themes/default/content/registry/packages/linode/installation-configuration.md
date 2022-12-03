---
title: Linode SDK Installation & Configuration
meta_desc: Provides an overview on how to setup the Linode SDK for Pulumi.
layout: installation
---

The Pulumi Linode provider uses the Linode SDK to manage and provision resources.

## Installation

The Linode provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/linode`](https://www.npmjs.com/package/@pulumi/linode)
* Python: [`pulumi-linode`](https://pypi.org/project/pulumi-linode/)
* Go: [`github.com/pulumi/pulumi-linode/sdk/v3/go/linode`](https://github.com/pulumi/pulumi-linode)
* .NET: [`Pulumi.Linode`](https://www.nuget.org/packages/Pulumi.Linode)

## Configuring Credentials

Pulumi relies on the Linode SDK to authenticate requests from your computer to Linode. Your credentials are never sent
to pulumi.com.
The Pulumi Linode Provider needs to be configured with Linode credentials
before it can be used to create resources.

Once the credentials are obtained, there are two ways to communicate your authorization tokens to Pulumi:

1. Set the environment variable `LINODE_TOKEN`:

    ```bash
    $ export LINODE_TOKEN=XXXXXXXXXXXXXX
    ```

2. Set them using configuration, if you prefer that they be stored alongside your Pulumi stack for easy multi-user access:

    ```bash
    $ pulumi config set linode:token XXXXXXXXXXXXXX --secret
    ```

Remember to pass `--secret` when setting `token` so that it is properly encrypted.
