---
title: threefold
meta_desc: Provides an overview of the threefold grid Provider for Pulumi.
layout: package
---

The Threefold Resource Provider for the [threefold grid](https://threefold.io) lets you manage your infrastructure using Pulumi.

## Example

{{< chooser language "go,yaml" >}}

{{% choosable language go %}}

```go
package main

import (
    "os"

    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
    "github.com/threefoldtech/pulumi-threefold/sdk/go/threefold"
    "github.com/threefoldtech/pulumi-threefold/sdk/go/threefold/provider"
)

func main() {
    pulumi.Run(func(ctx *pulumi.Context) error {
        mnemonic := os.Getenv("MNEMONIC")

        _, err := threefold.NewProvider(ctx, "grid provider", &threefold.ProviderArgs{
            Mnemonic: pulumi.String(mnemonic),
        })

        if err != nil {
            return err
        }

        _, err = provider.NewNetwork(ctx, "grid network", &provider.NetworkArgs{
        Description: pulumi.String("example network"),
            Ip_range:    pulumi.String("10.1.0.0/16"),
            Name:        pulumi.String("example"),
            Nodes:       pulumi.Array{pulumi.Int(14)},
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

```yml
name: pulumi-threefold
runtime: yaml

resources:
  provider:
    type: pulumi:providers:threefold
    properties:
      mnemonic:

  scheduler:
    type: threefold:provider:Scheduler
    options:
      provider: ${provider}
    properties:
      farm_ids: [1]

  network:
    type: threefold:provider:Network
    options:
      provider: ${provider}
      dependsOn:
        - ${scheduler}
    properties:
      name: testing
      description: test network
      nodes:
        - ${scheduler.nodes[0]}
      ip_range: 10.1.0.0/16

outputs:
  node_deployment_id: ${network.node_deployment_id}
  nodes_ip_range: ${network.nodes_ip_range}
```

{{% /choosable %}}

{{< /chooser >}}
