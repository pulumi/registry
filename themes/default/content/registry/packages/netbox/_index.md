---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/e-breuninger/netbox/3.11.0/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Netbox Provider
meta_desc: Provides an overview on how to configure the Pulumi Netbox provider.
layout: package
---

## Generate Provider

The Netbox provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider e-breuninger/netbox
```
## Overview

The Pulumi Netbox provider is a plugin for Pulumi that allows for the full lifecycle management of [Netbox](https://docs.netbox.dev/en/stable/) resources.

Use the navigation to the left to read about the available resources.
## Supported Netbox versions
Netbox often makes breaking API changes even in non-major releases. Check the table below to see which version this provider is compatible with your Netbox version. It is generally recommended to use the provider version matching your Netbox version.

| Netbox version  | Provider version |
|-----------------|------------------|
| v4.1.0 - 4.1.11 | v3.10.0 and up   |
| v4.0.0 - 4.0.11 | v3.9.0 - 3.9.2   |
| v3.7.0 - 3.7.8  | v3.8.0 - 3.8.9   |
| v3.6.0 - 3.6.9  | v3.7.0 - 3.7.7   |
| v3.5.1 - 3.5.9  | v3.6.x           |
| v3.4.3 - 3.4.10 | v3.5.x           |
| v3.3.0 - 3.4.2  | v3.0.x - 3.5.1   |
| v3.2.0 - 3.2.9  | v2.0.x           |
| v3.1.9          | v1.6.0 - 1.6.7   |
| v3.1.3          | v1.1.x - 1.5.2   |
| v3.0.9          | v1.0.x           |
| v2.11.12        | v0.3.x           |
| v2.10.10        | v0.2.x           |
| v2.9            | v0.1.x           |

Additionally, since version 1.6.6, each version of the provider has a built-in list of all Netbox versions it supports at release time. Upon initialization, the provider will probe your Netbox version and include a (non-blocking) warning if the used Netbox version is not supported.
## Configuration
You must configure the provider with proper credentials before you can use it. You can configure the provider via attributes in the provider configuration or via environment variables. See Schema for all configuration options
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    netbox:apiToken:
        value: <your api key>
    netbox:serverUrl:
        value: https://demo.netbox.dev

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
    netbox:apiToken:
        value: <your api key>
    netbox:serverUrl:
        value: https://demo.netbox.dev

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
    netbox:apiToken:
        value: <your api key>
    netbox:serverUrl:
        value: https://demo.netbox.dev

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
    netbox:apiToken:
        value: <your api key>
    netbox:serverUrl:
        value: https://demo.netbox.dev

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
    netbox:apiToken:
        value: <your api key>
    netbox:serverUrl:
        value: https://demo.netbox.dev

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
    netbox:apiToken:
        value: <your api key>
    netbox:serverUrl:
        value: https://demo.netbox.dev

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
### Required

- `apiToken` (String) Netbox API authentication token. Can be set via the `NETBOX_API_TOKEN` environment variable.
- `serverUrl` (String) Location of Netbox server including scheme (http or https) and optional port. Can be set via the `NETBOX_SERVER_URL` environment variable.

- `allowInsecureHttps` (Boolean) Flag to set whether to allow https with invalid certificates. Can be set via the `NETBOX_ALLOW_INSECURE_HTTPS` environment variable. Defaults to `false`.
- `defaultTags` (Set of String) Tags to add to every resource managed by this provider.
- `headers` (Map of String) Set these header on all requests to Netbox. Can be set via the `NETBOX_HEADERS` environment variable.
- `requestTimeout` (Number) Netbox API HTTP request timeout in seconds. Can be set via the `NETBOX_REQUEST_TIMEOUT` environment variable.
- `skipVersionCheck` (Boolean) If true, do not try to determine the running Netbox version at provider startup. Disables warnings about possibly unsupported Netbox version. Also useful for local testing on pulumi plans. Can be set via the `NETBOX_SKIP_VERSION_CHECK` environment variable. Defaults to `false`.
- `stripTrailingSlashesFromUrl` (Boolean) If true, strip trailing slashes from the `serverUrl` parameter and print a warning when doing so. Note that using trailing slashes in the `serverUrl` parameter will usually lead to errors. Can be set via the `NETBOX_STRIP_TRAILING_SLASHES_FROM_URL` environment variable. Defaults to `true`.