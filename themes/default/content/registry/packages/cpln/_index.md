---
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

const location = new cpln.Location("example", {
    name: "aws-us-west-2"
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumiverse_cpln as cpln

db = cpln.Location("example",
    name="aws-us-west-2"
)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"fmt"
	cpln "github.com/pulumiverse/pulumi-cpln/sdk/go/cpln"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		location, err := cpln.NewLocation(ctx, "example", &cpln.LocationArgs{
			Name: pulumi.String("aws-us-west-2"),
		})
		if err != nil {
			return fmt.Errorf("error creating location: %v", err)
		}

		ctx.Export("location.enabled", location.enabled)

		return nil
	})
}
```

{{% /choosable %}}
{{% choosable language csharp %}}

```csharp
using Pulumi;
using Pulumiverse.cpln;

class cpln : Stack
{
    public cpln()
    {
        var location = new Location("example", new LocationArgs{
            Name: "example"
        });
    }
}
```

{{% /choosable %}}

{{< /chooser >}}
