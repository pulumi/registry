---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-scm/v0.4.0/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Strata Cloud Manager Provider
meta_desc: Provides an overview on how to configure the Pulumi Strata Cloud Manager provider.
layout: package
---
## Installation

The Strata Cloud Manager provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/scm`](https://www.npmjs.com/package/@pulumi/scm)
* Python: [`pulumi-scm`](https://pypi.org/project/pulumi-scm/)
* Go: [`github.com/pulumi/pulumi-scm/sdk/go/scm`](https://github.com/pulumi/pulumi-scm)
* .NET: [`Pulumi.Scm`](https://www.nuget.org/packages/Pulumi.Scm)
* Java: [`com.pulumi/scm`](https://central.sonatype.com/artifact/com.pulumi/scm)
## Overview

The `scm` provider provides resources and functions to manage and query Strata Cloud Manager.

This provider covers the following aspects of Strata Cloud Manager:
* Unified Networking Security
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    scm:clientId:
        value: your-id@12345
    scm:clientSecret:
        value: secret
    scm:host:
        value: api.strata.paloaltonetworks.com
    scm:scope:
        value: tsg_id:12345

```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    scm:clientId:
        value: your-id@12345
    scm:clientSecret:
        value: secret
    scm:host:
        value: api.strata.paloaltonetworks.com
    scm:scope:
        value: tsg_id:12345

```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    scm:clientId:
        value: your-id@12345
    scm:clientSecret:
        value: secret
    scm:host:
        value: api.strata.paloaltonetworks.com
    scm:scope:
        value: tsg_id:12345

```

{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    scm:clientId:
        value: your-id@12345
    scm:clientSecret:
        value: secret
    scm:host:
        value: api.strata.paloaltonetworks.com
    scm:scope:
        value: tsg_id:12345

```

{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
config:
    scm:clientId:
        value: your-id@12345
    scm:clientSecret:
        value: secret
    scm:host:
        value: api.strata.paloaltonetworks.com
    scm:scope:
        value: tsg_id:12345

```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    scm:clientId:
        value: your-id@12345
    scm:clientSecret:
        value: secret
    scm:host:
        value: api.strata.paloaltonetworks.com
    scm:scope:
        value: tsg_id:12345

```

{{% /choosable %}}
{{< /chooser >}}
## Provider Parameter Priority

There are multiple ways to specify the provider's parameters.  If overlapping values are configured for the provider, then this is the resolution order:

1. Statically configured in the provider configuration
2. Environment variable (where applicable)
3. Taken from the JSON config file
## Configuration Reference

- `authFile` (String) The file path to the JSON file with auth creds for SCM.
- `authUrl` (String) The URL to send auth credentials to which will return a JWT. Default: `https://auth.apps.paloaltonetworks.com/auth/v1/oauth2/access_token`. Environment variable: `SCM_AUTH_URL`. JSON config file variable: `authUrl`.
- `clientId` (String) The client ID for the connection. Environment variable: `SCM_CLIENT_ID`. JSON config file variable: `clientId`.
- `clientSecret` (String, Sensitive) The client secret for the connection. Environment variable: `SCM_CLIENT_SECRET`. JSON config file variable: `clientSecret`.
- `headers` (Map of String) Custom HTTP headers to be sent with all API commands. Environment variable: `SCM_HEADERS`. JSON config file variable: `headers`.
- `host` (String) The hostname of Strata Cloud Manager API. Default: `api.sase.paloaltonetworks.com`. Environment variable: `SCM_HOST`. JSON config file variable: `host`.
- `logging` (String) The logging level of the provider and the underlying communication. Default: `quiet`. Environment variable: `SCM_LOGGING`. JSON config file variable: `logging`.
- `port` (Number) The port number to use for API commands, if non-standard for the given protocol. Environment variable: `SCM_PORT`. JSON config file variable: `port`.
- `protocol` (String) The protocol to use for SCM. This should be 'http' or 'https'. Default: `https`. Environment variable: `SCM_PROTOCOL`. JSON config file variable: `protocol`.
- `scope` (String) The client scope. Environment variable: `SCM_SCOPE`. JSON config file variable: `scope`.