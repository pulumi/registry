---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumiverse/pulumi-esxi-native/v1.0.0/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
title: ESXi Native
meta_desc: Provides an overview of the ESXi Native Provider for Pulumi.
layout: package
---

The ESXi Native provider is used to provision VMs directly on an ESXi hypervisor without a need for vCenter or vSphere.

## Example

{{< chooser language "typescript,python,go,csharp" >}}

{{% choosable language typescript %}}
```typescript
import * as esxi from "@pulumiverse/esxi-native";


export = async () => {
    const vm = new esxi.VirtualMachine("vm-test", {
        diskStore: "nvme-ssd-datastore",
        networkInterfaces: [
            {
                virtualNetwork: "default"
            }
        ]
    });

    return {
        "id": vm.id,
        "name": vm.name,
        "os": vm.os,
    };
}
```
{{% /choosable %}}

{{% choosable language python %}}
```python
import pulumi
from typing import Sequence
from pulumiverse_esxi_native import VirtualMachine, NetworkInterfaceArgs

vm = VirtualMachine("vm-test",
                    disk_store="nvme-ssd-datastore",
                    network_interfaces=Sequence[NetworkInterfaceArgs(
                        virtual_network="default"
                    )])

pulumi.export("id", vm.id)
pulumi.export("name", vm.name)
pulumi.export("os", vm.os)
```
{{% /choosable %}}

{{% choosable language go %}}
```go
package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-esxi-native/sdk/go/esxi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		vm, err := esxi.NewVirtualMachine(ctx, "vm-test", &esxi.VirtualMachineArgs{
			DiskStore:               pulumi.String("nvme-ssd-datastore"),
			NetworkInterfaces:       esxi.NetworkInterfaceArray{
				esxi.NetworkInterfaceArgs{
					VirtualNetwork: pulumi.String("default"),
				},
			},
		})
		if err != nil {
			return err
		}

		ctx.Export("id", vm.ID())
		ctx.Export("name", vm.Name)
		ctx.Export("os", vm.Os)
		return nil
	})
}
```
{{% /choosable %}}

{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using Pulumi;
using Pulumiverse.EsxiNative;
using Pulumiverse.EsxiNative.Inputs;

return await Deployment.RunAsync(() =>
{
    var vm = new VirtualMachine("vm-test", new VirtualMachineArgs
    {
        DiskStore = "nvme-ssd-datastore",
        NetworkInterfaces = new NetworkInterfaceArgs[]
        {
            new ()
            {
                VirtualNetwork = "default"
            }
        }
    });

    return new Dictionary<string, object?>
    {
        ["id"] = vm.Id,
        ["name"] = vm.Name,
        ["os"] = vm.Os,
    };
});

```
{{% /choosable %}}

{{< /chooser >}}
