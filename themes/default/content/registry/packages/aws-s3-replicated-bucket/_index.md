---
title: AWS S3 Replicated Bucket
meta_desc: Use Pulumi's Component for creating AWS S3 buckets that are replicated across regions using infrastructure as code.
layout: overview
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
using Pulumi;
using S3 = Pulumi.AwsS3ReplicatedBucket;
class MyStack : Stack
{
    [Output("srcBucket")] Output<string> SrcBucket { get; set; }
    [Output("dstBucket")] Output<string> DstBucket { get; set; }
    public MyStack()
    {
        var bucket = new S3.ReplicatedBucket(ctx, "bucket", new S3.ReplicatedBucketArgs{
            DestinationRegion = "us-east-1"
        });
        this.SrcBucket = bucket.SourceBucket.Arn;
        this.DstBucket = bucket.DestinationBucket.Arn;
    }
}
```

{{% /choosable %}}

{{< /chooser >}}
