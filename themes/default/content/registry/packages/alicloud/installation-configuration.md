---
title: Alibaba Cloud Setup
meta_desc: Provides an overview on how to configure credentials for the Pulumi Alibaba Cloud Provider.
layout: installation
---

The Pulumi Alibaba Cloud provider uses the Alibaba Cloud SDK to manage and provision resources.

## Installation

The Alibaba Cloud provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/alicloud`](https://www.npmjs.com/package/@pulumi/alicloud)
* Python: [`pulumi-alicloud`](https://pypi.org/project/pulumi-alicloud/)
* Go: [`github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud`](https://github.com/pulumi/pulumi-alicloud)
* .NET: [`Pulumi.Alicloud`](https://www.nuget.org/packages/Pulumi.Alicloud)

## Configuring Credentials

Pulumi relies on the Alibaba Cloud SDK to authenticate requests from your computer to Alibaba Cloud. Your credentials are never sent
to pulumi.com. The Pulumi Alibaba Cloud Provider needs to be configured with Alibaba Cloud credentials
before it can be used to create resources.

Once the credentials are obtained, there are two ways to communicate your authorization tokens to Pulumi:

1. Set the environment variables `ALICLOUD_ACCESS_KEY` and `ALICLOUD_SECRET_KEY`:

    ```bash
    $ export ALICLOUD_ACCESS_KEY=XXXXXXXXXXXXXX
    $ export ALICLOUD_SECRET_KEY=YYYYYYYYYYYYYY
    ```

2. Set them using configuration, if you prefer that they be stored alongside your Pulumi stack for easy multi-user access:

    ```bash
    $ pulumi config set alicloud:accessKey XXXXXXXXXXXXXX --secret
    $ pulumi config set alicloud:secretKey YYYYYYYYYYYYYY --secret
    ```

Remember to pass `--secret` when setting `alicloud:secretKey` and `alicloud:accessKe` so that they are properly encrypted.
The complete list of
configuration parameters is in the [Alibaba Cloud provider](https://github.com/pulumi/pulumi-alicloud/blob/master/README.md).
