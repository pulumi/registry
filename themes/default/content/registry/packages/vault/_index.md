---
title: HashiCorp Vault
meta_desc: Provides an overview of the HashiCorp Vault Provider for Pulumi.
layout: overview
---

{{% overview-description %}}
The HashiCorp Vault provider for Pulumi can be used to provision any of the resources available in [Vault](https://www.vaultproject.io/).
{{% /overview-description %}}

## Example

{{< chooser language "javascript,typescript,python,go,csharp" >}}

{{% choosable language javascript %}}

```javascript
const vault = require("@pulumi/vault")

const be = new vault.AuthBackend("example", {
  type: "github"
});
```

{{% /choosable %}}
{{% choosable language typescript %}}

```typescript
import * as vault from "@pulumi/vault";

const be = new vault.AuthBackend("example", {
  type: "github"
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_vault as vault

be = vault.AuthBackend("example",
  type='github'
)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	vault "github.com/pulumi/pulumi-vault/sdk/v4/go/vault"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		be, err := vault.NewAuthBackend(ctx, "example", &vault.AuthBackendArgs{
			Type: pulumi.String("github"),
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
using Pulumi.Vault;

class Program
{
    static Task Main() =>
        Deployment.Run(() => {
            var be = new Vault.AuthBackend("example", new Vault.AuthBackendArgs
            {
                Type = "github",
            });
        });
}
```

{{% /choosable %}}

{{< /chooser >}}
