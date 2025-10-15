---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/kong/konnect/3.3.0/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Konnect Provider
meta_desc: Provides an overview on how to configure the Pulumi Konnect provider.
layout: package
---

## Generate Provider

The Konnect provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider kong/konnect
```
## Overview

Konnect API: The Konnect platform API
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs

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

- `konnectAccessToken` (String, Sensitive) The Konnect access token is meant to be used by the Konnect dashboard and the decK CLI to authenticate with..
- `personalAccessToken` (String, Sensitive) The personal access token is meant to be used as an alternative to basic-auth when accessing Konnect via APIs. You can generate a Personal Access Token (PAT) from the personal access token page in the Konnect dashboard.. Configurable via environment variable `KONNECT_TOKEN`.
- `serverUrl` (String) Server URL (defaults to <https://global.api.konghq.com>)
- `serviceAccessToken` (String, Sensitive) The Service access token is meant to be used between internal services.
  .
- `systemAccountAccessToken` (String, Sensitive) The system account access token is meant for automations and integrations that are not directly associated with a human identity.
  You can generate a system account Access Token by creating a system account and then obtaining a system account access token for that account.
  The access token must be passed in the header of a request, for example:
  `curl -X GET 'https://global.api.konghq.com/v2/users/' --header 'Authorization: Bearer spat_i2Ej...'`
  . Configurable via environment variable `KONNECT_SPAT`.