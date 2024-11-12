---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumiverse/pulumi-aquasec/v0.8.29/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Aquasec Installation & Configuration
meta_desc: Information on how to install the Aquasec provider.
layout: package
---

## Installation

The Pulumi Aquasec provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumiverse/aquasec`](https://www.npmjs.com/package/@pulumiverse/aquasec)
* Python: [`pulumiverse_aquasec`](https://pypi.org/project/pulumiverse-aquasec/)
* Go: [`github.com/pulumiverse/pulumi-aquasec/sdk/go/aquasec`](https://pkg.go.dev/github.com/pulumiverse/pulumi-aquasec/sdk)
* .NET: [`Pulumiverse.Aquasec`](https://www.nuget.org/packages/Pulumiverse.Aquasec)

## Setup

To provision resources with the Pulumi Aquasec provider, you need to provide the `username`, `password` and `aqua_url`. 

## Configuration Options

Use `pulumi config set aquasec:<option>`.

| Option     | Required/Optional | Description                                                                                     |
|------------|-------------------|-------------------------------------------------------------------------------------------------|
| `username` | Required          | This is the user id that should be used to make the connection (environment: `AQUA_USER`).      |
| `password` | Required          | This is the password that should be used to make the connection (environment: `AQUA_PASSWORD`). |
| `aqua_url` | Required          | This is the base URL of your Aqua instance (environment: `AQUA_URL`).                           |
