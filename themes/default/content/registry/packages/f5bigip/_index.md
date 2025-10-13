---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-f5bigip/v3.19.1/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/pulumi/pulumi-f5bigip/blob/v3.19.1/docs/_index.md
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: F5bigip Provider
meta_desc: Provides an overview on how to configure the Pulumi F5bigip provider.
layout: package
---

## Installation

The F5bigip provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/f5bigip`](https://www.npmjs.com/package/@pulumi/f5bigip)
* Python: [`pulumi-f5bigip`](https://pypi.org/project/pulumi-f5bigip/)
* Go: [`github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip`](https://github.com/pulumi/pulumi-f5bigip)
* .NET: [`Pulumi.F5bigip`](https://www.nuget.org/packages/Pulumi.F5bigip)
* Java: [`com.pulumi/f5bigip`](https://central.sonatype.com/artifact/com.pulumi/f5bigip)

## Overview

Use the F5 BIG-IP Pulumi Provider to manage and provision your BIG-IP
configurations in Pulumi. Using BIG-IP Provider you can manage LTM(Local Traffic Manager),Network,System objects and it also supports AS3/DO integration.
### Requirements

This provider uses the iControlREST API. All the resources are validated with BigIP v12.1.1 and above.

> **NOTE** For AWAF resources, F5 BIG-IP version should be > v16.x , and ASM need to be provisioned.
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    f5bigip:address:
        value: 'TODO: var.hostname'
    f5bigip:password:
        value: 'TODO: var.password'
    f5bigip:username:
        value: 'TODO: var.username'

```
```typescript
import * as pulumi from "@pulumi/pulumi";

const config = new pulumi.Config();
const hostname = config.requireObject<any>("hostname");
const username = config.requireObject<any>("username");
const password = config.requireObject<any>("password");
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    f5bigip:address:
        value: 'TODO: var.hostname'
    f5bigip:password:
        value: 'TODO: var.password'
    f5bigip:username:
        value: 'TODO: var.username'

```
```python
import pulumi

config = pulumi.Config()
hostname = config.require_object("hostname")
username = config.require_object("username")
password = config.require_object("password")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    f5bigip:address:
        value: 'TODO: var.hostname'
    f5bigip:password:
        value: 'TODO: var.password'
    f5bigip:username:
        value: 'TODO: var.username'

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;

return await Deployment.RunAsync(() =>
{
    var config = new Config();
    var hostname = config.RequireObject<dynamic>("hostname");
    var username = config.RequireObject<dynamic>("username");
    var password = config.RequireObject<dynamic>("password");
});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    f5bigip:address:
        value: 'TODO: var.hostname'
    f5bigip:password:
        value: 'TODO: var.password'
    f5bigip:username:
        value: 'TODO: var.username'

```
```go
package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")
		hostname := cfg.RequireObject("hostname")
		username := cfg.RequireObject("username")
		password := cfg.RequireObject("password")
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
    f5bigip:address:
        value: 'TODO: var.hostname'
    f5bigip:password:
        value: 'TODO: var.password'
    f5bigip:username:
        value: 'TODO: var.username'

```
```yaml
configuration:
  hostname:
    type: dynamic
  username:
    type: dynamic
  password:
    type: dynamic
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    f5bigip:address:
        value: 'TODO: var.hostname'
    f5bigip:password:
        value: 'TODO: var.password'
    f5bigip:username:
        value: 'TODO: var.username'

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
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
        final var hostname = config.get("hostname");
        final var username = config.get("username");
        final var password = config.get("password");
    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

- `address` - (type `string`) Domain name or IP address of the BIG-IP. Can be set via the `BIGIP_HOST` environment variable.
- `username` - (type `string`) BIG-IP Username for authentication. Can be set via the `BIGIP_USER` environment variable.
- `password` - (type `string`) BIG-IP Password for authentication. Can be set via the `BIGIP_PASSWORD` environment variable.
- `tokenAuth` - (Optional, Default `true`) Enable to use token authentication. Can be set via the `BIGIP_TOKEN_AUTH` environment variable.
- `tokenValue` - (Optional) A token generated outside the provider, in place of password
- `apiTimeout` - (Optional, type `int`) A timeout for AS3 requests, represented as a number of seconds.
- `tokenTimeout` - (Optional, type `int`) A lifespan to request for the AS3 auth token, represented as a number of seconds.
- `apiRetries` - (Optional, type `int`) Amount of times to retry AS3 API requests.
- `loginRef` - (Optional,Default `tmos`) Login reference for token authentication (see BIG-IP REST docs for details). May be set via the `BIGIP_LOGIN_REF` environment variable.
- `port` - (Optional) Management Port to connect to BIG-IP,this is mainly required if we have single nic BIG-IP in AWS/Azure/GCP (or) Management port other than `443`. Can be set via `BIGIP_PORT` environment variable.
- `validateCertsDisable` - (Optional, Default `true`) If set to true, Disables TLS certificate check on BIG-IP. Can be set via the `BIGIP_VERIFY_CERT_DISABLE` environment variable.
- `trustedCertPath` - (type `string`) Provides Certificate Path to be used TLS Validate.It will be required only if `validateCertsDisable` set to `false`.Can be set via the `BIGIP_TRUSTED_CERT_PATH` environment variable.

> **Note** For BIG-IQ resources these provider credentials `address`,`username`,`password` can be set to BIG-IQ credentials.

> **Note** The F5 BIG-IP provider gathers non-identifiable usage data for the purposes of improving the product as outlined in the end user license agreement for BIG-IP. To opt out of data collection, use the following : `export TEEM_DISABLE=true`