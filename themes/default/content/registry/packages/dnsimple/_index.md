---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-dnsimple/v4.3.1/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Dnsimple Provider
meta_desc: Provides an overview on how to configure the Pulumi Dnsimple provider.
layout: package
---
## Installation

The Dnsimple provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/dnsimple`](https://www.npmjs.com/package/@pulumi/dnsimple)
* Python: [`pulumi-dnsimple`](https://pypi.org/project/pulumi-dnsimple/)
* Go: [`github.com/pulumi/pulumi-dnsimple/sdk/v4/go/dnsimple`](https://github.com/pulumi/pulumi-dnsimple)
* .NET: [`Pulumi.Dnsimple`](https://www.nuget.org/packages/Pulumi.Dnsimple)
* Java: [`com.pulumi/dnsimple`](https://central.sonatype.com/artifact/com.pulumi/dnsimple)
## Overview

The DNSimple provider is used to interact with the resources supported by DNSimple. The provider needs to be configured
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
    dnsimple:account:
        value: 'TODO: "${var.dnsimple_account}"'
    dnsimple:token:
        value: 'TODO: "${var.dnsimple_token}"'

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as dnsimple from "@pulumi/dnsimple";

// Create a record
const www = new dnsimple.ZoneRecord("www", {});
// Create an email forward
const hello = new dnsimple.EmailForward("hello", {});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    dnsimple:account:
        value: 'TODO: "${var.dnsimple_account}"'
    dnsimple:token:
        value: 'TODO: "${var.dnsimple_token}"'

```
```python
import pulumi
import pulumi_dnsimple as dnsimple

# Create a record
www = dnsimple.ZoneRecord("www")
# Create an email forward
hello = dnsimple.EmailForward("hello")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    dnsimple:account:
        value: 'TODO: "${var.dnsimple_account}"'
    dnsimple:token:
        value: 'TODO: "${var.dnsimple_token}"'

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using DNSimple = Pulumi.DNSimple;

return await Deployment.RunAsync(() =>
{
    // Create a record
    var www = new DNSimple.ZoneRecord("www");

    // Create an email forward
    var hello = new DNSimple.EmailForward("hello");

});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    dnsimple:account:
        value: 'TODO: "${var.dnsimple_account}"'
    dnsimple:token:
        value: 'TODO: "${var.dnsimple_token}"'

```
```go
package main

import (
	"github.com/pulumi/pulumi-dnsimple/sdk/v4/go/dnsimple"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a record
		_, err := dnsimple.NewZoneRecord(ctx, "www", nil)
		if err != nil {
			return err
		}
		// Create an email forward
		_, err = dnsimple.NewEmailForward(ctx, "hello", nil)
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
    dnsimple:account:
        value: 'TODO: "${var.dnsimple_account}"'
    dnsimple:token:
        value: 'TODO: "${var.dnsimple_token}"'

```
```yaml
resources:
  # Create a record
  www:
    type: dnsimple:ZoneRecord
  # Create an email forward
  hello:
    type: dnsimple:EmailForward
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    dnsimple:account:
        value: 'TODO: "${var.dnsimple_account}"'
    dnsimple:token:
        value: 'TODO: "${var.dnsimple_token}"'

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.dnsimple.ZoneRecord;
import com.pulumi.dnsimple.EmailForward;
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
        // Create a record
        var www = new ZoneRecord("www");

        // Create an email forward
        var hello = new EmailForward("hello");

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

The following configuration inputs are supported:

* `token` - (Required) The DNSimple API v2 token. It must be provided, but it can also be sourced from the `DNSIMPLE_TOKEN` environment variable. Please note that this must be an [API v2 token](https://support.dnsimple.com/articles/api-access-token/). You can use either an User or Account token, but an Account token is recommended.
* `account` - (Required) The ID of the account associated with the token. It must be provided, but it can also be sourced from the `DNSIMPLE_ACCOUNT` environment variable.
* `sandbox` - Set to true to connect to the API [sandbox environment](https://developer.dnsimple.com/sandbox/). `DNSIMPLE_SANDBOX` environment variable can also be used.
* `prefetch` - Set to true to enable prefetching `ZoneRecords` when dealing with large configurations. This is useful
  when you are dealing with API rate limitations given your number of zones and zone records. `DNSIMPLE_PREFETCH` environment variable can also be used.
* `userAgent` - (Optional) Custom string to append to the user agent used for sending HTTP requests to the API.