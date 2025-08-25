---
# WARNING: this file was fetched from https://raw.githubusercontent.com/vatesfr/pulumi-xenorchestra/v2.2.0/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Xenorchestra Installation & Configuration
meta_desc: Information on how to install the Xenorchestra provider.
layout: package
---

## Installation

The Pulumi `Xenorchestra` provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@vates/pulumi-xenorchestra`](https://www.npmjs.com/package/@vates/pulumi-xenorchestra)
* Python: [`pulumi-xenorchestra`](https://pypi.org/project/pulumi-xenorchestra/)
* Go: [`github.com/vatesfr/pulumi-xenorchestra/sdk`](https://pkg.go.dev/github.com/vatesfr/pulumi-xenorchestra/sdk)
* .NET: [`Pulumi.Xenorchestra`](https://www.nuget.org/packages/Pulumi.Xenorchestra)

### Provider Binary

The Xenorchestra provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource xenorchestra vX.Y.Z --server github://api.github.com/vatesfr/pulumi-xenorchestra
```

Replace the version string with your desired version.

## Setup

To provision resources with the Pulumi Xenorchestra provider, you need to have Xenorchestra credentials.
Your Xenorchestra credentials are never sent to pulumi.com. Pulumi uses the Xenorchestra API and the credentials in your environment to authenticate requests from your computer to your Xenorchestra instance.

### Get your credentials

Use `pulumi config set xenorchestra:<option>` (see [documentation](https://www.pulumi.com/docs/concepts/config/) or pass options to the constructor of new xenorchestra.Provider or use environment variables.

The following configuration points are available for the `xenorchestra` provider:

- `xenorchestra:url` (environment: `XOA_URL`) - the URL for the Xen Orchestra websockets endpoint. Starts with `wss://`
Set either:
- `xenorchestra:username` (environment: `XOA_USERNAME`) - the username for Xen Orchestra
- `xenorchestra:password` (environment: `XOA_PASSWORD`) - the password for Xen Orchestra
Or:
- `xenorchestra:token` (environment: `XOA_TOKEN`) - API token for Xen Orchestra

- `xenorchestra:insecure` (environment: `XOA_INSECURE`) - set to any value to disable SSL verification, false by default. Only use if you are using a self-signed certificate and know what you are doing.


To get a token you can use the [xo-cli](https://docs.xen-orchestra.com/architecture#xo-cli-cli) or the XO GUI (under User > Authentication tokens).

