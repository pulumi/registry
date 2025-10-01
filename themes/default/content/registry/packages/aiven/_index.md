---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-aiven/v6.43.1/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/pulumi/pulumi-aiven/blob/v6.43.1/docs/_index.md
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Aiven Provider
meta_desc: Provides an overview on how to configure the Pulumi Aiven provider.
layout: package
---

## Installation

The Aiven provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/aiven`](https://www.npmjs.com/package/@pulumi/aiven)
* Python: [`pulumi-aiven`](https://pypi.org/project/pulumi-aiven/)
* Go: [`github.com/pulumi/pulumi-aiven/sdk/v6/go/aiven`](https://github.com/pulumi/pulumi-aiven)
* .NET: [`Pulumi.Aiven`](https://www.nuget.org/packages/Pulumi.Aiven)
* Java: [`com.pulumi/aiven`](https://central.sonatype.com/artifact/com.pulumi/aiven)

## Overview

The Pulumi provider for [Aiven](https://aiven.io/), your AI-ready open source data platform.
## Authentication
Sign up for Aiven and [create a personal token](https://aiven.io/docs/platform/howto/create_authentication_token).

You can also create an [application user](https://aiven.io/docs/platform/howto/manage-application-users) and use its token for accessing the Aiven Provider.
## Example usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    aiven:apiToken:
        value: 'TODO: var.aiven_api_token'

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
    aiven:apiToken:
        value: 'TODO: var.aiven_api_token'

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
    aiven:apiToken:
        value: 'TODO: var.aiven_api_token'

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
    aiven:apiToken:
        value: 'TODO: var.aiven_api_token'

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
    aiven:apiToken:
        value: 'TODO: var.aiven_api_token'

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
    aiven:apiToken:
        value: 'TODO: var.aiven_api_token'

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
## Environment variables

* For authentication, you can set the `AIVEN_TOKEN` to your token value.
* To use beta resources, set `PROVIDER_AIVEN_ENABLE_BETA` to any value.
* To allow IP filters to be purged, set `AIVEN_ALLOW_IP_FILTER_PURGE` to any value. This feature prevents accidental purging of IP filters, which can cause you to lose access to services.
## Resource options
The list of options in this document is not comprehensive. However, most map directly to the [Aiven REST API](https://api.aiven.io/doc/) properties.