---
title: NS1 Setup
meta_desc: Provides an overview on how to configure credentials for the Pulumi NS1 Provider.
layout: installation
---

The Pulumi NS1 provider uses the NS1 SDK to manage and provision resources.

## Installation

The NS1 provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/ns1`](https://www.npmjs.com/package/@pulumi/ns1)
* Python: [`pulumi-ns1`](https://pypi.org/project/pulumi-ns1/)
* Go: [`github.com/pulumi/pulumi-ns1/sdk/v2/go/ns1`](https://github.com/pulumi/pulumi-ns1)
* .NET: [`Pulumi.NS1`](https://www.nuget.org/packages/Pulumi.Ns1)

## Configuring Credentials

Pulumi relies on the NS1 SDK to authenticate requests from your computer to NS1. Your credentials are never sent
to pulumi.com.
The Pulumi NS1 Provider needs to be configured with NS1 credentials
before it can be used to create resources.

Once the credentials are obtained, there are two ways to communicate your authorization tokens to Pulumi:

1. Set the environment variable `NS1_APIKEY`:

    ```bash
    $ export NS1_APIKEY=XXXXXXXXXXXXXX
    ```

2. Set them using configuration, if you prefer that they be stored alongside your Pulumi stack for easy multi-user access:

    ```bash
    $ pulumi config set ns1:apikey XXXXXXXXXXXXXX --secret
    ```

Remember to pass `--secret` when setting `ns1:apiKey` so that it is properly encrypted. The complete list of
configuration parameters is in the
[NS1 provider README](https://github.com/pulumi/pulumi-ns1/blob/master/README.md).
