---
title: Cilium
meta_desc: Provides an overview of the Cilium Provider for Pulumi.
layout: package
---

The Cilium provider for Pulumi can be used to provision any of the cloud resources available in Cilium.

The Cilium provider must be configured with credentials to deploy and update resources in Cilium.

## Example

{{< chooser language "python,go" />}}

{{% choosable language python %}}

```python
"""Deploy cilium"""
example_install = cilium.Install("exampleInstall",
    sets=[
        "ipam.mode=kubernetes",
        "ipam.operator.replicas=1",
        "tunnel=vxlan",
    ],
    version="1.14.5")
```

{{% /choosable %}}

{{% choosable language go %}}

```golang
// Deploy cilium
_, err = cilium.NewInstall(ctx, "exampleInstall", &cilium.InstallArgs{
        Sets: pulumi.StringArray{
                pulumi.String("ipam.mode=kubernetes"),
                pulumi.String("ipam.operator.replicas=1"),
                pulumi.String("tunnel=vxlan"),
        },
        Version: pulumi.String("1.14.5"),
})
if err != nil {
        return err
}
return nil
```

{{% /choosable %}}
