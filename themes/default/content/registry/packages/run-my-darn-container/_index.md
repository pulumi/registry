---
title: Run My Darn Container
meta_desc: Deploy your container to the 3 major cloud providers with just a few lines of code.
layout: overview
---

Run My Darn Container (RDC) is the multi cloud solution you've been waiting for. Deploy your container to the 3 major cloud providers with just a few lines of code. RDC allows you to ship a built Docker Image into the available container running engines for the major cloud providers. RDC is designed to have absolutely no configurability whatsoever, if you want to configure your application or image in any way at all, you'll want to hardcode it into your Docker image.

## Example

{{< chooser language "typescript,python,go,csharp" >}}

{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as rdc from "@pulumi/run-my-darn-container";

const awsapp = new rdc.AWSInstance("aws-example", {
    image: "public.ecr.aws/aws-containers/hello-app-runner:latest",
    port: 8000,
});

const azureapp = new rdc.AzureInstance("azure-example", {
    image: "mcr.microsoft.com/azuredocs/aci-helloworld",
    port: 80,
});

const gcpApp = new rdc.GCPInstance("gcp-example", {
    image: "gcr.io/cloudrun/hello:latest",
    port: 8000,
});

const urls = [];

urls.push(
    gcpApp.url,
    azureapp.url,
    awsapp.url
);

export { urls };
```

{{% /choosable %}}
{{% choosable language python %}}

```py
import pulumi
import pulumi_run_my_darn_container as rdc

aws = rdc.aws_instance(
    "aws-example",
    image="public.ecr.aws/aws-containers/hello-app-runner:latest",
    port=8000
)

azure = rdc.azure_instance(
    "azure-example",
    image="mcr.microsoft.com/azuredocs/aci-helloworld",
    port=8000,
)

gcp = rdc.gcp_instance(
    "gcp-example",
    image="gcr.io/cloudrun/hello:latest",
    port=80
)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
package main

import (
	container "github.com/pulumi/pulumi-run-my-darn-container/sdk/go/container"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		aws, err := container.NewAWSInstance(ctx, "aws-example", &container.AWSInstanceArgs{
			Image: pulumi.String("public.ecr.aws/aws-containers/hello-app-runner:latest"),
			Port:  pulumi.Int(8000),
		})
		if err != nil {
			return err
		}

		ctx.Export("AWSUrl", aws.Url)
		return nil
	})
}
```

{{% /choosable %}}

{{% choosable language csharp %}}

```csharp
using Pulumi;
using Pulumi.RunMyDarnContainer;

class MyStack : Stack
{
    public MyStack()
    {
        // Create a new AWS container instance.
        var aws = new Pulumi.RunMyDarnContainer.AWSInstance("example", new Pulumi.RunMyDarnContainer.AWSInstanceArgs{
            Image = "public.ecr.aws/aws-containers/hello-app-runner:latest",
			Port = 8000,
        });

        this.AWSUrl = aws.Url;
    }

    [Output]
    public Output<string> AWSUrl { get; set; }
}
```

{{% /choosable %}}

{{< /chooser >}}

## SDK packages

The Run My Darn Container component is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/run-my-darn-container`](https://www.npmjs.com/package/@pulumi/run-my-darn-container)
* Python: [`pulumi-run-my-darn-container`](https://pypi.org/project/pulumi-run-my-darn-container/)
* Go: [`github.com/pulumi/pulumi-run-my-darn-container/sdk/go/container`](github.com/pulumi/pulumi-run-my-darn-container)
* .NET: [`Pulumi.RunMyDarnContainer`](https://www.nuget.org/packages/Pulumi.RunMyDarnContainer)
