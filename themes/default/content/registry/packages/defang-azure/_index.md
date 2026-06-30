---
# WARNING: this file was fetched from https://raw.githubusercontent.com/DefangLabs/pulumi-defang/v2.1.1/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/DefangLabs/pulumi-defang/blob/v2.1.1/docs/_index.md
title: Defang Azure Provider
meta_desc: Take your app from Docker Compose to a secure and scalable Azure deployment with Pulumi.
layout: package
---
# Defang Azure Pulumi Provider

![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/DefangLabs/pulumi-defang?label=Version)

The Pulumi Provider for [Defang](https://defang.io) — Take your app from Docker Compose to a secure and scalable Azure deployment with Pulumi.

## Example usage

You can find complete working TypeScript, Python, Go, .NET, and Yaml code samples in the [`./examples`](https://github.com/DefangLabs/pulumi-defang/tree/main/examples) directory, and some example snippets below:

{{< chooser language "typescript,python,go,dotnet,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as defang_azure from "@defang-io/pulumi-defang-azure";

const azureDemo = new defang_azure.Project("azure-demo", {services: {
    app: {
        image: "nginx",
        ports: [{
            target: 80,
            mode: "ingress",
            appProtocol: "http",
        }],
    },
}});
export const endpoints = azureDemo.endpoints;
```

{{% /choosable %}}

{{% choosable language python %}}
```python
import pulumi
import pulumi_defang_azure as defang_azure

azure_demo = defang_azure.Project("azure-demo", services={
    "app": {
        "image": "nginx",
        "ports": [{
            "target": 80,
            "mode": "ingress",
            "app_protocol": "http",
        }],
    },
})
pulumi.export("endpoints", azure_demo.endpoints)
```

{{% /choosable %}}

{{% choosable language go %}}
```go
package main

import (
	defangazure "github.com/DefangLabs/pulumi-defang/sdk/v2/go/defang-azure"
	"github.com/DefangLabs/pulumi-defang/sdk/v2/go/defang-azure/compose"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		azureDemo, err := defangazure.NewProject(ctx, "azure-demo", &defangazure.ProjectArgs{
			Services: compose.ServiceConfigMap{
				"app": &compose.ServiceConfigArgs{
					Image: pulumi.String("nginx"),
					Ports: compose.ServicePortConfigArray{
						&compose.ServicePortConfigArgs{
							Target:      pulumi.Int(80),
							Mode:        pulumi.String("ingress"),
							AppProtocol: pulumi.String("http"),
						},
					},
				},
			},
		})
		if err != nil {
			return err
		}
		ctx.Export("endpoints", azureDemo.Endpoints)
		return nil
	})
}
```

{{% /choosable %}}

{{% choosable language dotnet %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using DefangAzure = DefangLabs.DefangAzure;

return await Deployment.RunAsync(() =>
{
    var azureDemo = new DefangAzure.Project("azure-demo", new()
    {
        Services =
        {
            { "app", new DefangAzure.Compose.Inputs.ServiceConfigArgs
            {
                Image = "nginx",
                Ports = new[]
                {
                    new DefangAzure.Compose.Inputs.ServicePortConfigArgs
                    {
                        Target = 80,
                        Mode = "ingress",
                        AppProtocol = "http",
                    },
                },
            } },
        },
    });

    return new Dictionary<string, object?>
    {
        ["endpoints"] = azureDemo.Endpoints,
    };
});
```

{{% /choosable %}}

{{% choosable language yaml %}}
```yaml
name: defang-azure
runtime: yaml
description: Example using defang-azure to deploy services to Azure

plugins:
  providers:
    - name: defang-azure
      path: ../../bin

resources:
  azure-demo:
    type: defang-azure:index:Project
    properties:
      services:
        app:
          image: nginx
          ports:
            - target: 80
              mode: ingress
              appProtocol: http

outputs:
  endpoints: ${azure-demo.endpoints}
```

{{% /choosable %}}
{{< /chooser >}}

## Installation and Configuration

See our [Installation and Configuration](https://www.pulumi.com/registry/packages/defang-azure/installation-configuration/) docs

## Development

See the [Contributing](https://github.com/DefangLabs/pulumi-defang/blob/main/CONTRIBUTING.md) doc.
