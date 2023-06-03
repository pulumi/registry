---
title: AzureDevOps Installation & Configuration
meta_desc: Provides an overview on how to configure credentials for the Pulumi AzureDevOps Provider.
layout: package
---

The Pulumi AzureDevOps provider uses the AzureDevOps SDK to manage and provision resources.

## Installation

The AzureDevOps provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/azuredevops`](https://www.npmjs.com/package/@pulumi/azuredevops)
* Python: [`pulumi-azuredevops`](https://pypi.org/project/pulumi-azuredevops/)
* Go: [`github.com/pulumi/pulumi-azuredevops/sdk/v2/go/azuredevops`](https://github.com/pulumi/pulumi-azuredevops)
* .NET: [`Pulumi.AzureDevOps`](https://www.nuget.org/packages/Pulumi.AzureDevOps)

## Configuring Credentials

Pulumi relies on the AzureDevOps SDK to authenticate requests from your computer to AzureDevOps. Your credentials are never sent to pulumi.com.
The Pulumi AzureDevOps Provider needs to be configured with AzureDevOps credentials
before it can be used to create resources.

Once the credentials are obtained, there are two ways to communicate your authorization tokens to Pulumi:

1. Set the environment variables `AZDO_ORG_SERVICE_URL` and `AZDO_PERSONAL_ACCESS_TOKEN`:

    ```bash
    $ export AZDO_ORG_SERVICE_URL=XXXXXXXXXXXXXX
    $ export AZDO_PERSONAL_ACCESS_TOKEN=YYYYYYYYYYYYYY
    ```

2. Set them using configuration, if you prefer that they be stored alongside your Pulumi stack for easy multi-user access:

    ```bash
    $ pulumi config set azuredevops:orgServiceUrl XXXXXXXXXXXXXX --secret
    $ pulumi config set azuredevops:personalAccessToken YYYYYYYYYYYYYY --secret
    ```

Remember to pass `--secret` when setting `azuredevops:personalAccessToken` so that it is properly encrypted. The complete list of
configuration parameters is in the [AzureDevOps provider README](https://github.com/pulumi/pulumi-azuredevops/blob/master/README.md).
