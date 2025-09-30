---
# WARNING: this file was fetched from https://raw.githubusercontent.com/dirien/pulumi-vultr/v2.27.1/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/dirien/pulumi-vultr/blob/v2.27.1/docs/installation-configuration.md
title: Vultr Installation & Configuration
meta_desc: Information on how to install the Vultr provider.
layout: package
---

## Installation

The Pulumi Vultr provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@ediri/vultr`](https://www.npmjs.com/package/@ediri/vultr)
* Python: [`ediri_vultr`](https://pypi.org/project/ediri-vultr/)
* Go: [`github.com/dirien/pulumi-vultr/sdk/v2`](https://pkg.go.dev/github.com/dirien/pulumi-vultr/sdk/v2)
* .NET: [`ediri.Vultr`](https://www.nuget.org/packages/ediri.Vultr)

## Setup

To provision resources with the Pulumi Vultr provider, you need to provide the `apiKey`.

## Configuration Options

Use `pulumi config set vultr:<option>`.

The following configuration points are available for the `vultr` provider:

- `vultr:apiKey` (environment: `VULTR_API_KEY`) -  This is the Vultr API key. It can be found in the [Vultr API section](https://my.vultr.com/settings/#settingsapi).
- `vultr:rateLimit` - Vultr limits API calls to 30 calls per second. This field lets you configure how the rate limit using milliseconds. The default value if this field is omitted is 500 milliseconds per call.
- `vultr:retryLimit` - This field lets you configure how many retries should be attempted on a failed call. The default value if this field is omitted is 3 retries.
