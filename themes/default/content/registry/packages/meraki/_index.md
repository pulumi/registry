---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-meraki/v0.4.4/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/pulumi/pulumi-meraki/blob/v0.4.4/docs/_index.md
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Cisco Meraki Provider
meta_desc: Provides an overview on how to configure the Pulumi Cisco Meraki provider.
layout: package
---

## Installation

The Cisco Meraki provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/meraki`](https://www.npmjs.com/package/@pulumi/meraki)
* Python: [`pulumi-meraki`](https://pypi.org/project/pulumi-meraki/)
* Go: [`github.com/pulumi/pulumi-meraki/sdk/go/meraki`](https://github.com/pulumi/pulumi-meraki)
* .NET: [`Pulumi.Meraki`](https://www.nuget.org/packages/Pulumi.Meraki)
* Java: [`com.pulumi/meraki`](https://central.sonatype.com/artifact/com.pulumi/meraki)

## Configuration Reference
### Required
- `merakiDashboardApiKey` (String, Sensitive) Cisco  merakiDashboardApiKey to authenticate. If not set, it uses the MERAKI_DASHBOARD_API_KEY environment variable.

- `merakiBaseUrl` (String) Cisco Meraki base URL, FQDN or IP. If not set, it uses the MERAKI_BASE_URL environment variable defaults is (<https://api.meraki.com/)>.
- `merakiDebug` (String) Flag for Cisco Meraki to enable debugging. If not set, it uses the MERAKI_DEBUG environment variable defaults to `false`.
- `merakiRequestsPerSecond` (Int) Setting requests per second, default is 10.