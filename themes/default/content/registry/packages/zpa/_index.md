---
# WARNING: this file was fetched from https://raw.githubusercontent.com/zscaler/pulumi-zpa/v1.0.3/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/zscaler/pulumi-zpa/blob/v1.0.3/docs/_index.md
title: Zscaler Private Access (ZPA)
meta_desc: Provides an overview of the ZPA Provider for Pulumi.
layout: overview
---

The Zscaler Private Access (ZPA) provider for Pulumi can be used to provision any of the cloud resources available in [Zscaler Private Access](https://help.zscaler.com/zpa).
The Zscaler Private Access (ZPA) provider must be configured with credentials to deploy and update resources in the ZPA Cloud.

## Example

{{< chooser language "typescript,python,go,csharp" >}}
{{% choosable language typescript %}}

```typescript
import * as zpa from "@bdzscaler/pulumi-zpa";
const segmentGroup = new zpa.ZPASegmentGroup("segment-group", {
    name:          "Pulumi Segment Group",
    description:   "Pulumi Segment Group",
    enabled: true,
    policyMigrated: true,
    tcpKeepAliveEnabled: "1",
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import zscaler_pulumi_zpa as zpa

segmentGroup = zpa.ZPASegmentGroup("segment-group",
  name="Pulumi-Segment-Group",
  description="Pulumi-Segment-Group",
  enabled=True,
  policy_migrated=True,
  tcp_keep_alive_enabled="1"
)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	zpa "github.com/zscaler/pulumi-zpa/sdk/go/zpa"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		segmentGroup, err := zpa.NewZPASegmentGroup(ctx, "segment-group-example", &zpa.ZPASegmentGroupArgs{
			Name:                   pulumi.String("Pulumi Segment Group"),
			Description:            pulumi.String("Pulumi Segment Group"),
			Enabled:                pulumi.Bool(true),
			PolicyMigrated:         pulumi.Bool(true),
			TcpKeepAliveEnabled: pulumi.String("1"),
		})
		if err != nil {
			return fmt.Errorf("error creating segment group: %v", err)
		}

		ctx.Export("segment_group", segmentGroup.Name)

		return nil
	})
}
```

{{% /choosable %}}
{{% choosable language csharp %}}

```csharp
using Pulumi;
using zscaler.PulumiPackage.Zpa;

await Deployment.RunAsync(() =>
{
	var group = new Zpa.ZPASegmentGroup("segment-group-example", new()
	{
		  Name = "Pulumi Segment Group",
        Description = "Pulumi Segment Group",
        Enabled = true,
        TcpKeepAliveEnabled = "1",
	});
});
```

{{% /choosable %}}

{{< /chooser >}}
