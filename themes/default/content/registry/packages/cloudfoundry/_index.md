---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/cloudfoundry-community/cloudfoundry/0.54.0/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Cloudfoundry Provider
meta_desc: Provides an overview on how to configure the Pulumi Cloudfoundry provider.
layout: package
---

## Generate Provider

The Cloudfoundry provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider cloudfoundry-community/cloudfoundry
```
## Overview

The Cloud Foundry (cloudfoundry) provider is used to interact with a
Cloud Foundry target to perform administrative configuration of platform
resources, or user actions (such as pushing a cf app).

Use the navigation to the left to read about the available resources.
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    cloudfoundry:apiUrl:
        value: 'TODO: var.api_url'
    cloudfoundry:appLogsMax:
        value: 30
    cloudfoundry:password:
        value: 'TODO: var.password'
    cloudfoundry:skipSslValidation:
        value: true
    cloudfoundry:uaaClientId:
        value: admin
    cloudfoundry:uaaClientSecret:
        value: 'TODO: var.uaa_admin_client_secret'
    cloudfoundry:user:
        value: admin

```
```typescript
import * as pulumi from "@pulumi/pulumi";

const config = new pulumi.Config();
const apiUrl = config.requireObject("apiUrl");
const password = config.requireObject("password");
const uaaAdminClientSecret = config.requireObject("uaaAdminClientSecret");
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    cloudfoundry:apiUrl:
        value: 'TODO: var.api_url'
    cloudfoundry:appLogsMax:
        value: 30
    cloudfoundry:password:
        value: 'TODO: var.password'
    cloudfoundry:skipSslValidation:
        value: true
    cloudfoundry:uaaClientId:
        value: admin
    cloudfoundry:uaaClientSecret:
        value: 'TODO: var.uaa_admin_client_secret'
    cloudfoundry:user:
        value: admin

```
```python
import pulumi

config = pulumi.Config()
api_url = config.require_object("apiUrl")
password = config.require_object("password")
uaa_admin_client_secret = config.require_object("uaaAdminClientSecret")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    cloudfoundry:apiUrl:
        value: 'TODO: var.api_url'
    cloudfoundry:appLogsMax:
        value: 30
    cloudfoundry:password:
        value: 'TODO: var.password'
    cloudfoundry:skipSslValidation:
        value: true
    cloudfoundry:uaaClientId:
        value: admin
    cloudfoundry:uaaClientSecret:
        value: 'TODO: var.uaa_admin_client_secret'
    cloudfoundry:user:
        value: admin

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;

return await Deployment.RunAsync(() =>
{
    var config = new Config();
    var apiUrl = config.RequireObject<dynamic>("apiUrl");
    var password = config.RequireObject<dynamic>("password");
    var uaaAdminClientSecret = config.RequireObject<dynamic>("uaaAdminClientSecret");
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
        value: 'TODO: var.api_url'
    cloudfoundry:appLogsMax:
        value: 30
    cloudfoundry:password:
        value: 'TODO: var.password'
    cloudfoundry:skipSslValidation:
        value: true
    cloudfoundry:uaaClientId:
        value: admin
    cloudfoundry:uaaClientSecret:
        value: 'TODO: var.uaa_admin_client_secret'
    cloudfoundry:user:
        value: admin

```
```go
package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")
		apiUrl := cfg.RequireObject("apiUrl")
		password := cfg.RequireObject("password")
		uaaAdminClientSecret := cfg.RequireObject("uaaAdminClientSecret")
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
        value: 'TODO: var.api_url'
    cloudfoundry:appLogsMax:
        value: 30
    cloudfoundry:password:
        value: 'TODO: var.password'
    cloudfoundry:skipSslValidation:
        value: true
    cloudfoundry:uaaClientId:
        value: admin
    cloudfoundry:uaaClientSecret:
        value: 'TODO: var.uaa_admin_client_secret'
    cloudfoundry:user:
        value: admin

```
```yaml
configuration:
  # Set the variable values in *.tfvars file
  # or using -var="api_url=..." CLI option
  apiUrl:
    type: dynamic
  password:
    type: dynamic
  uaaAdminClientSecret:
    type: dynamic
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    cloudfoundry:apiUrl:
        value: 'TODO: var.api_url'
    cloudfoundry:appLogsMax:
        value: 30
    cloudfoundry:password:
        value: 'TODO: var.password'
    cloudfoundry:skipSslValidation:
        value: true
    cloudfoundry:uaaClientId:
        value: admin
    cloudfoundry:uaaClientSecret:
        value: 'TODO: var.uaa_admin_client_secret'
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
        final var config = ctx.config();
        final var apiUrl = config.get("apiUrl");
        final var password = config.get("password");
        final var uaaAdminClientSecret = config.get("uaaAdminClientSecret");
    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

The following configuration inputs are supported:

* `apiUrl` - (Required) API endpoint (e.g. [https://api.local.pcfdev.io](https://api.local.pcfdev.io)). This can also be specified
  with the `CF_API_URL` shell environment variable.

* `origin` - (Optional) Indicates the identity provider to be used for login. This can also be specified
  with the `CF_ORIGIN` shell environment variable.

* `user` - (Optional) Cloud Foundry user. Defaults to "admin". This can also be specified
  with the `CF_USER` shell environment variable. Unless mentionned explicitly in a resource, CF admin permissions are not required.

* `password` - (Optional) Cloud Foundry user's password. This can also be specified
  with the `CF_PASSWORD` shell environment variable.

* `ssoPasscode` - (Optional) A passcode provided by UAA single sign on. The equivalent of `cf login --sso-passcode`. This can also be specified
  with the `CF_SSO_PASSCODE` shell environment variable.

* `cfClientId` - (Optional) The cf client ID to make request with a client instead of user. This can also be specified
  with the `CF_CLIENT_ID` shell environment variable.

* `cfClientSecret` - (Optional) The cf client secret to make request with a client instead of user. This can also be specified
  with the `CF_CLIENT_SECRET` shell environment variable.

* `uaaClientId` - (Optional) The UAA admin client ID. Defaults to "admin". This can also be specified
  with the `CF_UAA_CLIENT_ID` shell environment variable.

* `uaaClientSecret` - (Optional) This secret of the UAA admin client. This can also be specified
  with the `CF_UAA_CLIENT_SECRET` shell environment variable.

* `skipSslValidation` - (Optional) Skip verification of the API endpoint - Not recommended!. Defaults to "false". This can also be specified
  with the `CF_SKIP_SSL_VALIDATION` shell environment variable.

* `defaultQuotaName` - (Optional, Default: `default`) Change the name of your default quota . This can also be specified
  with the `CF_DEFAULT_QUOTA_NAME` shell environment variable.

* `appLogsMax` - (Optional) Number of logs message which can be see when app creation is errored (-1 means all messages stored). Defaults to "30". This can also be specified
  with the `CF_APP_LOGS_MAX` shell environment variable.

* `purgeWhenDelete` - (Optional) Set to true to purge when deleting a resource (e.g.: service instance, service broker) . This can also be specified
  with the `CF_PURGE_WHEN_DELETE` shell environment variable.

* `storeTokensPath` - (Optional) Path to a file to store tokens used for login. (this is useful for sso, this avoid
  requiring each time sso passcode) . This can also be specified with the `CF_STORE_TOKENS_PATH` shell environment variable.

* `forceBrokerNotFailWhenCatalogNotAccessible` (Optional) Set to true to enforce `failWhenCatalogNotAccessible` to `true` to all broker for avoiding being
  stuck if broker has been deleted for example. This can also be specified with the `CF_FORCE_BROKER_NOT_FAIL_CATALOG` shell environment variable.