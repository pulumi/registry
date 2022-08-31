---
title: CloudAMQP Setup
meta_desc: Provides an overview on how to configure credentials for the Pulumi CloudAMQP Provider.
layout: installation
---

The Pulumi CloudAMQP provider uses the CloudAMQP SDK to manage and provision resources.

## Installation

The CloudAMQP provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/cloudamqp`](https://www.npmjs.com/package/@pulumi/cloudamqp)
* Python: [`pulumi-cloudamqp`](https://pypi.org/project/pulumi-cloudamqp/)
* Go: [`github.com/pulumi/pulumi-cloudamqp/sdk/v3/go/cloudamqp`](https://github.com/pulumi/pulumi-cloudamqp)
* .NET: [`Pulumi.Cloudamqp`](https://www.nuget.org/packages/Pulumi.Cloudamqp)

## Configuring Credentials

Pulumi relies on the CloudAMQP SDK to authenticate requests from your computer to CloudAMQP. Your credentials are never sent
to pulumi.com. The Pulumi CloudAMQP Provider needs to be configured with CloudAMQP credentials
before it can be used to create resources. Once the credentials are obtained, there are two ways to communicate your authorization tokens to Pulumi:

1. Set the environment variable `CLOUDAMQP_APIKEY`:

    ```bash
    $ export CLOUDAMQP_APIKEY=XXXXXXXXXXXXXX
    ```

2. Set them using configuration, if you prefer that they be stored alongside your Pulumi stack for easy multi-user access:

    ```bash
    $ pulumi config set cloudamqp:apikey XXXXXXXXXXXXXX --secret
    ```

Remember to pass `--secret` when setting `cloudamqp:apikey` so that it is properly encrypted. The complete list of
configuration parameters is in the [CloudAMQP README](https://github.com/pulumi/pulumi-cloudamqp/blob/master/README.md).
