---
title: Alibaba Cloud
meta_desc: This page provides an overview of the Alibaba Cloud Provider for Pulumi.
menu:
  intro:
    parent: cloud-providers
    identifier: clouds-alicloud
    weight: 2
layout: overview
---

The Alibaba Cloud provider for Pulumi can be used to provision any of the cloud resources available in [Alibaba Cloud](https://www.alibabacloud.com/).
The Alibaba Cloud provider must be configured with credentials to deploy and update resources in Alibaba Cloud.

## Example

{{< chooser language "javascript,typescript,python,go,csharp" >}}

{{% choosable language javascript %}}

```javascript
const alicloud = require("@pulumi/alicloud")

const vpc = new alicloud.vpc.Network("my-vpc", {
    cidrBlock: "10.0.0.0/16",
});
```

{{% /choosable %}}
{{% choosable language typescript %}}

```typescript
import * as alicloud from "@pulumi/alicloud";

const vpc = new alicloud.vpc.Network("my-vpc", {
    cidrBlock: "10.0.0.0/16",
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_alicloud as alicloud

vpc = alicloud.vpc.Network("my-vpc",
    cidr_block="10.0.0.0/16"
)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	vpc "github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		vpc, err := vpc.NewVpc(ctx, "demo-instance", &vpc.VpcArgs{
			CidrBlock: pulumi.String("10.0.0.0/16"),
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
using Pulumi.Alicloud.Vpc;

class Program
{
    static Task Main() =>
        Deployment.Run(() => {
            var vpc = new Vpc("demo-instance", new VpcArgs
            {
                CidrBlock = "10.0.0.0/16"
            });
        });
}
```

{{% /choosable %}}

{{< /chooser >}}

## Libraries

The following packages are available in package managers:

* JavaScript/TypeScript: [`@pulumi/alicloud`](https://www.npmjs.com/package/@pulumi/alicloud)
* Python: [`pulumi-alicloud`](https://pypi.org/project/pulumi-alicloud/)
* Go: [`github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud`](https://github.com/pulumi/pulumi-alicloud)
* .NET: [`Pulumi.Alicloud`](https://www.nuget.org/packages/Pulumi.Alicloud)

The Alibaba Cloud provider is open source and available in the [pulumi/pulumi-alicloud](https://github.com/pulumi/pulumi-alicloud) repo.
