---
title: "Enhanced Region Support"
h1: "Enhanced Region Support"
meta_desc: Practitioner level instructions for utilizing the new region functionality in AWS v7
layout: package
---

Version 7.0.0 of the Pulumi AWS Provider adds `region` to most resources making it significantly easier to manage infrastructure across AWS Regions without requiring multiple provider configurations.

<!-- mtoc-start -->

* [What's new](#whats-new)
* [What's _not_ changing](#whats-_not_-changing)
* [Can I use `region` in every resource?](#can-i-use-region-in-every-resource)
* [Why make this change](#why-make-this-change)
* [How `region` works](#how-region-works)
* [Migrating from multiple provider configurations](#migrating-from-multiple-provider-configurations)
* [Before and after examples using `region`](#before-and-after-examples-using-region)
* [Non-region-aware resources](#non-region-aware-resources)

<!-- mtoc-end -->

## What's new

As of v7.0.0, most existing resources and data sources are now [Region-aware](#non-region-aware-resources), meaning they support a new top-level `region`. This allows you to manage a resource in a Region different from the one specified in the provider configuration without requiring multiple provider definitions. See [How `region` works](#how-region-works) for details.

For example, if your provider is configured for `us-east-1`, you can now manage a VPC in `us-west-2` without defining an additional provider:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}

{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const peer = new aws.ec2.Vpc("peer", {
    region: "us-west-2",
    cidrBlock: "10.1.0.0/16",
});

```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi
import pulumi_aws as aws

peer = aws.ec2.Vpc("peer",
    region="us-west-2",
    cidr_block="10.1.0.0/16")

```

{{% /choosable %}}

{{% choosable language go %}}

```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2.NewVpc(ctx, "peer", &ec2.VpcArgs{
			Region:    pulumi.String("us-west-2"),
			CidrBlock: pulumi.String("10.1.0.0/16"),
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
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var peer = new Aws.Ec2.Vpc("peer", new()
    {
        Region = "us-west-2",
        CidrBlock = "10.1.0.0/16",
    });

});


```

{{% /choosable %}}

{{% choosable language java %}}

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.ec2.Vpc;
import com.pulumi.aws.ec2.VpcArgs;
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
        var peer = new Vpc("peer", VpcArgs.builder()
            .region("us-west-2")
            .cidrBlock("10.1.0.0/16")
            .build());

    }
}

```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
name: example
runtime: yaml
resources:
  peer:
    type: aws:ec2:Vpc
    properties:
      region: us-west-2
      cidrBlock: 10.1.0.0/16
```

{{% /choosable %}}

{{< /chooser >}}

## What's _not_ changing

_Pre-v7.0.0 configurations that use provider configurations per Region remain valid in v7.0.0 and are not deprecated._

You can still define the Region at the provider level using any of the existing methods—for example, through the AWS [config file](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html), [provider configuration](https://www.pulumi.com/registry/packages/aws/installation-configuration/#configuration), [environment variables](https://www.pulumi.com/registry/packages/aws/installation-configuration/#set-credentials-as-environment-variables), [shared configuration files](https://www.pulumi.com/registry/packages/aws/installation-configuration/#create-a-shared-credentials-file), or explicitly using the `provider`'s [`region`](https://www.pulumi.com/registry/packages/aws/installation-configuration/#configuration-options).

## Can I use `region` in every resource?

No. While most resources are now Region-aware, there are exceptions. These include a few resources that already had a `region` and resources that are inherently global. See [Non–region-aware resources](#non-region-aware-resources).

## Why make this change

Before version 7.0.0, managing infrastructure across multiple Regions required a separate provider configuration for each Region. This approach led to complex and repetitive configurations, especially for large infrastructures—AWS currently operates in [36 Regions](https://aws.amazon.com/about-aws/global-infrastructure/), with more announced. Additionally, each provider configuration adds overhead in terms of memory and compute resources.

See the [examples](#before-and-after-examples-using-region) below for a comparison of configurations before and after introducing `region`.

## How `region` works

The new top-level `region` is _Optional_, and defaults to the Region specified in the provider configuration. Its value is validated to ensure it belongs to the configured [partition](https://docs.aws.amazon.com/whitepapers/latest/aws-fault-isolation-boundaries/partitions.html). **Changing the value of `region` will force resource replacement.**

To [import](https://www.pulumi.com/docs/iac/adopting-pulumi/import/) a resource in a specific Region, append `@<region>` to the [import ID](https://www.pulumi.com/docs/iac/adopting-pulumi/import/#how-resource-import-works)—for example:

```sh
pulumi import aws:ec2/vpc:Vpc test_vpc vpc-a01106c2@eu-west-1
```

## Migrating from multiple provider configurations

To migrate from a separate provider configuration for each Region to a single provider configuration and per-resource `region` values you must ensure that Pulumi state is refreshed:

1. Upgrade to v7.0.0 following the upgrade guide.
2. Modify the affected resource configurations, replacing the [`provider` resource option](https://www.pulumi.com/docs/iac/concepts/options/provider/) with a `region` argument
3. Run Pulumi with refresh -- `pulumi up --refresh`

## Before and after examples using `region`

### Cross-region VPC peering

**Before, Pre-v7.0.0**

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}

{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const awsPeerProvider = new aws.Provider("awsPeerProvider", {region: "us-west-2"});
const peerIdentity = aws.getCallerIdentityOutput({}, {
    provider: awsPeerProvider,
});
const awsProvider = new aws.Provider("awsProvider", {region: "us-east-1"});
const main = new aws.ec2.Vpc("main", {cidrBlock: "10.0.0.0/16"}, {
    provider: awsProvider,
});
const peer = new aws.ec2.Vpc("peer", {cidrBlock: "10.1.0.0/16"}, {
    provider: awsPeerProvider,
});
const peerConnection = new aws.ec2.VpcPeeringConnection("peerConnection", {
    vpcId: main.id,
    peerVpcId: peer.id,
    peerOwnerId: peerIdentity.apply(peerIdentity => peerIdentity.accountId),
    peerRegion: "us-west-2",
    autoAccept: false,
}, {
    provider: awsProvider,
});
const peerAccepter = new aws.ec2.VpcPeeringConnectionAccepter("peerAccepter", {
    vpcPeeringConnectionId: peerConnection.id,
    autoAccept: true,
}, {
    provider: awsPeerProvider,
});

```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi
import pulumi_aws as aws

aws_peer_provider = aws.Provider("awsPeerProvider", region="us-west-2")
peer_identity = aws.get_caller_identity_output(, opts=pulumi.InvokeOutputOptions(provider=aws_peer_provider))
aws_provider = aws.Provider("awsProvider", region="us-east-1")
main = aws.ec2.Vpc("main", cidr_block="10.0.0.0/16",
opts = pulumi.ResourceOptions(provider=aws_provider))
peer = aws.ec2.Vpc("peer", cidr_block="10.1.0.0/16",
opts = pulumi.ResourceOptions(provider=aws_peer_provider))
peer_connection = aws.ec2.VpcPeeringConnection("peerConnection",
    vpc_id=main.id,
    peer_vpc_id=peer.id,
    peer_owner_id=peer_identity.account_id,
    peer_region="us-west-2",
    auto_accept=False,
    opts = pulumi.ResourceOptions(provider=aws_provider))
peer_accepter = aws.ec2.VpcPeeringConnectionAccepter("peerAccepter",
    vpc_peering_connection_id=peer_connection.id,
    auto_accept=True,
    opts = pulumi.ResourceOptions(provider=aws_peer_provider))

```

{{% /choosable %}}

{{% choosable language go %}}

```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		awsPeerProvider, err := aws.NewProvider(ctx, "awsPeerProvider", &aws.ProviderArgs{
			Region: pulumi.String("us-west-2"),
		})
		if err != nil {
			return err
		}
		peerIdentity := aws.GetCallerIdentityOutput(ctx, aws.GetCallerIdentityOutputArgs{}, pulumi.Provider(awsPeerProvider))
		awsProvider, err := aws.NewProvider(ctx, "awsProvider", &aws.ProviderArgs{
			Region: pulumi.String("us-east-1"),
		})
		if err != nil {
			return err
		}
		main, err := ec2.NewVpc(ctx, "main", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.0.0.0/16"),
		}, pulumi.Provider(awsProvider))
		if err != nil {
			return err
		}
		peer, err := ec2.NewVpc(ctx, "peer", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.1.0.0/16"),
		}, pulumi.Provider(awsPeerProvider))
		if err != nil {
			return err
		}
		peerConnection, err := ec2.NewVpcPeeringConnection(ctx, "peerConnection", &ec2.VpcPeeringConnectionArgs{
			VpcId:     main.ID(),
			PeerVpcId: peer.ID(),
			PeerOwnerId: pulumi.String(peerIdentity.ApplyT(func(peerIdentity aws.GetCallerIdentityResult) (*string, error) {
				return &peerIdentity.AccountId, nil
			}).(pulumi.StringPtrOutput)),
			PeerRegion: pulumi.String("us-west-2"),
			AutoAccept: pulumi.Bool(false),
		}, pulumi.Provider(awsProvider))
		if err != nil {
			return err
		}
		_, err = ec2.NewVpcPeeringConnectionAccepter(ctx, "peerAccepter", &ec2.VpcPeeringConnectionAccepterArgs{
			VpcPeeringConnectionId: peerConnection.ID(),
			AutoAccept:             pulumi.Bool(true),
		}, pulumi.Provider(awsPeerProvider))
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
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var awsPeerProvider = new Aws.Provider("awsPeerProvider", new()
    {
        Region = "us-west-2",
    });

    var peerIdentity = Aws.GetCallerIdentity.Invoke(, new() {
        Provider = awsPeerProvider,
    });

    var awsProvider = new Aws.Provider("awsProvider", new()
    {
        Region = "us-east-1",
    });

    var main = new Aws.Ec2.Vpc("main", new()
    {
        CidrBlock = "10.0.0.0/16",
    }, new CustomResourceOptions
    {
        Provider = awsProvider,
    });

    var peer = new Aws.Ec2.Vpc("peer", new()
    {
        CidrBlock = "10.1.0.0/16",
    }, new CustomResourceOptions
    {
        Provider = awsPeerProvider,
    });

    var peerConnection = new Aws.Ec2.VpcPeeringConnection("peerConnection", new()
    {
        VpcId = main.Id,
        PeerVpcId = peer.Id,
        PeerOwnerId = peerIdentity.Apply(getCallerIdentityResult => getCallerIdentityResult.AccountId),
        PeerRegion = "us-west-2",
        AutoAccept = false,
    }, new CustomResourceOptions
    {
        Provider = awsProvider,
    });

    var peerAccepter = new Aws.Ec2.VpcPeeringConnectionAccepter("peerAccepter", new()
    {
        VpcPeeringConnectionId = peerConnection.Id,
        AutoAccept = true,
    }, new CustomResourceOptions
    {
        Provider = awsPeerProvider,
    });

});


```

{{% /choosable %}}

{{% choosable language java %}}

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.Provider;
import com.pulumi.aws.ProviderArgs;
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetCallerIdentityArgs;
import com.pulumi.aws.ec2.Vpc;
import com.pulumi.aws.ec2.VpcArgs;
import com.pulumi.aws.ec2.VpcPeeringConnection;
import com.pulumi.aws.ec2.VpcPeeringConnectionArgs;
import com.pulumi.aws.ec2.VpcPeeringConnectionAccepter;
import com.pulumi.aws.ec2.VpcPeeringConnectionAccepterArgs;
import com.pulumi.deployment.InvokeOptions;
import com.pulumi.resources.CustomResourceOptions;
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
        var awsPeerProvider = new Provider("awsPeerProvider", ProviderArgs.builder()
            .region("us-west-2")
            .build());

        final var peerIdentity = AwsFunctions.getCallerIdentity(GetCallerIdentityArgs.builder()
            .build(), InvokeOptions.builder()
                .provider(awsPeerProvider)
                .build());

        var awsProvider = new Provider("awsProvider", ProviderArgs.builder()
            .region("us-east-1")
            .build());

        var main = new Vpc("main", VpcArgs.builder()
            .cidrBlock("10.0.0.0/16")
            .build(), CustomResourceOptions.builder()
                .provider(awsProvider)
                .build());

        var peer = new Vpc("peer", VpcArgs.builder()
            .cidrBlock("10.1.0.0/16")
            .build(), CustomResourceOptions.builder()
                .provider(awsPeerProvider)
                .build());

        var peerConnection = new VpcPeeringConnection("peerConnection", VpcPeeringConnectionArgs.builder()
            .vpcId(main.id())
            .peerVpcId(peer.id())
            .peerOwnerId(peerIdentity.applyValue(_peerIdentity -> _peerIdentity.accountId()))
            .peerRegion("us-west-2")
            .autoAccept(false)
            .build(), CustomResourceOptions.builder()
                .provider(awsProvider)
                .build());

        var peerAccepter = new VpcPeeringConnectionAccepter("peerAccepter", VpcPeeringConnectionAccepterArgs.builder()
            .vpcPeeringConnectionId(peerConnection.id())
            .autoAccept(true)
            .build(), CustomResourceOptions.builder()
                .provider(awsPeerProvider)
                .build());

    }
}

```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
name: vpc-peering
runtime: yaml

variables:
  peerIdentity:
    fn::invoke:
      function: aws:getCallerIdentity
      arguments: {}
      options:
        provider: ${awsPeerProvider}

resources:
  # AWS Provider for us-east-1
  awsProvider:
    type: pulumi:providers:aws
    properties:
      region: "us-east-1"

  # AWS Provider for us-west-2
  awsPeerProvider:
    type: pulumi:providers:aws
    properties:
      region: "us-west-2"

  # Main VPC in us-east-1
  main:
    type: aws:ec2:Vpc
    properties:
      cidrBlock: "10.0.0.0/16"
    options:
      provider: ${awsProvider}

  # Peer VPC in us-west-2
  peer:
    type: aws:ec2:Vpc
    properties:
      cidrBlock: "10.1.0.0/16"
    options:
      provider: ${awsPeerProvider}

  # VPC Peering Connection (requester side)
  peerConnection:
    type: aws:ec2:VpcPeeringConnection
    properties:
      vpcId: ${main.id}
      peerVpcId: ${peer.id}
      peerOwnerId: ${peerIdentity.accountId}
      peerRegion: "us-west-2"
      autoAccept: false
    options:
      provider: ${awsProvider}

  # VPC Peering Connection Accepter (accepter side)
  peerAccepter:
    type: aws:ec2:VpcPeeringConnectionAccepter
    properties:
      vpcPeeringConnectionId: ${peerConnection.id}
      autoAccept: true
    options:
      provider: ${awsPeerProvider}
```

{{% /choosable %}}

{{< /chooser >}}

**After, v7.0.0+**

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}

{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const awsProvider = new aws.Provider("awsProvider", {region: "us-east-1"});
const main = new aws.ec2.Vpc("main", {cidrBlock: "10.0.0.0/16"}, {
    provider: awsProvider,
});
const peer = new aws.ec2.Vpc("peer", {
    region: "us-west-2",
    cidrBlock: "10.1.0.0/16",
}, {
    provider: awsProvider,
});
const peerConnection = new aws.ec2.VpcPeeringConnection("peerConnection", {
    vpcId: main.id,
    peerVpcId: peer.id,
    peerRegion: "us-west-2",
    autoAccept: false,
}, {
    provider: awsProvider,
});
const peerAccepter = new aws.ec2.VpcPeeringConnectionAccepter("peerAccepter", {
    region: "us-west-2",
    vpcPeeringConnectionId: peerConnection.id,
    autoAccept: true,
}, {
    provider: awsProvider,
});

```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi
import pulumi_aws as aws

aws_provider = aws.Provider("awsProvider", region="us-east-1")
main = aws.ec2.Vpc("main", cidr_block="10.0.0.0/16",
opts = pulumi.ResourceOptions(provider=aws_provider))
peer = aws.ec2.Vpc("peer",
    region="us-west-2",
    cidr_block="10.1.0.0/16",
    opts = pulumi.ResourceOptions(provider=aws_provider))
peer_connection = aws.ec2.VpcPeeringConnection("peerConnection",
    vpc_id=main.id,
    peer_vpc_id=peer.id,
    peer_region="us-west-2",
    auto_accept=False,
    opts = pulumi.ResourceOptions(provider=aws_provider))
peer_accepter = aws.ec2.VpcPeeringConnectionAccepter("peerAccepter",
    region="us-west-2",
    vpc_peering_connection_id=peer_connection.id,
    auto_accept=True,
    opts = pulumi.ResourceOptions(provider=aws_provider))

```

{{% /choosable %}}

{{% choosable language go %}}

```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		awsProvider, err := aws.NewProvider(ctx, "awsProvider", &aws.ProviderArgs{
			Region: pulumi.String("us-east-1"),
		})
		if err != nil {
			return err
		}
		main, err := ec2.NewVpc(ctx, "main", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.0.0.0/16"),
		}, pulumi.Provider(awsProvider))
		if err != nil {
			return err
		}
		peer, err := ec2.NewVpc(ctx, "peer", &ec2.VpcArgs{
			Region:    pulumi.String("us-west-2"),
			CidrBlock: pulumi.String("10.1.0.0/16"),
		}, pulumi.Provider(awsProvider))
		if err != nil {
			return err
		}
		peerConnection, err := ec2.NewVpcPeeringConnection(ctx, "peerConnection", &ec2.VpcPeeringConnectionArgs{
			VpcId:      main.ID(),
			PeerVpcId:  peer.ID(),
			PeerRegion: pulumi.String("us-west-2"),
			AutoAccept: pulumi.Bool(false),
		}, pulumi.Provider(awsProvider))
		if err != nil {
			return err
		}
		_, err = ec2.NewVpcPeeringConnectionAccepter(ctx, "peerAccepter", &ec2.VpcPeeringConnectionAccepterArgs{
			Region:                 pulumi.String("us-west-2"),
			VpcPeeringConnectionId: peerConnection.ID(),
			AutoAccept:             pulumi.Bool(true),
		}, pulumi.Provider(awsProvider))
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
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var awsProvider = new Aws.Provider("awsProvider", new()
    {
        Region = "us-east-1",
    });

    var main = new Aws.Ec2.Vpc("main", new()
    {
        CidrBlock = "10.0.0.0/16",
    }, new CustomResourceOptions
    {
        Provider = awsProvider,
    });

    var peer = new Aws.Ec2.Vpc("peer", new()
    {
        Region = "us-west-2",
        CidrBlock = "10.1.0.0/16",
    }, new CustomResourceOptions
    {
        Provider = awsProvider,
    });

    var peerConnection = new Aws.Ec2.VpcPeeringConnection("peerConnection", new()
    {
        VpcId = main.Id,
        PeerVpcId = peer.Id,
        PeerRegion = "us-west-2",
        AutoAccept = false,
    }, new CustomResourceOptions
    {
        Provider = awsProvider,
    });

    var peerAccepter = new Aws.Ec2.VpcPeeringConnectionAccepter("peerAccepter", new()
    {
        Region = "us-west-2",
        VpcPeeringConnectionId = peerConnection.Id,
        AutoAccept = true,
    }, new CustomResourceOptions
    {
        Provider = awsProvider,
    });

});


```

{{% /choosable %}}

{{% choosable language java %}}

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.Provider;
import com.pulumi.aws.ProviderArgs;
import com.pulumi.aws.ec2.Vpc;
import com.pulumi.aws.ec2.VpcArgs;
import com.pulumi.aws.ec2.VpcPeeringConnection;
import com.pulumi.aws.ec2.VpcPeeringConnectionArgs;
import com.pulumi.aws.ec2.VpcPeeringConnectionAccepter;
import com.pulumi.aws.ec2.VpcPeeringConnectionAccepterArgs;
import com.pulumi.resources.CustomResourceOptions;
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
        var awsProvider = new Provider("awsProvider", ProviderArgs.builder()
            .region("us-east-1")
            .build());

        var main = new Vpc("main", VpcArgs.builder()
            .cidrBlock("10.0.0.0/16")
            .build(), CustomResourceOptions.builder()
                .provider(awsProvider)
                .build());

        var peer = new Vpc("peer", VpcArgs.builder()
            .region("us-west-2")
            .cidrBlock("10.1.0.0/16")
            .build(), CustomResourceOptions.builder()
                .provider(awsProvider)
                .build());

        var peerConnection = new VpcPeeringConnection("peerConnection", VpcPeeringConnectionArgs.builder()
            .vpcId(main.id())
            .peerVpcId(peer.id())
            .peerRegion("us-west-2")
            .autoAccept(false)
            .build(), CustomResourceOptions.builder()
                .provider(awsProvider)
                .build());

        var peerAccepter = new VpcPeeringConnectionAccepter("peerAccepter", VpcPeeringConnectionAccepterArgs.builder()
            .region("us-west-2")
            .vpcPeeringConnectionId(peerConnection.id())
            .autoAccept(true)
            .build(), CustomResourceOptions.builder()
                .provider(awsProvider)
                .build());

    }
}

```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
name: vpc-peering
runtime: yaml

resources:
  # AWS Provider for us-east-1
  awsProvider:
    type: pulumi:providers:aws
    properties:
      region: "us-east-1"

  # Main VPC in us-east-1
  main:
    type: aws:ec2:Vpc
    properties:
      cidrBlock: "10.0.0.0/16"
    options:
      provider: ${awsProvider}

  # Peer VPC in us-west-2
  peer:
    type: aws:ec2:Vpc
    properties:
      region: "us-west-2"
      cidrBlock: "10.1.0.0/16"
    options:
      provider: ${awsProvider}

  # VPC Peering Connection (requester side)
  peerConnection:
    type: aws:ec2:VpcPeeringConnection
    properties:
      vpcId: ${main.id}
      peerVpcId: ${peer.id}
      peerRegion: "us-west-2"
      autoAccept: false
    options:
      provider: ${awsProvider}

  # VPC Peering Connection Accepter (accepter side)
  peerAccepter:
    type: aws:ec2:VpcPeeringConnectionAccepter
    properties:
      region: "us-west-2"
      vpcPeeringConnectionId: ${peerConnection.id}
      autoAccept: true
    options:
      provider: ${awsProvider}
```

{{% /choosable %}}

{{< /chooser >}}

### KMS replica key

**Before, Pre-v7.0.0**

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}

{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const awsPrimaryProvider = new aws.Provider("awsPrimaryProvider", {region: "us-east-1"});
const awsProvider = new aws.Provider("awsProvider", {region: "us-west-2"});
const primary = new aws.kms.Key("primary", {
    description: "Multi-Region primary key",
    deletionWindowInDays: 30,
    multiRegion: true,
}, {
    provider: awsPrimaryProvider,
});
const replica = new aws.kms.ReplicaKey("replica", {
    description: "Multi-Region replica key",
    deletionWindowInDays: 7,
    primaryKeyArn: primary.arn,
}, {
    provider: awsProvider,
});

```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi
import pulumi_aws as aws

aws_primary_provider = aws.Provider("awsPrimaryProvider", region="us-east-1")
aws_provider = aws.Provider("awsProvider", region="us-west-2")
primary = aws.kms.Key("primary",
    description="Multi-Region primary key",
    deletion_window_in_days=30,
    multi_region=True,
    opts = pulumi.ResourceOptions(provider=aws_primary_provider))
replica = aws.kms.ReplicaKey("replica",
    description="Multi-Region replica key",
    deletion_window_in_days=7,
    primary_key_arn=primary.arn,
    opts = pulumi.ResourceOptions(provider=aws_provider))

```

{{% /choosable %}}

{{% choosable language go %}}

```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws/kms"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		awsPrimaryProvider, err := aws.NewProvider(ctx, "awsPrimaryProvider", &aws.ProviderArgs{
			Region: pulumi.String("us-east-1"),
		})
		if err != nil {
			return err
		}
		awsProvider, err := aws.NewProvider(ctx, "awsProvider", &aws.ProviderArgs{
			Region: pulumi.String("us-west-2"),
		})
		if err != nil {
			return err
		}
		primary, err := kms.NewKey(ctx, "primary", &kms.KeyArgs{
			Description:          pulumi.String("Multi-Region primary key"),
			DeletionWindowInDays: pulumi.Int(30),
			MultiRegion:          pulumi.Bool(true),
		}, pulumi.Provider(awsPrimaryProvider))
		if err != nil {
			return err
		}
		_, err = kms.NewReplicaKey(ctx, "replica", &kms.ReplicaKeyArgs{
			Description:          pulumi.String("Multi-Region replica key"),
			DeletionWindowInDays: pulumi.Int(7),
			PrimaryKeyArn:        primary.Arn,
		}, pulumi.Provider(awsProvider))
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
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var awsPrimaryProvider = new Aws.Provider("awsPrimaryProvider", new()
    {
        Region = "us-east-1",
    });

    var awsProvider = new Aws.Provider("awsProvider", new()
    {
        Region = "us-west-2",
    });

    var primary = new Aws.Kms.Key("primary", new()
    {
        Description = "Multi-Region primary key",
        DeletionWindowInDays = 30,
        MultiRegion = true,
    }, new CustomResourceOptions
    {
        Provider = awsPrimaryProvider,
    });

    var replica = new Aws.Kms.ReplicaKey("replica", new()
    {
        Description = "Multi-Region replica key",
        DeletionWindowInDays = 7,
        PrimaryKeyArn = primary.Arn,
    }, new CustomResourceOptions
    {
        Provider = awsProvider,
    });

});


```

{{% /choosable %}}

{{% choosable language java %}}

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.Provider;
import com.pulumi.aws.ProviderArgs;
import com.pulumi.aws.kms.Key;
import com.pulumi.aws.kms.KeyArgs;
import com.pulumi.aws.kms.ReplicaKey;
import com.pulumi.aws.kms.ReplicaKeyArgs;
import com.pulumi.resources.CustomResourceOptions;
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
        var awsPrimaryProvider = new Provider("awsPrimaryProvider", ProviderArgs.builder()
            .region("us-east-1")
            .build());

        var awsProvider = new Provider("awsProvider", ProviderArgs.builder()
            .region("us-west-2")
            .build());

        var primary = new Key("primary", KeyArgs.builder()
            .description("Multi-Region primary key")
            .deletionWindowInDays(30)
            .multiRegion(true)
            .build(), CustomResourceOptions.builder()
                .provider(awsPrimaryProvider)
                .build());

        var replica = new ReplicaKey("replica", ReplicaKeyArgs.builder()
            .description("Multi-Region replica key")
            .deletionWindowInDays(7)
            .primaryKeyArn(primary.arn())
            .build(), CustomResourceOptions.builder()
                .provider(awsProvider)
                .build());

    }
}

```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
name: kms-multi-region
runtime: yaml

resources:
  # AWS Provider for us-east-1 (primary)
  awsPrimaryProvider:
    type: pulumi:providers:aws
    properties:
      region: "us-east-1"

  # AWS Provider for us-west-2 (default)
  awsProvider:
    type: pulumi:providers:aws
    properties:
      region: "us-west-2"

  # Multi-Region primary KMS key in us-east-1
  primary:
    type: aws:kms:Key
    properties:
      description: "Multi-Region primary key"
      deletionWindowInDays: 30
      multiRegion: true
    options:
      provider: ${awsPrimaryProvider}

  # Multi-Region replica KMS key in us-west-2
  replica:
    type: aws:kms:ReplicaKey
    properties:
      description: "Multi-Region replica key"
      deletionWindowInDays: 7
      primaryKeyArn: ${primary.arn}
    options:
      provider: ${awsProvider}
```

{{% /choosable %}}

{{< /chooser >}}

**After, v7.0.0**

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}

{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const awsProvider = new aws.Provider("awsProvider", {region: "us-west-2"});
const primary = new aws.kms.Key("primary", {
    region: "us-east-1",
    description: "Multi-Region primary key",
    deletionWindowInDays: 30,
    multiRegion: true,
}, {
    provider: awsProvider,
});
const replica = new aws.kms.ReplicaKey("replica", {
    description: "Multi-Region replica key",
    deletionWindowInDays: 7,
    primaryKeyArn: primary.arn,
}, {
    provider: awsProvider,
});

```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi
import pulumi_aws as aws

aws_provider = aws.Provider("awsProvider", region="us-west-2")
primary = aws.kms.Key("primary",
    region="us-east-1",
    description="Multi-Region primary key",
    deletion_window_in_days=30,
    multi_region=True,
    opts = pulumi.ResourceOptions(provider=aws_provider))
replica = aws.kms.ReplicaKey("replica",
    description="Multi-Region replica key",
    deletion_window_in_days=7,
    primary_key_arn=primary.arn,
    opts = pulumi.ResourceOptions(provider=aws_provider))

```

{{% /choosable %}}

{{% choosable language go %}}

```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws/kms"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		awsProvider, err := aws.NewProvider(ctx, "awsProvider", &aws.ProviderArgs{
			Region: pulumi.String("us-west-2"),
		})
		if err != nil {
			return err
		}
		primary, err := kms.NewKey(ctx, "primary", &kms.KeyArgs{
			Region:               pulumi.String("us-east-1"),
			Description:          pulumi.String("Multi-Region primary key"),
			DeletionWindowInDays: pulumi.Int(30),
			MultiRegion:          pulumi.Bool(true),
		}, pulumi.Provider(awsProvider))
		if err != nil {
			return err
		}
		_, err = kms.NewReplicaKey(ctx, "replica", &kms.ReplicaKeyArgs{
			Description:          pulumi.String("Multi-Region replica key"),
			DeletionWindowInDays: pulumi.Int(7),
			PrimaryKeyArn:        primary.Arn,
		}, pulumi.Provider(awsProvider))
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
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var awsProvider = new Aws.Provider("awsProvider", new()
    {
        Region = "us-west-2",
    });

    var primary = new Aws.Kms.Key("primary", new()
    {
        Region = "us-east-1",
        Description = "Multi-Region primary key",
        DeletionWindowInDays = 30,
        MultiRegion = true,
    }, new CustomResourceOptions
    {
        Provider = awsProvider,
    });

    var replica = new Aws.Kms.ReplicaKey("replica", new()
    {
        Description = "Multi-Region replica key",
        DeletionWindowInDays = 7,
        PrimaryKeyArn = primary.Arn,
    }, new CustomResourceOptions
    {
        Provider = awsProvider,
    });

});


```

{{% /choosable %}}

{{% choosable language java %}}

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.Provider;
import com.pulumi.aws.ProviderArgs;
import com.pulumi.aws.kms.Key;
import com.pulumi.aws.kms.KeyArgs;
import com.pulumi.aws.kms.ReplicaKey;
import com.pulumi.aws.kms.ReplicaKeyArgs;
import com.pulumi.resources.CustomResourceOptions;
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
        var awsProvider = new Provider("awsProvider", ProviderArgs.builder()
            .region("us-west-2")
            .build());

        var primary = new Key("primary", KeyArgs.builder()
            .region("us-east-1")
            .description("Multi-Region primary key")
            .deletionWindowInDays(30)
            .multiRegion(true)
            .build(), CustomResourceOptions.builder()
                .provider(awsProvider)
                .build());

        var replica = new ReplicaKey("replica", ReplicaKeyArgs.builder()
            .description("Multi-Region replica key")
            .deletionWindowInDays(7)
            .primaryKeyArn(primary.arn())
            .build(), CustomResourceOptions.builder()
                .provider(awsProvider)
                .build());

    }
}

```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
name: kms-multi-region
runtime: yaml

resources:
  # AWS Provider for us-west-2
  awsProvider:
    type: pulumi:providers:aws
    properties:
      region: "us-west-2"

  # Multi-Region primary KMS key in us-east-1
  primary:
    type: aws:kms:Key
    properties:
      region: "us-east-1"
      description: "Multi-Region primary key"
      deletionWindowInDays: 30
      multiRegion: true
    options:
      provider: ${awsProvider}

  # Multi-Region replica KMS key in us-west-2
  replica:
    type: aws:kms:ReplicaKey
    properties:
      description: "Multi-Region replica key"
      deletionWindowInDays: 7
      primaryKeyArn: ${primary.arn}
    options:
      provider: ${awsProvider}
```

{{% /choosable %}}

{{< /chooser >}}


### S3 bucket replication configuration

**Before, Pre-v7.0.0**

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}

{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const assumeRoleDocument = aws.iam.getPolicyDocumentOutput({
    statements: [{
        effect: "Allow",
        principals: [{
            type: "Service",
            identifiers: ["s3.amazonaws.com"],
        }],
        actions: ["sts:AssumeRole"],
    }],
});
const awsProvider = new aws.Provider("awsProvider", {region: "eu-west-1"});
const destination = new aws.s3.Bucket("destination", {}, {
    provider: awsProvider,
});
const awsCentralProvider = new aws.Provider("awsCentralProvider", {region: "eu-central-1"});
const source = new aws.s3.Bucket("source", {}, {
    provider: awsCentralProvider,
});
const replicationPolicyDocument = aws.iam.getPolicyDocumentOutput({
    statements: [
        {
            effect: "Allow",
            actions: [
                "s3:GetReplicationConfiguration",
                "s3:ListBucket",
            ],
            resources: [source.arn],
        },
        {
            effect: "Allow",
            actions: [
                "s3:GetObjectVersionForReplication",
                "s3:GetObjectVersionAcl",
                "s3:GetObjectVersionTagging",
            ],
            resources: [pulumi.interpolate`${source.arn}/*`],
        },
        {
            effect: "Allow",
            actions: [
                "s3:ReplicateObject",
                "s3:ReplicateDelete",
                "s3:ReplicateTags",
            ],
            resources: [pulumi.interpolate`${destination.arn}/*`],
        },
    ],
});
const replicationRole = new aws.iam.Role("replicationRole", {assumeRolePolicy: assumeRoleDocument.apply(assumeRoleDocument => assumeRoleDocument.json)}, {
    provider: awsProvider,
});
const replicationPolicy = new aws.iam.Policy("replicationPolicy", {policy: replicationPolicyDocument.apply(replicationPolicyDocument => replicationPolicyDocument.json)}, {
    provider: awsProvider,
});
const replicationPolicyAttachment = new aws.iam.RolePolicyAttachment("replicationPolicyAttachment", {
    role: replicationRole.name,
    policyArn: replicationPolicy.arn,
}, {
    provider: awsProvider,
});
const destinationVersioning = new aws.s3.BucketVersioning("destinationVersioning", {
    bucket: destination.id,
    versioningConfiguration: {
        status: "Enabled",
    },
}, {
    provider: awsProvider,
});
const sourceBucketAcl = new aws.s3.BucketAcl("sourceBucketAcl", {
    bucket: source.id,
    acl: "private",
}, {
    provider: awsCentralProvider,
});
const sourceVersioning = new aws.s3.BucketVersioning("sourceVersioning", {
    bucket: source.id,
    versioningConfiguration: {
        status: "Enabled",
    },
}, {
    provider: awsCentralProvider,
});
const replication = new aws.s3.BucketReplicationConfig("replication", {
    role: replicationRole.arn,
    bucket: source.id,
    rules: [{
        id: "examplerule",
        filter: {
            prefix: "example",
        },
        status: "Enabled",
        destination: {
            bucket: destination.arn,
            storageClass: "STANDARD",
        },
    }],
}, {
    provider: awsCentralProvider,
    dependsOn: [sourceVersioning],
});

```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi
import pulumi_aws as aws

assume_role_document = aws.iam.get_policy_document_output(statements=[{
    "effect": "Allow",
    "principals": [{
        "type": "Service",
        "identifiers": ["s3.amazonaws.com"],
    }],
    "actions": ["sts:AssumeRole"],
}])
aws_provider = aws.Provider("awsProvider", region="eu-west-1")
destination = aws.s3.Bucket("destination", opts = pulumi.ResourceOptions(provider=aws_provider))
aws_central_provider = aws.Provider("awsCentralProvider", region="eu-central-1")
source = aws.s3.Bucket("source", opts = pulumi.ResourceOptions(provider=aws_central_provider))
replication_policy_document = aws.iam.get_policy_document_output(statements=[
    {
        "effect": "Allow",
        "actions": [
            "s3:GetReplicationConfiguration",
            "s3:ListBucket",
        ],
        "resources": [source.arn],
    },
    {
        "effect": "Allow",
        "actions": [
            "s3:GetObjectVersionForReplication",
            "s3:GetObjectVersionAcl",
            "s3:GetObjectVersionTagging",
        ],
        "resources": [source.arn.apply(lambda arn: f"{arn}/*")],
    },
    {
        "effect": "Allow",
        "actions": [
            "s3:ReplicateObject",
            "s3:ReplicateDelete",
            "s3:ReplicateTags",
        ],
        "resources": [destination.arn.apply(lambda arn: f"{arn}/*")],
    },
])
replication_role = aws.iam.Role("replicationRole", assume_role_policy=assume_role_document.json,
opts = pulumi.ResourceOptions(provider=aws_provider))
replication_policy = aws.iam.Policy("replicationPolicy", policy=replication_policy_document.json,
opts = pulumi.ResourceOptions(provider=aws_provider))
replication_policy_attachment = aws.iam.RolePolicyAttachment("replicationPolicyAttachment",
    role=replication_role.name,
    policy_arn=replication_policy.arn,
    opts = pulumi.ResourceOptions(provider=aws_provider))
destination_versioning = aws.s3.BucketVersioning("destinationVersioning",
    bucket=destination.id,
    versioning_configuration={
        "status": "Enabled",
    },
    opts = pulumi.ResourceOptions(provider=aws_provider))
source_bucket_acl = aws.s3.BucketAcl("sourceBucketAcl",
    bucket=source.id,
    acl="private",
    opts = pulumi.ResourceOptions(provider=aws_central_provider))
source_versioning = aws.s3.BucketVersioning("sourceVersioning",
    bucket=source.id,
    versioning_configuration={
        "status": "Enabled",
    },
    opts = pulumi.ResourceOptions(provider=aws_central_provider))
replication = aws.s3.BucketReplicationConfig("replication",
    role=replication_role.arn,
    bucket=source.id,
    rules=[{
        "id": "examplerule",
        "filter": {
            "prefix": "example",
        },
        "status": "Enabled",
        "destination": {
            "bucket": destination.arn,
            "storage_class": "STANDARD",
        },
    }],
    opts = pulumi.ResourceOptions(provider=aws_central_provider,
        depends_on=[source_versioning]))

```

{{% /choosable %}}

{{% choosable language go %}}

```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		assumeRoleDocument := iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
			Statements: iam.GetPolicyDocumentStatementArray{
				&iam.GetPolicyDocumentStatementArgs{
					Effect: pulumi.String("Allow"),
					Principals: iam.GetPolicyDocumentStatementPrincipalArray{
						&iam.GetPolicyDocumentStatementPrincipalArgs{
							Type: pulumi.String("Service"),
							Identifiers: pulumi.StringArray{
								pulumi.String("s3.amazonaws.com"),
							},
						},
					},
					Actions: pulumi.StringArray{
						pulumi.String("sts:AssumeRole"),
					},
				},
			},
		}, nil)
		awsProvider, err := aws.NewProvider(ctx, "awsProvider", &aws.ProviderArgs{
			Region: pulumi.String("eu-west-1"),
		})
		if err != nil {
			return err
		}
		destination, err := s3.NewBucket(ctx, "destination", nil, pulumi.Provider(awsProvider))
		if err != nil {
			return err
		}
		awsCentralProvider, err := aws.NewProvider(ctx, "awsCentralProvider", &aws.ProviderArgs{
			Region: pulumi.String("eu-central-1"),
		})
		if err != nil {
			return err
		}
		source, err := s3.NewBucket(ctx, "source", nil, pulumi.Provider(awsCentralProvider))
		if err != nil {
			return err
		}
		replicationPolicyDocument := iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
			Statements: iam.GetPolicyDocumentStatementArray{
				&iam.GetPolicyDocumentStatementArgs{
					Effect: pulumi.String("Allow"),
					Actions: pulumi.StringArray{
						pulumi.String("s3:GetReplicationConfiguration"),
						pulumi.String("s3:ListBucket"),
					},
					Resources: pulumi.StringArray{
						source.Arn,
					},
				},
				&iam.GetPolicyDocumentStatementArgs{
					Effect: pulumi.String("Allow"),
					Actions: pulumi.StringArray{
						pulumi.String("s3:GetObjectVersionForReplication"),
						pulumi.String("s3:GetObjectVersionAcl"),
						pulumi.String("s3:GetObjectVersionTagging"),
					},
					Resources: pulumi.StringArray{
						source.Arn.ApplyT(func(arn string) (string, error) {
							return fmt.Sprintf("%v/*", arn), nil
						}).(pulumi.StringOutput),
					},
				},
				&iam.GetPolicyDocumentStatementArgs{
					Effect: pulumi.String("Allow"),
					Actions: pulumi.StringArray{
						pulumi.String("s3:ReplicateObject"),
						pulumi.String("s3:ReplicateDelete"),
						pulumi.String("s3:ReplicateTags"),
					},
					Resources: pulumi.StringArray{
						destination.Arn.ApplyT(func(arn string) (string, error) {
							return fmt.Sprintf("%v/*", arn), nil
						}).(pulumi.StringOutput),
					},
				},
			},
		}, nil)
		replicationRole, err := iam.NewRole(ctx, "replicationRole", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.String(assumeRoleDocument.ApplyT(func(assumeRoleDocument iam.GetPolicyDocumentResult) (*string, error) {
				return &assumeRoleDocument.Json, nil
			}).(pulumi.StringPtrOutput)),
		}, pulumi.Provider(awsProvider))
		if err != nil {
			return err
		}
		replicationPolicy, err := iam.NewPolicy(ctx, "replicationPolicy", &iam.PolicyArgs{
			Policy: pulumi.String(replicationPolicyDocument.ApplyT(func(replicationPolicyDocument iam.GetPolicyDocumentResult) (*string, error) {
				return &replicationPolicyDocument.Json, nil
			}).(pulumi.StringPtrOutput)),
		}, pulumi.Provider(awsProvider))
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicyAttachment(ctx, "replicationPolicyAttachment", &iam.RolePolicyAttachmentArgs{
			Role:      replicationRole.Name,
			PolicyArn: replicationPolicy.Arn,
		}, pulumi.Provider(awsProvider))
		if err != nil {
			return err
		}
		_, err = s3.NewBucketVersioning(ctx, "destinationVersioning", &s3.BucketVersioningArgs{
			Bucket: destination.ID(),
			VersioningConfiguration: &s3.BucketVersioningVersioningConfigurationArgs{
				Status: pulumi.String("Enabled"),
			},
		}, pulumi.Provider(awsProvider))
		if err != nil {
			return err
		}
		_, err = s3.NewBucketAcl(ctx, "sourceBucketAcl", &s3.BucketAclArgs{
			Bucket: source.ID(),
			Acl:    pulumi.String("private"),
		}, pulumi.Provider(awsCentralProvider))
		if err != nil {
			return err
		}
		sourceVersioning, err := s3.NewBucketVersioning(ctx, "sourceVersioning", &s3.BucketVersioningArgs{
			Bucket: source.ID(),
			VersioningConfiguration: &s3.BucketVersioningVersioningConfigurationArgs{
				Status: pulumi.String("Enabled"),
			},
		}, pulumi.Provider(awsCentralProvider))
		if err != nil {
			return err
		}
		_, err = s3.NewBucketReplicationConfig(ctx, "replication", &s3.BucketReplicationConfigArgs{
			Role:   replicationRole.Arn,
			Bucket: source.ID(),
			Rules: s3.BucketReplicationConfigRuleArray{
				&s3.BucketReplicationConfigRuleArgs{
					Id: pulumi.String("examplerule"),
					Filter: &s3.BucketReplicationConfigRuleFilterArgs{
						Prefix: pulumi.String("example"),
					},
					Status: pulumi.String("Enabled"),
					Destination: &s3.BucketReplicationConfigRuleDestinationArgs{
						Bucket:       destination.Arn,
						StorageClass: pulumi.String("STANDARD"),
					},
				},
			},
		}, pulumi.Provider(awsCentralProvider), pulumi.DependsOn([]pulumi.Resource{
			sourceVersioning,
		}))
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
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var assumeRoleDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Effect = "Allow",
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Type = "Service",
                        Identifiers = new[]
                        {
                            "s3.amazonaws.com",
                        },
                    },
                },
                Actions = new[]
                {
                    "sts:AssumeRole",
                },
            },
        },
    });

    var awsProvider = new Aws.Provider("awsProvider", new()
    {
        Region = "eu-west-1",
    });

    var destination = new Aws.S3.Bucket("destination", new()
    {
    }, new CustomResourceOptions
    {
        Provider = awsProvider,
    });

    var awsCentralProvider = new Aws.Provider("awsCentralProvider", new()
    {
        Region = "eu-central-1",
    });

    var source = new Aws.S3.Bucket("source", new()
    {
    }, new CustomResourceOptions
    {
        Provider = awsCentralProvider,
    });

    var replicationPolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Effect = "Allow",
                Actions = new[]
                {
                    "s3:GetReplicationConfiguration",
                    "s3:ListBucket",
                },
                Resources = new[]
                {
                    source.Arn,
                },
            },
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Effect = "Allow",
                Actions = new[]
                {
                    "s3:GetObjectVersionForReplication",
                    "s3:GetObjectVersionAcl",
                    "s3:GetObjectVersionTagging",
                },
                Resources = new[]
                {
                    $"{source.Arn}/*",
                },
            },
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Effect = "Allow",
                Actions = new[]
                {
                    "s3:ReplicateObject",
                    "s3:ReplicateDelete",
                    "s3:ReplicateTags",
                },
                Resources = new[]
                {
                    $"{destination.Arn}/*",
                },
            },
        },
    });

    var replicationRole = new Aws.Iam.Role("replicationRole", new()
    {
        AssumeRolePolicy = assumeRoleDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    }, new CustomResourceOptions
    {
        Provider = awsProvider,
    });

    var replicationPolicy = new Aws.Iam.Policy("replicationPolicy", new()
    {
        PolicyDocument = replicationPolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    }, new CustomResourceOptions
    {
        Provider = awsProvider,
    });

    var replicationPolicyAttachment = new Aws.Iam.RolePolicyAttachment("replicationPolicyAttachment", new()
    {
        Role = replicationRole.Name,
        PolicyArn = replicationPolicy.Arn,
    }, new CustomResourceOptions
    {
        Provider = awsProvider,
    });

    var destinationVersioning = new Aws.S3.BucketVersioning("destinationVersioning", new()
    {
        Bucket = destination.Id,
        VersioningConfiguration = new Aws.S3.Inputs.BucketVersioningVersioningConfigurationArgs
        {
            Status = "Enabled",
        },
    }, new CustomResourceOptions
    {
        Provider = awsProvider,
    });

    var sourceBucketAcl = new Aws.S3.BucketAcl("sourceBucketAcl", new()
    {
        Bucket = source.Id,
        Acl = "private",
    }, new CustomResourceOptions
    {
        Provider = awsCentralProvider,
    });

    var sourceVersioning = new Aws.S3.BucketVersioning("sourceVersioning", new()
    {
        Bucket = source.Id,
        VersioningConfiguration = new Aws.S3.Inputs.BucketVersioningVersioningConfigurationArgs
        {
            Status = "Enabled",
        },
    }, new CustomResourceOptions
    {
        Provider = awsCentralProvider,
    });

    var replication = new Aws.S3.BucketReplicationConfig("replication", new()
    {
        Role = replicationRole.Arn,
        Bucket = source.Id,
        Rules = new[]
        {
            new Aws.S3.Inputs.BucketReplicationConfigRuleArgs
            {
                Id = "examplerule",
                Filter = new Aws.S3.Inputs.BucketReplicationConfigRuleFilterArgs
                {
                    Prefix = "example",
                },
                Status = "Enabled",
                Destination = new Aws.S3.Inputs.BucketReplicationConfigRuleDestinationArgs
                {
                    Bucket = destination.Arn,
                    StorageClass = "STANDARD",
                },
            },
        },
    }, new CustomResourceOptions
    {
        Provider = awsCentralProvider,
        DependsOn =
        {
            sourceVersioning,
        },
    });

});


```

{{% /choosable %}}

{{% choosable language java %}}

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.Provider;
import com.pulumi.aws.ProviderArgs;
import com.pulumi.aws.s3.Bucket;
import com.pulumi.aws.s3.BucketArgs;
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.iam.Policy;
import com.pulumi.aws.iam.PolicyArgs;
import com.pulumi.aws.iam.RolePolicyAttachment;
import com.pulumi.aws.iam.RolePolicyAttachmentArgs;
import com.pulumi.aws.s3.BucketVersioning;
import com.pulumi.aws.s3.BucketVersioningArgs;
import com.pulumi.aws.s3.inputs.BucketVersioningVersioningConfigurationArgs;
import com.pulumi.aws.s3.BucketAcl;
import com.pulumi.aws.s3.BucketAclArgs;
import com.pulumi.aws.s3.BucketReplicationConfig;
import com.pulumi.aws.s3.BucketReplicationConfigArgs;
import com.pulumi.aws.s3.inputs.BucketReplicationConfigRuleArgs;
import com.pulumi.aws.s3.inputs.BucketReplicationConfigRuleFilterArgs;
import com.pulumi.aws.s3.inputs.BucketReplicationConfigRuleDestinationArgs;
import com.pulumi.resources.CustomResourceOptions;
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
        final var assumeRoleDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("Service")
                    .identifiers("s3.amazonaws.com")
                    .build())
                .actions("sts:AssumeRole")
                .build())
            .build());

        var awsProvider = new Provider("awsProvider", ProviderArgs.builder()
            .region("eu-west-1")
            .build());

        var destination = new Bucket("destination", BucketArgs.Empty, CustomResourceOptions.builder()
            .provider(awsProvider)
            .build());

        var awsCentralProvider = new Provider("awsCentralProvider", ProviderArgs.builder()
            .region("eu-central-1")
            .build());

        var source = new Bucket("source", BucketArgs.Empty, CustomResourceOptions.builder()
            .provider(awsCentralProvider)
            .build());

        final var replicationPolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(            
                GetPolicyDocumentStatementArgs.builder()
                    .effect("Allow")
                    .actions(                    
                        "s3:GetReplicationConfiguration",
                        "s3:ListBucket")
                    .resources(source.arn())
                    .build(),
                GetPolicyDocumentStatementArgs.builder()
                    .effect("Allow")
                    .actions(                    
                        "s3:GetObjectVersionForReplication",
                        "s3:GetObjectVersionAcl",
                        "s3:GetObjectVersionTagging")
                    .resources(source.arn().applyValue(_arn -> String.format("%s/*", _arn)))
                    .build(),
                GetPolicyDocumentStatementArgs.builder()
                    .effect("Allow")
                    .actions(                    
                        "s3:ReplicateObject",
                        "s3:ReplicateDelete",
                        "s3:ReplicateTags")
                    .resources(destination.arn().applyValue(_arn -> String.format("%s/*", _arn)))
                    .build())
            .build());

        var replicationRole = new Role("replicationRole", RoleArgs.builder()
            .assumeRolePolicy(assumeRoleDocument.applyValue(_assumeRoleDocument -> _assumeRoleDocument.json()))
            .build(), CustomResourceOptions.builder()
                .provider(awsProvider)
                .build());

        var replicationPolicy = new Policy("replicationPolicy", PolicyArgs.builder()
            .policy(replicationPolicyDocument.applyValue(_replicationPolicyDocument -> _replicationPolicyDocument.json()))
            .build(), CustomResourceOptions.builder()
                .provider(awsProvider)
                .build());

        var replicationPolicyAttachment = new RolePolicyAttachment("replicationPolicyAttachment", RolePolicyAttachmentArgs.builder()
            .role(replicationRole.name())
            .policyArn(replicationPolicy.arn())
            .build(), CustomResourceOptions.builder()
                .provider(awsProvider)
                .build());

        var destinationVersioning = new BucketVersioning("destinationVersioning", BucketVersioningArgs.builder()
            .bucket(destination.id())
            .versioningConfiguration(BucketVersioningVersioningConfigurationArgs.builder()
                .status("Enabled")
                .build())
            .build(), CustomResourceOptions.builder()
                .provider(awsProvider)
                .build());

        var sourceBucketAcl = new BucketAcl("sourceBucketAcl", BucketAclArgs.builder()
            .bucket(source.id())
            .acl("private")
            .build(), CustomResourceOptions.builder()
                .provider(awsCentralProvider)
                .build());

        var sourceVersioning = new BucketVersioning("sourceVersioning", BucketVersioningArgs.builder()
            .bucket(source.id())
            .versioningConfiguration(BucketVersioningVersioningConfigurationArgs.builder()
                .status("Enabled")
                .build())
            .build(), CustomResourceOptions.builder()
                .provider(awsCentralProvider)
                .build());

        var replication = new BucketReplicationConfig("replication", BucketReplicationConfigArgs.builder()
            .role(replicationRole.arn())
            .bucket(source.id())
            .rules(BucketReplicationConfigRuleArgs.builder()
                .id("examplerule")
                .filter(BucketReplicationConfigRuleFilterArgs.builder()
                    .prefix("example")
                    .build())
                .status("Enabled")
                .destination(BucketReplicationConfigRuleDestinationArgs.builder()
                    .bucket(destination.arn())
                    .storageClass("STANDARD")
                    .build())
                .build())
            .build(), CustomResourceOptions.builder()
                .provider(awsCentralProvider)
                .dependsOn(sourceVersioning)
                .build());

    }
}

