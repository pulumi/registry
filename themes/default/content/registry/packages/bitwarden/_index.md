---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/maxlaverse/bitwarden/0.24.0/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Bitwarden Provider
meta_desc: Provides an overview on how to configure the Pulumi Bitwarden provider.
layout: package
---

## Generate Provider

The Bitwarden provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider maxlaverse/bitwarden
```
## Overview

Use the Bitwarden provider to manage your [Password Manager](https://bitwarden.com/products/personal/) Logins and Secure Notes, and [Secrets Manager](https://bitwarden.com/products/secrets-manager/) Secrets.
You must configure the provider with proper credentials before you can use it.
If you're not using `clientImplementation = "embedded"`, you also need a [Password Manager CLI](https://bitwarden.com/help/article/cli/#download-and-install) or [Secrets Manager CLI](https://bitwarden.com/help/secrets-manager-cli/) installed locally.
To read credentials without storing them in plan or state, use ephemeral resources.

> **Note:**  or OpenTofu 1.12 or later (Protocol 6).
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    bitwarden:email:
        value: pulumi@example.com

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
    bitwarden:email:
        value: pulumi@example.com

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
    bitwarden:email:
        value: pulumi@example.com

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
    bitwarden:email:
        value: pulumi@example.com

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
    bitwarden:email:
        value: pulumi@example.com

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
    bitwarden:email:
        value: pulumi@example.com

```
```java
Example currently unavailable in this language
```
{{% /choosable %}}
{{< /chooser >}}
## Guides

* Authentication — Password Manager and Secrets Manager credentials, including provider aliases.
* Client Implementation — official CLIs versus the embedded client, and where data is stored locally.
* Using Ephemeral Resources — read secrets and items without persisting them in plan or state.
## Configuration
Configuration for the Bitwarden Provider can be derived from two sources:
* Parameters in the provider configuration
* Environment variables
### Parameters
Credentials can be provided by adding a combination of `email`, `masterPassword`, `clientId`, `clientSecret`, `accessToken` or `sessionKey` to the bitwarden provider configuration.
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    bitwarden:clientId:
        value: my-client-id
    bitwarden:clientSecret:
        value: my-client-secret
    bitwarden:email:
        value: pulumi@example.com
    bitwarden:masterPassword:
        value: my-master-password

```
### Environment variables
Credentials can be provided by using a combination of `BW_EMAIL`, `BW_PASSWORD`, `BW_CLIENTID`, `BW_CLIENTSECRET`, `BWS_ACCESS_TOKEN` or `BW_SESSION` environment variables.

For example:
```bitwarden
provider "bitwarden" {}
```

```console
export BW_EMAIL="pulumi@example.com"
export BW_PASSWORD="my-master-password"
export BW_CLIENTID="my-client-id"
export BW_CLIENTSECRET="my-client-secret"
```
## Configuration Reference

- `accessToken` (String, Sensitive) Machine Account Access Token (env: `BWS_ACCESS_TOKEN`)).
- `clientId` (String) Client ID (env: `BW_CLIENTID`)
- `clientImplementation` (String) Client implementation type. Valid values are "embedded" (use embedded client) or "cli" (use CLI binaries, default).
- `clientSecret` (String, Sensitive) Client Secret (env: `BW_CLIENTSECRET`). Do not commit this information in Git unless you know what you're doing. Prefer using a Pulumi `variable {}` in order to inject this value from the environment.
- `email` (String) Login Email of the Vault (env: `BW_EMAIL`).
- `experimental` (Block Set) Enable experimental features. (see below for nested schema)
- `extraCaCerts` (String) Extends the well known 'root' CAs (like VeriSign) with the extra certificates in file (env: `NODE_EXTRA_CA_CERTS`, doesn't work with embedded client).
- `masterPassword` (String, Sensitive) Master password of the Vault (env: `BW_PASSWORD`). Do not commit this information in Git unless you know what you're doing. Prefer using a Pulumi `variable {}` in order to inject this value from the environment.
- `server` (String) Bitwarden Server URL (default: `https://vault.bitwarden.com`, env: `BW_URL` or `BWS_SERVER_URL`).
- `sessionKey` (String, Sensitive) A Bitwarden Session Key (env: `BW_SESSION`)
- `vaultPath` (String) Alternative directory for storing the Vault locally (default: `.bitwarden/`, env: `BITWARDENCLI_APPDATA_DIR`; set to empty string to use CLI default).

<a id="nestedblock--experimental"></a>
### Nested Schema for `experimental`

Optional:

- `disableSyncAfterWriteVerification` (Boolean) Skip verification of server-side modifications (like timestamp updates) after write operations - useful when the Bitwarden server makes minor, non-functional changes to objects.
- `embeddedClient` (Boolean, Deprecated) Use the embedded client instead of an external binary.