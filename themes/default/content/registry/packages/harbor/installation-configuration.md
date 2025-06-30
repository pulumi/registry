---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumiverse/pulumi-harbor/v3.10.21/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Harbor Installation & Configuration
meta_desc: Information on how to install the Harbor provider.
layout: installation
---

## Installation

The Pulumi `Harbor` provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumiverse/harbor`](https://www.npmjs.com/package/@pulumiverse/harbor)
* Python: [`pulumiverse-harbor`](https://pypi.org/project/pulumiverse-harbor/)
* Go: [`github.com/pulumiverse/pulumi-harbor/sdk/v3`](https://pkg.go.dev/github.com/pulumiverse/pulumi-harbor/sdk/v3)
* .NET: [`Pulumiverse.Harbor`](https://www.nuget.org/packages/Pulumiverse.Harbor)

## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @pulumiverse/harbor
```

or `yarn`:

```bash
yarn add @pulumiverse/harbor
```

### Python

To use from Python, install using `pip`:

```bash
pip install pulumiverse-harbor
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/pulumiverse/pulumi-harbor/sdk/v3
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package Pulumiverse.Harbor
```

## Configuration

The following configuration points are available for the `harbor` provider:

- `harbor:url` - (Required) The url of the harbor.
- `harbor:username` - (Required) The username to be used to access harbor.
- `harbor:password` - (Required) The password to be used to access harbor.
- `harbor:insecure` - (Optional) Choose to ignore certificate errors
- `harbor:apiVersion` - (Optional) Choose which version of the api you would like to use `1` or `2`. Default is `2`.
