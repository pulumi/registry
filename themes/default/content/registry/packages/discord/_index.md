---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/lucky3028/discord/2.2.0/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Discord Provider
meta_desc: Provides an overview on how to configure the Pulumi Discord provider.
layout: package
---

## Generate Provider

The Discord provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider lucky3028/discord
```
## Overview

This is a fork of Chaotic-Logic/pulumi-provider-discord.

The Discord provider is used to interact with the Discord API. It requires proper credentials before it can be used.

Use the navigation on the left to read more about the resources and functions.
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    discord:token:
        value: 'TODO: var.discord_token'

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as discord from "@pulumi/discord";

const logo = discord.getLocalImage({
    file: "logo.png",
});
const myServer = new discord.Server("my_server", {
    name: "My Awesome Server",
    region: "us-west",
    defaultMessageNotifications: 0,
    iconDataUri: logo.then(logo => logo.dataUri),
});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    discord:token:
        value: 'TODO: var.discord_token'

```
```python
import pulumi
import pulumi_discord as discord

logo = discord.get_local_image(file="logo.png")
my_server = discord.Server("my_server",
    name="My Awesome Server",
    region="us-west",
    default_message_notifications=0,
    icon_data_uri=logo.data_uri)
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    discord:token:
        value: 'TODO: var.discord_token'

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Discord = Pulumi.Discord;

return await Deployment.RunAsync(() =>
{
    var logo = Discord.GetLocalImage.Invoke(new()
    {
        File = "logo.png",
    });

    var myServer = new Discord.Server("my_server", new()
    {
        Name = "My Awesome Server",
        Region = "us-west",
        DefaultMessageNotifications = 0,
        IconDataUri = logo.Apply(getLocalImageResult => getLocalImageResult.DataUri),
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
    discord:token:
        value: 'TODO: var.discord_token'

```
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/discord/v2/discord"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		logo, err := discord.GetLocalImage(ctx, &discord.GetLocalImageArgs{
			File: "logo.png",
		}, nil)
		if err != nil {
			return err
		}
		_, err = discord.NewServer(ctx, "my_server", &discord.ServerArgs{
			Name:                        pulumi.String("My Awesome Server"),
			Region:                      pulumi.String("us-west"),
			DefaultMessageNotifications: pulumi.Float64(0),
			IconDataUri:                 pulumi.String(logo.DataUri),
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
    discord:token:
        value: 'TODO: var.discord_token'

```
```yaml
resources:
  myServer:
    type: discord:Server
    name: my_server
    properties:
      name: My Awesome Server
      region: us-west
      defaultMessageNotifications: 0
      iconDataUri: ${logo.dataUri}
variables:
  logo:
    fn::invoke:
      function: discord:getLocalImage
      arguments:
        file: logo.png
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    discord:token:
        value: 'TODO: var.discord_token'

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.discord.DiscordFunctions;
import com.pulumi.discord.inputs.GetLocalImageArgs;
import com.pulumi.discord.Server;
import com.pulumi.discord.ServerArgs;
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
        final var logo = DiscordFunctions.getLocalImage(GetLocalImageArgs.builder()
            .file("logo.png")
            .build());

        var myServer = new Server("myServer", ServerArgs.builder()
            .name("My Awesome Server")
            .region("us-west")
            .defaultMessageNotifications(0)
            .iconDataUri(logo.applyValue(getLocalImageResult -> getLocalImageResult.dataUri()))
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

- `clientId` (String) OAuth app client ID. Currently unused.
- `secret` (String) OAuth app secret. Currently unused.
- `token` (String) Discord API token, without the `Bot` prefix. This can be found in the Discord Developer Portal. This can also be set via the `DISCORD_TOKEN` environment variable.