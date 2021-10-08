---
title: AWS API Gateway
meta_desc: Use Pulumi's Component for managing AWS API Gateway resources using infrastructure as code.
layout: overview
---

Amazon API Gateway is a fully managed service that makes it easy for developers to create, publish, maintain, monitor, and secure APIs at any scale. Use this Pulumi Component to build, deploy, and manage API Gateways as part of your application.

## Example

<!-- Provide a simple example of how to use your package, ideally in all languages. -->

{{< chooser language "typescript,python,go,csharp" >}}

{{% choosable language typescript %}}

```typescript
import * as apigateway from "@pulumi/aws-api-gateway";
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
        method: apigateway.Method.GET,
        eventHandler: f,
    }],
});

export const url = api.url;
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi
import pulumi_aws_api_gateway as apigateway

api = apigateway.RestAPI('api', routes=[
    apigateway.RouteArgs(path="/", method=apigateway.Method.GET, event_handler=f),
    apigateway.RouteArgs(path="/www", method=apigateway.Method.GET, local_path="www", index=False),
    apigateway.RouteArgs(path="/integration",
        target=apigateway.TargetArgs(uri="https://www.google.com", type=apigateway.IntegrationType.Http_proxy))
])

pulumi.export('url', api.url)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
package main

import (
	"github.com/pulumi/pulumi-aws-api-gateway/sdk/go/apigateway"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create cluster with default settings
		api, err := apigateway.NewRestAPI(ctx, "api", &apigateway.RestAPIArgs{
            Routes: []apigateway.RouteArgs{
                &apigateway.RouteArgs{
                    Path:         pulumi.String("/"),
                    Method:       apigateway.MethodGET,
                    EventHandler: f,
                },
                &apigateway.RouteArgs{
                    Path:      pulumi.String("/"),
                    Method:    apigateway.MethodGET,
                    LocalPath: pulumi.String("www"),
                    Index:     pulumi.Bool(false),
                },
                &apigateway.RouteArgs{
                    Path:   pulumi.String("/integration"),
                    Method: apigateway.MethodGET,
                    Target: &apigateway.TargetArgs{
                        Uri:  pulumi.String("https://www.google.com"),
                        Type: apigateway.IntegrationType.Http_proxy,
                    },
                },
            }
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
{{% choosable language csharp %}}

```csharp
using Pulumi;
using ApiGateway = Pulumi.AwsApiGateway;

class MyStack : Stack
{

    [Output("url")] Output<string> Url { get; set; }

    public MyStack()
    {
        var api = new ApiGateway.RestAPI("api", new ApiGateway.RestAPIArgs
        {
            Routes =
            {
                new ApiGateway.RouteArgs{ Path = "/", Method = ApiGateway.Method.GET, EventHandler = f },
                new ApiGateway.RouteArgs{
                    Path = "/www",
                    Method = ApiGateway.Method.GET,
                    LocalPath = "www",
                    Index = false
                },
                new ApiGateway.RouteArgs{
                    Path = "/integration",
                    Target = new ApiGateway.TargetArgs{
                        Uri = "https://www.google.com",
                        Type = ApiGateway.IntegrationType.Http_proxy,
                    }
                },
            },
        });

        this.Url = api.Url;
    }
}
```

{{% /choosable %}}

{{< /chooser >}}
