---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/dome9/dome9/1.40.4/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Dome9 Provider
meta_desc: Provides an overview on how to configure the Pulumi Dome9 provider.
layout: package
---

## Generate Provider

The Dome9 provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider dome9/dome9
```
## Overview

The Check Point CloudGuard Dome9 provider is used to interact with [CloudGuard Dome9](https://www.checkpoint.com/dome9/) ([https://www.checkpoint.com/dome9/](https://www.checkpoint.com/dome9/)) security posture platform
to onboard cloud environments and configure compliance policies.
To use this  provider, you must create CloudGuard API credentials.<br/>
Use the navigation menu on the left to read about the available resources.
## Authentication

This provider requires a CloudGuard API Key and Key secret in order to manage the resources.
- For accounts onboarded through the Infinity Portal, you can generate these credentials at [https://portal.checkpoint.com/dashboard/cloudguard#/settings/credentials](https://portal.checkpoint.com/dashboard/cloudguard#/settings/credentials) . For more details, see the  [CloudGuard documentation](https://sc1.checkpoint.com/documents/Infinity_Portal/WebAdminGuides/EN/CloudGuard-PM-Admin-Guide/Documentation/Settings/Credentials.htm?cshid=API_V2) for the Infinity Portal.

- For accounts onboarded through the CloudGuard Portal, use [https://secure.dome9.com/v2/settings/credentials](https://secure.dome9.com/v2/settings/credentials). For more details, see  [CloudGuard documentation](https://sc1.checkpoint.com/documents/CloudGuard_Dome9/Default.htm#cshid=API_V2)

To manage the full selection of resources, provide the credentials from an account with administrative access permissions.

You can use the Key and Secret in the following ways:

- On the CLI, omit the provider configuration from your tf file, and the CLI will prompt for proper credentials.
  [CLI config file](https://www.terraform.io/docs/commands/cli-config.html#credentials).
- Set the `DOME9_ACCESS_ID` and `DOME9_SECRET_KEY` environment variables.
- [Optional] The provider works by default with US region. Set 'base_url' with one of the following
  URLs for working with other supported regions.
  Support regions URLs list:
  - N.Virginia [DEFAULT]: 'https://api.dome9.com/v2/'
  - Ireland : 'https://api.eu1.dome9.com/v2/'
  - Singapore : 'https://api.ap1.dome9.com/v2/'
  - Sydney : 'https://api.ap2.dome9.com/v2/'
  - Mumbai : 'https://api.ap3.dome9.com/v2/'
  - Canada : 'https://api.cace1.dome9.com/v2/'
- Fill the provider configuration with the appropriate arguments:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    dome9:dome9AccessId:
        value: 'TODO: "${var.access_id}"'
    dome9:dome9SecretKey:
        value: 'TODO: "${var.secret_key}"'

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as dome9 from "@pulumi/dome9";

// Create an organization
const account = new dome9.CloudaccountAws("account", {});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    dome9:dome9AccessId:
        value: 'TODO: "${var.access_id}"'
    dome9:dome9SecretKey:
        value: 'TODO: "${var.secret_key}"'

```
```python
import pulumi
import pulumi_dome9 as dome9

# Create an organization
account = dome9.CloudaccountAws("account")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    dome9:dome9AccessId:
        value: 'TODO: "${var.access_id}"'
    dome9:dome9SecretKey:
        value: 'TODO: "${var.secret_key}"'

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Dome9 = Pulumi.Dome9;

return await Deployment.RunAsync(() =>
{
    // Create an organization
    var account = new Dome9.CloudaccountAws("account");

});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    dome9:dome9AccessId:
        value: 'TODO: "${var.access_id}"'
    dome9:dome9SecretKey:
        value: 'TODO: "${var.secret_key}"'

```
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/dome9/dome9"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create an organization
		_, err := dome9.NewCloudaccountAws(ctx, "account", nil)
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
    dome9:dome9AccessId:
        value: 'TODO: "${var.access_id}"'
    dome9:dome9SecretKey:
        value: 'TODO: "${var.secret_key}"'

```
```yaml
resources:
  # Create an organization
  account:
    type: dome9:CloudaccountAws
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    dome9:dome9AccessId:
        value: 'TODO: "${var.access_id}"'
    dome9:dome9SecretKey:
        value: 'TODO: "${var.secret_key}"'

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.dome9.CloudaccountAws;
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
        // Create an organization
        var account = new CloudaccountAws("account");

    }
}
```
{{% /choosable %}}
{{< /chooser >}}

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    dome9:baseUrl:
        value: 'TODO: "${var.base_url}"'
    dome9:dome9AccessId:
        value: 'TODO: "${var.access_id}"'
    dome9:dome9SecretKey:
        value: 'TODO: "${var.secret_key}"'

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as dome9 from "@pulumi/dome9";

// Create an organization
const account = new dome9.CloudaccountAws("account", {});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    dome9:baseUrl:
        value: 'TODO: "${var.base_url}"'
    dome9:dome9AccessId:
        value: 'TODO: "${var.access_id}"'
    dome9:dome9SecretKey:
        value: 'TODO: "${var.secret_key}"'

```
```python
import pulumi
import pulumi_dome9 as dome9

# Create an organization
account = dome9.CloudaccountAws("account")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    dome9:baseUrl:
        value: 'TODO: "${var.base_url}"'
    dome9:dome9AccessId:
        value: 'TODO: "${var.access_id}"'
    dome9:dome9SecretKey:
        value: 'TODO: "${var.secret_key}"'

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Dome9 = Pulumi.Dome9;

return await Deployment.RunAsync(() =>
{
    // Create an organization
    var account = new Dome9.CloudaccountAws("account");

});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    dome9:baseUrl:
        value: 'TODO: "${var.base_url}"'
    dome9:dome9AccessId:
        value: 'TODO: "${var.access_id}"'
    dome9:dome9SecretKey:
        value: 'TODO: "${var.secret_key}"'

```
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/dome9/dome9"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create an organization
		_, err := dome9.NewCloudaccountAws(ctx, "account", nil)
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
    dome9:baseUrl:
        value: 'TODO: "${var.base_url}"'
    dome9:dome9AccessId:
        value: 'TODO: "${var.access_id}"'
    dome9:dome9SecretKey:
        value: 'TODO: "${var.secret_key}"'

```
```yaml
resources:
  # Create an organization
  account:
    type: dome9:CloudaccountAws
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    dome9:baseUrl:
        value: 'TODO: "${var.base_url}"'
    dome9:dome9AccessId:
        value: 'TODO: "${var.access_id}"'
    dome9:dome9SecretKey:
        value: 'TODO: "${var.secret_key}"'

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.dome9.CloudaccountAws;
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
        // Create an organization
        var account = new CloudaccountAws("account");

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
### Configuration Reference

* `dome9AccessId` - (Required) the Dome9 API Key
* `dome9SecretKey` - (Required) the Dome9  key secret