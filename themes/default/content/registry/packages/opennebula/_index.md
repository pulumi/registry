---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/opennebula/opennebula/1.4.1/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Opennebula Provider
meta_desc: Provides an overview on how to configure the Pulumi Opennebula provider.
layout: package
---

## Generate Provider

The Opennebula provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider opennebula/opennebula
```
## Overview

The OpenNebula provider is used to interact with OpenNebula cluster resources.

The provider allows you to manage your OpenNebula clusters resources.
It needs to be configured with proper credentials before it can be used.
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    opennebula:endpoint:
        value: https://example.com:2633/RPC2

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as opennebula from "@pulumi/opennebula";

const group = new opennebula.Group("group", {name: "OpenNebula"});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    opennebula:endpoint:
        value: https://example.com:2633/RPC2

```
```python
import pulumi
import pulumi_opennebula as opennebula

group = opennebula.Group("group", name="OpenNebula")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    opennebula:endpoint:
        value: https://example.com:2633/RPC2

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Opennebula = Pulumi.Opennebula;

return await Deployment.RunAsync(() =>
{
    var @group = new Opennebula.Group("group", new()
    {
        Name = "OpenNebula",
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
    opennebula:endpoint:
        value: https://example.com:2633/RPC2

```
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/opennebula/opennebula"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := opennebula.NewGroup(ctx, "group", &opennebula.GroupArgs{
			Name: pulumi.String("OpenNebula"),
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
    opennebula:endpoint:
        value: https://example.com:2633/RPC2

```
```yaml
resources:
  group:
    type: opennebula:Group
    properties:
      name: OpenNebula
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    opennebula:endpoint:
        value: https://example.com:2633/RPC2

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.opennebula.Group;
import com.pulumi.opennebula.GroupArgs;
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
        var group = new Group("group", GroupArgs.builder()
            .name("OpenNebula")
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Authentication and Configuration

The configuration of the OpenNebula Provider can be set by the provider configuration attributes or by the environment variables.
### Provider attributes

* `endpoint` - (Required) The URL of OpenNebula XML-RPC Endpoint API (for example, `http://example.com:2633/RPC2`).
* `flowEndpoint` - (Optional) The OneFlow HTTP Endpoint API (for example, `http://example.com:2474`).
* `username` - (Required) The OpenNebula username.
* `password` - (Required) The Opennebula password matching the username.
* `insecure` - (Optional) Allow insecure connexion (skip TLS verification).
* `defaultTags` - (Optional) Apply default custom tags to resources supporting `tags`. Theses tags can be overriden in the `tags` section of the resource. See Using tags below for more details.

!> **Warning:** Hard-coded credentials are not recommended in any Pulumi configuration file and should not be commited in a public repository you might prefer Environment variables instead.
### Environment variables

The provider can also read the following environment variables if no value is set in the the provider configuration attributes:

* `OPENNEBULA_ENDPOINT`
* `OPENNEBULA_FLOW_ENDPOINT`
* `OPENNEBULA_USERNAME`
* `OPENNEBULA_PASSWORD`
* `OPENNEBULA_INSECURE`
### Example

```bash
export OPENNEBULA_ENDPOINT="https://example.com:2633/RPC2"
export OPENNEBULA_FLOW_ENDPOINT="https://example.com:2474"
export OPENNEBULA_USERNAME="me"
export OPENNEBULA_PASSWORD="p@s5w0rD"
export OPENNEBULA_INSECURE="true"
```

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as opennebula from "@pulumi/opennebula";

const group = new opennebula.Group("group", {});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python

```
```python
import pulumi
import pulumi_opennebula as opennebula

group = opennebula.Group("group")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Opennebula = Pulumi.Opennebula;

return await Deployment.RunAsync(() =>
{
    var @group = new Opennebula.Group("group");

});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go

```
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/opennebula/opennebula"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := opennebula.NewGroup(ctx, "group", nil)
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

```
```yaml
resources:
  group:
    type: opennebula:Group
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.opennebula.Group;
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
        var group = new Group("group");

    }
}
```
{{% /choosable %}}
{{< /chooser >}}

```bash
pulumi up
pulumi preview
```
## Using tags
### Resource tags

Some resources can support the `tags` attribute. In consists of a map of key-value elements. In the OpenNebula language, these are called 'attributes'.
### Default tags

The provider's `defaultTags` attribute allows to set default tags for all resources supporting tags.

`defaultTags` supports the following arguments:

* `tags` - (Optional) Map of tags.

When a tag is added to a resource, it overrides the one present in `defaultTags` from the provider configuration if the key matches. In the example bellow, the resource will have `environment = "production"` and `deploymentMethod = "pulumi"` tags.

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    opennebula:endpoint:
        value: https://example.com:2633/RPC2

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as opennebula from "@pulumi/opennebula";

const group = new opennebula.Group("group", {tags: {
    environment: "production",
}});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    opennebula:endpoint:
        value: https://example.com:2633/RPC2

```
```python
import pulumi
import pulumi_opennebula as opennebula

group = opennebula.Group("group", tags={
    "environment": "production",
})
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    opennebula:endpoint:
        value: https://example.com:2633/RPC2

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Opennebula = Pulumi.Opennebula;

return await Deployment.RunAsync(() =>
{
    var @group = new Opennebula.Group("group", new()
    {
        Tags =
        {
            { "environment", "production" },
        },
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
    opennebula:endpoint:
        value: https://example.com:2633/RPC2

```
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/opennebula/opennebula"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := opennebula.NewGroup(ctx, "group", &opennebula.GroupArgs{
			Tags: pulumi.StringMap{
				"environment": pulumi.String("production"),
			},
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
    opennebula:endpoint:
        value: https://example.com:2633/RPC2

```
```yaml
resources:
  group:
    type: opennebula:Group
    properties:
      tags:
        environment: production
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    opennebula:endpoint:
        value: https://example.com:2633/RPC2

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.opennebula.Group;
import com.pulumi.opennebula.GroupArgs;
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
        var group = new Group("group", GroupArgs.builder()
            .tags(Map.of("environment", "production"))
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}