---
title: Azure Native Installation & Configuration
meta_desc: How to set up credentials to use the Pulumi Azure Native Provider and choose configuration options to tailor the provider to suit your use case.
layout: package
---

To provision resources with the Pulumi Azure Native provider, you need to have Azure credentials.
Your Azure credentials are never sent to Pulumi.com.
Pulumi uses the Azure SDK and the credentials in your environment to authenticate requests from your computer to Azure.

## Installation

The Azure Native provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/azure-native`](https://www.npmjs.com/package/@pulumi/azure-native)
* Python: [`pulumi-azure-native`](https://pypi.org/project/pulumi-azure-native/)
* Go: [`github.com/pulumi/pulumi-azure-native-sdk`](https://github.com/pulumi/pulumi-azure-native-sdk)
* .NET: [`Pulumi.AzureNative`](https://www.nuget.org/packages/Pulumi.AzureNative)
* Java: [`com.pulumi.azurenative`](https://search.maven.org/search?q=com.pulumi.azure-native)


## Authentication Methods

Pulumi can authenticate to Azure via several methods:
- Azure CLI (the default)
- Default Azure Credential
- OpenID Connect (OIDC)
- Service Principal with a client secret or certificate
- Managed Service Identity (MSI)

If you're running the Pulumi CLI locally, in a developer scenario, we recommend using the Azure CLI.  For team
environments, particularly in Continuous Integration, we recommend using Azure Default Credential (`useDefaultAzureCredential`).

{{% notes type="info" %}}
Authenticating using the CLI will not work for Service Principal logins (e.g.,
`az login --service-principal`).  For such cases, authenticate using the Service Principal method instead.
{{% /notes %}}


### Authenticate using the CLI

The CLI instructions assume you're using the [Azure CLI](https://github.com/Azure/azure-cli) (`az`).

Log in to the Azure CLI and Pulumi will automatically use your credentials:

```bash
$ az login
A web browser has been opened at https://login.microsoftonline.com/organizations/oauth2/v2.0/authorize. Please continue the login in the web browser. If no web browser is available or if the web browser fails to open, use device code flow with `az login --use-device-code`.
```

Do as instructed to log in.  After completed, `az login` will return and you are ready to go.

{{% notes type="info" %}}
If you're using Government or China Clouds, you'll need to configure the Azure CLI to work with that cloud.  Do so by running `az cloud set --name <Cloud>`, where `<Cloud>` is one of `AzureUSGovernment` or `AzureChinaCloud`.
{{% /notes %}}

The Azure CLI, and thus Pulumi, will use the current subscription by default. You can activate a different subscription by setting your subscription ID to the `id` output from `az account list`'s output:

```bash
$ az account list
```

Pick out the `<id>` from the list and run:

```bash
$ az account set --subscription=<id>
```

#### Authenticate with Default Azure Credential

The "Default Azure Credential" mode is a preconfigured chain of credentials. It's designed to support many environments, 
along with the most common authentication flows and developer tools. Use this mode to automatically detect credentials
available in Azure hosting environments and to automatically use the Azure CLI during local development.

See ["DefaultAzureCredential Overview"](https://learn.microsoft.com/en-us/azure/developer/go/sdk/authentication/credential-chains#defaultazurecredential-overview)
for a complete description of the types and ordering of credentials.

You also need to obtain a Subscription ID. To retrieve your current Subscription ID, you can use:

```bash
$ az account show --query id -o tsv
```

Once you have your Subscription ID, configure the provider in one of two ways.

##### Set configuration using `pulumi config`

    ```bash
    $ pulumi config set azure-native:useDefaultAzureCredential true
    $ pulumi config set azure-native:subscriptionId <subscriptionId>
    ```

##### Set configuration using environment variables

{{< chooser os "linux,macos,windows" >}}
{{% choosable os linux %}}

```bash
$ export ARM_USE_DEFAULT_AZURE_CREDENTIAL=true
$ export ARM_SUBSCRIPTION_ID=<YOUR_ARM_SUBSCRIPTION_ID>
```

{{% /choosable %}}

{{% choosable os macos %}}

```bash
$ export ARM_USE_DEFAULT_AZURE_CREDENTIAL=true
$ export ARM_SUBSCRIPTION_ID=<YOUR_ARM_SUBSCRIPTION_ID>
```

{{% /choosable %}}

{{% choosable os windows %}}

```powershell
> $env:ARM_USE_DEFAULT_AZURE_CREDENTIAL = "true"
> $env:ARM_SUBSCRIPTION_ID = "<YOUR_ARM_SUBSCRIPTION_ID>"
```

{{% /choosable %}}
{{< /chooser >}}

##### Use a specific credential
To exclude all credentials except for one, set environment variable `AZURE_TOKEN_CREDENTIALS` to one of:

- `AzureCLICredential`
- `AzureDeveloperCLICredential`
- `EnvironmentCredential`
- `ManagedIdentityCredential`
- `WorkloadIdentityCredential`

##### Exclude a credential type category

To exclude all "developer tool" or "deployed service" credentials, set environment variable `AZURE_TOKEN_CREDENTIALS` to `prod` or `dev`, respectively.

### Authenticate with OpenID Connect (OIDC)

OIDC allows you to establish a trust relationship between Azure and another identity provider such as GitHub or Azure DevOps. Once
established, your program can exchange an ID token issued by the identity provider for an Azure token. Your Pulumi program running in
the identity provider's service, for instance, GitHub Actions CI or Azure DevOps Pipelines, can then access Azure, without storing any
secrets in GitHub.

#### OIDC Azure Configuration

To configure the trust relationship in Azure, please refer to
[this guide](https://learn.microsoft.com/en-us/azure/active-directory/workload-identities/workload-identity-federation-create-trust?pivots=identity-wif-apps-methods-azp#github-actions).
This needs to be set up only once.

#### OIDC Pulumi Provider Configuration

To use OIDC, either set the Pulumi configuration `useOidc` via `pulumi config set azure-native:useOidc true` or set the
environment variable `ARM_USE_OIDC` to "true".

Next, supply the Pulumi provider with the ID token to exchange for an Azure token. This step depends on the service
(identity provider) your program will run on.
- On GitHub, you don't need to configure anything since
[GitHub sets the relevant environment variables](https://docs.github.com/en/actions/deployment/security-hardening-your-deployments/about-security-hardening-with-openid-connect) by default and the Pulumi provider reads them automatically.
- Other identity providers offer a way to access the ID token. For instance, in GitLab CI/CD jobs, the ID token is available
via the environment variable `GITLAB_OIDC_TOKEN`. Configure the Pulumi provider to use the ID token by assigning it to
the Pulumi configuration `azure-native:oidcToken` or the environment variable `ARM_OIDC_TOKEN`.
- Some identity providers, such as Azure Workload Identity, supply the ID token in a file. In this case, set the
Pulumi configuration `azure-native:oidcTokenFilePath` or the environment variable `ARM_OIDC_TOKEN_FILE_PATH` to the path
of the file.

{{% notes type="info" %}}
If your identity provider does not offer an ID token directly but it does offer a way to exchange a local bearer token
for an ID token, you can configure this exchange as well. This is a rare case that you won't need unless the identity
provider's documentation explicitly requests it. In that case, set one of the following pairs:
  - both the `azure-native:oidcRequestToken` and `azure-native:oidcRequestUrl` Pulumi configuration values, **or**
  - both the `ARM_OIDC_REQUEST_TOKEN` and `ARM_OIDC_REQUEST_URL` environment variables.
{{% /notes %}}

Finally, configure the client and tenant IDs of your Azure Active Directory application. Refer to the
[above Azure documentation](https://learn.microsoft.com/en-us/azure/active-directory/workload-identities/workload-identity-federation-create-trust?pivots=identity-wif-apps-methods-azp)
on how to retrieve the IDs, and set them via Pulumi config as `azure-native:clientId` and `azure-native:tenantId` or via environment
variables as `ARM_CLIENT_ID` and `ARM_TENANT_ID`.

{{% notes type="info" %}}
If you get the error "_AADSTS70021: No matching federated identity record found for presented assertion_", this points
to a configuration issue with the _entity type_ and _environment name_ described in the Azure documentation. Make sure
they match your setup, e.g., the type "branch" and the correct branch name if CI runs against a fixed branch.
{{% /notes %}}

#### OIDC Dynamic Credentials with Pulumi ESC

In addition to configuring the Azure Native provider locally, you also have the option to centralize your configurations using [Pulumi ESC (Environments, Secrets, and Configuration)](/docs/pulumi-cloud/esc/). Using this service will enable you to run Pulumi CLI commands with dynamically generated credentials, removing the need to configure and manage your credentials locally.

To do this, you will need to complete the following steps:

##### Configure OIDC between Pulumi and Azure

Refer to the [Configuring OpenID Connect for Azure Guide](/docs/pulumi-cloud/oidc/azure/) for the step-by-step process on how to do this.

##### [Optional] Move Pulumi config to your ESC environment

With Pulumi ESC, you can define and expose environment variables as shown below:

```yaml
values:
  azure:
    login:
      fn::open::azure-login:
        clientId: <your-client-id>
        tenantId: <your-tenant-id>
        subscriptionId: <your-subscription-id>
        oidc: true
  environmentVariables:
    ARM_USE_OIDC: 'true'
    ARM_CLIENT_ID: ${azure.login.clientId}
    ARM_TENANT_ID: ${azure.login.tenantId}
    ARM_OIDC_TOKEN: ${azure.login.oidc.token}
    ARM_SUBSCRIPTION_ID: ${azure.login.subscriptionId}
