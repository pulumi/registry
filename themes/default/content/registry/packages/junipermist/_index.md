---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-junipermist/v0.1.1/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Juniper Mist
meta_desc: The Juniper Mist provider for Pulumi can be used to provision any of the cloud resources available in Juniper Mist.
layout: package
---

The Juniper Mist Provider allows Pulumi to manage Juniper Mist Organizations.

The Juniper Mist provider must be configured with credentials to deploy and update resources; see [Installation & Configuration](./installation-configuration/) for instructions.

## Example

{{< chooser language "typescript,python" >}}

{{% choosable language typescript %}}

```typescript
import * as junipermist from "@pulumi/juniper-mist";

new junipermist.site.Wlan("wlan-one", {
  ssid: "wlan_one",
  siteId: site.id,
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_juniper_mist as junipermist

wlan = junipermist.site.Wlan(
    "wlan-one", site_id=site.id, ssid="wlan_one"
)
```

{{% /choosable %}}

{{< /chooser >}}
