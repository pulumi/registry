---
title: "Azure Quickstart: Container Registry Geo-Replicated"
meta_desc: Create a geo-replicated Azure Container Registry instance that automatically synchronizes registry content across more than one Azure region.
layout: overview
---

<!-- Write a brief description of your package. -->

## Example

<!-- Provide a simple example of how to use your package, ideally in all languages. -->

{{< chooser language "typescript,python,go,csharp" >}}

{{% choosable language typescript %}}

```typescript
import * as acr from "@pulumi/azure-quickstart-acr-geo-replication";
import * as resources from "@pulumi/azure-native/resources";

const resourceGroup = new resources.ResourceGroup("resourceGroup");

const registry = new acr.ReplicatedRegistry("registry", {
    name: "registry",
    replicationLocation: "westus",
    resourceGroupName: resourceGroup.name,
});

export const login = registry.loginServer;
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi
from pulumi_azure_native import resources
from pulumi_azure_quickstart_acr_geo_replication import ReplicatedRegistry

# Create an Azure Resource Group
resource_group = resources.ResourceGroup('resource_group')

# Create an Azure resource (Storage Account)
registry = ReplicatedRegistry('registry',
    name="registry",
    resource_group_name=resource_group.name,
    replication_location="westus")

pulumi.export("login", registry.login_server)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
package main

import (
	"github.com/pulumi/pulumi-azure-native/sdk/go/azure/resources"
	"github.com/pulumi/pulumi-azure-quickstart-acr-geo-replication/sdk/go/acr"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		resourceGroup, err := resources.NewResourceGroup(ctx, "resourceGroup", nil)
		if err != nil {
			return err
		}

		registry, err := acr.NewReplicatedRegistry(ctx, "registry", &acr.ReplicatedRegistryArgs{
			Name:                "registry",
			ReplicationLocation: "westus",
			ResourceGroupName:   resourceGroup.Name,
		})

		if err != nil {
			return err
		}

		ctx.Export("login", registry.LoginServer)

		return nil
	})
}
```

{{% /choosable %}}
{{% choosable language csharp %}}

```csharp
using Pulumi;
using AzureNative = Pulumi.AzureNative;
using Acr = Pulumi.AzureQuickstartAcrGeoReplicated;

class MyStack : Stack
{
    [Output("login")] Output<string> Login { get; set; }

    public MyStack()
    {
        var resourceGroup = new AzureNative.Resources.ResourceGroup("resourceGroup");

        var registry = new Acr.ReplicatedRegistry(ctx, "bucket", new Acr.ReplicatedRegistryArgs{
            Name = "registry",
            ReplicationLocation = "westus",
            ResourceGroupName = resourceGroup.Name
        });

        this.Login = registry.LoginServer;
    }
}
```

{{% /choosable %}}

{{< /chooser >}}
