---
title: Grafana Cloud
meta_desc: Provides an overview of the Grafana Provider for Pulumi.
layout: overview
---

The Grafana provider for Pulumi can be used to provision any of the cloud resources available in [Grafana Cloud](https://grafana.com/products/cloud/) or a self hosted Grafana instance
The Grafana provider must be configured with credentials to deploy and update resources in Grafana.

## Example

{{< chooser language "typescript,python" >}}
{{% choosable language typescript %}}

```typescript
import * as grafana from "@lbrlabs/pulumi-grafana";

const sa = new grafana.ServiceAccount("example", {
    role: "Viewer",
})

```

{{% /choosable %}}
{{% choosable language python %}}

```python
import lbrlabs_pulumi_grafana as grafana

service_account = grafana.ServiceAccount(
    "example",
    role="Viewer"
)
```

{{% /choosable %}}
{{< /chooser >}}
