---
title: OVH
meta_desc: Provides an overview of the OVH Provider for Pulumi.
layout: overview
---

The `OVH` provider for Pulumi can be used to provision any of the resources available in [OVHcloud](https://www.ovhcloud.com/fr/).
The `OVH` provider must be configured with credentials to deploy and update resources in `OVH`.

## Information

Note that the [lbrlabs Pulumi OVH provider](https://github.com/lbrlabs/pulumi-ovh) is replaced by this official one. See the
[Migration Guide](./how-to-guides/migration-from-lbrlabs-package.md) for more information.

## Example

{{< chooser language "go,typescript,python,csharp" >}}
{{% choosable language go %}}

```golang
// Deploy a new Kubernetes cluster
myKube, err := cloudproject.NewKube(ctx, "my_desired_cluster", &cloudproject.KubeArgs{
    ServiceName: pulumi.String("xxxxxxxxxxxxx-xxxx-xxxx-xxxxxxxxx"),
    Name:        pulumi.String("my_desired_cluster"),
    Region:      pulumi.String("GRA5"),
})
if err != nil {
    return err
}

// Export kubeconfig file to a secret
ctx.Export("kubeconfig", pulumi.ToSecret(myKube.Kubeconfig))
```

{{% choosable language python %}}

```python
"""Get a Kubernetes cluster version"""
import pulumi
import pulumi_ovh as ovh

config = pulumi.Config();
service_name = config.require('serviceName')

print(service_name);

my_kube_cluster = ovh.cloudproject.get_kube(service_name=service_name,
    kube_id="xxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx");
pulumi.export("version", my_kube_cluster.version)
```

{{% /choosable %}}

{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";

import * as ovh from "@ovh-devrelteam/pulumi-ovh"

let config = new pulumi.Config();
let serviceName = config.require("serviceName")

console.log(serviceName)

// Get a Kubernetes cluster version
let myKubeCluster = ovh.cloudproject.getKube({
    serviceName: serviceName,
    kubeId: "xxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxx"
}) 

export const version = myKubeCluster.then(myKubeCluster => myKubeCluster.version);
```

{{% /choosable %}}

{{% choosable language csharp %}}

```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Ovh = Pulumi.Ovh;
using System;

return await Deployment.RunAsync(() => 
{

    var config = new Pulumi.Config();
    var serviceName = config.Require("serviceName");

    System.Console.WriteLine(serviceName);

    var myKubeCluster = Ovh.CloudProject.GetKube.Invoke(new()
    {
        ServiceName = serviceName,
        KubeId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxx",
    });

    return new Dictionary<string, object?>
    {
        ["version"] = myKubeCluster.Apply(getKubeResult => getKubeResult.Version),
    };
});
```

{{% /choosable %}}

{{< /chooser >}}
