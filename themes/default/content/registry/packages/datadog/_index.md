---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-datadog/v4.59.0/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/pulumi/pulumi-datadog/blob/v4.59.0/docs/_index.md
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Datadog Provider
meta_desc: Provides an overview on how to configure the Pulumi Datadog provider.
layout: package
---

## Installation

The Datadog provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/datadog`](https://www.npmjs.com/package/@pulumi/datadog)
* Python: [`pulumi-datadog`](https://pypi.org/project/pulumi-datadog/)
* Go: [`github.com/pulumi/pulumi-datadog/sdk/v4/go/datadog`](https://github.com/pulumi/pulumi-datadog)
* .NET: [`Pulumi.Datadog`](https://www.nuget.org/packages/Pulumi.Datadog)
* Java: [`com.pulumi/datadog`](https://central.sonatype.com/artifact/com.pulumi/datadog)

## Overview

The [Datadog](https://www.datadoghq.com) provider is used to interact with the resources supported by Datadog. The provider needs to be configured with the proper credentials before it can be used. 

Use the navigation to the left to read about the available resources.
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    datadog:apiKey:
        value: 'TODO: var.datadog_api_key'
    datadog:appKey:
        value: 'TODO: var.datadog_app_key'

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
    datadog:apiKey:
        value: 'TODO: var.datadog_api_key'
    datadog:appKey:
        value: 'TODO: var.datadog_app_key'

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
    datadog:apiKey:
        value: 'TODO: var.datadog_api_key'
    datadog:appKey:
        value: 'TODO: var.datadog_app_key'

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
    datadog:apiKey:
        value: 'TODO: var.datadog_api_key'
    datadog:appKey:
        value: 'TODO: var.datadog_app_key'

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
    datadog:apiKey:
        value: 'TODO: var.datadog_api_key'
    datadog:appKey:
        value: 'TODO: var.datadog_app_key'

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
    datadog:apiKey:
        value: 'TODO: var.datadog_api_key'
    datadog:appKey:
        value: 'TODO: var.datadog_app_key'

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

- `apiKey` (String, Sensitive) (Required unless validate is false) Datadog API key. This can also be set via the DD_API_KEY environment variable.
- `apiUrl` (String) The API URL. This can also be set via the DD_HOST environment variable, and defaults to `https://api.datadoghq.com`. Note that this URL must not end with the `/api/` path. For example, `https://api.datadoghq.com/` is a correct value, while `https://api.datadoghq.com/api/` is not. And if you're working with "EU" version of Datadog, use `https://api.datadoghq.eu/`. Other Datadog region examples: `https://api.us5.datadoghq.com/`, `https://api.us3.datadoghq.com/` and `https://api.ddog-gov.com/`. See <https://docs.datadoghq.com/getting_started/site/> for all available regions.
- `appKey` (String, Sensitive) (Required unless validate is false) Datadog APP key. This can also be set via the DD_APP_KEY environment variable.
- `awsAccessKeyId` (String, Sensitive) The AWS access key ID; used for cloud-provider-based authentication. This can also be set using the `AWS_ACCESS_KEY_ID` environment variable. Required when using `cloudProviderType` set to `aws`.
- `awsSecretAccessKey` (String, Sensitive) The AWS secret access key; used for cloud-provider-based authentication. This can also be set using the `AWS_SECRET_ACCESS_KEY` environment variable. Required when using `cloudProviderType` set to `aws`.
- `awsSessionToken` (String, Sensitive) The AWS session token; used for cloud-provider-based authentication. This can also be set using the `AWS_SESSION_TOKEN` environment variable. Required when using `cloudProviderType` set to `aws` and using temporary credentials.
- `cloudProviderRegion` (String) The cloud provider region specifier; used for cloud-provider-based authentication. For example, `us-east-1` for AWS.
- `cloudProviderType` (String) Specifies the cloud provider used for cloud-provider-based authentication, enabling keyless access without API or app keys. Only [`aws`] is supported. This feature is in Preview. If you'd like to enable it for your organization, contact [support](https://docs.datadoghq.com/help/).
- `defaultTags` (Block List, Max: 1) [Experimental - Logs Pipelines, Monitors Security Monitoring Rules, and Service Level Objectives only] Configuration block containing settings to apply default resource tags across all resources. (see below for nested schema)
- `httpClientRetryBackoffBase` (Number) The HTTP request retry back off base. Defaults to 2.
- `httpClientRetryBackoffMultiplier` (Number) The HTTP request retry back off multiplier. Defaults to 2.
- `httpClientRetryEnabled` (String) Enables request retries on HTTP status codes 429 and 5xx. Valid values are [`true`, `false`]. Defaults to `true`.
- `httpClientRetryMaxRetries` (Number) The HTTP request maximum retry number. Defaults to 3.
- `httpClientRetryTimeout` (Number) The HTTP request retry timeout period. Defaults to 60 seconds.
- `orgUuid` (String) The organization UUID; used for cloud-provider-based authentication. See the [Datadog API documentation](https://docs.datadoghq.com/api/v1/organizations/) for more information.
- `validate` (String) Enables validation of the provided API key during provider initialization. Valid values are [`true`, `false`]. Default is true. When false, apiKey won't be checked.

<a id="nestedblock--default_tags"></a>
### Nested Schema for `defaultTags`

Optional:

- `tags` (Map of String) [Experimental - Logs Pipelines, Monitors Security Monitoring Rules, and Service Level Objectives only] Resource tags to be applied by default across all resources.