---
title: Civo Setup
meta_desc: Provides an overview on how to configure credentials for the Pulumi Civo Provider.
layout: installation
---

The Pulumi Civo provider uses the Civo SDK to manage and provision resources.

## Installation

The Civo provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/civo`](https://www.npmjs.com/package/@pulumi/civo)
* Python: [`pulumi-civo`](https://pypi.org/project/pulumi-civo/)
* Go: [`github.com/pulumi/pulumi-civo/sdk/go/civo`](https://github.com/pulumi/pulumi-civo)
* .NET: [`Pulumi.Civo`](https://www.nuget.org/packages/Pulumi.Civo)

## Configuring Credentials

Pulumi relies on the Civo SDK to authenticate requests from your computer to Civo. Your credentials are never sent
to pulumi.com. The Pulumi Civo Provider needs to be configured with Civo credentials
before it can be used to create resources. Once the credentials are obtained, there are two ways to communicate your authorization tokens to Pulumi:

1. Set the environment variable `CIVO_TOKEN`:

    ```bash
    $ export CIVO_TOKEN=XXXXXXXXXXXXXX
    ```

2. Set them using configuration, if you prefer that they be stored alongside your Pulumi stack for easy multi-user access:

    ```bash
    $ pulumi config set civo:token XXXXXXXXXXXXXX --secret
    ```

Remember to pass `--secret` when setting `civo:token` so that it is properly encrypted. The complete list of
configuration parameters is in the [Civo provider README](https://github.com/pulumi/pulumi-civo/blob/master/README.md).
