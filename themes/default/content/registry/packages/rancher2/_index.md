---
title: Rancher2
meta_desc: Provides an overview of the Rancher2 Provider for Pulumi.
layout: overview
---

{{% overview-description %}}
The Rancher2 provider for Pulumi can be used to provision any of the cloud resources available via [Rancher](https://rancher.com/).
The Rancher2 provider must be configured with credentials to deploy and update resources for Rancher.
{{% /overview-description %}}

## Example

{{< chooser language "javascript,typescript,python,go,csharp" >}}

{{% choosable language javascript %}}

```javascript
const rancher2 = require("@pulumi/rancher2")

const myUser = new rancher2.User("my-user", {
  name: "Foo user",
  username: "foo",
  password: "initialPassw0rd",
  enabled: true,
});
```

{{% /choosable %}}
{{% choosable language typescript %}}

```typescript
import * as rancher2 from "@pulumi/rancher2";

const myUser = new rancher2.User("my-user", {
  name: "Foo user",
  username: "foo",
  password: "initialPassw0rd",
  enabled: true,
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_rancher2 as rancher2

user = rancher2.User("my-user",
  name="Foo user",
  username="foo",
  password="initialPassw0rd",
  enabled=true
)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	rancher2 "github.com/pulumi/pulumi-rancher2/sdk/v3/go/rancher2"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		user, err := rancher2.NewUser(ctx, "my-user", &rancher2.UserArgs{
			Name:     pulumi.String("Foo user"),
			Username: pulumi.String("foo"),
			Password: pulumi.String("initialPassw0rd"),
			Enabled:  pulumi.Bool(true),
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
using Pulumi.Rancher2;

class Program
{
    static Task Main() =>
        Deployment.Run(() => {
            var user = new User("my-user", new UserArgs
            {
                Name = "Foo user",
                Username = "foo",
                Password = "initialPassw0rd",
                Enabled = true,
            });
        });
}
```

{{% /choosable %}}

{{< /chooser >}}
