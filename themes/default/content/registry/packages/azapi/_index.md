---
# WARNING: this file was fetched from https://raw.githubusercontent.com/dirien/pulumi-azapi/v1.12.2/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
title: AzAPI
meta_desc: Provides an overview of the AzAPI Provider for Pulumi.
layout: package
---

The AzAPI Resource Provider lets you manage [AzAPI](https://learn.microsoft.com/en-us/azure/developer/terraform/overview-azapi-provider) resources.


## Example

{{< chooser language "typescript,python,go,csharp" >}}


{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as azure from "@pulumi/azure";
import * as azapi from "@ediri/azapi"
import {jsonStringify} from "@pulumi/pulumi";

const location = "West Europe";

// Create an Azure Resource Group
const resourceGroup = new azure.core.ResourceGroup("resource-group", {
    location: location,
});

const appServicePlan = new azapi.Resource("app-service-plan", {
    type: "Microsoft.Web/serverfarms@2020-06-01",
    name: "app-service-plan",
    parentId: resourceGroup.id,
    ignoreCasing: true,
    body: jsonStringify(
        {
            sku: {
                name: "F1",
            },
            properties: {
                reserved: true,
            },
            kind: "linux",
            location: location,
        }
    ),
    responseExportValues: [
        "*"
    ]
})

const site = new azapi.Resource("site", {
    type: "Microsoft.Web/sites@2021-02-01",
    parentId: resourceGroup.id,
    ignoreCasing: true,
    body: jsonStringify(
        {
            properties: {
                serverFarmId: appServicePlan.id,
                siteConfig: {
                    linuxFxVersion: "node|14-lts",
                },
                httpsOnly: true,
            },
            kind: "app,linux",
            location: location,
        }
    ),
    identity: {
        type: "SystemAssigned"
    },
    responseExportValues: [
        "*"
    ]
})

new azapi.UpdateResource("source-control", {
    type: "Microsoft.Web/sites/sourcecontrols@2022-03-01",
    name: "web",
    parentId: site.id,
    ignoreCasing: true,
    body: jsonStringify(
        {
            properties: {
                repoUrl: "https://github.com/Azure-Samples/nodejs-docs-hello-world",
                branch: "main",
                isManualIntegration: true,
            }
        })
}, {
    dependsOn: [
        site
    ]
})

export const url = pulumi.interpolate`https://${site.name}.azurewebsites.net`;
```

{{% /choosable %}}

{{% choosable language python %}}

```python
"""An Azure Python Pulumi program"""

import pulumi
from pulumi_azure import core
from ediri_azapi import Resource, ResourceArgs, ResourceIdentityArgs, UpdateResource, UpdateResourceArgs

location = 'West Europe'

resource_group = core.ResourceGroup('resource-group', location=location)

app_service_plan = Resource('app-service-plan', ResourceArgs(
    type="Microsoft.Web/serverfarms@2020-06-01",
    name="app-service-plan",
    parent_id=resource_group.id,
    ignore_casing=True,
    body=pulumi.Output.json_dumps({
        "sku": {
            "name": "F1",
        },
        "properties": {
            "reserved": True,
        },
        "kind": "linux",
        "location": location,
    }),
    response_export_values=[
        "*"
    ]
))

site = Resource('site', ResourceArgs(
    type="Microsoft.Web/sites@2021-02-01",
    parent_id=resource_group.id,
    ignore_casing=True,
    body=pulumi.Output.json_dumps({
        "properties": {
            "serverFarmId": app_service_plan.id,
            "siteConfig": {
                "linuxFxVersion": "node|14-lts",
            },
            "httpsOnly": True,
        },
        "kind": "app,linux",
        "location": location,
    }),
    identity=ResourceIdentityArgs(
        type="SystemAssigned"
    ),
    response_export_values=[
        "*"
    ]
))

UpdateResource('source-control', UpdateResourceArgs(
    type="Microsoft.Web/sites/sourcecontrols@2022-03-01",
    name="web",
    parent_id=site.id,
    ignore_casing=True,
    body=pulumi.Output.json_dumps({
        "properties": {
            "repoUrl": "https://github.com/Azure-Samples/nodejs-docs-hello-world",
            "branch": "main",
            "isManualIntegration": True,
        },
    }),
), depends_on=[site])

pulumi.export('url', site.name.apply(
    lambda name: f'https://{name}.azurewebsites.net'))
```

{{% /choosable %}}
{{% choosable language go %}}

```go
package main

import (
	"github.com/dirien/pulumi-azapi/go/azapi"
	"github.com/pulumi/pulumi-azure/sdk/v5/go/azure/core"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		location := pulumi.String("West Europe")

		resourceGroup, err := core.NewResourceGroup(ctx, "resource-group", &core.ResourceGroupArgs{
			Location: location,
		})
		if err != nil {
			return err
		}

		appServicePlan, err := azapi.NewResource(ctx, "app-service-plan", &azapi.ResourceArgs{
			Type:     pulumi.String("Microsoft.Web/serverfarms@2020-06-01"),
			Name:     pulumi.String("app-service-plan"),
			ParentId: resourceGroup.ID(),

			IgnoreCasing: pulumi.Bool(true),
			Body: pulumi.JSONMarshal(pulumi.Map{
				"sku": pulumi.Map{
					"name": pulumi.String("F1"),
				},
				"properties": pulumi.Map{
					"reserved": pulumi.Bool(true),
				},
				"kind":     pulumi.String("linux"),
				"location": location,
			}),
			ResponseExportValues: pulumi.StringArray{
				pulumi.String("*"),
			},
		})
		if err != nil {
			return err
		}

		site, err := azapi.NewResource(ctx, "site", &azapi.ResourceArgs{
			Type:         pulumi.String("Microsoft.Web/sites@2021-02-01"),
			ParentId:     resourceGroup.ID(),
			IgnoreCasing: pulumi.Bool(true),
			Body: pulumi.JSONMarshal(pulumi.Map{
				"properties": pulumi.Map{
					"serverFarmId": appServicePlan.ID(),
					"siteConfig": pulumi.Map{
						"linuxFxVersion": pulumi.String("node|14-lts"),
					},
					"httpsOnly": pulumi.Bool(true),
				},
				"kind":     pulumi.String("app,linux"),
				"location": location,
			}),
			Identity: &azapi.ResourceIdentityArgs{
				Type: pulumi.String("SystemAssigned"),
			},
			ResponseExportValues: pulumi.StringArray{
				pulumi.String("*"),
			},
		})
		if err != nil {
			return err
		}

		_, err = azapi.NewUpdateResource(ctx, "source-control", &azapi.UpdateResourceArgs{
			Type:         pulumi.String("Microsoft.Web/sites/sourcecontrols@2022-03-01"),
			Name:         pulumi.String("web"),
			ParentId:     site.ID(),
			IgnoreCasing: pulumi.Bool(true),
			Body: pulumi.JSONMarshal(pulumi.Map{
				"properties": pulumi.Map{
					"repoUrl":             pulumi.String("https://github.com/Azure-Samples/nodejs-docs-hello-world"),
					"branch":              pulumi.String("main"),
					"isManualIntegration": pulumi.Bool(true),
				},
			}),
		}, pulumi.DependsOn([]pulumi.Resource{site}))
		if err != nil {
			return err
		}
		ctx.Export("url", pulumi.Sprintf("https://%s.azurewebsites.net", site.Name))
		return nil
	})
}
```

{{% /choosable %}}

{{% choosable language csharp %}}

```csharp
using Pulumi;
using Azure = Pulumi.Azure;
using Azapi = ediri.Azapi;
using System.Collections.Generic;
using System.Text.Json;


