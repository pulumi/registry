---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-coreweave/v1.0.1/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/pulumi/pulumi-coreweave/blob/v1.0.1/docs/_index.md
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: CoreWeave Provider
meta_desc: Provides an overview on how to configure the Pulumi CoreWeave provider.
layout: package
---

## Installation

The CoreWeave provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/coreweave`](https://www.npmjs.com/package/@pulumi/coreweave)
* Python: [`pulumi-coreweave`](https://pypi.org/project/pulumi-coreweave/)
* Go: [`github.com/pulumi/pulumi-coreweave/sdk/go/coreweave`](https://github.com/pulumi/pulumi-coreweave)
* .NET: [`Pulumi.Coreweave`](https://www.nuget.org/packages/Pulumi.Coreweave)
* Java: [`com.pulumi/coreweave`](https://central.sonatype.com/artifact/com.pulumi/coreweave)

## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as coreweave from "@pulumi/coreweave";

const _default = new coreweave.NetworkingVpc("default", {
    name: "default",
    zone: "US-EAST-04A",
    vpcPrefixes: [
        {
            name: "pod cidr",
            value: "10.0.0.0/13",
        },
        {
            name: "service cidr",
            value: "10.16.0.0/22",
        },
        {
            name: "internal lb cidr",
            value: "10.32.4.0/22",
        },
    ],
});
const defaultCksCluster = new coreweave.CksCluster("default", {
    name: "default",
    version: "v1.35",
    zone: "US-EAST-04A",
    vpcId: _default.id,
    "public": true,
    podCidrName: "pod cidr",
    serviceCidrName: "service cidr",
    internalLbCidrNames: ["internal lb cidr"],
});
```

