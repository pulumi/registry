---
title: Cisco Meraki Installation & Configuration
meta_desc: Information on how to install the Meraki provider.
layout: installation
---

## Installation

The Pulumi Cisco Meraki provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/meraki`](https://www.npmjs.com/package/@pulumi/meraki)
* Python: [`pulumi_meraki`](https://pypi.org/project/pulumi_meraki/)
* Go: [`github.com/pulumi/pulumi-meraki/sdk/go/meraki`](https://pkg.go.dev/github.com/pulumi/pulumi-meraki/sdk/go/meraki)
* .NET: [`Pulumi.Meraki`](https://www.nuget.org/packages/Pulumi.Meraki)


## Configuration

The following configuration points are available for the `meraki` provider:

- `meraki:merakiBaseUrl` (environment: `MERAKI_BASE_URL`) - base URL, FQDN or IP, default is https://api.meraki.com/
- `meraki:merakiDashboardApiKey` (environment: `MERAKI_DASHBOARD_API_KEY`) - dashboard API key to authenticate
- `meraki:merakiDebug` (environment: `MERAKI_DEBUG`) - flag to enable debugging, default is false
- `meraki:merakiRequestsPerSecond` - requests per second allowed for client, default is 10
