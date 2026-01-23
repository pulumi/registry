---
title: Equinix Metal (Deprecated)
meta_desc: Provides an overview of the Equinix Metal provider for Pulumi.
layout: package
---

{{% notes type="info" %}}
This provider has been deprecated as of April 2023. It is recommended to use the [Official Equinix Provider](/registry/packages/equinix) as a replacement.
Unfortunately, there is no upgrade path from this provider to the Official Equinix provider, but you can take advantage of the [Pulumi Import](/docs/guides/adopting/import) to help achieve the migration.
{{% /notes %}}

The Equinix Metal provider for Pulumi can be used to provision any of the cloud resources available in [Equinix Metal](https://metal.equinix.com/).
The Equinix Metal provider must be configured with credentials to deploy and update resources in Equinix Metal.

## Example

{{< chooser language "typescript,python,go,csharp" >}}

{{% choosable language typescript %}}

```typescript
import * as metal from "@pulumi/equinix-metal";

const project = new metal.Project("my-project", {
  name: "DevelopmentEnvironment"
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_equinix_metal as metal
project = metal.Project("my-project",
  name='DevelopmentEnvironment'
)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	metal "github.com/pulumi/pulumi-equinix-metal/sdk/v2/go/equinix"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		project, err := metal.NewProject(ctx, "test", &metal.ProjectArgs{
			Name: pulumi.String("DevelopmentEnvironment"),
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
using Pulumi.EquinixMetal;

class Program
{
    static Task Main() =>
        Deployment.Run(() => {
            var project = new EquinixMetal.Project("my-project", new EquinixMetal.ProjectArgs
            {
                Name = "DevelopmentEnvironment"
            });
        });
}
```

{{% /choosable %}}

{{< /chooser >}}
