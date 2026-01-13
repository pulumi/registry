---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/vmware/avi/31.2.1/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Avi Provider
meta_desc: Provides an overview on how to configure the Pulumi Avi provider.
layout: package
---

## Generate Provider

The Avi provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider vmware/avi
```
## Overview

The AVI provider is used to interact with [AVI Controller](https://avinetworks.com/).
The provider needs
to be configured with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.
# Integrating Pulumi with Avi Vantage
Each Avi Vantage REST resource is exposed as a resource in Pulumi. For instance, you can setup a pool by using avi.Pool as the resource type in Pulumi. There are more than 50 different REST resources supported.

Each Pulumi resource maps directly to a corresponding Avi Vantage API as defined in the Swagger Spec (available on any Avi Vantage Controller at https://<controller>/swagger/) or the API documentation (available [here](https://avinetworks.com/docs/latest/api-guide/)). There is a 1:1 mapping between the fields in the Pulumi schema and the corresponding Avi Vantage object definition.

Avi Pulumi provider is a native integration into the Pulumi to setup all Avi configuration that is part of Avi REST API. With the Avi Pulumi provider, a pulumi preview written in HCL syntax can specify any Avi configuration and it will be reflected on the Avi Controller. For instance, you can use a pulumi preview with the following code to setup a pool in Avi Vantage.

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as avi from "@pulumi/avi";

const testpool = new avi.Pool("testpool", {
    name: "pool-42",
    healthMonitorRefs: [testHm1.id],
    tenantRef: defaultTenant.id,
    cloudRef: defaultCloud.id,
    servers: [{
        ips: [{
            type: "V4",
            addr: "x.x.x.x",
        }],
        port: "8080",
    }],
    failActions: [{
        type: "FAIL_ACTION_CLOSE_CONN",
    }],
});
```
{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import pulumi_avi as avi

testpool = avi.Pool("testpool",
    name="pool-42",
    health_monitor_refs=[test_hm1["id"]],
    tenant_ref=default_tenant["id"],
    cloud_ref=default_cloud["id"],
    servers=[{
        "ips": [{
            "type": "V4",
            "addr": "x.x.x.x",
        }],
        "port": "8080",
    }],
    fail_actions=[{
        "type": "FAIL_ACTION_CLOSE_CONN",
    }])
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Avi = Pulumi.Avi;

return await Deployment.RunAsync(() =>
{
    var testpool = new Avi.Pool("testpool", new()
    {
        Name = "pool-42",
        HealthMonitorRefs = new[]
        {
            testHm1.Id,
        },
        TenantRef = defaultTenant.Id,
        CloudRef = defaultCloud.Id,
        Servers = new[]
        {
            new Avi.Inputs.PoolServerArgs
            {
                Ips = new[]
                {
                    new Avi.Inputs.PoolServerIpArgs
                    {
                        Type = "V4",
                        Addr = "x.x.x.x",
                    },
                },
                Port = "8080",
            },
        },
        FailActions = new[]
        {
            new Avi.Inputs.PoolFailActionArgs
            {
                Type = "FAIL_ACTION_CLOSE_CONN",
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
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/avi/v31/avi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := avi.NewPool(ctx, "testpool", &avi.PoolArgs{
			Name: pulumi.String("pool-42"),
			HealthMonitorRefs: pulumi.StringArray{
				testHm1.Id,
			},
			TenantRef: pulumi.Any(defaultTenant.Id),
			CloudRef:  pulumi.Any(defaultCloud.Id),
			Servers: avi.PoolServerArray{
				&avi.PoolServerArgs{
					Ips: avi.PoolServerIpArray{
						&avi.PoolServerIpArgs{
							Type: pulumi.String("V4"),
							Addr: pulumi.String("x.x.x.x"),
						},
					},
					Port: pulumi.String("8080"),
				},
			},
			FailActions: avi.PoolFailActionArray{
				&avi.PoolFailActionArgs{
					Type: pulumi.String("FAIL_ACTION_CLOSE_CONN"),
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
  testpool:
    type: avi:Pool
    properties:
      name: pool-42
      healthMonitorRefs:
        - ${testHm1.id}
      tenantRef: ${defaultTenant.id}
      cloudRef: ${defaultCloud.id}
      servers:
        - ips:
            - type: V4
              addr: x.x.x.x
          port: 8080
      failActions:
        - type: FAIL_ACTION_CLOSE_CONN
```
{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.avi.Pool;
import com.pulumi.avi.PoolArgs;
import com.pulumi.avi.inputs.PoolServerArgs;
import com.pulumi.avi.inputs.PoolFailActionArgs;
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
        var testpool = new Pool("testpool", PoolArgs.builder()
            .name("pool-42")
            .healthMonitorRefs(testHm1.id())
            .tenantRef(defaultTenant.id())
            .cloudRef(defaultCloud.id())
            .servers(PoolServerArgs.builder()
                .ips(PoolServerIpArgs.builder()
                    .type("V4")
                    .addr("x.x.x.x")
                    .build())
                .port("8080")
                .build())
            .failActions(PoolFailActionArgs.builder()
                .type("FAIL_ACTION_CLOSE_CONN")
                .build())
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}

The above code breaks down as follows::
* The line below declares a resource of type ‘pool’ with name ‘testpool’:

* The line below shows a field that is a reference to another Pulumi resource declared elsewhere in the plan:

* The line below shows a field that is a reference to a read-only function corresponding to an existing Avi object not defined through the pulumi preview:

* Read-only functions can be imported either by name or by UUID. For example, the “System-HTTP” application profile can be imported as a function by name using the following:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as avi from "@pulumi/avi";

const systemHttpProfile = avi.getApplicationprofile({
    name: "System-HTTP",
});
```
{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import pulumi_avi as avi

system_http_profile = avi.get_applicationprofile(name="System-HTTP")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Avi = Pulumi.Avi;

return await Deployment.RunAsync(() =>
{
    var systemHttpProfile = Avi.GetApplicationprofile.Invoke(new()
    {
        Name = "System-HTTP",
    });

});

```
{{% /choosable %}}
{{% choosable language go %}}
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/avi/v31/avi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := avi.LookupApplicationprofile(ctx, &avi.LookupApplicationprofileArgs{
			Name: pulumi.StringRef("System-HTTP"),
		}, nil)
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
variables:
  systemHttpProfile:
    fn::invoke:
      function: avi:getApplicationprofile
      arguments:
        name: System-HTTP
```
{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.avi.AviFunctions;
import com.pulumi.avi.inputs.GetApplicationprofileArgs;
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
        final var systemHttpProfile = AviFunctions.getApplicationprofile(GetApplicationprofileArgs.builder()
            .name("System-HTTP")
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}

* Alternatively, a function can be imported based on the object’s UUID rather than it’s name, for example:
  {{< chooser language "typescript,python,go,csharp,java,yaml" >}}
  {{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as avi from "@pulumi/avi";

const systemHttpProfile = avi.getApplicationprofile({
    uuid: "applicationprofile-xxxxxxx",
});
```
{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import pulumi_avi as avi

system_http_profile = avi.get_applicationprofile(uuid="applicationprofile-xxxxxxx")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Avi = Pulumi.Avi;

return await Deployment.RunAsync(() =>
{
    var systemHttpProfile = Avi.GetApplicationprofile.Invoke(new()
    {
        Uuid = "applicationprofile-xxxxxxx",
    });

});

```
{{% /choosable %}}
{{% choosable language go %}}
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/avi/v31/avi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := avi.LookupApplicationprofile(ctx, &avi.LookupApplicationprofileArgs{
			Uuid: pulumi.StringRef("applicationprofile-xxxxxxx"),
		}, nil)
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
variables:
  systemHttpProfile:
    fn::invoke:
      function: avi:getApplicationprofile
      arguments:
        uuid: applicationprofile-xxxxxxx
```
{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.avi.AviFunctions;
import com.pulumi.avi.inputs.GetApplicationprofileArgs;
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
        final var systemHttpProfile = AviFunctions.getApplicationprofile(GetApplicationprofileArgs.builder()
            .uuid("applicationprofile-xxxxxxx")
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
# Usage Avi Pulumi
The following are the steps to use the provider:
1. Pulumi module must declare which providers it requires, so that Pulumi can install and use them. Provider requirements are declared in a requiredProviders block. Starting with Pulumi version 0.13+, Avi Pulumi provider has been migrated to Pulumi registry. In order to use it, you need to add the below block in versions.tf file.

where,
* source — Source of Pulumi provider. Pulumi will pull the released Avi Pulumi provider from Pulumi registry from this namespace.

* version — Avi Pulumi provider release version in Pulumi registry. If this field is skipped then Pulumi will pull the latest release of Avi Pulumi provider from Pulumi registry.

2. Create Avi provider in pulumi preview. For instance, to restrict template version in 0.1.x, you can use the following CLI:
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    avi:aviController:
        value: x.x.x.x
    avi:aviPassword:
        value: password
    avi:aviTenant:
        value: admin
    avi:aviUsername:
        value: admin
    avi:aviVersion:
        value: 21.1.1

```

3. Create Avi virtual service and pool example. Here the VS and Pool depends on read only tenant, cloud, applicationprofile, sslkeyandcertificate, sslprofile functions and another networkprofile, vsvip, health monitor defined as a resource in the pulumi preview.
   {{< chooser language "typescript,python,go,csharp,java,yaml" >}}
   {{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as avi from "@pulumi/avi";

const defaultTenant = avi.getTenant({
    name: "admin",
});
const systemHttpProfile = avi.getApplicationprofile({
    name: "System-HTTP",
});
const defaultCloud = avi.getCloud({
    name: "Defaul-Cloud",
});
const systemDefaultCert = avi.getSslkeyandcertificate({
    name: "System-Default-Cert",
});
const systemStandardSslprofile = avi.getSslprofile({
    name: "System-Standard",
});
const networkProfile1 = new avi.Networkprofile("network_profile1", {
    name: "tf-network-profile",
    profiles: [{
        type: "PROTOCOL_TYPE_TCP_PROXY",
        tcpProxyProfiles: [{
            ignoreTimeWait: "false",
            timeWaitDelay: "2000",
            maxRetransmissions: "8",
            maxSynRetransmissions: "8",
            automatic: "true",
            receiveWindow: "64",
            naglesAlgorithm: "false",
            ipDscp: "0",
            reorderThreshold: "10",
            minRexmtTimeout: "100",
            idleConnectionType: "KEEP_ALIVE",
            idleConnectionTimeout: "600",
            useInterfaceMtu: "true",
            ccAlgo: "CC_ALGO_NEW_RENO",
            aggressiveCongestionAvoidance: "false",
            slowStartScalingFactor: "1",
            congestionRecoveryScalingFactor: "2",
            reassemblyQueueSize: "0",
            keepaliveInHalfcloseState: "true",
            autoWindowGrowth: "true",
        }],
    }],
    connectionMirror: "false",
    tenantRef: defaultTenant.then(defaultTenant => defaultTenant.id),
});
const testHm1 = new avi.Healthmonitor("test_hm_1", {
    name: "pulumi-monitor",
    type: "HEALTH_MONITOR_HTTP",
    tenantRef: defaultTenant.then(defaultTenant => defaultTenant.id),
});
const testpool = new avi.Pool("testpool", {
    name: "pool-42",
    healthMonitorRefs: [testHm1.healthmonitorId],
    tenantRef: defaultTenant.then(defaultTenant => defaultTenant.id),
    cloudRef: defaultCloud.then(defaultCloud => defaultCloud.id),
    servers: [{
        ips: [{
            type: "V4",
            addr: "x.x.x.x",
        }],
        port: "8080",
    }],
    failActions: [{
        type: "FAIL_ACTION_CLOSE_CONN",
    }],
});
const testVsvip = new avi.Vsvip("test_vsvip", {
    name: "pulumi-vip-1",
    vips: [{
        vipId: "0",
        ipAddresses: [{
            type: "V4",
            addr: "x.x.x.x",
        }],
    }],
    cloudRef: defaultCloud.then(defaultCloud => defaultCloud.id),
    tenantRef: defaultTenant.then(defaultTenant => defaultTenant.id),
});
const httpsVs = new avi.Virtualservice("https_vs", {
    name: "tf_vs",
    poolRef: testpool.poolId,
    tenantRef: defaultTenant.then(defaultTenant => defaultTenant.id),
    vsvipRef: testVsvip.vsvipId,
    cloudRef: defaultCloud.then(defaultCloud => defaultCloud.id),
    sslKeyAndCertificateRefs: [systemDefaultCert.then(systemDefaultCert => systemDefaultCert.id)],
    sslProfileRef: systemStandardSslprofile.then(systemStandardSslprofile => systemStandardSslprofile.id),
    applicationProfileRef: systemHttpProfile.then(systemHttpProfile => systemHttpProfile.id),
    networkProfileRef: networkProfile1.networkprofileId,
    cloudType: "CLOUD_VCENTER",
    services: [{
        port: "8443",
        enableSsl: "true",
    }],
});
```
{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import pulumi_avi as avi

default_tenant = avi.get_tenant(name="admin")
system_http_profile = avi.get_applicationprofile(name="System-HTTP")
default_cloud = avi.get_cloud(name="Defaul-Cloud")
system_default_cert = avi.get_sslkeyandcertificate(name="System-Default-Cert")
system_standard_sslprofile = avi.get_sslprofile(name="System-Standard")
network_profile1 = avi.Networkprofile("network_profile1",
    name="tf-network-profile",
    profiles=[{
        "type": "PROTOCOL_TYPE_TCP_PROXY",
        "tcp_proxy_profiles": [{
            "ignore_time_wait": "false",
            "time_wait_delay": "2000",
            "max_retransmissions": "8",
            "max_syn_retransmissions": "8",
            "automatic": "true",
            "receive_window": "64",
            "nagles_algorithm": "false",
            "ip_dscp": "0",
            "reorder_threshold": "10",
            "min_rexmt_timeout": "100",
            "idle_connection_type": "KEEP_ALIVE",
            "idle_connection_timeout": "600",
            "use_interface_mtu": "true",
            "cc_algo": "CC_ALGO_NEW_RENO",
            "aggressive_congestion_avoidance": "false",
            "slow_start_scaling_factor": "1",
            "congestion_recovery_scaling_factor": "2",
            "reassembly_queue_size": "0",
            "keepalive_in_halfclose_state": "true",
            "auto_window_growth": "true",
        }],
    }],
    connection_mirror="false",
    tenant_ref=default_tenant.id)
test_hm1 = avi.Healthmonitor("test_hm_1",
    name="pulumi-monitor",
    type="HEALTH_MONITOR_HTTP",
    tenant_ref=default_tenant.id)
testpool = avi.Pool("testpool",
    name="pool-42",
    health_monitor_refs=[test_hm1.healthmonitor_id],
    tenant_ref=default_tenant.id,
    cloud_ref=default_cloud.id,
    servers=[{
        "ips": [{
            "type": "V4",
            "addr": "x.x.x.x",
        }],
        "port": "8080",
    }],
    fail_actions=[{
        "type": "FAIL_ACTION_CLOSE_CONN",
    }])
test_vsvip = avi.Vsvip("test_vsvip",
    name="pulumi-vip-1",
    vips=[{
        "vip_id": "0",
        "ip_addresses": [{
            "type": "V4",
            "addr": "x.x.x.x",
        }],
    }],
    cloud_ref=default_cloud.id,
    tenant_ref=default_tenant.id)
https_vs = avi.Virtualservice("https_vs",
    name="tf_vs",
    pool_ref=testpool.pool_id,
    tenant_ref=default_tenant.id,
    vsvip_ref=test_vsvip.vsvip_id,
    cloud_ref=default_cloud.id,
    ssl_key_and_certificate_refs=[system_default_cert.id],
    ssl_profile_ref=system_standard_sslprofile.id,
    application_profile_ref=system_http_profile.id,
    network_profile_ref=network_profile1.networkprofile_id,
    cloud_type="CLOUD_VCENTER",
    services=[{
        "port": "8443",
        "enable_ssl": "true",
    }])
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Avi = Pulumi.Avi;

return await Deployment.RunAsync(() =>
{
    var defaultTenant = Avi.GetTenant.Invoke(new()
    {
        Name = "admin",
    });

    var systemHttpProfile = Avi.GetApplicationprofile.Invoke(new()
    {
        Name = "System-HTTP",
    });

    var defaultCloud = Avi.GetCloud.Invoke(new()
    {
        Name = "Defaul-Cloud",
    });

    var systemDefaultCert = Avi.GetSslkeyandcertificate.Invoke(new()
    {
        Name = "System-Default-Cert",
    });

    var systemStandardSslprofile = Avi.GetSslprofile.Invoke(new()
    {
        Name = "System-Standard",
    });

    var networkProfile1 = new Avi.Networkprofile("network_profile1", new()
    {
        Name = "tf-network-profile",
        Profiles = new[]
        {
            new Avi.Inputs.NetworkprofileProfileArgs
            {
                Type = "PROTOCOL_TYPE_TCP_PROXY",
                TcpProxyProfiles = new[]
                {
                    new Avi.Inputs.NetworkprofileProfileTcpProxyProfileArgs
                    {
                        IgnoreTimeWait = "false",
                        TimeWaitDelay = "2000",
                        MaxRetransmissions = "8",
                        MaxSynRetransmissions = "8",
                        Automatic = "true",
                        ReceiveWindow = "64",
                        NaglesAlgorithm = "false",
                        IpDscp = "0",
                        ReorderThreshold = "10",
                        MinRexmtTimeout = "100",
                        IdleConnectionType = "KEEP_ALIVE",
                        IdleConnectionTimeout = "600",
                        UseInterfaceMtu = "true",
                        CcAlgo = "CC_ALGO_NEW_RENO",
                        AggressiveCongestionAvoidance = "false",
                        SlowStartScalingFactor = "1",
                        CongestionRecoveryScalingFactor = "2",
                        ReassemblyQueueSize = "0",
                        KeepaliveInHalfcloseState = "true",
                        AutoWindowGrowth = "true",
                    },
                },
            },
        },
        ConnectionMirror = "false",
        TenantRef = defaultTenant.Apply(getTenantResult => getTenantResult.Id),
    });

    var testHm1 = new Avi.Healthmonitor("test_hm_1", new()
    {
        Name = "pulumi-monitor",
        Type = "HEALTH_MONITOR_HTTP",
        TenantRef = defaultTenant.Apply(getTenantResult => getTenantResult.Id),
    });

    var testpool = new Avi.Pool("testpool", new()
    {
        Name = "pool-42",
        HealthMonitorRefs = new[]
        {
            testHm1.HealthmonitorId,
        },
        TenantRef = defaultTenant.Apply(getTenantResult => getTenantResult.Id),
        CloudRef = defaultCloud.Apply(getCloudResult => getCloudResult.Id),
        Servers = new[]
        {
            new Avi.Inputs.PoolServerArgs
            {
                Ips = new[]
                {
                    new Avi.Inputs.PoolServerIpArgs
                    {
                        Type = "V4",
                        Addr = "x.x.x.x",
                    },
                },
                Port = "8080",
            },
        },
        FailActions = new[]
        {
            new Avi.Inputs.PoolFailActionArgs
            {
                Type = "FAIL_ACTION_CLOSE_CONN",
            },
        },
    });

    var testVsvip = new Avi.Vsvip("test_vsvip", new()
    {
        Name = "pulumi-vip-1",
        Vips = new[]
        {
            new Avi.Inputs.VsvipVipArgs
            {
                VipId = "0",
                IpAddresses = new[]
                {
                    new Avi.Inputs.VsvipVipIpAddressArgs
                    {
                        Type = "V4",
                        Addr = "x.x.x.x",
                    },
                },
            },
        },
        CloudRef = defaultCloud.Apply(getCloudResult => getCloudResult.Id),
        TenantRef = defaultTenant.Apply(getTenantResult => getTenantResult.Id),
    });

    var httpsVs = new Avi.Virtualservice("https_vs", new()
    {
        Name = "tf_vs",
        PoolRef = testpool.PoolId,
        TenantRef = defaultTenant.Apply(getTenantResult => getTenantResult.Id),
        VsvipRef = testVsvip.VsvipId,
        CloudRef = defaultCloud.Apply(getCloudResult => getCloudResult.Id),
        SslKeyAndCertificateRefs = new[]
        {
            systemDefaultCert.Apply(getSslkeyandcertificateResult => getSslkeyandcertificateResult.Id),
        },
        SslProfileRef = systemStandardSslprofile.Apply(getSslprofileResult => getSslprofileResult.Id),
        ApplicationProfileRef = systemHttpProfile.Apply(getApplicationprofileResult => getApplicationprofileResult.Id),
        NetworkProfileRef = networkProfile1.NetworkprofileId,
        CloudType = "CLOUD_VCENTER",
        Services = new[]
        {
            new Avi.Inputs.VirtualserviceServiceArgs
            {
                Port = "8443",
                EnableSsl = "true",
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
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/avi/v31/avi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		defaultTenant, err := avi.LookupTenant(ctx, &avi.LookupTenantArgs{
			Name: pulumi.StringRef("admin"),
		}, nil)
		if err != nil {
			return err
		}
		systemHttpProfile, err := avi.LookupApplicationprofile(ctx, &avi.LookupApplicationprofileArgs{
			Name: pulumi.StringRef("System-HTTP"),
		}, nil)
		if err != nil {
			return err
		}
		defaultCloud, err := avi.LookupCloud(ctx, &avi.LookupCloudArgs{
			Name: pulumi.StringRef("Defaul-Cloud"),
		}, nil)
		if err != nil {
			return err
		}
		systemDefaultCert, err := avi.LookupSslkeyandcertificate(ctx, &avi.LookupSslkeyandcertificateArgs{
			Name: pulumi.StringRef("System-Default-Cert"),
		}, nil)
		if err != nil {
			return err
		}
		systemStandardSslprofile, err := avi.LookupSslprofile(ctx, &avi.LookupSslprofileArgs{
			Name: pulumi.StringRef("System-Standard"),
		}, nil)
		if err != nil {
			return err
		}
		networkProfile1, err := avi.NewNetworkprofile(ctx, "network_profile1", &avi.NetworkprofileArgs{
			Name: pulumi.String("tf-network-profile"),
			Profiles: avi.NetworkprofileProfileArray{
				&avi.NetworkprofileProfileArgs{
					Type: pulumi.String("PROTOCOL_TYPE_TCP_PROXY"),
					TcpProxyProfiles: avi.NetworkprofileProfileTcpProxyProfileArray{
						&avi.NetworkprofileProfileTcpProxyProfileArgs{
							IgnoreTimeWait:                  pulumi.String("false"),
							TimeWaitDelay:                   pulumi.String("2000"),
							MaxRetransmissions:              pulumi.String("8"),
							MaxSynRetransmissions:           pulumi.String("8"),
							Automatic:                       pulumi.String("true"),
							ReceiveWindow:                   pulumi.String("64"),
							NaglesAlgorithm:                 pulumi.String("false"),
							IpDscp:                          pulumi.String("0"),
							ReorderThreshold:                pulumi.String("10"),
							MinRexmtTimeout:                 pulumi.String("100"),
							IdleConnectionType:              pulumi.String("KEEP_ALIVE"),
							IdleConnectionTimeout:           pulumi.String("600"),
							UseInterfaceMtu:                 pulumi.String("true"),
							CcAlgo:                          pulumi.String("CC_ALGO_NEW_RENO"),
							AggressiveCongestionAvoidance:   pulumi.String("false"),
							SlowStartScalingFactor:          pulumi.String("1"),
							CongestionRecoveryScalingFactor: pulumi.String("2"),
							ReassemblyQueueSize:             pulumi.String("0"),
							KeepaliveInHalfcloseState:       pulumi.String("true"),
							AutoWindowGrowth:                pulumi.String("true"),
						},
					},
				},
			},
			ConnectionMirror: pulumi.String("false"),
			TenantRef:        pulumi.String(defaultTenant.Id),
		})
		if err != nil {
			return err
		}
		testHm1, err := avi.NewHealthmonitor(ctx, "test_hm_1", &avi.HealthmonitorArgs{
			Name:      pulumi.String("pulumi-monitor"),
			Type:      pulumi.String("HEALTH_MONITOR_HTTP"),
			TenantRef: pulumi.String(defaultTenant.Id),
		})
		if err != nil {
			return err
		}
		testpool, err := avi.NewPool(ctx, "testpool", &avi.PoolArgs{
			Name: pulumi.String("pool-42"),
			HealthMonitorRefs: pulumi.StringArray{
				testHm1.HealthmonitorId,
			},
			TenantRef: pulumi.String(defaultTenant.Id),
			CloudRef:  pulumi.String(defaultCloud.Id),
			Servers: avi.PoolServerArray{
				&avi.PoolServerArgs{
					Ips: avi.PoolServerIpArray{
						&avi.PoolServerIpArgs{
							Type: pulumi.String("V4"),
							Addr: pulumi.String("x.x.x.x"),
						},
					},
					Port: pulumi.String("8080"),
				},
			},
			FailActions: avi.PoolFailActionArray{
				&avi.PoolFailActionArgs{
					Type: pulumi.String("FAIL_ACTION_CLOSE_CONN"),
				},
			},
		})
		if err != nil {
			return err
		}
		testVsvip, err := avi.NewVsvip(ctx, "test_vsvip", &avi.VsvipArgs{
			Name: pulumi.String("pulumi-vip-1"),
			Vips: avi.VsvipVipArray{
				&avi.VsvipVipArgs{
					VipId: pulumi.String("0"),
					IpAddresses: avi.VsvipVipIpAddressArray{
						&avi.VsvipVipIpAddressArgs{
							Type: pulumi.String("V4"),
							Addr: pulumi.String("x.x.x.x"),
						},
					},
				},
			},
			CloudRef:  pulumi.String(defaultCloud.Id),
			TenantRef: pulumi.String(defaultTenant.Id),
		})
		if err != nil {
			return err
		}
		_, err = avi.NewVirtualservice(ctx, "https_vs", &avi.VirtualserviceArgs{
			Name:      pulumi.String("tf_vs"),
			PoolRef:   testpool.PoolId,
			TenantRef: pulumi.String(defaultTenant.Id),
			VsvipRef:  testVsvip.VsvipId,
			CloudRef:  pulumi.String(defaultCloud.Id),
			SslKeyAndCertificateRefs: pulumi.StringArray{
				pulumi.String(systemDefaultCert.Id),
			},
			SslProfileRef:         pulumi.String(systemStandardSslprofile.Id),
			ApplicationProfileRef: pulumi.String(systemHttpProfile.Id),
			NetworkProfileRef:     networkProfile1.NetworkprofileId,
			CloudType:             pulumi.String("CLOUD_VCENTER"),
			Services: avi.VirtualserviceServiceArray{
				&avi.VirtualserviceServiceArgs{
					Port:      pulumi.String("8443"),
					EnableSsl: pulumi.String("true"),
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
  networkProfile1:
    type: avi:Networkprofile
    name: network_profile1
    properties:
      name: tf-network-profile
      profiles:
        - type: PROTOCOL_TYPE_TCP_PROXY
          tcpProxyProfiles:
            - ignoreTimeWait: false
              timeWaitDelay: 2000
              maxRetransmissions: 8
              maxSynRetransmissions: 8
              automatic: true
              receiveWindow: 64
              naglesAlgorithm: false
              ipDscp: 0
              reorderThreshold: 10
              minRexmtTimeout: 100
              idleConnectionType: KEEP_ALIVE
              idleConnectionTimeout: 600
              useInterfaceMtu: true
              ccAlgo: CC_ALGO_NEW_RENO
              aggressiveCongestionAvoidance: false
              slowStartScalingFactor: 1
              congestionRecoveryScalingFactor: 2
              reassemblyQueueSize: 0
              keepaliveInHalfcloseState: true
              autoWindowGrowth: true
      connectionMirror: false
      tenantRef: ${defaultTenant.id}
  testHm1:
    type: avi:Healthmonitor
    name: test_hm_1
    properties:
      name: pulumi-monitor
      type: HEALTH_MONITOR_HTTP
      tenantRef: ${defaultTenant.id}
  testpool:
    type: avi:Pool
    properties:
      name: pool-42
      healthMonitorRefs:
        - ${testHm1.healthmonitorId}
      tenantRef: ${defaultTenant.id}
      cloudRef: ${defaultCloud.id}
      servers:
        - ips:
            - type: V4
              addr: x.x.x.x
          port: 8080
      failActions:
        - type: FAIL_ACTION_CLOSE_CONN
  testVsvip:
    type: avi:Vsvip
    name: test_vsvip
    properties:
      name: pulumi-vip-1
      vips:
        - vipId: '0'
          ipAddresses:
            - type: V4
              addr: x.x.x.x
      cloudRef: ${defaultCloud.id}
      tenantRef: ${defaultTenant.id}
  httpsVs:
    type: avi:Virtualservice
    name: https_vs
    properties:
      name: tf_vs
      poolRef: ${testpool.poolId}
      tenantRef: ${defaultTenant.id}
      vsvipRef: ${testVsvip.vsvipId}
      cloudRef: ${defaultCloud.id}
      sslKeyAndCertificateRefs:
        - ${systemDefaultCert.id}
      sslProfileRef: ${systemStandardSslprofile.id}
      applicationProfileRef: ${systemHttpProfile.id}
      networkProfileRef: ${networkProfile1.networkprofileId}
      cloudType: CLOUD_VCENTER
      services:
        - port: 8443
          enableSsl: true
variables:
  defaultTenant:
    fn::invoke:
      function: avi:getTenant
      arguments:
        name: admin
  systemHttpProfile:
    fn::invoke:
      function: avi:getApplicationprofile
      arguments:
        name: System-HTTP
  defaultCloud:
    fn::invoke:
      function: avi:getCloud
      arguments:
        name: Defaul-Cloud
  systemDefaultCert:
    fn::invoke:
      function: avi:getSslkeyandcertificate
      arguments:
        name: System-Default-Cert
  systemStandardSslprofile:
    fn::invoke:
      function: avi:getSslprofile
      arguments:
        name: System-Standard
```
{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.avi.AviFunctions;
import com.pulumi.avi.inputs.GetTenantArgs;
import com.pulumi.avi.inputs.GetApplicationprofileArgs;
import com.pulumi.avi.inputs.GetCloudArgs;
import com.pulumi.avi.inputs.GetSslkeyandcertificateArgs;
import com.pulumi.avi.inputs.GetSslprofileArgs;
import com.pulumi.avi.Networkprofile;
import com.pulumi.avi.NetworkprofileArgs;
import com.pulumi.avi.inputs.NetworkprofileProfileArgs;
import com.pulumi.avi.Healthmonitor;
import com.pulumi.avi.HealthmonitorArgs;
import com.pulumi.avi.Pool;
import com.pulumi.avi.PoolArgs;
import com.pulumi.avi.inputs.PoolServerArgs;
import com.pulumi.avi.inputs.PoolFailActionArgs;
import com.pulumi.avi.Vsvip;
import com.pulumi.avi.VsvipArgs;
import com.pulumi.avi.inputs.VsvipVipArgs;
import com.pulumi.avi.Virtualservice;
import com.pulumi.avi.VirtualserviceArgs;
import com.pulumi.avi.inputs.VirtualserviceServiceArgs;
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
        final var defaultTenant = AviFunctions.getTenant(GetTenantArgs.builder()
            .name("admin")
            .build());

        final var systemHttpProfile = AviFunctions.getApplicationprofile(GetApplicationprofileArgs.builder()
            .name("System-HTTP")
            .build());

        final var defaultCloud = AviFunctions.getCloud(GetCloudArgs.builder()
            .name("Defaul-Cloud")
            .build());

        final var systemDefaultCert = AviFunctions.getSslkeyandcertificate(GetSslkeyandcertificateArgs.builder()
            .name("System-Default-Cert")
            .build());

        final var systemStandardSslprofile = AviFunctions.getSslprofile(GetSslprofileArgs.builder()
            .name("System-Standard")
            .build());

        var networkProfile1 = new Networkprofile("networkProfile1", NetworkprofileArgs.builder()
            .name("tf-network-profile")
            .profiles(NetworkprofileProfileArgs.builder()
                .type("PROTOCOL_TYPE_TCP_PROXY")
                .tcpProxyProfiles(NetworkprofileProfileTcpProxyProfileArgs.builder()
                    .ignoreTimeWait("false")
                    .timeWaitDelay("2000")
                    .maxRetransmissions("8")
                    .maxSynRetransmissions("8")
                    .automatic("true")
                    .receiveWindow("64")
                    .naglesAlgorithm("false")
                    .ipDscp("0")
                    .reorderThreshold("10")
                    .minRexmtTimeout("100")
                    .idleConnectionType("KEEP_ALIVE")
                    .idleConnectionTimeout("600")
                    .useInterfaceMtu("true")
                    .ccAlgo("CC_ALGO_NEW_RENO")
                    .aggressiveCongestionAvoidance("false")
                    .slowStartScalingFactor("1")
                    .congestionRecoveryScalingFactor("2")
                    .reassemblyQueueSize("0")
                    .keepaliveInHalfcloseState("true")
                    .autoWindowGrowth("true")
                    .build())
                .build())
            .connectionMirror("false")
            .tenantRef(defaultTenant.id())
            .build());

        var testHm1 = new Healthmonitor("testHm1", HealthmonitorArgs.builder()
            .name("pulumi-monitor")
            .type("HEALTH_MONITOR_HTTP")
            .tenantRef(defaultTenant.id())
            .build());

        var testpool = new Pool("testpool", PoolArgs.builder()
            .name("pool-42")
            .healthMonitorRefs(testHm1.healthmonitorId())
            .tenantRef(defaultTenant.id())
            .cloudRef(defaultCloud.id())
            .servers(PoolServerArgs.builder()
                .ips(PoolServerIpArgs.builder()
                    .type("V4")
                    .addr("x.x.x.x")
                    .build())
                .port("8080")
                .build())
            .failActions(PoolFailActionArgs.builder()
                .type("FAIL_ACTION_CLOSE_CONN")
                .build())
            .build());

        var testVsvip = new Vsvip("testVsvip", VsvipArgs.builder()
            .name("pulumi-vip-1")
            .vips(VsvipVipArgs.builder()
                .vipId("0")
                .ipAddresses(VsvipVipIpAddressArgs.builder()
                    .type("V4")
                    .addr("x.x.x.x")
                    .build())
                .build())
            .cloudRef(defaultCloud.id())
            .tenantRef(defaultTenant.id())
            .build());

        var httpsVs = new Virtualservice("httpsVs", VirtualserviceArgs.builder()
            .name("tf_vs")
            .poolRef(testpool.poolId())
            .tenantRef(defaultTenant.id())
            .vsvipRef(testVsvip.vsvipId())
            .cloudRef(defaultCloud.id())
            .sslKeyAndCertificateRefs(systemDefaultCert.id())
            .sslProfileRef(systemStandardSslprofile.id())
            .applicationProfileRef(systemHttpProfile.id())
            .networkProfileRef(networkProfile1.networkprofileId())
            .cloudType("CLOUD_VCENTER")
            .services(VirtualserviceServiceArgs.builder()
                .port("8443")
                .enableSsl("true")
                .build())
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Authentication

The AVI provider offers following means of providing credentials for
authentication:

- Static credentials
- Environment variable
### Static credentials

Static credentials can be provided by adding an `aviUsername`, `aviPassword`, `aviControllerIp` and `avi.Tenant` in-line in the
AVI provider configuration:

Usage:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    avi:aviController:
        value: x.x.x.x
    avi:aviPassword:
        value: password
    avi:aviTenant:
        value: admin
    avi:aviUsername:
        value: admin
    avi:aviVersion:
        value: 21.1.1

```

If a Saas controller is used , then csp authentication is done by adding an `aviCspHost` and `aviCspToken` in-line in the
AVI provider configuration:
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    avi:aviController:
        value: x.x.x.x
    avi:aviCspHost:
        value: csp_host
    avi:aviCspToken:
        value: csp_token
    avi:aviTenant:
        value: admin
    avi:aviVersion:
        value: 21.1.1

```
### Environment variables

You can provide your credentials via the `AVI_USERNAME`, `AVI_PASSWORD`, `AVI_CONTROLLER` , `AVI_VERSION` and `AVI_TENANT`
environment variables, representing your AVI username, password, controller, api version and tenant, respectively.

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:

```
### Usage:

```sh
$ export AVI_USERNAME = username
$ export AVI_PASSWORD = password
$ export AVI_CONTROLLER = x.x.x.x
$ export AVI_TENANT = foo
$ export AVI_VERSION=21.1.1
$ export CSP_HOST = host
$ export CSP_TOKEN = token
$ pulumi up
$ pulumi preview
```

You can export `AVI_SUPPRESS_SENSITIVE_FIELDS_DIFF` environment variable as `AVI_SUPPRESS_SENSITIVE_FIELDS_DIFF=true`
if you want Pulumi to suppress the difference for sensitive fields during the plan update. However, if you want to
intentionally update the sensitive fields in the plan update then you need to un-export the environment variable or set
it to false, i.e. export `AVI_SUPPRESS_SENSITIVE_FIELDS_DIFF=false` or unset `AVI_SUPPRESS_SENSITIVE_FIELDS_DIFF`.
### Usage:
```sh
$ export AVI_SUPPRESS_SENSITIVE_FIELDS_DIFF = true
```
# Examples

|         Name          |                                  Link                                   |                                                                                                                                                                                            Description                                                                                                                                                                                             |
|-----------------------|-------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Virtual Services      | Basic VS <br>HTTPS VS <br> HTTP VS <br> HTTP VS <br> DNS VS <br> TCP VS | Using these examples we can configure HTTPS Virtual Service on AVI controller with dependent resources including Pool, Pool Group, SSL Certificates, Health Monitor, VsVip, Application Persistence Profile, Network Profile and functions including Application Profiles, Tenant, Cloud, Service Engine Group, Network Profile, Analytics Profile, SSL certificates, SSl Profiles, VRF Context. |
| VMware Environment    | Vcenter Cloud <br> Vcenter Controller Deployment                        | Using Vcenter Cloud example we can configure cloud of type CLOUD_VCENTER on AVI controller. <br> Using Controller Deployment examples we can deploy AVI Controller on the Vcenter using Ova or Ovf or content library.                                                                                                                                                                             |
| NSXT Environment      | nsx-t Cloud                                                             | Using this example we can integrate AVI Solution with NSX-T. This pulumi preview will create NSX-T cloud using AVI Pulumi provider.                                                                                                                                                                                                                                                             |
| GCP Environment       | GCP Cloud and Controller                                                | These examples deploy AVI controller on GCP, configures AVI controller cluster and cloud of type GCP.                                                                                                                                                                                                                                                                                              |
| Azure Environment     | Azure Cloud and Controller                                              | These examples deploy AVI controller on Azure, configures AVI controller cluster, cloud of type Azure, Virtual Services and its dependent resources.                                                                                                                                                                                                                                               |
| AWS Environment       | AWS Cloud and Controller                                                | These examples deploy AVI controller on AWS, configures AVI controller cluster, cloud of type AWS, Virtual Services and its dependent resources.                                                                                                                                                                                                                                                   |
| WAF                   | Waf Profile/Policy                                                      | Using this example we can configure WAF Policy on AVI controller with resources including Waf Profile, WAF Policy, Virtual Service, Pool, Pool Group, SSL profile, VsVip, Network Profile and functions including WAF Application Signature Provider, Application Profiles, Tenant, Cloud, SSL certificates, SSl Profiles, VRF Context.                                                          |
| Horizon               | Horizon VS                                                              | This Pulumi example can be used to configure AVI for Horizon in a Shared VIP with L7 and L4 Virtual Services.                                                                                                                                                                                                                                                                                   |
| Openstack Environment | Openstack Cloud and Controller                                          | Using this example we can launch AVI controller instance in openstack environment and configure application(Virtual Service) on controller.                                                                                                                                                                                                                                                        |