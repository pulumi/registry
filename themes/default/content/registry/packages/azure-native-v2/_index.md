---
title: Azure Native (v2 Preview)
meta_desc: Learn how to use Pulumi's Azure Native Provider to reduce the complexity of managing and provisioning Azure resources with Azure Resource Manager (ARM) APIs.
layout: package
---

The Azure Native provider for Pulumi can be used to provision all of the cloud resources available in [Azure](https://azure.microsoft.com/en-us/). It manages and provisions resources using the [Azure Resource Manager (ARM) APIs](https://docs.microsoft.com/en-us/rest/api/resources/).

Azure Native must be configured with credentials to deploy and update resources in Azure; see [Installation & Configuration](./installation-configuration) for instructions.

**New to Pulumi and Azure?** [Get started with Azure using our tutorial](/docs/get-started/azure)

## Example

{{< chooser language "typescript,python,csharp,go,java,yaml" >}}

{{% choosable language typescript %}}

```typescript
import * as resources from "@pulumi/azure-native/resources";

const resourceGroup = new resources.ResourceGroup("resourceGroup");
```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi_azure_native as azure_native

resource_group = azure_native.resources.ResourceGroup("resourceGroup")
```

{{% /choosable %}}

{{% choosable language csharp %}}

```csharp
using Pulumi;
using Pulumi.AzureNative.Resources;

await Deployment.RunAsync(() =>
{
    var resourceGroup = new ResourceGroup("resourceGroup");
});
```

{{% /choosable %}}

{{% choosable language go %}}

```go
package main

import (
    "github.com/pulumi/pulumi-azure-native/sdk/go/azure/resources"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
    pulumi.Run(func(ctx *pulumi.Context) error {
        _, err := resources.NewResourceGroup(ctx, "resourceGroup", nil)
        if err != nil {
            return err
        }
        return nil
    })
}
```

{{% /choosable %}}

{{% choosable language java %}}

```java
import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.azurenative.resources.ResourceGroup;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    private static void stack(Context ctx) {
        var resourceGroup = new ResourceGroup("resourceGroup");
        ctx.export("resourceGroupName", resourceGroup.name());
    }
}
```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
resources:
  resourceGroup:
    type: azure-native:resources:ResourceGroup
```

{{% /choosable %}}

{{< /chooser >}}