```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
name: s3-replication
runtime: yaml

variables:
  assumeRoleDocument:
    fn::invoke:
      function: aws:iam:getPolicyDocument
      arguments:
        statements:
          - effect: "Allow"
            principals:
              - type: "Service"
                identifiers:
                  - s3.amazonaws.com
            actions:
              - sts:AssumeRole

  replicationPolicyDocument:
    fn::invoke:
      function: aws:iam:getPolicyDocument
      arguments:
        statements:
          - effect: "Allow"
            actions:
              - s3:GetReplicationConfiguration
              - s3:ListBucket
            resources:
              - ${source.arn}
          - effect: "Allow"
            actions:
              - s3:GetObjectVersionForReplication
              - s3:GetObjectVersionAcl
              - s3:GetObjectVersionTagging
            resources:
              - "${source.arn}/*"
          - effect: "Allow"
            actions:
              - s3:ReplicateObject
              - s3:ReplicateDelete
              - s3:ReplicateTags
            resources:
              - "${destination.arn}/*"

resources:
  # AWS Provider for eu-west-1 (default)
  awsProvider:
    type: pulumi:providers:aws
    properties:
      region: "eu-west-1"

  # AWS Provider for eu-central-1
  awsCentralProvider:
    type: pulumi:providers:aws
    properties:
      region: "eu-central-1"

  # IAM Role for replication
  replicationRole:
    type: aws:iam:Role
    properties:
      assumeRolePolicy: ${assumeRoleDocument.json}
    options:
      provider: ${awsProvider}

  # Destination bucket in eu-west-1
  destination:
    type: aws:s3:Bucket
    options:
      provider: ${awsProvider}

  # Source bucket in eu-central-1
  source:
    type: aws:s3:Bucket
    options:
      provider: ${awsCentralProvider}

  # IAM Policy for replication
  replicationPolicy:
    type: aws:iam:Policy
    properties:
      policy: ${replicationPolicyDocument.json}
    options:
      provider: ${awsProvider}

  # Attach policy to role
  replicationPolicyAttachment:
    type: aws:iam:RolePolicyAttachment
    properties:
      role: ${replicationRole.name}
      policyArn: ${replicationPolicy.arn}
    options:
      provider: ${awsProvider}

  # Enable versioning on destination bucket
  destinationVersioning:
    type: aws:s3:BucketVersioning
    properties:
      bucket: ${destination.id}
      versioningConfiguration:
        status: "Enabled"
    options:
      provider: ${awsProvider}

  # Set ACL on source bucket
  sourceBucketAcl:
    type: aws:s3:BucketAcl
    properties:
      bucket: ${source.id}
      acl: "private"
    options:
      provider: ${awsCentralProvider}

  # Enable versioning on source bucket
  sourceVersioning:
    type: aws:s3:BucketVersioning
    properties:
      bucket: ${source.id}
      versioningConfiguration:
        status: "Enabled"
    options:
      provider: ${awsCentralProvider}

  # Configure bucket replication
  replication:
    type: aws:s3:BucketReplicationConfig
    properties:
      role: ${replicationRole.arn}
      bucket: ${source.id}
      rules:
        - id: "examplerule"
          filter:
            prefix: "example"
          status: "Enabled"
          destination:
            bucket: ${destination.arn}
            storageClass: "STANDARD"
    options:
      provider: ${awsCentralProvider}
      dependsOn:
        - ${sourceVersioning}
```

