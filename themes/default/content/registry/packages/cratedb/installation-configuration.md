---
title: CrateDB Installation & Configuration
meta_desc: Information on how to install the CrateDB provider.
layout: package
---

## Installation

The Pulumi CrateDB provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@komminarlabs/cratedb`](https://www.npmjs.com/package/@komminarlabs/cratedb)
* Python: [`komminarlabs_cratedb`](https://pypi.org/project/komminarlabs_cratedb/)
* Go: [`github.com/komminarlabs/pulumi-cratedb/sdk/go/cratedb`](https://pkg.go.dev/github.com/komminarlabs/pulumi-cratedb/sdk/go/cratedb)
* .NET: [`KomminarLabs.CrateDB`](https://www.nuget.org/packages/KomminarLabs.CrateDB)


## Configuration

The following configuration points are available for the `cratedb` provider:

- `cratedb:api_key` (environment: `CRATEDB_API_KEY`) - The API key
- `cratedb:api_secret` (environment: `CRATEDB_API_SECRET`) - The API secret
- `cratedb:url` (environment: `CRATEDB_URL`) - The CrateDB Cloud URL

### Provider Binary

The CrateDB provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource cratedb <version> --server github://api.github.com/komminarlabs
```

Replace the version string `<version>` with your desired version.
