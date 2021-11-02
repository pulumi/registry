---
title: Okta Setup
meta_desc: Provides an overview on how to configure the Pulumi Okta Provider.
layout: installation
---

The Pulumi Okta provider uses the Okta SDK to manage resources.
Pulumi relies on the Okta SDK to authenticate requests from your computer to Okta. Your credentials are never sent
to pulumi.com.

## Installation

The Okta provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/okta`](https://www.npmjs.com/package/@pulumi/okta)
* Python: [`pulumi-okta`](https://pypi.org/project/pulumi-okta/)
* Go: [`github.com/pulumi/pulumi-okta/sdk/v3/go/okta`](https://github.com/pulumi/pulumi-okta)
* .NET: [`Pulumi.Okta`](https://www.nuget.org/packages/Pulumi.Okta)

## Configuring The Provider

Once the credetials are obtained, there are two ways to communicate your configuration tokens to Pulumi:

1. Set the environment variables `OKTA_ORG_NAME`, `OKTA_BASE_URL` and `OKTA_API_TOKEN`:

    ```bash
    $ export OKTA_ORG_NAME=XXXXXX
    $ export OKTA_BASE_URL=YYYYYY
    $ export OKTA_API_TOKEN=ZZZZZZ
    ```

2. Set them using configuration, if you prefer that they be stored alongside your Pulumi stack for easy multi-user access:

    ```bash
    $ pulumi config set okta:orgName XXXXXX
    $ pulumi config set okta:baseUrl YYYYYY
    $ pulumi config set --secret okta:apiToken ZZZZZZ
    ```

Remember to pass `--secret` when setting `apiToken` so that it is properly encrypted.
