---
title: Kafka Setup
meta_desc: Provides an overview on how to configure credentials for the Pulumi Kafka Provider.
layout: installation
---

The Pulumi Kafka provider uses the Kafka SDK to manage and provision resources.

## Installation

The Kafka Cloud provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/kafka`](https://www.npmjs.com/package/@pulumi/kafka)
* Python: [`pulumi-kafka`](https://pypi.org/project/pulumi-kafka/)
* Go: [`github.com/pulumi/pulumi-kafka/sdk/v3/go/kafka`](https://github.com/pulumi/pulumi-kafka)
* .NET: [`Pulumi.Kafka`](https://www.nuget.org/packages/Pulumi.Kafka)

## Configuring Credentials

Pulumi relies on the Kafka SDK to authenticate requests from your computer to Kafka. Your credentials are never sent
to pulumi.com. The Pulumi Kafka Provider needs to be configured with Kafka credentials
before it can be used to create resources.

In order to communicate your configuration details to Pulumi:

1. Set them using configuration, if you prefer that they be stored alongside your Pulumi stack for easy multi-user access:

    ```bash
    $ pulumi config set kafka:bootstrapServers ["127.0.0.1:9092"]
    ```

The complete list of
configuration parameters is in the [Kafka provider README](https://github.com/pulumi/pulumi-kafka/blob/master/README.md).
