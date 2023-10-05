---
title: Cisco NX OS Resource Provider
meta_desc: Provides an overview of the Cisco NX OS Provider for Pulumi.
layout: package
---

The Cisco NX OS provider for Pulumi can be used to provision any of the cloud resources available on Cisco NX OS.
The provider must be configured with credentials to deploy and update resources on a Cisco NX OS Device.

## Example

{{< chooser language "typescript,python" >}}
{{% choosable language typescript %}}

```typescript
import * as nxos from "@lbrlabs/pulumi-nxos";
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import lbrlabs_pulumi_nxos as nxos
```

{{% /choosable %}}
{{< /chooser >}}
