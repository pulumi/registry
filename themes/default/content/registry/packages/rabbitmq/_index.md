---
title: RabbitMQ
meta_desc: This page provides an overview of the RabbitMQ Provider for Pulumi.
layout: overview
---

The RabbitMQ provider for Pulumi can be used to provision any of the resources available for RabbitMQ.
The RabbitMQ provider must be configured with credentials to deploy and update resources in Fastly.

## Example

{{< chooser language "javascript,typescript,python,go,csharp" >}}

{{% choosable language javascript %}}

```javascript
const rabbitmq = require("@pulumi/rabbitmq")

const user = new rabbitmq.User("user", {
  password: "MyPassword1234!"
});
```

{{% /choosable %}}
{{% choosable language typescript %}}

```typescript
import * as rabbitmq from "@pulumi/rabbitmq";

const user = new rabbitmq.User("user", {
  password: "MyPassword1234!"
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_rabbitmq as rabbitmq

user = rabbitmq.User("user",
  password="MyPassword1234!"
)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	rabbitmq "github.com/pulumi/pulumi-rabbitmq/sdk/v3/go/rabbitmq"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		user, err := rabbitmq.NewUser(ctx, "user", &rabbitmq.UserArgs{
			Password: pulumi.String("MyPassword1234!"),
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
using Pulumi.Rabbitmq;

class Program
{
    static Task Main() =>
        Deployment.Run(() => {
            var user = new User("user", new UserArgs
            {
                Password = "MyPassword1234!"
            });
        });
}
```

{{% /choosable %}}

{{< /chooser >}}
