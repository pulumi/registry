---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumiverse/pulumi-cpln/v0.0.73/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/pulumiverse/pulumi-cpln/blob/v0.0.73/docs/_index.md
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Control Plane Provider
meta_desc: Provides an overview on how to configure the Pulumi Control Plane provider.
layout: package
---

## Installation

The Control Plane provider is available as a package in all Pulumi languages:

- JavaScript/TypeScript: [`@pulumiserve/cpln`](https://www.npmjs.com/package/@pulumiverse/cpln)
- Python: [`pulumiverse-cpln`](https://pypi.org/project/pulumiverse-cpln/)
- Go: [`github.com/pulumiverse/pulumi-cpln/sdk/go/cpln`](https://pkg.go.dev/github.com/pulumiverse/pulumi-cpln/sdk)
- .NET: [`Pulumiverse.cpln`](https://www.nuget.org/packages/Pulumiverse.cpln)

## Overview

The Control Plane Pulumi Provider Plugin enables the scaffolding of any Control Plane object as code using HCL. It enables infrastructure as code with all the added benefit of the global virtual cloud (GVC). You can build your VPCs, subnets, databases, queues, caches, etc. and overlay them with a multi-cloud/multi-region universal compute workloads that span regions and clouds. Nearly everything you can do using the Control Plane CLI, UI or API is available using Pulumi.

Each header below (i.e., `cpln.Agent`) corresponds to a resource within the Control Plane Pulumi provider.

## Authentication

Authenticate using one of the following methods:

`1. CLI`

- [Install the CLI](https://docs.controlplane.com/reference/cli#installation) and execute the command `cpln login`. After a successful login, the Pulumi provider will use the `default` profile to authenticate. To use a different profile, set the `profile` variable when initializing the provider or set the `CPLN_PROFILE` environment variable.

`2. Token`

- The `token` variable can be set when initializing the provider or by setting the `CPLN_TOKEN` environment variable.
- The value of `token` can be either:
  - The output of running the command `cpln profile token PROFILE_NAME`, or
  - In the case of a [Service Account](https://docs.controlplane.com/reference/serviceaccount), the value of one of it's [keys](https://docs.controlplane.com/reference/serviceaccount#keys)

`3. Refresh Token`

- The `refreshToken` variable is used when the provider is required to create an org or update the `authConfig` property using the `cpln.Org` resource. The `refreshToken` variable can be set when initializing the provider or by setting the `CPLN_REFRESH_TOKEN` environment variable.
- When creating an org, the `refreshToken` **must** belong to a user that has the `orgCreator` role for the associated account.
- When updating the org `authConfig` property, the `refreshToken` **must** belong to a user that was authenticated using SAML.
- The `refreshToken` can be obtained by following these steps:
  - Using the CLI, authenticate with a user account by executing `cpln login`.
  - Browser to the path `~/.config/cpln/profiles`. This path will contain JSON files corresponding to the name of the profile (i.e., `default.json`).
  - The contents of the JSON file will contain a key named `refreshToken`. Use the value of this key for the `refreshToken` variable.

> **Note** To perform automated tasks using Pulumi, the preferred method is to use a `Service Account` and one of it's `keys` as the `token` value.

## Provider Declaration

### Required

- **org** (String) The Control Plane org that this provider will perform actions against. Can be specified with the `CPLN_ORG` environment variable.

- **endpoint** (String) The Control Plane Data Service API endpoint. Default is: `https://api.cpln.io`. Can be specified with the `CPLN_ENDPOINT` environment variable.
- **profile** (String) The user/service account profile that this provider will use to authenticate to the data service. Can be specified with the `CPLN_PROFILE` environment variable.
- **token** (String) A generated token that can be used to authenticate to the data service API. Can be specified with the `CPLN_TOKEN` environment variable.
- **refresh_token** (String) A generated token that can be used to authenticate to the data service API. Can be specified with the `CPLN_REFRESH_TOKEN` environment variable. Used when the provider is required to create an org or update the `authConfig` property. Refer to the section above on how to obtain the refresh token.

> **Note** If the `token` or `refreshToken` value is empty, the Control Plane CLI (cpln) must be installed and the `cpln login` command must be used to authenticate.

## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
  cpln:endpoint:
    value: "TODO: var.endpoint"
  cpln:org:
    value: "TODO: var.org"
  cpln:profile:
    value: "TODO: var.profile"
  cpln:refreshToken:
    value: "TODO: var.refresh_token"
  cpln:token:
    value: "TODO: var.token"
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
  cpln:endpoint:
    value: "TODO: var.endpoint"
  cpln:org:
    value: "TODO: var.org"
  cpln:profile:
    value: "TODO: var.profile"
  cpln:refreshToken:
    value: "TODO: var.refresh_token"
  cpln:token:
    value: "TODO: var.token"
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
  cpln:endpoint:
    value: "TODO: var.endpoint"
  cpln:org:
    value: "TODO: var.org"
  cpln:profile:
    value: "TODO: var.profile"
  cpln:refreshToken:
    value: "TODO: var.refresh_token"
  cpln:token:
    value: "TODO: var.token"
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
  cpln:endpoint:
    value: "TODO: var.endpoint"
  cpln:org:
    value: "TODO: var.org"
  cpln:profile:
    value: "TODO: var.profile"
  cpln:refreshToken:
    value: "TODO: var.refresh_token"
  cpln:token:
    value: "TODO: var.token"
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
  cpln:endpoint:
    value: "TODO: var.endpoint"
  cpln:org:
    value: "TODO: var.org"
  cpln:profile:
    value: "TODO: var.profile"
  cpln:refreshToken:
    value: "TODO: var.refresh_token"
  cpln:token:
    value: "TODO: var.token"
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
  cpln:endpoint:
    value: "TODO: var.endpoint"
  cpln:org:
    value: "TODO: var.org"
  cpln:profile:
    value: "TODO: var.profile"
  cpln:refreshToken:
    value: "TODO: var.refresh_token"
  cpln:token:
    value: "TODO: var.token"
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
