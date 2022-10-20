---
title: Exoscale
meta_desc: Provides an overview of the Exoscale Provider for Pulumi.
layout: overview
---

The Exoscale provider for Pulumi can be used to provision any of the cloud resources available in [Exoscale](https://www.exoscale.com/).
The Exoscale provider must be configured with credentials to deploy and update resources in Exoscale.

## Example

{{< chooser language "typescript,python,go,csharp" >}}
{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as exoscale from "@pulumiverse/exoscale"


let cluster = new exoscale.SKSCluster("cluster", {
  zone: "ch-gva-2",
  name: "my-sks-cluster"
})

export const endpoint = cluster.endpoint
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumiverse_exoscale as exoscale

cluster = exoscale.SKSCluster("cluster",
    zone="ch-gva-2",
    name="my-sks-cluster",
)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-exoscale/sdk/go/exoscale"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cluster, err := exoscale.NewSKSCluster(ctx, "cluster", &exoscale.SKSClusterArgs{
            Zone:   pulumi.String("ch-gva-2""),
			Name:   pulumi.String("my-sks-cluster"),
		})
		if err != nil {
			return err
		}
		ctx.Export("endpoint", cluster.Endpoint)
		return nil
	})
}
```

{{% /choosable %}}

{{< /chooser >}}
