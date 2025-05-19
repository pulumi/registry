---
# WARNING: this file was fetched from https://raw.githubusercontent.com/lbrlabs/pulumi-harness/v0.7.4/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Harness Provider
meta_desc: Provides an overview on how to configure the Pulumi Harness provider.
layout: package
---
## Installation

The Harness provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/harness`](https://www.npmjs.com/package/@pulumi/harness)
* Python: [`pulumi-harness`](https://pypi.org/project/pulumi-harness/)
* Go: [`github.com/pulumi/pulumi-harness/sdk/go/harness`](https://github.com/pulumi/pulumi-harness)
* .NET: [`Pulumi.Harness`](https://www.nuget.org/packages/Pulumi.Harness)
* Java: [`com.pulumi/harness`](https://central.sonatype.com/artifact/com.pulumi/harness)
## Overview

For an explanation on how to use this Provider along with code samples, refer to the Harness Pulumi Provider Quickstart Guide <https://docs.harness.io/article/7cude5tvzh-harness-pulumi-provider>.
## Example Usage

Configure the Harness provider for First Gen resources
{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    harness:accountId:
        value: '....'
    harness:apiKey:
        value: '......'
    harness:endpoint:
        value: https://app.harness.io/gateway

```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    harness:accountId:
        value: '....'
    harness:apiKey:
        value: '......'
    harness:endpoint:
        value: https://app.harness.io/gateway

```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    harness:accountId:
        value: '....'
    harness:apiKey:
        value: '......'
    harness:endpoint:
        value: https://app.harness.io/gateway

```

{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    harness:accountId:
        value: '....'
    harness:apiKey:
        value: '......'
    harness:endpoint:
        value: https://app.harness.io/gateway

```

{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
config:
    harness:accountId:
        value: '....'
    harness:apiKey:
        value: '......'
    harness:endpoint:
        value: https://app.harness.io/gateway

```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    harness:accountId:
        value: '....'
    harness:apiKey:
        value: '......'
    harness:endpoint:
        value: https://app.harness.io/gateway

```

{{% /choosable %}}
{{< /chooser >}}

Configure the Harness provider for Next Gen resources
{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    harness:accountId:
        value: '....'
    harness:endpoint:
        value: https://app.harness.io/gateway
    harness:platformApiKey:
        value: '......'

```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    harness:accountId:
        value: '....'
    harness:endpoint:
        value: https://app.harness.io/gateway
    harness:platformApiKey:
        value: '......'

```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    harness:accountId:
        value: '....'
    harness:endpoint:
        value: https://app.harness.io/gateway
    harness:platformApiKey:
        value: '......'

```

{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    harness:accountId:
        value: '....'
    harness:endpoint:
        value: https://app.harness.io/gateway
    harness:platformApiKey:
        value: '......'

```

{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
config:
    harness:accountId:
        value: '....'
    harness:endpoint:
        value: https://app.harness.io/gateway
    harness:platformApiKey:
        value: '......'

```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    harness:accountId:
        value: '....'
    harness:endpoint:
        value: https://app.harness.io/gateway
    harness:platformApiKey:
        value: '......'

```

{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

- `accountId` (String) The Harness account id. This can also be set using the `HARNESS_ACCOUNT_ID` environment variable.
- `apiKey` (String) The Harness API key. This can also be set using the `HARNESS_API_KEY` environment variable. For more information to create an API key in FirstGen, see <https://docs.harness.io/article/smloyragsm-api-keys#create_an_api_key>.
- `endpoint` (String) The URL of the Harness API endpoint. The default is `https://app.harness.io/gateway`. This can also be set using the `HARNESS_ENDPOINT` environment variable.
- `platformApiKey` (String) The API key for the Harness next gen platform. This can also be set using the `HARNESS_PLATFORM_API_KEY` environment variable. For more information to create an API key in NextGen, see <https://docs.harness.io/article/tdoad7xrh9-add-and-manage-api-keys>.