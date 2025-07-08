---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-splunk/v1.2.19/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Splunk Provider
meta_desc: Provides an overview on how to configure the Pulumi Splunk provider.
layout: package
---
## Installation

The Splunk provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/splunk`](https://www.npmjs.com/package/@pulumi/splunk)
* Python: [`pulumi-splunk`](https://pypi.org/project/pulumi-splunk/)
* Go: [`github.com/pulumi/pulumi-splunk/sdk/go/splunk`](https://github.com/pulumi/pulumi-splunk)
* .NET: [`Pulumi.Splunk`](https://www.nuget.org/packages/Pulumi.Splunk)
* Java: [`com.pulumi/splunk`](https://central.sonatype.com/artifact/com.pulumi/splunk)
## Overview

The Splunk provider can interact with the resources supported by Splunk. The provider needs to be configured with the proper credentials before it can be used.
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    splunk:insecureSkipVerify:
        value: true
    splunk:password:
        value: changeme
    splunk:url:
        value: localhost:8089
    splunk:username:
        value: admin

```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    splunk:insecureSkipVerify:
        value: true
    splunk:password:
        value: changeme
    splunk:url:
        value: localhost:8089
    splunk:username:
        value: admin

```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    splunk:insecureSkipVerify:
        value: true
    splunk:password:
        value: changeme
    splunk:url:
        value: localhost:8089
    splunk:username:
        value: admin

```

{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    splunk:insecureSkipVerify:
        value: true
    splunk:password:
        value: changeme
    splunk:url:
        value: localhost:8089
    splunk:username:
        value: admin

```

{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
config:
    splunk:insecureSkipVerify:
        value: true
    splunk:password:
        value: changeme
    splunk:url:
        value: localhost:8089
    splunk:username:
        value: admin

```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    splunk:insecureSkipVerify:
        value: true
    splunk:password:
        value: changeme
    splunk:url:
        value: localhost:8089
    splunk:username:
        value: admin

```

{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

Below arguments for the provider can also be set as environment variables.

* `url` or `SPLUNK_URL` - (Required) The URL for the Splunk instance to be configured. (The provider uses `https` as the default schema as prefix to the URL)
* `username` or `SPLUNK_USERNAME`  - (Optional) The username to access the Splunk instance to be configured.
* `password` or `SPLUNK_PASSWORD` - (Optional) The password to access the Splunk instance to be configured.
* `authToken` or `SPLUNK_AUTH_TOKEN` - (Optional) Use auth token instead of username and password to configure Splunk instance.
  If specified, auth token takes priority over username/password.
* `insecureSkipVerify` or `SPLUNK_INSECURE_SKIP_VERIFY` - (Optional) Insecure skip verification flag (Defaults to `true`)
* `timeout` or `SPLUNK_TIMEOUT` -  (Optional) Timeout when making calls to Splunk server. (Defaults to `60 seconds`)

(NOTE: Auth token can only be used with certain type of Splunk deployments.
Read more on authentication with tokens here: <https://docs.splunk.com/Documentation/Splunk/latest/Security/Setupauthenticationwithtokens>)