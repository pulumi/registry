---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/ucloud/ucloud/1.39.2/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Ucloud Provider
meta_desc: Provides an overview on how to configure the Pulumi Ucloud provider.
layout: package
---

## Generate Provider

The Ucloud provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider ucloud/ucloud
```
## Overview

> **NOTE:** This guide requires an available UCloud account or sub-account with project to create resources.

The UCloud provider is used to interact with the
resources supported by UCloud. The provider needs to be configured
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
    ucloud:privateKey:
        value: 'TODO: var.ucloud_private_key'
    ucloud:projectId:
        value: 'TODO: var.ucloud_project_id'
    ucloud:publicKey:
        value: 'TODO: var.ucloud_public_key'
    ucloud:region:
        value: cn-bj2

```
```typescript
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    ucloud:privateKey:
        value: 'TODO: var.ucloud_private_key'
    ucloud:projectId:
        value: 'TODO: var.ucloud_project_id'
    ucloud:publicKey:
        value: 'TODO: var.ucloud_public_key'
    ucloud:region:
        value: cn-bj2

```
```python
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    ucloud:privateKey:
        value: 'TODO: var.ucloud_private_key'
    ucloud:projectId:
        value: 'TODO: var.ucloud_project_id'
    ucloud:publicKey:
        value: 'TODO: var.ucloud_public_key'
    ucloud:region:
        value: cn-bj2

```
```csharp
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    ucloud:privateKey:
        value: 'TODO: var.ucloud_private_key'
    ucloud:projectId:
        value: 'TODO: var.ucloud_project_id'
    ucloud:publicKey:
        value: 'TODO: var.ucloud_public_key'
    ucloud:region:
        value: cn-bj2

```
```go
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
config:
    ucloud:privateKey:
        value: 'TODO: var.ucloud_private_key'
    ucloud:projectId:
        value: 'TODO: var.ucloud_project_id'
    ucloud:publicKey:
        value: 'TODO: var.ucloud_public_key'
    ucloud:region:
        value: cn-bj2

```
```yaml
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    ucloud:privateKey:
        value: 'TODO: var.ucloud_private_key'
    ucloud:projectId:
        value: 'TODO: var.ucloud_project_id'
    ucloud:publicKey:
        value: 'TODO: var.ucloud_public_key'
    ucloud:region:
        value: cn-bj2

```
```java
Example currently unavailable in this language
```
{{% /choosable %}}
{{< /chooser >}}
## Authentication

The UCloud provider offers a flexible means of providing credentials for
authentication. The following methods are supported, in this order, and
explained below:

- Static credentials
- Environment variables
### Static credentials

Static credentials can be provided by adding an `publicKey` and `privateKey` in-line in the
UCloud provider configuration:

Usage:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    ucloud:privateKey:
        value: your_private_key
    ucloud:projectId:
        value: your_project_id
    ucloud:publicKey:
        value: your_public_key
    ucloud:region:
        value: cn-bj2

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
    ucloud:privateKey:
        value: your_private_key
    ucloud:projectId:
        value: your_project_id
    ucloud:publicKey:
        value: your_public_key
    ucloud:region:
        value: cn-bj2

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
    ucloud:privateKey:
        value: your_private_key
    ucloud:projectId:
        value: your_project_id
    ucloud:publicKey:
        value: your_public_key
    ucloud:region:
        value: cn-bj2

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
    ucloud:privateKey:
        value: your_private_key
    ucloud:projectId:
        value: your_project_id
    ucloud:publicKey:
        value: your_public_key
    ucloud:region:
        value: cn-bj2

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
    ucloud:privateKey:
        value: your_private_key
    ucloud:projectId:
        value: your_project_id
    ucloud:publicKey:
        value: your_public_key
    ucloud:region:
        value: cn-bj2

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
    ucloud:privateKey:
        value: your_private_key
    ucloud:projectId:
        value: your_project_id
    ucloud:publicKey:
        value: your_public_key
    ucloud:region:
        value: cn-bj2

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
### Environment variables

You can provide your credentials via `UCLOUD_PUBLIC_KEY` and `UCLOUD_PRIVATE_KEY`
environment variables, representing your UCloud public key and private key respectively.
`UCLOUD_REGION` and `UCLOUD_PROJECT_ID` are also used, if applicable:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:

```

Usage:

```sh
$ export UCLOUD_PUBLIC_KEY="your_public_key"
$ export UCLOUD_PRIVATE_KEY="your_private_key"
$ export UCLOUD_REGION="cn-bj2"
$ export UCLOUD_PROJECT_ID="org-xxx"

$ pulumi preview
```
## Configuration Reference

In addition to generic `provider` arguments
(e.g. `alias` and `version`), the following arguments are supported in the UCloud
provider configuration:

* `publicKey` - (Required) This is the UCloud public key. You may refer to [get public key from console](https://console.ucloud.cn/uapi/apikey). It must be provided, but
  it can also be sourced from the `UCLOUD_PUBLIC_KEY` environment variable.

* `privateKey` - (Required) This is the UCloud private key. You may refer to [get private key from console](https://console.ucloud.cn/uapi/apikey). It must be provided, but
  it can also be sourced from the `UCLOUD_PRIVATE_KEY` environment variable.

* `region` - (Required) This is the UCloud region. It must be provided, but
  it can also be sourced from the `UCLOUD_REGION` environment variables.

* `projectId` - (Required) This is the UCloud project id. It must be provided, but
  it can also be sourced from the `UCLOUD_PROJECT_ID` environment variables.

* `maxRetries` - (Optional) This is the max retry attempts number. Default max retry attempts number is `0`.

* `insecure` - (Optional) This is a switch to disable/enable https.  (Default: `false`, means enable https).
  > **Note** this argument conflicts with `baseUrl`.

* `profile` - (Optional) This is the UCloud profile name as set in the shared credentials file, it can also be sourced from the `UCLOUD_PROFILE` environment variables.

* `sharedCredentialsFile` - (Optional) This is the path to the shared credentials file, it can also be sourced from the `UCLOUD_SHARED_CREDENTIAL_FILE` environment variables. If this is not set and a profile is specified, `~/.ucloud/credential.json` will be used.

* `baseUrl` - (Optional) This is the base url.(Default: `https://api.ucloud.cn`).
  > **Note** this argument conflicts with `insecure`.

* `assumeRole` - (Optional) Configuration block for assuming an IAM role. See the `assumeRole` Configuration Block section below. Only one `assumeRole` block may be in the configuration.
### assumeRole Configuration Block

The `assumeRole` configuration block supports the following arguments:

* `duration` - (Optional) Duration of the assume role session. Represented by a string such as `1h`, `2h45m`, or `30m15s`. Default is `900s`. The maximum value is the maximum session duration of the role
* `policy` - (Optional) IAM Policy JSON describing further restricting permissions for the IAM Role being assumed.
* `roleUrn` - (Required) URN of the IAM Role to assume.
* `sessionName` - (Required) Session name to use when assuming the role. The value can contain 1 to 64 letters, digits, underscores (_), hyphens (-), and periods (.).