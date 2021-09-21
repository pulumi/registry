---
title: MongoDB Atlas
meta_desc: This page provides an overview of the MongoDB Atlas Provider for Pulumi.
layout: overview
---

The MongoDB Atlas provider for Pulumi can be used to provision any of the resources available for MongoDB Atlas.
The MongoDB Atlas provider must be configured with credentials to deploy and update resources in MongoDB Atlas.

## Example

{{< chooser language "javascript,typescript,python,go,csharp" >}}

{{% choosable language javascript %}}

```javascript
const mongodbatlas = require("@pulumi/mongodbatlas")

const project = new mongodbatlas.Project("my-demo-project", {
    orgId: "12345",
});
```

{{% /choosable %}}
{{% choosable language typescript %}}

```typescript
import * as mongodbatlas from "@pulumi/mongodbatlas";

const project = new mongodbatlas.Project("my-demo-project", {
    orgId: "12345",
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_mongodbatlas as mongodbatlas

project = mongodbatlas.Project("my-demo-project", org_id="12345")
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
    mongodbatlas "github.com/pulumi/pulumi-mongodbatlas/sdk/v2/go/mongodbatlas"
)

func main() {
    pulumi.Run(func(ctx *pulumi.Context) error {
        project, err := mongodb.NewProject(ctx, "my-demo-project", &mongodb.ProjectArgs{
        OrgId: pulumi.String("12345"),
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
using Pulumi.Mongodbatlas;

class Program
{
    static Task Main() =>
        Deployment.Run(() => {
            project = new Project("my-demo-project", new ProjectArgs{
              OrgId = "12345",
            });
        });
}
```

{{% /choosable %}}

{{< /chooser >}}

## Libraries

The following packages are available in packager managers:

* JavaScript/TypeScript: [`@pulumi/mongodbatlas`](https://www.npmjs.com/package/@pulumi/mongodbatlas)
* Python: [`pulumi-mongodbatlas`](https://pypi.org/project/pulumi-mongodbatlas/)
* Go: [`github.com/pulumi/pulumi-mongodbatlas/sdk/v2/go/mongodbatlas`](https://github.com/pulumi/pulumi-mongodbatlas)
* .NET: [`Pulumi.Mongodbatlas`](https://www.nuget.org/packages/Pulumi.Mongodbatlas)

The MongoDB Atlas provider is open source and available in the [pulumi/pulumi-mongodbatlas](https://github.com/pulumi/pulumi-mongodbatlas) repo.
