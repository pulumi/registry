---
title: ACI
meta_desc: Provides an overview of the ACI Provider for Pulumi.
layout: overview
---

The ACI provider for Pulumi can be used to provision any of the network resources available in an ACI based network controlled by an APIC controller.
The ACI provider must be configured with credentials to deploy and update resources on APIC.

## Example

{{< chooser language "typescript,python,go,csharp" >}}
{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aci from "@netascode/aci";

const tenant = new aci.apic.Rest(
    "TENANT1",
    {
        dn: "uni/tn-TENANT1",
        class_name: "fvTenant",
        content: {
            "name": "TENANT1",
            "descr": "Tenant created by Pulumi"
        }
    });
```
 
{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_aci as aci

tenant = aci.apic.Rest(
    "TENANT1",
    dn="uni/tn-TENANT1",
    class_name="fvTenant",
    content={
        "name": "TENANT1",
        "descr": "Tenant created by Pulumi",
    },
)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"fmt"
	aci "github.com/pulumi/pulumi-aci/sdk/go/aci"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		tenant, err := aci.apic.NewRest(ctx, "TENANT1", &aci.apic.RestArgs{
            Dn: pulumi.String("uni/tn-TENANT1"),
            ClassName: pulumi.String("fvTenant"),
            Content: pulumi.StringMap{
                "name": pulumi.String("TENANT1"),
                "descr": pulumi.String("Tenant created by Pulumi"),
            },
		})
		if err != nil {
			return fmt.Errorf("error creating tenant: %v", err)
		}

		return nil
	})
}
```

{{% /choosable %}}
{{% choosable language csharp %}}

```csharp
using Pulumi;
using Pulumi.Aci;

class AciTenant : Stack
{
    public AciTenant()
    {
        var tenant = new Rest("TENANT1", new RestArgs{
            Dn: "uni/tn-TENANT1",
            ClassName: "fvTenant",
            Content: {
                { "name", "TENANT1" },
                { "descr", "Tenant created by Pulumi" },
            },
        });
    }
}
```

{{% /choosable %}}

{{< /chooser >}}