---
title: Azure Quickstart ACR Geo Replicated
meta_desc: Use Pulumi's Component for creating Azure ACR Geo Replicated Registries using infrastructure as code.
layout: overview
---

Easily create Azure ACR Registries that are replicated across Azure locations as a package available in all Pulumi languages.

Example:

{{< chooser language "typescript,python,go" >}}

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
export const underlying_registry_id = registry.registry.id;
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi
from pulumi_azure_native import resources
from pulumi_azure_quickstart_acr_geo_replication import Registry

# Create an Azure Resource Group
resource_group = resources.ResourceGroup('resource_group')

# Create an Azure resource (Storage Account)
registry = ReplicatedRegistry('registry',
                              name="registry",
                              resource_group_name=resource_group.name,
                              replication_location="westus")

pulumi.export("login", registry.login_server_out)
pulumi.export("underlying_registry_id", registry.registry.Id)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
package main

import (
	"github.com/pulumi/pulumi-azure-native/sdk/go/azure/resources"
	acr "github.com/pulumi/pulumi-azure-quickstart-acr-geo-replication/sdk/go/azure"
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

{{< /chooser >}}
