---
title: Fortios
meta_desc: Provides an overview of the Fortios Provider for Pulumi.
layout: package
---

The FortiOS provider is used to interact with the resources supported by FortiOS
and FortiManager. We need to configure the provider with the proper credentials
before it can be used.

Two products are supported now: FortiGate and FortiManager, please use the
navigation on the left to read more details about the available resources.

Detailed instructions how to configure the provider is available in the [Fortios Installation & Configuration](./installation-configuration).

## Issues

This is a community maintained provider. Please file issues and feature requests here:

[pulumiverse/pulumi-fortios](https://github.com/pulumiverse/pulumi-fortios/issues)

## Example

{{< chooser language "typescript,python,go,csharp,yaml" >}}

{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as fortios from "@pulumiverse/fortios";

const config = new pulumi.Config();
const fortiosCabundlefile = config.require("fortios:cabundlefile");
const fortiosHostname = config.require("fortios:hostname");
const fortiosInsecure = config.require("fortios:insecure");
const fortiosToken = config.require("fortios:token");

const fortiosProvider = new fortios.Provider("fortios-provider", {
    cabundlefile: fortiosCabundlefile,
    hostname: fortiosHostname,
    insecure: fortiosInsecure,
    token: fortiosToken,
});

const test1 = new fortios.networking.RouteStatic("test1", {
    dst: "110.2.2.122/32",
    gateway: "2.2.2.2",
    device: "device",
}, {
    provider: fortiosProvider,
});
```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi
import pulumiverse_fortios as fortios

config = pulumi.Config()

fortios_cabundlefile = config.require("fortios:cabundlefile")
fortios_hostname = config.require("fortios:hostname")
fortios_insecure = config.require("fortios:insecure")
fortios_token = config.require("fortios:token")

fortios_provider = fortios.Provider("fortios-provider",
    cabundlefile=fortios_cabundlefile,
    hostname=fortios_hostname,
    insecure=fortios_insecure,
    token=fortios_token)

test1 = fortios.networking.RouteStatic("test1",
    dst="110.2.2.122/32",
    gateway="2.2.2.2",
    device="device",
    opts=pulumi.ResourceOptions(provider=fortios_provider))
```

{{% /choosable %}}

{{% choosable language go %}}

```go
package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/networking"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")

		fortiosCabundlefile := cfg.Require("fortios:cabundlefile")
		fortiosHostname := cfg.Require("fortios:hostname")
		fortiosInsecure := cfg.Require("fortios:insecure")
		fortiosToken := cfg.Require("fortios:token")

		fortiosProvider, err := fortios.NewProvider(ctx, "fortios-provider", &fortios.ProviderArgs{
			Cabundlefile: pulumi.String(fortiosCabundlefile),
			Hostname:     pulumi.String(fortiosHostname),
			Insecure:     pulumi.String(fortiosInsecure),
			Token:        pulumi.String(fortiosToken),
		})
		if err != nil {
			return err
		}
		_, err = networking.NewRouteStatic(ctx, "test1", &networking.RouteStaticArgs{
			Dst:     pulumi.String("110.2.2.122/32"),
			Gateway: pulumi.String("2.2.2.2"),
			Device:  pulumi.String("device"),
		}, pulumi.Provider(fortiosProvider))
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
using Fortios = Pulumiverse.Fortios;

return await Deployment.RunAsync(() =>
{
    var config = new Config();

    var fortiosCabundlefile = config.Require("fortios:cabundlefile");
    var fortiosHostname = config.Require("fortios:hostname");
    var fortiosInsecure = config.Require("fortios:insecure");
    var fortiosToken = config.Require("fortios:token");

    var fortiosProvider = new Fortios.Provider("fortios-provider", new()
    {
        Cabundlefile = fortiosCabundlefile,
        Hostname = fortiosHostname,
        Insecure = fortiosInsecure,
        Token = fortiosToken,
    });

    var test1 = new Fortios.Networking.RouteStatic("test1", new()
    {
        Dst = "110.2.2.122/32",
        Gateway = "2.2.2.2",
        Device = "device",
    }, new CustomResourceOptions
    {
        Provider = fortiosProvider,
    });

});
```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
name: basic

runtime: yaml

config:
  fortios:cabundlefile:
    value: /path/yourCA.crt
  fortios:hostname:
    value: 192.168.52.177
  fortios:insecure:
    value: false
  fortios:token:
    value: jn3t3Nw7qckQzt955Htkfj5hwQ6jdb

resources:
  # Create provider instance
  fortios-provider:
    type: pulumi:providers:fortios
    properties:
      cabundlefile: ${fortios:cabundlefile}
      hostname: ${fortios:hostname}
      insecure: ${fortios:insecure}
      token: ${fortios:token}

  # Create a Static Route Item
  test1:
    type: fortios:networking:RouteStatic
    properties:
      dst: 110.2.2.122/32
      gateway: 2.2.2.2
      device: device
    options:
      provider: ${fortios-provider}
```

{{% /choosable %}}

{{< /chooser >}}
