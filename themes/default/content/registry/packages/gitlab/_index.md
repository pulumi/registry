---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-gitlab/v9.8.1/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/pulumi/pulumi-gitlab/blob/v9.8.1/docs/_index.md
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Gitlab Provider
meta_desc: Provides an overview on how to configure the Pulumi Gitlab provider.
layout: package
---

## Installation

The Gitlab provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/gitlab`](https://www.npmjs.com/package/@pulumi/gitlab)
* Python: [`pulumi-gitlab`](https://pypi.org/project/pulumi-gitlab/)
* Go: [`github.com/pulumi/pulumi-gitlab/sdk/v9/go/gitlab`](https://github.com/pulumi/pulumi-gitlab)
* .NET: [`Pulumi.Gitlab`](https://www.nuget.org/packages/Pulumi.Gitlab)
* Java: [`com.pulumi/gitlab`](https://central.sonatype.com/artifact/com.pulumi/gitlab)

## Overview

Use the GitLab provider to interact with GitLab resources, like
users, groups, projects and more. You must configure the provider with
the proper credentials before you can use it.

The provider uses the [`client-go`](https://gitlab.com/gitlab-org/api/client-go) library
to interact with the [GitLab REST API](https://docs.gitlab.com/api/api_resources/).

We support the following versions:

- Latest 3 patch releases within a major release. For example, if current is 17.8, we support 17.6-17.8. Or if current is 18.1, we support 18.0-18.1.
- We introduce any breaking changes on major releases only. For example, 17.0 or 18.0.
- We run tests against the latest 3 patch releases regardless of whether these cross a major release boundary. For example, if current is 17.8, we test 17.6-17.8. Or if current is 18.1, we test 17.11-18.1.

All other versions are best effort support.

We do not support experimental GitLab features until they are enabled by default or made Generally Available.

> Note, that the compatibility between a provider release and GitLab itself **cannot** be inferred from the
release version. New features added to GitLab may not be added to the provider until later versions.
Equally, features removed or deprecated in GitLab may not be removed or deprecated from the provider until later versions.

Each function and resource references the appropriate upstream GitLab REST API documentation,
which may be consumed to better understand the behavior of the API.

Use the navigation to the left to read about the valid functions and resources.




> Using a Project or Group access token may cause issues with some resources, as those token types don't
have full access to every API. This is also true when using a `CI_JOB_TOKEN`. Consider using a dedicated
Personal Access Token or Service Account if you are experiencing permission errors when adding resources.
## Authentication and Configuration

The configuration for the GitLab Provider can be derived from several sources,
which are applied in the following order:

1. Attributes in the provider configuration (see Schema section below)
2. Environment variables (see Schema section below)
3. (experimental) Configuration file (see <https://gitlab.com/gitlab-org/api/client-go#use-the-config-package-experimental>)
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    gitlab:token:
        value: 'TODO: var.gitlab_token'

```

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as gitlab from "@pulumi/gitlab";

// Add a project owned by the user
const sampleProject = new gitlab.Project("sample_project", {name: "example"});
// Add a hook to the project
const sampleProjectHook = new gitlab.ProjectHook("sample_project_hook", {
    project: sampleProject.id,
    url: "https://example.com/project_hook",
});
// Add a variable to the project
const sampleProjectVariable = new gitlab.ProjectVariable("sample_project_variable", {
    project: sampleProject.id,
    key: "project_variable_key",
    value: "project_variable_value",
});
// Add a deploy key to the project
const sampleDeployKey = new gitlab.DeployKey("sample_deploy_key", {
    project: sampleProject.id,
    title: "pulumi example",
    key: "ssh-ed25519 AAAA...",
});
// Add a group
const sampleGroup = new gitlab.Group("sample_group", {
    name: "example",
    path: "example",
    description: "An example group",
});
// Add a project to the group - example/example
const sampleGroupProject = new gitlab.Project("sample_group_project", {
    name: "example",
    namespaceId: sampleGroup.id,
});
```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    gitlab:token:
        value: 'TODO: var.gitlab_token'

```

