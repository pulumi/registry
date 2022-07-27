---
title: Confluent Cloud Setup
meta_desc: Provides an overview on how to configure credentials for the Pulumi Confluent Cloud Provider.
layout: installation
---

{{% notes type="info" %}}
This provider has been deprecated as of July 2022. It is recommended to use the [Official Confluent Provider]({{<relref "/registry/packages/confluentcloud">}}) as a replacement.
Unfortunately, there is no upgrade path from this provider to the Official Confluent provider, but you can take advantage of the [Pulumi Import]({{<relref "/docs/guides/adopting/import">}}) to help achieve the migration.
{{% /notes %}}

The Pulumi Confluent Cloud provider uses the Confluent Cloud SDK to manage and provision resources.

## Installation

The Confluent Cloud provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/confluent`](https://www.npmjs.com/package/@pulumi/confluent)
* Python: [`pulumi-confluent`](https://pypi.org/project/pulumi-confluent/)
* Go: [`github.com/pulumi/pulumi-confluent/sdk/go/confluent`](https://github.com/pulumi/pulumi-confluent)
* .NET: [`Pulumi.Confluent`](https://www.nuget.org/packages/Pulumi.Confluent)

## Configuring Credentials

Pulumi relies on the Confluent Cloud SDK to authenticate requests from your computer to Confluent Cloud. Your credentials are never sent
to pulumi.com. The Pulumi Confluent Cloud Provider needs to be configured with Confluent Cloud credentials
before it can be used to create resources. Once the credetials are obtained, there are two ways to communicate your authorization tokens to Pulumi:

1. Set the environment variables `CONFLUENT_CLOUD_USERNAME` and `CONFLUENT_CLOUD_PASSWORD`:

    ```bash
    $ export CONFLUENT_CLOUD_USERNAME=XXXXXXXXXXXXXX
    $ export CONFLUENT_CLOUD_PASSWORD=YYYYYYYYYYYYYY
    ```

2. Set them using configuration, if you prefer that they be stored alongside your Pulumi stack for easy multi-user access:

    ```bash
    $ pulumi config set confluent:username XXXXXXXXXXXXXX
    $ pulumi config set confluent:password YYYYYYYYYYYYYY --secret
    ```

Remember to pass `--secret` when setting `confluent:password` so that it is properly encrypted. The complete list of
configuration parameters is in the [Confluent Cloud provider README](https://github.com/pulumi/pulumi-confluent/blob/master/README.md).
