---
title: Wavefront Setup
meta_desc: Provides an overview on how to configure credentials for the Pulumi Wavefront Provider.
layout: installation
---

The Pulumi Wavefront provider uses the Wavefront SDK to manage and provision resources.

## Installation

The Wavefront provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/wavefront`](https://www.npmjs.com/package/@pulumi/wavefront)
* Python: [`pulumi-wavefront`](https://pypi.org/project/pulumi-wavefront/)
* Go: [`github.com/pulumi/pulumi-wavefront/sdk/go/wavefront`](https://github.com/pulumi/pulumi-wavefront)
* .NET: [`Pulumi.Wavefront`](https://www.nuget.org/packages/Pulumi.Wavefront)

## Configuring Credentials

Pulumi relies on the Wavefront SDK to authenticate requests from your computer to Wavefront. Your credentials are never sent
to pulumi.com.
The Pulumi Wavefront Provider needs to be configured with Wavefront credentials
before it can be used to create resources.

Once the credentials are obtained, there are two ways to communicate your authorization tokens to Pulumi:

1. Set the environment variables for `WAVEFRONT_TOKEN` and `WAVEFRONT_ADDRESS`:

    ```bash
    $ export WAVEFRONT_TOKEN=XXXXXXXXXXXXXX
    $ export WAVEFRONT_ADDRESS=YYYYYYYYYYYYYY
    ```

2. Set them using configuration, if you prefer that they be stored alongside your Pulumi stack for easy multi-user access:

    ```bash
    $ pulumi config set wavefront:token XXXXXXXXXXXXXX --secret
    $ pulumi config set wavefront:address YYYYYYYYYYYYYY
    ```

Remember to pass `--secret` when setting `wavefront:token` so that it is properly encrypted. The complete list of
configuration parameters is in the [Wavefront provider README](https://github.com/pulumi/pulumi-wavefront/blob/master/README.md).
