---
# WARNING: this file was fetched from https://raw.githubusercontent.com/dimeskigj/pulumi-dokploy/v0.3.1/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/dimeskigj/pulumi-dokploy/blob/v0.3.1/docs/_index.md
layout: package
title: Dokploy
meta_desc: Use the Pulumi Dokploy provider to manage Dokploy infrastructure and workloads.
description: Pulumi provider for managing Dokploy resources.
---

The Dokploy provider for Pulumi lets you manage [Dokploy](https://dokploy.com/)
projects, applications, databases, domains, backups, and related resources as
part of Pulumi programs. This is a **community-maintained** provider; neither
Dokploy nor Pulumi maintains this package.

## Installation

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}

```bash
npm install @dimeskigj/pulumi-dokploy
```

{{% /choosable %}}
{{% choosable language python %}}

```bash
pip install pulumi-dokploy
```

{{% /choosable %}}
{{% choosable language go %}}

```bash
go get github.com/dimeskigj/pulumi-dokploy/sdk/go/dokploy
```

{{% /choosable %}}
{{% choosable language csharp %}}

```bash
dotnet add package Dimeskigj.Pulumi.Dokploy
```

{{% /choosable %}}
{{% choosable language java %}}

Maven:

Use the provider version you are installing in place of `${DOKPLOY_VERSION}`.

```xml
<dependency>
  <groupId>net.dimeski.pulumi</groupId>
  <artifactId>dokploy</artifactId>
  <version>${DOKPLOY_VERSION}</version>
</dependency>
```

Gradle:

```groovy
implementation 'net.dimeski.pulumi:dokploy:${DOKPLOY_VERSION}'
```

{{% /choosable %}}
{{% choosable language yaml %}}

```bash
pulumi package add github.com/dimeskigj/pulumi-dokploy dokploy
```

{{% /choosable %}}
{{< /chooser >}}

## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as dokploy from "@dimeskigj/pulumi-dokploy";

const project = new dokploy.Project("example", {
  name: "registry-example",
  description: "Managed by Pulumi",
});

export const projectId = project.id;
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi
import pulumi_dokploy

project = pulumi_dokploy.Project(
    "example",
    name="registry-example",
    description="Managed by Pulumi",
)

pulumi.export("project_id", project.id)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
package main

import (
	"github.com/dimeskigj/pulumi-dokploy/sdk/go/dokploy"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		project, err := dokploy.NewProject(ctx, "example", &dokploy.ProjectArgs{
			Name: pulumi.String("registry-example"),
			Description: pulumi.StringPtr("Managed by Pulumi"),
		})
		if err != nil {
			return err
		}
		ctx.Export("projectId", project.ID())
		return nil
	})
}
```

{{% /choosable %}}
{{% choosable language csharp %}}

```csharp
using System.Collections.Generic;
using Pulumi;
using Dimeskigj.Pulumi.Dokploy;

return await Deployment.RunAsync(() =>
{
    var project = new Project("example", new ProjectArgs
    {
        Name = "registry-example",
        Description = "Managed by Pulumi",
    });

    return new Dictionary<string, object?>
    {
        ["projectId"] = project.Id,
    };
});
```

{{% /choosable %}}
{{% choosable language java %}}

```java
import com.pulumi.Context;
import com.pulumi.Pulumi;
import net.dimeski.pulumi.dokploy.Project;
import net.dimeski.pulumi.dokploy.ProjectArgs;

public class Main {
    public static void main(String[] args) {
        Pulumi.run(ctx -> {
            var project = new Project("example", ProjectArgs.builder()
                .name("registry-example")
                .description("Managed by Pulumi")
                .build());
            ctx.export("projectId", project.id());
        });
    }
}
```

{{% /choosable %}}
{{% choosable language yaml %}}

```yaml
name: dokploy-example
runtime: yaml
resources:
  project:
    type: dokploy:index:Project
    properties:
      name: registry-example
      description: Managed by Pulumi
outputs:
  projectId: ${project.id}
```

{{% /choosable %}}
{{< /chooser >}}

Configure the Dokploy endpoint and API key before running an example:

```bash
pulumi config set dokploy:endpoint https://dokploy.example.invalid
pulumi config set --secret dokploy:apiKey your-api-key
```

## Configuration

- `endpoint` (Required, Not secret) - The HTTP or HTTPS URL of the Dokploy instance. May also be set with `DOKPLOY_ENDPOINT`.
- `apiKey` (Required, Secret) - The API key used to authenticate to Dokploy. Set it with `pulumi config set --secret`; it may also be supplied with `DOKPLOY_API_KEY`.

Pulumi can acquire the matching provider plugin automatically from the GitHub
release metadata. See the [installation and configuration guide](https://www.pulumi.com/registry/packages/dokploy/installation-configuration/)
for additional secret-handling guidance.

Support and source code are available in the [repository](https://github.com/dimeskigj/pulumi-dokploy).
Report non-sensitive problems in [GitHub issues](https://github.com/dimeskigj/pulumi-dokploy/issues)
and see the [contributing guide](https://github.com/dimeskigj/pulumi-dokploy/blob/main/CONTRIBUTING.md).
