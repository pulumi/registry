---
title: Amazon EKS
meta_desc: Learn how you can use Pulumi's Amazon EKS component to set up an AWS Elastic Kubernetes Service cluster in just a few lines of code.
layout: package
---

Amazon EKS is a Pulumi Component that creates and manages the resources necessary to run an [EKS Kubernetes cluster](https://aws.amazon.com/eks/) in AWS. Use this component to quickly set up an EKS cluster in just a few lines of code. This component exposes the [Crosswalk for AWS](../../../docs/guides/crosswalk/aws/) functionality documented in the [Pulumi Amazon EKS guide](../../../docs/guides/crosswalk/aws/eks/).

Amazon EKS must be configured with credentials to deploy and update resources in AWS; see [Installation & Configuration](./installation-configuration) for instructions.

## Example

<!---
javascript removed
--->
{{< chooser language "typescript,python,go,csharp,java,yaml" >}}

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
using System.Collections.Generic;
using Pulumi;
using Pulumi.Eks;

await Deployment.RunAsync(() =>
{
  // Create an EKS cluster with default settings.
  var cluster = new Cluster("eks-cluster");

  // Export the cluster's kubeconfig.
  return new Dictionary<string, object?>
  {
    ["kubeconfig"] = cluster.Kubeconfig
  };
});
```

{{% /choosable %}}

{{% choosable language java %}}

```java
import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.eks.Cluster;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    private static void stack(Context ctx) {
		final var cluster = new Cluster("eks-cluster");
		ctx.export("kubeconfig", cluster.kubeconfig());
	}
}
```

{{% /choosable %}}
{{% choosable language yaml %}}

```yaml
resources:
  eks-cluster:
    type: eks:Cluster
outputs:
  kubeconfig: ${cluster.kubeconfig}
```

{{% /choosable %}}

{{< /chooser >}}
