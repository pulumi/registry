---
title: Kubernetes CoreDNS
meta_desc: Use Pulumi's Component for managing CoreDNS installations using infrastructure as code.
layout: overview
---

Easily manage CoreDNS installations as a package available in all Pulumi languages.

## Example

{{< chooser language "typescript,python" >}}

{{% choosable language typescript %}}

```typescript
import * as coredns from "@pulumi/kubernetes-coredns";

const dns = new coredns.CoreDNS("dns", {
    servers: [{
        zones: [
            {
                zone: "hello.world.",
                scheme: "tls://",
            },
            {
                zone: "foo.bar.",
                scheme: "dns://",
                use_tcp: true,
            },
        ],
        port: 12345,
        plugins: [
            {
                name: "kubernetes",
                parameters: "foo bar",
                configBlock: "hello world\nfoo bar",
            },
        ],
    }],
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi
from pulumi_kubernetes_coredns import CoreDNS, CoreDNSServerArgs, CoreDNSServerZoneArgs, CoreDNSServerPluginArgs

dns = CoreDNS('dns',
              servers=[
                  CoreDNSServerArgs(
                      zones=[
                          CoreDNSServerZoneArgs(
                              zone='hello.world.',
                              scheme='tls://',
                          ),
                          CoreDNSServerZoneArgs(
                              zone='foo.bar.',
                              scheme='dns://',
                              use_tcp=True,
                          ),
                      ],
                      port=12345,
                      plugins=[
                          CoreDNSServerPluginArgs(
                              name='kubernetes',
                              parameters='foo bar',
                              config_block='hello world\nfoo bar',
                          ),
                      ],
                  ),
              ],
              )
```

{{% /choosable %}}

{{< /chooser >}}