{{% /choosable %}}

{{< /chooser >}}


**After, v7.0.0**

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}

{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const assumeRoleDocument = aws.iam.getPolicyDocumentOutput({
    statements: [{
        effect: "Allow",
        principals: [{
            type: "Service",
            identifiers: ["s3.amazonaws.com"],
        }],
        actions: ["sts:AssumeRole"],
    }],
});
const awsProvider = new aws.Provider("awsProvider", {region: "eu-west-1"});
const destination = new aws.s3.Bucket("destination", {}, {
    provider: awsProvider,
});
const source = new aws.s3.Bucket("source", {region: "eu-central-1"}, {
    provider: awsProvider,
});
const replicationPolicyDocument = aws.iam.getPolicyDocumentOutput({
    statements: [
        {
            effect: "Allow",
            actions: [
                "s3:GetReplicationConfiguration",
                "s3:ListBucket",
            ],
            resources: [source.arn],
        },
        {
            effect: "Allow",
            actions: [
                "s3:GetObjectVersionForReplication",
                "s3:GetObjectVersionAcl",
                "s3:GetObjectVersionTagging",
            ],
            resources: [pulumi.interpolate`${source.arn}/*`],
        },
        {
            effect: "Allow",
            actions: [
                "s3:ReplicateObject",
                "s3:ReplicateDelete",
                "s3:ReplicateTags",
            ],
            resources: [pulumi.interpolate`${destination.arn}/*`],
        },
    ],
});
const replicationRole = new aws.iam.Role("replicationRole", {assumeRolePolicy: assumeRoleDocument.apply(assumeRoleDocument => assumeRoleDocument.json)}, {
    provider: awsProvider,
});
const replicationPolicy = new aws.iam.Policy("replicationPolicy", {policy: replicationPolicyDocument.apply(replicationPolicyDocument => replicationPolicyDocument.json)}, {
    provider: awsProvider,
});
const replicationPolicyAttachment = new aws.iam.RolePolicyAttachment("replicationPolicyAttachment", {
    role: replicationRole.name,
    policyArn: replicationPolicy.arn,
}, {
    provider: awsProvider,
});
const destinationVersioning = new aws.s3.BucketVersioning("destinationVersioning", {
    bucket: destination.id,
    versioningConfiguration: {
        status: "Enabled",
    },
}, {
    provider: awsProvider,
});
const sourceBucketAcl = new aws.s3.BucketAcl("sourceBucketAcl", {
    region: "eu-central-1",
    bucket: source.id,
    acl: "private",
}, {
    provider: awsProvider,
});
const sourceVersioning = new aws.s3.BucketVersioning("sourceVersioning", {
    region: "eu-central-1",
    bucket: source.id,
    versioningConfiguration: {
        status: "Enabled",
    },
}, {
    provider: awsProvider,
});
const replication = new aws.s3.BucketReplicationConfig("replication", {
    region: "eu-central-1",
    role: replicationRole.arn,
    bucket: source.id,
    rules: [{
        id: "examplerule",
        filter: {
            prefix: "example",
        },
        status: "Enabled",
        destination: {
            bucket: destination.arn,
            storageClass: "STANDARD",
        },
    }],
}, {
    provider: awsProvider,
    dependsOn: [sourceVersioning],
});

```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi
import pulumi_aws as aws

assume_role_document = aws.iam.get_policy_document_output(statements=[{
    "effect": "Allow",
    "principals": [{
        "type": "Service",
        "identifiers": ["s3.amazonaws.com"],
    }],
    "actions": ["sts:AssumeRole"],
}])
aws_provider = aws.Provider("awsProvider", region="eu-west-1")
destination = aws.s3.Bucket("destination", opts = pulumi.ResourceOptions(provider=aws_provider))
source = aws.s3.Bucket("source", region="eu-central-1",
opts = pulumi.ResourceOptions(provider=aws_provider))
replication_policy_document = aws.iam.get_policy_document_output(statements=[
    {
        "effect": "Allow",
        "actions": [
            "s3:GetReplicationConfiguration",
            "s3:ListBucket",
        ],
        "resources": [source.arn],
    },
    {
        "effect": "Allow",
        "actions": [
            "s3:GetObjectVersionForReplication",
            "s3:GetObjectVersionAcl",
            "s3:GetObjectVersionTagging",
        ],
        "resources": [source.arn.apply(lambda arn: f"{arn}/*")],
    },
    {
        "effect": "Allow",
        "actions": [
            "s3:ReplicateObject",
            "s3:ReplicateDelete",
            "s3:ReplicateTags",
        ],
        "resources": [destination.arn.apply(lambda arn: f"{arn}/*")],
    },
])
replication_role = aws.iam.Role("replicationRole", assume_role_policy=assume_role_document.json,
opts = pulumi.ResourceOptions(provider=aws_provider))
replication_policy = aws.iam.Policy("replicationPolicy", policy=replication_policy_document.json,
opts = pulumi.ResourceOptions(provider=aws_provider))
replication_policy_attachment = aws.iam.RolePolicyAttachment("replicationPolicyAttachment",
    role=replication_role.name,
    policy_arn=replication_policy.arn,
    opts = pulumi.ResourceOptions(provider=aws_provider))
destination_versioning = aws.s3.BucketVersioning("destinationVersioning",
    bucket=destination.id,
    versioning_configuration={
        "status": "Enabled",
    },
    opts = pulumi.ResourceOptions(provider=aws_provider))
source_bucket_acl = aws.s3.BucketAcl("sourceBucketAcl",
    region="eu-central-1",
    bucket=source.id,
    acl="private",
    opts = pulumi.ResourceOptions(provider=aws_provider))
source_versioning = aws.s3.BucketVersioning("sourceVersioning",
    region="eu-central-1",
    bucket=source.id,
    versioning_configuration={
        "status": "Enabled",
    },
    opts = pulumi.ResourceOptions(provider=aws_provider))
replication = aws.s3.BucketReplicationConfig("replication",
    region="eu-central-1",
    role=replication_role.arn,
    bucket=source.id,
    rules=[{
        "id": "examplerule",
        "filter": {
            "prefix": "example",
        },
        "status": "Enabled",
        "destination": {
            "bucket": destination.arn,
            "storage_class": "STANDARD",
        },
    }],
    opts = pulumi.ResourceOptions(provider=aws_provider,
        depends_on=[source_versioning]))

```

{{% /choosable %}}

{{% choosable language go %}}

```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		assumeRoleDocument := iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
			Statements: iam.GetPolicyDocumentStatementArray{
				&iam.GetPolicyDocumentStatementArgs{
					Effect: pulumi.String("Allow"),
					Principals: iam.GetPolicyDocumentStatementPrincipalArray{
						&iam.GetPolicyDocumentStatementPrincipalArgs{
							Type: pulumi.String("Service"),
							Identifiers: pulumi.StringArray{
								pulumi.String("s3.amazonaws.com"),
							},
						},
					},
					Actions: pulumi.StringArray{
						pulumi.String("sts:AssumeRole"),
					},
				},
			},
		}, nil)
		awsProvider, err := aws.NewProvider(ctx, "awsProvider", &aws.ProviderArgs{
			Region: pulumi.String("eu-west-1"),
		})
		if err != nil {
			return err
		}
		destination, err := s3.NewBucket(ctx, "destination", nil, pulumi.Provider(awsProvider))
		if err != nil {
			return err
		}
		source, err := s3.NewBucket(ctx, "source", &s3.BucketArgs{
			Region: pulumi.String("eu-central-1"),
		}, pulumi.Provider(awsProvider))
		if err != nil {
			return err
		}
		replicationPolicyDocument := iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
			Statements: iam.GetPolicyDocumentStatementArray{
				&iam.GetPolicyDocumentStatementArgs{
					Effect: pulumi.String("Allow"),
					Actions: pulumi.StringArray{
						pulumi.String("s3:GetReplicationConfiguration"),
						pulumi.String("s3:ListBucket"),
					},
					Resources: pulumi.StringArray{
						source.Arn,
					},
				},
				&iam.GetPolicyDocumentStatementArgs{
					Effect: pulumi.String("Allow"),
					Actions: pulumi.StringArray{
						pulumi.String("s3:GetObjectVersionForReplication"),
						pulumi.String("s3:GetObjectVersionAcl"),
						pulumi.String("s3:GetObjectVersionTagging"),
					},
					Resources: pulumi.StringArray{
						source.Arn.ApplyT(func(arn string) (string, error) {
							return fmt.Sprintf("%v/*", arn), nil
						}).(pulumi.StringOutput),
					},
				},
				&iam.GetPolicyDocumentStatementArgs{
					Effect: pulumi.String("Allow"),
					Actions: pulumi.StringArray{
						pulumi.String("s3:ReplicateObject"),
						pulumi.String("s3:ReplicateDelete"),
						pulumi.String("s3:ReplicateTags"),
					},
					Resources: pulumi.StringArray{
						destination.Arn.ApplyT(func(arn string) (string, error) {
							return fmt.Sprintf("%v/*", arn), nil
						}).(pulumi.StringOutput),
					},
				},
			},
		}, nil)
		replicationRole, err := iam.NewRole(ctx, "replicationRole", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.String(assumeRoleDocument.ApplyT(func(assumeRoleDocument iam.GetPolicyDocumentResult) (*string, error) {
				return &assumeRoleDocument.Json, nil
			}).(pulumi.StringPtrOutput)),
		}, pulumi.Provider(awsProvider))
		if err != nil {
			return err
		}
		replicationPolicy, err := iam.NewPolicy(ctx, "replicationPolicy", &iam.PolicyArgs{
			Policy: pulumi.String(replicationPolicyDocument.ApplyT(func(replicationPolicyDocument iam.GetPolicyDocumentResult) (*string, error) {
				return &replicationPolicyDocument.Json, nil
			}).(pulumi.StringPtrOutput)),
		}, pulumi.Provider(awsProvider))
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicyAttachment(ctx, "replicationPolicyAttachment", &iam.RolePolicyAttachmentArgs{
			Role:      replicationRole.Name,
			PolicyArn: replicationPolicy.Arn,
		}, pulumi.Provider(awsProvider))
		if err != nil {
			return err
		}
		_, err = s3.NewBucketVersioning(ctx, "destinationVersioning", &s3.BucketVersioningArgs{
			Bucket: destination.ID(),
			VersioningConfiguration: &s3.BucketVersioningVersioningConfigurationArgs{
				Status: pulumi.String("Enabled"),
			},
		}, pulumi.Provider(awsProvider))
		if err != nil {
			return err
		}
		_, err = s3.NewBucketAcl(ctx, "sourceBucketAcl", &s3.BucketAclArgs{
			Region: pulumi.String("eu-central-1"),
			Bucket: source.ID(),
			Acl:    pulumi.String("private"),
		}, pulumi.Provider(awsProvider))
		if err != nil {
			return err
		}
		sourceVersioning, err := s3.NewBucketVersioning(ctx, "sourceVersioning", &s3.BucketVersioningArgs{
			Region: pulumi.String("eu-central-1"),
			Bucket: source.ID(),
			VersioningConfiguration: &s3.BucketVersioningVersioningConfigurationArgs{
				Status: pulumi.String("Enabled"),
			},
		}, pulumi.Provider(awsProvider))
		if err != nil {
			return err
		}
		_, err = s3.NewBucketReplicationConfig(ctx, "replication", &s3.BucketReplicationConfigArgs{
			Region: pulumi.String("eu-central-1"),
			Role:   replicationRole.Arn,
			Bucket: source.ID(),
			Rules: s3.BucketReplicationConfigRuleArray{
				&s3.BucketReplicationConfigRuleArgs{
					Id: pulumi.String("examplerule"),
					Filter: &s3.BucketReplicationConfigRuleFilterArgs{
						Prefix: pulumi.String("example"),
					},
					Status: pulumi.String("Enabled"),
					Destination: &s3.BucketReplicationConfigRuleDestinationArgs{
						Bucket:       destination.Arn,
						StorageClass: pulumi.String("STANDARD"),
					},
				},
			},
		}, pulumi.Provider(awsProvider), pulumi.DependsOn([]pulumi.Resource{
			sourceVersioning,
		}))
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
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var assumeRoleDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Effect = "Allow",
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Type = "Service",
                        Identifiers = new[]
                        {
                            "s3.amazonaws.com",
                        },
                    },
                },
                Actions = new[]
                {
                    "sts:AssumeRole",
                },
            },
        },
    });

    var awsProvider = new Aws.Provider("awsProvider", new()
    {
        Region = "eu-west-1",
    });

    var destination = new Aws.S3.Bucket("destination", new()
    {
    }, new CustomResourceOptions
    {
        Provider = awsProvider,
    });

    var source = new Aws.S3.Bucket("source", new()
    {
        Region = "eu-central-1",
    }, new CustomResourceOptions
    {
        Provider = awsProvider,
    });

    var replicationPolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Effect = "Allow",
                Actions = new[]
                {
                    "s3:GetReplicationConfiguration",
                    "s3:ListBucket",
                },
                Resources = new[]
                {
                    source.Arn,
                },
            },
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Effect = "Allow",
                Actions = new[]
                {
                    "s3:GetObjectVersionForReplication",
                    "s3:GetObjectVersionAcl",
                    "s3:GetObjectVersionTagging",
                },
                Resources = new[]
                {
                    $"{source.Arn}/*",
                },
            },
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Effect = "Allow",
                Actions = new[]
                {
                    "s3:ReplicateObject",
                    "s3:ReplicateDelete",
                    "s3:ReplicateTags",
                },
                Resources = new[]
                {
                    $"{destination.Arn}/*",
                },
            },
        },
    });

    var replicationRole = new Aws.Iam.Role("replicationRole", new()
    {
        AssumeRolePolicy = assumeRoleDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    }, new CustomResourceOptions
    {
        Provider = awsProvider,
    });

    var replicationPolicy = new Aws.Iam.Policy("replicationPolicy", new()
    {
        PolicyDocument = replicationPolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    }, new CustomResourceOptions
    {
        Provider = awsProvider,
    });

    var replicationPolicyAttachment = new Aws.Iam.RolePolicyAttachment("replicationPolicyAttachment", new()
    {
        Role = replicationRole.Name,
        PolicyArn = replicationPolicy.Arn,
    }, new CustomResourceOptions
    {
        Provider = awsProvider,
    });

    var destinationVersioning = new Aws.S3.BucketVersioning("destinationVersioning", new()
    {
        Bucket = destination.Id,
        VersioningConfiguration = new Aws.S3.Inputs.BucketVersioningVersioningConfigurationArgs
        {
            Status = "Enabled",
        },
    }, new CustomResourceOptions
    {
        Provider = awsProvider,
    });

    var sourceBucketAcl = new Aws.S3.BucketAcl("sourceBucketAcl", new()
    {
        Region = "eu-central-1",
        Bucket = source.Id,
        Acl = "private",
    }, new CustomResourceOptions
    {
        Provider = awsProvider,
    });

    var sourceVersioning = new Aws.S3.BucketVersioning("sourceVersioning", new()
    {
        Region = "eu-central-1",
        Bucket = source.Id,
        VersioningConfiguration = new Aws.S3.Inputs.BucketVersioningVersioningConfigurationArgs
        {
            Status = "Enabled",
        },
    }, new CustomResourceOptions
    {
        Provider = awsProvider,
    });

    var replication = new Aws.S3.BucketReplicationConfig("replication", new()
    {
        Region = "eu-central-1",
        Role = replicationRole.Arn,
        Bucket = source.Id,
        Rules = new[]
        {
            new Aws.S3.Inputs.BucketReplicationConfigRuleArgs
            {
                Id = "examplerule",
                Filter = new Aws.S3.Inputs.BucketReplicationConfigRuleFilterArgs
                {
                    Prefix = "example",
                },
                Status = "Enabled",
                Destination = new Aws.S3.Inputs.BucketReplicationConfigRuleDestinationArgs
                {
                    Bucket = destination.Arn,
                    StorageClass = "STANDARD",
                },
            },
        },
    }, new CustomResourceOptions
    {
        Provider = awsProvider,
        DependsOn =
        {
            sourceVersioning,
        },
    });

});


```

{{% /choosable %}}

{{% choosable language java %}}

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.Provider;
import com.pulumi.aws.ProviderArgs;
import com.pulumi.aws.s3.Bucket;
import com.pulumi.aws.s3.BucketArgs;
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.iam.Policy;
import com.pulumi.aws.iam.PolicyArgs;
import com.pulumi.aws.iam.RolePolicyAttachment;
import com.pulumi.aws.iam.RolePolicyAttachmentArgs;
import com.pulumi.aws.s3.BucketVersioning;
import com.pulumi.aws.s3.BucketVersioningArgs;
import com.pulumi.aws.s3.inputs.BucketVersioningVersioningConfigurationArgs;
import com.pulumi.aws.s3.BucketAcl;
import com.pulumi.aws.s3.BucketAclArgs;
import com.pulumi.aws.s3.BucketReplicationConfig;
import com.pulumi.aws.s3.BucketReplicationConfigArgs;
import com.pulumi.aws.s3.inputs.BucketReplicationConfigRuleArgs;
import com.pulumi.aws.s3.inputs.BucketReplicationConfigRuleFilterArgs;
import com.pulumi.aws.s3.inputs.BucketReplicationConfigRuleDestinationArgs;
import com.pulumi.resources.CustomResourceOptions;
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
        final var assumeRoleDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("Service")
                    .identifiers("s3.amazonaws.com")
                    .build())
                .actions("sts:AssumeRole")
                .build())
            .build());

        var awsProvider = new Provider("awsProvider", ProviderArgs.builder()
            .region("eu-west-1")
            .build());

        var destination = new Bucket("destination", BucketArgs.Empty, CustomResourceOptions.builder()
            .provider(awsProvider)
            .build());

        var source = new Bucket("source", BucketArgs.builder()
            .region("eu-central-1")
            .build(), CustomResourceOptions.builder()
                .provider(awsProvider)
                .build());

        final var replicationPolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(            
                GetPolicyDocumentStatementArgs.builder()
                    .effect("Allow")
                    .actions(                    
                        "s3:GetReplicationConfiguration",
                        "s3:ListBucket")
                    .resources(source.arn())
                    .build(),
                GetPolicyDocumentStatementArgs.builder()
                    .effect("Allow")
                    .actions(                    
                        "s3:GetObjectVersionForReplication",
                        "s3:GetObjectVersionAcl",
                        "s3:GetObjectVersionTagging")
                    .resources(source.arn().applyValue(_arn -> String.format("%s/*", _arn)))
                    .build(),
                GetPolicyDocumentStatementArgs.builder()
                    .effect("Allow")
                    .actions(                    
                        "s3:ReplicateObject",
                        "s3:ReplicateDelete",
                        "s3:ReplicateTags")
                    .resources(destination.arn().applyValue(_arn -> String.format("%s/*", _arn)))
                    .build())
            .build());

        var replicationRole = new Role("replicationRole", RoleArgs.builder()
            .assumeRolePolicy(assumeRoleDocument.applyValue(_assumeRoleDocument -> _assumeRoleDocument.json()))
            .build(), CustomResourceOptions.builder()
                .provider(awsProvider)
                .build());

        var replicationPolicy = new Policy("replicationPolicy", PolicyArgs.builder()
            .policy(replicationPolicyDocument.applyValue(_replicationPolicyDocument -> _replicationPolicyDocument.json()))
            .build(), CustomResourceOptions.builder()
                .provider(awsProvider)
                .build());

        var replicationPolicyAttachment = new RolePolicyAttachment("replicationPolicyAttachment", RolePolicyAttachmentArgs.builder()
            .role(replicationRole.name())
            .policyArn(replicationPolicy.arn())
            .build(), CustomResourceOptions.builder()
                .provider(awsProvider)
                .build());

        var destinationVersioning = new BucketVersioning("destinationVersioning", BucketVersioningArgs.builder()
            .bucket(destination.id())
            .versioningConfiguration(BucketVersioningVersioningConfigurationArgs.builder()
                .status("Enabled")
                .build())
            .build(), CustomResourceOptions.builder()
                .provider(awsProvider)
                .build());

        var sourceBucketAcl = new BucketAcl("sourceBucketAcl", BucketAclArgs.builder()
            .region("eu-central-1")
            .bucket(source.id())
            .acl("private")
            .build(), CustomResourceOptions.builder()
                .provider(awsProvider)
                .build());

        var sourceVersioning = new BucketVersioning("sourceVersioning", BucketVersioningArgs.builder()
            .region("eu-central-1")
            .bucket(source.id())
            .versioningConfiguration(BucketVersioningVersioningConfigurationArgs.builder()
                .status("Enabled")
                .build())
            .build(), CustomResourceOptions.builder()
                .provider(awsProvider)
                .build());

        var replication = new BucketReplicationConfig("replication", BucketReplicationConfigArgs.builder()
            .region("eu-central-1")
            .role(replicationRole.arn())
            .bucket(source.id())
            .rules(BucketReplicationConfigRuleArgs.builder()
                .id("examplerule")
                .filter(BucketReplicationConfigRuleFilterArgs.builder()
                    .prefix("example")
                    .build())
                .status("Enabled")
                .destination(BucketReplicationConfigRuleDestinationArgs.builder()
                    .bucket(destination.arn())
                    .storageClass("STANDARD")
                    .build())
                .build())
            .build(), CustomResourceOptions.builder()
                .provider(awsProvider)
                .dependsOn(sourceVersioning)
                .build());

    }
}

