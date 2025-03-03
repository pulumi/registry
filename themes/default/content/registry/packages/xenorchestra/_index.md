---
# WARNING: this file was fetched from https://raw.githubusercontent.com/vatesfr/pulumi-xenorchestra/v1.4.0/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Xen Orchestra
meta_desc: Provides an overview of the Xen Orchestra Provider for Pulumi.
layout: overview
---

The Xen Orchestra Provider lets you manage [Xen Orchestra](https://github.com/vatesfr/xen-orchestra) resources.

## Information

Note that the Xen Orchestra Pulumi provider is a based on the [Terraform Xen Orchestra Provider](https://github.com/vatesfr/terraform-provider-xenorchestra)

# Examples

{{< chooser language "python,go,typescript" >}}
{{% choosable language python %}}

In your `__main__.py` set up the Xen Orchestra provider for pulumi in the following way:
```
import pulumi_xenorchestra as xoa

xoa_provider = xoa.Provider(
    "xenorchestra",
    xoa.ProviderArgs(
        url=Output.secret(secret_data["xoa_url"]),
        token=Output.secret(secret_data["xoa_token"]),
        insecure=False,
)
```

You can then une the provider to create a Virtual Machine:

```
xoa.Vm(
    resource_name=machine.name,
    name_label=machine.name,
    name_description="XOA VM",
    cpus=4,
    memory_max=8590000000,
    template=template.id,
    tags=["pulumi"],
    cloud_config=pathlib.Path("./config/cloudinit-xen-static.yaml").read_text()
    ),
    disks=[
        xoa.VmDiskArgs(
            sr_id=machine.storage_repository_id,
            name_label="OS",
            size=34361835520,
        ),
    ],
    networks=[
        xoa.VmNetworkArgs(
            network_id=network.id,
            mac_address=map_ip_to_mac(
                machine.ipv4_address,
                separator=":",
            ),
        ),
    ],
    opts=pulumi.ResourceOptions(
        provider=xoa_provider,
        # Do not recreate node with new credentials
        ignore_changes=["cloud_config"],
        parent=self,
    ),
)
```

{{% /choosable %}}

{{% choosable language go %}}

This examples uses `pulumi config set xenorchestra:token --secret` and `pulumi config set xenorchestra:url <url>` to configure the provider.


```go
package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	xoa "github.com/vatesfr/pulumi-xenorchestra/sdk/go/xenorchestra"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		poolId := "f4b3b3b3-3b3b-3b3b-3b3b-3b3b3b3b3b3b"
		templateId := "f4b3b3b3-3b3b-3b3b-3b3b-3b3b3b3b3b3b"

		net, err := xoa.LookupXoaNetwork(ctx, &xoa.LookupXoaNetworkArgs{
			NameLabel: "VLAN11",
			PoolId:    &poolId,
		})
		if err != nil {
			return err
		}

		vm, err := xoa.NewVm(ctx, "Pulumin go example", &xoa.VmArgs{
			NameLabel: pulumi.String("Pulumi go example"),
			MemoryMax: pulumi.Int(1073733632),
			Cpus:      pulumi.Int(1),
			Template:  pulumi.String(templateId),
			CloudConfig: pulumi.String(`
#cloud-config
ssh_authorized_keys:
  - ....
			`),
			Disks: xoa.VmDiskArray{
				xoa.VmDiskArgs{
					SrId:      pulumi.String("f4b3b3b3-3b3b-3b3b-3b3b-3b3b3b3b3b3b"),
					NameLabel: pulumi.String("OS"),
					Size:      pulumi.Int(4294967296),
				},
			},
			Networks: xoa.VmNetworkArray{
				xoa.VmNetworkArgs{
					NetworkId: pulumi.String(net.Id),
				},
			},
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

    const vm = new xoa.Vm("Pulumi example", {
        nameLabel: "Pulumi example",
        nameDescription: "Example with pulumi typescript provider",
        tags: ["pulumi"],
        cpus: 1,
        memoryMax: 1073733632,
        template: "template-id",
        cloudConfig: `
#cloud-config
ssh_authorized_keys:
    - ...
`,
        disks: [
            {
                nameLabel: "OS",
                size: 4294967296,
                srId: "local-storage-id",
            }
        ],
        networks: [
            {
                networkId: "network-id",
            }
        ],
        powerState: "Running"
    })
```

{{% /choosable %}}

{{< /chooser >}}
