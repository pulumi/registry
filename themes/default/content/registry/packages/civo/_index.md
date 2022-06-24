---
title: Civo
meta_desc: Provides an overview of the Civo Provider for Pulumi.
layout: overview
---

The Civo provider for Pulumi can be used to provision any of the cloud resources available in [Civo](https://www.civo.com/).
The Civo provider must be configured with credentials to deploy and update resources in Civo.

## Example

{{< chooser language "javascript,typescript,python,go,csharp" >}}

{{% choosable language javascript %}}

```javascript
const civo = require("@pulumi/civo")

const network = new civo.Network("demo", {
    label: "demo-network"
});
```

{{% /choosable %}}
{{% choosable language typescript %}}

```typescript
import * as civo from "@pulumi/civo";

const network = new civo.Network("demo", {
    label: "demo-network-typescript"
});

```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_civo as civo

network = civo.Network("demo",
  label="demo-network")
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi-civo/sdk/go/civo"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		network, err := civo.NewNetwork(ctx, "demo", &civo.NetworkArgs{
			Label: pulumi.String("demo-network"),
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
using Pulumi.Civo;

await Deployment.RunAsync(() => 
{
	var network = new Network("demo", new NetworkArgs
	{
		Label = "demo-network"
	});
});
```

{{% /choosable %}}

{{< /chooser >}}
