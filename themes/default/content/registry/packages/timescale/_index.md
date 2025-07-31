---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/timescale/timescale/2.4.0/index.md
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
When you log in to your [Timescale Account](https://console.cloud.timescale.com/), click on your project name on the upper left-hand side of the page and go to the `Project settings` page.
From here, you can create client credentials for programmatic usage. Click the `Create credentials` button to generate a new public/secret key pair.

Find more information on creating Client Credentials in the [Timescale docs](https://docs.timescale.com/use-timescale/latest/security/client-credentials/#creating-client-credentials).
### Project ID

To view the project ID, click on your project name on the upper left-hand side of the page.
### Example files and usage
#### Service with HA replica and pooler

> [!NOTE]The example file creates:
> * A single instance called `tf-test` that contains:
>   * 0.5 CPUs
>   * 2GB of RAM
>   * the region set to `us-west-2`
>   * an HA replica
>   * the connection pooler enabled
> * Outputs to display the connection info for:
>   * the primary hostname and port
>   * the ha-replica hostname and port
>   * the pooler hostname and port

Create a `main.tf` configuration file with the following content.
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
import * as timescale from "@pulumi/timescale";

const config = new pulumi.Config();
const tsProjectId = config.require("tsProjectId");
const tsAccessKey = config.require("tsAccessKey");
const tsSecretKey = config.require("tsSecretKey");
const tf_test = new timescale.Service("tf-test", {
    name: "tf-test",
    milliCpu: 500,
    memoryGb: 2,
    regionCode: "us-west-2",
    connectionPoolerEnabled: true,
    enableHaReplica: true,
});
export const hostAddr = tf_test.hostname;
export const hostPort = tf_test.port;
export const replicaAddr = tf_test.replicaHostname;
export const replicaPort = tf_test.replicaPort;
export const poolerAddr = tf_test.poolerHostname;
export const poolerPort = tf_test.poolerPort;
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
import pulumi_timescale as timescale

config = pulumi.Config()
ts_project_id = config.require("tsProjectId")
ts_access_key = config.require("tsAccessKey")
ts_secret_key = config.require("tsSecretKey")
tf_test = timescale.Service("tf-test",
    name="tf-test",
    milli_cpu=500,
    memory_gb=2,
    region_code="us-west-2",
    connection_pooler_enabled=True,
    enable_ha_replica=True)
pulumi.export("hostAddr", tf_test.hostname)
pulumi.export("hostPort", tf_test.port)
pulumi.export("replicaAddr", tf_test.replica_hostname)
pulumi.export("replicaPort", tf_test.replica_port)
pulumi.export("poolerAddr", tf_test.pooler_hostname)
pulumi.export("poolerPort", tf_test.pooler_port)
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
using Timescale = Pulumi.Timescale;

return await Deployment.RunAsync(() =>
{
    var config = new Config();
    var tsProjectId = config.Require("tsProjectId");
    var tsAccessKey = config.Require("tsAccessKey");
    var tsSecretKey = config.Require("tsSecretKey");
    var tf_test = new Timescale.Service("tf-test", new()
    {
        Name = "tf-test",
        MilliCpu = 500,
        MemoryGb = 2,
        RegionCode = "us-west-2",
        ConnectionPoolerEnabled = true,
        EnableHaReplica = true,
    });

    return new Dictionary<string, object?>
    {
        ["hostAddr"] = tf_test.Hostname,
        ["hostPort"] = tf_test.Port,
        ["replicaAddr"] = tf_test.ReplicaHostname,
        ["replicaPort"] = tf_test.ReplicaPort,
        ["poolerAddr"] = tf_test.PoolerHostname,
        ["poolerPort"] = tf_test.PoolerPort,
    };
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
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/timescale/v2/timescale"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")
		tsProjectId := cfg.Require("tsProjectId")
		tsAccessKey := cfg.Require("tsAccessKey")
		tsSecretKey := cfg.Require("tsSecretKey")
		tf_test, err := timescale.NewService(ctx, "tf-test", &timescale.ServiceArgs{
			Name:                    pulumi.String("tf-test"),
			MilliCpu:                pulumi.Float64(500),
			MemoryGb:                pulumi.Float64(2),
			RegionCode:              pulumi.String("us-west-2"),
			ConnectionPoolerEnabled: pulumi.Bool(true),
			EnableHaReplica:         pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		ctx.Export("hostAddr", tf_test.Hostname)
		ctx.Export("hostPort", tf_test.Port)
		ctx.Export("replicaAddr", tf_test.ReplicaHostname)
		ctx.Export("replicaPort", tf_test.ReplicaPort)
		ctx.Export("poolerAddr", tf_test.PoolerHostname)
		ctx.Export("poolerPort", tf_test.PoolerPort)
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
  tsProjectId:
    type: string
  tsAccessKey:
    type: string
  tsSecretKey:
    type: string
resources:
  tf-test:
    type: timescale:Service
    properties:
      name: tf-test
      milliCpu: 500
      memoryGb: 2
      regionCode: us-west-2
      connectionPoolerEnabled: true
      enableHaReplica: true
outputs:
  ## host connection info
  hostAddr: ${["tf-test"].hostname}
  hostPort: ${["tf-test"].port}
  ## ha-replica connection info
  replicaAddr: ${["tf-test"].replicaHostname}
  replicaPort: ${["tf-test"].replicaPort}
  ## pooler connection info
  poolerAddr: ${["tf-test"].poolerHostname}
  poolerPort: ${["tf-test"].poolerPort}
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
import com.pulumi.timescale.Service;
import com.pulumi.timescale.ServiceArgs;
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
        final var tsProjectId = config.get("tsProjectId");
        final var tsAccessKey = config.get("tsAccessKey");
        final var tsSecretKey = config.get("tsSecretKey");
        var tf_test = new Service("tf-test", ServiceArgs.builder()
            .name("tf-test")
            .milliCpu(500)
            .memoryGb(2)
            .regionCode("us-west-2")
            .connectionPoolerEnabled(true)
            .enableHaReplica(true)
            .build());

        ctx.export("hostAddr", tf_test.hostname());
        ctx.export("hostPort", tf_test.port());
        ctx.export("replicaAddr", tf_test.replicaHostname());
        ctx.export("replicaPort", tf_test.replicaPort());
        ctx.export("poolerAddr", tf_test.poolerHostname());
        ctx.export("poolerPort", tf_test.poolerPort());
    }
}
```
{{% /choosable %}}
{{< /chooser >}}

and define the `secret.tfvars` file:

> [!IMPORTANT]
> Replace the values above with the `tsProjectId`, the `tsAccessKey`, and `tsSecretKey`

Now use the `pulumi` cli with the `secrets.tfvars` file, for example:

```shell
pulumi preview --var-file=secrets.tfvars
```
#### VPC Peering

> [!NOTE]The example file creates:
> * A Timescale VPC with name `tf-test` in `us-east-1`
> * An AWS VPC in eu-central-1
> * A Peering connection between them (request and accept automatically)
>
> IMPORTANT: Update region, account ID and CIDRs as needed

Create a `main.tf` configuration file with the following content.

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    aws:region:
        value: eu-central-1
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
const tsProjectId = config.require("tsProjectId");
const tsAccessKey = config.require("tsAccessKey");
const tsSecretKey = config.require("tsSecretKey");
const ts_test = new timescale.Vpcs("ts-test", {
    cidr: "10.0.0.0/24",
    name: "tf-test",
    regionCode: "us-east-1",
});
// Creating a test VPC. Change to your VPC if you already have one in your AWS account.
const main = new aws.index.Vpc("main", {cidrBlock: "11.0.0.0/24"});
// Requester's side of the peering connection (Timescale).
const peer = new timescale.PeeringConnection("peer", {
    peerAccountId: "000000000000",
    peerRegionCode: "eu-central-1",
    peerVpcId: main.id,
    timescaleVpcId: ts_test.vpcsId,
});
// Acceptor's side of the peering connection (AWS).
const peerVpcPeeringConnectionAccepter = new aws.index.VpcPeeringConnectionAccepter("peer", {
    vpcPeeringConnectionId: peer.accepterProvisionedId,
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
    aws:region:
        value: eu-central-1
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
ts_project_id = config.require("tsProjectId")
ts_access_key = config.require("tsAccessKey")
ts_secret_key = config.require("tsSecretKey")
ts_test = timescale.Vpcs("ts-test",
    cidr="10.0.0.0/24",
    name="tf-test",
    region_code="us-east-1")
# Creating a test VPC. Change to your VPC if you already have one in your AWS account.
main = aws.index.Vpc("main", cidr_block=11.0.0.0/24)
# Requester's side of the peering connection (Timescale).
peer = timescale.PeeringConnection("peer",
    peer_account_id="000000000000",
    peer_region_code="eu-central-1",
    peer_vpc_id=main["id"],
    timescale_vpc_id=ts_test.vpcs_id)
# Acceptor's side of the peering connection (AWS).
peer_vpc_peering_connection_accepter = aws.index.VpcPeeringConnectionAccepter("peer",
    vpc_peering_connection_id=peer.accepter_provisioned_id,
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
    aws:region:
        value: eu-central-1
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
    var tsProjectId = config.Require("tsProjectId");
    var tsAccessKey = config.Require("tsAccessKey");
    var tsSecretKey = config.Require("tsSecretKey");
    var ts_test = new Timescale.Vpcs("ts-test", new()
    {
        Cidr = "10.0.0.0/24",
        Name = "tf-test",
        RegionCode = "us-east-1",
    });

    // Creating a test VPC. Change to your VPC if you already have one in your AWS account.
    var main = new Aws.Index.Vpc("main", new()
    {
        CidrBlock = "11.0.0.0/24",
    });

    // Requester's side of the peering connection (Timescale).
    var peer = new Timescale.PeeringConnection("peer", new()
    {
        PeerAccountId = "000000000000",
        PeerRegionCode = "eu-central-1",
        PeerVpcId = main.Id,
        TimescaleVpcId = ts_test.VpcsId,
    });

    // Acceptor's side of the peering connection (AWS).
    var peerVpcPeeringConnectionAccepter = new Aws.Index.VpcPeeringConnectionAccepter("peer", new()
    {
        VpcPeeringConnectionId = peer.AccepterProvisionedId,
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
    aws:region:
        value: eu-central-1
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
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/timescale/v2/timescale"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")
		tsProjectId := cfg.Require("tsProjectId")
		tsAccessKey := cfg.Require("tsAccessKey")
		tsSecretKey := cfg.Require("tsSecretKey")
		ts_test, err := timescale.NewVpcs(ctx, "ts-test", &timescale.VpcsArgs{
			Cidr:       pulumi.String("10.0.0.0/24"),
			Name:       pulumi.String("tf-test"),
			RegionCode: pulumi.String("us-east-1"),
		})
		if err != nil {
			return err
		}
		// Creating a test VPC. Change to your VPC if you already have one in your AWS account.
		main, err := aws.NewVpc(ctx, "main", &aws.VpcArgs{
			CidrBlock: "11.0.0.0/24",
		})
		if err != nil {
			return err
		}
		// Requester's side of the peering connection (Timescale).
		peer, err := timescale.NewPeeringConnection(ctx, "peer", &timescale.PeeringConnectionArgs{
			PeerAccountId:  pulumi.String("000000000000"),
			PeerRegionCode: pulumi.String("eu-central-1"),
			PeerVpcId:      main.Id,
			TimescaleVpcId: ts_test.VpcsId,
		})
		if err != nil {
			return err
		}
		// Acceptor's side of the peering connection (AWS).
		_, err = aws.NewVpcPeeringConnectionAccepter(ctx, "peer", &aws.VpcPeeringConnectionAccepterArgs{
			VpcPeeringConnectionId: peer.AccepterProvisionedId,
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
    aws:region:
        value: eu-central-1
    timescale:accessKey:
        value: 'TODO: var.ts_access_key'
    timescale:projectId:
        value: 'TODO: var.ts_project_id'
    timescale:secretKey:
        value: 'TODO: var.ts_secret_key'

```
```yaml
configuration:
  tsProjectId:
    type: string
  tsAccessKey:
    type: string
  tsSecretKey:
    type: string
resources:
  ts-test:
    type: timescale:Vpcs
    properties:
      cidr: 10.0.0.0/24
      name: tf-test
      regionCode: us-east-1
  # Creating a test VPC. Change to your VPC if you already have one in your AWS account.
  main:
    type: aws:Vpc
    properties:
      cidrBlock: 11.0.0.0/24
  # Requester's side of the peering connection (Timescale).
  peer:
    type: timescale:PeeringConnection
    properties:
      peerAccountId: '000000000000'
      peerRegionCode: eu-central-1
      peerVpcId: ${main.id}
      timescaleVpcId: ${["ts-test"].vpcsId}
  # Acceptor's side of the peering connection (AWS).
  peerVpcPeeringConnectionAccepter:
    type: aws:VpcPeeringConnectionAccepter
    name: peer
    properties:
      vpcPeeringConnectionId: ${peer.accepterProvisionedId}
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
    aws:region:
        value: eu-central-1
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
import com.pulumi.timescale.Vpcs;
import com.pulumi.timescale.VpcsArgs;
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
        final var tsProjectId = config.get("tsProjectId");
        final var tsAccessKey = config.get("tsAccessKey");
        final var tsSecretKey = config.get("tsSecretKey");
        var ts_test = new Vpcs("ts-test", VpcsArgs.builder()
            .cidr("10.0.0.0/24")
            .name("tf-test")
            .regionCode("us-east-1")
            .build());

        // Creating a test VPC. Change to your VPC if you already have one in your AWS account.
        var main = new Vpc("main", VpcArgs.builder()
            .cidrBlock("11.0.0.0/24")
            .build());

        // Requester's side of the peering connection (Timescale).
        var peer = new PeeringConnection("peer", PeeringConnectionArgs.builder()
            .peerAccountId("000000000000")
            .peerRegionCode("eu-central-1")
            .peerVpcId(main.id())
            .timescaleVpcId(ts_test.vpcsId())
            .build());

        // Acceptor's side of the peering connection (AWS).
        var peerVpcPeeringConnectionAccepter = new VpcPeeringConnectionAccepter("peerVpcPeeringConnectionAccepter", VpcPeeringConnectionAccepterArgs.builder()
            .vpcPeeringConnectionId(peer.accepterProvisionedId())
            .autoAccept(true)
            .build(), CustomResourceOptions.builder()
                .dependsOn(peer)
                .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
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
### Regions
Please reference the [docs](https://docs.timescale.com/use-timescale/latest/regions/) for a list of currently supported regions.
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
✅ AWS Transit Gateway peering <br />
✅ Connection pooling <br />
✅ Metric exporters <br />
✅ Log exporters <br />
## Billing
Services are currently billed for hourly usage. If a service is running for less than an hour,
it will still be charged for the full hour of usage.