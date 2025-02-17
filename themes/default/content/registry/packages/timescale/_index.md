---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/timescale/timescale/1.14.0/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Timescale Provider
meta_desc: Provides an overview on how to configure the Pulumi Timescale provider.
layout: package
---

## Generate Provider

The Timescale provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider timescale/timescale
```
## Overview

The Pulumi provider for [Timescale](https://www.timescale.com/cloud).
## Requirements
- Pulumi >= 1.0
## Quick Start
### Authorization
When you log in to your [Timescale Account](https://console.cloud.timescale.com/), navigate to the `Project settings` page.
From here, you can create client credentials for programmatic usage. Click the `Create credentials` button to generate a new public/secret key pair.
### Project ID
The project ID can be found on the `Project settings` page.

Create a `main.tf` configuration file with the following content.
### VPC Peering

Since v1.9.0 it is possible to peer Timescale VPCs to AWS VPCs using pulumi.

Below is a minimal working example:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    timescale:accessKey:
        value: 'TODO: var.ts_access_key'
    timescale:projectId:
        value: 'TODO: var.ts_project_id'
    timescale:secretKey:
        value: 'TODO: var.ts_secret_key'

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as timescale from "@pulumi/timescale";

const config = new pulumi.Config();
const awsAccountId = config.require("awsAccountId");
const awsRegion = config.get("awsRegion") || "us-east-1";
const tsProjectId = config.require("tsProjectId");
const tsAccessKey = config.require("tsAccessKey");
const tsSecretKey = config.require("tsSecretKey");
const tsRegion = config.get("tsRegion") || "us-east-1";
const main = new timescale.index/vpcs.Vpcs("main", {
    cidr: "10.10.0.0/16",
    name: "vpc_name",
    regionCode: tsRegion,
});
const mainVpc = new aws.index.Vpc("main", {cidrBlock: "10.0.1.0/24"});
// Requester's side of the peering connection.
const peer = new timescale.index/peeringConnection.PeeringConnection("peer", {
    peerAccountId: awsAccountId,
    peerRegionCode: awsRegion,
    peerVpcId: mainVpc.id,
    timescaleVpcId: main.vpcsId,
});
// Accepter's side of the peering connection.
const peerVpcPeeringConnectionAccepter = new aws.index.VpcPeeringConnectionAccepter("peer", {
    vpcPeeringConnectionId: peer.provisionedId,
    autoAccept: true,
}, {
    dependsOn: [peer],
});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    timescale:accessKey:
        value: 'TODO: var.ts_access_key'
    timescale:projectId:
        value: 'TODO: var.ts_project_id'
    timescale:secretKey:
        value: 'TODO: var.ts_secret_key'

```
```python
import pulumi
import pulumi_aws as aws
import pulumi_timescale as timescale

config = pulumi.Config()
aws_account_id = config.require("awsAccountId")
aws_region = config.get("awsRegion")
if aws_region is None:
    aws_region = "us-east-1"
ts_project_id = config.require("tsProjectId")
ts_access_key = config.require("tsAccessKey")
ts_secret_key = config.require("tsSecretKey")
ts_region = config.get("tsRegion")
if ts_region is None:
    ts_region = "us-east-1"
main = timescale.index.vpcs.Vpcs("main",
    cidr=10.10.0.0/16,
    name=vpc_name,
    region_code=ts_region)
main_vpc = aws.index.Vpc("main", cidr_block=10.0.1.0/24)
# Requester's side of the peering connection.
peer = timescale.index.peering_connection.PeeringConnection("peer",
    peer_account_id=aws_account_id,
    peer_region_code=aws_region,
    peer_vpc_id=main_vpc.id,
    timescale_vpc_id=main.vpcs_id)
# Accepter's side of the peering connection.
peer_vpc_peering_connection_accepter = aws.index.VpcPeeringConnectionAccepter("peer",
    vpc_peering_connection_id=peer.provisioned_id,
    auto_accept=True,
    opts = pulumi.ResourceOptions(depends_on=[peer]))
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    timescale:accessKey:
        value: 'TODO: var.ts_access_key'
    timescale:projectId:
        value: 'TODO: var.ts_project_id'
    timescale:secretKey:
        value: 'TODO: var.ts_secret_key'

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;
using Timescale = Pulumi.Timescale;

return await Deployment.RunAsync(() =>
{
    var config = new Config();
    var awsAccountId = config.Require("awsAccountId");
    var awsRegion = config.Get("awsRegion") ?? "us-east-1";
    var tsProjectId = config.Require("tsProjectId");
    var tsAccessKey = config.Require("tsAccessKey");
    var tsSecretKey = config.Require("tsSecretKey");
    var tsRegion = config.Get("tsRegion") ?? "us-east-1";
    var main = new Timescale.Index.Vpcs.Vpcs("main", new()
    {
        Cidr = "10.10.0.0/16",
        Name = "vpc_name",
        RegionCode = tsRegion,
    });

    var mainVpc = new Aws.Index.Vpc("main", new()
    {
        CidrBlock = "10.0.1.0/24",
    });

    // Requester's side of the peering connection.
    var peer = new Timescale.Index.PeeringConnection.PeeringConnection("peer", new()
    {
        PeerAccountId = awsAccountId,
        PeerRegionCode = awsRegion,
        PeerVpcId = mainVpc.Id,
        TimescaleVpcId = main.VpcsId,
    });

    // Accepter's side of the peering connection.
    var peerVpcPeeringConnectionAccepter = new Aws.Index.VpcPeeringConnectionAccepter("peer", new()
    {
        VpcPeeringConnectionId = peer.ProvisionedId,
        AutoAccept = true,
    }, new CustomResourceOptions
    {
        DependsOn =
        {
            peer,
        },
    });

});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    timescale:accessKey:
        value: 'TODO: var.ts_access_key'
    timescale:projectId:
        value: 'TODO: var.ts_project_id'
    timescale:secretKey:
        value: 'TODO: var.ts_secret_key'

```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/go/aws"
	"github.com/pulumi/pulumi-timescale/sdk/go/timescale"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")
		awsAccountId := cfg.Require("awsAccountId")
		awsRegion := "us-east-1"
		if param := cfg.Get("awsRegion"); param != "" {
			awsRegion = param
		}
		tsProjectId := cfg.Require("tsProjectId")
		tsAccessKey := cfg.Require("tsAccessKey")
		tsSecretKey := cfg.Require("tsSecretKey")
		tsRegion := "us-east-1"
		if param := cfg.Get("tsRegion"); param != "" {
			tsRegion = param
		}
		main, err := index / vpcs.NewVpcs(ctx, "main", &index/vpcs.VpcsArgs{
			Cidr:       "10.10.0.0/16",
			Name:       "vpc_name",
			RegionCode: tsRegion,
		})
		if err != nil {
			return err
		}
		mainVpc, err := aws.NewVpc(ctx, "main", &aws.VpcArgs{
			CidrBlock: "10.0.1.0/24",
		})
		if err != nil {
			return err
		}
		// Requester's side of the peering connection.
		peer, err := index / peeringconnection.NewPeeringConnection(ctx, "peer", &index/peeringconnection.PeeringConnectionArgs{
			PeerAccountId:  awsAccountId,
			PeerRegionCode: awsRegion,
			PeerVpcId:      mainVpc.Id,
			TimescaleVpcId: main.VpcsId,
		})
		if err != nil {
			return err
		}
		// Accepter's side of the peering connection.
		_, err = aws.NewVpcPeeringConnectionAccepter(ctx, "peer", &aws.VpcPeeringConnectionAccepterArgs{
			VpcPeeringConnectionId: peer.ProvisionedId,
			AutoAccept:             true,
		}, pulumi.DependsOn([]pulumi.Resource{
			peer,
		}))
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
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
config:
    timescale:accessKey:
        value: 'TODO: var.ts_access_key'
    timescale:projectId:
        value: 'TODO: var.ts_project_id'
    timescale:secretKey:
        value: 'TODO: var.ts_secret_key'

