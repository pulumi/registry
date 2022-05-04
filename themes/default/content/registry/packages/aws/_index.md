---
title: AWS Classic
meta_desc: Learn how you can use Pulumi's AWS Classic Provider to reduce the complexity of provisioning and managing resources on AWS.
layout: overview
---

The Amazon Web Services (AWS) provider for Pulumi can provision many of the cloud resources available in [AWS](https://aws.amazon.com/). It uses the AWS SDK to manage and provision resources.

The AWS provider must be configured with credentials to deploy and update resources in AWS; see [Installation & Configuration]({{<relref "./installation-configuration">}}) for instructions.

**New to Pulumi and AWS?** [Get started with AWS using our tutorial]({{<relref "/docs/get-started/aws">}}).

{{% notes %}}
Pulumi has a new AWS provider: the [Pulumi AWS Native Provider]({{<relref "/registry/packages/aws-native">}}). AWS Native gives you same-day access to all new AWS resources and includes coverage of all resources in the [AWS Cloud Control API](https://aws.amazon.com/blogs/aws/announcing-aws-cloud-control-api/).

Consider trying [AWS Native]({{<relref "/registry/packages/aws-native">}}) if you need AWS resources that aren't available in this provider.
{{% /notes %}}

## Example

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}

{{% choosable language typescript %}}

```typescript
const aws = require("@pulumi/aws");

const bucket = new aws.s3.Bucket("mybucket");
```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi
import pulumi_aws as aws

bucket = aws.s3.Bucket("bucket")
```

{{% /choosable %}}

{{% choosable language go %}}

```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := s3.NewBucket(ctx, "bucket", &s3.BucketArgs{})
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
using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var bucket = new Aws.S3.Bucket("bucket");
    }
}

```

{{% /choosable %}}

{{% choosable language java %}}

```java
import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.aws.s3.Bucket;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    private static void stack(Context ctx) {
        final var bucket = new Bucket("my-bucket");
        ctx.export("bucketName", bucket.name());
  }
}
```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
resources:
  mybucket:
    type: aws:s3:Bucket
outputs:
  bucketName: ${mybucket.name}
```

{{% /choosable %}}

{{< /chooser >}}

Visit the [How-to Guides]({{<relref "./how-to-guides">}}) to find step-by-step guides for specific scenarios like creating a serverless application or setting up Athena search.

## Components

Pulumi offers Components that provide simpler interfaces and higher-productivity APIs for many areas of AWS:

* [Amazon EKS]({{<relref "/registry/packages/eks">}})
* [Crosswalk for AWS]({{<relref "/docs/guides/crosswalk/aws">}}), which includes API Gateway, CloudWatch, Elastic Container Registry, Elastic Container Service, Elastic Kubernetes Service, Elastic Load Balancing, Identity & Access Management, Lambda, Virtual Private Cloud, and more
