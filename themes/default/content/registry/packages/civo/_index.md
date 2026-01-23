---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/civo/civo/1.1.7/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Civo Provider
meta_desc: Provides an overview on how to configure the Pulumi Civo provider.
layout: package
---

## Generate Provider

The Civo provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider civo/civo
```
~> **NOTE:** This provider was previously published as @pulumi/civo.
However, that package is no longer being updated.Going forward, it is available as a [Local Package](https://www.pulumi.com/blog/any-terraform-provider/) instead.
Please see the [provider's repository](https://github.com/pulumi/pulumi-civo) for details.

## Overview

The Civo provider is used to interact with the resources supported by Civo.

Use the navigation to the left to read about the available resources.
## Configuration

The provider will use the credentials of the [Civo CLI](https://github.com/civo/cli) (stored in ` ~/.civo.json`) if no other credentials have been set up. The provider will use credentials in the following order:

1. Environment variable (`CIVO_TOKEN`).
2. Token provided via a credentials file (See `credentialsFile` input below)
3. [CLI](https://github.com/civo/cli) configuration (`~/.civo.json`)

That means that if the `CIVO_TOKEN` variable is set, all other credentials will be ignored, and if the `credentialsFile` is set, that will be used over the CLI credentials.
### Obtaining a token

First you will need to create a [Civo Account](https://dashboard.civo.com/signup) and then you can do the following:

* If you can want to configure your credentials with the CLI, instructions are [here](https://www.civo.com/docs/overview/civo-cli#add-an-api-key-to-civo-cli)
* To fetch an API key go to the [security section](https://dashboard.civo.com/security) on the dashboard
### Using the CIVO_TOKEN variable

To use the Civo token, export the variable containing your token:

```bash
export CIVO_TOKEN=<your token>
```
### Using a credentials file

The format of the credentials file is as follows:

```json
{
	"apikeys": {
		"tf_key": "write-your-token-here"
	},
	"meta": {
		"current_apikey": "tf_key"
	}
}
```

you will then need to configure the `credentialsFile` input to the correct location, for example:

`credentialsFile = "/secure/path/civo.json"`
### Using the CLI

If you install the CLI and [configure a token](https://www.civo.com/docs/overview/civo-cli#add-an-api-key-to-civo-cli), there is nothing else you need to do if those are the credentials you wish to use, ideal for local usage.
## Example Usage
### Simplest usage

In this example the provider will look for credentials set by the CLI and use LON1 region as default, or use the environment variable `CIVO_TOKEN` if set.

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    civo:region:
        value: LON1

```
```typescript
import * as pulumi from "@pulumi/pulumi";

```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    civo:region:
        value: LON1

```
```python
import pulumi

```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    civo:region:
        value: LON1

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;

return await Deployment.RunAsync(() =>
{
});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    civo:region:
        value: LON1

```
```go
package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
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
    civo:region:
        value: LON1

```
```yaml
{}
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    civo:region:
        value: LON1

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
    }
}
```
{{% /choosable %}}
{{< /chooser >}}
### Example with credentials file

In this example we are providing a specific version of the pulumi provider and setting a credentials file to use. The credentials file will be used if the environment variable `CIVO_TOKEN` is not set.

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    civo:credentialsFile:
        value: /secure/path/civo.json
    civo:region:
        value: LON1

```
```typescript
import * as pulumi from "@pulumi/pulumi";

```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    civo:credentialsFile:
        value: /secure/path/civo.json
    civo:region:
        value: LON1

```
```python
import pulumi

```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    civo:credentialsFile:
        value: /secure/path/civo.json
    civo:region:
        value: LON1

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;

return await Deployment.RunAsync(() =>
{
});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    civo:credentialsFile:
        value: /secure/path/civo.json
    civo:region:
        value: LON1

```
```go
package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
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
    civo:credentialsFile:
        value: /secure/path/civo.json
    civo:region:
        value: LON1

