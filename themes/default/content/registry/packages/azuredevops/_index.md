---
title: AzureDevOps
meta_desc: Provides an overview of the AzureDevOps Provider for Pulumi.
layout: overview
---

The AzureDevOps provider for Pulumi can be used to provision any of the cloud resources available in [AzureDevOps](https://azure.microsoft.com/en-us/services/devops/).
The AzureDevOps provider must be configured with credentials to deploy and update resources in AzureDevOps.

## Example

{{< chooser language "javascript,typescript,python,go,csharp" >}}

{{% choosable language javascript %}}

```javascript
const ado = require("@pulumi/azuredevops")

const project = new ado.Core.Project("demo-project", {
    projectName: "my-project",
});
```

{{% /choosable %}}
{{% choosable language typescript %}}

```typescript
import * as ado from "@pulumi/azuredevops";

const project = new ado.Core.Project("demo-project", {
    projectName: "my-project",
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_azuredevops as ado

project = ado.core.Project("demo-project",
                           project_name="my-project")
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	ado "github.com/pulumi/pulumi-azuredevops/sdk/v2/go/azuredevops"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		project, err := ado.NewProject(ctx, "test", &ado.ProjectArgs{
			ProjectName: pulumi.String("my-project"),
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
using Pulumi.AzureDevOps.Core;

await Deployment.RunAsync(() => 
{
	var project = new Project("test", new ProjectArgs
	{
		ProjectName = "my-project"
	});
});
```

{{% /choosable %}}

{{< /chooser >}}
