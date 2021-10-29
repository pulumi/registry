---
title: Opsgenie
meta_desc: Provides an overview of the Opsgenie Provider for Pulumi.
layout: overview
---

The Opsgenie provider for Pulumi can be used to provision any of the resources available for Opsgenie.
The Opsgenie provider must be configured with credentials to deploy and update resources in Opsgenie.

## Example

{{< chooser language "javascript,typescript,python,go,csharp" >}}

{{% choosable language javascript %}}

```javascript
const opsgenie = require("@pulumi/opsgenie")

const team = new opsgenie.Team('test');
```

{{% /choosable %}}
{{% choosable language typescript %}}

```typescript
import * as opsgenie from "@pulumi/opsgenie";

const team = new opsgenie.Team('test');
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_opsgenie as opsgenie

team = opsgenie.Team("test")
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi-opsgenie/sdk/go/opsgenie"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		team, err := opsgenie.NewTeam(ctx, "my-team", &opsgenie.TeamArgs{})
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
using Pulumi.Opsgenie;

class Program
{
    static Task Main() =>
        Deployment.Run(() => {
            var team = new Team("my-team", new TeamArgs{});
        });
}
```

{{% /choosable %}}

{{< /chooser >}}
