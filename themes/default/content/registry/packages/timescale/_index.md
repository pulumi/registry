---
# WARNING: this file was fetched from https://ds7waww6k7rw2.cloudfront.net/docs/registry.opentofu.org/timescale/timescale/1.12.1/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Timescale Provider
meta_desc: Provides an overview on how to configure the Pulumi Timescale provider.
layout: package
---

## Generate Provider

The Timescale provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider timescale/timescale
```
## Overview

The Pulumi provider for [Timescale](https://www.timescale.com/cloud).
## Requirements
- Pulumi >= 1.0
## Quick Start
### Authorization
When you log in to your [Timescale Account](https://console.cloud.timescale.com/), navigate to the `Project settings` page.
From here, you can create client credentials for programmatic usage. Click the `Create credentials` button to generate a new public/secret key pair.
### Project ID
The project ID can be found on the `Project settings` page.

Create a `main.tf` configuration file with the following content.
### VPC Peering

Since v1.9.0 it is possible to peer Timescale VPCs using pulumi.

Below is an example configuration:

As of v1.9.2, the awsVpcPeeringConnection import id must be manually added. This
value, always starting with `pcx-...` will be available in `timescale_peering_connection.pc.provisioned_id` after a pulumi refresh.
## Supported Service Configurations
### Compute
- 500m CPU / 2 GB Memory
- 1000m CPU / 4 GB Memory
- 2000m CPU / 8 GB Memory
- 4000m CPU / 16 GB Memory
- 8000m CPU / 32 GB Memory
- 16000m CPU / 64 GB Memory
- 32000m CPU / 128 GB Memory
### Storage
Since June 2023, you no longer need to allocate a fixed storage volume or worry about managing your disk size, and you'll be billed only for the storage you actually use.
See more info in our [blogpost](https://www.timescale.com/blog/savings-unlocked-why-we-switched-to-a-pay-for-what-you-store-database-storage-model/)
## Supported Operations
✅ Create service <br />
✅ Rename service <br />
✅ Resize service <br />
✅ Pause/resume service <br />
✅ Delete service <br />
✅ Import service <br />
✅ Enable High Availability replicas <br />
✅ Enable read replicas <br />
✅ VPC peering <br />
✅ Connection pooling <br />
## Billing
Services are currently billed for hourly usage. If a service is running for less than an hour,
it will still be charged for the full hour of usage.