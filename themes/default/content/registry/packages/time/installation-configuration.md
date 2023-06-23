---
title: Time Installation & Configuration
meta_desc: Information on how to install the Pulumi Time provider.
layout: installation
---

## Installation

The Pulumi Time provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumiverse/time`](https://www.npmjs.com/package/@pulumiverse/time)
* Python: [`pulumiverse_time`](https://pypi.org/project/pulumiverse_time/)
* Go: [`github.com/pulumiverse/pulumi-time/sdk/go/time`](https://pkg.go.dev/github.com/pulumiverse/pulumi-time/sdk/go/time)
* .NET: [`Pulumiverse.Time`](https://www.nuget.org/packages/Pulumiverse.Time)

### Provider Binary

The Time provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource time <version>
```

Replace the `<version>` string with your desired version.
