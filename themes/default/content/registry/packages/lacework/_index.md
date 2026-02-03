---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/lacework/lacework/2.1.1/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Lacework Provider
meta_desc: Provides an overview on how to configure the Pulumi Lacework provider.
layout: package
---

## Generate Provider

The Lacework provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider lacework/lacework
```
## Overview

The Lacework provider is used to interact with the Lacework cloud security platform.
The provider needs to be configured with the proper credentials before it can be used.

Use the left navigation panel to read about the available resources.
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    lacework:profile:
        value: my-profile

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as lacework from "@pulumi/lacework";

// Connect an AWS account to Lacework for configuration and compliance assessment
const accountAbc = new lacework.IntegrationAwsCfg("account_abc", {});
// Configure Lacework to forward alerts to a Slack channel
const critical = new lacework.AlertChannelSlack("critical", {});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    lacework:profile:
        value: my-profile

```
```python
import pulumi
import pulumi_lacework as lacework

# Connect an AWS account to Lacework for configuration and compliance assessment
account_abc = lacework.IntegrationAwsCfg("account_abc")
# Configure Lacework to forward alerts to a Slack channel
critical = lacework.AlertChannelSlack("critical")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    lacework:profile:
        value: my-profile

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Lacework = Pulumi.Lacework;

return await Deployment.RunAsync(() =>
{
    // Connect an AWS account to Lacework for configuration and compliance assessment
    var accountAbc = new Lacework.IntegrationAwsCfg("account_abc");

    // Configure Lacework to forward alerts to a Slack channel
    var critical = new Lacework.AlertChannelSlack("critical");

});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    lacework:profile:
        value: my-profile

```
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/lacework/v2/lacework"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Connect an AWS account to Lacework for configuration and compliance assessment
		_, err := lacework.NewIntegrationAwsCfg(ctx, "account_abc", nil)
		if err != nil {
			return err
		}
		// Configure Lacework to forward alerts to a Slack channel
		_, err = lacework.NewAlertChannelSlack(ctx, "critical", nil)
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
    lacework:profile:
        value: my-profile

```
```yaml
resources:
  # Connect an AWS account to Lacework for configuration and compliance assessment
  accountAbc:
    type: lacework:IntegrationAwsCfg
    name: account_abc
  # Configure Lacework to forward alerts to a Slack channel
  critical:
    type: lacework:AlertChannelSlack
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    lacework:profile:
        value: my-profile

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.lacework.IntegrationAwsCfg;
import com.pulumi.lacework.AlertChannelSlack;
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
        // Connect an AWS account to Lacework for configuration and compliance assessment
        var accountAbc = new IntegrationAwsCfg("accountAbc");

        // Configure Lacework to forward alerts to a Slack channel
        var critical = new AlertChannelSlack("critical");

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
# Authentication

The Lacework provider can be configured with the proper credentials via the following supported methods:

* Static credentials
* Environment variables
* Configuration file
## Static credentials

!> **Warning:** Hard-coding credentials into any Pulumi configuration is not
recommended. Secrets could be leaked by committing this file to a public version
control system.

Static credentials can be provided by adding the `account`, `apiKey`, and `apiSecret` in-line in the
Lacework provider configuration:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    lacework:account:
        value: my-account
    lacework:apiKey:
        value: my-api-key
    lacework:apiSecret:
        value: my-api-secret

```
### Using an API Access Token

For short-lived workflows, Lacework provides an alternative method to configure the provider by using an
API access token (`apiToken`), note that all API tokens from the Lacework platform are short-lived which
means that once the token expires, a new one needs to be generated and configured.

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    lacework:account:
        value: my-account
    lacework:apiToken:
        value: my-api-token

```

