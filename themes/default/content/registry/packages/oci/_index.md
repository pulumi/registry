---
title: Oracle Cloud Infrastructure
meta_desc: Provides an overview of the Oracle Cloud Infrastructure Provider for Pulumi.
layout: package
---

The Oracle Cloud Infrastructure (OCI) provider for Pulumi can be used to provision any of the resources available in [OCI](https://www.oracle.com/cloud).
The OCI provider must be configured with credentials to deploy and update resources in OCI.

## Example

{{< chooser language "typescript,python,go,csharp,java" >}}

{{% choosable language javascript %}}

```typescript
import * as oci from "@pulumi/oci";

const testUser = new oci.identity.User("testUser", {
    compartmentId: _var.tenancy_ocid,
    description: _var.user_description,
    definedTags: {
        "Operations.CostCenter": "42",
    },
    email: _var.user_email,
    freeformTags: {
        Department: "Finance",
    },
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_oci as oci

test_user = oci.identity.User("testUser",
                              compartment_id=var["tenancy_ocid"],
                              description=var["user_description"],
                              defined_tags={
                                  "Operations.CostCenter": "42",
                              },
                              email=var["user_email"],
                              freeform_tags={
                                  "Department": "Finance",
                              })

```

{{% /choosable %}}
{{% choosable language go %}}

```go
package main

import (
	"github.com/pulumi/pulumi-oci/sdk/go/oci/Identity"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := Identity.NewUser(ctx, "testUser", &Identity.UserArgs{
			CompartmentId: pulumi.Any(_var.Tenancy_ocid),
			Description:   pulumi.Any(_var.User_description),
			DefinedTags: pulumi.AnyMap{
				"Operations.CostCenter": pulumi.Any("42"),
			},
			Email: pulumi.Any(_var.User_email),
			FreeformTags: pulumi.AnyMap{
				"Department": pulumi.Any("Finance"),
			},
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
using System.Linq;
using Pulumi;
using Oci = Pulumi.Oci;

return await Deployment.RunAsync(() => 
{
    var testUser = new Oci.Identity.User("testUser", new()
    {
        CompartmentId = @var.Tenancy_ocid,
        Description = @var.User_description,
        DefinedTags = 
        {
            { "Operations.CostCenter", "42" },
        },
        Email = @var.User_email,
        FreeformTags = 
        {
            { "Department", "Finance" },
        },
    });

});
```

{{% /choosable %}}
{{% choosable language java %}}
```java
import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.oci.Identity.User;
import com.pulumi.oci.Identity.UserArgs;
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        var testUser = new User("testUser", UserArgs.builder()        
            .compartmentId(var_.tenancy_ocid())
            .description(var_.user_description())
            .definedTags(Map.of("Operations.CostCenter", "42"))
            .email(var_.user_email())
            .freeformTags(Map.of("Department", "Finance"))
            .build());

    }
}
```
{{% /choosable %}}

{{% choosable language yaml %}}
```yaml
resources:
  testUser:
    type: oci:Identity:User
    properties:
      #Required
      compartmentId: ${var.tenancy_ocid}
      description: ${var.user_description}
      #Optional
      definedTags:
        Operations.CostCenter: '42'
      email: ${var.user_email}
      freeformTags:
        Department: Finance
```
{{% /choosable %}}

{{< /chooser >}}
