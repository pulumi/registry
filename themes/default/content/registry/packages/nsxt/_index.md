---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/vmware/nsxt/3.11.0/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Nsxt Provider
meta_desc: Provides an overview on how to configure the Pulumi Nsxt provider.
layout: package
---

## Generate Provider

The Nsxt provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider vmware/nsxt
```
## Overview

The NSX Pulumi provider gives the NSX administrator a way to automate NSX to
provide virtualized networking and security services using both ESXi and KVM
based hypervisor hosts as well as container networking and security.

More information on NSX can be found on the [NSX Product
Page](https://www.vmware.com/products/nsx.html)

Documentation on the NSX platform can be found on the [NSX Documentation
Page](https://docs.vmware.com/en/VMware-NSX-T/index.html)

Please use the navigation to the left to read about available functions and
resources.
## Basic Configuration of the NSX Pulumi Provider

In order to use the NSX Pulumi provider you must first configure the
provider to communicate with the VMware NSX manager. The NSX manager is the
system which serves the NSX REST API and provides a way to configure the
desired state of the NSX system. The configuration of the NSX provider requires
the IP address, hostname, or FQDN of the NSX manager.

The NSX provider offers several ways to authenticate to the NSX manager.
Credentials can be provided statically or provided as environment variables. In
addition, client certificates can be used for authentication. For
authentication with certificates Pulumi will require a certificate file and
private key file in PEM format. To use client certificates the client
certificate needs to be registered with NSX manager prior to invoking
Pulumi.

The provider also can accept both signed and self-signed server certificates.
It is recommended that in production environments you only use certificates
signed by a certificate authority. NSX ships by default with a self-signed
server certificates as the hostname of the NSX manager is not known until the
NSX administrator determines what name or IP to use.

Setting the `allowUnverifiedSsl` parameter to `true` will direct the
Pulumi client to skip server certificate verification. This is not
recommended in production deployments as it is recommended that you use trusted
connection using certificates signed by a certificate authority.

With the `caFile` parameter you can also specify a file that contains your
certificate authority certificate in PEM format to verify certificates with a
certificate authority.

There are also a number of other parameters that can be set to tune how the
provider connects to the NSX REST API. It is recommended you leave these to the
defaults unless you experience issues in which case they can be tuned to
optimize the system in your environment.

Note that with pulumi 0.14 onwards, `pulumi` block should be added to your
configuration:

Note that in all of the examples you will need to update attributes such as
`host`, `username`, `password`, `vmcToken` to match your NSX deployment.
### Example of Configuration with Credentials

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    nsxt:allowUnverifiedSsl:
        value: true
    nsxt:host:
        value: 192.168.110.41
    nsxt:maxRetries:
        value: 2
    nsxt:password:
        value: default
    nsxt:username:
        value: admin

```
### Example of Setting Environment Variables

```
export NSXT_MANAGER_HOST="192.168.110.41"
export NSXT_USERNAME="admin"
export NSXT_PASSWORD="default"
```
### Example using a Client Certificate

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    nsxt:allowUnverifiedSsl:
        value: true
    nsxt:clientAuthCertFile:
        value: mycert.pem
    nsxt:clientAuthKeyFile:
        value: mykey.pem
    nsxt:host:
        value: 192.168.110.41

```
### Example with Certificate Authority Certificate

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    nsxt:caFile:
        value: myca.pem
    nsxt:host:
        value: 10.160.94.11
    nsxt:password:
        value: qwerty
    nsxt:username:
        value: admin

```
### VMC Environment Example

Note that only a limited subset of policy resources are supported with VMC.

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    nsxt:allowUnverifiedSsl:
        value: true
    nsxt:enforcementPoint:
        value: vmc-enforcementpoint
    nsxt:host:
        value: x-54-200-54-5.rp.vmwarevmc.com/vmc/reverse-proxy/api/orgs/b003c3a5-3f68-4a8c-a74f-f79a0625da17/sddcs/d2f43050-f4e2-4989-ab52-2eb0b89d8487/sks-nsxt-manager
    nsxt:vmcToken:
        value: 5aVZEj6dJN1bQ6ZheakMyV0Qbj7P65sa2pYuhgx7Mp5glvgCkFKHcGxy3KmslllT

```
### VMC PCI Compliant Environment Example

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    nsxt:allowUnverifiedSsl:
        value: true
    nsxt:enforcementPoint:
        value: vmc-enforcementpoint
    nsxt:host:
        value: 10.4.14.23
    nsxt:password:
        value: qwerty
    nsxt:username:
        value: admin
    nsxt:vmcAuthMode:
        value: Basic

```
### Policy Global Manager Example

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    nsxt:globalManager:
        value: true
    nsxt:host:
        value: 192.168.110.41
    nsxt:maxRetries:
        value: 10
    nsxt:password:
        value: default
    nsxt:retryMaxDelay:
        value: 1000
    nsxt:retryMinDelay:
        value: 500
    nsxt:username:
        value: admin

```
## Configuration Reference

The following arguments are used to configure the provider:

* `host` - (Required) The host name or IP address of the NSX manager. Can also
  be specified with the `NSXT_MANAGER_HOST` environment variable. Do not include
  `http://` or `https://` in the host.
* `username` - (Required) The user name to connect to the NSX manager as. Can
  also be specified with the `NSXT_USERNAME` environment variable.
* `password` - (Required) The password for the NSX manager user. Can also be
  specified with the `NSXT_PASSWORD` environment variable.
* `clientAuthCertFile` - (Optional) The path to a certificate file for client
  certificate authorization. Can also be specified with the
  `NSXT_CLIENT_AUTH_CERT_FILE` environment variable.
* `clientAuthKeyFile` - (Optional) The path to a private key file for the
  certificate supplied to `clientAuthCertFile`. Can also be specified with
  the `NSXT_CLIENT_AUTH_KEY_FILE` environment variable.
* `clientAuthCert` - (Optional) Client certificate string.
  Can also be specified with the `NSXT_CLIENT_AUTH_CERT` environment variable.
* `clientAuthKey` - (Optional) Client certificate private key string.
  Can also be specified with the `NSXT_CLIENT_AUTH_KEY` environment variable.
* `allowUnverifiedSsl` - (Optional) Boolean that can be set to true to disable
  SSL certificate verification. This should be used with care as it could allow
  an attacker to intercept your auth token. If omitted, default value is
  `false`. Can also be specified with the `NSXT_ALLOW_UNVERIFIED_SSL`
  environment variable.
* `caFile` - (Optional) The path to an optional CA certificate file for SSL
  validation. Can also be specified with the `NSXT_CA_FILE` environment
  variable.
* `ca` - (Optional) CA certificate string for SSL validation.
  Can also be specified with the `NSXT_CA` environment variable.
* `maxRetries` - (Optional) The maximum number of retires before failing an API
  request. Default: `4` Can also be specified with the `NSXT_MAX_RETRIES`
  environment variable. For Global Manager, it is recommended to increase this value
  since slower realization times tend to delay resolution of some errors.
* `retryMinDelay` - (Optional) The minimum delay, in milliseconds, between
  retries. Default: `0`. For Global Manager, it is recommended to increase this value
  since slower realization times tend to delay resolution of some errors.
  Can also be specified with the `NSXT_RETRY_MIN_DELAY` environment variable.
* `retryMaxDelay` - (Optional) The maximum delay, in milliseconds, between
  retries. Default: `500`. For Global Manager, it is recommended to increase this
  value since slower realization times tend to delay resolution of some errors.
  Can also be specified with the `NSXT_RETRY_MAX_DELAY` environment variable.
* `retryOnStatusCodes` - (Optional) A list of HTTP status codes to retry on.
  By default, the provider supplies a set of status codes recommended for retry with
  policy resources: `400, 409, 429, 500, 503, 504`. Can also be specified with the
  `NSXT_RETRY_ON_STATUS_CODES` environment variable.
* `remoteAuth` - (Optional) Would trigger remote authorization instead of basic
  authorization. This is required for users based on vIDM authentication for early
  NSX versions.
* `sessionAuth` - (Optional) Creates session to avoid re-authentication for every
  request. Speeds up pulumi execution for vIDM based environments. Defaults to `true`
  The default for this flag is false. Can also be specified with the
  `NSXT_REMOTE_AUTH` environment variable.
* `toleratePartialSuccess` - (Optional) Setting this flag to true would treat
  partially successful realization as valid state and not fail apply.
* `vmcToken` - (Optional) Long-lived API token for authenticating with VMware
  Cloud Services APIs. This token will be used to short-lived token that is
  needed to communicate with NSX Manager in VMC environment. Can not be specified
  together with `vmcClientId` and `vmcClientSecret`.
  Note that only subset of policy resources are supported with VMC environment.
* `vmcClientId`- (Optional) ID of OAuth App associated with the VMC organization.
  The combination with `vmcClientSecret` is used to authenticate when calling
  VMware Cloud Services APIs. Can not be specified together with `vmcToken`.
* `vmcClientSecret` - (Optional) Secret of OAuth App associated with the VMC
  organization. The combination with `vmcClientId` is used to authenticate when
  calling VMware Cloud Services APIs. Can not be specified together with `vmcToken`.
  Note that only subset of policy resources are supported with VMC environment.
* `vmcAuthHost` - (Optional) URL for VMC authorization service that is used
  to obtain short-lived token for NSX manager access. Defaults to VMC
  console authorization URL.
* `vmcAuthMode` - (Optional) VMC authorization mode, that determines what HTTP
  header is used for authorization. Accepted values are `Default`, `Bearer`, `Basic`.
  For direct VMC connections with a token, use `Bearer` mode. For PCI mode with basic
  authentication, use `Basic`. Otherwise no need to specify this setting.
* `enforcementPoint` - (Optional) Enforcement point, mostly relevant for policy
  functions. For VMC environment, this should be set to `vmc-enforcementpoint`.
  For on-prem deployments, this setting should not be specified.
* `globalManager` - (Optional) True if this is a global manager endpoint.
  False by default.
* `licenseKeys` - (Optional) List of NSX-T license keys. License keys are applied
  during plan or apply commands. Note that the provider will not remove license keys if
  those are removed from provider config - please clean up licenses manually.
* `onDemandConnection` - (Optional) Avoid verification on NSX connectivity on provider
  startup. Instead, initialize the connection on demand. This setting can not be turned on
  for VMC environments, and is not supported with deprecated NSX manager resources and
  functions. Note - this setting is useful when NSX manager is not yet available at
  time of provider evaluation, and not recommended to be turned on otherwise.
## NSX Logical Networking

This release of the NSX-T Pulumi Provider extends to cover NSX-T declarative
API called Policy. This API aims to simplify the consumption of logical objects
and offer additional capabilities.The NSX-T Pulumi Provider covers most of NSX
functionality.
While you can still build topologies from the imperative API and existing config files
will continue to work, the recommendation
is to build logical topologies from the declarative API(Policy Objects).The resources
and functions using the policy API have *policy* in their name.
For more details on the NSX-T Policy API usage, please refer to NSX-T documentation.

The existing functions and resources are still available to consume but using
the new Policy based functions and resources are recommended.
### Logical Networking and Security Example Usage

The following example demonstrates using the NSX Pulumi provider to create
Tier-1 Gateways, Segments, DHCP Service, Static and Dynamic Groups, Firewall
rules and tags the VMs
#### Example variables.tf File

This file allows you to define some variables that can be reused in multiple
.tf files.

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";

