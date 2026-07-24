---
# WARNING: this file was fetched from https://raw.githubusercontent.com/rpothin/pulumi-powerplatform/v0.4.1/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/rpothin/pulumi-powerplatform/blob/v0.4.1/docs/_index.md
title: Power Platform
meta_desc: The Pulumi Power Platform provider manages Microsoft Power Platform resources.
layout: package
---

The Power Platform provider for Pulumi allows you to manage
[Microsoft Power Platform](https://powerplatform.microsoft.com/) resources
using infrastructure as code.

Power Platform must be configured with Azure AD credentials to deploy and manage resources. See [Installation & Configuration](./installation-configuration/) for instructions.

## Overview

The provider supports:

- **Environment** — create, update, and delete Power Platform environments
- **Environment Group** — organise environments into groups
- **Environment Settings** — govern per-environment settings (audit, logging, etc.)
- **Environment Backup** — manage environment backups
- **Managed Environment** — enable/disable managed-environment governance
- **Tenant Settings** — manage tenant-wide Power Platform settings
- **Data Record** — manage Dataverse table rows used by advanced resources and components
- **DLP Policy** — configure data loss prevention policies
- **Billing Policy** — manage pay-as-you-go billing policies
- **Pipeline Sharing** — share deployment pipelines with teams
- **Role Assignment** — assign security roles to principals
- **ISV Contract** — manage ISV contract resources
- **Component Resources** — AVM-aligned modules for environments, DLP policies, tenant settings, and deployment pipelines

Data-source functions are available for listing environments, connectors,
apps, flows, DLP policies, security roles, Dataverse records, and for
generating DLP policy migration configurations.

## Example

{{< chooser language "python,typescript,go,csharp,java,yaml" >}}

{{% choosable language python %}}

```python
import pulumi
import rpothin_powerplatform as pp

env = pp.Environment(
    "dev",
    display_name="Development",
    location="unitedstates",
    environment_type="Sandbox",
)

pulumi.export("envId", env.id)
```

{{% /choosable %}}

{{% choosable language typescript %}}

```typescript
import * as pp from "@rpothin/powerplatform";

const env = new pp.Environment("dev", {
    displayName: "Development",
    location: "unitedstates",
    environmentType: "Sandbox",
});

export const envId = env.id;
```

{{% /choosable %}}

{{% choosable language go %}}

```go
package main

import (
	pp "github.com/rpothin/pulumi-powerplatform/sdk/go/powerplatform"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		env, err := pp.NewEnvironment(ctx, "dev", &pp.EnvironmentArgs{
			DisplayName:     pulumi.String("Development"),
			Location:        pulumi.String("unitedstates"),
			EnvironmentType: pulumi.String("Sandbox"),
		})
		if err != nil {
			return err
		}
		ctx.Export("envId", env.ID())
		return nil
	})
}
```

{{% /choosable %}}

{{% choosable language csharp %}}

```csharp
using Pulumi;
using Pulumi.Powerplatform;

return await Deployment.RunAsync(() =>
{
    var env = new Environment("dev", new EnvironmentArgs
    {
        DisplayName = "Development",
        Location = "unitedstates",
        EnvironmentType = "Sandbox",
    });

    return new Dictionary<string, object?>
    {
        ["envId"] = env.Id,
    };
});
```

{{% /choosable %}}

{{% choosable language java %}}

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import io.github.rpothin.powerplatform.Environment;
import io.github.rpothin.powerplatform.EnvironmentArgs;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        var env = new Environment("dev", EnvironmentArgs.builder()
            .displayName("Development")
            .location("unitedstates")
            .environmentType("Sandbox")
            .build());

        ctx.export("envId", env.id());
    }
}
```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
name: my-powerplatform-stack
runtime: yaml
resources:
  dev:
    type: powerplatform:index:Environment
    properties:
      displayName: Development
      location: unitedstates
      environmentType: Sandbox
outputs:
  envId: ${dev.id}
```

{{% /choosable %}}

{{< /chooser >}}
