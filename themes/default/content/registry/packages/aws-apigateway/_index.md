---
title: AWS API Gateway
meta_desc: Easily create AWS API Gateway REST APIs using Pulumi.
layout: package
---

Easily create AWS API Gateway REST APIs using Pulumi. This component provides higher-level abstractions documented in the [Pulumi AWS API Gateway guide](/docs/clouds/aws/guides/api-gateway/) as a package available in all Pulumi languages.

## Example:

The TypeScript example defines the Lambda handler inline with `aws.lambda.CallbackFunction`, which serializes the callback into the function's code. That only works in Node.js, so the other languages package a handler from the `./handler` directory instead.

{{< chooser language "typescript,python,csharp,go,java,yaml,hcl" >}}

{{% choosable language typescript %}}

```typescript
import * as apigateway from "@pulumi/aws-apigateway";
import * as aws from "@pulumi/aws";

const f = new aws.lambda.CallbackFunction("f", {
    callback: async (ev, ctx) => {
        console.log(JSON.stringify(ev));
        return {
            statusCode: 200,
            body: "goodbye",
        };
    },
});

const api = new apigateway.RestAPI("api", {
    routes: [{
        path: "/",
        method: "GET",
        eventHandler: f,
    }],
});

export const url = api.url;
```

{{% /choosable %}}

{{% choosable language python %}}

```python
import json
import pulumi
import pulumi_aws as aws
import pulumi_aws_apigateway as apigateway

role = aws.iam.Role("mylambda-role",
    assume_role_policy=json.dumps({
        "Version": "2012-10-17",
        "Statement": [{
            "Effect": "Allow",
            "Principal": {"Service": "lambda.amazonaws.com"},
            "Action": "sts:AssumeRole",
        }],
    }))

policy = aws.iam.RolePolicy("mylambda-policy",
    role=role.id,
    policy=json.dumps({
        "Version": "2012-10-17",
        "Statement": [{
            "Action": ["logs:*", "cloudwatch:*"],
            "Resource": "*",
            "Effect": "Allow",
        }],
    }))

# Closure serialization is not supported in multi-lang components
# so we need to provide a handler function explicitly from the file-system.
# Refer to https://github.com/pulumi/pulumi-aws-apigateway/tree/main/examples/simple-py/handler
# for an example handler.
f = aws.lambda_.Function("mylambda",
    runtime=aws.lambda_.Runtime.PYTHON3D12,
    code=pulumi.FileArchive("./handler"),
    timeout=300,
    handler="handler.handler",
    role=role.arn,
    opts=pulumi.ResourceOptions(depends_on=[policy]))

api = apigateway.RestAPI("api", routes=[{
    "path": "/",
    "method": apigateway.Method.GET,
    "event_handler": f,
}])

pulumi.export("url", api.url)
```

{{% /choosable %}}

{{% choosable language csharp %}}

```csharp
using System.Collections.Generic;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;
using ApiGateway = Pulumi.AwsApiGateway;

return await Deployment.RunAsync(() =>
{
    var role = new Aws.Iam.Role("mylambda-role", new()
    {
        AssumeRolePolicy = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["Version"] = "2012-10-17",
            ["Statement"] = new[]
            {
                new Dictionary<string, object?>
                {
                    ["Effect"] = "Allow",
                    ["Principal"] = new Dictionary<string, object?>
                    {
                        ["Service"] = "lambda.amazonaws.com",
                    },
                    ["Action"] = "sts:AssumeRole",
                },
            },
        }),
    });

    var policy = new Aws.Iam.RolePolicy("mylambda-policy", new()
    {
        Role = role.Id,
        Policy = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["Version"] = "2012-10-17",
            ["Statement"] = new[]
            {
                new Dictionary<string, object?>
                {
                    ["Action"] = new[] { "logs:*", "cloudwatch:*" },
                    ["Resource"] = "*",
                    ["Effect"] = "Allow",
                },
            },
        }),
    });

    // Closure serialization is not supported in multi-lang components
    // so we need to provide a handler function explicitly from the file-system.
    // Refer to https://github.com/pulumi/pulumi-aws-apigateway/tree/main/examples/simple-py/handler
    // for an example handler.
    var f = new Aws.Lambda.Function("mylambda", new()
    {
        Runtime = Aws.Lambda.Runtime.Python3d12,
        Code = new FileArchive("./handler"),
        Timeout = 300,
        Handler = "handler.handler",
        Role = role.Arn,
    }, new CustomResourceOptions
    {
        DependsOn = { policy },
    });

    var api = new ApiGateway.RestAPI("api", new()
    {
        Routes =
        {
            new ApiGateway.Inputs.RouteArgs
            {
                Path = "/",
                Method = ApiGateway.Method.GET,
                EventHandler = f,
            },
        },
    });

    return new Dictionary<string, object?>
    {
        ["url"] = api.Url,
    };
});
```

