---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/vk-cs/vkcs/0.14.0/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Vkcs Provider
meta_desc: Provides an overview on how to configure the Pulumi Vkcs provider.
layout: package
---

## Generate Provider

The Vkcs provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider vk-cs/vkcs
```
## Overview

The VKCS provider is used to interact with [VKCS services](https://mcs.mail.ru/). The provider needs to be configured with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.
## Example Usage



{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as vkcs from "@pulumi/vkcs";

// Create new compute instance
const myinstance = new vkcs.ComputeInstance("myinstance", {});
```
{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import pulumi_vkcs as vkcs

# Create new compute instance
myinstance = vkcs.ComputeInstance("myinstance")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Vkcs = Pulumi.Vkcs;

return await Deployment.RunAsync(() =>
{
    // Create new compute instance
    var myinstance = new Vkcs.ComputeInstance("myinstance");

});

```
{{% /choosable %}}
{{% choosable language go %}}
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/vkcs/vkcs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create new compute instance
		_, err := vkcs.NewComputeInstance(ctx, "myinstance", nil)
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
resources:
  # Create new compute instance
  myinstance:
    type: vkcs:ComputeInstance
```
{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.vkcs.ComputeInstance;
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
        // Create new compute instance
        var myinstance = new ComputeInstance("myinstance");

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Authentication

The VKCS provider supports username/password authentication. Preconfigured provider file with `username` and `projectId` can be downloaded from [https://mcs.mail.ru/app/project](https://mcs.mail.ru/app/project) portal. Go to `Pulumi` tab > click on the "Download VKCS provider file".

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    vkcs:password:
        value: PASSWORD
    vkcs:projectId:
        value: PROJECT_ID
    vkcs:username:
        value: USERNAME

```
## Configuration Reference
- `accessToken` optional sensitive *string* &rarr;  A temporary token to use for authentication. You alternatively can use `OS_AUTH_TOKEN` environment variable. If both are specified, this attribute takes precedence. <br>**Note:** The token will not be renewed and will eventually expire, usually after 1 hour. If access is needed for longer than a token's lifetime, use credentials-based authentication.

- `authUrl` optional *string* &rarr;  The Identity authentication URL.

- `cloudContainersApiVersion` optional *string* &rarr;  Cloud Containers API version to use. <br>**Note:** Only for custom VKCS deployments.

- `endpointOverrides` optional &rarr;  Custom endpoints for corresponding APIs. If not specified, endpoints provided by the catalog will be used.
  - `backup` optional *string* &rarr;  Backup API custom endpoint.

  - `blockStorage` optional *string* &rarr;  Block Storage API custom endpoint.

  - `cdn` optional *string* &rarr;  CDN API custom endpoint.

  - `compute` optional *string* &rarr;  Compute API custom endpoint.

  - `containerInfra` optional *string* &rarr;  Cloud Containers API custom endpoint.

  - `containerInfraAddons` optional *string* &rarr;  Cloud Containers Addons API custom endpoint.

  - `dataPlatform` optional *string* &rarr;  Data Platform API custom endpoint.

  - `database` optional *string* &rarr;  Database API custom endpoint.

  - `iamServiceUsers` optional *string* &rarr;  IAM Service Users API custom endpoint.

  - `ics` optional *string* &rarr;  ICS API custom endpoint.

  - `image` optional *string* &rarr;  Image API custom endpoint.

  - `keyManager` optional *string* &rarr;  Key Manager API custom endpoint.

  - `loadBalancer` optional *string* &rarr;  Load Balancer API custom endpoint.

  - `mlPlatform` optional *string* &rarr;  ML Platform API custom endpoint.

  - `networking` optional *string* &rarr;  Networking API custom endpoint.

  - `publicDns` optional *string* &rarr;  Public DNS API custom endpoint.

  - `sharedFilesystem` optional *string* &rarr;  Shared Filesystem API custom endpoint.

  - `templater` optional *string* &rarr;  Templater API custom endpoint.

- `password` optional sensitive *string* &rarr;  Password to login with.

- `projectId` optional *string* &rarr;  The ID of Project to login with.

- `region` optional *string* &rarr;  A region to use.

- `skipClientAuth` optional *boolean* &rarr;  Skip authentication on client initialization. Only applicablie if `accessToken` is provided. <br>**Note:** If set to true, the endpoint catalog will not be used for discovery and all required endpoints must be provided via `endpointOverrides`.

- `userDomainId` optional *string* &rarr;  The id of the domain where the user resides.

- `userDomainName` optional *string* &rarr;  The name of the domain where the user resides.

- `username` optional *string* &rarr;  User name to login with.
## Working with VKCS Cloud Storage

VKCS provider does not support working with cloud storage.
To do this, we recommend an AWS provider, you can learn how to use it by following this documentation.