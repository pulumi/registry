---
title: Oracle Cloud Infrastructure
meta_desc: Provides an overview of the Oracle Cloud Infrastructure Provider for Pulumi.
layout: overview
---

The Oracle Cloud Infrastructure (OCI) provider for Pulumi can be used to provision any of the resources available in [OCI](https://www.oracle.com/cloud).
The OCI provider must be configured with credentials to deploy and update resources in OCI.

## Example

{{< chooser language "javascript,typescript,python,go,csharp" >}}

{{% choosable language javascript %}}

```javascript
const oci = require("@pulumi/oci")
const testUser = new oci.identity.User("testUser", {
    compartmentId: "<compartmentID>",
    email: "testuser@pulumi.com",
    description: "Test User Created by Pulumi"
});
```

{{% /choosable %}}
{{% choosable language typescript %}}

```typescript
import * as oci from "@pulumi/oci";
const testUser = new oci.identity.User("testUser", {
    compartmentId: "<compartmentID>",
    email: "testuser@pulumi.com",
    description: "Test User Created by Pulumi"
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_oci as oci
test_user = oci.identity.User("test-user",
                             compartment_id="<compartment_id>",
                             email="testuser@pulumi.com",
                             description="Test User Created by Pulumi")
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
    oci "github.com/pulumi/pulumi-oci/sdk/go/oci"
)
func main() {
    pulumi.Run(func(ctx *pulumi.Context) error {
        testUser, err := oci.Identity.NewUser(ctx, "test-user", &oci.Identity.UserArgs{
            CompartmentID: pulumi.String("<compartment_id>"),
            Email:         pulumi.String("testuser@pulumi.com"),
            Description:   pulumi.String("Test User Created by Pulumi"),
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
using Oci = Pulumi.Oci;
class Program
{
    static Task Main() =>
        Deployment.Run(() => {
            var testUser = new Oci.Identity.User("test-user", new Oci.Identity.UserArgs
            {
                CompartmentId = "<compartment_id>",
                Email = "testuser@pulumi.com",
                Description = "Test User Created by Pulumi",
            });
        });
}
```

{{% /choosable %}}

{{< /chooser >}}
