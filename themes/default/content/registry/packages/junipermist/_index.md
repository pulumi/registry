---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-junipermist/v0.7.3/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/pulumi/pulumi-junipermist/blob/v0.7.3/docs/_index.md
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Juniper Mist Provider
meta_desc: Provides an overview on how to configure the Pulumi Juniper Mist provider.
layout: package
---

## Installation

The Juniper Mist provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/junipermist`](https://www.npmjs.com/package/@pulumi/junipermist)
* Python: [`pulumi-junipermist`](https://pypi.org/project/pulumi-junipermist/)
* Go: [`github.com/pulumi/pulumi-junipermist/sdk/go/junipermist`](https://github.com/pulumi/pulumi-junipermist)
* .NET: [`Pulumi.Junipermist`](https://www.nuget.org/packages/Pulumi.Junipermist)
* Java: [`com.pulumi/junipermist`](https://central.sonatype.com/artifact/com.pulumi/junipermist)

## Overview

The Mist Provider allows Pulumi to manage Juniper Mist Organizations.

It enables Infrastructure as Code (IaC) automation to provision, deploy and manage Juniper Mist deployments. The provider can be used to manage Wi-Fi, Wired/Switch, WAN/SD-WAN and NAC resources in a Mist Organization.

Use the navigation tree to the left to read about the available resources and functions.
## Authentication

The provider supports two authentication methods:
* **API Token** (recommended): Use a Mist API token for authentication
* **Username/Password**: Use Mist account credentials (2FA must be disabled)

**Note:** Only one authentication method should be configured at a time.
## Supported Mist Clouds

This provider supports the following Mist Clouds:

**Global Clouds:**
* Global 01: `api.mist.com`
* Global 02: `api.gc1.mist.com`
* Global 03: `api.ac2.mist.com`
* Global 04: `api.gc2.mist.com`
* Global 05: `api.gc4.mist.com`

**EMEA Clouds:**
* EMEA 01: `api.eu.mist.com`
* EMEA 02: `api.gc3.mist.com`
* EMEA 03: `api.ac6.mist.com`
* EMEA 04: `api.gc6.mist.com`

**APAC Clouds:**
* APAC 01: `api.ac5.mist.com`
* APAC 02: `api.gc5.mist.com`
* APAC 03: `api.gc7.mist.com`
## Configuration
### Provider configuration example

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    mist:apitoken:
        value: xxxxxxxxxxxxxxxxxxxxxxxxxxx
    mist:host:
        value: api.mist.com

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
    mist:apitoken:
        value: xxxxxxxxxxxxxxxxxxxxxxxxxxx
    mist:host:
        value: api.mist.com

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
    mist:apitoken:
        value: xxxxxxxxxxxxxxxxxxxxxxxxxxx
    mist:host:
        value: api.mist.com

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
    mist:apitoken:
        value: xxxxxxxxxxxxxxxxxxxxxxxxxxx
    mist:host:
        value: api.mist.com

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
    mist:apitoken:
        value: xxxxxxxxxxxxxxxxxxxxxxxxxxx
    mist:host:
        value: api.mist.com

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
    mist:apitoken:
        value: xxxxxxxxxxxxxxxxxxxxxxxxxxx
    mist:host:
        value: api.mist.com

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
### Credentials

Users are encouraged to pass the API Token or the username and password via the
environment variables (see below). If authentication information are provided
in the provider configuration and in the environment variables, the Provider
configuration will be used.

Please consider whether writing credentials to a configuration file is
acceptable in your environment.
### Proxy Support

HTTP, HTTPS, and SOCKS5 proxies are supported through the `MIST_PROXY` environment
variables or the `proxy` provider configuration attribute.
## Configuration Reference

- `apiDebug` (Boolean) Enable API request/response debugging. When enabled, request and response bodies, headers, and other sensitive data may be logged. Can also be set via the `MIST_API_DEBUG` environment variable. Default: `false`.
- `apiTimeout` (Number) Timeout in seconds for API requests. Set to 0 for infinite timeout. Can also be set via the `MIST_API_TIMEOUT` environment variable. Default: `10` seconds.
- `apitoken` (String, Sensitive) Mist API Token for authentication. Can also be set via the `MIST_APITOKEN` environment variable. This is the recommended authentication method.
- `host` (String) URL of the Mist Cloud (e.g., `api.mist.com`). Can also be set via the `MIST_HOST` environment variable.
- `password` (String, Sensitive) Mist Account password for basic authentication. Can also be set via the `MIST_PASSWORD` environment variable. Requires `username` to be set.
- `proxy` (String) Proxy configuration for API requests. The value may be either a complete URL or `[username:password@]host[:port]` format. Supported schemes: `http`, `https`, and `socks5`. If no scheme is provided, `http` is assumed. Can also be set via the `MIST_PROXY` environment variable.
- `username` (String) Mist Account username for basic authentication. Can also be set via the `MIST_USERNAME` environment variable. Requires `password` to be set and 2FA to be disabled.
### Environment Variables

|   Variable Name    | Provider attribute |  Type  |                                                                                                                  Description                                                                                                                   |
|--------------------|--------------------|--------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `MIST_HOST`        | `host`             | String | URL of the Mist Cloud, e.g. `api.mist.com`. See above for the list of supported Clouds.                                                                                                                                                        |
| `MIST_API_TOKEN`   | `apitoken`         | String | For API Token authentication, the Mist API Token.                                                                                                                                                                                              |
| `MIST_USERNAME`    | `username`         | String | For username/password authentication, the Mist Account password.                                                                                                                                                                               |
| `MIST_PASSWORD`    | `password`         | String | For username/password authentication, the Mist Account password.                                                                                                                                                                               |
| `MIST_PROXY`       | `proxy`            | String | Requests use the configured proxy to reach the Mist Cloud. The value may be either a complete URL or a `[username:password@]host[:port]`, in which case the `http` scheme is assumed. The schemes `http`, `https`, and `socks5` are supported. |
| `MIST_API_TIMEOUT` | `apiTimeout`      | Int    | Timeout in seconds for completing API transactions with the Mist Cloud. Omit for default value of 10 seconds. Value of 0 results in infinite timeout.                                                                                          |