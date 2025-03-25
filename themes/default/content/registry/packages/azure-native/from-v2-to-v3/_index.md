---
title: Upgrade Guide for Azure Native v2 to v3
meta_desc: How to upgrade from v2 to v3 of the Pulumi Azure Native Provider.
layout: package
---

## About

The Pulumi Azure Native Provider v3 is now available. To upgrade to v3 and take advantages of the smaller, streamlined SDKs, there are a few changes you may need to make depending on your Pulumi programs.

## Upgrade Steps

### Prerequisites

We recommend upgrading to the [latest version](https://github.com/pulumi/pulumi-azure-native/releases/tag/v2.89.3) of the v2 provider as a prerequisite to the v3 upgrade. This will enable you to resolve any pre-existing deprecated versions or resource structure changes.

Review your program for warnings on missing imports or deprecated resources. If your program uses any deprecated explicit API versions where the import contains the API version, you will need to update these since explicit API versions are no longer included in the SDK. It's recommnded to migrate to the default version, if possible, or you can generate a local SDK for the desired API version as described in the [version guide](./version-guide).

### Upgrade Dependencies

In your Pulumi program, upgrade the package to point to the latest v3.x version.

* JavaScript/TypeScript: [`@pulumi/azure-native`](https://www.npmjs.com/package/@pulumi/azure-native/v/3.0.0)
* Python: [`pulumi-azure-native`](https://pypi.org/project/pulumi-azure-native/3.0.0/)
* Go: [`github.com/pulumi/pulumi-azure-native/sdk/go/azure`](https://github.com/pulumi/pulumi-azure-native/releases/tag/v3.0.0)
* .NET: [`Pulumi.AzureNative`](https://www.nuget.org/packages/Pulumi.AzureNative/3.0.0)
* Java: [`com.pulumi.azurenative`](https://central.sonatype.com/artifact/com.pulumi/azure-native/3.0.0)

{{< chooser language "typescript,python,csharp,go" >}}

{{% choosable language typescript %}}

```
-"@pulumi/azure-native": "^2.0.0",
- "resolved": "https://registry.npmjs.org/@pulumi/azure-native/-/azure-native-2.89.3.tgz",
- "resolved": "https://registry.npmjs.org/@pulumi/azure-native/-/azure-native-2.89.3.tgz",
+ @pulumi/azure-native": "v3.0.0",
+ "resolved": "https://registry.npmjs.org/@pulumi/azure-native/-/azure-native-3.0.0.tgz",
+ "resolved": "https://registry.npmjs.org/@pulumi/azure-native/-/azure-native-3.0.0.tgz",
```

{{% /choosable %}}
{{% choosable language python %}}

```python
- pulumi_azure_native==v2.89.3
+ pulumi_azure_native==v3.0.0
```

{{% /choosable %}}
{{% choosable language csharp %}}

```csharp
- <PackageReference Include="Pulumi.AzureNative" Version="2.89.3"
+ <PackageReference Include="Pulumi.AzureNative" Version="3.0.0"
```

{{% /choosable %}}
{{% choosable language go %}}

```go
- github.com/pulumi/pulumi-azure-native-sdk/storage v2.89.3
+ github.com/pulumi/pulumi-azure-native-sdk/v3/storage v3.0.0
```

{{% /choosable %}}

{{% /chooser %}}

### Upgrade Imports

Go programs will need all imports updated to include `v3` in the path.

```go
import (
- "github.com/pulumi/pulumi-azure-native-sdk/resources/v2"
- "github.com/pulumi/pulumi-azure-native-sdk/storage/v2"
+ "github.com/pulumi/pulumi-azure-native-sdk/resources/v3"
+ "github.com/pulumi/pulumi-azure-native-sdk/storage/v3"
)
```

If you do not replace the deprecated import, you can expect to see an error message:

```bash
    error: Running program '<path/to/project>' failed with an unhandled exception:
Error: Cannot find module '@pulumi/azure-native/resources/v20210501'
```

### Review Pulumi Diff

After updating all imports to reflect v3 included Azure API versions, run `pulumi preview` and review the output.

#### Pending Changes on Default Versions

You may see pending changes when using the default version as the shape of the resource may have changed. You can choose to accept the changes, update your program to modify the resource properties to mitigate changes, or choose an API version which is compatible with the older version of the resource.

A full list of default version changes can be found in the [resource versions table](./resource-versions).

To continue using the previous Azure API version of a resource:

1. Check the resource documentation in your IDE or our [registry API docs](https://www.pulumi.com/registry/packages/azure-native/) which identifies the previous version for each resource. For example: `Azure REST API Version: 2024-10-01. Prior API version in Azure Native 2.x: 2023-05-01`.
2. Generate and use the local SDK for the previous version of the resource as described in the [version guide](./version-guide).

#### New module structure and naming aligned closer to Azure SDKs

The Azure specification sometimes contains related but distinct services in one namespace. For example, `Microsoft.Network` contains `DNS`, `DnsResolver`, `Frontdoor`, `Network`, `PrivateDns`, and `TrafficManager`.

The Azure SDKs reflect this so that these different network-related services are logically grouped instead of mixed up in one `Network` namespace. In v3 of this provider, we have aligned several modules to match the Azure SDK layout. The changes are:

- `Cache` is split into `Redis` and `RedisEnterprise` which also reflects the current names of the services.
- `Devices` is split into `DeviceProvisioningServices` and `IotHub`.
- `DocumentDB` is not split but renamed to `CosmosDB` to reflect the current naming of the service.
- `Insights` is split into `ApplicationInsights` and `Monitor`.
- `Network` is split into `Dns`, `DnsResolver`, `Frontdoor`, `Network`, `PrivateDns`, and `TrafficManager`.

This changes require corresponding updates in your Pulumi programs. Existing resources will not be recreated, however, due to the use of aliases.

#### De-duplication of conflicting resources and types

The Azure specification sometimes uses the same name for different resources in the same module. For example, `CosmosDB.FirewallRule` can refer to both regular and MongoDB (`Microsoft.DocumentDB/mongoClusters`) instances. In v3, we have de-duplicated these to `FirewallRule` and `MongoClusterFirewallRule`. The complete list of changes are:

- In `CosmosDB`, `FirewallRule` was split into `FirewallRule` and `MongoClusterFirewallRule`.
- In `HDInsight`, `Cluster` was split into `Cluster` and `ClusterPoolCluster`.
- In `HybridContainerService`, `HybridIdentityMetadatum` was split into `HybridIdentityMetadata` and `ClusterInstanceHybridIdentityMetadatum`.
- In `NetApp`, resources relating to Capacity Pools are prefixed with `CapacityPool`, splitting `Backup` into `Backup` and `CapacityPoolBackup`.

#### Propagating `forceNew` from referenced types

Tracked by [#3006](https://github.com/pulumi/pulumi-azure-native/issues/3006), this change extended the [`replaceOnChanges`](https://www.pulumi.com/docs/iac/concepts/options/replaceonchanges/) parameter to be propagated from referenced types. This improves the correctness of the provider by correctly replacing resources when sub-properties of the resource that are annotated with `replaceOnChanges` are updated.

You can set the environment variable `PULUMI_FORCE_NEW_FROM_SUBTYPES` to `false` to disable this behavior if you observe unnecessary recreation of resources. Please [file an issue](https://github.com/pulumi/pulumi-azure-native/issues) in that case since that would be unexpected.

#### Improved flattening of resource shapes

Tracked by [#3195](https://github.com/pulumi/pulumi-azure-native/issues/3195), this change improved the correctness of the provider by avoiding to flatten nested properties, even when requested by the Azure API specification, when it would lead to the overwriting of properties. 

For instance, the `network:NetworkVirtualApplianceConnection` resource has a property called `Properties` which has in turn, slightly simplified, two properties: `Name` and `TunnelIdentifier`.

```
NetworkVirtualApplianceConnection:
  Name
  Properties:
    Name
    TunnelIdentifier
```

`Properties` is annotated in the API spec to be flattened into the outer `NetworkVirtualApplianceConnection`. That's what v2 of the provider does, with the result being:

```
NetworkVirtualApplianceConnection:
  Name
  TunnelIdentifier
```

Since `NetworkVirtualApplianceConnection` already has a `Name` property, the two `Name`s are in conflict. In v3, the provider does not flatten `Properties` due to this conflict and the result is:

```
NetworkVirtualApplianceConnection:
  Name
  Properties:
    Name
    TunnelIdentifier
```

All affected resources are listed below:

- `azure-native:authorization:AccessReviewHistoryDefinitionById`
- `azure-native:authorization:AccessReviewScheduleDefinitionById`
- `azure-native:authorization:ScopeAccessReviewHistoryDefinitionById`
- `azure-native:authorization:ScopeAccessReviewScheduleDefinitionById`
- `azure-native:devhub:Workflow`
- `azure-native:education:Lab`
- `azure-native:network:P2sVpnServerConfiguration`
- `azure-native:network:VpnServerConfiguration`
- `azure-native:security:DefenderForStorage`
- `azure-native:vmwarecloudsimple:DedicatedCloudNode`

#### MySQL and PostgreSQL Server and Flexible Server

Tracked by [#2753](https://github.com/pulumi/pulumi-azure-native/issues/2753), this change cleanly separates the different PostgreSQL offerings from Azure: Flexible Servers, Server Groups V2, and the deprecated Single Server offering.

TODO,tkappler elaborate with examples

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

#### Removed `__inputs` from state

This change is completely transparent to users, except for those who have custom code or scripts working with Pulumi stack state. In v2 and older, the provider added a field `__inputs` to the state of each resources which tracked the inputs of the resource at the time of its creation. This is not necessary anymore since the Pulumi engine [now supports sending the old inputs](https://github.com/pulumi/pulumi/pull/13139).

### Contributing

If you experience any unexpected issues during your migration or would like to contribute to our codebase, please visit our [respository](https://github.com/pulumi/pulumi-azure-native) to open an [issue](https://github.com/pulumi/pulumi-azure-native/issues) or submit a pull request.
