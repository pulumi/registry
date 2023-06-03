---
title: Spotinst
meta_desc: Provides an overview of the Spotinst Provider for Pulumi.
layout: package
---

The Spotinst provider for Pulumi can be used to provision any of the resources available in [Spotinst](https://spotinst.com/).

## Example

{{< chooser language "javascript,typescript,python,go,csharp" >}}

{{% choosable language javascript %}}

```javascript
const spotinst = require("@pulumi/spotinst")

const deployment = new spotinst.multai.Deployment("myDeployment", {});
```

{{% /choosable %}}
{{% choosable language typescript %}}

```typescript
import * as spotinst from "@pulumi/spotinst";

const deployment = new spotinst.multai.Deployment("myDeployment");
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_spotinst as spotinst

deployment = spotinst.multai.Deployment("my_deployment")
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	multai "github.com/pulumi/pulumi-spotinst/sdk/v3/go/spotinst/multai"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		deployment, err := multai.NewDeployment(ctx, "example", &multai.DeploymentArgs{})
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
using Pulumi.SpotInst.Multai;

class Program
{
    static Task Main() =>
        Deployment.Run(() => {
            var deployment = new Multai.Deployment("example", new Multai.DeploymentArgs{});
        });
}
```

{{% /choosable %}}

{{< /chooser >}}
