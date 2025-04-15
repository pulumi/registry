---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/flexibleenginecloud/flexibleengine/1.46.0/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Flexibleengine Provider
meta_desc: Provides an overview on how to configure the Pulumi Flexibleengine provider.
layout: package
---

## Generate Provider

The Flexibleengine provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider flexibleenginecloud/flexibleengine
```
## Overview

The FlexibleEngine provider is used to interact with the
many resources supported by FlexibleEngine. The provider needs to be configured
with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    flexibleengine:domainName:
        value: admin
    flexibleengine:password:
        value: pwd
    flexibleengine:region:
        value: eu-west-0
    flexibleengine:userName:
        value: admin

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as flexibleengine from "@pulumi/flexibleengine";

// Create a web server
const test_server = new flexibleengine.ComputeInstanceV2("test-server", {});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    flexibleengine:domainName:
        value: admin
    flexibleengine:password:
        value: pwd
    flexibleengine:region:
        value: eu-west-0
    flexibleengine:userName:
        value: admin

```
```python
import pulumi
import pulumi_flexibleengine as flexibleengine

# Create a web server
test_server = flexibleengine.ComputeInstanceV2("test-server")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    flexibleengine:domainName:
        value: admin
    flexibleengine:password:
        value: pwd
    flexibleengine:region:
        value: eu-west-0
    flexibleengine:userName:
        value: admin

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Flexibleengine = Pulumi.Flexibleengine;

return await Deployment.RunAsync(() =>
{
    // Create a web server
    var test_server = new Flexibleengine.ComputeInstanceV2("test-server");

});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    flexibleengine:domainName:
        value: admin
    flexibleengine:password:
        value: pwd
    flexibleengine:region:
        value: eu-west-0
    flexibleengine:userName:
        value: admin

```
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/flexibleengine/flexibleengine"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a web server
		_, err := flexibleengine.NewComputeInstanceV2(ctx, "test-server", nil)
		if err != nil {
			return err
		}
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
    flexibleengine:domainName:
        value: admin
    flexibleengine:password:
        value: pwd
    flexibleengine:region:
        value: eu-west-0
    flexibleengine:userName:
        value: admin

```
```yaml
resources:
  # Create a web server
  test-server:
    type: flexibleengine:ComputeInstanceV2
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    flexibleengine:domainName:
        value: admin
    flexibleengine:password:
        value: pwd
    flexibleengine:region:
        value: eu-west-0
    flexibleengine:userName:
        value: admin

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.flexibleengine.ComputeInstanceV2;
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
        // Create a web server
        var test_server = new ComputeInstanceV2("test-server");

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Authentication
### User name + Password

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    flexibleengine:domainName:
        value: 'TODO: var.domain_name'
    flexibleengine:password:
        value: 'TODO: var.password'
    flexibleengine:region:
        value: eu-west-0
    flexibleengine:userName:
        value: 'TODO: var.user_name'

```
### AKSK

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    flexibleengine:accessKey:
        value: 'TODO: var.access_key'
    flexibleengine:domainName:
        value: 'TODO: var.domain_name'
    flexibleengine:region:
        value: eu-west-0
    flexibleengine:secretKey:
        value: 'TODO: var.secret_key'

```
### Token

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    flexibleengine:domainName:
        value: 'TODO: var.domain_name'
    flexibleengine:region:
        value: eu-west-0
    flexibleengine:tenantName:
        value: 'TODO: var.tenant_name'
    flexibleengine:token:
        value: 'TODO: var.token'

```

> If token, aksk and password are set simultaneously, then it will authenticate in the order of Token, Password and AKSK.
### Federated

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    flexibleengine:accessKey:
        value: 'TODO: var.access_key'
    flexibleengine:domainName:
        value: 'TODO: var.domain_name'
    flexibleengine:region:
        value: eu-west-0
    flexibleengine:secretKey:
        value: 'TODO: var.secret_key'
    flexibleengine:securityToken:
        value: 'TODO: var.security_token'
    flexibleengine:tenantName:
        value: 'TODO: var.tenant_name'
    flexibleengine:token:
        value: 'TODO: var.token'

```
## Configuration Reference

The following configuration inputs are supported:

* `region` - (Required) The region of the FlexibleEngine cloud to use. It must be provided,
  but it can also be sourced from the `OS_REGION_NAME` environment variables.

* `accessKey` - (Optional) The access key of the FlexibleEngine cloud to use.
  If omitted, the `OS_ACCESS_KEY` environment variable is used.

* `secretKey` - (Optional) The secret key of the FlexibleEngine cloud to use.
  If omitted, the `OS_SECRET_KEY` environment variable is used.

* `userName` - (Optional) The User name to login with. If omitted, the
  `OS_USER_NAME` environment variable is used.

* `userId` - (Optional) The User ID to login with. If omitted, the
  `OS_USER_ID` environment variable is used.

* `password` - (Optional) The Password to login with. If omitted, the
  `OS_PASSWORD` environment variable is used.

* `domainId` - (Optional) The ID of the Domain to scope to.
  If omitted, the following environment variables are checked (in this order):
  `OS_USER_DOMAIN_ID`, `OS_PROJECT_DOMAIN_ID`, `OS_DOMAIN_ID`.

* `domainName` - (Optional) The Name of the Domain to scope to.
  If omitted, the following environment variables are checked (in this order):
  `OS_USER_DOMAIN_NAME`, `OS_PROJECT_DOMAIN_NAME`, `OS_DOMAIN_NAME`, `DEFAULT_DOMAIN`.

* `tenantId` - (Optional) The ID of the Project to login with.
  If omitted, the `OS_TENANT_ID` or `OS_PROJECT_ID` environment variables are used.

* `tenantName` - (Optional) The Name of the Project to login with.
  If omitted, the `OS_TENANT_NAME`, `OS_PROJECT_NAME` environment variable or `region` is used.

* `token` - (Optional) A token is an expiring, temporary means of access issued via the
  IAM service. By specifying a token, you do not have to specify a username/password
  combination, since the token was already created by a username/password out of
  band of Pulumi. If omitted, the `OS_AUTH_TOKEN` environment variable is used.

* `securityToken` - (Optional) The security token to authenticate with a temporary security credential.
  If omitted, the `OS_SECURITY_TOKEN` environment variable is used.

* `authUrl` - (Optional) The Identity authentication URL.
  If omitted, the `OS_AUTH_URL` environment variable is used.
  The default value is `https://iam.{{region}}.prod-cloud-ocb.orange-business.com/v3`.

* `maxRetries` - (Optional) This is the maximum number of times an API
  call is retried, in the case where requests are being throttled or
  experiencing transient failures. The delay between the subsequent API
  calls increases exponentially. The default value is `5`.
  If omitted, the `OS_MAX_RETRIES` environment variable is used.

* `insecure` - (Optional) Trust self-signed SSL certificates. If omitted, the
  `OS_INSECURE` environment variable is used.

* `cacertFile` - (Optional) Specify a custom CA certificate when communicating
  over SSL. You can specify either a path to the file or the contents of the
  certificate. If omitted, the `OS_CACERT` environment variable is used.

* `cert` - (Optional) Specify client certificate file for SSL client
  authentication. You can specify either a path to the file or the contents of
  the certificate. If omitted the `OS_CERT` environment variable is used.

* `key` - (Optional) Specify client private key file for SSL client
  authentication. You can specify either a path to the file or the contents of
  the key. If omitted the `OS_KEY` environment variable is used.