---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/paloaltonetworks/panos/2.0.2/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Panos Provider
meta_desc: Provides an overview on how to configure the Pulumi Panos provider.
layout: package
---

## Generate Provider

The Panos provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider paloaltonetworks/panos
```
## Overview

PAN-OS&reg; is the operating system for Palo Alto Networks&reg; NGFWs and Panorama&trade;. The panos provider allows you to manage various aspects of a firewall's or a Panorama's config, such as data interfaces and security policies.

Use the navigation to the left to read about the available Panorama and NGFW resources.

> **NOTE:** The provider requires the use of Pulumi 1.8 or later.

> **NOTE:** The panos provider resources and functions are auto-generated based on [specs](https://github.com/PaloAltoNetworks/pan-os-codegen/tree/main/specs) using [pan-os-codegen(https://github.com/PaloAltoNetworks/pan-os-codegen/tree/main)
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    panos:hostname:
        value: hostname
    panos:password:
        value: password
    panos:username:
        value: username

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
    panos:hostname:
        value: hostname
    panos:password:
        value: password
    panos:username:
        value: username

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
    panos:hostname:
        value: hostname
    panos:password:
        value: password
    panos:username:
        value: username

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
    panos:hostname:
        value: hostname
    panos:password:
        value: password
    panos:username:
        value: username

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
    panos:hostname:
        value: hostname
    panos:password:
        value: password
    panos:username:
        value: username

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
    panos:hostname:
        value: hostname
    panos:password:
        value: password
    panos:username:
        value: username

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
## Authentication

The following authentication methods are supported. From highest to lowest priority;

- Static credentials via provider configuration
- Environment variables
- Configuration file
### Static Credentials

!> **Warning:** Including hard-coded credentials in Pulumi configurations is discouraged due to the risk of secret leakage.

Static credentials can be provided through either username-password combinations or API key-based authentication.

Usage:
### Environment Variables

You can provide your credentials using the `PANOS_USERNAME` and `PANOS_PASSWORD` environment variables for username-password based authentication, or use `PANOS_API_KEY` for API key based authentication, along with `PANOS_HOSTNAME`.

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:

```

Usage:

```sh
$ export PANOS_HOSTNAME="1.2.3.4"
$ export PANOS_USERNAME="username"
$ export PANOS_PASSWORD="password"
$ pulumi preview
```
### Configuration Files

You can also supply configuration parameters for the provider using a JSON configuration file.

Usage:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    panos:authFile:
        value: /path/to/auth_file.json

```

```json
{
  "hostname": "1.2.3.4",
  "username": "username",
  "password": "password",
}
```

```json
{
  "hostname": "1.2.3.4",
  "api_key": "api_key"
}
```
## Configuration Reference

- `additionalHeaders` (Map of String) Additional HTTP headers to send with API calls Environment variable: `PANOS_HEADERS`. JSON config file variable: `additionalHeaders`.
- `apiKey` (String) The API key for PAN-OS. Either specify this or give both username and password. Environment variable: `PANOS_API_KEY`. JSON config file variable: `apiKey`.
- `apiKeyInRequest` (Boolean) Send the API key in the request body instead of using the authentication header. Environment variable: `PANOS_API_KEY_IN_REQUEST`. JSON config file variable: `apiKeyInRequest`.
- `authFile` (String) Filesystem path to a JSON config file that specifies the provider's params. JSON config file variable: `authFile`.
- `configFile` (String) (Local inspection mode) The PAN-OS config file to load read in using `file()` JSON config file variable: `configFile`.
- `hostname` (String) The hostname or IP address of the PAN-OS instance (NGFW or Panorama). Environment variable: `PANOS_HOSTNAME`. JSON config file variable: `hostname`.
- `multiConfigBatchSize` (Number) Number of operations to send as part of a single MultiConfig update Default: `500`. Environment variable: `PANOS_MULTI_CONFIG_BATCH_SIZE`. JSON config file variable: `multiConfigBatchSize`.
- `panosVersion` (String) (Local inspection mode) The version of PAN-OS that exported the config file. This is only used if the root 'config' block does not contain the 'detail-version' attribute. Example: `10.2.3`. JSON config file variable: `panosVersion`.
- `password` (String, Sensitive) The password.  This is required if the apiKey is not configured. Environment variable: `PANOS_PASSWORD`. JSON config file variable: `password`.
- `port` (Number) If the port is non-standard for the protocol, the port number to use. Environment variable: `PANOS_PORT`. JSON config file variable: `port`.
- `protocol` (String) The protocol (https or http). Default: `https`. Environment variable: `PANOS_PROTOCOL`. JSON config file variable: `protocol`.
- `sdkLogCategories` (String) Log categories to configure for the PAN-OS SDK library Environment variable: `PANOS_LOG_CATEGORIES`. JSON config file variable: `sdkLogCategories`.
- `sdkLogLevel` (String) SDK logging Level for categories Default: `INFO`. Environment variable: `PANOS_LOG_LEVEL`. JSON config file variable: `sdkLogLevel`.
- `skipVerifyCertificate` (Boolean) (For https protocol) Skip verifying the HTTPS certificate. Environment variable: `PANOS_SKIP_VERIFY_CERTIFICATE`. JSON config file variable: `skipVerifyCertificate`.
- `target` (String) Target setting (NGFW serial number). Environment variable: `PANOS_TARGET`. JSON config file variable: `target`.
- `username` (String) The username.  This is required if apiKey is not configured. Environment variable: `PANOS_USERNAME`. JSON config file variable: `username`.