---
title: MySQL Installation & Configuration
meta_desc: Provides an overview on how to configure credentials for the Pulumi MySQL Provider.
layout: installation
---

The Pulumi MySQL provider uses the MySQL SDK to manage and provision resources.

## Installation

The MySQL Atlas provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/mysql`](https://www.npmjs.com/package/@pulumi/mysql)
* Python: [`pulumi-mysql`](https://pypi.org/project/pulumi-mysql/)
* Go: [`github.com/pulumi/pulumi-mysql/sdk/v3/go/mysql`](https://github.com/pulumi/pulumi-mysql)
* .NET: [`Pulumi.Mysql`](https://www.nuget.org/packages/Pulumi.Mysql)

## Configuring Credentials

Pulumi relies on the MySQL SDK to authenticate requests from your computer to MySQL. Your credentials are never sent
to pulumi.com. The Pulumi MySQL Provider needs to be configured with MySQL credentials
before it can be used to manage resources.

In order to communicate your configuration details to Pulumi:

1. Set the environment variables `MYSQL_ENDPOINT` and `MYSQL_USERNAME`:

    ```bash
    $ export MYSQL_ENDPOINT=XXXXXXXXXXXXXX
    $ export MYSQL_USERNAME=YYYYYYYYYYYYYY
    ```

1. Set them using configuration, if you prefer that they be stored alongside your Pulumi stack for easy multi-user access:

    ```bash
    $ pulumi config set mysql:endpoint XXXXXXXXXXXXXX
    $ pulumi config set mysql:username YYYYYYYYYYYYYY
    ```

If you are going to set `mysql:password`, please remember to pass `--secret` so that it is properly encrypted.The complete list of
configuration parameters is in the [MySQL provider README](https://github.com/pulumi/pulumi-mysql/blob/master/README.md).
