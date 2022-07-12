---
title: Artifactory
meta_desc: Provides an overview of the Artifactory Provider for Pulumi.
layout: overview
---

The Artifactory provider for Pulumi can be used to provision any of the cloud resources available in [Artifactory](https://jfrog.com/artifactory/).
The Artifactory provider must be configured with credentials to deploy and update resources in Artifactory.

## Example

{{< chooser language "javascript,typescript,python,go,csharp" >}}

{{% choosable language javascript %}}

```javascript
const artifactory = require("@pulumi/artifactory")

new artifactory.LocalRepository('pulumipus', {
    key: "pulumipus", // this key can be the name of your repository
    packageType: "docker"
})
```

{{% /choosable %}}
{{% choosable language typescript %}}

```typescript
import * as artifactory from "@pulumi/artifactory";

new artifactory.LocalRepository('pulumipus', {
    key: "pulumipus", // this key can be the name of your repository
    packageType: "docker"
})
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_artifactory as artifactory

my_local = artifactory.LocalRepository("pylumipus",
                                       key="pylumipus",
                                       package_type="npm")
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
    artifactory "github.com/pulumi/pulumi-artifactory/sdk/go/artifactory"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		repo, err := artifactory.NewLocalRepository(ctx, "golumipus", &artifactory.LocalRepositoryArgs{
			Key:         pulumi.String("golumipus"),
			PackageType: pulumi.String("npm"),
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
using Pulumi.Artifactory;

await Deployment.RunAsync(() =>
{
    var repo = new LocalRepository("cslumipus", new LocalRepositoryArgs
    {
        Key = "cslumipus",
        PackageType = "npm",
    });
});
```

{{% /choosable %}}

{{< /chooser >}}
