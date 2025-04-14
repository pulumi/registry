---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/mrolla/circleci/0.6.1/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Circleci Provider
meta_desc: Provides an overview on how to configure the Pulumi Circleci provider.
layout: package
---

## Generate Provider

The Circleci provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider mrolla/circleci
```
## Overview

The CircleCI provider is used to interact with CircleCI API.
## Authentication

This provider requires a CircleCI API token in order to manage
resources.
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    circleci:apiToken:
        value: 'TODO: "${file("circleci_token")}"'
    circleci:organization:
        value: my_org
    circleci:vcsType:
        value: github

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as circleci from "@pulumi/circleci";

// Create a context
const build = new circleci.Context("build", {});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    circleci:apiToken:
        value: 'TODO: "${file("circleci_token")}"'
    circleci:organization:
        value: my_org
    circleci:vcsType:
        value: github

```
```python
import pulumi
import pulumi_circleci as circleci

# Create a context
build = circleci.Context("build")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    circleci:apiToken:
        value: 'TODO: "${file("circleci_token")}"'
    circleci:organization:
        value: my_org
    circleci:vcsType:
        value: github

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Circleci = Pulumi.Circleci;

return await Deployment.RunAsync(() =>
{
    // Create a context
    var build = new Circleci.Context("build");

});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    circleci:apiToken:
        value: 'TODO: "${file("circleci_token")}"'
    circleci:organization:
        value: my_org
    circleci:vcsType:
        value: github

```
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/circleci/circleci"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a context
		_, err := circleci.NewContext(ctx, "build", nil)
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
    circleci:apiToken:
        value: 'TODO: "${file("circleci_token")}"'
    circleci:organization:
        value: my_org
    circleci:vcsType:
        value: github

```
```yaml
resources:
  # Create a context
  build:
    type: circleci:Context
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    circleci:apiToken:
        value: 'TODO: "${file("circleci_token")}"'
    circleci:organization:
        value: my_org
    circleci:vcsType:
        value: github

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.circleci.Context;
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
        // Create a context
        var build = new Context("build");

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

The following configuration inputs are supported:

* `apiToken` - (Required) A CircleCI API token. This can also be set via the `CIRCLECI_TOKEN` environment variable.
* `vcsType` - (Optional) The version control system, either `"github"` or `"bitbucket"`. Defaults to `"github"`. This can also be set via the `CIRCLECI_VCS_TYPE` environment variable.
* `organization` - (Optional) The organization where resources will be created. If unset, an organization must be provided with each resource. This can also be set via the `CIRCLECI_ORGANIZATION` environment variable.
* `url` - (Optional) The URL for the the CircleCI API (v1). Defaults to `"https://circleci.com/api/v2/"`. This value should generally only be set for testing. This can also be set via the `CIRCLECI_URL` environment variable.