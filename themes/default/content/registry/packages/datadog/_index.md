---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-datadog/v4.49.0/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
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

{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

- `apiKey` (String, Sensitive) (Required unless validate is false) Datadog API key. This can also be set via the DD_API_KEY environment variable.
- `apiUrl` (String) The API URL. This can also be set via the DD_HOST environment variable, and defaults to `https://api.datadoghq.com`. Note that this URL must not end with the `/api/` path. For example, `https://api.datadoghq.com/` is a correct value, while `https://api.datadoghq.com/api/` is not. And if you're working with "EU" version of Datadog, use `https://api.datadoghq.eu/`. Other Datadog region examples: `https://api.us5.datadoghq.com/`, `https://api.us3.datadoghq.com/` and `https://api.ddog-gov.com/`. See <https://docs.datadoghq.com/getting_started/site/> for all available regions.
- `appKey` (String, Sensitive) (Required unless validate is false) Datadog APP key. This can also be set via the DD_APP_KEY environment variable.
- `defaultTags` (Block List, Max: 1) [Experimental - Monitors only] Configuration block containing settings to apply default resource tags across all resources. (see below for nested schema)
- `httpClientRetryBackoffBase` (Number) The HTTP request retry back off base. Defaults to 2.
- `httpClientRetryBackoffMultiplier` (Number) The HTTP request retry back off multiplier. Defaults to 2.
- `httpClientRetryEnabled` (String) Enables request retries on HTTP status codes 429 and 5xx. Valid values are [`true`, `false`]. Defaults to `true`.
- `httpClientRetryMaxRetries` (Number) The HTTP request maximum retry number. Defaults to 3.
- `httpClientRetryTimeout` (Number) The HTTP request retry timeout period. Defaults to 60 seconds.
- `validate` (String) Enables validation of the provided API key during provider initialization. Valid values are [`true`, `false`]. Default is true. When false, apiKey won't be checked.

<a id="nestedblock--default_tags"></a>
### Nested Schema for `defaultTags`

Optional:

- `tags` (Map of String) [Experimental - Monitors only] Resource tags to be applied by default across all resources.