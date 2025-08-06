---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/spectrocloud/spectrocloud/0.24.0/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Spectrocloud Provider
meta_desc: Provides an overview on how to configure the Pulumi Spectrocloud provider.
layout: package
---

## Generate Provider

The Spectrocloud provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider spectrocloud/spectrocloud
```
## Overview

The Spectro Cloud provider provides resources to interact with Palette and Palette VerteX through Infrastructure as code. The provider supports both SaaS and on-prem deployments of Palette and Palette VerteX.
## What is Palette?

Palette brings the managed Kubernetes experience to users' own unique enterprise
Kubernetes infrastructure stacks deployed in any public cloud, or private cloud environments. Palette allows users to
not have to trade-off between flexibility and manageability. Palette provides a platform-as-a-service experience
to users by automating the lifecycle of multiple Kubernetes clusters based on user-defined Kubernetes
infrastructure stacks.
## Palette Account

To get started with Palette, sign up for an account [here](https://www.spectrocloud.com/get-started).
Use your Palette [API key](https://docs.spectrocloud.com/user-management/authentication/api-key/create-api-key) to authenticate. For more details on the authentication, navigate to the authentication section.
## Example Usage

Create a `Pulumi.yaml` file with the following:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    spectrocloud:apiKey:
        value: 'TODO: var.sc_api_key'
    spectrocloud:host:
        value: 'TODO: var.sc_host'
    spectrocloud:projectName:
        value: 'TODO: var.sc_project_name'

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
    spectrocloud:apiKey:
        value: 'TODO: var.sc_api_key'
    spectrocloud:host:
        value: 'TODO: var.sc_host'
    spectrocloud:projectName:
        value: 'TODO: var.sc_project_name'

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
    spectrocloud:apiKey:
        value: 'TODO: var.sc_api_key'
    spectrocloud:host:
        value: 'TODO: var.sc_host'
    spectrocloud:projectName:
        value: 'TODO: var.sc_project_name'

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
    spectrocloud:apiKey:
        value: 'TODO: var.sc_api_key'
    spectrocloud:host:
        value: 'TODO: var.sc_host'
    spectrocloud:projectName:
        value: 'TODO: var.sc_project_name'

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
    spectrocloud:apiKey:
        value: 'TODO: var.sc_api_key'
    spectrocloud:host:
        value: 'TODO: var.sc_host'
    spectrocloud:projectName:
        value: 'TODO: var.sc_project_name'

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
    spectrocloud:apiKey:
        value: 'TODO: var.sc_api_key'
    spectrocloud:host:
        value: 'TODO: var.sc_host'
    spectrocloud:projectName:
        value: 'TODO: var.sc_project_name'

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

Copy `pulumi.template.tfvars` file to a `pulumi.tfvars` file and modify its content:

Be sure to populate the `scHost`, `scApiKey`, and other pulumi vars.

Copy one of the resource configuration files (e.g: spectrocloud_cluster_profile) from the *Resources* documentation. Be sure to specify
all required parameters.

Next, run pulumi using:

```console
pulumi up && pulumi up
```

Detailed schema definitions for each resource are listed in the *Resources* menu on the left.

For an end-to-end example of provisioning Spectro Cloud resources, visit:
Spectro Cloud E2E Examples.
## Environment Variables

Credentials and other configurations can be provided through environment variables. The following environment variables are availabe.

- `SPECTROCLOUD_HOST`
- `SPECTROCLOUD_APIKEY`
- `SPECTROCLOUD_TRACE`
- `SPECTROCLOUD_RETRY_ATTEMPTS`
## Authentication
You can use an API key to authenticate with Spectro Cloud. Visit the User Management API Key [documentation](https://docs.spectrocloud.com/user-management/user-authentication/#usingapikey) to learn more about Spectro Cloud API keys.
```shell
export SPECTROCLOUD_APIKEY=5b7aad.........
```

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:

```
## Import
Starting with Pulumi v1.5.0 and later, you can use an `import` block to import resources into your state file.

Each resource type has its own specific requirements for the import process. We recommend you refer to the documentation for each resource to better understand the exact format of the `id` and any other required parameters.
The `import` block specifies the resource you want to import and its unique identifier with the following structure:

- `<resource>`: The type of the resource you are importing.
- `<name>`: A name you assign to the resource within your Pulumi configuration.
- `<unique_identifier>`: The ID of the resource you are importing. This can include additional context if required.
### Examples

The following examples showcase how to import a resource. Some resource requires the context to be specified during the import action. The context refers to the Palette scope. Allowed values are either `project` or `tenant`.
#### Import With Context

When importing resources that require additional context, the `id` is followed by a context, separated by a colon.

You can also import a resource using the Pulumi CLI and the `import` command.

```console
pulumi import spectrocloud_cluster_aks.example example_id:project
```

    Specify' tenant' after the colon if you want to import a resource at the tenant scope.

Example of importing a resource with the tenant context through the Pulumi CLI.

```console
pulumi import spectrocloud_cluster_aks.example example_id:tenant
```

> Ensure you have tenant admin access when importing a resource at the tenant scope.
#### Import Without Context

For resources that do not require additional context, the `id` is the only provided argument. The following is an example of a resource that does not require the context and only provides the ID.

Below is an example of using the Pulumi CLI and the `import` command without specifying the context.

```console
pulumi import spectrocloud_cluster_profile.example id
```
## Support

For questions or issues with the provider, open up an issue in the provider GitHub discussion board.
## Configuration Reference

- `apiKey` (String, Sensitive) The Spectro Cloud API key. Can also be set with the `SPECTROCLOUD_APIKEY` environment variable.
- `host` (String) The Spectro Cloud API host url. Can also be set with the `SPECTROCLOUD_HOST` environment variable. Defaults to <https://api.spectrocloud.com>
- `ignoreInsecureTlsError` (Boolean) Ignore insecure TLS errors for Spectro Cloud API endpoints. Defaults to false.
- `projectName` (String) The Palette project the provider will target. If no value is provided, the `Default` Palette project is used. The default value is `Default`.
- `retryAttempts` (Number) Number of retry attempts. Can also be set with the `SPECTROCLOUD_RETRY_ATTEMPTS` environment variable. Defaults to 10.
- `trace` (Boolean) Enable HTTP request tracing. Can also be set with the `SPECTROCLOUD_TRACE` environment variable. To enable Pulumi debug logging, set `TF_LOG=DEBUG`. Visit the Pulumi documentation to learn more about Pulumi debugging.