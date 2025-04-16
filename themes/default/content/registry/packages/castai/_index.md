---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/castai/castai/7.45.0/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Castai Provider
meta_desc: Provides an overview on how to configure the Pulumi Castai provider.
layout: package
---

## Generate Provider

The Castai provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider castai/castai
```
## Overview

CAST AI provider can be used to onboard your cluster and manage resources supported by CAST AI.

> **Note** To use the provider, an API token first must be generated for your account at <https://console.cast.ai/>
## Provider configuration



{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    castai:apiToken:
        value: my-castai-api-token

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
    castai:apiToken:
        value: my-castai-api-token

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
    castai:apiToken:
        value: my-castai-api-token

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
    castai:apiToken:
        value: my-castai-api-token

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
    castai:apiToken:
        value: my-castai-api-token

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
    castai:apiToken:
        value: my-castai-api-token

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
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    castai:apiToken:
        value: 'TODO: var.castai_api_token'
    castai:apiUrl:
        value: 'TODO: var.castai_api_url'

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
    castai:apiToken:
        value: 'TODO: var.castai_api_token'
    castai:apiUrl:
        value: 'TODO: var.castai_api_url'

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
    castai:apiToken:
        value: 'TODO: var.castai_api_token'
    castai:apiUrl:
        value: 'TODO: var.castai_api_url'

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
    castai:apiToken:
        value: 'TODO: var.castai_api_token'
    castai:apiUrl:
        value: 'TODO: var.castai_api_url'

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
    castai:apiToken:
        value: 'TODO: var.castai_api_token'
    castai:apiUrl:
        value: 'TODO: var.castai_api_url'

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
    castai:apiToken:
        value: 'TODO: var.castai_api_token'
    castai:apiUrl:
        value: 'TODO: var.castai_api_url'

```
```java
Example currently unavailable in this language
```
{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference
### Required

- `apiToken` (String) The token used to connect to CAST AI API.

- `apiUrl` (String) CAST.AI API url.