---
title: Symbiosis Installation & Configuration
meta_desc: Information on how to install the Symbiosis provider.
layout: installation
---

## Installation

The Kuraudo.io Symbiosis provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@kuraudo-io/symbiosis`](https://www.npmjs.com/package/@kuraudo-io/symbiosis)
* Python: [`kuraudo_symbiosis`](https://pypi.org/project/kuraudo-symbiosis/)
* Go: [`github.com/kuraudo-io/pulumi-symbiosis/sdk/go/unifi`](https://pkg.go.dev/github.com/kuraudo-io/pulumi-symbiosis/sdk/go/symbiosis)
* .NET: [`Kuraudo.Symbiosis`](https://www.nuget.org/packages/Kuraudo.Symbiosis)

## Setup

To provision resources with the Pulumi Symbiosis provider, you need to have an API key from Symbiosis. You can generate this in the [api keys](https://app.symbiosis.host/api-keys) tab in the Symbiosis console.

## Configuration Options

Use `pulumi config set <option> (--secret) <value>`.

The following configuration points are available for the `symbiosis` provider:

- `symbiosis:apiKey` (environment: `SYMBIOSIS_API_KEY`) - The ApiKey used to authenticate requests towards Symbiosis.
- `symbiosis:endpoint` (environment: `SYMBIOSIS_ENDPOINT`) - Endpoint for reaching the symbiosis API. Used for debugging or when accessed through a proxy.

