---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/hpe/hpegl/0.4.18/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Hpegl Provider
meta_desc: Provides an overview on how to configure the Pulumi Hpegl provider.
layout: package
---

## Generate Provider

The Hpegl provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider hpe/hpegl
```
## Overview

This is the hpegl (HPE GreenLake) pulumi provider for use in managing HPE GreenLake
services.

Current supported services:
- VMaaS
- Containers
- Metal

This provider requires 64-bit versions of the pulumi binary to work properly.
## API Client

Note that an API client must be used with this provider.  The provider supports two different types
of API Client corresponding to two different version of GreenLake IAM:
- glcs
- glp

The version of IAM used by the provider is determined by the `HPEGL_IAM_VERSION` environment variable.
It can have two values corresponding to the two versions of IAM: `glcs` and `glp`. If the variable is not set,
the provider will default to `glcs`.

Note that most GreenLake installations use the `glcs` version of IAM.  The exceptions are "Disconnected"
installations which use the `glp` version of IAM.

At present only `metal` resources and data-sources are supported with the `glp` version of IAM.
### glcs API Client
For information on how to create a glcs API client see [here](http://www.hpe.com/info/greenlakecentral-create-api-client).

The following env-vars can be used to supply glcs API client creds and related information to
the provider:

```bash
export HPEGL_TENANT_ID=< tenant-id >
export HPEGL_USER_ID=< API client id >
export HPEGL_USER_SECRET=< API client secret >
export HPEGL_IAM_SERVICE_URL=< the "issuer" URL for the API client  >
```
### glp API Client
For information on glp API Clients see [here](https://developer.greenlake.hpe.com/docs/greenlake/services/#configuring-api-client-credentials)

The following env-vars can be used to supply glp API client creds and related information to
the provider:

```bash
export HPEGL_IAM_VERSION=glp
export HPEGL_USER_ID=< API client id >
export HPEGL_USER_SECRET=< API client secret >
export HPEGL_IAM_SERVICE_URL=< the "Token URL" for API clients, can be found on the API Client creation screen  >
```

Note that the `HPEGL_IAM_VERSION` environment variable must be set to `glp` to use the glp API client.
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

- `apiVendedServiceClient` (Boolean) Declare if the API client being used is an API-vended one or not.  Defaults to "true"
  i.e. the client is API-vended.  The value can be set using the HPEGL_API_VENDED_SERVICE_CLIENT env-var.
- `caas` (Block Set, Max: 1) (see below for nested schema)
- `iamServiceUrl` (String) The IAM service URL to be used to generate tokens.  In the case of GLCS API clients
  (the default) then this should be set to the "issuer url" for the client.  In the case of GLP
  API clients use the appropriate "Token URL" from the API screen. Can be set by HPEGL_IAM_SERVICE_URL env-var
- `iamToken` (String) The IAM token to be used with the client(s).  Note that in normal operation
  an API client is used.  Passing-in a token means that tokens will not be generated or refreshed.
- `iamVersion` (String) The IAM version to be used.  Can be set by HPEGL_IAM_VERSION env-var. Valid values are:
  [glcs glp]The default is glcs.
- `metal` (Block Set, Max: 1) (see below for nested schema)
- `tenantId` (String) The tenant-id to be used for GLCS IAM, can be set by HPEGL_TENANT_ID env-var
- `userId` (String) The user id to be used, can be set by HPEGL_USER_ID env-var
- `userSecret` (String) The user secret to be used, can be set by HPEGL_USER_SECRET env-var
- `vmaas` (Block Set, Max: 1) (see below for nested schema)

<a id="nestedblock--caas"></a>
### Nested Schema for `caas`

Optional:

- `apiUrl` (String) The URL to use for the CaaS API, can also be set with the HPEGL_CAAS_API_URL env var

<a id="nestedblock--metal"></a>
### Nested Schema for `metal`

Optional:

- `glToken` (Boolean) Field indicating whether the token is GreenLake (GLCS or GLP) IAM issued token or Metal Service issued one,
  can also be set with the HPEGL_METAL_GL_TOKEN env-var
- `glpRole` (String) Field indicating the GLP role to be used, can also be set with the HPEGL_METAL_GLP_ROLE env-var
- `glpWorkspace` (String) Field indicating the GLP workspace to be used, can also be set with the HPEGL_METAL_GLP_WORKSPACE env-var
- `projectId` (String) The Metal project-id to use, can also be set with the HPEGL_METAL_PROJECT_ID env-var
- `restUrl` (String) The Metal portal rest-url to use, can also be set with the HPEGL_METAL_REST_URL env-var
- `spaceName` (String) The space-name to use with Metal, only required for project creation operations,
  can also be set with the HPEGL_METAL_SPACE_NAME env-var

<a id="nestedblock--vmaas"></a>
### Nested Schema for `vmaas`

Optional:

- `allowInsecure` (Boolean) Not to be used in production. To perform client connection ignoring TLS, it can also be set with the INSECURE env var
- `brokerUrl` (String) The URL to use for the VMaaS Broker API, can also be set with the HPEGL_VMAAS_BROKER_URL env var
- `location` (String) Location of GL VMaaS Service, can also be set with the HPEGL_VMAAS_LOCATION env var.
- `morpheusToken` (String) The Morpheus token, can also be set with the HPEGL_MORPHEUS_TOKEN env var
- `morpheusUrl` (String) The Morpheus URL, can also be set with the HPEGL_MORPHEUS_URL env var
- `spaceName` (String) It can also be set with the HPEGL_VMAAS_SPACE_NAME env var. When `HPEGL_IAM_VERSION` is `glcs` it refers to IAM Space name of the GL VMaaS Service i.e., Default. When `HPEGL_IAM_VERSION` is `glp` it refers to GLP Workspace ID.