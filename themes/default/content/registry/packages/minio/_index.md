---
title: Minio
meta_desc: Provides an overview of the Minio Provider for Pulumi.
layout: overview
---

The Minio provider for Pulumi can be used to provision any of the cloud resources available in [Minio](https://min.io/).
The Minio provider must be configured with credentials to deploy and update resources in Minio.

## Example

{{< chooser language "javascript,typescript,python,go,csharp" >}}

{{% choosable language javascript %}}

```javascript
const minio = require("@pulumi/minio")

const user = new minio.IamUser("ts-user");
```

{{% /choosable %}}
{{% choosable language typescript %}}

```typescript
import * as minio from "@pulumi/minio";

const user = new minio.IamUser("ts-user")
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_minio as minio

user = minio.IamUser("python-user")
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
    minio "github.com/pulumi/pulumi-minio/sdk/go/minio"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		user, err := minio.NewIamUser(ctx, "go-user", nil)
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
using Pulumi.Minio;

class Program
{
    static Task Main() =>
        Deployment.Run(() => {
            var user = new Minio.IamUser("csharp-user");
        });
}
```

{{% /choosable %}}

{{< /chooser >}}
