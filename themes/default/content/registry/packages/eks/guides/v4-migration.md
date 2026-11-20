---
title: Amazon EKS v4 Migration Guide
meta_desc: How to upgrade from v3 to v4 of the Pulumi Amazon EKS package
layout: package
weight: 1
---

Pulumi EKS v4 moves the package to version 7 of the Pulumi AWS provider.
Apart from that, it has one breaking change: the `defaultAddonsToRemove` input of `Cluster` is removed.

## Upgrade the AWS provider to v7

Every SDK of EKS v4 depends on `pulumi-aws` 7.x, and every input and output that takes or returns an AWS resource or type is now a v7 type.
This includes the cluster's `instanceRoles`, `serviceRole`, and `clusterSecurityGroup`, and the node groups' launch templates and security groups.
Upgrade the AWS provider in the same change as EKS, and follow the [AWS 7.0 migration guide](/registry/packages/aws/guides/7-0-migration/) for any AWS resources your program creates directly.

In Go, both module paths change: import `github.com/pulumi/pulumi-eks/sdk/v4/go/eks` and `github.com/pulumi/pulumi-aws/sdk/v7/go/aws`.

## `defaultAddonsToRemove` is removed

`defaultAddonsToRemove` removed the `vpc-cni`, `coredns`, and `kube-proxy` addons that EKS installs when it creates a cluster.
It only took effect when the cluster was created, so removing it from your program doesn't change existing clusters.

To create a cluster without those addons in v4, set `bootstrapSelfManagedAddons` to `false`.
Like `defaultAddonsToRemove`, it only applies at creation: changing it on an existing cluster has no effect.
The `Cluster` component also installs its own managed versions of these addons by default, so turn those off too if you want a cluster with no addons:

{{< chooser language "typescript,python,go,csharp,java,yaml,hcl" >}}

{{% choosable language typescript %}}

```typescript
import * as eks from "@pulumi/eks";

const cluster = new eks.Cluster("cluster", {
    bootstrapSelfManagedAddons: false,
    useDefaultVpcCni: true,
    corednsAddonOptions: {
        enabled: false,
    },
    kubeProxyAddonOptions: {
        enabled: false,
    },
});
```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi_eks as eks

cluster = eks.Cluster("cluster",
    bootstrap_self_managed_addons=False,
    use_default_vpc_cni=True,
    coredns_addon_options={
        "enabled": False,
    },
    kube_proxy_addon_options={
        "enabled": False,
    })
```

{{% /choosable %}}

{{% choosable language go %}}

```go
package main

import (
	"github.com/pulumi/pulumi-eks/sdk/v4/go/eks"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := eks.NewCluster(ctx, "cluster", &eks.ClusterArgs{
			BootstrapSelfManagedAddons: pulumi.Bool(false),
			UseDefaultVpcCni:           pulumi.BoolRef(true),
			CorednsAddonOptions: &eks.CoreDnsAddonOptionsArgs{
				Enabled: pulumi.BoolRef(false),
			},
			KubeProxyAddonOptions: &eks.KubeProxyAddonOptionsArgs{
				Enabled: pulumi.BoolRef(false),
			},
		})
		return err
	})
}
```

{{% /choosable %}}

{{% choosable language csharp %}}

```csharp
using Pulumi;
using Eks = Pulumi.Eks;

return await Deployment.RunAsync(() =>
{
    var cluster = new Eks.Cluster("cluster", new()
    {
        BootstrapSelfManagedAddons = false,
        UseDefaultVpcCni = true,
        CorednsAddonOptions = new Eks.Inputs.CoreDnsAddonOptionsArgs
        {
            Enabled = false,
        },
        KubeProxyAddonOptions = new Eks.Inputs.KubeProxyAddonOptionsArgs
        {
            Enabled = false,
        },
    });
});
```

{{% /choosable %}}

{{% choosable language java %}}

```java
import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.eks.Cluster;
import com.pulumi.eks.ClusterArgs;
import com.pulumi.eks.inputs.CoreDnsAddonOptionsArgs;
import com.pulumi.eks.inputs.KubeProxyAddonOptionsArgs;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        var cluster = new Cluster("cluster", ClusterArgs.builder()
            .bootstrapSelfManagedAddons(false)
            .useDefaultVpcCni(true)
            .corednsAddonOptions(CoreDnsAddonOptionsArgs.builder()
                .enabled(false)
                .build())
            .kubeProxyAddonOptions(KubeProxyAddonOptionsArgs.builder()
                .enabled(false)
                .build())
            .build());
    }
}
```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
resources:
  cluster:
    type: eks:Cluster
    properties:
      bootstrapSelfManagedAddons: false
      useDefaultVpcCni: true
      corednsAddonOptions:
        enabled: false
      kubeProxyAddonOptions:
        enabled: false
```

{{% /choosable %}}

{{% choosable language hcl %}}

```hcl
terraform {
  required_providers {
    eks = {
      source = "pulumi/eks"
    }
  }
}

resource "eks_cluster" "cluster" {
  bootstrap_self_managed_addons = false
  use_default_vpc_cni           = true
  coredns_addon_options = {
    enabled = false
  }
  kube_proxy_addon_options = {
    enabled = false
  }
}
```

{{% /choosable %}}

{{< /chooser >}}

`bootstrapSelfManagedAddons` turns off the addons EKS installs itself, `useDefaultVpcCni` stops the component from creating its managed VPC CNI addon, and the two `enabled: false` options do the same for CoreDNS and kube-proxy.
