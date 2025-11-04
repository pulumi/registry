---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumiverse/pulumi-exoscale/v0.66.1/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/pulumiverse/pulumi-exoscale/blob/v0.66.1/docs/installation-configuration.md
title: Exoscale Installation & Configuration
meta_desc: Information on how to install the Exoscale provider.
layout: installation
---

## Installation

The Pulumi Exoscale provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumiverse/exoscale`](https://www.npmjs.com/package/@pulumiverse/exoscale)
* Python: [`pulumiverse_exoscale`](https://pypi.org/project/pulumiverse-exoscale/)
* Go: [`github.com/pulumiverse/pulumi-exoscale/sdk/go/exoscale`](https://pkg.go.dev/github.com/pulumiverse/pulumi-exoscale/sdk)
* .NET: [`Pulumiverse.Exoscale`](https://www.nuget.org/packages/Pulumiverse.Exoscale)

## Setup

To provision resources with the Pulumi Exoscale provider, you need to provide the `key` and `secret` (optional `timeout`). 

## Configuration Options

Use `pulumi config set exoscale:<option>`.

| Option    | Required/Optional | Description                                                       |
|-----------|-------------------|-------------------------------------------------------------------|
| `key`     | Required          | Exoscale account API key (environment: `EXOSCALE_API_KEY`).       |
| `secret`  | Required          | Exoscale account API secret (environment: `EXOSCALE_API_SECRET`). |
| `timeout` | Optional          | Global async operations waiting time in seconds (default: 300).   |
