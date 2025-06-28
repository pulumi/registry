---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-ise/v0.2.4/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Cisco ISE Provider
meta_desc: Provides an overview on how to configure the Pulumi Cisco ISE provider.
layout: package
---
## Installation

The Cisco ISE provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/ise`](https://www.npmjs.com/package/@pulumi/ise)
* Python: [`pulumi-ise`](https://pypi.org/project/pulumi-ise/)
* Go: [`github.com/pulumi/pulumi-ise/sdk/go/ise`](https://github.com/pulumi/pulumi-ise)
* .NET: [`Pulumi.Ise`](https://www.nuget.org/packages/Pulumi.Ise)
* Java: [`com.pulumi/ise`](https://central.sonatype.com/artifact/com.pulumi/ise)
## Overview

The ISE provider provides resources to interact with a Cisco ISE (Identity Service Engine) instance. It communicates with ISE via the REST API.

This provider uses both, the ERS and Open API. Instructions on how to enable API access can be found here: <https://developer.cisco.com/docs/identity-services-engine/latest/#!cisco-ise-api-framework>

All resources and functions have been tested with the following releases.

| Platform |    Version    |
|----------|---------------|
| ISE      | 3.2.0 Patch 4 |
| ISE      | 3.3.0         |
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    ise:password:
        value: password
    ise:url:
        value: https://10.1.1.1
    ise:username:
        value: admin

```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    ise:password:
        value: password
    ise:url:
        value: https://10.1.1.1
    ise:username:
        value: admin

```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    ise:password:
        value: password
    ise:url:
        value: https://10.1.1.1
    ise:username:
        value: admin

```

{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    ise:password:
        value: password
    ise:url:
        value: https://10.1.1.1
    ise:username:
        value: admin

```

{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
config:
    ise:password:
        value: password
    ise:url:
        value: https://10.1.1.1
    ise:username:
        value: admin

```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    ise:password:
        value: password
    ise:url:
        value: https://10.1.1.1
    ise:username:
        value: admin

```

{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

- `insecure` (Boolean) Allow insecure HTTPS client. This can also be set as the ISE_INSECURE environment variable. Defaults to `true`.
- `password` (String, Sensitive) Password for the ISE instance. This can also be set as the ISE_PASSWORD environment variable.
- `retries` (Number) Number of retries for REST API calls. This can also be set as the ISE_RETRIES environment variable. Defaults to `3`.
- `url` (String) URL of the Cisco ISE instance. This can also be set as the ISE_URL environment variable.
- `username` (String) Username for the ISE instance. This can also be set as the ISE_USERNAME environment variable.