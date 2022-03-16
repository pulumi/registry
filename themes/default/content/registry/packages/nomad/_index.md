---
title: HashiCorp Nomad
meta_desc: Provides an overview of the HashiCorp Nomad Provider for Pulumi.
layout: overview
---

{{% overview-description %}}
The HashiCorp Nomad provider for Pulumi can be used to provision any of the resources available in [Nomad](https://www.nomadproject.io/).
{{% /overview-description %}}

## Example

{{< chooser language "javascript,typescript,python,go,csharp" >}}

{{% choosable language javascript %}}

```javascript
const nomad = require("@pulumi/nomad")

const ns = new nomad.Namespace("test-ns-ts");
```

{{% /choosable %}}
{{% choosable language typescript %}}

```typescript
import * as nomad from "@pulumi/nomad";

const ns = new nomad.Namespace("test-ns-ts");
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi
import pulumi_nomad as nomad

ns = nomad.Namespace("test-ns-py")
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	nomad "github.com/pulumi/pulumi-nomad/sdk/go/nomad"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
        _, err := nomad.NewNamespace(ctx, "test-ns-go", nil)
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
using Pulumi.Nomad;

class Program
{
    static Task Main() =>
        Deployment.Run(() => {
            var ns = new Nomad.Namespace("test-ns-cs", new Nomad.NamespaceArgs{});
        });
}
```

{{% /choosable %}}

{{< /chooser >}}
