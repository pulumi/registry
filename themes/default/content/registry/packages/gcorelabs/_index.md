---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/g-core/gcorelabs/0.3.63/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Gcorelabs Provider
meta_desc: Provides an overview on how to configure the Pulumi Gcorelabs provider.
layout: package
---

## Generate Provider

The Gcorelabs provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider g-core/gcorelabs
```
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
import * as gcorelabs from "@pulumi/gcorelabs";

const kp = new gcorelabs.Keypair("kp", {
    projectId: 1,
    publicKey: "your oub key",
    sshkeyName: "testkey",
});
const network = new gcorelabs.Network("network", {
    name: "network_example",
    mtu: 1450,
    type: "vxlan",
    regionId: 1,
    projectId: 1,
});
const subnet = new gcorelabs.Subnet("subnet", {
    name: "subnet_example",
    cidr: "192.168.10.0/24",
    networkId: network.networkId,
    dnsNameservers: [
        "8.8.4.4",
        "1.1.1.1",
    ],
    hostRoutes: [{
        destination: "10.0.3.0/24",
        nexthop: "10.0.0.13",
    }],
    gatewayIp: "192.168.10.1",
    regionId: 1,
    projectId: 1,
});
const subnet2 = new gcorelabs.Subnet("subnet2", {
    name: "subnet2_example",
    cidr: "192.168.20.0/24",
    networkId: network.networkId,
    dnsNameservers: [
        "8.8.4.4",
        "1.1.1.1",
    ],
    hostRoutes: [{
        destination: "10.0.3.0/24",
        nexthop: "10.0.0.13",
    }],
    gatewayIp: "192.168.20.1",
    regionId: 1,
    projectId: 1,
});
const firstVolume = new gcorelabs.Volume("first_volume", {
    name: "boot volume",
    typeName: "ssd_hiiops",
    size: 6,
    imageId: "f4ce3d30-e29c-4cfd-811f-46f383b6081f",
    regionId: 1,
    projectId: 1,
});
const secondVolume = new gcorelabs.Volume("second_volume", {
    name: "second volume",
    typeName: "ssd_hiiops",
    imageId: "f4ce3d30-e29c-4cfd-811f-46f383b6081f",
    size: 6,
    regionId: 1,
    projectId: 1,
});
const thirdVolume = new gcorelabs.Volume("third_volume", {
    name: "third volume",
    typeName: "ssd_hiiops",
    size: 6,
    regionId: 1,
    projectId: 1,
});
const instance = new gcorelabs.Instance("instance", {
    flavorId: "g1-standard-2-4",
    name: "test",
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
        },
        {
            type: "subnet",
            networkId: network.networkId,
            subnetId: subnet2.subnetId,
        },
    ],
    securityGroups: [{
        id: "66988147-f1b9-43b2-aaef-dee6d009b5b7",
        name: "default",
    }],
    metadatas: [{
        key: "some_key",
        value: "some_data",
    }],
    configurations: [{
        key: "some_key",
        value: "some_data",
    }],
    regionId: 1,
    projectId: 1,
});
const lb = new gcorelabs.Loadbalancer("lb", {
    projectId: 1,
    regionId: 1,
    name: "test1",
    flavor: "lb1-1-2",
    listener: {
        name: "test",
        protocol: "HTTP",
        protocolPort: 80,
    },
});
const pl = new gcorelabs.Lbpool("pl", {
    projectId: 1,
    regionId: 1,
    name: "test_pool1",
    protocol: "HTTP",
    lbAlgorithm: "LEAST_CONNECTIONS",
    loadbalancerId: lb.loadbalancerId,
    listenerId: lb.listener.apply(listener => listener.id),
    healthMonitor: {
        type: "PING",
        delay: 60,
        maxRetries: 5,
        timeout: 10,
    },
    sessionPersistence: {
        type: "APP_COOKIE",
        cookieName: "test_new_cookie",
    },
});
const lbm = new gcorelabs.Lbmember("lbm", {
    projectId: 1,
    regionId: 1,
    poolId: pl.lbpoolId,
    instanceId: instance.instanceId,
    address: instance.interfaces.apply(interfaces => interfaces[0].ipAddress),
    protocolPort: 8081,
    weight: 5,
});
const instance2 = new gcorelabs.Instance("instance2", {
    flavorId: "g1-standard-2-4",
    name: "test2",
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
    }],
    securityGroups: [{
        id: "66988147-f1b9-43b2-aaef-dee6d009b5b7",
        name: "default",
    }],
    metadatas: [{
        key: "some_key",
        value: "some_data",
    }],
    configurations: [{
        key: "some_key",
        value: "some_data",
    }],
    regionId: 1,
    projectId: 1,
});
const lbm2 = new gcorelabs.Lbmember("lbm2", {
    projectId: 1,
    regionId: 1,
    poolId: pl.lbpoolId,
    instanceId: instance2.instanceId,
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
import pulumi_gcorelabs as gcorelabs

kp = gcorelabs.Keypair("kp",
    project_id=1,
    public_key="your oub key",
    sshkey_name="testkey")
network = gcorelabs.Network("network",
    name="network_example",
    mtu=1450,
    type="vxlan",
    region_id=1,
    project_id=1)
subnet = gcorelabs.Subnet("subnet",
    name="subnet_example",
    cidr="192.168.10.0/24",
    network_id=network.network_id,
    dns_nameservers=[
        "8.8.4.4",
        "1.1.1.1",
    ],
    host_routes=[{
        "destination": "10.0.3.0/24",
        "nexthop": "10.0.0.13",
    }],
    gateway_ip="192.168.10.1",
    region_id=1,
    project_id=1)
subnet2 = gcorelabs.Subnet("subnet2",
    name="subnet2_example",
    cidr="192.168.20.0/24",
    network_id=network.network_id,
    dns_nameservers=[
        "8.8.4.4",
        "1.1.1.1",
    ],
    host_routes=[{
        "destination": "10.0.3.0/24",
        "nexthop": "10.0.0.13",
    }],
    gateway_ip="192.168.20.1",
    region_id=1,
    project_id=1)
first_volume = gcorelabs.Volume("first_volume",
    name="boot volume",
    type_name="ssd_hiiops",
    size=6,
    image_id="f4ce3d30-e29c-4cfd-811f-46f383b6081f",
    region_id=1,
    project_id=1)
second_volume = gcorelabs.Volume("second_volume",
    name="second volume",
    type_name="ssd_hiiops",
    image_id="f4ce3d30-e29c-4cfd-811f-46f383b6081f",
    size=6,
    region_id=1,
    project_id=1)
third_volume = gcorelabs.Volume("third_volume",
    name="third volume",
    type_name="ssd_hiiops",
    size=6,
    region_id=1,
    project_id=1)
instance = gcorelabs.Instance("instance",
    flavor_id="g1-standard-2-4",
    name="test",
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
        },
        {
            "type": "subnet",
            "network_id": network.network_id,
            "subnet_id": subnet2.subnet_id,
        },
    ],
    security_groups=[{
        "id": "66988147-f1b9-43b2-aaef-dee6d009b5b7",
        "name": "default",
    }],
    metadatas=[{
        "key": "some_key",
        "value": "some_data",
    }],
    configurations=[{
        "key": "some_key",
        "value": "some_data",
    }],
    region_id=1,
    project_id=1)
lb = gcorelabs.Loadbalancer("lb",
    project_id=1,
    region_id=1,
    name="test1",
    flavor="lb1-1-2",
    listener={
        "name": "test",
        "protocol": "HTTP",
        "protocol_port": 80,
    })
pl = gcorelabs.Lbpool("pl",
    project_id=1,
    region_id=1,
    name="test_pool1",
    protocol="HTTP",
    lb_algorithm="LEAST_CONNECTIONS",
    loadbalancer_id=lb.loadbalancer_id,
    listener_id=lb.listener.id,
    health_monitor={
        "type": "PING",
        "delay": 60,
        "max_retries": 5,
        "timeout": 10,
    },
    session_persistence={
        "type": "APP_COOKIE",
        "cookie_name": "test_new_cookie",
    })
lbm = gcorelabs.Lbmember("lbm",
    project_id=1,
    region_id=1,
    pool_id=pl.lbpool_id,
    instance_id=instance.instance_id,
    address=instance.interfaces[0].ip_address,
    protocol_port=8081,
    weight=5)
instance2 = gcorelabs.Instance("instance2",
    flavor_id="g1-standard-2-4",
    name="test2",
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
    }],
    security_groups=[{
        "id": "66988147-f1b9-43b2-aaef-dee6d009b5b7",
        "name": "default",
    }],
    metadatas=[{
        "key": "some_key",
        "value": "some_data",
    }],
    configurations=[{
        "key": "some_key",
        "value": "some_data",
    }],
    region_id=1,
    project_id=1)
lbm2 = gcorelabs.Lbmember("lbm2",
    project_id=1,
    region_id=1,
    pool_id=pl.lbpool_id,
    instance_id=instance2.instance_id,
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
using Gcorelabs = Pulumi.Gcorelabs;

return await Deployment.RunAsync(() =>
{
    var kp = new Gcorelabs.Keypair("kp", new()
    {
        ProjectId = 1,
        PublicKey = "your oub key",
        SshkeyName = "testkey",
    });

    var network = new Gcorelabs.Network("network", new()
    {
        Name = "network_example",
        Mtu = 1450,
        Type = "vxlan",
        RegionId = 1,
        ProjectId = 1,
    });

    var subnet = new Gcorelabs.Subnet("subnet", new()
    {
        Name = "subnet_example",
        Cidr = "192.168.10.0/24",
        NetworkId = network.NetworkId,
        DnsNameservers = new[]
        {
            "8.8.4.4",
            "1.1.1.1",
        },
        HostRoutes = new[]
        {
            new Gcorelabs.Inputs.SubnetHostRouteArgs
            {
                Destination = "10.0.3.0/24",
                Nexthop = "10.0.0.13",
            },
        },
        GatewayIp = "192.168.10.1",
        RegionId = 1,
        ProjectId = 1,
    });

    var subnet2 = new Gcorelabs.Subnet("subnet2", new()
    {
        Name = "subnet2_example",
        Cidr = "192.168.20.0/24",
        NetworkId = network.NetworkId,
        DnsNameservers = new[]
        {
            "8.8.4.4",
            "1.1.1.1",
        },
        HostRoutes = new[]
        {
            new Gcorelabs.Inputs.SubnetHostRouteArgs
            {
                Destination = "10.0.3.0/24",
                Nexthop = "10.0.0.13",
            },
        },
        GatewayIp = "192.168.20.1",
        RegionId = 1,
        ProjectId = 1,
    });

    var firstVolume = new Gcorelabs.Volume("first_volume", new()
    {
        Name = "boot volume",
        TypeName = "ssd_hiiops",
        Size = 6,
        ImageId = "f4ce3d30-e29c-4cfd-811f-46f383b6081f",
        RegionId = 1,
        ProjectId = 1,
    });

    var secondVolume = new Gcorelabs.Volume("second_volume", new()
    {
        Name = "second volume",
        TypeName = "ssd_hiiops",
        ImageId = "f4ce3d30-e29c-4cfd-811f-46f383b6081f",
        Size = 6,
        RegionId = 1,
        ProjectId = 1,
    });

    var thirdVolume = new Gcorelabs.Volume("third_volume", new()
    {
        Name = "third volume",
        TypeName = "ssd_hiiops",
        Size = 6,
        RegionId = 1,
        ProjectId = 1,
    });

    var instance = new Gcorelabs.Instance("instance", new()
    {
        FlavorId = "g1-standard-2-4",
        Name = "test",
        KeypairName = kp.SshkeyName,
        Volumes = new[]
        {
            new Gcorelabs.Inputs.InstanceVolumeArgs
            {
                Source = "existing-volume",
                VolumeId = firstVolume.VolumeId,
                BootIndex = 0,
            },
        },
        Interfaces = new[]
        {
            new Gcorelabs.Inputs.InstanceInterfaceArgs
            {
                Type = "subnet",
                NetworkId = network.NetworkId,
                SubnetId = subnet.SubnetId,
            },
            new Gcorelabs.Inputs.InstanceInterfaceArgs
            {
                Type = "subnet",
                NetworkId = network.NetworkId,
                SubnetId = subnet2.SubnetId,
            },
        },
        SecurityGroups = new[]
        {
            new Gcorelabs.Inputs.InstanceSecurityGroupArgs
            {
                Id = "66988147-f1b9-43b2-aaef-dee6d009b5b7",
                Name = "default",
            },
        },
        Metadatas = new[]
        {
            new Gcorelabs.Inputs.InstanceMetadataArgs
            {
                Key = "some_key",
                Value = "some_data",
            },
        },
        Configurations = new[]
        {
            new Gcorelabs.Inputs.InstanceConfigurationArgs
            {
                Key = "some_key",
                Value = "some_data",
            },
        },
        RegionId = 1,
        ProjectId = 1,
    });

    var lb = new Gcorelabs.Loadbalancer("lb", new()
    {
        ProjectId = 1,
        RegionId = 1,
        Name = "test1",
        Flavor = "lb1-1-2",
        Listener = new Gcorelabs.Inputs.LoadbalancerListenerArgs
        {
            Name = "test",
            Protocol = "HTTP",
            ProtocolPort = 80,
        },
    });

    var pl = new Gcorelabs.Lbpool("pl", new()
    {
        ProjectId = 1,
        RegionId = 1,
        Name = "test_pool1",
        Protocol = "HTTP",
        LbAlgorithm = "LEAST_CONNECTIONS",
        LoadbalancerId = lb.LoadbalancerId,
        ListenerId = lb.Listener.Apply(listener => listener.Id),
        HealthMonitor = new Gcorelabs.Inputs.LbpoolHealthMonitorArgs
        {
            Type = "PING",
            Delay = 60,
            MaxRetries = 5,
            Timeout = 10,
        },
        SessionPersistence = new Gcorelabs.Inputs.LbpoolSessionPersistenceArgs
        {
            Type = "APP_COOKIE",
            CookieName = "test_new_cookie",
        },
    });

    var lbm = new Gcorelabs.Lbmember("lbm", new()
    {
        ProjectId = 1,
        RegionId = 1,
        PoolId = pl.LbpoolId,
        InstanceId = instance.InstanceId,
        Address = instance.Interfaces.Apply(interfaces => interfaces[0].IpAddress),
        ProtocolPort = 8081,
        Weight = 5,
    });

    var instance2 = new Gcorelabs.Instance("instance2", new()
    {
        FlavorId = "g1-standard-2-4",
        Name = "test2",
        KeypairName = kp.SshkeyName,
        Volumes = new[]
        {
            new Gcorelabs.Inputs.InstanceVolumeArgs
            {
                Source = "existing-volume",
                VolumeId = secondVolume.VolumeId,
                BootIndex = 0,
            },
            new Gcorelabs.Inputs.InstanceVolumeArgs
            {
                Source = "existing-volume",
                VolumeId = thirdVolume.VolumeId,
                BootIndex = 1,
            },
        },
        Interfaces = new[]
        {
            new Gcorelabs.Inputs.InstanceInterfaceArgs
            {
                Type = "subnet",
                NetworkId = network.NetworkId,
                SubnetId = subnet.SubnetId,
            },
        },
        SecurityGroups = new[]
        {
            new Gcorelabs.Inputs.InstanceSecurityGroupArgs
            {
                Id = "66988147-f1b9-43b2-aaef-dee6d009b5b7",
                Name = "default",
            },
        },
        Metadatas = new[]
        {
            new Gcorelabs.Inputs.InstanceMetadataArgs
            {
                Key = "some_key",
                Value = "some_data",
            },
        },
        Configurations = new[]
        {
            new Gcorelabs.Inputs.InstanceConfigurationArgs
            {
                Key = "some_key",
                Value = "some_data",
            },
        },
        RegionId = 1,
        ProjectId = 1,
    });

    var lbm2 = new Gcorelabs.Lbmember("lbm2", new()
    {
        ProjectId = 1,
        RegionId = 1,
        PoolId = pl.LbpoolId,
        InstanceId = instance2.InstanceId,
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
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/gcorelabs/gcorelabs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		kp, err := gcorelabs.NewKeypair(ctx, "kp", &gcorelabs.KeypairArgs{
			ProjectId:  pulumi.Float64(1),
			PublicKey:  pulumi.String("your oub key"),
			SshkeyName: pulumi.String("testkey"),
		})
		if err != nil {
			return err
		}
		network, err := gcorelabs.NewNetwork(ctx, "network", &gcorelabs.NetworkArgs{
			Name:      pulumi.String("network_example"),
			Mtu:       pulumi.Float64(1450),
			Type:      pulumi.String("vxlan"),
			RegionId:  pulumi.Float64(1),
			ProjectId: pulumi.Float64(1),
		})
		if err != nil {
			return err
		}
		subnet, err := gcorelabs.NewSubnet(ctx, "subnet", &gcorelabs.SubnetArgs{
			Name:      pulumi.String("subnet_example"),
			Cidr:      pulumi.String("192.168.10.0/24"),
			NetworkId: network.NetworkId,
			DnsNameservers: pulumi.StringArray{
				pulumi.String("8.8.4.4"),
				pulumi.String("1.1.1.1"),
			},
			HostRoutes: gcorelabs.SubnetHostRouteArray{
				&gcorelabs.SubnetHostRouteArgs{
					Destination: pulumi.String("10.0.3.0/24"),
					Nexthop:     pulumi.String("10.0.0.13"),
				},
			},
			GatewayIp: pulumi.String("192.168.10.1"),
			RegionId:  pulumi.Float64(1),
			ProjectId: pulumi.Float64(1),
		})
		if err != nil {
			return err
		}
		subnet2, err := gcorelabs.NewSubnet(ctx, "subnet2", &gcorelabs.SubnetArgs{
			Name:      pulumi.String("subnet2_example"),
			Cidr:      pulumi.String("192.168.20.0/24"),
			NetworkId: network.NetworkId,
			DnsNameservers: pulumi.StringArray{
				pulumi.String("8.8.4.4"),
				pulumi.String("1.1.1.1"),
			},
			HostRoutes: gcorelabs.SubnetHostRouteArray{
				&gcorelabs.SubnetHostRouteArgs{
					Destination: pulumi.String("10.0.3.0/24"),
					Nexthop:     pulumi.String("10.0.0.13"),
				},
			},
			GatewayIp: pulumi.String("192.168.20.1"),
			RegionId:  pulumi.Float64(1),
			ProjectId: pulumi.Float64(1),
		})
		if err != nil {
			return err
		}
		firstVolume, err := gcorelabs.NewVolume(ctx, "first_volume", &gcorelabs.VolumeArgs{
			Name:      pulumi.String("boot volume"),
			TypeName:  pulumi.String("ssd_hiiops"),
			Size:      pulumi.Float64(6),
			ImageId:   pulumi.String("f4ce3d30-e29c-4cfd-811f-46f383b6081f"),
			RegionId:  pulumi.Float64(1),
			ProjectId: pulumi.Float64(1),
		})
		if err != nil {
			return err
		}
		secondVolume, err := gcorelabs.NewVolume(ctx, "second_volume", &gcorelabs.VolumeArgs{
			Name:      pulumi.String("second volume"),
			TypeName:  pulumi.String("ssd_hiiops"),
			ImageId:   pulumi.String("f4ce3d30-e29c-4cfd-811f-46f383b6081f"),
			Size:      pulumi.Float64(6),
			RegionId:  pulumi.Float64(1),
			ProjectId: pulumi.Float64(1),
		})
		if err != nil {
			return err
		}
		thirdVolume, err := gcorelabs.NewVolume(ctx, "third_volume", &gcorelabs.VolumeArgs{
			Name:      pulumi.String("third volume"),
			TypeName:  pulumi.String("ssd_hiiops"),
			Size:      pulumi.Float64(6),
			RegionId:  pulumi.Float64(1),
			ProjectId: pulumi.Float64(1),
		})
		if err != nil {
			return err
		}
		instance, err := gcorelabs.NewInstance(ctx, "instance", &gcorelabs.InstanceArgs{
			FlavorId:    pulumi.String("g1-standard-2-4"),
			Name:        pulumi.String("test"),
			KeypairName: kp.SshkeyName,
			Volumes: gcorelabs.InstanceVolumeArray{
				&gcorelabs.InstanceVolumeArgs{
					Source:    pulumi.String("existing-volume"),
					VolumeId:  firstVolume.VolumeId,
					BootIndex: pulumi.Float64(0),
				},
			},
			Interfaces: gcorelabs.InstanceInterfaceArray{
				&gcorelabs.InstanceInterfaceArgs{
					Type:      pulumi.String("subnet"),
					NetworkId: network.NetworkId,
					SubnetId:  subnet.SubnetId,
				},
				&gcorelabs.InstanceInterfaceArgs{
					Type:      pulumi.String("subnet"),
					NetworkId: network.NetworkId,
					SubnetId:  subnet2.SubnetId,
				},
			},
			SecurityGroups: gcorelabs.InstanceSecurityGroupArray{
				&gcorelabs.InstanceSecurityGroupArgs{
					Id:   pulumi.String("66988147-f1b9-43b2-aaef-dee6d009b5b7"),
					Name: pulumi.String("default"),
				},
			},
			Metadatas: gcorelabs.InstanceMetadataArray{
				&gcorelabs.InstanceMetadataArgs{
					Key:   pulumi.String("some_key"),
					Value: pulumi.String("some_data"),
				},
			},
			Configurations: gcorelabs.InstanceConfigurationArray{
				&gcorelabs.InstanceConfigurationArgs{
					Key:   pulumi.String("some_key"),
					Value: pulumi.String("some_data"),
				},
			},
			RegionId:  pulumi.Float64(1),
			ProjectId: pulumi.Float64(1),
		})
		if err != nil {
			return err
		}
		lb, err := gcorelabs.NewLoadbalancer(ctx, "lb", &gcorelabs.LoadbalancerArgs{
			ProjectId: pulumi.Float64(1),
			RegionId:  pulumi.Float64(1),
			Name:      pulumi.String("test1"),
			Flavor:    pulumi.String("lb1-1-2"),
			Listener: &gcorelabs.LoadbalancerListenerArgs{
				Name:         pulumi.String("test"),
				Protocol:     pulumi.String("HTTP"),
				ProtocolPort: pulumi.Float64(80),
			},
		})
		if err != nil {
			return err
		}
		pl, err := gcorelabs.NewLbpool(ctx, "pl", &gcorelabs.LbpoolArgs{
			ProjectId:      pulumi.Float64(1),
			RegionId:       pulumi.Float64(1),
			Name:           pulumi.String("test_pool1"),
			Protocol:       pulumi.String("HTTP"),
			LbAlgorithm:    pulumi.String("LEAST_CONNECTIONS"),
			LoadbalancerId: lb.LoadbalancerId,
			ListenerId: pulumi.String(lb.Listener.ApplyT(func(listener gcorelabs.LoadbalancerListener) (*string, error) {
				return &listener.Id, nil
			}).(pulumi.StringPtrOutput)),
			HealthMonitor: &gcorelabs.LbpoolHealthMonitorArgs{
				Type:       pulumi.String("PING"),
				Delay:      pulumi.Float64(60),
				MaxRetries: pulumi.Float64(5),
				Timeout:    pulumi.Float64(10),
			},
			SessionPersistence: &gcorelabs.LbpoolSessionPersistenceArgs{
				Type:       pulumi.String("APP_COOKIE"),
				CookieName: pulumi.String("test_new_cookie"),
			},
		})
		if err != nil {
			return err
		}
		_, err = gcorelabs.NewLbmember(ctx, "lbm", &gcorelabs.LbmemberArgs{
			ProjectId:  pulumi.Float64(1),
			RegionId:   pulumi.Float64(1),
			PoolId:     pl.LbpoolId,
			InstanceId: instance.InstanceId,
			Address: pulumi.String(instance.Interfaces.ApplyT(func(interfaces []gcorelabs.InstanceInterface) (*string, error) {
				return &interfaces[0].IpAddress, nil
			}).(pulumi.StringPtrOutput)),
			ProtocolPort: pulumi.Float64(8081),
			Weight:       pulumi.Float64(5),
		})
		if err != nil {
			return err
		}
		instance2, err := gcorelabs.NewInstance(ctx, "instance2", &gcorelabs.InstanceArgs{
			FlavorId:    pulumi.String("g1-standard-2-4"),
			Name:        pulumi.String("test2"),
			KeypairName: kp.SshkeyName,
			Volumes: gcorelabs.InstanceVolumeArray{
				&gcorelabs.InstanceVolumeArgs{
					Source:    pulumi.String("existing-volume"),
					VolumeId:  secondVolume.VolumeId,
					BootIndex: pulumi.Float64(0),
				},
				&gcorelabs.InstanceVolumeArgs{
					Source:    pulumi.String("existing-volume"),
					VolumeId:  thirdVolume.VolumeId,
					BootIndex: pulumi.Float64(1),
				},
			},
			Interfaces: gcorelabs.InstanceInterfaceArray{
				&gcorelabs.InstanceInterfaceArgs{
					Type:      pulumi.String("subnet"),
					NetworkId: network.NetworkId,
					SubnetId:  subnet.SubnetId,
				},
			},
			SecurityGroups: gcorelabs.InstanceSecurityGroupArray{
				&gcorelabs.InstanceSecurityGroupArgs{
					Id:   pulumi.String("66988147-f1b9-43b2-aaef-dee6d009b5b7"),
					Name: pulumi.String("default"),
				},
			},
			Metadatas: gcorelabs.InstanceMetadataArray{
				&gcorelabs.InstanceMetadataArgs{
					Key:   pulumi.String("some_key"),
					Value: pulumi.String("some_data"),
				},
			},
			Configurations: gcorelabs.InstanceConfigurationArray{
				&gcorelabs.InstanceConfigurationArgs{
					Key:   pulumi.String("some_key"),
					Value: pulumi.String("some_data"),
				},
			},
			RegionId:  pulumi.Float64(1),
			ProjectId: pulumi.Float64(1),
		})
		if err != nil {
			return err
		}
		_, err = gcorelabs.NewLbmember(ctx, "lbm2", &gcorelabs.LbmemberArgs{
			ProjectId:  pulumi.Float64(1),
			RegionId:   pulumi.Float64(1),
			PoolId:     pl.LbpoolId,
			InstanceId: instance2.InstanceId,
			Address: pulumi.String(instance2.Interfaces.ApplyT(func(interfaces []gcorelabs.InstanceInterface) (*string, error) {
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
resources:
  kp:
    type: gcorelabs:Keypair
    properties:
      projectId: 1
      publicKey: your oub key
      sshkeyName: testkey
  network:
    type: gcorelabs:Network
    properties:
      name: network_example
      mtu: 1450
      type: vxlan
      regionId: 1
      projectId: 1
  subnet:
    type: gcorelabs:Subnet
    properties:
      name: subnet_example
      cidr: 192.168.10.0/24
      networkId: ${network.networkId}
      dnsNameservers:
        - 8.8.4.4
        - 1.1.1.1
      hostRoutes:
        - destination: 10.0.3.0/24
          nexthop: 10.0.0.13
      gatewayIp: 192.168.10.1
      regionId: 1
      projectId: 1
  subnet2:
    type: gcorelabs:Subnet
    properties:
      name: subnet2_example
      cidr: 192.168.20.0/24
      networkId: ${network.networkId}
      dnsNameservers:
        - 8.8.4.4
        - 1.1.1.1
      hostRoutes:
        - destination: 10.0.3.0/24
          nexthop: 10.0.0.13
      gatewayIp: 192.168.20.1
      regionId: 1
      projectId: 1
  firstVolume:
    type: gcorelabs:Volume
    name: first_volume
    properties:
      name: boot volume
      typeName: ssd_hiiops
      size: 6
      imageId: f4ce3d30-e29c-4cfd-811f-46f383b6081f
      regionId: 1
      projectId: 1
  secondVolume:
    type: gcorelabs:Volume
    name: second_volume
    properties:
      name: second volume
      typeName: ssd_hiiops
      imageId: f4ce3d30-e29c-4cfd-811f-46f383b6081f
      size: 6
      regionId: 1
      projectId: 1
  thirdVolume:
    type: gcorelabs:Volume
    name: third_volume
    properties:
      name: third volume
      typeName: ssd_hiiops
      size: 6
      regionId: 1
      projectId: 1
  instance:
    type: gcorelabs:Instance
    properties:
      flavorId: g1-standard-2-4
      name: test
      keypairName: ${kp.sshkeyName}
      volumes:
        - source: existing-volume
          volumeId: ${firstVolume.volumeId}
          bootIndex: 0
      interfaces:
        - type: subnet
          networkId: ${network.networkId}
          subnetId: ${subnet.subnetId}
        - type: subnet
          networkId: ${network.networkId}
          subnetId: ${subnet2.subnetId}
      securityGroups:
        - id: 66988147-f1b9-43b2-aaef-dee6d009b5b7
          name: default
      metadatas:
        - key: some_key
          value: some_data
      configurations:
        - key: some_key
          value: some_data
      regionId: 1
      projectId: 1
  lb:
    type: gcorelabs:Loadbalancer
    properties:
      projectId: 1
      regionId: 1
      name: test1
      flavor: lb1-1-2
      listener:
        name: test
        protocol: HTTP
        protocolPort: 80
  pl:
    type: gcorelabs:Lbpool
    properties:
      projectId: 1
      regionId: 1
      name: test_pool1
      protocol: HTTP
      lbAlgorithm: LEAST_CONNECTIONS
      loadbalancerId: ${lb.loadbalancerId}
      listenerId: ${lb.listener.id}
      healthMonitor:
        type: PING
        delay: 60
        maxRetries: 5
        timeout: 10
      sessionPersistence:
        type: APP_COOKIE
        cookieName: test_new_cookie
  lbm:
    type: gcorelabs:Lbmember
    properties:
      projectId: 1
      regionId: 1
      poolId: ${pl.lbpoolId}
      instanceId: ${instance.instanceId}
      address: ${instance.interfaces[0].ipAddress}
      protocolPort: 8081
      weight: 5
  instance2:
    type: gcorelabs:Instance
    properties:
      flavorId: g1-standard-2-4
      name: test2
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
        - id: 66988147-f1b9-43b2-aaef-dee6d009b5b7
          name: default
      metadatas:
        - key: some_key
          value: some_data
      configurations:
        - key: some_key
          value: some_data
      regionId: 1
      projectId: 1
  lbm2:
    type: gcorelabs:Lbmember
    properties:
      projectId: 1
      regionId: 1
      poolId: ${pl.lbpoolId}
      instanceId: ${instance2.instanceId}
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
import com.pulumi.gcorelabs.Keypair;
import com.pulumi.gcorelabs.KeypairArgs;
import com.pulumi.gcorelabs.Network;
import com.pulumi.gcorelabs.NetworkArgs;
import com.pulumi.gcorelabs.Subnet;
import com.pulumi.gcorelabs.SubnetArgs;
import com.pulumi.gcorelabs.inputs.SubnetHostRouteArgs;
import com.pulumi.gcorelabs.Volume;
import com.pulumi.gcorelabs.VolumeArgs;
import com.pulumi.gcorelabs.Instance;
import com.pulumi.gcorelabs.InstanceArgs;
import com.pulumi.gcorelabs.inputs.InstanceVolumeArgs;
import com.pulumi.gcorelabs.inputs.InstanceInterfaceArgs;
import com.pulumi.gcorelabs.inputs.InstanceSecurityGroupArgs;
import com.pulumi.gcorelabs.inputs.InstanceMetadataArgs;
import com.pulumi.gcorelabs.inputs.InstanceConfigurationArgs;
import com.pulumi.gcorelabs.Loadbalancer;
import com.pulumi.gcorelabs.LoadbalancerArgs;
import com.pulumi.gcorelabs.inputs.LoadbalancerListenerArgs;
import com.pulumi.gcorelabs.Lbpool;
import com.pulumi.gcorelabs.LbpoolArgs;
import com.pulumi.gcorelabs.inputs.LbpoolHealthMonitorArgs;
import com.pulumi.gcorelabs.inputs.LbpoolSessionPersistenceArgs;
import com.pulumi.gcorelabs.Lbmember;
import com.pulumi.gcorelabs.LbmemberArgs;
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
        var kp = new Keypair("kp", KeypairArgs.builder()
            .projectId(1)
            .publicKey("your oub key")
            .sshkeyName("testkey")
            .build());

        var network = new Network("network", NetworkArgs.builder()
            .name("network_example")
            .mtu(1450)
            .type("vxlan")
            .regionId(1)
            .projectId(1)
            .build());

        var subnet = new Subnet("subnet", SubnetArgs.builder()
            .name("subnet_example")
            .cidr("192.168.10.0/24")
            .networkId(network.networkId())
            .dnsNameservers(
                "8.8.4.4",
                "1.1.1.1")
            .hostRoutes(SubnetHostRouteArgs.builder()
                .destination("10.0.3.0/24")
                .nexthop("10.0.0.13")
                .build())
            .gatewayIp("192.168.10.1")
            .regionId(1)
            .projectId(1)
            .build());

        var subnet2 = new Subnet("subnet2", SubnetArgs.builder()
            .name("subnet2_example")
            .cidr("192.168.20.0/24")
            .networkId(network.networkId())
            .dnsNameservers(
                "8.8.4.4",
                "1.1.1.1")
            .hostRoutes(SubnetHostRouteArgs.builder()
                .destination("10.0.3.0/24")
                .nexthop("10.0.0.13")
                .build())
            .gatewayIp("192.168.20.1")
            .regionId(1)
            .projectId(1)
            .build());

        var firstVolume = new Volume("firstVolume", VolumeArgs.builder()
            .name("boot volume")
            .typeName("ssd_hiiops")
            .size(6)
            .imageId("f4ce3d30-e29c-4cfd-811f-46f383b6081f")
            .regionId(1)
            .projectId(1)
            .build());

        var secondVolume = new Volume("secondVolume", VolumeArgs.builder()
            .name("second volume")
            .typeName("ssd_hiiops")
            .imageId("f4ce3d30-e29c-4cfd-811f-46f383b6081f")
            .size(6)
            .regionId(1)
            .projectId(1)
            .build());

        var thirdVolume = new Volume("thirdVolume", VolumeArgs.builder()
            .name("third volume")
            .typeName("ssd_hiiops")
            .size(6)
            .regionId(1)
            .projectId(1)
            .build());

        var instance = new Instance("instance", InstanceArgs.builder()
            .flavorId("g1-standard-2-4")
            .name("test")
            .keypairName(kp.sshkeyName())
            .volumes(InstanceVolumeArgs.builder()
                .source("existing-volume")
                .volumeId(firstVolume.volumeId())
                .bootIndex(0)
                .build())
            .interfaces(
                InstanceInterfaceArgs.builder()
                    .type("subnet")
                    .networkId(network.networkId())
                    .subnetId(subnet.subnetId())
                    .build(),
                InstanceInterfaceArgs.builder()
                    .type("subnet")
                    .networkId(network.networkId())
                    .subnetId(subnet2.subnetId())
                    .build())
            .securityGroups(InstanceSecurityGroupArgs.builder()
                .id("66988147-f1b9-43b2-aaef-dee6d009b5b7")
                .name("default")
                .build())
            .metadatas(InstanceMetadataArgs.builder()
                .key("some_key")
                .value("some_data")
                .build())
            .configurations(InstanceConfigurationArgs.builder()
                .key("some_key")
                .value("some_data")
                .build())
            .regionId(1)
            .projectId(1)
            .build());

        var lb = new Loadbalancer("lb", LoadbalancerArgs.builder()
            .projectId(1)
            .regionId(1)
            .name("test1")
            .flavor("lb1-1-2")
            .listener(LoadbalancerListenerArgs.builder()
                .name("test")
                .protocol("HTTP")
                .protocolPort(80)
                .build())
            .build());

        var pl = new Lbpool("pl", LbpoolArgs.builder()
            .projectId(1)
            .regionId(1)
            .name("test_pool1")
            .protocol("HTTP")
            .lbAlgorithm("LEAST_CONNECTIONS")
            .loadbalancerId(lb.loadbalancerId())
            .listenerId(lb.listener().applyValue(listener -> listener.id()))
            .healthMonitor(LbpoolHealthMonitorArgs.builder()
                .type("PING")
                .delay(60)
                .maxRetries(5)
                .timeout(10)
                .build())
            .sessionPersistence(LbpoolSessionPersistenceArgs.builder()
                .type("APP_COOKIE")
                .cookieName("test_new_cookie")
                .build())
            .build());

        var lbm = new Lbmember("lbm", LbmemberArgs.builder()
            .projectId(1)
            .regionId(1)
            .poolId(pl.lbpoolId())
            .instanceId(instance.instanceId())
            .address(instance.interfaces().applyValue(interfaces -> interfaces[0].ipAddress()))
            .protocolPort(8081)
            .weight(5)
            .build());

        var instance2 = new Instance("instance2", InstanceArgs.builder()
            .flavorId("g1-standard-2-4")
            .name("test2")
            .keypairName(kp.sshkeyName())
            .volumes(
                InstanceVolumeArgs.builder()
                    .source("existing-volume")
                    .volumeId(secondVolume.volumeId())
                    .bootIndex(0)
                    .build(),
                InstanceVolumeArgs.builder()
                    .source("existing-volume")
                    .volumeId(thirdVolume.volumeId())
                    .bootIndex(1)
                    .build())
            .interfaces(InstanceInterfaceArgs.builder()
                .type("subnet")
                .networkId(network.networkId())
                .subnetId(subnet.subnetId())
                .build())
            .securityGroups(InstanceSecurityGroupArgs.builder()
                .id("66988147-f1b9-43b2-aaef-dee6d009b5b7")
                .name("default")
                .build())
            .metadatas(InstanceMetadataArgs.builder()
                .key("some_key")
                .value("some_data")
                .build())
            .configurations(InstanceConfigurationArgs.builder()
                .key("some_key")
                .value("some_data")
                .build())
            .regionId(1)
            .projectId(1)
            .build());

        var lbm2 = new Lbmember("lbm2", LbmemberArgs.builder()
            .projectId(1)
            .regionId(1)
            .poolId(pl.lbpoolId())
            .instanceId(instance2.instanceId())
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

- `apiEndpoint` (String) A single API endpoint for all products. Will be used when specific product API url is not defined.
- `gcoreApi` (String, Deprecated) Region API
- `gcoreCdnApi` (String) CDN API (define only if you want to override CDN API endpoint)
- `gcoreClientId` (String) Client id
- `gcoreCloudApi` (String) Region API (define only if you want to override Region API endpoint)
- `gcoreDnsApi` (String) DNS API (define only if you want to override DNS API endpoint)
- `gcorePlatform` (String, Deprecated) Platform URL is used for generate JWT
- `gcorePlatformApi` (String) Platform URL is used for generate JWT (define only if you want to override Platform API endpoint)
- `gcoreStorageApi` (String) Storage API (define only if you want to override Storage API endpoint)
- `ignoreCredsAuthError` (Boolean, Deprecated) Should be set to true when you are gonna to use storage resource with permanent API-token only.
- `password` (String, Deprecated)
- `permanentApiToken` (String, Sensitive) A permanent [API-token](https://support.gcorelabs.com/hc/en-us/articles/360018625617-API-tokens)
- `userName` (String, Deprecated)