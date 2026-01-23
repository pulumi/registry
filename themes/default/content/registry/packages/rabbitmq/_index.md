---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-rabbitmq/v3.4.0/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/pulumi/pulumi-rabbitmq/blob/v3.4.0/docs/_index.md
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Rabbitmq Provider
meta_desc: Provides an overview on how to configure the Pulumi Rabbitmq provider.
layout: package
---

## Installation

The Rabbitmq provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/rabbitmq`](https://www.npmjs.com/package/@pulumi/rabbitmq)
* Python: [`pulumi-rabbitmq`](https://pypi.org/project/pulumi-rabbitmq/)
* Go: [`github.com/pulumi/pulumi-rabbitmq/sdk/v3/go/rabbitmq`](https://github.com/pulumi/pulumi-rabbitmq)
* .NET: [`Pulumi.Rabbitmq`](https://www.nuget.org/packages/Pulumi.Rabbitmq)
* Java: [`com.pulumi/rabbitmq`](https://central.sonatype.com/artifact/com.pulumi/rabbitmq)

## Overview

[RabbitMQ](http://www.rabbitmq.com) is an AMQP message broker server. The
RabbitMQ provider exposes resources used to manage the configuration of
resources in a RabbitMQ server.

Use the navigation to the left to read about the available resources.
## Example Usage

The following is a minimal example:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    rabbitmq:endpoint:
        value: http://127.0.0.1
    rabbitmq:password:
        value: guest
    rabbitmq:username:
        value: guest

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as rabbitmq from "@pulumi/rabbitmq";

// Create a virtual host
const vhost1 = new rabbitmq.VHost("vhost_1", {name: "vhost_1"});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    rabbitmq:endpoint:
        value: http://127.0.0.1
    rabbitmq:password:
        value: guest
    rabbitmq:username:
        value: guest

```
```python
import pulumi
import pulumi_rabbitmq as rabbitmq

# Create a virtual host
vhost1 = rabbitmq.VHost("vhost_1", name="vhost_1")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    rabbitmq:endpoint:
        value: http://127.0.0.1
    rabbitmq:password:
        value: guest
    rabbitmq:username:
        value: guest

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using RabbitMQ = Pulumi.RabbitMQ;

return await Deployment.RunAsync(() =>
{
    // Create a virtual host
    var vhost1 = new RabbitMQ.VHost("vhost_1", new()
    {
        Name = "vhost_1",
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
    rabbitmq:endpoint:
        value: http://127.0.0.1
    rabbitmq:password:
        value: guest
    rabbitmq:username:
        value: guest

```
```go
package main

import (
	"github.com/pulumi/pulumi-rabbitmq/sdk/v3/go/rabbitmq"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a virtual host
		_, err := rabbitmq.NewVHost(ctx, "vhost_1", &rabbitmq.VHostArgs{
			Name: pulumi.String("vhost_1"),
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
    rabbitmq:endpoint:
        value: http://127.0.0.1
    rabbitmq:password:
        value: guest
    rabbitmq:username:
        value: guest

```
```yaml
resources:
  # Create a virtual host
  vhost1:
    type: rabbitmq:VHost
    name: vhost_1
    properties:
      name: vhost_1
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    rabbitmq:endpoint:
        value: http://127.0.0.1
    rabbitmq:password:
        value: guest
    rabbitmq:username:
        value: guest

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.rabbitmq.VHost;
import com.pulumi.rabbitmq.VHostArgs;
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
        // Create a virtual host
        var vhost1 = new VHost("vhost1", VHostArgs.builder()
            .name("vhost_1")
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Requirements

The RabbitMQ management plugin must be enabled on the server, to use this provider. You can
enable the plugin by doing something similar to:

```
$ sudo rabbitmq-plugins enable rabbitmq_management
```
## Configuration Reference

The following configuration inputs are supported:

* `endpoint` - (Required) The HTTP URL of the management plugin on the
  RabbitMQ server. This can also be sourced from the `RABBITMQ_ENDPOINT`
  Environment Variable. The RabbitMQ management plugin *must* be enabled in order
  to use this provider. *Note*: This is not the IP address or hostname of the
  RabbitMQ server that you would use to access RabbitMQ directly.
* `username` - (Required) Username to use to authenticate with the server.
  This can also be sourced from the `RABBITMQ_USERNAME` Environment Variable.
* `password` - (Optional) Password for the given user. This can also be sourced
  from the `RABBITMQ_PASSWORD` Environment Variable.
* `insecure` - (Optional) Trust self-signed certificates. This can also be sourced
  from the `RABBITMQ_INSECURE` Environment Variable.
* `cacertFile` - (Optional) The path to a custom CA / intermediate certificate.
  This can also be sourced from the `RABBITMQ_CACERT` Environment Variable.
* `clientcertFile` - (Optional) The path to the X.509 client certificate.
  This can also be sourced from the `RABBITMQ_CLIENTCERT` Environment Variable
* `clientkeyFile` - (Optional) The path to the private key.
  This can also be sourced from the `RABBITMQ_CLIENTKEY` Environment Variable
* `proxy` - (Optional) The URL of a proxy through which to send HTTP requests to
  the RabbitMQ server. This can also be sourced from the `RABBITMQ_PROXY`
  Environment Variable. If not set, the default `HTTP_PROXY`/`HTTPS_PROXY` will
  be used instead.