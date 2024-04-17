---
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
import * as pulumi from "@pulumi/pulumi";
import { Instance } from "@genesiscloud/pulumi-genesiscloud";

const example = new Instance("example", {
  image: "ubuntu-ml-nvidia-pytorch",
  region: "ARC-IS-HAF-1",
  sshKeyIds: ["my-ssh-key-id"],
  type: "vcpu-4_memory-12g_disk-80g_nvidia3080-1",
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi
import pulumi_genesiscloud as genesiscloud

example = genesiscloud.Instance("example",
  image="ubuntu-ml-nvidia-pytorch",
  region="ARC-IS-HAF-1",
  ssh_key_ids=["my-ssh-key-id"],
  type="vcpu-4_memory-12g_disk-80g_nvidia3080-1")
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"github.com/genesiscloud/pulumi-genesiscloud/sdk/go/genesiscloud"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
  pulumi.Run(func(ctx *pulumi.Context) error {
    _, err := genesiscloud.NewInstance(ctx, "example", &genesiscloud.InstanceArgs{
      Image:  pulumi.String("ubuntu-ml-nvidia-pytorch"),
      Region: pulumi.String("ARC-IS-HAF-1"),
      SshKeyIds: pulumi.StringArray{
        pulumi.String("my-ssh-key-id"),
      },
      Type: pulumi.String("vcpu-4_memory-12g_disk-80g_nvidia3080-1"),
    })
    if err != nil {
      return err
    }
    return nil
  })
}
```

{{% /choosable %}}
{{% choosable language csharp %}}

```csharp
using Pulumi;
using GenesisCloud.PulumiPackage.Genesiscloud;
using System.Threading.Tasks;

class GenesisCloudInstance : Stack
{
    public GenesisCloudInstance()
    {
        var instance = new Instance("my-pulumi-instance", new InstanceArgs
        {
            Name = "my-pulumi-instance",
            Region = region,
            Image = "ubuntu-ml-nvidia-pytorch",
            Type = "vcpu-4_memory-12g_disk-80g_nvidia3080-1",
            SshKeyIds = { "my-ssh-key-id" },
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
