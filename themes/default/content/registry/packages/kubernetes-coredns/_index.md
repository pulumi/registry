---
title: Kubernetes CoreDNS
meta_desc: Use Pulumi's Component for managing CoreDNS installations using infrastructure as code.
layout: package
---

Easily manage CoreDNS installations as a package available in all Pulumi languages.

## Example

{{< chooser language "typescript,python,go,csharp,java,yaml,hcl" >}}

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

{{% choosable language go %}}

```go
import (
	kubernetescoredns "github.com/pulumi/pulumi-kubernetes-coredns/sdk/go/kubernetes-coredns"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := kubernetescoredns.NewCoreDNS(ctx, "dns", &kubernetescoredns.CoreDNSArgs{
			Servers: kubernetescoredns.CoreDNSServerArray{
				&kubernetescoredns.CoreDNSServerArgs{
					Zones: kubernetescoredns.CoreDNSServerZoneArray{
						&kubernetescoredns.CoreDNSServerZoneArgs{
							Zone:   pulumi.String("hello.world."),
							Scheme: pulumi.String("tls://"),
						},
						&kubernetescoredns.CoreDNSServerZoneArgs{
							Zone:    pulumi.String("foo.bar."),
							Scheme:  pulumi.String("dns://"),
							Use_tcp: pulumi.Bool(true),
						},
					},
					Port: pulumi.Int(12345),
					Plugins: kubernetescoredns.CoreDNSServerPluginArray{
						&kubernetescoredns.CoreDNSServerPluginArgs{
							Name:        pulumi.String("kubernetes"),
							Parameters:  pulumi.String("foo bar"),
							ConfigBlock: pulumi.String("hello world\nfoo bar"),
						},
					},
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
```

{{% /choosable %}}

{{% choosable language csharp %}}

```csharp
using Pulumi;
using KubernetesCoreDNS = Pulumi.KubernetesCoreDNS;

await Deployment.RunAsync(() =>
{
    var dns = new KubernetesCoreDNS.CoreDNS("dns", new()
    {
        Servers = new[]
        {
            new KubernetesCoreDNS.Inputs.CoreDNSServerArgs
            {
                Zones = new[]
                {
                    new KubernetesCoreDNS.Inputs.CoreDNSServerZoneArgs
                    {
                        Zone = "hello.world.",
                        Scheme = "tls://",
                    },
                    new KubernetesCoreDNS.Inputs.CoreDNSServerZoneArgs
                    {
                        Zone = "foo.bar.",
                        Scheme = "dns://",
                        Use_tcp = true,
                    },
                },
                Port = 12345,
                Plugins = new[]
                {
                    new KubernetesCoreDNS.Inputs.CoreDNSServerPluginArgs
                    {
                        Name = "kubernetes",
                        Parameters = "foo bar",
                        ConfigBlock = @"hello world
foo bar",
                    },
                },
            },
        },
    });
});
```

{{% /choosable %}}

{{% choosable language java %}}

```java
import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.kubernetescoredns.CoreDNS;
import com.pulumi.kubernetescoredns.CoreDNSArgs;
import com.pulumi.kubernetescoredns.inputs.CoreDNSServerArgs;
import com.pulumi.kubernetescoredns.inputs.CoreDNSServerZoneArgs;
import com.pulumi.kubernetescoredns.inputs.CoreDNSServerPluginArgs;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        var dns = new CoreDNS("dns", CoreDNSArgs.builder()
            .servers(CoreDNSServerArgs.builder()
                .zones(
                    CoreDNSServerZoneArgs.builder()
                        .zone("hello.world.")
                        .scheme("tls://")
                        .build(),
                    CoreDNSServerZoneArgs.builder()
                        .zone("foo.bar.")
                        .scheme("dns://")
                        .use_tcp(true)
                        .build())
                .port(12345)
                .plugins(CoreDNSServerPluginArgs.builder()
                    .name("kubernetes")
                    .parameters("foo bar")
                    .configBlock("hello world\nfoo bar")
                    .build())
                .build())
            .build());
    }
}
```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
resources:
  dns:
    type: kubernetes-coredns:CoreDNS
    properties:
      servers:
        - zones:
            - zone: hello.world.
              scheme: tls://
            - zone: foo.bar.
              scheme: dns://
              use_tcp: true
          port: 12345
          plugins:
            - name: kubernetes
              parameters: foo bar
              configBlock: |-
                hello world
                foo bar
```

{{% /choosable %}}

{{% choosable language hcl %}}

```hcl
terraform {
  required_providers {
    kubernetes-coredns = {
      source = "pulumi/kubernetes-coredns"
    }
  }
}

resource "kubernetes-coredns_coredns" "dns" {
  servers {
    zones {
      zone   = "hello.world."
      scheme = "tls://"
    }
    zones {
      zone    = "foo.bar."
      scheme  = "dns://"
      use_tcp = true
    }
    port = 12345
    plugins {
      name         = "kubernetes"
      parameters   = "foo bar"
      config_block = "hello world\nfoo bar"
    }
  }
}
```

{{% /choosable %}}

{{< /chooser >}}
