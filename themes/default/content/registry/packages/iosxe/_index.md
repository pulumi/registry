---
title: Cisco IOS XE Resource Provider
meta_desc: Provides an overview of the Cisco IOS XE Provider for Pulumi.
layout: package
---

The Cisco IOS XE provider for Pulumi can be used to provision any of the cloud resources available on Cisco IOS XE.
The provider must be configured with credentials to deploy and update resources on a Cisco IOS XE Device.

## Example

{{< chooser language "typescript,python" >}}
{{% choosable language typescript %}}

```typescript
import * as iosxe from "@lbrlabs/pulumi-iosxe";
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import lbrlabs_pulumi_iosxe as iosxe
```

{{% /choosable %}}
{{< /chooser >}}
