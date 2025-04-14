---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/dell/powerflex/1.8.0/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Powerflex Provider
meta_desc: Provides an overview on how to configure the Pulumi Powerflex provider.
layout: package
---

## Generate Provider

The Powerflex provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider dell/powerflex
```
## Overview

The Pulumi provider for Dell PowerFlex can be used to interact with a Dell PowerFlex array in order to manage the array resources.
## Example Usage

provider.tf
{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    powerflex:endpoint:
        value: 'TODO: var.endpoint'
    powerflex:insecure:
        value: true
    powerflex:password:
        value: 'TODO: var.password'
    powerflex:timeout:
        value: 120
    powerflex:username:
        value: 'TODO: var.username'

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as powerflex from "@pulumi/powerflex";

const pd = new powerflex.ProtectionDomain("pd", {name: "domain_1"});
const sds1 = new powerflex.Sds("sds1", {
    name: "sds_1",
    protectionDomainId: pd.id,
    ipLists: [{
        ip: "10.10.10.1",
        role: "all",
    }],
}, {
    dependsOn: [pd],
});
const sds2 = new powerflex.Sds("sds2", {
    name: "sds_2",
    protectionDomainId: pd.id,
    ipLists: [{
        ip: "10.10.10.2",
        role: "all",
    }],
}, {
    dependsOn: [pd],
});
const sds3 = new powerflex.Sds("sds3", {
    name: "sds_3",
    protectionDomainId: pd.id,
    ipLists: [{
        ip: "10.10.10.3",
        role: "all",
    }],
}, {
    dependsOn: [pd],
});
const sp = new powerflex.StoragePool("sp", {
    name: "SP",
    protectionDomainId: pd.id,
    mediaType: "HDD",
    useRmcache: true,
    useRfcache: true,
});
const device1 = new powerflex.Device("device1", {
    name: "device1",
    devicePath: "/dev/sdb",
    sdsId: sds1.id,
    storagePoolId: sp.id,
    mediaType: "HDD",
    externalAccelerationType: "ReadAndWrite",
}, {
    dependsOn: [sp],
});
const device2 = new powerflex.Device("device2", {
    name: "device2",
    devicePath: "/dev/sdb",
    sdsId: sds2.id,
    storagePoolId: sp.id,
    mediaType: "HDD",
    externalAccelerationType: "ReadAndWrite",
}, {
    dependsOn: [sp],
});
const device3 = new powerflex.Device("device3", {
    name: "device3",
    devicePath: "/dev/sdb",
    sdsId: sds3.id,
    storagePoolId: sp.id,
    mediaType: "HDD",
    externalAccelerationType: "ReadAndWrite",
}, {
    dependsOn: [sp],
});
const volume = new powerflex.Volume("volume", {
    name: "volume1",
    protectionDomainId: pd.id,
    storagePoolId: sp.id,
    size: 16,
    volumeType: "ThinProvisioned",
}, {
    dependsOn: [
        device1,
        device2,
        device3,
    ],
});
const map = new powerflex.SdcVolumesMapping("map", {
    sdcVolumesMappingId: "e3d105e900000005",
    volumeLists: [{
        volumeId: volume.id,
        limitIops: 140,
        limitBwInMbps: 19,
        accessMode: "ReadOnly",
    }],
}, {
    dependsOn: [volume],
});
const upload_test = new powerflex.Package("upload-test", {filePaths: [
    "/root/powerflex_packages/PowerFlex_3.6.700.103_RHEL_OEL7/EMC-ScaleIO-lia-3.6-700.103.el7.x86_64.rpm",
    "/root/powerflex_packages/PowerFlex_3.6.700.103_RHEL_OEL7/EMC-ScaleIO-mdm-3.6-700.103.el7.x86_64.rpm",
    "/root/powerflex_packages/PowerFlex_3.6.700.103_RHEL_OEL7/EMC-ScaleIO-sds-3.6-700.103.el7.x86_64.rpm",
    "/root/powerflex_packages/PowerFlex_3.6.700.103_RHEL_OEL7/EMC-ScaleIO-sdc-3.6-700.103.el7.x86_64.rpm",
    "/root/powerflex_packages/PowerFlex_3.6.700.103_RHEL_OEL7/EMC-ScaleIO-sdr-3.6-700.103.el7.x86_64.rpm",
    "/root/powerflex_packages/PowerFlex_3.6.700.103_RHEL_OEL7/EMC-ScaleIO-sdt-3.6-700.103.el7.x86_64.rpm",
]});
const user = new powerflex.User("user", {
    name: "NewUser",
    role: "Monitor",
    password: "Password123",
});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    powerflex:endpoint:
        value: 'TODO: var.endpoint'
    powerflex:insecure:
        value: true
    powerflex:password:
        value: 'TODO: var.password'
    powerflex:timeout:
        value: 120
    powerflex:username:
        value: 'TODO: var.username'

```
```python
import pulumi
import pulumi_powerflex as powerflex

pd = powerflex.ProtectionDomain("pd", name="domain_1")
sds1 = powerflex.Sds("sds1",
    name="sds_1",
    protection_domain_id=pd.id,
    ip_lists=[{
        "ip": "10.10.10.1",
        "role": "all",
    }],
    opts = pulumi.ResourceOptions(depends_on=[pd]))
sds2 = powerflex.Sds("sds2",
    name="sds_2",
    protection_domain_id=pd.id,
    ip_lists=[{
        "ip": "10.10.10.2",
        "role": "all",
    }],
    opts = pulumi.ResourceOptions(depends_on=[pd]))
sds3 = powerflex.Sds("sds3",
    name="sds_3",
    protection_domain_id=pd.id,
    ip_lists=[{
        "ip": "10.10.10.3",
        "role": "all",
    }],
    opts = pulumi.ResourceOptions(depends_on=[pd]))
sp = powerflex.StoragePool("sp",
    name="SP",
    protection_domain_id=pd.id,
    media_type="HDD",
    use_rmcache=True,
    use_rfcache=True)
device1 = powerflex.Device("device1",
    name="device1",
    device_path="/dev/sdb",
    sds_id=sds1.id,
    storage_pool_id=sp.id,
    media_type="HDD",
    external_acceleration_type="ReadAndWrite",
    opts = pulumi.ResourceOptions(depends_on=[sp]))
device2 = powerflex.Device("device2",
    name="device2",
    device_path="/dev/sdb",
    sds_id=sds2.id,
    storage_pool_id=sp.id,
    media_type="HDD",
    external_acceleration_type="ReadAndWrite",
    opts = pulumi.ResourceOptions(depends_on=[sp]))
device3 = powerflex.Device("device3",
    name="device3",
    device_path="/dev/sdb",
    sds_id=sds3.id,
    storage_pool_id=sp.id,
    media_type="HDD",
    external_acceleration_type="ReadAndWrite",
    opts = pulumi.ResourceOptions(depends_on=[sp]))
volume = powerflex.Volume("volume",
    name="volume1",
    protection_domain_id=pd.id,
    storage_pool_id=sp.id,
    size=16,
    volume_type="ThinProvisioned",
    opts = pulumi.ResourceOptions(depends_on=[
            device1,
            device2,
            device3,
        ]))
map = powerflex.SdcVolumesMapping("map",
    sdc_volumes_mapping_id="e3d105e900000005",
    volume_lists=[{
        "volume_id": volume.id,
        "limit_iops": 140,
        "limit_bw_in_mbps": 19,
        "access_mode": "ReadOnly",
    }],
    opts = pulumi.ResourceOptions(depends_on=[volume]))
upload_test = powerflex.Package("upload-test", file_paths=[
    "/root/powerflex_packages/PowerFlex_3.6.700.103_RHEL_OEL7/EMC-ScaleIO-lia-3.6-700.103.el7.x86_64.rpm",
    "/root/powerflex_packages/PowerFlex_3.6.700.103_RHEL_OEL7/EMC-ScaleIO-mdm-3.6-700.103.el7.x86_64.rpm",
    "/root/powerflex_packages/PowerFlex_3.6.700.103_RHEL_OEL7/EMC-ScaleIO-sds-3.6-700.103.el7.x86_64.rpm",
    "/root/powerflex_packages/PowerFlex_3.6.700.103_RHEL_OEL7/EMC-ScaleIO-sdc-3.6-700.103.el7.x86_64.rpm",
    "/root/powerflex_packages/PowerFlex_3.6.700.103_RHEL_OEL7/EMC-ScaleIO-sdr-3.6-700.103.el7.x86_64.rpm",
    "/root/powerflex_packages/PowerFlex_3.6.700.103_RHEL_OEL7/EMC-ScaleIO-sdt-3.6-700.103.el7.x86_64.rpm",
])
user = powerflex.User("user",
    name="NewUser",
    role="Monitor",
    password="Password123")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    powerflex:endpoint:
        value: 'TODO: var.endpoint'
    powerflex:insecure:
        value: true
    powerflex:password:
        value: 'TODO: var.password'
    powerflex:timeout:
        value: 120
    powerflex:username:
        value: 'TODO: var.username'

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Powerflex = Pulumi.Powerflex;

return await Deployment.RunAsync(() =>
{
    var pd = new Powerflex.ProtectionDomain("pd", new()
    {
        Name = "domain_1",
    });

    var sds1 = new Powerflex.Sds("sds1", new()
    {
        Name = "sds_1",
        ProtectionDomainId = pd.Id,
        IpLists = new[]
        {
            new Powerflex.Inputs.SdsIpListArgs
            {
                Ip = "10.10.10.1",
                Role = "all",
            },
        },
    }, new CustomResourceOptions
    {
        DependsOn =
        {
            pd,
        },
    });

    var sds2 = new Powerflex.Sds("sds2", new()
    {
        Name = "sds_2",
        ProtectionDomainId = pd.Id,
        IpLists = new[]
        {
            new Powerflex.Inputs.SdsIpListArgs
            {
                Ip = "10.10.10.2",
                Role = "all",
            },
        },
    }, new CustomResourceOptions
    {
        DependsOn =
        {
            pd,
        },
    });

    var sds3 = new Powerflex.Sds("sds3", new()
    {
        Name = "sds_3",
        ProtectionDomainId = pd.Id,
        IpLists = new[]
        {
            new Powerflex.Inputs.SdsIpListArgs
            {
                Ip = "10.10.10.3",
                Role = "all",
            },
        },
    }, new CustomResourceOptions
    {
        DependsOn =
        {
            pd,
        },
    });

    var sp = new Powerflex.StoragePool("sp", new()
    {
        Name = "SP",
        ProtectionDomainId = pd.Id,
        MediaType = "HDD",
        UseRmcache = true,
        UseRfcache = true,
    });

    var device1 = new Powerflex.Device("device1", new()
    {
        Name = "device1",
        DevicePath = "/dev/sdb",
        SdsId = sds1.Id,
        StoragePoolId = sp.Id,
        MediaType = "HDD",
        ExternalAccelerationType = "ReadAndWrite",
    }, new CustomResourceOptions
    {
        DependsOn =
        {
            sp,
        },
    });

    var device2 = new Powerflex.Device("device2", new()
    {
        Name = "device2",
        DevicePath = "/dev/sdb",
        SdsId = sds2.Id,
        StoragePoolId = sp.Id,
        MediaType = "HDD",
        ExternalAccelerationType = "ReadAndWrite",
    }, new CustomResourceOptions
    {
        DependsOn =
        {
            sp,
        },
    });

    var device3 = new Powerflex.Device("device3", new()
    {
        Name = "device3",
        DevicePath = "/dev/sdb",
        SdsId = sds3.Id,
        StoragePoolId = sp.Id,
        MediaType = "HDD",
        ExternalAccelerationType = "ReadAndWrite",
    }, new CustomResourceOptions
    {
        DependsOn =
        {
            sp,
        },
    });

    var volume = new Powerflex.Volume("volume", new()
    {
        Name = "volume1",
        ProtectionDomainId = pd.Id,
        StoragePoolId = sp.Id,
        Size = 16,
        VolumeType = "ThinProvisioned",
    }, new CustomResourceOptions
    {
        DependsOn =
        {
            device1,
            device2,
            device3,
        },
    });

    var map = new Powerflex.SdcVolumesMapping("map", new()
    {
        SdcVolumesMappingId = "e3d105e900000005",
        VolumeLists = new[]
        {
            new Powerflex.Inputs.SdcVolumesMappingVolumeListArgs
            {
                VolumeId = volume.Id,
                LimitIops = 140,
                LimitBwInMbps = 19,
                AccessMode = "ReadOnly",
            },
        },
    }, new CustomResourceOptions
    {
        DependsOn =
        {
            volume,
        },
    });

    var upload_test = new Powerflex.Package("upload-test", new()
    {
        FilePaths = new[]
        {
            "/root/powerflex_packages/PowerFlex_3.6.700.103_RHEL_OEL7/EMC-ScaleIO-lia-3.6-700.103.el7.x86_64.rpm",
            "/root/powerflex_packages/PowerFlex_3.6.700.103_RHEL_OEL7/EMC-ScaleIO-mdm-3.6-700.103.el7.x86_64.rpm",
            "/root/powerflex_packages/PowerFlex_3.6.700.103_RHEL_OEL7/EMC-ScaleIO-sds-3.6-700.103.el7.x86_64.rpm",
            "/root/powerflex_packages/PowerFlex_3.6.700.103_RHEL_OEL7/EMC-ScaleIO-sdc-3.6-700.103.el7.x86_64.rpm",
            "/root/powerflex_packages/PowerFlex_3.6.700.103_RHEL_OEL7/EMC-ScaleIO-sdr-3.6-700.103.el7.x86_64.rpm",
            "/root/powerflex_packages/PowerFlex_3.6.700.103_RHEL_OEL7/EMC-ScaleIO-sdt-3.6-700.103.el7.x86_64.rpm",
        },
    });

    var user = new Powerflex.User("user", new()
    {
        Name = "NewUser",
        Role = "Monitor",
        Password = "Password123",
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
    powerflex:endpoint:
        value: 'TODO: var.endpoint'
    powerflex:insecure:
        value: true
    powerflex:password:
        value: 'TODO: var.password'
    powerflex:timeout:
        value: 120
    powerflex:username:
        value: 'TODO: var.username'

```
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/powerflex/powerflex"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		pd, err := powerflex.NewProtectionDomain(ctx, "pd", &powerflex.ProtectionDomainArgs{
			Name: pulumi.String("domain_1"),
		})
		if err != nil {
			return err
		}
		sds1, err := powerflex.NewSds(ctx, "sds1", &powerflex.SdsArgs{
			Name:               pulumi.String("sds_1"),
			ProtectionDomainId: pd.ID(),
			IpLists: powerflex.SdsIpListArray{
				&powerflex.SdsIpListArgs{
					Ip:   pulumi.String("10.10.10.1"),
					Role: pulumi.String("all"),
				},
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			pd,
		}))
		if err != nil {
			return err
		}
		sds2, err := powerflex.NewSds(ctx, "sds2", &powerflex.SdsArgs{
			Name:               pulumi.String("sds_2"),
			ProtectionDomainId: pd.ID(),
			IpLists: powerflex.SdsIpListArray{
				&powerflex.SdsIpListArgs{
					Ip:   pulumi.String("10.10.10.2"),
					Role: pulumi.String("all"),
				},
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			pd,
		}))
		if err != nil {
			return err
		}
		sds3, err := powerflex.NewSds(ctx, "sds3", &powerflex.SdsArgs{
			Name:               pulumi.String("sds_3"),
			ProtectionDomainId: pd.ID(),
			IpLists: powerflex.SdsIpListArray{
				&powerflex.SdsIpListArgs{
					Ip:   pulumi.String("10.10.10.3"),
					Role: pulumi.String("all"),
				},
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			pd,
		}))
		if err != nil {
			return err
		}
		sp, err := powerflex.NewStoragePool(ctx, "sp", &powerflex.StoragePoolArgs{
			Name:               pulumi.String("SP"),
			ProtectionDomainId: pd.ID(),
			MediaType:          pulumi.String("HDD"),
			UseRmcache:         pulumi.Bool(true),
			UseRfcache:         pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		device1, err := powerflex.NewDevice(ctx, "device1", &powerflex.DeviceArgs{
			Name:                     pulumi.String("device1"),
			DevicePath:               pulumi.String("/dev/sdb"),
			SdsId:                    sds1.ID(),
			StoragePoolId:            sp.ID(),
			MediaType:                pulumi.String("HDD"),
			ExternalAccelerationType: pulumi.String("ReadAndWrite"),
		}, pulumi.DependsOn([]pulumi.Resource{
			sp,
		}))
		if err != nil {
			return err
		}
		device2, err := powerflex.NewDevice(ctx, "device2", &powerflex.DeviceArgs{
			Name:                     pulumi.String("device2"),
			DevicePath:               pulumi.String("/dev/sdb"),
			SdsId:                    sds2.ID(),
			StoragePoolId:            sp.ID(),
			MediaType:                pulumi.String("HDD"),
			ExternalAccelerationType: pulumi.String("ReadAndWrite"),
		}, pulumi.DependsOn([]pulumi.Resource{
			sp,
		}))
		if err != nil {
			return err
		}
		device3, err := powerflex.NewDevice(ctx, "device3", &powerflex.DeviceArgs{
			Name:                     pulumi.String("device3"),
			DevicePath:               pulumi.String("/dev/sdb"),
			SdsId:                    sds3.ID(),
			StoragePoolId:            sp.ID(),
			MediaType:                pulumi.String("HDD"),
			ExternalAccelerationType: pulumi.String("ReadAndWrite"),
		}, pulumi.DependsOn([]pulumi.Resource{
			sp,
		}))
		if err != nil {
			return err
		}
		volume, err := powerflex.NewVolume(ctx, "volume", &powerflex.VolumeArgs{
			Name:               pulumi.String("volume1"),
			ProtectionDomainId: pd.ID(),
			StoragePoolId:      sp.ID(),
			Size:               pulumi.Float64(16),
			VolumeType:         pulumi.String("ThinProvisioned"),
		}, pulumi.DependsOn([]pulumi.Resource{
			device1,
			device2,
			device3,
		}))
		if err != nil {
			return err
		}
		_, err = powerflex.NewSdcVolumesMapping(ctx, "map", &powerflex.SdcVolumesMappingArgs{
			SdcVolumesMappingId: pulumi.String("e3d105e900000005"),
			VolumeLists: powerflex.SdcVolumesMappingVolumeListArray{
				&powerflex.SdcVolumesMappingVolumeListArgs{
					VolumeId:      volume.ID(),
					LimitIops:     pulumi.Float64(140),
					LimitBwInMbps: pulumi.Float64(19),
					AccessMode:    pulumi.String("ReadOnly"),
				},
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			volume,
		}))
		if err != nil {
			return err
		}
		_, err = powerflex.NewPackage(ctx, "upload-test", &powerflex.PackageArgs{
			FilePaths: pulumi.StringArray{
				pulumi.String("/root/powerflex_packages/PowerFlex_3.6.700.103_RHEL_OEL7/EMC-ScaleIO-lia-3.6-700.103.el7.x86_64.rpm"),
				pulumi.String("/root/powerflex_packages/PowerFlex_3.6.700.103_RHEL_OEL7/EMC-ScaleIO-mdm-3.6-700.103.el7.x86_64.rpm"),
				pulumi.String("/root/powerflex_packages/PowerFlex_3.6.700.103_RHEL_OEL7/EMC-ScaleIO-sds-3.6-700.103.el7.x86_64.rpm"),
				pulumi.String("/root/powerflex_packages/PowerFlex_3.6.700.103_RHEL_OEL7/EMC-ScaleIO-sdc-3.6-700.103.el7.x86_64.rpm"),
				pulumi.String("/root/powerflex_packages/PowerFlex_3.6.700.103_RHEL_OEL7/EMC-ScaleIO-sdr-3.6-700.103.el7.x86_64.rpm"),
				pulumi.String("/root/powerflex_packages/PowerFlex_3.6.700.103_RHEL_OEL7/EMC-ScaleIO-sdt-3.6-700.103.el7.x86_64.rpm"),
			},
		})
		if err != nil {
			return err
		}
		_, err = powerflex.NewUser(ctx, "user", &powerflex.UserArgs{
			Name:     pulumi.String("NewUser"),
			Role:     pulumi.String("Monitor"),
			Password: pulumi.String("Password123"),
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
    powerflex:endpoint:
        value: 'TODO: var.endpoint'
    powerflex:insecure:
        value: true
    powerflex:password:
        value: 'TODO: var.password'
    powerflex:timeout:
        value: 120
    powerflex:username:
        value: 'TODO: var.username'

```
```yaml
resources:
  pd:
    type: powerflex:ProtectionDomain
    properties:
      name: domain_1
  sds1:
    type: powerflex:Sds
    properties:
      name: sds_1
      protectionDomainId: ${pd.id}
      ipLists:
        - ip: 10.10.10.1
          role: all
    options:
      dependsOn:
        - ${pd}
  sds2:
    type: powerflex:Sds
    properties:
      name: sds_2
      protectionDomainId: ${pd.id}
      ipLists:
        - ip: 10.10.10.2
          role: all
    options:
      dependsOn:
        - ${pd}
  sds3:
    type: powerflex:Sds
    properties:
      name: sds_3
      protectionDomainId: ${pd.id}
      ipLists:
        - ip: 10.10.10.3
          role: all
    options:
      dependsOn:
        - ${pd}
  sp:
    type: powerflex:StoragePool
    properties:
      name: SP
      protectionDomainId: ${pd.id}
      mediaType: HDD
      useRmcache: true
      useRfcache: true
  device1:
    type: powerflex:Device
    properties:
      name: device1
      devicePath: /dev/sdb
      sdsId: ${sds1.id}
      storagePoolId: ${sp.id}
      mediaType: HDD
      externalAccelerationType: ReadAndWrite
    options:
      dependsOn:
        - ${sp}
  device2:
    type: powerflex:Device
    properties:
      name: device2
      devicePath: /dev/sdb
      sdsId: ${sds2.id}
      storagePoolId: ${sp.id}
      mediaType: HDD
      externalAccelerationType: ReadAndWrite
    options:
      dependsOn:
        - ${sp}
  device3:
    type: powerflex:Device
    properties:
      name: device3
      devicePath: /dev/sdb
      sdsId: ${sds3.id}
      storagePoolId: ${sp.id}
      mediaType: HDD
      externalAccelerationType: ReadAndWrite
    options:
      dependsOn:
        - ${sp}
  volume:
    type: powerflex:Volume
    properties:
      name: volume1
      protectionDomainId: ${pd.id}
      storagePoolId: ${sp.id}
      size: 16
      volumeType: ThinProvisioned
    options:
      dependsOn:
        - ${device1}
        - ${device2}
        - ${device3}
  map:
    type: powerflex:SdcVolumesMapping
    properties:
      sdcVolumesMappingId: e3d105e900000005
      volumeLists:
        - volumeId: ${volume.id}
          limitIops: 140
          limitBwInMbps: 19
          accessMode: ReadOnly
    options:
      dependsOn:
        - ${volume}
  upload-test:
    type: powerflex:Package
    properties:
      filePaths:
        - /root/powerflex_packages/PowerFlex_3.6.700.103_RHEL_OEL7/EMC-ScaleIO-lia-3.6-700.103.el7.x86_64.rpm
        - /root/powerflex_packages/PowerFlex_3.6.700.103_RHEL_OEL7/EMC-ScaleIO-mdm-3.6-700.103.el7.x86_64.rpm
        - /root/powerflex_packages/PowerFlex_3.6.700.103_RHEL_OEL7/EMC-ScaleIO-sds-3.6-700.103.el7.x86_64.rpm
        - /root/powerflex_packages/PowerFlex_3.6.700.103_RHEL_OEL7/EMC-ScaleIO-sdc-3.6-700.103.el7.x86_64.rpm
        - /root/powerflex_packages/PowerFlex_3.6.700.103_RHEL_OEL7/EMC-ScaleIO-sdr-3.6-700.103.el7.x86_64.rpm
        - /root/powerflex_packages/PowerFlex_3.6.700.103_RHEL_OEL7/EMC-ScaleIO-sdt-3.6-700.103.el7.x86_64.rpm
  user:
    type: powerflex:User
    properties:
      name: NewUser
      role: Monitor
      password: Password123
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    powerflex:endpoint:
        value: 'TODO: var.endpoint'
    powerflex:insecure:
        value: true
    powerflex:password:
        value: 'TODO: var.password'
    powerflex:timeout:
        value: 120
    powerflex:username:
        value: 'TODO: var.username'

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.powerflex.ProtectionDomain;
import com.pulumi.powerflex.ProtectionDomainArgs;
import com.pulumi.powerflex.Sds;
import com.pulumi.powerflex.SdsArgs;
import com.pulumi.powerflex.inputs.SdsIpListArgs;
import com.pulumi.powerflex.StoragePool;
import com.pulumi.powerflex.StoragePoolArgs;
import com.pulumi.powerflex.Device;
import com.pulumi.powerflex.DeviceArgs;
import com.pulumi.powerflex.Volume;
import com.pulumi.powerflex.VolumeArgs;
import com.pulumi.powerflex.SdcVolumesMapping;
import com.pulumi.powerflex.SdcVolumesMappingArgs;
import com.pulumi.powerflex.inputs.SdcVolumesMappingVolumeListArgs;
import com.pulumi.powerflex.Package;
import com.pulumi.powerflex.PackageArgs;
import com.pulumi.powerflex.User;
import com.pulumi.powerflex.UserArgs;
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
        var pd = new ProtectionDomain("pd", ProtectionDomainArgs.builder()
            .name("domain_1")
            .build());

        var sds1 = new Sds("sds1", SdsArgs.builder()
            .name("sds_1")
            .protectionDomainId(pd.id())
            .ipLists(SdsIpListArgs.builder()
                .ip("10.10.10.1")
                .role("all")
                .build())
            .build(), CustomResourceOptions.builder()
                .dependsOn(pd)
                .build());

        var sds2 = new Sds("sds2", SdsArgs.builder()
            .name("sds_2")
            .protectionDomainId(pd.id())
            .ipLists(SdsIpListArgs.builder()
                .ip("10.10.10.2")
                .role("all")
                .build())
            .build(), CustomResourceOptions.builder()
                .dependsOn(pd)
                .build());

        var sds3 = new Sds("sds3", SdsArgs.builder()
            .name("sds_3")
            .protectionDomainId(pd.id())
            .ipLists(SdsIpListArgs.builder()
                .ip("10.10.10.3")
                .role("all")
                .build())
            .build(), CustomResourceOptions.builder()
                .dependsOn(pd)
                .build());

        var sp = new StoragePool("sp", StoragePoolArgs.builder()
            .name("SP")
            .protectionDomainId(pd.id())
            .mediaType("HDD")
            .useRmcache(true)
            .useRfcache(true)
            .build());

        var device1 = new Device("device1", DeviceArgs.builder()
            .name("device1")
            .devicePath("/dev/sdb")
            .sdsId(sds1.id())
            .storagePoolId(sp.id())
            .mediaType("HDD")
            .externalAccelerationType("ReadAndWrite")
            .build(), CustomResourceOptions.builder()
                .dependsOn(sp)
                .build());

        var device2 = new Device("device2", DeviceArgs.builder()
            .name("device2")
            .devicePath("/dev/sdb")
            .sdsId(sds2.id())
            .storagePoolId(sp.id())
            .mediaType("HDD")
            .externalAccelerationType("ReadAndWrite")
            .build(), CustomResourceOptions.builder()
                .dependsOn(sp)
                .build());

        var device3 = new Device("device3", DeviceArgs.builder()
            .name("device3")
            .devicePath("/dev/sdb")
            .sdsId(sds3.id())
            .storagePoolId(sp.id())
            .mediaType("HDD")
            .externalAccelerationType("ReadAndWrite")
            .build(), CustomResourceOptions.builder()
                .dependsOn(sp)
                .build());

        var volume = new Volume("volume", VolumeArgs.builder()
            .name("volume1")
            .protectionDomainId(pd.id())
            .storagePoolId(sp.id())
            .size(16)
            .volumeType("ThinProvisioned")
            .build(), CustomResourceOptions.builder()
                .dependsOn(
                    device1,
                    device2,
                    device3)
                .build());

        var map = new SdcVolumesMapping("map", SdcVolumesMappingArgs.builder()
            .sdcVolumesMappingId("e3d105e900000005")
            .volumeLists(SdcVolumesMappingVolumeListArgs.builder()
                .volumeId(volume.id())
                .limitIops(140)
                .limitBwInMbps(19)
                .accessMode("ReadOnly")
                .build())
            .build(), CustomResourceOptions.builder()
                .dependsOn(volume)
                .build());

        var upload_test = new Package("upload-test", PackageArgs.builder()
            .filePaths(
                "/root/powerflex_packages/PowerFlex_3.6.700.103_RHEL_OEL7/EMC-ScaleIO-lia-3.6-700.103.el7.x86_64.rpm",
                "/root/powerflex_packages/PowerFlex_3.6.700.103_RHEL_OEL7/EMC-ScaleIO-mdm-3.6-700.103.el7.x86_64.rpm",
                "/root/powerflex_packages/PowerFlex_3.6.700.103_RHEL_OEL7/EMC-ScaleIO-sds-3.6-700.103.el7.x86_64.rpm",
                "/root/powerflex_packages/PowerFlex_3.6.700.103_RHEL_OEL7/EMC-ScaleIO-sdc-3.6-700.103.el7.x86_64.rpm",
                "/root/powerflex_packages/PowerFlex_3.6.700.103_RHEL_OEL7/EMC-ScaleIO-sdr-3.6-700.103.el7.x86_64.rpm",
                "/root/powerflex_packages/PowerFlex_3.6.700.103_RHEL_OEL7/EMC-ScaleIO-sdt-3.6-700.103.el7.x86_64.rpm")
            .build());

        var user = new User("user", UserArgs.builder()
            .name("NewUser")
            .role("Monitor")
            .password("Password123")
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}

variables.tf
{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";

const config = new pulumi.Config();
// Stores the username of PowerFlex host.
const username = config.require("username");
// Stores the password of PowerFlex host.
const password = config.require("password");
// Stores the endpoint of PowerFlex host. eg: https://10.1.1.1:443, here 443 is port where API requests are getting accepted
const endpoint = config.require("endpoint");
```
{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi

config = pulumi.Config()
# Stores the username of PowerFlex host.
username = config.require("username")
# Stores the password of PowerFlex host.
password = config.require("password")
# Stores the endpoint of PowerFlex host. eg: https://10.1.1.1:443, here 443 is port where API requests are getting accepted
endpoint = config.require("endpoint")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;

return await Deployment.RunAsync(() =>
{
    var config = new Config();
    // Stores the username of PowerFlex host.
    var username = config.Require("username");
    // Stores the password of PowerFlex host.
    var password = config.Require("password");
    // Stores the endpoint of PowerFlex host. eg: https://10.1.1.1:443, here 443 is port where API requests are getting accepted
    var endpoint = config.Require("endpoint");
});

```
{{% /choosable %}}
{{% choosable language go %}}
```go
package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")
		// Stores the username of PowerFlex host.
		username := cfg.Require("username")
		// Stores the password of PowerFlex host.
		password := cfg.Require("password")
		// Stores the endpoint of PowerFlex host. eg: https://10.1.1.1:443, here 443 is port where API requests are getting accepted
		endpoint := cfg.Require("endpoint")
		return nil
	})
}
```
{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
configuration:
  # /*
  # Copyright (c) 2023-2024 Dell Inc., or its subsidiaries. All Rights Reserved.

  # Licensed under the Mozilla Public License Version 2.0 (the "License");
  # you may not use this file except in compliance with the License.
  # You may obtain a copy of the License at

  #     http://mozilla.org/MPL/2.0/


  # Unless required by applicable law or agreed to in writing, software
  # distributed under the License is distributed on an "AS IS" BASIS,
  # WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  # See the License for the specific language governing permissions and
  # limitations under the License.
  # */
  username:
    type: string
  password:
    type: string
  endpoint:
    type: string
```
{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
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
        final var username = config.get("username");
        final var password = config.get("password");
        final var endpoint = config.get("endpoint");
    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference
### Required

- `endpoint` (String) The PowerFlex Gateway server URL (inclusive of the port).
- `password` (String, Sensitive) The password required for the authentication.
- `username` (String) The username required for authentication.

- `insecure` (Boolean) Specifies if the user wants to skip SSL verification.
- `timeout` (Number) HTTPS timeout.