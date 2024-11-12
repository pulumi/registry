---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumiverse/pulumi-fortios/v0.0.6/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Fortios Installation, Configuration and Usage
meta_desc: Information on how to install, configure use and debug the Fortios provider.
layout: package
---

## Installation

The Pulumi Fortios provider is available as a package in all Pulumi languages:

- JavaScript/TypeScript: [`@pulumiverse/fortios`](https://www.npmjs.com/package/@pulumiverse/fortios)
- Python: [`pulumiverse_fortios`](https://pypi.org/project/pulumiverse_fortios/)
- Go: [`github.com/pulumiverse/pulumi-fortios/sdk/go/fortios`](https://pkg.go.dev/github.com/pulumiverse/pulumi-fortios/sdk/go/fortios)
- .NET: [`Pulumiverse.Fortios`](https://www.nuget.org/packages/Pulumiverse.Fortios)

### Provider Binary

The Fortios provider binary is a third party binary. It can be installed using
the `pulumi plugin` command.

```bash
pulumi plugin install resource fortios <version> --server github://api.github.com/pulumiverse
```

Replace the version string `<version>` with your desired version.

## Configuration for FortiGate  

### Example usage

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

If it is used for testing, you can set `insecure` to `true` and unset
`cabundlefile` to quickly set the provider up, for example:

{{< chooser language "typescript,python,go,csharp,yaml" >}}

{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as fortios from "@pulumiverse/fortios";

const config = new pulumi.Config();
const fortiosHostname = config.require("fortios:hostname");
const fortiosToken = config.require("fortios:token");

const fortiosProvider = new fortios.Provider("fortios-provider", {
    cabundlefile: fortiosCabundlefile,
    hostname: fortiosHostname,
    insecure: true,
    token: fortiosToken,
});

```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi
import pulumiverse_fortios as fortios

config = pulumi.Config()

fortios_hostname = config.require("fortios:hostname")
fortios_token = config.require("fortios:token")

fortios_provider = fortios.Provider("fortios-provider",
    cabundlefile=fortios_cabundlefile,
    hostname=fortios_hostname,
    insecure=true,
    token=fortios_token)

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

		fortiosHostname := cfg.Require("fortios:hostname")
		fortiosToken := cfg.Require("fortios:token")

		fortiosProvider, err := fortios.NewProvider(ctx, "fortios-provider", &fortios.ProviderArgs{
			Cabundlefile: pulumi.String(fortiosCabundlefile),
			Hostname:     pulumi.String(fortiosHostname),
			Insecure:     pulumi.Bool(true),
			Token:        pulumi.String(fortiosToken),
		})
		if err != nil {
			return err
		}

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

    var fortiosHostname = config.Require("fortios:hostname");
    var fortiosToken = config.Require("fortios:token");

    var fortiosProvider = new Fortios.Provider("fortios-provider", new()
    {
        Hostname = fortiosHostname,
        Insecure = true,
        Token = fortiosToken,
    });

});
```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
name: basic

runtime: yaml

config:
  fortios:hostname:
    value: 192.168.52.177
  fortios:insecure:
    value: true
  fortios:token:
    value: jn3t3Nw7qckQzt955Htkfj5hwQ6jdb

resources:
  # Create provider instance
  fortios-provider:
    type: pulumi:providers:fortios
    properties:
      hostname: ${fortios:hostname}
      insecure: ${fortios:insecure}
      token: ${fortios:token}

```

{{% /choosable %}}

{{< /chooser >}}

Please refer to the Argument Reference below for more help on `insecure` and `cabundlefile`.

### Authentication

The FortiOS provider offers a means of providing credentials for authentication.
The following methods are supported:

- Static credentials
- Environment variables

#### Static credentials

Static credentials can be provided by adding credential keys in-line in the
FortiOS provider block. Default using `token` if both credential are provided.

There are two kinds of credentials supported:

- `token` based authentication (Recommended). User needs to generate an API
  token from FortiOS.
- `username/password` authentication. User provide the username and password of
  the administrator.

#### Generate an API token for FortiOS

A descriptions how to create an API token for Fortios, switch to chapter [Authentication](#authentication)

#### Environment variables

You can provide your credentials via the

- `FORTIOS_ACCESS_HOSTNAME`,
- `FORTIOS_ACCESS_TOKEN`
- `FORTIOS_ACCESS_USERNAME`
- `FORTIOS_ACCESS_PASSWORD`
- `FORTIOS_INSECURE` and
- `FORTIOS_CA_CABUNDLE`

environment variables, which releases you from configuring the provider
explicitly in Pulumi like shown above.

> Note that setting your FortiOS credentials using static credentials variables will
> override the environment variables.

**Usage:**

```shell
$ export "FORTIOS_ACCESS_HOSTNAME"="192.168.52.177"
$ export "FORTIOS_INSECURE"="false"
$ export "FORTIOS_ACCESS_TOKEN"="09m441wrwc10yGyrtQ4nk6mjbqcfz9"
$ export "FORTIOS_CA_CABUNDLE"="/path/yourCA.crt"
```

### VDOM

If the FortiGate unit is running in VDOM mode, the `vdom` configuration needs to
be added.

{{< chooser language "typescript,python,go,csharp,yaml" >}}

{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as fortios from "@pulumiverse/fortios";

const fortiosProvider = new fortios.Provider("fortios-provider", {
    cabundlefile: "/path/yourCA.crt",
    hostname: "192.168.52.177",
    insecure: false,
    token: "q3Hs49jxts195gkd9Hjsxnjtmr6k39",
    vdom: "vmdomtest",
});
const test1 = new fortios.networking.RouteStatic("test1", {
    dst: "120.2.2.122/32",
    gateway: "2.2.2.2",
    blackhole: "disable",
    distance: "22",
    weight: "3",
    priority: "3",
    device: "lbforvdomtest",
    comment: "Terraform test",
}, {
    provider: fortiosProvider,
});
```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumiverse_fortios as fortios

fortios_provider = fortios.Provider("fortios-provider",
    cabundlefile="/path/yourCA.crt",
    hostname="192.168.52.177",
    insecure=False,
    token="q3Hs49jxts195gkd9Hjsxnjtmr6k39",
    vdom="vmdomtest")

test1 = fortios.networking.RouteStatic("test1",
    dst="120.2.2.122/32",
    gateway="2.2.2.2",
    blackhole="disable",
    distance="22",
    weight="3",
    priority="3",
    device="lbforvdomtest",
    comment="Terraform test",
    opts=pulumi.ResourceOptions(provider=fortios_provider))
```

{{% /choosable %}}

{{% choosable language go %}}

```go
package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/networking"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		fortiosProvider, err := fortios.NewProvider(ctx, "fortios-provider", &fortios.ProviderArgs{
			Cabundlefile: pulumi.String("/path/yourCA.crt"),
			Hostname:     pulumi.String("192.168.52.177"),
			Insecure:     pulumi.Bool(false),
			Token:        pulumi.String("q3Hs49jxts195gkd9Hjsxnjtmr6k39"),
			Vdom:         pulumi.String("vmdomtest"),
		})
		if err != nil {
			return err
		}
		_, err = networking.NewRouteStatic(ctx, "test1", &networking.RouteStaticArgs{
			Dst:       pulumi.String("120.2.2.122/32"),
			Gateway:   pulumi.String("2.2.2.2"),
			Blackhole: pulumi.String("disable"),
			Distance:  pulumi.String("22"),
			Weight:    pulumi.String("3"),
			Priority:  pulumi.String("3"),
			Device:    pulumi.String("lbforvdomtest"),
			Comment:   pulumi.String("Terraform test"),
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
using Fortios = Pulumiverse.Fortios;

return await Deployment.RunAsync(() =>
{
    var fortiosProvider = new Fortios.Provider("fortios-provider", new()
    {
        Cabundlefile = "/path/yourCA.crt",
        Hostname = "192.168.52.177",
        Insecure = false,
        Token = "q3Hs49jxts195gkd9Hjsxnjtmr6k39",
        Vdom = "vmdomtest",
    });

    var test1 = new Fortios.Networking.RouteStatic("test1", new()
    {
        Dst = "120.2.2.122/32",
        Gateway = "2.2.2.2",
        Blackhole = "disable",
        Distance = "22",
        Weight = "3",
        Priority = "3",
        Device = "lbforvdomtest",
        Comment = "Terraform test",
    }, new CustomResourceOptions
    {
        Provider = fortiosProvider,
    });

});
```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
name: vdom
runtime: yaml

config: {}

resources:
  # Create provider instance
  fortios-provider:
    type: pulumi:providers:fortios
    properties:
      cabundlefile: /path/yourCA.crt
      hostname: 192.168.52.177
      insecure: false
      token: q3Hs49jxts195gkd9Hjsxnjtmr6k39
      vdom: vmdomtest

  test1:
    type: fortios:networking:RouteStatic
    properties:
      dst: 120.2.2.122/32
      gateway: 2.2.2.2
      blackhole: disable
      distance: '22'
      weight: '3'
      priority: '3'
      device: lbforvdomtest
      comment: Terraform test
    options:
      provider: ${fortios-provider}
```

{{% /choosable %}}

{{< /chooser >}}

By default, each resource inherits the provider's global vdom settings, but it
can also set its own vdom through the `vdomparam` of each resource. See the
`vdomparam` argument of each resource for details.

### Argument Reference

The following arguments are supported:

- `hostname` - (Optional) The hostname or IP address of FortiOS unit. It must be
  provided, but it can also be sourced from the `FORTIOS_ACCESS_HOSTNAME`
  environment variable.

- `token` - (Optional) The token of FortiOS unit. If omitted, the
  `FORTIOS_ACCESS_TOKEN` environment variable will be used. If neither is set,
  username/password will be used.

- `username` - (Optional) The username of FortiOS unit. If omitted, the
  `FORTIOS_ACCESS_USERNAME` environment variable will be used.

- `password` - (Optional) The password of FortiOS unit.If omitted, the
  `FORTIOS_ACCESS_PASSWORD` environment variable will be used.

- `insecure` - (Optional) Control whether the Provider to perform insecure SSL
  requests. If omitted, the `FORTIOS_INSECURE` environment variable is used. If
  neither is set, default value is `false`.

- `cabundlefile` - (Optional) The path of a custom CA bundle file. You can
  specify a path to the file, or you can specify it by the `FORTIOS_CA_CABUNDLE`
  environment variable.

- `cabundlecontent` - (Optional) The content of a custom CA bundle file. Note:
  `cabundlefile` and `cabundlecontent` cannot exist at the same time! Please
  only configure one of them.

- `vdom` - (Optional) If the FortiGate unit is running in VDOM mode, you can use
  this argument to specify the name of the vdom to be set .

- `http_proxy` - (Optional) HTTP proxy address. Set this argument to `ENV` if
  you want to use environment settings. By setting this argument to `ENV`, the
  provider will get the environment variable `HTTPS_PROXY` or `HTTP_PROXY`.
  Default is empty, which means no HTTP proxy.

## Configuration for FortiManager

### Example Usage

{{< chooser language "typescript,python,go,csharp,yaml" >}}

{{% choosable language typescript %}}

```typescript
import * as fortios from "@pulumiverse/fortios";

const fortiosProvider = new fortios.Provider("fortios-provider", {
    fmgCabundlefile: "/path/yourCA.crt",
    fmgHostname: "192.168.88.100",
    fmgInsecure: false,
    fmgPasswd: "admin",
    fmgUsername: "APIUser",
});

const test1 = new fortios.fmg.SystemDns("test1", {
    primary: "208.91.112.52",
    secondary: "208.91.112.54",
}, {
    provider: fortiosProvider,
});
```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumiverse_fortios as fortios

fortios_provider = fortios.Provider("fortios-provider",
    fmg_cabundlefile="/path/yourCA.crt",
    fmg_hostname="192.168.88.100",
    fmg_insecure=False,
    fmg_passwd="admin",
    fmg_username="APIUser")

test1 = fortios.fmg.SystemDns("test1",
    primary="208.91.112.52",
    secondary="208.91.112.54",
    opts=pulumi.ResourceOptions(provider=fortios_provider))
```

{{% /choosable %}}

{{% choosable language go %}}

```go
package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/fmg"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		fortiosProvider, err := fortios.NewProvider(ctx, "fortios-provider", &fortios.ProviderArgs{
			FmgCabundlefile: pulumi.String("/path/yourCA.crt"),
			FmgHostname:     pulumi.String("192.168.88.100"),
			FmgInsecure:     pulumi.Bool(false),
			FmgPasswd:       pulumi.String("admin"),
			FmgUsername:     pulumi.String("APIUser"),
		})
		if err != nil {
			return err
		}
		_, err = fmg.NewSystemDns(ctx, "test1", &fmg.SystemDnsArgs{
			Primary:   pulumi.String("208.91.112.52"),
			Secondary: pulumi.String("208.91.112.54"),
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
    var fortiosProvider = new Fortios.Provider("fortios-provider", new()
    {
        FmgCabundlefile = "/path/yourCA.crt",
        FmgHostname = "192.168.88.100",
        FmgInsecure = false,
        FmgPasswd = "admin",
        FmgUsername = "APIUser",
    });

    var test1 = new Fortios.Fmg.SystemDns("test1", new()
    {
        Primary = "208.91.112.52",
        Secondary = "208.91.112.54",
    }, new CustomResourceOptions
    {
        Provider = fortiosProvider,
    });

});

```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
name: fortimanager
runtime: yaml

config: {}

resources:
  fortios-provider:
    type: pulumi:providers:fortios
    properties:
      fmgCabundlefile: /path/yourCA.crt
      fmgHostname: 192.168.88.100
      fmgInsecure: false
      fmgPasswd: admin
      fmgUsername: APIUser

  test1:
    type: fortios:fmg:SystemDns
    properties:
      primary: 208.91.112.52
      secondary: 208.91.112.54
    options:
      provider: ${fortios-provider}
```

{{% /choosable %}}

{{< /chooser >}}

If it is used for testing, you can set `insecure` to true and unset
`cabundlefile` to quickly set the provider up.

Please refer to the Argument Reference below for more help on `insecure` and
`cabundlefile`.

### Authentication

As the same to provider for FortiGate, the following two methods are supported:

- Static credentials
- Environment variables

#### Static credentials

Static credentials can be provided by adding the `fmg_hostname`, `fmg_username`
and `fmg_passwd` key in-line in the FortiOS provider block.

#### Environment variables

You can provide your credentials via the

- `FORTIOS_FMG_HOSTNAME`
- `FORTIOS_FMG_USERNAME`
- `FORTIOS_FMG_PASSWORD`
- `FORTIOS_FMG_INSECURE` and
- `FORTIOS_FMG_CABUNDLE`

environment variables.

> Note that setting your FortiOS credentials using static credentials variables
> will override the environment variables.

**Usage:**

```shell
$ export "FORTIOS_FMG_HOSTNAME"="192.168.88.100"
$ export "FORTIOS_FMG_USERNAME"="admin"
$ export "FORTIOS_FMG_PASSWORD"="admin"
$ export "FORTIOS_FMG_INSECURE"="false"
$ export "FORTIOS_FMG_CABUNDLE"="/path/yourCA.crt"
```

### Multi-Adom

Multi-Adom feature is supported in case of using FortiManager, just take the
following example as a reference:

{{< chooser language "typescript,python,go,csharp,yaml" >}}

{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as fortios from "@pulumiverse/fortios";

const script = new fortios.fmg.DevicemanagerScript("script", {
    name: "config-intf3",
    description: "configure interface3",
    content: `config system interface
 edit port3
	 set vdom "root"
	 set ip 10.10.0.200 255.255.0.0
	 set allowaccess ping http https
	 next
 end`,
    target: "device_database",
    adom: "test-adom",
});

const scriptExecute = new fortios.fmg.DevicemanagerScriptExecute("script", {
    scriptName: script.name,
    targetDevname: "FGVM64-test",
    adom: "test-adom",
    vdom: "root",
});

_ = new fortios.fmg.DevicemanagerInstallDevice("script", {
    targetDevname: scriptExecute.targetDevname,
    adom: "test-adom",
    vdom: "root",
});

```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi
import pulumiverse_fortios as fortios

script = fortios.fmg.DevicemanagerScript("script",
    name="config-intf3",
    description="configure interface3",
    content="""config system interface
 edit port3
	 set vdom "root"
	 set ip 10.10.0.200 255.255.0.0
	 set allowaccess ping http https
	 next
 end""",
    target="device_database",
    adom="test-adom")

script_execute = fortios.fmg.DevicemanagerScriptExecute("script",
    script_name=script.name,
    target_devname="FGVM64-test",
    adom="test-adom",
    vdom="root")

_ = fortios.fmg.DevicemanagerInstallDevice("script",
    target_devname=script_execute.target_devname,
    adom="test-adom",
    vdom="root")

```

{{% /choosable %}}

{{% choosable language go %}}

```go
package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/fmg"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		script, err := fmg.NewDevicemanagerScript(ctx, "script", &fmg.DevicemanagerScriptArgs{
			Name:        pulumi.String("config-intf3"),
			Description: pulumi.String("configure interface3"),
			Content: pulumi.String(`config system interface
 edit port3
	 set vdom "root"
	 set ip 10.10.0.200 255.255.0.0
	 set allowaccess ping http https
	 next
 end`),
			Target: pulumi.String("device_database"),
			Adom:   pulumi.String("test-adom"),
		})
		if err != nil {
			return err
		}
		scriptExecute, err := fmg.NewDevicemanagerScriptExecute(ctx, "script", &fmg.DevicemanagerScriptExecuteArgs{
			ScriptName:    script.Name,
			TargetDevname: pulumi.String("FGVM64-test"),
			Adom:          pulumi.String("test-adom"),
			Vdom:          pulumi.String("root"),
		})
		if err != nil {
			return err
		}
		_, err = fmg.NewDevicemanagerInstallDevice(ctx, "script", &fmg.DevicemanagerInstallDeviceArgs{
			TargetDevname: scriptExecute.TargetDevname,
			Adom:          pulumi.String("test-adom"),
			Vdom:          pulumi.String("root"),
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

```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Fortios = Pulumiverse.Fortios;

return await Deployment.RunAsync(() =>
{
    var script = new Fortios.Fmg.DevicemanagerScript("script", new()
    {
        Name = "config-intf3",
        Description = "configure interface3",
        Content = @"config system interface
 edit port3
	 set vdom ""root""
	 set ip 10.10.0.200 255.255.0.0
	 set allowaccess ping http https
	 next
 end",
        Target = "device_database",
        Adom = "test-adom",
    });

    var scriptExecute = new Fortios.Fmg.DevicemanagerScriptExecute("script", new()
    {
        ScriptName = script.Name,
        TargetDevname = "FGVM64-test",
        Adom = "test-adom",
        Vdom = "root",
    });

    var installDevice = new Fortios.Fmg.DevicemanagerInstallDevice("script", new()
    {
        TargetDevname = scriptExecute.TargetDevname,
        Adom = "test-adom",
        Vdom = "root",
    });

});
```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
name: terraform
runtime: yaml

config:
  fortios:fmgHostname:
    value: 192.168.88.200
  fortios:fmgInsecure:
    value: true
  fortios:fmgPasswd:
    value: admin
  fortios:fmgUsername:
    value: APIUser

resources:
  script:
    type: fortios:fmg:DevicemanagerScript
    properties:
      name: config-intf3
      description: configure interface3
      content: <|
        config system interface
        edit port3
          set vdom "root"
          set ip 10.10.0.200 255.255.0.0
          set allowaccess ping http https
          next
        end
      target: device_database
      adom: test-adom

  script-execute:
    type: fortios:fmg:DevicemanagerScriptExecute
    properties:
      scriptName: ${script.name}
      targetDevname: FGVM64-test
      adom: test-adom
      vdom: root

  install-device:
    type: fortios:fmg:DevicemanagerInstallDevice
    properties:
      targetDevname: ${script-execute.targetDevname}
      adom: test-adom
      vdom: root
```

{{% /choosable %}}

{{< /chooser >}}

This will install the script from the FMG(test-adom) to the FGT(root).

> Note that one resource supports Multi-Adom feature if it has 'adom' argument.

### Argument Reference

The following arguments are supported:

- `fmg_hostname` - (Optional) The hostname or IP address of FortiManager. It
  must be provided, but it can also be sourced from the `FORTIOS_FMG_HOSTNAME`
  environment variable.

- `fmg_username` - (Optional) The username of FortiManager. It must be provided,
  but it can also be sourced from the `FORTIOS_FMG_USERNAME` environment
  variable.

- `fmg_passwd` - (Optional) The password of FortiManager, it can also be sourced
  from the `FORTIOS_FMG_PASSWORD` environment variable.

- `fmg_insecure` - (Optional) Control whether the Provider to perform insecure
  SSL requests. If omitted, the `FORTIOS_FMG_INSECURE` environment variable is
  used. If neither is set, default value is `false`.

- `fmg_cabundlefile` - (Optional) The path of a custom CA bundle file. You can
  specify a path to the file, or you can specify it by the
  `FORTIOS_FMG_CABUNDLE` environment variable.

## Generate an API token for FortiOS

FortiOS Provider requires an API token to be authenticated. An API token is
generated by creating a new REST API admin.

> Note: The steps below must be performed from the FortiGate GUI as an
> Administrator with the super_admin profile and are customized for this
> tutorial

- Step 1: Determine your Source Address

  The source address is needed to ensure the API token can only be used from
  trusted hosts. This step can be skipped if the trusted host IP address is
  already known.

  On the FortiGate GUI, select the `Status` dashboard and locate the `Administrators` widget.

  Click `your-userid > Show active administrator sessions`.

  Make note of the `Source Address` for `your-userid` as it will be needed to create the Trusted Host in `Step 3: Create the REST API Admin`.

- Step 2: Create an Administrator profile

  In this tutorial, we will create a profile that only has Read access to the
  firewall address permission group.

  On the FortiGate GUI, select `System > Admin Profiles > Create New`.
  Populate the fields according to your environment.

- Step 3: Create the REST API Admin

  Use the steps below to create the FortiOS API admin.

  On the FortiGate GUI, select `System > Administrators > Create New > REST API Admin`.

  The Trusted Host must be specified to ensure that your local host can reach
  the FortiGate. For example, to restrict requests as coming from only
  10.20.100.99, enter 10.20.100.99/32. The Trusted Host is created from the
  Source Address obtained in `Step 1: Determine your Source Address`.

  In Administrator Profile field, select profile from Step2.

  > Note: If you want
  > to configure VDOM or resources related to administrator user permissions, you
  > need to set the field with the System predefined Administrator Profile
  > `super_admin` by CLI.\*

  Click OK and an API token will be generated.

  Make note of the API token as it is only shown once and cannot be retrieved.
  It will be needed for the rest of the tutorial.

  Click Close to complete creation of the REST API Admin.

For more about the configuration of token, please refer to the `system->system
api-user` and `execute->api-user` chapters of the `FortiOS Handbook - CLI
Reference`.

## Debugging for FortiGate

Verbose output can help with debugging. To enable verbose debugging, use the following commands in the FortiGate CLI:

```shell
$ diagnose debug enable
$ diagnose debug application httpsd -1
$ diagnose debug cli 8
```

Debug messages will be displayed for 30 minutes and will include debug messages for all requests to/from the FortiOS web interface. CLI commands for any configuration changes will be logged to the console.

When the REST API for IPv4 policy traffic statistics is queried, REST API related debug messages will be similar to the following:

```
[httpsd 4137 - 1588331103     info] ap_invoke_handler[593] -- new request (handler='api_cmdb_v2-handler', uri='/api/v2/cmdb/system/global?access_token=******************************', method='PUT')
[httpsd 4137 - 1588331103     info] ap_invoke_handler[597] -- User-Agent: Go-http-client/1.1
[httpsd 4137 - 1588331103     info] ap_invoke_handler[600] -- Source: 192.168.52.1:64379 Destination: 192.168.52.177:443
[httpsd 4137 - 1588331103     info] api_cmdb_v2_handler[2048] -- received api_cmdb_v2_request from '192.168.52.1'
[httpsd 4137 - 1588331103  warning] api_access_check_for_api_key[957] -- API Key request authorized for test-apiuser from 192.168.52.1.
[httpsd 4137 - 1588331103     info] api_store_parameter[227] -- add API parameter 'access_token': '********' (type=string)
[httpsd 4137 - 1588331103     info] api_store_parameter[227] -- add API parameter 'admintimeout': '"65"' (type=string)
[httpsd 4137 - 1588331103     info] api_store_parameter[227] -- add API parameter 'timezone': '"04"' (type=string)
[httpsd 4137 - 1588331103     info] api_store_parameter[227] -- add API parameter 'hostname': '"mytestfirewall"' (type=string)
[httpsd 4137 - 1588331103     info] api_store_parameter[227] -- add API parameter 'admin-sport': '"443"' (type=string)
[httpsd 4137 - 1588331103     info] api_store_parameter[227] -- add API parameter 'admin-ssh-port': '"22"' (type=string)
[httpsd 4137 - 1588331103     info] handle_cli_req_v2_vdom[1950] -- new CMDB API request (vdom='root',user='test-apiuser')
[httpsd 4137 - 1588331103     info] api_cmdb_request_init_by_path[1351] -- new CMDB query (path='system',name='global')
[httpsd 4137 - 1588331103     info] api_return_cmdb_revision[779] -- ETag check for system.global
[httpsd 4137 - 1588331103     info] _api_cmdb_v2_config[1136] -- editing CLI object (append=0, auto_key=0, path=system, name=global, mkey=(null), flags=0)
[httpsd 4137 - 1588331103     info] api_set_cmdb_attr[1685] -- 'admintimeout': '65'
[httpsd 4137 - 1588331103     info] api_set_cmdb_attr[1685] -- 'timezone': '04'
[httpsd 4137 - 1588331103     info] api_set_cmdb_attr[1685] -- 'hostname': 'mytestfirewall'
[httpsd 4137 - 1588331103     info] api_set_cmdb_attr[1685] -- 'admin-sport': '443'
[httpsd 4137 - 1588331103     info] api_set_cmdb_attr[1685] -- 'admin-ssh-port': '22'
0: config system global
0: set hostname "mytestfirewall"
0: end
[httpsd 4137 - 1588331103     info] cmdb_save_with_children[269] -- appended main node (nret=0, is_new=0)
[httpsd 4137 - 1588331103     info] ap_invoke_handler[616] -- request completed (handler='api_cmdb_v2-handler' result==0)
[httpsd 4137 - 1588331103     info] ap_invoke_handler[593] -- new request (handler='api_cmdb_v2-handler', uri='/api/v2/cmdb/system/global?access_token=******************************', method='GET')
[httpsd 4137 - 1588331103     info] ap_invoke_handler[597] -- User-Agent: Go-http-client/1.1
[httpsd 4137 - 1588331103     info] ap_invoke_handler[600] -- Source: 192.168.52.1:64379 Destination: 192.168.52.177:443
[httpsd 4137 - 1588331103     info] api_cmdb_v2_handler[2048] -- received api_cmdb_v2_request from '192.168.52.1'
[httpsd 4137 - 1588331103  warning] api_access_check_for_api_key[957] -- API Key request authorized for test-apiuser from 192.168.52.1.
[httpsd 4137 - 1588331103     info] api_store_parameter[227] -- add API parameter 'access_token': '********' (type=string)
[httpsd 4137 - 1588331103     info] handle_cli_req_v2_vdom[1950] -- new CMDB API request (vdom='root',user='test-apiuser')
[httpsd 4137 - 1588331103     info] api_cmdb_request_init_by_path[1351] -- new CMDB query (path='system',name='global')
[httpsd 4137 - 1588331103     info] api_cmdb_select_etag[2062] -- ETag check for system.global
[httpsd 4137 - 1588331103     info] api_return_cmdb_revision[779] -- ETag check for system.global
[httpsd 4137 - 1588331103     info] api_add_etag[864] -- no If-None-Match header
[httpsd 4137 - 1588331103     info] api_return_cmdb_revision[779] -- ETag check for system.global
[httpsd 4137 - 1588331103     info] ap_invoke_handler[616] -- request completed (handler='api_cmdb_v2-handler' result==0)
write config file success, prepare to save in flash
[__create_file_new_version:257] the new version config file '/data/./config/sys_global.conf.gz.v000000172' is created
[symlink_config_file:324] a new version of '/data/./config/sys_global.conf.gz' is created: /data/./config/sys_global.conf.gz.v000000172
[symlink_config_file:367] the old version '/data/./config/sys_global.conf.gz.v000000171' is deleted
[symlink_config_file:370] '/data/./config/sys_global.conf.gz' has been symlink'ed to the new version '/data/./config/sys_global.conf.gz.v000000172'. The old version '/data/./config/sys_global.conf.gz.v000000171' has been deleted
zip config file /data/./config/sys_global.conf.gz success!
```

## Filter results of datasource

Filter results of the list type datasources.

### Syntax

**Syntax**: `[key][operator][pattern]`

The following filter operators are supported:

- `==` - Pattern must be identical to the value, case-sensitive.
- `=*` - Pattern must be identical to the value, case-insensitive.
- `!=` - Pattern does not match the value, case-sensitive.
- `!*` - Pattern does not match the value, case-insensitive.
- `=@` - Pattern found within value, case-insensitive.
- `!@` - Pattern not found within value, case-insensitive.
- `<=` - Value must be less than or equal to pattern.
- `<` - Value must be less than pattern.
- `>=` - Value must be greater than or equal to pattern.
- `>` - Value must be greater than pattern.

#### Examples:

```typescript
import * as fortios from "@pulumiverse/fortios";

const sample1 = fortios.firewall.getPolicylistOutput({
    filter: "schedule==always",
});

const sample2 = fortios.system.getInterfacelistOutput({
    filter: "name!=port1",
});

export const sample1Output = sample1.apply(sample1 => sample1.policyidlists);
export const sample2Output = sample2.apply(sample2 => sample2.namelists);

```

### Combination

To create a complex query, filters can be combined as follows:

- `Logical OR` - Separate filters using commas `,`.
- `Logical AND` - Separate filters using ampersands `&`.
- `Combining AND and OR` - Separate filters using commas `,` and ampersands `&`.

#### Examples:

```typescript
import * as fortios from "@pulumiverse/fortios";

const sample3 = fortios.firewall.getPolicylistOutput({
    filter: "schedule==always,schedule==once",
});
const sample4 = fortios.firewall.getPolicylistOutput({
    filter: "schedule==always&action==accept",
});
const sample5 = fortios.firewall.getPolicylistOutput({
    filter: "schedule==always&action==accept,action==deny",
});
const sample6Var = fortios.firewall.getPolicylistOutput({
    filter: "policyid>3",
});
const sample7Var = fortios.firewall.getAddresslistOutput({
    filter: "name!=google-play",
});

export const sample3Output = sample3.apply(sample3 => sample3.policyidlists);
export const sample4Output = sample4.apply(sample4 => sample4.policyidlists);
export const sample5Output = sample5.apply(sample5 => sample5.policyidlists);
export const sample6 = sample6Var.apply(sample6Var => sample6Var.policyidlists);
export const sample7 = sample7Var.apply(sample7Var => sample7Var.namelists);

const grp1 = new fortios.firewall.Addrgrp("grp1", {
    name: "addrgrptest1",
    members: sample7Var.apply(sample7Var => sample7Var.namelists.map(name => {
        return {
            name: name,
        }
    }))
})
```

## Change the admin default port to the custom port

To access the FortiGate via api user, port 443 is used for HTTPS (by default).
If the port is changed or intended to be changed, refer to the details below
(Let us assume that the customized port is 8443):

### Option I

1. Configure the Firewall's admin_sport to 8443 manually.

1. Configure the Provider part of the configuration file, add the 8443 port
number to the hostname/IP and separate them with a colon, as follows in the
`Pulumi.yaml` file:

```yaml
name: Fortios
runtime: yaml
config:
  fortios:hostname:
    value: 192.168.52.111:8443
  fortios:token:
    value: nx6nbGn8tnFddaa3Qy79jpjfsyw1
```

### Option II

1. Configure the admin_sport port as follows:

    Pulumi.yaml:

    ```yaml
    name: Fortios
    runtime: yaml
    config:
      fortios:hostname:
        value: 192.168.52.111
      fortios:insecure:
        value: "true"
      fortios:token:
        value: nx6nbGn8tnFddaa3Qy79jpjfsyw1
    ```

    Language dependent resource:

    {{< chooser language "typescript,python,go,csharp,yaml" >}}

    {{% choosable language typescript %}}

    ```typescript
    import * as pulumi from "@pulumi/pulumi";
    import * as fortios from "@pulumiverse/fortios";

    const fglobal = new fortios.system.Global("fglobal", {
        admintimeout: 240,
        hostname: "testhostname",
        timezone: "33",
        adminSport: 8443,
        adminSshPort: 22,
    });
    ```

    {{% /choosable %}}

    {{% choosable language python %}}

    ```python
    import pulumi
    import pulumiverse_fortios as fortios

    fglobal = fortios.system.Global("fglobal",
        admintimeout=240,
        hostname="testhostname",
        timezone="33",
        admin_sport=8443,
        admin_ssh_port=22)
    ```

    {{% /choosable %}}

    {{% choosable language go %}}

    ```go
    package main

    import (
      "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
      "github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/system"
    )

    func main() {
      pulumi.Run(func(ctx *pulumi.Context) error {
        _, err := system.NewGlobal(ctx, "fglobal", &system.GlobalArgs{
          Admintimeout: pulumi.Int(240),
          Hostname:     pulumi.String("testhostname"),
          Timezone:     pulumi.String("33"),
          AdminSport:   pulumi.Int(8443),
          AdminSshPort: pulumi.Int(22),
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

    ```csharp
    using System.Collections.Generic;
    using System.Linq;
    using Pulumi;
    using Fortios = Pulumiverse.Fortios;

    return await Deployment.RunAsync(() => 
    {
        var fglobal = new Fortios.System.Global("fglobal", new()
        {
            Admintimeout = 240,
            Hostname = "testhostname",
            Timezone = "33",
            AdminSport = 8443,
            AdminSshPort = 22,
        });

    });
    ```

    {{% /choosable %}}

    {{% choosable language yaml %}}

    ```yaml
    resources:
      fglobal:
        type: fortios:system:Global
        properties:
          admintimeout: 240
          hostname: testhostname
          timezone: '33'
          adminSport: 8443
          adminSshPort: 22
    ```

    {{% /choosable %}}

    {{< /chooser >}}

    Then execute "Pulumi init; Pulumi plan; Pulumi apply". After a few seconds
    or more (depending on your network situation), then press `Ctrl+C` to end
    the apply command. If your network is okay, the `adminSport` will be
    successfully set to `8443`.

2. Change the Provider part of the configuration file `Pulumi.yaml`, add the
   8443 port number to the hostname and separate them with a colon, as follows,
   then re-apply Pulumi:

    ```yaml
    name: Fortios
    runtime: yaml
    config:
      fortios:hostname:
        value: 192.168.52.111:8443
      fortios:insecure:
        value: "true"
      fortios:token:
        value: nx6nbGn8tnFddaa3Qy79jpjfsyw1
    ```

## Sort security policies

Pulumi is a parallel system, that means when Pulumi walks the dependency
tree, it will create as many resources in parallel as it can, so Pulumi can
figure out the most efficient way to make it happen. We can make resources be
submitted to the device in order with the help of Pulumi's dependsOn
feature in combination with [ComponentResource](https://www.pulumi.com/docs/concepts/resources/components/).

For example, let's suppose there are the following ComponentResources:

`ComponentResource` m1 will create policies with policyid 1,2,3,
`ComponentResource` m2 will create policies 4,5,6, and `ComponentResource` m3
will create policies 7,8,9.

If we want them to be submitted to the device in the order of 1, 2, 3, 4, 5, 6,
7, 8, 9, the configuration is as follows:

```typescript
import * as pulumi from '@pulumi/pulumi'
import * as fortios from '@pulumi/fortios'

class M1 extends pulumi.ComponentResource {
  constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
    super('components:index:M1', name, {}, opts)
    const trname1 = new fortios.index.FirewallPolicy(
      `${name}-trname1`,
      {
        action: 'accept',
        name: 'policy1',
        policyid: 1,
        status: 'enable',
        dstaddr: [
          {
            name: 'all',
          },
        ],
        dstintf: [
          {
            name: 'port4',
          },
        ],
        service: [
          {
            name: 'HTTP',
          },
        ],
        srcaddr: [
          {
            name: 'all',
          },
        ],
        srcintf: [
          {
            name: 'port3',
          },
        ],
      },
      {
        parent: this,
      },
    )

    const trname2 = new fortios.index.FirewallPolicy(
      `${name}-trname2`,
      {
        action: 'accept',
        name: 'policy2',
        policyid: 2,
        status: 'enable',
        dstaddr: [
          {
            name: 'swscan.apple.com',
          },
        ],
        dstintf: [
          {
            name: 'port2',
          },
        ],
        service: [
          {
            name: 'AFS3',
          },
        ],
        srcaddr: [
          {
            name: 'myaddress',
          },
        ],
        srcintf: [
          {
            name: 'port1',
          },
        ],
      },
      {
        parent: this,
      },
    )

    const trname3 = new fortios.index.FirewallPolicy(
      `${name}-trname3`,
      {
        action: 'accept',
        name: 'policy3',
        policyid: 3,
        status: 'enable',
        dstaddr: [
          {
            name: 'all',
          },
        ],
        dstintf: [
          {
            name: 'port2',
          },
        ],
        service: [
          {
            name: 'ALL_TCP',
          },
        ],
        srcaddr: [
          {
            name: 'myaddress',
          },
        ],
        srcintf: [
          {
            name: 'port1',
          },
        ],
      },
      {
        parent: this,
      },
    )
  }
}

class M2 extends pulumi.ComponentResource {
  constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
    super('components:index:M2', name, {}, opts)
    const trname4 = new fortios.index.FirewallPolicy(
      `${name}-trname4`,
      {
        action: 'accept',
        name: 'policy4',
        policyid: 4,
        status: 'enable',
        dstaddr: [
          {
            name: 'swscan.apple.com',
          },
        ],
        dstintf: [
          {
            name: 'port2',
          },
        ],
        service: [
          {
            name: 'ALL_ICMP',
          },
        ],
        srcaddr: [
          {
            name: 'google-play',
          },
        ],
        srcintf: [
          {
            name: 'port1',
          },
        ],
      },
      {
        parent: this,
      },
    )

    const trname5 = new fortios.index.FirewallPolicy(
      `${name}-trname5`,
      {
        action: 'accept',
        name: 'policy5',
        policyid: 5,
        status: 'enable',
        dstaddr: [
          {
            name: 'autoupdate.opera.com',
          },
        ],
        dstintf: [
          {
            name: 'port3',
          },
        ],
        service: [
          {
            name: 'ALL_TCP',
          },
        ],
        srcaddr: [
          {
            name: 'myaddress',
          },
        ],
        srcintf: [
          {
            name: 'port1',
          },
        ],
      },
      {
        parent: this,
      },
    )

    const trname6 = new fortios.index.FirewallPolicy(
      `${name}-trname6`,
      {
        action: 'accept',
        name: 'policy6',
        policyid: 6,
        status: 'enable',
        dstaddr: [
          {
            name: 'all',
          },
        ],
        dstintf: [
          {
            name: 'port2',
          },
        ],
        service: [
          {
            name: 'ALL_TCP',
          },
        ],
        srcaddr: [
          {
            name: 'autoupdate.opera.com',
          },
        ],
        srcintf: [
          {
            name: 'port2',
          },
        ],
      },
      {
        parent: this,
      },
    )
  }
}

class M3 extends pulumi.ComponentResource {
  constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
    super('components:index:M3', name, {}, opts)
    const trname7 = new fortios.index.FirewallPolicy(
      `${name}-trname7`,
      {
        action: 'accept',
        name: 'policy7',
        policyid: 7,
        status: 'enable',
        dstaddr: [
          {
            name: 'swscan.apple.com',
          },
        ],
        dstintf: [
          {
            name: 'port2',
          },
        ],
        service: [
          {
            name: 'ALL_TCP',
          },
        ],
        srcaddr: [
          {
            name: 'google-play',
          },
        ],
        srcintf: [
          {
            name: 'port1',
          },
        ],
      },
      {
        parent: this,
      },
    )

    const trname8 = new fortios.index.FirewallPolicy(
      `${name}-trname8`,
      {
        action: 'accept',
        name: 'policy8',
        policyid: 8,
        status: 'enable',
        dstaddr: [
          {
            name: 'autoupdate.opera.com',
          },
        ],
        dstintf: [
          {
            name: 'port3',
          },
        ],
        service: [
          {
            name: 'ALL_ICMP',
          },
        ],
        srcaddr: [
          {
            name: 'myaddress',
          },
        ],
        srcintf: [
          {
            name: 'port1',
          },
        ],
      },
      {
        parent: this,
      },
    )

    const trname9 = new fortios.index.FirewallPolicy(
      `${name}-trname9`,
      {
        action: 'accept',
        name: 'policy9',
        policyid: 9,
        status: 'enable',
        dstaddr: [
          {
            name: 'all',
          },
        ],
        dstintf: [
          {
            name: 'port2',
          },
        ],
        service: [
          {
            name: 'ALL_ICMP',
          },
        ],
        srcaddr: [
          {
            name: 'autoupdate.opera.com',
          },
        ],
        srcintf: [
          {
            name: 'port2',
          },
        ],
      },
      {
        parent: this,
      },
    )
  }
}

/*****
 * MAIN
 */
const m1 = new M1('m1')
const m2 = new M2('m2', {
  dependsOn: [m1],
})
const m3 = new M3('m3', {
  dependsOn: [m2],
})

```

Note the `dependsOn` in the `ComponentResource` configuration.

We will find that these policies have been submitted to the device in the order
of 1, 2, 3, 4, 5, 6, 7, 8, 9 after executing `pulumi up`.

> You find the complete (runnable) example in the
> [examples/sorting_polcies/typescript](https://github.com/pulumiverse/pulumi-fortios/tree/main/examples/sorting_policies/typescript) folder.
