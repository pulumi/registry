---
title: AzureAD Installation & Configuration
meta_desc: Provides an overview on how to configure credentials for the Pulumi AzureAD Provider.
layout: package
---

The Pulumi AzureAD provider uses the AzureAD SDK to manage and provision resources.
To provision resources with the Pulumi AzureAD provider, you need to have Azure credentials.
Your Azure credentials are never sent to Pulumi.com.
Pulumi uses the Azure SDK and the credentials in your environment to authenticate requests from your computer to Azure.

## Installation

The AzureAD provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/azuread`](https://www.npmjs.com/package/@pulumi/azuread)
* Python: [`pulumi-azuread`](https://pypi.org/project/pulumi-azuread/)
* Go: [`github.com/pulumi/pulumi-azuread/sdk/v4/go/azuread`](https://github.com/pulumi/pulumi-azuread)
* .NET: [`Pulumi.AzureAD`](https://www.nuget.org/packages/Pulumi.AzureAD)
* Java: [`com.pulumi.azuread`](https://search.maven.org/search?q=com.pulumi.azuread)

## Credentials

Pulumi can authenticate to Azure via several methods:

* Azure CLI
* OpenID Connect (OIDC)
* Service Principal with a client secret or certificate
* Managed Service Identity (MSI)

If you're running the Pulumi CLI locally, in a developer scenario, we recommend using the Azure CLI.  For team
environments, particularly in Continuous Integration, one of the other options is strongly recommended.

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
az account list
```

Pick out the `<id>` from the list and run:

```bash
az account set --subscription=<id>
```

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

To use OIDC, either set the Pulumi configuration `useOidc` via `pulumi config set azuread:useOidc true` or set the
environment variable `ARM_USE_OIDC` to "true".

