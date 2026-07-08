---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumiverse/pulumi-matchbox/v0.3.0/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/pulumiverse/pulumi-matchbox/blob/v0.3.0/docs/installation-configuration.md
title: Matchbox Setup
meta_desc: Information on how to install the Matchbox provider.
layout: package
---

## Installation

The Pulumi Matchbox provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumiverse/matchbox`](https://www.npmjs.com/package/@pulumiverse/matchbox)
* Python: [`pulumiverse_matchbox`](https://pypi.org/project/pulumiverse-matchbox/)
* Go: [`github.com/pulumiverse/pulumi-matchbox/sdk/go/matchbox`](https://pkg.go.dev/github.com/pulumiverse/pulumi-matchbox/sdk/go/matchbox)
* .NET: [`Pulumiverse.Matchbox`](https://www.nuget.org/packages/Pulumiverse.Matchbox)

## Setup

To provision resources with the Pulumi Matchbox provider, you need to have a certificate. 

## Configuration Options

Use `pulumi config set matchbox:<option>`.

| Option | Required/Optional | Description |
|-----|------|----|
| `endpoint`| Required | Endpoint to the Matchbox server in <host>:<port> format |
| `client_cert`| Optional | Client certificate |
| `client_key`| Optional | Client key |
| `ca`| Optional | Certificate Authority chain |
