---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumiverse/pulumi-mssql/v0.1.0/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
title: MSSQL Installation & Configuration
meta_desc: Information on how to install the MSSQL provider.
layout: package
---

## Installation

The Pulumi MSSQL provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumiverse/mssql`](https://www.npmjs.com/package/@pulumiverse/mssql)
* Python: [`pulumiverse_mssql`](https://pypi.org/project/pulumiverse_mssql/)
* Go: [`github.com/pulumiverse/pulumi-mssql/sdk/go/mssql`](https://pkg.go.dev/github.com/pulumiverse/pulumi-mssql/sdk/go/mssql)
* .NET: [`Pulumiverse.Mssql`](https://www.nuget.org/packages/Pulumiverse.Mssql)

### Authentication Methods

#### SQL auth

Provider username and password, using `sql_auth` field:

{{< chooser language "typescript,python,go,csharp,yaml" >}}

{{% choosable language typescript %}}
```typescript
const providerMssql = new mssql.Provider("provider-mssql", {
    hostname: "localhost",
    port: 1433,
    sqlAuth: {
        username: "sa",
        password: "sa_password",
    },
});
```
{{% /choosable %}}

{{% choosable language python %}}
```python
provider_mssql = mssql.Provider("provider-mssql",
    hostname="localhost",
    port=1433,
    sql_auth=mssql.ProviderSqlAuthArgs(
        username="sa",
        password="sa_password",
    ))
```
{{% /choosable %}}

{{% choosable language go %}}
```go
_, err := mssql.NewProvider(ctx, "provider-mssql", &mssql.ProviderArgs{
  Hostname: pulumi.String("localhost"),
  Port:     pulumi.Int(1433),
  SqlAuth: &mssql.ProviderSqlAuthArgs{
    Username: pulumi.String("sa"),
    Password: pulumi.String("sa_password"),
  },
})
if err != nil {
  return err
}
```
{{% /choosable %}}

{{% choosable language csharp %}}
```csharp
var providerMssql = new Mssql.Provider("provider-mssql", new()
{
    Hostname = "localhost",
    Port = 1433,
    SqlAuth = new Mssql.Inputs.ProviderSqlAuthArgs
    {
        Username = "sa",
        Password = "sa_password",
    },
});
```
{{% /choosable %}}

{{% choosable language yaml %}}
```yaml
resources:
  provider-mssql:
    type: pulumi:providers:mssql
    properties:
      hostname: "localhost"
      port: 1433
      sqlAuth:
        username: "sa"
        password: "sa_password"
```
{{% /choosable %}}

{{< /chooser >}}

#### Azure AD
In case of Azure SQL connections, Azure Active Directory auth tokens can be used to authenticate.

##### Service Principal

Provide client ID (a.k.a. application ID), secret and tenant ID, using `azure_auth` field:

{{< chooser language "typescript,python,go,csharp,yaml" >}}

{{% choosable language typescript %}}
```typescript
const providerMssql = new mssql.Provider("provider-mssql", {
    hostname: "example.database.windows.net",
    port: 1433,
    azureAuth: {
        clientId: "94e8d55d-cbbc-4e41-b21a-8923d83f9a85",
        clientSecret: "client_secret",
        tenantId: "a352c914-bfd9-4b7e-8b1d-554a58353f22",
    },
});
```
{{% /choosable %}}

{{% choosable language python %}}
```python
provider_mssql = mssql.Provider("provider-mssql",
    hostname="example.database.windows.net",
    port=1433,
    azure_auth=mssql.ProviderAzureAuthArgs(
        client_id="94e8d55d-cbbc-4e41-b21a-8923d83f9a85",
        client_secret="client_secret",
        tenant_id="a352c914-bfd9-4b7e-8b1d-554a58353f22",
    ))
