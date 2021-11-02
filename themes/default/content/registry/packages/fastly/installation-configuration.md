---
title: Fastly Setup
meta_desc: Provides an overview on how to configure credentials for the Pulumi Fastly Provider.
layout: installation
---

The Pulumi Fastly provider uses the Fastly SDK to manage and provision resources.

## Installation

The Fastly provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/fastly`](https://www.npmjs.com/package/@pulumi/fastly)
* Python: [`pulumi-fastly`](https://pypi.org/project/pulumi-fastly/)
* Go: [`github.com/pulumi/pulumi-fastly/sdk/v3/go/fastly`](https://github.com/pulumi/pulumi-fastly)
* .NET: [`Pulumi.Fastly`](https://www.nuget.org/packages/Pulumi.Fastly)

## Configuring Credentials

Pulumi relies on the Fastly SDK to authenticate requests from your computer to Fastly. Your credentials are never sent
to pulumi.com.
The Pulumi Fastly Provider needs to be configured with Fastly credentials
before it can be used to create resources.

Once the credetials are obtained, there are two ways to communicate your authorization tokens to Pulumi:

1. Set the environment variable `FASTLY_API_KEY`:

    ```bash
    $ export FASTLY_API_KEY=XXXXXXXXXXXXXX
    ```

2. Set them using configuration, if you prefer that they be stored alongside your Pulumi stack for easy multi-user access:

    ```bash
    $ pulumi config set fastly:apiKey XXXXXXXXXXXXXX --secret
    ```

Remember to pass `--secret` when setting `fastly:apiKey` so that it is properly encrypted. The complete list of
configuration parameters is in the [Fastly provider README](https://github.com/pulumi/pulumi-fastly/blob/master/README.md).
