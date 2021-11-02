---
title: RabbitMQ Setup
meta_desc: Provides an overview on how to configure credentials for the Pulumi RabbitMQ Provider.
layout: installation
---

The Pulumi RabbitMQ provider uses the RabbitMQ SDK to manage and provision resources.
Pulumi relies on the RabbitMQ SDK to authenticate requests from your computer to RabbitMQ. Your credentials are never sent
to pulumi.com.
The Pulumi RabbitMQ Provider needs to be configured with RabbitMQ credentials
before it can be used to create resources.

## Installation

The RabbitMQ provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/rabbitmq`](https://www.npmjs.com/package/@pulumi/rabbitmq)
* Python: [`pulumi-rabbitmq`](https://pypi.org/project/pulumi-rabbitmq/)
* Go: [`github.com/pulumi/pulumi-rabbitmq/sdk/v3/go/rabbitmq`](https://github.com/pulumi/pulumi-rabbitmq)
* .NET: [`Pulumi.Rabbitmq`](https://www.nuget.org/packages/Pulumi.Rabbitmq)

## Configuring Credentials

Once the credetials are obtained, there are two ways to communicate your authorization tokens to Pulumi:

1. Set the environment variables `RABBITMQ_ENDPOINT`, `RABBITMQ_USERNAME` and `RABBITMQ_PASSWORD`:

    ```bash
    $ export RABBITMQ_ENDPOINT=XXXXXXXXXXXXXX
    $ export RABBITMQ_USERNAME=YYYYYYYYYYYYYY
    $ export RABBITMQ_PASSWORD=ZZZZZZZZZZZZZZ
    ```

2. Set them using configuration, if you prefer that they be stored alongside your Pulumi stack for easy multi-user access:

    ```bash
    $ pulumi config set rabbitmq:endpoint XXXXXXXXXXXXXX
    $ pulumi config set rabbitmq:username YYYYYYYYYYYYYY --secret
    $ pulumi config set rabbitmq:password ZZZZZZZZZZZZZZ --secret
    ```

Remember to pass `--secret` when setting `rabbitmq:username` and `rabbitmq:password` so that they are properly encrypted. The complete list of
configuration parameters is in the [RabbitMQ provider README](https://github.com/pulumi/pulumi-rabbitmq/blob/master/README.md).
