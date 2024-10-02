---
title: Adobe Experience Manager Provider Installation & Configuration
meta_desc: Information on how to install the Adobe Experience Manager provider.
layout: package
---

## Installation

The Adobe Experience Manager provider is available as a package in following Pulumi languages:

* JavaScript/TypeScript: [`@wttech/aem`](https://www.npmjs.com/package/@wttech/aem)
* Python: [`wttech_aem`](https://pypi.org/project/wttech-aem/)
* Go: [`github.com/wttech/pulumi-aem/sdk/go/aem`](https://pkg.go.dev/github.com/wttech/pulumi-aem/sdk)
* .NET: [`WTTech.Aem`](https://www.nuget.org/packages/WTTech.Aem)

### Provider Binary

The Adobe Experience Manager provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource aem <version> --server github://api.github.com/wttech
```

Replace the version string with your desired version.

## Setup

To provision resources with the Adobe Experience Manager provider, you need to have correct configuration towards your AWS or other cloud controller.

### Set environment variables

The Adobe Experience Manager provider does not use any specific environment variables.

## Configuration Options

The Adobe Experience Manager provider does not require any additional options.
