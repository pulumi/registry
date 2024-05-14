---
title: Meraki Installation & Configuration
meta_desc: Information on how to install the Meraki provider.
layout: installation
---

## Installation

The Pulumi Meraki provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/meraki`](https://www.npmjs.com/package/@pulumi/meraki)
* Python: [`pulumi_meraki`](https://pypi.org/project/pulumi_meraki/)
* Go: [`github.com/pulumi/pulumi-meraki/sdk/go/meraki`](https://pkg.go.dev/github.com/pulumi/pulumi-meraki/sdk/go/meraki)
* .NET: [`Pulumi.Meraki`](https://www.nuget.org/packages/Pulumi.Meraki)


## Configuration

> Note:  
> Replace the following **sample content**, with the configuration options
> of the wrapped Terraform provider and remove this note.

The following configuration points are available for the `meraki` provider:

- `meraki:apiKey` (environment: `meraki_API_KEY`) - the API key for `meraki`
- `meraki:region` (environment: `meraki_REGION`) - the region in which to deploy resources

### Provider Binary

The Meraki provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource meraki <version>
```

Replace the version string `<version>` with your desired version.
