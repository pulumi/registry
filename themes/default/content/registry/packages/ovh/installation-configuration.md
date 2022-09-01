---
title: Ovh Setup
meta_desc: Information on how to install the Pulumi Ovh provider.
layout: installation
---

## Installation

The Pulumi Ovh provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumiverse/time`](https://www.npmjs.com/package/@pulumiverse/ovh)
* Python: [`pulumiverse_time`](https://pypi.org/project/pulumiverse_ovh/)
* Go: [`github.com/pulumiverse/pulumi-time/sdk/go/time`](https://pkg.go.dev/github.com/pulumiverse/pulumi-ovh/sdk)
* .NET: [`Pulumiverse.Time`](https://www.nuget.org/packages/Pulumiverse.Ovh)

### Provider Binary

The Ovh provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource ovh v0.0.1
```

Replace the version string with your desired version.