{{% /choosable %}}

{{% choosable language go %}}

```go
package main

import (
	apigateway "github.com/pulumi/pulumi-aws-apigateway/sdk/v3/go/apigateway"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws/lambda"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		role, err := iam.NewRole(ctx, "mylambda-role", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.String(`{
				"Version": "2012-10-17",
				"Statement": [{
					"Effect": "Allow",
					"Principal": { "Service": "lambda.amazonaws.com" },
					"Action": "sts:AssumeRole"
				}]
			}`),
		})
		if err != nil {
			return err
		}

		policy, err := iam.NewRolePolicy(ctx, "mylambda-policy", &iam.RolePolicyArgs{
			Role: role.ID(),
			Policy: pulumi.String(`{
				"Version": "2012-10-17",
				"Statement": [{
					"Action": ["logs:*", "cloudwatch:*"],
					"Resource": "*",
					"Effect": "Allow"
				}]
			}`),
		})
		if err != nil {
			return err
		}

		// Closure serialization is not supported in multi-lang components
		// so we need to provide a handler function explicitly from the file-system.
		// Refer to https://github.com/pulumi/pulumi-aws-apigateway/tree/main/examples/simple-go/handler
		// for an example handler.
		f, err := lambda.NewFunction(ctx, "mylambda", &lambda.FunctionArgs{
			Runtime: pulumi.String(lambda.RuntimePython3d12),
			Code:    pulumi.NewFileArchive("./handler"),
			Timeout: pulumi.Int(300),
			Handler: pulumi.String("handler.handler"),
			Role:    role.Arn,
		}, pulumi.DependsOn([]pulumi.Resource{policy}))
		if err != nil {
			return err
		}

		// Method is optional on a route, so the Go SDK takes it as a pointer.
		getMethod := apigateway.MethodGET
		api, err := apigateway.NewRestAPI(ctx, "api", &apigateway.RestAPIArgs{
			Routes: []apigateway.RouteArgs{
				{
					Path:         "/",
					Method:       &getMethod,
					EventHandler: f,
				},
			},
		})
		if err != nil {
			return err
		}

		ctx.Export("url", api.Url)
		return nil
	})
}
```

{{% /choosable %}}

{{% choosable language java %}}

