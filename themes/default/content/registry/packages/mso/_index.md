---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/ciscodevnet/mso/1.5.1/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Mso Provider
meta_desc: Provides an overview on how to configure the Pulumi Mso provider.
layout: package
---

## Generate Provider

The Mso provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider ciscodevnet/mso
```
## MSO Provider

Cisco ACI Multi-Site Orchestrator (MSO) is responsible for provisioning, health monitoring, and managing the full lifecycle of Cisco ACI networking policies and tenant policies across Cisco ACI sites around the world. Pulumi provider MSO is a Pulumi plugin which will be used to manage the MSO Fabric Constructs on the Cisco MSO platform with leveraging advantages of Pulumi. The provider needs to be configured with the proper credentials before it can be used.
## Authentication

Authentication with user-id and password.
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    mso:domain:
        value: domain-name
    mso:insecure:
        value: true
    mso:password:
        value: password
    mso:platform:
        value: nd
    mso:retries:
        value: 4
    mso:url:
        value: https://173.36.219.193/
    mso:username:
        value: admin

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as mso from "@pulumi/mso";

const fooSchema = new mso.Schema("foo_schema", {
    name: "nkp12",
    templateName: "template1",
    tenantId: "5ea000bd2c000058f90a26ab",
});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    mso:domain:
        value: domain-name
    mso:insecure:
        value: true
    mso:password:
        value: password
    mso:platform:
        value: nd
    mso:retries:
        value: 4
    mso:url:
        value: https://173.36.219.193/
    mso:username:
        value: admin

```
```python
import pulumi
import pulumi_mso as mso

foo_schema = mso.Schema("foo_schema",
    name="nkp12",
    template_name="template1",
    tenant_id="5ea000bd2c000058f90a26ab")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    mso:domain:
        value: domain-name
    mso:insecure:
        value: true
    mso:password:
        value: password
    mso:platform:
        value: nd
    mso:retries:
        value: 4
    mso:url:
        value: https://173.36.219.193/
    mso:username:
        value: admin

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Mso = Pulumi.Mso;

return await Deployment.RunAsync(() =>
{
    var fooSchema = new Mso.Schema("foo_schema", new()
    {
        Name = "nkp12",
        TemplateName = "template1",
        TenantId = "5ea000bd2c000058f90a26ab",
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
    mso:domain:
        value: domain-name
    mso:insecure:
        value: true
    mso:password:
        value: password
    mso:platform:
        value: nd
    mso:retries:
        value: 4
    mso:url:
        value: https://173.36.219.193/
    mso:username:
        value: admin

```
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/mso/mso"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := mso.NewSchema(ctx, "foo_schema", &mso.SchemaArgs{
			Name:         pulumi.String("nkp12"),
			TemplateName: pulumi.String("template1"),
			TenantId:     pulumi.String("5ea000bd2c000058f90a26ab"),
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
    mso:domain:
        value: domain-name
    mso:insecure:
        value: true
    mso:password:
        value: password
    mso:platform:
        value: nd
    mso:retries:
        value: 4
    mso:url:
        value: https://173.36.219.193/
    mso:username:
        value: admin

```
```yaml
resources:
  fooSchema:
    type: mso:Schema
    name: foo_schema
    properties:
      name: nkp12
      templateName: template1
      tenantId: 5ea000bd2c000058f90a26ab
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    mso:domain:
        value: domain-name
    mso:insecure:
        value: true
    mso:password:
        value: password
    mso:platform:
        value: nd
    mso:retries:
        value: 4
    mso:url:
        value: https://173.36.219.193/
    mso:username:
        value: admin

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.mso.Schema;
import com.pulumi.mso.SchemaArgs;
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
        var fooSchema = new Schema("fooSchema", SchemaArgs.builder()
            .name("nkp12")
            .templateName("template1")
            .tenantId("5ea000bd2c000058f90a26ab")
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

Following arguments are supported with Cisco MSO pulumi provider.

* `username` - (Required) This is the Cisco MSO username, which is required to authenticate with CISCO MSO.
* `password` - (Required) Password of the user mentioned in username argument. It is required when you want to use token basedauthentication.
* `url` - (Required) URL for CISCO MSO.
* `insecure` - (Optional) This determines whether to use insecure HTTP connection or not. Default value is `true`.
* `domain`- (Optional) Name of domain. Use this parameter to provide domain name in case of using remote user with the Pulumi provider. Defaults to `Local`.
* `platform`- (Optional) Parameter is used to check the platform from which MSO is accessed. Defaults to `mso`.
* `retries` - (Optional) Number of retries for REST API calls. Defaults to `2`.