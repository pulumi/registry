---
title: Kong Setup
meta_desc: Provides an overview on how to configure credentials for the Pulumi Kong Provider.
layout: installation
---

The [Pulumi Kong provider]({{< relref "./" >}}) uses the Kong SDK to manage and provision resources.
Pulumi relies on the Kong SDK to authenticate requests from your computer to Kong. Your credentials are never sent
to pulumi.com.
The [Pulumi Kong Provider]({{< relref "./" >}}) needs to be configured with Kong credentials
before it can be used to create resources.

## Installation

The Kong provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/kong`](https://www.npmjs.com/package/@pulumi/kong)
* Python: [`pulumi-kong`](https://pypi.org/project/pulumi-kong/)
* Go: [`github.com/pulumi/pulumi-kong/sdk/v4/go/kong`](https://github.com/pulumi/pulumi-kong)
* .NET: [`Pulumi.Kong`](https://www.nuget.org/packages/Pulumi.Kong)

## Configuring Credentials

Once obtained, there are two ways to communicate your authorization tokens to Pulumi:

1. Set the environment variable `KONG_ADMIN_ADDR`, `KONG_ADMIN_USERNAME` and `KONG_ADMIN_PASSWORD`:

    ```bash
    $ export KONG_ADMIN_ADDR=XXXXXXXXXXXXXX
    $ export KONG_ADMIN_USERNAME=YYYYYYYYYYYYYY
    $ export KONG_ADMIN_PASSWORD=YYYYYYYYYYYYYY
    ```

2. Set them using configuration, if you prefer that they be stored alongside your Pulumi stack for easy multi-user access:

    ```bash
    $ pulumi config set kong:kongAdminUri XXXXXXXXXXXXXX
    $ pulumi config set kong:kongAdminUsername YYYYYYYYYYYYYY
    $ pulumi config set kong:kongAdminPassword YYYYYYYYYYYYYY --secret
    ```

Remember to pass `--secret` when setting `kong:kongAdminPassword` so that it is properly encrypted. The complete list of
configuration parameters is in the [Kong provider README](https://github.com/pulumi/pulumi-kong/blob/master/README.md).
