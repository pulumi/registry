---
title: onelogin
meta_desc: Provides an overview of the onelogin Provider for Pulumi.
layout: overview
---

The onelogin provider for Pulumi can be used to provision any of the cloud resources available in [onelogin](https://www.onelogin.com/).
The onelogin provider must be configured with credentials to deploy and update resources in onelogin.

## Example

{{< chooser language "javascript,typescript,python,go,csharp" >}}

{{% choosable language javascript %}}

```javascript
const onelogin = require("@pulumi/onelogin")

const user = new onelogin.User("user", {
    email: "user@mycompany.com",
    username: "myusername"
});
```

{{% /choosable %}}
{{% choosable language typescript %}}

```typescript
import * as onelogin from "@pulumi/onelogin";

const user = new onelogin.User("user", {
    email: "user@mycompany.com",
    username: "myusername"
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_onelogin as onelogin

zone = onelogin.User("user",
  email="user@mycompany.com",
  user="myusername"
)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	onelogin "github.com/pulumi/pulumi-onelogin/sdk/go/onelogin"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		user, err := onelogin.NewUser(ctx, "user", &onelogin.UserArgs{
			Email: pulumi.String("user@mycompany.com"),
			Username: pulumi.String("myusername")
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
using Pulumi.Onelogin;

class Program
{
    static Task Main() =>
        Deployment.Run(() => {
            var user = new User("user", new UserArgs
            {
                Email = "user@mycompany.com",
                Username = "myusername"
            });
        });
}
```

{{% /choosable %}}

{{< /chooser >}}
