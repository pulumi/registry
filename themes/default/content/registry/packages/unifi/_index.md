---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumiverse/pulumi-unifi/v0.2.0/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/pulumiverse/pulumi-unifi/blob/v0.2.0/docs/_index.md
title: Unifi
meta_desc: Provides an overview of the Unifi Provider for Pulumi.
layout: overview
---

The Unifi provider for Pulumi can be used to provision any of the network resources available in a Unifi based network controlled by a Unifi controller.
The Unifi provider must be configured with credentials to deploy and update resources in Unifi.

## Example

{{< chooser language "typescript,python,go,csharp" >}}
{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as unifi from "@pulumiverse/unifi";

const mysite = new unifi.Site("mysite", {
    description: "mysite",
});
```
 
{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumiverse_unifi as unifi

db = unifi.Site("mysite",
    description="mysite"
)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"fmt"
	unifi "github.com/pulumiverse/pulumi-unifi/sdk/go/unifi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		site, err := unifi.NewSite(ctx, "mysite", &unifi.SiteArgs{
            Description: pulumi.String("mysite"),
		})
		if err != nil {
			return fmt.Errorf("error creating site: %v", err)
		}

		ctx.Export("dbId", site.Id)

		return nil
	})
}
```

{{% /choosable %}}
{{% choosable language csharp %}}

```csharp
using Pulumi;
using Pulumiverse.Unifi;

class UnifiNet : Stack
{
    public UnifiNet()
    {
        var db = new Site("mysite", new SiteArgs{
            Descriptino: "mysite"
        });
    }
}
```

{{% /choosable %}}

{{< /chooser >}}
