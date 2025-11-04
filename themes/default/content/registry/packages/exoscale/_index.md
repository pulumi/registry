---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumiverse/pulumi-exoscale/v0.66.1/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/pulumiverse/pulumi-exoscale/blob/v0.66.1/docs/_index.md
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
import pulumi
import pulumiverse_exoscale as exoscale

cluster = exoscale.SKSCluster("cluster",
                              zone="ch-gva-2",
                              name="my-sks-cluster",
                              )

pulumi.export("endpoint", cluster.endpoint)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-exoscale/sdk/go/exoscale"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cluster, err := exoscale.NewSKSCluster(ctx, "sks-cluster", &exoscale.SKSClusterArgs{
			Name: pulumi.String("my-sks-cluster"),
			Zone: pulumi.String("ch-gva-2"),
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
