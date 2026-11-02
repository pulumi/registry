---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/cloudfoundry/cloudfoundry/1.18.0/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Cloudfoundry Provider
meta_desc: Provides an overview on how to configure the Pulumi Cloudfoundry provider.
layout: package
---

## Generate Provider

The Cloudfoundry provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider cloudfoundry/cloudfoundry
```
## Overview

The Cloud Foundry Pulumi plugin is an integration that allows users to leverage Pulumi, an infrastructure as code tool, to define and provision infrastructure resources within the Cloud Foundry platform.
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    cloudfoundry:apiUrl:
        value: https://api.cf.example.com
    cloudfoundry:password:
        value: admin
    cloudfoundry:user:
        value: admin

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
config:
    cloudfoundry:apiUrl:
        value: https://api.cf.example.com
    cloudfoundry:password:
        value: admin
    cloudfoundry:user:
        value: admin

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
config:
    cloudfoundry:apiUrl:
        value: https://api.cf.example.com
    cloudfoundry:password:
        value: admin
    cloudfoundry:user:
        value: admin

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
config:
    cloudfoundry:apiUrl:
        value: https://api.cf.example.com
    cloudfoundry:password:
        value: admin
    cloudfoundry:user:
        value: admin

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
config:
    cloudfoundry:apiUrl:
        value: https://api.cf.example.com
    cloudfoundry:password:
        value: admin
    cloudfoundry:user:
        value: admin

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
config:
    cloudfoundry:apiUrl:
        value: https://api.cf.example.com
    cloudfoundry:password:
        value: admin
    cloudfoundry:user:
        value: admin

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

- `accessToken` (String, Sensitive) OAuth token to authenticate with Cloud Foundry
- `apiUrl` (String) Specific URL representing the entry point for communication between the client and a Cloud Foundry instance.
- `assertionToken` (String, Sensitive) OAuth JWT assertion token. Used for OAuth 2.0 JWT Bearer Assertion Grant flow to authenticate with Cloud Foundry. Typically used with a custom origin.
- `cfClientId` (String) Unique identifier for a client application used in authentication and authorization processes
- `cfClientSecret` (String, Sensitive) A confidential string used by a client application for secure authentication and authorization, requires cfClientId to authenticate
- `origin` (String) Indicates the identity provider to be used for login
- `password` (String, Sensitive) A confidential alphanumeric code associated with a user account on the Cloud Foundry platform, requires user to authenticate.
- `refreshToken` (String, Sensitive) Token to refresh the access token, requires accessToken
- `skipSslValidation` (Boolean) Allows the client to disregard SSL certificate validation when connecting to the Cloud Foundry API
- `user` (String) A unique identifier associated with an individual or entity for authentication & authorization purposes.

**Note**

All parameter values for the provider can be injected by setting environment variables `CF_API_URL`, `CF_USER`, `CF_PASSWORD`, `CF_ORIGIN`, `CF_CLIENT_ID`, `CF_CLIENT_SECRET`, `CF_ACCESS_TOKEN`, `CF_REFRESH_TOKEN`.
Alternatively, one can even log in to their CF landscape via CF-CLI and the provider will pick the credentials from the config.json present in CF Home in case no attributes are given in the provider configuration or if no environment variables are set.
## Custom User-Agent Information

By default, the underlying Cloud Foundry client used by the Pulumi Cloud Foundry Provider creates requests with User-Agent headers that include information about Pulumi and Cloud Foundry Pulumi provider versions. To add more details to the User-Agent headers, the `CF_APPEND_USER_AGENT` environment variable can be set, and its value will be directly added to HTTP requests. E.g.,

```bash
% export CF_APPEND_USER_AGENT="Optional_Extra_Information"
```
## Enabling Application Log Tracing

To assist with debugging application failures, you can enable log tracing by setting the following environment variable:

```bash
% export ENABLE_CF_APP_LOG_TRACE="true"
```
When enabled, recent error logs from the application will be printed to the console if a deployment of application fails.
## Get Started

If you're not familiar with Pulumi yet, see the Fundamentals section with a lot of helpful tutorials.
## Migration Guide

The newer v3 APIs have brought several changes to  many of the resources created. Some are even deprecated and replaced by
different resources.

Please refer to the migration guide to help transitioning from the existing community CF provider to the current CF provider. The migration guide
highlights the differences and maps to the equivalent resources/functions as well.
## Best Practices

For the best experience using the Cloud Foundry provider, we recommend applying the common best practices for Pulumi adoption as described in the Pulumi documentation. For example, see Phases of Pulumi Adoption.
## Drift Detection

You find further information on drift detection in the Pulumi provider repository.
## Import

The documentation of each resource shows if the resource can be imported, and which keys are needed. To provide a quick overview we listed the resources in the Pulumi provider repository. To import a resource whose ID contains commas or other special characters, you should consider to use the import block with a resource identity if the resource supports identity-based import.