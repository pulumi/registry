---
title: Matchbox
meta_desc: Provides an overview of the Matchbox Provider for Pulumi.
layout: package
---

The Matchbox provider for Pulumi can be used to provision the [Matchbox](https://matchbox.psdn.io) iPXE server.
The Matchbox provider must be configured with certificates to connect correctly to the server.

## Example

{{< chooser language "typescript" >}}
{{% choosable language typescript %}}

```typescript
import fs from 'fs';

import * as pulumi from "@pulumi/pulumi"
import * as matchbox from "@pulumiverse/matchbox"

const matchboxConfig = new pulumi.Config("matchbox")
const matchBoxEndpoint = matchboxConfig.require("endpoint")

const config = new pulumi.Config()
// Set stack config item "stream" to a value applicable for the os releases, e,g:
// "Fedora CoreOS release stream (e.g. testing, stable)"
const osStream = config.get("stream") || "stable"

// Set stack config item "versioin" to a value applicable for the os version, e,g:
// "Fedora CoreOS version to PXE and install (e.g. 32.20200715.3.0)"
const osVersion = config.require("version")

const ignitionData = fs.readFileSync('worker.yaml', 'utf8');

const pxeProfile = new matchbox.Profile("pxeProfile", {
    initrds: [
        pulumi.interpolate`https://builds.coreos.fedoraproject.org/prod/streams/${osStream}/builds/${osVersion}/x86_64/fedora-coreos-${osVersion}-live-initramfs.x86_64.img`
    ],
    args: [
        "ip=dhcp",
        "rd.neednet=1",
        pulumi.interpolate`initrd=fedora-coreos-${osVersion}-live-initramfs.x86_64.img`,
        pulumi.interpolate`coreos.inst.image_url=https://builds.coreos.fedoraproject.org/prod/streams/${osStream}/builds/${osVersion}/x86_64/fedora-coreos-${osVersion}-metal.x86_64.raw.xz`,
        pulumi.interpolate`coreos.inst.ignition_url=${matchBoxEndpoint}/ignition?uuid=\${uuid}&mac=\${mac:hexhyp}`,
        "coreos.inst.install_dev=sda",
        "console=tty0",
        "console=ttyS0",
    ],
    rawIgnition: ignitionData
})

const pxeGroup = new matchbox.Group("node1", {
    profile: pxeProfile.id,
    selector: {
        mac: "52:54:00:a1:9c:ae"
    },
    metadata: {
        customVariable: "machine_specific_value_here"
    }
})
```

{{% /choosable %}}

{{< /chooser >}}
