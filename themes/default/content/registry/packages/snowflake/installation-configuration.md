---
title: Snowflake Installation & Configuration
meta_desc: Provides an overview on how to configure credentials for the Pulumi Snowflake Provider.
layout: package
---

The Pulumi Snowflake provider uses the Snowflake SDK to manage and provision resources.

## Installation

The Snowflake provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/snowflake`](https://www.npmjs.com/package/@pulumi/snowflake)
* Python: [`pulumi-snowflake`](https://pypi.org/project/pulumi-snowflake/)
* Go: [`github.com/pulumi/pulumi-snowflake/sdk/go/snowflake`](https://github.com/pulumi/pulumi-snowflake)
* .NET: [`Pulumi.Snowflake`](https://www.nuget.org/packages/Pulumi.Snowflake)

## Authentication against Snowflake

The Snowflake provider support multiple ways to authenticate:

* Password
* OAuth Access Token
* OAuth Refresh Token
* Browser Auth
* Private Key

## Configuring Credentials

Pulumi relies on the Snowflake SDK to authenticate requests from your computer to Snowflake. Your credentials are never sent
to pulumi.com.
The Pulumi Snowflake Provider needs to be configured with Snowflake credentials
before it can be used to create resources.

Once the credentials are obtained, there are two ways to communicate your authorization tokens to Pulumi:

1. Set the environment variables `SNOWFLAKE_ACCOUNT`, `SNOWFLAKE_REGION` and `SNOWFLAKE_USERNAME` with the correct combination of authentication variables:

    ```bash
    $ export SNOWFLAKE_ACCOUNT=XXXXXXXXXXXXXX
    $ export SNOWFLAKE_REGION=YYYYYYYYYYYYYY
    $ export SNOWFLAKE_USERNAME=ZZZZZZZZZZZZZZ
    ```

2. Set them using configuration, if you prefer that they be stored alongside your Pulumi stack for easy multi-user access:

    ```bash
    $ pulumi config set snowflake:account XXXXXXXXXXXXXX
    $ pulumi config set snowflake:region YYYYYYYYYYYYYY
    $ pulumi config set snowflake:username ZZZZZZZZZZZZZZ
    ```

Remember to pass `--secret` when setting any secret keys so that they are properly encrypted. The complete list of
configuration parameters is in the [Snowflake Resource provider README](https://github.com/pulumi/pulumi-snowflake/blob/master/README.md).
