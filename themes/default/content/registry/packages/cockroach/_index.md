---
title: CockroachDB Cloud
meta_desc: Provides an overview of the CockroachDB provider for Pulumi.
layout: package
---

The Cockroach provider for Pulumi can be used to provision any of the cloud resources available in [CockroachDB](https://www.cockroachlabs.com) or a self hosted CockroachDB instance.

The Cockroach provider must be configured with credentials to deploy and update resources in CockroachDB.

## Example

{{< chooser language "typescript,python" >}}
{{% choosable language typescript %}}

```typescript
import * as cockroach from "@lbrlabs/pulumi-cockroach";

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
import lbrlabs_pulumi_cockroach as cockroach

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
