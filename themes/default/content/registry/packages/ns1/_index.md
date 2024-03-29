---
title: NS1
meta_desc: Provides an overview of the NS1 Provider for Pulumi.
layout: package
---

The NS1 provider for Pulumi can be used to provision any of the cloud resources available in [NS1](https://www.ns1.com/).
The NS1 provider must be configured with credentials to deploy and update resources in NS1.

## Example

{{< chooser language "javascript,typescript,python,go,csharp" >}}

{{% choosable language javascript %}}

```javascript
const ns1 = require("@pulumi/ns1")

const myZone = new ns1.Zone("demo-zone", {
    zone: "demo.pulumi.io"
});
```

{{% /choosable %}}
{{% choosable language typescript %}}

```typescript
import * as ns1 from "@pulumi/ns1";

const myZone = new ns1.Zone("demo-zone", {
    zone: "demo.pulumi.io"
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_ns1 as ns1

zone = ns1.Zone("demo-zone",
  zone="demo.pulumi.io",
)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	ns1 "github.com/pulumi/pulumi-ns1/sdk/v2/go/ns1"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		zone, err := ns1.NewZone(ctx, "demo-zone", &ns1.ZoneArgs{
			Zone: pulumi.String("demo.pulumi.io"),
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
using Pulumi.Ns1;

class Program
{
    static Task Main() =>
        Deployment.Run(() => {
            var zone = new Zone("demo-zone", new ZoneArgs
            {
                ZoneName = "demo.pulumi.io",
            });
        });
}
```

{{% /choosable %}}

{{< /chooser >}}
