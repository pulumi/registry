---
title: Azure Native version guide
meta_desc: How Azure API versions are represented in the native Azure provider for Pulumi.
layout: package
---

The native Azure provider SDKs follow the semantic versioning convention, similarly to other Pulumi providers. At the same time, Azure API is structured around date-based API versions like `2020-03-01` defined per Azure resource provider.

This guide describes how we combine both versioning approaches to achieve the following design goals:

- Provide access to the entire API surface of Azure resources.
- Ship new updates as soon as Microsoft publishes them.
- Provide stable SDKs without unexpected breaking changes.
- Keep the SDK size manageable.

## Azure Native Semantic Versioning

Releases of the Azure Native SDKs follow the semantic versioning convention, where each version is identified with three numbers: `major.minor.patch` (for example, `1.2.1`).

- Patch versions contain bug fixes only.
- Minor versions may ship fixes and new functionality that doesn't break existing code: new resources, functions, properties.
- Major versions include breaking changes, where existing user code may need to be modified.

Microsoft has defined [Breaking changes guidelines](https://github.com/Azure/azure-rest-api-specs/blob/master/documentation/Breaking%20changes%20guidelines.md) for all the Azure API specifications. The native Azure provider generates SDKs from the Open API specifications, so we rely on Microsoft to follow their breaking changes guideline.

Occasionally, Microsoft ships breaking changes. For example, they may rename a misspelled property in the spec to match it with the actual service behavior. We treat these occurrences as bug fixes and ship them in patch or minor provider versions without revving the major version.

## Azure Default API Versions

Azure's APIs are versioned. API versions are defined per Azure service (resource provider) and are published frequently. API versions are based on a date, for instance, `2020-03-01` or `2019-05-15-preview`. Breaking changes may occur between API versions.

For each resource, Pulumi chooses a default API version. The default API version is often, but not always, the same for all resources in a resource provider. We choose the default version through a combination of factors, including comparing with Azure SDKs, preferring newer versions, and preferring stable versions over preview versions unless the latest stable version is rather old.

The SDKs represent these default versions of the resources. The [API reference docs](/registry/packages/azure-native/api-docs/) document them.

Following the semantic versioning outlined above, patch versions will not contain API version changes. Minor versions contain API version changes when we determine that they don't cause breaking changes, as far as we can tell from the Azure specification. Therefore, regular minor version upgrades allow users to stay current with Azure API changes while minimizing the risk of breaking changes.

## Other API versions

Other API versions are still available through the two methods detailed below, local packages and the generic resource.

Due to the large size of the Azure API surface, we remove some API versions from our metadata. We determine these versions by comparing them to adjacent ones and removing the ones that are not significantly different. We remove only older API versions. These removed versions are not available via local packages, but they are available via the generic resource.

## Accessing any API Version via Local Packages

We expect most users to use the standard resources with the default API versions most of the time, without the need to specify an explicit API version. There are cases, however, where you may want to use a specific API version:
- you have a resource on an older API version and cannot migrate yet,
- the default version has an issue,
- you want to pin an API version for extra stability , or
- you need a new preview version for access to bleeding edge features.

For such cases, the Pulumi provider allows you to generate a local SDK for a specific API version and use it in your project. This is enabled by our new Local Packages feature, which lets you generate Pulumi Packages locally into your Pulumi project, instead of relying only on packages published to the Pulumi Registry.

A single command generates a local package for a specific API version of an Azure resource provider:

```bash
pulumi package add azure-native storage v20240101
```

This command generates a local package for the `storage` resource provider with the API version `2024-01-01` and adds it to your project.

You can then update your Pulumi program to use the new, local SDK:

{{< chooser language "typescript,python,csharp,go,java,yaml" >}}

{{% choosable language typescript %}}

```typescript
import * as storage_v20240101 from "@pulumi/azure-native_storage_v20240101";

const storageAccount = new storage_v20240101.storage.v20240101.StorageAccount(...)
```

{{% /choosable %}}

{{% choosable language python %}}

```python
TODO
```

{{% /choosable %}}

{{% choosable language csharp %}}

```csharp
TODO
```

{{% /choosable %}}

{{% choosable language go %}}

```go
TODO
```

{{% /choosable %}}

{{% choosable language java %}}

TODO Java is not supported.

{{% /choosable %}}


{{% choosable language yaml %}}

TODO
```yaml
resources:
  my-cluster:
   type: "azure-native:containerservice/v20220301:ManagedCluster"
   properties:
     # ...
```

{{% /choosable %}}

{{< /chooser >}}

## Using the generic resource

The generic resource [azure-native.resources.Resource](../api-docs/resources/resource/) is a special resource that allows you to access any Azure resource at any API version. The downside is that it is not explicitly typed and type-safe like the specific resources. It's useful when you need to access a resource that is not yet supported by the Pulumi provider or when an issue prevents you from using a specific resource.

The example below creates two storage accounts, one using the generic resource and one using the dedicated `StorageAccount` resource for comparison.

{{< chooser language "typescript,python,csharp,go,java,yaml" >}}

{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as azure_native from "@pulumi/azure-native";

const resourceGroup = new azure_native.resources.ResourceGroup("resourceGroup", {});

// Create a Storage Account
const sa = new azure_native.storage.StorageAccount("sa", {
    resourceGroupName: resourceGroup.name,
    sku: {
        name: azure_native.storage.SkuName.Standard_LRS,
    },
    kind: azure_native.storage.Kind.StorageV2,
});

// Create another Storage Account using the generic resource
const generic = new azure_native.resources.Resource("generic", {
    resourceGroupName: resourceGroup.name,
    resourceProviderNamespace: "Microsoft.Storage",
    resourceType: "storageAccounts",
    apiVersion: "2022-09-01",
    // This property is required even when empty.
    parentResourcePath: "",

    // These properties are optional and are defined on the generic resource because many resources have a kind or a SKU.
    kind: "StorageV2",
    sku: {
        name: "Standard_LRS",
    },

    // This is a generic property bag for all other properties of the targeted resource.
    properties: {
        allowBlobPublicAccess: false,
    },
});
```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi
import pulumi_azure_native as azure_native

resource_group = azure_native.resources.ResourceGroup("resourceGroup")

# Create a Storage Account
sa = azure_native.storage.StorageAccount("sa",
    resource_group_name=resource_group.name,
    sku={
        "name": azure_native.storage.SkuName.STANDARD_LRS,
    },
    kind=azure_native.storage.Kind.STORAGE_V2)

# Create another Storage Account using the generic resource
generic = azure_native.resources.Resource("generic",
    resource_group_name=resource_group.name,
    resource_provider_namespace="Microsoft.Storage",
    resource_type="storageAccounts",
    api_version="2022-09-01",
    # This property is required even when empty.
    parent_resource_path="",

    # These properties are optional and are defined on the generic resource
    # because many resources have a kind or a SKU.
    kind="StorageV2",
    sku={
        "name": "Standard_LRS",
    },

    # This is a generic property bag for all other properties of the targeted resource.
    properties={
        "allowBlobPublicAccess": False,
    })
```

{{% /choosable %}}

{{% choosable language csharp %}}

```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using AzureNative = Pulumi.AzureNative;

return await Deployment.RunAsync(() => 
{
    var resourceGroup = new AzureNative.Resources.ResourceGroup("resourceGroup");

    // Create a Storage Account
    var sa = new AzureNative.Storage.StorageAccount("sa", new()
    {
        ResourceGroupName = resourceGroup.Name,
        Sku = new AzureNative.Storage.Inputs.SkuArgs
        {
            Name = AzureNative.Storage.SkuName.Standard_LRS,
        },
        Kind = AzureNative.Storage.Kind.StorageV2,
    });

    // Create another Storage Account using the generic resource
    var generic = new AzureNative.Resources.Resource("generic", new()
    {
        ResourceGroupName = resourceGroup.Name,
        ResourceProviderNamespace = "Microsoft.Storage",
        ResourceType = "storageAccounts",
        ApiVersion = "2022-09-01",
        // This property is required even when empty.
        ParentResourcePath = "",

        // These properties are optional and are defined on the generic
        // resource because many resources have a kind or a SKU.
        Kind = "StorageV2",
        Sku = new AzureNative.Resources.Inputs.SkuArgs
        {
            Name = "Standard_LRS",
        },

        // This is a generic property bag for all other properties of the
        // targeted resource.
        Properties = new Dictionary<string, object?>
        {
            ["allowBlobPublicAccess"] = false,
        },
    });

});
```

{{% /choosable %}}

{{% choosable language go %}}

```go
package main

import (
	resources "github.com/pulumi/pulumi-azure-native-sdk/resources/v2"
	storage "github.com/pulumi/pulumi-azure-native-sdk/storage/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		resourceGroup, err := resources.NewResourceGroup(ctx, "resourceGroup", nil)
		if err != nil {
			return err
		}

    // Create a Storage Account
		_, err = storage.NewStorageAccount(ctx, "sa", &storage.StorageAccountArgs{
			ResourceGroupName: resourceGroup.Name,
			Sku: &storage.SkuArgs{
				Name: pulumi.String(storage.SkuName_Standard_LRS),
			},
			Kind: pulumi.String(storage.KindStorageV2),
		})
		if err != nil {
			return err
		}

    // Create another Storage Account using the generic resource
		_, err = resources.NewResource(ctx, "generic", &resources.ResourceArgs{
			ResourceGroupName:         resourceGroup.Name,
			ResourceProviderNamespace: pulumi.String("Microsoft.Storage"),
			ResourceType:              pulumi.String("storageAccounts"),
			ApiVersion:                pulumi.String("2022-09-01"),
      // This property is required even when empty.
			ParentResourcePath:        pulumi.String(""),

      // These properties are optional and are defined on the generic
      // resource because many resources have a kind or a SKU.
			Kind:                      pulumi.String("StorageV2"),
			Sku: &resources.SkuArgs{
				Name: pulumi.String("Standard_LRS"),
			},

      // This is a generic property bag for all other properties of the
      // targeted resource.
			Properties: pulumi.Any(map[string]interface{}{
				"allowBlobPublicAccess": false,
			}),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
```

{{% /choosable %}}

{{% choosable language java %}}

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.azurenative.resources.ResourceGroup;
import com.pulumi.azurenative.storage.StorageAccount;
import com.pulumi.azurenative.storage.StorageAccountArgs;
import com.pulumi.azurenative.storage.inputs.SkuArgs;
import com.pulumi.azurenative.resources.Resource;
import com.pulumi.azurenative.resources.ResourceArgs;
import java.util.Map;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        var resourceGroup = new ResourceGroup("resourceGroup");

        // Create a Storage Account
        var sa = new StorageAccount("sa", StorageAccountArgs.builder()
            .resourceGroupName(resourceGroup.name())
            .sku(SkuArgs.builder()
                .name("Standard_LRS")
                .build())
            .kind("StorageV2")
            .build());

        // Create another Storage Account using the generic resource
        var generic = new Resource("generic", ResourceArgs.builder()
            .resourceGroupName(resourceGroup.name())
            .resourceProviderNamespace("Microsoft.Storage")
            .resourceType("storageAccounts")
            .apiVersion("2022-09-01")
            // This property is required even when empty.
            .parentResourcePath("")

            // These properties are optional and are defined on the generic
            // resource because many resources have a kind or a SKU.
            .kind("StorageV2")
            .sku(com.pulumi.azurenative.resources.inputs.SkuArgs.builder()
                .name("Standard_LRS")
                .build())

            // This is a generic property bag for all other properties of the
            // targeted resource.
            .properties(Map.of("allowBlobPublicAccess", false))
            .build());
    }
}
```

{{% /choosable %}}


{{% choosable language yaml %}}

```yaml
resources:
  resourceGroup:
    type: azure-native:resources:ResourceGroup

  # Create a Storage Account
  sa:
    type: azure-native:storage:StorageAccount
    properties:
      resourceGroupName: ${resourceGroup.name}
      sku:
        name: Standard_LRS
      kind: StorageV2

  # Create another Storage Account using the generic resource
  generic:
    type: azure-native:resources:Resource
    properties:
      # These properties are required.
      resourceGroupName: ${resourceGroup.name}
      resourceProviderNamespace: Microsoft.Storage
      resourceType: storageAccounts
      apiVersion: "2022-09-01"
      # This property is required even when empty.
      parentResourcePath: ""

      # These properties are optional and are defined on the generic resource
      # because many resources have a kind or a SKU.
      kind: StorageV2
      sku:
        name: Standard_LRS

      # This is a generic property bag for all other properties of the targeted resource.
      properties:
        allowBlobPublicAccess: false
```

{{% /choosable %}}

{{< /chooser >}}