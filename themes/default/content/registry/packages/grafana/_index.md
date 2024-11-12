---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumiverse/pulumi-grafana/v0.7.0/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Grafana Cloud
meta_desc: Provides an overview of the Grafana Provider for Pulumi.
layout: package
---

The Grafana provider for Pulumi can be used to provision any of the cloud resources available in [Grafana Cloud](https://grafana.com/products/cloud/) or a self hosted Grafana instance
The Grafana provider must be configured with credentials to deploy and update resources in Grafana.

## Example

{{< chooser language "typescript,python" >}}
{{% choosable language typescript %}}

```typescript
import * as grafana from "@pulumiverse/grafana";

const sa = new grafana.ServiceAccount("example", {
    role: "Viewer",
})

```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumiverse_grafana as grafana

service_account = grafana.ServiceAccount(
    "example",
    role="Viewer"
)
```

{{% /choosable %}}
{{< /chooser >}}

## Issues

This is a community maintained provider. Please file issues and feature requests here:

[pulumiverse/pulumi-grafana](https://github.com/pulumiverse/pulumi-grafana/issues)
