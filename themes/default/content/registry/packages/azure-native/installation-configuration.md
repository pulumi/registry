---
title: Azure Native Setup
meta_desc: How to set up credentials to use the Pulumi Azure Native Provider and choose configuration options to tailor the provider to suit your use case.
layout: installation
---

To provision resources with the Pulumi Azure Native provider, you need to have Azure credentials.
These instructions assume you're using the [Azure CLI 2.0](https://github.com/Azure/azure-cli).
Your Azure credentials are never sent to Pulumi.com.
Pulumi uses the Azure SDK and the credentials in your environment to authenticate requests from your computer to Azure.

## Installation

The Azure Native provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/azure-native`](https://www.npmjs.com/package/@pulumi/azure-native)
* Python: [`pulumi-azure-native`](https://pypi.org/project/pulumi-azure-native/)
* Go: [`github.com/pulumi/pulumi-azure-native/sdk/go/azure`](https://github.com/pulumi/pulumi-azure-native)
* .NET: [`Pulumi.AzureNative`](https://www.nuget.org/packages/Pulumi.AzureNative)
* Java: [`com.pulumi.azurenative`](https://search.maven.org/search?q=com.pulumi.azure-native)

## Credentials

Pulumi can authenticate to Azure via the Azure CLI or using a Service Principal.

If you're running the Pulumi CLI locally, in a developer scenario, we recommend using the Azure CLI.  For team
environments, particularly in CI, a Service Principal is recommended.

{{% notes type="info" %}}
Authenticating using the CLI will not work for Service Principal logins (e.g.,
`az login --service-principal`).  For such cases, authenticate using the Service Principal method instead.
{{% /notes %}}

### Authenticate using the CLI

Log in to the Azure CLI and Pulumi will automatically use your credentials:

```bash
$ az login
A web browser has been opened at https://login.microsoftonline.com/organizations/oauth2/v2.0/authorize. Please continue the login in the web browser. If no web browser is available or if the web browser fails to open, use device code flow with `az login --use-device-code`.
```

Do as instructed to log in.  After completed, `az login` will return and you are ready to go.

{{% notes type="info" %}}
If you're using Government or China Clouds, you'll need to configure the Azure CLI to work with that cloud.  Do so by running `az cloud set --name <Cloud>`, where `<Cloud>` is one of `AzureUSGovernment` or `AzureChinaCloud`.
{{% /notes %}}

The Azure CLI, and thus Pulumi, will use the Default Subscription by default. You can override the subscription by setting your subscription ID to the `id` output from `az account list`'s output:

```bash
$ az account list
```

Pick out the `<id>` from the list and run:

```bash
$ az account set --subscription=<id>
```

#### Authenticate using a Service Principal

A Service Principal is an application in Azure Active Directory with three authorization tokens: a client ID, a client secret, and a tenant ID. Using a Service Principal is the recommended way to connect Pulumi to Azure in a team or CI setting.

#### Create your Service Principal and get your tokens

To use a Service Principal, you must first create one. If you already have one, skip this section.

You can create a Service Principal [using the Azure CLI](https://docs.microsoft.com/en-us/cli/azure/create-an-azure-service-principal-azure-cli?view=azure-cli-latest), [using the Azure Cloud Shell](https://shell.azure.com/), or [using the Azure Portal](https://docs.microsoft.com/en-us/azure/azure-resource-manager/resource-group-create-service-principal-portal?view=azure-cli-latest).

After creating a Service Principal, you will obtain three important tokens:

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

##### Set configuration using `pulumi config`

Remember to pass `--secret` when setting `clientSecret` so that it is properly encrypted:

    ```bash
    $ pulumi config set azure-native:clientId <clientID>
    $ pulumi config set azure-native:clientSecret <clientSecret> --secret
    $ pulumi config set azure-native:tenantId <tenantID>
    $ pulumi config set azure-native:subscriptionId <subscriptionId>
    $ pulumi config set azure-native:location <locationName>
    ```

##### Set configuration using environment variables

{{< chooser os "linux,macos,windows" >}}
{{% choosable os linux %}}

```bash
$ export ARM_CLIENT_ID=<YOUR_ARM_CLIENT_ID>
$ export ARM_CLIENT_SECRET=<YOUR_ARM_CLIENT_SECRET>
$ export ARM_TENANT_ID=<YOUR_ARM_TENANT_ID>
$ export ARM_SUBSCRIPTION_ID=<YOUR_ARM_SUBSCRIPTION_ID>
$ export ARM_LOCATION_NAME=<YOUR_ARM_LOCATION_NAME>
```

{{% /choosable %}}

{{% choosable os macos %}}

```bash
$ export ARM_CLIENT_ID=<YOUR_ARM_CLIENT_ID>
$ export ARM_CLIENT_SECRET=<YOUR_ARM_CLIENT_SECRET>
$ export ARM_TENANT_ID=<YOUR_ARM_TENANT_ID>
$ export ARM_SUBSCRIPTION_ID=<YOUR_ARM_SUBSCRIPTION_ID>
$ export ARM_LOCATION_NAME=<YOUR_ARM_LOCATION_NAME>
```

{{% /choosable %}}

{{% choosable os windows %}}

```powershell
> $env:ARM_CLIENT_ID = "<YOUR_ARM_CLIENT_ID>"
> $env:ARM_CLIENT_SECRET = "<YOUR_ARM_CLIENT_SECRET>"
> $env:ARM_TENANT_ID = "<YOUR_ARM_TENANT_ID>"
> $env:ARM_SUBSCRIPTION_ID = "<YOUR_ARM_SUBSCRIPTION_ID>"
> $env:ARM_LOCATION_NAME = "<YOUR_ARM_LOCATION_NAME>"
```

{{% /choosable %}}
{{< /chooser >}}

## Configuration options

Use `pulumi config set azure-native:<option>` or pass options to the [constructor of `new azure-native.Provider`]({{< relref "/registry/packages/azure-native/api-docs/provider" >}}).

| Option                      | Required/Optional | Description                                                                                                                                                                                                                                                                                            |
|-----------------------------|-------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `environment`               | Required          | The cloud environment to use. It can also be sourced from the ARM_ENVIRONMENT environment variable. Supported values are: `public` (default), `usgovernment`, `china`.                                                                                                                                 |
| `location`                  | Optional          | The location to use. ResourceGroups will consult this property for a default location, if one was not supplied explicitly. This must be set in either the resource or the environment vars                                                                                                                                                                        |
| `clientId`                  | Optional          | The client ID to use. It can also be sourced from the `ARM_CLIENT_ID` environment variable.                                                                                                                                                                                                            |
| `clientSecret`              | Optional          | The client secret to use. It can also be sourced from the `ARM_CLIENT_SECRET` environment variable.                                                                                                                                                                                                    |
| `msiEndpoint`               | Optional          | The REST endpoint to retrieve an MSI token from. Pulumi will attempt to discover this automatically but it can be specified manually here. It can also be sourced from the `ARM_MSI_ENDPOINT` environment variable.                                                                                    |
| `skipCredentialsValidation` | Optional          | Prevents the provider from validating the given credentials. When set to true, `skip_provider_registration` is assumed. It can also be sourced from the `ARM_SKIP_CREDENTIALS_VALIDATION` environment variable; defaults to `false`.                                                                   |
| `skipProviderRegistration`  | Optional          | Prevents the provider from registering the ARM provider namespaces, this can be used if you don't wish to give the Active Directory Application permission to register resource providers. It can also be sourced from the `ARM_SKIP_PROVIDER_REGISTRATION` environment variable; defaults to `false`. |
| `subscriptionId`            | Optional          | The subscription ID to use. It can also be sourced from the `ARM_SUBSCRIPTION_ID` environment variable.                                                                                                                                                                                                |
| `tenantId`                  | Optional          | The tenant ID to use. It can also be sourced from the `ARM_TENANT_ID` environment variable.                                                                                                                                                                                                            |
| `useMsi`                    | Optional          | Set to true to authenticate using managed service identity. It can also be sourced from the `ARM_USE_MSI` environment variable.                                                                                                                                                                        |
