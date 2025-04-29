---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-kong/v4.5.9/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Kong Provider
meta_desc: Provides an overview on how to configure the Pulumi Kong provider.
layout: package
---
## Installation

The Kong provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/kong`](https://www.npmjs.com/package/@pulumi/kong)
* Python: [`pulumi-kong`](https://pypi.org/project/pulumi-kong/)
* Go: [`github.com/pulumi/pulumi-kong/sdk/v4/go/kong`](https://github.com/pulumi/pulumi-kong)
* .NET: [`Pulumi.Kong`](https://www.nuget.org/packages/Pulumi.Kong)
* Java: [`com.pulumi/kong`](https://central.sonatype.com/artifact/com.pulumi/kong)
## Overview

The Kong Pulumi Provider tested against real Kong (using Docker)!

Pulumi provider tested to work against Kong 2.X.
## Usage

To configure the provider:
{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    kong:kongAdminUri:
        value: http://localhost:8001

```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    kong:kongAdminUri:
        value: http://localhost:8001

```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    kong:kongAdminUri:
        value: http://localhost:8001

```

{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    kong:kongAdminUri:
        value: http://localhost:8001

```

{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
config:
    kong:kongAdminUri:
        value: http://localhost:8001

```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    kong:kongAdminUri:
        value: http://localhost:8001

```

{{% /choosable %}}
{{< /chooser >}}

Optionally you can configure Username and Password for BasicAuth:
{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    kong:kongAdminPassword:
        value: yourpass
    kong:kongAdminUri:
        value: http://localhost:8001
    kong:kongAdminUsername:
        value: youruser

```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    kong:kongAdminPassword:
        value: yourpass
    kong:kongAdminUri:
        value: http://localhost:8001
    kong:kongAdminUsername:
        value: youruser

```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    kong:kongAdminPassword:
        value: yourpass
    kong:kongAdminUri:
        value: http://localhost:8001
    kong:kongAdminUsername:
        value: youruser

```

{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    kong:kongAdminPassword:
        value: yourpass
    kong:kongAdminUri:
        value: http://localhost:8001
    kong:kongAdminUsername:
        value: youruser

```

{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
config:
    kong:kongAdminPassword:
        value: yourpass
    kong:kongAdminUri:
        value: http://localhost:8001
    kong:kongAdminUsername:
        value: youruser

```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    kong:kongAdminPassword:
        value: yourpass
    kong:kongAdminUri:
        value: http://localhost:8001
    kong:kongAdminUsername:
        value: youruser

```

{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

In addition to generic provider arguments (e.g. alias and version), the following arguments are supported in the Kong provider configuration:

* `kongAdminUri` - (Required) The URI of the Kong admin API, can be sourced from the `KONG_ADMIN_ADDR` environment variable
* `kongAdminUsername` - (Optional) The username for the Kong admin API if set, can be sourced from the `KONG_ADMIN_USERNAME` environment variable
* `kongAdminPassword` - (Optional) The password for the Kong admin API if set, can be sourced from the `KONG_ADMIN_PASSWORD` environment variable
* `tlsSkipVerify` - (Optional) Whether to skip TLS certificate verification for the kong api when using https, can be sourced from the `TLS_SKIP_VERIFY` environment variable
* `kongApiKey` - (Optional) API key used to secure the kong admin API, can be sourced from the `KONG_API_KEY` environment variable
* `kongAdminToken` - (Optional) API key used to secure the kong admin API in the Enterprise Edition, can be sourced from the `KONG_ADMIN_TOKEN` environment variable
* `kongWorkspace` - (Optional) Workspace context (Enterprise Edition)
* `strictPluginsMatch` - (Optional) Should plugins `configJson` field strictly match plugin configuration