---
title: Authress
meta_desc: Provides an overview of the Authress Provider for Pulumi.
layout: overview
---

The Authress provider for Pulumi can be used to provision any of the cloud resources available in [Authress](https://authress.io/).
The Authress provider must be configured with credentials to deploy and update resources in Authress.

## Example

{{< chooser language "typescript,python,go,csharp" >}}
{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as Authress from "@pulumi/authress"


const authressRole = new Authress.Role("TestRole", {
    roleId: 'test-role',
    name: "Test Role",
    description: "An example description for this Role",
    permissions: {
        "documents:read": {
            allow: true
        }
    }
});

export let authressTestRoleId = authressRole.roleId;
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi
import pulumi_authress as Authress

authressRole = Authress.Role("TestRole",
                              roleId="test-role",
                              name="Test Role",
                              )

pulumi.export("authressTestRoleId", authressRole.roleId)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	Authress "github.com/Authress/pulumi-authress/sdk/go/authress"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		authressRole, err := Authress.Role(ctx, "sks-cluster", &Authress.RoleArgs{
			RoleId: pulumi.String("test-role"),
            Name: pulumi.String("Test Role"),
		})
		if err != nil {
			return err
		}
		ctx.Export("authressTestRoleId", authressRole.RoleId)
		return nil
	})
}
```

{{% /choosable %}}

{{< /chooser >}}