```

{{< notes type="info" >}}
To learn more about projecting environment variables in Pulumi ESC, refer to the [relevant Pulumi ESC documentation](/docs/pulumi-cloud/esc/environments/#projecting-environment-variables).
{{< /notes >}}

To [expose configuration values to Pulumi IaC](/docs/pulumi-cloud/esc/environments/#using-environments-with-pulumi-iac), you will need to add a second-level key named `pulumiConfig` and define your desired values underneath it. Further, if your workflow does not require the exposure of environment variables, you can also define those variables under the `pulumiConfig` block as shown below:

```yaml
values:
  azure:
    login:
      fn::open::azure-login:
        clientId: <your-client-id>
        tenantId: <your-tenant-id>
        subscriptionId: <your-subscription-id>
        oidc: true
  pulumiConfig:
    azure-native:useOidc: 'true'
    azure:useOidc: 'true'
    azure-native:environment: <your-environment>
    azure-native:clientId: ${azure.login.clientId}
    azure-native:tenantId: ${azure.login.tenantId}
    azure-native:subscriptionId: ${azure.login.subscriptionId}
    azure-native:oidcToken: ${azure.login.oidc.token}
```

This will ensure that those values are scoped only to your `pulumi` run.

{{< notes type="info" >}}
The configuration values under `pulumiConfig` can also be referenced directly from within your Pulumi program code. This is done using the same method to reference values from your project's stack settings file. You can see examples of how to do this in the [Accessing Configuration from Code](/docs/concepts/config/#code) section of the Pulumi documentation.
{{< /notes >}}

##### Import your environment

The last step is to update your project's stack settings file (`Pulumi.<stack-name>.yaml`) to import your ESC environment as shown below:

```yaml
environment:
  - <your-environment-name>
