---
title: Vultr
meta_desc: Provides an overview of the Vultr Provider for Pulumi.
layout: package
---

The Vultr Resource Provider lets you manage [Vultr](https://www.vultr.com/) resources.

## Example

{{< chooser language "javascript,typescript,python,go,csharp" >}}


{{% choosable language javascript %}}

```javascript
"use strict";
const vultr = require("@ediri/vultr");


const vke = new vultr.Kubernetes("vke", {
    region: "fra",
    version: "v1.25.4+1",
    label: "pulumi-vultr",
    nodePools: {
        nodeQuantity: 1,
        plan: "vc2-2c-4gb",
        label: "pulumi-vultr-nodepool",
    },
})

exports.kubeConfig = vke.kubeConfig
```

{{% /choosable %}}

{{% choosable language typescript %}}

```typescript
import * as vultr from "@ediri/vultr";

let vke = new vultr.Kubernetes("vke", {
    region: "fra",
    version: "v1.25.4+1",
    label: "pulumi-vultr",
    nodePools: {
        nodeQuantity: 1,
        plan: "vc2-2c-4gb",
        label: "pulumi-vultr-nodepool",
    },
})

export const kubeconfig = vke.kubeConfig;
```

{{% /choosable %}}

{{% choosable language python %}}

```python
from pulumi import export
import ediri_vultr as vultr

vke = vultr.Kubernetes('vke', region='fra', version='v1.25.4+1',
                       label='pulumi-vke',
                       node_pools=vultr.KubernetesNodePoolsArgs(node_quantity=1,
                                                                plan='vc2-2c-4gb', label='pulumi-vultr-nodepool'))

export('kubeConfig', vke.kube_config)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
package main

import (
	"github.com/dirien/pulumi-vultr/v2/go/vultr"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		vke, err := vultr.NewKubernetes(ctx, "vke", &vultr.KubernetesArgs{
			Region:  pulumi.String("fra"),
			Version: pulumi.String("v1.25.4+1"),
			Label:   pulumi.String("pulumi-vultr"),
			NodePools: vultr.KubernetesNodePoolsTypeArgs{
				NodeQuantity: pulumi.Int(1),
				Plan:         pulumi.String("vc2-2c-4gb"),
				Label:        pulumi.String("pulumi-vultr-nodepool"),
			},
		})
		if err != nil {
			return err
		}
		ctx.Export("kubeconfig", vke.KubeConfig)
		return nil
	})
}
```

{{% /choosable %}}

{{% choosable language csharp %}}

```csharp
using System.Collections.Generic;
using Pulumi;
using ediri.Vultr; 

return await Deployment.RunAsync(() =>
{
   var vke = new Kubernetes("vke", new KubernetesArgs{
      Region = "fra",
      Version = "v1.25.4+1",
      Label = "pulumi-vultr",
      NodePools = new ediri.Vultr.Inputs.KubernetesNodePoolsArgs{
           NodeQuantity = 1,
           Plan = "vc2-2c-4gb",
           Label = "pulumi-vultr",
         },
      }
   );
   return new Dictionary<string, object?>
   {
      ["kubeConfig"] = vke.KubeConfig,
   };
});
```

{{% /choosable %}}

{{< /chooser >}}
