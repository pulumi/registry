---
title: Flux Installation & Configuration
meta_desc: Information on how to install the Flux provider.
layout: installation
---

## Installation

The Pulumi Flux provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@worawat/flux`](https://www.npmjs.com/package/@worawat/flux)
* Python: [`pulumi-flux`](https://pypi.org/project/pulumi-flux/)
* Go: [`github.com/oun/pulumi-flux/sdk/go/flux`](https://pkg.go.dev/github.com/oun/pulumi-flux/sdk)
* .NET: [`Pulumi.Flux`](https://www.nuget.org/packages/Pulumi.Flux)

### Provider Binary

The Flux provider binary is a third party binary. If you use [Pulumi v3.35.3 or up](https://www.pulumi.com/docs/guides/pulumi-packages/how-to-author/#support-for-github-releases), the provider binary should install automatically when using one of the language SDKs. If you are using an older version or it fails to install, the binary can be installed using the `pulumi plugin` command:

```bash
pulumi plugin install resource flux <version> --server https://github.com/oun/pulumi-flux/releases/download/<version>
```

Replace the version string with your desired version.
