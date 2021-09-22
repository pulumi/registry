---
title: DigitalOcean
meta_desc: This page provides an overview of the DigitalOcean Provider for Pulumi.
layout: overview
aliases:
    - "/docs/reference/clouds/digitalocean/"
    - "/docs/intro/cloud-providers/digitalocean/"
---

The DigitalOcean provider for Pulumi can be used to provision any of the cloud resources available in [DigitalOcean](https://www.digitalocean.com/).
The DigitalOcean provider must be configured with credentials to deploy and update resources in a DigitalOcean cloud.

## Example

{{< chooser language "javascript,typescript,python,go,csharp" >}}

{{% choosable language javascript %}}

```javascript
const do = require("@pulumi/digitalocean")

const domain = new do.Domain("test", {
  name: "mydomain.com",
  ipAddress: "192.168.10.10",
});
```

{{% /choosable %}}
{{% choosable language typescript %}}

```typescript
import * as digitalocean from "@pulumi/digitalocean";

const instance = new digitalocean.Domain("test", {
  name: "mydomain.com",
  ipAddress: "192.168.10.10",
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_digitalocean as do

instance = do.Domain("test",
  name='mydomain.com',
  ip_address='192.168.10.10'
)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	do "github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		domain, err := do.NewDomain(ctx, "test", &do.DomainArgs{
			Name:      pulumi.String("mydomain.com"),
			IpAddress: pulumi.String("192.168.10.10"),
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
using Pulumi.DigitalOcean;

class Program
{
    static Task Main() =>
        Deployment.Run(() => {
            var instance = new DigitalOcean.Domain("test", new DigitalOcean.DomainArgs
            {
                Name = "mydomain.com",
                IpAddress = "192.168.10.10",
            });
        });
}
```

{{% /choosable %}}

{{< /chooser >}}

## Libraries

The following packages are available in packager managers:

* JavaScript/TypeScript: [`@pulumi/digitalocean`](https://www.npmjs.com/package/@pulumi/digitalocean)
* Python: [`pulumi-digitalocean`](https://pypi.org/project/pulumi-digitalocean/)
* Go: [`github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean`](https://github.com/pulumi/pulumi-digitalocean)
* .NET: [`Pulumi.DigitalOcean`](https://www.nuget.org/packages/Pulumi.DigitalOcean)

The DigitalOcean provider is open source and available in the [pulumi/pulumi-digitalocean](https://github.com/pulumi/pulumi-digitalocean) repo.