{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import pulumi_coreweave as coreweave

default = coreweave.NetworkingVpc("default",
    name="default",
    zone="US-EAST-04A",
    vpc_prefixes=[
        {
            "name": "pod cidr",
            "value": "10.0.0.0/13",
        },
        {
            "name": "service cidr",
            "value": "10.16.0.0/22",
        },
        {
            "name": "internal lb cidr",
            "value": "10.32.4.0/22",
        },
    ])
default_cks_cluster = coreweave.CksCluster("default",
    name="default",
    version="v1.35",
    zone="US-EAST-04A",
    vpc_id=default.id,
    public=True,
    pod_cidr_name="pod cidr",
    service_cidr_name="service cidr",
    internal_lb_cidr_names=["internal lb cidr"])
```

{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using CoreWeave = Pulumi.CoreWeave;

return await Deployment.RunAsync(() =>
{
    var @default = new CoreWeave.NetworkingVpc("default", new()
    {
        Name = "default",
        Zone = "US-EAST-04A",
        VpcPrefixes = new[]
        {
            new CoreWeave.Inputs.NetworkingVpcVpcPrefixArgs
            {
                Name = "pod cidr",
                Value = "10.0.0.0/13",
            },
            new CoreWeave.Inputs.NetworkingVpcVpcPrefixArgs
            {
                Name = "service cidr",
                Value = "10.16.0.0/22",
            },
            new CoreWeave.Inputs.NetworkingVpcVpcPrefixArgs
            {
                Name = "internal lb cidr",
                Value = "10.32.4.0/22",
            },
        },
    });

    var defaultCksCluster = new CoreWeave.CksCluster("default", new()
    {
        Name = "default",
        Version = "v1.35",
        Zone = "US-EAST-04A",
        VpcId = @default.Id,
        Public = true,
        PodCidrName = "pod cidr",
        ServiceCidrName = "service cidr",
        InternalLbCidrNames = new[]
        {
            "internal lb cidr",
        },
    });

});

```

{{% /choosable %}}
{{% choosable language go %}}
```go
package main

import (
	"github.com/pulumi/pulumi-coreweave/sdk/go/coreweave"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_default, err := coreweave.NewNetworkingVpc(ctx, "default", &coreweave.NetworkingVpcArgs{
			Name: pulumi.String("default"),
			Zone: pulumi.String("US-EAST-04A"),
			VpcPrefixes: coreweave.NetworkingVpcVpcPrefixArray{
				&coreweave.NetworkingVpcVpcPrefixArgs{
					Name:  pulumi.String("pod cidr"),
					Value: pulumi.String("10.0.0.0/13"),
				},
				&coreweave.NetworkingVpcVpcPrefixArgs{
					Name:  pulumi.String("service cidr"),
					Value: pulumi.String("10.16.0.0/22"),
				},
				&coreweave.NetworkingVpcVpcPrefixArgs{
					Name:  pulumi.String("internal lb cidr"),
					Value: pulumi.String("10.32.4.0/22"),
				},
			},
		})
		if err != nil {
			return err
		}
		_, err = coreweave.NewCksCluster(ctx, "default", &coreweave.CksClusterArgs{
			Name:            pulumi.String("default"),
			Version:         pulumi.String("v1.35"),
			Zone:            pulumi.String("US-EAST-04A"),
			VpcId:           _default.ID(),
			Public:          pulumi.Bool(true),
			PodCidrName:     pulumi.String("pod cidr"),
			ServiceCidrName: pulumi.String("service cidr"),
			InternalLbCidrNames: pulumi.StringArray{
				pulumi.String("internal lb cidr"),
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
{{% choosable language yaml %}}
```yaml
resources:
  default:
    type: coreweave:NetworkingVpc
    properties:
      name: default
      zone: US-EAST-04A
      vpcPrefixes:
        - name: pod cidr
          value: 10.0.0.0/13
        - name: service cidr
          value: 10.16.0.0/22
        - name: internal lb cidr
          value: 10.32.4.0/22
  defaultCksCluster:
    type: coreweave:CksCluster
    name: default
    properties:
      name: default
      version: v1.35
      zone: US-EAST-04A
      vpcId: ${default.id}
      public: true
      podCidrName: pod cidr
      serviceCidrName: service cidr
      internalLbCidrNames:
        - internal lb cidr
```

{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.coreweave.NetworkingVpc;
import com.pulumi.coreweave.NetworkingVpcArgs;
import com.pulumi.coreweave.inputs.NetworkingVpcVpcPrefixArgs;
import com.pulumi.coreweave.CksCluster;
import com.pulumi.coreweave.CksClusterArgs;
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        var default_ = new NetworkingVpc("default", NetworkingVpcArgs.builder()
            .name("default")
            .zone("US-EAST-04A")
            .vpcPrefixes(
                NetworkingVpcVpcPrefixArgs.builder()
                    .name("pod cidr")
                    .value("10.0.0.0/13")
                    .build(),
                NetworkingVpcVpcPrefixArgs.builder()
                    .name("service cidr")
                    .value("10.16.0.0/22")
                    .build(),
                NetworkingVpcVpcPrefixArgs.builder()
                    .name("internal lb cidr")
                    .value("10.32.4.0/22")
                    .build())
            .build());

        var defaultCksCluster = new CksCluster("defaultCksCluster", CksClusterArgs.builder()
            .name("default")
            .version("v1.35")
            .zone("US-EAST-04A")
            .vpcId(default_.id())
            .public_(true)
            .podCidrName("pod cidr")
            .serviceCidrName("service cidr")
            .internalLbCidrNames("internal lb cidr")
            .build());

    }
}
```

{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

- `endpoint` (String) CoreWeave API Endpoint. This can also be set via the COREWEAVE_API_ENDPOINT environment variable, which takes precedence. Defaults to `https://api.coreweave.com/`
- `httpTimeout` (String) Timeout duration for the HTTP client to use. This can also be set via the COREWEAVE_HTTP_TIMEOUT environment variable, which takes precedence. If unset, defaults to 10 seconds
- `s3Endpoint` (String) CoreWeave S3 Endpoint, used for CoreWeave Object Storage. This can also be set via the COREWEAVE_S3_ENDPOINT environment variable, which takes precedence. Defaults to `https://cwobject.com`
- `token` (String, Sensitive) CoreWeave API Token in the form `CW-SECRET-<secret>`. This can also be set via the COREWEAVE_API_TOKEN environment variable, which takes precedence.