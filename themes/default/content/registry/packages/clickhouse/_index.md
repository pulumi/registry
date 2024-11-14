---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumiverse/pulumi-clickhouse/v1.0.1/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Clickhouse
meta_desc: Provides an overview of the Clickhouse Provider for Pulumi.
layout: package
---

The Clickhouse provider for Pulumi can be used to provision any of the cloud resources available in [Clickhouse Cloud](https://clickhouse.cloud).

The Clickhouse provider must be configured with credentials to deploy and update resources in Clickhouse Cloud.

## Example

{{< chooser language "typescript,python" >}}
{{% choosable language typescript %}}

```typescript
import * as clickhouse from "@pulumiverse/clickhouse";

new clickhouse.Service("example", {
    region: "us-central1",
    cloudProvider: "gcp",
    tier: "development",
    password: "1234",
    ipAccesses: [{
        source: "0.0.0.0",
        description: "Test IP"
      }]
});

```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumiverse_clickhouse as clickhouse

service_account = clickhouse.Service(
    "example",
    region = "us-central1",
    cloud_provider = "gcp",
    tier = "development",
    password =  "1234",
    ip_accesses = [{
        "source": "0.0.0.0",
        "description": "Test IP"
      }]
)
```

{{% /choosable %}}
{{< /chooser >}}

## Issues

This is a community maintained provider. Please file issues and feature requests here:

[pulumiverse/pulumi-clickhouse](https://github.com/pulumiverse/pulumi-clickhouse/issues)