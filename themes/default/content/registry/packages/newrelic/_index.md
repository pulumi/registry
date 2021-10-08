---
title: New Relic
meta_desc: This page provides an overview of the New Relic Provider for Pulumi.
layout: overview
---

The New Relic provider for Pulumi can be used to provision any of the cloud resources available in [New Relic](https://newrelic.com/).
The New Relic provider must be configured with credentials to deploy and update resources in New Relic.

## Example

{{< chooser language "javascript,typescript,python,go,csharp" >}}

{{% choosable language javascript %}}

```javascript
const newrelic = require("@pulumi/newrelic")

const policy = new newrelic.AlertPolicy("my-policy");
```

{{% /choosable %}}
{{% choosable language typescript %}}

```typescript
import * as newrelic from "@pulumi/newrelic";

const policu = new newrelic.AlertPolicy("my-policy");
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_newrelic as newrelic

policy = newrelic.AlertPolicy("my-policy")
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	newrelic "github.com/pulumi/pulumi-newrelic/sdk/v4/go/newrelic"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		policy, err := newrelic.NewAlertPolicy(ctx, "my-policy", &newrelic.AlertPolicyArgs{})
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
using Pulumi.Newrelic;

class Program
{
    static Task Main() =>
        Deployment.Run(() => {
            var policy = new AlertPolicy("my-policy");
        });
}
```

{{% /choosable %}}

{{< /chooser >}}

## Libraries

The following packages are available in packager managers:

* JavaScript/TypeScript: [`@pulumi/newrelic`](https://www.npmjs.com/package/@pulumi/newrelic)
* Python: [`pulumi-newrelic`](https://pypi.org/project/pulumi-newrelic/)
* Go: [`github.com/pulumi/pulumi-newrelic/sdk/v4/go/newrelic`](https://github.com/pulumi/pulumi-newrelic)
* .NET: [`Pulumi.Newrelic`](https://www.nuget.org/packages/Pulumi.Newrelic)

The New Relic provider is open source and available in the [pulumi/pulumi-newrelic](https://github.com/pulumi/pulumi-newrelic) repo.
