---
title: "BucketV2 Migration Guide"
h1: "Migrating from aws.s3.Bucket to aws.s3.BucketV2"
meta_desc: Practitioner level instructions for migrating from aws.s3.Bucket to aws.s3.BucketV2 resources
layout: package
---

## Migrating from aws.s3.Bucket to aws.s3.BucketV2

In the upcoming AWS Classic major release (v7), `aws.s3.Bucket` will be discontinued in favor of `BucketV2`. Users of
`aws.s3.Bucket` resource are encouraged to migrate early. The migration is straightforward for simple use cases but
requires some care for advanced configurations. Specifically `BucketV2` inputs such as `policy` and `accelerationStatus`
are deprecated and using the requires migrating to side-by-side resources (see Removed inputs). This guide aims to cover
all relevant scenarios with precise migration instructions.

To migrate existing `aws.s3.Bucket` resources to `aws.s3.BucketV2`:

1. Edit your Pulumi program sources to replace removed inputs (see below) with equivalent side-by-side resources.

2. Perform `pulumi up` to replace (delete and re-create) the buckets in AWS.

If replacement is not acceptable, it is possible to perform a manual migration with `pulumi import` (see Avoiding
replacement).

## Migrating deprecated inputs

Inputs such as `policy` and `accelerationStatus` are deprecated on the `aws.s3.BucketV2` resource. Instead of using
these inputs, it is recommended to use side-by-side resources to achieve the same effect.

### policy input

To specify a
[bucket policy](https://docs.aws.amazon.com/AmazonS3/latest/dev/example-bucket-policies.html),
instead of:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}

