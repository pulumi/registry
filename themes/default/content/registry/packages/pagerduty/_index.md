---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-pagerduty/v4.28.0/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/pulumi/pulumi-pagerduty/blob/v4.28.0/docs/_index.md
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Pagerduty Provider
meta_desc: Provides an overview on how to configure the Pulumi Pagerduty provider.
layout: package
---

## Installation

The Pagerduty provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/pagerduty`](https://www.npmjs.com/package/@pulumi/pagerduty)
* Python: [`pulumi-pagerduty`](https://pypi.org/project/pulumi-pagerduty/)
* Go: [`github.com/pulumi/pulumi-pagerduty/sdk/v4/go/pagerduty`](https://github.com/pulumi/pulumi-pagerduty)
* .NET: [`Pulumi.Pagerduty`](https://www.nuget.org/packages/Pulumi.Pagerduty)
* Java: [`com.pulumi/pagerduty`](https://central.sonatype.com/artifact/com.pulumi/pagerduty)

## Overview

[PagerDuty](https://www.pagerduty.com/) is an incident management platform that provides reliable notifications, automatic escalations, on-call scheduling, and other functionality to help teams detect and address unplanned work in real-time.

Use the navigation to the left to read about the available resources.
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    pagerduty:token:
        value: 'TODO: var.pagerduty_token'

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as pagerduty from "@pulumi/pagerduty";

// Create a PagerDuty team
const engineering = new pagerduty.Team("engineering", {
    name: "Engineering",
    description: "All engineering",
});
// Create a PagerDuty user
const earline = new pagerduty.User("earline", {
    name: "Earline Greenholt",
    email: "125.greenholt.earline@graham.name",
});
// Create a team membership
const earlineEngineering = new pagerduty.TeamMembership("earline_engineering", {
    userId: earline.id,
    teamId: engineering.id,
});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    pagerduty:token:
        value: 'TODO: var.pagerduty_token'

```
```python
import pulumi
import pulumi_pagerduty as pagerduty

# Create a PagerDuty team
engineering = pagerduty.Team("engineering",
    name="Engineering",
    description="All engineering")
# Create a PagerDuty user
earline = pagerduty.User("earline",
    name="Earline Greenholt",
    email="125.greenholt.earline@graham.name")
# Create a team membership
earline_engineering = pagerduty.TeamMembership("earline_engineering",
    user_id=earline.id,
    team_id=engineering.id)
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    pagerduty:token:
        value: 'TODO: var.pagerduty_token'

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Pagerduty = Pulumi.Pagerduty;

return await Deployment.RunAsync(() =>
{
    // Create a PagerDuty team
    var engineering = new Pagerduty.Team("engineering", new()
    {
        Name = "Engineering",
        Description = "All engineering",
    });

    // Create a PagerDuty user
    var earline = new Pagerduty.User("earline", new()
    {
        Name = "Earline Greenholt",
        Email = "125.greenholt.earline@graham.name",
    });

    // Create a team membership
    var earlineEngineering = new Pagerduty.TeamMembership("earline_engineering", new()
    {
        UserId = earline.Id,
        TeamId = engineering.Id,
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
    pagerduty:token:
        value: 'TODO: var.pagerduty_token'

```
```go
package main

import (
	"github.com/pulumi/pulumi-pagerduty/sdk/v4/go/pagerduty"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a PagerDuty team
		engineering, err := pagerduty.NewTeam(ctx, "engineering", &pagerduty.TeamArgs{
			Name:        pulumi.String("Engineering"),
			Description: pulumi.String("All engineering"),
		})
		if err != nil {
			return err
		}
		// Create a PagerDuty user
		earline, err := pagerduty.NewUser(ctx, "earline", &pagerduty.UserArgs{
			Name:  pulumi.String("Earline Greenholt"),
			Email: pulumi.String("125.greenholt.earline@graham.name"),
		})
		if err != nil {
			return err
		}
		// Create a team membership
		_, err = pagerduty.NewTeamMembership(ctx, "earline_engineering", &pagerduty.TeamMembershipArgs{
			UserId: earline.ID(),
			TeamId: engineering.ID(),
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
    pagerduty:token:
        value: 'TODO: var.pagerduty_token'

```
```yaml
resources:
  # Create a PagerDuty team
  engineering:
    type: pagerduty:Team
    properties:
      name: Engineering
      description: All engineering
  # Create a PagerDuty user
  earline:
    type: pagerduty:User
    properties:
      name: Earline Greenholt
      email: 125.greenholt.earline@graham.name
  # Create a team membership
  earlineEngineering:
    type: pagerduty:TeamMembership
    name: earline_engineering
    properties:
      userId: ${earline.id}
      teamId: ${engineering.id}
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    pagerduty:token:
        value: 'TODO: var.pagerduty_token'

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.pagerduty.Team;
import com.pulumi.pagerduty.TeamArgs;
import com.pulumi.pagerduty.User;
import com.pulumi.pagerduty.UserArgs;
import com.pulumi.pagerduty.TeamMembership;
import com.pulumi.pagerduty.TeamMembershipArgs;
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
        // Create a PagerDuty team
        var engineering = new Team("engineering", TeamArgs.builder()
            .name("Engineering")
            .description("All engineering")
            .build());

        // Create a PagerDuty user
        var earline = new User("earline", UserArgs.builder()
            .name("Earline Greenholt")
            .email("125.greenholt.earline@graham.name")
            .build());

        // Create a team membership
        var earlineEngineering = new TeamMembership("earlineEngineering", TeamMembershipArgs.builder()
            .userId(earline.id())
            .teamId(engineering.id())
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

The following configuration inputs are supported:

* `token` - (Optional) The v2 authorization token. It can also be sourced from the `PAGERDUTY_TOKEN` environment variable. See [API Documentation](https://developer.pagerduty.com/docs/ZG9jOjExMDI5NTUx-authentication)for more information.
* `userToken` - (Optional) The v2 user level authorization token. It can also be sourced from the `PAGERDUTY_USER_TOKEN` environment variable. See [API Documentation](https://developer.pagerduty.com/docs/ZG9jOjExMDI5NTUx-authentication) for more information.
* `useAppOauthScopedToken` - (Optional) Defines the configuration needed for making use of [App Oauth Scoped API token](https://developer.pagerduty.com/docs/e518101fde5f3-obtaining-an-app-o-auth-token) for authenticating API calls.
* `skipCredentialsValidation` - (Optional) Skip validation of the token against the PagerDuty API.
* `serviceRegion` - (Optional) The PagerDuty service region to use. Default to empty (uses US region). Supported value: `eu`. This setting also affects configuration of `useAppOauthScopedToken` for setting Region of *App Oauth token credentials*. It can also be sourced from the `PAGERDUTY_SERVICE_REGION` environment variable.
* `apiUrlOverride` - (Optional) It can be used to set a custom proxy endpoint as PagerDuty client api url overriding `serviceRegion` setup.
* `insecureTls` - (Optional) Can be used to disable TLS certificate checking when calling the PagerDuty API. This can be useful if you're behind a corporate proxy.

The `useAppOauthScopedToken` input has the following nested fields:

* `pdClientId` - (Required) An identifier issued when the Scoped OAuth client was added to a PagerDuty App. It can also be sourced from the `PAGERDUTY_CLIENT_ID` environment variable.
* `pdClientSecret` - (Required) A secret issued when the Scoped OAuth client was added to a PagerDuty App. It can also be sourced from the `PAGERDUTY_CLIENT_SECRET` environment variable.
* `pdSubdomain` - (Required) Your PagerDuty account subdomain; i.e: If the *URL* shown by the Browser when you are in your PagerDuty account is some like: <https://acme.pagerduty.com>, then your PagerDuty subdomain is `acme`. It can also be sourced from the `PAGERDUTY_SUBDOMAIN` environment variable.
## Example using App Oauth scoped token

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as pagerduty from "@pulumi/pagerduty";

// Create a PagerDuty team
const engineering = new pagerduty.Team("engineering", {
    name: "Engineering",
    description: "All engineering",
});
// Create a PagerDuty user
const earline = new pagerduty.User("earline", {
    name: "Earline Greenholt",
    email: "125.greenholt.earline@graham.name",
});
// Create a team membership
const earlineEngineering = new pagerduty.TeamMembership("earline_engineering", {
    userId: earline.id,
    teamId: engineering.id,
});
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
import pulumi_pagerduty as pagerduty

# Create a PagerDuty team
engineering = pagerduty.Team("engineering",
    name="Engineering",
    description="All engineering")
# Create a PagerDuty user
earline = pagerduty.User("earline",
    name="Earline Greenholt",
    email="125.greenholt.earline@graham.name")
# Create a team membership
earline_engineering = pagerduty.TeamMembership("earline_engineering",
    user_id=earline.id,
    team_id=engineering.id)
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
using Pagerduty = Pulumi.Pagerduty;

return await Deployment.RunAsync(() =>
{
    // Create a PagerDuty team
    var engineering = new Pagerduty.Team("engineering", new()
    {
        Name = "Engineering",
        Description = "All engineering",
    });

    // Create a PagerDuty user
    var earline = new Pagerduty.User("earline", new()
    {
        Name = "Earline Greenholt",
        Email = "125.greenholt.earline@graham.name",
    });

    // Create a team membership
    var earlineEngineering = new Pagerduty.TeamMembership("earline_engineering", new()
    {
        UserId = earline.Id,
        TeamId = engineering.Id,
    });

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
	"github.com/pulumi/pulumi-pagerduty/sdk/v4/go/pagerduty"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a PagerDuty team
		engineering, err := pagerduty.NewTeam(ctx, "engineering", &pagerduty.TeamArgs{
			Name:        pulumi.String("Engineering"),
			Description: pulumi.String("All engineering"),
		})
		if err != nil {
			return err
		}
		// Create a PagerDuty user
		earline, err := pagerduty.NewUser(ctx, "earline", &pagerduty.UserArgs{
			Name:  pulumi.String("Earline Greenholt"),
			Email: pulumi.String("125.greenholt.earline@graham.name"),
		})
		if err != nil {
			return err
		}
		// Create a team membership
		_, err = pagerduty.NewTeamMembership(ctx, "earline_engineering", &pagerduty.TeamMembershipArgs{
			UserId: earline.ID(),
			TeamId: engineering.ID(),
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

```
```yaml
resources:
  # Create a PagerDuty team
  engineering:
    type: pagerduty:Team
    properties:
      name: Engineering
      description: All engineering
  # Create a PagerDuty user
  earline:
    type: pagerduty:User
    properties:
      name: Earline Greenholt
      email: 125.greenholt.earline@graham.name
  # Create a team membership
  earlineEngineering:
    type: pagerduty:TeamMembership
    name: earline_engineering
    properties:
      userId: ${earline.id}
      teamId: ${engineering.id}
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
import com.pulumi.pagerduty.Team;
import com.pulumi.pagerduty.TeamArgs;
import com.pulumi.pagerduty.User;
import com.pulumi.pagerduty.UserArgs;
import com.pulumi.pagerduty.TeamMembership;
import com.pulumi.pagerduty.TeamMembershipArgs;
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
        // Create a PagerDuty team
        var engineering = new Team("engineering", TeamArgs.builder()
            .name("Engineering")
            .description("All engineering")
            .build());

        // Create a PagerDuty user
        var earline = new User("earline", UserArgs.builder()
            .name("Earline Greenholt")
            .email("125.greenholt.earline@graham.name")
            .build());

        // Create a team membership
        var earlineEngineering = new TeamMembership("earlineEngineering", TeamMembershipArgs.builder()
            .userId(earline.id())
            .teamId(engineering.id())
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}