---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/glesys/glesys/0.15.0/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Glesys Provider
meta_desc: Provides an overview on how to configure the Pulumi Glesys provider.
layout: package
---

## Generate Provider

The Glesys provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider glesys/glesys
```
## Overview

The Glesys Provider is used to interact with the resources supported by GleSYS.
The provider needs to be configured with the proper credentials before it can be
used.

Use the navigation to the left to read about the available resources.
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    glesys:token:
        value: ABC123
    glesys:userid:
        value: CL12345

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as glesys from "@pulumi/glesys";

// Create a server resource
const www = new glesys.Server("www", {});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    glesys:token:
        value: ABC123
    glesys:userid:
        value: CL12345

```
```python
import pulumi
import pulumi_glesys as glesys

# Create a server resource
www = glesys.Server("www")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    glesys:token:
        value: ABC123
    glesys:userid:
        value: CL12345

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Glesys = Pulumi.Glesys;

return await Deployment.RunAsync(() =>
{
    // Create a server resource
    var www = new Glesys.Server("www");

});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    glesys:token:
        value: ABC123
    glesys:userid:
        value: CL12345

```
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/glesys/glesys"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a server resource
		_, err := glesys.NewServer(ctx, "www", nil)
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
    glesys:token:
        value: ABC123
    glesys:userid:
        value: CL12345

```
```yaml
resources:
  # Create a server resource
  www:
    type: glesys:Server
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    glesys:token:
        value: ABC123
    glesys:userid:
        value: CL12345

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.glesys.Server;
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
        // Create a server resource
        var www = new Server("www");

    }
}
```
{{% /choosable %}}
{{< /chooser >}}

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:

```

Execute in your shell:
```sh
$ pulumi up
$ pulumi preview
```
## Configuration Reference

- `apiEndpoint` (String) The base URL to use for the Glesys API requests. (Defaults to the value of the `GLESYS_API_URL` environment variable or `https://api.glesys.com` if unset.
- `token` (String) User token for the Glesys API. Alternatively, this can be set using the `GLESYS_TOKEN` environment variable
- `userid` (String) UserId for the Glesys API. Alternatively, this can be set using the `GLESYS_USERID` environment variable
## Authentication
### Static credentials

!> **Warning** Hard-coding credentials into any Pulumi configuration is not
recommended, and risks secret leakage should this file ever be committed to a
public version control system.

Static credentials can be provided by adding `token` and `userid` in-line in the
Glesys provider configuration:

Usage:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    glesys:token:
        value: ABC123
    glesys:userid:
        value: CL12345

```
### Environment variables

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:

```

Usage:

```sh
$ export GLESYS_TOKEN="ABC123"
$ export GLESYS_USERID="CL12345"
$ pulumi preview
```