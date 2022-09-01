---
title: Gandi
meta_desc: Provides an overview of the Gandi Provider for Pulumi.
layout: overview
---

The Gandi provider for Pulumi can be used to provision any of the resources available in [Gandi](https://gandi.net/).

## Example

{{< chooser language "typescript" >}}

{{% choosable language typescript %}}

```typescript
import * as gandi from "@pulumiverse/gandi";

const domain_nameservers = new gandi.domains.Nameservers("my-domain", {
    domain: "my-domain.com",
    servers: [
        "1.1.1.1",
        "2.2.2.2",
    ],
});
```

{{% /choosable %}}
{{< /chooser >}}

> You could find more complete and detailed examples in the [pulumi-gandi repository](https://github.com/pulumiverse/pulumi-gandi/tree/main/examples)
