---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/sysdiglabs/sysdig/1.56.2/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Sysdig Provider
meta_desc: Provides an overview on how to configure the Pulumi Sysdig provider.
layout: package
---

## Generate Provider

The Sysdig provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider sysdiglabs/sysdig
```
## Overview

> ⚠️ content deprecated, use <https://docs.sysdig.com/en/docs/developer-tools/pulumi-provider/>
## Introduction
### What is Pulumi

Pulumi is a tool for building, changing, and versioning infrastructure safely and efficiently.
Pulumi can manage existing and popular service providers as well as custom in-house solutions.

Configuration files describe to Pulumi the components needed to run a single application or
your entire datacenter. Pulumi generates an execution plan describing what it will do to reach the
desired state, and then executes it to build the described infrastructure or configuration.

As the configuration changes, Pulumi is able to determine what changed and create incremental execution
plans which can be applied.
### Pulumi Provider for Sysdig

The Pulumi Provider for Sysdig allows you to manage your configuration in Sysdig Secure
and Sysdig Monitor as code, allowing you to synchronize your declarative configuration with
the configuration at the Platform.

You can instrument several use cases like:
- Backup/restore
- Disaster recovery
- Configuration version management
## Installation

To use the provider, first you need to install Pulumi, which is the main executable that
interacts with the provider.

Download the Pulumi executable for your OS/Architecture from
here: <https://www.pulumi.io/downloads.html>
### Pulumi v0.13+

As of Pulumi 0.13, the new block `requiredProviders` was added,
making it easier to use community providers, since they are automatically
downloaded from the Pulumi Registry.

You can tell Pulumi to download and use `sysdiglabs/sysdig` as the `sysdig`
provider by defining this block in one of your .tf files.
### Pulumi v0.12

In older Pulumi versions, you need to download the
latest version of the Pulumi Provider for Sysdig
for your OS/Architecture, extract it and move the executable under `$HOME/.pulumi.d/plugins` (you need to create
this directory if it does not exist yet) as this link suggests:
<https://www.pulumi.io/docs/configuration/providers.html#third-party-plugins> .
## Usage example

Pulumi will use the Sysdig provider when you specify a resource
or function with a name starting with `sysdig_*` (i.e.: `sysdig.User`)

But in order to actually create valid requests to the API and create/update/remove those resources,
you need to specify a correct API token for the product.

You can do so in 2 ways:
1. Using environment variables
2. Using a tfvars file.
### Configure the provider: Using env vars

You can configure the following environment variables to specify the API token:
- `SYSDIG_SECURE_API_TOKEN`
- `SYSDIG_MONITOR_API_TOKEN`

For example:

```sh
$ export SYSDIG_SECURE_API_TOKEN=323232323-3232-3232-32323232
$ export SYSDIG_MONITOR_API_TOKEN=343434343-3434-3434-34343434
```

Once you execute Pulumi and apply the manifests, that env vars will be used to configure
the provider and create API calls.
### Configure the provider: Using a tfvars file

To use a tfvars file
you need to first create it, and specify the API tokens as variables, for example:

```
# File: pulumi.tfvars

secure_token = "323232323-3232-3232-32323232"
monitor_token = "343434343-3434-3434-34343434"
```

Then, you can reference it in the provider configuration block:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    sysdig:sysdigMonitorApiToken:
        value: 'TODO: var.monitor_token'
    sysdig:sysdigSecureApiToken:
        value: 'TODO: var.secure_token'

```
### Creating resources with Pulumi

This is an example to create a pair of rules able to detect SSH connections and
shells spawned in containers.

Start by defining a couple of rules in the `rules.tf` file. One rule will detect inbound and outbound connections
made to the port 22, and the other will detect a shell process being spawned.

For more information about the configuration blocks, see: <https://www.pulumi.io/docs/configuration/syntax.html>

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as sysdig from "@pulumi/sysdig";

