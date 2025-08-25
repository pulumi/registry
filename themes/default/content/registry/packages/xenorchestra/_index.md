---
# WARNING: this file was fetched from https://raw.githubusercontent.com/vatesfr/pulumi-xenorchestra/v2.2.0/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Xen Orchestra
meta_desc: Provides an overview of the Xen Orchestra Provider for Pulumi.
layout: package
---

The Xen Orchestra Provider lets you manage [Xen Orchestra](https://github.com/vatesfr/xen-orchestra) resources.

## Information

Note that the Xen Orchestra Pulumi provider is a based on the [Terraform Xen Orchestra Provider](https://github.com/vatesfr/terraform-provider-xenorchestra)

# Examples

Those examples uses `pulumi config set xenorchestra:token --secret` and `pulumi config set xenorchestra:url <url>` to configure the provider.

{{< chooser language "python,go,typescript,csharp,yaml" >}}
{{% choosable language python %}}

```python
import pulumi
import pulumi_xenorchestra as xoa
import pathlib

template = xoa.get_xoa_template(
    name_label="Debian 12 Cloud-init (Hub)",
    pool_id=pool.id
)

xoa.Vm(
    resource_name="Pulumi example",
    name_label="Pulumi example",
    name_description="Example with pulumi python provider",
    tags=["pulumi"],
    cpus=1,
    memory_max=1073733632,
    template=template.id,
    cloud_config=pathlib.Path("./config/cloudinit-xen-static.yaml").read_text(),
    disks=[
        xoa.VmDiskArgs(
            name_label="OS",
            size=4294967296,
            sr_id="sr-id"
        ),
    ],
    networks=[
        xoa.VmNetworkArgs(
            network_id="network-id",
        ),
    ],
    power_state="Running",
)

pulumi.export("vm_ip", vm.ipv4_addresses)
```

{{% /choosable %}}

{{% choosable language go %}}

```go
package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	xoa "github.com/vatesfr/pulumi-xenorchestra/sdk/go/xenorchestra"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		poolId := "pool-id"

        template, err := xoa.GetXoaTemplate(ctx, &xoa.GetXoaTemplateArgs{
			NameLabel: "Debian 12 Cloud-init (Hub)",
			PoolId:    &poolId,
		})
		if err != nil {
			return err
		}

		vm, err := xoa.NewVm(ctx, "Pulumi example", &xoa.VmArgs{
			NameLabel:       pulumi.String("Pulumi example"),
            NameDescription: pulumi.String("Example with pulumi golang provider"),
            Tags:            pulumi.StringArray{pulumi.String("pulumi")},
			Cpus:            pulumi.Int(1),
			MemoryMax:       pulumi.Float64(1073733632),
			Template:        pulumi.String(template.Id),
			CloudConfig: pulumi.String(`
#cloud-config
ssh_authorized_keys:
  - ....
			`),
			Disks: xoa.VmDiskArray{
				xoa.VmDiskArgs{
					NameLabel: pulumi.String("OS"),
					Size:      pulumi.Float64(4294967296),
					SrId:      pulumi.String("sr-id"),
				},
			},
			Networks: xoa.VmNetworkArray{
				xoa.VmNetworkArgs{
					NetworkId: pulumi.String(net.Id),
				},
			},
			PowerState: pulumi.String("Running"),
		})
		if err != nil {
			return err
		}

		ctx.Export("vp_ip", vm.Ipv4Addresses)
		return nil
	})
}

```

{{% /choosable %}}

{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as xoa from "@pulumi/xenorchestra"

async function getPrerequisites(){
    const template = await xoa.getXoaTemplate({
        nameLabel: "Debian 12 Cloud-init (Hub)",
        poolId: "pool-id"
    })

    return {
        template: template
    }
}

export const out = getPrerequisites().then(prerequisites => {
    const vm = new xoa.Vm("Pulumi example", {
        nameLabel: "Pulumi example",
        nameDescription: "Example with pulumi typescript provider",
        tags: ["pulumi"],
        cpus: 1,
        memoryMax: 1073733632,
        template: prerequisites.template.id,
        cloudConfig: `
#cloud-config
ssh_authorized_keys:
    - ...
`,
        disks: [
            {
                nameLabel: "OS",
                size: 4294967296,
                srId: "sr-id",
            }
        ],
        networks: [
            {
                networkId: "network-id",
            }
        ],
        powerState: "Running"
    })
    return vm.ipv4Addresses
})
```

{{% /choosable %}}

{{% choosable language csharp %}}

```csharp
using System.Collections.Generic;
using Pulumi;
using Pulumi.Xenorchestra;
using Pulumi.Xenorchestra.Inputs;

return await Deployment.RunAsync(() =>
{    
    var template = GetXoaTemplate.Invoke(new GetXoaTemplateInvokeArgs
    {
        NameLabel = "Debian 12 Cloud-init (Hub)",
        PoolId = "pool-id"
    });
    var vm = new Vm("vm", new VmArgs
    {
        NameLabel = "Pulumi example",
        NameDescription = "Example with pulumi dotnet provider",
        Tags = ["pulumi"],
        Cpus = 1,
        MemoryMax = 1073733632,
        Template = template.Apply(getXoaTemplateResult => getXoaTemplateResult.Id),
        CloudConfig = """
                    #cloud-config
                    ssh_authorized_keys:
                        - ...
                    """
        Disks = new VmDiskArgs[] {
            new VmDiskArgs {
                NameLabel = "OS",
                Size = 4294967296,
                SrId = "sr-id"
            }
        },
        Networks = new VmNetworkArgs[] {
            new VmNetworkArgs {
                NetworkId = "network-id"
            }
        },
        PowerState = "Running"
    });

    return new Dictionary<string, object?>
    {
        ["vp_ip"] = vm.Ipv4Addresses.Apply(ipv4Addresses => ipv4Addresses)
    };
});
```
{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
name: test-yaml
description: A minimal Pulumi YAML program
runtime: yaml
config: {'pulumi:tags': {value: {'pulumi:template': yaml}}}
variables: 
  templateId:
    fn::invoke:
      function: xenorchestra:getXoaTemplate
      arguments:
        nameLabel: "Debian 12 Cloud-init (Hub)"
        poolId: "pool-id"
      return: id
resources:
  vm:
    type: xenorchestra:Vm
    properties:
      nameLabel: "Pulumi example"
      nameDescription: "Example with pulumi yam provider"
      tags:
        - pulumi
      cpus: 1
      memoryMax: 1073733632
      template: ${templateId}
      cloudConfig: |
        #cloud-config
        ssh_authorized_keys:
            - ...
      disks:
        - nameLabel: "OS"
          size: 4294967296
          srId: "sr-id"
      networks:
        - networkId: "network-id"
      powerState: "Running"

outputs: 
  vmIp: ${vm.ipv4Addresses}
```

{{% /choosable %}}

{{< /chooser >}}
