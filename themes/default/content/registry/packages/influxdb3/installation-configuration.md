---
# WARNING: this file was fetched from https://raw.githubusercontent.com/komminarlabs/pulumi-influxdb3/v1.3.0/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
title: InfluxDB V3 Installation & Configuration
meta_desc: Information on how to install the InfluxDB V3 provider.
layout: package
---

## Installation

The Pulumi InfluxDB V3 provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@komminarlabs/influxdb3`](https://www.npmjs.com/package/@komminarlabs/influxdb3)
* Python: [`komminarlabs_influxdb3`](https://pypi.org/project/komminarlabs_influxdb3/)
* Go: [`github.com/komminarlabs/pulumi-influxdb3/sdk/go/influxdb3`](https://pkg.go.dev/github.com/komminarlabs/pulumi-influxdb3/sdk/go/influxdb3)
* .NET: [`KomminarLabs.InfluxDB3`](https://www.nuget.org/packages/KomminarLabs.InfluxDB3)


## Configuration

The following configuration points are available for the `InfluxDB3` provider:

- `influxdb3:account_id` (environment: `INFLUXDB3_ACCOUNT_ID`) - The ID of the account that the cluster belongs to
- `influxdb3:cluster_id` (environment: `INFLUXDB3_CLUSTER_ID`) - The ID of the cluster that you want to manage
- `influxdb3:token` (environment: `INFLUXDB3_TOKEN`) - The InfluxDB management token

### Provider Binary

The InfluxDB V3 provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource influxdb3 <version> --server github://api.github.com/komminarlabs
```

Replace the version string `<version>` with your desired version.
