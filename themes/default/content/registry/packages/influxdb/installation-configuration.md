---
# WARNING: this file was fetched from https://raw.githubusercontent.com/komminarlabs/pulumi-influxdb/v1.6.0/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/komminarlabs/pulumi-influxdb/blob/v1.6.0/docs/installation-configuration.md
title: InfluxDB Installation & Configuration
meta_desc: Information on how to install the InfluxDB provider.
layout: package
---

## Installation

The Pulumi InfluxDB provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@thulasirajkomminar/influxdb`](https://www.npmjs.com/package/@thulasirajkomminar/influxdb)
* Python: [`thulasirajkomminar_influxdb`](https://pypi.org/project/thulasirajkomminar_influxdb/)
* Go: [`github.com/thulasirajkomminar/pulumi-influxdb/sdk/go/influxdb`](https://pkg.go.dev/github.com/thulasirajkomminar/pulumi-influxdb/sdk/go/influxdb)
* .NET: [`ThulasirajKomminar.InfluxDB`](https://www.nuget.org/packages/thulasirajkomminar.InfluxDB)


## Configuration

The following configuration points are available for the `InfluxDB` provider:

- `influxdb:token` (environment: `INFLUXDB_TOKEN`) - An InfluxDB token string
- `influxdb:url` (environment: `INFLUXDB_URL`) - The InfluxDB server URL

### Provider Binary

The InfluxDB provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource influxdb <version> --server github://api.github.com/thulasirajkomminar
```

Replace the version string `<version>` with your desired version.
