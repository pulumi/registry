---
title: AWS API Gateway
meta_desc: Easily create AWS API Gateway REST APIs using Pulumi.
layout: overview
---

Easily create AWS API Gateway REST APIs using Pulumi. This component exposes the Crosswalk for AWS functionality documented at https://www.pulumi.com/docs/guides/crosswalk/aws/api-gateway/ as a Pulumi Package available in all Pulumi languages.

Example:

{{< chooser language "typescript,python,go" >}}

{{% choosable language typescript %}}

```ts
import * as apigateway from "@pulumi/aws-apigateway";
import * as aws from "@pulumi/aws";

// Create a Lambda Function
const f = new aws.lambda.CallbackFunction("f", {
    callback: async (ev, ctx) => {
        console.log(JSON.stringify(ev));
        return {
            statusCode: 200,
            body: "goodbye",
        };
    },
});

// Create a REST API that invokes the Lambda Function
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

```py
# Create a Lambda Function
# f = ...

# Create a REST API that invokes the Lambda Function and uses a few other route kinds
api = apigateway.RestAPI('api', routes=[
    apigateway.RouteArgs(path="/", method="GET", event_handler=f),
    apigateway.RouteArgs(path="/www", method="GET", local_path="www", index=False),
    apigateway.RouteArgs(path="/integration", target=apigateway.TargetArgs(uri="https://www.google.com", type="http_proxy"))
])

pulumi.export('url', api.url)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
// Create a Lambda Function
// f, err := ...

// Create a REST API that invokes the Lambda Function
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
```

{{% /choosable %}}

{{< /chooser >}}

## SDK packages

The AWS API Gateway provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/aws-apigateway`](https://www.npmjs.com/package/@pulumi/aws)
* Python: [`pulumi-aws-apigateway`](https://pypi.org/project/pulumi-aws-apigateway/)
* Go: [`github.com/pulumi/pulumi-aws-apigateway/sdk/go/apigateway`](https://github.com/pulumi/pulumi-aws-apigateway)
* .NET: [`Pulumi.AwsApiGateway`](https://www.nuget.org/packages/Pulumi.AwsApiGateway)
