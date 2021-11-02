---
title: Mailgun Setup
meta_desc: Provides an overview on how to configure credentials for the Pulumi Mailgun Provider.
layout: installation
---

The Pulumi Mailgun provider uses the Mailgun SDK to manage and provision resources.
Pulumi relies on the Mailgun SDK to authenticate requests from your computer to Mailgun. Your credentials are never sent
to pulumi.com.
The Pulumi Mailgun Provider needs to be configured with Mailgun credentials
before it can be used to create resources.

## Installation

The Mailgun provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/mailgun`](https://www.npmjs.com/package/@pulumi/mailgun)
* Python: [`pulumi-mailgun`](https://pypi.org/project/pulumi-mailgun/)
* Go: [`github.com/pulumi/pulumi-mailgun/sdk/v3/go/mailgun`](https://github.com/pulumi/pulumi-mailgun)
* .NET: [`Pulumi.Mailgun`](https://www.nuget.org/packages/Pulumi.Mailgun)

## Configuring Credentials

Once the credetials are obtained, there are two ways to communicate your authorization tokens to Pulumi:

1. Set the environment variable `MAILGUN_API_KEY`:

    ```bash
    $ export MAILGUN_API_KEY=XXXXXXXXXXXXXX
    ```

2. Set them using configuration, if you prefer that they be stored alongside your Pulumi stack for easy multi-user access:

    ```bash
    $ pulumi config set mailgun:apiKey XXXXXXXXXXXXXX --secret
    ```

Remember to pass `--secret` when setting `mailgun:apiKey` so that it is properly encrypted. The complete list of
configuration parameters is in the [Mailgun provider README](https://github.com/pulumi/pulumi-mailgun/blob/master/README.md).
