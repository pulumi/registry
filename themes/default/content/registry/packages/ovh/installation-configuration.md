---
# WARNING: this file was fetched from https://raw.githubusercontent.com/ovh/pulumi-ovh/v1.5.0/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
title: OVH Installation & Configuration
meta_desc: Information on how to install the OVH provider.
layout: installation
---

## Information

Note that the [lbrlabs Pulumi OVH provider](https://github.com/lbrlabs/pulumi-ovh) is replaced by this official one.

## Installation

The Pulumi `OVH` provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@ovhcloud/pulumi-ovh`](https://www.npmjs.com/package/@ovhcloud/pulumi-ovh)
* Python: [`pulumi_ovh`](https://pypi.org/project/pulumi-ovh/)
* Go: [`github.com/ovh/pulumi-ovh/sdk/go/ovh`](https://pkg.go.dev/github.com/ovh/pulumi-ovh/sdk)
* .NET: [`Pulumi.Ovh`](https://www.nuget.org/packages/Pulumi.Ovh)

### Provider Binary

The OVH provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource ovh vX.Y.Z --server github://api.github.com/ovh/pulumi-ovh
```

Replace the version string with your desired version.

## Setup

To provision resources with the Pulumi OVH provider, you need to have OVH credentials.
Your OVH credentials are never sent to pulumi.com. Pulumi uses the OVH API and the credentials in your environment to authenticate requests from your computer to OVH.

### Get your credentials

The "OVH provider" needs to be configured with a set of credentials, which can be set using
[Pulumi stack configuration](https://www.pulumi.com/docs/concepts/config/) or environment variables:

* `ovh:endpoint` (environment variable: `OVH_ENDPOINT`)
* `ovh:applicationKey` (environment variable: `OVH_APPLICATION_KEY`)
* `ovh:applicationSecret` (secret) (environment variable: `OVH_APPLICATION_SECRET`)
* `ovh:consumerKey` (environment variable: `OVH_CONSUMER_KEY`)

Why?

Because, behind the scenes, the provider is doing requests to OVHcloud APIs. 

In order to retrieve this necessary information, please follow [First steps with the OVHcloud APIs](https://docs.ovh.com/gb/en/customer/first-steps-with-ovh-api/) tutorial.

Concretely, you have to generate these credentials via the [OVH token generation page](https://api.ovh.com/createToken/?GET=/*&POST=/*&PUT=/*&DELETE=/*) with the following rights:

* GET `/`
* POST `/*`
* PUT `/*`
* DELETE `/*`

When you have successfully generated your OVH tokens, please keep them. You'll have to define them in the coming minutes ;-).

### Set environment variables

Once you have provisioned these credentials, you can set environment variables to provision resources in Grafana:

{{< chooser os "linux,macos,windows" >}}
{{% choosable os linux %}}

```bash
$ export OVH_ENDPOINT="<the Ovh endpoint, for example ovh-eu>"
$ export OVH_APPLICATION_KEY="<the Ovh application key>"
$ export OVH_APPLICATION_SECRET="<the Ovh application secret>"
$ export OVH_CONSUMER_KEY="<the Ovh consumer key>"
```

{{% /choosable %}}

{{% choosable os macos %}}

```bash
$ export OVH_ENDPOINT="<the Ovh endpoint, for example ovh-eu>"
$ export OVH_APPLICATION_KEY="<the Ovh application key>"
$ export OVH_APPLICATION_SECRET="<the Ovh application secret>"
$ export OVH_CONSUMER_KEY="<the Ovh consumer key>"
```

{{% /choosable %}}

{{% choosable os windows %}}

```powershell
> $env:OVH_ENDPOINT = "<the Ovh endpoint, for example ovh-eu>"
> $env:OVH_APPLICATION_KEY = "<the Ovh application key>"
> $env:OVH_APPLICATION_SECRET = "<the Ovh application secret>"
> $env:OVH_CONSUMER_KEY = "<the Ovh consumer key>"
```

{{% /choosable %}}
{{< /chooser >}}
