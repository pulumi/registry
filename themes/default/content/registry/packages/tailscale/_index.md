---
title: Tailscale
meta_desc: Provides an overview of the Tailscale Provider for Pulumi.
layout: overview
---

The Tailscale provider for Pulumi can be used to provision any of the resources available in [Tailscale](https://tailscale.com/).

## Example

{{< chooser language "javascript,typescript,python,go,csharp" >}}

{{% choosable language javascript %}}

```javascript
const tailscale = require("@pulumi/tailscale")

const nameservers = new tailscale.DnsNameservers("ts-demo", {
    nameservers: ["1.1.1.1"]
});
```

{{% /choosable %}}
{{% choosable language typescript %}}

```typescript
import * as tailscale from "@pulumi/tailscale";

const nameservers = new tailscale.DnsNameservers("ts-demo", {
    nameservers: ["1.1.1.1"]
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import json
import pulumi_tailscale as tailscale

acl = tailscale.Acl("demo-py",
                    acl=json.dumps({
                        "ACLs": [{
                            "Action": "accept",
                            "User": ["*"],
                            "Ports": ["*:*"],
                        }]
                    }))
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	tailscale "github.com/pulumi/pulumi-tailscale/sdk/go/tailscale"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		ns, err := tailscale.NewDnsNameservers(ctx, "sample-record", &tailscale.DnsNameserversArgs{
			NameServers: pulumi.StringArray{
                pulumi.String("1.1.1.1"),
			}.
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
using Pulumi.Tailscale;

class Program
{
    static Task Main() =>
        Deployment.Run(() => {
            var sampleNameservers = new Tailscale.DnsNameservers("sampleNameservers", new Tailscale.DnsNameserversArgs
            {
                Nameservers =
                {
                    "1.1.1.1",
                },
            });
        });
}
```

{{% /choosable %}}

{{< /chooser >}}
