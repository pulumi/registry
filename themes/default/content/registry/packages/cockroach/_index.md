---
title: Cockroach DB Cloud
meta_desc: Provides an overview of the Cockroach DB Provider for Pulumi.
layout: overview
---

The Cockroach provider for Pulumi can be used to provision any of the cloud resources available in [Cockroach DB](https://www.cockroachlabs.com) or a self hosted Cockroach DB instance
The Cockroach provider must be configured with credentials to deploy and update resources in Cockroach DB.

## Example

{{< chooser language "typescript,python" >}}
{{% choosable language typescript %}}

```typescript
import * as cockroach from "@lbrlabs/pulumi-cockroach";

const cluster = new cockroach.Cluster("example", {
  cloudProvider: "aws",
  name: "cockroach-provider-ts",
  regions: [
    {
      name: "us-west-2",
    },
  ],
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import lbrlabs_pulumi_cockroach as cockroach

cluster = cockroach.Cluster(
    "example",
    cloud_provider="aws",
    name="cockroach-provider-py",
)
```

{{% /choosable %}}
{{< /chooser >}}
