---
# WARNING: this file was fetched from https://raw.githubusercontent.com/shirasakaren/pulumi-biznetgio/v0.1.7/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/shirasakaren/pulumi-biznetgio/blob/v0.1.7/docs/_index.md
title: BiznetGIO
meta_desc: Provides an overview of the BiznetGIO Provider for Pulumi.
layout: package
---

Unofficial [Pulumi](https://www.pulumi.com) provider for [BiznetGIO](https://www.biznetgio.com), an Indonesian
cloud platform. It manages NEO Metal (bare metal), NEO Lite / NEO Lite Pro (VMs), NEO GPU, and NEO Object Storage
as code, backed by the [BiznetGIO Portal API](https://api.portal.biznetgio.com/v1/docs).

Not affiliated with or endorsed by PT Biznet Gio Nusantara. See the
[full docs site](https://biznetgio.creations.ren) for guides, every resource, and the complete API reference.

## Example

{{< chooser language "typescript,python,go,csharp,java" >}}
{{% choosable language typescript %}}

```typescript
import * as biznetgio from "@shirasakaren/biznetgio";

const keypair = new biznetgio.NeoliteKeypair("deploy", { name: "deploy-key" });

const vm = new biznetgio.NeoliteVm("web", {
    vmName: "web-1",
    productId: 123,
    selectOs: "Ubuntu 22.04",
    keypairId: keypair.keypairId,
    cycle: "m",
    consolePassword: "change-this-now!",
    sshAndConsoleUser: "root",
    // defaults to true: bills the card on file immediately.
    // set false to keep the order pending until paid manually in the portal.
    payWithCreditCard: true,
});

export const vmStatus = vm.status;
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi
import pulumi_biznetgio as biznetgio

keypair = biznetgio.NeoliteKeypair("deploy", name="deploy-key")

vm = biznetgio.NeoliteVm("web",
    vm_name="web-1",
    product_id=123,
    select_os="Ubuntu 22.04",
    keypair_id=keypair.keypair_id,
    cycle="m",
    console_password="change-this-now!",
    ssh_and_console_user="root",
    # defaults to True: bills the card on file immediately.
    # set False to keep the order pending until paid manually in the portal.
    pay_with_credit_card=True,
)

pulumi.export("vmStatus", vm.status)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	biznetgio "github.com/shirasakaren/pulumi-biznetgio/sdk/go/pulumi-biznetgio"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		keypair, err := biznetgio.NewNeoliteKeypair(ctx, "deploy", &biznetgio.NeoliteKeypairArgs{
			Name: pulumi.String("deploy-key"),
		})
		if err != nil {
			return err
		}

		vm, err := biznetgio.NewNeoliteVm(ctx, "web", &biznetgio.NeoliteVmArgs{
			VmName:            pulumi.String("web-1"),
			ProductId:         pulumi.Int(123),
			SelectOs:          pulumi.String("Ubuntu 22.04"),
			KeypairId:         keypair.KeypairId,
			Cycle:             pulumi.String("m"),
			ConsolePassword:   pulumi.String("change-this-now!"),
			SshAndConsoleUser: pulumi.String("root"),
			// defaults to true: bills the card on file immediately.
			// set false to keep the order pending until paid manually in the portal.
			PayWithCreditCard: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}

		ctx.Export("vmStatus", vm.Status)
		return nil
	})
}
```

{{% /choosable %}}
{{% choosable language csharp %}}

```csharp
using System.Collections.Generic;
using System.Threading.Tasks;
using Pulumi;
using Shirasakaren.Biznetgio;

return await Deployment.RunAsync(() =>
{
    var keypair = new NeoliteKeypair("deploy", new NeoliteKeypairArgs
    {
        Name = "deploy-key",
    });

    var vm = new NeoliteVm("web", new NeoliteVmArgs
    {
        VmName = "web-1",
        ProductId = 123,
        SelectOs = "Ubuntu 22.04",
        KeypairId = keypair.KeypairId,
        Cycle = "m",
        ConsolePassword = "change-this-now!",
        SshAndConsoleUser = "root",
        // defaults to true: bills the card on file immediately.
        // set false to keep the order pending until paid manually in the portal.
        PayWithCreditCard = true,
    });

    return new Dictionary<string, object?>
    {
        ["vmStatus"] = vm.Status,
    };
});
```

{{% /choosable %}}
{{% choosable language java %}}

```java
package myproject;

import com.pulumi.Pulumi;
import ren.shirasaka.biznetgio.NeoliteKeypair;
import ren.shirasaka.biznetgio.NeoliteKeypairArgs;
import ren.shirasaka.biznetgio.NeoliteVm;
import ren.shirasaka.biznetgio.NeoliteVmArgs;

public class App {
    public static void main(String[] args) {
        Pulumi.run(ctx -> {
            var keypair = new NeoliteKeypair("deploy", NeoliteKeypairArgs.builder()
                    .name("deploy-key")
                    .build());

            var vm = new NeoliteVm("web", NeoliteVmArgs.builder()
                    .vmName("web-1")
                    .productId(123)
                    .selectOs("Ubuntu 22.04")
                    .keypairId(keypair.keypairId())
                    .cycle("m")
                    .consolePassword("change-this-now!")
                    .sshAndConsoleUser("root")
                    // defaults to true: bills the card on file immediately.
                    // set false to keep the order pending until paid manually in the portal.
                    .payWithCreditCard(true)
                    .build());

            ctx.export("vmStatus", vm.status());
        });
    }
}
```

{{% /choosable %}}
{{< /chooser >}}

> **Billing note**: `payWithCreditCard` defaults to `true`, so the first `pulumi up` places a real order and may
> charge the credit card on file. Set it to `false` to leave the order `Pending` until it's paid manually in the
> portal.

## What's covered

- **NEO Metal** - bare metal servers, keypairs, additional IPs, elastic storage.
- **NEO Lite / NEO Lite Pro** - VMs, keypairs, snapshots, extra disks, and one-way Lite-to-Pro migration.
- **NEO GPU** - GPU instances (subscription or on-demand billing) and keypairs.
- **NEO Object Storage** - S3-compatible storage instances, buckets, credentials, and objects.

Every resource also exposes a secret-marked `raw` output with the full last-read API response, so you're never
blocked on a field the provider hasn't modeled yet.
