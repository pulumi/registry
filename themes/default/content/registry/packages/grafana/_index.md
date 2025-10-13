---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumiverse/pulumi-grafana/v2.11.0/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/pulumiverse/pulumi-grafana/blob/v2.11.0/docs/_index.md
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

const stack = new grafana.cloud.Stack("stack", {
    name: "<stack-name>",
    slug: "<stack-name>",
    // Regions https://grafana.com/api/stack-regions
    regionSlug: "<region>",
});

const serviceAccount = new grafana.cloud.StackServiceAccount("serviceAccount", {
    stackSlug: stack.slug,
    name: "Service account",
    role: "Admin",
});

const serviceAccountToken = new grafana.cloud.StackServiceAccountToken("serviceAccountToken", {
    stackSlug: stack.slug,
    name: "Service account token",
    serviceAccountId: serviceAccount.id,
});

const serviceAccountProvider = new grafana.Provider("serviceAccountProvider", {
    auth: serviceAccountToken.key,
    url: stack.url.apply((url) => url as string),
});

const testFolder = new grafana.oss.Folder(
    "folder",
    { title: "Folder title" },
    { provider: serviceAccountProvider }
);

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
