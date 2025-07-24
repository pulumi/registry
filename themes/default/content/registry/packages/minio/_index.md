---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-minio/v0.16.5/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Minio Provider
meta_desc: Provides an overview on how to configure the Pulumi Minio provider.
layout: package
---
## Installation

The Minio provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/minio`](https://www.npmjs.com/package/@pulumi/minio)
* Python: [`pulumi-minio`](https://pypi.org/project/pulumi-minio/)
* Go: [`github.com/pulumi/pulumi-minio/sdk/go/minio`](https://github.com/pulumi/pulumi-minio)
* .NET: [`Pulumi.Minio`](https://www.nuget.org/packages/Pulumi.Minio)
* Java: [`com.pulumi/minio`](https://central.sonatype.com/artifact/com.pulumi/minio)
## Overview

This is a pulumi provider plugin for managing [Minio](https://min.io/) S3 buckets and IAM users.
## Example Provider Configuration

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    minio:minioApiVersion:
        value: '...'
    minio:minioPassword:
        value: '...'
    minio:minioRegion:
        value: '...'
    minio:minioServer:
        value: '...'
    minio:minioSsl:
        value: '...'
    minio:minioUser:
        value: '...'

```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    minio:minioApiVersion:
        value: '...'
    minio:minioPassword:
        value: '...'
    minio:minioRegion:
        value: '...'
    minio:minioServer:
        value: '...'
    minio:minioSsl:
        value: '...'
    minio:minioUser:
        value: '...'

```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    minio:minioApiVersion:
        value: '...'
    minio:minioPassword:
        value: '...'
    minio:minioRegion:
        value: '...'
    minio:minioServer:
        value: '...'
    minio:minioSsl:
        value: '...'
    minio:minioUser:
        value: '...'

```

{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    minio:minioApiVersion:
        value: '...'
    minio:minioPassword:
        value: '...'
    minio:minioRegion:
        value: '...'
    minio:minioServer:
        value: '...'
    minio:minioSsl:
        value: '...'
    minio:minioUser:
        value: '...'

```

{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
config:
    minio:minioApiVersion:
        value: '...'
    minio:minioPassword:
        value: '...'
    minio:minioRegion:
        value: '...'
    minio:minioServer:
        value: '...'
    minio:minioSsl:
        value: '...'
    minio:minioUser:
        value: '...'

```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    minio:minioApiVersion:
        value: '...'
    minio:minioPassword:
        value: '...'
    minio:minioRegion:
        value: '...'
    minio:minioServer:
        value: '...'
    minio:minioSsl:
        value: '...'
    minio:minioUser:
        value: '...'

```

{{% /choosable %}}
{{< /chooser >}}
## Authentication

The Minio provider offers the following methods of providing credentials for
authentication, in this order, and explained below:

- Static API key
- Environment variables
### Static API Key

Static credentials can be provided by adding the `minio-server`, `minioUser` and `minioPassword` variables in-line in the
Minio provider configuration:

Usage:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    minio:minioPassword:
        value: '...'
    minio:minioServer:
        value: '...'
    minio:minioUser:
        value: '...'

```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    minio:minioPassword:
        value: '...'
    minio:minioServer:
        value: '...'
    minio:minioUser:
        value: '...'

```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    minio:minioPassword:
        value: '...'
    minio:minioServer:
        value: '...'
    minio:minioUser:
        value: '...'

```

{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    minio:minioPassword:
        value: '...'
    minio:minioServer:
        value: '...'
    minio:minioUser:
        value: '...'

```

{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
config:
    minio:minioPassword:
        value: '...'
    minio:minioServer:
        value: '...'
    minio:minioUser:
        value: '...'

```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    minio:minioPassword:
        value: '...'
    minio:minioServer:
        value: '...'
    minio:minioUser:
        value: '...'

```

{{% /choosable %}}
{{< /chooser >}}
### Environment variables

You can provide your configuration via the environment variables representing your minio credentials:

```
$ export MINIO_ENDPOINT="http://myendpoint"
$ export MINIO_USER="244tefewg"
$ export MINIO_PASSWORD="xgwgwqqwv"
```

When using this method, you may omit the
minio provider configuration entirely:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as minio from "@pulumi/minio";

const statePulumiS3 = new minio.S3Bucket("state_pulumi_s3", {});
```
{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import pulumi_minio as minio

state_pulumi_s3 = minio.S3Bucket("state_pulumi_s3")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Minio = Pulumi.Minio;

return await Deployment.RunAsync(() =>
{
    var statePulumiS3 = new Minio.S3Bucket("state_pulumi_s3");

});

```
{{% /choosable %}}
{{% choosable language go %}}
```go
package main

import (
	"github.com/pulumi/pulumi-minio/sdk/go/minio"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := minio.NewS3Bucket(ctx, "state_pulumi_s3", nil)
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
  statePulumiS3:
    type: minio:S3Bucket
    name: state_pulumi_s3
```
{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.minio.S3Bucket;
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
        var statePulumiS3 = new S3Bucket("statePulumiS3");

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

The following configuration inputs are supported in the provider configuration:

* `minioServer` - (Required) Minio Host and Port. It must be provided, but
  it can also be sourced from the `MINIO_ENDPOINT` environment variable

* `minioUser` - (Required) Minio User. It must be provided, but
  it can also be sourced from the `MINIO_USER` environment variable

* `minioPassword` - (Required) Minio Password. It must be provided, but
  it can also be sourced from the `MINIO_PASSWORD` environment variable

* `minioRegion` - (Optional) Minio Region (`default: us-east-1`).

* `minioApiVersion` - (Optional) Minio API Version (type: string, options: `v2` or `v4`, default: `v4`).

* `minioSsl` - (Optional) Minio SSL enabled (default: `false`). It can also be sourced from the
  `MINIO_ENABLE_HTTPS` environment variable