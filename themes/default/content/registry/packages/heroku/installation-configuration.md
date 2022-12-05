---
title: Heroku Installation & Configuration
meta_desc: Information on how to install the Pulumi Heroku provider.
layout: installation
---

## Installation

The Pulumi Heroku provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumiverse/heroku`](https://www.npmjs.com/package/@pulumiverse/heroku)
* Python: [`pulumiverse_heroku`](https://pypi.org/project/pulumiverse_heroku/)
* Go: [`github.com/pulumiverse/pulumi-heroku/sdk/go/heroku`](https://github.com/pulumiverse/pulumi-heroku/sdk/go/heroku)
* .NET: [`Pulumiverse.Heroku`](https://www.nuget.org/packages/Pulumiverse.Heroku)

### Provider Binary

The Heroku provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource heroku v0.1.0
```

Replace the version string with your desired version.
