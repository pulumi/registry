---
# WARNING: this file was fetched from https://raw.githubusercontent.com/johnneerdael/pulumi-netskope-publisher/v0.3.4/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/johnneerdael/pulumi-netskope-publisher/blob/v0.3.4/docs/_index.md
title: Netskope Publisher
meta_desc: Pulumi components for provisioning Netskope Private Access Publishers.
layout: package
---

The Netskope Publisher package provides Pulumi component resources for
provisioning Netskope Private Access Publishers on AWS, Azure, Google
Cloud, Kubernetes, vSphere, ESXi, Hcloud, Nutanix, OpenStack, OVH,
Scaleway, OCI, Alicloud, Proxmox VE, DigitalOcean, Vultr, Exoscale,
UpCloud, Stackit, Equinix Metal, Outscale, OpenTelekomCloud,
TencentCloud, and Yandex. The package mirrors the
Terraform module pattern:
register or reuse Netskope publisher records, generate per-publisher
cloud-init, and create the virtual machines that run the publisher
appliance. On Kubernetes, it installs the
`kubernetes-netskope-publisher` Helm chart.

## Components

- `AwsPublisher` creates EC2-backed publishers.
- `AzurePublisher` creates Azure virtual machine backed publishers.
- `GcpPublisher` creates Google Compute Engine backed publishers.
- `KubernetesPublisher` installs Helm chart backed publishers.
- `VspherePublisher` creates vSphere virtual machine backed publishers.
- `EsxiPublisher` creates direct-host ESXi backed publishers.
- `HcloudPublisher`, `NutanixPublisher`, `OpenstackPublisher`,
  `OvhPublisher`, `ScalewayPublisher`, `OciPublisher`, and
  `AlicloudPublisher` create Ubuntu 22.04 bootstrap-mode publishers.
- `ProxmoxvePublisher` clones a Proxmox VE Ubuntu template and attaches
  bootstrap cloud-init through a snippets datastore.
- `DigitaloceanPublisher`, `VultrPublisher`, `ExoscalePublisher`,
  `UpcloudPublisher`, `StackitPublisher`, `EquinixPublisher`,
  `OutscalePublisher`, `OpentelekomcloudPublisher`,
  `TencentcloudPublisher`, and `YandexPublisher` create Ubuntu 22.04
  bootstrap-mode publishers.
- `HypervPublisher` is experimental and remains opt-in.

Each component accepts either Netskope tenant credentials for automatic
publisher registration or pre-created registration tokens keyed by
publisher name.

The Java SDK is also published as `com.pulumi:netskope-publisher` from
the configured Maven-compatible repository, defaulting to GitHub
Packages for this repository.

## TypeScript example

```typescript
import { AwsPublisher } from "@johninnl/pulumi-netskope-publisher";

const publisher = new AwsPublisher("publisher", {
  namePrefix: "pub-eu",
  replicas: 2,
  tenantUrl: "https://example.goskope.com",
  apiToken: "ns-api-token",
  subnetId: "subnet-0123456789abcdef0",
  securityGroupIds: ["sg-0123456789abcdef0"],
});
```

## Python example

```python
from pulumi_netskope_publisher import AwsPublisher

publisher = AwsPublisher(
    "publisher",
    name_prefix="pub-eu",
    replicas=2,
    tenant_url="https://example.goskope.com",
    api_token="ns-api-token",
    subnet_id="subnet-0123456789abcdef0",
    security_group_ids=["sg-0123456789abcdef0"],
)
```

## C# example

```csharp
using Pulumi;
using Pulumi.NetskopePublisher;

return await Deployment.RunAsync(() =>
{
    var publisher = new AwsPublisher("publisher", new AwsPublisherArgs
    {
        NamePrefix = "pub-eu",
        Replicas = 2,
        TenantUrl = "https://example.goskope.com",
        ApiToken = "ns-api-token",
        SubnetId = "subnet-0123456789abcdef0",
        SecurityGroupIds = new[] { "sg-0123456789abcdef0" },
    });
});
```

## Go example

```go
package main

import (
	"github.com/johnneerdael/pulumi-netskope-publisher/sdk/go/netskopepublisher"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := netskopepublisher.NewAwsPublisher(ctx, "publisher", &netskopepublisher.AwsPublisherArgs{
			NamePrefix: pulumi.StringPtr("pub-eu"),
			Replicas: pulumi.IntPtr(2),
			TenantUrl: pulumi.StringPtr("https://example.goskope.com"),
			ApiToken: pulumi.StringPtr("ns-api-token"),
			SubnetId: pulumi.String("subnet-0123456789abcdef0"),
			SecurityGroupIds: pulumi.StringArray{
				pulumi.String("sg-0123456789abcdef0"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
```

## Java example

```java
package myproject;

import com.pulumi.Pulumi;
import com.pulumi.netskopepublisher.AwsPublisher;
import com.pulumi.netskopepublisher.AwsPublisherArgs;

public class App {
    public static void main(String[] args) {
        Pulumi.run(ctx -> {
            new AwsPublisher("publisher", AwsPublisherArgs.builder()
                .namePrefix("pub-eu")
                .replicas(2)
                .tenantUrl("https://example.goskope.com")
                .apiToken("ns-api-token")
                .subnetId("subnet-0123456789abcdef0")
                .securityGroupIds("sg-0123456789abcdef0")
                .build());
        });
    }
}
```

Provider-specific examples are available in the repository under
`examples/`.
