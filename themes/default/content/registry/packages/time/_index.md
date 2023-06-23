---
title: Time
meta_desc: Provides an overview of the Time Provider for Pulumi.
layout: package
---

The time provider is used to interact with time-based resources. 
The provider itself has no configuration options.

## Example

{{< chooser language "typescript,python,go,csharp,yaml" >}}


{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as time from "@pulumiverse/time";

export = async () => {
    const config = new pulumi.Config();
    const timeTrigger = config.require("time-trigger");

    const provider = new time.Provider("provider", {});

    const rotating = new time.Rotating("rotating", {
        rotationDays: 30,
        triggers: {
            trigger1: timeTrigger,
        },
    });

    const offset = new time.Offset("offset", {offsetDays: 7});

    return {
        "rotating-output-unix": rotating.unix,
        "rotating-output-rfc3339": rotating.rfc3339,
        "offset-date": pulumi.interpolate`${offset.day}-${offset.month}-${offset.year}`,
    };
}
```
{{% /choosable %}}

{{% choosable language python %}}
```python
import pulumi
import pulumiverse_time as time

config = pulumi.Config()

time_trigger = config.require("time-trigger")

provider = time.Provider("provider")

rotating = time.Rotating("rotating",
    rotation_days=30,
    triggers={
        "trigger1": time_trigger,
    })

offset = time.Offset("offset", offset_days=7)

pulumi.export("rotating-output-unix", rotating.unix)
pulumi.export("rotating-output-rfc3339", rotating.rfc3339)
pulumi.export("offset-date", pulumi.Output.all(offset.day, offset.month, offset.year).apply(lambda day, month, year: f"{day}-{month}-{year}"))
```
{{% /choosable %}}

{{% choosable language go %}}
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
	"github.com/pulumiverse/pulumi-time/sdk/go/time"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")
		timeTrigger := cfg.Require("time-trigger")
		_, err := time.NewProvider(ctx, "provider", nil)
		if err != nil {
			return err
		}
		rotating, err := time.NewRotating(ctx, "rotating", &time.RotatingArgs{
			RotationDays: pulumi.Int(30),
			Triggers: pulumi.StringMap{
				"trigger1": pulumi.String(timeTrigger),
			},
		})
		if err != nil {
			return err
		}
		offset, err := time.NewOffset(ctx, "offset", &time.OffsetArgs{
			OffsetDays: pulumi.Int(7),
		})
		if err != nil {
			return err
		}
		ctx.Export("rotating-output-unix", rotating.Unix)
		ctx.Export("rotating-output-rfc3339", rotating.Rfc3339)
		ctx.Export("offset-date", pulumi.All(offset.Day, offset.Month, offset.Year).ApplyT(func(_args []interface{}) (string, error) {
			day := _args[0].(int)
			month := _args[1].(int)
			year := _args[2].(int)
			return fmt.Sprintf("%v-%v-%v", day, month, year), nil
		}).(pulumi.StringOutput))
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
using Time = Pulumiverse.Time;

return await Deployment.RunAsync(() =>
{
    var config = new Config();
    var timeTrigger = config.Require("time-trigger");
    var provider = new Time.Provider("provider");

    var rotating = new Time.Rotating("rotating", new()
    {
        RotationDays = 30,
        Triggers =
        {
            { "trigger1", timeTrigger },
        },
    });

    var offset = new Time.Offset("offset", new()
    {
        OffsetDays = 7,
    });

    return new Dictionary<string, object?>
    {
        ["rotating-output-unix"] = rotating.Unix,
        ["rotating-output-rfc3339"] = rotating.Rfc3339,
        ["offset-date"] = Output.Tuple(offset.Day, offset.Month, offset.Year).Apply(values =>
        {
            var day = values.Item1;
            var month = values.Item2;
            var year = values.Item3;
            return $"{day}-{month}-{year}";
        }),
    };
});
```
{{% /choosable %}}

{{% choosable language yaml %}}
```yaml
config:
  time-trigger:
    type: string

resources:
  provider:
    type: pulumi:providers:time

  rotating:
    type: time:Rotating
    properties:
      rotationDays: 30
      triggers:
        trigger1: ${time-trigger}

  offset:
    type: time:Offset
    properties:
      offsetDays: 7

outputs:
  rotating-output-unix: ${rotating.unix}
  rotating-output-rfc3339: ${rotating.rfc3339}
  offset-date: "${offset.day}-${offset.month}-${offset.year}"
```
{{% /choosable %}}

{{< /chooser >}}
