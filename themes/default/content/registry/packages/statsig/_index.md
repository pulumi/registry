---
# WARNING: this file was fetched from https://raw.githubusercontent.com/statsig-io/pulumi-statsig/v0.0.1/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Statsig Provider
meta_desc: Provides an overview on how to configure the Pulumi Statsig provider.
layout: overview
---

The Statsig provider is used to interact with resources supported by
Statsig. The provider needs to be configured with the proper credentials
before it can be used.

## Example Usage

{{< chooser language "typescript,python,go,csharp" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    statsig:consoleApiKey:
        value: 'TODO: var.console_api_key'

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as statsig from "@statsig/pulumi-statsig";

// Create a Feature Gate
const gate = new Statsig.Gate("my-gate", {});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    statsig:consoleApiKey:
        value: 'TODO: var.console_api_key'

```
```python
import pulumi
import pulumi_statsig as statsig

# Create a Feature Gate
gate = statsig.Gate("my-gate")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    statsig:consoleApiKey:
        value: 'TODO: var.console_api_key'

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Statsig = Statsig.Pulumi;

return await Deployment.RunAsync(() =>
{
    // Create a Feature Gate
    var gate = new Statsig.Gate("my-gate");

});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    statsig:consoleApiKey:
        value: 'TODO: var.console_api_key'

```
```go
package main

import (
  "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/statsig-io/pulumi-statsig/sdk/go/statsig"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a Feature Gate
		_, err := statsig.NewGate(ctx, "my-gate", nil)
		if err != nil {
			return err
		}
		return nil
	})
}
```
{{% /choosable %}}
{{< /chooser >}}
