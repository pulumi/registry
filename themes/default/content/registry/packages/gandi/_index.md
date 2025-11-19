---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumiverse/pulumi-gandi/v2.3.2/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/pulumiverse/pulumi-gandi/blob/v2.3.2/docs/_index.md
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Gandi Provider
meta_desc: Provides an overview on how to configure the Pulumi Gandi provider.
layout: package
---

## Installation

The Gandi provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/gandi`](https://www.npmjs.com/package/@pulumi/gandi)
* Python: [`pulumi-gandi`](https://pypi.org/project/pulumi-gandi/)
* Go: [`github.com/pulumiverse/pulumi-gandi/sdk/v2/go/gandi`](https://github.com/pulumi/pulumi-gandi)
* .NET: [`Pulumi.Gandi`](https://www.nuget.org/packages/Pulumi.Gandi)
* Java: [`com.pulumi/gandi`](https://central.sonatype.com/artifact/com.pulumi/gandi)

## Overview

The Gandi provider enables the purchasing and management of the
following Gandi resources:

- [DNS zones](https://api.gandi.net/docs/domains/)
- [LiveDNS service](https://api.gandi.net/docs/livedns/)
- [Email](https://api.gandi.net/docs/email/)
- [SimpleHosting](https://api.gandi.net/docs/simplehosting/)

The provider needs to be configured with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    gandi:personalAccessToken:
        value: MY_PERSONAL_ACCESS_TOKEN

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as gandi from "@pulumiverse/gandi";

const exampleCom = new gandi.domains.Domain("example_com", {name: "example.com"});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    gandi:personalAccessToken:
        value: MY_PERSONAL_ACCESS_TOKEN

```
```python
import pulumi
import pulumiverse_gandi as gandi

example_com = gandi.domains.Domain("example_com", name="example.com")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    gandi:personalAccessToken:
        value: MY_PERSONAL_ACCESS_TOKEN

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Gandi = Pulumi.Gandi;

return await Deployment.RunAsync(() =>
{
    var exampleCom = new Gandi.Domains.Domain("example_com", new()
    {
        Name = "example.com",
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
    gandi:personalAccessToken:
        value: MY_PERSONAL_ACCESS_TOKEN

```
```go
package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-gandi/sdk/v2/go/gandi/domains"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := domains.NewDomain(ctx, "example_com", &domains.DomainArgs{
			Name: pulumi.String("example.com"),
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
    gandi:personalAccessToken:
        value: MY_PERSONAL_ACCESS_TOKEN

```
```yaml
resources:
  exampleCom:
    type: gandi:domains:Domain
    name: example_com
    properties:
      name: example.com
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    gandi:personalAccessToken:
        value: MY_PERSONAL_ACCESS_TOKEN

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.gandi.domains.Domain;
import com.pulumi.gandi.domains.DomainArgs;
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
        var exampleCom = new Domain("exampleCom", DomainArgs.builder()
            .name("example.com")
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Authentication

The Gandi provider supports a couple of different methods for providing authentication credentials.

The recommended way is to create a Personal Access Token. Read more about these tokens in the [Gandi public API documentation](https://api.gandi.net/docs/authentication/).

The previous method of using an API key is now deprecated and should not be used anymore, though it is still supported by this provider for now. When using an API Key, you could also provide a Sharing ID to specify an organization. If set, the Sharing ID indicates the organization that will pay for any ordered products, and will filter collections.
### Static Credentials

!> Hard-coding credentials into any Pulumi configuration is not recommended, and risks leaking secrets should the configuration be committed to public version control.

Usage:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    gandi:personalAccessToken:
        value: MY_PERSONAL_ACCESS_TOKEN

```
### Environment Variables

You can provide your credentials via the `GANDI_PERSONAL_ACCESS_TOKEN` environment variable, representing the Personal Access Token.

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:

```

Usage: