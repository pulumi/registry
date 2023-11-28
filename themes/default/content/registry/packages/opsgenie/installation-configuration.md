---
title: Opsgenie Installation & Configuration
meta_desc: Provides an overview on how to configure credentials for the Pulumi Opsgenie Provider.
layout: package
---

The Pulumi Opsgenie provider uses the Opsgenie SDK to manage and provision resources.

## Installation

The Opsgenie provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/opsgenie`](https://www.npmjs.com/package/@pulumi/opsgenie)
* Python: [`pulumi-opsgenie`](https://pypi.org/project/pulumi-opsgenie/)
* Go: [`github.com/pulumi/pulumi-opsgenie/sdk/go/opsgenie`](https://github.com/pulumi/pulumi-opsgenie)
* .NET: [`Pulumi.Opsgenie`](https://www.nuget.org/packages/Pulumi.Opsgenie)
* Java: [`com.pulumi/opsgenie`](https://central.sonatype.com/artifact/com.pulumi/opsgenie)

## Configuring Credentials

Pulumi relies on the Opsgenie SDK to authenticate requests from your computer to Opsgenie. Your credentials are never sent
to pulumi.com.
The Pulumi Opsgenie Provider needs to be configured with Opsgenie credentials
before it can be used to manage resources.

In order to communicate your configuration details to Pulumi:

1. Set the environment variables `OPSGENIE_API_KEY` and (an optional) `OPSGENIE_API_URL`:

    ```bash
    $ export OPSGENIE_API_KEY=XXXXXXXXXXXXXX
    $ export OPSGENIE_API_URL=YYYYYYYYYYYYYY
    ```

1. Set them using configuration, if you prefer that they be stored alongside your Pulumi stack for easy multi-user access:

    ```bash
    $ pulumi config set opsgenie:apiKey XXXXXXXXXXXXXX --secret
    $ pulumi config set opsgenie:apiUrl YYYYYYYYYYYYYY
    ```

If you are going to set `opsgenie:apiKey`, please remember to pass `--secret` so that it is properly encrypted. The complete list of
configuration parameters is in the [Opsgenie provider README](https://github.com/pulumi/pulumi-opsgenie/blob/master/README.md).