> **Note:** You can use the [Lacework CLI](https://docs.lacework.com/cli) command `lacework access-token` to
generate an API access token and the command `lacework configure show account` to display your configured account.
## Environment Variables

You can provide your credentials via the `LW_ACCOUNT`, `LW_API_KEY`, and `LW_API_SECRET` environment
variables, they represent your Lacework account subdomain of URL, Lacework API access key, and Lacework
API access secret, respectively.

> **Note:** Setting your Lacework credentials using these environment variables will override the use of `LW_PROFILE`.

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:

```

**Bash:**
```
export LW_ACCOUNT="my-account"
export LW_API_KEY="my-api-key"
export LW_API_SECRET="my-api-secret"
```

**Powershell:**
```
$Env:LW_ACCOUNT = "my-account"
$Env:LW_API_KEY = "my-api-key"
$Env:LW_API_SECRET = "my-api-secret"
```
### Using an API Access Token

For short-lived workflows, Lacework provides an alternative method to configure the provider by using an
API access token (`LW_API_TOKEN`), note that all API tokens from the Lacework platform are short-lived which
means that once the token expires, a new one needs to be generated and configured.

**Bash:**
```
export LW_ACCOUNT="my-account"
export LW_API_TOKEN="my-api-token"
```

**Powershell:**
```
$Env:LW_ACCOUNT = "my-account"
$Env:LW_API_TOKEN = "my-api-token"
```

> **Note:** You can use the [Lacework CLI](https://docs.lacework.com/cli) command `lacework access-token` to
generate an API access token and the command `lacework configure show account` to display your configured account.
## Configuration file

It is possible to use credentials from the Lacework configuration file. The default location on Linux and OS X
is `$HOME/.lacework.toml`, and for Windows users is `"%USERPROFILE%\.lacework.toml"`. This configuration file
can be easily managed using the [Lacework CLI](https://docs.lacework.com/cli). This
method also supports a `profile` configuration and matching `LW_PROFILE` environment variable.

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    lacework:profile:
        value: custom-profile

```
# Organizational Accounts

An organization can contain multiple accounts so you can manage components such as alerts, resource groups,
team members, and audit logs at a more granular level inside an organization. A team member may have access
to multiple accounts and can easily switch between them.

> **Note:** To enroll your account in an organization follow the [Organization Enrollment Process](https://docs.lacework.com/console/organization-overview) documentation.

Use the argument `subaccount` to switch to a different account inside your organizational account.

For example, having a `default` profile that has access to your primary account named `my-company`:
```toml
[default]
  account = "my-company"
  api_key = "my-api-key"
  api_secret = "my-api-secret"
  version = 2
```

To access your sub-account named `business-unit` you would specify the argument `subaccount`.

From there, you can pass the `alias` meta-argument to any resource or module to switch between accounts.

For a resource or function, set its `provider` meta-argument to a `lacework.<ALIAS>` reference:
{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as lacework from "@pulumi/lacework";

const primaryCritical = new lacework.AlertChannelSlack("primary_critical", {});
const businessUnitCritical = new lacework.AlertChannelSlack("business_unit_critical", {});
```
{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import pulumi_lacework as lacework

primary_critical = lacework.AlertChannelSlack("primary_critical")
business_unit_critical = lacework.AlertChannelSlack("business_unit_critical")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Lacework = Pulumi.Lacework;

return await Deployment.RunAsync(() =>
{
    var primaryCritical = new Lacework.AlertChannelSlack("primary_critical");

    var businessUnitCritical = new Lacework.AlertChannelSlack("business_unit_critical");

});

```
{{% /choosable %}}
{{% choosable language go %}}
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/lacework/v2/lacework"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lacework.NewAlertChannelSlack(ctx, "primary_critical", nil)
		if err != nil {
			return err
		}
		_, err = lacework.NewAlertChannelSlack(ctx, "business_unit_critical", nil)
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
  primaryCritical:
    type: lacework:AlertChannelSlack
    name: primary_critical
  businessUnitCritical:
    type: lacework:AlertChannelSlack
    name: business_unit_critical
```
{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.lacework.AlertChannelSlack;
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
        var primaryCritical = new AlertChannelSlack("primaryCritical");

        var businessUnitCritical = new AlertChannelSlack("businessUnitCritical");

    }
}
```
{{% /choosable %}}
{{< /chooser >}}

For a module, use its `providers` meta-argument to specify which provider configurations should be mapped to which
local provider names inside the module:

!> **Warning:** To manage multiple accounts, your user should have the Organization Administrator Role.
## Organization Level Access

Organization administrators can access organization level data sets by setting the `organization` argument to `true`.
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    lacework:organization:
        value: true

```

!> **Warning:** When accessing organization level data sets, the `subaccount` argument is ignored.

Using this type of configuration is intended for managing resources such as alerts, resource groups, team members,
cloud accounts, and more, at the organization level.
## Migrating existing resources to the Organization level

When attempting to migrate an existing resource from one of your Lacework accounts to the organization level,
you need to delete the resource, update the Lacework provider to access the organization level data set, and
run `pulumi up` to create a new resource at the organization level.
# Configuration Reference

The following configuration inputs are supported in the provider configuration:

* `profile` - (Optional) This is the Lacework profile name to use, profiles are configured
  at `$HOME/.lacework.toml` via the [Lacework CLI](https://docs.lacework.com/cli). It can also
  be sourced from the `LW_PROFILE` environment variable.

* `account` - (Optional) This is the Lacework account subdomain of URL (i.e. `<ACCOUNT>`
  .lacework.net). It must be provided, but it can also be sourced from the `LW_ACCOUNT`
  environment variable, or via the configuration file if `profile` is specified.

* `apiKey` - (Optional) This is a Lacework API access key. It must be provided when an `apiToken`
  is not used. It can also be sourced from the `LW_API_KEY` environment variable, or via the
  configuration file if `profile` is specified.

* `apiSecret` - (Optional) This is a Lacework API access secret. It must be provided when an
  `apiToken` is not used. It can also be sourced from the `LW_API_SECRET` environment variable,
  or via the configuration file if `profile` is specified.

* `apiToken` - (Optional) This is a Lacework API access token. It must be provided when neither
  the `apiKey` nor the `apiSecret` are used. It can also be sourced from the `LW_API_TOKEN`
  environment variable. Note that all API access tokens from the Lacework platform are short-lived
  which means that once the token expires, a new one needs to be generated and configured.

* `subaccount` - (Optional) The sub-account name inside your organization (for organization
  administrators only). It can also be sourced from the `LW_SUBACCOUNT` environment variable,
  or via the configuration file if `profile` is specified.

* `organization` - (Optional) Set this argument to `true` to access organization level data
  sets (for organization administrators only). It can also be sourced from the `LW_ORGANIZATION`
  environment variable.

> **Note:** For more information about creating a set of API access keys, see [Generate API Access Keys and Tokens](https://docs.lacework.com/console/generate-api-access-keys-and-tokens).