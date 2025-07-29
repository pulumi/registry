---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/g-core/gcore/0.27.1/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Gcore Provider
meta_desc: Provides an overview on how to configure the Pulumi Gcore provider.
layout: package
---

## Generate Provider

The Gcore provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider g-core/gcore
```
## Overview

Gcore Pulumi Provider allows you to automate the provisioning, management, and testing of your Gcore resources programatically.
## Authentication and Configuration

To start using the Gcore Pulumi Provider you need to configure the provider with the proper credentials.

Configuration for the provider can be derived from multiple sources, which are applied in the following order:

1. Parameters in the provider configuration
2. Environment variables
### Provider Configuration

!> Warning: Hard-coded credentials are not recommended in any Pulumi configuration and risk secret leakage should it ever be committed to a public version control system.

The [permanent API token](https://gcore.com/docs/account-settings/create-use-or-delete-a-permanent-api-token) can be provided by adding a `permanentApiToken` argument to the `gcore` provider configuration.

Example:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    gcore:permanentApiToken:
        value: 251$d3361.............1b35f26d8

```

If needed, the **deprecated** `username` and `password` arguments can be added to the `gcore` provider configuration instead of a permanent API token.

Other settings that can be configured include:

- `apiEndpoint`
- `gcoreCdnApi`
- `gcoreClientId`
- `gcoreCloudApi`
- `gcoreDnsApi`
- `gcorePlatformApi`
- `gcoreStorageApi`
### Environment Variables

