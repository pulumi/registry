---
title: Lambda Builders
meta_desc: Learn how you can use Pulumi's Lambda Builders Component to simplify building Lambda Function deployment packages.
layout: package
---

Pulumi Lambda Builders is a library that provides utilities for easily
building/bundling Lambda Function code. The library currently supports building
`go` Lambdas, but will eventually support the below languages/build tools.

- Java with Gradle
- Java with Maven
- Dotnet with amazon.lambda.tools
- Python with Pip
- Javascript with Npm
- Typescript with esbuild
- Ruby with Bundler
- Go with Mod
- Rust with Cargo

This library integrates with the
[aws-lambda-builders](https://github.com/aws/aws-lambda-builders) library which
provides the building utilities.

## Examples

### Build Golang Lambda

Builds a Golang Lambda Function into a Pulumi Asset that can be deployed.

The below example uses a folder structure like this:
```tree
examples/simple-go
├── Pulumi.yaml
├── cmd
│   └── simple
│       └── main.go
├── go.mod
├── go.sum
└── main.go
```

The output of `buildGo` produces an asset that can be passed to the
`aws.Lambda` `Code` property.

{{< chooser language "typescript,python,go,csharp,java,yaml" / >}}

{{% choosable language "javascript,typescript" %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as lambdaBuilders from "@pulumi/lambda-builders";

const builder = lambdaBuilders.buildGo({
    architecture: "arm64",
    code: "cmd/simple",
});
const lambdaRolePolicy = aws.iam.getPolicyDocumentOutput({
    statements: [{
        actions: ["sts:AssumeRole"],
        principals: [{
            type: "Service",
            identifiers: ["lambda.amazonaws.com"],
        }],
    }],
});
const role = new aws.iam.Role("role", {
   assumeRolePolicy: lambdaRolePolicy.apply(lambdaRolePolicy => lambdaRolePolicy.json),
});
new aws.lambda.Function("function", {
    code: builder.asset,
    architectures: ["arm64"],
    handler: "bootstrap",
    role: role.arn,
    runtime: aws.lambda.Runtime.CustomAL2023,
});
```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi
import pulumi_aws as aws
import pulumi_lambda_builders as lambda_builders

builder = lambda_builders.build_go(architecture="arm64",
    code="cmd/simple")
lambda_role_policy = aws.iam.get_policy_document_output(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    actions=["sts:AssumeRole"],
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="Service",
        identifiers=["lambda.amazonaws.com"],
    )],
)])
role = aws.iam.Role("role", assume_role_policy=lambda_role_policy.json)
function = aws.lambda_.Function("function",
    code=builder.asset,
    architectures=["arm64"],
    handler="bootstrap",
    role=role.arn,
    runtime=aws.lambda_.Runtime.CUSTOM_AL2023)
```

{{% /choosable %}}

{{% choosable language go %}}

```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lambda"
	lambdabuilders "github.com/pulumi/pulumi-lambda-builders/sdk/go/lambda-builders"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		builder, err := lambdabuilders.BuildGo(ctx, &lambdabuilders.BuildGoArgs{
			Architecture: pulumi.StringRef("arm64"),
			Code:         pulumi.StringRef("cmd/simple"),
		}, nil)
		if err != nil {
			return err
		}
		lambdaRolePolicy, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Actions: []string{
						"sts:AssumeRole",
					},
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Type: "Service",
							Identifiers: []string{
								"lambda.amazonaws.com",
							},
						},
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		role, err := iam.NewRole(ctx, "role", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.String(lambdaRolePolicy.Json),
		})
		if err != nil {
			return err
		}
		_, err = lambda.NewFunction(ctx, "function", &lambda.FunctionArgs{
			Code: builder.Asset,
			Architectures: pulumi.StringArray{
				pulumi.String("arm64"),
			},
			Handler: pulumi.String("bootstrap"),
			Role:    role.Arn,
			Runtime: pulumi.String(lambda.RuntimeCustomAL2023),
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
using LambdaBuilders = Pulumi.LambdaBuilders;

return await Deployment.RunAsync(() => 
{
    var builder = LambdaBuilders.BuildGo.Invoke(new()
    {
        Architecture = "arm64",
        Code = "cmd/simple",
    });

    var lambdaRolePolicy = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Actions = new[]
                {
                    "sts:AssumeRole",
                },
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Type = "Service",
                        Identifiers = new[]
                        {
                            "lambda.amazonaws.com",
                        },
                    },
                },
            },
        },
    });

    var role = new Aws.Iam.Role("role", new()
    {
        AssumeRolePolicy = lambdaRolePolicy.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var function = new Aws.Lambda.Function("function", new()
    {
        Code = builder.Apply(buildGoResult => buildGoResult.Asset),
        Architectures = new[]
        {
            "arm64",
        },
        Handler = "bootstrap",
        Role = role.Arn,
        Runtime = Aws.Lambda.Runtime.CustomAL2023,
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
import com.pulumi.lambdabuilders.LambdabuildersFunctions;
import com.pulumi.lambdabuilders.inputs.BuildGoArgs;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.lambda.Function;
import com.pulumi.aws.lambda.FunctionArgs;
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
        final var builder = Lambda-buildersFunctions.buildGo(BuildGoArgs.builder()
            .architecture("arm64")
            .code("cmd/simple")
            .build());

        final var lambdaRolePolicy = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .actions("sts:AssumeRole")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("Service")
                    .identifiers("lambda.amazonaws.com")
                    .build())
                .build())
            .build());

        var role = new Role("role", RoleArgs.builder()
            .assumeRolePolicy(lambdaRolePolicy.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var function = new Function("function", FunctionArgs.builder()
            .code(buildGoResult.asset())
            .architectures("arm64")
            .handler("bootstrap")
            .role(role.arn())
            .runtime("provided.al2023")
            .build());

    }
}
```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
  role:
    type: aws:iam:Role
    properties:
      assumeRolePolicy: ${lambdaRolePolicy.json}
  function:
    type: aws:lambda:Function
    properties:
      code: ${builder.asset}
      architectures:
        - arm64
      handler: bootstrap
      role: ${role.arn}
      runtime: provided.al2023
variables:
  builder:
    fn::invoke:
      Function: lambda-builders:index:buildGo
      Arguments:
        architecture: arm64
        code: cmd/simple
  lambdaRolePolicy:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - actions:
              - sts:AssumeRole
            principals:
              - type: Service
                identifiers:
                  - lambda.amazonaws.com
```

{{% /choosable %}}
