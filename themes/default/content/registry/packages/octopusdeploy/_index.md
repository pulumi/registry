---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/octopusdeploylabs/octopusdeploy/0.43.0/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Octopusdeploy Provider
meta_desc: Provides an overview on how to configure the Pulumi Octopusdeploy provider.
layout: package
---

## Generate Provider

The Octopusdeploy provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider octopusdeploylabs/octopusdeploy
```
## Warning

This provider is under active development. Its functionalty can and will change; it is a `v0.*` product until its robustness can be assured. Please be aware that types like resources can and will be modified over time. It is strongly recommended to `validate` and `plan` configuration prior to committing changes via `apply`.
## Overview

This provider is used to configure resources in Octopus Deploy. The provider must be configured with the proper credentials before it can be used.
## Configuration
### Authentication Methods
The provider supports authenticating to an Octopus Server instance via either:
* API Key
* OIDC Access Token

These are mutually exclusive options - use either, not both. For backward compatibility, API Key will always be preferred over OIDC, when an API Key is present.
### Default Space

Octopus Deploy supports the concept of a Default Space. This is the first space that is automatically created on server setup. If you do not specify a Space when configuring the Octopus Deploy Pulumi provider it will use the Default Space.

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    octopusdeploy:address:
        value: https://octopus.example.com
    octopusdeploy:apiKey:
        value: API-XXXXXXXXXXXXX

```
### Scoped to a specific Space

You can configure the Pulumi Provider for Octopus Deploy to target a
particular space. If this configuration is specified, resources managed by the
provider will be scoped to this space. To scope the provider to a space, simply
provide the *ID* of the space.

Scoping the provider by the ID of a space is done as follows:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    octopusdeploy:address:
        value: https://octopus.example.com
    octopusdeploy:apiKey:
        value: API-XXXXXXXXXXXXX
    octopusdeploy:spaceId:
        value: Spaces-321

```
### Multiple Spaces

To manage resources in multiple spaces you can specify the spaceId on the resource directly:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    octopusdeploy:address:
        value: https://octopus.example.com
    octopusdeploy:apiKey:
        value: API-XXXXXXXXXXXXX

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as octopusdeploy from "@pulumi/octopusdeploy";

const dev = octopusdeploy.getSpace({
    name: "Product Development",
});
//This resource will use the default space
const development_environment = new octopusdeploy.Environment("development-environment", {name: "TestEnv1"});
//This resource will be scoped to the space named "Product Development".
const env3 = new octopusdeploy.Environment("Env3", {
    spaceId: dev.then(dev => dev.id),
    name: "TestEnv3",
});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    octopusdeploy:address:
        value: https://octopus.example.com
    octopusdeploy:apiKey:
        value: API-XXXXXXXXXXXXX

```
```python
import pulumi
import pulumi_octopusdeploy as octopusdeploy

dev = octopusdeploy.get_space(name="Product Development")
#This resource will use the default space
development_environment = octopusdeploy.Environment("development-environment", name="TestEnv1")
#This resource will be scoped to the space named "Product Development".
env3 = octopusdeploy.Environment("Env3",
    space_id=dev.id,
    name="TestEnv3")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    octopusdeploy:address:
        value: https://octopus.example.com
    octopusdeploy:apiKey:
        value: API-XXXXXXXXXXXXX

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Octopusdeploy = Pulumi.Octopusdeploy;

return await Deployment.RunAsync(() =>
{
    var dev = Octopusdeploy.GetSpace.Invoke(new()
    {
        Name = "Product Development",
    });

    //This resource will use the default space
    var development_environment = new Octopusdeploy.Environment("development-environment", new()
    {
        Name = "TestEnv1",
    });

    //This resource will be scoped to the space named "Product Development".
    var env3 = new Octopusdeploy.Environment("Env3", new()
    {
        SpaceId = dev.Apply(getSpaceResult => getSpaceResult.Id),
        Name = "TestEnv3",
    });

});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    octopusdeploy:address:
        value: https://octopus.example.com
    octopusdeploy:apiKey:
        value: API-XXXXXXXXXXXXX

```
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/octopusdeploy/octopusdeploy"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		dev, err := octopusdeploy.LookupSpace(ctx, &octopusdeploy.LookupSpaceArgs{
			Name: pulumi.StringRef("Product Development"),
		}, nil)
		if err != nil {
			return err
		}
		// This resource will use the default space
		_, err = octopusdeploy.NewEnvironment(ctx, "development-environment", &octopusdeploy.EnvironmentArgs{
			Name: pulumi.String("TestEnv1"),
		})
		if err != nil {
			return err
		}
		// This resource will be scoped to the space named "Product Development".
		_, err = octopusdeploy.NewEnvironment(ctx, "Env3", &octopusdeploy.EnvironmentArgs{
			SpaceId: pulumi.String(dev.Id),
			Name:    pulumi.String("TestEnv3"),
		})
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
    octopusdeploy:address:
        value: https://octopus.example.com
    octopusdeploy:apiKey:
        value: API-XXXXXXXXXXXXX

```
```yaml
resources:
  # /*
  # This resource will use the default space
  # */
  development-environment:
    type: octopusdeploy:Environment
    properties:
      name: TestEnv1
  # /*
  # This resource will be scoped to the space named "Product Development".
  # */
  env3:
    type: octopusdeploy:Environment
    name: Env3
    properties:
      spaceId: ${dev.id}
      name: TestEnv3
variables:
  dev:
    fn::invoke:
      function: octopusdeploy:getSpace
      arguments:
        name: Product Development
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    octopusdeploy:address:
        value: https://octopus.example.com
    octopusdeploy:apiKey:
        value: API-XXXXXXXXXXXXX

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.octopusdeploy.OctopusdeployFunctions;
import com.pulumi.octopusdeploy.inputs.GetSpaceArgs;
import com.pulumi.octopusdeploy.Environment;
import com.pulumi.octopusdeploy.EnvironmentArgs;
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
        final var dev = OctopusdeployFunctions.getSpace(GetSpaceArgs.builder()
            .name("Product Development")
            .build());

        //This resource will use the default space
        var development_environment = new Environment("development-environment", EnvironmentArgs.builder()
            .name("TestEnv1")
            .build());

        //This resource will be scoped to the space named "Product Development".
        var env3 = new Environment("env3", EnvironmentArgs.builder()
            .spaceId(dev.applyValue(getSpaceResult -> getSpaceResult.id()))
            .name("TestEnv3")
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

- `accessToken` (String) The OIDC Access Token to use with the Octopus REST API
- `address` (String) The endpoint of the Octopus REST API
- `apiKey` (String) The API key to use with the Octopus REST API
- `spaceId` (String) The space ID to target