Next, supply the Pulumi provider with the ID token to exchange for an Azure token. This step depends on the service
(identity provider) your program will run on.
* On GitHub, you don't need to configure anything since
[GitHub sets the relevant environment variables](https://docs.github.com/en/actions/deployment/security-hardening-your-deployments/about-security-hardening-with-openid-connect) by default and the Pulumi provider reads them automatically.
* Other identity providers offer a way to access the ID token. For instance, in GitLab CI/CD jobs, the ID token is available
via the environment variable `GITLAB_OIDC_TOKEN`. Configure the Pulumi provider to use the ID token by assigning it to
the Pulumi configuration `azuread:oidcToken` or the environment variable `ARM_OIDC_TOKEN`.

{{% notes type="info" %}}
If your identity provider does not offer an ID token directly but it does offer a way to exchange a local bearer token
for an ID token, you can configure this exchange as well. This is a rare case that you won't need unless the identity
provider's documentation explicitly requests it. In that case, set one of the following pairs:

* both the `azuread:oidcRequestToken` and `azuread:oidcRequestUrl` Pulumi configuration values, **or**
* both the `ARM_OIDC_REQUEST_TOKEN` and `ARM_OIDC_REQUEST_URL` environment variables.
{{% /notes %}}

Finally, configure the client and tenant IDs of your Azure Active Directory application. Refer to the
[above Azure documentation](https://learn.microsoft.com/en-us/azure/active-directory/workload-identities/workload-identity-federation-create-trust?pivots=identity-wif-apps-methods-azp)
on how to retrieve the IDs, and set them via Pulumi config as `azuread:clientId` and `azuread:tenantId` or via environment
variables as `ARM_CLIENT_ID` and `ARM_TENANT_ID`.

{{% notes type="info" %}}
If you get the error "_AADSTS70021: No matching federated identity record found for presented assertion_", this points
to a configuration issue with the _entity type_ and _environment name_ described in the Azure documentation. Make sure
they match your setup, e.g., the type "branch" and the correct branch name if CI runs against a fixed branch.
{{% /notes %}}

### Authenticate using a Service Principal

A Service Principal is an application in Azure Active Directory with a client ID and a tenant ID, exactly like the one
used in the OIDC scenario. In this scenario, instead of a pre-configured trust relationship, a client secret is used to
authenticate with Azure.

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

#### Make tokens available to Pulumi

Once you have the Service Principal's authorization tokens, choose one of the ways below to make them available to Pulumi:

##### Set configuration using `pulumi config`

Remember to pass `--secret` when setting `clientSecret` so that it is properly encrypted:

    ```bash
    $ pulumi config set azuread:clientId <clientID>
    $ pulumi config set azuread:clientSecret <clientSecret> --secret
    $ pulumi config set azuread:tenantId <tenantID>
    ```

##### Set configuration using environment variables

{{< chooser os "linux,macos,windows" >}}
{{% choosable os linux %}}

```bash
export ARM_CLIENT_ID=<YOUR_ARM_CLIENT_ID>
export ARM_CLIENT_SECRET=<YOUR_ARM_CLIENT_SECRET>
export ARM_TENANT_ID=<YOUR_ARM_TENANT_ID>
```

{{% /choosable %}}

{{% choosable os macos %}}

```bash
export ARM_CLIENT_ID=<YOUR_ARM_CLIENT_ID>
export ARM_CLIENT_SECRET=<YOUR_ARM_CLIENT_SECRET>
export ARM_TENANT_ID=<YOUR_ARM_TENANT_ID>
```

{{% /choosable %}}

{{% choosable os windows %}}

```powershell
> $env:ARM_CLIENT_ID = "<YOUR_ARM_CLIENT_ID>"
> $env:ARM_CLIENT_SECRET = "<YOUR_ARM_CLIENT_SECRET>"
> $env:ARM_TENANT_ID = "<YOUR_ARM_TENANT_ID>"
```

{{% /choosable %}}
{{< /chooser >}}

## Configuration options

Use `pulumi config set azuread:<option>` or pass options to the [constructor of `new azuread.Provider`](/registry/packages/azuread/api-docs/provider).

| Option | Required? | Description |
| - | - | - |
| `environment` | Required | The cloud environment to use. It can also be sourced from the ARM_ENVIRONMENT environment variable. Supported values are: `public` (default), `usgovernment`, `china`. |
| `location` | Optional | The location to use. ResourceGroups will consult this property for a default location, if one was not supplied explicitly. |
| `clientId` | Optional | The client ID to use. It can also be sourced from the `ARM_CLIENT_ID` environment variable. |
| `clientSecret` | Optional | The client secret to use. It can also be sourced from the `ARM_CLIENT_SECRET` environment variable. |
| `msiEndpoint` | Optional | The REST endpoint to retrieve an MSI token from. Pulumi will attempt to discover this automatically but it can be specified manually here. It can also be sourced from the `ARM_MSI_ENDPOINT` environment variable. |
| `oidcRequestToken` | Optional | The bearer token to exchange for an OIDC ID token. It can also be sourced from the `ARM_OIDC_REQUEST_TOKEN` environment variable. Requires `oidcRequestUrl` and should not be combined with `oidcToken`. |
| `oidcRequestUrl` | Optional | The bearer token exchange URL for OIDC authentication. It can also be sourced from the `ARM_OIDC_REQUEST_URL` environment variable. Requires `oidcRequestToken` and should not be combined with `oidcToken`. |
| `oidcToken` | Optional | The ID token for OIDC authentication. It can also be sourced from the `ARM_OIDC_TOKEN` environment variable. |
| `skipCredentialsValidation` | Optional | Prevents the provider from validating the given credentials. When set to true, `skip_provider_registration` is assumed. It can also be sourced from the `ARM_SKIP_CREDENTIALS_VALIDATION` environment variable; defaults to `false`.
| `skipProviderRegistration` | Optional | Prevents the provider from registering the ARM provider namespaces, this can be used if you don't wish to give the Active Directory Application permission to register resource providers. It can also be sourced from the `ARM_SKIP_PROVIDER_REGISTRATION` environment variable; defaults to `false`. |
| `tenantId` | Optional | The tenant ID to use. It can also be sourced from the `ARM_TENANT_ID` environment variable. |
| `useMsi` | Optional | Set to true to authenticate using managed service identity. It can also be sourced from the `ARM_USE_MSI` environment variable.  |
| `useOidc` | Optional | Set to true to authenticate using OIDC. It can also be sourced from the `ARM_USE_OIDC` environment variable. |
