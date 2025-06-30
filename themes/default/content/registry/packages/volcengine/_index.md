---
# WARNING: this file was fetched from https://raw.githubusercontent.com/volcengine/pulumi-volcengine/v0.0.33/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Volcengine
meta_desc: Provides an overview of the Volcengine Provider for Pulumi.
layout: package
---

The Volcengine provider for Pulumi can be used to provision any of the cloud resources available in [Volcengine](https://volcengine.com).
The Volcengine provider must be configured with credentials to deploy and update resources in Volcengine.

## Example

{{< chooser language "typescript,python,go,csharp" >}}
{{% choosable language typescript %}}

```typescript
import * as volcengine from "@volcengine/pulumi";

const myVpc = new volcengine.vpc.Vpc("myVpc", {
    cidrBlock: "172.16.0.0/16",
    dnsServers: [
        "8.8.8.8",
        "114.114.114.114",
    ],
    vpcName: "pulumi-vpc-demo"
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_volcengine as volcengine

my_vpc = volcengine.vpc.Vpc("myVpc",
    cidr_block="172.16.0.0/16",
    dns_servers=[
        "8.8.8.8",
        "114.114.114.114",
    ],
    vpc_name="pulumi-vpc-demo"
)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
package main

import (
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
    "github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/vpc"
)

func main() {
    pulumi.Run(func(ctx *pulumi.Context) error {
        _, err := vpc.NewVpc(ctx, "myVpc", &vpc.VpcArgs{
            CidrBlock: pulumi.String("172.16.0.0/16"),
            DnsServers: pulumi.StringArray{
                pulumi.String("8.8.8.8"),
                pulumi.String("114.114.114.114"),
            },
            VpcName: pulumi.String("pulumi-vpc-demo"),
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
using Pulumi;
using Pulumi.Volcengine.Vpc;

return await Deployment.RunAsync(() =>
{
   var myVpc = new Vpc("myVpc", new VpcArgs
   {
      CidrBlock = "172.16.0.0/16",
      DnsServers = new InputList<string>
      {
         "8.8.8.8",
         "114.114.114.114"
      },
      VpcName = "pulumi-vpc-demo"
   });

   return new Dictionary<string, object?>
   {
   };
});
```

{{% /choosable %}}

{{< /chooser >}}