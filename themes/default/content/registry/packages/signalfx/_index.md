---
title: SignalFx
meta_desc: Provides an overview of the SignalFx Provider for Pulumi.
layout: overview
---

{{% overview-description %}}
The SignalFx provider for Pulumi can be used to provision any of the cloud resources available in [SignalFx](https://signalfx.com/).
The SignalFx provider must be configured with credentials to deploy and update resources in SignalFx.
{{% /overview-description %}}

## Example

{{< chooser language "javascript,typescript,python,go,csharp" >}}

{{% choosable language javascript %}}

```javascript
const signalfx = require("@pulumi/signalfx")

const group = new signalfx.DashboardGroup("my-group", {
  description: "my demo dashboard group"
});
```

{{% /choosable %}}
{{% choosable language typescript %}}

```typescript
import * as signalfx from "@pulumi/signalfx";

const group = new signalfx.DashboardGroup("my-group", {
  description: "my demo dashboard group"
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_signalfx as signalfx

group = signalfx.DashboardGroup("my-group",
  description="my demo dashboard group"
)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	signalfx "github.com/pulumi/pulumi-signalfx/sdk/v5/go/signalfx"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		group, err := signalfx.NewDashboardGroup(ctx, "my-group", &signalfx.DashboardGroupArgs{
			Description: pulumi.String("my demo dashboard group"),
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
using System.Collections.Generic;
using System.Threading.Tasks;
using Pulumi;
using Pulumi.Signalfx;

class Program
{
    static Task Main() =>
        Deployment.Run(() => {
            var group = new DashboardGroup("my-group", new DashboardGroupArgs
            {
                Description = "my demo dashboard group"
            });
        });
}
```

{{% /choosable %}}

{{< /chooser >}}
