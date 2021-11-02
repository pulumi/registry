---
title: Datadog Setup
meta_desc: Provides an overview on how to configure credentials for the Pulumi Datadog Provider.
layout: installation
---

The Pulumi Datadog provider uses the Datadog SDK to manage and provision resources.

## Installation

The Datadog provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/datadog`](https://www.npmjs.com/package/@pulumi/datadog)
* Python: [`pulumi-datadog`](https://pypi.org/project/pulumi-datadog/)
* Go: [`github.com/pulumi/pulumi-datadog/sdk/v3/go/datadog`](https://github.com/pulumi/pulumi-datadog)
* .NET: [`Pulumi.Datadog`](https://www.nuget.org/packages/Pulumi.Datadog)

## Configuring Credentials

Pulumi relies on the Datadog SDK to authenticate requests from your computer to Datadog. Your credentials are never sent
to pulumi.com. The Pulumi Datadog Provider needs to be configured with Datadog credentials
before it can be used to create resources.

Once the credetials are obtained, there are two ways to communicate your authorization tokens to Pulumi:

1. Set the environment variables `DATADOG_API_KEY` and `DATADOG_APP_KEY`:

    ```bash
    $ export DATADOG_API_KEY=XXXXXXXXXXXXXX
    $ export DATADOG_APP_KEY=YYYYYYYYYYYYYY
    ```

2. Set them using configuration, if you prefer that they be stored alongside your Pulumi stack for easy multi-user access:

    ```bash
    $ pulumi config set datadog:apiKey XXXXXXXXXXXXXX --secret
    $ pulumi config set datadog:appKey YYYYYYYYYYYYYY --secret
    ```

Remember to pass `--secret` when setting `datadog:apiKey` and `datadog:appKey` so that they are properly encrypted. The complete list of
configuration parameters is in the [Datadog provider README](https://github.com/pulumi/pulumi-datadog/blob/master/README.md).
