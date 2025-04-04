---
title: Terraform
meta_desc: Consume Terraform state in Pulumi
layout: package
---

The Terraform provider for Pulumi is a powerful integration that enables developers and infrastructure engineers to seamlessly incorporate Terraform state management into their Pulumi-driven infrastructure as code workflows. This provider allows you to directly consume and interact with the outputs contained in Terraform state files, whether they reside locally or in remote backends, right from within your Pulumi programs.

By leveraging this provider, you can easily transition existing Terraform-managed infrastructure into your Pulumi deployments or augment your current Pulumi projects with resources managed by Terraform. The provider is designed to simplify the access to Terraform state, offering flexibility in consumption through specialized functions tailored for different backend types, which significantly streamlines hybrid infrastructure management scenarios.

With comprehensive support for multiple programming languages including C#, Go, Java, Node.js, and Python, this provider ensures that developers can integrate Terraform state data into their preferred coding environments. It adheres to the native idioms and best practices of each language, providing a natural and efficient experience for Pulumi users across diverse technical ecosystems.

Whether you're aiming to unify your infrastructure management workflows, gradually adopt Pulumi features, or maintain existing Terraform states while extending functionality with Pulumi's rich cloud APIs, the Terraform provider serves as a bridge that facilitates these transitions with minimal friction and robust functionality.


## Example

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}

{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as terraform from "@pulumi/terraform";

const state = pulumi.output(terraform.localStateReference({
    path: "./terraform.0-12-24.tfstate",
}));

export const outputs = state.outputs;
```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi
from pulumi_terraform import local_state_reference

state = local_state_reference.LocalStateReference(
    path="./terraform.0-12-24.tfstate"
)

pulumi.export("state", state.outputs)
```

{{% /choosable %}}

{{% choosable language go %}}

```go
package main

import (
	"github.com/pulumi/pulumi-terraform/sdk/go/terraform"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		state, err := terraform.LocalStateReference(ctx,
			&terraform.LocalStateReferenceArgs{
				Path: pulumi.String("./terraform.0-12-24.tfstate"),
			},
		)
		if err != nil {
			return err
		}

		ctx.Export("state", state.Outputs)
		return nil
	})
}
```

{{% /choosable %}}

{{% choosable language csharp %}}

```csharp
using Pulumi;
using Pulumi.Terraform;

class MyStack : Stack
{
    public MyStack()
    {
        var state = Output.Create(InvokeAsync<Inputs.TerraformLocalStateReferenceInputs, Outputs.TerraformLocalStateReferenceOutputs>(
            "terraform:index:localStateReference",
            new Inputs.TerraformLocalStateReferenceInputs
            {
                Path = "./terraform.0-12-24.tfstate",
            })).Apply(result => result.Outputs);
    }
}
```

{{% /choosable %}}

{{% choosable language java %}}

```java
import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.deployment.InvokeOptions;
import com.pulumi.terraform.TerraformFunctions;
import com.pulumi.terraform.inputs.LocalStateReferenceArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.ResourceType;

@ResourceType(type = "terraform:index:localStateReference")
public class TerraformLocalStateWithYaml {

    public static void main(String[] args) {
        Pulumi.run(TerraformLocalStateWithYaml::main);
    }

    public static void main(Context ctx) {
        Output<Object> state = TerraformFunctions.localStateReference(LocalStateReferenceArgs.builder()
                .path("./terraform.0-12-24.tfstate")
                .build(), InvokeOptions.Empty);

        ctx.export("state", state);
    }
}
```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
name: terraform-local-state-with-yaml
runtime: yaml
outputs:
  state:
    fn::invoke:
      function: terraform:index:localStateReference
      arguments:
        path: ./terraform.0-12-24.tfstate
      return: outputs
```

{{% /choosable %}}

{{< /chooser >}}
