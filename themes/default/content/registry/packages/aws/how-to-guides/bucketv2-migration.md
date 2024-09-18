---
title: "BucketV2 Migration Guide"
h1: "Migrating from aws.s3.Bucket to aws.s3.BucketV2"
meta_desc: Practitioner level instructions for migrating from aws.s3.Bucket to aws.s3.BucketV2 resources
layout: package
---

In the upcoming AWS Classic major release (v7), `aws.s3.Bucket` will be discontinued in favor of `BucketV2`. Users of
`aws.s3.Bucket` resource are encouraged to migrate early. The migration requires a significant refactor to the source
code and additional steps for avoiding data loss. This guide aims to cover all relevant scenarios with precise migration
instructions.

To migrate existing `aws.s3.Bucket` resources to `aws.s3.BucketV2`:

1. Edit your Pulumi program sources to replace removed inputs with equivalent side-by-side resources. Specifically
   `BucketV2` inputs such as `policy` and `accelerationStatus` are to be replaced by side-by-side resources
   `aws.s3.BucketPolicy` and `aws.s3.BucketAccelerateConfigurationV2`.

2. Perform `pulumi up` to replace the buckets in AWS. **WARNING**: replacing the buckets will delete and re-create them,
   losing any data stored in these buckets. If this is not acceptable, consider using `pulumi import` to migrate
   manually instead as outlined in the "Avoiding replacement" section.

## Migrating deprecated inputs

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
    bucket: myBucket.bucket,
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
    bucket=my_bucket.bucket,
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
			Bucket: myBucket.Bucket,
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
        Bucket = myBucket.Bucket,
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
            .bucket(myBucket.bucket())
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
      bucket: ${my-bucket.bucket}
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

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}

{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const myBucket = new aws.s3.Bucket("my-bucket", {accelerationStatus: "Enabled"});

```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi
import pulumi_aws as aws

my_bucket = aws.s3.Bucket("my-bucket", acceleration_status="Enabled")

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
			AccelerationStatus: pulumi.String("Enabled"),
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
        AccelerationStatus = "Enabled",
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
            .accelerationStatus("Enabled")
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
      accelerationStatus: Enabled
```

{{% /choosable %}}

{{< /chooser >}}




Use the `aws.s3.BucketAccelerateConfigurationV2` resource:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}

{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const myBucket = new aws.s3.BucketV2("my-bucket", {});
const myBucketAcceleration = new aws.s3.BucketAccelerateConfigurationV2("my-bucket-acceleration", {
    bucket: myBucket.bucket,
    status: "Enabled",
});

```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi
import pulumi_aws as aws

my_bucket = aws.s3.BucketV2("my-bucket")
my_bucket_acceleration = aws.s3.BucketAccelerateConfigurationV2("my-bucket-acceleration",
    bucket=my_bucket.bucket,
    status="Enabled")

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
		_, err = s3.NewBucketAccelerateConfigurationV2(ctx, "my-bucket-acceleration", &s3.BucketAccelerateConfigurationV2Args{
			Bucket: myBucket.Bucket,
			Status: pulumi.String("Enabled"),
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

    var myBucketAcceleration = new Aws.S3.BucketAccelerateConfigurationV2("my-bucket-acceleration", new()
    {
        Bucket = myBucket.Bucket,
        Status = "Enabled",
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
import com.pulumi.aws.s3.BucketAccelerateConfigurationV2;
import com.pulumi.aws.s3.BucketAccelerateConfigurationV2Args;
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

        var myBucketAcceleration = new BucketAccelerateConfigurationV2("myBucketAcceleration", BucketAccelerateConfigurationV2Args.builder()
            .bucket(myBucket.bucket())
            .status("Enabled")
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
  my-bucket-acceleration:
    type: aws:s3:BucketAccelerateConfigurationV2
    properties:
      bucket: ${my-bucket.bucket}
      status: Enabled
```

{{% /choosable %}}

{{< /chooser >}}


### corsRules input

To configure [Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html),
instead of:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}

{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const myBucket = new aws.s3.Bucket("my-bucket", {corsRules: [{
    allowedHeaders: ["*"],
    allowedMethods: ["GET"],
    allowedOrigins: ["*"],
    exposeHeaders: ["ETag"],
    maxAgeSeconds: 3000,
}]});

```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi
import pulumi_aws as aws

my_bucket = aws.s3.Bucket("my-bucket", cors_rules=[{
    "allowed_headers": ["*"],
    "allowed_methods": ["GET"],
    "allowed_origins": ["*"],
    "expose_headers": ["ETag"],
    "max_age_seconds": 3000,
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
		_, err := s3.NewBucket(ctx, "my-bucket", &s3.BucketArgs{
			CorsRules: s3.BucketCorsRuleArray{
				&s3.BucketCorsRuleArgs{
					AllowedHeaders: pulumi.StringArray{
						pulumi.String("*"),
					},
					AllowedMethods: pulumi.StringArray{
						pulumi.String("GET"),
					},
					AllowedOrigins: pulumi.StringArray{
						pulumi.String("*"),
					},
					ExposeHeaders: pulumi.StringArray{
						pulumi.String("ETag"),
					},
					MaxAgeSeconds: pulumi.Int(3000),
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
        CorsRules = new[]
        {
            new Aws.S3.Inputs.BucketCorsRuleArgs
            {
                AllowedHeaders = new[]
                {
                    "*",
                },
                AllowedMethods = new[]
                {
                    "GET",
                },
                AllowedOrigins = new[]
                {
                    "*",
                },
                ExposeHeaders = new[]
                {
                    "ETag",
                },
                MaxAgeSeconds = 3000,
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
import com.pulumi.aws.s3.inputs.BucketCorsRuleArgs;
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
            .corsRules(BucketCorsRuleArgs.builder()
                .allowedHeaders("*")
                .allowedMethods("GET")
                .allowedOrigins("*")
                .exposeHeaders("ETag")
                .maxAgeSeconds(3000)
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
      corsRules:
        - allowedHeaders:
            - "*"
          allowedMethods:
            - "GET"
          allowedOrigins:
            - "*"
          exposeHeaders:
            - "ETag"
          maxAgeSeconds: 3000
```

{{% /choosable %}}

{{< /chooser >}}



Use the `aws.s3.BucketCorsConfigurationV2` resource:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}

