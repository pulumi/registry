---
# WARNING: this file was fetched from https://raw.githubusercontent.com/azaurus1/pulumi-kafka-connect/v0.0.12/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Kafka Connect Setup
meta_desc: Information on how to install the Kafka Connect Provider for Pulumi.
layout: package
---

## Installation
The Kafka Connect Pulumi provider is available as a package in all Pulumi languages:
- JavaScript/TypeScript: [`@azaurus/kafkaconnect`]("https://www.npmjs.com/package/@azaurus/kafkaconnect")
- Python: [`pulumi-kafkaconnect`]("https://pypi.org/project/pulumi-kafkaconnect/")
- Go: [`github.com/azaurus1/pulumi-kafka-connect/sdk/go/kafkaconnect`]("https://github.com/azaurus1/pulumi-kafka-connect/sdk/go/kafkaconnect")
- .NET: [`Azaurus1.Kafkaconnect`]("https://www.nuget.org/packages/Azaurus1.Kafkaconnect/")

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @azaurus/kafkaconnect
```

or `yarn`:

```bash
yarn add @azaurus/kafkaconnect
```

### Python

To use from Python, install using `pip`:

```bash
pip install pulumi-kafkaconnect
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/azaurus1/pulumi-kafka-connect/sdk/go/kafkaconnect
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package Azaurus1.Kafkaconnect
```

## Configuration
The following configuration points are available for the `kafka-connect` provider:
- `kafkaconnect:url` - This is the URL of the Connect cluster (Env: `KAFKA_CONNECT_URL`)

### Configuring the Provider

#### Using `pulumi config set`
You can configure the provider using Pulumi's `config set` command. For example:

```bash
pulumi config set kafkaconnect:url https://my-kafka-connect-cluster:8083
```

#### Using Environment Variables
You can also configure the provider by setting environment variables before running your Pulumi program. For example:

```bash
export KAFKA_CONNECT_URL=https://my-kafka-connect-cluster:8083
```

If both `pulumi config` and environment variables are set, the environment variable will take precedence.
