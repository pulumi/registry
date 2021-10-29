---
title: Linode
meta_desc: This page provides an overview of the Linode SDK for Pulumi.
layout: overview
---

The Linode provider for Pulumi can be used to provision any of the cloud resources available in [Linode](https://www.linode.com).
The Linode provider must be configured with credentials to deploy and update resources in Linode.

## Example

{{< chooser language "javascript,typescript,python,go,csharp" >}}

{{% choosable language javascript %}}

```javascript
const linode = require("@pulumi/linode")
const domain = new linode.Domain("my-domain", {
  domain: "foobar.example",
  soaEmail: "example@foobar.example",
  type: "master",
});
```

{{% /choosable %}}
{{% choosable language typescript %}}

```typescript
import * as linode from "@pulumi/linode";
const domain = new linode.Domain("my-domain", {
  domain: "foobar.example",
  soaEmail: "example@foobar.example",
  type: "master",
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_linode as linode
domain = linode.Domain("my-domain",
  domain='foobar.example',
  soa_email='example@foobar.example',
  type='master',
)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	linode "github.com/pulumi/pulumi-linode/sdk/v3/go/linode"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		domain, err := linode.NewDomain(ctx, "test", &linode.DomainArgs{
			Domain:   pulumi.String("foobar.example"),
			SoaEmail: pulumi.String("example@foobar.example"),
			Type:     pulumi.String("master"),
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
using Pulumi.Linode;

class Program
{
    static Task Main() =>
        Deployment.Run(() => {
            var domain = new Linode.Domain("my-domain", new Linode.DomainArgs
            {
                Domain = "foobar.example",
                SoaEmail = "example@foobar.example",
                Type = "master",
            });
        });
}
```

{{% /choosable %}}

{{< /chooser >}}
