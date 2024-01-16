---
title: Pinecone
meta_desc: Provides an overview of the Pinecone Provider for Pulumi.
layout: package
---

This Pulumi Pinecone Provider enables you to manage your [Pinecone](https://www.pinecone.io/) collections and indexes using any language of Pulumi Infrastructure as Code.

## Example

{{< chooser language "javascript,typescript,python,go,csharp,yaml" >}}


{{% choosable language javascript %}}

```javascript
"use strict";
const pulumi = require("@pulumi/pulumi");
const pinecone = require("@pinecone-database/pulumi");

const myExampleIndex = new pinecone.PineconeIndex("my-example-index", {
    name: "my-example-index",
    metric: pinecone.IndexMetric.Cosine,
    spec: {
        serverless: {
            cloud: pinecone.ServerlessSpecCloud.Aws,
            region: "us-west-2",
        }
    }
});

exports.host = myExampleIndex.host;
```

{{% /choosable %}}

{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as pinecone from "@pinecone-database/pulumi";

const myExampleIndex = new pinecone.PineconeIndex("my-example-index", {
    name: "example-index-ts",
    metric: pinecone.IndexMetric.Cosine,
    spec: {
        serverless: {
            cloud: pinecone.ServerlessSpecCloud.Aws,
            region: "us-west-2",
        },
    },
});
export const host = myExampleIndex.host;
```

{{% /choosable %}}

{{% choosable language python %}}

```python
"""A Python Pulumi program"""
import pulumi
import pinecone_pulumi as pinecone

my_pinecone_index = pinecone.PineconeIndex("myPineconeIndex",
   name="example-index",
   metric=pinecone.IndexMetric.COSINE,
   spec=pinecone.PineconeSpecArgs(
       serverless=pinecone.PineconeServerlessSpecArgs(
           cloud=pinecone.ServerlessSpecCloud.AWS,
           region="us-west-2",
       ),
   ))
pulumi.export("output", {
    "value": my_pinecone_index.host,
})
```

{{% /choosable %}}

{{% choosable language go %}}

```go
package main

import (
	"github.com/pinecone-io/pulumi-pinecone/sdk/go/pinecone"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		myExampleIndex, err := pinecone.NewPineconeIndex(ctx, "my-example-index", &pinecone.PineconeIndexArgs{
			Name:   pulumi.String("example-index-go"),
			Metric: pinecone.IndexMetricCosine,
			Spec: &pinecone.PineconeSpecArgs{
				Serverless: &pinecone.PineconeServerlessSpecArgs{
					Cloud:  pinecone.ServerlessSpecCloudAws,
					Region: pulumi.String("us-west-2"),
				},
			},
		})
		if err != nil {
			return err
		}
		ctx.Export("myPineconeIndexHost", myExampleIndex.Host)

		return nil
	})
}
```

{{% /choosable %}}

{{% choosable language csharp %}}

```csharp
using System.Collections.Generic;
using PineconeDatabase.Pinecone.Inputs;
using Pulumi;
using Pinecone = PineconeDatabase.Pinecone;

return await Deployment.RunAsync(() =>
{
    var myExampleIndex = new Pinecone.PineconeIndex("myExampleIndex", new Pinecone.PineconeIndexArgs
    {
        Name = "example-index-csharp",
        Metric= Pinecone.IndexMetric.Cosine,
        Spec= new Pinecone.Inputs.PineconeSpecArgs {
            Serverless= new PineconeServerlessSpecArgs{
                Cloud= Pinecone.ServerlessSpecCloud.Aws,
                Region= "us-west-2",
        }
    },
    });

    return new Dictionary<string, object?>
    {
        ["myPineconeIndexHost"] = myExampleIndex.Host
    };
});
```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
name: pinecone-serverless-yaml
description: A minimal Pinecone Serverless Pulumi YAML program
runtime: yaml

resources:
  myExampleIndex:
    type: pinecone:index:PineconeIndex
    properties:
      name: "example-index"
      metric: "cosine"
      spec:
        serverless:
          cloud: aws
          region: us-west-2

outputs:
  output:
    value: ${myExampleIndex.host}
```

{{% /choosable %}}

{{< /chooser >}}