```
```yaml
configuration:
  awsAccountId:
    type: string
  awsRegion:
    type: string
    default: us-east-1
  tsProjectId:
    type: string
  tsAccessKey:
    type: string
  tsSecretKey:
    type: string
  tsRegion:
    type: string
    default: us-east-1
resources:
  main:
    type: timescale:Vpcs
    properties:
      cidr: 10.10.0.0/16
      name: vpc_name
      regionCode: ${tsRegion}
  mainVpc:
    type: aws:Vpc
    name: main
    properties:
      cidrBlock: 10.0.1.0/24
  # Requester's side of the peering connection.
  peer:
    type: timescale:PeeringConnection
    properties:
      peerAccountId: ${awsAccountId}
      peerRegionCode: ${awsRegion}
      peerVpcId: ${mainVpc.id}
      timescaleVpcId: ${main.vpcsId}
  # Accepter's side of the peering connection.
  peerVpcPeeringConnectionAccepter:
    type: aws:VpcPeeringConnectionAccepter
    name: peer
    properties:
      vpcPeeringConnectionId: ${peer.provisionedId}
      autoAccept: true
    options:
      dependsOn:
        - ${peer}
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    timescale:accessKey:
        value: 'TODO: var.ts_access_key'
    timescale:projectId:
        value: 'TODO: var.ts_project_id'
    timescale:secretKey:
        value: 'TODO: var.ts_secret_key'

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.timescale.index_vpcs.Vpcs;
import com.pulumi.timescale.index_vpcs.VpcsArgs;
import com.pulumi.aws.Vpc;
import com.pulumi.aws.VpcArgs;
import com.pulumi.timescale.PeeringConnection;
import com.pulumi.timescale.PeeringConnectionArgs;
import com.pulumi.aws.VpcPeeringConnectionAccepter;
import com.pulumi.aws.VpcPeeringConnectionAccepterArgs;
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
        final var config = ctx.config();
        final var awsAccountId = config.get("awsAccountId");
        final var awsRegion = config.get("awsRegion").orElse("us-east-1");
        final var tsProjectId = config.get("tsProjectId");
        final var tsAccessKey = config.get("tsAccessKey");
        final var tsSecretKey = config.get("tsSecretKey");
        final var tsRegion = config.get("tsRegion").orElse("us-east-1");
        var main = new Vpcs("main", VpcsArgs.builder()
            .cidr("10.10.0.0/16")
            .name("vpc_name")
            .regionCode(tsRegion)
            .build());

        var mainVpc = new Vpc("mainVpc", VpcArgs.builder()
            .cidrBlock("10.0.1.0/24")
            .build());

        // Requester's side of the peering connection.
        var peer = new PeeringConnection("peer", PeeringConnectionArgs.builder()
            .peerAccountId(awsAccountId)
            .peerRegionCode(awsRegion)
            .peerVpcId(mainVpc.id())
            .timescaleVpcId(main.vpcsId())
            .build());

        // Accepter's side of the peering connection.
        var peerVpcPeeringConnectionAccepter = new VpcPeeringConnectionAccepter("peerVpcPeeringConnectionAccepter", VpcPeeringConnectionAccepterArgs.builder()
            .vpcPeeringConnectionId(peer.provisionedId())
            .autoAccept(true)
            .build(), CustomResourceOptions.builder()
                .dependsOn(peer)
                .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}

