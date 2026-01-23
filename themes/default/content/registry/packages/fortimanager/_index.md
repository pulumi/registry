---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/fortinetdev/fortimanager/1.15.0/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Fortimanager Provider
meta_desc: Provides an overview on how to configure the Pulumi Fortimanager provider.
layout: package
---

## Generate Provider

The Fortimanager provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider fortinetdev/fortimanager
```
## Overview

The FortiManager provider is used to interact with the resources supported by FortiManager. We need to configure the provider with the proper credentials before it can be used. Please use the navigation on the left to read more details about the available resources.
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    fortimanager:adom:
        value: root
    fortimanager:cabundlefile:
        value: /path/yourCA.crt
    fortimanager:hostname:
        value: 192.168.52.178
    fortimanager:insecure:
        value: "false"
    fortimanager:password:
        value: admin
    fortimanager:scopetype:
        value: adom
    fortimanager:username:
        value: admin

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as fortimanager from "@pulumi/fortimanager";

// Create a firewall vip object
const trname1 = new fortimanager.ObjectFirewallVip("trname1", {
    scopetype: "inherit",
    adom: "root",
    extintf: "any",
    extip: "1.1.1.1-2.1.1.1",
    mappedips: ["12.1.1.1-13.1.1.1"],
    name: "viptest",
});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    fortimanager:adom:
        value: root
    fortimanager:cabundlefile:
        value: /path/yourCA.crt
    fortimanager:hostname:
        value: 192.168.52.178
    fortimanager:insecure:
        value: "false"
    fortimanager:password:
        value: admin
    fortimanager:scopetype:
        value: adom
    fortimanager:username:
        value: admin

```
```python
import pulumi
import pulumi_fortimanager as fortimanager

# Create a firewall vip object
trname1 = fortimanager.ObjectFirewallVip("trname1",
    scopetype="inherit",
    adom="root",
    extintf="any",
    extip="1.1.1.1-2.1.1.1",
    mappedips=["12.1.1.1-13.1.1.1"],
    name="viptest")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    fortimanager:adom:
        value: root
    fortimanager:cabundlefile:
        value: /path/yourCA.crt
    fortimanager:hostname:
        value: 192.168.52.178
    fortimanager:insecure:
        value: "false"
    fortimanager:password:
        value: admin
    fortimanager:scopetype:
        value: adom
    fortimanager:username:
        value: admin

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Fortimanager = Pulumi.Fortimanager;

return await Deployment.RunAsync(() =>
{
    // Create a firewall vip object
    var trname1 = new Fortimanager.ObjectFirewallVip("trname1", new()
    {
        Scopetype = "inherit",
        Adom = "root",
        Extintf = "any",
        Extip = "1.1.1.1-2.1.1.1",
        Mappedips = new[]
        {
            "12.1.1.1-13.1.1.1",
        },
        Name = "viptest",
    });

});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    fortimanager:adom:
        value: root
    fortimanager:cabundlefile:
        value: /path/yourCA.crt
    fortimanager:hostname:
        value: 192.168.52.178
    fortimanager:insecure:
        value: "false"
    fortimanager:password:
        value: admin
    fortimanager:scopetype:
        value: adom
    fortimanager:username:
        value: admin

```
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/fortimanager/fortimanager"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a firewall vip object
		_, err := fortimanager.NewObjectFirewallVip(ctx, "trname1", &fortimanager.ObjectFirewallVipArgs{
			Scopetype: pulumi.String("inherit"),
			Adom:      pulumi.String("root"),
			Extintf:   pulumi.String("any"),
			Extip:     pulumi.String("1.1.1.1-2.1.1.1"),
			Mappedips: pulumi.StringArray{
				pulumi.String("12.1.1.1-13.1.1.1"),
			},
			Name: pulumi.String("viptest"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
```
{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
config:
    fortimanager:adom:
        value: root
    fortimanager:cabundlefile:
        value: /path/yourCA.crt
    fortimanager:hostname:
        value: 192.168.52.178
    fortimanager:insecure:
        value: "false"
    fortimanager:password:
        value: admin
    fortimanager:scopetype:
        value: adom
    fortimanager:username:
        value: admin

```
```yaml
resources:
  # Create a firewall vip object
  trname1:
    type: fortimanager:ObjectFirewallVip
    properties:
      scopetype: inherit
      adom: root
      extintf: any
      extip: 1.1.1.1-2.1.1.1
      mappedips:
        - 12.1.1.1-13.1.1.1
      name: viptest
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    fortimanager:adom:
        value: root
    fortimanager:cabundlefile:
        value: /path/yourCA.crt
    fortimanager:hostname:
        value: 192.168.52.178
    fortimanager:insecure:
        value: "false"
    fortimanager:password:
        value: admin
    fortimanager:scopetype:
        value: adom
    fortimanager:username:
        value: admin

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.fortimanager.ObjectFirewallVip;
import com.pulumi.fortimanager.ObjectFirewallVipArgs;
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
        // Create a firewall vip object
        var trname1 = new ObjectFirewallVip("trname1", ObjectFirewallVipArgs.builder()
            .scopetype("inherit")
            .adom("root")
            .extintf("any")
            .extip("1.1.1.1-2.1.1.1")
            .mappedips("12.1.1.1-13.1.1.1")
            .name("viptest")
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}

Before using this provider, the permission level for rpc-permit need to be set. See `Guides->To Set the Permission Level for RPC-Permit` for details.

If it is used for testing, you can set `insecure` to "true" and unset `cabundlefile` to quickly set the provider up, for example:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    fortimanager:adom:
        value: root
    fortimanager:hostname:
        value: 192.168.52.178
    fortimanager:insecure:
        value: "true"
    fortimanager:password:
        value: admin
    fortimanager:scopetype:
        value: adom
    fortimanager:username:
        value: admin

```

Please refer to the Configuration Reference below for more help on `insecure` and `cabundlefile`.
## Authentication

The FortiManager provider offers a means of providing credentials for authentication. The following methods are supported:

- Static credentials
- Environment variables
### Static credentials

Static credentials can be provided by adding credential keys in-line in the FortiManager provider configuration.

There are two kinds of credentials supported for on-prem FortiManager.
- `token` based authentication (Recommanded). User needs to generate an API token from FortiManager. *Note: Only FortiManager version >= v7.2.2 supports Token based authentication.*
- `username/password` authentication. User provide the username and password of the administrator.

There are two kinds of credentials supported for FortiManager Cloud.
- Provide `fmgCloudToken` directly. User needs to generate an FortiCloud token. *Note: The Token could be expired. Make sure the Token provided is valid.*
- `username/password` authentication. User provide the username and password of the FortiCloud API user. The provider will generate the FortiCloud token based on username/password.

Usage:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    fortimanager:adom:
        value: root
    fortimanager:hostname:
        value: 192.168.52.178
    fortimanager:insecure:
        value: "true"
    fortimanager:scopetype:
        value: adom
    fortimanager:token:
        value: 4ktitbsdiuye6ja43aoxafuqcb15kzse

```
#### Generate an API token for FortiManager

See the left navigation: `Guides` > `Generate an API token for FortiManager`.
#### Create FortiCloud API user and generate FortiCloud token for FortiManager Cloud

See the left navigation: `Guides` > `Generate an FortiCloud token for FortiManager Cloud`.
### Environment variables

You can provide your credentials via the `FORTIMANAGER_ACCESS_HOSTNAME`, `FORTIMANAGER_ACCESS_TOKEN`, `FORTIMANAGER_CLOUD_ACCESS_TOKEN`,`FORTIMANAGER_ACCESS_USERNAME`, `FORTIMANAGER_ACCESS_PASSWORD`, `FORTIMANAGER_INSECURE` and `FORTIMANAGER_CA_CABUNDLE` environment variables. Note that setting your FortiManager credentials using static credentials variables will override the environment variables.

Usage:

```shell
$ export "FORTIMANAGER_ACCESS_HOSTNAME"="192.168.52.178"
$ export "FORTIMANAGER_ACCESS_USERNAME"="admin"
$ export "FORTIMANAGER_ACCESS_PASSWORD"="admin"
$ export "FORTIMANAGER_ACCESS_TOKEN"="4ktitbsdiuye6ja43aoxafuqcb15kzse"
$ export "FORTIMANAGER_INSECURE"="false"
$ export "FORTIMANAGER_CA_CABUNDLE"="/path/yourCA.crt"
```

Then configure the FortiManager Provider as following:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    fortimanager:adom:
        value: root
    fortimanager:scopetype:
        value: adom

```
## Configuration Reference

The following configuration inputs are supported:

* `hostname` - (String | Optional) The hostname or IP address of FortiManager unit. It must be provided, but it can also be sourced from the `FORTIMANAGER_ACCESS_HOSTNAME` environment variable.

* `token` - (String | Optional) The token of FortiManager unit. If omitted, the `FORTIMANAGER_ACCESS_TOKEN` environment variable will be used. If neither is set, username/password will be used.

* `fmgCloudToken` - (String | Optional) The access token of FortiManager Cloud. If omitted, the `FORTIMANAGER_CLOUD_ACCESS_TOKEN` environment variable will be used. If neither is set, username/password will be used. Available only when `fmgType` set to `forticloud`.

* `fmgType` - (String | Optional) FortiManager type. Valid values: `on-prem`, `forticloud`. Default is `on-prem`. Set to `forticloud` if using FortiManager Cloud under FortiCloud.

* `username` - (String | Optional) FortiManager username if `fmgType` is `on-prem`, and it is required. FortiCloud API username if `fmgType` is `forticloud`, and it is optional. It can also be sourced from the `FORTIMANAGER_ACCESS_USERNAME` environment variable.

* `password` - (String | Optional) FortiManager password if `fmgType` is `on-prem`, and it is required. FortiCloud API password if `fmgType` is `forticloud`, and it is optional. It can also be sourced from the `FORTIMANAGER_ACCESS_PASSWORD` environment variable.

* `insecure` - (Bool | Optional) Control whether the Provider to perform insecure SSL requests. If omitted, the `FORTIMANAGER_INSECURE` environment variable is used. If neither is set, default value is `false`.

* `cabundlefile` - (String | Optional) The path of a custom CA bundle file. You can specify a path to the file, or you can specify it by the `FORTIMANAGER_CA_CABUNDLE` environment variable.

* `scopetype` - (String | Optional) The option is used to set the default scope of application of those resources managed by the provider. Valid values: `adom`, `global`. The default value is `adom`. Each resource can also set its own scope as needed, see the description of each resource for details.

* `adom` - (String | Optional) Adom. This value is valid only when the `scopetype` is set to `adom`. The option is used to set the default adom of the resources managed by the provider. The default value is `root`. Each resource can also set its own adom as needed, see the description of each resource for details.

* `importOptions` - (List | Optional) This parameter is only used for import in some special cases. When the resource to be imported includes pkg parameter, you need to assign a value to the parameter here, for example:

  ```yaml
  ```
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
fortimanager:adom:
value: root
fortimanager:hostname:
value: 192.168.52.178
fortimanager:importOptions:
value:
- pkg=default
fortimanager:insecure:
value: "true"
fortimanager:password:
value: admin
fortimanager:scopetype:
value: adom
fortimanager:username:
value: admin

```
``

* `logsession` - (Bool | Optional) Save the session to a local file. Used to assist fortimanager_exec_workspace_action resource. Valid values: `true`: log to file, `false`: do not log to file. Default is `false`. See `Guides -> To Lock for Restricting Configuration Changes` for details.

* `presession` - (String | Optional) The session saved earlier and within the validity period, used to reuse the previous session and assist fortimanager_exec_workspace_action resource. See `Guides -> To Lock for Restricting Configuration Changes` for details. Default is empty.

* `clean_session` - (Bool | Optional) Whether clean sessions. **Only works on workspace mode set to `disabled` if using resource fortimanager_exec_workspace_action.** If set to `true`, the provider will generate and logout the session for each HTTPS request. If set to `false`, the provider will generate a session for each Pulumi operation. But the session will stay exist until it expires. Default is `false`.

* `workspace_mode` - (String | Optional) FortiManger workspace mode. If set to `normal`, FortiManager will be locked/unlocked for each resource operation. Conflict with resource `fortimanager_exec_workspace_action`. Valid opeions: `normal`, `disabled`. Default is `disabled`.
## Release
Check out the FortiManager provider release notes and additional information from: the FortiManager provider releases.
## FortiManager best practices

FortiManager is an integrated platform for the centralized management of products in a Fortinet security infrastructure, including FortiGates.

Once FortiGates are managed by a FortiManager that is operating in Normal Mode, whenever possible, configuration changes should be made on the FortiManager and not the FortiGate. This recommendation also applies when the configuration of FortiGates and FortiManager is executed through the FortiOS and FortiManager providers.

To help you get the most out of your FortiManager products, maximize performance, and avoid potential problems, please refer to the [FortiManager documentation](https://docs.fortinet.com/product/fortimanager), including the Administration Guide and Best Practices documents.

Fortinet also provides a developer community to help administrators and advanced users enhance and increase the effectiveness of Fortinet products. The [Fortinet Developer Network (FNDN)](https://fndn.fortinet.net/) provides the official documentation and advanced tools for developing custom solutions using Fortinet products.
## Versioning

The provider can cover FortiManager 6.4 to 7.4 versions, the configuration of all parameters should be based on the relevant FortiManager version manual and FortiManager API guide.
```