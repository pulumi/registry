---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/qdrant/qdrant-cloud/1.15.1/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Qdrant-Cloud Provider
meta_desc: Provides an overview on how to configure the Pulumi Qdrant-Cloud provider.
layout: package
---

## Generate Provider

The Qdrant-Cloud provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider qdrant/qdrant-cloud
```
## Overview

The Pulumi Qdrant Cloud provider is a plugin for Pulumi that allows for the full lifecycle management of Qdrant Cloud resources.
Note that it's not intended to manage the content of the database itself.

Below is a sample that creates a [vector database] cluster and a token for accessing the cluster.

The ID and version of the cluster as well the URL (endpoint of the database cluster) are displayed (as output).

The access token created is scoped for the created cluster and is displayed as well.
Note that this token should be kept secret, with this token the database can be manipulated and viewed (CRUD operations).

To view the cluster itself, please visit (in a web-browser)
`<cluster_url>:6333/dashboard`
This will ask the token to grant you access.

The url and token can be used in client libraries as well

For more info, please visit <https://qdrant.tech/>

---
## Versioning and Compatibility

Please always use the latest version of the provider.

Versions below `1.1.0` are deprecated and should not be used.

---
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    qdrant-cloud:accountId:
        value: ""
    qdrant-cloud:apiKey:
        value: ""
    qdrant-cloud:apiUrl:
        value: ""

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
    qdrant-cloud:accountId:
        value: ""
    qdrant-cloud:apiKey:
        value: ""
    qdrant-cloud:apiUrl:
        value: ""

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
    qdrant-cloud:accountId:
        value: ""
    qdrant-cloud:apiKey:
        value: ""
    qdrant-cloud:apiUrl:
        value: ""

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
    qdrant-cloud:accountId:
        value: ""
    qdrant-cloud:apiKey:
        value: ""
    qdrant-cloud:apiUrl:
        value: ""

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
    qdrant-cloud:accountId:
        value: ""
    qdrant-cloud:apiKey:
        value: ""
    qdrant-cloud:apiUrl:
        value: ""

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
    qdrant-cloud:accountId:
        value: ""
    qdrant-cloud:apiKey:
        value: ""
    qdrant-cloud:apiUrl:
        value: ""

```
```java
Example currently unavailable in this language
```
{{% /choosable %}}
{{< /chooser >}}

The Qdrant Cloud API Key and Account ID can also be provided using these environment variables:

* `QDRANT_CLOUD_API_KEY`
* `QDRANT_CLOUD_ACCOUNT_ID`
## Configuration Reference
### Required

- `apiKey` (String, Sensitive) The API Key for Qdrant Cloud API operations.

- `accountId` (String) Default Account Identifier for the Qdrant cloud
- `apiUrl` (String) The URL of the Qdrant Cloud API.
- `insecure` (Boolean) Allow insecure gRPC connections. This is useful for development environments with self-signed certificates. Defaults to false.