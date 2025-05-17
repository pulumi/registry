---
# WARNING: this file was fetched from https://raw.githubusercontent.com/genesiscloud/pulumi-genesiscloud/v0.0.35/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Genesis Cloud
meta_desc: Provides an overview of the Genesis Cloud Provider for Pulumi.
layout: package
---

The Genesis Cloud provider for Pulumi can be used to provision any of the cloud resources available in [Genesis Cloud](https://www.genesiscloud.com).
The Genesis Cloud provider must be configured with credentials to deploy and update resources.

## Example

{{< chooser language "typescript,python,go,csharp" >}}
{{% choosable language typescript %}}

```typescript
import {
  SecurityGroup,
  SSHKey,
  Instance,
} from "@genesiscloud/pulumi-genesiscloud";

const region = "ARC-IS-HAF-1";

const sshKey = new SSHKey("ssh-key", {
  name: "ssh-key",
  publicKey: "<YOUR_SSH_PUBLIC_KEY>",
});

const allowSSH = new SecurityGroup("allow-ssh", {
  name: "allow-ssh",
  region,
  description: "Allow SSH access",
  rules: [
    {
      direction: "ingress",
      protocol: "tcp",
      portRangeMin: 22,
      portRangeMax: 22,
    },
  ],
});

const firstPulumiInstance = new Instance("first-pulumi-instance", {
  name: "first-pulumi-instance",
  region,
  image: "ubuntu-ml-nvidia-pytorch",
  type: "vcpu-4_memory-12g_disk-80g_nvidia3080-1",
  sshKeyIds: [sshKey.id],
  securityGroupIds: [allowSSH.id],
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_genesiscloud as gc

region = "ARC-IS-HAF-1"

ssh_key = gc.SSHKey(
    "ssh-key",
    name="ssh-key",
    public_key="<YOUR_SSH_PUBLIC_KEY>",
)

allow_ssh = gc.SecurityGroup(
    "allow-ssh",
    name="allow-ssh",
    description="Allow SSH",
    region=region,
    rules=[
        gc.SecurityGroupRuleArgs(
            direction="ingress",
            protocol="tcp",
            port_range_min=22,
            port_range_max=22,
        )
    ],
)

instance = gc.Instance(
    "my-pulumi-instance",
    name="my-pulumi-instance",
    region=region,
    image="ubuntu-ml-nvidia-pytorch",
    type="vcpu-4_memory-12g_disk-80g_nvidia3080-1",
    ssh_key_ids=[ssh_key.id],
    security_group_ids=[allow_ssh.id],
)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
package main

import (
  "github.com/genesiscloud/pulumi-genesiscloud/sdk/go/genesiscloud"
  "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
  pulumi.Run(func(ctx *pulumi.Context) error {
    region := pulumi.String("ARC-IS-HAF-1")
    sshKey, _ := genesiscloud.NewSSHKey(ctx, "ssh-key", &genesiscloud.SSHKeyArgs{
      PublicKey: pulumi.String("<YOUR_SSH_PUBLIC_KEY>"),
    })

		allowSSH, _ := genesiscloud.NewSecurityGroup(ctx, "allow-ssh",
			&genesiscloud.SecurityGroupArgs{
				Name:        pulumi.String("allow-ssh"),
				Description: pulumi.String("Allow SSH"),
				Region:      region,
				Rules: genesiscloud.SecurityGroupRuleArray{
					genesiscloud.SecurityGroupRuleArgs{
						Direction:    pulumi.String("ingress"),
						Protocol:     pulumi.String("tcp"),
						PortRangeMin: pulumi.Int(22),
						PortRangeMax: pulumi.Int(22),
					},
				},
			})

		genesiscloud.NewInstance(ctx, "my-pulumi-instance", &genesiscloud.InstanceArgs{
			Region:           region,
			Image:            pulumi.String("ubuntu-ml-nvidia-pytorch"),
			Name:             pulumi.String("my-pulumi-instance"),
			Type:             pulumi.String("vcpu-4_memory-12g_disk-80g_nvidia3080-1"),
			SshKeyIds:        pulumi.StringArray{sshKey.ID()},
			SecurityGroupIds: pulumi.StringArray{allowSSH.ID()},
		})

		return nil
	})
}
```

{{% /choosable %}}
{{% choosable language csharp %}}

```csharp
using Pulumi;
using GenesisCloud.PulumiPackage.Genesiscloud;
using GenesisCloud.PulumiPackage.Genesiscloud.Inputs;
using System.Threading.Tasks;

class GenesisCloudInstance : Stack
{
    public GenesisCloudInstance()
    {
        var region = "ARC-IS-HAF-1";

        var ssh_key = new SSHKey("ssh-key", new SSHKeyArgs
        {
            Name = "ssh-key",
            PublicKey = "<YOUR_SSH_PUBLIC_KEY>"
        });

        var allow_ssh = new SecurityGroup("allow_ssh", new SecurityGroupArgs
        {
            Name = "allow-ssh",
            Description = "Allow SSH",
            Region = region,
            Rules = {
              new SecurityGroupRuleArgs
              {
                Direction = "ingress",
                Protocol = "tcp",
                PortRangeMin = 22,
                PortRangeMax = 22,
              }
            }
        });

        var instance = new Instance("my-pulumi-instance", new InstanceArgs
        {
            Name = "my-pulumi-instance",
            Region = region,
            Image = "ubuntu-ml-nvidia-pytorch",
            Type = "vcpu-4_memory-12g_disk-80g_nvidia3080-1",
            SshKeyIds = { ssh_key.Id },
            SecurityGroupIds = { allow_ssh.Id },
        });
    }
}

class Program
{
    static Task<int> Main(string[] args) => Deployment.RunAsync<GenesisCloudInstance>();
}
```

{{% /choosable %}}

{{< /chooser >}}
