---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/propelauth/propelauth/0.4.1/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Propelauth Provider
meta_desc: Provides an overview on how to configure the Pulumi Propelauth provider.
layout: package
---

## Generate Provider

The Propelauth provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider propelauth/propelauth
```
## Overview

Manage your PropelAuth integration for authentication, B2B authorization, and user management.
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    propelauth:apiKey:
        value: <PROPELAUTH_API_KEY>
    propelauth:projectId:
        value: <PROPELAUTH_PROJECT_ID>
    propelauth:tenantId:
        value: <PROPELAUTH_TENANT_ID>

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
    propelauth:apiKey:
        value: <PROPELAUTH_API_KEY>
    propelauth:projectId:
        value: <PROPELAUTH_PROJECT_ID>
    propelauth:tenantId:
        value: <PROPELAUTH_TENANT_ID>

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
    propelauth:apiKey:
        value: <PROPELAUTH_API_KEY>
    propelauth:projectId:
        value: <PROPELAUTH_PROJECT_ID>
    propelauth:tenantId:
        value: <PROPELAUTH_TENANT_ID>

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
    propelauth:apiKey:
        value: <PROPELAUTH_API_KEY>
    propelauth:projectId:
        value: <PROPELAUTH_PROJECT_ID>
    propelauth:tenantId:
        value: <PROPELAUTH_TENANT_ID>

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
    propelauth:apiKey:
        value: <PROPELAUTH_API_KEY>
    propelauth:projectId:
        value: <PROPELAUTH_PROJECT_ID>
    propelauth:tenantId:
        value: <PROPELAUTH_TENANT_ID>

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
    propelauth:apiKey:
        value: <PROPELAUTH_API_KEY>
    propelauth:projectId:
        value: <PROPELAUTH_PROJECT_ID>
    propelauth:tenantId:
        value: <PROPELAUTH_TENANT_ID>

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

- `apiKey` (String, Sensitive) A PropelAuth Infrastructure Integration Key for your project. You can generate one on the Infrastructure Integration page of the PropelAuth Dashboard. If not provided, the provider will attempt to use the PROPELAUTH_API_KEY environment variable.
- `projectId` (String) Your PropelAuth Project ID. This can be retrieved from Infrastructure Integration page of the PropelAuth Dashboard. If not provided, the provider will attempt to use the PROPELAUTH_PROJECT_ID environment variable.
- `tenantId` (String) Your PropelAuth Tenant ID. This can be retrieved from Infrastructure Integration page of the PropelAuth Dashboard. If not provided, the provider will attempt to use the PROPELAUTH_TENANT_ID environment variable.