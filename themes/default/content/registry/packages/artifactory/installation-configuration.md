---
title: Artifactory Installation & Configuration
meta_desc: Provides an overview on how to configure credentials for the Pulumi Artifactory Provider.
layout: installation
---

The Pulumi Artifactory provider uses the Artifactory SDK to manage and provision resources.

## Installation

The Artifactory provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/artifactory`](https://www.npmjs.com/package/@pulumi/artifactory)
* Python: [`pulumi-artifactory`](https://pypi.org/project/pulumi-artifactory/)
* Go: [`github.com/pulumi/pulumi-artifactory/sdk/go/artifactory`](https://github.com/pulumi/pulumi-artifactory)
* .NET: [`Pulumi.Artifactory`](https://www.nuget.org/packages/Pulumi.Artifactory)

## Configuring Credentials

Pulumi relies on the Artifactory SDK to authenticate requests from your computer to Artifactory. Your credentials are never sent
to pulumi.com.
The Pulumi Artifactory Provider needs to be configured with Artifactory credentials
before it can be used to create resources.

Once the credentials are obtained, there are two ways to communicate your authorization tokens to Pulumi:

1. Set the environment variables `ARTIFACTORY_URL`, and `ARTIFACTORY_ACCESS_TOKEN` (`ARTIFACTORY_USERNAME` and `ARTIFACTORY_PASSWORD`):

    ```bash
    $ export ARTIFACTORY_URL=XXXXXXXXXXXXXX
    $ export ARTIFACTORY_ACCESS_TOKEN=YYYYYYYYYYYYYY
    ```

2. Set them using configuration, if you prefer that they be stored alongside your Pulumi stack for easy multi-user access:

    ```bash
    $ pulumi config set artifactory:url XXXXXXXXXXXXXX
    $ pulumi config set artifactory:accessToken YYYYYYYYYYYYYY --secret
    ```

Remember to pass `--secret` when setting `artifactory:accessToken` so that it is properly encrypted. The complete list of
configuration parameters is in the [Artifactory provider README](https://github.com/pulumi/pulumi-artifactory/blob/main/README.md).
