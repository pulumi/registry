---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-junipermist/v0.5.4/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
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

It is mainly focusing on day 0 and day 1 operations (provisioning and deployment) but will be completed over time.

Use the navigation tree to the left to read about the available resources and functions.

It is possible to use API Token or Username/Password authentication (without 2FA), but only one method should be configured.
## Supported Mist Clouds

This provider can be used with the following Mist Clouds:
* Global 01 (api.mist.com)
* Global 02 (api.gc1.mist.com)
* Global 03 (api.ac2.mist.com)
* Global 04 (api.gc2.mist.com)
* EMEA 01 (api.eu.mist.com)
* EMEA 02 (api.gc3.mist.com)
* EMEA 03 (api.ac6.mist.com)
* APAC 01 (api.ac5.mist.com)
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

- `apiDebug` (Boolean) Flag to enable debugging API calls. Default is false.
- `apiTimeout` (Number) Timeout in seconds for completing API transactions with the Mist Cloud. Omit for default value of 10 seconds. Value of 0 results in infinite timeout.
- `apitoken` (String, Sensitive) For API Token authentication, the Mist API Token.
- `host` (String) URL of the Mist Cloud, e.g. `api.mist.com`.
- `password` (String, Sensitive) For username/password authentication, the Mist Account password.
- `proxy` (String) Requests use the configured proxy to reach the Mist Cloud.
  The value may be either a complete URL or a `[username:password@]host[:port]`, in which case the `http` scheme is assumed. The schemes `http`, `https`, and `socks5` are supported.
- `username` (String) For username/password authentication, the Mist Account username.
### Environment Variables

|   Varibale Name    | Provider attribute |  Type  |                                                                                                                  Description                                                                                                                   |
|--------------------|--------------------|--------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `MIST_HOST`        | `host`             | String | URL of the Mist Cloud, e.g. `api.mist.com`. See above for the list of supported Clouds.                                                                                                                                                        |
| `MIST_API_TOKEN`   | `apitoken`         | String | For API Token authentication, the Mist API Token.                                                                                                                                                                                              |
| `MIST_USERNAME`    | `username`         | String | For username/password authentication, the Mist Account password.                                                                                                                                                                               |
| `MIST_PASSWORD`    | `password`         | String | For username/password authentication, the Mist Account password.                                                                                                                                                                               |
| `MIST_PROXY`       | `proxy`            | String | Requests use the configured proxy to reach the Mist Cloud. The value may be either a complete URL or a `[username:password@]host[:port]`, in which case the `http` scheme is assumed. The schemes `http`, `https`, and `socks5` are supported. |
| `MIST_API_TIMEOUT` | `apiTimeout`      | Int    | Timeout in seconds for completing API transactions with the Mist Cloud. Omit for default value of 10 seconds. Value of 0 results in infinite timeout.                                                                                          |