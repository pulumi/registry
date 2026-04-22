---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/g-core/gcore/2.0.0-alpha.5/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Gcore Provider
meta_desc: Provides an overview on how to configure the Pulumi Gcore provider.
layout: package
---

## Generate Provider

The Gcore provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider g-core/gcore
```
## Overview

The Gcore provider allows you to configure your [Gcore](https://gcore.com) infrastructure.

!> This provider is a complete rewrite with **breaking changes** compared to v0.x. Resource names, attribute schemas, and import IDs have changed. It is currently in **alpha** and under active development. The previous provider remains fully functional and is recommended for production workloads until v2 reaches GA. Install it from `registry.pulumi.io/providers/G-Core/gcore`.

> **Need help?** If you encounter any issues or have questions about the Gcore Pulumi provider, please reach out to [Gcore Support](https://support.gcore.com/hc/) for assistance.
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    gcore:apiKey:
        value: 'TODO: var.gcore_api_key'

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
    gcore:apiKey:
        value: 'TODO: var.gcore_api_key'

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
    gcore:apiKey:
        value: 'TODO: var.gcore_api_key'

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
    gcore:apiKey:
        value: 'TODO: var.gcore_api_key'

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
    gcore:apiKey:
        value: 'TODO: var.gcore_api_key'

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
    gcore:apiKey:
        value: 'TODO: var.gcore_api_key'

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

- `apiKey` (String, Sensitive) The API key for authenticating with the Gcore API. Can also be set via the `GCORE_API_KEY` environment variable.
- `baseUrl` (String) Set the base url that the provider connects to.
- `cloudPollingIntervalSeconds` (Number) Interval in seconds between polling requests for long-running cloud operations. Defaults to `3`.
- `cloudPollingTimeoutSeconds` (Number) Timeout in seconds for polling long-running cloud operations. Defaults to `7200`.
- `cloudProjectId` (Number) Default cloud project ID to use for cloud resources. Serves as a convenience fallback for local development; for production, prefer setting `projectId` explicitly on each resource. Can also be set via the `GCORE_CLOUD_PROJECT_ID` environment variable.
- `cloudRegionId` (Number) Default cloud region ID to use for cloud resources. Serves as a convenience fallback for local development; for production, prefer setting `regionId` explicitly on each resource. Can also be set via the `GCORE_CLOUD_REGION_ID` environment variable.
## Best Practices
### Project and Region Configuration

Cloud resources accept `projectId` and `regionId` as explicit attributes. The provider-level
`cloudProjectId` / `cloudRegionId` (and their `GCORE_CLOUD_PROJECT_ID` / `GCORE_CLOUD_REGION_ID`
environment variable counterparts) serve as **convenience fallbacks for local development and testing**.

For production configurations, always set `projectId` and `regionId` explicitly on each resource:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as gcore from "@pulumi/gcore";

const example = new gcore.CloudVolume("example", {
    projectId: 1,
    regionId: 1,
    name: "my-volume",
});
```
{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import pulumi_gcore as gcore

example = gcore.CloudVolume("example",
    project_id=1,
    region_id=1,
    name="my-volume")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Gcore = Pulumi.Gcore;

return await Deployment.RunAsync(() =>
{
    var example = new Gcore.CloudVolume("example", new()
    {
        ProjectId = 1,
        RegionId = 1,
        Name = "my-volume",
    });

});

```
{{% /choosable %}}
{{% choosable language go %}}
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/gcore/v2/gcore"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := gcore.NewCloudVolume(ctx, "example", &gcore.CloudVolumeArgs{
			ProjectId: pulumi.Float64(1),
			RegionId:  pulumi.Float64(1),
			Name:      pulumi.String("my-volume"),
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
resources:
  example:
    type: gcore:CloudVolume
    properties:
      projectId: 1
      regionId: 1
      name: my-volume
```
{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.gcore.CloudVolume;
import com.pulumi.gcore.CloudVolumeArgs;
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
        var example = new CloudVolume("example", CloudVolumeArgs.builder()
            .projectId(1.0)
            .regionId(1.0)
            .name("my-volume")
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}

This makes your Pulumi configuration self-contained, portable, and unambiguous — it won't
silently change behavior when run in a different environment.