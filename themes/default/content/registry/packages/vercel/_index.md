---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumiverse/pulumi-vercel/v3.2.1/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Vercel Provider
meta_desc: Provides an overview on how to configure the Pulumi Vercel provider.
layout: package
---
## Installation

The Vercel provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/vercel`](https://www.npmjs.com/package/@pulumi/vercel)
* Python: [`pulumi-vercel`](https://pypi.org/project/pulumi-vercel/)
* Go: [`github.com/pulumiverse/pulumi-vercel/sdk/v3/go/vercel`](https://github.com/pulumi/pulumi-vercel)
* .NET: [`Pulumi.Vercel`](https://www.nuget.org/packages/Pulumi.Vercel)
* Java: [`com.pulumi/vercel`](https://central.sonatype.com/artifact/com.pulumi/vercel)
## Overview

The Vercel provider is used to interact with resources supported by Vercel.
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
    vercel:apiToken:
        value: 'TODO: var.vercel_api_token'
    vercel:team:
        value: your_team_slug_or_id

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
    vercel:apiToken:
        value: 'TODO: var.vercel_api_token'
    vercel:team:
        value: your_team_slug_or_id

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
    vercel:apiToken:
        value: 'TODO: var.vercel_api_token'
    vercel:team:
        value: your_team_slug_or_id

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
    vercel:apiToken:
        value: 'TODO: var.vercel_api_token'
    vercel:team:
        value: your_team_slug_or_id

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
    vercel:apiToken:
        value: 'TODO: var.vercel_api_token'
    vercel:team:
        value: your_team_slug_or_id

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
    vercel:apiToken:
        value: 'TODO: var.vercel_api_token'
    vercel:team:
        value: your_team_slug_or_id

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

- `apiToken` (String, Sensitive) The Vercel API Token to use. This can also be specified with the `VERCEL_API_TOKEN` shell environment variable. Tokens can be created from your [Vercel settings](https://vercel.com/account/tokens).
- `team` (String) The default Vercel Team to use when creating resources or reading functions. This can be provided as either a team slug, or team ID. The slug and ID are both available from the Team Settings page in the Vercel dashboard.