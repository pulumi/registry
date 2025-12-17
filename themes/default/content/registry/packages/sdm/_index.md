---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pierskarsenbarg/pulumi-sdm/v1.33.0/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/pierskarsenbarg/pulumi-sdm/blob/v1.33.0/docs/_index.md
title: StrongDM
meta_desc: Provides an overview of the StrongDM Provider for Pulumi.
layout: package
---

The StrongDM provider for Pulumi can be used to provision any of the cloud resources available in [StrongDM](https://www.strongdm.com/).
The StrongDM provider must be configured with credentials to deploy and update resources in StrongDM.

## Example

{{< chooser language "typescript,python,go,csharp" >}}

{{% choosable language typescript %}}

```typescript
import * as sdm from "@pierskarsenbarg/sdm";

const account = new sdm.Account("account", {
    user: {
        firstName: "Alice",
        lastName: "Bob",
        email: "alicebob@email.com"
    }
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pierskarsenbarg_pulumi_sdm as sdm

account = sdm.Account("account",
  user=AccountUserArgs(
    first_name="Alice",
    last_name="Bob",
    email="alicebob@email.com"
  )
)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"github.com/pierskarsenbarg/pulumi-sdm/sdk/go/sdm"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sdm.NewAccount(ctx, "account", &sdm.AccountArgs{
			User: &sdm.AccountUserArgs{
				FirstName: pulumi.String("Alice"),
				LastName:  pulumi.String("Bob"),
				Email:     pulumi.String("alicebob@email.com"),
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
using Pulumi;
using PiersKarsenbarg.Sdm;
using PiersKarsenbarg.Sdm.Inputs;

return await Deployment.RunAsync(() =>
{
   var account = new Account("account", new()
   {
      User = new AccountUserArgs
      {
         FirstName = "Alice",
         LastName = "Bob",
         Email = "alicebob@email.com"
      }
   });
});
```

{{% /choosable %}}

{{< /chooser >}}
