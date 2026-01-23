---
# WARNING: this file was fetched from https://raw.githubusercontent.com/threefoldtech/pulumi-threefold/v0.8.2/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/threefoldtech/pulumi-threefold/blob/v0.8.2/docs/_index.md
title: threefold
meta_desc: Provides an overview of the threefold grid Provider for Pulumi.
layout: package
---

The Threefold Resource Provider for the [threefold grid](https://threefold.io) lets you manage your infrastructure using Pulumi.

## Example

### Network resource

{{< chooser language "go,python,yaml,nodejs" >}}

{{% choosable language go %}}

```go
package main

import (
  "os"

  "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
  "github.com/threefoldtech/pulumi-threefold/sdk/go/threefold"
)

func main() {
  pulumi.Run(func(ctx *pulumi.Context) error {
    tfProvider, err := threefold.NewProvider(ctx, "provider", &threefold.ProviderArgs{
      Mnemonic: pulumi.String(os.Getenv("MNEMONIC")),
    })
    if err != nil {
      return err
    }

    scheduler, err := threefold.NewScheduler(ctx, "scheduler", &threefold.SchedulerArgs{
      Farm_ids: pulumi.IntArray{
        pulumi.Int(1),
      },
    }, pulumi.Provider(tfProvider))
    if err != nil {
      return err
    }

    network, err := threefold.NewNetwork(ctx, "network", &threefold.NetworkArgs{
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

{{% choosable language python %}}

```python
import os
import pulumi
import pulumi_threefold as threefold

mnemonic = os.environ['MNEMONIC']
network = os.environ['NETWORK']

provider = threefold.Provider("provider", mnemonic=mnemonic, network=network)
scheduler = threefold.Scheduler("scheduler", farm_ids=[1],
opts=pulumi.ResourceOptions(provider=provider))
network = threefold.Network("network",
    name="example",
    description="example network",
    nodes=[scheduler.nodes[0]],
    ip_range="10.1.0.0/16",
    opts=pulumi.ResourceOptions(provider=provider,
        depends_on=[scheduler]))
pulumi.export("node_deployment_id", network.node_deployment_id)
pulumi.export("nodes_ip_range", network.nodes_ip_range)
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
    type: threefold:Scheduler
    options:
      provider: ${provider}
    properties:
      farm_ids: [1]

  network:
    type: threefold:Network
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

{{% choosable language nodejs %}}

```ts
import * as threefold from "@threefold/pulumi";

const provider = new threefold.Provider("provider", {mnemonic: process.env.MNEMONIC, network: process.env.NETWORK});
const scheduler = new threefold.Scheduler("scheduler", {farm_ids: [1]}, {
    provider: provider,
});
const network = new threefold.Network("network", {
    name: "testing",
    description: "test network",
    nodes: [scheduler.nodes[0]],
    ip_range: "10.1.0.0/16",
}, {
    provider: provider,
    dependsOn: [scheduler],
});
export const nodeDeploymentId = network.node_deployment_id;
export const nodesIpRange = network.nodes_ip_range;
```

{{% /choosable %}}

{{< /chooser >}}

### Virtual machine resource

{{< chooser language "go,python,yaml,nodejs" >}}

{{% choosable language go %}}

```go
package main

import (
  "os"

  "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
  "github.com/threefoldtech/pulumi-threefold/sdk/go/threefold"
)

func main() {
  pulumi.Run(func(ctx *pulumi.Context) error {
    tfProvider, err := threefold.NewProvider(ctx, "provider", &threefold.ProviderArgs{
      Mnemonic: pulumi.String(os.Getenv("MNEMONIC")),
    })
    if err != nil {
      return err
    }
    scheduler, err := threefold.NewScheduler(ctx, "scheduler", &threefold.SchedulerArgs{
      Mru: pulumi.Int(1),
      Sru: pulumi.Int(2),
      Farm_ids: pulumi.IntArray{
        pulumi.Int(1),
      },
    }, pulumi.Provider(tfProvider))
    if err != nil {
      return err
    }
    network, err := threefold.NewNetwork(ctx, "network", &threefold.NetworkArgs{
      Name:        pulumi.String("test"),
      Description: pulumi.String("test network"),
      Nodes: pulumi.Array{
        scheduler.Nodes.ApplyT(func(nodes []int) (int, error) {
          return nodes[0], nil
        }).(pulumi.IntOutput),
      },
      Ip_range: pulumi.String("10.1.0.0/16"),
      Mycelium: pulumi.Bool(true),
    }, pulumi.Provider(tfProvider), pulumi.DependsOn([]pulumi.Resource{
      scheduler,
    }))
    if err != nil {
      return err
    }
    deployment, err := threefold.NewDeployment(ctx, "deployment", &threefold.DeploymentArgs{
      Node_id: scheduler.Nodes.ApplyT(func(nodes []int) (int, error) {
        return nodes[0], nil
      }).(pulumi.IntOutput),
      Name:         pulumi.String("deployment"),
      Network_name: pulumi.String("test"),
      Vms: threefold.VMInputArray{
        &threefold.VMInputArgs{
          Name:         pulumi.String("vm"),
          Node_id: scheduler.Nodes.ApplyT(func(nodes []int) (int, error) {
            return nodes[0], nil
          }).(pulumi.IntOutput),
          Flist:        pulumi.String("https://hub.grid.tf/tf-official-apps/base:latest.flist"),
          Entrypoint:   pulumi.String("/sbin/zinit init"),
          Network_name: pulumi.String("test"),
          Cpu:          pulumi.Int(2),
          Memory:       pulumi.Int(256),
          Planetary:    pulumi.Bool(true),
          Mycelium:     pulumi.Bool(true),
          Mounts: threefold.MountArray{
            &threefold.MountArgs{
              Name:   pulumi.String("data"),
              Mount_point: pulumi.String("/app"),
            },
          },
          Env_vars: pulumi.StringMap{
            "SSH_KEY": nil,
          },
        },
      },
      Disks: threefold.DiskArray{
        &threefold.DiskArgs{
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
    ctx.Export("planetary_ip", deployment.Vms_computed.ApplyT(func(vms_computed []threefold.VMComputed) (*string, error) {
      return &vms_computed[0].Planetary_ip, nil
    }).(pulumi.StringPtrOutput))
    ctx.Export("mycelium_ip", deployment.Vms_computed.ApplyT(func(vms_computed []threefold.VMComputed) (*string, error) {
      return &vms_computed[0].Mycelium_ip, nil
    }).(pulumi.StringPtrOutput))
    return nil
  })
}
```

{{% /choosable %}}

{{% choosable language python %}}

```python
import os
import pulumi
import pulumi_threefold as threefold

mnemonic = os.environ['MNEMONIC']
network = os.environ['NETWORK']

provider = threefold.Provider("provider", mnemonic=mnemonic, network=network)
scheduler = threefold.Scheduler("scheduler", farm_ids=[1],
opts=pulumi.ResourceOptions(provider=provider))
network = threefold.Network("network",
    name="example",
    description="example network",
    nodes=[scheduler.nodes[0]],
    ip_range="10.1.0.0/16",
    mycelium=True,
    opts=pulumi.ResourceOptions(provider=provider,
        depends_on=[scheduler]))

deployment = threefold.Deployment("deployment",
    node_id=scheduler.nodes[0],
    name="deployment",
    network_name="example",
    vms=[threefold.VMInputArgs(
        name="vm",
        node_id=scheduler.nodes[0],
        flist="https://hub.grid.tf/tf-official-apps/base:latest.flist",
        entrypoint="/sbin/zinit init",
        network_name="test",
        cpu=2,
        memory=256, #MB
        mycelium=True,
        mounts=[threefold.MountArgs(
            name="data",
            mount_point="/app",
        )],
        env_vars={
            "SSH_KEY": None,
        },
    )],
    disks=[threefold.DiskArgs(
        name="data",
        size=2,
    )],
    opts=pulumi.ResourceOptions(provider=provider,
        depends_on=[network]))

pulumi.export("node_deployment_id", deployment.node_deployment_id)
pulumi.export("planetary_ip", deployment.vms_computed[0].planetary_ip)
pulumi.export("mycelium_ip", deployment.vms_computed[0].mycelium_ip)
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
    type: threefold:Scheduler
    options:
      provider: ${provider}
    properties:
      mru: 0.25 # 256 megabytes
      sru: 2
      farm_ids: [1]

  network:
    type: threefold:Network
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
    type: threefold:Deployment
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
          node_id: ${scheduler.nodes[0]}
          flist: https://hub.grid.tf/tf-official-apps/base:latest.flist
          entrypoint: "/sbin/zinit init"
          network_name: test
          cpu: 2
          memory: 256
          planetary: true
          mycelium: true
          mounts:
            - name: data
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

{{% choosable language nodejs %}}

```ts
import * as threefold from "@threefold/pulumi";

const provider = new threefold.Provider("provider", {mnemonic: process.env.MNEMONIC, network: process.env.NETWORK});
const scheduler = new threefold.Scheduler("scheduler", {
    mru: 0.25,
    sru: 2,
    farm_ids: [1],
}, {
    provider: provider,
});
const network = new threefold.Network("network", {
    name: "test",
    description: "test network",
    nodes: [scheduler.nodes[0]],
    ip_range: "10.1.0.0/16",
    mycelium: true,
}, {
    provider: provider,
    dependsOn: [scheduler],
});
const deployment = new threefold.Deployment("deployment", {
    node_id: scheduler.nodes[0],
    name: "deployment",
    network_name: "test",
    vms: [{
        name: "vm",
        node_id: scheduler.nodes[0],
        flist: "https://hub.grid.tf/tf-official-apps/base:latest.flist",
        entrypoint: "/sbin/zinit init",
        network_name: "test",
        cpu: 2,
        memory: 256,
        planetary: true,
        mycelium: true,
        mounts: [{
            name: "data",
            mount_point: "/app",
        }],
        env_vars: {
            SSH_KEY: "",
        },
    }],
    disks: [{
        name: "data",
        size: 2,
    }],
}, {
    provider: provider,
    dependsOn: [network],
});
export const nodeDeploymentId = deployment.node_deployment_id;
export const planetaryIp = deployment.vms_computed.apply(vms_computed => vms_computed[0].planetary_ip);
export const myceliumIp = deployment.vms_computed.apply(vms_computed => vms_computed[0].mycelium_ip);
```

{{% /choosable %}}

{{< /chooser >}}

### Kubernetes resource

{{< chooser language "go,python,yaml,nodejs" >}}

{{% choosable language go %}}

```go
package main

import (
  "os"

  "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
  "github.com/threefoldtech/pulumi-threefold/sdk/go/threefold"
)

func main() {
  pulumi.Run(func(ctx *pulumi.Context) error {
    tfProvider, err := threefold.NewProvider(ctx, "provider", &threefold.ProviderArgs{
      Mnemonic: pulumi.String(os.Getenv("MNEMONIC")),
    })
    if err != nil {
      return err
    }
    scheduler, err := threefold.NewScheduler(ctx, "scheduler", &threefold.SchedulerArgs{
      Mru: pulumi.Int(6),
      Sru: pulumi.Int(6),
      Farm_ids: pulumi.IntArray{
        pulumi.Int(1),
      },
    }, pulumi.Provider(tfProvider))
    if err != nil {
      return err
    }
    network, err := threefold.NewNetwork(ctx, "network", &threefold.NetworkArgs{
      Name:        pulumi.String("test"),
      Description: pulumi.String("test network"),
      Nodes: pulumi.Array{
        scheduler.Nodes.ApplyT(func(nodes []int) (int, error) {
          return nodes[0], nil
        }).(pulumi.IntOutput),
      },
      Ip_range: pulumi.String("10.1.0.0/16"),
      Mycelium: pulumi.Bool(true),
    }, pulumi.Provider(tfProvider), pulumi.DependsOn([]pulumi.Resource{
      scheduler,
    }))
    if err != nil {
      return err
    }
    kubernetes, err := threefold.NewKubernetes(ctx, "kubernetes", &threefold.KubernetesArgs{
      Master: &threefold.K8sNodeInputArgs{
        Name:         pulumi.String("kubernetes"),
        Network_name: pulumi.String("test"),
        NodeID: scheduler.Nodes.ApplyT(func(nodes []int) (int, error) {
          return nodes[0], nil
        }).(pulumi.IntOutput),
        Planetary: pulumi.Bool(true),
        Mycelium:  pulumi.Bool(true),
        Cpu:       pulumi.Int(2),
        Memory:    pulumi.Int(2048),
        Disk_size: pulumi.Int(2),
      },
      Workers: threefold.K8sNodeInputArray{
        &threefold.K8sNodeInputArgs{
          Name: pulumi.String("worker1"),
          Network_name: pulumi.String("test"),
          Node_id: scheduler.Nodes.ApplyT(func(nodes []int) (int, error) {
            return nodes[0], nil
          }).(pulumi.IntOutput),
          Disk_size: pulumi.Int(2),
          Cpu:       pulumi.Int(2),
          Memory:    pulumi.Int(2048),
        },
        &threefold.K8sNodeInputArgs{
          Name: pulumi.String("worker2"),
          Network_name: pulumi.String("test"),
          Node_id: scheduler.Nodes.ApplyT(func(nodes []int) (int, error) {
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
    ctx.Export("planetary_ip", kubernetes.Master_computed.ApplyT(func(master_computed threefold.K8sNodeComputed) (*string, error) {
      return &master_computed.Planetary_ip, nil
    }).(pulumi.StringPtrOutput))
    ctx.Export("mycelium_ip", kubernetes.Master_computed.ApplyT(func(master_computed threefold.K8sNodeComputed) (*string, error) {
      return &master_computed.Mycelium_ip, nil
    }).(pulumi.StringPtrOutput))
    return nil
  })
}
```

{{% /choosable %}}

{{% choosable language python %}}

```python
import os
import pulumi
import pulumi_threefold as threefold

mnemonic = os.environ['MNEMONIC']
network = os.environ['NETWORK']

provider = threefold.Provider("provider", mnemonic=mnemonic, network=network)
scheduler = threefold.Scheduler("scheduler",
    mru=6,
    sru=6,
    farm_ids=[1],
    opts = pulumi.ResourceOptions(provider=provider))
network = threefold.Network("network",
    name="example",
    description="example network",
    nodes=[scheduler.nodes[0]],
    ip_range="10.1.0.0/16",
    mycelium=True,
    opts = pulumi.ResourceOptions(provider=provider,
        depends_on=[scheduler]))
kubernetes = threefold.Kubernetes("kubernetes",
    master=threefold.K8sNodeInputArgs(
        name="kubernetes",
        network_name="example",
        node_id=scheduler.nodes[0],
        disk_size=2,
        planetary=True,
        mycelium=True,
        cpu=2,
        memory=2048,
    ),
    workers=[
        threefold.K8sNodeInputArgs(
            name="worker1",
            network_name="example",
            node_id=scheduler.nodes[0],
            disk_size=2,
            cpu=2,
            memory=2048,
        ),
        threefold.K8sNodeInputArgs(
            name="worker2",
            network_name="example",
            node_id=scheduler.nodes[0],
            disk_size=2,
            cpu=2,
            memory=2048,
        ),
    ],
    token="t123456789",
    network_name="example",
    ssh_key=None,
    opts = pulumi.ResourceOptions(provider=provider,
        depends_on=[network]))

pulumi.export("node_deployment_id", kubernetes.node_deployment_id)
pulumi.export("master", kubernetes.master_computed)
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
    type: threefold:Scheduler
    options:
      provider: ${provider}
    properties:
      mru: 6
      sru: 6
      farm_ids: [1]

  network:
    type: threefold:Network
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

  kubernetes:
    type: threefold:Kubernetes
    options:
      provider: ${provider}
      dependsOn:
        - ${network}
    properties:
      master:
        name: kubernetes
        network_name: test
        node_id: ${scheduler.nodes[0]}
        disk_size: 2
        planetary: true
        mycelium: true
        cpu: 2
        memory: 2048

      workers:
        - name: worker1
          network_name: test
          node_id: ${scheduler.nodes[0]}
          disk_size: 2
          cpu: 2
          memory: 2048
          mycelium: true
        - name: worker2
          network_name: test
          node_id: ${scheduler.nodes[0]}
          disk_size: 2
          cpu: 2
          memory: 2048
          mycelium: true

      token: t123456789
      network_name: test
      ssh_key:

outputs:
  node_deployment_id: ${kubernetes.node_deployment_id}
  planetary_ip: ${kubernetes.master_computed.planetary_ip}
  mycelium_ip: ${kubernetes.master_computed.mycelium_ip}
```

{{% /choosable %}}

{{% choosable language nodejs %}}

```ts
import * as threefold from "@threefold/pulumi";

const provider = new threefold.Provider("provider", {mnemonic: process.env.MNEMONIC, network: process.env.NETWORK});
const scheduler = new threefold.Scheduler("scheduler", {
    mru: 6,
    sru: 6,
    farm_ids: [1],
}, {
    provider: provider,
});
const network = new threefold.Network("network", {
    name: "test",
    description: "test network",
    nodes: [scheduler.nodes[0]],
    ip_range: "10.1.0.0/16",
    mycelium: true,
}, {
    provider: provider,
    dependsOn: [scheduler],
});
const kubernetes = new threefold.Kubernetes("kubernetes", {
    master: {
        name: "kubernetes",
        network_name: "test",
        node_id: scheduler.nodes[0],
        disk_size: 2,
        planetary: true,
        mycelium: true,
        cpu: 2,
        memory: 2048,
    },
    workers: [
        {
            name: "worker1",
            network_name: "test",
            node_id: scheduler.nodes[0],
            disk_size: 2,
            cpu: 2,
            memory: 2048,
            mycelium: true,
        },
        {
            name: "worker2",
            network_name: "test",
            node_id: scheduler.nodes[0],
            disk_size: 2,
            cpu: 2,
            memory: 2048,
            mycelium: true,
        },
    ],
    token: "t123456789",
    network_name: "test",
    ssh_key: undefined,
}, {
    provider: provider,
    dependsOn: [network],
});
export const nodeDeploymentId = kubernetes.node_deployment_id;
export const planetaryIp = kubernetes.master_computed.apply(master_computed => master_computed.planetary_ip);
export const myceliumIp = kubernetes.master_computed.apply(master_computed => master_computed.mycelium_ip);
```

{{% /choosable %}}

{{< /chooser >}}

### Name gateway resource

{{< chooser language "go,python,yaml,nodejs" >}}

{{% choosable language go %}}

```go
package main

import (
  "os"

  "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
  "github.com/threefoldtech/pulumi-threefold/sdk/go/threefold"
)

func main() {
  pulumi.Run(func(ctx *pulumi.Context) error {
    tfProvider, err := threefold.NewProvider(ctx, "provider", &threefold.ProviderArgs{
      Mnemonic: pulumi.String(os.Getenv("MNEMONIC")),
    })
    if err != nil {
      return err
    }
    scheduler, err := threefold.NewScheduler(ctx, "scheduler", &threefold.SchedulerArgs{
      Farm_ids: pulumi.IntArray{
        pulumi.Int(1),
      },
      Ipv4:     pulumi.Bool(true),
      Free_ips: pulumi.Int(1),
    }, pulumi.Provider(tfProvider))
    if err != nil {
      return err
    }
    gatewayName, err := threefold.NewGatewayName(ctx, "gatewayName", &threefold.GatewayNameArgs{
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

{{% choosable language python %}}

```python
import os
import pulumi
import pulumi_threefold as threefold

mnemonic = os.environ['MNEMONIC']
network = os.environ['NETWORK']

provider = threefold.Provider("provider", mnemonic=mnemonic, network=network)
scheduler = threefold.Scheduler("scheduler",
    farm_ids=[1],
    ipv4=True,
    free_ips=1,
    opts = pulumi.ResourceOptions(provider=provider))
gateway_name = threefold.GatewayName("gatewayName",
    name="pulumi",
    node_id=scheduler.nodes[0],
    backends=["http://69.164.223.208"],
    opts = pulumi.ResourceOptions(provider=provider,
        depends_on=[scheduler]))
pulumi.export("node_deployment_id", gateway_name.node_deployment_id)
pulumi.export("fqdn", gateway_name.fqdn)
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
    type: threefold:Scheduler
    options:
      provider: ${provider}
    properties:
      farm_ids: [1]
      ipv4: true
      free_ips: 1

  gatewayName:
    type: threefold:GatewayName
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

{{% choosable language nodejs %}}

```ts
import * as threefold from "@threefold/pulumi";

const provider = new threefold.Provider("provider", {mnemonic: process.env.MNEMONIC, network: process.env.NETWORK});
const scheduler = new threefold.Scheduler("scheduler", {
    farm_ids: [1],
    ipv4: true,
    free_ips: 1,
}, {
    provider: provider,
});
const gatewayName = new threefold.GatewayName("gatewayName", {
    name: "pulumi",
    node_id: scheduler.nodes[0],
    backends: ["http://69.164.223.208"],
}, {
    provider: provider,
    dependsOn: [scheduler],
});
export const nodeDeploymentId = gatewayName.node_deployment_id;
export const fqdn = gatewayName.fqdn;
```

{{% /choosable %}}

{{< /chooser >}}

### FQDN gateway resource

{{< chooser language "go,python,yaml,nodejs" >}}

{{% choosable language go %}}

```go
package main

import (
  "fmt"
  "os"

  "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
  "github.com/threefoldtech/pulumi-threefold/sdk/go/threefold"
)

func main() {
  pulumi.Run(func(ctx *pulumi.Context) error {
    tfProvider, err := threefold.NewProvider(ctx, "provider", &threefold.ProviderArgs{
      Mnemonic: pulumi.String(os.Getenv("MNEMONIC")),
    })
    if err != nil {
      return err
    }
    scheduler, err := threefold.NewScheduler(ctx, "scheduler", &threefold.SchedulerArgs{
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
    network, err := threefold.NewNetwork(ctx, "network", &threefold.NetworkArgs{
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
    deployment, err := threefold.NewDeployment(ctx, "deployment", &threefold.DeploymentArgs{
      Node_id: scheduler.Nodes.ApplyT(func(nodes []int) (int, error) {
        return nodes[0], nil
      }).(pulumi.IntOutput),
      Name:         pulumi.String("deployment"),
      Network_name: pulumi.String("test"),
      Vms: threefold.VMInputArray{
        &threefold.VMInputArgs{
          Name:         pulumi.String("vm"),
          Node_id: scheduler.Nodes.ApplyT(func(nodes []int) (int, error) {
            return nodes[0], nil
          }).(pulumi.IntOutput),
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
    gatewayFQDN, err := threefold.NewGatewayFQDN(ctx, "gatewayFQDN", &threefold.GatewayFQDNArgs{
      Name:    pulumi.String("testing"),
      Node_id: pulumi.Any(14),
      Fqdn:    pulumi.String("remote.omar.grid.tf"),
      Backends: pulumi.StringArray{
        deployment.Vms_computed.ApplyT(func(vms_computed []threefold.VMComputed) (string, error) {
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

{{% choosable language python %}}

```python
import os
import pulumi
import pulumi_threefold as threefold

mnemonic = os.environ['MNEMONIC']
network = os.environ['NETWORK']

provider = threefold.Provider("provider", mnemonic=mnemonic, network=network)
scheduler = threefold.Scheduler("scheduler",
    mru=0.25,
    farm_ids=[1],
    ipv4=True,
    free_ips=1,
    opts = pulumi.ResourceOptions(provider=provider))
network = threefold.Network("network",
    name="example",
    description="example network",
    nodes=[scheduler.nodes[0]],
    ip_range="10.1.0.0/16",
    opts = pulumi.ResourceOptions(provider=provider,
        depends_on=[scheduler]))
deployment = threefold.Deployment("deployment",
    node_id=scheduler.nodes[0],
    name="deployment",
    network_name="example",
    vms=[threefold.VMInputArgs(
        name="vm",
        node_id=scheduler.nodes[0],
        flist="https://hub.grid.tf/tf-official-apps/base:latest.flist",
        network_name="example",
        cpu=2,
        memory=256,
        planetary=True,
    )],
    opts = pulumi.ResourceOptions(provider=provider,
        depends_on=[network]))
gateway_fqdn = threefold.GatewayFQDN("gatewayFQDN",
    name="testing",
    node_id=14,
    fqdn="remote.omar.grid.tf",
    backends=[deployment.vms_computed.apply(lambda vms_computed: f"http://[{vms_computed[0].planetary_ip}]:9000")],
    opts = pulumi.ResourceOptions(provider=provider,
        depends_on=[deployment]))
pulumi.export("node_deployment_id", gateway_fqdn.node_deployment_id)
pulumi.export("fqdn", gateway_fqdn.fqdn)
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
    type: threefold:Scheduler
    options:
      provider: ${provider}
    properties:
      mru: 0.25 # 256 megabytes
      farm_ids: [1]
      ipv4: true
      free_ips: 1

  network:
    type: threefold:Network
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
    type: threefold:Deployment
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
          node_id: ${scheduler.nodes[0]}
          flist: https://hub.grid.tf/tf-official-apps/base:latest.flist
          network_name: test
          cpu: 2
          memory: 256
          planetary: true

  gatewayFQDN:
    type: threefold:GatewayFQDN
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

{{% choosable language nodejs %}}

```ts
import * as pulumi from "@pulumi/pulumi";
import * as threefold from "@threefold/pulumi";

const provider = new threefold.Provider("provider", {mnemonic: process.env.MNEMONIC, network: process.env.NETWORK});
const scheduler = new threefold.Scheduler("scheduler", {
    mru: 0.25,
    farm_ids: [1],
    ipv4: true,
    free_ips: 1,
}, {
    provider: provider,
});
const network = new threefold.Network("network", {
    name: "test",
    description: "test network",
    nodes: [scheduler.nodes[0]],
    ip_range: "10.1.0.0/16",
}, {
    provider: provider,
    dependsOn: [scheduler],
});
const deployment = new threefold.Deployment("deployment", {
    node_id: scheduler.nodes[0],
    name: "deployment",
    network_name: "test",
    vms: [{
        name: "vm",
        node_id: scheduler.nodes[0],
        flist: "https://hub.grid.tf/tf-official-apps/base:latest.flist",
        network_name: "test",
        cpu: 2,
        memory: 256,
        planetary: true,
    }],
}, {
    provider: provider,
    dependsOn: [network],
});
const gatewayFQDN = new threefold.GatewayFQDN("gatewayFQDN", {
    name: "testing",
    node_id: 14,
    fqdn: "remote.omar.grid.tf",
    backends: [pulumi.interpolate `http://[${deployment.vms_computed[0].planetary_ip}]:9000`],
}, {
    provider: provider,
    dependsOn: [deployment],
});
export const nodeDeploymentId = gatewayFQDN.node_deployment_id;
export const fqdn = gatewayFQDN.fqdn;
```

{{% /choosable %}}

{{< /chooser >}}

### ZDB resource

{{< chooser language "go,python,yaml,nodejs" >}}

{{% choosable language go %}}

```go
package main

import (
  "fmt"
  "os"

  "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
  "github.com/threefoldtech/pulumi-threefold/sdk/go/threefold"
)

func main() {
  pulumi.Run(func(ctx *pulumi.Context) error {
    tfProvider, err := threefold.NewProvider(ctx, "provider", &threefold.ProviderArgs{
      Mnemonic: pulumi.String(os.Getenv("MNEMONIC")),
    })
    if err != nil {
      return err
    }
    scheduler, err := threefold.NewScheduler(ctx, "scheduler", &threefold.SchedulerArgs{
      Mru: pulumi.Int(1),
      Sru: pulumi.Int(2),
      Farm_ids: pulumi.IntArray{
        pulumi.Int(1),
      },
    }, pulumi.Provider(tfProvider))
    if err != nil {
      return err
    }
    deployment, err := threefold.NewDeployment(ctx, "deployment", &threefold.DeploymentArgs{
      Node_id: scheduler.Nodes.ApplyT(func(nodes []int) (int, error) {
        return nodes[0], nil
      }).(pulumi.IntOutput),
      Name: pulumi.String("zdb"),
      Zdbs: threefold.ZDBInputArray{
        &threefold.ZDBInputArgs{
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
      deploymentZdbs_computed := _args[0].([]threefold.ZDBComputed)
      deploymentZdbs_computed1 := _args[1].([]threefold.ZDBComputed)
      return fmt.Sprintf("[%v]:%v", deploymentZdbs_computed[0].Ips[1], deploymentZdbs_computed1[0].Port), nil
    }).(pulumi.StringOutput))
    ctx.Export("zdb_namespace", deployment.Zdbs_computed.ApplyT(func(zdbs_computed []threefold.ZDBComputed) (*string, error) {
      return &zdbs_computed[0].Namespace, nil
    }).(pulumi.StringPtrOutput))
    return nil
  })
}
```

{{% /choosable %}}

{{% choosable language python %}}

```python
import os
import pulumi
import pulumi_threefold as threefold

mnemonic = os.environ['MNEMONIC']
network = os.environ['NETWORK']

provider = threefold.Provider("provider", mnemonic=mnemonic, network=network)
scheduler = threefold.Scheduler("scheduler",
    mru=0.25,
    sru=2,
    farm_ids=[1],
    opts = pulumi.ResourceOptions(provider=provider))
deployment = threefold.Deployment("deployment",
    node_id=scheduler.nodes[0],
    name="zdb",
    zdbs=[threefold.ZDBInputArgs(
        name="zdbsTest",
        size=2,
        password="123456",
    )],
    opts = pulumi.ResourceOptions(provider=provider))

pulumi.export("node_deployment_id", deployment.node_deployment_id)
pulumi.export("zdb_endpoint_ip", deployment.zdbs_computed[0].ips[1])
pulumi.export("zdb_endpoint_port", deployment.zdbs_computed[0].port)
pulumi.export("zdb_namespace", deployment.zdbs_computed[0].namespace)
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
    type: threefold:Scheduler
    options:
      provider: ${provider}
    properties:
      mru: 0.25 # 256 megabytes
      sru: 2
      farm_ids: [1]

  deployment:
    type: threefold:Deployment
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

{{% choosable language nodejs %}}

```ts
import * as pulumi from "@pulumi/pulumi";
import * as threefold from "@threefold/pulumi";

const provider = new threefold.Provider("provider", {mnemonic: process.env.MNEMONIC, network: process.env.NETWORK});
const scheduler = new threefold.Scheduler("scheduler", {
    mru: 0.25,
    sru: 2,
    farm_ids: [1],
}, {
    provider: provider,
});
const deployment = new threefold.Deployment("deployment", {
    node_id: scheduler.nodes[0],
    name: "zdb",
    zdbs: [{
        name: "zdbsTest",
        size: 2,
        password: "123456",
    }],
}, {
    provider: provider,
});
export const nodeDeploymentId = deployment.node_deployment_id;
export const zdbEndpoint = pulumi.all([deployment.zdbs_computed, deployment.zdbs_computed]).apply(([deploymentZdbs_computed, deploymentZdbs_computed1]) => `[${deploymentZdbs_computed[0].ips?.[1]}]:${deploymentZdbs_computed1[0].port}`);
export const zdbNamespace = deployment.zdbs_computed.apply(zdbs_computed => zdbs_computed[0].namespace);
```

{{% /choosable %}}

{{< /chooser >}}