return await Pulumi.Deployment.RunAsync(() =>
{
    var location = "West Europe";
    // Create an Azure Resource Group
    var resourceGroup = new Azure.Core.ResourceGroup("resource-group", new Azure.Core.ResourceGroupArgs
    {
        Location = location
    });

    var appServicePlan = new Azapi.Resource("app-service-plan", new Azapi.ResourceArgs
    {
        Type = "Microsoft.Web/serverfarms@2020-06-01",
        Name = "app-service-plan",
        ParentId = resourceGroup.Id,
        IgnoreCasing = true,
        Body=  JsonSerializer.Serialize(new
        {
            kind = "linux",
            location = location,
            sku = new
            {
                name = "F1",
            },
            properties = new
            {
                reserved = true,
            }
        }),
        ResponseExportValues = new List<string> { 
            "*" 
        }
    });

    var site = new Azapi.Resource("site", new Azapi.ResourceArgs
    {
        Type = "Microsoft.Web/sites@2021-02-01",
        ParentId = resourceGroup.Id,
        IgnoreCasing = true,
        Body=  appServicePlan.Id.Apply(id => JsonSerializer.Serialize(new
        {
            kind = "app,linux",
            location = location,
            properties = new
            {
                serverFarmId = id,
                siteConfig = new
                {
                    linuxFxVersion = "node|14-lts",
                },
                httpsOnly = true,
            }
        })),
        Identity = new Azapi.Inputs.ResourceIdentityArgs
        {
            Type = "SystemAssigned"
        },
        ResponseExportValues = new List<string> { 
            "*"
        }
    });

    new Azapi.UpdateResource("source-control", new Azapi.UpdateResourceArgs
    {
        Type = "Microsoft.Web/sites/sourcecontrols@2022-03-01",
        Name = "web",
        ParentId= site.Id,
        IgnoreCasing = true,
        Body = JsonSerializer.Serialize(new
        {
            properties = new
            {
                repoUrl = "https://github.com/Azure-Samples/nodejs-docs-hello-world",
                branch = "main",
                isManualIntegration = true,
            }
        })
    }, new CustomResourceOptions
    {
        DependsOn = { site }
    }); 

    // Export the primary key of the Storage Account
    return new Dictionary<string, object?>
    {
        ["url"] = Pulumi.Output.Format($"https://{site.Name}.azurewebsites.net")
    };
});
```

{{% /choosable %}}

{{< /chooser >}}
