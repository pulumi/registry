---
title: MinIO Installation & Configuration
meta_desc: Provides an overview on how to configure credentials for the Pulumi MinIO Provider.
layout: installation
---

The Pulumi MinIO provider uses the MinIO SDK to manage and provision resources.

## Installation

The MinIO provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/minio`](https://www.npmjs.com/package/@pulumi/minio)
* Python: [`pulumi-minio`](https://pypi.org/project/pulumi-minio/)
* Go: [`github.com/pulumi/pulumi-minio/sdk/go/minio`](https://github.com/pulumi/pulumi-minio)
* .NET: [`Pulumi.Minio`](https://www.nuget.org/packages/Pulumi.Minio)

## Configuring Credentials

Pulumi relies on the MinIO SDK to authenticate requests from your computer to MinIO. Your credentials are never sent
to pulumi.com.
The Pulumi MinIO Provider needs to be configured with MinIO credentials
before it can be used to create resources.

Once the credentials are obtained, there are two ways to communicate your authorization tokens to Pulumi:

1. Set the environment variables `MINIO_ENDPOINT`, `MINIO_ACCESS_KEY` and `MINIO_SECRET_KEY`:

    ```bash
    $ export MINIO_ENDPOINT=XXXXXXXXXXXXXX
    $ export MINIO_ACCESS_KEY=YYYYYYYYYYYYYY
    $ export MINIO_SECRET_KEY=ZZZZZZZZZZZZZZ
    ```

2. Set them using configuration, if you prefer that they be stored alongside your Pulumi stack for easy multi-user access:

    ```bash
    $ pulumi config set minio:minioServer XXXXXXXXXXXXXX
    $ pulumi config set minio:minioAccessKey YYYYYYYYYYYYYY --secret
    $ pulumi config set minio:minioSecretKey ZZZZZZZZZZZZZZ --secret
    ```

Remember to pass `--secret` when setting `minio:minioAccessKey` and `minio:minioSecretKey` so that they are properly encrypted. The complete list of
configuration parameters is in the [MinIO provider README](https://github.com/pulumi/pulumi-minio/blob/master/README.md).
