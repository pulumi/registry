---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-opsgenie/v1.3.17/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/pulumi/pulumi-opsgenie/blob/v1.3.17/docs/_index.md
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Opsgenie Provider
meta_desc: Provides an overview on how to configure the Pulumi Opsgenie provider.
layout: package
---

## Installation

The Opsgenie provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/opsgenie`](https://www.npmjs.com/package/@pulumi/opsgenie)
* Python: [`pulumi-opsgenie`](https://pypi.org/project/pulumi-opsgenie/)
* Go: [`github.com/pulumi/pulumi-opsgenie/sdk/go/opsgenie`](https://github.com/pulumi/pulumi-opsgenie)
* .NET: [`Pulumi.Opsgenie`](https://www.nuget.org/packages/Pulumi.Opsgenie)
* Java: [`com.pulumi/opsgenie`](https://central.sonatype.com/artifact/com.pulumi/opsgenie)

## Overview

The Opsgenie provider is used to interact with the
many resources supported by Opsgenie. The provider needs to be configured
with the proper credentials before it can be used.

**Breaking Change - v0.6.0**

With 0.6.0 version provider adopted Pulumi Plugin SDK v2 therefore some resources reads has changed.
If you encounter any problems you can contact us via Github

Use the navigation to the left to read about the available resources.
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    opsgenie:apiKey:
        value: key
    opsgenie:apiUrl:
        value: api.eu.opsgenie.com

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as opsgenie from "@pulumi/opsgenie";

// Create a user
const test = new opsgenie.User("test", {});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    opsgenie:apiKey:
        value: key
    opsgenie:apiUrl:
        value: api.eu.opsgenie.com

```
```python
import pulumi
import pulumi_opsgenie as opsgenie

# Create a user
test = opsgenie.User("test")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    opsgenie:apiKey:
        value: key
    opsgenie:apiUrl:
        value: api.eu.opsgenie.com

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Opsgenie = Pulumi.Opsgenie;

return await Deployment.RunAsync(() =>
{
    // Create a user
    var test = new Opsgenie.User("test");

});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    opsgenie:apiKey:
        value: key
    opsgenie:apiUrl:
        value: api.eu.opsgenie.com

```
```go
package main

import (
	"github.com/pulumi/pulumi-opsgenie/sdk/go/opsgenie"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a user
		_, err := opsgenie.NewUser(ctx, "test", nil)
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
    opsgenie:apiKey:
        value: key
    opsgenie:apiUrl:
        value: api.eu.opsgenie.com

```
```yaml
resources:
  # Create a user
  test:
    type: opsgenie:User
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    opsgenie:apiKey:
        value: key
    opsgenie:apiUrl:
        value: api.eu.opsgenie.com

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.opsgenie.User;
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
        // Create a user
        var test = new User("test");

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

The following configuration inputs are supported:

* `apiKey` - (Required) The API Key for the Opsgenie Integration. If omitted, the
  `OPSGENIE_API_KEY` environment variable is used.

* `apiUrl` - (Optional) The API url for the Opsgenie.

You can generate an API Key within Opsgenie by creating a new API Integration with Read/Write permissions.