const disallowedSshConnection = new sysdig.SecureRuleNetwork("disallowed_ssh_connection", {
    name: "Disallowed SSH Connection detected",
    description: "Detect any new ssh connection to a host",
    tags: ["network"],
    blockInbound: true,
    blockOutbound: true,
    tcps: [{
        matching: true,
        ports: [22],
    }],
});
const terminalShell = new sysdig.SecureRuleProcess("terminal_shell", {
    name: "Terminal shell detected",
    description: "A shell was used as the entrypoint/exec point",
    tags: ["shell"],
    processes: [
        "ash",
        "bash",
        "csh",
        "ksh",
        "sh",
        "tcsh",
        "zsh",
        "dash",
    ],
});
```
{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import pulumi_sysdig as sysdig

disallowed_ssh_connection = sysdig.SecureRuleNetwork("disallowed_ssh_connection",
    name="Disallowed SSH Connection detected",
    description="Detect any new ssh connection to a host",
    tags=["network"],
    block_inbound=True,
    block_outbound=True,
    tcps=[{
        "matching": True,
        "ports": [22],
    }])
terminal_shell = sysdig.SecureRuleProcess("terminal_shell",
    name="Terminal shell detected",
    description="A shell was used as the entrypoint/exec point",
    tags=["shell"],
    processes=[
        "ash",
        "bash",
        "csh",
        "ksh",
        "sh",
        "tcsh",
        "zsh",
        "dash",
    ])
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Sysdig = Pulumi.Sysdig;

return await Deployment.RunAsync(() =>
{
    var disallowedSshConnection = new Sysdig.SecureRuleNetwork("disallowed_ssh_connection", new()
    {
        Name = "Disallowed SSH Connection detected",
        Description = "Detect any new ssh connection to a host",
        Tags = new[]
        {
            "network",
        },
        BlockInbound = true,
        BlockOutbound = true,
        Tcps = new[]
        {
            new Sysdig.Inputs.SecureRuleNetworkTcpArgs
            {
                Matching = true,
                Ports = new[]
                {
                    22,
                },
            },
        },
    });

    var terminalShell = new Sysdig.SecureRuleProcess("terminal_shell", new()
    {
        Name = "Terminal shell detected",
        Description = "A shell was used as the entrypoint/exec point",
        Tags = new[]
        {
            "shell",
        },
        Processes = new[]
        {
            "ash",
            "bash",
            "csh",
            "ksh",
            "sh",
            "tcsh",
            "zsh",
            "dash",
        },
    });

});

```
{{% /choosable %}}
{{% choosable language go %}}
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/sysdig/sysdig"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sysdig.NewSecureRuleNetwork(ctx, "disallowed_ssh_connection", &sysdig.SecureRuleNetworkArgs{
			Name:        pulumi.String("Disallowed SSH Connection detected"),
			Description: pulumi.String("Detect any new ssh connection to a host"),
			Tags: pulumi.StringArray{
				pulumi.String("network"),
			},
			BlockInbound:  pulumi.Bool(true),
			BlockOutbound: pulumi.Bool(true),
			Tcps: sysdig.SecureRuleNetworkTcpArray{
				&sysdig.SecureRuleNetworkTcpArgs{
					Matching: pulumi.Bool(true),
					Ports: pulumi.Float64Array{
						pulumi.Float64(22),
					},
				},
			},
		})
		if err != nil {
			return err
		}
		_, err = sysdig.NewSecureRuleProcess(ctx, "terminal_shell", &sysdig.SecureRuleProcessArgs{
			Name:        pulumi.String("Terminal shell detected"),
			Description: pulumi.String("A shell was used as the entrypoint/exec point"),
			Tags: pulumi.StringArray{
				pulumi.String("shell"),
			},
			Processes: pulumi.StringArray{
				pulumi.String("ash"),
				pulumi.String("bash"),
				pulumi.String("csh"),
				pulumi.String("ksh"),
				pulumi.String("sh"),
				pulumi.String("tcsh"),
				pulumi.String("zsh"),
				pulumi.String("dash"),
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
  disallowedSshConnection:
    type: sysdig:SecureRuleNetwork
    name: disallowed_ssh_connection
    properties:
      name: Disallowed SSH Connection detected
      description: Detect any new ssh connection to a host
      tags:
        - network
      blockInbound: true
      blockOutbound: true
      tcps:
        - matching: true
          ports:
            - 22
  terminalShell:
    type: sysdig:SecureRuleProcess
    name: terminal_shell
    properties:
      name: Terminal shell detected
      description: A shell was used as the entrypoint/exec point
      tags:
        - shell
      processes:
        - ash
        - bash
        - csh
        - ksh
        - sh
        - tcsh
        - zsh
        - dash
```
{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.sysdig.SecureRuleNetwork;
import com.pulumi.sysdig.SecureRuleNetworkArgs;
import com.pulumi.sysdig.inputs.SecureRuleNetworkTcpArgs;
import com.pulumi.sysdig.SecureRuleProcess;
import com.pulumi.sysdig.SecureRuleProcessArgs;
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
        var disallowedSshConnection = new SecureRuleNetwork("disallowedSshConnection", SecureRuleNetworkArgs.builder()
            .name("Disallowed SSH Connection detected")
            .description("Detect any new ssh connection to a host")
            .tags("network")
            .blockInbound(true)
            .blockOutbound(true)
            .tcps(SecureRuleNetworkTcpArgs.builder()
                .matching(true)
                .ports(22)
                .build())
            .build());

        var terminalShell = new SecureRuleProcess("terminalShell", SecureRuleProcessArgs.builder()
            .name("Terminal shell detected")
            .description("A shell was used as the entrypoint/exec point")
            .tags("shell")
            .processes(
                "ash",
                "bash",
                "csh",
                "ksh",
                "sh",
                "tcsh",
                "zsh",
                "dash")
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}

Now create a policy in a file called `policy.tf` to define how these rules
are applied. The policy will stop the affected container and trigger a capture for
further troubleshooting.

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as sysdig from "@pulumi/sysdig";

const terminalShellOrSshInContainer = new sysdig.SecurePolicy("terminal_shell_or_ssh_in_container", {
    name: "Terminal shell or SSH detected in container",
    description: "Detects a terminal shell or a ssh spawned in a container",
    enabled: true,
    severity: 0,
    scope: "container.id != \"\"",
    ruleNames: [
        disallowedSshConnection.name,
        terminalShell.name,
    ],
    actions: [{
        container: "stop",
        captures: [{
            secondsBeforeEvent: 5,
            secondsAfterEvent: 10,
        }],
    }],
});
```
{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import pulumi_sysdig as sysdig

terminal_shell_or_ssh_in_container = sysdig.SecurePolicy("terminal_shell_or_ssh_in_container",
    name="Terminal shell or SSH detected in container",
    description="Detects a terminal shell or a ssh spawned in a container",
    enabled=True,
    severity=0,
    scope="container.id != \"\"",
    rule_names=[
        disallowed_ssh_connection["name"],
        terminal_shell["name"],
    ],
    actions=[{
        "container": "stop",
        "captures": [{
            "seconds_before_event": 5,
            "seconds_after_event": 10,
        }],
    }])
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Sysdig = Pulumi.Sysdig;

return await Deployment.RunAsync(() =>
{
    var terminalShellOrSshInContainer = new Sysdig.SecurePolicy("terminal_shell_or_ssh_in_container", new()
    {
        Name = "Terminal shell or SSH detected in container",
        Description = "Detects a terminal shell or a ssh spawned in a container",
        Enabled = true,
        Severity = 0,
        Scope = "container.id != \"\"",
        RuleNames = new[]
        {
            disallowedSshConnection.Name,
            terminalShell.Name,
        },
        Actions = new[]
        {
            new Sysdig.Inputs.SecurePolicyActionArgs
            {
                Container = "stop",
                Captures = new[]
                {
                    new Sysdig.Inputs.SecurePolicyActionCaptureArgs
                    {
                        SecondsBeforeEvent = 5,
                        SecondsAfterEvent = 10,
                    },
                },
            },
        },
    });

});

```
{{% /choosable %}}
{{% choosable language go %}}
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/sysdig/sysdig"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sysdig.NewSecurePolicy(ctx, "terminal_shell_or_ssh_in_container", &sysdig.SecurePolicyArgs{
			Name:        pulumi.String("Terminal shell or SSH detected in container"),
			Description: pulumi.String("Detects a terminal shell or a ssh spawned in a container"),
			Enabled:     pulumi.Bool(true),
			Severity:    pulumi.Float64(0),
			Scope:       pulumi.String("container.id != \"\""),
			RuleNames: pulumi.StringArray{
				disallowedSshConnection.Name,
				terminalShell.Name,
			},
			Actions: sysdig.SecurePolicyActionArray{
				&sysdig.SecurePolicyActionArgs{
					Container: pulumi.String("stop"),
					Captures: sysdig.SecurePolicyActionCaptureArray{
						&sysdig.SecurePolicyActionCaptureArgs{
							SecondsBeforeEvent: pulumi.Float64(5),
							SecondsAfterEvent:  pulumi.Float64(10),
						},
					},
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
resources:
  terminalShellOrSshInContainer:
    type: sysdig:SecurePolicy
    name: terminal_shell_or_ssh_in_container
    properties:
      name: Terminal shell or SSH detected in container
      description: Detects a terminal shell or a ssh spawned in a container
      enabled: true
      severity: 0 # HIGH
      scope: container.id != ""
      ruleNames:
        - ${disallowedSshConnection.name}
        - ${terminalShell.name}
      actions:
        - container: stop
          captures:
            - secondsBeforeEvent: 5
              secondsAfterEvent: 10
```
{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.sysdig.SecurePolicy;
import com.pulumi.sysdig.SecurePolicyArgs;
import com.pulumi.sysdig.inputs.SecurePolicyActionArgs;
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
        var terminalShellOrSshInContainer = new SecurePolicy("terminalShellOrSshInContainer", SecurePolicyArgs.builder()
            .name("Terminal shell or SSH detected in container")
            .description("Detects a terminal shell or a ssh spawned in a container")
            .enabled(true)
            .severity(0)
            .scope("container.id != \"\"")
            .ruleNames(
                disallowedSshConnection.name(),
                terminalShell.name())
            .actions(SecurePolicyActionArgs.builder()
                .container("stop")
                .captures(SecurePolicyActionCaptureArgs.builder()
                    .secondsBeforeEvent(5)
                    .secondsAfterEvent(10)
                    .build())
                .build())
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}

With the given `scope`, the policy will only be applied to processes being executed inside containers:

```
scope = "container.id != \"\""
```

Using `pulumi up` the resources are applied in the backend:

!pulumi up creates the resources

Pulumi tells us that is going to create 3 resources, which matches what we defined in `rules.tf` and `policy.tf`.

!Pulumi application completes successfully

After applying the plan, Pulumi reports that the 3 resources have been successfully created. The policy uses the
rules created before, that’s why it’s the last one being created.

The resources have been created, this is how they look in Sysdig Secure:

!Pulumi rules created in Sysdig Secure

!Pulumi policy created in Sysdig Secure

But now the problem is that, if this policy triggers there's no alert notice unless notification channels are defined.
Creating two notification channels, one for the email and another one for slack in a file called `notification.tf`,
will alert us when the policy is triggered:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as sysdig from "@pulumi/sysdig";

const devops_email = new sysdig.SecureNotificationChannelEmail("devops-email", {
    name: "DevOps e-mail",
    enabled: true,
    recipients: "devops@example.com",
    notifyWhenOk: false,
    notifyWhenResolved: false,
});
const devops_slack = new sysdig.SecureNotificationChannelSlack("devops-slack", {
    name: "DevOps Slack",
    enabled: true,
    url: "https://hooks.slack.com/services/xxxxxx/xxxxxxx/xxxxxxxxxxx",
    channel: "#devops",
    notifyWhenOk: false,
    notifyWhenResolved: false,
});
```
{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import pulumi_sysdig as sysdig

devops_email = sysdig.SecureNotificationChannelEmail("devops-email",
    name="DevOps e-mail",
    enabled=True,
    recipients="devops@example.com",
    notify_when_ok=False,
    notify_when_resolved=False)
devops_slack = sysdig.SecureNotificationChannelSlack("devops-slack",
    name="DevOps Slack",
    enabled=True,
    url="https://hooks.slack.com/services/xxxxxx/xxxxxxx/xxxxxxxxxxx",
    channel="#devops",
    notify_when_ok=False,
    notify_when_resolved=False)
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Sysdig = Pulumi.Sysdig;

return await Deployment.RunAsync(() =>
{
    var devops_email = new Sysdig.SecureNotificationChannelEmail("devops-email", new()
    {
        Name = "DevOps e-mail",
        Enabled = true,
        Recipients = "devops@example.com",
        NotifyWhenOk = false,
        NotifyWhenResolved = false,
    });

    var devops_slack = new Sysdig.SecureNotificationChannelSlack("devops-slack", new()
    {
        Name = "DevOps Slack",
        Enabled = true,
        Url = "https://hooks.slack.com/services/xxxxxx/xxxxxxx/xxxxxxxxxxx",
        Channel = "#devops",
        NotifyWhenOk = false,
        NotifyWhenResolved = false,
    });

});

```
{{% /choosable %}}
{{% choosable language go %}}
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/sysdig/sysdig"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sysdig.NewSecureNotificationChannelEmail(ctx, "devops-email", &sysdig.SecureNotificationChannelEmailArgs{
			Name:               pulumi.String("DevOps e-mail"),
			Enabled:            pulumi.Bool(true),
			Recipients:         pulumi.StringArray("devops@example.com"),
			NotifyWhenOk:       pulumi.Bool(false),
			NotifyWhenResolved: pulumi.Bool(false),
		})
		if err != nil {
			return err
		}
		_, err = sysdig.NewSecureNotificationChannelSlack(ctx, "devops-slack", &sysdig.SecureNotificationChannelSlackArgs{
			Name:               pulumi.String("DevOps Slack"),
			Enabled:            pulumi.Bool(true),
			Url:                pulumi.String("https://hooks.slack.com/services/xxxxxx/xxxxxxx/xxxxxxxxxxx"),
			Channel:            pulumi.String("#devops"),
			NotifyWhenOk:       pulumi.Bool(false),
			NotifyWhenResolved: pulumi.Bool(false),
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
  devops-email:
    type: sysdig:SecureNotificationChannelEmail
    properties:
      name: DevOps e-mail
      enabled: true
      recipients: devops@example.com
      notifyWhenOk: false
      notifyWhenResolved: false
  devops-slack:
    type: sysdig:SecureNotificationChannelSlack
    properties:
      name: DevOps Slack
      enabled: true
      url: https://hooks.slack.com/services/xxxxxx/xxxxxxx/xxxxxxxxxxx
      channel: '#devops'
      notifyWhenOk: false
      notifyWhenResolved: false
```
{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.sysdig.SecureNotificationChannelEmail;
import com.pulumi.sysdig.SecureNotificationChannelEmailArgs;
import com.pulumi.sysdig.SecureNotificationChannelSlack;
import com.pulumi.sysdig.SecureNotificationChannelSlackArgs;
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
        var devops_email = new SecureNotificationChannelEmail("devops-email", SecureNotificationChannelEmailArgs.builder()
            .name("DevOps e-mail")
            .enabled(true)
            .recipients("devops@example.com")
            .notifyWhenOk(false)
            .notifyWhenResolved(false)
            .build());

        var devops_slack = new SecureNotificationChannelSlack("devops-slack", SecureNotificationChannelSlackArgs.builder()
            .name("DevOps Slack")
            .enabled(true)
            .url("https://hooks.slack.com/services/xxxxxx/xxxxxxx/xxxxxxxxxxx")
            .channel("#devops")
            .notifyWhenOk(false)
            .notifyWhenResolved(false)
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}

Bind them to the policy, modifying the file `policy.tf`; note the `notificationChannels` property:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as sysdig from "@pulumi/sysdig";

const terminalShellOrSshInContainer = new sysdig.SecurePolicy("terminal_shell_or_ssh_in_container", {
    name: "Terminal shell or SSH detected in container",
    description: "Detects a terminal shell or a ssh spawned in a container",
    enabled: true,
    severity: 0,
    scope: "container.id != \"\"",
    ruleNames: [
        disallowedSshConnection.name,
        terminalShell.name,
    ],
    actions: [{
        container: "stop",
        captures: [{
            secondsBeforeEvent: 5,
            secondsAfterEvent: 10,
        }],
    }],
    notificationChannels: [
        devops_email.id,
        devops_slack.id,
    ],
});
```
{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import pulumi_sysdig as sysdig

terminal_shell_or_ssh_in_container = sysdig.SecurePolicy("terminal_shell_or_ssh_in_container",
    name="Terminal shell or SSH detected in container",
    description="Detects a terminal shell or a ssh spawned in a container",
    enabled=True,
    severity=0,
    scope="container.id != \"\"",
    rule_names=[
        disallowed_ssh_connection["name"],
        terminal_shell["name"],
    ],
    actions=[{
        "container": "stop",
        "captures": [{
            "seconds_before_event": 5,
            "seconds_after_event": 10,
        }],
    }],
    notification_channels=[
        devops_email["id"],
        devops_slack["id"],
    ])
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Sysdig = Pulumi.Sysdig;

return await Deployment.RunAsync(() =>
{
    var terminalShellOrSshInContainer = new Sysdig.SecurePolicy("terminal_shell_or_ssh_in_container", new()
    {
        Name = "Terminal shell or SSH detected in container",
        Description = "Detects a terminal shell or a ssh spawned in a container",
        Enabled = true,
        Severity = 0,
        Scope = "container.id != \"\"",
        RuleNames = new[]
        {
            disallowedSshConnection.Name,
            terminalShell.Name,
        },
        Actions = new[]
        {
            new Sysdig.Inputs.SecurePolicyActionArgs
            {
                Container = "stop",
                Captures = new[]
                {
                    new Sysdig.Inputs.SecurePolicyActionCaptureArgs
                    {
                        SecondsBeforeEvent = 5,
                        SecondsAfterEvent = 10,
                    },
                },
            },
        },
        NotificationChannels = new[]
        {
            devops_email.Id,
            devops_slack.Id,
        },
    });

});

```
{{% /choosable %}}
{{% choosable language go %}}
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/sysdig/sysdig"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sysdig.NewSecurePolicy(ctx, "terminal_shell_or_ssh_in_container", &sysdig.SecurePolicyArgs{
			Name:        pulumi.String("Terminal shell or SSH detected in container"),
			Description: pulumi.String("Detects a terminal shell or a ssh spawned in a container"),
			Enabled:     pulumi.Bool(true),
			Severity:    pulumi.Float64(0),
			Scope:       pulumi.String("container.id != \"\""),
			RuleNames: pulumi.StringArray{
				disallowedSshConnection.Name,
				terminalShell.Name,
			},
			Actions: sysdig.SecurePolicyActionArray{
				&sysdig.SecurePolicyActionArgs{
					Container: pulumi.String("stop"),
					Captures: sysdig.SecurePolicyActionCaptureArray{
						&sysdig.SecurePolicyActionCaptureArgs{
							SecondsBeforeEvent: pulumi.Float64(5),
							SecondsAfterEvent:  pulumi.Float64(10),
						},
					},
				},
			},
			NotificationChannels: pulumi.Float64Array{
				devops_email.Id,
				devops_slack.Id,
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
  terminalShellOrSshInContainer:
    type: sysdig:SecurePolicy
    name: terminal_shell_or_ssh_in_container
    properties:
      name: Terminal shell or SSH detected in container
      description: Detects a terminal shell or a ssh spawned in a container
      enabled: true
      severity: 0 # HIGH
      scope: container.id != ""
      ruleNames:
        - ${disallowedSshConnection.name}
        - ${terminalShell.name}
      actions:
        - container: stop
          captures:
            - secondsBeforeEvent: 5
              secondsAfterEvent: 10
      notificationChannels:
        - ${["devops-email"].id}
        - ${["devops-slack"].id}
```
{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.sysdig.SecurePolicy;
import com.pulumi.sysdig.SecurePolicyArgs;
import com.pulumi.sysdig.inputs.SecurePolicyActionArgs;
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
        var terminalShellOrSshInContainer = new SecurePolicy("terminalShellOrSshInContainer", SecurePolicyArgs.builder()
            .name("Terminal shell or SSH detected in container")
            .description("Detects a terminal shell or a ssh spawned in a container")
            .enabled(true)
            .severity(0)
            .scope("container.id != \"\"")
            .ruleNames(
                disallowedSshConnection.name(),
                terminalShell.name())
            .actions(SecurePolicyActionArgs.builder()
                .container("stop")
                .captures(SecurePolicyActionCaptureArgs.builder()
                    .secondsBeforeEvent(5)
                    .secondsAfterEvent(10)
                    .build())
                .build())
            .notificationChannels(
                devops_email.id(),
                devops_slack.id())
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}

Finally, doing a `pulumi up`, it will inform that it will create 2 new resources and modify the existing policy:

!pulumi up updates the resources

After inputting **yes**, Pulumi will create the notification channels and bind them to the policy,
ensuring that the state in Monitor and Secure matches our state defined in the code.

This is how the resources appear on the Sysdig Secure UI:

!pulumi up creates new notification channels

!Pulumi updates the policy resource

Now, if someone tries to update it manually, by re-applying the policies, Pulumi will
restore the desired status from the `.tf` manifests.
## Full Pulumi resources documentation

Check all the available resources and functions for the Pulumi Provider for Sysdig here:

Pulumi provider for Sysdig Functions

---
!Sysdig logo