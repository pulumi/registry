---
title: Symbiosis
meta_desc: Provides an overview of the Symbiosis Provider for Pulumi.
layout: overview
---

The Symbiosis provider for Pulumi can be used to provision any of the resources available in a Symbiosis cloud account.
The Symbiosis provider must be configured with credentials to deploy and update resources in Symbiosis.

## Example

{{< chooser language "typescript,python,go,csharp" >}}
{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as symbiosis from "@kuraudo-io/symbiosis";

const mysite = new symbiosis.Cluster("example", {
    region: "germany-1"
});
```
 
{{% /choosable %}}
{{% choosable language python %}}

```python
import kuraudo_symbiosis as symbiosis

db = symbiosis.Cluster("example",
    region="germany-1"
)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"github.com/kuraudo-io/pulumi-symbiosis/sdk/go/symbiosis"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := symbiosis.NewCluster(ctx, "example", &symbiosis.ClusterArgs{
			Region: pulumi.String("germany-1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
```

{{% /choosable %}}
{{% choosable language csharp %}}

```csharp
using Pulumi;
using Kuraudo.Symbiosis;

class ExampleCluster : Stack
{
    public ExampleCluster()
    {
        var cluster = new Cluster("example", new ClusterArgs{
            Region: "germany-1"
        });
    }
}
```

{{% /choosable %}}

{{< /chooser >}}
