---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/ionos-cloud/ionoscloud/6.7.7/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Ionoscloud Provider
meta_desc: Provides an overview on how to configure the Pulumi Ionoscloud provider.
layout: package
---

## Generate Provider

The Ionoscloud provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider ionos-cloud/ionoscloud
```
## Overview

The IonosCloud provider gives the ability to deploy and configure resources using the IonosCloud APIs.

Use the navigation to the left to read about the available functions and resources.
## Using the provider

The provider needs to be configured with proper credentials before it can be used.

You can use token authentication. We strongly suggest to use token authentication for security purposes. Details on how to generate your token [here](https://github.com/ionos-cloud/sdk-go#token-authentication):

```bash
export IONOS_TOKEN="token"
```

You can set the environment variables for HTTP basic authentication:

```bash
export IONOS_USERNAME="username"
export IONOS_PASSWORD="password"
```
For managing IONOS Object Storage resources you need to set the following environment variables with your credentials
```bash
export IONOS_S3_ACCESS_KEY="accesskey"
export IONOS_S3_SECRET_KEY="secretkey"
```

Also, you can overwrite the api endpoint: `api.ionos.com` via the following environment variable:

```bash
export IONOS_API_URL="api-url"
```

Note: if `IONOS_API_URL` environment variable is set, it is recommended to have the `api.ionos.com` value (not `api.ionos.com/cloudapi/v5`).

Another way of configuring it, is by providing your credentials/api_url in a `.tf` configuration file in the provider configuration as shown in the below example.

⚠️ **Note:** It's NOT usually necessary to set `endpoint` field. The SDKs the pulumi uses know how to route requests to the correct endpoints in the API.

You can either explicitly write them in the .tf file or use var.name as in the example above. For setting the var.name, environment variables can be used. The environment variables must be in the format TF_VAR_name and this will be checked last for a value. For example:

```bash
export TF_VAR_ionos_token="token"
#export TF_VAR_ionos_username="username"
#export TF_VAR_ionos_password="password"
#export TF_VAR_ionos_s3_access_key="accesskey"
#export TF_VAR_ionos_s3_secret_key="secretkey"
#export TF_VAR_ionos_s3_region="region"
```
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    ionoscloud:token:
        value: 'TODO: var.ionos_token'

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as ionoscloud from "@pulumi/ionoscloud";

const main = new ionoscloud.Datacenter("main", {});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    ionoscloud:token:
        value: 'TODO: var.ionos_token'

```
```python
import pulumi
import pulumi_ionoscloud as ionoscloud

main = ionoscloud.Datacenter("main")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    ionoscloud:token:
        value: 'TODO: var.ionos_token'

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Ionoscloud = Pulumi.Ionoscloud;

return await Deployment.RunAsync(() =>
{
    var main = new Ionoscloud.Datacenter("main");

});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    ionoscloud:token:
        value: 'TODO: var.ionos_token'

```
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/ionoscloud/v6/ionoscloud"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ionoscloud.NewDatacenter(ctx, "main", nil)
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
    ionoscloud:token:
        value: 'TODO: var.ionos_token'

```
```yaml
resources:
  main:
    type: ionoscloud:Datacenter
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    ionoscloud:token:
        value: 'TODO: var.ionos_token'

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.ionoscloud.Datacenter;
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
        var main = new Datacenter("main");

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
### Important Notes
* The `requiredProviders` section needs to be specified in order for pulumi to be
  able to find and download the ionoscloud provider
* The credentials provided in a `.tf` file will override the credentials from environment variables.
### Configuration Reference

The following configuration inputs are supported:

- `token`    - Required if username and password are not set. If omitted, the `IONOS_TOKEN` environment variable is used.

- `username` - Required if token is not set. If omitted, the `IONOS_USERNAME` environment variable is used. The username is generally an e-mail address in 'username@domain.tld' format.

- `password` - Required if token is not set. If omitted, the `IONOS_PASSWORD` environment variable is used.

- `endpoint` - (Optional) Usually not necessary to be set, SDks know internally how to route requests to the API. If omitted, the `IONOS_API_URL` environment variable is used, or it defaults to the current Cloud API release.

- `retries` - (Deprecated) Number of retries while waiting for a resource to be provisioned. Default value is 50. **Note**: This argument has been deprecated and replaced by the implementation of resource timeouts described below.

- `contractNumber` - "To be set only for reseller accounts. Allows to run pulumi on a contract number under a reseller account.",

- `s3AccessKey` - Required for managing IONOS Object Storage resources.

- `s3SecretKey` - Required for managing IONOS Object Storage resources.
- `s3Region` - Optional, defines the region of the Object Storage resource.
### Environment Variables