Note that this configuration may fail on first apply, as the value of
`timescale_peering_connection.peer.provisioned_id` (starting with `pcx-`) may
not be immediately available. This typically happens due to the asynchronous
nature of the VPC peering request and its acceptance process. In this case, a
second `pulumi up` can be run to ensure everything is applied.
## Supported Service Configurations
### Compute
- 500m CPU / 2 GB Memory
- 1000m CPU / 4 GB Memory
- 2000m CPU / 8 GB Memory
- 4000m CPU / 16 GB Memory
- 8000m CPU / 32 GB Memory
- 16000m CPU / 64 GB Memory
- 32000m CPU / 128 GB Memory
### Storage
Since June 2023, you no longer need to allocate a fixed storage volume or worry about managing your disk size, and you'll be billed only for the storage you actually use.
See more info in our [blogpost](https://www.timescale.com/blog/savings-unlocked-why-we-switched-to-a-pay-for-what-you-store-database-storage-model/)
## Supported Operations
✅ Create service <br />
✅ Rename service <br />
✅ Resize service <br />
✅ Pause/resume service <br />
✅ Delete service <br />
✅ Import service <br />
✅ Enable High Availability replicas <br />
✅ Enable read replicas <br />
✅ VPC peering <br />
✅ Connection pooling <br />
## Billing
Services are currently billed for hourly usage. If a service is running for less than an hour,
it will still be charged for the full hour of usage.