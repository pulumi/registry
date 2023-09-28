---
title: Zitadel Installation & Configuration
meta_desc: Information on how to install the Zitadel provider.
layout: installation
---

## Installation

The Pulumi Zitadel provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumiserve/zitadel`](https://www.npmjs.com/package/@pulumiverse/zitadel)
* Python: [`pulumiverse_zitadel`](https://pypi.org/project/pulumiverse-zitadel/)
* Go: [`github.com/pulumiverse/pulumi-zitadel/sdk/go/zitadel`](https://pkg.go.dev/github.com/pulumiverse/pulumi-zitadel/sdk)
* .NET: [`Pulumiverse.Zitadel`](https://www.nuget.org/packages/Pulumiverse.Zitadel)

## Setup

To provision resources with the Pulumi Zitadel provider, you need to have a Zitadel domain and JWT profile.

## Configuration Options

Use `pulumi config set zitadel:<option>`.

| Option | Required/Optional | Description |
|-----|------|----|
| `domain`| Required | Domain used to connect to the Zitadel instance |
| `insecure`| Optional | Use insecure connection |
| `jwtProfileFile`| Optional |  Path to the file containing credentials to connect to Zitadel. Either `jwtProfileFile` or `jwtProfileJson` is required |
| `jwtProfileJson`| Optional | JSON value of credentials to connect to Zitadel. Either `jwtProfileFile` or `jwtProfileJson` is required |
| `port`| Optional |  Used port if not the default ports 80 or 443 are configured |
| `jwtProfileJson`| Optional | Domain used to connect to the Zitadel instance |
| `token`| Optional | Path to the file containing credentials to connect to Zitadel |