{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const myBucket = new aws.s3.BucketV2("my-bucket", {});
const myBucketCors = new aws.s3.BucketCorsConfigurationV2("my-bucket-cors", {
    bucket: myBucket.bucket,
    corsRules: [{
        allowedHeaders: ["*"],
        allowedMethods: ["GET"],
        allowedOrigins: ["*"],
        exposeHeaders: ["ETag"],
        maxAgeSeconds: 3000,
    }],
});

```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi
import pulumi_aws as aws

my_bucket = aws.s3.BucketV2("my-bucket")
my_bucket_cors = aws.s3.BucketCorsConfigurationV2("my-bucket-cors",
    bucket=my_bucket.bucket,
    cors_rules=[{
        "allowed_headers": ["*"],
        "allowed_methods": ["GET"],
        "allowed_origins": ["*"],
        "expose_headers": ["ETag"],
        "max_age_seconds": 3000,
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
		_, err = s3.NewBucketCorsConfigurationV2(ctx, "my-bucket-cors", &s3.BucketCorsConfigurationV2Args{
			Bucket: myBucket.Bucket,
			CorsRules: s3.BucketCorsConfigurationV2CorsRuleArray{
				&s3.BucketCorsConfigurationV2CorsRuleArgs{
					AllowedHeaders: pulumi.StringArray{
						pulumi.String("*"),
					},
					AllowedMethods: pulumi.StringArray{
						pulumi.String("GET"),
					},
					AllowedOrigins: pulumi.StringArray{
						pulumi.String("*"),
					},
					ExposeHeaders: pulumi.StringArray{
						pulumi.String("ETag"),
					},
					MaxAgeSeconds: pulumi.Int(3000),
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

    var myBucketCors = new Aws.S3.BucketCorsConfigurationV2("my-bucket-cors", new()
    {
        Bucket = myBucket.Bucket,
        CorsRules = new[]
        {
            new Aws.S3.Inputs.BucketCorsConfigurationV2CorsRuleArgs
            {
                AllowedHeaders = new[]
                {
                    "*",
                },
                AllowedMethods = new[]
                {
                    "GET",
                },
                AllowedOrigins = new[]
                {
                    "*",
                },
                ExposeHeaders = new[]
                {
                    "ETag",
                },
                MaxAgeSeconds = 3000,
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
import com.pulumi.aws.s3.BucketCorsConfigurationV2;
import com.pulumi.aws.s3.BucketCorsConfigurationV2Args;
import com.pulumi.aws.s3.inputs.BucketCorsConfigurationV2CorsRuleArgs;
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

        var myBucketCors = new BucketCorsConfigurationV2("myBucketCors", BucketCorsConfigurationV2Args.builder()
            .bucket(myBucket.bucket())
            .corsRules(BucketCorsConfigurationV2CorsRuleArgs.builder()
                .allowedHeaders("*")
                .allowedMethods("GET")
                .allowedOrigins("*")
                .exposeHeaders("ETag")
                .maxAgeSeconds(3000)
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
  my-bucket-cors:
    type: aws:s3:BucketCorsConfigurationV2
    properties:
      bucket: ${my-bucket.bucket}
      corsRules:
        - allowedHeaders:
            - "*"
          allowedMethods:
            - "GET"
          allowedOrigins:
            - "*"
          exposeHeaders:
            - "ETag"
          maxAgeSeconds: 3000
```

{{% /choosable %}}

{{< /chooser >}}



### acl and grants inputs

The `aws.s3.BucketAcl` resource is now required to configure either
[canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)
or [ACL policy grant](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#sample-acl),
and it replaces the use of `acl` and `grants` inputs. To illustrate migrating the uses of `grants`, instead of:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}

{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const currentUser = aws.s3.getCanonicalUserIdOutput({});
const myBucket = new aws.s3.Bucket("my-bucket", {grants: [{
    permissions: ["READ"],
    type: "CanonicalUser",
    id: currentUser.apply(currentUser => currentUser.id),
}]});

```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi
import pulumi_aws as aws

current_user = aws.s3.get_canonical_user_id_output()
my_bucket = aws.s3.Bucket("my-bucket", grants=[{
    "permissions": ["READ"],
    "type": "CanonicalUser",
    "id": current_user.id,
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
		currentUser, err := s3.GetCanonicalUserId(ctx, nil, nil)
		if err != nil {
			return err
		}
		_, err = s3.NewBucket(ctx, "my-bucket", &s3.BucketArgs{
			Grants: s3.BucketGrantArray{
				&s3.BucketGrantArgs{
					Permissions: pulumi.StringArray{
						pulumi.String("READ"),
					},
					Type: pulumi.String("CanonicalUser"),
					Id:   pulumi.String(currentUser.Id),
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
    var currentUser = Aws.S3.GetCanonicalUserId.Invoke();

    var myBucket = new Aws.S3.Bucket("my-bucket", new()
    {
        Grants = new[]
        {
            new Aws.S3.Inputs.BucketGrantArgs
            {
                Permissions = new[]
                {
                    "READ",
                },
                Type = "CanonicalUser",
                Id = currentUser.Apply(getCanonicalUserIdResult => getCanonicalUserIdResult.Id),
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
import com.pulumi.aws.s3.S3Functions;
import com.pulumi.aws.s3.Bucket;
import com.pulumi.aws.s3.BucketArgs;
import com.pulumi.aws.s3.inputs.BucketGrantArgs;
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
        final var currentUser = S3Functions.getCanonicalUserId();

        var myBucket = new Bucket("myBucket", BucketArgs.builder()
            .grants(BucketGrantArgs.builder()
                .permissions("READ")
                .type("CanonicalUser")
                .id(currentUser.applyValue(getCanonicalUserIdResult -> getCanonicalUserIdResult.id()))
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
variables:
  currentUser:
    fn::invoke:
      function: aws:s3:getCanonicalUserId
      arguments: {}
resources:
  my-bucket:
    type: aws:s3:Bucket
    properties:
      grants:
        - permissions:
            - "READ"
          type: "CanonicalUser"
          id: ${currentUser.id}
```

{{% /choosable %}}

{{< /chooser >}}



Use the `aws.s3.BucketAcl` resource like this:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}

{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const currentUser = aws.s3.getCanonicalUserIdOutput({});
const myBucket = new aws.s3.BucketV2("my-bucket", {});
const myBucketOwnership = new aws.s3.BucketOwnershipControls("my-bucket-ownership", {
    bucket: myBucket.id,
    rule: {
        objectOwnership: "BucketOwnerPreferred",
    },
});
const myBucketAcl = new aws.s3.BucketAclV2("my-bucket-acl", {
    bucket: myBucket.bucket,
    accessControlPolicy: {
        owner: {
            id: currentUser.apply(currentUser => currentUser.id),
        },
        grants: [{
            permission: "READ",
            grantee: {
                type: "CanonicalUser",
                id: currentUser.apply(currentUser => currentUser.id),
            },
        }],
    },
}, {
    dependsOn: [myBucketOwnership],
});

```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi
import pulumi_aws as aws

current_user = aws.s3.get_canonical_user_id_output()
my_bucket = aws.s3.BucketV2("my-bucket")
my_bucket_ownership = aws.s3.BucketOwnershipControls("my-bucket-ownership",
    bucket=my_bucket.id,
    rule={
        "object_ownership": "BucketOwnerPreferred",
    })
my_bucket_acl = aws.s3.BucketAclV2("my-bucket-acl",
    bucket=my_bucket.bucket,
    access_control_policy={
        "owner": {
            "id": current_user.id,
        },
        "grants": [{
            "permission": "READ",
            "grantee": {
                "type": "CanonicalUser",
                "id": current_user.id,
            },
        }],
    },
    opts = pulumi.ResourceOptions(depends_on=[my_bucket_ownership]))

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
		currentUser, err := s3.GetCanonicalUserId(ctx, nil, nil)
		if err != nil {
			return err
		}
		myBucket, err := s3.NewBucketV2(ctx, "my-bucket", nil)
		if err != nil {
			return err
		}
		myBucketOwnership, err := s3.NewBucketOwnershipControls(ctx, "my-bucket-ownership", &s3.BucketOwnershipControlsArgs{
			Bucket: myBucket.ID(),
			Rule: &s3.BucketOwnershipControlsRuleArgs{
				ObjectOwnership: pulumi.String("BucketOwnerPreferred"),
			},
		})
		if err != nil {
			return err
		}
		_, err = s3.NewBucketAclV2(ctx, "my-bucket-acl", &s3.BucketAclV2Args{
			Bucket: myBucket.Bucket,
			AccessControlPolicy: &s3.BucketAclV2AccessControlPolicyArgs{
				Owner: &s3.BucketAclV2AccessControlPolicyOwnerArgs{
					Id: pulumi.String(currentUser.Id),
				},
				Grants: s3.BucketAclV2AccessControlPolicyGrantArray{
					&s3.BucketAclV2AccessControlPolicyGrantArgs{
						Permission: pulumi.String("READ"),
						Grantee: &s3.BucketAclV2AccessControlPolicyGrantGranteeArgs{
							Type: pulumi.String("CanonicalUser"),
							Id:   pulumi.String(currentUser.Id),
						},
					},
				},
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			myBucketOwnership,
		}))
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
    var currentUser = Aws.S3.GetCanonicalUserId.Invoke();

    var myBucket = new Aws.S3.BucketV2("my-bucket");

    var myBucketOwnership = new Aws.S3.BucketOwnershipControls("my-bucket-ownership", new()
    {
        Bucket = myBucket.Id,
        Rule = new Aws.S3.Inputs.BucketOwnershipControlsRuleArgs
        {
            ObjectOwnership = "BucketOwnerPreferred",
        },
    });

    var myBucketAcl = new Aws.S3.BucketAclV2("my-bucket-acl", new()
    {
        Bucket = myBucket.Bucket,
        AccessControlPolicy = new Aws.S3.Inputs.BucketAclV2AccessControlPolicyArgs
        {
            Owner = new Aws.S3.Inputs.BucketAclV2AccessControlPolicyOwnerArgs
            {
                Id = currentUser.Apply(getCanonicalUserIdResult => getCanonicalUserIdResult.Id),
            },
            Grants = new[]
            {
                new Aws.S3.Inputs.BucketAclV2AccessControlPolicyGrantArgs
                {
                    Permission = "READ",
                    Grantee = new Aws.S3.Inputs.BucketAclV2AccessControlPolicyGrantGranteeArgs
                    {
                        Type = "CanonicalUser",
                        Id = currentUser.Apply(getCanonicalUserIdResult => getCanonicalUserIdResult.Id),
                    },
                },
            },
        },
    }, new CustomResourceOptions
    {
        DependsOn =
        {
            myBucketOwnership,
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
import com.pulumi.aws.s3.S3Functions;
import com.pulumi.aws.s3.BucketV2;
import com.pulumi.aws.s3.BucketOwnershipControls;
import com.pulumi.aws.s3.BucketOwnershipControlsArgs;
import com.pulumi.aws.s3.inputs.BucketOwnershipControlsRuleArgs;
import com.pulumi.aws.s3.BucketAclV2;
import com.pulumi.aws.s3.BucketAclV2Args;
import com.pulumi.aws.s3.inputs.BucketAclV2AccessControlPolicyArgs;
import com.pulumi.aws.s3.inputs.BucketAclV2AccessControlPolicyOwnerArgs;
import com.pulumi.resources.CustomResourceOptions;
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
        final var currentUser = S3Functions.getCanonicalUserId();

        var myBucket = new BucketV2("myBucket");

        var myBucketOwnership = new BucketOwnershipControls("myBucketOwnership", BucketOwnershipControlsArgs.builder()
            .bucket(myBucket.id())
            .rule(BucketOwnershipControlsRuleArgs.builder()
                .objectOwnership("BucketOwnerPreferred")
                .build())
            .build());

        var myBucketAcl = new BucketAclV2("myBucketAcl", BucketAclV2Args.builder()
            .bucket(myBucket.bucket())
            .accessControlPolicy(BucketAclV2AccessControlPolicyArgs.builder()
                .owner(BucketAclV2AccessControlPolicyOwnerArgs.builder()
                    .id(currentUser.applyValue(getCanonicalUserIdResult -> getCanonicalUserIdResult.id()))
                    .build())
                .grants(BucketAclV2AccessControlPolicyGrantArgs.builder()
                    .permission("READ")
                    .grantee(BucketAclV2AccessControlPolicyGrantGranteeArgs.builder()
                        .type("CanonicalUser")
                        .id(currentUser.applyValue(getCanonicalUserIdResult -> getCanonicalUserIdResult.id()))
                        .build())
                    .build())
                .build())
            .build(), CustomResourceOptions.builder()
                .dependsOn(myBucketOwnership)
                .build());

    }
}

```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
name: example
runtime: yaml
variables:
  currentUser:
    fn::invoke:
      function: aws:s3:getCanonicalUserId
      arguments: {}
resources:
  my-bucket:
    type: aws:s3:BucketV2
  my-bucket-ownership:
    type: aws:s3:BucketOwnershipControls
    properties:
      bucket: ${my-bucket.id}
      rule:
        objectOwnership: "BucketOwnerPreferred"
  my-bucket-acl:
    type: aws:s3:BucketAclV2
    properties:
      bucket: ${my-bucket.bucket}
      accessControlPolicy:
        owner:
          id: ${currentUser.id}
        grants:
          - permission: "READ"
            grantee:
              type: "CanonicalUser"
              id: ${currentUser.id}
    options:
      dependsOn:
        - ${my-bucket-ownership}
```

{{% /choosable %}}

{{< /chooser >}}




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

- Find URNs for legacy Bucket Pulumi resources using `pulumi stack export`
- Determine the actual bucket name(s)
- Remove the legacy Bucket code from your Pulumi program source
- Remove the legacy Bucket resources from state using `pulumi state delete $bucketURN`
- Determine which side-by-side resources will be needed for each bucket
- Construct an `pulumi-import.json` file listing the buckets and their side-by-side resources
- Run `pulumi import --file import-file.json` using the [Bulk Importing](https://www.pulumi.com/tutorials/importing/bulk-importing/) feature
- Add the suggested code into your Pulumi program source
- Run `pulumi preview` to confirm a no-change plan
- If warnings are generated, edit the program to remove deprecated inputs from BucketV2
- Run `pulumi preview` one more time to confirm a no-change plan on the final program

Consider a concrete example, suppose you have provisioned a bucket as follows:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}

{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const myBucket = new aws.s3.Bucket("my-bucket", {
    bucket: "my-bucket-26224917",
    serverSideEncryptionConfiguration: {
        rule: {
            applyServerSideEncryptionByDefault: {
                sseAlgorithm: "AES256",
            },
        },
    },
    lifecycleRules: [{
        enabled: true,
        expiration: {
            days: 30,
        },
    }],
    policy: JSON.stringify({
        Version: "2012-10-17",
        Id: "PutObjPolicy",
        Statement: [{
            Sid: "DenyObjectsThatAreNotSSEKMS",
            Principal: "*",
            Effect: "Deny",
            Action: "s3:PutObject",
            Resource: "arn:aws:s3:::my-bucket-26224917/*",
            Condition: {
                Null: {
                    "s3:x-amz-server-side-encryption-aws-kms-key-id": "true",
                },
            },
        }],
    }),
    tags: {
        Environment: "Dev",
    },
    objectLockConfiguration: {
        objectLockEnabled: "Enabled",
    },
    versioning: {
        enabled: true,
    },
});

```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi
import json
import pulumi_aws as aws

my_bucket = aws.s3.Bucket("my-bucket",
    bucket="my-bucket-26224917",
    server_side_encryption_configuration={
        "rule": {
            "apply_server_side_encryption_by_default": {
                "sse_algorithm": "AES256",
            },
        },
    },
    lifecycle_rules=[{
        "enabled": True,
        "expiration": {
            "days": 30,
        },
    }],
    policy=json.dumps({
        "Version": "2012-10-17",
        "Id": "PutObjPolicy",
        "Statement": [{
            "Sid": "DenyObjectsThatAreNotSSEKMS",
            "Principal": "*",
            "Effect": "Deny",
            "Action": "s3:PutObject",
            "Resource": "arn:aws:s3:::my-bucket-26224917/*",
            "Condition": {
                "Null": {
                    "s3:x-amz-server-side-encryption-aws-kms-key-id": "true",
                },
            },
        }],
    }),
    tags={
        "Environment": "Dev",
    },
    object_lock_configuration={
        "object_lock_enabled": "Enabled",
    },
    versioning={
        "enabled": True,
    })

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
			"Statement": []map[string]interface{}{
				map[string]interface{}{
					"Sid":       "DenyObjectsThatAreNotSSEKMS",
					"Principal": "*",
					"Effect":    "Deny",
					"Action":    "s3:PutObject",
					"Resource":  "arn:aws:s3:::my-bucket-26224917/*",
					"Condition": map[string]interface{}{
						"Null": map[string]interface{}{
							"s3:x-amz-server-side-encryption-aws-kms-key-id": "true",
						},
					},
				},
			},
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		_, err = s3.NewBucket(ctx, "my-bucket", &s3.BucketArgs{
			Bucket: pulumi.String("my-bucket-26224917"),
			ServerSideEncryptionConfiguration: &s3.BucketServerSideEncryptionConfigurationArgs{
				Rule: &s3.BucketServerSideEncryptionConfigurationRuleArgs{
					ApplyServerSideEncryptionByDefault: &s3.BucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultArgs{
						SseAlgorithm: pulumi.String("AES256"),
					},
				},
			},
			LifecycleRules: s3.BucketLifecycleRuleArray{
				&s3.BucketLifecycleRuleArgs{
					Enabled: pulumi.Bool(true),
					Expiration: &s3.BucketLifecycleRuleExpirationArgs{
						Days: pulumi.Int(30),
					},
				},
			},
			Policy: pulumi.String(json0),
			Tags: pulumi.StringMap{
				"Environment": pulumi.String("Dev"),
			},
			ObjectLockConfiguration: &s3.BucketObjectLockConfigurationArgs{
				ObjectLockEnabled: pulumi.String("Enabled"),
			},
			Versioning: &s3.BucketVersioningArgs{
				Enabled: pulumi.Bool(true),
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
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var myBucket = new Aws.S3.Bucket("my-bucket", new()
    {
        BucketName = "my-bucket-26224917",
        ServerSideEncryptionConfiguration = new Aws.S3.Inputs.BucketServerSideEncryptionConfigurationArgs
        {
            Rule = new Aws.S3.Inputs.BucketServerSideEncryptionConfigurationRuleArgs
            {
                ApplyServerSideEncryptionByDefault = new Aws.S3.Inputs.BucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultArgs
                {
                    SseAlgorithm = "AES256",
                },
            },
        },
        LifecycleRules = new[]
        {
            new Aws.S3.Inputs.BucketLifecycleRuleArgs
            {
                Enabled = true,
                Expiration = new Aws.S3.Inputs.BucketLifecycleRuleExpirationArgs
                {
                    Days = 30,
                },
            },
        },
        Policy = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["Version"] = "2012-10-17",
            ["Id"] = "PutObjPolicy",
            ["Statement"] = new[]
            {
                new Dictionary<string, object?>
                {
                    ["Sid"] = "DenyObjectsThatAreNotSSEKMS",
                    ["Principal"] = "*",
                    ["Effect"] = "Deny",
                    ["Action"] = "s3:PutObject",
                    ["Resource"] = "arn:aws:s3:::my-bucket-26224917/*",
                    ["Condition"] = new Dictionary<string, object?>
                    {
                        ["Null"] = new Dictionary<string, object?>
                        {
                            ["s3:x-amz-server-side-encryption-aws-kms-key-id"] = "true",
                        },
                    },
                },
            },
        }),
        Tags = 
        {
            { "Environment", "Dev" },
        },
        ObjectLockConfiguration = new Aws.S3.Inputs.BucketObjectLockConfigurationArgs
        {
            ObjectLockEnabled = "Enabled",
        },
        Versioning = new Aws.S3.Inputs.BucketVersioningArgs
        {
            Enabled = true,
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
import com.pulumi.aws.s3.inputs.BucketLifecycleRuleArgs;
import com.pulumi.aws.s3.inputs.BucketLifecycleRuleExpirationArgs;
import com.pulumi.aws.s3.inputs.BucketObjectLockConfigurationArgs;
import com.pulumi.aws.s3.inputs.BucketVersioningArgs;
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
            .bucket("my-bucket-26224917")
            .serverSideEncryptionConfiguration(BucketServerSideEncryptionConfigurationArgs.builder()
                .rule(BucketServerSideEncryptionConfigurationRuleArgs.builder()
                    .applyServerSideEncryptionByDefault(BucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultArgs.builder()
                        .sseAlgorithm("AES256")
                        .build())
                    .build())
                .build())
            .lifecycleRules(BucketLifecycleRuleArgs.builder()
                .enabled(true)
                .expiration(BucketLifecycleRuleExpirationArgs.builder()
                    .days(30)
                    .build())
                .build())
            .policy(serializeJson(
                jsonObject(
                    jsonProperty("Version", "2012-10-17"),
                    jsonProperty("Id", "PutObjPolicy"),
                    jsonProperty("Statement", jsonArray(jsonObject(
                        jsonProperty("Sid", "DenyObjectsThatAreNotSSEKMS"),
                        jsonProperty("Principal", "*"),
                        jsonProperty("Effect", "Deny"),
                        jsonProperty("Action", "s3:PutObject"),
                        jsonProperty("Resource", "arn:aws:s3:::my-bucket-26224917/*"),
                        jsonProperty("Condition", jsonObject(
                            jsonProperty("Null", jsonObject(
                                jsonProperty("s3:x-amz-server-side-encryption-aws-kms-key-id", "true")
                            ))
                        ))
                    )))
                )))
            .tags(Map.of("Environment", "Dev"))
            .objectLockConfiguration(BucketObjectLockConfigurationArgs.builder()
                .objectLockEnabled("Enabled")
                .build())
            .versioning(BucketVersioningArgs.builder()
                .enabled(true)
                .build())
            .build());

    }
}

```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
name: y2
runtime: yaml
resources:
  my-bucket:
    type: aws:s3:Bucket
    properties:
      bucket: "my-bucket-26224917"
      serverSideEncryptionConfiguration:
        rule:
          applyServerSideEncryptionByDefault:
            sseAlgorithm: "AES256"
      lifecycleRules:
         - enabled: true
           expiration:
             days: 30
      policy:
        fn::toJSON:
          Version: "2012-10-17"
          Id: "PutObjPolicy"
          Statement:
            - Sid: "DenyObjectsThatAreNotSSEKMS"
              Principal: "*"
              Effect: "Deny"
              Action: "s3:PutObject"
              Resource: "arn:aws:s3:::my-bucket-26224917/*"
              Condition:
                "Null":
                  "s3:x-amz-server-side-encryption-aws-kms-key-id": "true"
      tags:
        Environment: "Dev"
      objectLockConfiguration:
        objectLockEnabled: "Enabled"
      versioning:
        enabled: true
```

{{% /choosable %}}

{{< /chooser >}}

Migrate as follows:

- Scanning through the state in `pulumi stack export`, observe and note its URN is
  `"urn:pulumi:bucket1::y2::aws:s3/bucket:Bucket::my-bucket"`

- The state file should also include the actual cloud name for the bucket such as `"bucket": "my-bucket-36224917"`

- Run `pulumi state delete "urn:pulumi:bucket1::y2::aws:s3/bucket:Bucket::my-bucket"` to remove the old bucket from the
  state

- Delete the code for the old bucket from the sources.

- This bucket will require the following side-by-side resources:

  aws:s3:BucketServerSideEncryptionConfigurationV2
  aws:s3:BucketLifecycleConfigurationV2
  aws:s3:BucketPolicy
  aws:s3:BucketObjectLockConfigurationV2
  aws:s3:BucketVersioningV2

- The import file should therefore look like this:

    ```json
    {
      "resources": [
        {
          "type": "aws:s3/bucketV2:BucketV2",
          "name": "my-bucket2",
          "id": "my-bucket-36224917"
        },
        {
          "type": "aws:s3/bucketServerSideEncryptionConfigurationV2:BucketServerSideEncryptionConfigurationV2",
          "name": "my-bucket-encryption-configuration",
          "id": "my-bucket-36224917"
        },
        {
          "type": "aws:s3/bucketPolicy:BucketPolicy",
          "name": "my-bucket-policy",
          "id": "my-bucket-36224917"
        },
        {
          "type": "aws:s3/bucketLifecycleConfigurationV2:BucketLifecycleConfigurationV2",
          "name": "my-bucket-policy",
          "id": "my-bucket-36224917"
        },
        {
          "type": "aws:s3/bucketObjectLockConfigurationV2:BucketObjectLockConfigurationV2",
          "name": "my-bucket-object-lock-configuration",
          "id": "my-bucket-36224917"
        },
        {
          "type": "aws:s3/bucketVersioningV2:BucketVersioningV2",
          "name": "my-bucket-versioning",
          "id": "my-bucket-36224917"
        }
      ]
    }
    ```

- `pulumi import --file import-file.json` will suggest new code to include in your program, for example:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}

{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const myBucket = new aws.s3.BucketV2("my-bucket", {
    bucket: "my-bucket-36224917",
    grants: [{
        id: "e07865a5679c7977370948f1f1e51c21b12d8cfdd396a7e3172275d9164e01b8",
        permissions: ["FULL_CONTROL"],
        type: "CanonicalUser",
    }],
    lifecycleRules: [{
        enabled: true,
        expirations: [{
            days: 30,
        }],
        id: "pu-s3-lifecycle-20240918194815251100000001",
    }],
    objectLockConfiguration: {
        objectLockEnabled: "Enabled",
    },
    policy: "{\"Id\":\"PutObjPolicy\",\"Statement\":[{\"Action\":\"s3:PutObject\",\"Condition\":{\"Null\":{\"s3:x-amz-server-side-encryption-aws-kms-key-id\":\"true\"}},\"Effect\":\"Deny\",\"Principal\":\"*\",\"Resource\":\"arn:aws:s3:::my-bucket-36224917/*\",\"Sid\":\"DenyObjectsThatAreNotSSEKMS\"}],\"Version\":\"2012-10-17\"}",
    requestPayer: "BucketOwner",
    serverSideEncryptionConfigurations: [{
        rules: [{
            applyServerSideEncryptionByDefaults: [{
                sseAlgorithm: "AES256",
            }],
        }],
    }],
    tags: {
        Environment: "Dev",
    },
    versionings: [{
        enabled: true,
    }],
}, {
    protect: true,
});
const myBucketEncryptionConfiguration = new aws.s3.BucketServerSideEncryptionConfigurationV2("my-bucket-encryption-configuration", {
    bucket: "my-bucket-36224917",
    rules: [{
        applyServerSideEncryptionByDefault: {
            sseAlgorithm: "AES256",
        },
    }],
}, {
    protect: true,
});
const myBucketPolicy = new aws.s3.BucketPolicy("my-bucket-policy", {
    bucket: "my-bucket-36224917",
    policy: "{\"Id\":\"PutObjPolicy\",\"Statement\":[{\"Action\":\"s3:PutObject\",\"Condition\":{\"Null\":{\"s3:x-amz-server-side-encryption-aws-kms-key-id\":\"true\"}},\"Effect\":\"Deny\",\"Principal\":\"*\",\"Resource\":\"arn:aws:s3:::my-bucket-36224917/*\",\"Sid\":\"DenyObjectsThatAreNotSSEKMS\"}],\"Version\":\"2012-10-17\"}",
}, {
    protect: true,
});
const myBucketLifecycle = new aws.s3.BucketLifecycleConfigurationV2("my-bucket-lifecycle", {
    bucket: "my-bucket-36224917",
    rules: [{
        expiration: {
            days: 30,
        },
        id: "pu-s3-lifecycle-20240918194815251100000001",
        status: "Enabled",
    }],
}, {
    protect: true,
});
const myBucketObjectLockConfiguration = new aws.s3.BucketObjectLockConfigurationV2("my-bucket-object-lock-configuration", {
    bucket: "my-bucket-36224917",
    objectLockEnabled: "Enabled",
}, {
    protect: true,
});
const myBucketVersioning = new aws.s3.BucketVersioningV2("my-bucket-versioning", {
    bucket: "my-bucket-36224917",
    versioningConfiguration: {
        mfaDelete: "Disabled",
        status: "Enabled",
    },
}, {
    protect: true,
});

```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi
import pulumi_aws as aws

my_bucket = aws.s3.BucketV2("my-bucket",
    bucket="my-bucket-36224917",
    grants=[{
        "id": "e07865a5679c7977370948f1f1e51c21b12d8cfdd396a7e3172275d9164e01b8",
        "permissions": ["FULL_CONTROL"],
        "type": "CanonicalUser",
    }],
    lifecycle_rules=[{
        "enabled": True,
        "expirations": [{
            "days": 30,
        }],
        "id": "pu-s3-lifecycle-20240918194815251100000001",
    }],
    object_lock_configuration={
        "object_lock_enabled": "Enabled",
    },
    policy="{\"Id\":\"PutObjPolicy\",\"Statement\":[{\"Action\":\"s3:PutObject\",\"Condition\":{\"Null\":{\"s3:x-amz-server-side-encryption-aws-kms-key-id\":\"true\"}},\"Effect\":\"Deny\",\"Principal\":\"*\",\"Resource\":\"arn:aws:s3:::my-bucket-36224917/*\",\"Sid\":\"DenyObjectsThatAreNotSSEKMS\"}],\"Version\":\"2012-10-17\"}",
    request_payer="BucketOwner",
    server_side_encryption_configurations=[{
        "rules": [{
            "apply_server_side_encryption_by_defaults": [{
                "sse_algorithm": "AES256",
            }],
        }],
    }],
    tags={
        "Environment": "Dev",
    },
    versionings=[{
        "enabled": True,
    }],
    opts = pulumi.ResourceOptions(protect=True))
my_bucket_encryption_configuration = aws.s3.BucketServerSideEncryptionConfigurationV2("my-bucket-encryption-configuration",
    bucket="my-bucket-36224917",
    rules=[{
        "apply_server_side_encryption_by_default": {
            "sse_algorithm": "AES256",
        },
    }],
    opts = pulumi.ResourceOptions(protect=True))
my_bucket_policy = aws.s3.BucketPolicy("my-bucket-policy",
    bucket="my-bucket-36224917",
    policy="{\"Id\":\"PutObjPolicy\",\"Statement\":[{\"Action\":\"s3:PutObject\",\"Condition\":{\"Null\":{\"s3:x-amz-server-side-encryption-aws-kms-key-id\":\"true\"}},\"Effect\":\"Deny\",\"Principal\":\"*\",\"Resource\":\"arn:aws:s3:::my-bucket-36224917/*\",\"Sid\":\"DenyObjectsThatAreNotSSEKMS\"}],\"Version\":\"2012-10-17\"}",
    opts = pulumi.ResourceOptions(protect=True))
my_bucket_lifecycle = aws.s3.BucketLifecycleConfigurationV2("my-bucket-lifecycle",
    bucket="my-bucket-36224917",
    rules=[{
        "expiration": {
            "days": 30,
        },
        "id": "pu-s3-lifecycle-20240918194815251100000001",
        "status": "Enabled",
    }],
    opts = pulumi.ResourceOptions(protect=True))
my_bucket_object_lock_configuration = aws.s3.BucketObjectLockConfigurationV2("my-bucket-object-lock-configuration",
    bucket="my-bucket-36224917",
    object_lock_enabled="Enabled",
    opts = pulumi.ResourceOptions(protect=True))
my_bucket_versioning = aws.s3.BucketVersioningV2("my-bucket-versioning",
    bucket="my-bucket-36224917",
    versioning_configuration={
        "mfa_delete": "Disabled",
        "status": "Enabled",
    },
    opts = pulumi.ResourceOptions(protect=True))

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
		_, err := s3.NewBucketV2(ctx, "my-bucket", &s3.BucketV2Args{
			Bucket: pulumi.String("my-bucket-36224917"),
			Grants: s3.BucketV2GrantArray{
				&s3.BucketV2GrantArgs{
					Id: pulumi.String("e07865a5679c7977370948f1f1e51c21b12d8cfdd396a7e3172275d9164e01b8"),
					Permissions: pulumi.StringArray{
						pulumi.String("FULL_CONTROL"),
					},
					Type: pulumi.String("CanonicalUser"),
				},
			},
			LifecycleRules: s3.BucketV2LifecycleRuleArray{
				&s3.BucketV2LifecycleRuleArgs{
					Enabled: pulumi.Bool(true),
					Expirations: s3.BucketV2LifecycleRuleExpirationArray{
						&s3.BucketV2LifecycleRuleExpirationArgs{
							Days: pulumi.Int(30),
						},
					},
					Id: pulumi.String("pu-s3-lifecycle-20240918194815251100000001"),
				},
			},
			ObjectLockConfiguration: &s3.BucketV2ObjectLockConfigurationArgs{
				ObjectLockEnabled: pulumi.String("Enabled"),
			},
			Policy:       pulumi.String("{\"Id\":\"PutObjPolicy\",\"Statement\":[{\"Action\":\"s3:PutObject\",\"Condition\":{\"Null\":{\"s3:x-amz-server-side-encryption-aws-kms-key-id\":\"true\"}},\"Effect\":\"Deny\",\"Principal\":\"*\",\"Resource\":\"arn:aws:s3:::my-bucket-36224917/*\",\"Sid\":\"DenyObjectsThatAreNotSSEKMS\"}],\"Version\":\"2012-10-17\"}"),
			RequestPayer: pulumi.String("BucketOwner"),
			ServerSideEncryptionConfigurations: s3.BucketV2ServerSideEncryptionConfigurationArray{
				&s3.BucketV2ServerSideEncryptionConfigurationArgs{
					Rules: s3.BucketV2ServerSideEncryptionConfigurationRuleArray{
						&s3.BucketV2ServerSideEncryptionConfigurationRuleArgs{
							ApplyServerSideEncryptionByDefaults: s3.BucketV2ServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultArray{
								&s3.BucketV2ServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultArgs{
									SseAlgorithm: pulumi.String("AES256"),
								},
							},
						},
					},
				},
			},
			Tags: pulumi.StringMap{
				"Environment": pulumi.String("Dev"),
			},
			Versionings: s3.BucketV2VersioningArray{
				&s3.BucketV2VersioningArgs{
					Enabled: pulumi.Bool(true),
				},
			},
		}, pulumi.Protect(true))
		if err != nil {
			return err
		}
		_, err = s3.NewBucketServerSideEncryptionConfigurationV2(ctx, "my-bucket-encryption-configuration", &s3.BucketServerSideEncryptionConfigurationV2Args{
			Bucket: pulumi.String("my-bucket-36224917"),
			Rules: s3.BucketServerSideEncryptionConfigurationV2RuleArray{
				&s3.BucketServerSideEncryptionConfigurationV2RuleArgs{
					ApplyServerSideEncryptionByDefault: &s3.BucketServerSideEncryptionConfigurationV2RuleApplyServerSideEncryptionByDefaultArgs{
						SseAlgorithm: pulumi.String("AES256"),
					},
				},
			},
		}, pulumi.Protect(true))
		if err != nil {
			return err
		}
		_, err = s3.NewBucketPolicy(ctx, "my-bucket-policy", &s3.BucketPolicyArgs{
			Bucket: pulumi.String("my-bucket-36224917"),
			Policy: pulumi.Any("{\"Id\":\"PutObjPolicy\",\"Statement\":[{\"Action\":\"s3:PutObject\",\"Condition\":{\"Null\":{\"s3:x-amz-server-side-encryption-aws-kms-key-id\":\"true\"}},\"Effect\":\"Deny\",\"Principal\":\"*\",\"Resource\":\"arn:aws:s3:::my-bucket-36224917/*\",\"Sid\":\"DenyObjectsThatAreNotSSEKMS\"}],\"Version\":\"2012-10-17\"}"),
		}, pulumi.Protect(true))
		if err != nil {
			return err
		}
		_, err = s3.NewBucketLifecycleConfigurationV2(ctx, "my-bucket-lifecycle", &s3.BucketLifecycleConfigurationV2Args{
			Bucket: pulumi.String("my-bucket-36224917"),
			Rules: s3.BucketLifecycleConfigurationV2RuleArray{
				&s3.BucketLifecycleConfigurationV2RuleArgs{
					Expiration: &s3.BucketLifecycleConfigurationV2RuleExpirationArgs{
						Days: pulumi.Int(30),
					},
					Id:     pulumi.String("pu-s3-lifecycle-20240918194815251100000001"),
					Status: pulumi.String("Enabled"),
				},
			},
		}, pulumi.Protect(true))
		if err != nil {
			return err
		}
		_, err = s3.NewBucketObjectLockConfigurationV2(ctx, "my-bucket-object-lock-configuration", &s3.BucketObjectLockConfigurationV2Args{
			Bucket:            pulumi.String("my-bucket-36224917"),
			ObjectLockEnabled: pulumi.String("Enabled"),
		}, pulumi.Protect(true))
		if err != nil {
			return err
		}
		_, err = s3.NewBucketVersioningV2(ctx, "my-bucket-versioning", &s3.BucketVersioningV2Args{
			Bucket: pulumi.String("my-bucket-36224917"),
			VersioningConfiguration: &s3.BucketVersioningV2VersioningConfigurationArgs{
				MfaDelete: pulumi.String("Disabled"),
				Status:    pulumi.String("Enabled"),
			},
		}, pulumi.Protect(true))
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
    var myBucket = new Aws.S3.BucketV2("my-bucket", new()
    {
        Bucket = "my-bucket-36224917",
        Grants = new[]
        {
            new Aws.S3.Inputs.BucketV2GrantArgs
            {
                Id = "e07865a5679c7977370948f1f1e51c21b12d8cfdd396a7e3172275d9164e01b8",
                Permissions = new[]
                {
                    "FULL_CONTROL",
                },
                Type = "CanonicalUser",
            },
        },
        LifecycleRules = new[]
        {
            new Aws.S3.Inputs.BucketV2LifecycleRuleArgs
            {
                Enabled = true,
                Expirations = new[]
                {
                    new Aws.S3.Inputs.BucketV2LifecycleRuleExpirationArgs
                    {
                        Days = 30,
                    },
                },
                Id = "pu-s3-lifecycle-20240918194815251100000001",
            },
        },
        ObjectLockConfiguration = new Aws.S3.Inputs.BucketV2ObjectLockConfigurationArgs
        {
            ObjectLockEnabled = "Enabled",
        },
        Policy = "{\"Id\":\"PutObjPolicy\",\"Statement\":[{\"Action\":\"s3:PutObject\",\"Condition\":{\"Null\":{\"s3:x-amz-server-side-encryption-aws-kms-key-id\":\"true\"}},\"Effect\":\"Deny\",\"Principal\":\"*\",\"Resource\":\"arn:aws:s3:::my-bucket-36224917/*\",\"Sid\":\"DenyObjectsThatAreNotSSEKMS\"}],\"Version\":\"2012-10-17\"}",
        RequestPayer = "BucketOwner",
        ServerSideEncryptionConfigurations = new[]
        {
            new Aws.S3.Inputs.BucketV2ServerSideEncryptionConfigurationArgs
            {
                Rules = new[]
                {
                    new Aws.S3.Inputs.BucketV2ServerSideEncryptionConfigurationRuleArgs
                    {
                        ApplyServerSideEncryptionByDefaults = new[]
                        {
                            new Aws.S3.Inputs.BucketV2ServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultArgs
                            {
                                SseAlgorithm = "AES256",
                            },
                        },
                    },
                },
            },
        },
        Tags = 
        {
            { "Environment", "Dev" },
        },
        Versionings = new[]
        {
            new Aws.S3.Inputs.BucketV2VersioningArgs
            {
                Enabled = true,
            },
        },
    }, new CustomResourceOptions
    {
        Protect = true,
    });

    var myBucketEncryptionConfiguration = new Aws.S3.BucketServerSideEncryptionConfigurationV2("my-bucket-encryption-configuration", new()
    {
        Bucket = "my-bucket-36224917",
        Rules = new[]
        {
            new Aws.S3.Inputs.BucketServerSideEncryptionConfigurationV2RuleArgs
            {
                ApplyServerSideEncryptionByDefault = new Aws.S3.Inputs.BucketServerSideEncryptionConfigurationV2RuleApplyServerSideEncryptionByDefaultArgs
                {
                    SseAlgorithm = "AES256",
                },
            },
        },
    }, new CustomResourceOptions
    {
        Protect = true,
    });

    var myBucketPolicy = new Aws.S3.BucketPolicy("my-bucket-policy", new()
    {
        Bucket = "my-bucket-36224917",
        Policy = "{\"Id\":\"PutObjPolicy\",\"Statement\":[{\"Action\":\"s3:PutObject\",\"Condition\":{\"Null\":{\"s3:x-amz-server-side-encryption-aws-kms-key-id\":\"true\"}},\"Effect\":\"Deny\",\"Principal\":\"*\",\"Resource\":\"arn:aws:s3:::my-bucket-36224917/*\",\"Sid\":\"DenyObjectsThatAreNotSSEKMS\"}],\"Version\":\"2012-10-17\"}",
    }, new CustomResourceOptions
    {
        Protect = true,
    });

    var myBucketLifecycle = new Aws.S3.BucketLifecycleConfigurationV2("my-bucket-lifecycle", new()
    {
        Bucket = "my-bucket-36224917",
        Rules = new[]
        {
            new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleArgs
            {
                Expiration = new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleExpirationArgs
                {
                    Days = 30,
                },
                Id = "pu-s3-lifecycle-20240918194815251100000001",
                Status = "Enabled",
            },
        },
    }, new CustomResourceOptions
    {
        Protect = true,
    });

    var myBucketObjectLockConfiguration = new Aws.S3.BucketObjectLockConfigurationV2("my-bucket-object-lock-configuration", new()
    {
        Bucket = "my-bucket-36224917",
        ObjectLockEnabled = "Enabled",
    }, new CustomResourceOptions
    {
        Protect = true,
    });

    var myBucketVersioning = new Aws.S3.BucketVersioningV2("my-bucket-versioning", new()
    {
        Bucket = "my-bucket-36224917",
        VersioningConfiguration = new Aws.S3.Inputs.BucketVersioningV2VersioningConfigurationArgs
        {
            MfaDelete = "Disabled",
            Status = "Enabled",
        },
    }, new CustomResourceOptions
    {
        Protect = true,
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
import com.pulumi.aws.s3.BucketV2Args;
import com.pulumi.aws.s3.inputs.BucketV2GrantArgs;
import com.pulumi.aws.s3.inputs.BucketV2LifecycleRuleArgs;
import com.pulumi.aws.s3.inputs.BucketV2ObjectLockConfigurationArgs;
import com.pulumi.aws.s3.inputs.BucketV2ServerSideEncryptionConfigurationArgs;
import com.pulumi.aws.s3.inputs.BucketV2VersioningArgs;
import com.pulumi.aws.s3.BucketServerSideEncryptionConfigurationV2;
import com.pulumi.aws.s3.BucketServerSideEncryptionConfigurationV2Args;
import com.pulumi.aws.s3.inputs.BucketServerSideEncryptionConfigurationV2RuleArgs;
import com.pulumi.aws.s3.inputs.BucketServerSideEncryptionConfigurationV2RuleApplyServerSideEncryptionByDefaultArgs;
import com.pulumi.aws.s3.BucketPolicy;
import com.pulumi.aws.s3.BucketPolicyArgs;
import com.pulumi.aws.s3.BucketLifecycleConfigurationV2;
import com.pulumi.aws.s3.BucketLifecycleConfigurationV2Args;
import com.pulumi.aws.s3.inputs.BucketLifecycleConfigurationV2RuleArgs;
import com.pulumi.aws.s3.inputs.BucketLifecycleConfigurationV2RuleExpirationArgs;
import com.pulumi.aws.s3.BucketObjectLockConfigurationV2;
import com.pulumi.aws.s3.BucketObjectLockConfigurationV2Args;
import com.pulumi.aws.s3.BucketVersioningV2;
import com.pulumi.aws.s3.BucketVersioningV2Args;
import com.pulumi.aws.s3.inputs.BucketVersioningV2VersioningConfigurationArgs;
import com.pulumi.resources.CustomResourceOptions;
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
        var myBucket = new BucketV2("myBucket", BucketV2Args.builder()
            .bucket("my-bucket-36224917")
            .grants(BucketV2GrantArgs.builder()
                .id("e07865a5679c7977370948f1f1e51c21b12d8cfdd396a7e3172275d9164e01b8")
                .permissions("FULL_CONTROL")
                .type("CanonicalUser")
                .build())
            .lifecycleRules(BucketV2LifecycleRuleArgs.builder()
                .enabled(true)
                .expirations(BucketV2LifecycleRuleExpirationArgs.builder()
                    .days(30)
                    .build())
                .id("pu-s3-lifecycle-20240918194815251100000001")
                .build())
            .objectLockConfiguration(BucketV2ObjectLockConfigurationArgs.builder()
                .objectLockEnabled("Enabled")
                .build())
            .policy("{\"Id\":\"PutObjPolicy\",\"Statement\":[{\"Action\":\"s3:PutObject\",\"Condition\":{\"Null\":{\"s3:x-amz-server-side-encryption-aws-kms-key-id\":\"true\"}},\"Effect\":\"Deny\",\"Principal\":\"*\",\"Resource\":\"arn:aws:s3:::my-bucket-36224917/*\",\"Sid\":\"DenyObjectsThatAreNotSSEKMS\"}],\"Version\":\"2012-10-17\"}")
            .requestPayer("BucketOwner")
            .serverSideEncryptionConfigurations(BucketV2ServerSideEncryptionConfigurationArgs.builder()
                .rules(BucketV2ServerSideEncryptionConfigurationRuleArgs.builder()
                    .applyServerSideEncryptionByDefaults(BucketV2ServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultArgs.builder()
                        .sseAlgorithm("AES256")
                        .build())
                    .build())
                .build())
            .tags(Map.of("Environment", "Dev"))
            .versionings(BucketV2VersioningArgs.builder()
                .enabled(true)
                .build())
            .build(), CustomResourceOptions.builder()
                .protect(true)
                .build());

        var myBucketEncryptionConfiguration = new BucketServerSideEncryptionConfigurationV2("myBucketEncryptionConfiguration", BucketServerSideEncryptionConfigurationV2Args.builder()
            .bucket("my-bucket-36224917")
            .rules(BucketServerSideEncryptionConfigurationV2RuleArgs.builder()
                .applyServerSideEncryptionByDefault(BucketServerSideEncryptionConfigurationV2RuleApplyServerSideEncryptionByDefaultArgs.builder()
                    .sseAlgorithm("AES256")
                    .build())
                .build())
            .build(), CustomResourceOptions.builder()
                .protect(true)
                .build());

        var myBucketPolicy = new BucketPolicy("myBucketPolicy", BucketPolicyArgs.builder()
            .bucket("my-bucket-36224917")
            .policy("{\"Id\":\"PutObjPolicy\",\"Statement\":[{\"Action\":\"s3:PutObject\",\"Condition\":{\"Null\":{\"s3:x-amz-server-side-encryption-aws-kms-key-id\":\"true\"}},\"Effect\":\"Deny\",\"Principal\":\"*\",\"Resource\":\"arn:aws:s3:::my-bucket-36224917/*\",\"Sid\":\"DenyObjectsThatAreNotSSEKMS\"}],\"Version\":\"2012-10-17\"}")
            .build(), CustomResourceOptions.builder()
                .protect(true)
                .build());

        var myBucketLifecycle = new BucketLifecycleConfigurationV2("myBucketLifecycle", BucketLifecycleConfigurationV2Args.builder()
            .bucket("my-bucket-36224917")
            .rules(BucketLifecycleConfigurationV2RuleArgs.builder()
                .expiration(BucketLifecycleConfigurationV2RuleExpirationArgs.builder()
                    .days(30)
                    .build())
                .id("pu-s3-lifecycle-20240918194815251100000001")
                .status("Enabled")
                .build())
            .build(), CustomResourceOptions.builder()
                .protect(true)
                .build());

        var myBucketObjectLockConfiguration = new BucketObjectLockConfigurationV2("myBucketObjectLockConfiguration", BucketObjectLockConfigurationV2Args.builder()
            .bucket("my-bucket-36224917")
            .objectLockEnabled("Enabled")
            .build(), CustomResourceOptions.builder()
                .protect(true)
                .build());

        var myBucketVersioning = new BucketVersioningV2("myBucketVersioning", BucketVersioningV2Args.builder()
            .bucket("my-bucket-36224917")
            .versioningConfiguration(BucketVersioningV2VersioningConfigurationArgs.builder()
                .mfaDelete("Disabled")
                .status("Enabled")
                .build())
            .build(), CustomResourceOptions.builder()
                .protect(true)
                .build());

    }
}

```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
name: y2
runtime: yaml
resources:
  my-bucket:
    type: aws:s3:BucketV2
    properties:
      bucket: my-bucket-36224917
      grants:
        - id: e07865a5679c7977370948f1f1e51c21b12d8cfdd396a7e3172275d9164e01b8
          permissions:
            - FULL_CONTROL
          type: CanonicalUser
      lifecycleRules:
        - enabled: true
          expirations:
            - days: 30
          id: pu-s3-lifecycle-20240918194815251100000001
      objectLockConfiguration:
        objectLockEnabled: Enabled
      policy: '{"Id":"PutObjPolicy","Statement":[{"Action":"s3:PutObject","Condition":{"Null":{"s3:x-amz-server-side-encryption-aws-kms-key-id":"true"}},"Effect":"Deny","Principal":"*","Resource":"arn:aws:s3:::my-bucket-36224917/*","Sid":"DenyObjectsThatAreNotSSEKMS"}],"Version":"2012-10-17"}'
      requestPayer: BucketOwner
      serverSideEncryptionConfigurations:
        - rules:
            - applyServerSideEncryptionByDefaults:
                - sseAlgorithm: AES256
      tags:
        Environment: Dev
      versionings:
        - enabled: true
    options:
      protect: true
  my-bucket-encryption-configuration:
    type: aws:s3:BucketServerSideEncryptionConfigurationV2
    properties:
      bucket: my-bucket-36224917
      rules:
        - applyServerSideEncryptionByDefault:
            sseAlgorithm: AES256
    options:
      protect: true
  my-bucket-policy:
    type: aws:s3:BucketPolicy
    properties:
      bucket: my-bucket-36224917
      policy: '{"Id":"PutObjPolicy","Statement":[{"Action":"s3:PutObject","Condition":{"Null":{"s3:x-amz-server-side-encryption-aws-kms-key-id":"true"}},"Effect":"Deny","Principal":"*","Resource":"arn:aws:s3:::my-bucket-36224917/*","Sid":"DenyObjectsThatAreNotSSEKMS"}],"Version":"2012-10-17"}'
    options:
      protect: true
  my-bucket-lifecycle:
    type: aws:s3:BucketLifecycleConfigurationV2
    properties:
      bucket: my-bucket-36224917
      rules:
        - expiration:
            days: 30
          id: pu-s3-lifecycle-20240918194815251100000001
          status: Enabled
    options:
      protect: true
  my-bucket-object-lock-configuration:
    type: aws:s3:BucketObjectLockConfigurationV2
    properties:
      bucket: my-bucket-36224917
      objectLockEnabled: Enabled
    options:
      protect: true
  my-bucket-versioning:
    type: aws:s3:BucketVersioningV2
    properties:
      bucket: my-bucket-36224917
      versioningConfiguration:
        mfaDelete: Disabled
        status: Enabled
    options:
      protect: true
```

{{% /choosable %}}

{{< /chooser >}}

- `pulumi preview` should result in `Resources: N unchanged` to confirm everything went well.

- At this point you may see warnings from deprecated inputs like this:

```shell
warning: urn:pulumi:bucket1::y2::aws:s3/bucketV2:BucketV2::my-bucket verification warning: Use the top-level parameter object_lock_enabled instead
warning: urn:pulumi:bucket1::y2::aws:s3/bucketV2:BucketV2::my-bucket verification warning: Use the top-level parameter object_lock_enabled and the aws_s3_bucket_object_lock_configuration resource instead
warning: urn:pulumi:bucket1::y2::aws:s3/bucketV2:BucketV2::my-bucket verification warning: Use the aws_s3_bucket_lifecycle_configuration resource instead
warning: urn:pulumi:bucket1::y2::aws:s3/bucketV2:BucketV2::my-bucket verification warning: Use the aws_s3_bucket_versioning resource instead
warning: urn:pulumi:bucket1::y2::aws:s3/bucketV2:BucketV2::my-bucket verification warning: Use the aws_s3_bucket_acl resource instead
warning: urn:pulumi:bucket1::y2::aws:s3/bucketV2:BucketV2::my-bucket verification warning: Use the aws_s3_bucket_request_payment_configuration resource instead
warning: urn:pulumi:bucket1::y2::aws:s3/bucketV2:BucketV2::my-bucket verification warning: Use the aws_s3_bucket_policy resource instead
warning: urn:pulumi:bucket1::y2::aws:s3/bucketV2:BucketV2::my-bucket verification warning: Use the aws_s3_bucket_server_side_encryption_configuration resource instead
```

- To mitigate, edit the program to remove the deprecated inputs, leaving the final simplify program that should still
  result in an empty `pulumi preview`:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}

{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const myBucket = new aws.s3.BucketV2("my-bucket", {
    bucket: "my-bucket-36224917",
    tags: {
        Environment: "Dev",
    },
}, {
    protect: true,
});
const myBucketEncryptionConfiguration = new aws.s3.BucketServerSideEncryptionConfigurationV2("my-bucket-encryption-configuration", {
    bucket: "my-bucket-36224917",
    rules: [{
        applyServerSideEncryptionByDefault: {
            sseAlgorithm: "AES256",
        },
    }],
}, {
    protect: true,
});
const myBucketPolicy = new aws.s3.BucketPolicy("my-bucket-policy", {
    bucket: "my-bucket-36224917",
    policy: "{\"Id\":\"PutObjPolicy\",\"Statement\":[{\"Action\":\"s3:PutObject\",\"Condition\":{\"Null\":{\"s3:x-amz-server-side-encryption-aws-kms-key-id\":\"true\"}},\"Effect\":\"Deny\",\"Principal\":\"*\",\"Resource\":\"arn:aws:s3:::my-bucket-36224917/*\",\"Sid\":\"DenyObjectsThatAreNotSSEKMS\"}],\"Version\":\"2012-10-17\"}",
}, {
    protect: true,
});
const myBucketLifecycle = new aws.s3.BucketLifecycleConfigurationV2("my-bucket-lifecycle", {
    bucket: "my-bucket-36224917",
    rules: [{
        expiration: {
            days: 30,
        },
        id: "pu-s3-lifecycle-20240918194815251100000001",
        status: "Enabled",
    }],
}, {
    protect: true,
});
const myBucketObjectLockConfiguration = new aws.s3.BucketObjectLockConfigurationV2("my-bucket-object-lock-configuration", {
    bucket: "my-bucket-36224917",
    objectLockEnabled: "Enabled",
}, {
    protect: true,
});
const myBucketVersioning = new aws.s3.BucketVersioningV2("my-bucket-versioning", {
    bucket: "my-bucket-36224917",
    versioningConfiguration: {
        mfaDelete: "Disabled",
        status: "Enabled",
    },
}, {
    protect: true,
});

```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi
import pulumi_aws as aws

my_bucket = aws.s3.BucketV2("my-bucket",
    bucket="my-bucket-36224917",
    tags={
        "Environment": "Dev",
    },
    opts = pulumi.ResourceOptions(protect=True))
my_bucket_encryption_configuration = aws.s3.BucketServerSideEncryptionConfigurationV2("my-bucket-encryption-configuration",
    bucket="my-bucket-36224917",
    rules=[{
        "apply_server_side_encryption_by_default": {
            "sse_algorithm": "AES256",
        },
    }],
    opts = pulumi.ResourceOptions(protect=True))
my_bucket_policy = aws.s3.BucketPolicy("my-bucket-policy",
    bucket="my-bucket-36224917",
    policy="{\"Id\":\"PutObjPolicy\",\"Statement\":[{\"Action\":\"s3:PutObject\",\"Condition\":{\"Null\":{\"s3:x-amz-server-side-encryption-aws-kms-key-id\":\"true\"}},\"Effect\":\"Deny\",\"Principal\":\"*\",\"Resource\":\"arn:aws:s3:::my-bucket-36224917/*\",\"Sid\":\"DenyObjectsThatAreNotSSEKMS\"}],\"Version\":\"2012-10-17\"}",
    opts = pulumi.ResourceOptions(protect=True))
my_bucket_lifecycle = aws.s3.BucketLifecycleConfigurationV2("my-bucket-lifecycle",
    bucket="my-bucket-36224917",
    rules=[{
        "expiration": {
            "days": 30,
        },
        "id": "pu-s3-lifecycle-20240918194815251100000001",
        "status": "Enabled",
    }],
    opts = pulumi.ResourceOptions(protect=True))
my_bucket_object_lock_configuration = aws.s3.BucketObjectLockConfigurationV2("my-bucket-object-lock-configuration",
    bucket="my-bucket-36224917",
    object_lock_enabled="Enabled",
    opts = pulumi.ResourceOptions(protect=True))
my_bucket_versioning = aws.s3.BucketVersioningV2("my-bucket-versioning",
    bucket="my-bucket-36224917",
    versioning_configuration={
        "mfa_delete": "Disabled",
        "status": "Enabled",
    },
    opts = pulumi.ResourceOptions(protect=True))

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
		_, err := s3.NewBucketV2(ctx, "my-bucket", &s3.BucketV2Args{
			Bucket: pulumi.String("my-bucket-36224917"),
			Tags: pulumi.StringMap{
				"Environment": pulumi.String("Dev"),
			},
		}, pulumi.Protect(true))
		if err != nil {
			return err
		}
		_, err = s3.NewBucketServerSideEncryptionConfigurationV2(ctx, "my-bucket-encryption-configuration", &s3.BucketServerSideEncryptionConfigurationV2Args{
			Bucket: pulumi.String("my-bucket-36224917"),
			Rules: s3.BucketServerSideEncryptionConfigurationV2RuleArray{
				&s3.BucketServerSideEncryptionConfigurationV2RuleArgs{
					ApplyServerSideEncryptionByDefault: &s3.BucketServerSideEncryptionConfigurationV2RuleApplyServerSideEncryptionByDefaultArgs{
						SseAlgorithm: pulumi.String("AES256"),
					},
				},
			},
		}, pulumi.Protect(true))
		if err != nil {
			return err
		}
		_, err = s3.NewBucketPolicy(ctx, "my-bucket-policy", &s3.BucketPolicyArgs{
			Bucket: pulumi.String("my-bucket-36224917"),
			Policy: pulumi.Any("{\"Id\":\"PutObjPolicy\",\"Statement\":[{\"Action\":\"s3:PutObject\",\"Condition\":{\"Null\":{\"s3:x-amz-server-side-encryption-aws-kms-key-id\":\"true\"}},\"Effect\":\"Deny\",\"Principal\":\"*\",\"Resource\":\"arn:aws:s3:::my-bucket-36224917/*\",\"Sid\":\"DenyObjectsThatAreNotSSEKMS\"}],\"Version\":\"2012-10-17\"}"),
		}, pulumi.Protect(true))
		if err != nil {
			return err
		}
		_, err = s3.NewBucketLifecycleConfigurationV2(ctx, "my-bucket-lifecycle", &s3.BucketLifecycleConfigurationV2Args{
			Bucket: pulumi.String("my-bucket-36224917"),
			Rules: s3.BucketLifecycleConfigurationV2RuleArray{
				&s3.BucketLifecycleConfigurationV2RuleArgs{
					Expiration: &s3.BucketLifecycleConfigurationV2RuleExpirationArgs{
						Days: pulumi.Int(30),
					},
					Id:     pulumi.String("pu-s3-lifecycle-20240918194815251100000001"),
					Status: pulumi.String("Enabled"),
				},
			},
		}, pulumi.Protect(true))
		if err != nil {
			return err
		}
		_, err = s3.NewBucketObjectLockConfigurationV2(ctx, "my-bucket-object-lock-configuration", &s3.BucketObjectLockConfigurationV2Args{
			Bucket:            pulumi.String("my-bucket-36224917"),
			ObjectLockEnabled: pulumi.String("Enabled"),
		}, pulumi.Protect(true))
		if err != nil {
			return err
		}
		_, err = s3.NewBucketVersioningV2(ctx, "my-bucket-versioning", &s3.BucketVersioningV2Args{
			Bucket: pulumi.String("my-bucket-36224917"),
			VersioningConfiguration: &s3.BucketVersioningV2VersioningConfigurationArgs{
				MfaDelete: pulumi.String("Disabled"),
				Status:    pulumi.String("Enabled"),
			},
		}, pulumi.Protect(true))
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
    var myBucket = new Aws.S3.BucketV2("my-bucket", new()
    {
        Bucket = "my-bucket-36224917",
        Tags = 
        {
            { "Environment", "Dev" },
        },
    }, new CustomResourceOptions
    {
        Protect = true,
    });

    var myBucketEncryptionConfiguration = new Aws.S3.BucketServerSideEncryptionConfigurationV2("my-bucket-encryption-configuration", new()
    {
        Bucket = "my-bucket-36224917",
        Rules = new[]
        {
            new Aws.S3.Inputs.BucketServerSideEncryptionConfigurationV2RuleArgs
            {
                ApplyServerSideEncryptionByDefault = new Aws.S3.Inputs.BucketServerSideEncryptionConfigurationV2RuleApplyServerSideEncryptionByDefaultArgs
                {
                    SseAlgorithm = "AES256",
                },
            },
        },
    }, new CustomResourceOptions
    {
        Protect = true,
    });

    var myBucketPolicy = new Aws.S3.BucketPolicy("my-bucket-policy", new()
    {
        Bucket = "my-bucket-36224917",
        Policy = "{\"Id\":\"PutObjPolicy\",\"Statement\":[{\"Action\":\"s3:PutObject\",\"Condition\":{\"Null\":{\"s3:x-amz-server-side-encryption-aws-kms-key-id\":\"true\"}},\"Effect\":\"Deny\",\"Principal\":\"*\",\"Resource\":\"arn:aws:s3:::my-bucket-36224917/*\",\"Sid\":\"DenyObjectsThatAreNotSSEKMS\"}],\"Version\":\"2012-10-17\"}",
    }, new CustomResourceOptions
    {
        Protect = true,
    });

    var myBucketLifecycle = new Aws.S3.BucketLifecycleConfigurationV2("my-bucket-lifecycle", new()
    {
        Bucket = "my-bucket-36224917",
        Rules = new[]
        {
            new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleArgs
            {
                Expiration = new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleExpirationArgs
                {
                    Days = 30,
                },
                Id = "pu-s3-lifecycle-20240918194815251100000001",
                Status = "Enabled",
            },
        },
    }, new CustomResourceOptions
    {
        Protect = true,
    });

    var myBucketObjectLockConfiguration = new Aws.S3.BucketObjectLockConfigurationV2("my-bucket-object-lock-configuration", new()
    {
        Bucket = "my-bucket-36224917",
        ObjectLockEnabled = "Enabled",
    }, new CustomResourceOptions
    {
        Protect = true,
    });

    var myBucketVersioning = new Aws.S3.BucketVersioningV2("my-bucket-versioning", new()
    {
        Bucket = "my-bucket-36224917",
        VersioningConfiguration = new Aws.S3.Inputs.BucketVersioningV2VersioningConfigurationArgs
        {
            MfaDelete = "Disabled",
            Status = "Enabled",
        },
    }, new CustomResourceOptions
    {
        Protect = true,
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
import com.pulumi.aws.s3.BucketV2Args;
import com.pulumi.aws.s3.BucketServerSideEncryptionConfigurationV2;
import com.pulumi.aws.s3.BucketServerSideEncryptionConfigurationV2Args;
import com.pulumi.aws.s3.inputs.BucketServerSideEncryptionConfigurationV2RuleArgs;
import com.pulumi.aws.s3.inputs.BucketServerSideEncryptionConfigurationV2RuleApplyServerSideEncryptionByDefaultArgs;
import com.pulumi.aws.s3.BucketPolicy;
import com.pulumi.aws.s3.BucketPolicyArgs;
import com.pulumi.aws.s3.BucketLifecycleConfigurationV2;
import com.pulumi.aws.s3.BucketLifecycleConfigurationV2Args;
import com.pulumi.aws.s3.inputs.BucketLifecycleConfigurationV2RuleArgs;
import com.pulumi.aws.s3.inputs.BucketLifecycleConfigurationV2RuleExpirationArgs;
import com.pulumi.aws.s3.BucketObjectLockConfigurationV2;
import com.pulumi.aws.s3.BucketObjectLockConfigurationV2Args;
import com.pulumi.aws.s3.BucketVersioningV2;
import com.pulumi.aws.s3.BucketVersioningV2Args;
import com.pulumi.aws.s3.inputs.BucketVersioningV2VersioningConfigurationArgs;
import com.pulumi.resources.CustomResourceOptions;
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
        var myBucket = new BucketV2("myBucket", BucketV2Args.builder()
            .bucket("my-bucket-36224917")
            .tags(Map.of("Environment", "Dev"))
            .build(), CustomResourceOptions.builder()
                .protect(true)
                .build());

        var myBucketEncryptionConfiguration = new BucketServerSideEncryptionConfigurationV2("myBucketEncryptionConfiguration", BucketServerSideEncryptionConfigurationV2Args.builder()
            .bucket("my-bucket-36224917")
            .rules(BucketServerSideEncryptionConfigurationV2RuleArgs.builder()
                .applyServerSideEncryptionByDefault(BucketServerSideEncryptionConfigurationV2RuleApplyServerSideEncryptionByDefaultArgs.builder()
                    .sseAlgorithm("AES256")
                    .build())
                .build())
            .build(), CustomResourceOptions.builder()
                .protect(true)
                .build());

        var myBucketPolicy = new BucketPolicy("myBucketPolicy", BucketPolicyArgs.builder()
            .bucket("my-bucket-36224917")
            .policy("{\"Id\":\"PutObjPolicy\",\"Statement\":[{\"Action\":\"s3:PutObject\",\"Condition\":{\"Null\":{\"s3:x-amz-server-side-encryption-aws-kms-key-id\":\"true\"}},\"Effect\":\"Deny\",\"Principal\":\"*\",\"Resource\":\"arn:aws:s3:::my-bucket-36224917/*\",\"Sid\":\"DenyObjectsThatAreNotSSEKMS\"}],\"Version\":\"2012-10-17\"}")
            .build(), CustomResourceOptions.builder()
                .protect(true)
                .build());

        var myBucketLifecycle = new BucketLifecycleConfigurationV2("myBucketLifecycle", BucketLifecycleConfigurationV2Args.builder()
            .bucket("my-bucket-36224917")
            .rules(BucketLifecycleConfigurationV2RuleArgs.builder()
                .expiration(BucketLifecycleConfigurationV2RuleExpirationArgs.builder()
                    .days(30)
                    .build())
                .id("pu-s3-lifecycle-20240918194815251100000001")
                .status("Enabled")
                .build())
            .build(), CustomResourceOptions.builder()
                .protect(true)
                .build());

        var myBucketObjectLockConfiguration = new BucketObjectLockConfigurationV2("myBucketObjectLockConfiguration", BucketObjectLockConfigurationV2Args.builder()
            .bucket("my-bucket-36224917")
            .objectLockEnabled("Enabled")
            .build(), CustomResourceOptions.builder()
                .protect(true)
                .build());

        var myBucketVersioning = new BucketVersioningV2("myBucketVersioning", BucketVersioningV2Args.builder()
            .bucket("my-bucket-36224917")
            .versioningConfiguration(BucketVersioningV2VersioningConfigurationArgs.builder()
                .mfaDelete("Disabled")
                .status("Enabled")
                .build())
            .build(), CustomResourceOptions.builder()
                .protect(true)
                .build());

    }
}

```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
name: y2
runtime: yaml
resources:
  my-bucket:
    type: aws:s3:BucketV2
    properties:
      bucket: my-bucket-36224917
      tags:
        Environment: Dev
    options:
      protect: true
  my-bucket-encryption-configuration:
    type: aws:s3:BucketServerSideEncryptionConfigurationV2
    properties:
      bucket: my-bucket-36224917
      rules:
        - applyServerSideEncryptionByDefault:
            sseAlgorithm: AES256
    options:
      protect: true
  my-bucket-policy:
    type: aws:s3:BucketPolicy
    properties:
      bucket: my-bucket-36224917
      policy: '{"Id":"PutObjPolicy","Statement":[{"Action":"s3:PutObject","Condition":{"Null":{"s3:x-amz-server-side-encryption-aws-kms-key-id":"true"}},"Effect":"Deny","Principal":"*","Resource":"arn:aws:s3:::my-bucket-36224917/*","Sid":"DenyObjectsThatAreNotSSEKMS"}],"Version":"2012-10-17"}'
    options:
      protect: true
  my-bucket-lifecycle:
    type: aws:s3:BucketLifecycleConfigurationV2
    properties:
      bucket: my-bucket-36224917
      rules:
        - expiration:
            days: 30
          id: pu-s3-lifecycle-20240918194815251100000001
          status: Enabled
    options:
      protect: true
  my-bucket-object-lock-configuration:
    type: aws:s3:BucketObjectLockConfigurationV2
    properties:
      bucket: my-bucket-36224917
      objectLockEnabled: Enabled
    options:
      protect: true
  my-bucket-versioning:
    type: aws:s3:BucketVersioningV2
    properties:
      bucket: my-bucket-36224917
      versioningConfiguration:
        mfaDelete: Disabled
        status: Enabled
    options:
      protect: true
```

{{% /choosable %}}

{{< /chooser >}}

## Historical context

The distinction between Bucket and BucketV2 originates from breaking changes introduced in the V4 release of the
Terraform AWS Provider. In the Pulumi AWS provider projection BucketV2 represents the latest version of the upstream
resource, while Bucket is maintained by Pulumi to enable backwards compatibility.

See also:

- [Announcing v5.0.0 of the Pulumi AWS Provider](https://www.pulumi.com/blog/announcing-v5.0.0-of-the-pulumi-aws-provider/)
- [Terraform AWS Provider Version 4 Upgrade Guide](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/guides/version-4-upgrade#s3-bucket-refactor)
