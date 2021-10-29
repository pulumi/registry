---
title: Snowflake
meta_desc: This page provides an overview of the Snowflake Provider for Pulumi.
layout: overview
---

The Snowflake provider for Pulumi can be used to provision any of the cloud resources available in [Snowflake](https://www.snowflake.com/).
The Snowflake provider must be configured with credentials to deploy and update resources in Snowflake.

## Example

{{< chooser language "javascript,typescript,python,go,csharp" >}}

{{% choosable language javascript %}}

```javascript
const snowflake = require("@pulumi/snowflake")

const user = new snowflake.User("ts-user")
```

{{% /choosable %}}
{{% choosable language typescript %}}

```typescript
import * as snowflake from "@pulumi/snowflake";

const user = new snowflake.User("ts-user")
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_snowflake as snowflake

user = snowflake.User("py-user")
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
    snowflake "github.com/pulumi/pulumi-snowflake/sdk/go/snowflake"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
        user, err := snowflake.NewUser(ctx, "go-user", nil)
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
using Pulumi.Snowflake;

class Program
{
    static Task Main() =>
        Deployment.Run(() => {
            var user = new Snowflake.User("cs-user", new Snowflake.UserArgs{});
        });
}
```

{{% /choosable %}}

{{< /chooser >}}
