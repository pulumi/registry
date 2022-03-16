---
title: TLS
meta_desc: Provides an overview of the TLS Provider for Pulumi.
layout: overview
---

{{% overview-description %}}
The TLS provider for Pulumi can be used to help to create TLS keys and certitifcate for use with Pulumi resources.
{{% /overview-description %}}

## Example

{{< chooser language "javascript,typescript,python,go,csharp" >}}

{{% choosable language javascript %}}

```javascript
const tls = require("@pulumi/tls")

const key = new tls.PrivateKey("my-private-key", {
    algorithm: "ECDSA",
    ecdsaCurve: "P384",
});
```

{{% /choosable %}}
{{% choosable language typescript %}}

```typescript
import * as tls from "@pulumi/tls";

const key = new tls.PrivateKey("my-private-key", {
    algorithm: "ECDSA",
    ecdsaCurve: "P384",
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_tls as tls

key = tls.PrivateKey("my-private-key",
  algorithm="ECDSA",
  ecdsa_curve="P384"
)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	tls "github.com/pulumi/pulumi-tls/sdk/v4/go/tls"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		key, err := tls.NewPrivateKey(ctx, "my-private-key", &tls.PrivateKeyArgs{
			Algorithm:  pulumi.String("ECDSA"),
			EcdsaCurve: pulumi.String("P384"),
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
using Pulumi.Tls;

class Program
{
    static Task Main() =>
        Deployment.Run(() => {
            var key = new PrivateKey("my-private-key", new PrivateKeyArgs{
                Algorithm = "ECDSA",
                EcdsaCurve = "P384",
            });
        });
}
```

{{% /choosable %}}

{{< /chooser >}}
