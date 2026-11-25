---
title: Upgrade Guide for Azure Native v2 to v3
meta_desc: How to upgrade from v2 to v3 of the Pulumi Azure Native Provider.
layout: package
linkTitle: Upgrade to v3
weight: 2
aliases:
  - /registry/packages/azure-native/from-v2-to-v3/
---

## About

This guide covers upgrading from v2 to v3 of the Pulumi Azure Native Provider. To upgrade to v3 and take advantage of the smaller SDKs and other improvements, there are a few changes you may need to make depending on your Pulumi programs.

## Upgrade Steps

### Prerequisites

We recommend upgrading to the [latest version](https://github.com/pulumi/pulumi-azure-native/releases/tag/v2.89.3) of the v2 provider as a prerequisite to the v3 upgrade. This will enable you to resolve any pre-existing deprecated versions or resource structure changes.

Review your program for warnings on missing imports or deprecated resources. If your program uses any explicit API versions where the import contains the API version, you will need to update these since explicit API versions are no longer included in the SDK. It's recommended to migrate to the default version, if possible, or you can generate a local SDK for the desired API version as described in the [version guide](https://www.pulumi.com/registry/packages/azure-native/guides/version-guide/#accessing-any-api-version-via-local-packages).

### Upgrade Dependencies

In your Pulumi program, upgrade the package to point to the latest v3.x version.

* JavaScript/TypeScript: [`@pulumi/azure-native`](https://www.npmjs.com/package/@pulumi/azure-native/v/3.0.0)
* Python: [`pulumi-azure-native`](https://pypi.org/project/pulumi-azure-native/3.0.0/)
* Go: [`github.com/pulumi/pulumi-azure-native-sdk`](https://github.com/pulumi/pulumi-azure-native-sdk), one module per Azure service, such as `github.com/pulumi/pulumi-azure-native-sdk/storage/v3`
* .NET: [`Pulumi.AzureNative`](https://www.nuget.org/packages/Pulumi.AzureNative/3.0.0)
* Java: [`com.pulumi.azurenative`](https://central.sonatype.com/artifact/com.pulumi/azure-native/3.0.0)

{{< chooser language "typescript,python,csharp,go" >}}

{{% choosable language typescript %}}

```
-"@pulumi/azure-native": "^2.0.0",
- "resolved": "https://registry.npmjs.org/@pulumi/azure-native/-/azure-native-2.89.3.tgz",
+"@pulumi/azure-native": "v3.0.0",
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

In go.mod:
```go
- github.com/pulumi/pulumi-azure-native-sdk/storage v2.89.3
+ github.com/pulumi/pulumi-azure-native-sdk/storage/v3 v3.0.0
```

Go programs will also need all imports updated to include `v3` in the path.
```go
import (
- "github.com/pulumi/pulumi-azure-native-sdk/resources/v2"
- "github.com/pulumi/pulumi-azure-native-sdk/storage/v2"
+ "github.com/pulumi/pulumi-azure-native-sdk/resources/v3"
+ "github.com/pulumi/pulumi-azure-native-sdk/storage/v3"
)
```

{{% /choosable %}}

{{% /chooser %}}

### Review Pulumi Diff

After updating all imports to reflect v3 included Azure API versions, run `pulumi preview` and review the output.

#### Pending Changes on Default Versions

You may see pending changes when using the default version as the shape of the resource may have changed. You can choose to accept the changes, update your program to modify the resource properties to mitigate changes, or choose an API version which is compatible with the older version of the resource.

A full list of default version changes can be found in the [resource versions table](/registry/packages/azure-native/guides/from-v2-to-v3/resource-versions/).

To continue using the previous Azure API version of a resource:

1. Check the resource documentation in your IDE or our [registry API docs](https://www.pulumi.com/registry/packages/azure-native/) which identifies the previous version for each resource. For example:
  > _Uses Azure REST API version 2024-10-01. In version 2.x of the Azure Native provider, it used API version 2023-05-01._
2. Generate and use the local SDK for the previous version of the resource as described in the [version guide](/registry/packages/azure-native/guides/version-guide/).

#### New module structure and naming aligned closer to Azure SDKs

The Azure specification sometimes contains related but distinct services in one namespace. For example, `Microsoft.Network` contains `DNS`, `DnsResolver`, `Frontdoor`, `Network`, `PrivateDns`, and `TrafficManager`.

The Azure SDKs reflect this so that these different network-related services are logically grouped instead of mixed up in one `Network` namespace. In v3 of this provider, we have aligned several modules to match the Azure SDK layout. The changes are:

- `Cache` is split into `Redis` and `RedisEnterprise` which also reflects the current names of the services.
- `Devices` is split into `DeviceProvisioningServices` and `IotHub`.
- `DocumentDB` is not split but renamed to `CosmosDB` to reflect the current naming of the service.
- `Insights` is split into `ApplicationInsights` and `Monitor`.
- `Network` is split into `Dns`, `DnsResolver`, `Frontdoor`, `Network`, `PrivateDns`, and `TrafficManager`.

You'll need to update your code to use the new module names. The resources in these modules have been aliased to support this migration with no replacements of existing resources.

#### De-duplication of conflicting resources and types

The Azure specification sometimes uses the same name for different resources in the same module. For example, `CosmosDB.FirewallRule` can refer to both regular and MongoDB (`Microsoft.DocumentDB/mongoClusters`) instances. In v3, we have de-duplicated these to `FirewallRule` and `MongoClusterFirewallRule`. The complete list of changes are:

- In `CosmosDB`, `FirewallRule` was split into `FirewallRule` and `MongoClusterFirewallRule`.
- In `HDInsight`, `Cluster` was split into `Cluster` and `ClusterPoolCluster`.
- In `HybridContainerService`, `HybridIdentityMetadatum` was split into `HybridIdentityMetadata` and `ClusterInstanceHybridIdentityMetadatum`.
- In `NetApp`, resources relating to Capacity Pools are prefixed with `CapacityPool`, splitting `Backup` into `Backup` and `CapacityPoolBackup`.

You'll need to update your code to use the new names. The resources in these modules have been aliased to support this migration with no replacements of existing resources.

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

#### MySQL and PostgreSQL Single Server and Azure Cosmos DB for PostgreSQL

Tracked by [#2753](https://github.com/pulumi/pulumi-azure-native/issues/2753). The `dbformysql` and `dbforpostgresql` modules cover more than one Azure offering, and some resources from different offerings share a name in the Azure API, so in v2 some of them were missing from the default modules. In v3, they have distinct names:

- Single Server resources (API version `2017-12-01`) are in the default modules with a `SingleServer` prefix, such as `SingleServer`, `SingleServerDatabase`, and `SingleServerFirewallRule`. In v2, these were only available from explicit API-version modules such as `dbformysql/v20171201`, which v3 doesn't include.
- Azure Cosmos DB for PostgreSQL resources in `dbforpostgresql` have a `ServerGroup` prefix: `Cluster` is now `ServerGroupCluster`, and `Role` is now `ServerGroupRole`. `ServerGroupFirewallRule` and `ServerGroupPrivateEndpointConnection` are new.

Resources without a prefix, such as `Server` and `FirewallRule`, refer to Flexible Server, as they did in v2. The renamed resources have aliases for their v2 types, so updating your program doesn't replace them.

{{< chooser language "typescript,python,csharp,go,yaml" >}}

{{% choosable language typescript %}}

```typescript
- import * as mysqldb from "@pulumi/azure-native/dbformysql/v20171201";
- const server = new mysqldb.Server("server", { /* ... */ });
+ import * as mysqldb from "@pulumi/azure-native/dbformysql";
+ const server = new mysqldb.SingleServer("server", { /* ... */ });

- const cluster = new postgresqldb.Cluster("cluster", { /* ... */ });
+ const cluster = new postgresqldb.ServerGroupCluster("cluster", { /* ... */ });
```

{{% /choosable %}}
{{% choosable language python %}}

```python
- from pulumi_azure_native.dbformysql import v20171201 as dbformysql
- server = dbformysql.Server("server", ...)
+ from pulumi_azure_native import dbformysql
+ server = dbformysql.SingleServer("server", ...)

- cluster = dbforpostgresql.Cluster("cluster", ...)
+ cluster = dbforpostgresql.ServerGroupCluster("cluster", ...)
```

{{% /choosable %}}
{{% choosable language csharp %}}

```csharp
- using MySQLDB = Pulumi.AzureNative.DBforMySQL.V20171201;
- var server = new MySQLDB.Server("server", new() { /* ... */ });
+ using MySQLDB = Pulumi.AzureNative.DBforMySQL;
+ var server = new MySQLDB.SingleServer("server", new() { /* ... */ });

- var cluster = new PostgreSQLDB.Cluster("cluster", new() { /* ... */ });
+ var cluster = new PostgreSQLDB.ServerGroupCluster("cluster", new() { /* ... */ });
```

{{% /choosable %}}
{{% choosable language go %}}

```go
- import mysqldb "github.com/pulumi/pulumi-azure-native-sdk/dbformysql/v2/v20171201"
- server, err := mysqldb.NewServer(ctx, "server", &mysqldb.ServerArgs{ /* ... */ })
+ import mysqldb "github.com/pulumi/pulumi-azure-native-sdk/dbformysql/v3"
+ server, err := mysqldb.NewSingleServer(ctx, "server", &mysqldb.SingleServerArgs{ /* ... */ })

- cluster, err := postgresqldb.NewCluster(ctx, "cluster", &postgresqldb.ClusterArgs{ /* ... */ })
+ cluster, err := postgresqldb.NewServerGroupCluster(ctx, "cluster", &postgresqldb.ServerGroupClusterArgs{ /* ... */ })
```

{{% /choosable %}}
{{% choosable language yaml %}}

```yaml
- type: azure-native:dbformysql/v20171201:Server
+ type: azure-native:dbformysql:SingleServer

- type: azure-native:dbforpostgresql:Cluster
+ type: azure-native:dbforpostgresql:ServerGroupCluster
```

{{% /choosable %}}
{{< /chooser >}}

#### Removed `__inputs` from state

This change is completely transparent to users, except for those who have custom code or scripts working with Pulumi stack state. In v2 and older, the provider added a field `__inputs` to the state of each resources which tracked the inputs of the resource at the time of its creation. This is not necessary anymore since the Pulumi engine [now supports sending the old inputs](https://github.com/pulumi/pulumi/pull/13139).

### Contributing

If you experience any unexpected issues during your migration or would like to contribute to our codebase, please visit our [respository](https://github.com/pulumi/pulumi-azure-native) to open an [issue](https://github.com/pulumi/pulumi-azure-native/issues) or submit a pull request.
