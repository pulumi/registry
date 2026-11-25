---
title: AWS Cloud Control
meta_desc: Learn how you can use Pulumi's AWS Cloud Control Provider to reduce the complexity of provisioning and managing resources on AWS.
layout: package
aliases:
    - "/docs/reference/clouds/aws-native/"
    - "/docs/intro/cloud-providers/aws-native/"
---

{{% notes type="info" %}}
AWS Cloud Control provides coverage of all resources in the [AWS Cloud Control API](https://aws.amazon.com/blogs/aws/announcing-aws-cloud-control-api/), including same-day access to all new AWS resources. However, some AWS resources are not yet available in AWS Cloud Control.

For new projects, we recommend starting with our primary [AWS Provider](/registry/packages/aws) and adding AWS Cloud Control resources on an as needed basis.
{{% /notes %}}

The AWS Cloud Control provider for Pulumi can provision many of the cloud resources available in [AWS](https://aws.amazon.com/). It manages and provisions resources using the [AWS Cloud Control API](https://aws.amazon.com/blogs/aws/announcing-aws-cloud-control-api/), which typically supports new AWS features on the day of launch. Resources available in the Pulumi AWS Cloud Control provider are based on the resources defined in the [AWS CloudFormation Registry](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/registry.html).

[Hundreds of AWS resources](/registry/packages/aws-native/api-docs) are available in AWS Cloud Control. As AWS Cloud Control API adds resources, we will update AWS Cloud Control to include them.

AWS Cloud Control must be configured with credentials to deploy and update resources in AWS; see [Installation & Configuration](./installation-configuration) for instructions.

## Example

Create an Object Lambda access point that transforms object requests to a bucket:

{{< chooser language "typescript,python,csharp,go,java,yaml,hcl" >}}

{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as awsnative from "@pulumi/aws-native";

// The ARN of the Lambda function that transforms objects.
const config = new pulumi.Config();
const functionArn = config.require("functionArn");

const myBucket = new awsnative.s3.Bucket("myBucket");

const ap = new awsnative.s3.AccessPoint("ap", {
    bucket: myBucket.id,
});

const objectLambdaAp = new awsnative.s3objectlambda.AccessPoint("objectLambdaAp", {
    objectLambdaConfiguration: {
        supportingAccessPoint: ap.arn,
        transformationConfigurations: [{
            actions: ["GetObject"],
            contentTransformation: {
                awsLambda: {
                    functionArn: functionArn,
                },
            },
        }],
    },
});
```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi
import pulumi_aws_native as aws_native

# The ARN of the Lambda function that transforms objects.
config = pulumi.Config()
function_arn = config.require("functionArn")

my_bucket = aws_native.s3.Bucket("myBucket")

ap = aws_native.s3.AccessPoint("ap", bucket=my_bucket.id)

object_lambda_ap = aws_native.s3objectlambda.AccessPoint("objectLambdaAp", object_lambda_configuration={
    "supporting_access_point": ap.arn,
    "transformation_configurations": [{
        "actions": ["GetObject"],
        "content_transformation": {
            "aws_lambda": {
                "function_arn": function_arn,
            },
        },
    }],
})
```

{{% /choosable %}}

{{% choosable language csharp %}}

```csharp
using Pulumi;
using AwsNative = Pulumi.AwsNative;

return await Deployment.RunAsync(() =>
{
    // The ARN of the Lambda function that transforms objects.
    var config = new Config();
    var functionArn = config.Require("functionArn");

    var myBucket = new AwsNative.S3.Bucket("myBucket");

    var ap = new AwsNative.S3.AccessPoint("ap", new()
    {
        Bucket = myBucket.Id,
    });

    var objectLambdaAp = new AwsNative.S3ObjectLambda.AccessPoint("objectLambdaAp", new()
    {
        ObjectLambdaConfiguration = new AwsNative.S3ObjectLambda.Inputs.AccessPointObjectLambdaConfigurationArgs
        {
            SupportingAccessPoint = ap.Arn,
            TransformationConfigurations =
            {
                new AwsNative.S3ObjectLambda.Inputs.AccessPointTransformationConfigurationArgs
                {
                    Actions = { "GetObject" },
                    ContentTransformation = new AwsNative.S3ObjectLambda.Inputs.AccessPointTransformationConfigurationContentTransformationPropertiesArgs
                    {
                        AwsLambda = new AwsNative.S3ObjectLambda.Inputs.AccessPointAwsLambdaArgs
                        {
                            FunctionArn = functionArn,
                        },
                    },
                },
            },
        },
    });
});
```

{{% /choosable %}}

{{% choosable language go %}}

```go
package main

import (
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/s3"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/s3objectlambda"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// The ARN of the Lambda function that transforms objects.
		cfg := config.New(ctx, "")
		functionArn := cfg.Require("functionArn")

		myBucket, err := s3.NewBucket(ctx, "myBucket", nil)
		if err != nil {
			return err
		}

		ap, err := s3.NewAccessPoint(ctx, "ap", &s3.AccessPointArgs{
			Bucket: myBucket.ID().ToStringOutput(),
		})
		if err != nil {
			return err
		}

		_, err = s3objectlambda.NewAccessPoint(ctx, "objectLambdaAp", &s3objectlambda.AccessPointArgs{
			ObjectLambdaConfiguration: &s3objectlambda.AccessPointObjectLambdaConfigurationArgs{
				SupportingAccessPoint: ap.Arn,
				TransformationConfigurations: s3objectlambda.AccessPointTransformationConfigurationArray{
					&s3objectlambda.AccessPointTransformationConfigurationArgs{
						Actions: pulumi.StringArray{
							pulumi.String("GetObject"),
						},
						ContentTransformation: &s3objectlambda.AccessPointTransformationConfigurationContentTransformationPropertiesArgs{
							AwsLambda: &s3objectlambda.AccessPointAwsLambdaArgs{
								FunctionArn: pulumi.String(functionArn),
							},
						},
					},
				},
			},
		})
		return err
	})
}
```

{{% /choosable %}}

{{% choosable language java %}}

```java
import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.awsnative.s3.AccessPoint;
import com.pulumi.awsnative.s3.AccessPointArgs;
import com.pulumi.awsnative.s3.Bucket;
import com.pulumi.awsnative.s3objectlambda.inputs.AccessPointAwsLambdaArgs;
import com.pulumi.awsnative.s3objectlambda.inputs.AccessPointObjectLambdaConfigurationArgs;
import com.pulumi.awsnative.s3objectlambda.inputs.AccessPointTransformationConfigurationArgs;
import com.pulumi.awsnative.s3objectlambda.inputs.AccessPointTransformationConfigurationContentTransformationPropertiesArgs;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        // The ARN of the Lambda function that transforms objects.
        final var functionArn = ctx.config().require("functionArn");

        var myBucket = new Bucket("myBucket");

        var ap = new AccessPoint("ap", AccessPointArgs.builder()
            .bucket(myBucket.id())
            .build());

        var objectLambdaAp = new com.pulumi.awsnative.s3objectlambda.AccessPoint("objectLambdaAp",
            com.pulumi.awsnative.s3objectlambda.AccessPointArgs.builder()
                .objectLambdaConfiguration(AccessPointObjectLambdaConfigurationArgs.builder()
                    .supportingAccessPoint(ap.arn())
                    .transformationConfigurations(AccessPointTransformationConfigurationArgs.builder()
                        .actions("GetObject")
                        .contentTransformation(AccessPointTransformationConfigurationContentTransformationPropertiesArgs.builder()
                            .awsLambda(AccessPointAwsLambdaArgs.builder()
                                .functionArn(functionArn)
                                .build())
                            .build())
                        .build())
                    .build())
                .build());
    }
}
```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
config:
  # The ARN of the Lambda function that transforms objects.
  functionArn:
    type: string
resources:
  myBucket:
    type: aws-native:s3:Bucket
  ap:
    type: aws-native:s3:AccessPoint
    properties:
      bucket: ${myBucket}
  objectLambdaAp:
    type: aws-native:s3objectlambda:AccessPoint
    properties:
      objectLambdaConfiguration:
        supportingAccessPoint: ${ap.arn}
        transformationConfigurations:
          - actions:
              - GetObject
            contentTransformation:
              awsLambda:
                functionArn: ${functionArn}
```

{{% /choosable %}}

{{% choosable language hcl %}}

```hcl
terraform {
  required_providers {
    aws-native = {
      source = "pulumi/aws-native"
    }
  }
}

# The ARN of the Lambda function that transforms objects.
variable "function_arn" {
  type = string
}

resource "aws-native_s3_bucket" "myBucket" {
}

resource "aws-native_s3_access_point" "ap" {
  bucket = aws-native_s3_bucket.myBucket.id
}

resource "aws-native_s3objectlambda_access_point" "objectLambdaAp" {
  object_lambda_configuration = {
    supporting_access_point = aws-native_s3_access_point.ap.arn
    transformation_configurations = [{
      actions = ["GetObject"]
      content_transformation = {
        aws_lambda = {
          function_arn = var.function_arn
        }
      }
    }]
  }
}
```

{{% /choosable %}}

{{% /chooser %}}

## Third Party Resources

The SDK for the AWS Cloud Control provider only includes resources for which Amazon have published the corresponding CloudFormation specifications. AWS also supports [CloudFormation extensions via the CloudFormation registry](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/registry.html) which can also be accessed via Cloud Control.

If you want to manage resources using Pulumi's AWS Cloud Control Provider which aren't in the SDK, you can use the [ExtensionResource](https://www.pulumi.com/registry/packages/aws-native/api-docs/extensionresource/) resource within root of the SDK. The input properties and output properties are untyped allowing passing in any arbitrary values which will be passed to the Cloud Control API and tracked with Pulumi's managed state model.

Here's a very simple demonstration of using the ExtensionResource to create an S3 bucket:

{{< chooser language "typescript,python,go,csharp,java,yaml,hcl" / >}}

{{% choosable language "typescript" %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws_native from "@pulumi/aws-native";

const myBucket = new aws_native.ExtensionResource("myBucket", {
    type: "AWS::S3::Bucket",
    properties: {
        BucketName: "my-bucket",
    },
});
```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi
import pulumi_aws_native as aws_native

my_bucket = aws_native.ExtensionResource("myBucket",
    type="AWS::S3::Bucket",
    properties={
        "BucketName": "my-bucket",
    })
```

{{% /choosable %}}

{{% choosable language go %}}

```go
package main

import (
    awsnative "github.com/pulumi/pulumi-aws-native/sdk/go/aws"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
    pulumi.Run(func(ctx *pulumi.Context) error {
        _, err := awsnative.NewExtensionResource(ctx, "myBucket", &awsnative.ExtensionResourceArgs{
            Type: pulumi.String("AWS::S3::Bucket"),
            Properties: pulumi.Map{
                "BucketName": pulumi.Any("my-bucket"),
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
using AwsNative = Pulumi.AwsNative;

return await Deployment.RunAsync(() =>
{
    var myBucket = new AwsNative.ExtensionResource("myBucket", new()
    {
        Type = "AWS::S3::Bucket",
        Properties =
        {
            { "BucketName", "my-bucket" },
        },
    });

});
```

{{% /choosable %}}

{{% choosable language java %}}

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.awsnative.ExtensionResource;
import com.pulumi.awsnative.ExtensionResourceArgs;
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        var myBucket = new ExtensionResource("myBucket", ExtensionResourceArgs.builder()
            .type("AWS::S3::Bucket")
            .properties(Map.of("BucketName", "my-bucket"))
            .build());

    }
}
```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
resources:
  myBucket:
    type: 'aws-native:index:ExtensionResource'
    properties:
      type: 'AWS::S3::Bucket'
      properties:
        BucketName: my-bucket
```

{{% /choosable %}}

{{% choosable language hcl %}}

```hcl
terraform {
  required_providers {
    aws-native = {
      source = "pulumi/aws-native"
    }
  }
}

resource "aws-native_extension_resource" "myBucket" {
  type = "AWS::S3::Bucket"
  properties = {
    BucketName = "my-bucket"
  }
}
```

{{% /choosable %}}
