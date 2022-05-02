---
title: Twingate
meta_desc: Provides an overview of the Twingate Provider for Pulumi.
layout: overview
---

The Twingate provider for Pulumi can be used to provision any of the cloud resources available in [Twingate](https://www.twingate.com/).
The Twingate provider must be configured with credentials to deploy and update resources in Scaleway.

## Example

{{< chooser language "typescript,python,csharp" >}}
{{% choosable language typescript %}}

```typescript
import * as twingate from "@twingate-labs/pulumi-twingate";

const remoteNetwork = new twingate.TwingateRemoteNetwork("test-network")

new twingate.TwingateResource("test-resource", {
    name: "Pulumi Website",
    address: "www.pulumi.com",
    remoteNetworkId: remoteNetwork.id,
    groupIds: ["R3JvdXA6MzA2MDk="]
})
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_twingate as twingate

remote_network = twingate.TwingateRemoteNetwork("test-network")

twingate_resource = twingate.TwingateResource("example",
    name="Pulumi Website",
    address="www.pulumi.com",
    remote_network_id=remote_network.id,
    group_ids=["R3JvdXA6MzA2MDk="]
)
```

{{% /choosable %}}

{{% choosable language csharp %}}

```csharp
using Pulumi;
using Twingate = TwingateLabs.Twingate;

class TwingateStack : Stack
{
    var remoteNetwork = new TwingateRemoteNetwork("example", new TwingateRemoteNetworkArgs{});

    var twingateResource = new TwingateResource("example", new InstanceServerArgs{
        Name = "Pulumi Website",
        Address = "www.pulumi.com",
        RemoteNetworkId = remoteNetwork.Id,
        GroupIds =
        {
            "R3JvdXA6MzA2MDk=",
        },
    });
}
```

{{% /choosable %}}

{{< /chooser >}}
