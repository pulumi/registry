---
title: Splunk Setup
meta_desc: Provides an overview on how to configure credentials for the Pulumi Splunk Provider.
layout: installation
---

The Pulumi Splunk provider uses the Splunk SDK to manage and provision resources.

## Installation

The Splunk provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/splunk`](https://www.npmjs.com/package/@pulumi/splunk)
* Python: [`pulumi-splunk`](https://pypi.org/project/pulumi-splunk/)
* Go: [`github.com/pulumi/pulumi-splunk/sdk/go/splunk`](https://github.com/pulumi/pulumi-splunk)
* .NET: [`Pulumi.Splunk`](https://www.nuget.org/packages/Pulumi.Splunk)

## Configuring Credentials

Pulumi relies on the Splunk SDK to authenticate requests from your computer to Splunk. Your credentials are never sent
to pulumi.com.
The Pulumi Splunk Provider needs to be configured with Splunk credentials
before it can be used to create resources.

Once the credentials are obtained, there are two ways to communicate your authorization tokens to Pulumi:

1. Set the environment variables for `SPLUNK_URL`, `SPLUNK_USERNAME` and `SPLUNK_PASSWORD`:

    ```bash
    $ export SPLUNK_URL=XXXXXXXXXXXXXX
    $ export SPLUNK_USERNAME=YYYYYYYYYYYYYY
    $ export SPLUNK_PASSWORD=ZZZZZZZZZZZZZZ
    ```

2. Set them using configuration, if you prefer that they be stored alongside your Pulumi stack for easy multi-user access:

    ```bash
    $ pulumi config set splunk:url XXXXXXXXXXXXXX
    $ pulumi config set splunk:username YYYYYYYYYYYYYY
    $ pulumi config set splunk:password ZZZZZZZZZZZZZZ --secret
    ```

Remember to pass `--secret` when setting `splunk:password` so that it is properly encrypted. The complete list of
configuration parameters is in the [Splunk provider README](https://github.com/pulumi/pulumi-splunk/blob/master/README.md).