|  Environment Variable   |                                                                                Description                                                                                |
|-------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `IONOS_USERNAME`        | Specify the username used to login, to authenticate against the IONOS Cloud API                                                                                           |
| `IONOS_PASSWORD`        | Specify the password used to login, to authenticate against the IONOS Cloud API                                                                                           |
| `IONOS_TOKEN`           | Specify the token used to login, if a token is being used instead of username and password                                                                                |
| `IONOS_API_URL`         | Specify the API URL. It will overwrite the API endpoint default value `api.ionos.com`.  It is not necessary to override this value unless you have special routing config |
| `IONOS_LOG_LEVEL`       | Specify the Log Level used to log messages. Possible values: Off, Debug, Trace                                                                                            |
| `IONOS_PINNED_CERT`     | Specify the SHA-256 public fingerprint here, enables certificate pinning                                                                                                  |
| `IONOS_CONTRACT_NUMBER` | Specify the contract number on which you wish to provision. Only valid for reseller accounts, for other types of accounts the header will be ignored                      |
| `IONOS_S3_ACCESS_KEY`   | Specify the access key used to authenticate against the IONOS Object Storage API                                                                                          |
| `IONOS_S3_SECRET_KEY`   | Specify the secret key used to authenticate against the IONOS Object Storage API                                                                                          |
| `IONOS_S3_REGION`       | Region for IONOS Object Storage operations. Default value: eu-central-3. **If you use IONOS_API_URL_OBJECT_STORAGE, `IONOS_S3_REGION` is mandatory**                      |
## Resource Timeout

Individual resources may provide a `timeouts` block to configure the amount of time a specific operation is allowed to take before being considered an error. Each resource may provide configurable timeouts for the `create`, `update`, and `delete` operations. Each resource that supports timeouts will have or inherit default values for that operation.
Users can overwrite the default values for a specific resource in the configuration.

The default `timeouts` values are:

- create - (Default 60 minutes) Used for creating a resource.
- update - (Default 60 minutes) Used for updating a resource .
- delete - (Default 60 minutes) Used for destroying a resource.
- default - (Default 60 minutes) Used for every other action on a resource.

