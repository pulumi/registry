---
title: Yandex Cloud Installation & Configuration
meta_desc: Provides an overview on how to configure credentials for the Pulumi Yandex Cloud Provider.
layout: package
---

The Pulumi Yandex Cloud provider uses the Yandex Cloud SDK to manage and provision resources.

## Installation

The Yandex provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/yandex`](https://www.npmjs.com/package/@pulumi/yandex)
* Python: [`pulumi-yandex`](https://pypi.org/project/pulumi-yandex/)
* Go: [`github.com/pulumi/pulumi-yandex/sdk/go/yandex`](https://github.com/pulumi/pulumi-yandex)
* .NET: [`Pulumi.Yandex`](https://www.nuget.org/packages/Pulumi.Yandex)

## Configuring The Provider

Pulumi relies on the Yandex Cloud SDK to authenticate requests from your computer to Yandex Cloud. Your credentials are never sent
to pulumi.com.
The Pulumi Yandex Cloud Provider needs to be configured with Yandex Cloud credentials
before it can be used to create resources.

Once the credentials are obtained, there are two ways to communicate your configuration to Pulumi:

1. Set the environment variables `YC_TOKEN` or `YC_SERVICE_ACCOUNT_KEY_FILE` and `YC_CLOUD_ID` and `YC_FOLDER_ID`:

    ```bash
    $ export YC_TOKEN=XXXXXX
    $ export YC_CLOUD_ID=YYYYYY
    $ export YC_FOLDER_ID=ZZZZZZ
    ```

2. Set them using configuration, if you prefer that they be stored alongside your Pulumi stack for easy multi-user access:

    ```bash
    $ pulumi config set yandex:token XXXXXX --secret
    $ pulumi config set yandex:cloudId YYYYYY
    $ pulumi config set yandex:folderId ZZZZZZ
    ```

The complete list of
configuration parameters is in the [Yandex Cloud provider README](https://github.com/pulumi/pulumi-yandex/blob/master/README.md).
Remember to pass `--secret` when setting `token` so that it is properly encrypted.
