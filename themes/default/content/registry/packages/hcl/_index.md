---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-hcl/ac67c9cfa212ea64dda1946dadda9ad91e0ca828/registry/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/pulumi/pulumi-hcl/blob/ac67c9cfa212ea64dda1946dadda9ad91e0ca828/registry/_index.md
title: Any HCL Module
meta_desc: Instantiate any Terraform or OpenTofu module as a Pulumi component resource
layout: package
---

The `hcl` package lets you instantiate any Terraform or OpenTofu module as a
Pulumi [component resource](https://www.pulumi.com/docs/concepts/resources/components/).
Point the `Module` resource at a module source — a Terraform registry reference,
a Git URL, or a local path — pass its input variables, and consume its outputs
like any other Pulumi resource, from any supported language.

The module runs under Pulumi's engine: its resources are tracked in your Pulumi
state, its provider credentials come from your Pulumi configuration, and its
sensitive outputs are handled as Pulumi secrets.

## Example

The following program instantiates the community
[`terraform-aws-modules/vpc/aws`](https://registry.terraform.io/modules/terraform-aws-modules/vpc/aws)
module and exports the resulting VPC id.

{{< chooser language "typescript,python,go,csharp,java,yaml" / >}}

{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as hcl from "@pulumi/hcl";

const vpc = new hcl.Module("vpc", {
    source: "terraform-aws-modules/vpc/aws",
    version: "5.0.0",
    inputs: {
        name: "example-vpc",
        cidr: "10.0.0.0/16",
    },
});

export const vpcId = vpc.outputs.apply(o => o["vpc_id"]);
```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi
import pulumi_hcl as hcl

vpc = hcl.Module("vpc",
    source="terraform-aws-modules/vpc/aws",
    version="5.0.0",
    inputs={
        "name": "example-vpc",
        "cidr": "10.0.0.0/16",
    })

pulumi.export("vpc_id", vpc.outputs["vpc_id"])
```

{{% /choosable %}}

{{% choosable language go %}}

```go
package main

import (
	"github.com/pulumi/pulumi-hcl/sdk/go/hcl"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		vpc, err := hcl.NewModule(ctx, "vpc", &hcl.ModuleArgs{
			Source:  "terraform-aws-modules/vpc/aws",
			Version: pulumi.StringRef("5.0.0"),
			Inputs: pulumi.Map{
				"name": pulumi.String("example-vpc"),
				"cidr": pulumi.String("10.0.0.0/16"),
			},
		})
		if err != nil {
			return err
		}
		ctx.Export("vpcId", vpc.Outputs.MapIndex(pulumi.String("vpc_id")))
		return nil
	})
}
```

{{% /choosable %}}

{{% choosable language csharp %}}

```csharp
using System.Collections.Generic;
using Pulumi;
using Pulumi.Hcl;

return await Deployment.RunAsync(() =>
{
    var vpc = new Module("vpc", new ModuleArgs
    {
        Source = "terraform-aws-modules/vpc/aws",
        Version = "5.0.0",
        Inputs =
        {
            { "name", "example-vpc" },
            { "cidr", "10.0.0.0/16" },
        },
    });

    return new Dictionary<string, object?>
    {
        ["vpcId"] = vpc.Outputs.Apply(o => o["vpc_id"]),
    };
});
```

{{% /choosable %}}

{{% choosable language java %}}

```java
package myproject;

import com.pulumi.Pulumi;
import com.pulumi.hcl.Module;
import com.pulumi.hcl.ModuleArgs;
import java.util.Map;

public class App {
    public static void main(String[] args) {
        Pulumi.run(ctx -> {
            var vpc = new Module("vpc", ModuleArgs.builder()
                .source("terraform-aws-modules/vpc/aws")
                .version("5.0.0")
                .inputs(Map.of(
                    "name", "example-vpc",
                    "cidr", "10.0.0.0/16"))
                .build());

            ctx.export("vpcId", vpc.outputs().applyValue(o -> o.get("vpc_id")));
        });
    }
}
```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
name: example
runtime: yaml
resources:
  vpc:
    type: hcl:index:Module
    properties:
      source: terraform-aws-modules/vpc/aws
      version: "5.0.0"
      inputs:
        name: example-vpc
        cidr: 10.0.0.0/16
outputs:
  vpcId: ${vpc.outputs["vpc_id"]}
```

{{% /choosable %}}

## Strongly typed SDKs

The `Module` resource above takes its `source` at runtime, so its inputs and
outputs are untyped maps. To get a strongly typed SDK for a specific module,
generate a parameterized provider for it:

```bash
pulumi package add hcl module <source> [version] [--name <name>]
```

For example, `pulumi package add hcl module terraform-aws-modules/vpc/aws 5.0.0`
generates a local SDK whose resource exposes that module's variables and outputs
as first-class, typed properties in your language of choice.

By default the package is named after the module name and its system (`vpc-aws` in the
example above). Pass `--name` to choose a different package name:

```bash
pulumi package add hcl module terraform-aws-modules/vpc/aws 5.0.0 --name vpc
```
