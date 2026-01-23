---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/temporalio/temporalcloud/1.1.2/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Temporalcloud Provider
meta_desc: Provides an overview on how to configure the Pulumi Temporalcloud provider.
layout: package
---

## Generate Provider

The Temporalcloud provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider temporalio/temporalcloud
```
## Overview

Use the `temporalcloud` provider to interact with resources supported by [Temporal Cloud](https://temporal.io/cloud).

Use the navigation to the left to learn about the available resources supported by this provider.
## Provider Configuration

Credentials for Temporal Cloud can be provided by adding an `apiKey` property or by setting the environment variable `TEMPORAL_CLOUD_API_KEY`.
You can generate an API key for Temporal Cloud by following [these instructions](https://docs.temporal.io/cloud/api-keys).

!> Hard-coded credentials are not recommended in any Pulumi configuration and should not be committed
in version control. We recommend passing credentials to this provider via environment variables.
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    temporalcloud:allowInsecure:
        value: false
    temporalcloud:allowedAccountId:
        value: my-temporalcloud-account-id
    temporalcloud:apiKey:
        value: my-temporalcloud-api-key
    temporalcloud:endpoint:
        value: saas-api.tmprl.cloud:443

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
    temporalcloud:allowInsecure:
        value: false
    temporalcloud:allowedAccountId:
        value: my-temporalcloud-account-id
    temporalcloud:apiKey:
        value: my-temporalcloud-api-key
    temporalcloud:endpoint:
        value: saas-api.tmprl.cloud:443

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
    temporalcloud:allowInsecure:
        value: false
    temporalcloud:allowedAccountId:
        value: my-temporalcloud-account-id
    temporalcloud:apiKey:
        value: my-temporalcloud-api-key
    temporalcloud:endpoint:
        value: saas-api.tmprl.cloud:443

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
    temporalcloud:allowInsecure:
        value: false
    temporalcloud:allowedAccountId:
        value: my-temporalcloud-account-id
    temporalcloud:apiKey:
        value: my-temporalcloud-api-key
    temporalcloud:endpoint:
        value: saas-api.tmprl.cloud:443

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
    temporalcloud:allowInsecure:
        value: false
    temporalcloud:allowedAccountId:
        value: my-temporalcloud-account-id
    temporalcloud:apiKey:
        value: my-temporalcloud-api-key
    temporalcloud:endpoint:
        value: saas-api.tmprl.cloud:443

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
    temporalcloud:allowInsecure:
        value: false
    temporalcloud:allowedAccountId:
        value: my-temporalcloud-account-id
    temporalcloud:apiKey:
        value: my-temporalcloud-api-key
    temporalcloud:endpoint:
        value: saas-api.tmprl.cloud:443

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

- `allowInsecure` (Boolean) If set to True, it allows for an insecure connection to the Temporal Cloud API. This should never be set to 'true' in production and defaults to false.
- `allowedAccountId` (String) The ID of the account to operate on. Prevents accidental mutation of accounts other than that provided.
- `apiKey` (String, Sensitive) The API key for Temporal Cloud. See [this documentation](https://docs.temporal.io/cloud/api-keys) for information on how to obtain an API key.
- `endpoint` (String) The endpoint for the Temporal Cloud API. Defaults to `saas-api.tmprl.cloud:443`.