---
title: Dynatrace Installation & Configuration
meta_desc: Information on how to install the Dynatrace provider.
layout: package
---

## Installation

The Pulumi Dynatrace provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumiverse/pulumi-dynatrace`](https://www.npmjs.com/package/@pulumiverse/pulumi-dynatrace)
* Python: [`dynatrace`](https://pypi.org/project/pulumiverse-pulumi-dynatrace/)
* Go: [`github.com/pulumiverse/pulumi-dynatrace/sdk/go/dynatrace`](https://pkg.go.dev/github.com/pulumiverse/pulumi-dynatrace/sdk)
* .NET: [`Pulumiverse.PulumiPackage.Dynatrace`](https://www.nuget.org/packages/Pulumiverse.PulumiPackage.Dynatrace)

### Provider Binary

The Dynatrace provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource dynatrace <version>
```

Replace the version string with your desired version.

## Setup

To provision resources with the Pulumi Dynatrace provider, you need to have Dynatrace credentials. 

### Set environment variables

Coming soon
