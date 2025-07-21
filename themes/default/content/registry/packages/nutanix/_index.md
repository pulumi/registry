---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pierskarsenbarg/pulumi-nutanix/v0.9.0/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Nutanix
meta_desc: Provides an overview of the Nutanix Provider for Pulumi.
layout: package
---

The Nutanix provider for Pulumi can be used to provision any of the cloud resources available in [Nutanix](https://www.nutanix.com/).
The Nutanix provider must be configured with credentials to deploy and update resources in Nutanix.

## Example

{{< chooser language "typescript,python,go,csharp" >}}

{{% choosable language typescript %}}

```typescript
import * as nutanix from "@pierskarsenbarg/nutanix";

const vpc = new nutanix.Vpc("vpc", {
  commonDomainNameServerIpLists: [
    {
      ip: "8.8.8.8",
    },
    {
      ip: "8.8.8.9",
    },
  ],
  externalSubnetReferenceNames: ["test-Ext1", "test-ext2"],
  externallyRoutablePrefixLists: [
    {
      ip: "192.43.0.0",
      prefixLength: 16,
    },
  ],
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_nutanix as nutanix

vpc = nutanix.Vpc("vpc",
    common_domain_name_server_ip_lists=[
        nutanix.VpcCommonDomainNameServerIpListArgs(
            ip="8.8.8.8",
        ),
        nutanix.VpcCommonDomainNameServerIpListArgs(
            ip="8.8.8.9",
        ),
    ],
    external_subnet_reference_names=[
        "test-Ext1",
        "test-ext2",
    ],
    externally_routable_prefix_lists=[nutanix.VpcExternallyRoutablePrefixListArgs(
        ip="192.43.0.0",
        prefix_length=16,
    )])
```

{{% /choosable %}}
{{% choosable language go %}}

```go
package main

import (
	"github.com/pierskarsenbarg/pulumi-nutanix/sdk/go/nutanix"
)

func main() {
    pulumi.Run(func(ctx *pulumi.Context) error {
        _, err := nutanix.NewVpc(ctx, "vpc", &nutanix.VpcArgs{
            CommonDomainNameServerIpLists: nutanix.VpcCommonDomainNameServerIpListArray{
                &nutanix.VpcCommonDomainNameServerIpListArgs{
                    Ip: pulumi.String("8.8.8.8"),
                },
                &nutanix.VpcCommonDomainNameServerIpListArgs{
                    Ip: pulumi.String("8.8.8.9"),
                },
            },
            ExternalSubnetReferenceNames: pulumi.StringArray{
                pulumi.String("test-Ext1"),
                pulumi.String("test-ext2"),
            },
            ExternallyRoutablePrefixLists: nutanix.VpcExternallyRoutablePrefixListArray{
                &nutanix.VpcExternallyRoutablePrefixListArgs{
                    Ip:           pulumi.String("192.43.0.0"),
                    PrefixLength: pulumi.Int(16),
                },
            },
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
using System.Linq;
using Pulumi;
using Nutanix = PiersKarsenbarg.Nutanix;

return await Deployment.RunAsync(() =>
{
    var vpc = new Nutanix.Vpc("vpc", new()
    {
        CommonDomainNameServerIpLists = new[]
        {
            new Nutanix.Inputs.VpcCommonDomainNameServerIpListArgs
            {
                Ip = "8.8.8.8",
            },
            new Nutanix.Inputs.VpcCommonDomainNameServerIpListArgs
            {
                Ip = "8.8.8.9",
            },
        },
        ExternalSubnetReferenceNames = new[]
        {
            "test-Ext1",
            "test-ext2",
        },
        ExternallyRoutablePrefixLists = new[]
        {
            new Nutanix.Inputs.VpcExternallyRoutablePrefixListArgs
            {
                Ip = "192.43.0.0",
                PrefixLength = 16,
            },
        },
    });

});
```

{{% /choosable %}}

{{< /chooser >}}
