---
title: Azure Native Setup
meta_desc: How to set up credentials to use the Pulumi Azure Native Provider and choose configuration options to tailor the provider to suit your use case.
layout: installation
---

{{% notes type="info" %}}
Pulumi Azure also works with the legacy [Azure xPlat CLI](https://github.com/Azure/azure-xplat-cli) but we recommend using the CLI 2.0 for the best experience.
{{% /notes %}}

## Installation

The Azure Native provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/azure-native`](https://www.npmjs.com/package/@pulumi/azure-native)
* Python: [`pulumi-azure-native`](https://pypi.org/project/pulumi-azure-native/)
* Go: [`github.com/pulumi/pulumi-azure-native/sdk/go/azure`](https://github.com/pulumi/pulumi-azure-native)
* .NET: [`Pulumi.AzureNative`](https://www.nuget.org/packages/Pulumi.AzureNative)
* Java: [`com.pulumi.azurenative`](https://search.maven.org/search?q=com.pulumi.azurenative)

## Setup

To provision resources with the Pulumi Azure provider, you need to have an Azure CLI installed and you need to have Azure credentials. These instructions assume you're using the [Azure CLI 2.0](https://github.com/Azure/azure-cli). Your Azure credentials are never sent to Pulumi.com. Pulumi uses the Azure SDK and the credentials in your environment to authenticate requests from your computer to Azure.

Pulumi can authenticate to Azure using a Service Principal or the Azure CLI.

If you're running the Pulumi CLI locally, in a developer scenario, we recommend using the Azure CLI. For team environments, particularly in CI, we recommend using a Service Principal.

> **Note:** Authenticating using the CLI will not work for Service Principal logins (e.g.,
> `az login --service-principal`).  For such cases, authenticate using the Service Principal method instead.

### Option 1: Use the CLI

Simply login to the Azure CLI and Pulumi will automatically use your credentials:

```bash
$ az login
The default web browser has been opened at https://login.microsoftonline.com/common/oauth2/authorize. Please continue the login in the web browser. If no web browser is available or if the web browser fails to open, use device code flow with `az login --use-device-code`.
```

Do as instructed to login.  After completed, `az login` will return and you are ready to go.

> **Note:** If you're using Government, China, or German Clouds, you'll need to configure the Azure CLI to work with that cloud.  Do so by running `az cloud set --name <Cloud>`, where `<Cloud>` is one of `AzureUSGovernment`, `AzureChinaCloud`, or `AzureGermanCloud`.

The Azure CLI, and thus Pulumi, will use the Default Subscription by default. You can override the subscription by setting your subscription ID to the `id` output from `az account list`'s output:

```bash
$ az account list
```

Pick out the `<id>` from the list and run:

```bash
$ az account set --subscription=<id>
```

### Option 2: Use a Service Principal

A Service Principal is an application in Azure Active Directory with three authorization tokens: a client ID, a client secret, and a tenant ID. Using a Service Principal is the recommended way to connect Pulumi to Azure in a team or CI setting.

#### Create your Service Principal and get your tokens

To use a Service Principal, you must first create one. If you already have one, skip this section.

You can create a Service Principal [using the Azure CLI](https://docs.microsoft.com/en-us/cli/azure/create-an-azure-service-principal-azure-cli?view=azure-cli-latest), [using the Azure Cloud Shell](https://shell.azure.com/), or [using the Azure Portal](https://docs.microsoft.com/en-us/azure/azure-resource-manager/resource-group-create-service-principal-portal?view=azure-cli-latest).

After creating a Service Principal, you will obtain three important tokens, mapping to the three shown earlier:

* `appId` is the client ID
* `password` is the client secret
* `tenant` is the tenant ID

For example, a common Service Principal as displayed by the Azure CLI looks something like this:

```json
{
  "appId": "WWWWWWWW-WWWW-WWWW-WWWW-WWWWWWWWWWWW",
  "displayName": "ServicePrincipalName",
  "name": "http://ServicePrincipalName",
  "password": "XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX",
  "tenant": "YYYYYYYY-YYYY-YYYY-YYYY-YYYYYYYYYYYY"
}
```

You also need to obtain a Subscription ID. To retrieve your current Subscription ID, you can use:

```bash
$ az account show --query id -o tsv
```

To list all available subscriptions, you can use:

```bash
$ az account list --query '[].{subscriptionName:name,subscriptionId:id}' -o tsv
```

#### Make tokens available to Pulumi

Once you have the Service Principal's authorization tokens, choose one of the ways below to make them available to Pulumi:

1. Set them using configuration (and remember to pass `--secret` when setting `clientSecret` so that it is properly encrypted):

    ```bash
    $ pulumi config set azure-native:clientId <clientID>
    $ pulumi config set azure-native:clientSecret <clientSecret> --secret
    $ pulumi config set azure-native:tenantId <tenantID>
    $ pulumi config set azure-native:subscriptionId <subscriptionId>
    ```

1. Set the environment variables `ARM_CLIENT_ID`, `ARM_CLIENT_SECRET`, `ARM_TENANT_ID`, and `ARM_SUBSCRIPTION_ID`

## Configuration options

Use `pulumi config set azure-native:<option>` or pass options to the [constructor of `new azure-native.Provider`]({{< relref "/registry/packages/azure-native/api-docs/provider" >}}).

| Option                      | Required/Optional | Description                                                                                                                                                                                                                                         |
|-----------------------------|-------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `auxiliaryTenantIds`        | Optional          | It can also be sourced from the following environment variable: ARM_AUXILIARY_TENANT_IDS                                                                                                                                                            |
| `clientCertificatePassword` | Optional          | The password associated with the Client Certificate. For use when authenticating as a Service Principal using a Client Certificate It can also be sourced from the following environment variable: ARM_CLIENT_CERTIFICATE_PASSWORD                  |
| `clientCertificatePath`     | Optional          | The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service Principal using a Client Certificate. It can also be sourced from the following environment variable: ARM_CLIENT_CERTIFICATE_PATH |
| `clientId`                  | Optional          | The Client ID which should be used. It can also be sourced from the following environment variable: ARM_CLIENT_ID                                                                                                                                   |
| `clientSecret`              | Optional          | The Client Secret which should be used. For use When authenticating as a Service Principal using a Client Secret. It can also be sourced from the following environment variable: ARM_CLIENT_SECRET                                                 |
| `disablePulumiPartnerId`    | Optional          | This will disable the Pulumi Partner ID which is used if a custom partnerId isn’t specified. It can also be sourced from the following environment variable: ARM_DISABLE_PULUMI_PARTNER_ID                                                          |
| `environment`               | Optional          | The Cloud Environment which should be used. Possible values are public, usgovernment, german, and china. Defaults to public. It can also be sourced from the following environment variable: ARM_ENVIRONMENT                                        |
| `msiEndpoint`               | Optional          | The path to a custom endpoint for Managed Service Identity - in most circumstances this should be detected automatically. It can also be sourced from the following environment variable: ARM_MSI_ENDPOINT                                          |
| `partnerId`                 | Optional          | A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution. It can also be sourced from the following environment variable: ARM_PARTNER_ID                                                                      |
| `subscriptionId`            | Optional          | The Subscription ID which should be used. It can also be sourced from the following environment variable: ARM_SUBSCRIPTION_ID                                                                                                                       |
| `tenantId`                  | Optional          | The Tenant ID which should be used. It can also be sourced from the following environment variable: ARM_TENANT_ID                                                                                                                                   |
| `useMsi`                    | Optional          | Allowed Managed Service Identity be used for Authentication. It can also be sourced from the following environment variable: ARM_USE_MSI                                                                                                            |
