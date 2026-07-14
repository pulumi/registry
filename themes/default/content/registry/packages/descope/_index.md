---
# WARNING: this file was fetched from https://raw.githubusercontent.com/descope/pulumi-descope/v0.3.13/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/descope/pulumi-descope/blob/v0.3.13/docs/_index.md
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Descope Provider
meta_desc: Provides an overview on how to configure the Pulumi Descope provider.
layout: package
---

## Installation

The Descope provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/descope`](https://www.npmjs.com/package/@pulumi/descope)
* Python: [`pulumi-descope`](https://pypi.org/project/pulumi-descope/)
* Go: [`github.com/descope/pulumi-descope/sdk/go/descope`](https://github.com/pulumi/pulumi-descope)
* .NET: [`Pulumi.Descope`](https://www.nuget.org/packages/Pulumi.Descope)
* Java: [`com.pulumi/descope`](https://central.sonatype.com/artifact/com.pulumi/descope)

## Overview

The [Descope](https://www.descope.com) Pulumi Provider lets you manage your Descope project configuration as infrastructure-as-code. Configure authentication methods, define roles and permissions, set up third-party connectors, manage flows, and more—all declaratively in Pulumi.

Descope is an authentication and user management platform. The Pulumi provider manages *project configuration* (how your project behaves), not users or tenants (use the [Descope Management API](https://docs.descope.com/api/openapi) or [SDKs](https://docs.descope.com) for those).
## Requirements

- Pulumi 1.0 or later
- A Descope **Pro or Enterprise** plan
- A **Management Key** from [Company Settings](https://app.descope.com/settings/company) in the Descope console
## Authentication

The provider authenticates with the Descope API using a management key. You can create management keys in the [Company Settings](https://app.descope.com/settings/company) section of the Descope console.

> **Security Note:** Never hardcode your management key in Pulumi configuration files, as this risks exposing it in version control. Use environment variables or a secrets manager instead.
### Environment Variables

|         Variable         |                        Description                        |
|--------------------------|-----------------------------------------------------------|
| `DESCOPE_MANAGEMENT_KEY` | A valid management key for your Descope company           |
| `DESCOPE_BASE_URL`       | Override the Descope API base URL (optional, for testing) |

```shell
export DESCOPE_MANAGEMENT_KEY="K2..."
pulumi preview
```
## Example Usage
### Minimal Configuration

{{< chooser language "typescript,python,go,csharp,java,yaml,hcl" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs

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

```

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import java.util.ArrayList;
import java.util.Arrays;
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
{{% choosable language hcl %}}
```hcl
Example currently unavailable in this language
```

{{% /choosable %}}
{{< /chooser >}}
### Creating a Project

{{< chooser language "typescript,python,go,csharp,java,yaml,hcl" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as descope from "@descope/pulumi-descope";

const example = new descope.Project("example", {
    name: "my-app",
    environment: "production",
    tags: ["prod"],
});
```

{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import descope_pulumi as descope

example = descope.Project("example",
    name="my-app",
    environment="production",
    tags=["prod"])
```

{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Descope = Descope.Pulumi.Descope;

return await Deployment.RunAsync(() =>
{
    var example = new Descope.Project("example", new()
    {
        Name = "my-app",
        Environment = "production",
        Tags = new[]
        {
            "prod",
        },
    });

});

```

{{% /choosable %}}
{{% choosable language go %}}
```go
package main

import (
	"github.com/descope/pulumi-descope/sdk/go/descope"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := descope.NewProject(ctx, "example", &descope.ProjectArgs{
			Name:        pulumi.String("my-app"),
			Environment: pulumi.String("production"),
			Tags: pulumi.StringArray{
				pulumi.String("prod"),
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
resources:
  example:
    type: descope:Project
    properties:
      name: my-app
      environment: production
      tags:
        - prod
```

{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.descope.pulumi.descope.Project;
import com.descope.pulumi.descope.ProjectArgs;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Map;
import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        var example = new Project("example", ProjectArgs.builder()
            .name("my-app")
            .environment("production")
            .tags("prod")
            .build());

    }
}
```

{{% /choosable %}}
{{% choosable language hcl %}}
```hcl
pulumi {
  required_providers {
    descope = {
      source = "pulumi/descope"
    }
  }
}

resource "descope_project" "example" {
  name        = "my-app"
  environment = "production"
  tags        = ["prod"]
}
```

{{% /choosable %}}
{{< /chooser >}}
### Full Provider Configuration

If you need to explicitly configure the provider (e.g. in a module or when not using environment variables):

{{< chooser language "typescript,python,go,csharp,java,yaml,hcl" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    descope:managementKey:
        value: 'TODO: var.descope_management_key'

```

```typescript
import * as pulumi from "@pulumi/pulumi";

const config = new pulumi.Config();
const descopeManagementKey = config.require("descopeManagementKey");
```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    descope:managementKey:
        value: 'TODO: var.descope_management_key'

```

```python
import pulumi

config = pulumi.Config()
descope_management_key = config.require("descopeManagementKey")
```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    descope:managementKey:
        value: 'TODO: var.descope_management_key'

```

```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;

return await Deployment.RunAsync(() =>
{
    var config = new Config();
    var descopeManagementKey = config.Require("descopeManagementKey");
});

```

{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    descope:managementKey:
        value: 'TODO: var.descope_management_key'

```

```go
package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")
		descopeManagementKey := cfg.Require("descopeManagementKey")
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
    descope:managementKey:
        value: 'TODO: var.descope_management_key'

```

```yaml
configuration:
  descopeManagementKey:
    type: string
```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    descope:managementKey:
        value: 'TODO: var.descope_management_key'

```

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import java.util.ArrayList;
import java.util.Arrays;
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
        final var descopeManagementKey = config.require("descopeManagementKey");
    }
}
```

{{% /choosable %}}
{{% choosable language hcl %}}
```hcl
variable "descopeManagementKey" {
  type = string
}
```

{{% /choosable %}}
{{< /chooser >}}
## Resources

|                       Resource                       |                                  Description                                  |
|------------------------------------------------------|-------------------------------------------------------------------------------|
| `descope.Project`               | Full project configuration: authentication, RBAC, connectors, flows, and more |
| `descope.ManagementKey` | Programmatic management keys with scoped permissions                          |
| `descope.Descoper`             | Descope console user accounts with role assignments                           |
## Guides

- Quickstart – Set up the provider and manage your first project
## Configuration Reference

- `baseUrl` (String) An optional base URL for the Descope API
- `managementKey` (String, Sensitive) A valid management key for your Descope company
- `projectId` (String, Deprecated)