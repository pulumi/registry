---
# WARNING: this file was fetched from https://raw.githubusercontent.com/descope/pulumi-descope/v0.3.4/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/descope/pulumi-descope/blob/v0.3.4/docs/installation-configuration.md
title: Descope Installation & Configuration
meta_desc: Information on how to install the Descope provider.
layout: package
---

## Installation

The Pulumi `Descope` provider is available as a package in all Pulumi languages:

- JavaScript/TypeScript: [`@descope/pulumi-descope`](https://www.npmjs.com/package/@descope/pulumi-descope)
- Python: [`descope-pulumi`](https://pypi.org/project/descope-pulumi/)
- Go: [`github.com/descope/pulumi-descope/sdk`](https://pkg.go.dev/github.com/descope/pulumi-descope/sdk)
- .NET: [`Descope.Pulumi.Descope`](https://www.nuget.org/packages/Descope.Pulumi.Descope)

## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @descope/pulumi-descope
```

or `yarn`:

```bash
yarn add @descope/pulumi-descope
```

### Python

To use from Python, install using `pip`:

```bash
pip install descope-pulumi
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/descope/pulumi-descope/sdk/go/...
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package Descope.Pulumi.Descope
```

### Provider Binary

The Descope provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource descope <version> --server github://api.github.com/descope
```

Replace the version string with your desired version, or omit it to install the latest version.

## Configuration

The following configuration points are available for the `descope` provider:

- `descope:projectId` (environment: `DESCOPE_PROJECT_ID`) - Descope Project ID
- `descope:managementKey` (environment: `DESCOPE_MANAGEMENT_KEY`) - Descope Management Key

both available from the Project and Company Settings page in Descope dashboard.
