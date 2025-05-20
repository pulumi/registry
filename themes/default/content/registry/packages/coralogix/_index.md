---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/coralogix/coralogix/2.0.20/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Coralogix Provider
meta_desc: Provides an overview on how to configure the Pulumi Coralogix provider.
layout: package
---

## Generate Provider

The Coralogix provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider coralogix/coralogix
```
## Overview

Coralogix is a cloud-based, SaaS analytics and monitoring platform that combines logs, metrics, and traces to gain full observability into your system using one tool. The platform ingests data from any digital source and transforms it using our core features, allowing you to fully understand your system, analyze that data efficiently, and respond to incidents before they become problems.

Manage your Coralogix account from Pulumi, including alerts, dashboards, and more. First, sign up for an account at [Coralogix.com](https://coralogix.com/) and create an API key. With that key and your region you can then configure the provider as follows:

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

- `apiKey` (String, Sensitive) A key for using coralogix APIs (Auto Generated), appropriate for the defined environment. environment variable 'CORALOGIX_API_KEY' can be defined instead.
- `domain` (String) The Coralogix domain. Conflict With 'env'. environment variable 'CORALOGIX_DOMAIN' can be defined instead.
- `env` (String) The Coralogix API environment. can be one of ["AP1" "AP2" "AP3" "APAC1" "APAC2" "APAC3" "EU1" "EU2" "EUROPE1" "EUROPE2" "US1" "US2" "USA1" "USA2"]. environment variable 'CORALOGIX_ENV' can be defined instead.# Getting Started

Check out our examples for how to configure the various resources offered by the provider. If you already have Coralogix set up and want to import any existing resources, check out our migration script: pulumi-importer.
# Additional Notes
## Upgrading from V1.x.x to V2.x.x

In this version upgrade we changed the schema of our alerts, which are now incompatible to previous versions. You can ease the transition process by using the importer tool mentioned above so your state is safely upgraded. Note that for existing Coralogix users an additional process is required for upgrading your account. Please reach out to customer support to receive more guidance.