---
title: F5 BIG-IP
meta_desc: Provides an overview of the F5 BIG-IP Provider for Pulumi.
layout: package
---

The F5 BIG-IP provider for Pulumi can be used to provision any of the resources available with [F5 BIG-IP](https://www.f5.com/products/big-ip-services).
The F5 BIG-IP provider must be configured with credentials to deploy and update the resources.

## Example

{{< chooser language "javascript,typescript,python,go,csharp" >}}

{{% choosable language javascript %}}

```javascript
const f5bigip = require("@pulumi/f5bigip")

const monitor = new f5bigip.ltm.Monitor("backend", {
    name: "/Common/backend",
    parent: "/Common/http",
    send: "GET /\r\n",
    timeout: 5,
    interval: 10,
});
```

{{% /choosable %}}
{{% choosable language typescript %}}

```typescript
import * as f5bigip from "@pulumi/f5bigip";

const monitor = new f5bigip.ltm.Monitor("backend", {
    name: "/Common/backend",
    parent: "/Common/http",
    send: "GET /\r\n",
    timeout: 5,
    interval: 10,
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_f5bigip as f5bigip

monitor = f5bigip.ltm.Monitor("backend",
  name="/Common/backend",
  parent="/Common/http",
  send="GET /\r\n",
  timeout=5,
  interval=10,
)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	ltm "github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/ltm"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		monitor, err := ltm.NewMonitor(ctx, "backend", &ltm.MonitorArgs{
			Name:     pulumi.String("/Common/backend"),
			Parent:   pulumi.String("/Common/http"),
			Send:     pulumi.String("GET /\r\n"),
			Timeout:  pulumi.Int(5),
			Interval: pulumi.Int(10),
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
using Pulumi.F5bigip.Ltm;

class Program
{
    static Task Main() =>
        Deployment.Run(() => {
            var monitor = new Monitor("backend", new MonitorArgs
            {
                Name = "/Common/backend",
                Parent = "/Common/http",
                Send = "GET /\r\n",
                Timeout = 5,
                Interval = 100,
            });
        });
}
```

{{% /choosable %}}

{{< /chooser >}}
