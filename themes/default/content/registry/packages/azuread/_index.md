---
title: Azure Active Directory (Azure AD)
meta_desc: Provides an overview of the Azure Active Directory (Azure AD) Provider for Pulumi.
layout: package
---

The AzureAD provider for Pulumi can be used to provision any of the Azure Active Directory resources available in [Azure](https://azure.microsoft.com/en-us/).
The AzureAD provider must be configured with credentials to deploy and update resources in Azure.

## Example

<!---
javascript removed
--->
{{< chooser language "typescript,python,go,csharp,java,yaml" >}}

{{% choosable language typescript %}}

```typescript
import * as azad from "@pulumi/azuread";

const group = new azad.Group("my-group", {
    name: "my-group",
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_azuread as azad

group = azad.Group("my-group",
            name="my-group")
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	azad "github.com/pulumi/pulumi-azuread/sdk/v4/go/azuread"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		group, err := azad.NewGroup(ctx, "my-group", &azad.GroupArgs{
			Name: pulumi.String("my-group"),
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
using Pulumi;
using Pulumi.AzureAD;

await Deployment.RunAsync(() =>
{
    var group = new Group("my-group", new GroupArgs
    {
        Name = "my-group",
    });
});
```

{{% /choosable %}}
{{% choosable language java %}}

```java
import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.azuread.Group;
import com.pulumi.azuread.GroupArgs;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    private static void stack(Context ctx) {
		final var group = new Group("my-group", GroupArgs.builder()
			.name("my-group")
			.build());
	}
}
```

{{% /choosable %}}
{{% choosable language yaml %}}

```yaml
resources:
  my-group:
    type: azuread:Group
    properties:
      name: my-group
```

{{% /choosable %}}

{{< /chooser >}}
