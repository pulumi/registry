---
title: AWS Native
meta_desc: Learn how you can use Pulumi's AWS Native Provider to reduce the complexity of provisioning and managing resources on AWS.
layout: overview
aliases:
    - "/docs/reference/clouds/aws-native/"
    - "/docs/intro/cloud-providers/aws-native/"
---

{{% notes type="info" %}}
AWS Native is in public preview. AWS Native provides coverage of all resources in the [AWS Cloud Control API](https://aws.amazon.com/blogs/aws/announcing-aws-cloud-control-api/), including same-day access to all new AWS resources. However, some AWS resources are not yet available in AWS Native.

For new projects, we recommend using AWS Native and [AWS Classic]({{<relref "/registry/packages/aws">}}) side-by-side so you can get the speed and correctness benefits of AWS Native where possible. For existing projects, [AWS Classic]({{<relref "/registry/packages/aws">}}) remains fully supported; at this time, we recommend waiting to migrate existing projects to AWS Native.
{{% /notes %}}

The AWS Native provider for Pulumi can provision many of the cloud resources available in [AWS](https://aws.amazon.com/). It manages and provisions resources using the [AWS Cloud Control API](https://aws.amazon.com/blogs/aws/announcing-aws-cloud-control-api/), which typically supports new AWS features on the day of launch. Resources available in the Pulumi AWS Native provider are based on the resources defined in the [AWS CloudFormation Registry](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/registry.html).

[Hundreds of AWS resources]({{<relref "/registry/packages/aws-native/api-docs">}}) are available in AWS Native. As AWS Cloud Control API adds resources, we will update AWS Native to include them.

AWS Native must be configured with credentials to deploy and update resources in AWS; see [Installation & Configuration]({{<relref "./installation-configuration">}}) for instructions.

{{< youtube oKxaZCyu2OQ >}}

## Examples

Check out the examples below to try out AWS Native:

* [Create an ECS Cluster on AWS]({{<relref "/registry/packages/aws-native/how-to-guides/aws-native-ts-ecs">}})
* [Host a Static Website on Amazon S3]({{<relref "/registry/packages/aws-native/how-to-guides/aws-native-ts-s3-folder">}})
* [Launch a Simple AWS Step Function State Machine With Lambda Functions]({{<relref "/registry/packages/aws-native/how-to-guides/aws-native-ts-stepfunctions">}})
* Create an Object Lambda access point that transforms object requests to a bucket:

{{< chooser language "typescript,python,csharp,go,java,yaml" >}}

{{% choosable language typescript %}}

```typescript
import * as awsnative from "@pulumi/aws-native";

const bucket = new awsnative.s3.Bucket("source");

const accessPoint = new awsnative.s3.AccessPoint("ap", {
   bucket: bucket.id,
});

const objectlambda = new awsnative.s3objectlambda.AccessPoint("objectlambda-ap", {
   objectLambdaConfiguration: {
       supportingAccessPoint: accessPoint.arn,
       transformationConfigurations: [{
           actions: ["GetObject"],
           contentTransformation: {
               AwsLambda: {
                   FunctionArn: fn.arn,
               },
           },
       }]
   }
});
```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi
import pulumi_aws_native as aws_native

my_bucket = aws_native.s3.Bucket("myBucket")

ap = aws_native.s3.AccessPoint("ap", bucket=my_bucket.id)

objectlambdaap = aws_native.s3objectlambda.AccessPoint("objectlambdaap", object_lambda_configuration=aws_native.s3objectlambda.AccessPointObjectLambdaConfigurationArgs(
    supporting_access_point=ap.arn,
    transformation_configurations=[aws_native.s3objectlambda.AccessPointTransformationConfigurationArgs(
        actions=["GetObject"],
        content_transformation={
            "AwsLambda": {
                "FunctionArn": fn.arn,
            },
        },
    )],
))
```

{{% /choosable %}}

{{% choosable language csharp %}}

```csharp
var bucket = new AwsNative.S3.Bucket("my-bucket");

var accessPoint = new AwsNative.S3.AccessPoint("ap", new AwsNative.S3.AccessPointArgs
{
    Bucket = bucket.Id
});

var objectLambda = new AwsNative.S3ObjectLambda.AccessPoint("objectlambda-ap", new AwsNative.S3ObjectLambda.AccessPointArgs
{
    ObjectLambdaConfiguration = new AwsNative.S3ObjectLambda.Inputs.AccessPointObjectLambdaConfigurationArgs
    {
        SupportingAccessPoint = accessPoint.Arn,
        TransformationConfigurations =
        {
            new AwsNative.S3ObjectLambda.Inputs.AccessPointTransformationConfigurationArgs
            {
                Actions = { "GetObject" },
                ContentTransformation = fn.Arn.Apply(arn => new Dictionary<string, object>
                {
                    ["AwsLambda"] = new Dictionary<string, object>
                    {
                        ["FunctionArn"] = arn
                    }
                }
            }
        }
    }
});
```

{{% /choosable %}}

{{% choosable language go %}}

```go
func main() {
    pulumi.Run(func(ctx *pulumi.Context) error {
        myBucket, err := s3.NewBucket(ctx, "myBucket", nil)
        if err != nil {
            return err
        }
        ap, err := s3.NewAccessPoint(ctx, "ap", &s3.AccessPointArgs{
            Bucket: myBucket.ID(),
        })
        if err != nil {
            return err
        }

        _, err = s3objectlambda.NewAccessPoint(ctx, "objectlambdaap", &s3objectlambda.AccessPointArgs{
            ObjectLambdaConfiguration: &s3objectlambda.AccessPointObjectLambdaConfigurationArgs{
                SupportingAccessPoint: ap.Arn,
                TransformationConfigurations: s3objectlambda.AccessPointTransformationConfigurationArray{
                    &s3objectlambda.AccessPointTransformationConfigurationArgs{
                        Actions: pulumi.StringArray{
                            pulumi.String("GetObject"),
                        },
                        ContentTransformation: pulumi.Map{
                            "AwsLambda": pulumi.Map{
                                "FunctionArn": fn.Arn,
                            },
                        },
                    },
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

{{% choosable language java %}}

```java
import java.util.Map;

import com.pulumi.Context;
import com.pulumi.Exports;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;

import com.pulumi.awsnative.s3.Bucket;
import com.pulumi.awsnative.s3.AccessPoint;
import com.pulumi.awsnative.s3.AccessPointArgs;
import com.pulumi.awsnative.s3objectlambda.inputs.AccessPointObjectLambdaConfigurationArgs;
import com.pulumi.awsnative.s3objectlambda.inputs.AccessPointTransformationConfigurationArgs;


public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    private static void stack(Context ctx) {

        final var bucket = new Bucket("source");

        final var accessPoint = new AccessPoint("ap", AccessPointArgs.builder()
            .bucket(bucket.getId())
            .build());

        final Output<String> functionArn = /* your function arn */;

        final var contentTransform = functionArn.applyValue(arn ->
            Map.of("AwsLambda", Map.of("FunctionArn", arn))
        );

        final var objectLambdaConfig = AccessPointObjectLambdaConfigurationArgs
            .builder()
            .supportingAccessPoint(accessPoint.arn())
            .transformationConfigurations(
                AccessPointTransformationConfigurationArgs
                    .builder()
                    .actions("GetObject")
                    .contentTransformation(contentTransform)
                    .build()
            )
            .build();

        final var objectLambdaArgs = com.pulumi.awsnative.s3objectlambda.AccessPointArgs
            .builder()
            .objectLambdaConfiguration(objectLambdaConfig)
            .build();

        final var objectLambda = new com.pulumi.awsnative.s3objectlambda.AccessPoint("objectlambda-ap", objectLambdaArgs);
        ctx.export("objectLambdaArn", objectLambda.arn());
    }
}
```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
resources:
  myBucket:
    type: 'aws-native:s3:Bucket'
  ap:
    type: 'aws-native:s3:AccessPoint'
    properties:
      bucket: '${myBucket}'
  action:
    type: 'aws-native:s3objectlambda:AccessPoint'
    properties:
      objectLambdaConfiguration:
        supportingAccessPoint: '${ap.Arn}'
        transformationConfigurations:
          - actions:
              - GetObject
            contentTransformation:
              AwsLambda:
                FunctionArn: '${fn.Arn}'
```

{{% /choosable %}}

{{% /chooser %}}

Visit the [How-to Guides]({{<relref "./how-to-guides">}}) to find more step-by-step guides for specific scenarios like setting up an EKS cluster.
