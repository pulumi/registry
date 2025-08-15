---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/honeycombio/honeycombio/0.38.1/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Honeycombio Provider
meta_desc: Provides an overview on how to configure the Pulumi Honeycombio provider.
layout: package
---

## Generate Provider

The Honeycombio provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider honeycombio/honeycombio
```
## Overview

[Honeycomb](https://honeycomb.io) provides observability for high-performance engineering teams so they can quickly understand what their code does in the hands of real users in unpredictable and highly complex cloud environments.
Honeycomb customers stop wasting precious time on engineering mysteries because they can quickly solve them and know exactly how to create fast, reliable, and great customer experiences.

In order to use this provider, you must have a Honeycomb account. You can get started today with a free account.

Use the navigation to the left to read about the available resources and functions.
## Example usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as honeycombio from "@pulumi/honeycombio";

const config = new pulumi.Config();
const dataset = config.require("dataset");
// Create a marker
const hello = new honeycombio.Marker("hello", {
    message: "Hello world!",
    dataset: dataset,
});
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
import pulumi_honeycombio as honeycombio

config = pulumi.Config()
dataset = config.require("dataset")
# Create a marker
hello = honeycombio.Marker("hello",
    message="Hello world!",
    dataset=dataset)
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
using Honeycombio = Pulumi.Honeycombio;

return await Deployment.RunAsync(() =>
{
    var config = new Config();
    var dataset = config.Require("dataset");
    // Create a marker
    var hello = new Honeycombio.Marker("hello", new()
    {
        Message = "Hello world!",
        Dataset = dataset,
    });

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
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/honeycombio/honeycombio"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")
		dataset := cfg.Require("dataset")
		// Create a marker
		_, err := honeycombio.NewMarker(ctx, "hello", &honeycombio.MarkerArgs{
			Message: pulumi.String("Hello world!"),
			Dataset: pulumi.String(dataset),
		})
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

```
```yaml
configuration:
  dataset:
    type: string
resources:
  # Create a marker
  hello:
    type: honeycombio:Marker
    properties:
      message: Hello world!
      dataset: ${dataset}
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
import com.pulumi.honeycombio.Marker;
import com.pulumi.honeycombio.MarkerArgs;
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
        final var dataset = config.get("dataset");
        // Create a marker
        var hello = new Marker("hello", MarkerArgs.builder()
            .message("Hello world!")
            .dataset(dataset)
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}

More advanced examples can be found in the example directory.
## A note on "Datasets"

Several resources in this provider accept a `dataset` or `datasets` argument to specify which Honeycomb Dataset the resource belongs to.
These resources include but aren't limited to:
* queries
* triggers
* slos
* markers
* columns
* boards

Whenever a resource accepts a `dataset` or `datasets` argument, the argument is expected to be a Dataset **slug**, not a Dataset name or ID.
Dataset slugs can be found in the URL of the dataset in the Honeycomb UI, or in the `slug` field of the [Dataset API](https://api-docs.honeycomb.io/api/datasets/createdataset#datasets/createdataset/t=response&c=200&path=slug).
### Configuring the provider for Honeycomb EU

If you are a Honeycomb EU customer, to use the provider you must override the default API host.
This can be done with a provider configuration or by setting the `HONEYCOMB_API_ENDPOINT` environment variable.

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    honeycombio:apiUrl:
        value: https://api.eu1.honeycomb.io

```
## Authentication

The Honeycomb provider requires an API key to communicate with the Honeycomb APIs.
The provider can make calls to v1 and v2 APIs and requires specific key configurations for each.
For more information about API Keys, check out [Best Practices for API Keys](https://docs.honeycomb.io/get-started/best-practices/api-keys/).

A single instance of the provider can be configured with both key types.
At least one of the v1 or v2 API key configuration is required.
### v1 APIs

v1 APIs require Configuration Keys.
Their permissions can be managed in *Environment settings*.
Most resources and functions call v1 APIs today.

The key can be set with the `apiKey` argument or via the `HONEYCOMB_API_KEY` or `HONEYCOMBIO_APIKEY` environment variable.

`HONEYCOMB_API_KEY` environment variable will take priority over the `HONEYCOMBIO_APIKEY` environment variable.
### v2 APIs

v2 APIs require a Mangement Key.
Their permissions can be managed in *Team settings*.
Resources and functions that call v2 APIs will be noted along with the scope required to use the resource or function.

The key pair can be set with the `apiKeyId` and `apiKeySecret` arguments, or via the `HONEYCOMB_KEY_ID` and `HONEYCOMB_KEY_SECRET` environment variables.

> **Note** Hard-coding API keys in any Pulumi configuration is not recommended. Consider using the one of the environment variable options.
## Configuration Reference

Arguments accepted by this provider include:

* `apiKey` - (Optional) The Honeycomb API key to use. It can also be set using `HONEYCOMB_API_KEY` or `HONEYCOMBIO_APIKEY` environment variables.
* `apiKeyId` - (Optional) The ID portion of the Honeycomb Management API key to use. It can also be set via the `HONEYCOMB_KEY_ID` environment variable.
* `apiKeySecret` - (Optional) The secret portion of the Honeycomb Management API key to use. It can also be set via the `HONEYCOMB_KEY_SECRET` environment variable.
* `apiUrl` - (Optional) Override the URL of the Honeycomb.io API. It can also be set using `HONEYCOMB_API_ENDPOINT`. Defaults to `https://api.honeycomb.io`.
* `debug` - (Optional) Enable to log additional debug information. To view the logs, set `TF_LOG` to at least debug.
* `features` - (Optional) The features block allows customization of the behavior of the Honeycomb Provider. Full details documented below.

At least one of `apiKey`, or the `apiKeyId` and `apiKeySecret` pair must be configured.
## Features Block

The Honeycomb Provider allows the behavior of certain resources to be modified using the features block.

This allows different users to select the behavior they require for their use case while preserving default, "Pulumi-y" behavior.
### Example Usage

If you wish to use the default behaviors of the Honeycomb provider, then nothing needs to be done to your configuration at all.

Each of the blocks defined below can be optionally specified to configure the behaviour as needed - this example shows all the possible behaviors which can be configured:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:

```
### Configuration Reference Reference

The `features` block supports the following:

* `column` - (Optional) A `column` block as defined below.
* `dataset` - (Optional) A `dataset` block as defined below.

---
The `column` block supports the following:
* `importOnConflict` - (Optional) This changes the creation behavior of the column resource to import and update an existing column if it already exists, rather than erroring out. Defaults to `false`.
  This is potentially dangerous if the type changes on the update -- switching from `string` to `boolean` and causing dataloss, for example -- and should be used with caution.

---
The `dataset` block supports the following:
* `importOnConflict` - (Optional) This changes the creation behavior of the dataset resource to import and update an existing dataset if it already exists, rather than erroring out. Defaults to `false`.