---
title: Kong
meta_desc: Provides an overview of the Kong Provider for Pulumi.
layout: overview
---

The Kong provider for Pulumi can be used to provision any of the cloud resources available in [Kong](https://konghq.com/kong).
The Kong provider must be configured with credentials to deploy and update resources in Mailgun.

## Example

{{< chooser language "javascript,typescript,python,go,csharp" >}}

{{% choosable language javascript %}}

```javascript
const kong = require("@pulumi/kong")

const consumer = new kong.Consumer("my-consumer", {
  username: "my-username-1",
  customId: "123"
});
```

{{% /choosable %}}
{{% choosable language typescript %}}

```typescript
import * as kong from "@pulumi/kong";

const consumer = new kong.Consumer("my-consumer", {
  username: "my-username-1",
  customId: "123"
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_kong as kong

consumer = kong.Consumer("my-consumer",
  username="my-username-1",
  custom_id="123")
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	kong "github.com/pulumi/pulumi-kong/sdk/v4/go/kong"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		consumer, err := kong.NewConsumer(ctx, "test-route", &kong.ConsumerArgs{
            CustomId: pulumi.String("123"),
            Username: pulumi.String("my-username-1"),
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
using Pulumi.Kong;

class Program
{
    static Task Main() =>
        Deployment.Run(() => {
            var consumer = new Kong.Consumer("consumer", new Kong.ConsumerArgs
            {
                CustomId = "123",
                Username = "my-username-1",
            });
        });
}
```

{{% /choosable %}}

{{< /chooser >}}