const config = new pulumi.Config();
const nsxManager = config.requireObject<any>("nsxManager");
const nsxUsername = config.requireObject<any>("nsxUsername");
const nsxPassword = config.requireObject<any>("nsxPassword");
```
{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi

config = pulumi.Config()
nsx_manager = config.require_object("nsxManager")
nsx_username = config.require_object("nsxUsername")
nsx_password = config.require_object("nsxPassword")
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
    var nsxManager = config.RequireObject<dynamic>("nsxManager");
    var nsxUsername = config.RequireObject<dynamic>("nsxUsername");
    var nsxPassword = config.RequireObject<dynamic>("nsxPassword");
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
		nsxManager := cfg.RequireObject("nsxManager")
		nsxUsername := cfg.RequireObject("nsxUsername")
		nsxPassword := cfg.RequireObject("nsxPassword")
		return nil
	})
}
```
{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
configuration:
  nsxManager:
    type: dynamic
  nsxUsername:
    type: dynamic
  nsxPassword:
    type: dynamic
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
        final var nsxManager = config.get("nsxManager");
        final var nsxUsername = config.get("nsxUsername");
        final var nsxPassword = config.get("nsxPassword");
    }
}
```
{{% /choosable %}}
{{< /chooser >}}
#### Example pulumi.tfvars File

This file allows you to set some variables that can be reused in multiple .tf
files.
#### Example nsx.tf file

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    nsxt:allowUnverifiedSsl:
        value: true
    nsxt:host:
        value: 'TODO: var.nsx_manager'
    nsxt:maxRetries:
        value: 10
    nsxt:password:
        value: 'TODO: var.nsx_password'
    nsxt:retryMaxDelay:
        value: 5000
    nsxt:retryMinDelay:
        value: 500
    nsxt:retryOnStatusCodes:
        value:
            - 429
    nsxt:username:
        value: 'TODO: var.nsx_username'

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as nsxt from "@pulumi/nsxt";

const config = new pulumi.Config();
const nsxTagScope = config.get("nsxTagScope") || "project";
const nsxTag = config.get("nsxTag") || "pulumi-demo";
//
// This part of the example shows some functions we will need to refer to
// later in the .tf file. They include the transport zone, tier 0 router and
// edge cluster.
// There Tier-0 (T0) Gateway is considered a "provider" router that is pre-created
// by the NSX Admin. A T0 Gateway is used for north/south connectivity between
// the logical networking space and the physical networking space. Many Tier1
// Gateways will be connected to the T0 Gateway
//
const demo = nsxt.getPolicyEdgeCluster({
    displayName: "Edge-Cluster",
});
const overlayTz = nsxt.getPolicyTransportZone({
    displayName: "Overlay-TZ",
});
const t0Gateway = nsxt.getPolicyTier0Gateway({
    displayName: "TF-T0-Gateway",
});
//
// Create a DHCP Profile that is used later
// Note, this resource is only in NSX 3.0.0+
const tierDhcp = new nsxt.PolicyDhcpServer("tier_dhcp", {
    displayName: "tier_dhcp",
    description: "DHCP server servicing all 3 Segments",
    serverAddresses: ["12.12.99.2/24"],
});
//
// In this part of the example, the settings required to create a Tier1 Gateway
// are defined. In NSX a Tier1 Gateway is often used on a per user, tenant,
// department or application basis. Each application may have it's own Tier1
// Gateway. The Tier1 Gateway provides the default gateway for virtual machines
// connected to the Segments on the Tier1 Gateway
//
const t1Gateway = new nsxt.PolicyTier1Gateway("t1_gateway", {
    displayName: "TF_T1",
    description: "Tier1 provisioned by Pulumi",
    edgeClusterPath: demo.then(demo => demo.path),
    dhcpConfigPath: tierDhcp.path,
    failoverMode: "PREEMPTIVE",
    defaultRuleLogging: false,
    enableFirewall: true,
    enableStandbyRelocation: false,
    forceWhitelisting: false,
    tier0Path: t0Gateway.then(t0Gateway => t0Gateway.path),
    routeAdvertisementTypes: [
        "TIER1_STATIC_ROUTES",
        "TIER1_CONNECTED",
    ],
    poolAllocation: "ROUTING",
    tags: [{
        scope: nsxTagScope,
        tag: nsxTag,
    }],
    routeAdvertisementRules: [{
        name: "rule1",
        action: "DENY",
        subnets: [
            "20.0.0.0/24",
            "21.0.0.0/24",
        ],
        prefixOperator: "GE",
        routeAdvertisementTypes: ["TIER1_CONNECTED"],
    }],
});
//
// This shows the settings required to create NSX Segment (Logical Switch) to
// which you can attach Virtual Machines (VMs)
//
const web = new nsxt.PolicySegment("web", {
    displayName: "web-tier",
    description: "Pulumi provisioned Web Segment",
    connectivityPath: t1Gateway.path,
    transportZonePath: overlayTz.then(overlayTz => overlayTz.path),
    subnets: [{
        cidr: "12.12.1.1/24",
        dhcpRanges: ["12.12.1.100-12.12.1.160"],
        dhcpV4Config: {
            serverAddress: "12.12.1.2/24",
            leaseTime: 36000,
            dhcpOption121s: [{
                network: "6.6.6.0/24",
                nextHop: "1.1.1.21",
            }],
        },
    }],
    tags: [
        {
            scope: nsxTagScope,
            tag: nsxTag,
        },
        {
            scope: "tier",
            tag: "web",
        },
    ],
});
const app = new nsxt.PolicySegment("app", {
    displayName: "app-tier",
    description: "Pulumi provisioned App Segment",
    connectivityPath: t1Gateway.path,
    transportZonePath: overlayTz.then(overlayTz => overlayTz.path),
    subnets: [{
        cidr: "12.12.2.1/24",
        dhcpRanges: ["12.12.2.100-12.12.2.160"],
        dhcpV4Config: {
            serverAddress: "12.12.2.2/24",
            leaseTime: 36000,
            dhcpOption121s: [{
                network: "6.6.6.0/24",
                nextHop: "1.1.1.21",
            }],
        },
    }],
    tags: [
        {
            scope: nsxTagScope,
            tag: nsxTag,
        },
        {
            scope: "tier",
            tag: "app",
        },
    ],
});
const db = new nsxt.PolicySegment("db", {
    displayName: "db-tier",
    description: "Pulumi provisioned DB Segment",
    connectivityPath: t1Gateway.path,
    transportZonePath: overlayTz.then(overlayTz => overlayTz.path),
    subnets: [{
        cidr: "12.12.3.1/24",
        dhcpRanges: ["12.12.3.100-12.12.3.160"],
        dhcpV4Config: {
            serverAddress: "12.12.3.2/24",
            leaseTime: 36000,
            dhcpOption121s: [{
                network: "6.6.6.0/24",
                nextHop: "1.1.1.21",
            }],
        },
    }],
    tags: [
        {
            scope: nsxTagScope,
            tag: nsxTag,
        },
        {
            scope: "tier",
            tag: "db",
        },
    ],
});
//
// This part of the example shows creating Groups with dynamic membership
// criteria
//
// All Virtual machines with specific tag and scope
const allVms = new nsxt.PolicyGroup("all_vms", {
    displayName: "All_VMs",
    description: "Group consisting of ALL VMs",
    criterias: [{
        conditions: [{
            memberType: "VirtualMachine",
            operator: "CONTAINS",
            key: "Tag",
            value: nsxTag,
        }],
    }],
});
// All WEB VMs
const webGroup = new nsxt.PolicyGroup("web_group", {
    displayName: "Web_VMs",
    description: "Group consisting of Web VMs",
    criterias: [{
        conditions: [{
            memberType: "VirtualMachine",
            operator: "CONTAINS",
            key: "Tag",
            value: "web",
        }],
    }],
    tags: [{
        scope: nsxTagScope,
        tag: nsxTag,
    }],
});
// All App VMs
const appGroup = new nsxt.PolicyGroup("app_group", {
    displayName: "App_VMs",
    description: "Group consisting of App VMs",
    criterias: [{
        conditions: [{
            memberType: "VirtualMachine",
            operator: "CONTAINS",
            key: "Tag",
            value: "app",
        }],
    }],
    tags: [{
        scope: nsxTagScope,
        tag: nsxTag,
    }],
});
// All DB VMs
const dbGroup = new nsxt.PolicyGroup("db_group", {
    displayName: "DB_VMs",
    description: "Group consisting of DB VMs",
    criterias: [{
        conditions: [{
            memberType: "VirtualMachine",
            operator: "CONTAINS",
            key: "Tag",
            value: "db",
        }],
    }],
    tags: [{
        scope: nsxTagScope,
        tag: nsxTag,
    }],
});
// Static Group of IP addresses
const ipSet = new nsxt.PolicyGroup("ip_set", {
    displayName: "external_IPs",
    description: "Group containing all external IPs",
    criterias: [{
        ipaddressExpression: {
            ipAddresses: [
                "211.1.1.1",
                "212.1.1.1",
                "192.168.1.1-192.168.1.100",
            ],
        },
    }],
    tags: [{
        scope: nsxTagScope,
        tag: nsxTag,
    }],
});
//
// An example for Service for App that listens on port 8443
//
const appService = new nsxt.PolicyService("app_service", {
    displayName: "app_service_8443",
    description: "Service for App that listens on port 8443",
    l4PortSetEntries: [{
        description: "TCP Port 8443",
        protocol: "TCP",
        destinationPorts: ["8443"],
    }],
    tags: [{
        scope: nsxTagScope,
        tag: nsxTag,
    }],
});
//
// Here we have examples of create functions for Services
//
const https = nsxt.getPolicyService({
    displayName: "HTTPS",
});
const mysql = nsxt.getPolicyService({
    displayName: "MySQL",
});
const ssh = nsxt.getPolicyService({
    displayName: "SSH",
});
//
// In this section, we have example to create Firewall sections and rules
// All rules in this section will be applied to VMs that are part of the
// Gropus we created earlier
//
const firewallSection = new nsxt.PolicySecurityPolicy("firewall_section", {
    displayName: "DFW Section",
    description: "Firewall section created by Pulumi",
    scopes: [allVms.path],
    category: "Application",
    locked: false,
    stateful: true,
    tags: [{
        scope: nsxTagScope,
        tag: nsxTag,
    }],
    rules: [
        {
            displayName: "Allow HTTPS",
            description: "In going rule",
            action: "ALLOW",
            logged: false,
            ipVersion: "IPV4",
            destinationGroups: [webGroup.path],
            services: [https.then(https => https.path)],
        },
        {
            displayName: "Allow SSH",
            description: "In going rule",
            action: "ALLOW",
            logged: false,
            ipVersion: "IPV4",
            destinationGroups: [webGroup.path],
            services: [ssh.then(ssh => ssh.path)],
        },
        {
            displayName: "Allow Web to App",
            description: "Web to App communication",
            action: "ALLOW",
            logged: false,
            ipVersion: "IPV4",
            sourceGroups: [webGroup.path],
            destinationGroups: [appGroup.path],
            services: [appService.path],
        },
        {
            displayName: "Allow App to DB",
            description: "App to DB communication",
            action: "ALLOW",
            logged: false,
            ipVersion: "IPV4",
            sourceGroups: [appGroup.path],
            destinationGroups: [dbGroup.path],
            services: [mysql.then(mysql => mysql.path)],
        },
        {
            displayName: "Allow Infrastructure",
            description: "Allow DNS and Management servers",
            action: "ALLOW",
            logged: true,
            ipVersion: "IPV4",
            sourceGroups: [ipSet.path],
            destinationGroups: [allVms.path],
        },
        {
            displayName: "Allow out",
            description: "Outgoing rule",
            action: "ALLOW",
            logged: true,
            ipVersion: "IPV4",
            sourceGroups: [allVms.path],
        },
        {
            displayName: "Deny ANY",
            description: "Default Deny the traffic",
            action: "REJECT",
            logged: true,
            ipVersion: "IPV4",
        },
    ],
});
//
// Here we have examples for creating NAT rules. The example here assumes
// the Web IP addresses are reachable from outside and no NAT is required.
//
const rule1 = new nsxt.PolicyNatRule("rule1", {
    displayName: "App 1-to-1 In",
    action: "SNAT",
    translatedNetworks: ["102.10.22.1"],
    sourceNetworks: ["12.12.2.0/24"],
    gatewayPath: t1Gateway.path,
    tags: [{
        scope: nsxTagScope,
        tag: nsxTag,
    }],
});
const rule2 = new nsxt.PolicyNatRule("rule2", {
    displayName: "App 1-to-1 Out",
    action: "DNAT",
    translatedNetworks: ["102.10.22.2"],
    destinationNetworks: ["102.10.22.1/32"],
    sourceNetworks: ["12.12.2.0/24"],
    gatewayPath: t1Gateway.path,
    tags: [{
        scope: nsxTagScope,
        tag: nsxTag,
    }],
});
const rule3 = new nsxt.PolicyNatRule("rule3", {
    displayName: "DB 1-to-1 In",
    action: "SNAT",
    translatedNetworks: ["102.10.23.1"],
    sourceNetworks: ["12.12.3.0/24"],
    gatewayPath: t1Gateway.path,
    tags: [{
        scope: nsxTagScope,
        tag: nsxTag,
    }],
});
const rule4 = new nsxt.PolicyNatRule("rule4", {
    displayName: "App 1-to-1 Out",
    action: "DNAT",
    translatedNetworks: ["102.10.23.3"],
    destinationNetworks: ["102.10.23.1/32"],
    sourceNetworks: ["12.12.3.0/24"],
    gatewayPath: t1Gateway.path,
    tags: [{
        scope: nsxTagScope,
        tag: nsxTag,
    }],
});
// The 3 VMs available in the NSX Inventory
const webVm = nsxt.getPolicyVm({
    displayName: "web-vm",
});
const appVm = nsxt.getPolicyVm({
    displayName: "app-vm",
});
const dbVm = nsxt.getPolicyVm({
    displayName: "db-vm",
});
// Assign the right tags to the VMs so that they get included in the
// dynamic groups created above
const webVmTag = new nsxt.PolicyVmTags("web_vm_tag", {
    instanceId: webVm.then(webVm => webVm.instanceId),
    tags: [
        {
            scope: "tier",
            tag: "web",
        },
        {
            scope: nsxTagScope,
            tag: nsxTag,
        },
    ],
});
const appVmTag = new nsxt.PolicyVmTags("app_vm_tag", {
    instanceId: appVm.then(appVm => appVm.instanceId),
    tags: [
        {
            scope: "tier",
            tag: "app",
        },
        {
            scope: nsxTagScope,
            tag: nsxTag,
        },
    ],
});
const dbVmTag = new nsxt.PolicyVmTags("db_vm_tag", {
    instanceId: dbVm.then(dbVm => dbVm.instanceId),
    tags: [
        {
            scope: "tier",
            tag: "db",
        },
        {
            scope: nsxTagScope,
            tag: nsxTag,
        },
    ],
});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    nsxt:allowUnverifiedSsl:
        value: true
    nsxt:host:
        value: 'TODO: var.nsx_manager'
    nsxt:maxRetries:
        value: 10
    nsxt:password:
        value: 'TODO: var.nsx_password'
    nsxt:retryMaxDelay:
        value: 5000
    nsxt:retryMinDelay:
        value: 500
    nsxt:retryOnStatusCodes:
        value:
            - 429
    nsxt:username:
        value: 'TODO: var.nsx_username'

```
```python
import pulumi
import pulumi_nsxt as nsxt

config = pulumi.Config()
nsx_tag_scope = config.get("nsxTagScope")
if nsx_tag_scope is None:
    nsx_tag_scope = "project"
nsx_tag = config.get("nsxTag")
if nsx_tag is None:
    nsx_tag = "pulumi-demo"
#
# This part of the example shows some functions we will need to refer to
# later in the .tf file. They include the transport zone, tier 0 router and
# edge cluster.
# There Tier-0 (T0) Gateway is considered a "provider" router that is pre-created
# by the NSX Admin. A T0 Gateway is used for north/south connectivity between
# the logical networking space and the physical networking space. Many Tier1
# Gateways will be connected to the T0 Gateway
#
demo = nsxt.get_policy_edge_cluster(display_name="Edge-Cluster")
overlay_tz = nsxt.get_policy_transport_zone(display_name="Overlay-TZ")
t0_gateway = nsxt.get_policy_tier0_gateway(display_name="TF-T0-Gateway")
#
# Create a DHCP Profile that is used later
# Note, this resource is only in NSX 3.0.0+
tier_dhcp = nsxt.PolicyDhcpServer("tier_dhcp",
    display_name="tier_dhcp",
    description="DHCP server servicing all 3 Segments",
    server_addresses=["12.12.99.2/24"])
#
# In this part of the example, the settings required to create a Tier1 Gateway
# are defined. In NSX a Tier1 Gateway is often used on a per user, tenant,
# department or application basis. Each application may have it's own Tier1
# Gateway. The Tier1 Gateway provides the default gateway for virtual machines
# connected to the Segments on the Tier1 Gateway
#
t1_gateway = nsxt.PolicyTier1Gateway("t1_gateway",
    display_name="TF_T1",
    description="Tier1 provisioned by Pulumi",
    edge_cluster_path=demo.path,
    dhcp_config_path=tier_dhcp.path,
    failover_mode="PREEMPTIVE",
    default_rule_logging=False,
    enable_firewall=True,
    enable_standby_relocation=False,
    force_whitelisting=False,
    tier0_path=t0_gateway.path,
    route_advertisement_types=[
        "TIER1_STATIC_ROUTES",
        "TIER1_CONNECTED",
    ],
    pool_allocation="ROUTING",
    tags=[{
        "scope": nsx_tag_scope,
        "tag": nsx_tag,
    }],
    route_advertisement_rules=[{
        "name": "rule1",
        "action": "DENY",
        "subnets": [
            "20.0.0.0/24",
            "21.0.0.0/24",
        ],
        "prefix_operator": "GE",
        "route_advertisement_types": ["TIER1_CONNECTED"],
    }])
#
# This shows the settings required to create NSX Segment (Logical Switch) to
# which you can attach Virtual Machines (VMs)
#
web = nsxt.PolicySegment("web",
    display_name="web-tier",
    description="Pulumi provisioned Web Segment",
    connectivity_path=t1_gateway.path,
    transport_zone_path=overlay_tz.path,
    subnets=[{
        "cidr": "12.12.1.1/24",
        "dhcp_ranges": ["12.12.1.100-12.12.1.160"],
        "dhcp_v4_config": {
            "server_address": "12.12.1.2/24",
            "lease_time": 36000,
            "dhcp_option121s": [{
                "network": "6.6.6.0/24",
                "next_hop": "1.1.1.21",
            }],
        },
    }],
    tags=[
        {
            "scope": nsx_tag_scope,
            "tag": nsx_tag,
        },
        {
            "scope": "tier",
            "tag": "web",
        },
    ])
app = nsxt.PolicySegment("app",
    display_name="app-tier",
    description="Pulumi provisioned App Segment",
    connectivity_path=t1_gateway.path,
    transport_zone_path=overlay_tz.path,
    subnets=[{
        "cidr": "12.12.2.1/24",
        "dhcp_ranges": ["12.12.2.100-12.12.2.160"],
        "dhcp_v4_config": {
            "server_address": "12.12.2.2/24",
            "lease_time": 36000,
            "dhcp_option121s": [{
                "network": "6.6.6.0/24",
                "next_hop": "1.1.1.21",
            }],
        },
    }],
    tags=[
        {
            "scope": nsx_tag_scope,
            "tag": nsx_tag,
        },
        {
            "scope": "tier",
            "tag": "app",
        },
    ])
db = nsxt.PolicySegment("db",
    display_name="db-tier",
    description="Pulumi provisioned DB Segment",
    connectivity_path=t1_gateway.path,
    transport_zone_path=overlay_tz.path,
    subnets=[{
        "cidr": "12.12.3.1/24",
        "dhcp_ranges": ["12.12.3.100-12.12.3.160"],
        "dhcp_v4_config": {
            "server_address": "12.12.3.2/24",
            "lease_time": 36000,
            "dhcp_option121s": [{
                "network": "6.6.6.0/24",
                "next_hop": "1.1.1.21",
            }],
        },
    }],
    tags=[
        {
            "scope": nsx_tag_scope,
            "tag": nsx_tag,
        },
        {
            "scope": "tier",
            "tag": "db",
        },
    ])
#
# This part of the example shows creating Groups with dynamic membership
# criteria
#
# All Virtual machines with specific tag and scope
all_vms = nsxt.PolicyGroup("all_vms",
    display_name="All_VMs",
    description="Group consisting of ALL VMs",
    criterias=[{
        "conditions": [{
            "member_type": "VirtualMachine",
            "operator": "CONTAINS",
            "key": "Tag",
            "value": nsx_tag,
        }],
    }])
# All WEB VMs
web_group = nsxt.PolicyGroup("web_group",
    display_name="Web_VMs",
    description="Group consisting of Web VMs",
    criterias=[{
        "conditions": [{
            "member_type": "VirtualMachine",
            "operator": "CONTAINS",
            "key": "Tag",
            "value": "web",
        }],
    }],
    tags=[{
        "scope": nsx_tag_scope,
        "tag": nsx_tag,
    }])
# All App VMs
app_group = nsxt.PolicyGroup("app_group",
    display_name="App_VMs",
    description="Group consisting of App VMs",
    criterias=[{
        "conditions": [{
            "member_type": "VirtualMachine",
            "operator": "CONTAINS",
            "key": "Tag",
            "value": "app",
        }],
    }],
    tags=[{
        "scope": nsx_tag_scope,
        "tag": nsx_tag,
    }])
# All DB VMs
db_group = nsxt.PolicyGroup("db_group",
    display_name="DB_VMs",
    description="Group consisting of DB VMs",
    criterias=[{
        "conditions": [{
            "member_type": "VirtualMachine",
            "operator": "CONTAINS",
            "key": "Tag",
            "value": "db",
        }],
    }],
    tags=[{
        "scope": nsx_tag_scope,
        "tag": nsx_tag,
    }])
# Static Group of IP addresses
ip_set = nsxt.PolicyGroup("ip_set",
    display_name="external_IPs",
    description="Group containing all external IPs",
    criterias=[{
        "ipaddress_expression": {
            "ip_addresses": [
                "211.1.1.1",
                "212.1.1.1",
                "192.168.1.1-192.168.1.100",
            ],
        },
    }],
    tags=[{
        "scope": nsx_tag_scope,
        "tag": nsx_tag,
    }])
#
# An example for Service for App that listens on port 8443
#
app_service = nsxt.PolicyService("app_service",
    display_name="app_service_8443",
    description="Service for App that listens on port 8443",
    l4_port_set_entries=[{
        "description": "TCP Port 8443",
        "protocol": "TCP",
        "destination_ports": ["8443"],
    }],
    tags=[{
        "scope": nsx_tag_scope,
        "tag": nsx_tag,
    }])
#
# Here we have examples of create functions for Services
#
https = nsxt.get_policy_service(display_name="HTTPS")
mysql = nsxt.get_policy_service(display_name="MySQL")
ssh = nsxt.get_policy_service(display_name="SSH")
#
# In this section, we have example to create Firewall sections and rules
# All rules in this section will be applied to VMs that are part of the
# Gropus we created earlier
#
firewall_section = nsxt.PolicySecurityPolicy("firewall_section",
    display_name="DFW Section",
    description="Firewall section created by Pulumi",
    scopes=[all_vms.path],
    category="Application",
    locked=False,
    stateful=True,
    tags=[{
        "scope": nsx_tag_scope,
        "tag": nsx_tag,
    }],
    rules=[
        {
            "display_name": "Allow HTTPS",
            "description": "In going rule",
            "action": "ALLOW",
            "logged": False,
            "ip_version": "IPV4",
            "destination_groups": [web_group.path],
            "services": [https.path],
        },
        {
            "display_name": "Allow SSH",
            "description": "In going rule",
            "action": "ALLOW",
            "logged": False,
            "ip_version": "IPV4",
            "destination_groups": [web_group.path],
            "services": [ssh.path],
        },
        {
            "display_name": "Allow Web to App",
            "description": "Web to App communication",
            "action": "ALLOW",
            "logged": False,
            "ip_version": "IPV4",
            "source_groups": [web_group.path],
            "destination_groups": [app_group.path],
            "services": [app_service.path],
        },
        {
            "display_name": "Allow App to DB",
            "description": "App to DB communication",
            "action": "ALLOW",
            "logged": False,
            "ip_version": "IPV4",
            "source_groups": [app_group.path],
            "destination_groups": [db_group.path],
            "services": [mysql.path],
        },
        {
            "display_name": "Allow Infrastructure",
            "description": "Allow DNS and Management servers",
            "action": "ALLOW",
            "logged": True,
            "ip_version": "IPV4",
            "source_groups": [ip_set.path],
            "destination_groups": [all_vms.path],
        },
        {
            "display_name": "Allow out",
            "description": "Outgoing rule",
            "action": "ALLOW",
            "logged": True,
            "ip_version": "IPV4",
            "source_groups": [all_vms.path],
        },
        {
            "display_name": "Deny ANY",
            "description": "Default Deny the traffic",
            "action": "REJECT",
            "logged": True,
            "ip_version": "IPV4",
        },
    ])
#
# Here we have examples for creating NAT rules. The example here assumes
# the Web IP addresses are reachable from outside and no NAT is required.
#
rule1 = nsxt.PolicyNatRule("rule1",
    display_name="App 1-to-1 In",
    action="SNAT",
    translated_networks=["102.10.22.1"],
    source_networks=["12.12.2.0/24"],
    gateway_path=t1_gateway.path,
    tags=[{
        "scope": nsx_tag_scope,
        "tag": nsx_tag,
    }])
rule2 = nsxt.PolicyNatRule("rule2",
    display_name="App 1-to-1 Out",
    action="DNAT",
    translated_networks=["102.10.22.2"],
    destination_networks=["102.10.22.1/32"],
    source_networks=["12.12.2.0/24"],
    gateway_path=t1_gateway.path,
    tags=[{
        "scope": nsx_tag_scope,
        "tag": nsx_tag,
    }])
rule3 = nsxt.PolicyNatRule("rule3",
    display_name="DB 1-to-1 In",
    action="SNAT",
    translated_networks=["102.10.23.1"],
    source_networks=["12.12.3.0/24"],
    gateway_path=t1_gateway.path,
    tags=[{
        "scope": nsx_tag_scope,
        "tag": nsx_tag,
    }])
rule4 = nsxt.PolicyNatRule("rule4",
    display_name="App 1-to-1 Out",
    action="DNAT",
    translated_networks=["102.10.23.3"],
    destination_networks=["102.10.23.1/32"],
    source_networks=["12.12.3.0/24"],
    gateway_path=t1_gateway.path,
    tags=[{
        "scope": nsx_tag_scope,
        "tag": nsx_tag,
    }])
# The 3 VMs available in the NSX Inventory
web_vm = nsxt.get_policy_vm(display_name="web-vm")
app_vm = nsxt.get_policy_vm(display_name="app-vm")
db_vm = nsxt.get_policy_vm(display_name="db-vm")
# Assign the right tags to the VMs so that they get included in the
# dynamic groups created above
web_vm_tag = nsxt.PolicyVmTags("web_vm_tag",
    instance_id=web_vm.instance_id,
    tags=[
        {
            "scope": "tier",
            "tag": "web",
        },
        {
            "scope": nsx_tag_scope,
            "tag": nsx_tag,
        },
    ])
app_vm_tag = nsxt.PolicyVmTags("app_vm_tag",
    instance_id=app_vm.instance_id,
    tags=[
        {
            "scope": "tier",
            "tag": "app",
        },
        {
            "scope": nsx_tag_scope,
            "tag": nsx_tag,
        },
    ])
db_vm_tag = nsxt.PolicyVmTags("db_vm_tag",
    instance_id=db_vm.instance_id,
    tags=[
        {
            "scope": "tier",
            "tag": "db",
        },
        {
            "scope": nsx_tag_scope,
            "tag": nsx_tag,
        },
    ])
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    nsxt:allowUnverifiedSsl:
        value: true
    nsxt:host:
        value: 'TODO: var.nsx_manager'
    nsxt:maxRetries:
        value: 10
    nsxt:password:
        value: 'TODO: var.nsx_password'
    nsxt:retryMaxDelay:
        value: 5000
    nsxt:retryMinDelay:
        value: 500
    nsxt:retryOnStatusCodes:
        value:
            - 429
    nsxt:username:
        value: 'TODO: var.nsx_username'

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Nsxt = Pulumi.Nsxt;

return await Deployment.RunAsync(() =>
{
    var config = new Config();
    var nsxTagScope = config.Get("nsxTagScope") ?? "project";
    var nsxTag = config.Get("nsxTag") ?? "pulumi-demo";
    //
    // This part of the example shows some functions we will need to refer to
    // later in the .tf file. They include the transport zone, tier 0 router and
    // edge cluster.
    // There Tier-0 (T0) Gateway is considered a "provider" router that is pre-created
    // by the NSX Admin. A T0 Gateway is used for north/south connectivity between
    // the logical networking space and the physical networking space. Many Tier1
    // Gateways will be connected to the T0 Gateway
    //
    var demo = Nsxt.GetPolicyEdgeCluster.Invoke(new()
    {
        DisplayName = "Edge-Cluster",
    });

    var overlayTz = Nsxt.GetPolicyTransportZone.Invoke(new()
    {
        DisplayName = "Overlay-TZ",
    });

    var t0Gateway = Nsxt.GetPolicyTier0Gateway.Invoke(new()
    {
        DisplayName = "TF-T0-Gateway",
    });

    //
    // Create a DHCP Profile that is used later
    // Note, this resource is only in NSX 3.0.0+
    var tierDhcp = new Nsxt.PolicyDhcpServer("tier_dhcp", new()
    {
        DisplayName = "tier_dhcp",
        Description = "DHCP server servicing all 3 Segments",
        ServerAddresses = new[]
        {
            "12.12.99.2/24",
        },
    });

    //
    // In this part of the example, the settings required to create a Tier1 Gateway
    // are defined. In NSX a Tier1 Gateway is often used on a per user, tenant,
    // department or application basis. Each application may have it's own Tier1
    // Gateway. The Tier1 Gateway provides the default gateway for virtual machines
    // connected to the Segments on the Tier1 Gateway
    //
    var t1Gateway = new Nsxt.PolicyTier1Gateway("t1_gateway", new()
    {
        DisplayName = "TF_T1",
        Description = "Tier1 provisioned by Pulumi",
        EdgeClusterPath = demo.Apply(getPolicyEdgeClusterResult => getPolicyEdgeClusterResult.Path),
        DhcpConfigPath = tierDhcp.Path,
        FailoverMode = "PREEMPTIVE",
        DefaultRuleLogging = false,
        EnableFirewall = true,
        EnableStandbyRelocation = false,
        ForceWhitelisting = false,
        Tier0Path = t0Gateway.Apply(getPolicyTier0GatewayResult => getPolicyTier0GatewayResult.Path),
        RouteAdvertisementTypes = new[]
        {
            "TIER1_STATIC_ROUTES",
            "TIER1_CONNECTED",
        },
        PoolAllocation = "ROUTING",
        Tags = new[]
        {
            new Nsxt.Inputs.PolicyTier1GatewayTagArgs
            {
                Scope = nsxTagScope,
                Tag = nsxTag,
            },
        },
        RouteAdvertisementRules = new[]
        {
            new Nsxt.Inputs.PolicyTier1GatewayRouteAdvertisementRuleArgs
            {
                Name = "rule1",
                Action = "DENY",
                Subnets = new[]
                {
                    "20.0.0.0/24",
                    "21.0.0.0/24",
                },
                PrefixOperator = "GE",
                RouteAdvertisementTypes = new[]
                {
                    "TIER1_CONNECTED",
                },
            },
        },
    });

    //
    // This shows the settings required to create NSX Segment (Logical Switch) to
    // which you can attach Virtual Machines (VMs)
    //
    var web = new Nsxt.PolicySegment("web", new()
    {
        DisplayName = "web-tier",
        Description = "Pulumi provisioned Web Segment",
        ConnectivityPath = t1Gateway.Path,
        TransportZonePath = overlayTz.Apply(getPolicyTransportZoneResult => getPolicyTransportZoneResult.Path),
        Subnets = new[]
        {
            new Nsxt.Inputs.PolicySegmentSubnetArgs
            {
                Cidr = "12.12.1.1/24",
                DhcpRanges = new[]
                {
                    "12.12.1.100-12.12.1.160",
                },
                DhcpV4Config = new Nsxt.Inputs.PolicySegmentSubnetDhcpV4ConfigArgs
                {
                    ServerAddress = "12.12.1.2/24",
                    LeaseTime = 36000,
                    DhcpOption121s = new[]
                    {
                        new Nsxt.Inputs.PolicySegmentSubnetDhcpV4ConfigDhcpOption121Args
                        {
                            Network = "6.6.6.0/24",
                            NextHop = "1.1.1.21",
                        },
                    },
                },
            },
        },
        Tags = new[]
        {
            new Nsxt.Inputs.PolicySegmentTagArgs
            {
                Scope = nsxTagScope,
                Tag = nsxTag,
            },
            new Nsxt.Inputs.PolicySegmentTagArgs
            {
                Scope = "tier",
                Tag = "web",
            },
        },
    });

    var app = new Nsxt.PolicySegment("app", new()
    {
        DisplayName = "app-tier",
        Description = "Pulumi provisioned App Segment",
        ConnectivityPath = t1Gateway.Path,
        TransportZonePath = overlayTz.Apply(getPolicyTransportZoneResult => getPolicyTransportZoneResult.Path),
        Subnets = new[]
        {
            new Nsxt.Inputs.PolicySegmentSubnetArgs
            {
                Cidr = "12.12.2.1/24",
                DhcpRanges = new[]
                {
                    "12.12.2.100-12.12.2.160",
                },
                DhcpV4Config = new Nsxt.Inputs.PolicySegmentSubnetDhcpV4ConfigArgs
                {
                    ServerAddress = "12.12.2.2/24",
                    LeaseTime = 36000,
                    DhcpOption121s = new[]
                    {
                        new Nsxt.Inputs.PolicySegmentSubnetDhcpV4ConfigDhcpOption121Args
                        {
                            Network = "6.6.6.0/24",
                            NextHop = "1.1.1.21",
                        },
                    },
                },
            },
        },
        Tags = new[]
        {
            new Nsxt.Inputs.PolicySegmentTagArgs
            {
                Scope = nsxTagScope,
                Tag = nsxTag,
            },
            new Nsxt.Inputs.PolicySegmentTagArgs
            {
                Scope = "tier",
                Tag = "app",
            },
        },
    });

    var db = new Nsxt.PolicySegment("db", new()
    {
        DisplayName = "db-tier",
        Description = "Pulumi provisioned DB Segment",
        ConnectivityPath = t1Gateway.Path,
        TransportZonePath = overlayTz.Apply(getPolicyTransportZoneResult => getPolicyTransportZoneResult.Path),
        Subnets = new[]
        {
            new Nsxt.Inputs.PolicySegmentSubnetArgs
            {
                Cidr = "12.12.3.1/24",
                DhcpRanges = new[]
                {
                    "12.12.3.100-12.12.3.160",
                },
                DhcpV4Config = new Nsxt.Inputs.PolicySegmentSubnetDhcpV4ConfigArgs
                {
                    ServerAddress = "12.12.3.2/24",
                    LeaseTime = 36000,
                    DhcpOption121s = new[]
                    {
                        new Nsxt.Inputs.PolicySegmentSubnetDhcpV4ConfigDhcpOption121Args
                        {
                            Network = "6.6.6.0/24",
                            NextHop = "1.1.1.21",
                        },
                    },
                },
            },
        },
        Tags = new[]
        {
            new Nsxt.Inputs.PolicySegmentTagArgs
            {
                Scope = nsxTagScope,
                Tag = nsxTag,
            },
            new Nsxt.Inputs.PolicySegmentTagArgs
            {
                Scope = "tier",
                Tag = "db",
            },
        },
    });

    //
    // This part of the example shows creating Groups with dynamic membership
    // criteria
    //
    // All Virtual machines with specific tag and scope
    var allVms = new Nsxt.PolicyGroup("all_vms", new()
    {
        DisplayName = "All_VMs",
        Description = "Group consisting of ALL VMs",
        Criterias = new[]
        {
            new Nsxt.Inputs.PolicyGroupCriteriaArgs
            {
                Conditions = new[]
                {
                    new Nsxt.Inputs.PolicyGroupCriteriaConditionArgs
                    {
                        MemberType = "VirtualMachine",
                        Operator = "CONTAINS",
                        Key = "Tag",
                        Value = nsxTag,
                    },
                },
            },
        },
    });

    // All WEB VMs
    var webGroup = new Nsxt.PolicyGroup("web_group", new()
    {
        DisplayName = "Web_VMs",
        Description = "Group consisting of Web VMs",
        Criterias = new[]
        {
            new Nsxt.Inputs.PolicyGroupCriteriaArgs
            {
                Conditions = new[]
                {
                    new Nsxt.Inputs.PolicyGroupCriteriaConditionArgs
                    {
                        MemberType = "VirtualMachine",
                        Operator = "CONTAINS",
                        Key = "Tag",
                        Value = "web",
                    },
                },
            },
        },
        Tags = new[]
        {
            new Nsxt.Inputs.PolicyGroupTagArgs
            {
                Scope = nsxTagScope,
                Tag = nsxTag,
            },
        },
    });

    // All App VMs
    var appGroup = new Nsxt.PolicyGroup("app_group", new()
    {
        DisplayName = "App_VMs",
        Description = "Group consisting of App VMs",
        Criterias = new[]
        {
            new Nsxt.Inputs.PolicyGroupCriteriaArgs
            {
                Conditions = new[]
                {
                    new Nsxt.Inputs.PolicyGroupCriteriaConditionArgs
                    {
                        MemberType = "VirtualMachine",
                        Operator = "CONTAINS",
                        Key = "Tag",
                        Value = "app",
                    },
                },
            },
        },
        Tags = new[]
        {
            new Nsxt.Inputs.PolicyGroupTagArgs
            {
                Scope = nsxTagScope,
                Tag = nsxTag,
            },
        },
    });

    // All DB VMs
    var dbGroup = new Nsxt.PolicyGroup("db_group", new()
    {
        DisplayName = "DB_VMs",
        Description = "Group consisting of DB VMs",
        Criterias = new[]
        {
            new Nsxt.Inputs.PolicyGroupCriteriaArgs
            {
                Conditions = new[]
                {
                    new Nsxt.Inputs.PolicyGroupCriteriaConditionArgs
                    {
                        MemberType = "VirtualMachine",
                        Operator = "CONTAINS",
                        Key = "Tag",
                        Value = "db",
                    },
                },
            },
        },
        Tags = new[]
        {
            new Nsxt.Inputs.PolicyGroupTagArgs
            {
                Scope = nsxTagScope,
                Tag = nsxTag,
            },
        },
    });

    // Static Group of IP addresses
    var ipSet = new Nsxt.PolicyGroup("ip_set", new()
    {
        DisplayName = "external_IPs",
        Description = "Group containing all external IPs",
        Criterias = new[]
        {
            new Nsxt.Inputs.PolicyGroupCriteriaArgs
            {
                IpaddressExpression = new Nsxt.Inputs.PolicyGroupCriteriaIpaddressExpressionArgs
                {
                    IpAddresses = new[]
                    {
                        "211.1.1.1",
                        "212.1.1.1",
                        "192.168.1.1-192.168.1.100",
                    },
                },
            },
        },
        Tags = new[]
        {
            new Nsxt.Inputs.PolicyGroupTagArgs
            {
                Scope = nsxTagScope,
                Tag = nsxTag,
            },
        },
    });

    //
    // An example for Service for App that listens on port 8443
    //
    var appService = new Nsxt.PolicyService("app_service", new()
    {
        DisplayName = "app_service_8443",
        Description = "Service for App that listens on port 8443",
        L4PortSetEntries = new[]
        {
            new Nsxt.Inputs.PolicyServiceL4PortSetEntryArgs
            {
                Description = "TCP Port 8443",
                Protocol = "TCP",
                DestinationPorts = new[]
                {
                    "8443",
                },
            },
        },
        Tags = new[]
        {
            new Nsxt.Inputs.PolicyServiceTagArgs
            {
                Scope = nsxTagScope,
                Tag = nsxTag,
            },
        },
    });

    //
    // Here we have examples of create functions for Services
    //
    var https = Nsxt.GetPolicyService.Invoke(new()
    {
        DisplayName = "HTTPS",
    });

    var mysql = Nsxt.GetPolicyService.Invoke(new()
    {
        DisplayName = "MySQL",
    });

    var ssh = Nsxt.GetPolicyService.Invoke(new()
    {
        DisplayName = "SSH",
    });

    //
    // In this section, we have example to create Firewall sections and rules
    // All rules in this section will be applied to VMs that are part of the
    // Gropus we created earlier
    //
    var firewallSection = new Nsxt.PolicySecurityPolicy("firewall_section", new()
    {
        DisplayName = "DFW Section",
        Description = "Firewall section created by Pulumi",
        Scopes = new[]
        {
            allVms.Path,
        },
        Category = "Application",
        Locked = false,
        Stateful = true,
        Tags = new[]
        {
            new Nsxt.Inputs.PolicySecurityPolicyTagArgs
            {
                Scope = nsxTagScope,
                Tag = nsxTag,
            },
        },
        Rules = new[]
        {
            new Nsxt.Inputs.PolicySecurityPolicyRuleArgs
            {
                DisplayName = "Allow HTTPS",
                Description = "In going rule",
                Action = "ALLOW",
                Logged = false,
                IpVersion = "IPV4",
                DestinationGroups = new[]
                {
                    webGroup.Path,
                },
                Services = new[]
                {
                    https.Apply(getPolicyServiceResult => getPolicyServiceResult.Path),
                },
            },
            new Nsxt.Inputs.PolicySecurityPolicyRuleArgs
            {
                DisplayName = "Allow SSH",
                Description = "In going rule",
                Action = "ALLOW",
                Logged = false,
                IpVersion = "IPV4",
                DestinationGroups = new[]
                {
                    webGroup.Path,
                },
                Services = new[]
                {
                    ssh.Apply(getPolicyServiceResult => getPolicyServiceResult.Path),
                },
            },
            new Nsxt.Inputs.PolicySecurityPolicyRuleArgs
            {
                DisplayName = "Allow Web to App",
                Description = "Web to App communication",
                Action = "ALLOW",
                Logged = false,
                IpVersion = "IPV4",
                SourceGroups = new[]
                {
                    webGroup.Path,
                },
                DestinationGroups = new[]
                {
                    appGroup.Path,
                },
                Services = new[]
                {
                    appService.Path,
                },
            },
            new Nsxt.Inputs.PolicySecurityPolicyRuleArgs
            {
                DisplayName = "Allow App to DB",
                Description = "App to DB communication",
                Action = "ALLOW",
                Logged = false,
                IpVersion = "IPV4",
                SourceGroups = new[]
                {
                    appGroup.Path,
                },
                DestinationGroups = new[]
                {
                    dbGroup.Path,
                },
                Services = new[]
                {
                    mysql.Apply(getPolicyServiceResult => getPolicyServiceResult.Path),
                },
            },
            new Nsxt.Inputs.PolicySecurityPolicyRuleArgs
            {
                DisplayName = "Allow Infrastructure",
                Description = "Allow DNS and Management servers",
                Action = "ALLOW",
                Logged = true,
                IpVersion = "IPV4",
                SourceGroups = new[]
                {
                    ipSet.Path,
                },
                DestinationGroups = new[]
                {
                    allVms.Path,
                },
            },
            new Nsxt.Inputs.PolicySecurityPolicyRuleArgs
            {
                DisplayName = "Allow out",
                Description = "Outgoing rule",
                Action = "ALLOW",
                Logged = true,
                IpVersion = "IPV4",
                SourceGroups = new[]
                {
                    allVms.Path,
                },
            },
            new Nsxt.Inputs.PolicySecurityPolicyRuleArgs
            {
                DisplayName = "Deny ANY",
                Description = "Default Deny the traffic",
                Action = "REJECT",
                Logged = true,
                IpVersion = "IPV4",
            },
        },
    });

    //
    // Here we have examples for creating NAT rules. The example here assumes
    // the Web IP addresses are reachable from outside and no NAT is required.
    //
    var rule1 = new Nsxt.PolicyNatRule("rule1", new()
    {
        DisplayName = "App 1-to-1 In",
        Action = "SNAT",
        TranslatedNetworks = new[]
        {
            "102.10.22.1",
        },
        SourceNetworks = new[]
        {
            "12.12.2.0/24",
        },
        GatewayPath = t1Gateway.Path,
        Tags = new[]
        {
            new Nsxt.Inputs.PolicyNatRuleTagArgs
            {
                Scope = nsxTagScope,
                Tag = nsxTag,
            },
        },
    });

    var rule2 = new Nsxt.PolicyNatRule("rule2", new()
    {
        DisplayName = "App 1-to-1 Out",
        Action = "DNAT",
        TranslatedNetworks = new[]
        {
            "102.10.22.2",
        },
        DestinationNetworks = new[]
        {
            "102.10.22.1/32",
        },
        SourceNetworks = new[]
        {
            "12.12.2.0/24",
        },
        GatewayPath = t1Gateway.Path,
        Tags = new[]
        {
            new Nsxt.Inputs.PolicyNatRuleTagArgs
            {
                Scope = nsxTagScope,
                Tag = nsxTag,
            },
        },
    });

    var rule3 = new Nsxt.PolicyNatRule("rule3", new()
    {
        DisplayName = "DB 1-to-1 In",
        Action = "SNAT",
        TranslatedNetworks = new[]
        {
            "102.10.23.1",
        },
        SourceNetworks = new[]
        {
            "12.12.3.0/24",
        },
        GatewayPath = t1Gateway.Path,
        Tags = new[]
        {
            new Nsxt.Inputs.PolicyNatRuleTagArgs
            {
                Scope = nsxTagScope,
                Tag = nsxTag,
            },
        },
    });

    var rule4 = new Nsxt.PolicyNatRule("rule4", new()
    {
        DisplayName = "App 1-to-1 Out",
        Action = "DNAT",
        TranslatedNetworks = new[]
        {
            "102.10.23.3",
        },
        DestinationNetworks = new[]
        {
            "102.10.23.1/32",
        },
        SourceNetworks = new[]
        {
            "12.12.3.0/24",
        },
        GatewayPath = t1Gateway.Path,
        Tags = new[]
        {
            new Nsxt.Inputs.PolicyNatRuleTagArgs
            {
                Scope = nsxTagScope,
                Tag = nsxTag,
            },
        },
    });

    // The 3 VMs available in the NSX Inventory
    var webVm = Nsxt.GetPolicyVm.Invoke(new()
    {
        DisplayName = "web-vm",
    });

    var appVm = Nsxt.GetPolicyVm.Invoke(new()
    {
        DisplayName = "app-vm",
    });

    var dbVm = Nsxt.GetPolicyVm.Invoke(new()
    {
        DisplayName = "db-vm",
    });

    // Assign the right tags to the VMs so that they get included in the
    // dynamic groups created above
    var webVmTag = new Nsxt.PolicyVmTags("web_vm_tag", new()
    {
        InstanceId = webVm.Apply(getPolicyVmResult => getPolicyVmResult.InstanceId),
        Tags = new[]
        {
            new Nsxt.Inputs.PolicyVmTagsTagArgs
            {
                Scope = "tier",
                Tag = "web",
            },
            new Nsxt.Inputs.PolicyVmTagsTagArgs
            {
                Scope = nsxTagScope,
                Tag = nsxTag,
            },
        },
    });

    var appVmTag = new Nsxt.PolicyVmTags("app_vm_tag", new()
    {
        InstanceId = appVm.Apply(getPolicyVmResult => getPolicyVmResult.InstanceId),
        Tags = new[]
        {
            new Nsxt.Inputs.PolicyVmTagsTagArgs
            {
                Scope = "tier",
                Tag = "app",
            },
            new Nsxt.Inputs.PolicyVmTagsTagArgs
            {
                Scope = nsxTagScope,
                Tag = nsxTag,
            },
        },
    });

    var dbVmTag = new Nsxt.PolicyVmTags("db_vm_tag", new()
    {
        InstanceId = dbVm.Apply(getPolicyVmResult => getPolicyVmResult.InstanceId),
        Tags = new[]
        {
            new Nsxt.Inputs.PolicyVmTagsTagArgs
            {
                Scope = "tier",
                Tag = "db",
            },
            new Nsxt.Inputs.PolicyVmTagsTagArgs
            {
                Scope = nsxTagScope,
                Tag = nsxTag,
            },
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
    nsxt:allowUnverifiedSsl:
        value: true
    nsxt:host:
        value: 'TODO: var.nsx_manager'
    nsxt:maxRetries:
        value: 10
    nsxt:password:
        value: 'TODO: var.nsx_password'
    nsxt:retryMaxDelay:
        value: 5000
    nsxt:retryMinDelay:
        value: 500
    nsxt:retryOnStatusCodes:
        value:
            - 429
    nsxt:username:
        value: 'TODO: var.nsx_username'

```
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/nsxt/v3/nsxt"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")
		nsxTagScope := "project"
		if param := cfg.Get("nsxTagScope"); param != "" {
			nsxTagScope = param
		}
		nsxTag := "pulumi-demo"
		if param := cfg.Get("nsxTag"); param != "" {
			nsxTag = param
		}
		// This part of the example shows some functions we will need to refer to
		// later in the .tf file. They include the transport zone, tier 0 router and
		// edge cluster.
		// There Tier-0 (T0) Gateway is considered a "provider" router that is pre-created
		// by the NSX Admin. A T0 Gateway is used for north/south connectivity between
		// the logical networking space and the physical networking space. Many Tier1
		// Gateways will be connected to the T0 Gateway
		demo, err := nsxt.LookupPolicyEdgeCluster(ctx, &nsxt.LookupPolicyEdgeClusterArgs{
			DisplayName: pulumi.StringRef("Edge-Cluster"),
		}, nil)
		if err != nil {
			return err
		}
		overlayTz, err := nsxt.LookupPolicyTransportZone(ctx, &nsxt.LookupPolicyTransportZoneArgs{
			DisplayName: pulumi.StringRef("Overlay-TZ"),
		}, nil)
		if err != nil {
			return err
		}
		t0Gateway, err := nsxt.LookupPolicyTier0Gateway(ctx, &nsxt.LookupPolicyTier0GatewayArgs{
			DisplayName: pulumi.StringRef("TF-T0-Gateway"),
		}, nil)
		if err != nil {
			return err
		}
		// Create a DHCP Profile that is used later
		// Note, this resource is only in NSX 3.0.0+
		tierDhcp, err := nsxt.NewPolicyDhcpServer(ctx, "tier_dhcp", &nsxt.PolicyDhcpServerArgs{
			DisplayName: pulumi.String("tier_dhcp"),
			Description: pulumi.String("DHCP server servicing all 3 Segments"),
			ServerAddresses: pulumi.StringArray{
				pulumi.String("12.12.99.2/24"),
			},
		})
		if err != nil {
			return err
		}
		// In this part of the example, the settings required to create a Tier1 Gateway
		// are defined. In NSX a Tier1 Gateway is often used on a per user, tenant,
		// department or application basis. Each application may have it's own Tier1
		// Gateway. The Tier1 Gateway provides the default gateway for virtual machines
		// connected to the Segments on the Tier1 Gateway
		t1Gateway, err := nsxt.NewPolicyTier1Gateway(ctx, "t1_gateway", &nsxt.PolicyTier1GatewayArgs{
			DisplayName:             pulumi.String("TF_T1"),
			Description:             pulumi.String("Tier1 provisioned by Pulumi"),
			EdgeClusterPath:         pulumi.String(demo.Path),
			DhcpConfigPath:          tierDhcp.Path,
			FailoverMode:            pulumi.String("PREEMPTIVE"),
			DefaultRuleLogging:      pulumi.Bool(false),
			EnableFirewall:          pulumi.Bool(true),
			EnableStandbyRelocation: pulumi.Bool(false),
			ForceWhitelisting:       pulumi.Bool(false),
			Tier0Path:               pulumi.String(t0Gateway.Path),
			RouteAdvertisementTypes: pulumi.StringArray{
				pulumi.String("TIER1_STATIC_ROUTES"),
				pulumi.String("TIER1_CONNECTED"),
			},
			PoolAllocation: pulumi.String("ROUTING"),
			Tags: nsxt.PolicyTier1GatewayTagArray{
				&nsxt.PolicyTier1GatewayTagArgs{
					Scope: pulumi.String(nsxTagScope),
					Tag:   pulumi.String(nsxTag),
				},
			},
			RouteAdvertisementRules: nsxt.PolicyTier1GatewayRouteAdvertisementRuleArray{
				&nsxt.PolicyTier1GatewayRouteAdvertisementRuleArgs{
					Name:   pulumi.String("rule1"),
					Action: pulumi.String("DENY"),
					Subnets: pulumi.StringArray{
						pulumi.String("20.0.0.0/24"),
						pulumi.String("21.0.0.0/24"),
					},
					PrefixOperator: pulumi.String("GE"),
					RouteAdvertisementTypes: pulumi.StringArray{
						pulumi.String("TIER1_CONNECTED"),
					},
				},
			},
		})
		if err != nil {
			return err
		}
		// This shows the settings required to create NSX Segment (Logical Switch) to
		// which you can attach Virtual Machines (VMs)
		_, err = nsxt.NewPolicySegment(ctx, "web", &nsxt.PolicySegmentArgs{
			DisplayName:       pulumi.String("web-tier"),
			Description:       pulumi.String("Pulumi provisioned Web Segment"),
			ConnectivityPath:  t1Gateway.Path,
			TransportZonePath: pulumi.String(overlayTz.Path),
			Subnets: nsxt.PolicySegmentSubnetArray{
				&nsxt.PolicySegmentSubnetArgs{
					Cidr: pulumi.String("12.12.1.1/24"),
					DhcpRanges: pulumi.StringArray{
						pulumi.String("12.12.1.100-12.12.1.160"),
					},
					DhcpV4Config: &nsxt.PolicySegmentSubnetDhcpV4ConfigArgs{
						ServerAddress: pulumi.String("12.12.1.2/24"),
						LeaseTime:     pulumi.Float64(36000),
						DhcpOption121s: nsxt.PolicySegmentSubnetDhcpV4ConfigDhcpOption121Array{
							&nsxt.PolicySegmentSubnetDhcpV4ConfigDhcpOption121Args{
								Network: pulumi.String("6.6.6.0/24"),
								NextHop: pulumi.String("1.1.1.21"),
							},
						},
					},
				},
			},
			Tags: nsxt.PolicySegmentTagArray{
				&nsxt.PolicySegmentTagArgs{
					Scope: pulumi.String(nsxTagScope),
					Tag:   pulumi.String(nsxTag),
				},
				&nsxt.PolicySegmentTagArgs{
					Scope: pulumi.String("tier"),
					Tag:   pulumi.String("web"),
				},
			},
		})
		if err != nil {
			return err
		}
		_, err = nsxt.NewPolicySegment(ctx, "app", &nsxt.PolicySegmentArgs{
			DisplayName:       pulumi.String("app-tier"),
			Description:       pulumi.String("Pulumi provisioned App Segment"),
			ConnectivityPath:  t1Gateway.Path,
			TransportZonePath: pulumi.String(overlayTz.Path),
			Subnets: nsxt.PolicySegmentSubnetArray{
				&nsxt.PolicySegmentSubnetArgs{
					Cidr: pulumi.String("12.12.2.1/24"),
					DhcpRanges: pulumi.StringArray{
						pulumi.String("12.12.2.100-12.12.2.160"),
					},
					DhcpV4Config: &nsxt.PolicySegmentSubnetDhcpV4ConfigArgs{
						ServerAddress: pulumi.String("12.12.2.2/24"),
						LeaseTime:     pulumi.Float64(36000),
						DhcpOption121s: nsxt.PolicySegmentSubnetDhcpV4ConfigDhcpOption121Array{
							&nsxt.PolicySegmentSubnetDhcpV4ConfigDhcpOption121Args{
								Network: pulumi.String("6.6.6.0/24"),
								NextHop: pulumi.String("1.1.1.21"),
							},
						},
					},
				},
			},
			Tags: nsxt.PolicySegmentTagArray{
				&nsxt.PolicySegmentTagArgs{
					Scope: pulumi.String(nsxTagScope),
					Tag:   pulumi.String(nsxTag),
				},
				&nsxt.PolicySegmentTagArgs{
					Scope: pulumi.String("tier"),
					Tag:   pulumi.String("app"),
				},
			},
		})
		if err != nil {
			return err
		}
		_, err = nsxt.NewPolicySegment(ctx, "db", &nsxt.PolicySegmentArgs{
			DisplayName:       pulumi.String("db-tier"),
			Description:       pulumi.String("Pulumi provisioned DB Segment"),
			ConnectivityPath:  t1Gateway.Path,
			TransportZonePath: pulumi.String(overlayTz.Path),
			Subnets: nsxt.PolicySegmentSubnetArray{
				&nsxt.PolicySegmentSubnetArgs{
					Cidr: pulumi.String("12.12.3.1/24"),
					DhcpRanges: pulumi.StringArray{
						pulumi.String("12.12.3.100-12.12.3.160"),
					},
					DhcpV4Config: &nsxt.PolicySegmentSubnetDhcpV4ConfigArgs{
						ServerAddress: pulumi.String("12.12.3.2/24"),
						LeaseTime:     pulumi.Float64(36000),
						DhcpOption121s: nsxt.PolicySegmentSubnetDhcpV4ConfigDhcpOption121Array{
							&nsxt.PolicySegmentSubnetDhcpV4ConfigDhcpOption121Args{
								Network: pulumi.String("6.6.6.0/24"),
								NextHop: pulumi.String("1.1.1.21"),
							},
						},
					},
				},
			},
			Tags: nsxt.PolicySegmentTagArray{
				&nsxt.PolicySegmentTagArgs{
					Scope: pulumi.String(nsxTagScope),
					Tag:   pulumi.String(nsxTag),
				},
				&nsxt.PolicySegmentTagArgs{
					Scope: pulumi.String("tier"),
					Tag:   pulumi.String("db"),
				},
			},
		})
		if err != nil {
			return err
		}
		// This part of the example shows creating Groups with dynamic membership
		// criteria
		//
		// All Virtual machines with specific tag and scope
		allVms, err := nsxt.NewPolicyGroup(ctx, "all_vms", &nsxt.PolicyGroupArgs{
			DisplayName: pulumi.String("All_VMs"),
			Description: pulumi.String("Group consisting of ALL VMs"),
			Criterias: nsxt.PolicyGroupCriteriaArray{
				&nsxt.PolicyGroupCriteriaArgs{
					Conditions: nsxt.PolicyGroupCriteriaConditionArray{
						&nsxt.PolicyGroupCriteriaConditionArgs{
							MemberType: pulumi.String("VirtualMachine"),
							Operator:   pulumi.String("CONTAINS"),
							Key:        pulumi.String("Tag"),
							Value:      pulumi.String(nsxTag),
						},
					},
				},
			},
		})
		if err != nil {
			return err
		}
		// All WEB VMs
		webGroup, err := nsxt.NewPolicyGroup(ctx, "web_group", &nsxt.PolicyGroupArgs{
			DisplayName: pulumi.String("Web_VMs"),
			Description: pulumi.String("Group consisting of Web VMs"),
			Criterias: nsxt.PolicyGroupCriteriaArray{
				&nsxt.PolicyGroupCriteriaArgs{
					Conditions: nsxt.PolicyGroupCriteriaConditionArray{
						&nsxt.PolicyGroupCriteriaConditionArgs{
							MemberType: pulumi.String("VirtualMachine"),
							Operator:   pulumi.String("CONTAINS"),
							Key:        pulumi.String("Tag"),
							Value:      pulumi.String("web"),
						},
					},
				},
			},
			Tags: nsxt.PolicyGroupTagArray{
				&nsxt.PolicyGroupTagArgs{
					Scope: pulumi.String(nsxTagScope),
					Tag:   pulumi.String(nsxTag),
				},
			},
		})
		if err != nil {
			return err
		}
		// All App VMs
		appGroup, err := nsxt.NewPolicyGroup(ctx, "app_group", &nsxt.PolicyGroupArgs{
			DisplayName: pulumi.String("App_VMs"),
			Description: pulumi.String("Group consisting of App VMs"),
			Criterias: nsxt.PolicyGroupCriteriaArray{
				&nsxt.PolicyGroupCriteriaArgs{
					Conditions: nsxt.PolicyGroupCriteriaConditionArray{
						&nsxt.PolicyGroupCriteriaConditionArgs{
							MemberType: pulumi.String("VirtualMachine"),
							Operator:   pulumi.String("CONTAINS"),
							Key:        pulumi.String("Tag"),
							Value:      pulumi.String("app"),
						},
					},
				},
			},
			Tags: nsxt.PolicyGroupTagArray{
				&nsxt.PolicyGroupTagArgs{
					Scope: pulumi.String(nsxTagScope),
					Tag:   pulumi.String(nsxTag),
				},
			},
		})
		if err != nil {
			return err
		}
		// All DB VMs
		dbGroup, err := nsxt.NewPolicyGroup(ctx, "db_group", &nsxt.PolicyGroupArgs{
			DisplayName: pulumi.String("DB_VMs"),
			Description: pulumi.String("Group consisting of DB VMs"),
			Criterias: nsxt.PolicyGroupCriteriaArray{
				&nsxt.PolicyGroupCriteriaArgs{
					Conditions: nsxt.PolicyGroupCriteriaConditionArray{
						&nsxt.PolicyGroupCriteriaConditionArgs{
							MemberType: pulumi.String("VirtualMachine"),
							Operator:   pulumi.String("CONTAINS"),
							Key:        pulumi.String("Tag"),
							Value:      pulumi.String("db"),
						},
					},
				},
			},
			Tags: nsxt.PolicyGroupTagArray{
				&nsxt.PolicyGroupTagArgs{
					Scope: pulumi.String(nsxTagScope),
					Tag:   pulumi.String(nsxTag),
				},
			},
		})
		if err != nil {
			return err
		}
		// Static Group of IP addresses
		ipSet, err := nsxt.NewPolicyGroup(ctx, "ip_set", &nsxt.PolicyGroupArgs{
			DisplayName: pulumi.String("external_IPs"),
			Description: pulumi.String("Group containing all external IPs"),
			Criterias: nsxt.PolicyGroupCriteriaArray{
				&nsxt.PolicyGroupCriteriaArgs{
					IpaddressExpression: &nsxt.PolicyGroupCriteriaIpaddressExpressionArgs{
						IpAddresses: pulumi.StringArray{
							pulumi.String("211.1.1.1"),
							pulumi.String("212.1.1.1"),
							pulumi.String("192.168.1.1-192.168.1.100"),
						},
					},
				},
			},
			Tags: nsxt.PolicyGroupTagArray{
				&nsxt.PolicyGroupTagArgs{
					Scope: pulumi.String(nsxTagScope),
					Tag:   pulumi.String(nsxTag),
				},
			},
		})
		if err != nil {
			return err
		}
		// An example for Service for App that listens on port 8443
		appService, err := nsxt.NewPolicyService(ctx, "app_service", &nsxt.PolicyServiceArgs{
			DisplayName: pulumi.String("app_service_8443"),
			Description: pulumi.String("Service for App that listens on port 8443"),
			L4PortSetEntries: nsxt.PolicyServiceL4PortSetEntryArray{
				&nsxt.PolicyServiceL4PortSetEntryArgs{
					Description: pulumi.String("TCP Port 8443"),
					Protocol:    pulumi.String("TCP"),
					DestinationPorts: pulumi.StringArray{
						pulumi.String("8443"),
					},
				},
			},
			Tags: nsxt.PolicyServiceTagArray{
				&nsxt.PolicyServiceTagArgs{
					Scope: pulumi.String(nsxTagScope),
					Tag:   pulumi.String(nsxTag),
				},
			},
		})
		if err != nil {
			return err
		}
		// Here we have examples of create functions for Services
		https, err := nsxt.LookupPolicyService(ctx, &nsxt.LookupPolicyServiceArgs{
			DisplayName: pulumi.StringRef("HTTPS"),
		}, nil)
		if err != nil {
			return err
		}
		mysql, err := nsxt.LookupPolicyService(ctx, &nsxt.LookupPolicyServiceArgs{
			DisplayName: pulumi.StringRef("MySQL"),
		}, nil)
		if err != nil {
			return err
		}
		ssh, err := nsxt.LookupPolicyService(ctx, &nsxt.LookupPolicyServiceArgs{
			DisplayName: pulumi.StringRef("SSH"),
		}, nil)
		if err != nil {
			return err
		}
		// In this section, we have example to create Firewall sections and rules
		// All rules in this section will be applied to VMs that are part of the
		// Gropus we created earlier
		_, err = nsxt.NewPolicySecurityPolicy(ctx, "firewall_section", &nsxt.PolicySecurityPolicyArgs{
			DisplayName: pulumi.String("DFW Section"),
			Description: pulumi.String("Firewall section created by Pulumi"),
			Scopes: pulumi.StringArray{
				allVms.Path,
			},
			Category: pulumi.String("Application"),
			Locked:   pulumi.Bool(false),
			Stateful: pulumi.Bool(true),
			Tags: nsxt.PolicySecurityPolicyTagArray{
				&nsxt.PolicySecurityPolicyTagArgs{
					Scope: pulumi.String(nsxTagScope),
					Tag:   pulumi.String(nsxTag),
				},
			},
			Rules: nsxt.PolicySecurityPolicyRuleTypeArray{
				&nsxt.PolicySecurityPolicyRuleTypeArgs{
					DisplayName: pulumi.String("Allow HTTPS"),
					Description: pulumi.String("In going rule"),
					Action:      pulumi.String("ALLOW"),
					Logged:      pulumi.Bool(false),
					IpVersion:   pulumi.String("IPV4"),
					DestinationGroups: pulumi.StringArray{
						webGroup.Path,
					},
					Services: pulumi.StringArray{
						pulumi.String(https.Path),
					},
				},
				&nsxt.PolicySecurityPolicyRuleTypeArgs{
					DisplayName: pulumi.String("Allow SSH"),
					Description: pulumi.String("In going rule"),
					Action:      pulumi.String("ALLOW"),
					Logged:      pulumi.Bool(false),
					IpVersion:   pulumi.String("IPV4"),
					DestinationGroups: pulumi.StringArray{
						webGroup.Path,
					},
					Services: pulumi.StringArray{
						pulumi.String(ssh.Path),
					},
				},
				&nsxt.PolicySecurityPolicyRuleTypeArgs{
					DisplayName: pulumi.String("Allow Web to App"),
					Description: pulumi.String("Web to App communication"),
					Action:      pulumi.String("ALLOW"),
					Logged:      pulumi.Bool(false),
					IpVersion:   pulumi.String("IPV4"),
					SourceGroups: pulumi.StringArray{
						webGroup.Path,
					},
					DestinationGroups: pulumi.StringArray{
						appGroup.Path,
					},
					Services: pulumi.StringArray{
						appService.Path,
					},
				},
				&nsxt.PolicySecurityPolicyRuleTypeArgs{
					DisplayName: pulumi.String("Allow App to DB"),
					Description: pulumi.String("App to DB communication"),
					Action:      pulumi.String("ALLOW"),
					Logged:      pulumi.Bool(false),
					IpVersion:   pulumi.String("IPV4"),
					SourceGroups: pulumi.StringArray{
						appGroup.Path,
					},
					DestinationGroups: pulumi.StringArray{
						dbGroup.Path,
					},
					Services: pulumi.StringArray{
						pulumi.String(mysql.Path),
					},
				},
				&nsxt.PolicySecurityPolicyRuleTypeArgs{
					DisplayName: pulumi.String("Allow Infrastructure"),
					Description: pulumi.String("Allow DNS and Management servers"),
					Action:      pulumi.String("ALLOW"),
					Logged:      pulumi.Bool(true),
					IpVersion:   pulumi.String("IPV4"),
					SourceGroups: pulumi.StringArray{
						ipSet.Path,
					},
					DestinationGroups: pulumi.StringArray{
						allVms.Path,
					},
				},
				&nsxt.PolicySecurityPolicyRuleTypeArgs{
					DisplayName: pulumi.String("Allow out"),
					Description: pulumi.String("Outgoing rule"),
					Action:      pulumi.String("ALLOW"),
					Logged:      pulumi.Bool(true),
					IpVersion:   pulumi.String("IPV4"),
					SourceGroups: pulumi.StringArray{
						allVms.Path,
					},
				},
				&nsxt.PolicySecurityPolicyRuleTypeArgs{
					DisplayName: pulumi.String("Deny ANY"),
					Description: pulumi.String("Default Deny the traffic"),
					Action:      pulumi.String("REJECT"),
					Logged:      pulumi.Bool(true),
					IpVersion:   pulumi.String("IPV4"),
				},
			},
		})
		if err != nil {
			return err
		}
		// Here we have examples for creating NAT rules. The example here assumes
		// the Web IP addresses are reachable from outside and no NAT is required.
		_, err = nsxt.NewPolicyNatRule(ctx, "rule1", &nsxt.PolicyNatRuleArgs{
			DisplayName: pulumi.String("App 1-to-1 In"),
			Action:      pulumi.String("SNAT"),
			TranslatedNetworks: pulumi.StringArray{
				pulumi.String("102.10.22.1"),
			},
			SourceNetworks: pulumi.StringArray{
				pulumi.String("12.12.2.0/24"),
			},
			GatewayPath: t1Gateway.Path,
			Tags: nsxt.PolicyNatRuleTagArray{
				&nsxt.PolicyNatRuleTagArgs{
					Scope: pulumi.String(nsxTagScope),
					Tag:   pulumi.String(nsxTag),
				},
			},
		})
		if err != nil {
			return err
		}
		_, err = nsxt.NewPolicyNatRule(ctx, "rule2", &nsxt.PolicyNatRuleArgs{
			DisplayName: pulumi.String("App 1-to-1 Out"),
			Action:      pulumi.String("DNAT"),
			TranslatedNetworks: pulumi.StringArray{
				pulumi.String("102.10.22.2"),
			},
			DestinationNetworks: pulumi.StringArray{
				pulumi.String("102.10.22.1/32"),
			},
			SourceNetworks: pulumi.StringArray{
				pulumi.String("12.12.2.0/24"),
			},
			GatewayPath: t1Gateway.Path,
			Tags: nsxt.PolicyNatRuleTagArray{
				&nsxt.PolicyNatRuleTagArgs{
					Scope: pulumi.String(nsxTagScope),
					Tag:   pulumi.String(nsxTag),
				},
			},
		})
		if err != nil {
			return err
		}
		_, err = nsxt.NewPolicyNatRule(ctx, "rule3", &nsxt.PolicyNatRuleArgs{
			DisplayName: pulumi.String("DB 1-to-1 In"),
			Action:      pulumi.String("SNAT"),
			TranslatedNetworks: pulumi.StringArray{
				pulumi.String("102.10.23.1"),
			},
			SourceNetworks: pulumi.StringArray{
				pulumi.String("12.12.3.0/24"),
			},
			GatewayPath: t1Gateway.Path,
			Tags: nsxt.PolicyNatRuleTagArray{
				&nsxt.PolicyNatRuleTagArgs{
					Scope: pulumi.String(nsxTagScope),
					Tag:   pulumi.String(nsxTag),
				},
			},
		})
		if err != nil {
			return err
		}
		_, err = nsxt.NewPolicyNatRule(ctx, "rule4", &nsxt.PolicyNatRuleArgs{
			DisplayName: pulumi.String("App 1-to-1 Out"),
			Action:      pulumi.String("DNAT"),
			TranslatedNetworks: pulumi.StringArray{
				pulumi.String("102.10.23.3"),
			},
			DestinationNetworks: pulumi.StringArray{
				pulumi.String("102.10.23.1/32"),
			},
			SourceNetworks: pulumi.StringArray{
				pulumi.String("12.12.3.0/24"),
			},
			GatewayPath: t1Gateway.Path,
			Tags: nsxt.PolicyNatRuleTagArray{
				&nsxt.PolicyNatRuleTagArgs{
					Scope: pulumi.String(nsxTagScope),
					Tag:   pulumi.String(nsxTag),
				},
			},
		})
		if err != nil {
			return err
		}
		// The 3 VMs available in the NSX Inventory
		webVm, err := nsxt.GetPolicyVm(ctx, &nsxt.GetPolicyVmArgs{
			DisplayName: pulumi.StringRef("web-vm"),
		}, nil)
		if err != nil {
			return err
		}
		appVm, err := nsxt.GetPolicyVm(ctx, &nsxt.GetPolicyVmArgs{
			DisplayName: pulumi.StringRef("app-vm"),
		}, nil)
		if err != nil {
			return err
		}
		dbVm, err := nsxt.GetPolicyVm(ctx, &nsxt.GetPolicyVmArgs{
			DisplayName: pulumi.StringRef("db-vm"),
		}, nil)
		if err != nil {
			return err
		}
		// Assign the right tags to the VMs so that they get included in the
		// dynamic groups created above
		_, err = nsxt.NewPolicyVmTags(ctx, "web_vm_tag", &nsxt.PolicyVmTagsArgs{
			InstanceId: pulumi.String(webVm.InstanceId),
			Tags: nsxt.PolicyVmTagsTagArray{
				&nsxt.PolicyVmTagsTagArgs{
					Scope: pulumi.String("tier"),
					Tag:   pulumi.String("web"),
				},
				&nsxt.PolicyVmTagsTagArgs{
					Scope: pulumi.String(nsxTagScope),
					Tag:   pulumi.String(nsxTag),
				},
			},
		})
		if err != nil {
			return err
		}
		_, err = nsxt.NewPolicyVmTags(ctx, "app_vm_tag", &nsxt.PolicyVmTagsArgs{
			InstanceId: pulumi.String(appVm.InstanceId),
			Tags: nsxt.PolicyVmTagsTagArray{
				&nsxt.PolicyVmTagsTagArgs{
					Scope: pulumi.String("tier"),
					Tag:   pulumi.String("app"),
				},
				&nsxt.PolicyVmTagsTagArgs{
					Scope: pulumi.String(nsxTagScope),
					Tag:   pulumi.String(nsxTag),
				},
			},
		})
		if err != nil {
			return err
		}
		_, err = nsxt.NewPolicyVmTags(ctx, "db_vm_tag", &nsxt.PolicyVmTagsArgs{
			InstanceId: pulumi.String(dbVm.InstanceId),
			Tags: nsxt.PolicyVmTagsTagArray{
				&nsxt.PolicyVmTagsTagArgs{
					Scope: pulumi.String("tier"),
					Tag:   pulumi.String("db"),
				},
				&nsxt.PolicyVmTagsTagArgs{
					Scope: pulumi.String(nsxTagScope),
					Tag:   pulumi.String(nsxTag),
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
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
config:
    nsxt:allowUnverifiedSsl:
        value: true
    nsxt:host:
        value: 'TODO: var.nsx_manager'
    nsxt:maxRetries:
        value: 10
    nsxt:password:
        value: 'TODO: var.nsx_password'
    nsxt:retryMaxDelay:
        value: 5000
    nsxt:retryMinDelay:
        value: 500
    nsxt:retryOnStatusCodes:
        value:
            - 429
    nsxt:username:
        value: 'TODO: var.nsx_username'

```
```yaml
configuration:
  #
  # Here we show that you define a NSX tag which can be used later to easily to
  # search for the created objects in NSX.
  #
  nsxTagScope:
    type: string
    default: project
  nsxTag:
    type: string
    default: pulumi-demo
resources:
  #
  # Create a DHCP Profile that is used later
  # Note, this resource is only in NSX 3.0.0+
  tierDhcp:
    type: nsxt:PolicyDhcpServer
    name: tier_dhcp
    properties:
      displayName: tier_dhcp
      description: DHCP server servicing all 3 Segments
      serverAddresses:
        - 12.12.99.2/24
  #
  # In this part of the example, the settings required to create a Tier1 Gateway
  # are defined. In NSX a Tier1 Gateway is often used on a per user, tenant,
  # department or application basis. Each application may have it's own Tier1
  # Gateway. The Tier1 Gateway provides the default gateway for virtual machines
  # connected to the Segments on the Tier1 Gateway
  #
  t1Gateway:
    type: nsxt:PolicyTier1Gateway
    name: t1_gateway
    properties:
      displayName: TF_T1
      description: Tier1 provisioned by Pulumi
      edgeClusterPath: ${demo.path}
      dhcpConfigPath: ${tierDhcp.path}
      failoverMode: PREEMPTIVE
      defaultRuleLogging: 'false'
      enableFirewall: 'true'
      enableStandbyRelocation: 'false'
      forceWhitelisting: 'false'
      tier0Path: ${t0Gateway.path}
      routeAdvertisementTypes:
        - TIER1_STATIC_ROUTES
        - TIER1_CONNECTED
      poolAllocation: ROUTING
      tags:
        - scope: ${nsxTagScope}
          tag: ${nsxTag}
      routeAdvertisementRules:
        - name: rule1
          action: DENY
          subnets:
            - 20.0.0.0/24
            - 21.0.0.0/24
          prefixOperator: GE
          routeAdvertisementTypes:
            - TIER1_CONNECTED
  #
  # This shows the settings required to create NSX Segment (Logical Switch) to
  # which you can attach Virtual Machines (VMs)
  #
  web:
    type: nsxt:PolicySegment
    properties:
      displayName: web-tier
      description: Pulumi provisioned Web Segment
      connectivityPath: ${t1Gateway.path}
      transportZonePath: ${overlayTz.path}
      subnets:
        - cidr: 12.12.1.1/24
          dhcpRanges:
            - 12.12.1.100-12.12.1.160
          dhcpV4Config:
            serverAddress: 12.12.1.2/24
            leaseTime: 36000
            dhcpOption121s:
              - network: 6.6.6.0/24
                nextHop: 1.1.1.21
      tags:
        - scope: ${nsxTagScope}
          tag: ${nsxTag}
        - scope: tier
          tag: web
  app:
    type: nsxt:PolicySegment
    properties:
      displayName: app-tier
      description: Pulumi provisioned App Segment
      connectivityPath: ${t1Gateway.path}
      transportZonePath: ${overlayTz.path}
      subnets:
        - cidr: 12.12.2.1/24
          dhcpRanges:
            - 12.12.2.100-12.12.2.160
          dhcpV4Config:
            serverAddress: 12.12.2.2/24
            leaseTime: 36000
            dhcpOption121s:
              - network: 6.6.6.0/24
                nextHop: 1.1.1.21
      tags:
        - scope: ${nsxTagScope}
          tag: ${nsxTag}
        - scope: tier
          tag: app
  db:
    type: nsxt:PolicySegment
    properties:
      displayName: db-tier
      description: Pulumi provisioned DB Segment
      connectivityPath: ${t1Gateway.path}
      transportZonePath: ${overlayTz.path}
      subnets:
        - cidr: 12.12.3.1/24
          dhcpRanges:
            - 12.12.3.100-12.12.3.160
          dhcpV4Config:
            serverAddress: 12.12.3.2/24
            leaseTime: 36000
            dhcpOption121s:
              - network: 6.6.6.0/24
                nextHop: 1.1.1.21
      tags:
        - scope: ${nsxTagScope}
          tag: ${nsxTag}
        - scope: tier
          tag: db
  #
  # This part of the example shows creating Groups with dynamic membership
  # criteria
  #
  # All Virtual machines with specific tag and scope
  allVms:
    type: nsxt:PolicyGroup
    name: all_vms
    properties:
      displayName: All_VMs
      description: Group consisting of ALL VMs
      criterias:
        - conditions:
            - memberType: VirtualMachine
              operator: CONTAINS
              key: Tag
              value: ${nsxTag}
  # All WEB VMs
  webGroup:
    type: nsxt:PolicyGroup
    name: web_group
    properties:
      displayName: Web_VMs
      description: Group consisting of Web VMs
      criterias:
        - conditions:
            - memberType: VirtualMachine
              operator: CONTAINS
              key: Tag
              value: web
      tags:
        - scope: ${nsxTagScope}
          tag: ${nsxTag}
  # All App VMs
  appGroup:
    type: nsxt:PolicyGroup
    name: app_group
    properties:
      displayName: App_VMs
      description: Group consisting of App VMs
      criterias:
        - conditions:
            - memberType: VirtualMachine
              operator: CONTAINS
              key: Tag
              value: app
      tags:
        - scope: ${nsxTagScope}
          tag: ${nsxTag}
  # All DB VMs
  dbGroup:
    type: nsxt:PolicyGroup
    name: db_group
    properties:
      displayName: DB_VMs
      description: Group consisting of DB VMs
      criterias:
        - conditions:
            - memberType: VirtualMachine
              operator: CONTAINS
              key: Tag
              value: db
      tags:
        - scope: ${nsxTagScope}
          tag: ${nsxTag}
  # Static Group of IP addresses
  ipSet:
    type: nsxt:PolicyGroup
    name: ip_set
    properties:
      displayName: external_IPs
      description: Group containing all external IPs
      criterias:
        - ipaddressExpression:
            ipAddresses:
              - 211.1.1.1
              - 212.1.1.1
              - 192.168.1.1-192.168.1.100
      tags:
        - scope: ${nsxTagScope}
          tag: ${nsxTag}
  #
  # An example for Service for App that listens on port 8443
  #
  appService:
    type: nsxt:PolicyService
    name: app_service
    properties:
      displayName: app_service_8443
      description: Service for App that listens on port 8443
      l4PortSetEntries:
        - description: TCP Port 8443
          protocol: TCP
          destinationPorts:
            - '8443'
      tags:
        - scope: ${nsxTagScope}
          tag: ${nsxTag}
  #
  # In this section, we have example to create Firewall sections and rules
  # All rules in this section will be applied to VMs that are part of the
  # Gropus we created earlier
  #
  firewallSection:
    type: nsxt:PolicySecurityPolicy
    name: firewall_section
    properties:
      displayName: DFW Section
      description: Firewall section created by Pulumi
      scopes:
        - ${allVms.path}
      category: Application
      locked: 'false'
      stateful: 'true'
      tags:
        - scope: ${nsxTagScope}
          tag: ${nsxTag}
      rules:
        - displayName: Allow HTTPS
          description: In going rule
          action: ALLOW
          logged: 'false'
          ipVersion: IPV4
          destinationGroups:
            - ${webGroup.path}
          services:
            - ${https.path}
        - displayName: Allow SSH
          description: In going rule
          action: ALLOW
          logged: 'false'
          ipVersion: IPV4
          destinationGroups:
            - ${webGroup.path}
          services:
            - ${ssh.path}
        - displayName: Allow Web to App
          description: Web to App communication
          action: ALLOW
          logged: 'false'
          ipVersion: IPV4
          sourceGroups:
            - ${webGroup.path}
          destinationGroups:
            - ${appGroup.path}
          services:
            - ${appService.path}
        - displayName: Allow App to DB
          description: App to DB communication
          action: ALLOW
          logged: 'false'
          ipVersion: IPV4
          sourceGroups:
            - ${appGroup.path}
          destinationGroups:
            - ${dbGroup.path}
          services:
            - ${mysql.path}
        - displayName: Allow Infrastructure
          description: Allow DNS and Management servers
          action: ALLOW
          logged: 'true'
          ipVersion: IPV4
          sourceGroups:
            - ${ipSet.path}
          destinationGroups:
            - ${allVms.path}
        - displayName: Allow out
          description: Outgoing rule
          action: ALLOW
          logged: 'true'
          ipVersion: IPV4
          sourceGroups:
            - ${allVms.path}
        - displayName: Deny ANY
          description: Default Deny the traffic
          action: REJECT
          logged: 'true'
          ipVersion: IPV4
  #
  # Here we have examples for creating NAT rules. The example here assumes
  # the Web IP addresses are reachable from outside and no NAT is required.
  #
  rule1:
    type: nsxt:PolicyNatRule
    properties:
      displayName: App 1-to-1 In
      action: SNAT
      translatedNetworks: # NAT IP
        - 102.10.22.1
      sourceNetworks:
        - 12.12.2.0/24
      gatewayPath: ${t1Gateway.path}
      tags:
        - scope: ${nsxTagScope}
          tag: ${nsxTag}
  rule2:
    type: nsxt:PolicyNatRule
    properties:
      displayName: App 1-to-1 Out
      action: DNAT
      translatedNetworks:
        - 102.10.22.2
      destinationNetworks:
        - 102.10.22.1/32
      sourceNetworks:
        - 12.12.2.0/24
      gatewayPath: ${t1Gateway.path}
      tags:
        - scope: ${nsxTagScope}
          tag: ${nsxTag}
  rule3:
    type: nsxt:PolicyNatRule
    properties:
      displayName: DB 1-to-1 In
      action: SNAT
      translatedNetworks: # NAT IP
        - 102.10.23.1
      sourceNetworks:
        - 12.12.3.0/24
      gatewayPath: ${t1Gateway.path}
      tags:
        - scope: ${nsxTagScope}
          tag: ${nsxTag}
  rule4:
    type: nsxt:PolicyNatRule
    properties:
      displayName: App 1-to-1 Out
      action: DNAT
      translatedNetworks:
        - 102.10.23.3
      destinationNetworks:
        - 102.10.23.1/32
      sourceNetworks:
        - 12.12.3.0/24
      gatewayPath: ${t1Gateway.path}
      tags:
        - scope: ${nsxTagScope}
          tag: ${nsxTag}
  # Assign the right tags to the VMs so that they get included in the
  # dynamic groups created above
  webVmTag:
    type: nsxt:PolicyVmTags
    name: web_vm_tag
    properties:
      instanceId: ${webVm.instanceId}
      tags:
        - scope: tier
          tag: web
        - scope: ${nsxTagScope}
          tag: ${nsxTag}
  appVmTag:
    type: nsxt:PolicyVmTags
    name: app_vm_tag
    properties:
      instanceId: ${appVm.instanceId}
      tags:
        - scope: tier
          tag: app
        - scope: ${nsxTagScope}
          tag: ${nsxTag}
  dbVmTag:
    type: nsxt:PolicyVmTags
    name: db_vm_tag
    properties:
      instanceId: ${dbVm.instanceId}
      tags:
        - scope: tier
          tag: db
        - scope: ${nsxTagScope}
          tag: ${nsxTag}
variables:
  #
  # This part of the example shows some functions we will need to refer to
  # later in the .tf file. They include the transport zone, tier 0 router and
  # edge cluster.
  # There Tier-0 (T0) Gateway is considered a "provider" router that is pre-created
  # by the NSX Admin. A T0 Gateway is used for north/south connectivity between
  # the logical networking space and the physical networking space. Many Tier1
  # Gateways will be connected to the T0 Gateway
  #
  demo:
    fn::invoke:
      function: nsxt:getPolicyEdgeCluster
      arguments:
        displayName: Edge-Cluster
  overlayTz:
    fn::invoke:
      function: nsxt:getPolicyTransportZone
      arguments:
        displayName: Overlay-TZ
  t0Gateway:
    fn::invoke:
      function: nsxt:getPolicyTier0Gateway
      arguments:
        displayName: TF-T0-Gateway
  #
  # Here we have examples of create functions for Services
  #
  https:
    fn::invoke:
      function: nsxt:getPolicyService
      arguments:
        displayName: HTTPS
  mysql:
    fn::invoke:
      function: nsxt:getPolicyService
      arguments:
        displayName: MySQL
  ssh:
    fn::invoke:
      function: nsxt:getPolicyService
      arguments:
        displayName: SSH
  # The 3 VMs available in the NSX Inventory
  webVm:
    fn::invoke:
      function: nsxt:getPolicyVm
      arguments:
        displayName: web-vm
  appVm:
    fn::invoke:
      function: nsxt:getPolicyVm
      arguments:
        displayName: app-vm
  dbVm:
    fn::invoke:
      function: nsxt:getPolicyVm
      arguments:
        displayName: db-vm
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    nsxt:allowUnverifiedSsl:
        value: true
    nsxt:host:
        value: 'TODO: var.nsx_manager'
    nsxt:maxRetries:
        value: 10
    nsxt:password:
        value: 'TODO: var.nsx_password'
    nsxt:retryMaxDelay:
        value: 5000
    nsxt:retryMinDelay:
        value: 500
    nsxt:retryOnStatusCodes:
        value:
            - 429
    nsxt:username:
        value: 'TODO: var.nsx_username'

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.nsxt.NsxtFunctions;
import com.pulumi.nsxt.inputs.GetPolicyEdgeClusterArgs;
import com.pulumi.nsxt.inputs.GetPolicyTransportZoneArgs;
import com.pulumi.nsxt.inputs.GetPolicyTier0GatewayArgs;
import com.pulumi.nsxt.PolicyDhcpServer;
import com.pulumi.nsxt.PolicyDhcpServerArgs;
import com.pulumi.nsxt.PolicyTier1Gateway;
import com.pulumi.nsxt.PolicyTier1GatewayArgs;
import com.pulumi.nsxt.inputs.PolicyTier1GatewayTagArgs;
import com.pulumi.nsxt.inputs.PolicyTier1GatewayRouteAdvertisementRuleArgs;
import com.pulumi.nsxt.PolicySegment;
import com.pulumi.nsxt.PolicySegmentArgs;
import com.pulumi.nsxt.inputs.PolicySegmentSubnetArgs;
import com.pulumi.nsxt.inputs.PolicySegmentSubnetDhcpV4ConfigArgs;
import com.pulumi.nsxt.inputs.PolicySegmentTagArgs;
import com.pulumi.nsxt.PolicyGroup;
import com.pulumi.nsxt.PolicyGroupArgs;
import com.pulumi.nsxt.inputs.PolicyGroupCriteriaArgs;
import com.pulumi.nsxt.inputs.PolicyGroupTagArgs;
import com.pulumi.nsxt.inputs.PolicyGroupCriteriaIpaddressExpressionArgs;
import com.pulumi.nsxt.PolicyService;
import com.pulumi.nsxt.PolicyServiceArgs;
import com.pulumi.nsxt.inputs.PolicyServiceL4PortSetEntryArgs;
import com.pulumi.nsxt.inputs.PolicyServiceTagArgs;
import com.pulumi.nsxt.inputs.GetPolicyServiceArgs;
import com.pulumi.nsxt.PolicySecurityPolicy;
import com.pulumi.nsxt.PolicySecurityPolicyArgs;
import com.pulumi.nsxt.inputs.PolicySecurityPolicyTagArgs;
import com.pulumi.nsxt.inputs.PolicySecurityPolicyRuleArgs;
import com.pulumi.nsxt.PolicyNatRule;
import com.pulumi.nsxt.PolicyNatRuleArgs;
import com.pulumi.nsxt.inputs.PolicyNatRuleTagArgs;
import com.pulumi.nsxt.inputs.GetPolicyVmArgs;
import com.pulumi.nsxt.PolicyVmTags;
import com.pulumi.nsxt.PolicyVmTagsArgs;
import com.pulumi.nsxt.inputs.PolicyVmTagsTagArgs;
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
        final var nsxTagScope = config.get("nsxTagScope").orElse("project");
        final var nsxTag = config.get("nsxTag").orElse("pulumi-demo");
        //
        // This part of the example shows some functions we will need to refer to
        // later in the .tf file. They include the transport zone, tier 0 router and
        // edge cluster.
        // There Tier-0 (T0) Gateway is considered a "provider" router that is pre-created
        // by the NSX Admin. A T0 Gateway is used for north/south connectivity between
        // the logical networking space and the physical networking space. Many Tier1
        // Gateways will be connected to the T0 Gateway
        //
        final var demo = NsxtFunctions.getPolicyEdgeCluster(GetPolicyEdgeClusterArgs.builder()
            .displayName("Edge-Cluster")
            .build());

        final var overlayTz = NsxtFunctions.getPolicyTransportZone(GetPolicyTransportZoneArgs.builder()
            .displayName("Overlay-TZ")
            .build());

        final var t0Gateway = NsxtFunctions.getPolicyTier0Gateway(GetPolicyTier0GatewayArgs.builder()
            .displayName("TF-T0-Gateway")
            .build());

        //
        // Create a DHCP Profile that is used later
        // Note, this resource is only in NSX 3.0.0+
        var tierDhcp = new PolicyDhcpServer("tierDhcp", PolicyDhcpServerArgs.builder()
            .displayName("tier_dhcp")
            .description("DHCP server servicing all 3 Segments")
            .serverAddresses("12.12.99.2/24")
            .build());

        //
        // In this part of the example, the settings required to create a Tier1 Gateway
        // are defined. In NSX a Tier1 Gateway is often used on a per user, tenant,
        // department or application basis. Each application may have it's own Tier1
        // Gateway. The Tier1 Gateway provides the default gateway for virtual machines
        // connected to the Segments on the Tier1 Gateway
        //
        var t1Gateway = new PolicyTier1Gateway("t1Gateway", PolicyTier1GatewayArgs.builder()
            .displayName("TF_T1")
            .description("Tier1 provisioned by Pulumi")
            .edgeClusterPath(demo.path())
            .dhcpConfigPath(tierDhcp.path())
            .failoverMode("PREEMPTIVE")
            .defaultRuleLogging(false)
            .enableFirewall(true)
            .enableStandbyRelocation(false)
            .forceWhitelisting(false)
            .tier0Path(t0Gateway.path())
            .routeAdvertisementTypes(
                "TIER1_STATIC_ROUTES",
                "TIER1_CONNECTED")
            .poolAllocation("ROUTING")
            .tags(PolicyTier1GatewayTagArgs.builder()
                .scope(nsxTagScope)
                .tag(nsxTag)
                .build())
            .routeAdvertisementRules(PolicyTier1GatewayRouteAdvertisementRuleArgs.builder()
                .name("rule1")
                .action("DENY")
                .subnets(
                    "20.0.0.0/24",
                    "21.0.0.0/24")
                .prefixOperator("GE")
                .routeAdvertisementTypes("TIER1_CONNECTED")
                .build())
            .build());

        //
        // This shows the settings required to create NSX Segment (Logical Switch) to
        // which you can attach Virtual Machines (VMs)
        //
        var web = new PolicySegment("web", PolicySegmentArgs.builder()
            .displayName("web-tier")
            .description("Pulumi provisioned Web Segment")
            .connectivityPath(t1Gateway.path())
            .transportZonePath(overlayTz.path())
            .subnets(PolicySegmentSubnetArgs.builder()
                .cidr("12.12.1.1/24")
                .dhcpRanges("12.12.1.100-12.12.1.160")
                .dhcpV4Config(PolicySegmentSubnetDhcpV4ConfigArgs.builder()
                    .serverAddress("12.12.1.2/24")
                    .leaseTime(36000.0)
                    .dhcpOption121s(PolicySegmentSubnetDhcpV4ConfigDhcpOption121Args.builder()
                        .network("6.6.6.0/24")
                        .nextHop("1.1.1.21")
                        .build())
                    .build())
                .build())
            .tags(
                PolicySegmentTagArgs.builder()
                    .scope(nsxTagScope)
                    .tag(nsxTag)
                    .build(),
                PolicySegmentTagArgs.builder()
                    .scope("tier")
                    .tag("web")
                    .build())
            .build());

        var app = new PolicySegment("app", PolicySegmentArgs.builder()
            .displayName("app-tier")
            .description("Pulumi provisioned App Segment")
            .connectivityPath(t1Gateway.path())
            .transportZonePath(overlayTz.path())
            .subnets(PolicySegmentSubnetArgs.builder()
                .cidr("12.12.2.1/24")
                .dhcpRanges("12.12.2.100-12.12.2.160")
                .dhcpV4Config(PolicySegmentSubnetDhcpV4ConfigArgs.builder()
                    .serverAddress("12.12.2.2/24")
                    .leaseTime(36000.0)
                    .dhcpOption121s(PolicySegmentSubnetDhcpV4ConfigDhcpOption121Args.builder()
                        .network("6.6.6.0/24")
                        .nextHop("1.1.1.21")
                        .build())
                    .build())
                .build())
            .tags(
                PolicySegmentTagArgs.builder()
                    .scope(nsxTagScope)
                    .tag(nsxTag)
                    .build(),
                PolicySegmentTagArgs.builder()
                    .scope("tier")
                    .tag("app")
                    .build())
            .build());

        var db = new PolicySegment("db", PolicySegmentArgs.builder()
            .displayName("db-tier")
            .description("Pulumi provisioned DB Segment")
            .connectivityPath(t1Gateway.path())
            .transportZonePath(overlayTz.path())
            .subnets(PolicySegmentSubnetArgs.builder()
                .cidr("12.12.3.1/24")
                .dhcpRanges("12.12.3.100-12.12.3.160")
                .dhcpV4Config(PolicySegmentSubnetDhcpV4ConfigArgs.builder()
                    .serverAddress("12.12.3.2/24")
                    .leaseTime(36000.0)
                    .dhcpOption121s(PolicySegmentSubnetDhcpV4ConfigDhcpOption121Args.builder()
                        .network("6.6.6.0/24")
                        .nextHop("1.1.1.21")
                        .build())
                    .build())
                .build())
            .tags(
                PolicySegmentTagArgs.builder()
                    .scope(nsxTagScope)
                    .tag(nsxTag)
                    .build(),
                PolicySegmentTagArgs.builder()
                    .scope("tier")
                    .tag("db")
                    .build())
            .build());

        //
        // This part of the example shows creating Groups with dynamic membership
        // criteria
        //
        // All Virtual machines with specific tag and scope
        var allVms = new PolicyGroup("allVms", PolicyGroupArgs.builder()
            .displayName("All_VMs")
            .description("Group consisting of ALL VMs")
            .criterias(PolicyGroupCriteriaArgs.builder()
                .conditions(PolicyGroupCriteriaConditionArgs.builder()
                    .memberType("VirtualMachine")
                    .operator("CONTAINS")
                    .key("Tag")
                    .value(nsxTag)
                    .build())
                .build())
            .build());

        // All WEB VMs
        var webGroup = new PolicyGroup("webGroup", PolicyGroupArgs.builder()
            .displayName("Web_VMs")
            .description("Group consisting of Web VMs")
            .criterias(PolicyGroupCriteriaArgs.builder()
                .conditions(PolicyGroupCriteriaConditionArgs.builder()
                    .memberType("VirtualMachine")
                    .operator("CONTAINS")
                    .key("Tag")
                    .value("web")
                    .build())
                .build())
            .tags(PolicyGroupTagArgs.builder()
                .scope(nsxTagScope)
                .tag(nsxTag)
                .build())
            .build());

        // All App VMs
        var appGroup = new PolicyGroup("appGroup", PolicyGroupArgs.builder()
            .displayName("App_VMs")
            .description("Group consisting of App VMs")
            .criterias(PolicyGroupCriteriaArgs.builder()
                .conditions(PolicyGroupCriteriaConditionArgs.builder()
                    .memberType("VirtualMachine")
                    .operator("CONTAINS")
                    .key("Tag")
                    .value("app")
                    .build())
                .build())
            .tags(PolicyGroupTagArgs.builder()
                .scope(nsxTagScope)
                .tag(nsxTag)
                .build())
            .build());

        // All DB VMs
        var dbGroup = new PolicyGroup("dbGroup", PolicyGroupArgs.builder()
            .displayName("DB_VMs")
            .description("Group consisting of DB VMs")
            .criterias(PolicyGroupCriteriaArgs.builder()
                .conditions(PolicyGroupCriteriaConditionArgs.builder()
                    .memberType("VirtualMachine")
                    .operator("CONTAINS")
                    .key("Tag")
                    .value("db")
                    .build())
                .build())
            .tags(PolicyGroupTagArgs.builder()
                .scope(nsxTagScope)
                .tag(nsxTag)
                .build())
            .build());

        // Static Group of IP addresses
        var ipSet = new PolicyGroup("ipSet", PolicyGroupArgs.builder()
            .displayName("external_IPs")
            .description("Group containing all external IPs")
            .criterias(PolicyGroupCriteriaArgs.builder()
                .ipaddressExpression(PolicyGroupCriteriaIpaddressExpressionArgs.builder()
                    .ipAddresses(
                        "211.1.1.1",
                        "212.1.1.1",
                        "192.168.1.1-192.168.1.100")
                    .build())
                .build())
            .tags(PolicyGroupTagArgs.builder()
                .scope(nsxTagScope)
                .tag(nsxTag)
                .build())
            .build());

        //
        // An example for Service for App that listens on port 8443
        //
        var appService = new PolicyService("appService", PolicyServiceArgs.builder()
            .displayName("app_service_8443")
            .description("Service for App that listens on port 8443")
            .l4PortSetEntries(PolicyServiceL4PortSetEntryArgs.builder()
                .description("TCP Port 8443")
                .protocol("TCP")
                .destinationPorts("8443")
                .build())
            .tags(PolicyServiceTagArgs.builder()
                .scope(nsxTagScope)
                .tag(nsxTag)
                .build())
            .build());

        //
        // Here we have examples of create functions for Services
        //
        final var https = NsxtFunctions.getPolicyService(GetPolicyServiceArgs.builder()
            .displayName("HTTPS")
            .build());

        final var mysql = NsxtFunctions.getPolicyService(GetPolicyServiceArgs.builder()
            .displayName("MySQL")
            .build());

        final var ssh = NsxtFunctions.getPolicyService(GetPolicyServiceArgs.builder()
            .displayName("SSH")
            .build());

        //
        // In this section, we have example to create Firewall sections and rules
        // All rules in this section will be applied to VMs that are part of the
        // Gropus we created earlier
        //
        var firewallSection = new PolicySecurityPolicy("firewallSection", PolicySecurityPolicyArgs.builder()
            .displayName("DFW Section")
            .description("Firewall section created by Pulumi")
            .scopes(allVms.path())
            .category("Application")
            .locked(false)
            .stateful(true)
            .tags(PolicySecurityPolicyTagArgs.builder()
                .scope(nsxTagScope)
                .tag(nsxTag)
                .build())
            .rules(
                PolicySecurityPolicyRuleArgs.builder()
                    .displayName("Allow HTTPS")
                    .description("In going rule")
                    .action("ALLOW")
                    .logged(false)
                    .ipVersion("IPV4")
                    .destinationGroups(webGroup.path())
                    .services(https.path())
                    .build(),
                PolicySecurityPolicyRuleArgs.builder()
                    .displayName("Allow SSH")
                    .description("In going rule")
                    .action("ALLOW")
                    .logged(false)
                    .ipVersion("IPV4")
                    .destinationGroups(webGroup.path())
                    .services(ssh.path())
                    .build(),
                PolicySecurityPolicyRuleArgs.builder()
                    .displayName("Allow Web to App")
                    .description("Web to App communication")
                    .action("ALLOW")
                    .logged(false)
                    .ipVersion("IPV4")
                    .sourceGroups(webGroup.path())
                    .destinationGroups(appGroup.path())
                    .services(appService.path())
                    .build(),
                PolicySecurityPolicyRuleArgs.builder()
                    .displayName("Allow App to DB")
                    .description("App to DB communication")
                    .action("ALLOW")
                    .logged(false)
                    .ipVersion("IPV4")
                    .sourceGroups(appGroup.path())
                    .destinationGroups(dbGroup.path())
                    .services(mysql.path())
                    .build(),
                PolicySecurityPolicyRuleArgs.builder()
                    .displayName("Allow Infrastructure")
                    .description("Allow DNS and Management servers")
                    .action("ALLOW")
                    .logged(true)
                    .ipVersion("IPV4")
                    .sourceGroups(ipSet.path())
                    .destinationGroups(allVms.path())
                    .build(),
                PolicySecurityPolicyRuleArgs.builder()
                    .displayName("Allow out")
                    .description("Outgoing rule")
                    .action("ALLOW")
                    .logged(true)
                    .ipVersion("IPV4")
                    .sourceGroups(allVms.path())
                    .build(),
                PolicySecurityPolicyRuleArgs.builder()
                    .displayName("Deny ANY")
                    .description("Default Deny the traffic")
                    .action("REJECT")
                    .logged(true)
                    .ipVersion("IPV4")
                    .build())
            .build());

        //
        // Here we have examples for creating NAT rules. The example here assumes
        // the Web IP addresses are reachable from outside and no NAT is required.
        //
        var rule1 = new PolicyNatRule("rule1", PolicyNatRuleArgs.builder()
            .displayName("App 1-to-1 In")
            .action("SNAT")
            .translatedNetworks("102.10.22.1")
            .sourceNetworks("12.12.2.0/24")
            .gatewayPath(t1Gateway.path())
            .tags(PolicyNatRuleTagArgs.builder()
                .scope(nsxTagScope)
                .tag(nsxTag)
                .build())
            .build());

        var rule2 = new PolicyNatRule("rule2", PolicyNatRuleArgs.builder()
            .displayName("App 1-to-1 Out")
            .action("DNAT")
            .translatedNetworks("102.10.22.2")
            .destinationNetworks("102.10.22.1/32")
            .sourceNetworks("12.12.2.0/24")
            .gatewayPath(t1Gateway.path())
            .tags(PolicyNatRuleTagArgs.builder()
                .scope(nsxTagScope)
                .tag(nsxTag)
                .build())
            .build());

        var rule3 = new PolicyNatRule("rule3", PolicyNatRuleArgs.builder()
            .displayName("DB 1-to-1 In")
            .action("SNAT")
            .translatedNetworks("102.10.23.1")
            .sourceNetworks("12.12.3.0/24")
            .gatewayPath(t1Gateway.path())
            .tags(PolicyNatRuleTagArgs.builder()
                .scope(nsxTagScope)
                .tag(nsxTag)
                .build())
            .build());

        var rule4 = new PolicyNatRule("rule4", PolicyNatRuleArgs.builder()
            .displayName("App 1-to-1 Out")
            .action("DNAT")
            .translatedNetworks("102.10.23.3")
            .destinationNetworks("102.10.23.1/32")
            .sourceNetworks("12.12.3.0/24")
            .gatewayPath(t1Gateway.path())
            .tags(PolicyNatRuleTagArgs.builder()
                .scope(nsxTagScope)
                .tag(nsxTag)
                .build())
            .build());

        // The 3 VMs available in the NSX Inventory
        final var webVm = NsxtFunctions.getPolicyVm(GetPolicyVmArgs.builder()
            .displayName("web-vm")
            .build());

        final var appVm = NsxtFunctions.getPolicyVm(GetPolicyVmArgs.builder()
            .displayName("app-vm")
            .build());

        final var dbVm = NsxtFunctions.getPolicyVm(GetPolicyVmArgs.builder()
            .displayName("db-vm")
            .build());

        // Assign the right tags to the VMs so that they get included in the
        // dynamic groups created above
        var webVmTag = new PolicyVmTags("webVmTag", PolicyVmTagsArgs.builder()
            .instanceId(webVm.instanceId())
            .tags(
                PolicyVmTagsTagArgs.builder()
                    .scope("tier")
                    .tag("web")
                    .build(),
                PolicyVmTagsTagArgs.builder()
                    .scope(nsxTagScope)
                    .tag(nsxTag)
                    .build())
            .build());

        var appVmTag = new PolicyVmTags("appVmTag", PolicyVmTagsArgs.builder()
            .instanceId(appVm.instanceId())
            .tags(
                PolicyVmTagsTagArgs.builder()
                    .scope("tier")
                    .tag("app")
                    .build(),
                PolicyVmTagsTagArgs.builder()
                    .scope(nsxTagScope)
                    .tag(nsxTag)
                    .build())
            .build());

        var dbVmTag = new PolicyVmTags("dbVmTag", PolicyVmTagsArgs.builder()
            .instanceId(dbVm.instanceId())
            .tags(
                PolicyVmTagsTagArgs.builder()
                    .scope("tier")
                    .tag("db")
                    .build(),
                PolicyVmTagsTagArgs.builder()
                    .scope(nsxTagScope)
                    .tag(nsxTag)
                    .build())
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}

In order to be able to connect VMs to the newly created logical switch a new
`vpshereNetwork` function need to be defined.

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language python %}}
```python
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language go %}}
```go
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language java %}}
```java
Example currently unavailable in this language
```
{{% /choosable %}}
{{< /chooser >}}

