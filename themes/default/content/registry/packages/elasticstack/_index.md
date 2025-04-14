---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/elastic/elasticstack/0.11.14/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Elasticstack Provider
meta_desc: Provides an overview on how to configure the Pulumi Elasticstack provider.
layout: package
---

## Generate Provider

The Elasticstack provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider elastic/elasticstack
```
## Overview

The Elasticstack provider provides the resources to interact with Elastic stack products.

It is recommended to setup at least minimum security, <https://www.elastic.co/guide/en/elasticsearch/reference/current/security-minimal-setup.html>
in order to interact with the Elasticsearch and be able to use the provider's full capabilities

The provider uses Pulumi protocol version 6 that is compatible with Pulumi CLI version 1.0 and later.
## Authentication

The Elasticstack provider offers few different ways of providing credentials for authentication.
The following methods are supported:

* Static credentials
* Environment variables
* Each `elasticsearch` resource supports an `elasticsearchConnection` block, allowing use of the same provider to configure many different clusters at the same time
### Static credentials
#### Elasticsearch

Default static credentials can be provided by adding the `username`, `password` and `endpoints` in the `elasticsearch` block:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs

```
```typescript
import * as pulumi from "@pulumi/pulumi";

```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python

```
```python
import pulumi

```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;

return await Deployment.RunAsync(() =>
{
});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go

```
```go
package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		return nil
	})
}
```
{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml

```
```yaml
{}
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
    }
}
```
{{% /choosable %}}
{{< /chooser >}}

Alternatively an `apiKey` can be specified instead of `username` and `password`:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:

```
#### Kibana

Default static credentials can be provided by adding the `username`, `password` and `endpoints` in the `kibana` block:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:

```

If no credentials are supplied the provider will fall back to using those provided in the `elasticsearch` block.
### Environment Variables

The provider configuration can be specified through environment variables.

For Elasticsearch resources, you can use the following variables:
- `ELASTICSEARCH_USERNAME` - The username to use for Elasticsearch authentication
- `ELASTICSEARCH_PASSWORD` - The password to use for Elasticsearch authentication
- `ELASTICSEARCH_ENDPOINTS` - A comma separated list of Elasticsearch hosts to connect to
- `ELASTICSEARCH_API_KEY` - An Elasticsearch API key to use instead of `ELASTICSEARCH_USERNAME` and `ELASTICSEARCH_PASSWORD`
- `ELASTICSEARCH_BEARER_TOKEN` - A bearer token to use for Elasticsearch authorization header.
- `ELASTICSEARCH_ES_CLIENT_AUTHENTICATION` - The shared secret for the Elasticsearch authorization header.

Kibana resources will re-use any Elasticsearch credentials specified, these may be overridden with the following variables:
- `KIBANA_USERNAME` - The username to use for Kibana authentication
- `KIBANA_PASSWORD` - The password to use for Kibana authentication
- `KIBANA_ENDPOINT` - The Kibana host to connect to
- `KIBANA_API_KEY` - An Elasticsearch API key to use instead of `KIBANA_USERNAME` and `KIBANA_PASSWORD`

Fleet resources will re-use any Kibana or Elasticsearch credentials specified, these may be overridden with the following variables:
- `FLEET_USERNAME` - The username to use for Kibana authentication
- `FLEET_PASSWORD` - The password to use for Kibana authentication
- `FLEET_ENDPOINT` - The Kibana host to connect to. ** Note the Fleet API is hosted within Kibana. This must be a Kibana HTTP host **
- `FLEET_API_KEY` - API key to use for authentication to Fleet

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:

```
### Per resource credentials

See docs related to the specific resources.
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs

```
```typescript
import * as pulumi from "@pulumi/pulumi";

```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python

```
```python
import pulumi

```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;

return await Deployment.RunAsync(() =>
{
});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go

```
```go
package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		return nil
	})
}
```
{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml

```
```yaml
{}
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

- `elasticsearch` (Block List, Max: 1) Elasticsearch connection configuration block. (see below for nested schema)
- `fleet` (Block List, Max: 1) Fleet connection configuration block. (see below for nested schema)
- `kibana` (Block List, Max: 1) Kibana connection configuration block. (see below for nested schema)

<a id="nestedblock--elasticsearch"></a>
### Nested Schema for `elasticsearch`

Optional:

- `apiKey` (String, Sensitive) API Key to use for authentication to Elasticsearch
- `bearerToken` (String, Sensitive) Bearer Token to use for authentication to Elasticsearch
- `caData` (String) PEM-encoded custom Certificate Authority certificate
- `caFile` (String) Path to a custom Certificate Authority certificate
- `certData` (String) PEM encoded certificate for client auth
- `certFile` (String) Path to a file containing the PEM encoded certificate for client auth
- `endpoints` (List of String, Sensitive) A list of endpoints where the pulumi provider will point to, this must include the http(s) schema and port number.
- `esClientAuthentication` (String, Sensitive) ES Client Authentication field to be used with the JWT token
- `insecure` (Boolean) Disable TLS certificate validation
- `keyData` (String, Sensitive) PEM encoded private key for client auth
- `keyFile` (String) Path to a file containing the PEM encoded private key for client auth
- `password` (String, Sensitive) Password to use for API authentication to Elasticsearch.
- `username` (String) Username to use for API authentication to Elasticsearch.

<a id="nestedblock--fleet"></a>
### Nested Schema for `fleet`

Optional:

- `apiKey` (String, Sensitive) API Key to use for authentication to Fleet.
- `caCerts` (List of String) A list of paths to CA certificates to validate the certificate presented by the Fleet server.
- `endpoint` (String, Sensitive) The Fleet server where the pulumi provider will point to, this must include the http(s) schema and port number.
- `insecure` (Boolean) Disable TLS certificate validation
- `password` (String, Sensitive) Password to use for API authentication to Fleet.
- `username` (String) Username to use for API authentication to Fleet.

<a id="nestedblock--kibana"></a>
### Nested Schema for `kibana`

Optional:

- `apiKey` (String, Sensitive) API Key to use for authentication to Kibana
- `caCerts` (List of String) A list of paths to CA certificates to validate the certificate presented by the Kibana server.
- `endpoints` (List of String, Sensitive) A comma-separated list of endpoints where the pulumi provider will point to, this must include the http(s) schema and port number.
- `insecure` (Boolean) Disable TLS certificate validation
- `password` (String, Sensitive) Password to use for API authentication to Kibana.
- `username` (String) Username to use for API authentication to Kibana.