---
title: GCP Global CloudRun
meta_desc: Easily create a globally load balanced CloudRun application in Google Cloud.
layout: overview
---

Easily create a globally load-balanced Google Cloud Run application using this component, which is available in all Pulumi languages.

Example:

{{< chooser language "typescript,python,go" >}}

{{% choosable language typescript %}}

```ts
import * as pulumi from "@pulumi/pulumi";
import * as cloudrun from "@pulumi/gcp-global-cloudrun";

const conf = new pulumi.Config()
const project = conf.require("project")

const deployment = new cloudrun.Deployment("my-sample-deployment", {
    projectId: project,

    imageName: "gcr.io/ahmetb-public/zoneprinter",
    serviceName: "demo-service-ts"
});

export const ip = deployment.ipAddress;
```

{{% /choosable %}}
{{% choosable language python %}}

```py
import pulumi
import pulumi_gcp_global_cloudrun as cloudrun

config = pulumi.Config()
project = config.require("project")

deployment = cloudrun.Deployment("my-sample-deployment",
                                 project_id=project,
                                 image_name="gcr.io/ahmetb-public/zoneprinter",
                                 service_name="demo-service-py")

pulumi.export('ip', deployment.ip_address)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
	cloudrun "github.com/pulumi/pulumi-gcp-global-cloudrun/sdk/go/gcp"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		c := config.New(ctx, "")
		project := c.Require("project")

		deployment, err := cloudrun.NewDeployment(ctx, "demo-deployment-go", &cloudrun.DeploymentArgs{
			ImageName:   pulumi.String("gcr.io/ahmetb-public/zoneprinter"),
			ServiceName: "demo-service-ts",
			ProjectId:   project,
		})
		if err != nil {
			return err
		}

		ctx.Export("ip", deployment.IpAddress)

		return nil
	})
}
```

{{% /choosable %}}

{{< /chooser >}}

## SDK packages

The AWS API Gateway provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/aws-gcp-global-cloudrun`](https://www.npmjs.com/package/@pulumi/gcp-global-cloudrun)
* Python: [`pulumi-gcp-global-cloudrun`](https://pypi.org/project/pulumi-gcp-global-cloudrun/)
* Go: [`github.com/pulumi/pulumi-gcp-global-cloudrun/sdk/go/gcp`](https://github.com/pulumi/pulumi-gcp-global-cloudrun)
* .NET: [`Pulumi.GcpGlobalCloudRun`](https://www.nuget.org/packages/Pulumi.GcpGlobalCloudRun)
