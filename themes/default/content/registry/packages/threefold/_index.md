---
title: threefold
meta_desc: Provides an overview of the threefold grid Provider for Pulumi.
layout: package
---

The Threefold Resource Provider for the [threefold grid](https://threefold.io) lets you manage your infrastructure using Pulumi.

## Example

### Network resource

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
    tfProvider, err := threefold.NewProvider(ctx, "provider", &threefold.ProviderArgs{
      Mnemonic: pulumi.String(os.Getenv("MNEMONIC")),
    })
    if err != nil {
      return err
    }

    scheduler, err := provider.NewScheduler(ctx, "scheduler", &provider.SchedulerArgs{
      Farm_ids: pulumi.IntArray{
        pulumi.Int(1),
      },
    }, pulumi.Provider(tfProvider))
    if err != nil {
      return err
    }

    network, err := provider.NewNetwork(ctx, "network", &provider.NetworkArgs{
      Name:        pulumi.String("testing"),
      Description: pulumi.String("test network"),
      Nodes: pulumi.Array{
        scheduler.Nodes.ApplyT(func(nodes []int) (int, error) {
          return nodes[0], nil
        }).(pulumi.IntOutput),
      },
      Ip_range: pulumi.String("10.1.0.0/16"),
    }, pulumi.Provider(tfProvider), pulumi.DependsOn([]pulumi.Resource{
      scheduler,
    }))
    if err != nil {
      return err
    }

    ctx.Export("node_deployment_id", network.Node_deployment_id)
    ctx.Export("nodes_ip_range", network.Nodes_ip_range)
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

### Virtual machine resource

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
    tfProvider, err := threefold.NewProvider(ctx, "provider", &threefold.ProviderArgs{
      Mnemonic: pulumi.String(os.Getenv("MNEMONIC")),
    })
    if err != nil {
      return err
    }
    scheduler, err := provider.NewScheduler(ctx, "scheduler", &provider.SchedulerArgs{
      Mru: pulumi.Int(1),
      Sru: pulumi.Int(2),
      Farm_ids: pulumi.IntArray{
        pulumi.Int(1),
      },
    }, pulumi.Provider(tfProvider))
    if err != nil {
      return err
    }
    network, err := provider.NewNetwork(ctx, "network", &provider.NetworkArgs{
      Name:        pulumi.String("test"),
      Description: pulumi.String("test network"),
      Nodes: pulumi.Array{
        scheduler.Nodes.ApplyT(func(nodes []int) (int, error) {
          return nodes[0], nil
        }).(pulumi.IntOutput),
      },
      Ip_range: pulumi.String("10.1.0.0/16"),
    }, pulumi.Provider(tfProvider), pulumi.DependsOn([]pulumi.Resource{
      scheduler,
    }))
    if err != nil {
      return err
    }
    deployment, err := provider.NewDeployment(ctx, "deployment", &provider.DeploymentArgs{
      Node_id: scheduler.Nodes.ApplyT(func(nodes []int) (int, error) {
        return nodes[0], nil
      }).(pulumi.IntOutput),
      Name:         pulumi.String("deployment"),
      Network_name: pulumi.String("test"),
      Vms: provider.VMInputArray{
        &provider.VMInputArgs{
          Name:         pulumi.String("vm"),
          Flist:        pulumi.String("https://hub.grid.tf/tf-official-apps/base:latest.flist"),
          Entrypoint:   pulumi.String("/sbin/zinit init"),
          Network_name: pulumi.String("test"),
          Cpu:          pulumi.Int(2),
          Memory:       pulumi.Int(256),
          Planetary:    pulumi.Bool(true),
          Mounts: provider.MountArray{
            &provider.MountArgs{
              Disk_name:   pulumi.String("data"),
              Mount_point: pulumi.String("/app"),
            },
          },
          Env_vars: pulumi.StringMap{
            "SSH_KEY": nil,
          },
        },
      },
      Disks: provider.DiskArray{
        &provider.DiskArgs{
          Name: pulumi.String("data"),
          Size: pulumi.Int(2),
        },
      },
    }, pulumi.Provider(tfProvider), pulumi.DependsOn([]pulumi.Resource{
      network,
    }))
    if err != nil {
      return err
    }
    ctx.Export("node_deployment_id", deployment.Node_deployment_id)
    ctx.Export("planetary_ip", deployment.Vms_computed.ApplyT(func(vms_computed []provider.VMComputed) (*string, error) {
      return &vms_computed[0].Planetary_ip, nil
    }).(pulumi.StringPtrOutput))
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
    options:
      pluginDownloadURL: github://api.github.com/threefoldtech/pulumi-threefold # optional
    properties:
      mnemonic:

  scheduler:
    type: threefold:provider:Scheduler
    options:
      provider: ${provider}
    properties:
      mru: 0.25 # 256 megabytes
      sru: 2
      farm_ids: [1]

  network:
    type: threefold:provider:Network
    options:
      provider: ${provider}
      dependsOn:
        - ${scheduler}
    properties:
      name: test
      description: test network
      nodes:
        - ${scheduler.nodes[0]}
      ip_range: 10.1.0.0/16
      mycelium: true

  deployment:
    type: threefold:provider:Deployment
    options:
      provider: ${provider}
      dependsOn:
        - ${network}
    properties:
      node_id: ${scheduler.nodes[0]}
      name: deployment
      network_name: test
      vms:
        - name: vm
          flist: https://hub.grid.tf/tf-official-apps/base:latest.flist
          entrypoint: "/sbin/zinit init"
          network_name: test
          cpu: 2
          memory: 256
          planetary: true
          mycelium: true
          mounts:
            - disk_name: data
              mount_point: /app
          env_vars:
            SSH_KEY:

      disks:
        - name: data
          size: 2

outputs:
  node_deployment_id: ${deployment.node_deployment_id}
  planetary_ip: ${deployment.vms_computed[0].planetary_ip}
  mycelium_ip: ${deployment.vms_computed[0].mycelium_ip}
```

{{% /choosable %}}

{{< /chooser >}}

### Kubernetes resource

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
    tfProvider, err := threefold.NewProvider(ctx, "provider", &threefold.ProviderArgs{
      Mnemonic: pulumi.String(os.Getenv("MNEMONIC")),
    })
    if err != nil {
      return err
    }
    scheduler, err := provider.NewScheduler(ctx, "scheduler", &provider.SchedulerArgs{
      Mru: pulumi.Int(6),
      Sru: pulumi.Int(6),
      Farm_ids: pulumi.IntArray{
        pulumi.Int(1),
      },
    }, pulumi.Provider(tfProvider))
    if err != nil {
      return err
    }
    network, err := provider.NewNetwork(ctx, "network", &provider.NetworkArgs{
      Name:        pulumi.String("test"),
      Description: pulumi.String("test network"),
      Nodes: pulumi.Array{
        scheduler.Nodes.ApplyT(func(nodes []int) (int, error) {
          return nodes[0], nil
        }).(pulumi.IntOutput),
      },
      Ip_range: pulumi.String("10.1.0.0/16"),
    }, pulumi.Provider(tfProvider), pulumi.DependsOn([]pulumi.Resource{
      scheduler,
    }))
    if err != nil {
      return err
    }
    kubernetes, err := provider.NewKubernetes(ctx, "kubernetes", &provider.KubernetesArgs{
      Master: &provider.K8sNodeInputArgs{
        Name: pulumi.String("kubernetes"),
        Node: scheduler.Nodes.ApplyT(func(nodes []int) (int, error) {
          return nodes[0], nil
        }).(pulumi.IntOutput),
        Disk_size: pulumi.Int(2),
        Planetary: pulumi.Bool(true),
        Cpu:       pulumi.Int(2),
        Memory:    pulumi.Int(2048),
      },
      Workers: provider.K8sNodeInputArray{
        &provider.K8sNodeInputArgs{
          Name: pulumi.String("worker1"),
          Node: scheduler.Nodes.ApplyT(func(nodes []int) (int, error) {
            return nodes[0], nil
          }).(pulumi.IntOutput),
          Disk_size: pulumi.Int(2),
          Cpu:       pulumi.Int(2),
          Memory:    pulumi.Int(2048),
        },
        &provider.K8sNodeInputArgs{
          Name: pulumi.String("worker2"),
          Node: scheduler.Nodes.ApplyT(func(nodes []int) (int, error) {
            return nodes[0], nil
          }).(pulumi.IntOutput),
          Disk_size: pulumi.Int(2),
          Cpu:       pulumi.Int(2),
          Memory:    pulumi.Int(2048),
        },
      },
      Token:        pulumi.String("t123456789"),
      Network_name: pulumi.String("test"),
      Ssh_key:      nil,
    }, pulumi.Provider(tfProvider), pulumi.DependsOn([]pulumi.Resource{
      network,
    }))
    if err != nil {
      return err
    }
    ctx.Export("node_deployment_id", kubernetes.Node_deployment_id)
    ctx.Export("planetary_ip", kubernetes.Master_computed.ApplyT(func(master_computed provider.K8sNodeComputed) (*string, error) {
      return &master_computed.Planetary_ip, nil
    }).(pulumi.StringPtrOutput))
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
    options:
      pluginDownloadURL: github://api.github.com/threefoldtech/pulumi-threefold # optional
    properties:
      mnemonic:

  scheduler:
    type: threefold:provider:Scheduler
    options:
      provider: ${provider}
    properties:
      mru: 6
      sru: 6
      farm_ids: [1]

  network:
    type: threefold:provider:Network
    options:
      provider: ${provider}
      dependsOn:
        - ${scheduler}
    properties:
      name: test
      description: test network
      nodes:
        - ${scheduler.nodes[0]}
      ip_range: 10.1.0.0/16

  kubernetes:
    type: threefold:provider:Kubernetes
    options:
      provider: ${provider}
      dependsOn:
        - ${network}
    properties:
      master:
        name: kubernetes
        node: ${scheduler.nodes[0]}
        disk_size: 2
        planetary: true
        cpu: 2
        memory: 2048

      workers:
        - name: worker1
          node: ${scheduler.nodes[0]}
          disk_size: 2
          cpu: 2
          memory: 2048
        - name: worker2
          node: ${scheduler.nodes[0]}
          disk_size: 2
          cpu: 2
          memory: 2048

      token: t123456789
      network_name: test
      ssh_key:

outputs:
  node_deployment_id: ${kubernetes.node_deployment_id}
  planetary_ip: ${kubernetes.master_computed.planetary_ip}
```

{{% /choosable %}}

{{< /chooser >}}

### Name gateway resource

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
    tfProvider, err := threefold.NewProvider(ctx, "provider", &threefold.ProviderArgs{
      Mnemonic: pulumi.String(os.Getenv("MNEMONIC")),
    })
    if err != nil {
      return err
    }
    scheduler, err := provider.NewScheduler(ctx, "scheduler", &provider.SchedulerArgs{
      Farm_ids: pulumi.IntArray{
        pulumi.Int(1),
      },
      Ipv4:     pulumi.Bool(true),
      Free_ips: pulumi.Int(1),
    }, pulumi.Provider(tfProvider))
    if err != nil {
      return err
    }
    gatewayName, err := provider.NewGatewayName(ctx, "gatewayName", &provider.GatewayNameArgs{
      Name: pulumi.String("pulumi"),
      Node_id: scheduler.Nodes.ApplyT(func(nodes []int) (int, error) {
        return nodes[0], nil
      }).(pulumi.IntOutput),
      Backends: pulumi.StringArray{
        pulumi.String("http://69.164.223.208"),
      },
    }, pulumi.Provider(tfProvider), pulumi.DependsOn([]pulumi.Resource{
      scheduler,
    }))
    if err != nil {
      return err
    }
    ctx.Export("node_deployment_id", gatewayName.Node_deployment_id)
    ctx.Export("fqdn", gatewayName.Fqdn)
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
    options:
      pluginDownloadURL: github://api.github.com/threefoldtech/pulumi-threefold # optional
    properties:
      mnemonic:

  scheduler:
    type: threefold:provider:Scheduler
    options:
      provider: ${provider}
    properties:
      farm_ids: [1]
      ipv4: true
      free_ips: 1

  gatewayName:
    type: threefold:provider:GatewayName
    options:
      provider: ${provider}
      dependsOn:
        - ${scheduler}
    properties:
      name: pulumi
      node_id: ${scheduler.nodes[0]}
      backends:
        - "http://69.164.223.208"

outputs:
  node_deployment_id: ${gatewayName.node_deployment_id}
  fqdn: ${gatewayName.fqdn}
```

{{% /choosable %}}

{{< /chooser >}}

### FQDN gateway resource

{{< chooser language "go,yaml" >}}

{{% choosable language go %}}

```go
package main

import (
  "fmt"
  "os"

  "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
  "github.com/threefoldtech/pulumi-threefold/sdk/go/threefold"
  "github.com/threefoldtech/pulumi-threefold/sdk/go/threefold/provider"
)

func main() {
  pulumi.Run(func(ctx *pulumi.Context) error {
    tfProvider, err := threefold.NewProvider(ctx, "provider", &threefold.ProviderArgs{
      Mnemonic: pulumi.String(os.Getenv("MNEMONIC")),
    })
    if err != nil {
      return err
    }
    scheduler, err := provider.NewScheduler(ctx, "scheduler", &provider.SchedulerArgs{
      Mru: pulumi.Int(1),
      Farm_ids: pulumi.IntArray{
        pulumi.Int(1),
      },
      Ipv4:     pulumi.Bool(true),
      Free_ips: pulumi.Int(1),
    }, pulumi.Provider(tfProvider))
    if err != nil {
      return err
    }
    network, err := provider.NewNetwork(ctx, "network", &provider.NetworkArgs{
      Name:        pulumi.String("test"),
      Description: pulumi.String("test network"),
      Nodes: pulumi.Array{
        scheduler.Nodes.ApplyT(func(nodes []int) (int, error) {
          return nodes[0], nil
        }).(pulumi.IntOutput),
      },
      Ip_range: pulumi.String("10.1.0.0/16"),
    }, pulumi.Provider(tfProvider), pulumi.DependsOn([]pulumi.Resource{
      scheduler,
    }))
    if err != nil {
      return err
    }
    deployment, err := provider.NewDeployment(ctx, "deployment", &provider.DeploymentArgs{
      Node_id: scheduler.Nodes.ApplyT(func(nodes []int) (int, error) {
        return nodes[0], nil
      }).(pulumi.IntOutput),
      Name:         pulumi.String("deployment"),
      Network_name: pulumi.String("test"),
      Vms: provider.VMInputArray{
        &provider.VMInputArgs{
          Name:         pulumi.String("vm"),
          Flist:        pulumi.String("https://hub.grid.tf/tf-official-apps/base:latest.flist"),
          Network_name: pulumi.String("test"),
          Cpu:          pulumi.Int(2),
          Memory:       pulumi.Int(256),
          Planetary:    pulumi.Bool(true),
        },
      },
    }, pulumi.Provider(tfProvider), pulumi.DependsOn([]pulumi.Resource{
      network,
    }))
    if err != nil {
      return err
    }
    gatewayFQDN, err := provider.NewGatewayFQDN(ctx, "gatewayFQDN", &provider.GatewayFQDNArgs{
      Name:    pulumi.String("testing"),
      Node_id: pulumi.Any(14),
      Fqdn:    pulumi.String("remote.omar.grid.tf"),
      Backends: pulumi.StringArray{
        deployment.Vms_computed.ApplyT(func(vms_computed []provider.VMComputed) (string, error) {
          return fmt.Sprintf("http://[%v]:9000", vms_computed[0].Planetary_ip), nil
        }).(pulumi.StringOutput),
      },
    }, pulumi.Provider(tfProvider), pulumi.DependsOn([]pulumi.Resource{
      deployment,
    }))
    if err != nil {
      return err
    }
    ctx.Export("node_deployment_id", gatewayFQDN.Node_deployment_id)
    ctx.Export("fqdn", gatewayFQDN.Fqdn)
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
    options:
      pluginDownloadURL: github://api.github.com/threefoldtech/pulumi-threefold # optional
    properties:
      mnemonic:

  scheduler:
    type: threefold:provider:Scheduler
    options:
      provider: ${provider}
    properties:
      mru: 0.25 # 256 megabytes
      farm_ids: [1]
      ipv4: true
      free_ips: 1

  network:
    type: threefold:provider:Network
    options:
      provider: ${provider}
      dependsOn:
        - ${scheduler}
    properties:
      name: test
      description: test network
      nodes:
        - ${scheduler.nodes[0]}
      ip_range: 10.1.0.0/16

  deployment:
    type: threefold:provider:Deployment
    options:
      provider: ${provider}
      dependsOn:
        - ${network}
    properties:
      node_id: ${scheduler.nodes[0]}
      name: deployment
      network_name: test
      vms:
        - name: vm
          flist: https://hub.grid.tf/tf-official-apps/base:latest.flist
          network_name: test
          cpu: 2
          memory: 256
          planetary: true

  gatewayFQDN:
    type: threefold:provider:GatewayFQDN
    options:
      provider: ${provider}
      dependsOn:
        - ${deployment}
    properties:
      name: testing
      node_id: 14
      fqdn: remote.omar.grid.tf
      backends:
        - http://[${deployment.vms_computed[0].planetary_ip}]:9000

outputs:
  node_deployment_id: ${gatewayFQDN.node_deployment_id}
  fqdn: ${gatewayFQDN.fqdn}
```

{{% /choosable %}}

{{< /chooser >}}

### ZDB resource

{{< chooser language "go,yaml" >}}

{{% choosable language go %}}

```go
package main

import (
  "fmt"
  "os"

  "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
  "github.com/threefoldtech/pulumi-threefold/sdk/go/threefold"
  "github.com/threefoldtech/pulumi-threefold/sdk/go/threefold/provider"
)

func main() {
  pulumi.Run(func(ctx *pulumi.Context) error {
    tfProvider, err := threefold.NewProvider(ctx, "provider", &threefold.ProviderArgs{
      Mnemonic: pulumi.String(os.Getenv("MNEMONIC")),
    })
    if err != nil {
      return err
    }
    scheduler, err := provider.NewScheduler(ctx, "scheduler", &provider.SchedulerArgs{
      Mru: pulumi.Int(1),
      Sru: pulumi.Int(2),
      Farm_ids: pulumi.IntArray{
        pulumi.Int(1),
      },
    }, pulumi.Provider(tfProvider))
    if err != nil {
      return err
    }
    deployment, err := provider.NewDeployment(ctx, "deployment", &provider.DeploymentArgs{
      Node_id: scheduler.Nodes.ApplyT(func(nodes []int) (int, error) {
        return nodes[0], nil
      }).(pulumi.IntOutput),
      Name: pulumi.String("zdb"),
      Zdbs: provider.ZDBInputArray{
        &provider.ZDBInputArgs{
          Name:     pulumi.String("zdbsTest"),
          Size:     pulumi.Int(2),
          Password: pulumi.String("123456"),
        },
      },
    }, pulumi.Provider(tfProvider))
    if err != nil {
      return err
    }
    ctx.Export("node_deployment_id", deployment.Node_deployment_id)
    ctx.Export("zdb_endpoint", pulumi.All(deployment.Zdbs_computed, deployment.Zdbs_computed).ApplyT(func(_args []interface{}) (string, error) {
      deploymentZdbs_computed := _args[0].([]provider.ZDBComputed)
      deploymentZdbs_computed1 := _args[1].([]provider.ZDBComputed)
      return fmt.Sprintf("[%v]:%v", deploymentZdbs_computed[0].Ips[1], deploymentZdbs_computed1[0].Port), nil
    }).(pulumi.StringOutput))
    ctx.Export("zdb_namespace", deployment.Zdbs_computed.ApplyT(func(zdbs_computed []provider.ZDBComputed) (*string, error) {
      return &zdbs_computed[0].Namespace, nil
    }).(pulumi.StringPtrOutput))
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
    options:
      pluginDownloadURL: github://api.github.com/threefoldtech/pulumi-threefold # optional
    properties:
      mnemonic:

  scheduler:
    type: threefold:provider:Scheduler
    options:
      provider: ${provider}
    properties:
      mru: 0.25 # 256 megabytes
      sru: 2
      farm_ids: [1]

  deployment:
    type: threefold:provider:Deployment
    options:
      provider: ${provider}
    properties:
      node_id: ${scheduler.nodes[0]}
      name: zdb
      zdbs:
        - name: zdbsTest
          size: 2
          password: "123456"

outputs:
  node_deployment_id: ${deployment.node_deployment_id}
  zdb_endpoint: "[${deployment.zdbs_computed[0].ips[1]}]:${deployment.zdbs_computed[0].port}"
  zdb_namespace: ${deployment.zdbs_computed[0].namespace}
```

{{% /choosable %}}

{{< /chooser >}}
