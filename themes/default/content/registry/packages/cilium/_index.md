---
# WARNING: this file was fetched from https://raw.githubusercontent.com/littlejo/pulumi-cilium/v0.2.1/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
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