```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
name: s3-replication
runtime: yaml

variables:
  assumeRoleDocument:
    fn::invoke:
      function: aws:iam:getPolicyDocument
      arguments:
        statements:
          - effect: "Allow"
            principals:
              - type: "Service"
                identifiers:
                  - s3.amazonaws.com
            actions:
              - sts:AssumeRole

  replicationPolicyDocument:
    fn::invoke:
      function: aws:iam:getPolicyDocument
      arguments:
        statements:
          - effect: "Allow"
            actions:
              - s3:GetReplicationConfiguration
              - s3:ListBucket
            resources:
              - ${source.arn}
          - effect: "Allow"
            actions:
              - s3:GetObjectVersionForReplication
              - s3:GetObjectVersionAcl
              - s3:GetObjectVersionTagging
            resources:
              - "${source.arn}/*"
          - effect: "Allow"
            actions:
              - s3:ReplicateObject
              - s3:ReplicateDelete
              - s3:ReplicateTags
            resources:
              - "${destination.arn}/*"

resources:
  # AWS Provider for eu-west-1
  awsProvider:
    type: pulumi:providers:aws
    properties:
      region: "eu-west-1"

  # IAM Role for replication
  replicationRole:
    type: aws:iam:Role
    properties:
      assumeRolePolicy: ${assumeRoleDocument.json}
    options:
      provider: ${awsProvider}

  # Destination bucket in eu-west-1
  destination:
    type: aws:s3:Bucket
    options:
      provider: ${awsProvider}

  # Source bucket in eu-central-1
  source:
    type: aws:s3:Bucket
    properties:
      region: "eu-central-1"
    options:
      provider: ${awsProvider}

  # IAM Policy for replication
  replicationPolicy:
    type: aws:iam:Policy
    properties:
      policy: ${replicationPolicyDocument.json}
    options:
      provider: ${awsProvider}

  # Attach policy to role
  replicationPolicyAttachment:
    type: aws:iam:RolePolicyAttachment
    properties:
      role: ${replicationRole.name}
      policyArn: ${replicationPolicy.arn}
    options:
      provider: ${awsProvider}

  # Enable versioning on destination bucket
  destinationVersioning:
    type: aws:s3:BucketVersioning
    properties:
      bucket: ${destination.id}
      versioningConfiguration:
        status: "Enabled"
    options:
      provider: ${awsProvider}

  # Set ACL on source bucket
  sourceBucketAcl:
    type: aws:s3:BucketAcl
    properties:
      region: "eu-central-1"
      bucket: ${source.id}
      acl: "private"
    options:
      provider: ${awsProvider}

  # Enable versioning on source bucket
  sourceVersioning:
    type: aws:s3:BucketVersioning
    properties:
      region: "eu-central-1"
      bucket: ${source.id}
      versioningConfiguration:
        status: "Enabled"
    options:
      provider: ${awsProvider}

  # Configure bucket replication
  replication:
    type: aws:s3:BucketReplicationConfig
    properties:
      region: "eu-central-1"
      role: ${replicationRole.arn}
      bucket: ${source.id}
      rules:
        - id: "examplerule"
          filter:
            prefix: "example"
          status: "Enabled"
          destination:
            bucket: ${destination.arn}
            storageClass: "STANDARD"
    options:
      provider: ${awsProvider}
      dependsOn:
        - ${sourceVersioning}
```

