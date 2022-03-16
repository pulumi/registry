---
title: GitLab
meta_desc: Provides an overview of the GitLab Provider for Pulumi.
layout: overview
---

The GitLab provider for Pulumi can be used to provision any of the cloud resources available in [GitLab](https://about.gitlab.com/).
The GitLab provider must be configured with credentials to deploy and update resources in GitLab.

## Example

{{< chooser language "javascript,typescript,python,go,csharp" >}}

{{% choosable language javascript %}}

```javascript
const gitlab = require("@pulumi/gitlab")

const project = new gitlab.Project("my-project", {
  description: "example project created by Pulumi",
  visibilityLevel: "public",
});
```

{{% /choosable %}}
{{% choosable language typescript %}}

```typescript
import * as gitlab from "@pulumi/gitlab";

const project = new gitlab.Project("my-project", {
  description: "example project created by Pulumi",
  visibilityLevel: "public",
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_gitlab as gitlab

project = gitlab.Project("my-project",
  description="example project created by Pulumi",
  visibility_level="public",
)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	gitlab "github.com/pulumi/pulumi-gitlab/sdk/v4/go/gitlab"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		project, err := gitlab.NewProject(ctx, "test", &gitlab.ProjectArgs{
			Description:     pulumi.String("example project created by Pulumi"),
			VisibilityLevel: pulumi.String("public"),
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
using Pulumi.Gitlab;

class Program
{
    static Task Main() =>
        Deployment.Run(() => {
            var project = new Project("test", new ProjectArgs
            {
                Description = "example project created by Pulumi",
                VisibilityLevel = "public",
            });
        });
}
```

{{% /choosable %}}

{{< /chooser >}}
