---
title: Libvirt Setup
meta_desc: Provides an overview on how to setup the libvert Provider for Pulumi.
layout: installation
---

### Requirements

Please note, there is a requirement to have [libvirt](https://libvirt.org/) on the machine using the libvirt provider. You can
install this as follows:

{{< chooser os "linux,macos,windows" >}}

{{% choosable os linux %}}

* Refer to [Libvirt downloads](https://libvirt.org/downloads.html)

{{% /choosable %}}

{{% choosable os macos %}}

```bash
$ brew install libvirt
```

{{% /choosable %}}

{{% choosable os windows %}}

* Refer to [Libvirt Windows Instructions](https://libvirt.org/windows.html)

{{% /choosable %}}
{{< /chooser >}}

## Installation

The Libvirt provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/libvirt`](https://www.npmjs.com/package/@pulumi/libvirt)
* Python: [`pulumi-libvirt`](https://pypi.org/project/pulumi-libvirt/)
* Go: [`github.com/pulumi/pulumi-libvirt/sdk/go/libvirt`](https://github.com/pulumi/pulumi-libvirt)
* .NET: [`Pulumi.Libvirt`](https://www.nuget.org/packages/Pulumi.Libvirt)
