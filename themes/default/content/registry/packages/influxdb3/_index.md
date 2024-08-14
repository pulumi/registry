---
title: InfluxDB V3
meta_desc: Provides an overview of the InfluxDB V3 Provider for Pulumi.
layout: package
---

The InfluxDB V3 provider for Pulumi can be used to provision the resources available in [InfluxDB V3](https://www.influxdata.com/).

The InfluxDB V3 provider must be configured with credentials to deploy and update resources in InfluxDB; see [Installation & Configuration](./installation-configuration) for instructions.

## Example

{{< chooser language "typescript,python,go" >}}
{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as influxdb from "@komminarlabs/influxdb3";

// Create a new Bucket
export const database = new influxdb.Database("signals", {
    name: "signals",
    retentionPeriod: 604800,
});

// Get the id of the new bucket as an output
export const databaseId = database.id;
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import komminarlabs_influxdb3 as influxdb3

database = influxdb3.Database(
    "signals",
    name="signals",
    retention_period=604800,
)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
package main

import (
	"github.com/komminarlabs/pulumi-influxdb3/sdk/go/influxdb3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		signals, err := influxdb3.NewDatabase(ctx, "signals", &influxdb3.DatabaseArgs{
			Name:            pulumi.String("signals"),
			RetentionPeriod: pulumi.Int(604800),
		})
		if err != nil {
			return err
		}

		ctx.Export("databaseId", signals.ID())
		return nil
	})
}
```

{{% /choosable %}}

{{< /chooser >}}
