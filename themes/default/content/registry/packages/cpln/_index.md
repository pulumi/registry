---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumiverse/pulumi-cpln/v0.0.61/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Control Plane (cpln)
meta_desc: Provides an overview of the Control Plane (cpln) Provider for Pulumi.
layout: package
---

The Control Plane (cpln) Pulumi provider enables the scaffolding of any Control Plane (https://controlplane.com/) object as code. It enables infrastructure as code with all the added benefit of the global virtual cloud (GVC). You can build your VPCs, subnets, databases, queues, caches, etc., and overlay them with a multi-cloud/multi-region universal compute workloads that span regions and clouds. Nearly everything you can do using the Control Plane CLI, UI or API is available using Pulumi.

## Example

{{< chooser language "typescript,python,go,csharp" >}}
{{% choosable language typescript %}}

```typescript
import * as cpln from "@pulumiverse/cpln";

const group = new cpln.Group("example", {
    description: "example"
});

export const groupId = group.id;
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi
import pulumiverse_cpln as cpln

group = cpln.Group("example",
    description="example"
)

pulumi.export("group.id", group.id)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	cpln "github.com/pulumiverse/pulumi-cpln/sdk/go/cpln"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		group, err := cpln.NewGroup(ctx, "example", &cpln.GroupArgs{
			Description: pulumi.String("example"),
		})
		if err != nil {
			return fmt.Errorf("error creating a group: %v", err)
		}

		ctx.Export("group.id", group.ID())

		return nil
	})
}
```

{{% /choosable %}}
{{% choosable language csharp %}}

```csharp
using System.Collections.Generic;
using Pulumi;
using Pulumiverse.Cpln;

return await Deployment.RunAsync(() =>
{
    var group = new Group("example", new GroupArgs{
        Description = "example"
    });

    return new Dictionary<string, object?>
    {
        ["group.id"] = group.Id
    };
});
```

{{% /choosable %}}

{{< /chooser >}}
