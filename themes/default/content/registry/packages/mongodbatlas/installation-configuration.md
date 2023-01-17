---
title: MongoDB Atlas Installation & Configuration
meta_desc: Provides an overview on how to configure credentials for the Pulumi MongoDB Atlas Provider.
layout: installation
---

The Pulumi MongoDB Atlas provider uses the MongoDB Atlas SDK to manage and provision resources.

## Installation

The MongoDB Atlas provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/mongodbatlas`](https://www.npmjs.com/package/@pulumi/mongodbatlas)
* Python: [`pulumi-mongodbatlas`](https://pypi.org/project/pulumi-mongodbatlas/)
* Go: [`github.com/pulumi/pulumi-mongodbatlas/sdk/v3/go/mongodbatlas`](https://github.com/pulumi/pulumi-mongodbatlas)
* .NET: [`Pulumi.Mongodbatlas`](https://www.nuget.org/packages/Pulumi.Mongodbatlas)

## Configuring Credentials

Pulumi relies on the MongoDB Atlas SDK to authenticate requests from your computer to MongoDB Atlas. Your credentials are never sent
to pulumi.com. The Pulumi MongoDB Atlas Provider needs to be configured with MongoDB Atlas credentials
before it can be used to manage resources.

In order to communicate your configuration details to Pulumi:

1. Set the environment variables `MONGODB_ATLAS_PUBLIC_KEY` and `MONGODB_ATLAS_PRIVATE_KEY`:

    ```bash
    $ export MONGODB_ATLAS_PUBLIC_KEY=XXXXXXXXXXXXXX
    $ export MONGODB_ATLAS_PRIVATE_KEY=YYYYYYYYYYYYYY
    ```

1. Set them using configuration, if you prefer that they be stored alongside your Pulumi stack for easy multi-user access:

    ```bash
    $ pulumi config set mongodbatlas:publicKey XXXXXXXXXXXXXX --secret
    $ pulumi config set mongodbatlas:privateKey YYYYYYYYYYYYYY --secret
    ```

If you are going to set `mongodbatlas:privateKey` and `mongodbatlas:publicKey`, please remember to pass `--secret` so that they is properly encrypted. The complete list of
configuration parameters is in the [MongoDB Atlas provider README](https://github.com/pulumi/pulumi-mongodbatlas/blob/master/README.md).
