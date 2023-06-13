---
title: Migration from Azure Native v1 to v2
meta_desc: How to migrate from v1 to v2 of the Pulumi Azure Native Provider.
---

Azure Native v2 is currently in preview status. This guide 

## Key Differences between v1 and v2

### Upgraded Default Versions
This release upgrades all resources to use the latest available versions of Azure services. This gives access to all the latest features. This includes some breaking changes in the properties of some resources from the previous version.

A full list of default version changes can be found in the [top-level resource versions](top-level-resource-versions).

Where there are breaking changes these may appear as either a compilation error in a program or by showing as requiring replacement during preview due to the structure of the resources changing.

To continue using the previous version of a resource:

1. Check the documentation which lists the previous version for each resource. For example: `Azure REST API Version: 2022-05-01. Prior API version in Azure Native 1.x: 2020-07-17-preview`
2. Import the previous version of the resource. These are available in the version-specific sub-folders of the SDK.

Below are examples of changing an import to use a specific version in each language.

{{< chooser language "typescript,python,csharp,go,java,yaml" >}}

{{% choosable language typescript %}}

```typescript
- import { EventSubscription } from "@pulumi/azure-native/eventgrid";
+ import { EventSubscription } from "@pulumi/azure-native/eventgrid/v20200601";
```

{{% /choosable %}}

{{% choosable language csharp %}}

```csharp
- using EventGrid = Pulumi.AzureNative.EventGrid;
+ using EventGrid20200601 = Pulumi.AzureNative.EventGrid.V20200601;
```

{{% /choosable %}}

{{% choosable language python %}}

```python
- from pulumi_azure_native import eventgrid
+ from pulumi_azure_native.eventgrid import v20200601 as eventgrid20200601
```

{{% /choosable %}}

{{% choosable language go %}}

```go
- import eventgrid "github.com/pulumi/pulumi-azure-native-sdk/eventgrid/v2"
+ import eventgrid20200601 "github.com/pulumi/pulumi-azure-native-sdk/eventgrid/v2/v20200601”
```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
- type: azure-native:eventgrid:Topic
+ type: azure-native:eventgrid/v20200601:Topic
Remove Deprecated Explicit Versions
```

{{% /choosable %}}
{{< /chooser >}}

### Removed Deprecated Versions

Since v1.103.0, deprecation notices have been introduced on specific older versions of Azure’s REST API. With the release of Azure Native 2.0, resources marked as deprecated in 1.x will not be available in 2.0 and later. We recommend migrating any affected resources to a newer version of Azure’s REST API.

The primary reason for these removals is to limit the growth rate of the SDKs. Most languages have limits on the size of an SDK; and we’ve reached a few already which have required some work to mitigate. We will continue to look at long term solutions to address these limitations but this is necessary work to ensure the stability of the provider. Additionally, we've included compatible versions of these resources which will be present in 2.0 in the above deprecation warnings. This should provide a painless upgrade process to 1.103 or later, and similarly for 2.0.


### Title-Casing Resource Names

All resource names are now consistently in title-case, starting with an upper-case letter. About 40 resources that started with a lower-case letter in v1 were renamed. This change is mostly internal, but some function names changed in the process.

This is a breaking change in user programs using these functions. Fixing it is as simple as updating the function names. For instance, in the TypeScript SDK, `azure-native.aadiam.getazureADMetric` changed to `azure-native.aadiam.getAzureADMetric`.

The old lower-case resource names are aliased so all Pulumi state is backwards-compatible. Upgrading to the new names will not result in changes to existing stacks.

The complete list of affected resources is [in the PR](https://github.com/pulumi/pulumi-azure-native/pull/2366).

### MySQL and PostgreSQL Server and Flexible Server

Both [Azure Database for MySQL](https://azure.microsoft.com/en-us/products/mysql) and [Azure Database for PostgreSQL](https://azure.microsoft.com/en-us/products/postgresql) are available in a “Single Server” and a “Flexible Server” variant. The Single Server variants are on the retirement path ([MySQL](https://learn.microsoft.com/en-us/azure/mysql/single-server/whats-happening-to-mysql-single-server), [PostgreSQL](https://learn.microsoft.com/en-us/azure/postgresql/single-server/whats-happening-to-postgresql-single-server)). The MySQL single servers already cannot be created in the Azure portal anymore, although the API will continue to work until September 2024. All new servers are encouraged to be of the flexible kind.

The Azure specification introduced the new Flexible Server variants in a way that shares some names with the existing Single Server variants, leading to conflicts for these resources:
- `Configuration`
- `Database`
- `FirewallRule`
- `Server`
- `PrivateEndpointConnection`

Given that the Single Server variants are being retired, we decided to use these conflicting resource names for the new Flexible Server. Single Server will not be available in the default API version at all, but will remain available via explicitly importing older API versions.

For existing Pulumi programs using v1 of the Azure Native provider, this can mean that upgrading to v2 will result in a change of resource type for the above resources. For instance, `azure-native.dbformysql.Server` would previously have referred to a Single Server but will now refer to a Flexible Server, which is a different resource type. This will result in a replacement of the resource during the next `pulumi up`. However, the properties of flexible servers are sufficiently different that, in a typed language, the program will not compile.


### Simpler User Assigned Identities Inputs
“User assigned identity” inputs are now represented as a simple list of strings in each language.

Previously this was a map type where the value was the identity identifier and the value must be set to an empty property bag.

A list of affected resources can be found in the pull request summary.

## Configuration

We recommend upgrading to the latest version of the v1 provider as a prequisiste to the v2 upgrade. This will enable you to resolve any pre-existing deprecated versions or resource structure changes.

## Mix and match versions in the same project

It is possible to use multiple versions of the Azure Native provider or ARM resource versions within the same Pulumi program.

TODO: Add code sample