An example of overwriting the `create`, `update`, and `delete` timeouts:
{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as ionoscloud from "@pulumi/ionoscloud";
import * as random from "@pulumi/random";

const serverImagePassword = new random.index.Password("server_image_password", {
    length: 16,
    special: false,
});
const example = new ionoscloud.Server("example", {
    name: "Server Example",
    datacenterId: exampleIonoscloudDatacenter.id,
    cores: 1,
    ram: 1024,
    imageName: exampleIonoscloudImage.id,
    imagePassword: serverImagePassword.result,
    type: "ENTERPRISE",
    volume: {
        name: "system",
        size: 5,
        diskType: "SSD Standard",
        userData: "foo",
        bus: "VIRTIO",
    },
    nic: {
        lan: exampleIonoscloudLan.id,
        name: "system",
        dhcp: true,
        firewallActive: true,
        firewallType: "BIDIRECTIONAL",
        ips: [
            exampleIonoscloudIpblock.ips[0],
            exampleIonoscloudIpblock.ips[1],
        ],
        firewalls: [{
            protocol: "TCP",
            name: "SSH",
            portRangeStart: 22,
            portRangeEnd: 22,
            sourceMac: "00:0a:95:9d:68:17",
            sourceIp: exampleIonoscloudIpblock.ips[2],
            targetIp: exampleIonoscloudIpblock.ips[3],
            type: "EGRESS",
        }],
    },
});
```
{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import pulumi_ionoscloud as ionoscloud
import pulumi_random as random

server_image_password = random.index.Password("server_image_password",
    length=16,
    special=False)
example = ionoscloud.Server("example",
    name="Server Example",
    datacenter_id=example_ionoscloud_datacenter["id"],
    cores=1,
    ram=1024,
    image_name=example_ionoscloud_image["id"],
    image_password=server_image_password["result"],
    type="ENTERPRISE",
    volume={
        "name": "system",
        "size": 5,
        "disk_type": "SSD Standard",
        "user_data": "foo",
        "bus": "VIRTIO",
    },
    nic={
        "lan": example_ionoscloud_lan["id"],
        "name": "system",
        "dhcp": True,
        "firewall_active": True,
        "firewall_type": "BIDIRECTIONAL",
        "ips": [
            example_ionoscloud_ipblock["ips"],
            example_ionoscloud_ipblock["ips"],
        ],
        "firewalls": [{
            "protocol": "TCP",
            "name": "SSH",
            "port_range_start": 22,
            "port_range_end": 22,
            "source_mac": "00:0a:95:9d:68:17",
            "source_ip": example_ionoscloud_ipblock["ips"],
            "target_ip": example_ionoscloud_ipblock["ips"],
            "type": "EGRESS",
        }],
    })
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Ionoscloud = Pulumi.Ionoscloud;
using Random = Pulumi.Random;

return await Deployment.RunAsync(() =>
{
    var serverImagePassword = new Random.Index.Password("server_image_password", new()
    {
        Length = 16,
        Special = false,
    });

    var example = new Ionoscloud.Server("example", new()
    {
        Name = "Server Example",
        DatacenterId = exampleIonoscloudDatacenter.Id,
        Cores = 1,
        Ram = 1024,
        ImageName = exampleIonoscloudImage.Id,
        ImagePassword = serverImagePassword.Result,
        Type = "ENTERPRISE",
        Volume = new Ionoscloud.Inputs.ServerVolumeArgs
        {
            Name = "system",
            Size = 5,
            DiskType = "SSD Standard",
            UserData = "foo",
            Bus = "VIRTIO",
        },
        Nic = new Ionoscloud.Inputs.ServerNicArgs
        {
            Lan = exampleIonoscloudLan.Id,
            Name = "system",
            Dhcp = true,
            FirewallActive = true,
            FirewallType = "BIDIRECTIONAL",
            Ips = new[]
            {
                exampleIonoscloudIpblock.Ips[0],
                exampleIonoscloudIpblock.Ips[1],
            },
            Firewalls = new[]
            {
                new Ionoscloud.Inputs.ServerNicFirewallArgs
                {
                    Protocol = "TCP",
                    Name = "SSH",
                    PortRangeStart = 22,
                    PortRangeEnd = 22,
                    SourceMac = "00:0a:95:9d:68:17",
                    SourceIp = exampleIonoscloudIpblock.Ips[2],
                    TargetIp = exampleIonoscloudIpblock.Ips[3],
                    Type = "EGRESS",
                },
            },
        },
    });

});

```
{{% /choosable %}}
{{% choosable language go %}}
```go
package main

import (
	"github.com/pulumi/pulumi-random/sdk/go/random"
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/ionoscloud/v6/ionoscloud"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		serverImagePassword, err := random.NewPassword(ctx, "server_image_password", &random.PasswordArgs{
			Length:  16,
			Special: false,
		})
		if err != nil {
			return err
		}
		_, err = ionoscloud.NewServer(ctx, "example", &ionoscloud.ServerArgs{
			Name:          pulumi.String("Server Example"),
			DatacenterId:  pulumi.Any(exampleIonoscloudDatacenter.Id),
			Cores:         pulumi.Float64(1),
			Ram:           pulumi.Float64(1024),
			ImageName:     pulumi.Any(exampleIonoscloudImage.Id),
			ImagePassword: serverImagePassword.Result,
			Type:          pulumi.String("ENTERPRISE"),
			Volume: &ionoscloud.ServerVolumeArgs{
				Name:     pulumi.String("system"),
				Size:     pulumi.Float64(5),
				DiskType: pulumi.String("SSD Standard"),
				UserData: pulumi.String("foo"),
				Bus:      pulumi.String("VIRTIO"),
			},
			Nic: &ionoscloud.ServerNicArgs{
				Lan:            pulumi.Any(exampleIonoscloudLan.Id),
				Name:           pulumi.String("system"),
				Dhcp:           pulumi.Bool(true),
				FirewallActive: pulumi.Bool(true),
				FirewallType:   pulumi.String("BIDIRECTIONAL"),
				Ips: pulumi.StringArray{
					exampleIonoscloudIpblock.Ips[0],
					exampleIonoscloudIpblock.Ips[1],
				},
				Firewalls: ionoscloud.ServerNicFirewallArray{
					&ionoscloud.ServerNicFirewallArgs{
						Protocol:       pulumi.String("TCP"),
						Name:           pulumi.String("SSH"),
						PortRangeStart: pulumi.Float64(22),
						PortRangeEnd:   pulumi.Float64(22),
						SourceMac:      pulumi.String("00:0a:95:9d:68:17"),
						SourceIp:       pulumi.Any(exampleIonoscloudIpblock.Ips[2]),
						TargetIp:       pulumi.Any(exampleIonoscloudIpblock.Ips[3]),
						Type:           pulumi.String("EGRESS"),
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
{{% choosable language yaml %}}
```yaml
resources:
  example:
    type: ionoscloud:Server
    properties:
      name: Server Example
      datacenterId: ${exampleIonoscloudDatacenter.id}
      cores: 1
      ram: 1024
      imageName: ${exampleIonoscloudImage.id}
      imagePassword: ${serverImagePassword.result}
      type: ENTERPRISE
      volume:
        name: system
        size: 5
        diskType: SSD Standard
        userData: foo
        bus: VIRTIO
      nic:
        lan: ${exampleIonoscloudLan.id}
        name: system
        dhcp: true
        firewallActive: true
        firewallType: BIDIRECTIONAL
        ips:
          - ${exampleIonoscloudIpblock.ips[0]}
          - ${exampleIonoscloudIpblock.ips[1]}
        firewalls:
          - protocol: TCP
            name: SSH
            portRangeStart: 22
            portRangeEnd: 22
            sourceMac: 00:0a:95:9d:68:17
            sourceIp: ${exampleIonoscloudIpblock.ips[2]}
            targetIp: ${exampleIonoscloudIpblock.ips[3]}
            type: EGRESS
  serverImagePassword:
    type: random:Password
    name: server_image_password
    properties:
      length: 16
      special: false
```
{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.random.Password;
import com.pulumi.random.PasswordArgs;
import com.pulumi.ionoscloud.Server;
import com.pulumi.ionoscloud.ServerArgs;
import com.pulumi.ionoscloud.inputs.ServerVolumeArgs;
import com.pulumi.ionoscloud.inputs.ServerNicArgs;
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
        var serverImagePassword = new Password("serverImagePassword", PasswordArgs.builder()
            .length(16)
            .special(false)
            .build());

        var example = new Server("example", ServerArgs.builder()
            .name("Server Example")
            .datacenterId(exampleIonoscloudDatacenter.id())
            .cores(1)
            .ram(1024)
            .imageName(exampleIonoscloudImage.id())
            .imagePassword(serverImagePassword.result())
            .type("ENTERPRISE")
            .volume(ServerVolumeArgs.builder()
                .name("system")
                .size(5)
                .diskType("SSD Standard")
                .userData("foo")
                .bus("VIRTIO")
                .build())
            .nic(ServerNicArgs.builder()
                .lan(exampleIonoscloudLan.id())
                .name("system")
                .dhcp(true)
                .firewallActive(true)
                .firewallType("BIDIRECTIONAL")
                .ips(
                    exampleIonoscloudIpblock.ips()[0],
                    exampleIonoscloudIpblock.ips()[1])
                .firewalls(ServerNicFirewallArgs.builder()
                    .protocol("TCP")
                    .name("SSH")
                    .portRangeStart(22)
                    .portRangeEnd(22)
                    .sourceMac("00:0a:95:9d:68:17")
                    .sourceIp(exampleIonoscloudIpblock.ips()[2])
                    .targetIp(exampleIonoscloudIpblock.ips()[3])
                    .type("EGRESS")
                    .build())
                .build())
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}

Valid units of time should be expressed in "s", "m", "h" for "seconds", "minutes", and "hours" respectively.

Individual resources must opt-in to providing configurable `timeouts`, and attempting to configure values for a resource that does not support `timeouts`, or overwriting a specific action that the resource does not specify as an option, will result in an error.

> **Note:** Pulumi does not automatically rollback in the face of errors.
Instead, your Pulumi state file will be partially updated with
any resources that successfully completed.
## Migrating from the ProfitBricks provider

This area is not necessary unless you were using the old `profitbricks` provider
### Provider Name in HCL files

The provider name changed from `profitbricks` to `ionoscloud`.
This reflects in the following change in your pulumi hcl files:
`provider "profitbricks"` becomes `provider "ionoscloud"`
### Environment Variables

The following env variables have changed:

|   Old Variable Name   | New Variable Name |
|-----------------------|-------------------|
| PROFITBRICKS_USERNAME | IONOS_USERNAME    |
| PROFITBRICKS_PASSWORD | IONOS_PASSWORD    |
| PROFITBRICKS_TOKEN    | IONOS_TOKEN       |
| PROFITBRICKS_API_URL  | IONOS_API_URL     |
### Resources and Functions in HCL files

The migration affects resource names and function names.
Every resource and function changed its prefix from `profitbricks_` to `ionoscloud_`.

In order to accommodate that, the pulumi hcl files must be updated.

This can be done with a simple find and replace procedure.
For example, on Linux, sed can be used:
```bash
$ sed -i 's/profitbricks_/ionoscloud_/g' ./main.tf
```

On OSX the same command becomes:
```bash
$ sed -i bak 's/profitbricks_/ionoscloud_/g' ./main.tf
```
### Pulumi State

Because of the name changes of resources and functions, the pulumi state must also be updated.
The local state, in json format, can be updated by replacing `profitbricks_` with `ionoscloud_` directly in the state file.
For example, on Linux, using:
```bash
$ sed -i 's/profitbricks_/ionoscloud_/g' ./pulumi.tfstate
```

On OSX the same command becomes:
```bash
$ sed -i bak 's/profitbricks_/ionoscloud_/g' ./pulumi.tfstate
```

The `provider` entries must also be updated. For example:
```
"provider": "provider[\"registry.pulumi.io/pulumi/profitbricks\"]"
```
becomes
```
"provider": "provider[\"registry.pulumi.io/pulumi/ionoscloud\"]"
```

If you manage your state using remote backends you need to take the appropriate steps specific to your backend.