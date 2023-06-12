---
title: Migration from Azure Native v1 to v2
meta_desc: How to migrate from v1 to v2 of the Pulumi Azure Native Provider.
---

Azure Native v2 is currently in preview status. This guide 

## Key Differences between v1 and v2

### Upgraded Default Versions
This release upgrades all resources to use the latest available versions of Azure services. This gives access to all the latest features. This includes some breaking changes in the properties of some resources from the previous version.

A full list of default version changes can be found …

Where there are breaking changes these may appear as either a compilation error in a program or by showing as requiring replacement during preview due to the structure of the resources changing.

To continue using the previous version of a resource:
1. Check the documentation which lists the previous version for each resource. For example: `"API Version: 2022-09-01. Previous API Version: 2019-05-01.`
2. Import the previous version of the resource. These are available in the version-specific sub-folders of the SDK.

Below are examples of changing an import to use a specific version in each language.

{{< chooser language "typescript,python,csharp,go,java,yaml" >}}

{{% choosable language typescript %}}

```typescript
- import { EventSubscription } from "@pulumi/azure-native/eventgrid";
+ import { EventSubscription } from "@pulumi/azure-native/eventgrid/v20200601";
```

{{% /choosable %}}

```C#
- using EventGrid = Pulumi.AzureNative.EventGrid;
+ using EventGrid20200601 = Pulumi.AzureNative.EventGrid.V20200601;
```

{{% choosable language python %}}

```Python
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

### Removed Deprecated Versions

Since v1.103.0, deprecation notices have been introduced on specific older versions of Azure’s REST API. With the release of Azure Native 2.0, resources marked as deprecated in 1.x will not be available in 2.0 and later. We recommend migrating any affected resources to a newer version of Azure’s REST API.

The primary reason for these removals is to limit the growth rate of the SDKs. Most languages have limits on the size of an SDK; and we’ve reached a few already which have required some work to mitigate. We will continue to look at long term solutions to address these limitations but this is necessary work to ensure the stability of the provider. Additionally, we've included compatible versions of these resources which will be present in 2.0 in the above deprecation warnings. This should provide a painless upgrade process to 1.103 or later, and similarly for 2.0.


### Title-Casing Resource Names

All resource names are now title-cased. Some ~40 that started with a lower-case letter in v1 were renamed, which is a breaking change for some existing user programs.

The previous names are aliased to the new names so upgrading to the new resource names should result in no changes to existing stacks.

The complete list of renamed resources is in this PR: https://github.com/pulumi/pulumi-azure-native/pull/2366

Note that for Go and Java there’s no change as these languages require types to begin with an upper-case letter anyway.

### MySQL and PostgreSQL Server and Flexible Server

Both “Azure Database for MySQL/PostgreSQL” are available in a “Single Server” and a “Flexible Server” variant. The Single Server variant is on the retirement path (mysql, postgres). The mysql ones already cannot be created in the Azure portal anymore, although the API will continue to work until September 2024. All new servers are encouraged to be of the flexible kind.

The following resources that share the same name but a different path between single and flexible variants:
- Configuration
- Database
- FirewallRule
- Server
- PrivateEndpointConnection

### Simpler User Assigned Identities Inputs
“User assigned identity” inputs are now represented as a simple list of strings in each language.

Previously this was a map type where the value was the identity identifier and the value must be set to an empty property bag.

A list of affected resources can be found in the pull request summary.

## Configuration

We recommend upgrading to the latest version of the v1 provider as a prequisiste to the v2 upgrade. This will enable you to resolve any pre-existing deprecated versions or resource structure changes.

## Mix and match versions in the same project

It is possible to use multiple versions of the Azure Native provider or ARM resource versions within the same Pulumi program.

TODO: Add code sample