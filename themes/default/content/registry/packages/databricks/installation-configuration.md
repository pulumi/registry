---
title: Databricks Installation & Configuration
meta_desc: Provides an overview on how to configure credentials for the Pulumi Databricks Provider.
layout: installation
---

The Pulumi Databricks provider uses the Databricks SDK to manage and provision resources.

## Installation

The Databricks provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/databricks`](https://www.npmjs.com/package/@pulumi/databricks)
* Python: [`pulumi-databricks`](https://pypi.org/project/pulumi-databricks/)
* Go: [`github.com/pulumi/pulumi-databricks/sdk/go/databricks`](https://github.com/pulumi/pulumi-databricks)
* .NET: [`Pulumi.Databricks`](https://www.nuget.org/packages/Pulumi.Databricks)

## Configuring Credentials

Pulumi relies on the Databricks SDK to authenticate requests from your computer to Databricks. Your credentials are never sent
to pulumi.com. The Pulumi Databricks Provider needs to be configured with Databricks credentials
before it can be used to create resources. Once the credentials obtained, there are two ways to communicate your authorization tokens to Pulumi:

1. Set the environment variables `DATABRICKS_HOST` and `DATABRICKS_TOKEN`:

    ```bash
    $ export DATABRICKS_HOST=XXXXXXXXXXXXXX
    $ export DATABRICKS_TOKEN=YYYYYYYYYYYYYY
    ```

2. Set them using configuration, if you prefer that they be stored alongside your Pulumi stack for easy multi-user access:

    ```bash
    $ pulumi config set databricks:token YYYYYYYYYYYYYY --secret
    $ pulumi config set databricks:host XXXXXXXXXXXXXX
    ```

Remember to pass `--secret` when setting `databricks:token` so that it is properly encrypted. The complete list of
configuration parameters is in the [Databricks Provider README](https://github.com/pulumi/pulumi-databricks/blob/master/README.md).
