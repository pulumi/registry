---
title: Amazon EKS
meta_desc: Learn how you can use Pulumi's Amazon EKS component to set up an AWS Elastic Kubernetes Service cluster in just a few lines of code.
layout: overview
---

[![Deploy](https://get.pulumi.com/new/button.svg)](https://app.pulumi.com/new?template=https:%2F%2Fgithub.com%2Fpulumi%2Fexamples%2Ftree%2Fmaster%2Faws-ts-eks)

Amazon EKS is a Pulumi Component that creates and manages the resources necessary to run an [EKS Kubernetes cluster](https://aws.amazon.com/eks/) in AWS. Use this component to quickly set up an EKS cluster in just a few lines of code.

Amazon EKS must be configured with credentials to deploy and update resources in AWS; see the [Installation & Configuration page]({{<relref "./installation-configuration">}}) for instructions.

## Example

{{< chooser language "javascript,typescript,python,go,csharp" >}}

{{% choosable language javascript %}}

```javascript
import * as eks from "@pulumi/eks";

// Create an EKS cluster with the default configuration.
const cluster = new eks.Cluster("eks-cluster");

// Export the cluster's kubeconfig.
export const kubeconfig = cluster.kubeconfig;
```

{{% /choosable %}}
{{% choosable language typescript %}}

```typescript
import * as eks from "@pulumi/eks";

// Create an EKS cluster with the default configuration.
const cluster = new eks.Cluster("eks-cluster");

// Export the cluster's kubeconfig.
export const kubeconfig = cluster.kubeconfig;
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi
import pulumi_eks as eks

# Create an EKS cluster with the default configuration.
cluster = eks.Cluster("eks-cluster")

# Export the cluster's kubeconfig.
pulumi.export("kubeconfig", cluster.kubeconfig)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
package main

import (
	"github.com/pulumi/pulumi-eks/sdk/go/eks"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create an EKS cluster with default settings.
		cluster, err := eks.NewCluster(ctx, "eks-cluster", nil)
		if err != nil {
			return err
		}

		// Export the cluster's kubeconfig.
		ctx.Export("kubeconfig", cluster.Kubeconfig)
		return nil
	})
}
```

{{% /choosable %}}
{{% choosable language csharp %}}

```csharp
using System.Threading.Tasks;
using Pulumi;
using Pulumi.Eks;

class Program
{
    [Output("kubeconfig")]
    public Output<object> Kubeconfig { get; set; }

    static Task Main() =>
        Deployment.Run(() => {
            // Create an EKS cluster with default settings.
            var cluster = new Cluster("eks-cluster");

            // Export the cluster's kubeconfig.
            Kubeconfig = cluster.Kubeconfig;
        });
}
```

{{% /choosable %}}

{{< /chooser >}}

## SDK packages

The Amazon EKS provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/eks`](https://www.npmjs.com/package/@pulumi/eks)
* Python: [`pulumi_eks`](https://pypi.org/project/pulumi-eks//)
* Go: [`github.com/pulumi/pulumi-eks/sdk/go/eks`](https://github.com/pulumi/pulumi-eks)
* .NET: [`Pulumi.Eks`](https://www.nuget.org/packages/Pulumi.Eks)
