---
# WARNING: this file was fetched from https://raw.githubusercontent.com/lbrlabs/pulumi-lbrlabs-eks/v1.3.4/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
title: LBr Labs EKS
meta_desc: Information on how to install the LBr Labs EKS component.
layout: package
---

## Installation

The Pulumi LBr Labs EKS provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@lbrlabs/pulumi-eks`](https://www.npmjs.com/package/@lbrlabs/pulumi-eks)
* Python: [`lbrlabs_pulumi_eks`](https://pypi.org/project/lbrlabs-pulumi-eks/)
* Go: [`github.com/lbrlabs/pulumi-eks/sdk/go/`](https://pkg.go.dev/github.com/lbrlabs/pulumi-lbrlabs-eks/sdk/go/eks)
* .NET: [`Lbrlabs.PulumiPackage.eks`](https://www.nuget.org/packages/Lbrlabs.PulumiPackage.Eks)

### Provider Binary

The provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource lbrlabs-eks <version> --server github://api.github.com/lbrlabs
```

Replace the version string with your desired version.

