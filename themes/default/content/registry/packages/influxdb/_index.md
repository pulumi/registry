---
# WARNING: this file was fetched from https://raw.githubusercontent.com/komminarlabs/pulumi-influxdb/v1.6.0/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/komminarlabs/pulumi-influxdb/blob/v1.6.0/docs/_index.md
title: InfluxDB
meta_desc: Provides an overview of the InfluxDB Provider for Pulumi.
layout: package
---

The InfluxDB provider for Pulumi can be used to provision the resources available in [InfluxDB](https://www.influxdata.com/).

The InfluxDB provider must be configured with credentials to deploy and update resources in InfluxDB; see [Installation & Configuration](./installation-configuration) for instructions.

## Supported InfluxDB flavours

### v3

* [InfluxDB Cloud Serverless](https://www.influxdata.com/products/influxdb-cloud/serverless/)

### v2

* [InfluxDB Cloud TSM](https://docs.influxdata.com/influxdb/cloud/)
* [InfluxDB OSS](https://docs.influxdata.com/influxdb/v2/)

## Example

{{< chooser language "typescript,python,go,csharp" >}}
{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as influxdb from "@thulasirajkomminar/influxdb";

// Create a new Bucket
export const orgId = influxdb.getOrganizationOutput({ name: "IoT" }).id;

export const bucket = new influxdb.Bucket("signals", {
    orgId: orgId,
    name: "signals",
    description: "This is a bucket to store signals",
    retentionPeriod: 604800,
});

// Get the id of the new bucket as an output
export const bucketId = bucket.id;
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import thulasirajkomminar_influxdb as influxdb

org_id = influxdb.get_organization(name="IoT").id

bucket = influxdb.Bucket(
    "signals",
    org_id=org_id,
    name="signals",
    description="This is a bucket to store signals",
    retention_period=604800,
)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	influxdb "github.com/thulasirajkomminar/pulumi-influxdb/sdk/go/influxdb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		org, err := influxdb.LookupOrganization(ctx, &influxdb.LookupOrganizationArgs{Name: "IoT"})
		if err != nil {
			return err
		}

		signals, err := influxdb.NewBucket(ctx, "signals", &influxdb.BucketArgs{
			OrgId:           pulumi.String(org.Id),
			Name:            pulumi.String("signals"),
			Description:     pulumi.String("Bucket for storing signal data"),
			RetentionPeriod: pulumi.Int(604800),
		})
		if err != nil {
			return err
		}

		ctx.Export("bucketId", signals.ID())
		return nil
	})
}
```

{{% /choosable %}}

{{< /chooser >}}
