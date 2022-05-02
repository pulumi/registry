---
title: Azure Native
meta_desc: Learn how to use Pulumi's Azure Native Provider to reduce the complexity of managing and provisioning Azure resources with Azure Resource Manager (ARM) APIs.
layout: overview
---

The Azure Native provider for Pulumi can be used to provision all of the cloud resources available in [Azure](https://azure.microsoft.com/en-us/). It manages and provisions resources using the [Azure Resource Manager (ARM) APIs](https://docs.microsoft.com/en-us/rest/api/resources/).

Azure Native must be configured with credentials to deploy and update resources in Azure; see [Installation & Configuration]({{<relref "./installation-configuration">}}) for instructions.

**New to Pulumi and Azure?** [Get started with Azure using our tutorial]({{<relref "/docs/get-started/azure">}})

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

class MyStack : Stack
{
    public MyStack()
    {
        var resourceGroup = new ResourceGroup("resourceGroup");
    }

}

class Program
{
    static Task<int> Main(string[] args) => Deployment.RunAsync<MyStack>();
}
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

Visit the [How-to Guides]({{<relref "./how-to-guides">}}) to find step-by-step guides for specific scenarios like running an app in Azure App Service or setting up a serverless Azure Function.

## Migration

### From Azure Classic

If you're already using the [Pulumi Azure Classic Provider]({{<relref "/registry/packages/azure">}}) and would like to migrate to Azure Native, use the [migration guide]({{<relref "./from-classic">}}).

### From Azure Resource Manager (ARM) templates

If you have Azure Resource Manager (ARM) templates that you'd like to migrate to Pulumi, use the [Migrate From Azure Resource Manager guide]({{<relref "/docs/guides/adopting/from_azure">}}).

## Manage incompatible resources using the Azure SDK

Some Azure resources aren't included in Azure Native because they're not compatible with the [Pulumi resource model]({{<relref "docs/intro/concepts/how-pulumi-works">}}). If you need to manage these kinds of resources, you can use convenience helpers provided by Azure Native to set up an Azure SDK client and credentials in your preferred language. Use these how-to guides to get started:

* [Typescript]({{<relref "/registry/packages/azure-native/how-to-guides/azure-ts-call-azure-sdk">}})
* [Go]({{<relref "/registry/packages/azure-native/how-to-guides/azure-go-call-azure-sdk">}})
* [C#]({{<relref "/registry/packages/azure-native/how-to-guides/azure-cs-call-azure-api">}})
* [Python]({{<relref "/registry/packages/azure-native/how-to-guides/azure-py-call-azure-sdk">}})

## How resources are versioned

Azure Native provides access to all API versions of each Azure resource so that you can access the entire Azure API surface and pin to the version you prefer.

Read the [version guide]({{< relref "./version-guide" >}}) to learn more about how you can manage the Azure API versions you're using, including both module-per-version and top-level-resources approaches.
