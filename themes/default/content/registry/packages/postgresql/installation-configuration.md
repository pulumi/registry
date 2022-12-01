---
title: PostgreSQL Installation & Configuration
meta_desc: Provides an overview on how to configure credentials for the Pulumi PostgreSQL Provider.
layout: installation
---

The Pulumi PostgreSQL provider uses the PostgreSQL SDK to manage and provision resources.

## Installation

The PostgreSQL provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/postgresql`](https://www.npmjs.com/package/@pulumi/postgresql)
* Python: [`pulumi-postgresql`](https://pypi.org/project/pulumi-postgresql/)
* Go: [`github.com/pulumi/pulumi-postgresql/sdk/v3/go/postgresql`](https://github.com/pulumi/pulumi-postgresql)
* .NET: [`Pulumi.Postgresql`](https://www.nuget.org/packages/Pulumi.Postgresql)

## Configuring Credentials

Pulumi relies on the PostgreSQL SDK to authenticate requests from your computer to PostgreSQL. Your credentials are never sent
to pulumi.com.
The Pulumi PostgreSQL Provider needs to be configured with PostgreSQL credentials
before it can be used to manage resources.

In order to communicate your configuration details to Pulumi:

1. Set the environment variables `PGHOST` and `PGUSER`:

    ```bash
    $ export PGHOST=XXXXXXXXXXXXXX
    $ export PGUSER=YYYYYYYYYYYYYY
    ```

1. Set them using configuration, if you prefer that they be stored alongside your Pulumi stack for easy multi-user access:

    ```bash
    $ pulumi config set postgresql:host XXXXXXXXXXXXXX
    $ pulumi config set postgresql:username YYYYYYYYYYYYYY
    ```

If you are going to set `postgresql:password`, please remember to pass `--secret` so that it is properly encrypted. The complete list of
configuration parameters is in the [PostgreSQL provider README](https://github.com/pulumi/pulumi-postgresql/blob/master/README.md).
