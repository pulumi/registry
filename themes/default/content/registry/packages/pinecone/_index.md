---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pinecone-io/pulumi-pinecone/v2.0.2/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/pinecone-io/pulumi-pinecone/blob/v2.0.2/docs/_index.md
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

const myPineconeIndex = new pinecone.Index("myPineconeIndex", {
    name: "example-index",
    dimension: 10,
    spec: {
        serverless: {
            cloud: "aws",
            region: "us-east-1",
        },
    },
});
exports.name = myPineconeIndex.name;
exports.host = myPineconeIndex.host;
```

{{% /choosable %}}

{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as pinecone from "@pinecone-database/pulumi";

const myPineconeIndex = new pinecone.Index("myPineconeIndex", {
    name: "example-index",
    dimension: 10,
    spec: {
        serverless: {
            cloud: "aws",
            region: "us-east-1",
        },
    },
});
export const name = myPineconeIndex.name;
export const host = myPineconeIndex.host;
```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi
import pinecone_pulumi as pinecone

my_pinecone_index = pinecone.Index("myPineconeIndex",
   name="example-index",
   dimension=10,
   spec={
       "serverless": {
           "cloud": "aws",
           "region": "us-east-1",
       },
   })
pulumi.export("name", my_pinecone_index.name)
pulumi.export("host", my_pinecone_index.host)

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

		myPineconeIndex, err := pinecone.NewIndex(ctx, "myPineconeIndex", &pinecone.IndexArgs{
			Name:      pulumi.String("example-index2"),
			Dimension: pulumi.Int(10),
			Spec: pinecone.IndexSpecArgs{
				Serverless: pinecone.IndexSpecServerlessArgs{
					Cloud:  pulumi.String("aws"),
					Region: pulumi.String("us-east-1"),
				},
			},
		})
		if err != nil {
			return err
		}
		ctx.Export("name", myPineconeIndex.Name)
		ctx.Export("host", myPineconeIndex.Host)

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
using Pinecone = PineconeDatabase.Pinecone;

return await Deployment.RunAsync(() => 
{
    var myPineconeIndex = new Pinecone.Index("myPineconeIndex", new()
    {
        Name = "example-index",
        Dimension = 10,
        Spec = new Pinecone.Inputs.IndexSpecArgs
        {
            Serverless = new Pinecone.Inputs.IndexSpecServerlessArgs
            {
                Cloud = "aws",
                Region = "us-east-1",
            },
        },
    });

    return new Dictionary<string, object?>
    {
        ["name"] = myPineconeIndex.Name,
        ["host"] = myPineconeIndex.Host,
    };
});
```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
# File: examples/yaml/Pulumi.yaml
name: pinecone-yaml
runtime: yaml

resources:
  myPineconeIndex:
    type: pinecone:index:Index
    properties:
      name: "example-index"
      dimension: 10
      spec:
        serverless:
          cloud: aws
          region: us-east-1

outputs:
  name: ${myPineconeIndex.name}
  host: ${myPineconeIndex.host}
```

{{% /choosable %}}

{{< /chooser >}}
