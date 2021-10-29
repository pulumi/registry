---
title: Splunk
meta_desc: Provides an overview of the Splunk Provider for Pulumi.
layout: overview
---

The Splunk provider for Pulumi can be used to provision any of the cloud resources available in [Splunk](https://www.splunk.com/).
The Splunk provider must be configured with credentials to deploy and update resources in Splunk.

## Example

{{< chooser language "javascript,typescript,python,go,csharp" >}}

{{% choosable language javascript %}}

```javascript
const splunk = require("@pulumi/splunk")

const adminSamlGroup = new splunk.AdminSamlGroups("demo-ts-group", {
  roles: ["admin", "power"]
});
```

{{% /choosable %}}
{{% choosable language typescript %}}

```typescript
import * as splunk from "@pulumi/splunk";

const adminSamlGroup = new splunk.AdminSamlGroups("demo-ts-group", {
  roles: ["admin", "power"]
});

```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_splunk as splunk

saml_group = splunk.AdminSamlGroups("demo-py-group", roles=[
  "admin",
  "power",
])
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi-splunk/sdk/go/splunk"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		group, err := splunk.NewAdminSamlGroups(ctx, "demo", &splunk.AdminSamlGroupsArgs{
            Roles: pulumi.StringArray{
                pulumi.String("admin"),
                pulumi.String("power"),
            }
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
using Pulumi.Splunk;

class Program
{
    static Task Main() =>
        Deployment.Run(() => {
            var group = new AdminSamlGroups("saml-group", new AdminSamlGroupsArgs
            {
                Roles =
                {
                  "admin",
                  "power",
                }
            });
        });
}
```

{{% /choosable %}}

{{< /chooser >}}
