---
title: ACI
meta_desc: Provides an overview of the Cisco ACI Provider for Pulumi.
layout: overview
---

## Application Centric Infrastructure (ACI)

The Cisco Application Centric Infrastructure ([ACI](https://www.cisco.com/go/aci)) allows application requirements to define the network. This architecture simplifies, optimizes, and accelerates the entire application deployment life cycle.

## Application Policy Infrastructure Controller (APIC)

The APIC manages the scalable [ACI](https://www.cisco.com/go/aci) multi-tenant fabric. The APIC provides a unified point of automation and management, policy programming, application deployment, and health monitoring for the fabric. The APIC, which is implemented as a replicated synchronized clustered controller, optimizes performance, supports any application anywhere, and provides unified operation of the physical and virtual infrastructure. The APIC enables network administrators to easily define the optimal network for applications. Data center operators can clearly see how applications consume network resources, easily isolate and troubleshoot application and infrastructure problems, and monitor and profile resource usage patterns. The Cisco Application Policy Infrastructure Controller (APIC) API enables applications to directly connect with a secure, shared, high-performance resource pool that includes network, compute, and storage capabilities.

## Cisco ACI Provider for Pulumi

The Cisco ACI provider for Pulumi can be used to provision any of the network resources available in an [ACI](https://www.cisco.com/go/aci) based network controlled by an APIC controller. The ACI provider must be configured with credentials to deploy and update resources on APIC.

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