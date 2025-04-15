---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/terraform-routeros/routeros/1.82.0/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Routeros Provider
meta_desc: Provides an overview on how to configure the Pulumi Routeros provider.
layout: package
---

## Generate Provider

The Routeros provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider terraform-routeros/routeros
```
## Overview

To get started with the provider, you first need to enable the REST API on your router. [You can follow the Mikrotik documentation on this](https://help.mikrotik.com/docs/display/ROS/REST+API), but the gist is to create an SSL cert (in `/system/certificates`) and enable the `web-ssl` service (in `/ip/services`) which uses that certificate.
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    routeros:caCertificate:
        value: /path/to/ca/certificate.pem
    routeros:hosturl:
        value: https://router.local
    routeros:insecure:
        value: true
    routeros:password:
        value: ""
    routeros:username:
        value: admin

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as routeros from "@pulumi/routeros";

const greHq = new routeros.InterfaceGre("gre_hq", {
    name: "gre-hq-1",
    remoteAddress: "10.77.3.26",
    disabled: true,
});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    routeros:caCertificate:
        value: /path/to/ca/certificate.pem
    routeros:hosturl:
        value: https://router.local
    routeros:insecure:
        value: true
    routeros:password:
        value: ""
    routeros:username:
        value: admin

```
```python
import pulumi
import pulumi_routeros as routeros

gre_hq = routeros.InterfaceGre("gre_hq",
    name="gre-hq-1",
    remote_address="10.77.3.26",
    disabled=True)
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    routeros:caCertificate:
        value: /path/to/ca/certificate.pem
    routeros:hosturl:
        value: https://router.local
    routeros:insecure:
        value: true
    routeros:password:
        value: ""
    routeros:username:
        value: admin

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Routeros = Pulumi.Routeros;

return await Deployment.RunAsync(() =>
{
    var greHq = new Routeros.InterfaceGre("gre_hq", new()
    {
        Name = "gre-hq-1",
        RemoteAddress = "10.77.3.26",
        Disabled = true,
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
    routeros:caCertificate:
        value: /path/to/ca/certificate.pem
    routeros:hosturl:
        value: https://router.local
    routeros:insecure:
        value: true
    routeros:password:
        value: ""
    routeros:username:
        value: admin

```
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/routeros/routeros"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := routeros.NewInterfaceGre(ctx, "gre_hq", &routeros.InterfaceGreArgs{
			Name:          pulumi.String("gre-hq-1"),
			RemoteAddress: pulumi.String("10.77.3.26"),
			Disabled:      pulumi.Bool(true),
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
    routeros:caCertificate:
        value: /path/to/ca/certificate.pem
    routeros:hosturl:
        value: https://router.local
    routeros:insecure:
        value: true
    routeros:password:
        value: ""
    routeros:username:
        value: admin

```
```yaml
resources:
  greHq:
    type: routeros:InterfaceGre
    name: gre_hq
    properties:
      name: gre-hq-1
      remoteAddress: 10.77.3.26
      disabled: true
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    routeros:caCertificate:
        value: /path/to/ca/certificate.pem
    routeros:hosturl:
        value: https://router.local
    routeros:insecure:
        value: true
    routeros:password:
        value: ""
    routeros:username:
        value: admin

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.routeros.InterfaceGre;
import com.pulumi.routeros.InterfaceGreArgs;
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
        var greHq = new InterfaceGre("greHq", InterfaceGreArgs.builder()
            .name("gre-hq-1")
            .remoteAddress("10.77.3.26")
            .disabled(true)
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference
### Required

- `hosturl` (String) URL of the MikroTik router, default is TLS connection to REST.
  * API: api[s]://host[:port]
    * api://router.local
    * apis://router.local:8729
  * REST: http[s]://host
    * <http://router.local>
    * <https://router.local>
    * router.local
    * 127.0.0.1

  export ROS_HOSTURL=router.local or export MIKROTIK_HOST=router.local
- `username` (String) Username for the MikroTik WEB/Winbox.

  export ROS_USERNAME=admin or export MIKROTIK_USER=admin

- `caCertificate` (String) Path to MikroTik's certificate authority file (env: ROS_CA_CERTIFICATE | MIKROTIK_CA_CERTIFICATE).
- `insecure` (Boolean) Whether to verify the SSL certificate or not (env: ROS_INSECURE | MIKROTIK_INSECURE).
- `password` (String, Sensitive) Password for the MikroTik user (env: ROS_PASSWORD | MIKROTIK_PASSWORD).
- `routerosVersion` (String) RouterOS version for which resource schemes will be adapted. The version obtained from MikroTik will be used if not specified (env: ROS_VERSION).
- `suppressSysoDelWarn` (Boolean) Suppress the system object deletion warning (env: ROS_SUPPRESS_SYSO_DEL_WARN).