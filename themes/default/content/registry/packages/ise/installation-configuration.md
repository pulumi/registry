---
title: Ise Installation & Configuration
meta_desc: Information on how to install the Ise provider.
layout: installation
---

## Installation

The Pulumi Ise provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/ise`](https://www.npmjs.com/package/@pulumi/ise)
* Python: [`pulumi-ise`](https://pypi.org/project/pulumi-ise/)
* Go: [`github.com/pulumi/pulumi-ise/sdk/go/ise`](https://pkg.go.dev/github.com/pulumi/pulumi-ise/sdk/go/ise)
* .NET: [`Pulumi.Ise`](https://www.nuget.org/packages/Pulumi.Ise)


## Configuration

TODO

### Provider Binary

The Ise provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource ise <version>
```

Replace the version string `<version>` with your desired version.
