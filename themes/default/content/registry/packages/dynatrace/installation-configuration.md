---
title: Dynatrace Installation & Configuration
meta_desc: Information on how to install the Dynatrace provider.
layout: installation
---

## Installation

The Pulumi Dynatrace provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@lbrlabs/pulumi-dynatrace`](https://www.npmjs.com/package/@lbrlabs/pulumi-dynatrace)
* Python: [`dynatrace`](https://pypi.org/project/lbrlabs-pulumi-dynatrace/)
* Go: [`github.com/lbrlabs/pulumi-dynatrace/sdk/go/dynatrace`](https://pkg.go.dev/github.com/lbrlabs/pulumi-dynatrace/sdk)
* .NET: [`Lbrlabs.PulumiPackage.Dynatrace`](https://www.nuget.org/packages/Lbrlabs.PulumiPackage.Dynatrace)

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
