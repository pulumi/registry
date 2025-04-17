---
# WARNING: this file was fetched from https://raw.githubusercontent.com/port-labs/pulumi-port/v2.6.1/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Port Setup
meta_desc: Information on how to install the Port provider.
layout: package
---

## Installation

The Pulumi Port provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@port-labs/port`](https://www.npmjs.com/package/@port-labs/port)
* Python: [`port_pulumi`](https://pypi.org/project/port_pulumi/)
* Go: [`github.com/port-labs/pulumi-port/sdk/go/port`](https://github.com/port-labs/pulumi-port/)

## Configuration

The following configuration points are available for the `port` provider:

- `port:clientId` - This is the Port client ID. (environment: PORT_CLIENT_ID)
- `port:secret` - This is the Port secret. (environment: PORT_CLIENT_SECRET)
- `port:baseUrl` (optional) - This is the Port base URL. (environment: PORT_BASE_URL)
- `port:token` - (optional) This is the Port token.
