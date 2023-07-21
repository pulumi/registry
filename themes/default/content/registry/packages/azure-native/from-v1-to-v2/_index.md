---
title: Upgrade Guide for Azure Native v1 to v2
meta_desc: How to upgrade from v1 to v2 of the Pulumi Azure Native Provider.
layout: package
---

## About

The Pulumi Azure Native Provider v2 is now available. You can start taking advantages of the smaller, streamlined SDKs right away while also incorporating other enhancements including simplified user assigned identity handling, refreshed default versions, and consistent title casing. To upgrade to v2, there are a few changes you may need to make depending on your Pulumi programs.

## Upgrade Steps

### Prerequisites

We recommend upgrading to the [latest version](https://github.com/pulumi/pulumi-azure-native/releases/tag/v1.104.0) of the v1 provider as a prerequisite to the v2 upgrade. This will enable you to resolve any pre-existing deprecated versions or resource structure changes.

Review your program for warnings on missing imports or deprecated resources. If your program contains any deprecated explicit versions, you will need to update these to a newer version. The suggested version is shown in the deprecation message.

![azure-deprecation-cli](./azure-deprecation-cli.png)

![azure-deprecation-ide](./azure-deprecation-ide.png)

### Upgrade Dependencies

In your Pulumi program, upgrade the package to point to the latest v2.x version.

* JavaScript/TypeScript: [`@pulumi/azure-native`](https://www.npmjs.com/package/@pulumi/azure-native/v/2.0.0)
* Python: [`pulumi-azure-native`](https://pypi.org/project/pulumi-azure-native/2.0.0/)
* Go: [`github.com/pulumi/pulumi-azure-native/sdk/go/azure`](https://github.com/pulumi/pulumi-azure-native/releases/tag/v2.0.0)
* .NET: [`Pulumi.AzureNative`](https://www.nuget.org/packages/Pulumi.AzureNative/2.0.0)
* Java: [`com.pulumi.azurenative`](https://central.sonatype.com/artifact/com.pulumi/azure-native/2.0.0)

{{< chooser language "typescript,python,csharp,go" >}}

{{% choosable language typescript %}}

```
-"@pulumi/azure-native": "^1.0.0",
- "resolved": "https://registry.npmjs.org/@pulumi/azure-native/-/azure-native-1.104.0.tgz",
- "resolved": "https://registry.npmjs.org/@pulumi/azure-native/-/azure-native-1.104.0.tgz",
+ @pulumi/azure-native": "v2.0.0",
+ "resolved": "https://registry.npmjs.org/@pulumi/azure-native/-/azure-native-2.0.0.tgz",
+ "resolved": "https://registry.npmjs.org/@pulumi/azure-native/-/azure-native-2.0.0.tgz",
```

{{% /choosable %}}
{{% choosable language python %}}

```python
- pulumi_azure_native==v1.104.0
+ pulumi_azure_native==v2.0.0
```

{{% /choosable %}}
{{% choosable language csharp %}}

```csharp
- <PackageReference Include="Pulumi.AzureNative" Version="1.104.0"
+ <PackageReference Include="Pulumi.AzureNative" Version="2.0.0"
```

{{% /choosable %}}
{{% choosable language go %}}

```go
- github.com/pulumi/pulumi-azure-native-sdk/storage v1.104.0
+ github.com/pulumi/pulumi-azure-native-sdk/v2/storage v2.0.0
```

{{% /choosable %}}

{{% /chooser %}}

### Upgrade Imports

Go programs will need all imports updated to include `v2` in the path.

```go
import (
- "github.com/pulumi/pulumi-azure-native-sdk/resources"
- "github.com/pulumi/pulumi-azure-native-sdk/storage"
+ "github.com/pulumi/pulumi-azure-native-sdk/resources/v2"
+ "github.com/pulumi/pulumi-azure-native-sdk/storage/v2"
)
```

If you do not replace the deprecated import, you can expect to see an error message:

```bash
    error: Running program '<path/to/project>' failed with an unhandled exception:
Error: Cannot find module '@pulumi/azure-native/resources/v20210501'
```

### Review Pulumi Diff

After updating all imports to reflect v2 included Azure API versions, run `pulumi preview` and review the output.

#### Pending Changes on Default Versions

You may see pending changes when using the default version as the shape of the resource may have changed. You can choose to accept the changes, update your program to modify the resource properties to mitigate changes, or continue using the previous default version from v1.

A full list of default version changes can be found in the [top-level resource versions](./top-level-resource-versions).

To continue using the previous Azure API version of a resource:

1. Check the documentation in your IDE or our [registry API docs](https://www.pulumi.com/registry/packages/azure-native/) which identifies the previous version for each resource. For example: `Azure REST API Version: 2022-06-15. Prior API version in Azure Native 1.x: 2020-06-01`
2. Import the previous version of the resource. These are available in the version-specific sub-folders of the SDK.

Below are examples of changing an import to use an explicit version in each language.

{{< chooser language "typescript,python,csharp,go,yaml" >}}

{{% choosable language typescript %}}

```typescript
- import { EventSubscription } from "@pulumi/azure-native/eventgrid";
+ import { EventSubscription } from "@pulumi/azure-native/eventgrid/v20200601";
```

{{% /choosable %}}
{{% choosable language python %}}

```python
- from pulumi_azure_native import eventgrid
+ from pulumi_azure_native.eventgrid import v20200601 as eventgrid
```

{{% /choosable %}}
{{% choosable language csharp %}}

```csharp
- using EventGrid = Pulumi.AzureNative.EventGrid;
+ using EventGrid = Pulumi.AzureNative.EventGrid.V20200601;
```

{{% /choosable %}}
{{% choosable language go %}}

```go
- import eventgrid "github.com/pulumi/pulumi-azure-native-sdk/eventgrid/v2"
+ import eventgrid "github.com/pulumi/pulumi-azure-native-sdk/eventgrid/v2/v20200601”
```

{{% /choosable %}}
{{% choosable language yaml %}}

```yaml
- type: azure-native:eventgrid:Topic
+ type: azure-native:eventgrid/v20200601:Topic
```

{{% /choosable %}}
{{< /chooser >}}

#### User Assigned Identity Inputs

`User assigned identity` inputs are now represented as a simple string array in each language instead of a map type. Where you are referencing user assigned identity inputs, you will need to update the syntax to resolve the error.

```typescript
import * as resources from "@pulumi/azure-native/resources";
import * as managedidentity from "@pulumi/azure-native/managedidentity";
import * as storage from "@pulumi/azure-native/storage";

// Create a resource group
const resourceGroup = new resources.ResourceGroup("my-resource-group");

// Create a user-assigned managed identity
const userAssignedIdentity = new managedidentity.UserAssignedIdentity("my-user-assigned-identity", {
    resourceGroupName: resourceGroup.name,
});

// Create a storage account that references the user-assigned managed identity
const storageAccount = new storage.StorageAccount("mystorageaccount", {
    resourceGroupName: resourceGroup.name,
    kind: "StorageV2",
    location: resourceGroup.location,
    identity: {
        type: "UserAssigned",
        userAssignedIdentities: [userAssignedIdentity.id],
    },
    sku: {
        name: "Standard_LRS",
    },
});
```

Before this change, the identity block would need to use an apply as follows:

```typescript
user_assigned_identity.id.apply(lambda id: {id: {}})
```

#### Title Case Changes

All resource names are now consistently in title case, starting with an uppercase letter. The previous lowercase resource names are aliased so all Pulumi state is backwards-compatible. While functions are not aliased, resolving the name discrepancy is as simple as updating the function names. For instance, in the TypeScript SDK, `azure-native.aadiam.getazureADMetric` changed to `azure-native.aadiam.getAzureADMetric`.

The complete list of affected resources and functions is [in this PR](https://github.com/pulumi/pulumi-azure-native/pull/2366).

#### MySQL and PostgreSQL Server and Flexible Server

Both [Azure Database for MySQL](https://azure.microsoft.com/en-us/products/mysql) and [Azure Database for PostgreSQL](https://azure.microsoft.com/en-us/products/postgresql) are available in a `Single Server` and a `Flexible Server` variant. The `Single Server` variants are on the retirement path ([MySQL](https://learn.microsoft.com/en-us/azure/mysql/single-server/whats-happening-to-mysql-single-server), [PostgreSQL](https://learn.microsoft.com/en-us/azure/postgresql/single-server/whats-happening-to-postgresql-single-server)). Azure recommends that all new servers are created as a `Flexible Server` variant.

In v2, the following resources are now associated with a `Flexible Server` variant instead of `Single Server` as they were in v1:

* `Configuration`
* `Database`
* `FirewallRule`
* `Server`
* `PrivateEndpointConnection`

Existing v1 programs upgrading to v2 using the default version will result in a change of resource type. For instance, `azure-native.dbformysql.Server` would previously have referred to a `Single Server` but will now refer to a `Flexible Server` and will result in a replacement of the resource during the next `pulumi up`. However, the properties of flexible servers are sufficiently different that, in a typed language, the program will not compile. If you would like to continue using the previous Azure API version, you may do so by using the previous explicit version `2017-12-01`.

{{< chooser language "typescript,python,csharp,go,yaml" >}}

{{% choosable language typescript %}}

```typescript
- import * as mysqldb from "@pulumi/azure-native/dbformysql"
+ import * as mysqldb from "@pulumi/azure-native/dbformysql/v20171201"

- import * as postgresqldb from "@pulumi/azure-native/dbforpostgresql"
+ import * as postgresqldb from "@pulumi/azure-native/dbforpostgresql/v20171201"
```

{{% /choosable %}}
{{% choosable language python %}}

```python
- from pulumi_azure_native import dbformysql
+ from pulumi_azure_native.dbformysql import v20171201 as dbformysql

- from pulumi_azure_native import dbforpostgresql
+ from pulumi_azure_native.dbforpostgresql import v20171201 as dbforpostgresql
```

{{% /choosable %}}
{{% choosable language csharp %}}

```csharp
- using MySQLDB = Pulumi.AzureNative.DBforMySQL;
+ using MySQLDB = Pulumi.AzureNative.DBforMySQL.V20171201;

- using PostgreSQLDB = Pulumi.AzureNative.DBforPostgreSQL;
+ using PostgreSQLDB = Pulumi.AzureNative.DBforPostgreSQL.V20171201;
```

{{% /choosable %}}
{{% choosable language go %}}

```go
- import mysqldb "github.com/pulumi/pulumi-azure-native-sdk/dbformysql"
+ import mysqldb "github.com/pulumi/pulumi-azure-native-sdk/dbformysql/v2/v20200601"

- import postgresqldb "github.com/pulumi/pulumi-azure-native-sdk/dbforpostgresql"
+ import postgresqldb "github.com/pulumi/pulumi-azure-native-sdk/dbforpostgresql/v2/20171201"
```

{{% /choosable %}}
{{% choosable language yaml %}}

```yaml
- type: azure-native:dbformysql:Server
+ type: azure-native:dbformysql/v20171201:Server

- type: azure-native:dbforpostgresql:Server
+ type: azure-native:dbforpostgresql/v20171201:Server
```

{{% /choosable %}}
{{< /chooser >}}

### Contributing

If you experience any unexpected issues during your migration or would like to contribute to our codebase, please visit our [respository](https://github.com/pulumi/pulumi-azure-native) to open an [issue](https://github.com/pulumi/pulumi-azure-native/issues) or submit a pull request.
