---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/cyralinc/cyral/4.16.3/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Cyral Provider
meta_desc: Provides an overview on how to configure the Pulumi Cyral provider.
layout: package
---

## Generate Provider

The Cyral provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider cyralinc/cyral
```
## Overview

> If you are upgrading from versions `v2` or `v3` of this provider, please refer to our
4.0 migration guide.
The Cyral Pulumi Provider contains resources that can be used to interact
with the Cyral API through Pulumi code. It allows customers to maintain a
history of changes in Cyral environment by transforming configurations into
code and use Pulumi to control state changes.

Our provider uses the same naming conventions and organization as stated in
Pulumi guidelines for writing custom providers.
## Version history, upgrades and compatibility

Please refer to our Change Log to learn about our version
history with its features, bug fixes and Control Plane compatibility.

> If you are upgrading from versions `v2` or `v3` of this provider, please refer to our
4.0 migration guide.
## How to use this provider

To install this provider, copy and paste this code into your Pulumi
configuration. Then, run `pulumi up`.

- Pulumi 0.13+

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    cyral:clientId:
        value: ""
    cyral:clientSecret:
        value: ""
    cyral:controlPlane:
        value: '[TENANT].app.cyral.com'

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
    cyral:clientId:
        value: ""
    cyral:clientSecret:
        value: ""
    cyral:controlPlane:
        value: '[TENANT].app.cyral.com'

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
    cyral:clientId:
        value: ""
    cyral:clientSecret:
        value: ""
    cyral:controlPlane:
        value: '[TENANT].app.cyral.com'

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
    cyral:clientId:
        value: ""
    cyral:clientSecret:
        value: ""
    cyral:controlPlane:
        value: '[TENANT].app.cyral.com'

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
    cyral:clientId:
        value: ""
    cyral:clientSecret:
        value: ""
    cyral:controlPlane:
        value: '[TENANT].app.cyral.com'

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
    cyral:clientId:
        value: ""
    cyral:clientSecret:
        value: ""
    cyral:controlPlane:
        value: '[TENANT].app.cyral.com'

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
### Provider Credentials - UI
#### New Credentials

An `API Client` must be created in order to use the provider.
It can be created through the control plane UI, accessing the `API Access Keys`
option in the left menu and clicking on the `+` button. Choose a name for the
new access key and select the following roles so you can use all the provider
functions:

<img src="https://raw.githubusercontent.com/cyralinc/pulumi-provider-cyral/main/images/create_api_client.png">

Confirm the creation by clicking on the `CREATE` button. This will generate a
`Client ID` and a `Client Secret` that should be used in the
provider configuration block.
#### Rotate Credentials

To rotate secrets for existing `API Clients`, access the `API Access Keys` menu
in the UI, select the desired account, click `edit` and finally click
`ROTATE CLIENT SECRET` as the image below suggests:

<img src="https://raw.githubusercontent.com/cyralinc/pulumi-provider-cyral/main/images/rotate_client_secret.png">

That will generate a new `Client Secret` that you can copy and use to replace
the old one.
## Configuration Reference
### Required

- `controlPlane` (String) Control plane host and API port (ex: `tenant.app.cyral.com`)

- `clientId` (String, Sensitive) Client id used to authenticate against the control plane. Can be ommited and declared using the environment variable `CYRAL_TF_CLIENT_ID`.
- `clientSecret` (String, Sensitive) Client secret used to authenticate against the control plane. Can be ommited and declared using the environment variable `CYRAL_TF_CLIENT_SECRET`.
- `tlsSkipVerify` (Boolean) Specifies if the client will verify the TLS server certificate used by the control plane. If set to `true`, the client will not verify the server certificate, hence, it will allow insecure connections to be established. This should be set only for testing and is not recommended to be used in production environments. Can be set through the `CYRAL_TF_TLS_SKIP_VERIFY` environment variable. Defaults to `false`.