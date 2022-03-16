---
title: vSphere
meta_desc: Provides an overview of the vSphere provider for Pulumi.
layout: overview
---

{{% overview-description %}}
The vSphere provider for Pulumi can be used to provision any of the cloud resources available in [vSphere](https://www.vmware.com/products/vsphere.html).
The vSphere provider must be configured with credentials to deploy and update resources in vSphere.
{{% /overview-description %}}

## Example

{{< chooser language "javascript,typescript,python,go,csharp" >}}

{{% choosable language javascript %}}

```javascript
const vsphere = require("@pulumi/vsphere")
const dc = new vsphere.Datacenter("my-dc", {
  name: "Production-DataCenter",
});
```

{{% /choosable %}}
{{% choosable language typescript %}}

```typescript
import * as vsphere from "@pulumi/vsphere";
const dc = new vsphere.Datacenter("my-dc", {
  name: "Production-DataCenter",
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_vsphere as vsphere
dc = vsphere.Datacenter("my-dc",
  name='Production-DataCenter',
)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	vsphere "github.com/pulumi/pulumi-vsphere/sdk/v3/go/vsphere"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		dc, err := vsphere.NewDatacenter(ctx, "test", &vsphere.DatacenterArgs{
			Name: pulumi.String("Production-DataCenter"),
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
using System.Threading.Tasks;
using Pulumi;
using Pulumi.Vsphere;

class Program
{
    static Task Main() =>
        Deployment.Run(() => {
            var dc = new Vsphere.Datacenter("my-dc", new Vsphere.DatacenterArgs
            {
                Name = "Production-DataCenter",
            });
        });
}
```

{{% /choosable %}}

{{< /chooser >}}
