---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/jfrog/platform/2.2.5/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Platform Provider
meta_desc: Provides an overview on how to configure the Pulumi Platform provider.
layout: package
---

## Generate Provider

The Platform provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider jfrog/platform
```
## Overview

The [JFrog](https://jfrog.com/) Platform provider is used to interact with the features from [JFrog Platform REST API](https://jfrog.com/help/r/jfrog-rest-apis/jfrog-platform-rest-apis). The provider needs to be configured with the proper credentials before it can be used.

Links to documentation for specific resources can be found in the table of contents to the left.

This provider requires access to JFrog Platform APIs, which are only available in the *licensed* pro and enterprise editions. You can determine which license you have by accessing the following URL `${host}/artifactory/api/system/licenses/`

You can either access it via API, or web browser - it requires admin level credentials.

```bash
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
    artifactory:url:
        value: 'TODO: "${var.jfrog_url}"'
    platform:url:
        value: 'TODO: "${var.jfrog_url}"'

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as artifactory from "@pulumi/artifactory";
import * as platform from "@pulumi/platform";

const config = new pulumi.Config();
const jfrogUrl = config.get("jfrogUrl") || "http://localhost:8081";
const my_global_role = new platform.GlobalRole("my-global-role", {
    name: "my-global-role",
    description: "Test description",
    type: "CUSTOM_GLOBAL",
    environments: ["DEV"],
    actions: [
        "READ_REPOSITORY",
        "READ_BUILD",
    ],
});
const my_generic_local = new artifactory.index.LocalGenericRepository("my-generic-local", {key: "my-generic-local"});
const my_workers_service = new platform.WorkersService("my-workers-service", {
    key: "my-workers-service",
    enabled: true,
    description: "My workers service",
    sourceCode: "export default async (context: PlatformContext, data: BeforeDownloadRequest): Promise<BeforeDownloadResponse> => { console.log(await context.clients.platformHttp.get('/artifactory/api/system/ping')); console.log(await axios.get('https://my.external.resource')); return { status: 'DOWNLOAD_PROCEED', message: 'proceed', } }",
    action: "BEFORE_DOWNLOAD",
    filterCriteria: {
        artifactFilterCriteria: {
            repoKeys: [my_generic_local.key],
        },
    },
    secrets: [
        {
            key: "my-secret-key-1",
            value: "my-secret-value-1",
        },
        {
            key: "my-secret-key-2",
            value: "my-secret-value-2",
        },
    ],
});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    artifactory:url:
        value: 'TODO: "${var.jfrog_url}"'
    platform:url:
        value: 'TODO: "${var.jfrog_url}"'

```
```python
import pulumi
import pulumi_artifactory as artifactory
import pulumi_platform as platform

config = pulumi.Config()
jfrog_url = config.get("jfrogUrl")
if jfrog_url is None:
    jfrog_url = "http://localhost:8081"
my_global_role = platform.GlobalRole("my-global-role",
    name="my-global-role",
    description="Test description",
    type="CUSTOM_GLOBAL",
    environments=["DEV"],
    actions=[
        "READ_REPOSITORY",
        "READ_BUILD",
    ])
my_generic_local = artifactory.index.LocalGenericRepository("my-generic-local", key=my-generic-local)
my_workers_service = platform.WorkersService("my-workers-service",
    key="my-workers-service",
    enabled=True,
    description="My workers service",
    source_code="export default async (context: PlatformContext, data: BeforeDownloadRequest): Promise<BeforeDownloadResponse> => { console.log(await context.clients.platformHttp.get('/artifactory/api/system/ping')); console.log(await axios.get('https://my.external.resource')); return { status: 'DOWNLOAD_PROCEED', message: 'proceed', } }",
    action="BEFORE_DOWNLOAD",
    filter_criteria={
        "artifact_filter_criteria": {
            "repo_keys": [my_generic_local["key"]],
        },
    },
    secrets=[
        {
            "key": "my-secret-key-1",
            "value": "my-secret-value-1",
        },
        {
            "key": "my-secret-key-2",
            "value": "my-secret-value-2",
        },
    ])
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    artifactory:url:
        value: 'TODO: "${var.jfrog_url}"'
    platform:url:
        value: 'TODO: "${var.jfrog_url}"'

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Artifactory = Pulumi.Artifactory;
using Platform = Pulumi.Platform;