```java
import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.asset.FileArchive;
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.iam.RolePolicy;
import com.pulumi.aws.iam.RolePolicyArgs;
import com.pulumi.aws.lambda.Function;
import com.pulumi.aws.lambda.FunctionArgs;
import com.pulumi.awsapigateway.RestAPI;
import com.pulumi.awsapigateway.RestAPIArgs;
import com.pulumi.awsapigateway.enums.Method;
import com.pulumi.awsapigateway.inputs.RouteArgs;
import com.pulumi.resources.CustomResourceOptions;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        var role = new Role("mylambda-role", RoleArgs.builder()
            .assumeRolePolicy("{"
                + "\"Version\": \"2012-10-17\","
                + "\"Statement\": [{"
                + "\"Effect\": \"Allow\","
                + "\"Principal\": { \"Service\": \"lambda.amazonaws.com\" },"
                + "\"Action\": \"sts:AssumeRole\""
                + "}]"
                + "}")
            .build());

        var policy = new RolePolicy("mylambda-policy", RolePolicyArgs.builder()
            .role(role.id())
            .policy("{"
                + "\"Version\": \"2012-10-17\","
                + "\"Statement\": [{"
                + "\"Action\": [\"logs:*\", \"cloudwatch:*\"],"
                + "\"Resource\": \"*\","
                + "\"Effect\": \"Allow\""
                + "}]"
                + "}")
            .build());

        // Closure serialization is not supported in multi-lang components
        // so we need to provide a handler function explicitly from the file-system.
        // Refer to https://github.com/pulumi/pulumi-aws-apigateway/tree/main/examples/simple-py/handler
        // for an example handler.
        var f = new Function("mylambda", FunctionArgs.builder()
            .runtime("python3.12")
            .code(new FileArchive("./handler"))
            .timeout(300)
            .handler("handler.handler")
            .role(role.arn())
            .build(), CustomResourceOptions.builder()
                .dependsOn(policy)
                .build());

        var api = new RestAPI("api", RestAPIArgs.builder()
            .routes(RouteArgs.builder()
                .path("/")
                .method(Method.GET)
                .eventHandler(f)
                .build())
            .build());

        ctx.export("url", api.url());
    }
}
```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
resources:
  mylambda-role:
    type: aws:iam:Role
    properties:
      assumeRolePolicy:
        fn::toJSON:
          Version: 2012-10-17
          Statement:
            - Effect: Allow
              Principal:
                Service: lambda.amazonaws.com
              Action: sts:AssumeRole
  mylambda-policy:
    type: aws:iam:RolePolicy
    properties:
      role: ${mylambda-role.id}
      policy:
        fn::toJSON:
          Version: 2012-10-17
          Statement:
            - Action:
                - logs:*
                - cloudwatch:*
              Resource: "*"
              Effect: Allow
  # Closure serialization is not supported in multi-lang components
  # so we need to provide a handler function explicitly from the file-system.
  # Refer to https://github.com/pulumi/pulumi-aws-apigateway/tree/main/examples/simple-py/handler
  # for an example handler.
  mylambda:
    type: aws:lambda:Function
    properties:
      runtime: python3.12
      code:
        fn::fileArchive: ./handler
      timeout: 300
      handler: handler.handler
      role: ${mylambda-role.arn}
    options:
      dependsOn:
        - ${mylambda-policy}
  api:
    type: aws-apigateway:RestAPI
    properties:
      routes:
        - path: /
          method: GET
          eventHandler: ${mylambda}
outputs:
  url: ${api.url}
```

{{% /choosable %}}

{{% choosable language hcl %}}

```hcl
terraform {
  required_providers {
    aws = {
      source = "pulumi/aws"
    }
    aws-apigateway = {
      source = "pulumi/aws-apigateway"
    }
  }
}

resource "aws_iam_role" "mylambda-role" {
  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Effect = "Allow"
      Principal = {
        Service = "lambda.amazonaws.com"
      }
      Action = "sts:AssumeRole"
    }]
  })
}

resource "aws_iam_role_policy" "mylambda-policy" {
  role = aws_iam_role.mylambda-role.id
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Action   = ["logs:*", "cloudwatch:*"]
      Resource = "*"
      Effect   = "Allow"
    }]
  })
}

# Closure serialization is not supported in multi-lang components
# so we need to provide a handler function explicitly from the file-system.
# Refer to https://github.com/pulumi/pulumi-aws-apigateway/tree/main/examples/simple-py/handler
# for an example handler.
resource "aws_lambda_function" "mylambda" {
  runtime  = "python3.12"
  filename = filearchive("./handler")
  timeout  = 300
  handler  = "handler.handler"
  role     = aws_iam_role.mylambda-role.arn

  depends_on = [aws_iam_role_policy.mylambda-policy]
}

resource "aws-apigateway_rest_api" "api" {
  routes {
    path          = "/"
    method        = "GET"
    event_handler = aws_lambda_function.mylambda
  }
}

output "url" {
  value = aws-apigateway_rest_api.api.url
}
```

{{% /choosable %}}

{{< /chooser >}}