```

Make sure to replace `<your-environment-name>` with the name of the ESC environment you created in the previous steps.

You can test that your configuration is working by running the `pulumi preview` command. This will validate that your Azure resources can be deployed using the dynamically generated credentials in your environment file.

{{< notes type="info" >}}
Make sure that your local environment does not have Azure credentials configured before running this command. You can logout by running the `az logout` command.
{{< /notes >}}

### Authenticate using a Service Principal

A Service Principal is an application in Azure Active Directory with a client ID and a tenant ID, exactly like the one
used in the OIDC scenario. In this scenario, instead of a pre-configured trust relationship, a client secret is used to
authenticate with Azure. 

#### Create your Service Principal and get your tokens

To use a Service Principal, you must first create one. If you already have one, skip this section.

You can create a Service Principal [using the Azure CLI](https://docs.microsoft.com/en-us/cli/azure/create-an-azure-service-principal-azure-cli?view=azure-cli-latest), [using the Azure Cloud Shell](https://shell.azure.com/), or [using the Azure Portal](https://docs.microsoft.com/en-us/azure/azure-resource-manager/resource-group-create-service-principal-portal?view=azure-cli-latest).

After creating a Service Principal, you will obtain three important tokens:

* `appId` is the client ID
* `tenant` is the tenant ID
* `password` is the client secret

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
    # optional default location, otherwise set in code
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

Alternatively, you can use a certificate instead of a password as client secret. In this case, configure the provider
as shown above except use `clientCertPath`/`ARM_CLIENT_CERTIFICATE_PATH` instead of `clientSecret`/`ARM_CLIENT_SECRET`.


### Authenticate using Managed Service Identity (MSI)

MSI is [configured in Azure](https://learn.microsoft.com/en-us/azure/active-directory/managed-identities-azure-resources/overview). Once that's done, tell the provider to use it by setting the `useMsi` config or the `ARM_USE_MSI` environment variable to `true`.

When using a system-assigned identity, or when using a user-assigned identity _and there's only a single one_, you don't need to do anything else. The provider will automatically use the identity.

When one or more resources in your program have multiple user-assigned identities, you need to set the `clientId` config or the `ARM_CLIENT_ID` environment variable to the client ID of the identity you want to use.

## Connecting to a different Azure environment

The provider connects to the Azure public cloud by default. To connect to another cloud,
set it via Pulumi config as `azure-native:environment` or via environment
variable as `ARM_ENVIRONMENT`.  Use one of the following values:

- `AzureCloud`
- `AzureUSGovernment`
- `AzureChinaCloud`

## Configuration options

Use `pulumi config set azure-native:<option>` or pass options to the [constructor of `new azure-native.Provider`](/registry/packages/azure-native/api-docs/provider).

All configuration parameters are optional.

| Option | Description |
| - | - |
| `auxiliaryTenantIds` | Any additional Tenant IDs which should be used for authentication. It can also be sourced from the `ARM_AUXILIARY_TENANT_IDS` environment variable. |
| `clientCertificatePassword` | The password associated with the Client Certificate. For use when authenticating as a Service Principal using a Client Certificate. It can also be sourced from the `ARM_CLIENT_CERTIFICATE_PASSWORD` environment variable. |
| `clientCertificatePath` | The path to a certificate to use as client secret for Service Principal authentication. It can also be sourced from the `ARM_CLIENT_CERTIFICATE_PATH` environment variable. |
| `clientId` | The client ID to use for OIDC, Service Principal, or user-assigned identity authentication. It can also be sourced from the `ARM_CLIENT_ID` environment variable. |
| `clientSecret` | The client secret to use for Service Principal authentication. It can also be sourced from the `ARM_CLIENT_SECRET` environment variable. |
| `disablePulumiPartnerId` | This will disable the Pulumi Partner ID which is used if a custom `partnerId` isn't specified. It can also be sourced from the `ARM_DISABLE_PULUMI_PARTNER_ID` environment variable. |
| `environment` | The cloud environment to use. It can also be sourced from the ARM_ENVIRONMENT environment variable. Supported values are (case-insensitive): `public` or `AzureCloud` (default), `usgovernment` or `AzureUSGovernment`, `china` or `AzureChinaCloud`. |
| `location` | The location to use. ResourceGroups will consult this property for a default location, if one was not supplied explicitly when defining the resource. |
| `metadataHost` | The REST endpoint for the Azure Instance Metadata Service. Pulumi will attempt to discover this automatically but it can be specified manually here. It can also be sourced from the `ARM_METADATA_HOSTNAME` environment variable. |
| `msiEndpoint` | The REST endpoint to retrieve an MSI token from. Pulumi will attempt to discover this automatically but it can be specified manually here. It can also be sourced from the `ARM_MSI_ENDPOINT` environment variable. |
| `oidcAudience` | The audience to use for OIDC authentication. Typically not needed; defaults to the standard `api://AzureADTokenExchange`. It can also be sourced from the `ARM_OIDC_AUDIENCE` environment variable. |
| `oidcRequestToken` | Your cloud service or provider’s bearer token to exchange for an OIDC ID token. It can also be sourced from the `ARM_OIDC_REQUEST_TOKEN` environment variable. |
| `oidcRequestUrl` | The token exchange URL for OIDC authentication. It can also be sourced from the `ARM_OIDC_REQUEST_URL` environment variable. |
| `oidcToken` | The token to exchange for OIDC authentication. It can also be sourced from the `ARM_OIDC_TOKEN` environment variable. |
| `oidcTokenFilePath` | The path to a file containing a token to exchange for OIDC authentication. It can also be sourced from the `ARM_OIDC_TOKEN_FILE_PATH` environment variable. |
| `partnerId` | A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution. It can also be sourced from the `ARM_PARTNER_ID` environment variable. |
| `subscriptionId` | The subscription ID to use. It can also be sourced from the `ARM_SUBSCRIPTION_ID` environment variable. |
| `tenantId` | The tenant ID to use for OIDC or Service Principal authentication. It can also be sourced from the `ARM_TENANT_ID` environment variable. |
| `useDefaultAzureCredential` | Set to true to authenticate using Default Azure Credential. It can also be sourced from the `ARM_USE_DEFAULT_AZURE_CREDENTIAL` environment variable. |
| `useMsi` | Set to true to authenticate using managed service identity. It can also be sourced from the `ARM_USE_MSI` environment variable. |
| `useOidc` | Set to true to authenticate using OIDC. It can also be sourced from the `ARM_USE_OIDC` environment variable. |
