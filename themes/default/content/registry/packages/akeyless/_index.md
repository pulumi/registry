---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/akeyless-community/akeyless/1.9.0/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Akeyless Provider
meta_desc: Provides an overview on how to configure the Pulumi Akeyless provider.
layout: package
---

## Generate Provider

The Akeyless provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider akeyless-community/akeyless
```
## Overview

The Akeyless provider provides resources to interact with Akeyless services.
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    akeyless:apiGatewayAddress:
        value: https://api.akeyless.io

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
    akeyless:apiGatewayAddress:
        value: https://api.akeyless.io

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
    akeyless:apiGatewayAddress:
        value: https://api.akeyless.io

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
    akeyless:apiGatewayAddress:
        value: https://api.akeyless.io

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
    akeyless:apiGatewayAddress:
        value: https://api.akeyless.io

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
    akeyless:apiGatewayAddress:
        value: https://api.akeyless.io

```
```java
Example currently unavailable in this language
```
{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

- `apiGatewayAddress` (String) Origin URL of the API Gateway server. This is a URL with a scheme, a hostname and a port.
- `apiKeyLogin` (Block List) A configuration block, described below, that attempts to authenticate using API-Key. (see below for nested schema)
- `awsIamLogin` (Block List) A configuration block, described below, that attempts to authenticate using AWS-IAM authentication credentials. (see below for nested schema)
- `azureAdLogin` (Block List) A configuration block, described below, that attempts to authenticate using Azure Active Directory authentication. (see below for nested schema)
- `certLogin` (Block List) A configuration block, described below, that attempts to authenticate using Certificate authentication.  The Certificate and the Private key can be provided as a command line variable or it will be pulled out of an environment variable named AKEYLESS_AUTH_CERT and AKEYLESS_AUTH_KEY. (see below for nested schema)
- `emailLogin` (Block List) A configuration block, described below, that attempts to authenticate using email and password. (see below for nested schema)
- `gcpLogin` (Block List) A configuration block, described below, that attempts to authenticate using GCP-IAM authentication credentials. (see below for nested schema)
- `jwtLogin` (Block List) A configuration block, described below, that attempts to authenticate using JWT authentication.  The JWT can be provided as a command line variable or it will be pulled out of an environment variable named AKEYLESS_AUTH_JWT. (see below for nested schema)
- `tokenLogin` (Block List) A configuration block, described below, that attempts to authenticate using akeyless token. The token can be provided as a command line variable or it will be pulled out of an environment variable named AKEYLESS_AUTH_TOKEN. (see below for nested schema)
- `uidLogin` (Block List) A configuration block, described below, that attempts to authenticate using Universal Identity authentication. (see below for nested schema)

<a id="nestedblock--api_key_login"></a>
### Nested Schema for `apiKeyLogin`

Optional:

- `accessId` (String)
- `accessKey` (String, Sensitive)

<a id="nestedblock--aws_iam_login"></a>
### Nested Schema for `awsIamLogin`

Required:

- `accessId` (String)

<a id="nestedblock--azure_ad_login"></a>
### Nested Schema for `azureAdLogin`

Required:

- `accessId` (String)

<a id="nestedblock--cert_login"></a>
### Nested Schema for `certLogin`

Required:

- `accessId` (String)

Optional:

- `certData` (String, Sensitive)
- `certFileName` (String)
- `keyData` (String, Sensitive)
- `keyFileName` (String)

<a id="nestedblock--email_login"></a>
### Nested Schema for `emailLogin`

Required:

- `adminEmail` (String)
- `adminPassword` (String)

<a id="nestedblock--gcp_login"></a>
### Nested Schema for `gcpLogin`

Required:

- `accessId` (String)

Optional:

- `audience` (String)

<a id="nestedblock--jwt_login"></a>
### Nested Schema for `jwtLogin`

Required:

- `accessId` (String)
- `jwt` (String, Sensitive)

<a id="nestedblock--token_login"></a>
### Nested Schema for `tokenLogin`

Required:

- `token` (String, Sensitive)

<a id="nestedblock--uid_login"></a>
### Nested Schema for `uidLogin`

Required:

- `uidToken` (String, Sensitive)

Optional:

- `accessId` (String)