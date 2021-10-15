---
title: AWS Miniflux
meta_desc: Easily create a Miniflux installation on AWS using Pulumi.
layout: overview
---

Easily create an installation of Miniflux, the excellent open-source RSS server, using Pulumi.

## Example

{{< chooser language "typescript,python,go,csharp" >}}

{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as miniflux from "@pulumi/aws-miniflux";

const config = new pulumi.Config();
const adminPassword = config.requireSecret("adminPassword");
const dbPassword = config.requireSecret("adminPassword");

// Create a new Miniflux service.
const service = new miniflux.MinifluxService("service", {
    adminPassword,
    dbPassword,
});

// Export the URL of the service.
export const endpoint = pulumi.interpolate`http://${service.endpoint}`;
```

{{% /choosable %}}
{{% choosable language python %}}

```py
import pulumi
from pulumi_aws import s3
from pulumi_aws_miniflux import miniflux_service

config = pulumi.Config();
admin_password = config.get_secret("adminPassword")
db_password = config.get_secret("dbPassword")

# Create a new Miniflux service.
service = miniflux_service.MinifluxService("service",
        admin_password = admin_password,
        db_password = db_password
    )

# Export the URL of the service.
pulumi.export("endpoint", service.endpoint)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
package main

import (
	"github.com/pulumi/pulumi-aws-miniflux/sdk/go/miniflux"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		conf := config.New(ctx, "")
		adminPassword := conf.RequireSecret("adminPassword")
		dbPassword := conf.RequireSecret("dbPassword")

		// Create a new Miniflux service.
		service, err := miniflux.NewMinifluxService(ctx, "service", &miniflux.MinifluxServiceArgs{
			AdminPassword: adminPassword,
			DbPassword:    dbPassword,
		})
		if err != nil {
			return nil
		}

		// Export the URL of the service.
		ctx.Export("endpoint", pulumi.Sprintf("http://%s", service.Endpoint))
		return nil
	})
}
```

{{% /choosable %}}

{{% choosable language csharp %}}

```csharp
using Pulumi;
using Pulumi.Aws.S3;
using Pulumi.AwsMiniflux;

class MyStack : Stack
{
    public MyStack()
    {
        var config = new Pulumi.Config();
        var adminPassword = config.RequireSecret("adminPassword");
        var dbPassword = config.RequireSecret("dbPassword");

        // Create a new Miniflux service.
        var service = new Pulumi.AwsMiniflux.MinifluxService("service", new Pulumi.AwsMiniflux.MinifluxServiceArgs{
            AdminPassword = adminPassword,
            DbPassword = dbPassword,
        });

        // Export the URL of the service.
        this.Endpoint = Output.Format($"http://{service.Endpoint}");
    }

    [Output]
    public Output<string> Endpoint { get; set; }
}
```

{{% /choosable %}}

{{< /chooser >}}

## SDK packages

The AWS Miniflux component is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/aws-miniflux`](https://www.npmjs.com/package/@pulumi/aws-miniflux)
* Python: [`pulumi-aws-miniflux`](https://pypi.org/project/pulumi-aws-miniflux/)
* Go: [`github.com/pulumi/pulumi-aws-miniflux/sdk/go/miniflux`](https://github.com/pulumi/pulumi-aws-miniflux)
* .NET: [`Pulumi.AwsMiniflux`](https://www.nuget.org/packages/Pulumi.AwsMiniflux)
