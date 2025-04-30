---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-wavefront/v3.1.9/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Wavefront Provider
meta_desc: Provides an overview on how to configure the Pulumi Wavefront provider.
layout: package
---
## Installation

The Wavefront provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/wavefront`](https://www.npmjs.com/package/@pulumi/wavefront)
* Python: [`pulumi-wavefront`](https://pypi.org/project/pulumi-wavefront/)
* Go: [`github.com/pulumi/pulumi-wavefront/sdk/v3/go/wavefront`](https://github.com/pulumi/pulumi-wavefront)
* .NET: [`Pulumi.Wavefront`](https://www.nuget.org/packages/Pulumi.Wavefront)
* Java: [`com.pulumi/wavefront`](https://central.sonatype.com/artifact/com.pulumi/wavefront)
## Overview

The Wavefront provider is used to interact with the Wavefront monitoring service. The
provider needs to be configured with the proper credentials before it can be used.

Use the navigation on the left to read about the available resources.
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as wavefront from "@pulumi/wavefront";

const testAlert = new wavefront.Alert("test_alert", {
    name: "High CPU Alert",
    condition: "100-ts(\"cpu.usage_idle\", environment=preprod and cpu=cpu-total ) > 80",
    additionalInformation: "This is an Alert",
    displayExpression: "100-ts(\"cpu.usage_idle\", environment=preprod and cpu=cpu-total )",
    minutes: 5,
    severity: "WARN",
    tags: [
        "env.preprod",
        "cpu.total",
    ],
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
import pulumi_wavefront as wavefront

test_alert = wavefront.Alert("test_alert",
    name="High CPU Alert",
    condition="100-ts(\"cpu.usage_idle\", environment=preprod and cpu=cpu-total ) > 80",
    additional_information="This is an Alert",
    display_expression="100-ts(\"cpu.usage_idle\", environment=preprod and cpu=cpu-total )",
    minutes=5,
    severity="WARN",
    tags=[
        "env.preprod",
        "cpu.total",
    ])
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
using Wavefront = Pulumi.Wavefront;

return await Deployment.RunAsync(() =>
{
    var testAlert = new Wavefront.Alert("test_alert", new()
    {
        Name = "High CPU Alert",
        Condition = "100-ts(\"cpu.usage_idle\", environment=preprod and cpu=cpu-total ) > 80",
        AdditionalInformation = "This is an Alert",
        DisplayExpression = "100-ts(\"cpu.usage_idle\", environment=preprod and cpu=cpu-total )",
        Minutes = 5,
        Severity = "WARN",
        Tags = new[]
        {
            "env.preprod",
            "cpu.total",
        },
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
	"github.com/pulumi/pulumi-wavefront/sdk/v3/go/wavefront"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := wavefront.NewAlert(ctx, "test_alert", &wavefront.AlertArgs{
			Name:                  pulumi.String("High CPU Alert"),
			Condition:             pulumi.String("100-ts(\"cpu.usage_idle\", environment=preprod and cpu=cpu-total ) > 80"),
			AdditionalInformation: pulumi.String("This is an Alert"),
			DisplayExpression:     pulumi.String("100-ts(\"cpu.usage_idle\", environment=preprod and cpu=cpu-total )"),
			Minutes:               pulumi.Int(5),
			Severity:              pulumi.String("WARN"),
			Tags: pulumi.StringArray{
				pulumi.String("env.preprod"),
				pulumi.String("cpu.total"),
			},
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
resources:
  testAlert:
    type: wavefront:Alert
    name: test_alert
    properties:
      name: High CPU Alert
      condition: 100-ts("cpu.usage_idle", environment=preprod and cpu=cpu-total ) > 80
      additionalInformation: This is an Alert
      displayExpression: 100-ts("cpu.usage_idle", environment=preprod and cpu=cpu-total )
      minutes: 5
      severity: WARN
      tags:
        - env.preprod
        - cpu.total
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
import com.pulumi.wavefront.Alert;
import com.pulumi.wavefront.AlertArgs;
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
        var testAlert = new Alert("testAlert", AlertArgs.builder()
            .name("High CPU Alert")
            .condition("100-ts(\"cpu.usage_idle\", environment=preprod and cpu=cpu-total ) > 80")
            .additionalInformation("This is an Alert")
            .displayExpression("100-ts(\"cpu.usage_idle\", environment=preprod and cpu=cpu-total )")
            .minutes(5)
            .severity("WARN")
            .tags(
                "env.preprod",
                "cpu.total")
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Authentication

The Wavefront provider offers two ways of providing credentials for authentication.

* Static credentials
* Environment variables
### Static credentials

Static credentials can be provided by adding an `address` and `token` in-line in
the Wavefront provider configuration.

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    wavefront:address:
        value: cluster.wavefront.com
    wavefront:token:
        value: your-wf-token-secret

```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    wavefront:address:
        value: cluster.wavefront.com
    wavefront:token:
        value: your-wf-token-secret

```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    wavefront:address:
        value: cluster.wavefront.com
    wavefront:token:
        value: your-wf-token-secret

```

{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    wavefront:address:
        value: cluster.wavefront.com
    wavefront:token:
        value: your-wf-token-secret

```

{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
config:
    wavefront:address:
        value: cluster.wavefront.com
    wavefront:token:
        value: your-wf-token-secret

```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    wavefront:address:
        value: cluster.wavefront.com
    wavefront:token:
        value: your-wf-token-secret

```

{{% /choosable %}}
{{< /chooser >}}
### Environment Variables

You can provide your credentials by using the `WAVEFRONT_ADDRESS` and `WAVEFRONT_TOKEN`,
environment variables.

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs

```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python

```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet

```

{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go

```

{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml

```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java

```

{{% /choosable %}}
{{< /chooser >}}

Usage:

```sh
$ export WAVEFRONT_ADDRESS="cluster.wavefront.com"
$ export WAVEFRONT_TOKEN="your-wf-token-secret"
$ pulumi preview
```
## Configuration Reference

In addition to generic `provider` arguments
(e.g. `alias` and `version`), the following arguments are supported in the Wavefront
provider configuration:

* `address` - (Optional) The URL of your Wavefront cluster that you access Wavefront from without the
  leading `https://` or trailing `/` (e.g. `https://longboard.wavefront.com/` becomes `longboard.wavefront.com`)

* `token` - (Optional) Either a User Account token or Service Account token with the permissions necessary
  to manage your Wavefront account.

* `httpProxy` - (Optional) The proxy type is determined by the URL scheme. `http`, `https`, and `socks5` are supported.
  If the scheme is empty `http` is assumed.