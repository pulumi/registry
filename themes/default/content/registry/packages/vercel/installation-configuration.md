---
title: Vercel Installation & Configuration
meta_desc: Information on how to install the Vercel provider.
layout: package
---

## Installation

The Pulumi `Vercel` provider is available as a package in all Pulumi languages:

- JavaScript/TypeScript: [`@pulumiverse/vercel`](https://www.npmjs.com/package/@pulumiverse/vercel)
- Python: [`pulumiverse-vercel`](https://pypi.org/project/pulumiverse-vercel/)
- Go: [`github.com/pulumiverse/pulumi-vercel/sdk`](https://pkg.go.dev/github.com/pulumiverse/pulumi-vercel/sdk)
- .NET: [`Pulumiverse.vercel`](https://www.nuget.org/packages/Pulumiverse.vercel)

## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @pulumiverse/vercel
```

or `yarn`:

```bash
yarn add @pulumiverse/vercel
```

### Python

To use from Python, install using `pip`:

```bash
pip install pulumiverse-vercel
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/pulumiverse/pulumi-vercel/sdk/go/...
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package Pulumiverse.vercel
```

## Configuration

The following configuration points are available for the `vercel` provider:

- `vercel:apiToken` (environment: `VERCEL_API_TOKEN`) - the API key for `vercel`
