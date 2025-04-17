---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/maxlaverse/bitwarden/0.13.6/index.md
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
If you're not trying out the experimental `embeddedClient` feature, you also need a [Bitwarden CLI](https://bitwarden.com/help/article/cli/#download-and-install) installed locally.
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
## Authentication
Depending on the type of credentials you use, you'll be able to connect either with a Password Manager or Secret Manager.
If you want your workspace to interact with both, have a look at [provider aliases](https://developer.pulumi.com/pulumi/language/providers/configuration#alias-multiple-provider-configurations).
### Password Manager
The Password Manager accepts different combinations of credentials to authenticate:
* *[Personal API Key](https://bitwarden.com/help/personal-api-key/)* (requires `masterPassword`, `clientId` and `clientSecret` to be set).
* *Email and Password* (requires `email` and `masterPassword` to be set) (prefer *Personal API keys* instead).
* User-provided *Session Key* (requires `sessionKey` to be set), which only works with a pre-downloaded Vault (See *Generating a Session Key*).
#### Generating a Client ID and Secret
The recommended way to interact with your Password Manager Vault using the Bitwarden Provider Pulumi plugin is to generate an API key.
This allows you to easily revoke access to your Vault without having to change your master password.

In order to generate a pair of Client ID and Secret, you need to:
1. Connect to your Vault on <https://vault.bitwarden.com>, or your self-hosted instance
2. Click on *Settings* and then *My Account*
3. Scroll down to the *API Key* section
4. Click on *View API Key* (or maybe another label if it's the first time)
5. Save the API credentials somewhere safe
#### Generating a Session Key

If you don't want to use an API key, you can use a Session Key instead.
When doing so, it's your responsibility to:
* ensure the validity of the Session Key
* keep the Session Key safe
* revoke it when you don't need it anymore

You can generate a Session Key by running the following command in your Pulumi Stack:
```console
BITWARDENCLI_APPDATA_DIR=.bitwarden bw login

# or if you use a custom vault path
BITWARDENCLI_APPDATA_DIR=<vault_path> bw login
```

A Session Key is bound to a local copy of a Vault. It's therefore important that you set the right `BITWARDENCLI_APPDATA_DIR` to the path where your Vault is stored.
### Secrets Manager
The Secrets Manager only accepts [Access Tokens](https://bitwarden.com/help/access-tokens/) (requires `accessToken` to be set).

In order to generate an Access Token you need to:
1. Connect to your Vault on <https://vault.bitwarden.com>
2. Ensure the *Secrets Manager* section is selected (bottom left)
3. Click on *Machine accounts*
4. Click on *New*
5. Click on your generated Machine Account
6. Select the *Access Tokens* tab
7. Created a new Access Token and save it somewhere safe
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

- `accessToken` (String) Machine Account Access Token (env: `BWS_ACCESS_TOKEN`)).
- `clientId` (String) Client ID (env: `BW_CLIENTID`)
- `clientSecret` (String) Client Secret (env: `BW_CLIENTSECRET`). Do not commit this information in Git unless you know what you're doing. Prefer using a Pulumi `variable {}` in order to inject this value from the environment.
- `email` (String) Login Email of the Vault (env: `BW_EMAIL`).
- `experimental` (Block Set) Enable experimental features. (see below for nested schema)
- `extraCaCerts` (String) Extends the well known 'root' CAs (like VeriSign) with the extra certificates in file (env: `NODE_EXTRA_CA_CERTS`, doesn't work with embedded client).
- `masterPassword` (String) Master password of the Vault (env: `BW_PASSWORD`). Do not commit this information in Git unless you know what you're doing. Prefer using a Pulumi `variable {}` in order to inject this value from the environment.
- `server` (String) Bitwarden Server URL (default: `https://vault.bitwarden.com`, env: `BW_URL`).
- `sessionKey` (String) A Bitwarden Session Key (env: `BW_SESSION`)
- `vaultPath` (String) Alternative directory for storing the Vault locally (default: `.bitwarden/`, env: `BITWARDENCLI_APPDATA_DIR`).

<a id="nestedblock--experimental"></a>
### Nested Schema for `experimental`

Optional:

- `disableSyncAfterWriteVerification` (Boolean) Skip verification of server-side modifications (like timestamp updates) after write operations - useful when the Bitwarden server makes minor, non-functional changes to objects.
- `embeddedClient` (Boolean) Use the embedded client instead of an external binary.