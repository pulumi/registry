---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-slack/v0.4.14/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/pulumi/pulumi-slack/blob/v0.4.14/docs/_index.md
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Slack Provider
meta_desc: Provides an overview on how to configure the Pulumi Slack provider.
layout: package
---

## Installation

The Slack provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/slack`](https://www.npmjs.com/package/@pulumi/slack)
* Python: [`pulumi-slack`](https://pypi.org/project/pulumi-slack/)
* Go: [`github.com/pulumi/pulumi-slack/sdk/go/slack`](https://github.com/pulumi/pulumi-slack)
* .NET: [`Pulumi.Slack`](https://www.nuget.org/packages/Pulumi.Slack)
* Java: [`com.pulumi/slack`](https://central.sonatype.com/artifact/com.pulumi/slack)

## Overview

The Slack provider is used to interact with Slack resources supported by Slack.
The provider needs to be configured with a valid token before it can be used.

Use the navigation to the left to read about the available resources.
## Example Usage



{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    slack:token:
        value: 'TODO: var.slack_token'

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as slack from "@pulumi/slack";

const testUser00 = slack.getUser({
    name: "contact_test-user-ter",
});
// Create a User Group
const myGroup = new slack.Usergroup("my_group", {
    name: "TestGroup",
    handle: "test",
    description: "Test user group",
    users: [testUser00.then(testUser00 => testUser00.id)],
});
// Create a Slack channel
const test = new slack.Conversation("test", {
    name: "my-channel",
    topic: "The topic for my channel",
    permanentMembers: myGroup.users,
    isPrivate: true,
});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    slack:token:
        value: 'TODO: var.slack_token'

```
```python
import pulumi
import pulumi_slack as slack

test_user00 = slack.get_user(name="contact_test-user-ter")
# Create a User Group
my_group = slack.Usergroup("my_group",
    name="TestGroup",
    handle="test",
    description="Test user group",
    users=[test_user00.id])
# Create a Slack channel
test = slack.Conversation("test",
    name="my-channel",
    topic="The topic for my channel",
    permanent_members=my_group.users,
    is_private=True)
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    slack:token:
        value: 'TODO: var.slack_token'

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Slack = Pulumi.Slack;

return await Deployment.RunAsync(() =>
{
    var testUser00 = Slack.GetUser.Invoke(new()
    {
        Name = "contact_test-user-ter",
    });

    // Create a User Group
    var myGroup = new Slack.Usergroup("my_group", new()
    {
        Name = "TestGroup",
        Handle = "test",
        Description = "Test user group",
        Users = new[]
        {
            testUser00.Apply(getUserResult => getUserResult.Id),
        },
    });

    // Create a Slack channel
    var test = new Slack.Conversation("test", new()
    {
        Name = "my-channel",
        Topic = "The topic for my channel",
        PermanentMembers = myGroup.Users,
        IsPrivate = true,
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
    slack:token:
        value: 'TODO: var.slack_token'

```
```go
package main

import (
	"github.com/pulumi/pulumi-slack/sdk/go/slack"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		testUser00, err := slack.GetUser(ctx, &slack.GetUserArgs{
			Name: pulumi.StringRef("contact_test-user-ter"),
		}, nil)
		if err != nil {
			return err
		}
		// Create a User Group
		myGroup, err := slack.NewUsergroup(ctx, "my_group", &slack.UsergroupArgs{
			Name:        pulumi.String("TestGroup"),
			Handle:      pulumi.String("test"),
			Description: pulumi.String("Test user group"),
			Users: pulumi.StringArray{
				pulumi.String(testUser00.Id),
			},
		})
		if err != nil {
			return err
		}
		// Create a Slack channel
		_, err = slack.NewConversation(ctx, "test", &slack.ConversationArgs{
			Name:             pulumi.String("my-channel"),
			Topic:            pulumi.String("The topic for my channel"),
			PermanentMembers: myGroup.Users,
			IsPrivate:        pulumi.Bool(true),
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
    slack:token:
        value: 'TODO: var.slack_token'

```
```yaml
resources:
  # Create a User Group
  myGroup:
    type: slack:Usergroup
    name: my_group
    properties:
      name: TestGroup
      handle: test
      description: Test user group
      users:
        - ${testUser00.id}
  # Create a Slack channel
  test:
    type: slack:Conversation
    properties:
      name: my-channel
      topic: The topic for my channel
      permanentMembers: ${myGroup.users}
      isPrivate: true
variables:
  testUser00:
    fn::invoke:
      function: slack:getUser
      arguments:
        name: contact_test-user-ter
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    slack:token:
        value: 'TODO: var.slack_token'

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.slack.SlackFunctions;
import com.pulumi.slack.inputs.GetUserArgs;
import com.pulumi.slack.Usergroup;
import com.pulumi.slack.UsergroupArgs;
import com.pulumi.slack.Conversation;
import com.pulumi.slack.ConversationArgs;
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
        final var testUser00 = SlackFunctions.getUser(GetUserArgs.builder()
            .name("contact_test-user-ter")
            .build());

        // Create a User Group
        var myGroup = new Usergroup("myGroup", UsergroupArgs.builder()
            .name("TestGroup")
            .handle("test")
            .description("Test user group")
            .users(testUser00.id())
            .build());

        // Create a Slack channel
        var test = new Conversation("test", ConversationArgs.builder()
            .name("my-channel")
            .topic("The topic for my channel")
            .permanentMembers(myGroup.users())
            .isPrivate(true)
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Authentication

The Slack provider requires an Slack API token. It can be provided by different
means:

- Static token
- Environment variables
### Static Token

!> **Warning:** Hard-coding credentials into any Pulumi configuration is not
recommended, and risks secret leakage should this file ever be committed to a
public version control system.

A static can be provided by adding `token` in-line in the Slack provider configuration:

Usage:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    slack:token:
        value: 'TODO: var.slack_token'

```
### Environment Variables

You can provide your token via the `SLACK_TOKEN` environment variable:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:

```

Usage:

```sh
export SLACK_TOKEN="my-token"
pulumi preview
```
## Configuration Reference

In addition to generic `provider` arguments
(e.g. `alias` and `version`), the following arguments are supported in the Slack
provider configuration:

- `token` - (Mandatory) The Slack token. It must be provided,
  but it can also be sourced from the `SLACK_TOKEN` environment variable.