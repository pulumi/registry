---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-spotinst/v3.120.2/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Spotinst Provider
meta_desc: Provides an overview on how to configure the Pulumi Spotinst provider.
layout: package
---
## Installation

The Spotinst provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/spotinst`](https://www.npmjs.com/package/@pulumi/spotinst)
* Python: [`pulumi-spotinst`](https://pypi.org/project/pulumi-spotinst/)
* Go: [`github.com/pulumi/pulumi-spotinst/sdk/v3/go/spotinst`](https://github.com/pulumi/pulumi-spotinst)
* .NET: [`Pulumi.Spotinst`](https://www.nuget.org/packages/Pulumi.Spotinst)
* Java: [`com.pulumi/spotinst`](https://central.sonatype.com/artifact/com.pulumi/spotinst)
## Overview

The Spotinst provider is used to interact with the
resources supported by Spotinst. The provider needs to be configured
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
    spotinst:account:
        value: 'TODO: "${var.spotinst_account}"'
    spotinst:token:
        value: 'TODO: "${var.spotinst_token}"'

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as spotinst from "@pulumi/spotinst";

// Create an Elastigroup
const foo = new spotinst.aws.Elastigroup("foo", {});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    spotinst:account:
        value: 'TODO: "${var.spotinst_account}"'
    spotinst:token:
        value: 'TODO: "${var.spotinst_token}"'

```
```python
import pulumi
import pulumi_spotinst as spotinst

# Create an Elastigroup
foo = spotinst.aws.Elastigroup("foo")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    spotinst:account:
        value: 'TODO: "${var.spotinst_account}"'
    spotinst:token:
        value: 'TODO: "${var.spotinst_token}"'

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using SpotInst = Pulumi.SpotInst;

return await Deployment.RunAsync(() =>
{
    // Create an Elastigroup
    var foo = new SpotInst.Aws.Elastigroup("foo");

});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    spotinst:account:
        value: 'TODO: "${var.spotinst_account}"'
    spotinst:token:
        value: 'TODO: "${var.spotinst_token}"'

```
```go
package main

import (
	"github.com/pulumi/pulumi-spotinst/sdk/v3/go/spotinst/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create an Elastigroup
		_, err := aws.NewElastigroup(ctx, "foo", nil)
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
    spotinst:account:
        value: 'TODO: "${var.spotinst_account}"'
    spotinst:token:
        value: 'TODO: "${var.spotinst_token}"'

```
```yaml
resources:
  # Create an Elastigroup
  foo:
    type: spotinst:aws:Elastigroup
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    spotinst:account:
        value: 'TODO: "${var.spotinst_account}"'
    spotinst:token:
        value: 'TODO: "${var.spotinst_token}"'

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.spotinst.aws.Elastigroup;
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
        // Create an Elastigroup
        var foo = new Elastigroup("foo");

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

The following configuration inputs are supported:

* `enabled` - (Optional) Boolean value to enable or disable the provider. Default is `true`.
* `token` - (Required) A Personal API Access Token issued by Spotinst. It can be sourced from the `SPOTINST_TOKEN` environment variable.
* `account` - (Optional) A valid Spotinst account ID. It can be sourced from the `SPOTINST_ACCOUNT` environment variable.
* `featureFlags` - (Optional) Spotinst SDK feature flags. They can be sourced from the `SPOTINST_FEATURE_FLAGS` environment variable.
## Credential Precedence

Credentials will be set given the following precedence:
1. credentials defined in the provider configuration of the template
2. credentials defined as environment variables
3. credentials defined in ~/.spotinst/credentials

The credentials can be merge in the chain by enabling the `MergeCredentialsChain` feature flag.

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    spotinst:account:
        value: 'TODO: "${var.spotinst_account}"'
    spotinst:featureFlags:
        value: MergeCredentialsChain=true

```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    spotinst:account:
        value: 'TODO: "${var.spotinst_account}"'
    spotinst:featureFlags:
        value: MergeCredentialsChain=true

```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    spotinst:account:
        value: 'TODO: "${var.spotinst_account}"'
    spotinst:featureFlags:
        value: MergeCredentialsChain=true

```

{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    spotinst:account:
        value: 'TODO: "${var.spotinst_account}"'
    spotinst:featureFlags:
        value: MergeCredentialsChain=true

```

{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
config:
    spotinst:account:
        value: 'TODO: "${var.spotinst_account}"'
    spotinst:featureFlags:
        value: MergeCredentialsChain=true

```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    spotinst:account:
        value: 'TODO: "${var.spotinst_account}"'
    spotinst:featureFlags:
        value: MergeCredentialsChain=true

```

{{% /choosable %}}
{{< /chooser >}}

Please note that if you omit the Spotinst account, resources will be created using the default account for your organization.