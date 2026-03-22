---
# WARNING: this file was fetched from https://raw.githubusercontent.com/lokkju/pulumi-improvmx/v0.1.10/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/lokkju/pulumi-improvmx/blob/v0.1.10/docs/installation-configuration.md
title: ImprovMX Installation & Configuration
meta_desc: How to install and configure the ImprovMX Pulumi provider.
layout: package
---

## Installation

The ImprovMX provider is available as a package in the following Pulumi languages:

### Node.js (JavaScript/TypeScript)

```bash
npm install pulumi-improvmx
```

### Python

```bash
pip install pulumi-improvmx
```

### .NET

```bash
dotnet add package Pulumi.Improvmx
```

### Go

```bash
go get github.com/lokkju/improvmx/sdk/go/improvmx
```

## Configuration

The following configuration options are available:

| Option | Required | Description | Environment Variable |
|--------|----------|-------------|---------------------|
| `improvmx:apiToken` | Yes | ImprovMX API token | `IMPROVMX_API_TOKEN` |

### Set via Pulumi config

```bash
pulumi config set improvmx:apiToken --secret sk_xxxxx
```

### Set via environment variable

```bash
export IMPROVMX_API_TOKEN=sk_xxxxx
```

You can get your API token from the [ImprovMX account settings](https://app.improvmx.com/account/api).
