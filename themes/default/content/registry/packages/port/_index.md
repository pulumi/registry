---
title: Port
meta_desc: Provides an overview of the Port Provider for Pulumi.
layout: package
---

The Port Resource Provider lets you manage [Port](https://www.getport.io) resources.

## Example

{{< chooser language "javascript,typescript,python,go" >}}


{{% choosable language javascript %}}

```javascript
"use strict";
const pulumi = require("@pulumi/pulumi");
const port = require("@port-labs/port");

const entity = new port.Entity("entity", {
    identifier: "monolith",
    title: "monolith",
    blueprint: "microservice",
    properties: [
        {
            name: "language",
            value: "Node",
        }
    ]
});

exports.title = entity.title;
```

{{% /choosable %}}

{{% choosable language typescript %}}

```typescript
import * as port from "@port-labs/port";

export const blueprint = new port.Blueprint("microservice", {
    identifier: "microservice",
    title: "Microservice",
    properties: [
        {
            identifier: "language",
            title: "Language",
            type: "string",
            required: false,
        }
    ]
});

export const entity = new port.Entity("monolith", {
    identifier: "monolith",
    title: "monolith",
    blueprint: blueprint.identifier,
    properties: [
        {
            name: "language",
            value: "Node",
        }
    ]
});
```

{{% /choosable %}}

{{% choosable language python %}}

```python
"""A Python Pulumi program"""

import pulumi
from port_pulumi import Entity

entity = Entity("port_pulumi", title="monolith", blueprint="microservice",
                properties=[{"name": "language", "value": "Python"}])
```

{{% /choosable %}}
{{% choosable language go %}}

```go
package main

import (
	"github.com/port-labs/pulumi-port/go/port"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		entity, err := port.NewEntity(ctx, "entity", &port.EntityArgs{
			Title:     pulumi.String("monolith"),
			Blueprint: pulumi.String("microservice"),
			Properties: port.EntityPropertyArray{
				&port.EntityPropertyArgs{
					Name:  pulumi.String("language"),
					Value: pulumi.String("GO"),
				},
			},
		})
		ctx.Export("entity", entity.Title)
		if err != nil {
			return err
		}
		return nil
	})
}

```

{{% /choosable %}}

{{< /chooser >}}
