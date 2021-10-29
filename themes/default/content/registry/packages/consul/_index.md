---
title: HashiCorp Consul
meta_desc: Provides an overview of the HashiCorp Consul Provider for Pulumi.
layout: overview
---

The HashiCorp Consul provider for Pulumi can be used to provision any of the resources available in [Consul](https://www.consul.io/).

## Example

{{< chooser language "javascript,typescript,python,go,csharp" >}}

{{% choosable language javascript %}}

```javascript
const consul = require("@pulumi/consul")

const node = new consul.Node("compute", {
  address: "www.google.com"
});
```

{{% /choosable %}}
{{% choosable language typescript %}}

```typescript
import * as consul from "@pulumi/consul";

const node = new consul.Node("compute", {
  address: "www.google.com"
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_consul as consul

node = consul.Node("compute",
  address='www.google.com'
)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	consul "github.com/pulumi/pulumi-consul/sdk/v3/go/consul"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		node, err := consul.NewNode(ctx, "compute", &consul.NodeArgs{
			Address: pulumi.String("www.google.com"),
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
using Pulumi.Consul;

class Program
{
    static Task Main() =>
        Deployment.Run(() => {
            var node = new Consul.Node("compute", new Consul.NodeArgs
            {
                Address = "www.google.com",
            });
        });
}
```

{{% /choosable %}}

{{< /chooser >}}
