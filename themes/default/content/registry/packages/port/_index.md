---
# WARNING: this file was fetched from https://raw.githubusercontent.com/port-labs/pulumi-port/v2.11.0/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
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
    properties: {
        stringProps: {
            "language": "typescript",
        }
    }
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
    properties: {
        stringProps: {
            "language": {
                default: "Go",
            }
        }
    }
});

export const entity = new port.Entity("monolith", {
    identifier: "monolith",
    title: "monolith",
    blueprint: blueprint.identifier,
    properties: {
        stringProps: {
            "language": "Node",
        }
    }
});
```

{{% /choosable %}}

{{% choosable language python %}}

```python
"""A Python Pulumi program"""

import pulumi
from port_pulumi import Entity,EntityPropertiesArgs

entity = Entity("port_pulumi", title="monolith", blueprint="microservice",
                properties=EntityPropertiesArgs(string_props={"language": "python"}),
                )
```

{{% /choosable %}}
{{% choosable language go %}}

```go
package main

import (
	"github.com/port-labs/pulumi-port/sdk/go/port"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		entity, err := port.NewEntity(ctx, "entity", &port.EntityArgs{
			Title:     pulumi.String("monolith"),
			Blueprint: pulumi.String("microservice"),
			Properties: port.EntityPropertiesArgs{
				StringProps: pulumi.StringMap{
					"language": pulumi.String("Go"),
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
