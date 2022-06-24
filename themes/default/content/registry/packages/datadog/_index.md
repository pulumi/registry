---
title: Datadog
meta_desc: Provides an overview of the Datadog Provider for Pulumi.
layout: overview
---

The Datadog provider for Pulumi can be used to provision any of the cloud resources available in [Datadog](https://datadoghq.com/).
The Datadog provider must be configured with credentials to deploy and update resources in Datadog.

## Example

{{< chooser language "javascript,typescript,python,go,csharp" >}}

{{% choosable language javascript %}}

```javascript
const datadog = require("@pulumi/datadog")

const user = new datadog.User("my-user", {
  email: "new@example.com",
  name: "New User",
});
```

{{% /choosable %}}
{{% choosable language typescript %}}

```typescript
import * as datadog from "@pulumi/datadog";

const user = new datadog.User("my-policy", {
  email: "new@example.com",
  handle: "new@example.com",
  name: "New User",
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_datadog as datadog

user = datadog.User("my-policy",
  email="new@example.com",
  handle="new@example.com",
  name="New User",
)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	datadog "github.com/pulumi/pulumi-datadog/sdk/v3/go/datadog"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		user, err := datadog.NewUser(ctx, "my-user", &datadog.UserArgs{
			Email:  pulumi.String("new@example.com"),
			Handle: pulumi.String("new@example.com"),
			Name:   pulumi.String("New User"),
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
using Pulumi.Datadog;

await Deployment.RunAsync(() => 
{
  var user = new User("my-user", new UserArgs
  {
      Name = "New User",
      Email = "new@example.com",
      Handle = "new@example.com"
  });
});
```

{{% /choosable %}}

{{< /chooser >}}
