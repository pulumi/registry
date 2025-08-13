---
# WARNING: this file was fetched from https://raw.githubusercontent.com/splightplatform/pulumi-splight/v1.2.18/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Splight
meta_desc: Provides an overview of the Splight Provider for Pulumi.
layout: package
---

## Splight Provider for Pulumi

The Splight provider enables seamless interaction with resources supported by [Splight](https://www.splight-ai.com/).

## Example

{{< chooser language "typescript,python,go,csharp" >}}
{{% choosable language typescript %}}

```typescript
import * as splight from "@splightplatform/pulumi-splight";

new splight.Asset("MyAsset", {
  name: "MyAsset",
  description: "My Asset Description",
  geometry: JSON.stringify({
    type: "GeometryCollection",
    geometries: [
      {
        type: "GeometryCollection",
        geometries: [
          {
            type: "Point",
            coordinates: [0, 0],
          },
        ],
      },
    ],
  }),
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import geojson
import pulumi_splight as splight


splight.Asset(
    resource_name="MyAsset",
    name="MyAsset",
    description="My Asset Description",
    geometry=geojson.dumps(
        geojson.GeometryCollection(
            geometries=[{"type": "Point", "coordinates": [0, 0]}]
        )
    ),
)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
package main

import (
	"encoding/json"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	splight "github.com/splightplatform/pulumi-splight/sdk/go/splight"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		geometry := map[string]interface{}{
			"type": "GeometryCollection",
			"geometries": []interface{}{
				map[string]interface{}{
					"type":        "Point",
					"coordinates": []float64{0.0, 0.0},
				},
			},
		}

		geometryJSON, err := json.Marshal(geometry)

		if err != nil {
			return err
		}

		_, err = splight.NewAsset(ctx, "MyAsset",
			&splight.AssetArgs{
				Name:        pulumi.String("MyAsset"),
				Description: pulumi.String("My Asset Description"),
				Geometry:    pulumi.String(geometryJSON),
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

```cs
using System.Text.Json;
using Pulumi;

class Program
{
    static Task<int> Main() => Deployment.RunAsync<MyStack>();
}

class MyStack : Stack
{
    public MyStack()
    {
        var geometry = new Dictionary<string, object>
        {
            { "type", "GeometryCollection" },
            { "geometries", new List<Dictionary<string, object>>
                {
                    new Dictionary<string, object>
                    {
                        { "type", "Point" },
                        { "coordinates", new List<double> { 0.0, 0.0 } }
                    }
                }
            }
        };

        var geometryJson = JsonSerializer.Serialize(geometry);

        var myAsset = new Splight.Splight.Asset("MyAsset", args: new Splight.Splight.AssetArgs
        {
            Name = "MyAsset",
            Description = "My Asset Description",
            Geometry = geometryJson
        });
    }
}
```

{{% /choosable %}}
{{< /chooser >}}
