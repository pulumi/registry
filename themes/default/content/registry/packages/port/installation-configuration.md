---
title: Port Setup
meta_desc: Information on how to install the Port provider.
layout: package
---

## Installation

The Pulumi Port provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@port-labs/port`](https://www.npmjs.com/package/@port-labs/port)
* Python: [`port_pulumi`](https://pypi.org/project/port_pulumi/)
* Go: [`github.com/port-labs/pulumi-port/sdk/go/port`](https://github.com/port-labs/pulumi-port/sdk/go/port)

## Configuration

The following configuration points are available for the `port` provider:

- `port:baseUrl` (environment: `PORT_BASE_URL`) - This is the Port base URL.
- `port:clientId` (environment: `PORT_CLIENT_ID`) - This is the Port client ID.
- `port:secret` (environment: `PORT_CLIENT_SECRET`) - This is the Port secret.
- `port:token` - This is the Port token.
