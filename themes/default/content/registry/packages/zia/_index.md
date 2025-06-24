---
# WARNING: this file was fetched from https://raw.githubusercontent.com/zscaler/pulumi-zia/v1.1.1/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Zscaler Internet Access (ZIA)
meta_desc: Provides an overview of the ZIA Provider for Pulumi.
layout: overview
---

The Zscaler Internet Access (ZIA) provider for Pulumi can be used to provision any of the cloud resources available in [Zscaler Internet Access](https://help.zscaler.com/zia).
The ZIA provider must be configured with credentials to deploy and update resources in the ZIA Cloud.

## Example

{{< chooser language "typescript,python,go,csharp" >}}
{{% choosable language typescript %}}

```typescript
import * as zia from "@bdzscaler/pulumi-zia";
const staticIP = new zia.ZIATrafficForwardingStaticIP("static_ip_example", {
    comment: "Pulumi Traffic Forwarding Static IP",
    geoOverride: true,
    ipAddress: "123.234.244.245",
    latitude: -36.848461,
    longitude: 174.763336,
    routableIp: true,
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import zscaler_pulumi_zia as zia

staticIP = zia.ZIATrafficForwardingStaticIPArgs("static_ip_example",
  comment="Pulumi Traffic Forwarding Static IP",
  geoOverride=True,
  routableIp=True,
  ipAddress="123.234.244.245",
  latitude=-36.848461,
  longitude=174.763336,
)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	zia "github.com/zscaler/pulumi-zia/sdk/go/zia"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		staticIP, err := zia.NewZIATrafficForwardingStaticIP(ctx, "static_ip_example", &zia.ZIATrafficForwardingStaticIPArgs{
			Comment:     pulumi.String("Pulumi Traffic Forwarding Static IP"),
			RoutableIp:  pulumi.Bool(true),
			GeoOverride: pulumi.Bool(true),
			IpAddress:   pulumi.String("123.234.244.245"),
			Latitude:    pulumi.Float64Ptr(37.3382082),
			Longitude:   pulumi.Float64Ptr(-121.8863286),
		})
		if err != nil {
			return fmt.Errorf("error creating zia static ip: %v", err)
		}

		ctx.Export("staticIP", staticIP.IpAddress)

		return nil
	})
}
```

{{% /choosable %}}
{{% choosable language csharp %}}

```csharp
    using System.Collections.Generic;
    using Pulumi;
    using zia = zscaler.PulumiPackage.Zia;

    return await Deployment.RunAsync(() =&gt;
    {
        // ZIA Traffic Forwarding - Static IP
     var example = new zia.ZIATrafficForwardingStaticIP("static_ip_example", new()
            {
             Comment = "Pulumi Traffic Forwarding Static IP",
             GeoOverride = true,
             RoutableIp = true,
             IpAddress = "123.234.244.245",
             Latitude = -36.848461,
             Longitude = 174.763336,
         });

     });
```

{{% /choosable %}}

{{< /chooser >}}