```python
import pulumi
import pulumi_gitlab as gitlab

# Add a project owned by the user
sample_project = gitlab.Project("sample_project", name="example")
# Add a hook to the project
sample_project_hook = gitlab.ProjectHook("sample_project_hook",
    project=sample_project.id,
    url="https://example.com/project_hook")
# Add a variable to the project
sample_project_variable = gitlab.ProjectVariable("sample_project_variable",
    project=sample_project.id,
    key="project_variable_key",
    value="project_variable_value")
# Add a deploy key to the project
sample_deploy_key = gitlab.DeployKey("sample_deploy_key",
    project=sample_project.id,
    title="pulumi example",
    key="ssh-ed25519 AAAA...")
# Add a group
sample_group = gitlab.Group("sample_group",
    name="example",
    path="example",
    description="An example group")
# Add a project to the group - example/example
sample_group_project = gitlab.Project("sample_group_project",
    name="example",
    namespace_id=sample_group.id)
```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    gitlab:token:
        value: 'TODO: var.gitlab_token'

```

```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using GitLab = Pulumi.GitLab;

return await Deployment.RunAsync(() =>
{
    // Add a project owned by the user
    var sampleProject = new GitLab.Project("sample_project", new()
    {
        Name = "example",
    });

    // Add a hook to the project
    var sampleProjectHook = new GitLab.ProjectHook("sample_project_hook", new()
    {
        Project = sampleProject.Id,
        Url = "https://example.com/project_hook",
    });

    // Add a variable to the project
    var sampleProjectVariable = new GitLab.ProjectVariable("sample_project_variable", new()
    {
        Project = sampleProject.Id,
        Key = "project_variable_key",
        Value = "project_variable_value",
    });

    // Add a deploy key to the project
    var sampleDeployKey = new GitLab.DeployKey("sample_deploy_key", new()
    {
        Project = sampleProject.Id,
        Title = "pulumi example",
        Key = "ssh-ed25519 AAAA...",
    });

    // Add a group
    var sampleGroup = new GitLab.Group("sample_group", new()
    {
        Name = "example",
        Path = "example",
        Description = "An example group",
    });

    // Add a project to the group - example/example
    var sampleGroupProject = new GitLab.Project("sample_group_project", new()
    {
        Name = "example",
        NamespaceId = sampleGroup.Id,
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
    gitlab:token:
        value: 'TODO: var.gitlab_token'

```

```go
package main

import (
	"github.com/pulumi/pulumi-gitlab/sdk/v9/go/gitlab"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Add a project owned by the user
		sampleProject, err := gitlab.NewProject(ctx, "sample_project", &gitlab.ProjectArgs{
			Name: pulumi.String("example"),
		})
		if err != nil {
			return err
		}
		// Add a hook to the project
		_, err = gitlab.NewProjectHook(ctx, "sample_project_hook", &gitlab.ProjectHookArgs{
			Project: sampleProject.ID(),
			Url:     pulumi.String("https://example.com/project_hook"),
		})
		if err != nil {
			return err
		}
		// Add a variable to the project
		_, err = gitlab.NewProjectVariable(ctx, "sample_project_variable", &gitlab.ProjectVariableArgs{
			Project: sampleProject.ID(),
			Key:     pulumi.String("project_variable_key"),
			Value:   pulumi.String("project_variable_value"),
		})
		if err != nil {
			return err
		}
		// Add a deploy key to the project
		_, err = gitlab.NewDeployKey(ctx, "sample_deploy_key", &gitlab.DeployKeyArgs{
			Project: sampleProject.ID(),
			Title:   pulumi.String("pulumi example"),
			Key:     pulumi.String("ssh-ed25519 AAAA..."),
		})
		if err != nil {
			return err
		}
		// Add a group
		sampleGroup, err := gitlab.NewGroup(ctx, "sample_group", &gitlab.GroupArgs{
			Name:        pulumi.String("example"),
			Path:        pulumi.String("example"),
			Description: pulumi.String("An example group"),
		})
		if err != nil {
			return err
		}
		// Add a project to the group - example/example
		_, err = gitlab.NewProject(ctx, "sample_group_project", &gitlab.ProjectArgs{
			Name:        pulumi.String("example"),
			NamespaceId: sampleGroup.ID(),
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
    gitlab:token:
        value: 'TODO: var.gitlab_token'

```

```yaml
resources:
  # Add a project owned by the user
  sampleProject:
    type: gitlab:Project
    name: sample_project
    properties:
      name: example
  # Add a hook to the project
  sampleProjectHook:
    type: gitlab:ProjectHook
    name: sample_project_hook
    properties:
      project: ${sampleProject.id}
      url: https://example.com/project_hook
  # Add a variable to the project
  sampleProjectVariable:
    type: gitlab:ProjectVariable
    name: sample_project_variable
    properties:
      project: ${sampleProject.id}
      key: project_variable_key
      value: project_variable_value
  # Add a deploy key to the project
  sampleDeployKey:
    type: gitlab:DeployKey
    name: sample_deploy_key
    properties:
      project: ${sampleProject.id}
      title: pulumi example
      key: ssh-ed25519 AAAA...
  # Add a group
  sampleGroup:
    type: gitlab:Group
    name: sample_group
    properties:
      name: example
      path: example
      description: An example group
  # Add a project to the group - example/example
  sampleGroupProject:
    type: gitlab:Project
    name: sample_group_project
    properties:
      name: example
      namespaceId: ${sampleGroup.id}
```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    gitlab:token:
        value: 'TODO: var.gitlab_token'

```

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.gitlab.Project;
import com.pulumi.gitlab.ProjectArgs;
import com.pulumi.gitlab.ProjectHook;
import com.pulumi.gitlab.ProjectHookArgs;
import com.pulumi.gitlab.ProjectVariable;
import com.pulumi.gitlab.ProjectVariableArgs;
import com.pulumi.gitlab.DeployKey;
import com.pulumi.gitlab.DeployKeyArgs;
import com.pulumi.gitlab.Group;
import com.pulumi.gitlab.GroupArgs;
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
        // Add a project owned by the user
        var sampleProject = new Project("sampleProject", ProjectArgs.builder()
            .name("example")
            .build());

        // Add a hook to the project
        var sampleProjectHook = new ProjectHook("sampleProjectHook", ProjectHookArgs.builder()
            .project(sampleProject.id())
            .url("https://example.com/project_hook")
            .build());

        // Add a variable to the project
        var sampleProjectVariable = new ProjectVariable("sampleProjectVariable", ProjectVariableArgs.builder()
            .project(sampleProject.id())
            .key("project_variable_key")
            .value("project_variable_value")
            .build());

        // Add a deploy key to the project
        var sampleDeployKey = new DeployKey("sampleDeployKey", DeployKeyArgs.builder()
            .project(sampleProject.id())
            .title("pulumi example")
            .key("ssh-ed25519 AAAA...")
            .build());

        // Add a group
        var sampleGroup = new Group("sampleGroup", GroupArgs.builder()
            .name("example")
            .path("example")
            .description("An example group")
            .build());

        // Add a project to the group - example/example
        var sampleGroupProject = new Project("sampleGroupProject", ProjectArgs.builder()
            .name("example")
            .namespaceId(sampleGroup.id())
            .build());

    }
}
```

{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

- `baseUrl` (String) This is the target GitLab base API endpoint. Providing a value is a requirement when working with GitLab CE or GitLab Enterprise e.g. `https://my.gitlab.server/api/v4/`. It is optional to provide this value and it can also be sourced from the `GITLAB_BASE_URL` environment variable. The value must end with a slash.
- `cacertFile` (String) This is a file containing the ca cert to verify the gitlab instance. This is available for use when working with GitLab CE or Gitlab Enterprise with a locally-issued or self-signed certificate chain.
- `clientCert` (String) File path to client certificate when GitLab instance is behind company proxy. File must contain PEM encoded data.
- `clientKey` (String) File path to client key when GitLab instance is behind company proxy. File must contain PEM encoded data. Required when `clientCert` is set.
- `configFile` (String) The path to the configuration file to use. It may be sourced from the `GITLAB_CONFIG_FILE` environment variable.
- `context` (String) The context to use for authentication and configuration. The context must exist in the configuration file. It may be sourced from the `GITLAB_CONTEXT` environment variable.
- `earlyAuthCheck` (Boolean) (Experimental) By default the provider does a dummy request to get the current user in order to verify that the provider configuration is correct and the GitLab API is reachable. Set this to `false` to skip this check. This may be useful if the GitLab instance does not yet exist and is created within the same pulumi module. It may be sourced from the `GITLAB_EARLY_AUTH_CHECK`. This is an experimental feature and may change in the future. Please make sure to always keep backups of your state.
- `enableAutoCiSupport` (Boolean) If automatic CI support should be enabled or not. This only works when not providing a token.
- `headers` (Map of String) A map of headers to append to all API request to the GitLab instance.
- `insecure` (Boolean) When set to true this disables SSL verification of the connection to the GitLab instance.
- `retries` (Number) The number of retries to execute when receiving a 429 Rate Limit error. Each retry will exponentially back off.
- `token` (String, Sensitive) The OAuth2 Token, Project, Group, Personal Access Token or CI Job Token used to connect to GitLab. The OAuth method is used in this provider for authentication (using Bearer authorization token). See <https://docs.gitlab.com/api/#authentication> for details. It may be sourced from the `GITLAB_TOKEN` environment variable.