```
{{% /choosable %}}

{{% choosable language go %}}
```go
_, err := mssql.NewProvider(ctx, "provider-mssql", &mssql.ProviderArgs{
	Hostname: pulumi.String("example.database.windows.net"),
	Port:     pulumi.Int(1433),
	AzureAuth: &mssql.ProviderAzureAuthArgs{
		ClientId:     pulumi.String("94e8d55d-cbbc-4e41-b21a-8923d83f9a85"),
		ClientSecret: pulumi.String("client_secret"),
		TenantId:     pulumi.String("a352c914-bfd9-4b7e-8b1d-554a58353f22"),
	},
})
if err != nil {
	return err
}
return nil
```
{{% /choosable %}}

{{% choosable language csharp %}}
```csharp
var providerMssql = new Mssql.Provider("provider-mssql", new()
{
    Hostname = "example.database.windows.net",
    Port = 1433,
    AzureAuth = new Mssql.Inputs.ProviderAzureAuthArgs
    {
        ClientId = "94e8d55d-cbbc-4e41-b21a-8923d83f9a85",
        ClientSecret = "client_secret",
        TenantId = "a352c914-bfd9-4b7e-8b1d-554a58353f22",
    },
});
```
{{% /choosable %}}

{{% choosable language yaml %}}
```yaml
resources:
  provider-mssql:
    type: pulumi:providers:mssql
    properties:
      hostname: "example.database.windows.net"
      port: 1433
      azureAuth:
        clientId: "94e8d55d-cbbc-4e41-b21a-8923d83f9a85"
        clientSecret: "client_secret"
        tenantId: "a352c914-bfd9-4b7e-8b1d-554a58353f22"
```
{{% /choosable %}}

{{< /chooser >}}

##### Default chained credentials

When `azure_auth` value is set to empty object (`{}`), the provider uses chained credentials built from `EnvironmentCredential` -> `ManagedIdentityCredential` -> `AzureCLICredential`.
See [DefaultAzureCredential docs](https://github.com/Azure/azure-sdk-for-go/wiki/Set-up-Your-Environment-for-Authentication#configure-defaultazurecredential) for details.

##### Environment variables

When `azure_auth` value is set to empty object (`{}`) and following environment variables are set, the env variable values will be used for authentication, taking precedence over `DefaultAzureCredential`.

- `ARM_CLIENT_ID` or `AZURE_CLIENT_ID` 
- `ARM_CLIENT_SECRET`or `AZURE_CLIENT_SECRET` 
- `ARM_TENANT_ID`or `AZURE_TENANT_ID` 

{{< chooser language "typescript,python,go,csharp,yaml" >}}

{{% choosable language typescript %}}
```typescript
const providerMssql = new mssql.Provider("provider-mssql", {
    hostname: "example.database.windows.net",
    port: 1433,
    azureAuth: {},
});
```
{{% /choosable %}}

{{% choosable language python %}}
```python
provider_mssql = mssql.Provider("provider-mssql",
    hostname="example.database.windows.net",
    port=1433,
    azure_auth=mssql.ProviderAzureAuthArgs())
```
{{% /choosable %}}

{{% choosable language go %}}
```go
_, err := mssql.NewProvider(ctx, "provider-mssql", &mssql.ProviderArgs{
	Hostname:  pulumi.String("example.database.windows.net"),
	Port:      pulumi.Int(1433),
	AzureAuth: AzureAuth: mssql.ProviderAzureAuthArgs{},
})
if err != nil {
	return err
}
return nil
})
```
{{% /choosable %}}

{{% choosable language csharp %}}
```csharp
var providerMssql = new Mssql.Provider("provider-mssql", new()
{
    Hostname = "example.database.windows.net",
    Port = 1433,
    AzureAuth = new Mssql.Inputs.ProviderAzureAuthArgs(),
});
```
{{% /choosable %}}

{{% choosable language yaml %}}
```yaml
resources:
  provider-mssql:
    type: pulumi:providers:mssql
    properties:
      hostname: "example.database.windows.net"
      port: 1433
      azureAuth: {}
```
{{% /choosable %}}

{{< /chooser >}}

### Provider Binary

The MSSQL provider binary is a third party binary. It can be installed using the `pulumi plugin` command.