The [permanent API token](https://gcore.com/docs/account-settings/create-use-or-delete-a-permanent-api-token) can be provided by setting the `GCORE_PERMANENT_TOKEN` environment variable.

For example:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:

```

```shell
export GCORE_PERMANENT_TOKEN='251$d3361.............1b35f26d8'
pulumi preview
```

If needed, the **deprecated** username / password authentication can be used by setting the `GCORE_USERNAME` and `GCORE_PASSWORD` environment variables.

Other supported environment variables include:

- `GCORE_API_ENDPOINT`
- `GCORE_CDN_API`
- `GCORE_CLIENT_ID`
- `GCORE_CLOUD_API`
- `GCORE_DNS_API`
- `GCORE_PLATFORM_API`
- `GCORE_STORAGE_API`
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    gcore:permanentApiToken:
        value: 251$d3361.............1b35f26d8

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as gcore from "@pulumi/gcore";

const config = new pulumi.Config();
const projectId = config.getNumber("projectId") || 1;
const regionId = config.getNumber("regionId") || 76;
const kp = new gcore.Keypair("kp", {
    projectId: projectId,
    publicKey: "ssh-ed25519 AAAA...CjZ user@example.com",
    sshkeyName: "test_key",
});
const network = new gcore.Network("network", {
    name: "test_network",
    type: "vxlan",
    regionId: regionId,
    projectId: projectId,
});
const subnet = new gcore.Subnet("subnet", {
    name: "test_subnet",
    cidr: "192.168.10.0/24",
    networkId: network.networkId,
    dnsNameservers: [
        "8.8.4.4",
        "1.1.1.1",
    ],
    regionId: regionId,
    projectId: projectId,
});
const subnet2 = new gcore.Subnet("subnet2", {
    name: "test_subnet_2",
    cidr: "192.168.20.0/24",
    networkId: network.networkId,
    dnsNameservers: [
        "8.8.4.4",
        "1.1.1.1",
    ],
    regionId: regionId,
    projectId: projectId,
});
const firstVolume = new gcore.Volume("first_volume", {
    name: "test_boot_volume_1",
    typeName: "ssd_hiiops",
    imageId: "8f0900ba-2002-4f79-b866-390444caa19e",
    size: 10,
    regionId: regionId,
    projectId: projectId,
});
const secondVolume = new gcore.Volume("second_volume", {
    name: "test_boot_volume_2",
    typeName: "ssd_hiiops",
    imageId: "8f0900ba-2002-4f79-b866-390444caa19e",
    size: 10,
    regionId: regionId,
    projectId: projectId,
});
const thirdVolume = new gcore.Volume("third_volume", {
    name: "test_data_volume",
    typeName: "ssd_hiiops",
    size: 6,
    regionId: regionId,
    projectId: projectId,
});
const instance = new gcore.Instancev2("instance", {
    flavorId: "g1-standard-2-4",
    name: "test_instance_1",
    keypairName: kp.sshkeyName,
    volumes: [{
        source: "existing-volume",
        volumeId: firstVolume.volumeId,
        bootIndex: 0,
    }],
    interfaces: [
        {
            type: "subnet",
            networkId: network.networkId,
            subnetId: subnet.subnetId,
            securityGroups: ["11384ae2-2677-439c-8618-f350da006163"],
        },
        {
            type: "subnet",
            networkId: network.networkId,
            subnetId: subnet2.subnetId,
            securityGroups: ["11384ae2-2677-439c-8618-f350da006163"],
        },
    ],
    metadataMap: {
        owner: "username",
    },
    regionId: regionId,
    projectId: projectId,
});
const lb = new gcore.Loadbalancerv2("lb", {
    projectId: projectId,
    regionId: regionId,
    name: "test_loadbalancer",
    flavor: "lb1-1-2",
});
const listener = new gcore.Lblistener("listener", {
    projectId: projectId,
    regionId: regionId,
    name: "test_listener",
    protocol: "HTTP",
    protocolPort: 80,
    loadbalancerId: lb.loadbalancerv2Id,
});
const pl = new gcore.Lbpool("pl", {
    projectId: projectId,
    regionId: regionId,
    name: "test_pool",
    protocol: "HTTP",
    lbAlgorithm: "LEAST_CONNECTIONS",
    loadbalancerId: lb.loadbalancerv2Id,
    listenerId: listener.lblistenerId,
    healthMonitor: {
        type: "PING",
        delay: 60,
        maxRetries: 5,
        timeout: 10,
    },
});
const lbm = new gcore.Lbmember("lbm", {
    projectId: projectId,
    regionId: regionId,
    poolId: pl.lbpoolId,
    instanceId: instance.instancev2Id,
    address: instance.interfaces.apply(interfaces => interfaces[0].ipAddress),
    protocolPort: 8081,
});
const instance2 = new gcore.Instancev2("instance2", {
    flavorId: "g1-standard-2-4",
    name: "test_instance_2",
    keypairName: kp.sshkeyName,
    volumes: [
        {
            source: "existing-volume",
            volumeId: secondVolume.volumeId,
            bootIndex: 0,
        },
        {
            source: "existing-volume",
            volumeId: thirdVolume.volumeId,
            bootIndex: 1,
        },
    ],
    interfaces: [{
        type: "subnet",
        networkId: network.networkId,
        subnetId: subnet.subnetId,
        securityGroups: ["11384ae2-2677-439c-8618-f350da006163"],
    }],
    metadataMap: {
        owner: "username",
    },
    regionId: regionId,
    projectId: projectId,
});
const lbm2 = new gcore.Lbmember("lbm2", {
    projectId: projectId,
    regionId: regionId,
    poolId: pl.lbpoolId,
    instanceId: instance2.instancev2Id,
    address: instance2.interfaces.apply(interfaces => interfaces[0].ipAddress),
    protocolPort: 8081,
    weight: 5,
});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    gcore:permanentApiToken:
        value: 251$d3361.............1b35f26d8

```
```python
import pulumi
import pulumi_gcore as gcore

config = pulumi.Config()
project_id = config.get_float("projectId")
if project_id is None:
    project_id = 1
region_id = config.get_float("regionId")
if region_id is None:
    region_id = 76
kp = gcore.Keypair("kp",
    project_id=project_id,
    public_key="ssh-ed25519 AAAA...CjZ user@example.com",
    sshkey_name="test_key")
network = gcore.Network("network",
    name="test_network",
    type="vxlan",
    region_id=region_id,
    project_id=project_id)
subnet = gcore.Subnet("subnet",
    name="test_subnet",
    cidr="192.168.10.0/24",
    network_id=network.network_id,
    dns_nameservers=[
        "8.8.4.4",
        "1.1.1.1",
    ],
    region_id=region_id,
    project_id=project_id)
subnet2 = gcore.Subnet("subnet2",
    name="test_subnet_2",
    cidr="192.168.20.0/24",
    network_id=network.network_id,
    dns_nameservers=[
        "8.8.4.4",
        "1.1.1.1",
    ],
    region_id=region_id,
    project_id=project_id)
first_volume = gcore.Volume("first_volume",
    name="test_boot_volume_1",
    type_name="ssd_hiiops",
    image_id="8f0900ba-2002-4f79-b866-390444caa19e",
    size=10,
    region_id=region_id,
    project_id=project_id)
second_volume = gcore.Volume("second_volume",
    name="test_boot_volume_2",
    type_name="ssd_hiiops",
    image_id="8f0900ba-2002-4f79-b866-390444caa19e",
    size=10,
    region_id=region_id,
    project_id=project_id)
third_volume = gcore.Volume("third_volume",
    name="test_data_volume",
    type_name="ssd_hiiops",
    size=6,
    region_id=region_id,
    project_id=project_id)
instance = gcore.Instancev2("instance",
    flavor_id="g1-standard-2-4",
    name="test_instance_1",
    keypair_name=kp.sshkey_name,
    volumes=[{
        "source": "existing-volume",
        "volume_id": first_volume.volume_id,
        "boot_index": 0,
    }],
    interfaces=[
        {
            "type": "subnet",
            "network_id": network.network_id,
            "subnet_id": subnet.subnet_id,
            "security_groups": ["11384ae2-2677-439c-8618-f350da006163"],
        },
        {
            "type": "subnet",
            "network_id": network.network_id,
            "subnet_id": subnet2.subnet_id,
            "security_groups": ["11384ae2-2677-439c-8618-f350da006163"],
        },
    ],
    metadata_map={
        "owner": "username",
    },
    region_id=region_id,
    project_id=project_id)
lb = gcore.Loadbalancerv2("lb",
    project_id=project_id,
    region_id=region_id,
    name="test_loadbalancer",
    flavor="lb1-1-2")
listener = gcore.Lblistener("listener",
    project_id=project_id,
    region_id=region_id,
    name="test_listener",
    protocol="HTTP",
    protocol_port=80,
    loadbalancer_id=lb.loadbalancerv2_id)
pl = gcore.Lbpool("pl",
    project_id=project_id,
    region_id=region_id,
    name="test_pool",
    protocol="HTTP",
    lb_algorithm="LEAST_CONNECTIONS",
    loadbalancer_id=lb.loadbalancerv2_id,
    listener_id=listener.lblistener_id,
    health_monitor={
        "type": "PING",
        "delay": 60,
        "max_retries": 5,
        "timeout": 10,
    })
lbm = gcore.Lbmember("lbm",
    project_id=project_id,
    region_id=region_id,
    pool_id=pl.lbpool_id,
    instance_id=instance.instancev2_id,
    address=instance.interfaces[0].ip_address,
    protocol_port=8081)
instance2 = gcore.Instancev2("instance2",
    flavor_id="g1-standard-2-4",
    name="test_instance_2",
    keypair_name=kp.sshkey_name,
    volumes=[
        {
            "source": "existing-volume",
            "volume_id": second_volume.volume_id,
            "boot_index": 0,
        },
        {
            "source": "existing-volume",
            "volume_id": third_volume.volume_id,
            "boot_index": 1,
        },
    ],
    interfaces=[{
        "type": "subnet",
        "network_id": network.network_id,
        "subnet_id": subnet.subnet_id,
        "security_groups": ["11384ae2-2677-439c-8618-f350da006163"],
    }],
    metadata_map={
        "owner": "username",
    },
    region_id=region_id,
    project_id=project_id)
lbm2 = gcore.Lbmember("lbm2",
    project_id=project_id,
    region_id=region_id,
    pool_id=pl.lbpool_id,
    instance_id=instance2.instancev2_id,
    address=instance2.interfaces[0].ip_address,
    protocol_port=8081,
    weight=5)
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    gcore:permanentApiToken:
        value: 251$d3361.............1b35f26d8

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Gcore = Pulumi.Gcore;

return await Deployment.RunAsync(() =>
{
    var config = new Config();
    var projectId = config.GetDouble("projectId") ?? 1;
    var regionId = config.GetDouble("regionId") ?? 76;
    var kp = new Gcore.Keypair("kp", new()
    {
        ProjectId = projectId,
        PublicKey = "ssh-ed25519 AAAA...CjZ user@example.com",
        SshkeyName = "test_key",
    });

    var network = new Gcore.Network("network", new()
    {
        Name = "test_network",
        Type = "vxlan",
        RegionId = regionId,
        ProjectId = projectId,
    });

    var subnet = new Gcore.Subnet("subnet", new()
    {
        Name = "test_subnet",
        Cidr = "192.168.10.0/24",
        NetworkId = network.NetworkId,
        DnsNameservers = new[]
        {
            "8.8.4.4",
            "1.1.1.1",
        },
        RegionId = regionId,
        ProjectId = projectId,
    });

    var subnet2 = new Gcore.Subnet("subnet2", new()
    {
        Name = "test_subnet_2",
        Cidr = "192.168.20.0/24",
        NetworkId = network.NetworkId,
        DnsNameservers = new[]
        {
            "8.8.4.4",
            "1.1.1.1",
        },
        RegionId = regionId,
        ProjectId = projectId,
    });

    var firstVolume = new Gcore.Volume("first_volume", new()
    {
        Name = "test_boot_volume_1",
        TypeName = "ssd_hiiops",
        ImageId = "8f0900ba-2002-4f79-b866-390444caa19e",
        Size = 10,
        RegionId = regionId,
        ProjectId = projectId,
    });

    var secondVolume = new Gcore.Volume("second_volume", new()
    {
        Name = "test_boot_volume_2",
        TypeName = "ssd_hiiops",
        ImageId = "8f0900ba-2002-4f79-b866-390444caa19e",
        Size = 10,
        RegionId = regionId,
        ProjectId = projectId,
    });

    var thirdVolume = new Gcore.Volume("third_volume", new()
    {
        Name = "test_data_volume",
        TypeName = "ssd_hiiops",
        Size = 6,
        RegionId = regionId,
        ProjectId = projectId,
    });

    var instance = new Gcore.Instancev2("instance", new()
    {
        FlavorId = "g1-standard-2-4",
        Name = "test_instance_1",
        KeypairName = kp.SshkeyName,
        Volumes = new[]
        {
            new Gcore.Inputs.Instancev2VolumeArgs
            {
                Source = "existing-volume",
                VolumeId = firstVolume.VolumeId,
                BootIndex = 0,
            },
        },
        Interfaces = new[]
        {
            new Gcore.Inputs.Instancev2InterfaceArgs
            {
                Type = "subnet",
                NetworkId = network.NetworkId,
                SubnetId = subnet.SubnetId,
                SecurityGroups = new[]
                {
                    "11384ae2-2677-439c-8618-f350da006163",
                },
            },
            new Gcore.Inputs.Instancev2InterfaceArgs
            {
                Type = "subnet",
                NetworkId = network.NetworkId,
                SubnetId = subnet2.SubnetId,
                SecurityGroups = new[]
                {
                    "11384ae2-2677-439c-8618-f350da006163",
                },
            },
        },
        MetadataMap =
        {
            { "owner", "username" },
        },
        RegionId = regionId,
        ProjectId = projectId,
    });

    var lb = new Gcore.Loadbalancerv2("lb", new()
    {
        ProjectId = projectId,
        RegionId = regionId,
        Name = "test_loadbalancer",
        Flavor = "lb1-1-2",
    });

    var listener = new Gcore.Lblistener("listener", new()
    {
        ProjectId = projectId,
        RegionId = regionId,
        Name = "test_listener",
        Protocol = "HTTP",
        ProtocolPort = 80,
        LoadbalancerId = lb.Loadbalancerv2Id,
    });

    var pl = new Gcore.Lbpool("pl", new()
    {
        ProjectId = projectId,
        RegionId = regionId,
        Name = "test_pool",
        Protocol = "HTTP",
        LbAlgorithm = "LEAST_CONNECTIONS",
        LoadbalancerId = lb.Loadbalancerv2Id,
        ListenerId = listener.LblistenerId,
        HealthMonitor = new Gcore.Inputs.LbpoolHealthMonitorArgs
        {
            Type = "PING",
            Delay = 60,
            MaxRetries = 5,
            Timeout = 10,
        },
    });

    var lbm = new Gcore.Lbmember("lbm", new()
    {
        ProjectId = projectId,
        RegionId = regionId,
        PoolId = pl.LbpoolId,
        InstanceId = instance.Instancev2Id,
        Address = instance.Interfaces.Apply(interfaces => interfaces[0].IpAddress),
        ProtocolPort = 8081,
    });

    var instance2 = new Gcore.Instancev2("instance2", new()
    {
        FlavorId = "g1-standard-2-4",
        Name = "test_instance_2",
        KeypairName = kp.SshkeyName,
        Volumes = new[]
        {
            new Gcore.Inputs.Instancev2VolumeArgs
            {
                Source = "existing-volume",
                VolumeId = secondVolume.VolumeId,
                BootIndex = 0,
            },
            new Gcore.Inputs.Instancev2VolumeArgs
            {
                Source = "existing-volume",
                VolumeId = thirdVolume.VolumeId,
                BootIndex = 1,
            },
        },
        Interfaces = new[]
        {
            new Gcore.Inputs.Instancev2InterfaceArgs
            {
                Type = "subnet",
                NetworkId = network.NetworkId,
                SubnetId = subnet.SubnetId,
                SecurityGroups = new[]
                {
                    "11384ae2-2677-439c-8618-f350da006163",
                },
            },
        },
        MetadataMap =
        {
            { "owner", "username" },
        },
        RegionId = regionId,
        ProjectId = projectId,
    });

    var lbm2 = new Gcore.Lbmember("lbm2", new()
    {
        ProjectId = projectId,
        RegionId = regionId,
        PoolId = pl.LbpoolId,
        InstanceId = instance2.Instancev2Id,
        Address = instance2.Interfaces.Apply(interfaces => interfaces[0].IpAddress),
        ProtocolPort = 8081,
        Weight = 5,
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
    gcore:permanentApiToken:
        value: 251$d3361.............1b35f26d8

```
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/gcore/gcore"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")
		projectId := float64(1)
		if param := cfg.GetFloat64("projectId"); param != 0 {
			projectId = param
		}
		regionId := float64(76)
		if param := cfg.GetFloat64("regionId"); param != 0 {
			regionId = param
		}
		kp, err := gcore.NewKeypair(ctx, "kp", &gcore.KeypairArgs{
			ProjectId:  pulumi.Float64(projectId),
			PublicKey:  pulumi.String("ssh-ed25519 AAAA...CjZ user@example.com"),
			SshkeyName: pulumi.String("test_key"),
		})
		if err != nil {
			return err
		}
		network, err := gcore.NewNetwork(ctx, "network", &gcore.NetworkArgs{
			Name:      pulumi.String("test_network"),
			Type:      pulumi.String("vxlan"),
			RegionId:  pulumi.Float64(regionId),
			ProjectId: pulumi.Float64(projectId),
		})
		if err != nil {
			return err
		}
		subnet, err := gcore.NewSubnet(ctx, "subnet", &gcore.SubnetArgs{
			Name:      pulumi.String("test_subnet"),
			Cidr:      pulumi.String("192.168.10.0/24"),
			NetworkId: network.NetworkId,
			DnsNameservers: pulumi.StringArray{
				pulumi.String("8.8.4.4"),
				pulumi.String("1.1.1.1"),
			},
			RegionId:  pulumi.Float64(regionId),
			ProjectId: pulumi.Float64(projectId),
		})
		if err != nil {
			return err
		}
		subnet2, err := gcore.NewSubnet(ctx, "subnet2", &gcore.SubnetArgs{
			Name:      pulumi.String("test_subnet_2"),
			Cidr:      pulumi.String("192.168.20.0/24"),
			NetworkId: network.NetworkId,
			DnsNameservers: pulumi.StringArray{
				pulumi.String("8.8.4.4"),
				pulumi.String("1.1.1.1"),
			},
			RegionId:  pulumi.Float64(regionId),
			ProjectId: pulumi.Float64(projectId),
		})
		if err != nil {
			return err
		}
		firstVolume, err := gcore.NewVolume(ctx, "first_volume", &gcore.VolumeArgs{
			Name:      pulumi.String("test_boot_volume_1"),
			TypeName:  pulumi.String("ssd_hiiops"),
			ImageId:   pulumi.String("8f0900ba-2002-4f79-b866-390444caa19e"),
			Size:      pulumi.Float64(10),
			RegionId:  pulumi.Float64(regionId),
			ProjectId: pulumi.Float64(projectId),
		})
		if err != nil {
			return err
		}
		secondVolume, err := gcore.NewVolume(ctx, "second_volume", &gcore.VolumeArgs{
			Name:      pulumi.String("test_boot_volume_2"),
			TypeName:  pulumi.String("ssd_hiiops"),
			ImageId:   pulumi.String("8f0900ba-2002-4f79-b866-390444caa19e"),
			Size:      pulumi.Float64(10),
			RegionId:  pulumi.Float64(regionId),
			ProjectId: pulumi.Float64(projectId),
		})
		if err != nil {
			return err
		}
		thirdVolume, err := gcore.NewVolume(ctx, "third_volume", &gcore.VolumeArgs{
			Name:      pulumi.String("test_data_volume"),
			TypeName:  pulumi.String("ssd_hiiops"),
			Size:      pulumi.Float64(6),
			RegionId:  pulumi.Float64(regionId),
			ProjectId: pulumi.Float64(projectId),
		})
		if err != nil {
			return err
		}
		instance, err := gcore.NewInstancev2(ctx, "instance", &gcore.Instancev2Args{
			FlavorId:    pulumi.String("g1-standard-2-4"),
			Name:        pulumi.String("test_instance_1"),
			KeypairName: kp.SshkeyName,
			Volumes: gcore.Instancev2VolumeArray{
				&gcore.Instancev2VolumeArgs{
					Source:    "existing-volume",
					VolumeId:  firstVolume.VolumeId,
					BootIndex: pulumi.Float64(0),
				},
			},
			Interfaces: gcore.Instancev2InterfaceArray{
				&gcore.Instancev2InterfaceArgs{
					Type:      pulumi.String("subnet"),
					NetworkId: network.NetworkId,
					SubnetId:  subnet.SubnetId,
					SecurityGroups: pulumi.StringArray{
						pulumi.String("11384ae2-2677-439c-8618-f350da006163"),
					},
				},
				&gcore.Instancev2InterfaceArgs{
					Type:      pulumi.String("subnet"),
					NetworkId: network.NetworkId,
					SubnetId:  subnet2.SubnetId,
					SecurityGroups: pulumi.StringArray{
						pulumi.String("11384ae2-2677-439c-8618-f350da006163"),
					},
				},
			},
			MetadataMap: pulumi.StringMap{
				"owner": pulumi.String("username"),
			},
			RegionId:  pulumi.Float64(regionId),
			ProjectId: pulumi.Float64(projectId),
		})
		if err != nil {
			return err
		}
		lb, err := gcore.NewLoadbalancerv2(ctx, "lb", &gcore.Loadbalancerv2Args{
			ProjectId: pulumi.Float64(projectId),
			RegionId:  pulumi.Float64(regionId),
			Name:      pulumi.String("test_loadbalancer"),
			Flavor:    pulumi.String("lb1-1-2"),
		})
		if err != nil {
			return err
		}
		listener, err := gcore.NewLblistener(ctx, "listener", &gcore.LblistenerArgs{
			ProjectId:      pulumi.Float64(projectId),
			RegionId:       pulumi.Float64(regionId),
			Name:           pulumi.String("test_listener"),
			Protocol:       pulumi.String("HTTP"),
			ProtocolPort:   pulumi.Float64(80),
			LoadbalancerId: lb.Loadbalancerv2Id,
		})
		if err != nil {
			return err
		}
		pl, err := gcore.NewLbpool(ctx, "pl", &gcore.LbpoolArgs{
			ProjectId:      pulumi.Float64(projectId),
			RegionId:       pulumi.Float64(regionId),
			Name:           pulumi.String("test_pool"),
			Protocol:       pulumi.String("HTTP"),
			LbAlgorithm:    pulumi.String("LEAST_CONNECTIONS"),
			LoadbalancerId: lb.Loadbalancerv2Id,
			ListenerId:     listener.LblistenerId,
			HealthMonitor: &gcore.LbpoolHealthMonitorArgs{
				Type:       pulumi.String("PING"),
				Delay:      pulumi.Float64(60),
				MaxRetries: pulumi.Float64(5),
				Timeout:    pulumi.Float64(10),
			},
		})
		if err != nil {
			return err
		}
		_, err = gcore.NewLbmember(ctx, "lbm", &gcore.LbmemberArgs{
			ProjectId:  pulumi.Float64(projectId),
			RegionId:   pulumi.Float64(regionId),
			PoolId:     pl.LbpoolId,
			InstanceId: instance.Instancev2Id,
			Address: pulumi.String(instance.Interfaces.ApplyT(func(interfaces []gcore.Instancev2Interface) (*string, error) {
				return &interfaces[0].IpAddress, nil
			}).(pulumi.StringPtrOutput)),
			ProtocolPort: pulumi.Float64(8081),
		})
		if err != nil {
			return err
		}
		instance2, err := gcore.NewInstancev2(ctx, "instance2", &gcore.Instancev2Args{
			FlavorId:    pulumi.String("g1-standard-2-4"),
			Name:        pulumi.String("test_instance_2"),
			KeypairName: kp.SshkeyName,
			Volumes: gcore.Instancev2VolumeArray{
				&gcore.Instancev2VolumeArgs{
					Source:    "existing-volume",
					VolumeId:  secondVolume.VolumeId,
					BootIndex: pulumi.Float64(0),
				},
				&gcore.Instancev2VolumeArgs{
					Source:    "existing-volume",
					VolumeId:  thirdVolume.VolumeId,
					BootIndex: pulumi.Float64(1),
				},
			},
			Interfaces: gcore.Instancev2InterfaceArray{
				&gcore.Instancev2InterfaceArgs{
					Type:      pulumi.String("subnet"),
					NetworkId: network.NetworkId,
					SubnetId:  subnet.SubnetId,
					SecurityGroups: pulumi.StringArray{
						pulumi.String("11384ae2-2677-439c-8618-f350da006163"),
					},
				},
			},
			MetadataMap: pulumi.StringMap{
				"owner": pulumi.String("username"),
			},
			RegionId:  pulumi.Float64(regionId),
			ProjectId: pulumi.Float64(projectId),
		})
		if err != nil {
			return err
		}
		_, err = gcore.NewLbmember(ctx, "lbm2", &gcore.LbmemberArgs{
			ProjectId:  pulumi.Float64(projectId),
			RegionId:   pulumi.Float64(regionId),
			PoolId:     pl.LbpoolId,
			InstanceId: instance2.Instancev2Id,
			Address: pulumi.String(instance2.Interfaces.ApplyT(func(interfaces []gcore.Instancev2Interface) (*string, error) {
				return &interfaces[0].IpAddress, nil
			}).(pulumi.StringPtrOutput)),
			ProtocolPort: pulumi.Float64(8081),
			Weight:       pulumi.Float64(5),
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
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
config:
    gcore:permanentApiToken:
        value: 251$d3361.............1b35f26d8

```
```yaml
configuration:
  projectId:
    type: number
    default: 1
  regionId:
    type: number
    default: 76
resources:
  kp:
    type: gcore:Keypair
    properties:
      projectId: ${projectId}
      publicKey: ssh-ed25519 AAAA...CjZ user@example.com
      sshkeyName: test_key
  network:
    type: gcore:Network
    properties:
      name: test_network
      type: vxlan
      regionId: ${regionId}
      projectId: ${projectId}
  subnet:
    type: gcore:Subnet
    properties:
      name: test_subnet
      cidr: 192.168.10.0/24
      networkId: ${network.networkId}
      dnsNameservers:
        - 8.8.4.4
        - 1.1.1.1
      regionId: ${regionId}
      projectId: ${projectId}
  subnet2:
    type: gcore:Subnet
    properties:
      name: test_subnet_2
      cidr: 192.168.20.0/24
      networkId: ${network.networkId}
      dnsNameservers:
        - 8.8.4.4
        - 1.1.1.1
      regionId: ${regionId}
      projectId: ${projectId}
  firstVolume:
    type: gcore:Volume
    name: first_volume
    properties:
      name: test_boot_volume_1
      typeName: ssd_hiiops
      imageId: 8f0900ba-2002-4f79-b866-390444caa19e
      size: 10
      regionId: ${regionId}
      projectId: ${projectId}
  secondVolume:
    type: gcore:Volume
    name: second_volume
    properties:
      name: test_boot_volume_2
      typeName: ssd_hiiops
      imageId: 8f0900ba-2002-4f79-b866-390444caa19e
      size: 10
      regionId: ${regionId}
      projectId: ${projectId}
  thirdVolume:
    type: gcore:Volume
    name: third_volume
    properties:
      name: test_data_volume
      typeName: ssd_hiiops
      size: 6
      regionId: ${regionId}
      projectId: ${projectId}
  instance:
    type: gcore:Instancev2
    properties:
      flavorId: g1-standard-2-4
      name: test_instance_1
      keypairName: ${kp.sshkeyName}
      volumes:
        - source: existing-volume
          volumeId: ${firstVolume.volumeId}
          bootIndex: 0
      interfaces:
        - type: subnet
          networkId: ${network.networkId}
          subnetId: ${subnet.subnetId}
          securityGroups:
            - 11384ae2-2677-439c-8618-f350da006163
        - type: subnet
          networkId: ${network.networkId}
          subnetId: ${subnet2.subnetId}
          securityGroups:
            - 11384ae2-2677-439c-8618-f350da006163
      metadataMap:
        owner: username
      regionId: ${regionId}
      projectId: ${projectId}
  lb:
    type: gcore:Loadbalancerv2
    properties:
      projectId: ${projectId}
      regionId: ${regionId}
      name: test_loadbalancer
      flavor: lb1-1-2
  listener:
    type: gcore:Lblistener
    properties:
      projectId: ${projectId}
      regionId: ${regionId}
      name: test_listener
      protocol: HTTP
      protocolPort: 80
      loadbalancerId: ${lb.loadbalancerv2Id}
  pl:
    type: gcore:Lbpool
    properties:
      projectId: ${projectId}
      regionId: ${regionId}
      name: test_pool
      protocol: HTTP
      lbAlgorithm: LEAST_CONNECTIONS
      loadbalancerId: ${lb.loadbalancerv2Id}
      listenerId: ${listener.lblistenerId}
      healthMonitor:
        type: PING
        delay: 60
        maxRetries: 5
        timeout: 10
  lbm:
    type: gcore:Lbmember
    properties:
      projectId: ${projectId}
      regionId: ${regionId}
      poolId: ${pl.lbpoolId}
      instanceId: ${instance.instancev2Id}
      address: ${instance.interfaces[0].ipAddress}
      protocolPort: 8081
  instance2:
    type: gcore:Instancev2
    properties:
      flavorId: g1-standard-2-4
      name: test_instance_2
      keypairName: ${kp.sshkeyName}
      volumes:
        - source: existing-volume
          volumeId: ${secondVolume.volumeId}
          bootIndex: 0
        - source: existing-volume
          volumeId: ${thirdVolume.volumeId}
          bootIndex: 1
      interfaces:
        - type: subnet
          networkId: ${network.networkId}
          subnetId: ${subnet.subnetId}
          securityGroups:
            - 11384ae2-2677-439c-8618-f350da006163
      metadataMap:
        owner: username
      regionId: ${regionId}
      projectId: ${projectId}
  lbm2:
    type: gcore:Lbmember
    properties:
      projectId: ${projectId}
      regionId: ${regionId}
      poolId: ${pl.lbpoolId}
      instanceId: ${instance2.instancev2Id}
      address: ${instance2.interfaces[0].ipAddress}
      protocolPort: 8081
      weight: 5
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    gcore:permanentApiToken:
        value: 251$d3361.............1b35f26d8

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.gcore.Keypair;
import com.pulumi.gcore.KeypairArgs;
import com.pulumi.gcore.Network;
import com.pulumi.gcore.NetworkArgs;
import com.pulumi.gcore.Subnet;
import com.pulumi.gcore.SubnetArgs;
import com.pulumi.gcore.Volume;
import com.pulumi.gcore.VolumeArgs;
import com.pulumi.gcore.Instancev2;
import com.pulumi.gcore.Instancev2Args;
import com.pulumi.gcore.inputs.Instancev2VolumeArgs;
import com.pulumi.gcore.inputs.Instancev2InterfaceArgs;
import com.pulumi.gcore.Loadbalancerv2;
import com.pulumi.gcore.Loadbalancerv2Args;
import com.pulumi.gcore.Lblistener;
import com.pulumi.gcore.LblistenerArgs;
import com.pulumi.gcore.Lbpool;
import com.pulumi.gcore.LbpoolArgs;
import com.pulumi.gcore.inputs.LbpoolHealthMonitorArgs;
import com.pulumi.gcore.Lbmember;
import com.pulumi.gcore.LbmemberArgs;
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
        final var projectId = config.get("projectId").orElse(1);
        final var regionId = config.get("regionId").orElse(76);
        var kp = new Keypair("kp", KeypairArgs.builder()
            .projectId(projectId)
            .publicKey("ssh-ed25519 AAAA...CjZ user@example.com")
            .sshkeyName("test_key")
            .build());

        var network = new Network("network", NetworkArgs.builder()
            .name("test_network")
            .type("vxlan")
            .regionId(regionId)
            .projectId(projectId)
            .build());

        var subnet = new Subnet("subnet", SubnetArgs.builder()
            .name("test_subnet")
            .cidr("192.168.10.0/24")
            .networkId(network.networkId())
            .dnsNameservers(
                "8.8.4.4",
                "1.1.1.1")
            .regionId(regionId)
            .projectId(projectId)
            .build());

        var subnet2 = new Subnet("subnet2", SubnetArgs.builder()
            .name("test_subnet_2")
            .cidr("192.168.20.0/24")
            .networkId(network.networkId())
            .dnsNameservers(
                "8.8.4.4",
                "1.1.1.1")
            .regionId(regionId)
            .projectId(projectId)
            .build());

        var firstVolume = new Volume("firstVolume", VolumeArgs.builder()
            .name("test_boot_volume_1")
            .typeName("ssd_hiiops")
            .imageId("8f0900ba-2002-4f79-b866-390444caa19e")
            .size(10)
            .regionId(regionId)
            .projectId(projectId)
            .build());

        var secondVolume = new Volume("secondVolume", VolumeArgs.builder()
            .name("test_boot_volume_2")
            .typeName("ssd_hiiops")
            .imageId("8f0900ba-2002-4f79-b866-390444caa19e")
            .size(10)
            .regionId(regionId)
            .projectId(projectId)
            .build());

        var thirdVolume = new Volume("thirdVolume", VolumeArgs.builder()
            .name("test_data_volume")
            .typeName("ssd_hiiops")
            .size(6)
            .regionId(regionId)
            .projectId(projectId)
            .build());

        var instance = new Instancev2("instance", Instancev2Args.builder()
            .flavorId("g1-standard-2-4")
            .name("test_instance_1")
            .keypairName(kp.sshkeyName())
            .volumes(Instancev2VolumeArgs.builder()
                .source("existing-volume")
                .volumeId(firstVolume.volumeId())
                .bootIndex(0)
                .build())
            .interfaces(
                Instancev2InterfaceArgs.builder()
                    .type("subnet")
                    .networkId(network.networkId())
                    .subnetId(subnet.subnetId())
                    .securityGroups("11384ae2-2677-439c-8618-f350da006163")
                    .build(),
                Instancev2InterfaceArgs.builder()
                    .type("subnet")
                    .networkId(network.networkId())
                    .subnetId(subnet2.subnetId())
                    .securityGroups("11384ae2-2677-439c-8618-f350da006163")
                    .build())
            .metadataMap(Map.of("owner", "username"))
            .regionId(regionId)
            .projectId(projectId)
            .build());

        var lb = new Loadbalancerv2("lb", Loadbalancerv2Args.builder()
            .projectId(projectId)
            .regionId(regionId)
            .name("test_loadbalancer")
            .flavor("lb1-1-2")
            .build());

        var listener = new Lblistener("listener", LblistenerArgs.builder()
            .projectId(projectId)
            .regionId(regionId)
            .name("test_listener")
            .protocol("HTTP")
            .protocolPort(80)
            .loadbalancerId(lb.loadbalancerv2Id())
            .build());

        var pl = new Lbpool("pl", LbpoolArgs.builder()
            .projectId(projectId)
            .regionId(regionId)
            .name("test_pool")
            .protocol("HTTP")
            .lbAlgorithm("LEAST_CONNECTIONS")
            .loadbalancerId(lb.loadbalancerv2Id())
            .listenerId(listener.lblistenerId())
            .healthMonitor(LbpoolHealthMonitorArgs.builder()
                .type("PING")
                .delay(60)
                .maxRetries(5)
                .timeout(10)
                .build())
            .build());

        var lbm = new Lbmember("lbm", LbmemberArgs.builder()
            .projectId(projectId)
            .regionId(regionId)
            .poolId(pl.lbpoolId())
            .instanceId(instance.instancev2Id())
            .address(instance.interfaces().applyValue(interfaces -> interfaces[0].ipAddress()))
            .protocolPort(8081)
            .build());

        var instance2 = new Instancev2("instance2", Instancev2Args.builder()
            .flavorId("g1-standard-2-4")
            .name("test_instance_2")
            .keypairName(kp.sshkeyName())
            .volumes(
                Instancev2VolumeArgs.builder()
                    .source("existing-volume")
                    .volumeId(secondVolume.volumeId())
                    .bootIndex(0)
                    .build(),
                Instancev2VolumeArgs.builder()
                    .source("existing-volume")
                    .volumeId(thirdVolume.volumeId())
                    .bootIndex(1)
                    .build())
            .interfaces(Instancev2InterfaceArgs.builder()
                .type("subnet")
                .networkId(network.networkId())
                .subnetId(subnet.subnetId())
                .securityGroups("11384ae2-2677-439c-8618-f350da006163")
                .build())
            .metadataMap(Map.of("owner", "username"))
            .regionId(regionId)
            .projectId(projectId)
            .build());

        var lbm2 = new Lbmember("lbm2", LbmemberArgs.builder()
            .projectId(projectId)
            .regionId(regionId)
            .poolId(pl.lbpoolId())
            .instanceId(instance2.instancev2Id())
            .address(instance2.interfaces().applyValue(interfaces -> interfaces[0].ipAddress()))
            .protocolPort(8081)
            .weight(5)
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

- `apiEndpoint` (String) A single API endpoint for all products. Will be used when specific product API url is not defined. Can also be set with the GCORE_API_ENDPOINT environment variable.
- `gcoreApi` (String, Deprecated) Region API.
- `gcoreCdnApi` (String) CDN API (define only if you want to override CDN API endpoint). Can also be set with the GCORE_CDN_API environment variable.
- `gcoreClientId` (String) Client ID. Can also be set with the GCORE_CLIENT_ID environment variable.
- `gcoreCloudApi` (String) Region API (define only if you want to override Region API endpoint). Can also be set with the GCORE_CLOUD_API environment variable.
- `gcoreDnsApi` (String) DNS API (define only if you want to override DNS API endpoint). Can also be set with the GCORE_DNS_API environment variable.
- `gcoreFastedgeApi` (String) FastEdge API (define only if you want to override FastEdge API endpoint). Can also be set with the GCORE_FASTEDGE_API environment variable.
- `gcorePlatform` (String, Deprecated) Platform URL is used for generate JWT.
- `gcorePlatformApi` (String) Platform URL is used for generate JWT (define only if you want to override Platform API endpoint). Can also be set with the GCORE_PLATFORM_API environment variable.
- `gcoreStorageApi` (String) Storage API (define only if you want to override Storage API endpoint). Can also be set with the GCORE_STORAGE_API environment variable.
- `gcoreWaapApi` (String) WAAP API (define only if you want to override WAAP API endpoint). Can also be set with the GCORE_WAAP_API environment variable.
- `ignoreCredsAuthError` (Boolean, Deprecated) Should be set to true when you are gonna to use storage resource with permanent API-token only.
- `password` (String, Deprecated) Gcore account password. Can also be set with the GCORE_PASSWORD environment variable.
- `permanentApiToken` (String, Sensitive) A permanent [API-token](https://gcore.com/docs/account-settings/create-use-or-delete-a-permanent-api-token). Can also be set with the GCORE_PERMANENT_TOKEN environment variable.
- `userName` (String, Deprecated) Gcore account username. Can also be set with the GCORE_USERNAME environment variable.