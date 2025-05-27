---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/edge-center/edgecenter/0.8.1/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Edgecenter Provider
meta_desc: Provides an overview on how to configure the Pulumi Edgecenter provider.
layout: package
---

## Generate Provider

The Edgecenter provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider edge-center/edgecenter
```
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    edgecenter:permanentApiToken:
        value: 251$d3361.............1b35f26d8

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as edgecenter from "@pulumi/edgecenter";

const kp = new edgecenter.Keypair("kp", {
    projectId: 1,
    publicKey: "your oub key",
    sshkeyName: "testkey",
});
const network = new edgecenter.Network("network", {
    name: "network_example",
    type: "vxlan",
    regionId: 1,
    projectId: 1,
});
const subnet = new edgecenter.Subnet("subnet", {
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
const subnet2 = new edgecenter.Subnet("subnet2", {
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
const firstVolume = new edgecenter.Volume("first_volume", {
    name: "boot volume",
    typeName: "ssd_hiiops",
    size: 6,
    imageId: "f4ce3d30-e29c-4cfd-811f-46f383b6081f",
    regionId: 1,
    projectId: 1,
});
const secondVolume = new edgecenter.Volume("second_volume", {
    name: "second volume",
    typeName: "ssd_hiiops",
    imageId: "f4ce3d30-e29c-4cfd-811f-46f383b6081f",
    size: 6,
    regionId: 1,
    projectId: 1,
});
const thirdVolume = new edgecenter.Volume("third_volume", {
    name: "third volume",
    typeName: "ssd_hiiops",
    size: 6,
    regionId: 1,
    projectId: 1,
});
const instance = new edgecenter.Instance("instance", {
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
const lb = new edgecenter.Loadbalancer("lb", {
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
const pl = new edgecenter.Lbpool("pl", {
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
const lbm = new edgecenter.Lbmember("lbm", {
    projectId: 1,
    regionId: 1,
    poolId: pl.lbpoolId,
    instanceId: instance.instanceId,
    address: instance.interfaces.apply(interfaces => interfaces[0].ipAddress),
    protocolPort: 8081,
    weight: 5,
});
const instance2 = new edgecenter.Instance("instance2", {
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
const lbm2 = new edgecenter.Lbmember("lbm2", {
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
    edgecenter:permanentApiToken:
        value: 251$d3361.............1b35f26d8

```
```python
import pulumi
import pulumi_edgecenter as edgecenter

kp = edgecenter.Keypair("kp",
    project_id=1,
    public_key="your oub key",
    sshkey_name="testkey")
network = edgecenter.Network("network",
    name="network_example",
    type="vxlan",
    region_id=1,
    project_id=1)
subnet = edgecenter.Subnet("subnet",
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
subnet2 = edgecenter.Subnet("subnet2",
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
first_volume = edgecenter.Volume("first_volume",
    name="boot volume",
    type_name="ssd_hiiops",
    size=6,
    image_id="f4ce3d30-e29c-4cfd-811f-46f383b6081f",
    region_id=1,
    project_id=1)
second_volume = edgecenter.Volume("second_volume",
    name="second volume",
    type_name="ssd_hiiops",
    image_id="f4ce3d30-e29c-4cfd-811f-46f383b6081f",
    size=6,
    region_id=1,
    project_id=1)
third_volume = edgecenter.Volume("third_volume",
    name="third volume",
    type_name="ssd_hiiops",
    size=6,
    region_id=1,
    project_id=1)
instance = edgecenter.Instance("instance",
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
lb = edgecenter.Loadbalancer("lb",
    project_id=1,
    region_id=1,
    name="test1",
    flavor="lb1-1-2",
    listener={
        "name": "test",
        "protocol": "HTTP",
        "protocol_port": 80,
    })
pl = edgecenter.Lbpool("pl",
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
lbm = edgecenter.Lbmember("lbm",
    project_id=1,
    region_id=1,
    pool_id=pl.lbpool_id,
    instance_id=instance.instance_id,
    address=instance.interfaces[0].ip_address,
    protocol_port=8081,
    weight=5)
instance2 = edgecenter.Instance("instance2",
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
lbm2 = edgecenter.Lbmember("lbm2",
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
    edgecenter:permanentApiToken:
        value: 251$d3361.............1b35f26d8

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Edgecenter = Pulumi.Edgecenter;

return await Deployment.RunAsync(() =>
{
    var kp = new Edgecenter.Keypair("kp", new()
    {
        ProjectId = 1,
        PublicKey = "your oub key",
        SshkeyName = "testkey",
    });

    var network = new Edgecenter.Network("network", new()
    {
        Name = "network_example",
        Type = "vxlan",
        RegionId = 1,
        ProjectId = 1,
    });

    var subnet = new Edgecenter.Subnet("subnet", new()
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
            new Edgecenter.Inputs.SubnetHostRouteArgs
            {
                Destination = "10.0.3.0/24",
                Nexthop = "10.0.0.13",
            },
        },
        GatewayIp = "192.168.10.1",
        RegionId = 1,
        ProjectId = 1,
    });

    var subnet2 = new Edgecenter.Subnet("subnet2", new()
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
            new Edgecenter.Inputs.SubnetHostRouteArgs
            {
                Destination = "10.0.3.0/24",
                Nexthop = "10.0.0.13",
            },
        },
        GatewayIp = "192.168.20.1",
        RegionId = 1,
        ProjectId = 1,
    });

    var firstVolume = new Edgecenter.Volume("first_volume", new()
    {
        Name = "boot volume",
        TypeName = "ssd_hiiops",
        Size = 6,
        ImageId = "f4ce3d30-e29c-4cfd-811f-46f383b6081f",
        RegionId = 1,
        ProjectId = 1,
    });

    var secondVolume = new Edgecenter.Volume("second_volume", new()
    {
        Name = "second volume",
        TypeName = "ssd_hiiops",
        ImageId = "f4ce3d30-e29c-4cfd-811f-46f383b6081f",
        Size = 6,
        RegionId = 1,
        ProjectId = 1,
    });

    var thirdVolume = new Edgecenter.Volume("third_volume", new()
    {
        Name = "third volume",
        TypeName = "ssd_hiiops",
        Size = 6,
        RegionId = 1,
        ProjectId = 1,
    });

    var instance = new Edgecenter.Instance("instance", new()
    {
        FlavorId = "g1-standard-2-4",
        Name = "test",
        KeypairName = kp.SshkeyName,
        Volumes = new[]
        {
            new Edgecenter.Inputs.InstanceVolumeArgs
            {
                Source = "existing-volume",
                VolumeId = firstVolume.VolumeId,
                BootIndex = 0,
            },
        },
        Interfaces = new[]
        {
            new Edgecenter.Inputs.InstanceInterfaceArgs
            {
                Type = "subnet",
                NetworkId = network.NetworkId,
                SubnetId = subnet.SubnetId,
            },
            new Edgecenter.Inputs.InstanceInterfaceArgs
            {
                Type = "subnet",
                NetworkId = network.NetworkId,
                SubnetId = subnet2.SubnetId,
            },
        },
        SecurityGroups = new[]
        {
            new Edgecenter.Inputs.InstanceSecurityGroupArgs
            {
                Id = "66988147-f1b9-43b2-aaef-dee6d009b5b7",
                Name = "default",
            },
        },
        Metadatas = new[]
        {
            new Edgecenter.Inputs.InstanceMetadataArgs
            {
                Key = "some_key",
                Value = "some_data",
            },
        },
        Configurations = new[]
        {
            new Edgecenter.Inputs.InstanceConfigurationArgs
            {
                Key = "some_key",
                Value = "some_data",
            },
        },
        RegionId = 1,
        ProjectId = 1,
    });

    var lb = new Edgecenter.Loadbalancer("lb", new()
    {
        ProjectId = 1,
        RegionId = 1,
        Name = "test1",
        Flavor = "lb1-1-2",
        Listener = new Edgecenter.Inputs.LoadbalancerListenerArgs
        {
            Name = "test",
            Protocol = "HTTP",
            ProtocolPort = 80,
        },
    });

    var pl = new Edgecenter.Lbpool("pl", new()
    {
        ProjectId = 1,
        RegionId = 1,
        Name = "test_pool1",
        Protocol = "HTTP",
        LbAlgorithm = "LEAST_CONNECTIONS",
        LoadbalancerId = lb.LoadbalancerId,
        ListenerId = lb.Listener.Apply(listener => listener.Id),
        HealthMonitor = new Edgecenter.Inputs.LbpoolHealthMonitorArgs
        {
            Type = "PING",
            Delay = 60,
            MaxRetries = 5,
            Timeout = 10,
        },
        SessionPersistence = new Edgecenter.Inputs.LbpoolSessionPersistenceArgs
        {
            Type = "APP_COOKIE",
            CookieName = "test_new_cookie",
        },
    });

    var lbm = new Edgecenter.Lbmember("lbm", new()
    {
        ProjectId = 1,
        RegionId = 1,
        PoolId = pl.LbpoolId,
        InstanceId = instance.InstanceId,
        Address = instance.Interfaces.Apply(interfaces => interfaces[0].IpAddress),
        ProtocolPort = 8081,
        Weight = 5,
    });

    var instance2 = new Edgecenter.Instance("instance2", new()
    {
        FlavorId = "g1-standard-2-4",
        Name = "test2",
        KeypairName = kp.SshkeyName,
        Volumes = new[]
        {
            new Edgecenter.Inputs.InstanceVolumeArgs
            {
                Source = "existing-volume",
                VolumeId = secondVolume.VolumeId,
                BootIndex = 0,
            },
            new Edgecenter.Inputs.InstanceVolumeArgs
            {
                Source = "existing-volume",
                VolumeId = thirdVolume.VolumeId,
                BootIndex = 1,
            },
        },
        Interfaces = new[]
        {
            new Edgecenter.Inputs.InstanceInterfaceArgs
            {
                Type = "subnet",
                NetworkId = network.NetworkId,
                SubnetId = subnet.SubnetId,
            },
        },
        SecurityGroups = new[]
        {
            new Edgecenter.Inputs.InstanceSecurityGroupArgs
            {
                Id = "66988147-f1b9-43b2-aaef-dee6d009b5b7",
                Name = "default",
            },
        },
        Metadatas = new[]
        {
            new Edgecenter.Inputs.InstanceMetadataArgs
            {
                Key = "some_key",
                Value = "some_data",
            },
        },
        Configurations = new[]
        {
            new Edgecenter.Inputs.InstanceConfigurationArgs
            {
                Key = "some_key",
                Value = "some_data",
            },
        },
        RegionId = 1,
        ProjectId = 1,
    });

    var lbm2 = new Edgecenter.Lbmember("lbm2", new()
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
    edgecenter:permanentApiToken:
        value: 251$d3361.............1b35f26d8

```
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/edgecenter/edgecenter"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		kp, err := edgecenter.NewKeypair(ctx, "kp", &edgecenter.KeypairArgs{
			ProjectId:  pulumi.Float64(1),
			PublicKey:  pulumi.String("your oub key"),
			SshkeyName: pulumi.String("testkey"),
		})
		if err != nil {
			return err
		}
		network, err := edgecenter.NewNetwork(ctx, "network", &edgecenter.NetworkArgs{
			Name:      pulumi.String("network_example"),
			Type:      pulumi.String("vxlan"),
			RegionId:  pulumi.Float64(1),
			ProjectId: pulumi.Float64(1),
		})
		if err != nil {
			return err
		}
		subnet, err := edgecenter.NewSubnet(ctx, "subnet", &edgecenter.SubnetArgs{
			Name:      pulumi.String("subnet_example"),
			Cidr:      pulumi.String("192.168.10.0/24"),
			NetworkId: network.NetworkId,
			DnsNameservers: pulumi.StringArray{
				pulumi.String("8.8.4.4"),
				pulumi.String("1.1.1.1"),
			},
			HostRoutes: edgecenter.SubnetHostRouteArray{
				&edgecenter.SubnetHostRouteArgs{
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
		subnet2, err := edgecenter.NewSubnet(ctx, "subnet2", &edgecenter.SubnetArgs{
			Name:      pulumi.String("subnet2_example"),
			Cidr:      pulumi.String("192.168.20.0/24"),
			NetworkId: network.NetworkId,
			DnsNameservers: pulumi.StringArray{
				pulumi.String("8.8.4.4"),
				pulumi.String("1.1.1.1"),
			},
			HostRoutes: edgecenter.SubnetHostRouteArray{
				&edgecenter.SubnetHostRouteArgs{
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
		firstVolume, err := edgecenter.NewVolume(ctx, "first_volume", &edgecenter.VolumeArgs{
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
		secondVolume, err := edgecenter.NewVolume(ctx, "second_volume", &edgecenter.VolumeArgs{
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
		thirdVolume, err := edgecenter.NewVolume(ctx, "third_volume", &edgecenter.VolumeArgs{
			Name:      pulumi.String("third volume"),
			TypeName:  pulumi.String("ssd_hiiops"),
			Size:      pulumi.Float64(6),
			RegionId:  pulumi.Float64(1),
			ProjectId: pulumi.Float64(1),
		})
		if err != nil {
			return err
		}
		instance, err := edgecenter.NewInstance(ctx, "instance", &edgecenter.InstanceArgs{
			FlavorId:    pulumi.String("g1-standard-2-4"),
			Name:        pulumi.String("test"),
			KeypairName: kp.SshkeyName,
			Volumes: edgecenter.InstanceVolumeArray{
				&edgecenter.InstanceVolumeArgs{
					Source:    pulumi.String("existing-volume"),
					VolumeId:  firstVolume.VolumeId,
					BootIndex: pulumi.Float64(0),
				},
			},
			Interfaces: edgecenter.InstanceInterfaceArray{
				&edgecenter.InstanceInterfaceArgs{
					Type:      pulumi.String("subnet"),
					NetworkId: network.NetworkId,
					SubnetId:  subnet.SubnetId,
				},
				&edgecenter.InstanceInterfaceArgs{
					Type:      pulumi.String("subnet"),
					NetworkId: network.NetworkId,
					SubnetId:  subnet2.SubnetId,
				},
			},
			SecurityGroups: []map[string]interface{}{
				map[string]interface{}{
					"id":   "66988147-f1b9-43b2-aaef-dee6d009b5b7",
					"name": "default",
				},
			},
			Metadatas: edgecenter.InstanceMetadataArray{
				&edgecenter.InstanceMetadataArgs{
					Key:   pulumi.String("some_key"),
					Value: pulumi.String("some_data"),
				},
			},
			Configurations: edgecenter.InstanceConfigurationArray{
				&edgecenter.InstanceConfigurationArgs{
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
		lb, err := edgecenter.NewLoadbalancer(ctx, "lb", &edgecenter.LoadbalancerArgs{
			ProjectId: pulumi.Float64(1),
			RegionId:  pulumi.Float64(1),
			Name:      pulumi.String("test1"),
			Flavor:    pulumi.String("lb1-1-2"),
			Listener: &edgecenter.LoadbalancerListenerArgs{
				Name:         pulumi.String("test"),
				Protocol:     pulumi.String("HTTP"),
				ProtocolPort: pulumi.Float64(80),
			},
		})
		if err != nil {
			return err
		}
		pl, err := edgecenter.NewLbpool(ctx, "pl", &edgecenter.LbpoolArgs{
			ProjectId:      pulumi.Float64(1),
			RegionId:       pulumi.Float64(1),
			Name:           pulumi.String("test_pool1"),
			Protocol:       pulumi.String("HTTP"),
			LbAlgorithm:    pulumi.String("LEAST_CONNECTIONS"),
			LoadbalancerId: lb.LoadbalancerId,
			ListenerId: pulumi.String(lb.Listener.ApplyT(func(listener edgecenter.LoadbalancerListener) (*string, error) {
				return &listener.Id, nil
			}).(pulumi.StringPtrOutput)),
			HealthMonitor: &edgecenter.LbpoolHealthMonitorArgs{
				Type:       pulumi.String("PING"),
				Delay:      pulumi.Float64(60),
				MaxRetries: pulumi.Float64(5),
				Timeout:    pulumi.Float64(10),
			},
			SessionPersistence: &edgecenter.LbpoolSessionPersistenceArgs{
				Type:       pulumi.String("APP_COOKIE"),
				CookieName: pulumi.String("test_new_cookie"),
			},
		})
		if err != nil {
			return err
		}
		_, err = edgecenter.NewLbmember(ctx, "lbm", &edgecenter.LbmemberArgs{
			ProjectId:  pulumi.Float64(1),
			RegionId:   pulumi.Float64(1),
			PoolId:     pl.LbpoolId,
			InstanceId: instance.InstanceId,
			Address: pulumi.String(instance.Interfaces.ApplyT(func(interfaces []edgecenter.InstanceInterface) (*string, error) {
				return &interfaces[0].IpAddress, nil
			}).(pulumi.StringPtrOutput)),
			ProtocolPort: pulumi.Float64(8081),
			Weight:       pulumi.Float64(5),
		})
		if err != nil {
			return err
		}
		instance2, err := edgecenter.NewInstance(ctx, "instance2", &edgecenter.InstanceArgs{
			FlavorId:    pulumi.String("g1-standard-2-4"),
			Name:        pulumi.String("test2"),
			KeypairName: kp.SshkeyName,
			Volumes: edgecenter.InstanceVolumeArray{
				&edgecenter.InstanceVolumeArgs{
					Source:    pulumi.String("existing-volume"),
					VolumeId:  secondVolume.VolumeId,
					BootIndex: pulumi.Float64(0),
				},
				&edgecenter.InstanceVolumeArgs{
					Source:    pulumi.String("existing-volume"),
					VolumeId:  thirdVolume.VolumeId,
					BootIndex: pulumi.Float64(1),
				},
			},
			Interfaces: edgecenter.InstanceInterfaceArray{
				&edgecenter.InstanceInterfaceArgs{
					Type:      pulumi.String("subnet"),
					NetworkId: network.NetworkId,
					SubnetId:  subnet.SubnetId,
				},
			},
			SecurityGroups: []map[string]interface{}{
				map[string]interface{}{
					"id":   "66988147-f1b9-43b2-aaef-dee6d009b5b7",
					"name": "default",
				},
			},
			Metadatas: edgecenter.InstanceMetadataArray{
				&edgecenter.InstanceMetadataArgs{
					Key:   pulumi.String("some_key"),
					Value: pulumi.String("some_data"),
				},
			},
			Configurations: edgecenter.InstanceConfigurationArray{
				&edgecenter.InstanceConfigurationArgs{
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
		_, err = edgecenter.NewLbmember(ctx, "lbm2", &edgecenter.LbmemberArgs{
			ProjectId:  pulumi.Float64(1),
			RegionId:   pulumi.Float64(1),
			PoolId:     pl.LbpoolId,
			InstanceId: instance2.InstanceId,
			Address: pulumi.String(instance2.Interfaces.ApplyT(func(interfaces []edgecenter.InstanceInterface) (*string, error) {
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
    edgecenter:permanentApiToken:
        value: 251$d3361.............1b35f26d8

```
```yaml
resources:
  kp:
    type: edgecenter:Keypair
    properties:
      projectId: 1
      publicKey: your oub key
      sshkeyName: testkey
  network:
    type: edgecenter:Network
    properties:
      name: network_example
      type: vxlan
      regionId: 1
      projectId: 1
  subnet:
    type: edgecenter:Subnet
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
    type: edgecenter:Subnet
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
    type: edgecenter:Volume
    name: first_volume
    properties:
      name: boot volume
      typeName: ssd_hiiops
      size: 6
      imageId: f4ce3d30-e29c-4cfd-811f-46f383b6081f
      regionId: 1
      projectId: 1
  secondVolume:
    type: edgecenter:Volume
    name: second_volume
    properties:
      name: second volume
      typeName: ssd_hiiops
      imageId: f4ce3d30-e29c-4cfd-811f-46f383b6081f
      size: 6
      regionId: 1
      projectId: 1
  thirdVolume:
    type: edgecenter:Volume
    name: third_volume
    properties:
      name: third volume
      typeName: ssd_hiiops
      size: 6
      regionId: 1
      projectId: 1
  instance:
    type: edgecenter:Instance
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
    type: edgecenter:Loadbalancer
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
    type: edgecenter:Lbpool
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
    type: edgecenter:Lbmember
    properties:
      projectId: 1
      regionId: 1
      poolId: ${pl.lbpoolId}
      instanceId: ${instance.instanceId}
      address: ${instance.interfaces[0].ipAddress}
      protocolPort: 8081
      weight: 5
  instance2:
    type: edgecenter:Instance
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
    type: edgecenter:Lbmember
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
    edgecenter:permanentApiToken:
        value: 251$d3361.............1b35f26d8

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.edgecenter.Keypair;
import com.pulumi.edgecenter.KeypairArgs;
import com.pulumi.edgecenter.Network;
import com.pulumi.edgecenter.NetworkArgs;
import com.pulumi.edgecenter.Subnet;
import com.pulumi.edgecenter.SubnetArgs;
import com.pulumi.edgecenter.inputs.SubnetHostRouteArgs;
import com.pulumi.edgecenter.Volume;
import com.pulumi.edgecenter.VolumeArgs;
import com.pulumi.edgecenter.Instance;
import com.pulumi.edgecenter.InstanceArgs;
import com.pulumi.edgecenter.inputs.InstanceVolumeArgs;
import com.pulumi.edgecenter.inputs.InstanceInterfaceArgs;
import com.pulumi.edgecenter.inputs.InstanceMetadataArgs;
import com.pulumi.edgecenter.inputs.InstanceConfigurationArgs;
import com.pulumi.edgecenter.Loadbalancer;
import com.pulumi.edgecenter.LoadbalancerArgs;
import com.pulumi.edgecenter.inputs.LoadbalancerListenerArgs;
import com.pulumi.edgecenter.Lbpool;
import com.pulumi.edgecenter.LbpoolArgs;
import com.pulumi.edgecenter.inputs.LbpoolHealthMonitorArgs;
import com.pulumi.edgecenter.inputs.LbpoolSessionPersistenceArgs;
import com.pulumi.edgecenter.Lbmember;
import com.pulumi.edgecenter.LbmemberArgs;
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
            .securityGroups(%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference))
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
            .securityGroups(%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference))
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
- `edgecenterApi` (String, Deprecated) Region API
- `edgecenterCdnApi` (String) CDN API (define only if you want to override CDN API endpoint)
- `edgecenterCloudApi` (String) Region API (define only if you want to override Region API endpoint)
- `edgecenterDnsApi` (String) DNS API (define only if you want to override DNS API endpoint)
- `edgecenterPlatform` (String, Deprecated) Platform URL is used for generate JWT
- `edgecenterPlatformApi` (String) Platform URL is used for generate JWT (define only if you want to override Platform API endpoint)
- `edgecenterStorageApi` (String) Storage API (define only if you want to override Storage API endpoint)
- `ignoreCredsAuthError` (Boolean, Deprecated) Should be set to true when you are gonna to use storage resource with permanent API-token only.
- `password` (String, Deprecated)
- `permanentApiToken` (String, Sensitive) A permanent [API-token](https://support.edgecenter.ru/knowledge_base/item/257788)
- `userName` (String, Deprecated)