{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const myBucket = new aws.s3.Bucket("my-bucket", {
    bucket: "my-bucket-26224916",
    policy: JSON.stringify({
        Version: "2012-10-17",
        Id: "PutObjPolicy",
        Statement: {
            Sid: "DenyObjectsThatAreNotSSEKMS",
            Principal: "*",
            Effect: "Deny",
            Action: "s3:PutObject",
            Resource: "arn:aws:s3:::my-bucket-26224916/*",
            Condition: {
                Null: {
                    "s3:x-amz-server-side-encryption-aws-kms-key-id": "true",
                },
            },
        },
    }),
});

```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi
import json
import pulumi_aws as aws

my_bucket = aws.s3.Bucket("my-bucket",
    bucket="my-bucket-26224916",
    policy=json.dumps({
        "Version": "2012-10-17",
        "Id": "PutObjPolicy",
        "Statement": {
            "Sid": "DenyObjectsThatAreNotSSEKMS",
            "Principal": "*",
            "Effect": "Deny",
            "Action": "s3:PutObject",
            "Resource": "arn:aws:s3:::my-bucket-26224916/*",
            "Condition": {
                "Null": {
                    "s3:x-amz-server-side-encryption-aws-kms-key-id": "true",
                },
            },
        },
    }))

```

{{% /choosable %}}

{{% choosable language go %}}

```go
package main

import (
	"encoding/json"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"Version": "2012-10-17",
			"Id":      "PutObjPolicy",
			"Statement": map[string]interface{}{
				"Sid":       "DenyObjectsThatAreNotSSEKMS",
				"Principal": "*",
				"Effect":    "Deny",
				"Action":    "s3:PutObject",
				"Resource":  "arn:aws:s3:::my-bucket-26224916/*",
				"Condition": map[string]interface{}{
					"Null": map[string]interface{}{
						"s3:x-amz-server-side-encryption-aws-kms-key-id": "true",
					},
				},
			},
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		_, err = s3.NewBucket(ctx, "my-bucket", &s3.BucketArgs{
			Bucket: pulumi.String("my-bucket-26224916"),
			Policy: pulumi.String(json0),
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
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var myBucket = new Aws.S3.Bucket("my-bucket", new()
    {
        BucketName = "my-bucket-26224916",
        Policy = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["Version"] = "2012-10-17",
            ["Id"] = "PutObjPolicy",
            ["Statement"] = new Dictionary<string, object?>
            {
                ["Sid"] = "DenyObjectsThatAreNotSSEKMS",
                ["Principal"] = "*",
                ["Effect"] = "Deny",
                ["Action"] = "s3:PutObject",
                ["Resource"] = "arn:aws:s3:::my-bucket-26224916/*",
                ["Condition"] = new Dictionary<string, object?>
                {
                    ["Null"] = new Dictionary<string, object?>
                    {
                        ["s3:x-amz-server-side-encryption-aws-kms-key-id"] = "true",
                    },
                },
            },
        }),
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
import com.pulumi.aws.s3.Bucket;
import com.pulumi.aws.s3.BucketArgs;
import static com.pulumi.codegen.internal.Serialization.*;
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
        var myBucket = new Bucket("myBucket", BucketArgs.builder()
            .bucket("my-bucket-26224916")
            .policy(serializeJson(
                jsonObject(
                    jsonProperty("Version", "2012-10-17"),
                    jsonProperty("Id", "PutObjPolicy"),
                    jsonProperty("Statement", jsonObject(
                        jsonProperty("Sid", "DenyObjectsThatAreNotSSEKMS"),
                        jsonProperty("Principal", "*"),
                        jsonProperty("Effect", "Deny"),
                        jsonProperty("Action", "s3:PutObject"),
                        jsonProperty("Resource", "arn:aws:s3:::my-bucket-26224916/*"),
                        jsonProperty("Condition", jsonObject(
                            jsonProperty("Null", jsonObject(
                                jsonProperty("s3:x-amz-server-side-encryption-aws-kms-key-id", "true")
                            ))
                        ))
                    ))
                )))
            .build());

    }
}

```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
name: example
runtime: yaml
resources:
  my-bucket:
    type: aws:s3:Bucket
    properties:
      bucket: "my-bucket-26224916"
      policy:
        fn::toJSON:
          Version: 2012-10-17
          Id: PutObjPolicy
          Statement:
            Sid: DenyObjectsThatAreNotSSEKMS
            Principal: "*"
            Effect: Deny
            Action: s3:PutObject
            Resource: "arn:aws:s3:::my-bucket-26224916/*"
            Condition:
              "Null":
                s3:x-amz-server-side-encryption-aws-kms-key-id: "true"
```

{{% /choosable %}}

{{< /chooser >}}


Use the `aws.s3.BucketPolicy` resource:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}

{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const myBucket = new aws.s3.BucketV2("my-bucket", {});
const myBucketPolicy = new aws.s3.BucketPolicy("my-bucket-policy", {
    bucket: "my-bucket-26224916",
    policy: pulumi.jsonStringify({
        Version: "2012-10-17",
        Id: "PutObjPolicy",
        Statement: {
            Sid: "DenyObjectsThatAreNotSSEKMS",
            Principal: "*",
            Effect: "Deny",
            Action: "s3:PutObject",
            Resource: pulumi.interpolate`arn:aws:s3:::${myBucket.bucket}/*`,
            Condition: {
                Null: {
                    "s3:x-amz-server-side-encryption-aws-kms-key-id": "true",
                },
            },
        },
    }),
});

```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi
import json
import pulumi_aws as aws

my_bucket = aws.s3.BucketV2("my-bucket")
my_bucket_policy = aws.s3.BucketPolicy("my-bucket-policy",
    bucket="my-bucket-26224916",
    policy=pulumi.Output.json_dumps({
        "Version": "2012-10-17",
        "Id": "PutObjPolicy",
        "Statement": {
            "Sid": "DenyObjectsThatAreNotSSEKMS",
            "Principal": "*",
            "Effect": "Deny",
            "Action": "s3:PutObject",
            "Resource": my_bucket.bucket.apply(lambda bucket: f"arn:aws:s3:::{bucket}/*"),
            "Condition": {
                "Null": {
                    "s3:x-amz-server-side-encryption-aws-kms-key-id": "true",
                },
            },
        },
    }))

```

{{% /choosable %}}

{{% choosable language go %}}

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		myBucket, err := s3.NewBucketV2(ctx, "my-bucket", nil)
		if err != nil {
			return err
		}
		_, err = s3.NewBucketPolicy(ctx, "my-bucket-policy", &s3.BucketPolicyArgs{
			Bucket: pulumi.String("my-bucket-26224916"),
			Policy: myBucket.Bucket.ApplyT(func(bucket string) (pulumi.String, error) {
				var _zero pulumi.String
				tmpJSON0, err := json.Marshal(map[string]interface{}{
					"Version": "2012-10-17",
					"Id":      "PutObjPolicy",
					"Statement": map[string]interface{}{
						"Sid":       "DenyObjectsThatAreNotSSEKMS",
						"Principal": "*",
						"Effect":    "Deny",
						"Action":    "s3:PutObject",
						"Resource":  fmt.Sprintf("arn:aws:s3:::%v/*", bucket),
						"Condition": map[string]interface{}{
							"Null": map[string]interface{}{
								"s3:x-amz-server-side-encryption-aws-kms-key-id": "true",
							},
						},
					},
				})
				if err != nil {
					return _zero, err
				}
				json0 := string(tmpJSON0)
				return pulumi.String(json0), nil
			}).(pulumi.StringOutput),
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
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var myBucket = new Aws.S3.BucketV2("my-bucket");

    var myBucketPolicy = new Aws.S3.BucketPolicy("my-bucket-policy", new()
    {
        Bucket = "my-bucket-26224916",
        Policy = Output.JsonSerialize(Output.Create(new Dictionary<string, object?>
        {
            ["Version"] = "2012-10-17",
            ["Id"] = "PutObjPolicy",
            ["Statement"] = new Dictionary<string, object?>
            {
                ["Sid"] = "DenyObjectsThatAreNotSSEKMS",
                ["Principal"] = "*",
                ["Effect"] = "Deny",
                ["Action"] = "s3:PutObject",
                ["Resource"] = myBucket.Bucket.Apply(bucket => $"arn:aws:s3:::{bucket}/*"),
                ["Condition"] = new Dictionary<string, object?>
                {
                    ["Null"] = new Dictionary<string, object?>
                    {
                        ["s3:x-amz-server-side-encryption-aws-kms-key-id"] = "true",
                    },
                },
            },
        })),
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
import com.pulumi.aws.s3.BucketV2;
import com.pulumi.aws.s3.BucketPolicy;
import com.pulumi.aws.s3.BucketPolicyArgs;
import static com.pulumi.codegen.internal.Serialization.*;
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
        var myBucket = new BucketV2("myBucket");

        var myBucketPolicy = new BucketPolicy("myBucketPolicy", BucketPolicyArgs.builder()
            .bucket("my-bucket-26224916")
            .policy(myBucket.bucket().applyValue(bucket -> serializeJson(
                jsonObject(
                    jsonProperty("Version", "2012-10-17"),
                    jsonProperty("Id", "PutObjPolicy"),
                    jsonProperty("Statement", jsonObject(
                        jsonProperty("Sid", "DenyObjectsThatAreNotSSEKMS"),
                        jsonProperty("Principal", "*"),
                        jsonProperty("Effect", "Deny"),
                        jsonProperty("Action", "s3:PutObject"),
                        jsonProperty("Resource", String.format("arn:aws:s3:::%s/*", bucket)),
                        jsonProperty("Condition", jsonObject(
                            jsonProperty("Null", jsonObject(
                                jsonProperty("s3:x-amz-server-side-encryption-aws-kms-key-id", "true")
                            ))
                        ))
                    ))
                ))))
            .build());

    }
}

```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
name: example
runtime: yaml
resources:
  my-bucket:
    type: aws:s3:BucketV2
  my-bucket-policy:
    type: aws:s3:BucketPolicy
    properties:
      bucket: "my-bucket-26224916"
      policy:
        fn::toJSON:
          Version: 2012-10-17
          Id: PutObjPolicy
          Statement:
            Sid: DenyObjectsThatAreNotSSEKMS
            Principal: "*"
            Effect: Deny
            Action: s3:PutObject
            Resource: "arn:aws:s3:::${my-bucket.bucket}/*"
            Condition:
              "Null":
                s3:x-amz-server-side-encryption-aws-kms-key-id: "true"
```

{{% /choosable %}}

{{< /chooser >}}


As a bonus, the policy can now more easily refer to the concrete name of the bucket.

### serverSideEncryptionConfiguration input

To specify [server-side encryption configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html),
instead of:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}

{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const myBucket = new aws.s3.Bucket("my-bucket", {serverSideEncryptionConfiguration: {
    rule: {
        applyServerSideEncryptionByDefault: {
            sseAlgorithm: "aws:kms",
        },
    },
}});

```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi
import pulumi_aws as aws

my_bucket = aws.s3.Bucket("my-bucket", server_side_encryption_configuration={
    "rule": {
        "apply_server_side_encryption_by_default": {
            "sse_algorithm": "aws:kms",
        },
    },
})

```

{{% /choosable %}}

{{% choosable language go %}}

```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := s3.NewBucket(ctx, "my-bucket", &s3.BucketArgs{
			ServerSideEncryptionConfiguration: &s3.BucketServerSideEncryptionConfigurationArgs{
				Rule: &s3.BucketServerSideEncryptionConfigurationRuleArgs{
					ApplyServerSideEncryptionByDefault: &s3.BucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultArgs{
						SseAlgorithm: pulumi.String("aws:kms"),
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

{{% choosable language csharp %}}

```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var myBucket = new Aws.S3.Bucket("my-bucket", new()
    {
        ServerSideEncryptionConfiguration = new Aws.S3.Inputs.BucketServerSideEncryptionConfigurationArgs
        {
            Rule = new Aws.S3.Inputs.BucketServerSideEncryptionConfigurationRuleArgs
            {
                ApplyServerSideEncryptionByDefault = new Aws.S3.Inputs.BucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultArgs
                {
                    SseAlgorithm = "aws:kms",
                },
            },
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
import com.pulumi.aws.s3.Bucket;
import com.pulumi.aws.s3.BucketArgs;
import com.pulumi.aws.s3.inputs.BucketServerSideEncryptionConfigurationArgs;
import com.pulumi.aws.s3.inputs.BucketServerSideEncryptionConfigurationRuleArgs;
import com.pulumi.aws.s3.inputs.BucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultArgs;
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
        var myBucket = new Bucket("myBucket", BucketArgs.builder()
            .serverSideEncryptionConfiguration(BucketServerSideEncryptionConfigurationArgs.builder()
                .rule(BucketServerSideEncryptionConfigurationRuleArgs.builder()
                    .applyServerSideEncryptionByDefault(BucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultArgs.builder()
                        .sseAlgorithm("aws:kms")
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
name: example
runtime: yaml
resources:
  my-bucket:
    type: aws:s3:Bucket
    properties:
      serverSideEncryptionConfiguration:
        rule:
          applyServerSideEncryptionByDefault:
            sseAlgorithm: "aws:kms"
```

{{% /choosable %}}

{{< /chooser >}}



Use the `aws.s3.BucketServerSideEncryptionConfiguration` resource:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}

{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const myBucket = new aws.s3.BucketV2("my-bucket", {});
const myBucketSseConfig = new aws.s3.BucketServerSideEncryptionConfigurationV2("my-bucket-sse-config", {
    bucket: myBucket.bucket,
    rules: [{
        applyServerSideEncryptionByDefault: {
            sseAlgorithm: "aws:kms",
        },
    }],
});

```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi
import pulumi_aws as aws

my_bucket = aws.s3.BucketV2("my-bucket")
my_bucket_sse_config = aws.s3.BucketServerSideEncryptionConfigurationV2("my-bucket-sse-config",
    bucket=my_bucket.bucket,
    rules=[{
        "apply_server_side_encryption_by_default": {
            "sse_algorithm": "aws:kms",
        },
    }])

```

{{% /choosable %}}

{{% choosable language go %}}

```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		myBucket, err := s3.NewBucketV2(ctx, "my-bucket", nil)
		if err != nil {
			return err
		}
		_, err = s3.NewBucketServerSideEncryptionConfigurationV2(ctx, "my-bucket-sse-config", &s3.BucketServerSideEncryptionConfigurationV2Args{
			Bucket: myBucket.Bucket,
			Rules: s3.BucketServerSideEncryptionConfigurationV2RuleArray{
				&s3.BucketServerSideEncryptionConfigurationV2RuleArgs{
					ApplyServerSideEncryptionByDefault: &s3.BucketServerSideEncryptionConfigurationV2RuleApplyServerSideEncryptionByDefaultArgs{
						SseAlgorithm: pulumi.String("aws:kms"),
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

{{% choosable language csharp %}}

```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var myBucket = new Aws.S3.BucketV2("my-bucket");

    var myBucketSseConfig = new Aws.S3.BucketServerSideEncryptionConfigurationV2("my-bucket-sse-config", new()
    {
        Bucket = myBucket.Bucket,
        Rules = new[]
        {
            new Aws.S3.Inputs.BucketServerSideEncryptionConfigurationV2RuleArgs
            {
                ApplyServerSideEncryptionByDefault = new Aws.S3.Inputs.BucketServerSideEncryptionConfigurationV2RuleApplyServerSideEncryptionByDefaultArgs
                {
                    SseAlgorithm = "aws:kms",
                },
            },
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
import com.pulumi.aws.s3.BucketV2;
import com.pulumi.aws.s3.BucketServerSideEncryptionConfigurationV2;
import com.pulumi.aws.s3.BucketServerSideEncryptionConfigurationV2Args;
import com.pulumi.aws.s3.inputs.BucketServerSideEncryptionConfigurationV2RuleArgs;
import com.pulumi.aws.s3.inputs.BucketServerSideEncryptionConfigurationV2RuleApplyServerSideEncryptionByDefaultArgs;
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
        var myBucket = new BucketV2("myBucket");

        var myBucketSseConfig = new BucketServerSideEncryptionConfigurationV2("myBucketSseConfig", BucketServerSideEncryptionConfigurationV2Args.builder()
            .bucket(myBucket.bucket())
            .rules(BucketServerSideEncryptionConfigurationV2RuleArgs.builder()
                .applyServerSideEncryptionByDefault(BucketServerSideEncryptionConfigurationV2RuleApplyServerSideEncryptionByDefaultArgs.builder()
                    .sseAlgorithm("aws:kms")
                    .build())
                .build())
            .build());

    }
}

```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
name: example
runtime: yaml
resources:
  my-bucket:
    type: aws:s3:BucketV2
  my-bucket-sse-config:
    type: aws:s3:BucketServerSideEncryptionConfigurationV2
    properties:
      bucket: ${my-bucket.bucket}
      rules:
        - applyServerSideEncryptionByDefault:
            sseAlgorithm: "aws:kms"
```

{{% /choosable %}}

{{< /chooser >}}



### acceleration input

To enable acceleration, instead of:

``` typescript
new aws.s3.Bucket("my-bucket", {
  accelerationStatus: "Enabled",
});
```

Use the `aws.s3.BucketAccelerationConfiguration` resource:

``` typescript
const myBucket = new aws.s3.BucketV2("my-new-bucket", {});

new aws.s3.BucketAccelerateConfigurationV2("my-bucket-acceleration", {
  bucket: myBucket.bucket,
  status: "Enabled",
});
```

### corsRules input

To configure [Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html),
instead of:

``` typescript
new aws.s3.Bucket("my-bucket", {
  corsRules: [
    {
      allowedHeaders: ["*"],
      allowedMethods: ["GET"],
      allowedOrigins: ["*"],
      exposeHeaders: ["ETag"],
      maxAgeSeconds: 3000,
    },
  ],
});
```

Use the `aws.s3.BucketCorsConfiguration` resource:

``` typescript
const myBucket = new aws.s3.BucketV2("my-new-bucket", {});

new aws.s3.BucketCorsConfigurationV2("my-bucket-cors", {
  bucket: myBucket.bucket,
  corsRules: [
    {
      allowedHeaders: ["*"],
      allowedMethods: ["GET"],
      allowedOrigins: ["*"],
      exposeHeaders: ["ETag"],
      maxAgeSeconds: 3000,
    },
  ]
});
```

### acl and grants inputs

The `aws.s3.BucketAcl` resource is now required to configure either
[canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)
or [ACL policy grant](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#sample-acl),
and it replaces the use of `acl` and `grants` inputs. To illustrate migrating the uses of `grants`, instead of:

``` typescript
const currentUser = aws.s3.getCanonicalUserIdOutput({});

new aws.s3.Bucket("my-bucket", {
  grants: [
    {
      permissions: ["READ"],
      type: "CanonicalUser",
      id: currentUser.id,
    },
  ],
});
```

Use the `aws.s3.BucketAcl` resource like this:

``` typescript
const myBucket = new aws.s3.BucketV2("my-new-bucket", {});

const myBucketOwnershipControls = new aws.s3.BucketOwnershipControls("my-new-bucket-oc", {
  bucket: myBucket.id,
  rule: {
    objectOwnership: "BucketOwnerPreferred",
  },
});

new aws.s3.BucketAclV2("grant", {
  bucket: myBucket.bucket,
  accessControlPolicy: {
    owner: {id: currentUser.id},
    grants: [
      {
        permission: "READ",
        grantee: {
          type: "CanonicalUser",
          id: currentUser.id,
        },
      },
    ],
  },
}, {
  dependsOn: [myBucketOwnershipControls]
});
```

Note that `dependsOn` on `BucketOwnershipControls` is required for Pulumi to correctly order the operations for this
infrastructure. Consult the documentation on `aws.s3.BucketAclV2` for more fully worked examples of configuring ACL.

### hostedZoneId input

This will only be available as an output. If your program specified this as an input, simply remove it. The input was
ignored by prior versions of the provider and was exposed in error.

### other inputs

All other removed inputs are treated similarly. Consult the documentation for the matching resource replacing the input
for more information:

| Input                    | Resource                                   |
|--------------------------|--------------------------------------------|
| lifecycleRules           | aws.s3.BucketLifecycleConfigurationV2      |
| loggings                 | aws.s3.BucketLoggingV2                     |
| objectLockConfiguration  | aws.s3.BucketObjectLockConfigurationV2     |
| replicationConfiguration | aws.s3.BucketReplicationConfig             |
| requestPayer             | aws.s3.BucketRequestPaymentConfigurationV2 |
| versionings              | aws.s3.BucketVersioningV2                  |
| website                  | aws.s3.BucketWebsiteConfigurationV2        |
| websiteDomain            | aws.s3.BucketWebsiteConfigurationV2        |
| websiteEndpoint          | aws.s3.BucketWebsiteConfigurationV2        |


## Avoiding replacement

In situations when replacing the bucket in the AWS account is not acceptable, it is possible to perform a manual
migration that changes Pulumi program and state to track buckets using the new resource without executing any changes
against the actual cloud account. While the details will vary depending on your use case, this procedure generally
involves the following steps:

1. Find URNs for legacy Bucket Pulumi resources using `pulumi stack export`
2. Determine the actual bucket name(s)
3. Determine which side-by-side resources will be needed for each bucket
4. Construct an `pulumi-import.json` file listing the buckets and their side-by-side resources
5. Run `pulumi import --file import-file.json` using the [Bulk Importing](https://www.pulumi.com/tutorials/importing/bulk-importing/) feature
6. Add the suggested code into your Pulumi program source
7. Remove the legacy Bucket code from your Pulumi program source
8. Remove the legacy Bucket resources from state using `pulumi state delete $bucketURN`
9. Run `pulumi preview` to confirm a no-change plan

Consider a concrete example, suppose you have provisioned a bucket with a `serverSideEncryptionConfiguration` as
follows:

```typescript
import * as aws from "@pulumi/aws";

new aws.s3.Bucket("my-bucket", {
  serverSideEncryptionConfiguration: {
    rule: {
      applyServerSideEncryptionByDefault: {
        sseAlgorithm: "aws:kms",
      },
    },
  },
});
```

1. Scanning through the state in `pulumi stack export`, observe and note its URN is
   `"urn:pulumi:dev::bucket-playground::aws:s3/bucket:Bucket::my-bucket"`

2. The state file should also include the actual cloud name for the bucket such as `"bucket": "my-bucket-cd24744"`

3. This bucket will require a `BucketServerSideEncryptionConfiguration` side-by-side resource

4. The import file should therefore look like this:

    ```json
    {
      "resources": [
        {
          "type": "aws:s3/bucketV2:BucketV2",
          "name": "my-bucket",
          "id": "my-bucket-cd24744"
        },
        {
          "type": "aws:s3/bucketServerSideEncryptionConfigurationV2:BucketServerSideEncryptionConfigurationV2",
          "name": "my-bucket-encryption-configuration",
          "id": "my-bucket-cd24744"
        }
      ]
    }
    ```

5. `pulumi import --file import-file.json` will suggest new code to include in your program, for example:

    ```typescript
    const my_bucket = new aws.s3.BucketV2("my-bucket", {
        bucket: "my-bucket-cd24744",
    }, {
        protect: true,
    });

    const my_bucket_encryption_configuration = new aws.s3.BucketServerSideEncryptionConfigurationV2("my-bucket-encryption-configuration", {
        bucket: "my-bucket-cd24744",
        rules: [{
            applyServerSideEncryptionByDefault: {
                sseAlgorithm: "aws:kms",
            },
        }],
    }, {
        protect: true,
    });
    ```

6. Run `pulumi state delete "urn:pulumi:dev::bucket-playground::aws:s3/bucket:Bucket::my-bucket"` to remove the old
   bucket from the state

7. Delete the code for the old bucket from the sources.

8. `pulumi preview` should result in `Resources: N unchanged` to confirm everything went well.
