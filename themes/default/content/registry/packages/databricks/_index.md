---
title: Databricks
meta_desc: Provides an overview of the Databricks Provider for Pulumi.
layout: overview
---

The Databricks provider for Pulumi can be used to provision any of the cloud resources available in [Databricks](https://databricks.com/).
The Databricks provider must be configured with credentials to deploy and update resources in Databricks.

## Example

{{< chooser language "javascript,typescript,python,go,csharp" >}}

{{% choosable language javascript %}}

```javascript
const databricks = require("@pulumi/databricks");

const group = new databricks.Group("js-group", {
    displayName: "my company group",
});
```

{{% /choosable %}}
{{% choosable language typescript %}}

```typescript
import * as databricks from "@pulumi/databricks";

const group = new databricks.Group("ts-group", {
    displayName: "my company group",
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_databricks as databricks

group = databricks.Group("py-group", display_name="my company group")
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	databricks "github.com/pulumi/pulumi-databricks/sdk/go/databricks"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		group, err := databricks.NewGroup(ctx, "go-group", &databricks.GroupArgs{
			DisplayName: pulumi.String("my company group"),
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
using Pulumi.Databricks;

class Program
{
    static Task Main() =>
        Deployment.Run(() => {
            var group = new Group("csharp-group", new GroupArgs
            {
                DisplayName = "my company group",
            });
        });
}
```

{{% /choosable %}}

{{< /chooser >}}
