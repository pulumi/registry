---
# WARNING: this file was fetched from https://raw.githubusercontent.com/lbrlabs/pulumi-dynatrace/v0.30.0/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Dynatrace Installation & Configuration
meta_desc: Information on how to install the Dynatrace provider.
layout: package
---

## Installation

The Pulumi Dynatrace provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumiverse/dynatrace`](https://www.npmjs.com/package/@pulumiverse/dynatrace)
* Python: [`dynatrace`](https://pypi.org/project/pulumiverse-dynatrace/)
* Go: [`github.com/pulumiverse/pulumi-dynatrace/sdk/go/dynatrace`](https://pkg.go.dev/github.com/pulumiverse/pulumi-dynatrace/sdk)
* .NET: [`Pulumiverse.Dynatrace`](https://www.nuget.org/packages/Pulumiverse.Dynatrace)

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
