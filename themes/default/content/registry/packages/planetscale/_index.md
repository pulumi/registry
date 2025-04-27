---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/planetscale/planetscale/0.6.0/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Planetscale Provider
meta_desc: Provides an overview on how to configure the Pulumi Planetscale provider.
layout: package
---

## Generate Provider

The Planetscale provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider planetscale/planetscale
```
## Overview

The PlanetScale provider allows using the OpenAPI surface of our public API. To use this provider, one of the following are required:

- access token credentials, configured or stored in the environment variable `PLANETSCALE_ACCESS_TOKEN`
- service token credentials, configured or stored in the environment variables `PLANETSCALE_SERVICE_TOKEN_ID` and `PLANETSCALE_SERVICE_TOKEN`

Note that the provider is not production ready and only for early testing at this time.

Known limitations:
- Support for deployments, deploy queues, deploy requests and reverts is not implemented at this time. If you have a use case for it, please let us know in the repository issues.
- When using service tokens (recommended), ensure the token has the `createDatabases` organization-level permission. This allows pulumi to create new databases and automatically grants the token all other permissions on the databases created by the token.
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    planetscale:serviceTokenName:
        value: 8fbddg0zlq0r

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
    planetscale:serviceTokenName:
        value: 8fbddg0zlq0r

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
    planetscale:serviceTokenName:
        value: 8fbddg0zlq0r

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
    planetscale:serviceTokenName:
        value: 8fbddg0zlq0r

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
    planetscale:serviceTokenName:
        value: 8fbddg0zlq0r

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
    planetscale:serviceTokenName:
        value: 8fbddg0zlq0r

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

- `accessToken` (String, Sensitive) Name of the service token to use. Alternatively, use `PLANETSCALE_SERVICE_TOKEN_ID`. Mutually exclusive with `serviceTokenId` and `serviceToken`.
- `endpoint` (String) If set, points the API client to a different endpoint than `https://api.planetscale.com/v1`.
- `serviceToken` (String, Sensitive) Value of the service token to use. Alternatively, use `PLANETSCALE_SERVICE_TOKEN`. Mutually exclusive with `accessToken`.
- `serviceTokenId` (String) ID of the service token to use. Alternatively, use `PLANETSCALE_SERVICE_TOKEN_ID`. Mutually exclusive with `accessToken`.
- `serviceTokenName` (String, Deprecated) Name of the service token to use. Alternatively, use `PLANETSCALE_SERVICE_TOKEN_NAME`. Mutually exclusive with `accessToken`. (Deprecated, use `serviceTokenId` instead)