---
title: libvirt
meta_desc: Provides an overview of the libvirt Provider for Pulumi.
layout: overview
---

{{% overview-description %}}
A Pulumi provider that lets you provision servers on a libvirt host using Pulumi.
{{% /overview-description %}}

## Example

{{< chooser language "javascript,typescript,python,go,csharp" >}}

{{% choosable language javascript %}}

```javascript
const libvirt = require("@pulumi/libvirt")

const myPool = new libvirt.Pool("test", {
  type: "dir",
  path: "/home/user/cluster_storage"
})
```

{{% /choosable %}}
{{% choosable language typescript %}}

```typescript
import * as libvirt from "@pulumi/libvirt";

const myPool = new libvirt.Pool("test", {
  type: "dir",
  path: "/home/user/cluster_storage"
})
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_libvirt as libvirt

pool = libvirt.Pool("demo-py-pool", type="dir", path="/home/user/cluster_storage")
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
  "github.com/pulumi/pulumi-libvirt/sdk/go/libvirt"
  "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
  pulumi.Run(func(ctx *pulumi.Context) error {
    pool, err := libvirt.NewPool(ctx, "cluster", &libvirt.PoolArgs{
        Type: pulumi.String("dir"),
        Path: pulumi.String("/home/user/cluster_storage"),
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
using System.Collections.Generic;
using System.Threading.Tasks;
using Pulumi;
using Pulumi.Libvirt;

class Program
{
    static Task Main() =>
        Deployment.Run(() => {
          var cluster = new Libvirt.Pool("cluster", new Libvirt.PoolArgs
          {
              Type = "dir",
              Path = "/home/user/cluster_storage",
          });
        });
}
```

{{% /choosable %}}

{{< /chooser >}}
