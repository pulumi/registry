---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumiverse/pulumi-clickhouse/v1.0.1/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Clickhouse Installation & Configuration
meta_desc: Information on how to install the Clickhouse provider.
layout: package
---

## Installation

The Pulumi Clickhouse provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumiverse/clickhouse`](https://www.npmjs.com/package/@pulumiverse/clickhouse)
* Python: [`pulumiverse_clickhouse`](https://pypi.org/project/pulumiverse_clickhouse/)
* Go: [`github.com/pulumiverse/pulumi-clickhouse/sdk/go/clickhouse`](https://pkg.go.dev/github.com/pulumiverse/pulumi-clickhouse/sdk/go/clickhouse)
* .NET: [`Pulumiverse.Clickhouse`](https://www.nuget.org/packages/Pulumiverse.Clickhouse)


## Configuration

> Note:  
> Replace the following **sample content**, with the configuration options
> of the wrapped Terraform provider and remove this note.

The following configuration points are available for the `clickhouse` provider:

- `clickhouse:organizationId` (environment: `CLICKHOUSE_ORG_ID`) - the organization id on clickhouse
- `clickhouse:tokenKey` (environment: `CLICKHOUSE_TOKEN_KEY`) - the token key
- `clickhouse:tokenSecret` (environment: `CLICKHOUSE_TOKEN_SECRET`) - the token secret


### Provider Binary

The Clickhouse provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource clickhouse <version>
```

Replace the version string `<version>` with your desired version.
