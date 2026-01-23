---
title: Yandex
meta_desc: Provides an overview of the Yandex Cloud Provider for Pulumi.
layout: package
---

The Yandex Cloud provider for Pulumi can be used to provision cloud resources available in [Yandex Cloud](https://cloud.yandex.com/).
The Yandex Cloud provider must be configured with credentials to deploy and update resources in Yandex Cloud.

## Example

{{< chooser language "typescript,python,go,csharp" >}}

{{% choosable language typescript %}}

```typescript
import * as yandex from "@pulumi/yandex";

const defaultVpcNetwork = new yandex.VpcNetwork("pulumi-acc-test");
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_yandex as yandex

default_vpc_network = yandex.VpcNetwork("default")
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"github.com/pulumi/pulumi-yandex/sdk/go/yandex"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
        _, err := yandex.NewVpcNetwork(ctx, "defaultVpcNetwork", nil)
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
using Yandex = Pulumi.Yandex;

class Program
{
    static Task Main() =>
        Deployment.Run(() => {
            var defaultVpcNetwork = new Yandex.VpcNetwork("default", new Yandex.VpcNetworkArgs{});
        });
}
```

{{% /choosable %}}

{{< /chooser >}}
