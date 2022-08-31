---
title: SignalFx Setup
meta_desc: Provides an overview on how to configure credentials for the Pulumi SignalFx Provider.
layout: installation
---

The Pulumi SignalFx provider uses the SignalFx SDK to manage and provision resources.

## Installation

The SignalFx provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/signalfx`](https://www.npmjs.com/package/@pulumi/signalfx)
* Python: [`pulumi-signalfx`](https://pypi.org/project/pulumi-signalfx/)
* Go: [`github.com/pulumi/pulumi-signalfx/sdk/v5/go/signalfx`](https://github.com/pulumi/pulumi-signalfx)
* .NET: [`Pulumi.Signalfx`](https://www.nuget.org/packages/Pulumi.Signalfx)

## Configuring Credentials

Pulumi relies on the SignalFx SDK to authenticate requests from your computer to SignalFx. Your credentials are never sent
to pulumi.com.
The Pulumi SignalFx Provider needs to be configured with SignalFx credentials
before it can be used to create resources.

Once the credentials are obtained, there are two ways to communicate your authorization tokens to Pulumi:

1. Set the environment variables `SFX_AUTH_TOKEN`:

    ```bash
    $ export SFX_AUTH_TOKEN=XXXXXXXXXXXXXX
    ```

2. Set them using configuration, if you prefer that they be stored alongside your Pulumi stack for easy multi-user access:

    ```bash
    $ pulumi config set signalfx:authToken XXXXXXXXXXXXXX --secret
    ```

Remember to pass `--secret` when setting `signalfx:authToken` so that it is properly encrypted. The complete list of
configuration parameters is in the [SignalFx provider README](https://github.com/pulumi/pulumi-signalfx/blob/master/README.md).
