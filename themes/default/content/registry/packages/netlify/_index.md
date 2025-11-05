---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/netlify/netlify/0.3.2/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Netlify Provider
meta_desc: Provides an overview on how to configure the Pulumi Netlify provider.
layout: package
---

## Generate Provider

The Netlify provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider netlify/netlify
```
## Overview

The Netlify provider provides resources to manage Netlify resources like site configuration, environment variables, and Advanced Web Security features.
## Authentication

To use the provider, you will need a [personal access token](https://docs.netlify.com/api/get-started/#authentication).
You can create a new token in the [Netlify app](https://app.netlify.com/user/applications#personal-access-tokens).

You can expose the token as an environment variable:
```bash
export NETLIFY_API_TOKEN="your-personal-access-token"
```

Or by creating a Pulumi variable:
{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    netlify:token:
        value: 'TODO: var.netlify_api_token'

```
```typescript
import * as pulumi from "@pulumi/pulumi";

const config = new pulumi.Config();
const netlifyApiToken = config.require("netlifyApiToken");
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    netlify:token:
        value: 'TODO: var.netlify_api_token'

```
```python
import pulumi

config = pulumi.Config()
netlify_api_token = config.require("netlifyApiToken")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    netlify:token:
        value: 'TODO: var.netlify_api_token'

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;

return await Deployment.RunAsync(() =>
{
    var config = new Config();
    var netlifyApiToken = config.Require("netlifyApiToken");
});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    netlify:token:
        value: 'TODO: var.netlify_api_token'

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
		netlifyApiToken := cfg.Require("netlifyApiToken")
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
    netlify:token:
        value: 'TODO: var.netlify_api_token'

```
```yaml
configuration:
  netlifyApiToken:
    type: string
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    netlify:token:
        value: 'TODO: var.netlify_api_token'

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
        final var netlifyApiToken = config.get("netlifyApiToken");
    }
}
```
{{% /choosable %}}
{{< /chooser >}}

and setting the variable's value as an environment variable (`TF_VAR_netlify_api_token`).
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    netlify:defaultTeamSlug:
        value: your-team-slug
    netlify:token:
        value: 'TODO: var.netlify_api_token'

```
```typescript
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    netlify:defaultTeamSlug:
        value: your-team-slug
    netlify:token:
        value: 'TODO: var.netlify_api_token'

```
```python
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    netlify:defaultTeamSlug:
        value: your-team-slug
    netlify:token:
        value: 'TODO: var.netlify_api_token'

```
```csharp
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    netlify:defaultTeamSlug:
        value: your-team-slug
    netlify:token:
        value: 'TODO: var.netlify_api_token'

```
```go
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
config:
    netlify:defaultTeamSlug:
        value: your-team-slug
    netlify:token:
        value: 'TODO: var.netlify_api_token'

```
```yaml
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    netlify:defaultTeamSlug:
        value: your-team-slug
    netlify:token:
        value: 'TODO: var.netlify_api_token'

```
```java
Example currently unavailable in this language
```
{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

- `defaultTeamId` (String) The default team ID to use for resources that require a team ID or a team slug. Warning: Changing this value may not trigger recreation of resources.
- `defaultTeamSlug` (String) The default team slug to use for resources that require a team ID or a team slug. Warning: Changing this value may not trigger recreation of resources.
- `endpoint` (String) Defaults to: <https://api.netlify.com>
- `token` (String, Sensitive) Read: <https://docs.netlify.com/api/get-started/#authentication> , will use the `NETLIFY_API_TOKEN` environment variable if not set.