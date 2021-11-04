---
title: Equinix Metal Setup
meta_desc: Provides an overview of how to setup the Equinix Metal SDK to manage and provision resources.
layout: installation
---

The Pulumi Equinix Metal provider uses the Equinix Metal SDK to manage and provision resources.

## Installation

The Equinix Metal provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/equinix-metal`](https://www.npmjs.com/package/@pulumi/equinix-metal)
* Python: [`pulumi-equinix-metal`](https://pypi.org/project/pulumi-equinix-metal/)
* Go: [`github.com/pulumi/pulumi-equinix-metal/sdk/v2/go/equinix`](https://github.com/pulumi/pulumi-equinix-metal)
* .NET: [`Pulumi.EquinixMetal`](https://www.nuget.org/packages/Pulumi.EquinixMetal)

## Configuring Credentials

Pulumi relies on the Equinix Metal SDK to authenticate requests from your computer to Equinix Metal. Your credentials are never sent
to pulumi.com. The Pulumi Equinix Metal Provider needs to be configured with Equinix Metal credentials
before it can be used to create resources.

Once the credetials are obtained, there are two ways to communicate your authorization tokens to Pulumi:

1. Set the environment variable `PACKET_AUTH_TOKEN`:

    ```bash
    $ export PACKET_AUTH_TOKEN=XXXXXXXXXXXXXX
    ```

2. Set them using configuration, if you prefer that they be stored alongside your Pulumi stack for easy multi-user access:

    ```bash
    $ pulumi config set equinix-metal:authToken XXXXXXXXXXXXXX --secret
    ```

Remember to pass `--secret` when setting `authToken` so that it is properly encrypted.
