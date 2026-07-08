---
# WARNING: this file was fetched from https://raw.githubusercontent.com/OlivierPaquien/pulumi-freebox/v0.3.11/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/OlivierPaquien/pulumi-freebox/blob/v0.3.11/docs/_index.md
title: Freebox
meta_desc: Manage Freebox home routers as infrastructure as code with Pulumi — port forwarding, VMs, VPN, LAN, DHCP, and more.
layout: package
---

The Freebox provider lets you manage a [Freebox](https://www.free.fr/freebox/) router through the [Freebox API](https://dev.freebox.fr/sdk/) as infrastructure as code. It is a native Pulumi provider implemented in Go with [pulumi-go-provider](https://github.com/pulumi/pulumi-go-provider), ported from [terraform-provider-freebox](https://github.com/NikolaLohinski/terraform-provider-freebox) without the Terraform Bridge.

## Resources

| Token | Description |
|-------|-------------|
| `freebox:fw:PortForwarding` | Port forwarding rule (NAT) |
| `freebox:virtual:Disk` | Virtual disk (qcow2/raw) on the Freebox |
| `freebox:virtual:Machine` | Virtual machine configuration and optional power state |
| `freebox:virtual:MachinePower` | VM power state only (`running` / `stopped`), independent of VM config |
| `freebox:dhcp:StaticLease` | DHCP static lease (reserve an IPv4 address for a MAC address) |
| `freebox:downloads:File` | File on the Freebox (download, copy, upload, extract) |
| `freebox:vpn:Server` | OpenVPN server configuration (singleton per Freebox) |
| `freebox:vpn:User` | OpenVPN user account |
| `freebox:lan:Config` | LAN configuration (singleton per Freebox) |

### Virtual machines: `Machine` vs `MachinePower`

- Use **`freebox:virtual:Machine`** when you manage VM configuration (disk, memory, vCPUs) and optionally power state via the `status` property (`running` or `stopped`, default `running`).
- Use **`freebox:virtual:MachinePower`** when you want to control power independently from configuration (for example start/stop on a schedule without changing VM settings).
- Do not manage the same VM with both `status` on `Machine` and a `MachinePower` resource — pick one approach per VM.

## Functions

| Token | Description |
|-------|-------------|
| `freebox:api:Version` | Freebox API discovery (version, model, etc.) |
| `freebox:virtual:getVirtualDisk` | Virtual disk information (type, sizes) |
| `freebox:dhcp:getLease` | DHCP static lease by MAC address |
| `freebox:dhcp:getLeases` | List all DHCP static leases |
| `freebox:lan:getConfig` | Read LAN configuration |
| `freebox:lan:getInterfaces` | List LAN interfaces |
| `freebox:lan:getInterfaceHosts` | List hosts on a LAN interface |
| `freebox:lan:getInterfaceHost` | Read a single LAN host |
| `freebox:virtual:getDistributions` | VM OS distributions available on the Freebox |
| `freebox:system:getInfo` | Freebox system information |

## Example

The example below creates a TCP port forwarding rule that exposes SSH on the WAN side and forwards traffic to a host on the LAN.

{{< chooser language "typescript,python,csharp,go,yaml" >}}

{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as freebox from "pulumi-freebox";

const pf = new freebox.PortForwarding("ssh", {
    enabled: true,
    ipProtocol: "tcp",
    portRangeStart: 22,
    targetIp: "192.168.1.10",
    comment: "SSH",
});
```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi
import pulumi_freebox as freebox

pf = freebox.PortForwarding(
    "ssh",
    enabled=True,
    ip_protocol="tcp",
    port_range_start=22,
    target_ip="192.168.1.10",
    comment="SSH",
)
```

{{% /choosable %}}

{{% choosable language csharp %}}

```csharp
using Pulumi;
using OlivierPaquien.Pulumi.Freebox;

var pf = new PortForwarding("ssh", new PortForwardingArgs
{
    Enabled = true,
    IpProtocol = "tcp",
    PortRangeStart = 22,
    TargetIp = "192.168.1.10",
    Comment = "SSH",
});
```

{{% /choosable %}}

{{% choosable language go %}}

```go
package main

import (
	"github.com/OlivierPaquien/pulumi-freebox/sdk/go/freebox"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := freebox.NewPortForwarding(ctx, "ssh", &freebox.PortForwardingArgs{
			Enabled:        pulumi.Bool(true),
			IpProtocol:     pulumi.String("tcp"),
			PortRangeStart: pulumi.Int(22),
			TargetIp:       pulumi.String("192.168.1.10"),
			Comment:        pulumi.String("SSH"),
		})
		return err
	})
}
```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
name: freebox-example
runtime: yaml
config:
  freebox:endpoint: http://mafreebox.freebox.fr
  freebox:appId: "your_app_id"
  freebox:token:
    secret: your_token

resources:
  ssh:
    type: freebox:fw:PortForwarding
    properties:
      enabled: true
      ipProtocol: tcp
      portRangeStart: 22
      targetIp: "192.168.1.10"
      comment: SSH
```

{{% /choosable %}}

{{< /chooser >}}

## Source

- Provider repository: [github.com/OlivierPaquien/pulumi-freebox](https://github.com/OlivierPaquien/pulumi-freebox)
- Freebox API client: [github.com/NikolaLohinski/free-go](https://github.com/NikolaLohinski/free-go)