{{% /choosable %}}

{{< /chooser >}}

## Non-region-aware resources

This section lists resources that are not Region-aware--meaning `region` has not been added to them.

Some resources, such as [IAM and STS](https://docs.aws.amazon.com/IAM/latest/UserGuide/programming.html#IAMEndpoints) are [global](https://docs.aws.amazon.com/whitepapers/latest/aws-fault-isolation-boundaries/global-services.html) and exist in all Regions within a partition.

Other resources are not Region-aware because they already had a top-level `region`, are inherently global, or because adding `region` would not be appropriate for other reasons.

### Resources deprecating `region`

The following regional resources and data sources had a top-level `region` prior to version 7.0.0. It is now deprecated and will be replaced in a future version to support the new Region-aware behavior.

* `aws.cloudformation.StackSetInstance` resource
* `aws.config.AggregateAuthorization` resource
* `aws.directconnect.HostedConnection` resource
* `aws.getRegion` function
* `aws.s3.getBucket` function
* `aws.servicequotas.Template` resource
* `aws.servicequotas.getTemplates` function
* `aws.ssmincidents.ReplicationSet` resource and `aws.ssmincidents.getReplicationSet` function
* `aws.ec2.getVpcEndpointService` function
* `aws.ec2.getVpcPeeringConnection` function

### Global services

All resources for the following services are considered _global_:

* Account Management (`aws.organizations.*`)
* Billing (`aws.billing.*`)
* Billing and Cost Management Data Exports (`aws.bcmdataexports.*`)
* Budgets (`aws.budgets.*`)
* CloudFront (`aws.cloudfront.*` and `aws.cloudfrontkeyvaluestore.*`)
* Cost Explorer (`aws.ce.*`)
* Cost Optimization Hub (`aws.costoptimizationhub.*`)
* Cost and Usage Report (`aws.cur.*`)
* Global Accelerator (`aws.globalaccelerator.*`)
* IAM (`aws.iam.*`, `aws.rolesanywhere.*` and `aws.getCallerIdentity`)
* Network Manager (`aws.networkmanager.*`)
* Organizations (`aws.organizations.*`)
* Price List (`aws.pricing.*`)
* Route 53 (`aws.route53.*` and `aws.route53domains.*`)
* Route 53 ARC (`aws.route53recoverycontrolconfig.*` and `aws.route53recoveryreadiness.*`)
* Shield Advanced (`aws.shield.*`)
* User Notifications (`aws.notifications.*`)
* User Notifications Contacts (`aws.notificationscontacts.*`)
* WAF Classic (`aws.waf.*`)

### Global resources in regional services

Some regional services have a subset of resources that are global:

| Service | Type | Name |
|---|---|---|
| Backup | Resource | `aws.backup.GlobalSettings` |
| Chime SDK Voice | Resource | `aws.chimesdkvoice.GlobalSettings` |
| CloudTrail | Resource | `aws.cloudtrail.OrganizationDelegatedAdminAccount` |
| Direct Connect | Resource | `aws.directconnect.Gateway` |
| Direct Connect | Function | `aws.directconnect.getGateway` |
| EC2 | Resource | `aws.ec2.ImageBlockPublicAccess` |
| Firewall Manager | Resource | `aws.fms.AdminAccount` |
| IPAM | Resource | `aws.ec2.VpcIpamOrganizationAdminAccount` |
| QuickSight | Resource | `aws.quicksight.AccountSettings` |
| Resource Access Manager | Resource | `aws.ram.SharingWithOrganization` |
| S3 | Function | `aws.s3.getCanonicalUserId` |
| S3 | Resource | `aws.s3.AccountPublicAccessBlock` |
| S3 | Function | `aws.s3.getAccountPublicAccessBlock` |
| Service Catalog | Resource | `aws.servicecatalog.OrganizationsAccess` |

### Meta data sources

The `aws.getDefaultTags`, `aws.getPartition`, and `aws.getRegions` functions are effectively global.

`region` of the `aws.getArn` function stays as-is.

### Policy Document Data Sources

Some data sources convert data into JSON policy documents and are effectively global:

* `aws.cloudwatch.getLogDataProtectionPolicyDocument`
* `aws.ecr.getLifecyclePolicyDocument`
