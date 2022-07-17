---
title: Astra
meta_desc: Provides an overview of the Datastax Astra Provider for Pulumi.
layout: overview
---

The Astra provider for Pulumi can be used to provision any of the cloud resources available in [Datastax Astra](https://www.datastax.com/products/datastax-astra).
The Astra provider must be configured with credentials to deploy and update resources in Astra.

## Example

{{< chooser language "typescript,python,go,csharp" >}}
{{% choosable language typescript %}}

```typescript
import * as astra from "@pulumiverse/astra";
const db = new astra.Database("example", {
    cloudProvider: "azure",
    keyspace: "default",
    regions: ["westus2"],
    name: "example-db"
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumiverse_astra as astra

db = astra.Database("example",
    cloud_provider="azure",
    keyspace="default",
    regions=["westus2"],
    name="example-db"
)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"fmt"
	astra "github.com/pulumiverse/pulumi-astra/sdk/go/astra"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		db, err := astra.NewDatabase(ctx, "example", &astra.DatabaseArgs{
            CloudProvider: pulumi.String("azure"),
            Keyspace: pulumi.String("default"),
            Regions: pulumi.StringArray{
                pulumi.String("westus2")
            },
            Name: pulumi.String("example-db"),
		})
		if err != nil {
			return fmt.Errorf("error creating instance server: %v", err)
		}

		ctx.Export("dbId", db.Id)

		return nil
	})
}
```

{{% /choosable %}}
{{% choosable language csharp %}}

```csharp
using Pulumi;
using Pulumiverse.Astra;

class AstraDb : Stack
{
    public AstraDb()
    {
        var db = new Database("example", new DatabaseArgs{
            CloudProvider: "azure",
            Keyspace: "default",
            Regions: new[] {"westus2"},
            Name: "example-db"
        });
    }
}
```

{{% /choosable %}}

{{< /chooser >}}
