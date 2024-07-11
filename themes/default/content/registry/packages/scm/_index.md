---
title: Strata Cloud Manager
meta_desc: Provides an overview of the Strata Cloud Manager Provider for Pulumi.
layout: overview
---

The Strata Cloud Manager Resource Provider lets you manage [Strata Cloud Manager](https://docs.paloaltonetworks.com/strata-cloud-manager) resources. This is a bridged provider from the terraform SCM provider, located at [github.com/PaloAltoNetworks/terraform-provider-scm](https://github.com/PaloAltoNetworks/terraform-provider-scm).

{{% notes %}}
Please note that, due to how Strata Cloud Manager works, any changes made using the Strata Cloud Manager Resource provider will only be applied to the candidate configuration. For these changes to take effect, you must manually push the candidate configuration to become the running configuration. This can be done either through the Strata Cloud Manager GUI or via the [Strata Cloud Manager API](https://pan.dev/access/api/prisma-access-config/post-sse-config-v-1-config-versions-candidate-push/). 

Please see [Strata Cloud Manager Manage: Push Config](https://docs.paloaltonetworks.com/strata-cloud-manager/getting-started/operations/push-config) for more details.
{{% /notes %}}

## Example

{{< chooser language "typescript,python,go,csharp" >}}
{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as scm from "@pulumi/scm";

const example = new scm.Service("example", {});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi
import pulumi_scm as scm

example = scm.Service("example")
```

{{% /choosable %}}
{{% choosable language go %}}

```go
package main

import (
	"github.com/pulumi/pulumi-scm/sdk/go/scm"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

)

func main() {
    pulumi.Run(func(ctx *pulumi.Context) error {
        _, err := scm.NewService(ctx, "example", nil)
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
using Pulumi;
using Scm = Pulumi.Scm;

return await Deployment.RunAsync(() =>
{
    var example = new Scm.Service("example");
});
```

{{% /choosable %}}

{{< /chooser >}}
