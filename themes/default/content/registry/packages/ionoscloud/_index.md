---
# WARNING: this file was fetched from https://raw.githubusercontent.com/ionos-cloud/pulumi-ionoscloud/v0.2.2/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Ionoscloud
meta_desc: Provides an overview of the Ionoscloud Provider for Pulumi.
layout: overview
---

The `Ionoscloud` provider for Pulumi can be used to provision most of the resources available in the Ionoscloud Terraform provider [Ionoscloud Terraform provider](https://github.com/ionos-cloud/terraform-provider-ionoscloud).

The provider needs to be configured with proper credentials before it can be used.

We strongly suggest to use token authentication for security purposes. You can find more information [here](https://docs.ionos.com/cloud/set-up-ionos-cloud/management/token-management).

## Example

{{< chooser language "go,typescript,python,csharp" >}}
{{% choosable language go %}}

```golang
_, err := compute.NewDatacenter(ctx, "example", &compute.DatacenterArgs{
    Name:              pulumi.String("Datacenter Example"),
    Location:          pulumi.String("us/las"),
    Description:       pulumi.String("datacenter description"),
    SecAuthProtection: pulumi.Bool(false),
})
if err != nil {
    return err
}
```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi
import pulumi_ionoscloud as ionoscloud

example = ionoscloud.compute.Datacenter("example",
    name="Datacenter Example",
    location="us/las",
    description="datacenter description",
    sec_auth_protection=False
)
```

{{% /choosable %}}

{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as ionoscloud from "@ionos-cloud/sdk-pulumi";

const example = new ionoscloud.compute.Datacenter("example", {
    name: "Datacenter Example",
    location: "us/las",
    description: "datacenter description",
    secAuthProtection: false,
});
```

{{% /choosable %}}

{{% choosable language csharp %}}

```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Ionoscloud = Ionoscloud.Pulumi.Ionoscloud;

return await Deployment.RunAsync(() =&gt; 
{
    var example = new Ionoscloud.Compute.Datacenter("example", new()
    {
        Name = "Datacenter Example",
        Location = "us/las",
        Description = "datacenter description",
        SecAuthProtection = false,
    });

});
```

{{% /choosable %}}
{{< /chooser >}}
