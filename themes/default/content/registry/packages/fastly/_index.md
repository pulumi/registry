---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-fastly/v9.1.0/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Fastly Provider
meta_desc: Provides an overview on how to configure the Pulumi Fastly provider.
layout: package
---
## Installation

The Fastly provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/fastly`](https://www.npmjs.com/package/@pulumi/fastly)
* Python: [`pulumi-fastly`](https://pypi.org/project/pulumi-fastly/)
* Go: [`github.com/pulumi/pulumi-fastly/sdk/v8/go/fastly`](https://github.com/pulumi/pulumi-fastly)
* .NET: [`Pulumi.Fastly`](https://www.nuget.org/packages/Pulumi.Fastly)
* Java: [`com.pulumi/fastly`](https://central.sonatype.com/artifact/com.pulumi/fastly)
## Overview

The Fastly provider is used to interact with the content delivery network (CDN)
provided by Fastly.

In order to use this Provider, you must have an active account with Fastly.
Pricing and signup information can be found at <https://www.fastly.com/signup>

Use the navigation to the left to read about the available resources.
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    fastly:apiKey:
        value: test

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as fastly from "@pulumi/fastly";

// Create a Service
const myservice = new fastly.ServiceVcl("myservice", {name: "myawesometestservice"});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    fastly:apiKey:
        value: test

```
```python
import pulumi
import pulumi_fastly as fastly

# Create a Service
myservice = fastly.ServiceVcl("myservice", name="myawesometestservice")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    fastly:apiKey:
        value: test

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Fastly = Pulumi.Fastly;

return await Deployment.RunAsync(() =>
{
    // Create a Service
    var myservice = new Fastly.ServiceVcl("myservice", new()
    {
        Name = "myawesometestservice",
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
    fastly:apiKey:
        value: test

```
```go
package main

import (
	"github.com/pulumi/pulumi-fastly/sdk/v8/go/fastly"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a Service
		_, err := fastly.NewServiceVcl(ctx, "myservice", &fastly.ServiceVclArgs{
			Name: pulumi.String("myawesometestservice"),
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
    fastly:apiKey:
        value: test

```
```yaml
resources:
  # Create a Service
  myservice:
    type: fastly:ServiceVcl
    properties:
      name: myawesometestservice
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    fastly:apiKey:
        value: test

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.fastly.ServiceVcl;
import com.pulumi.fastly.ServiceVclArgs;
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
        // Create a Service
        var myservice = new ServiceVcl("myservice", ServiceVclArgs.builder()
            .name("myawesometestservice")
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Authentication

The Fastly provider offers an API key based method of providing credentials for
authentication. The following methods are supported, in this order, and
explained below:

- Static API key
- Environment variables
### Static API Key

Static credentials can be provided by adding a `apiKey` in-line:

Usage:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    fastly:apiKey:
        value: test

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as fastly from "@pulumi/fastly";

const myservice = new fastly.ServiceVcl("myservice", {});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    fastly:apiKey:
        value: test

```
```python
import pulumi
import pulumi_fastly as fastly

myservice = fastly.ServiceVcl("myservice")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    fastly:apiKey:
        value: test

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Fastly = Pulumi.Fastly;

return await Deployment.RunAsync(() =>
{
    var myservice = new Fastly.ServiceVcl("myservice");

});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    fastly:apiKey:
        value: test

```
```go
package main

import (
	"github.com/pulumi/pulumi-fastly/sdk/v8/go/fastly"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := fastly.NewServiceVcl(ctx, "myservice", nil)
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
    fastly:apiKey:
        value: test

```
```yaml
resources:
  myservice:
    type: fastly:ServiceVcl
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    fastly:apiKey:
        value: test

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.fastly.ServiceVcl;
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
        var myservice = new ServiceVcl("myservice");

    }
}
```
{{% /choosable %}}
{{< /chooser >}}

You can create a credential on the Personal API Tokens page: <https://manage.fastly.com/account/personal/tokens>
### Environment variables

You can provide your API key via `FASTLY_API_KEY` environment variable,
representing your Fastly API key.

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as fastly from "@pulumi/fastly";

const myservice = new fastly.ServiceVcl("myservice", {});
```
{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import pulumi_fastly as fastly

myservice = fastly.ServiceVcl("myservice")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Fastly = Pulumi.Fastly;

return await Deployment.RunAsync(() =>
{
    var myservice = new Fastly.ServiceVcl("myservice");

});

```
{{% /choosable %}}
{{% choosable language go %}}
```go
package main

import (
	"github.com/pulumi/pulumi-fastly/sdk/v8/go/fastly"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := fastly.NewServiceVcl(ctx, "myservice", nil)
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
resources:
  myservice:
    type: fastly:ServiceVcl
```
{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.fastly.ServiceVcl;
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
        var myservice = new ServiceVcl("myservice");

    }
}
```
{{% /choosable %}}
{{< /chooser >}}

Usage:

```sh
$ export FASTLY_API_KEY="afastlyapikey"
$ pulumi preview
```
## Configuration Reference

- `apiKey` (String) Fastly API Key from <https://app.fastly.com/#account>
- `baseUrl` (String) Fastly API URL
- `forceHttp2` (Boolean) Set this to `true` to disable HTTP/1.x fallback mechanism that the underlying Go library will attempt upon connection to `api.fastly.com:443` by default. This may slightly improve the provider's performance and reduce unnecessary TLS handshakes. Default: `false`
- `noAuth` (Boolean) Set to `true` if your configuration only consumes functions that do not require authentication, such as `fastly.getFastlyIpRanges`