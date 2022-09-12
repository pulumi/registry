---
title: Ovh Setup
meta_desc: Information on how to install the Pulumi Ovh provider.
layout: installation
---

## Installation

The Pulumi Ovh provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@lbrlabs/time`](https://www.npmjs.com/package/@lbrlabs/ovh)
* Python: [`lbrlabs_time`](https://pypi.org/project/lbrlabs_ovh/)
* Go: [`github.com/lbrlabs/pulumi-time/sdk/go/time`](https://pkg.go.dev/github.com/lbrlabs/pulumi-ovh/sdk)
* .NET: [`lbrlabs.Time`](https://www.nuget.org/packages/lbrlabs.Ovh)

### Provider Binary

The Ovh provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource ovh v0.0.1
```

Replace the version string with your desired version.
