---
title: Sumo Logic
meta_desc: Provides an overview of the Sumo Logic Provider for Pulumi.
layout: overview
---

The  provider for Pulumi can be used to provision any of the cloud resources available in [Sumo Logic](https://www.sumologic.com/).
The Sumo Logic provider must be configured with credentials to deploy and update resources in Sumo Logic.

## Example

{{< chooser language "javascript,typescript,python,go,csharp" >}}

{{% choosable language javascript %}}

```javascript
const sumologic = require("@pulumi/sumologic")

const folder = new sumologic.Folder("folder", {
  description: "A test folder",
  parentId: "<personal folder id>",
});
```

{{% /choosable %}}
{{% choosable language typescript %}}

```typescript
import * as sumologic from "@pulumi/sumologic";

const folder = new sumologic.Folder("folder", {
  description: "A test folder",
  parentId: "<personal folder id>",
});

```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_sumologic as sumologic

folder = sumologic.Folder("folder",
                          description="A test folder",
                          parent_id="<personal folder id>")
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi-sumologic/sdk/go/sumologic"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sumologic.NewFolder(ctx, "folder", &sumologic.FolderArgs{
			Description: pulumi.String("A test folder"),
            ParentId:    pulumi.String("<personal folder id>"),
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
using Pulumi.SumoLogic;

class Program
{
    static Task Main() =>
        Deployment.Run(() => {
            var folder = new SumoLogic.Folder("folder", new SumoLogic.FolderArgs
            {
                Description = "A test folder",
                 ParentId = "<personal folder id>",
            });
        });
}
```

{{% /choosable %}}

{{< /chooser >}}
