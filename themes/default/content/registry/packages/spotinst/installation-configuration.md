---
title: Spotinst Setup
meta_desc: Provides an overview on how to configure the Pulumi Spotinst Provider.
layout: installation
---

The Pulumi Spotinst provider uses the Spotinst SDK to manage resources.
Pulumi relies on the Spotinst SDK to authenticate requests from your computer to Spotinst. Your credentials are never sent
to pulumi.com.

## Installation

The Spotinst provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/spotinst`](https://www.npmjs.com/package/@pulumi/spotinst)
* Python: [`pulumi-spotinst`](https://pypi.org/project/pulumi-spotinst/)
* Go: [`github.com/pulumi/pulumi-spotinst/sdk/v3/go/spotinst`](https://github.com/pulumi/pulumi-spotinst)
* .NET: [`Pulumi.Spotinst`](https://www.nuget.org/packages/Pulumi.Spotinst)

## Configuration

Once obtained, there are two ways to communicate your configuration tokens to Pulumi:

1. Set the environment variables `SPOTINST_ACCOUNT` and `SPOTINST_ACCOUNT`:

    ```bash
    $ export SPOTINST_ACCOUNT=XXXXXX
    $ export SPOTINST_TOKEN=YYYYYY
    ```

2. Set them using configuration, if you prefer that they be stored alongside your Pulumi stack for easy multi-user access:

    ```bash
    $ pulumi config set spotinst:account XXXXXX
    $ pulumi config set spotinst:token --secret
    ```

Remember to pass `--secret` when setting `token` so that it is properly encrypted.
