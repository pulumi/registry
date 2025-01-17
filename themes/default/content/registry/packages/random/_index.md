---
# WARNING: this file was fetched from https://d3uflda8hlgvg8.cloudfront.net/docs/registry.opentofu.org/hashicorp/random/3.6.3/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Random Provider
meta_desc: Provides an overview on how to configure the Pulumi Random provider.
layout: package
---

## Generate Provider

The Random provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider hashicorp/random
```
## Overview

The "random" provider allows the use of randomness within Pulumi
configurations. This is a *logical provider*, which means that it works
entirely within Pulumi's logic, and doesn't interact with any other
services.

Unconstrained randomness within a Pulumi configuration would not be very
useful, since Pulumi's goal is to converge on a fixed configuration by
applying a diff. Because of this, the "random" provider provides an idea of
*managed randomness*: it provides resources that generate random values during
their creation and then hold those values steady until the inputs are changed.

Even with these resources, it is advisable to keep the use of randomness within
Pulumi configuration to a minimum, and retain it for special cases only;
Pulumi works best when the configuration is well-defined, since its behavior
can then be more readily predicted.

Unless otherwise stated within the documentation of a specific resource, this
provider's results are **not** sufficiently random for cryptographic use.

For more information on the specific resources available, see the links in the
navigation bar. Read on for information on the general patterns that apply
to this provider's resources.
## Resource "Keepers"

As noted above, the random resources generate randomness only when they are
created; the results produced are stored in the Pulumi state and re-used
until the inputs change, prompting the resource to be recreated.

The resources all provide a map argument called `keepers` that can be populated
with arbitrary key/value pairs that should be selected such that they remain
the same until new random values are desired.

For example:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as random from "@pulumi/random";

const server = new random.Id("server", {
    keepers: {
        ami_id: amiId,
    },
    byteLength: 8,
});
const serverInstance = new aws.index.Instance("server", {
    tags: {
        name: `web-server ${server.hex}`,
    },
    ami: server.keepers?.amiId,
});
```
{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import pulumi_aws as aws
import pulumi_random as random

server = random.Id("server",
    keepers={
        "ami_id": ami_id,
    },
    byte_length=8)
server_instance = aws.index.Instance("server",
    tags={
        name: fweb-server {server.hex},
    },
    ami=server.keepers.ami_id)
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;
using Random = Pulumi.Random;

return await Deployment.RunAsync(() =>
{
    var server = new Random.Id("server", new()
    {
        Keepers =
        {
            { "ami_id", amiId },
        },
        ByteLength = 8,
    });

    var serverInstance = new Aws.Index.Instance("server", new()
    {
        Tags =
        {
            { "name", $"web-server {server.Hex}" },
        },
        Ami = server.Keepers?.AmiId,
    });

});

```
{{% /choosable %}}
{{% choosable language go %}}
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/go/aws"
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/random/v3/random"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		server, err := random.NewId(ctx, "server", &random.IdArgs{
			Keepers: pulumi.StringMap{
				"ami_id": pulumi.Any(amiId),
			},
			ByteLength: pulumi.Float64(8),
		})
		if err != nil {
			return err
		}
		_, err = aws.NewInstance(ctx, "server", &aws.InstanceArgs{
			Tags: map[string]interface{}{
				"name": pulumi.Sprintf("web-server %v", server.Hex),
			},
			Ami: server.Keepers.AmiId,
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
  server:
    type: random:Id
    properties:
      keepers:
        ami_id: ${amiId}
      byteLength: 8
  serverInstance:
    type: aws:Instance
    name: server
    properties:
      tags:
        name: web-server ${server.hex}
      ami: ${server.keepers.amiId}
```
{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.random.Id;
import com.pulumi.random.IdArgs;
import com.pulumi.aws.Instance;
import com.pulumi.aws.InstanceArgs;
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
        var server = new Id("server", IdArgs.builder()
            .keepers(Map.of("ami_id", amiId))
            .byteLength(8)
            .build());

        var serverInstance = new Instance("serverInstance", InstanceArgs.builder()
            .tags(%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference))
            .ami(server.keepers().amiId())
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}

Resource "keepers" are optional. The other arguments to each resource must
*also* remain constant in order to retain a random result.

`keepers` are *not* treated as sensitive attributes; a value used for `keepers` will be displayed in Pulumi UI output as plaintext.

To force a random result to be replaced, the `taint` command can be used to
produce a new result on the next run.