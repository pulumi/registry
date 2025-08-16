---
# WARNING: this file was fetched from https://raw.githubusercontent.com/formalco/pulumi-formal/v1.0.2/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Formal Provider
meta_desc: Provides an overview on how to configure the Pulumi Formal provider.
layout: package
---

## Installation

The Formal provider is available as a package in the following Pulumi languages:

* JavaScript/TypeScript: [`@formalco/pulumi`](https://www.npmjs.com/package/@formalco/pulumi)
* Python: [`pulumi-formal`](https://pypi.org/project/pulumi-formal/)
* Go: [`github.com/formalco/pulumi-formal/sdk/go/formal`](https://pkg.go.dev/github.com/formalco/pulumi-formal/sdk/go/formal)
* .NET: [`Formal.Pulumi`](https://www.nuget.org/packages/Formal.Pulumi)

## Overview

Use the Formal Pulumi Provider to interact with the
many resources supported by Formal.

Use the navigation to the left to read about the available resources.

## Authentication and Configuration

Configuration for the Formal Provider is derived from the API tokens you can generate via the Formal Console.

### Provider Configuration

!> **Warning:** Hard-coded credentials are not recommended in any Pulumi
configuration and risks secret leakage should this file ever be committed to a
public version control system.

Credentials can be provided by adding an `apiKey`:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    formal:apiKey:
        value: '<apiKey>'
```

You can also use `pulumi config set formal:apiKey <apiKey> --secret` to set the API key.

Credentials can also be provided by using the `FORMAL_API_KEY` environment variable.

For example:
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:

```

```bash
export FORMAL_API_KEY="some_api_key" pulumi up
```

## Examples

{{< chooser language "go,typescript,python,csharp" >}}
{{% choosable language go %}}
```go
package main

import (
    formal "github.com/formalco/pulumi-formal/sdk/go/formal"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
    pulumi.Run(func(ctx *pulumi.Context) error {
        // Create a new connector instance.
        _, err := formal.NewConnector(ctx, "demo-connector", &formal.ConnectorArgs{
            Name:                  pulumi.String("demo-connector"),
            SpaceId:               nil,
            TerminationProtection: pulumi.Bool(false),
        })

        return err
    })
}
```
{{% /choosable %}}

{{% choosable language typescript %}}
```typescript
import * as formal from "@formalco/pulumi";

new formal.Connector('demo-connector', {
    name: 'demo-connector',
    spaceId: undefined,
    terminationProtection: false,
})
```
{{% /choosable %}}

{{% choosable language python %}}
```python
import pulumi_formal as formal

formal.Connector('demo-connector',
    name='demo-connector',
    space_id=None,
    termination_protection=False,
)
```
{{% /choosable %}}

{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using Pulumi;
using Formal.Pulumi;

return await Deployment.RunAsync(() =>
{
    var connector = new Connector("demo-connector", new ConnectorArgs {
        Name = "demo-connector",
        SpaceId = null,
        TerminationProtection = false
    });

    // Export outputs here
    return new Dictionary<string, object?>
    {
        ["demo-connector"] = connector.Id
    };
});
```
{{% /choosable %}}
{{< /chooser >}}

More examples on how to deploy Formal resources are available in the [`examples/`](https://github.com/formalco/pulumi-formal/tree/main/examples) folder of the Formal Pulumi repository.