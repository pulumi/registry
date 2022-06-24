---
title: Auth0
meta_desc: Provides an overview of the Auth0 Provider for Pulumi.
layout: overview
---

The Auth0 provider for Pulumi can be used to provision any of the cloud resources available in [Auth0](https://auth0.com/).
The Auth0 provider must be configured with credentials to deploy and update resources in Auth0.

## Example

{{< chooser language "javascript,typescript,python,go,csharp" >}}

{{% choosable language javascript %}}

```javascript
const auth0 = require("@pulumi/auth0")

const user = new auth0.User("my-demo-user", {
    username: "joebloggs",
    password: "Password1234!",
    email: "joebloggs@mycompany.com",
    connectionName: "Username-Password-Authentication",
});
```

{{% /choosable %}}
{{% choosable language typescript %}}

```typescript
import * as auth0 from "@pulumi/auth0";

const user = new auth0.User("my-demo-user", {
    username: "joebloggs",
    password: "Password1234!",
    email: "joebloggs@mycompany.com",
    connectionName: "Username-Password-Authentication",
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_auth0 as auth0

user = auth0.User("my-demo-user",
  username="joebloggs",
  password="Password1234!",
  email="joebloggs@mycompany.com",
  connection_name="Username-Password-Authentication",
)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	auth0 "github.com/pulumi/pulumi-auth0/sdk/v2/go/auth0"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		user, err := auth.NewUser(ctx, "my-demo-user", &auth0.UserArgs{
			Username:       pulumi.String("joebloggs"),
			Password:       pulumi.String("Password1234!"),
			Email:          pulumi.String("joebloggs@mycompany.com"),
			ConnectionName: pulumi.String("Username-Password-Authentication"),
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
using Auth0 = Pulumi.Auth0;

await Deployment.RunAsync(() =>
{
    var user = new Auth0.User("test", new Auth0.UserArgs
    {
        Username = "joebloggs",
        Password = "Password1234!",
        Email = "joebloggs@mycompany.com",
        ConnectionName = "Username-Password-Authentication",
    });
});
```

{{% /choosable %}}

{{< /chooser >}}
