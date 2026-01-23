---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pinecone-io/pulumi-pinecone/v2.0.2/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/pinecone-io/pulumi-pinecone/blob/v2.0.2/docs/installation-configuration.md
title: Pinecone Setup
meta_desc: Information on how to install the Pinecone Provider for Pulumi.
layout: package
---

## Installation

The Pulumi Pinecone provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pinecone-database/pulumi`](https://www.npmjs.com/package/@pinecone-database/pulumi)
* Python: [`pinecone_pulumi`](https://pypi.org/project/pinecone_pulumi/)
* Go: [`github.com/pinecone-io/pulumi-pinecone/sdk/go/port`](https://github.com/pinecone-io/pulumi-pinecone)
* * .NET: [`PineconeDatabase.Pinecone`](https://www.nuget.org/packages/PineconeDatabase.Pinecone)

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @pinecone-database/pulumi
```

or `yarn`:

```bash
yarn add @pinecone-database/pinecone
```

### Python

To use from Python, install using `pip`:

```bash
pip install pinecone_pulumi
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/pinecone-io/pulumi-pinecone/sdk
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package PineconeDatabase.Pinecone
```

## Configuration

The following configuration points are available for the `pinecone` provider:

- `pinecone:APIKey` - This is the Pinecone API key. (environment: `PINECONE_API_KEY`)
