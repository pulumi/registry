---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/onelogin/onelogin/0.10.0/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Onelogin Provider
meta_desc: Provides an overview on how to configure the Pulumi Onelogin provider.
layout: package
---

## Generate Provider

The Onelogin provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider onelogin/onelogin
```
## Overview

The OneLogin provider is used to interact with OneLogin resources.

The provider allows you to manage your OneLogin organization's resources easily.
It needs to be configured with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.
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
import * as onelogin from "@pulumi/onelogin";

// Add an App to your account
const mySamlApp = new onelogin.index.SamlApp("my_saml_app", {});
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
import pulumi_onelogin as onelogin

# Add an App to your account
my_saml_app = onelogin.index.SamlApp("my_saml_app")
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
using Onelogin = Pulumi.Onelogin;

return await Deployment.RunAsync(() =>
{
    // Add an App to your account
    var mySamlApp = new Onelogin.Index.SamlApp("my_saml_app");

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
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/onelogin/onelogin"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Add an App to your account
		_, err := onelogin.NewSamlApp(ctx, "my_saml_app", nil)
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

```
```yaml
resources:
  # Add an App to your account
  mySamlApp:
    type: onelogin:SamlApp
    name: my_saml_app
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
import com.pulumi.onelogin.SamlApp;
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
        // Add an App to your account
        var mySamlApp = new SamlApp("mySamlApp");

    }
}
```
{{% /choosable %}}
{{< /chooser >}}

You're also welcome to leave the provider field blank and export your
credentials to your environment
## Configuration Reference

The following configuration inputs are supported in the provider configuration:

None: This provider reads API credentials from your environment. You need to export
your OneLogin credentials like so:

```
export ONELOGIN_CLIENT_ID=<your client id>
export ONELOGIN_CLIENT_SECRET=<your client secret>
export ONELOGIN_API_URL=<the complete api url, e.g., https://company.onelogin.com>
```

For backward compatibility, you can also use the subdomain approach, but this is deprecated:

```
export ONELOGIN_CLIENT_ID=<your client id>
export ONELOGIN_CLIENT_SECRET=<your client secret>
export ONELOGIN_SUBDOMAIN=<your onelogin subdomain, e.g., 'company' for company.onelogin.com>
```

Using the `ONELOGIN_API_URL` is now the recommended approach.