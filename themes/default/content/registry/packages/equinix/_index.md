---
# WARNING: this file was fetched from https://raw.githubusercontent.com/equinix/pulumi-equinix/v0.22.0/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Equinix
meta_desc: Provides an overview of the Equinix Provider for Pulumi.
layout: package
---

The Equinix provider for Pulumi can be used to provision [Equinix](https://deploy.Equinix.com) resources (Metal, Fabric, and Network Edge).

The Equinix provider must be configured with credentials to create and update resources in Equinix.

## Example

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}

{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as equinix from "@equinix-labs/pulumi-equinix";

const config = new pulumi.Config();
const projectId = config.require("projectId");
const web = new equinix.metal.Device("web", {
    hostname: "webserver1",
    plan: "c3.small.x86",
    operatingSystem: "ubuntu_20_04",
    metro: "sv",
    billingCycle: "hourly",
    projectId: projectId,
});
export const webPublicIp = pulumi.interpolate`http://${web.accessPublicIpv4}`;
```
{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi
import pulumi_equinix as equinix

config = pulumi.Config()
project_id = config.require("projectId")
web = equinix.metal.Device("web",
    hostname="webserver1",
    plan="c3.small.x86",
    operating_system="ubuntu_20_04",
    metro="sv",
    billing_cycle="hourly",
    project_id=project_id)
pulumi.export("webPublicIp", web.access_public_ipv4.apply(lambda access_public_ipv4: f"http://{access_public_ipv4}"))
```
{{% /choosable %}}

{{% choosable language go %}}
```go
package main

import (
	"fmt"

	"github.com/equinix/pulumi-equinix/sdk/go/equinix/metal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")
		projectId := cfg.Require("projectId")
		web, err := metal.NewDevice(ctx, "web", &metal.DeviceArgs{
			Hostname:        pulumi.String("webserver1"),
			Plan:            pulumi.String("c3.small.x86"),
			OperatingSystem: pulumi.String("ubuntu_20_04"),
			Metro:           pulumi.String("sv"),
			BillingCycle:    pulumi.String("hourly"),
			ProjectId:       pulumi.String(projectId),
		})
		if err != nil {
			return err
		}
		ctx.Export("webPublicIp", web.AccessPublicIpv4.ApplyT(func(accessPublicIpv4 string) (string, error) {
			return fmt.Sprintf("http://%v", accessPublicIpv4), nil
		}).(pulumi.StringOutput))
		return nil
	})
}
```
{{% /choosable %}}

{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using Pulumi;
using Equinix = Pulumi.Equinix;

return await Deployment.RunAsync(() => 
{
    var config = new Config();
    var projectId = config.Require("projectId");
    var web = new Equinix.Metal.Device("web", new()
    {
        Hostname = "webserver1",
        Plan = "c3.small.x86",
        OperatingSystem = "ubuntu_20_04",
        Metro = "sv",
        BillingCycle = "hourly",
        ProjectId = projectId,
    });

    return new Dictionary<string, object?>
    {
        ["webPublicIp"] = web.AccessPublicIpv4.Apply(accessPublicIpv4 => $"http://{accessPublicIpv4}"),
    };
});
```
{{% /choosable %}}

{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.equinix.pulumi.metal.Device;
import com.equinix.pulumi.metal.DeviceArgs;
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
        final var config = ctx.config();
        final var projectId = config.get("projectId");
        var web = new Device("web", DeviceArgs.builder()        
            .hostname("webserver1")
            .plan("c3.small.x86")
            .operatingSystem("ubuntu_20_04")
            .metro("sv")
            .billingCycle("hourly")
            .projectId(projectId)
            .build());

        ctx.export("webPublicIp", web.accessPublicIpv4().applyValue(accessPublicIpv4 -> String.format("http://%s", accessPublicIpv4)));
    }
}
```
{{% /choosable %}}

{{% choosable language yaml %}}
```yaml
name: equinix-metal-device
runtime: yaml
description: An Equinix Metal Device resource
config:
  projectId:
    type: string
resources:
  web:
    type: equinix:metal:Device
    properties:
      hostname: webserver1
      plan: c3.small.x86
      operatingSystem: ubuntu_20_04
      metro: sv
      billingCycle: hourly
      projectId: ${projectId}
outputs:
  webPublicIp: http://${web.accessPublicIpv4}
```
{{% /choosable %}}

{{< /chooser >}}
