---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumiverse/pulumi-heroku/v1.0.4/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Heroku Installation & Configuration
meta_desc: Information on how to install the Pulumi Heroku provider.
layout: installation
---

## Installation

The Pulumi Heroku provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumiverse/heroku`](https://www.npmjs.com/package/@pulumiverse/heroku)
* Python: [`pulumiverse_heroku`](https://pypi.org/project/pulumiverse_heroku/)
* Go: [`github.com/pulumiverse/pulumi-heroku/sdk/go/heroku`](https://github.com/pulumiverse/pulumi-heroku/tree/main/sdk/go/heroku)
* .NET: [`Pulumiverse.Heroku`](https://www.nuget.org/packages/Pulumiverse.Heroku)

### Provider Binary

The Heroku provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource heroku v1.0.0 --server github://api.github.com/pulumiverse
```

Replace the version string with your desired version.
