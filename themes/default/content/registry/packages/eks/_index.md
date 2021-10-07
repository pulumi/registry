---
title: Amazon EKS
meta_desc: This page provides an overview of the Pulumi Amazon EKS component.
layout: overview
---

The Pulumi EKS library provides a Pulumi component that creates and manages the resources necessary to run an [EKS Kubernetes cluster](https://aws.amazon.com/eks/) in AWS.
The Amazon EKS component must be configured with credentials to deploy and update the resources.

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

## Libraries

The following packages are available in packager managers:

* JavaScript/TypeScript: [`@pulumi/f5bigip`](https://www.npmjs.com/package/@pulumi/eks)
* Python: [`pulumi_eks`](https://pypi.org/project/pulumi-eks//)
* Go: [`github.com/pulumi/pulumi-eks/sdk/go/eks`](https://github.com/pulumi/pulumi-eks)
* .NET: [`Pulumi.Eks`](https://www.nuget.org/packages/Pulumi.Eks)

The Amazon EKS component is open source and available in the [pulumi/pulumi-eks](https://github.com/pulumi/pulumi-eks) repo.
