---
title: AWS S3 Replicated Bucket
meta_desc: Use Pulumi's Component for creating AWS S3 buckets that are replicated across regions using infrastructure as code.
layout: package
---

Easily create AWS S3 buckets that are replicated across AWS regions as a package available in all Pulumi languages.

## Example

{{< chooser language "typescript,python,go,csharp" >}}

{{% choosable language typescript %}}

```typescript
import * as s3 from "@pulumi/aws-s3-replicated-bucket";
const bucket = new s3.ReplicatedBucket("bucket", {
    destinationRegion: "us-east-1",
});
export const srcBucket = bucket.sourceBucket.arn;
export const dstBucket = bucket.destinationBucket.arn;
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_aws_s3_replicated_bucket as s3
bucket = s3.ReplicatedBucket("bucket", destination_region="us-east-1")
pulumi.export('srcBucket', bucket.source_bucket.arn)
pulumi.export('dstBucket', bucket.destination_bucket.arn)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
package main
import (
	"github.com/pulumi/pulumi-aws-s3-replicated-bucket/sdk/go/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
        bucket, err := s3.NewReplicatedBucket(ctx, "bucket", &s3.ReplicatedBucketArgs{
    	    DestinationRegion: pulumi.String("us-east-1"),
        })
        if err != nil {
			return err
		}
		ctx.Export("srcBucket", bucket.SourceBucket.Arn)
        ctx.Export("dstBucket", bucket.DestinationBucket.Arn)
		return nil
    }
}
```

{{% /choosable %}}
{{% choosable language csharp %}}

```csharp
using System.Collections.Generic;
using Pulumi;
using S3 = Pulumi.AwsS3ReplicatedBucket;

await Deployment.RunAsync(() =>
{
    var bucket = new S3.ReplicatedBucket("bucket", new S3.ReplicatedBucketArgs
    {
        DestinationRegion = "us-east-1"
    });

    return new Dictionary<string, object?>
    {
        ["srcBucket"] = bucket.SourceBucket.Arn,
        ["dstBucket"] = bucket.DestinationBucket.Arn
    };
});
```

{{% /choosable %}}

{{< /chooser >}}
