---
title: dbt Cloud
meta_desc: Provides an overview of the dbt Cloud Provider for Pulumi.
layout: package
---

The dbt Cloud Resource Provider lets you manage [dbt Cloud](https://www.getdbt.com/) resources. This is a bridged provider from the terraform dbt Cloud provider, located at:

[github.com/dbt-labs/terraform-provider-dbtcloud](https://github.com/dbt-labs/terraform-provider-dbtcloud/)

## Example

{{< chooser language "typescript,python,go,csharp" >}}
{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as dbtcloud from "@pulumi/dbtcloud";

const project =  new dbtcloud.Project("ts-project", {name: "ts-project"});

export const projectname = project.name;
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi
import pulumi_dbtcloud as dbtcloud

project = dbtcloud.Project("py-project", name="py-project")

pulumi.export("project_name", project.name)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
package main

import (
	"github.com/pulumi/pulumi-dbtcloud/sdk/go/dbtcloud"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		project, err := dbtcloud.NewProject(ctx, "go-project", &dbtcloud.ProjectArgs{
			Name: pulumi.String("go-project"),
		})
		if err != nil {
			return err
		}
		ctx.Export("project_name", project.Name)
		return nil
	})
}
```

{{% /choosable %}}
{{% choosable language csharp %}}

```csharp
using System.Collections.Generic;
using Pulumi;
using Pulumi.DbtCloud;

return await Deployment.RunAsync(() =>
{
    var project = new Project("cs-project", new ProjectArgs { Name = "cs-project" });

    return new Dictionary<string, object?>
    {
        ["ProjectName"] = project.Name
    };
});
```

{{% /choosable %}}

{{% /choosable %}}
{{% choosable language java %}}

```java
import java.util.Map;
import java.util.HashMap;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.resources.Deployment;
import com.pulumi.dbtcloud.Project;
import com.pulumi.dbtcloud.ProjectArgs;

public class Main {
    public static void main(String[] args) {
        Pulumi.run(ctx -> {
            var project = new Project("java-project", ProjectArgs.builder()
                .name("java-project")
                .build());

            ctx.export("ProjectName", project.name());
        });
    }
}
```

{{% /choosable %}}

{{% /choosable %}}
{{% choosable language yaml %}}

```yaml
resources:
  yaml-project:
    type: dbtcloud:index:Project
    properties:
      name: "yaml-project"
outputs:
  project_name: ${yaml-project.name}

```

{{% /choosable %}}

{{< /chooser >}}

