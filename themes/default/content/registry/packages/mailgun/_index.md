---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-mailgun/v3.5.10/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Mailgun Provider
meta_desc: Provides an overview on how to configure the Pulumi Mailgun provider.
layout: package
---
## Installation

The Mailgun provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/mailgun`](https://www.npmjs.com/package/@pulumi/mailgun)
* Python: [`pulumi-mailgun`](https://pypi.org/project/pulumi-mailgun/)
* Go: [`github.com/pulumi/pulumi-mailgun/sdk/v3/go/mailgun`](https://github.com/pulumi/pulumi-mailgun)
* .NET: [`Pulumi.Mailgun`](https://www.nuget.org/packages/Pulumi.Mailgun)
* Java: [`com.pulumi/mailgun`](https://central.sonatype.com/artifact/com.pulumi/mailgun)
## Overview

The Mailgun provider is used to interact with the
resources supported by Mailgun. The provider needs to be configured
with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    mailgun:apiKey:
        value: 'TODO: "${var.mailgun_api_key}"'

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as mailgun from "@pulumi/mailgun";

// Create a new domain
const _default = new mailgun.Domain("default", {});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    mailgun:apiKey:
        value: 'TODO: "${var.mailgun_api_key}"'

```
```python
import pulumi
import pulumi_mailgun as mailgun

# Create a new domain
default = mailgun.Domain("default")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    mailgun:apiKey:
        value: 'TODO: "${var.mailgun_api_key}"'

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Mailgun = Pulumi.Mailgun;

return await Deployment.RunAsync(() =>
{
    // Create a new domain
    var @default = new Mailgun.Domain("default");

});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    mailgun:apiKey:
        value: 'TODO: "${var.mailgun_api_key}"'

```
```go
package main

import (
	"github.com/pulumi/pulumi-mailgun/sdk/v3/go/mailgun"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a new domain
		_, err := mailgun.NewDomain(ctx, "default", nil)
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
    mailgun:apiKey:
        value: 'TODO: "${var.mailgun_api_key}"'

```
```yaml
resources:
  # Create a new domain
  default:
    type: mailgun:Domain
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    mailgun:apiKey:
        value: 'TODO: "${var.mailgun_api_key}"'

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.mailgun.Domain;
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
        // Create a new domain
        var default_ = new Domain("default");

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

The following configuration inputs are supported:

* `apiKey` - (Required) Mailgun API key