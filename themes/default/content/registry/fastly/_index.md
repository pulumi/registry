---
title: Fastly
meta_desc: This page provides an overview of the Fastly Provider for Pulumi.
layout: overview
---

The Fastly provider for Pulumi can be used to provision any of the cloud resources available in [Fastly](https://www.fastly.com/).
The Fastly provider must be configured with credentials to deploy and update resources in Fastly.

## Example

{{< chooser language "javascript,typescript,python,go,csharp" >}}

{{% choosable language javascript %}}

```javascript
const fastly = require("@pulumi/fastly")

const service = new fastly.Servicev1("my-service", {
  backends: [{
       address: "127.0.0.1",
       name: "localhost",
       port: 80
   }],
  domains: [{
       comment: "demo",
       name: "demo.pulumi.com",
  }],
  forceDestroy: true,
});
```

{{% /choosable %}}
{{% choosable language typescript %}}

```typescript
import * as fastly from "@pulumi/fastly";

const service = new fastly.Servicev1("my-service", {
  backends: [{
       address: "127.0.0.1",
       name: "localhost",
       port: 80
   }],
  domains: [{
       comment: "demo",
       name: "demo.pulumi.com",
  }],
  forceDestroy: true,
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_fastly as fastly

service = fastly.Servicev1("my-service",
  backends=[{
    "address": "127.0.0.1",
    "name": "localhost",
    "port": "80",
  }],
  domains=[{
    "comment": "demo",
    "name": "demo.pulumi.com",
  }],
  force_destroy="true",
)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	fastly "github.com/pulumi/pulumi-fastly/sdk/v3/go/fastly"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		service, err := fastly.NewServicev1(ctx, "test", &fastly.Servicev1Args{
			Backends: fastly.Servicev1BackendArray{
				fastly.Servicev1BackendArgs{
					Address: pulumi.String("127.0.0.1"),
					Name:    pulumi.String("localhost"),
					Port:    pulumi.Int(80),
				},
			},
			Domains: fastly.Servicev1DomainArray{
				fastly.Servicev1DomainArgs{
					Comment: pulumi.String("demo"),
					Name:    pulumi.String("demo.pulumi.com"),
				},
			},
			ForceDestroy: pulumi.Bool(true),
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
using Pulumi.Fastly;

class Program
{
    static Task Main() =>
        Deployment.Run(() => {
            var service = new Servicev1("test", new Servicev1Args
            {
                Backends = new Servicev1BackendsArgs{
                    {
                        Address="127.0.0.1",
                        Name="localhost",
                        Port=80
                    },
                },
                Domains = new Servicev1DomainsArgs{
                    {
                        Comment = "demo",
                        Name = "demo.pulumi.com"
                    },
                },
                ForceDestroy = true,
            });
        });
}
```

{{% /choosable %}}

{{< /chooser >}}

## Libraries

The following packages are available in packager managers:

* JavaScript/TypeScript: [`@pulumi/fastly`](https://www.npmjs.com/package/@pulumi/fastly)
* Python: [`pulumi-fastly`](https://pypi.org/project/pulumi-fastly/)
* Go: [`github.com/pulumi/pulumi-fastly/sdk/v3/go/fastly`](https://github.com/pulumi/pulumi-fastly)
* .NET: [`Pulumi.Fastly`](https://www.nuget.org/packages/Pulumi.Fastly)

The Fastly provider is open source and available in the [pulumi/pulumi-fastly](https://github.com/pulumi/pulumi-fastly) repo.
