---
title: F5 BIG-IP Setup
meta_desc: Provides an overview on how to configure credentials for the Pulumi F5 BIG-IP Provider.
layout: installation
---

The Pulumi F5 BIG-IP provider uses the F5 BIG-IP SDK to manage and provision resources.

> Pulumi relies on the F5 BIG-IP SDK to authenticate requests from your computer to the resources. Your credentials are never sent
> to pulumi.com.

The Pulumi F5 BIG-IP Provider needs to be configured with F5 BIG-IP credentials
before it can be used to create resources.

## Installation

The F5 BIG-IP provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/f5bigip`](https://www.npmjs.com/package/@pulumi/f5bigip)
* Python: [`pulumi-f5bigip`](https://pypi.org/project/pulumi-f5bigip/)
* Go: [`github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip`](https://github.com/pulumi/pulumi-f5bigip)
* .NET: [`Pulumi.F5bigip`](https://www.nuget.org/packages/Pulumi.F5bigip)

## Configuring Credentials

Once obtained, there are two ways to communicate your authorization tokens to Pulumi:

1. Set the environment variable `BIGIP_HOST`, `BIGIP_USER` and `BIGIP_PASSWORD`:

    ```bash
    $ export BIGIP_HOST=XXXXXXXXXXXXXX
    $ export BIGIP_USER=YYYYYYYYYYYYYY
    $ export BIGIP_PASSWORD=ZZZZZZZZZZZZZZ
    ```

2. Set them using configuration, if you prefer that they be stored alongside your Pulumi stack for easy multi-user access:

    ```bash
    $ pulumi config set f5bigip:address XXXXXXXXXXXXXX
    $ pulumi config set f5bigip:username YYYYYYYYYYYYYY
    $ pulumi config set f5bigip:password ZZZZZZZZZZZZZZ --secret
    ```

Remember to pass `--secret` when setting `f5bigip:password` so that it is properly encrypted. The complete list of
configuration parameters is in the [F5 BIG-IP provider README](https://github.com/pulumi/pulumi-f5bigip/blob/master/README.md).
