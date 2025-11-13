---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-artifactory/v8.9.2/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/pulumi/pulumi-artifactory/blob/v8.9.2/docs/_index.md
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Artifactory Provider
meta_desc: Provides an overview on how to configure the Pulumi Artifactory provider.
layout: package
---

## Installation

The Artifactory provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/artifactory`](https://www.npmjs.com/package/@pulumi/artifactory)
* Python: [`pulumi-artifactory`](https://pypi.org/project/pulumi-artifactory/)
* Go: [`github.com/pulumi/pulumi-artifactory/sdk/v8/go/artifactory`](https://github.com/pulumi/pulumi-artifactory)
* .NET: [`Pulumi.Artifactory`](https://www.nuget.org/packages/Pulumi.Artifactory)
* Java: [`com.pulumi/artifactory`](https://central.sonatype.com/artifact/com.pulumi/artifactory)

## Overview

The [Artifactory](https://jfrog.com/artifactory/) provider is used to interact with the resources supported by Artifactory. The provider needs to be configured with the proper credentials before it can be used.

Links to documentation for specific resources can be found in the table of contents to the left.

This provider requires access to Artifactory APIs, which are only available in the *licensed* pro and enterprise editions. You can determine which license you have by accessing the following the URL `${host}/artifactory/api/system/licenses/`.

You can either access it via API, or web browser - it require admin level credentials.

```sh
curl -sL ${host}/artifactory/api/system/licenses/ | jq .
{
  "type" : "Enterprise Plus Trial",
  "validThrough" : "Jan 29, 2022",
  "licensedTo" : "JFrog Ltd"
}
```
## Example Usage
{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    artifactory:accessToken:
        value: 'TODO: "${var.artifactory_access_token}"'
    artifactory:url:
        value: 'TODO: "${var.artifactory_url}/artifactory"'

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as artifactory from "@pulumi/artifactory";

// Create a new repository
const pypi_libs = new artifactory.LocalPypiRepository("pypi-libs", {
    key: "pypi-libs",
    repoLayoutRef: "simple-default",
    description: "A pypi repository for python packages",
});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    artifactory:accessToken:
        value: 'TODO: "${var.artifactory_access_token}"'
    artifactory:url:
        value: 'TODO: "${var.artifactory_url}/artifactory"'

```
```python
import pulumi
import pulumi_artifactory as artifactory

# Create a new repository
pypi_libs = artifactory.LocalPypiRepository("pypi-libs",
    key="pypi-libs",
    repo_layout_ref="simple-default",
    description="A pypi repository for python packages")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    artifactory:accessToken:
        value: 'TODO: "${var.artifactory_access_token}"'
    artifactory:url:
        value: 'TODO: "${var.artifactory_url}/artifactory"'

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Artifactory = Pulumi.Artifactory;

return await Deployment.RunAsync(() =>
{
    // Create a new repository
    var pypi_libs = new Artifactory.LocalPypiRepository("pypi-libs", new()
    {
        Key = "pypi-libs",
        RepoLayoutRef = "simple-default",
        Description = "A pypi repository for python packages",
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
    artifactory:accessToken:
        value: 'TODO: "${var.artifactory_access_token}"'
    artifactory:url:
        value: 'TODO: "${var.artifactory_url}/artifactory"'

```
```go
package main

import (
	"github.com/pulumi/pulumi-artifactory/sdk/v8/go/artifactory"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a new repository
		_, err := artifactory.NewLocalPypiRepository(ctx, "pypi-libs", &artifactory.LocalPypiRepositoryArgs{
			Key:           pulumi.String("pypi-libs"),
			RepoLayoutRef: pulumi.String("simple-default"),
			Description:   pulumi.String("A pypi repository for python packages"),
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
    artifactory:accessToken:
        value: 'TODO: "${var.artifactory_access_token}"'
    artifactory:url:
        value: 'TODO: "${var.artifactory_url}/artifactory"'

```
```yaml
resources:
  # Create a new repository
  pypi-libs:
    type: artifactory:LocalPypiRepository
    properties:
      key: pypi-libs
      repoLayoutRef: simple-default
      description: A pypi repository for python packages
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    artifactory:accessToken:
        value: 'TODO: "${var.artifactory_access_token}"'
    artifactory:url:
        value: 'TODO: "${var.artifactory_url}/artifactory"'

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.artifactory.LocalPypiRepository;
import com.pulumi.artifactory.LocalPypiRepositoryArgs;
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
        // Create a new repository
        var pypi_libs = new LocalPypiRepository("pypi-libs", LocalPypiRepositoryArgs.builder()
            .key("pypi-libs")
            .repoLayoutRef("simple-default")
            .description("A pypi repository for python packages")
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Authentication

The Artifactory provider supports two ways of authentication. The following methods are supported:
* Access Token
* API Key (deprecated)
### Access Token

Artifactory access tokens may be used via the Authorization header by providing the `accessToken` attribute to the provider configuration. Getting this value from the environment is supported with `JFROG_ACCESS_TOKEN` or `ARTIFACTORY_ACCESS_TOKEN` variables.

Usage:
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    artifactory:accessToken:
        value: abc...xy
    artifactory:url:
        value: artifactory.site.com/artifactory

```
### API Key (deprecated)

!>An upcoming version will support the option to block the usage/creation of API Keys (for admins to set on their platform). In a future version (scheduled for end of Q3, 2023), the option to disable the usage/creation of API Keys will be available and set to disabled by default. Admins will be able to enable the usage/creation of API Keys. By end of Q4 2024, API Keys will be deprecated all together and the option to use them will no longer be available. See [JFrog API Key Deprecation Process](https://jfrog.com/help/r/jfrog-platform-administration-documentation/jfrog-api-key-deprecation-process).

~>If `accessToken` attribute, `JFROG_ACCESS_TOKEN` or `ARTIFACTORY_ACCESS_TOKEN` environment variable is set, the provider will ignore `apiKey` attribute.

Artifactory API keys may be used via the `X-JFrog-Art-Api` header by providing the `apiKey` attribute in the provider configuration.

Usage:
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    artifactory:apiKey:
        value: abc...xy
    artifactory:url:
        value: artifactory.site.com/artifactory

```
## Configuration Reference

The following configuration inputs are supported:

* `url` - (Optional) URL of Artifactory. This can also be sourced from the `JFROG_URL` or `ARTIFACTORY_URL` environment variable.
* `accessToken` - (Optional) This can also be sourced from `JFROG_ACCESS_TOKEN` or `ARTIFACTORY_ACCESS_TOKEN` environment variables.
* `apiKey` - (Optional, deprecated) API key for api auth.
* `oidcProviderName` - (Optional) OIDC provider name. See [Configure an OIDC Integration](https://jfrog.com/help/r/jfrog-platform-administration-documentation/configure-an-oidc-integration) for more details.