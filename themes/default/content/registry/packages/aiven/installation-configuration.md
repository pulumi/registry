---
title: Aiven Setup
meta_desc: Provides an overview on how to configure credentials for the Pulumi Aiven Provider.
layout: installation
---

The Pulumi Aiven provider uses the Aiven SDK to manage and provision resources.

## Installation

The Aiven provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/aiven`](https://www.npmjs.com/package/@pulumi/aiven)
* Python: [`pulumi-aiven`](https://pypi.org/project/pulumi-aiven/)
* Go: [`github.com/pulumi/pulumi-aiven/sdk/v4/go/aiven`](https://github.com/pulumi/pulumi-aiven)
* .NET: [`Pulumi.Aiven`](https://www.nuget.org/packages/Pulumi.Aiven)

## Configuring Credentials

Pulumi relies on the Aiven SDK to authenticate requests from your computer to Aiven. Your credentials are never sent
to pulumi.com. The Pulumi Aiven Provider needs to be configured with Aiven credentials
before it can be used to create resources. Once the credentials obtained, there are two ways to communicate your authorization tokens to Pulumi:

1. Set the environment variable `AIVEN_TOKEN`:

    ```bash
    $ export AIVEN_TOKEN=XXXXXXXXXXXXXX
    ```

2. Set them using configuration, if you prefer that they be stored alongside your Pulumi stack for easy multi-user access:

    ```bash
    $ pulumi config set aiven:apiToken XXXXXXXXXXXXXX --secret
    ```

Remember to pass `--secret` when setting `aiven:apiToken` so that it is properly encrypted. The complete list of
configuration parameters is in the [Avien Provider README](https://github.com/pulumi/pulumi-aiven/blob/master/README.md).