return await Deployment.RunAsync(() =>
{
    var config = new Config();
    var jfrogUrl = config.Get("jfrogUrl") ?? "http://localhost:8081";
    var my_global_role = new Platform.GlobalRole("my-global-role", new()
    {
        Name = "my-global-role",
        Description = "Test description",
        Type = "CUSTOM_GLOBAL",
        Environments = new[]
        {
            "DEV",
        },
        Actions = new[]
        {
            "READ_REPOSITORY",
            "READ_BUILD",
        },
    });

    var my_generic_local = new Artifactory.Index.LocalGenericRepository("my-generic-local", new()
    {
        Key = "my-generic-local",
    });

    var my_workers_service = new Platform.WorkersService("my-workers-service", new()
    {
        Key = "my-workers-service",
        Enabled = true,
        Description = "My workers service",
        SourceCode = "export default async (context: PlatformContext, data: BeforeDownloadRequest): Promise<BeforeDownloadResponse> => { console.log(await context.clients.platformHttp.get('/artifactory/api/system/ping')); console.log(await axios.get('https://my.external.resource')); return { status: 'DOWNLOAD_PROCEED', message: 'proceed', } }",
        Action = "BEFORE_DOWNLOAD",
        FilterCriteria = new Platform.Inputs.WorkersServiceFilterCriteriaArgs
        {
            ArtifactFilterCriteria = new Platform.Inputs.WorkersServiceFilterCriteriaArtifactFilterCriteriaArgs
            {
                RepoKeys = new[]
                {
                    my_generic_local.Key,
                },
            },
        },
        Secrets = new[]
        {
            new Platform.Inputs.WorkersServiceSecretArgs
            {
                Key = "my-secret-key-1",
                Value = "my-secret-value-1",
            },
            new Platform.Inputs.WorkersServiceSecretArgs
            {
                Key = "my-secret-key-2",
                Value = "my-secret-value-2",
            },
        },
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
    artifactory:url:
        value: 'TODO: "${var.jfrog_url}"'
    platform:url:
        value: 'TODO: "${var.jfrog_url}"'

```
```go
package main

import (
	"github.com/pulumi/pulumi-artifactory/sdk/go/artifactory"
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/platform/v2/platform"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")
		jfrogUrl := "http://localhost:8081"
		if param := cfg.Get("jfrogUrl"); param != "" {
			jfrogUrl = param
		}
		_, err := platform.NewGlobalRole(ctx, "my-global-role", &platform.GlobalRoleArgs{
			Name:        pulumi.String("my-global-role"),
			Description: pulumi.String("Test description"),
			Type:        pulumi.String("CUSTOM_GLOBAL"),
			Environments: pulumi.StringArray{
				pulumi.String("DEV"),
			},
			Actions: pulumi.StringArray{
				pulumi.String("READ_REPOSITORY"),
				pulumi.String("READ_BUILD"),
			},
		})
		if err != nil {
			return err
		}
		my_generic_local, err := artifactory.NewLocalGenericRepository(ctx, "my-generic-local", &artifactory.LocalGenericRepositoryArgs{
			Key: "my-generic-local",
		})
		if err != nil {
			return err
		}
		_, err = platform.NewWorkersService(ctx, "my-workers-service", &platform.WorkersServiceArgs{
			Key:         pulumi.String("my-workers-service"),
			Enabled:     pulumi.Bool(true),
			Description: pulumi.String("My workers service"),
			SourceCode:  pulumi.String("export default async (context: PlatformContext, data: BeforeDownloadRequest): Promise<BeforeDownloadResponse> => { console.log(await context.clients.platformHttp.get('/artifactory/api/system/ping')); console.log(await axios.get('https://my.external.resource')); return { status: 'DOWNLOAD_PROCEED', message: 'proceed', } }"),
			Action:      pulumi.String("BEFORE_DOWNLOAD"),
			FilterCriteria: &platform.WorkersServiceFilterCriteriaArgs{
				ArtifactFilterCriteria: &platform.WorkersServiceFilterCriteriaArtifactFilterCriteriaArgs{
					RepoKeys: pulumi.StringArray{
						my_generic_local.Key,
					},
				},
			},
			Secrets: platform.WorkersServiceSecretArray{
				&platform.WorkersServiceSecretArgs{
					Key:   pulumi.String("my-secret-key-1"),
					Value: pulumi.String("my-secret-value-1"),
				},
				&platform.WorkersServiceSecretArgs{
					Key:   pulumi.String("my-secret-key-2"),
					Value: pulumi.String("my-secret-value-2"),
				},
			},
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
    artifactory:url:
        value: 'TODO: "${var.jfrog_url}"'
    platform:url:
        value: 'TODO: "${var.jfrog_url}"'

```
```yaml
configuration:
  jfrogUrl:
    type: string
    default: http://localhost:8081
resources:
  my-global-role:
    type: platform:GlobalRole
    properties:
      name: my-global-role
      description: Test description
      type: CUSTOM_GLOBAL
      environments:
        - DEV
      actions:
        - READ_REPOSITORY
        - READ_BUILD
  my-generic-local:
    type: artifactory:LocalGenericRepository
    properties:
      key: my-generic-local
  my-workers-service:
    type: platform:WorkersService
    properties:
      key: my-workers-service
      enabled: true
      description: My workers service
      sourceCode: 'export default async (context: PlatformContext, data: BeforeDownloadRequest): Promise<BeforeDownloadResponse> => { console.log(await context.clients.platformHttp.get(''/artifactory/api/system/ping'')); console.log(await axios.get(''https://my.external.resource'')); return { status: ''DOWNLOAD_PROCEED'', message: ''proceed'', } }'
      action: BEFORE_DOWNLOAD
      filterCriteria:
        artifactFilterCriteria:
          repoKeys:
            - ${["my-generic-local"].key}
      secrets:
        - key: my-secret-key-1
          value: my-secret-value-1
        - key: my-secret-key-2
          value: my-secret-value-2
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    artifactory:url:
        value: 'TODO: "${var.jfrog_url}"'
    platform:url:
        value: 'TODO: "${var.jfrog_url}"'

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.platform.GlobalRole;
import com.pulumi.platform.GlobalRoleArgs;
import com.pulumi.artifactory.LocalGenericRepository;
import com.pulumi.artifactory.LocalGenericRepositoryArgs;
import com.pulumi.platform.WorkersService;
import com.pulumi.platform.WorkersServiceArgs;
import com.pulumi.platform.inputs.WorkersServiceFilterCriteriaArgs;
import com.pulumi.platform.inputs.WorkersServiceFilterCriteriaArtifactFilterCriteriaArgs;
import com.pulumi.platform.inputs.WorkersServiceSecretArgs;
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
        final var jfrogUrl = config.get("jfrogUrl").orElse("http://localhost:8081");
        var my_global_role = new GlobalRole("my-global-role", GlobalRoleArgs.builder()
            .name("my-global-role")
            .description("Test description")
            .type("CUSTOM_GLOBAL")
            .environments("DEV")
            .actions(
                "READ_REPOSITORY",
                "READ_BUILD")
            .build());

        var my_generic_local = new LocalGenericRepository("my-generic-local", LocalGenericRepositoryArgs.builder()
            .key("my-generic-local")
            .build());

        var my_workers_service = new WorkersService("my-workers-service", WorkersServiceArgs.builder()
            .key("my-workers-service")
            .enabled(true)
            .description("My workers service")
            .sourceCode("export default async (context: PlatformContext, data: BeforeDownloadRequest): Promise<BeforeDownloadResponse> => { console.log(await context.clients.platformHttp.get('/artifactory/api/system/ping')); console.log(await axios.get('https://my.external.resource')); return { status: 'DOWNLOAD_PROCEED', message: 'proceed', } }")
            .action("BEFORE_DOWNLOAD")
            .filterCriteria(WorkersServiceFilterCriteriaArgs.builder()
                .artifactFilterCriteria(WorkersServiceFilterCriteriaArtifactFilterCriteriaArgs.builder()
                    .repoKeys(my_generic_local.key())
                    .build())
                .build())
            .secrets(
                WorkersServiceSecretArgs.builder()
                    .key("my-secret-key-1")
                    .value("my-secret-value-1")
                    .build(),
                WorkersServiceSecretArgs.builder()
                    .key("my-secret-key-2")
                    .value("my-secret-value-2")
                    .build())
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Authentication

The JFrog Platform provider supports for the following types of authentication:
* Scoped token
* Pulumi Cloud OIDC provider
### Scoped Token

JFrog scoped tokens may be used via the HTTP Authorization header by providing the `accessToken` field to the provider configuration. Getting this value from the environment is supported with the `JFROG_ACCESS_TOKEN` environment variable.

Usage:
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    platform:accessToken:
        value: abc...xy
    platform:url:
        value: my.jfrog.io

```
## Configuration Reference

- `accessToken` (String, Sensitive) This is a access token that can be given to you by your admin under `Platform Configuration > User Management > Access Tokens`. This can also be sourced from the `JFROG_ACCESS_TOKEN` environment variable.
- `oidcProviderName` (String) OIDC provider name. See [Configure an OIDC Integration](https://jfrog.com/help/r/jfrog-platform-administration-documentation/configure-an-oidc-integration) for more details.
- `tfcCredentialTagName` (String) Pulumi Cloud Workload Identity Token tag name. Use for generating multiple TFC workload identity tokens. When set, the provider will attempt to use env var with this tag name as suffix. **Note:** this is case sensitive, so if set to `JFROG`, then env var `TFC_WORKLOAD_IDENTITY_TOKEN_JFROG` is used instead of `TFC_WORKLOAD_IDENTITY_TOKEN`. See Generating Multiple Tokens on HCP Pulumi for more details.
- `url` (String) JFrog Platform URL. This can also be sourced from the `JFROG_URL` environment variable.