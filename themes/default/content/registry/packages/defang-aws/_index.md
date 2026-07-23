---
# WARNING: this file was fetched from https://raw.githubusercontent.com/DefangLabs/pulumi-defang/v2.5.2/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/DefangLabs/pulumi-defang/blob/v2.5.2/docs/_index.md
title: Defang Providers
meta_desc: Take your app from Docker Compose to a secure and scalable cloud deployment with Pulumi.
layout: package
---
# Defang Pulumi Providers

![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/DefangLabs/pulumi-defang?label=Version)

The Pulumi Providers for [Defang](https://defang.io) — Take your app from Docker Compose to a secure and scalable cloud deployment with Pulumi.

Defang ships one provider package per cloud, all with the same Compose-shaped inputs:

- **`defang-aws`** — deploys to AWS (ECS Fargate, ALB, RDS, ElastiCache, …)
- **`defang-gcp`** — deploys to Google Cloud (Cloud Run / Compute Engine, Cloud SQL, Memorystore, …)
- **`defang-azure`** — deploys to Azure (Container Apps, Azure Database, Azure Cache, …)

The examples below use `defang-aws`; swap in the `defang-gcp` or `defang-azure` package of your language to target another cloud — the `Project` inputs are the same.

## Example usage

You can find complete working TypeScript, Python, Go, .NET, and Yaml code samples for every cloud in the [`./examples`](https://github.com/DefangLabs/pulumi-defang/tree/main/examples) directory, and some example snippets below:

{{< chooser language "typescript,python,go,dotnet,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as defang_aws from "@defang-io/pulumi-defang-aws";

const awsDemo = new defang_aws.Project("aws-demo", {
    services: {
        app: {
            image: "nginx",
            ports: [{
                target: 80,
                mode: "ingress",
                appProtocol: "http",
            }],
        },
    },
});
export const endpoints = awsDemo.endpoints;
```

{{% /choosable %}}

{{% choosable language python %}}
```python
import pulumi
import pulumi_defang_aws as defang_aws

aws_demo = defang_aws.Project("aws-demo", services={
    "app": {
        "image": "nginx",
        "ports": [{
            "target": 80,
            "mode": "ingress",
            "app_protocol": "http",
        }],
    },
})
pulumi.export("endpoints", aws_demo.endpoints)
```

{{% /choosable %}}

{{% choosable language go %}}
```go
package main

import (
	defangaws "github.com/DefangLabs/pulumi-defang/sdk/v2/go/defang-aws"
	"github.com/DefangLabs/pulumi-defang/sdk/v2/go/defang-aws/compose"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		awsDemo, err := defangaws.NewProject(ctx, "aws-demo", &defangaws.ProjectArgs{
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
		ctx.Export("endpoints", awsDemo.Endpoints)
		return nil
	})
}
```

{{% /choosable %}}

{{% choosable language dotnet %}}
```csharp
using System.Collections.Generic;
using Pulumi;
using DefangAws = DefangLabs.DefangAws;

return await Deployment.RunAsync(() =>
{
    var awsDemo = new DefangAws.Project("aws-demo", new()
    {
        Services =
        {
            { "app", new DefangAws.Compose.Inputs.ServiceConfigArgs
            {
                Image = "nginx",
                Ports = new[]
                {
                    new DefangAws.Compose.Inputs.ServicePortConfigArgs
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
        ["endpoints"] = awsDemo.Endpoints,
    };
});
```

{{% /choosable %}}

{{% choosable language yaml %}}
```yaml
name: defang-aws-example
runtime: yaml
description: Example using defang-aws to deploy services to AWS

resources:
  aws-demo:
    type: defang-aws:index:Project
    properties:
      services:
        app:
          image: nginx
          ports:
            - target: 80
              mode: ingress
              appProtocol: http

outputs:
  endpoints: ${aws-demo.endpoints}
```

{{% /choosable %}}
{{< /chooser >}}

## Installation and Configuration

See our [Installation and Configuration](./installation-configuration/) docs

## Development

See the [Contributing](https://github.com/DefangLabs/pulumi-defang/blob/main/CONTRIBUTING.md) doc.
