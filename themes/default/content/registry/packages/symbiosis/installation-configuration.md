---
title: Symbiosis Installation & Configuration
meta_desc: Information on how to install the Symbiosis provider.
layout: package
---

## Installation

The Symbiosis provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@symbiosis-cloud/symbiosis-pulumi`](https://www.npmjs.com/package/@symbiosis-cloud/symbiosis-pulumi)
* Python: [`symbiosis-pulumi`](https://pypi.org/project/symbiosis-pulumi/)
* Go: [`github.com/symbiosis-cloud/pulumi-symbiosis/sdk/go/symbiosis`](https://pkg.go.dev/github.com/symbiosis-cloud/pulumi-symbiosis/sdk/go/symbiosis)
* .NET: [`Symbiosis.Pulumi.Symbiosis`](https://www.nuget.org/packages/Symbiosis.Pulumi.Symbiosis)

## Setup

To provision resources with the Pulumi Symbiosis provider, you need to have an API key from Symbiosis. You can generate this in the [api keys](https://app.symbiosis.host/api-keys) tab in the Symbiosis console.

## Configuration Options

Use `pulumi config set <option> (--secret) <value>`.

The following configuration points are available for the `symbiosis` provider:

- `symbiosis:apiKey` (environment: `SYMBIOSIS_API_KEY`) - The ApiKey used to authenticate requests towards Symbiosis.
- `symbiosis:endpoint` (environment: `SYMBIOSIS_ENDPOINT`) - Endpoint for reaching the symbiosis API. Used for debugging or when accessed through a proxy.

