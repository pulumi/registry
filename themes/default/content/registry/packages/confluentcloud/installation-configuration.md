---
title: Confluent Cloud Installation & Configuration
meta_desc: Provides an overview on how to configure credentials for the Pulumi Confluent Cloud Provider.
layout: installation
---

The Pulumi Confluent Cloud provider uses the Confluent Cloud SDK to manage and provision resources.

## Installation

The Confluent Cloud provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/confluentcloud`](https://www.npmjs.com/package/@pulumi/confluentcloud)
* Python: [`pulumi-confluentcloud`](https://pypi.org/project/pulumi-confluentcloud/)
* Go: [`github.com/pulumi/pulumi-confluentcloud/sdk/go/confluentcloud`](https://github.com/pulumi/pulumi-confluentcloud)
* .NET: [`Pulumi.ConfluentCloud`](https://www.nuget.org/packages/Pulumi.ConfluentCloud)

## Configuring Credentials

Pulumi relies on the Confluent Cloud SDK to authenticate requests from your computer to Confluent Cloud. Your credentials are never sent
to pulumi.com. The Pulumi Confluent Cloud Provider needs to be configured with Confluent Cloud credentials
before it can be used to create resources. Once the credentials are obtained, there are two ways to communicate your authorization tokens to Pulumi:

1. Set the environment variables `CONFLUENT_CLOUD_API_KEY` and `CONFLUENT_CLOUD_API_SECRET`:

    ```bash
    $ export CONFLUENT_CLOUD_API_KEY=XXXXXXXXXXXXXX
    $ export CONFLUENT_CLOUD_API_SECRET=YYYYYYYYYYYYYY
    ```

2. Set them using configuration, if you prefer that they be stored alongside your Pulumi stack for easy multi-user access:

    ```bash
    $ pulumi config set confluentcloud:cloudApiKey XXXXXXXXXXXXXX
    $ pulumi config set confluentcloud:clouaApiSecret YYYYYYYYYYYYYY --secret
    ```

Remember to pass `--secret` when setting `confluentcloud:cloudApiKey` and `confluentcloud:cloudApiSecret` so that it is properly encrypted. The complete list of
configuration parameters is in the [Confluent Cloud provider README](https://github.com/pulumi/pulumi-confluentcloud/blob/master/README.md).
