---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumiverse/pulumi-cockroach/v0.9.4/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
title: CockroachDB Cloud
meta_desc: Provides an overview of the CockroachDB provider for Pulumi.
layout: package
---

The Cockroach provider for Pulumi can be used to provision any of the cloud resources available in [CockroachDB](https://www.cockroachlabs.com) or a self hosted CockroachDB instance.

The Cockroach provider must be configured with credentials to deploy and update resources in CockroachDB.

**NOTE** this is a community package in [need of new maintainers](https://github.com/pulumiverse/pulumi-cockroach/issues/75).

## Example

{{< chooser language "typescript,python" >}}
{{% choosable language typescript %}}

```typescript
import * as cockroach from "@pulumiverse/cockroach";

const cluster = new cockroach.Cluster("example", {
    cloudProvider: "AWS",
    name: "cockroach-provider-ts",
    regions: [
        {
            name: "us-west-2",
        },
    ],
    serverless: {
      spendLimit: 0,
    },
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumiverse_cockroach as cockroach

cluster = cockroach.Cluster(
    "example",
    cloud_provider="AWS",
    name="cockroach-provider-py",
    regions=[
        cockroach.ClusterRegionArgs(
            name="us-west-2",
        ),
    ],
    serverless=cockroach.ClusterServerlessArgs(
        spend_limit=0,
    ),
)
```

{{% /choosable %}}
{{< /chooser >}}

## Issues

This is a community maintained provider. Please file issues and feature requests here:

[pulumiverse/pulumi-cockroach](https://github.com/pulumiverse/pulumi-cockroach/issues)
