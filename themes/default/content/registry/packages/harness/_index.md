---
title: Harness
meta_desc: Provides an overview of the Harness Provider for Pulumi.
layout: package
---

The Harness provider for Pulumi can be used to provision any of the cloud resources available in [Scaleway](https://www.harness.io/).
The Harness provider must be configured with credentials to deploy and update resources in Harness.

## Example

{{< chooser language "typescript" >}}
{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as harness from "@pulumi/harness";

const secret = new harness.platform.SecretText("example", {
    identifier: "ts_example",
    valueType: "Inline",
    value: "correct-horse-battery-stable",
    secretManagerIdentifier: "harnessSecretManager",
})
```
{{% /choosable %}}
{{< /chooser >}}
