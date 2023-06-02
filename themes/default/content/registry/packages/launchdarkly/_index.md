---
title: Launch Darkly
meta_desc: Provides an overview of the Launch Darkly Provider for Pulumi.
layout: package
---

The Launch Darkly provider for Pulumi can be used to provision any of the cloud resources available in [Launch Darkly](https://launchdarkly.com/).
The Launch Darkly provider must be configured with credentials to deploy and update resources in Launch Darkly.

## Example

{{< chooser language "typescript,python" >}}
{{% choosable language typescript %}}

```typescript
import * as launchdarkly from "@lbrlabs/pulumi-lauchdarkly";

const sa = new launchdarkly.AccessToken("example", {
    role: "Reader",
})

```

{{% /choosable %}}
{{% choosable language python %}}

```python
import lbrlabs_pulumi_launchdarkly as launchdarkly

service_account = launchdarkly.ServiceAccount(
    "example",
    role="Reader"
)
```

{{% /choosable %}}
{{< /chooser >}}