The function in the above example should be referred in `networkId` inside
`networkInterface` section for `vsphereVirtualMachine` resource.

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as vsphere from "@pulumi/vsphere";

const appvm = new vsphere.index.VirtualMachine("appvm", {networkInterface: [{
    networkId: tfApp.id,
}]});
```
{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import pulumi_vsphere as vsphere

appvm = vsphere.index.VirtualMachine("appvm", network_interface=[{
    networkId: tf_app.id,
}])
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Vsphere = Pulumi.Vsphere;

return await Deployment.RunAsync(() =>
{
    var appvm = new Vsphere.Index.VirtualMachine("appvm", new()
    {
        NetworkInterface = new[]
        {

            {
                { "networkId", tfApp.Id },
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
	"github.com/pulumi/pulumi-vsphere/sdk/go/vsphere"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := vsphere.NewVirtualMachine(ctx, "appvm", &vsphere.VirtualMachineArgs{
			NetworkInterface: []map[string]interface{}{
				map[string]interface{}{
					"networkId": tfApp.Id,
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
  appvm:
    type: vsphere:VirtualMachine
    properties:
      networkInterface:
        - networkId: ${tfApp.id}
```
{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.vsphere.VirtualMachine;
import com.pulumi.vsphere.VirtualMachineArgs;
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
        var appvm = new VirtualMachine("appvm", VirtualMachineArgs.builder()
            .networkInterface(List.of(Map.of("networkId", tfApp.id())))
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}