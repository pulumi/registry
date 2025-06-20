---
# WARNING: this file was fetched from https://raw.githubusercontent.com/komminarlabs/pulumi-influxdb/v1.4.0/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
title: InfluxDB Installation & Configuration
meta_desc: Information on how to install the InfluxDB provider.
layout: package
---

## Installation

The Pulumi InfluxDB provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@komminarlabs/influxdb`](https://www.npmjs.com/package/@komminarlabs/influxdb)
* Python: [`komminarlabs_influxdb`](https://pypi.org/project/komminarlabs_influxdb/)
* Go: [`github.com/komminarlabs/pulumi-influxdb/sdk/go/influxdb`](https://pkg.go.dev/github.com/komminarlabs/pulumi-influxdb/sdk/go/influxdb)
* .NET: [`KomminarLabs.InfluxDB`](https://www.nuget.org/packages/KomminarLabs.InfluxDB)


## Configuration

The following configuration points are available for the `InfluxDB` provider:

- `influxdb:token` (environment: `INFLUXDB_TOKEN`) - An InfluxDB token string
- `influxdb:url` (environment: `INFLUXDB_URL`) - The InfluxDB Cloud Dedicated server URL

### Provider Binary

The InfluxDB provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource influxdb <version> --server github://api.github.com/komminarlabs
```

Replace the version string `<version>` with your desired version.