```
```yaml
{}
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    civo:credentialsFile:
        value: /secure/path/civo.json
    civo:region:
        value: LON1

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
    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

- `apiEndpoint` (String) The Base URL to use for CIVO API.
- `region` (String) This sets the default region for all resources. If no default region is set, you will need to specify individually in every resource.
  <a id="credentialsFile"></a>
- `credentialsFile` (string) specify a location for a file containing your civo credentials token
- `token` (String) (**Deprecated**) for legacy reasons the user can still specify the token as an input, but in order to avoid storing that in pulumi state we have deprecated this and will be remove in future versions - don't use it.
## Configuring Modules

Pulumi modules allow you to group resources and reuse configurations efficiently. Below are examples of how to configure modules with the Civo provider.
### Example Module Usage

To use a module with the Civo provider, create a module directory and define your infrastructure in it.

**Module Definition** (`modules/civo-instance/main.tf`)

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    civo:region:
        value: 'TODO: var.region'

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as civo from "@pulumi/civo";

const config = new pulumi.Config();
const region = config.requireObject("region");
const instanceSize = config.requireObject("instanceSize");
const example = new civo.Instance("example", {
    hostname: "example-instance",
    size: instanceSize,
    region: region,
});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    civo:region:
        value: 'TODO: var.region'

```
```python
import pulumi
import pulumi_civo as civo

config = pulumi.Config()
region = config.require_object("region")
instance_size = config.require_object("instanceSize")
example = civo.Instance("example",
    hostname="example-instance",
    size=instance_size,
    region=region)
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    civo:region:
        value: 'TODO: var.region'

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Civo = Pulumi.Civo;

return await Deployment.RunAsync(() =>
{
    var config = new Config();
    var region = config.RequireObject<dynamic>("region");
    var instanceSize = config.RequireObject<dynamic>("instanceSize");
    var example = new Civo.Instance("example", new()
    {
        Hostname = "example-instance",
        Size = instanceSize,
        Region = region,
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
    civo:region:
        value: 'TODO: var.region'

```
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/civo/civo"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")
		region := cfg.RequireObject("region")
		instanceSize := cfg.RequireObject("instanceSize")
		_, err := civo.NewInstance(ctx, "example", &civo.InstanceArgs{
			Hostname: pulumi.String("example-instance"),
			Size:     pulumi.Any(instanceSize),
			Region:   pulumi.Any(region),
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
    civo:region:
        value: 'TODO: var.region'

```
```yaml
configuration:
  region:
    type: dynamic
  instanceSize:
    type: dynamic
resources:
  example:
    type: civo:Instance
    properties:
      hostname: example-instance
      size: ${instanceSize}
      region: ${region}
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    civo:region:
        value: 'TODO: var.region'

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.civo.Instance;
import com.pulumi.civo.InstanceArgs;
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
        final var region = config.get("region");
        final var instanceSize = config.get("instanceSize");
        var example = new Instance("example", InstanceArgs.builder()
            .hostname("example-instance")
            .size(instanceSize)
            .region(region)
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}

**Root Module** (`main.tf`)
```
module "civo_instance" {
    source        = "./modules/civo-instance"
    region        = "LON1"
    instance_size = "g3.small"
}
```

This approach helps in managing infrastructure more efficiently by keeping configurations modular and reusable.
**Please note since this is a non-pulumi maintained provider, every module should have its own provider configuration.**
Reference: Pulumi Modules
## Configuration Reference

`apiEndpoint` (String) The Base URL to use for CIVO API.

`region` (String) This sets the default region for all resources. If no default region is set, you will need to specify individually in every resource.

`credentialsFile` (string) Specify a location for a file containing your Civo credentials token.

`token` (String) (Deprecated) For legacy reasons, the user can still specify the token as an input, but in order to avoid storing that in Pulumi state, we have deprecated this and will remove it in future versions—don't use it.