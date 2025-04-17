---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-nomad/v2.5.0/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Nomad Provider
meta_desc: Provides an overview on how to configure the Pulumi Nomad provider.
layout: package
---
## Installation

The Nomad provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/nomad`](https://www.npmjs.com/package/@pulumi/nomad)
* Python: [`pulumi-nomad`](https://pypi.org/project/pulumi-nomad/)
* Go: [`github.com/pulumi/pulumi-nomad/sdk/v2/go/nomad`](https://github.com/pulumi/pulumi-nomad)
* .NET: [`Pulumi.Nomad`](https://www.nuget.org/packages/Pulumi.Nomad)
* Java: [`com.pulumi/nomad`](https://central.sonatype.com/artifact/com.pulumi/nomad)
## Overview

[HashiCorp Nomad](https://www.nomadproject.io) is an application scheduler. The
Nomad provider exposes resources to interact with a Nomad cluster.

Use the navigation to the left to read about the available resources.
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    nomad:address:
        value: http://nomad.mycompany.com:4646
    nomad:region:
        value: us-east-2

```
```typescript
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    nomad:address:
        value: http://nomad.mycompany.com:4646
    nomad:region:
        value: us-east-2

```
```python
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    nomad:address:
        value: http://nomad.mycompany.com:4646
    nomad:region:
        value: us-east-2

```
```csharp
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    nomad:address:
        value: http://nomad.mycompany.com:4646
    nomad:region:
        value: us-east-2

```
```go
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
config:
    nomad:address:
        value: http://nomad.mycompany.com:4646
    nomad:region:
        value: us-east-2

```
```yaml
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    nomad:address:
        value: http://nomad.mycompany.com:4646
    nomad:region:
        value: us-east-2

```
```java
Example currently unavailable in this language
```
{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

The following configuration inputs are supported:

- `address` `(string: "http://127.0.0.1:4646")` - The HTTP(S) API address of the
  Nomad agent. This must include the leading protocol (e.g. `https://`). This
  can also be specified as the `NOMAD_ADDR` environment variable.

- `region` `(string: "")` - The Nomad region to target. This can also be
  specified as the `NOMAD_REGION` environment variable.

- `httpAuth` `(string: "")` - HTTP Basic Authentication credentials to be used
  when communicating with Nomad, in the format of either `user` or `user:pass`.
  This can also be specified using the `NOMAD_HTTP_AUTH` environment variable.

- `caFile` `(string: "")` - A local file path to a PEM-encoded certificate
  authority used to verify the remote agent's certificate. This can also be
  specified as the `NOMAD_CACERT` environment variable.

- `caPem` `(string: "")` - PEM-encoded certificate authority used to verify
  the remote agent's certificate.

- `certFile` `(string: "")` - A local file path to a PEM-encoded certificate
  provided to the remote agent. If this is specified, `keyFile` or `keyPem`
  is also required. This can also be specified as the `NOMAD_CLIENT_CERT`
  environment variable.

- `certPem` `(string: "")` - PEM-encoded certificate provided to the remote
  agent. If this is specified, `keyFile` or `keyPem` is also required.

- `keyFile` `(string: "")` - A local file path to a PEM-encoded private key.
  This is required if `certFile` or `certPem` is specified. This can also be
  specified via the `NOMAD_CLIENT_KEY` environment variable.

- `keyPem` `(string: "")` - PEM-encoded private key. This is required if
  `certFile` or `certPem` is specified.

- `skipVerify` `(boolean: false)` - Set this to true if you want to skip TLS verification on the client side.
  This can also be specified via the `NOMAD_SKIP_VERIFY` environment variable.

- `headers` - (Optional) A configuration block, described below, that provides headers
  to be sent along with all requests to Nomad.  This block can be specified
  multiple times.

- `vaultToken` `(string: "")` - A Vault token used when [submitting the job](https://www.nomadproject.io/docs/job-specification/job#vault_token).
  This can also be specified as the `VAULT_TOKEN` environment variable or using a
  Vault token helper (see [Vault's documentation](https://www.vaultproject.io/docs/commands/token-helper.html)
  for more details). See below for strategies when
  multiple Vault tokens are required.

- `consulToken` `(string: "")` - A Consul token used when [submitting the job](https://www.nomadproject.io/docs/job-specification/job#consul_token).
  This can also be specified as the `CONSUL_HTTP_TOKEN` environment variable.
  See below for strategies when multiple Consul tokens are required.

- `secretId` `(string: "")` - The Secret ID of an ACL token to make requests with,
  for ACL-enabled clusters. This can also be specified via the `NOMAD_TOKEN`
  environment variable.

- `ignoreEnvVars` `(map[string]bool: {})` - A map of environment variables
  that are ignored by the provider when configuring the Nomad API client.
  Supported keys are: `NOMAD_NAMESPACE` and `NOMAD_REGION`.

The `headers` nested type accepts the following arguments:
* `name` - (Required) The name of the header.
* `value` - (Required) The value of the header.