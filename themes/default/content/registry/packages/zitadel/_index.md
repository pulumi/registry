---
title: Zitadel
meta_desc: Provides an overview of the Zitadel Provider for Pulumi.
layout: overview
---

The Zitadel provider for Pulumi can be used to provision any of the cloud resources available in [Zitadel](https://zitadel.com/).
The Zitadel provider must be configured with credentials to deploy and update resources in Zitadel.

## Example

{{< chooser language "typescript,python,go,csharp" >}}
{{% choosable language typescript %}}

```typescript
import * as zitadel from "@pulumiverse/zitadel";
const db = new zitadel.Org("example", {
    name: "example",
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumiverse_zitadel as zitadel

db = zitadel.Database("example",
    name="example"
)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"fmt"
	zitadel "github.com/pulumiverse/pulumi-zitadel/sdk/go/zitadel"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		org, err := zitadel.NewOrg(ctx, "example", &zitadel.OrgArgs{
            Name: pulumi.String("example"),
		})
		if err != nil {
			return fmt.Errorf("error creating org: %v", err)
		}

		ctx.Export("orgId", org.Id)

		return nil
	})
}
```

{{% /choosable %}}
{{% choosable language csharp %}}

```csharp
using Pulumi;
using Pulumiverse.Zitadel;

class Zitadel : Stack
{
    public Zitadel()
    {
        var org = new Org("example", new OrgArgs{
            Name: "example"
        });
    }
}
```

{{% /choosable %}}

{{< /chooser >}}