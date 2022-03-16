---
title: Keycloak
meta_desc: Provides an overview of the Keycloak Provider for Pulumi.
layout: overview
---

{{% overview-description %}}
The Keycloak provider for Pulumi can be used to provision any of the cloud resources available via [Keycloak](https://rancher.com/).
The Keycloak provider must be configured with credentials to deploy and update resources for Rancher.
{{% /overview-description %}}

## Example

{{< chooser language "javascript,typescript,python,go,csharp" >}}

{{% choosable language javascript %}}

```javascript
const keycloak = require("@pulumi/keycloak")

const realm = new keycloak.Realm("new-typescript-realm", {
  realm: "my-example-typescript-realm"
});
```

{{% /choosable %}}
{{% choosable language typescript %}}

```typescript
import * as keycloak from "@pulumi/keycloak";

const realm = new keycloak.Realm("new-typescript-realm", {
  realm: "my-example-typescript-realm"
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_keycloak as keycloak

realm = keycloak.Realm("new-python-realm",
                       realm="my-example-python-realm"
                       )
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	keycloak "github.com/pulumi/pulumi-keycloak/sdk/v4/go/keycloak"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
        realm, err := keycloak.NewRealm(ctx, "new-go-realm", &keycloak.RealmArgs{
            Realm: pulumi.String("my-example-go-realm"),
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
using Pulumi.Keycloak;

class Program
{
    static Task Main() =>
        Deployment.Run(() => {
            var realm = new Keycloak.Realm("new-csharp-realm", new Keycloak.RealmArgs
            {
                RealmName = "my-example-csharp-realm",
            });
        });
}
```

{{% /choosable %}}

{{< /chooser >}}
