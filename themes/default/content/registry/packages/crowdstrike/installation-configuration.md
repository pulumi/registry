---
# WARNING: this file was fetched from https://raw.githubusercontent.com/crowdstrike/pulumi-crowdstrike/v0.0.14/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
title: CrowdStrike Installation & Configuration
meta_desc: Information on how to install the CrowdStrike provider.
layout: installation
---

## Installation

The Pulumi CrowdStrike provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@crowdstrike/pulumi`](https://www.npmjs.com/package/@crowdstrike/pulumi)
* Python: [`crowdstrike_pulumi`](https://pypi.org/project/crowdstrike_pulumi/)
* Go: [`github.com/crowdstrike/pulumi-crowdstrike/sdk/go/crowdstrike`](https://pkg.go.dev/github.com/crowdstrike/pulumi-crowdstrike/sdk/go/crowdstrike)

## Environment Variables

The following configuration environment variables are required for the `CrowdStrike` provider:

* CROWDSTRIKE_FALCON_CLIENT_ID - the user generated client id for the `CrowdStrike` provider
* CROWDSTRIKE_FALCON_CLIENT_SECRET - the user generated client secret for the `CrowdStrike` provider

### Provider Binary

The CrowdStrike provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource crowdstrike <version>
```

Replace the version string `<version>` with your desired version.
