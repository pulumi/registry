---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-ec/v0.10.7/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: ElasticCloud (EC) Provider
meta_desc: Provides an overview on how to configure the Pulumi ElasticCloud (EC) provider.
layout: package
---
## Installation

The ElasticCloud (EC) provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/ec`](https://www.npmjs.com/package/@pulumi/ec)
* Python: [`pulumi-ec`](https://pypi.org/project/pulumi-ec/)
* Go: [`github.com/pulumi/pulumi-ec/sdk/go/ec`](https://github.com/pulumi/pulumi-ec)
* .NET: [`Pulumi.Ec`](https://www.nuget.org/packages/Pulumi.Ec)
* Java: [`com.pulumi/ec`](https://central.sonatype.com/artifact/com.pulumi/ec)
## Overview

---
## page_title: "Provider: Elastic Cloud"
The Elastic Cloud Pulumi provider can be used to configure and manage Elastic Cloud deployments using the Elastic Cloud
APIs. Use the navigation to the left to read about functions and resources supported by the Elastic Cloud provider. Elastic Cloud APIs are available for:

* Elasticsearch Service (ESS).
* Elastic Cloud Enterprise (ECE).
* Elasticsearch Service Private (ESSP).
## Authentication

The Elastic Cloud Pulumi provider offers two methods of authentication against the remote API: `apikey` or a combination of `username` and `password`. Depending on the environment, you may choose one over the other. The Public API of Elasticsearch Service (ESS) is the default `endpoint` that the provider will target.

Elasticsearch Service (ESS) only supports `apikey`. Elastic Cloud Enterprise (ECE) supports `apikey` or a combination of `username` and `password`.

!> **Warning:** Hard-coding credentials into a Pulumi configuration is not recommended, and risks secret leakage should this file ever be committed to a public version control system.
### API key authentication (recommended)

API keys are the recommended authentication method. They can be used to authenticate against Elasticsearch Service or Elastic Cloud Enterprise.
#### Generating an Elasticsearch Service (ESS) API Key

To generate an API key, follow these steps:

1. Open you browser and navigate to <https://cloud.elastic.co/login>.
2. Log in with your email and password.
3. Click on [Elasticsearch Service](https://cloud.elastic.co/deployments).
4. Navigate to [Features > API Keys](https://cloud.elastic.co/deployment-features/keys) and click on **Generate API Key**.
5. Choose a name for your API key.
6. Save your API key somewhere.
#### Using your API Key on the Elastic Cloud pulumi provider

After you've generated your API Key, you can make it available to the Pulumi provider by exporting it as the environment variable `EC_API_KEY` (recommended), or hardcoded in the provider `.tf` configuration file (supported but not recommended).

```sh
$ export EC_API_KEY="<apikey value>"
```

Or set the `apikey` field in the "ec" provider to the value of your generated API key.

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    ec:apikey:
        value: <apikey value>

```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    ec:apikey:
        value: <apikey value>

```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    ec:apikey:
        value: <apikey value>

```

{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    ec:apikey:
        value: <apikey value>

```

{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
config:
    ec:apikey:
        value: <apikey value>

```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    ec:apikey:
        value: <apikey value>

```

{{% /choosable %}}
{{< /chooser >}}
### Username and password login (ECE)

If you are targeting an ECE environment, you can also use a combination of `username` and `password` as authentication method.

They can either be hardcoded in the provider `.tf` configuration (not recommended), or specified with the following environment variables: `EC_USERNAME` or `EC_USER` and `EC_PASSWORD` or `EC_PASS`.

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    ec:endpoint:
        value: https://my.ece-environment.corp
    ec:insecure:
        value: true
    ec:password:
        value: my-password
    ec:username:
        value: my-username

```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    ec:endpoint:
        value: https://my.ece-environment.corp
    ec:insecure:
        value: true
    ec:password:
        value: my-password
    ec:username:
        value: my-username

```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    ec:endpoint:
        value: https://my.ece-environment.corp
    ec:insecure:
        value: true
    ec:password:
        value: my-password
    ec:username:
        value: my-username

```

{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    ec:endpoint:
        value: https://my.ece-environment.corp
    ec:insecure:
        value: true
    ec:password:
        value: my-password
    ec:username:
        value: my-username

```

{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
config:
    ec:endpoint:
        value: https://my.ece-environment.corp
    ec:insecure:
        value: true
    ec:password:
        value: my-password
    ec:username:
        value: my-username

```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    ec:endpoint:
        value: https://my.ece-environment.corp
    ec:insecure:
        value: true
    ec:password:
        value: my-password
    ec:username:
        value: my-username

```

{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

- `apikey` (String, Sensitive) API Key to use for API authentication. The only valid authentication mechanism for the Elasticsearch Service.
- `endpoint` (String) Endpoint where the pulumi provider will point to. Defaults to "https://api.elastic-cloud.com".
- `insecure` (Boolean) Allow the provider to skip TLS validation on its outgoing HTTP calls.
- `password` (String, Sensitive) Password to use for API authentication. Available only when targeting ECE Installations or Elasticsearch Service Private.
- `timeout` (String) Timeout used for individual HTTP calls. Defaults to "1m".
- `username` (String) Username to use for API authentication. Available only when targeting ECE Installations or Elasticsearch Service Private.
- `verbose` (Boolean) When set, a "request.log" file will be written with all outgoing HTTP requests. Defaults to "false".
- `verboseCredentials` (Boolean) When set with verbose, the contents of the Authorization header will not be redacted. Defaults to "false".
- `verboseFile` (String) Timeout used for individual HTTP calls. Defaults to "1m".