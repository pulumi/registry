---
# WARNING: this file was fetched from https://raw.githubusercontent.com/DefangLabs/pulumi-defang/v1.0.0/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Defang Provider
meta_desc: Take your app from Docker Compose to a secure and scalable cloud deployment with Pulumi.
layout: package
---
# Defang Pulumi Provider

![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/DefangLabs/pulumi-defang?label=Version)

The Pulumi Provider for [Defang](https://defang.io) — Take your app from Docker Compose to a secure and scalable cloud deployment with Pulumi.

## Example usage

You can find complete working TypeScript, Python, Go, .NET, and Yaml code samples in the [`./examples`](https://github.com/DefangLabs/pulumi-defang/tree/main/examples) directory, and some example snippets below:

{{< chooser language "typescript,python,go,dotnet,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as defang from "@defang-io/pulumi-defang";

const myProject = new defang.Project("myProject", {
    providerID: "aws",
    configPaths: ["compose.yaml"],
});
export const output = {
    albArn: myProject.albArn,
    etag: myProject.etag,
};
```

{{% /choosable %}}

{{% choosable language python %}}
```python
import pulumi
import pulumi_defang as defang

my_project = defang.Project("myProject",
    provider_id="aws",
    config_paths=["compose.yaml"])
pulumi.export("output", {
    "albArn": my_project.alb_arn,
    "etag": my_project.etag,
})
```

{{% /choosable %}}

{{% choosable language go %}}
```go
package main

import (
	"example.com/pulumi-defang/sdk/go/defang"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		myProject, err := defang.NewProject(ctx, "myProject", &defang.ProjectArgs{
			ProviderID: pulumi.String("aws"),
			ConfigPaths: pulumi.StringArray{
				pulumi.String("compose.yaml"),
			},
		})
		if err != nil {
			return err
		}
		ctx.Export("output", pulumi.StringMap{
			"albArn": myProject.AlbArn,
			"etag":   myProject.Etag,
		})
		return nil
	})
}
```

{{% /choosable %}}

{{% choosable language dotnet %}}
```dotnet
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Defang = DefangLabs.Defang;

return await Deployment.RunAsync(() =>
{
    var myProject = new Defang.Project("myProject", new()
    {
        ProviderID = "aws",
        ConfigPaths = new[]
        {
            "./compose.yaml",
        },
    });

    return new Dictionary<string, object?>
    {
        ["output"] =
        {
            { "albArn", myProject.AlbArn },
            { "etag", myProject.Etag },
        },
    };
});

```

{{% /choosable %}}

{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
config:
    defang:Project:
        providerID: aws
        configPaths:
            - ./compose.yaml
```

{{% /choosable %}}
{{< /chooser >}}

## Installation and Configuration

See our [Installation and Configuration](https://pulumi.com/registry/packages/defang/installation-configuration/) docs

## Development

See the [Contributing](https://github.com/DefangLabs/pulumi-defang/blob/main/CONTRIBUTING.md) doc.
