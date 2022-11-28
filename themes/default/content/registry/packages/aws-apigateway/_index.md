---
title: AWS API Gateway
meta_desc: Easily create AWS API Gateway REST APIs using Pulumi.
layout: overview
---

Easily create AWS API Gateway REST APIs using Pulumi. This component exposes the [Crosswalk for AWS](../../../docs/guides/crosswalk/aws/) functionality documented in the [Pulumi AWS API Gateway guide](../../../docs/guides/crosswalk/aws/api-gateway/) as a package available in all Pulumi languages.

Example:

{{< chooser language "typescript,python,csharp,go" >}}

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
            "Principal": { "Service": "lambda.amazonaws.com" },
            "Action": "sts:AssumeRole"
        }]
    })
)

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
    runtime=aws.lambda_.Runtime.PYTHON3D8,
    code=pulumi.AssetArchive({
        ".": pulumi.FileArchive("./handler"),
    }),
    timeout=300,
    handler="handler.handler",
    role=role.arn,
    opts=pulumi.ResourceOptions(depends_on=[policy]),
)

api = apigateway.RestAPI('api', routes=[
    apigateway.RouteArgs(path="/", method="GET", event_handler=f),
])

pulumi.export('url', api.url)
```

{{% /choosable %}}

{{% choosable language csharp %}}

```csharp
using Pulumi;
using System.Collections.Generic;
using ApiGW = Pulumi.AwsApiGateway;
using Lambda = Pulumi.Aws.Lambda;
using Iam = Pulumi.Aws.Iam;

class MyStack : Stack
{
    public MyStack()
    {
        var lambdaRole = new Iam.Role("mylambda-role", new Iam.RoleArgs
        {
            AssumeRolePolicy =
                @"{
                ""Version"": ""2012-10-17"",
                ""Statement"": [{
                    ""Effect"": ""Allow"",
                    ""Principal"": { ""Service"": ""lambda.amazonaws.com"" },
                    ""Action"": ""sts:AssumeRole""
                }]
            }"
        });

        var rolePolicy = new Iam.RolePolicy("mylambda-policy", new Iam.RolePolicyArgs
        {
            Role = lambdaRole.Id,
            Policy =
               @"{
                ""Version"": ""2012-10-17"",
                ""Statement"": [{
                    ""Action"": [""logs:*"", ""cloudwatch:*""],
                    ""Resource"": ""*"",
                    ""Effect"": ""Allow""
                }]
            }"
        });

        // Closure serialization is not supported in multi-lang components
        // so we need to provide a handler function explicitly from the file-system.
        // Refer to https://github.com/pulumi/pulumi-aws-apigateway/tree/main/examples/simple-cs/handler
        // for an example handler.
        var lambda = new Lambda.Function("lambda", new Lambda.FunctionArgs
        {
            Runtime = Lambda.Runtime.Python3d8,
            Code = new AssetArchive(new Dictionary<string, AssetOrArchive>{
                ["."] = new FileArchive("./handler"),
            }),
            Timeout = 300,
            Handler = "handler.handler",
            Role = lambdaRole.Arn
        }, new Pulumi.CustomResourceOptions { DependsOn = { rolePolicy } });

        var restAPI = new ApiGW.RestAPI("api", new ApiGW.RestAPIArgs
        {
            Routes = new List<ApiGW.Inputs.RouteArgs>{
            new ApiGW.Inputs.RouteArgs{Path="/", Method=ApiGW.Method.GET, EventHandler=lambda}}
        });

        this.Url = restAPI.Url;
    }

    [Output]
    public Output<string> Url { get; set; }
}
```

{{% /choosable %}}

{{% choosable language go %}}

```go
package main

import (
 apigateway "github.com/pulumi/pulumi-aws-apigateway/sdk/go/apigateway"
 "github.com/pulumi/pulumi-aws/sdk/v4/go/aws/iam"
 "github.com/pulumi/pulumi-aws/sdk/v4/go/aws/lambda"
 "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
 pulumi.Run(func(ctx *pulumi.Context) error {

  role, err := iam.NewRole(ctx, "lambda-role", &iam.RoleArgs{
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

  policy, err := iam.NewRolePolicy(ctx, "lambda-policy", &iam.RolePolicyArgs{
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
  f, err := lambda.NewFunction(ctx, "lambda", &lambda.FunctionArgs{
   Runtime: lambda.RuntimePython3d8,
   Code: pulumi.NewAssetArchive(map[string]interface{}{
    ".": pulumi.NewFileArchive("./handler"),
   }),
   Timeout: pulumi.Int(300),
   Handler: pulumi.String("handler.handler"),
   Role:    role.Arn,
  }, pulumi.DependsOn([]pulumi.Resource{policy}))
  if err != nil {
   return err
  }

  getMethod := apigateway.MethodGET
  restAPI, err := apigateway.NewRestAPI(ctx, "api", &apigateway.RestAPIArgs{
   Routes: []apigateway.RouteArgs{
    apigateway.RouteArgs{
     Path:         "/",
     Method:       &getMethod,
     EventHandler: f,
    },
   },
  })
  if err != nil {
   return err
  }

  ctx.Export("url", restAPI.Url)
  return nil
 })
}
```

{{% /choosable %}}

{{% choosable language java %}}

```java

```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml

```

{{% /choosable %}}

{{% /chooser %}}
