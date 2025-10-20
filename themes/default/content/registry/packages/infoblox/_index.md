---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/infobloxopen/infoblox/2.11.0/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Infoblox Provider
meta_desc: Provides an overview on how to configure the Pulumi Infoblox provider.
layout: package
---

## Generate Provider

The Infoblox provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider infobloxopen/infoblox
```
## Prerequisites

Whether you intend to use the published plug-in or the customized version that you have built yourself, you must complete the following prerequisites:

- Install and set up a physical or virtual Infoblox NIOS appliance and has necessary licenses installed. Configure the access permissions for Pulumi to interact with NIOS Grid objects.
- To use the Infoblox IPAM Plug-In for Pulumi with a cloud user, you must either define the following extensible attributes or install the Cloud Network Automation license in the NIOS Grid, which adds the extensible attributes by default:
```json
{
    "Tenant ID": "String Type",
    "CMP Type": "String Type",
    "Cloud API Owned": "List Type (Values True, False)"
}
```
### **Creating the Pulumi Internal ID Extensible Attribute**
Create the Pulumi Internal ID Extensible Attribute in NIOS using one of the following methods. Only a NIOS admin with superuser privileges can create extensible attributes in NIOS.
- Create the extensible attribute manually in Infoblox NIOS Grid Manager. For steps, refer to the Adding Extensible Attributes topic in the Infoblox NIOS Documentation.

  If the user you want to manage is a cloud member, then enable the following option for the extensible attribute:
  - In Grid Manager, on the **Administration tab > Extensible Attributes** tab, edit the extensible attribute.
  - On the **Additional Properties** tab, enable **Allow cloud members to have the following access to this extensible attribute** and select **Read/Write (and disallow Write access from the GUI and the standard API)**.
- Use the following cURL command to create the extensible attribute as a read-only attribute in NIOS:

  ```bash
  curl -k -u <user>:<password> -H "Content-Type: application/json" -X POST https://<Grid_IP>/wapi/v2.12.3/extensibleattributedef -d '{"name": "Pulumi Internal ID", "flags": "CR", "type": "STRING", "comment": "Internal ID for Pulumi Resource"}'
  ```

  - If the user you want to manage is a cloud member, then include the flag C for cloud API.
  - If you are using multiple flags in the command, ensure that the flags are written in correct order. For more information about flags, refer to the Extensible Attribute Definition object in the Infoblox WAPI Documentation.

- Enable IPAM Plug-In for Pulumi to automatically create the extensible attribute by configuring the pulumi Infoblox provider with credentials of a NIOS admin user with superuser privileges. For more information, see Configure the Access Permissions.

> **Note:**
> Either the Pulumi Internal ID extensible attribute definition must be present in NIOS or IPAM Plug-In for Pulumi
> must be configured with superuser access for it to automatically create the extensible attribute. If not, the connection
> to Pulumi will fail.
> If you choose to create the Pulumi Internal ID extensible attribute manually or by using the cURL command,
> the creation of the extensible attribute is not managed by IPAM Plug-In for Pulumi.
> You must not modify the Pulumi Internal ID for a resource under any circumstances. If it is modified, the resource
> will no longer be managed by Pulumi.
## Configuring Infoblox Pulumi IPAM Plug-In

Pulumi relies on an Infoblox provider to interact with NIOS Grid objects. You can either use the published Infoblox provider (Infoblox IPAM Plug-In for Pulumi) available on the Pulumi Registry page or develop a plug-in with features that are not available in the published plug-in.

As a prerequisite, configure provider authentication to set up the required access permissions for Pulumi to interact with NIOS Grid objects. Additionally, declare the version of IPAM Plug-In for Pulumi in the .tf file to allow Pulumi to automatically install the published plug-in available in the Pulumi Registry.

To configure IPAM Plug-In for Pulumi for use, complete the following steps:

In the .tf file, specify the plug-in version in the requiredProviders block as follows in .tf file:

Configure the credentials required to access the NIOS Grid as environment variables or provider configuration in .tf file:

```bash
 # Using environment variable
 $ export INFOBLOX_SERVER=<nios_ip-addr or nios_hostname>
 $ export INFOBLOX_USERNAME=<nios_username>
 $ export INFOBLOX_PASSWORD=<nios_password>
```

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    infoblox:password:
        value: 'TODO: var.password'
    infoblox:server:
        value: 'TODO: var.server'
    infoblox:username:
        value: 'TODO: var.username'

```

Add other environment variables that you intend to use.
You can set the following environment variables instead of defining them as attributes inside the provider configuration in the .tf file. Each of these environment variables has a corresponding attribute in the provider configuration.
```
PORT
SSLMODE
CONNECT_TIMEOUT
POOL_CONNECTIONS
WAPI_VERSION
```
> **Note:** Plugin version **v2.9.0** includes an upgrade to the base WAPI version to **v2.12.3**.

Run the pulumi up command in the directory where the .tf file is located to initialize the plug-in.
## Resources

There are resources for the following objects, supported by the plugin:

* Network view (`infoblox.NetworkView`)
* Network container (`infoblox.Ipv4NetworkContainer`, `infoblox.Ipv6NetworkContainer`)
* Network (`infoblox.Ipv4Network`, `infoblox.Ipv6Network`)
* A-record (`infoblox.ARecord`)
* AAAA-record (`infoblox.AaaaRecord`)
* DNS View (`infoblox.DnsView`)
* PTR-record (`infoblox.PtrRecord`)
* CNAME-record (`infoblox.CnameRecord`)
* MX-record (`infoblox.MxRecord`)
* TXT-record (`infoblox.TxtRecord`)
* SRV-record (`infoblox.SrvRecord`)
* Zone Auth (`infoblox.ZoneAuth`)
* Zone Forward (`infoblox.ZoneForward`)
* Host record (`infoblox.IpAllocation` / `infoblox.IpAssociation`)
* Zone Delegated (`infoblox.ZoneDelegated`)
* DTC LBDN (`infoblox.DtcLbdn`)
* DTC Pool (`infoblox.DtcPool`)
* DTC Server (`infoblox.DtcServer`)
* Alias-record (`infoblox.AliasRecord`)
* NS-record (`infoblox.NsRecord`)
* IPV4 Shared Network (`infoblox.Ipv4SharedNetwork`)
* IPV4 Fixed Address (`infoblox.Ipv4FixedAddress`)
* IPV4 Range (`infoblox.Ipv4Range`)
* IPV4 Range Template (`infoblox.Ipv4RangeTemplate`)

Network and network container resources have two versions: IPv4 and IPv6. In
addition, there are two operations which are implemented as resources:
IP address allocation and IP address association with a network host
(ex. VM in a cloud environment); they have three versions: IPv4
and IPv6 separately, and IPv4/IPv6 combined.

The recommendation is to avoid using separate IPv4 and IPv6 versions of
IP allocation and IP association resources.
This recommendation does not relate to network container and network-related resources.

To work with DNS records a user must ensure that appropriate DNS zones
exist on the NIOS side, because currently the plugin does not support
creating a DNS zone.

Every resource has common attributes: 'comment' and 'ext_attrs'.
'comment' is text which describes the resource. 'ext_attrs' is a set of
NIOS Extensible Attributes attached to the resource.

For DNS-related resources there is 'ttl' attribute as well, it specifies
TTL value (in seconds) for appropriate record. There is no default
value, zone's TTL is used by NIOS, if the value is omitted.
In this case, the field 'ttl' will be set to a negative value in the Pulumi's state
for the resource, just to indicate the absence of a value.
TTL value of 0 (zero) means caching should be disabled for this record.

Please note that anywhere in the documents about this plugin, the default DNS view or
the default network view means: a view with the name `default`.
Usually, this is the name for the view which is marked as the default view on NIOS side, but this may be overridden.
But the plugin does use the name `default` for the view, despite which view is marked as the default on NIOS side.
## Data sources

There are functions for the following objects:

* Network View (`infoblox.NetworkView`)
* IPv4 Network (`infoblox.Ipv4Network`)
* IPv6 Network (`infoblox.Ipv6Network`)
* IPv4 Network Container (`infoblox.Ipv4NetworkContainer`)
* IPv6 Network Container (`infoblox.Ipv6NetworkContainer`)
* A-record (`infoblox.ARecord`)
* AAAA-record (`infoblox.AaaaRecord`)
* CNAME-record (`infoblox.CnameRecord`)
* DNS View (`infoblox.DnsView`)
* PTR-record (`infoblox.PtrRecord`)
* MX-record (`infoblox.MxRecord`)
* TXT-record (`infoblox.TxtRecord`)
* SRV-record (`infoblox.SrvRecord`)
* Zone Auth (`infoblox.ZoneAuth`)
* Zone Forward (`infoblox.ZoneForward`)
* Host Record (`infoblox.getHostRecord`)
* Zone Delegated (`infoblox.ZoneDelegated`)
* DTC LBDN (`infoblox.DtcLbdn`)
* DTC Pool (`infoblox.DtcPool`)
* DTC Server (`infoblox.DtcServer`)
* Alias-record (`infoblox.AliasRecord`)
* NS-record (`infoblox.NsRecord`)
* IPV4 Shared Network (`infoblox.Ipv4SharedNetwork`)
* IPV4 Fixed Address (`infoblox.Ipv4FixedAddress`)
* IPV4 Range (`infoblox.Ipv4Range`)
* IPV4 Range Template (`infoblox.Ipv4RangeTemplate`)

!> From version 2.5.0, new feature filters are introduced. Now the functions support to populate more than one
matching NIOS objects.

* `filters`: the schema, with passing combination of searchable fields are supported by NIOS server, which
  returns one or more matching objects from the NIOS server.

For usage of filters, add the fields as keys and appropriate values to be passed to the keys like `name`, `view` corresponding to object.
### Example for using filters:
{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as infoblox from "@pulumi/infoblox";

const vipHost = new infoblox.ARecord("vip_host", {
    fqdn: "very-interesting-host.example.com",
    ipAddr: "10.3.1.65",
    comment: "special host",
    dnsView: "nondefault_dnsview2",
    ttl: 120,
    extAttrs: JSON.stringify({
        Location: "65.8665701230204, -37.00791763398113",
    }),
});
const aRecTemp = infoblox.getARecord({
    filters: {
        name: "very-interesting-host.example.com",
        ipv4addr: "10.3.1.65",
        view: "nondefault_dnsview2",
    },
});
export const aRecRes = aRecTemp;
export const aRecName = aRecTemp.then(aRecTemp => aRecTemp.results?.[0]?.fqdn);
```
{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import json
import pulumi_infoblox as infoblox

vip_host = infoblox.ARecord("vip_host",
    fqdn="very-interesting-host.example.com",
    ip_addr="10.3.1.65",
    comment="special host",
    dns_view="nondefault_dnsview2",
    ttl=120,
    ext_attrs=json.dumps({
        "Location": "65.8665701230204, -37.00791763398113",
    }))
a_rec_temp = infoblox.get_a_record(filters={
    "name": "very-interesting-host.example.com",
    "ipv4addr": "10.3.1.65",
    "view": "nondefault_dnsview2",
})
pulumi.export("aRecRes", a_rec_temp)
pulumi.export("aRecName", a_rec_temp.results[0].fqdn)
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Infoblox = Pulumi.Infoblox;

return await Deployment.RunAsync(() =>
{
    var vipHost = new Infoblox.ARecord("vip_host", new()
    {
        Fqdn = "very-interesting-host.example.com",
        IpAddr = "10.3.1.65",
        Comment = "special host",
        DnsView = "nondefault_dnsview2",
        Ttl = 120,
        ExtAttrs = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["Location"] = "65.8665701230204, -37.00791763398113",
        }),
    });

    var aRecTemp = Infoblox.GetARecord.Invoke(new()
    {
        Filters =
        {
            { "name", "very-interesting-host.example.com" },
            { "ipv4addr", "10.3.1.65" },
            { "view", "nondefault_dnsview2" },
        },
    });

    return new Dictionary<string, object?>
    {
        ["aRecRes"] = aRecTemp,
        ["aRecName"] = aRecTemp.Apply(getARecordResult => getARecordResult.Results[0]?.Fqdn),
    };
});

```
{{% /choosable %}}
{{% choosable language go %}}
```go
package main

import (
	"encoding/json"

	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/infoblox/v2/infoblox"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"Location": "65.8665701230204, -37.00791763398113",
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		_, err = infoblox.NewARecord(ctx, "vip_host", &infoblox.ARecordArgs{
			Fqdn:     pulumi.String("very-interesting-host.example.com"),
			IpAddr:   pulumi.String("10.3.1.65"),
			Comment:  pulumi.String("special host"),
			DnsView:  pulumi.String("nondefault_dnsview2"),
			Ttl:      pulumi.Float64(120),
			ExtAttrs: pulumi.String(json0),
		})
		if err != nil {
			return err
		}
		aRecTemp, err := infoblox.LookupARecord(ctx, &infoblox.LookupARecordArgs{
			Filters: map[string]interface{}{
				"name":     "very-interesting-host.example.com",
				"ipv4addr": "10.3.1.65",
				"view":     "nondefault_dnsview2",
			},
		}, nil)
		if err != nil {
			return err
		}
		ctx.Export("aRecRes", aRecTemp)
		ctx.Export("aRecName", aRecTemp.Results[0].Fqdn)
		return nil
	})
}
```
{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
resources:
  vipHost:
    type: infoblox:ARecord
    name: vip_host
    properties:
      fqdn: very-interesting-host.example.com
      ipAddr: 10.3.1.65
      comment: special host
      dnsView: nondefault_dnsview2
      ttl: 120 # 120s
      extAttrs:
        fn::toJSON:
          Location: 65.8665701230204, -37.00791763398113
variables:
  aRecTemp:
    fn::invoke:
      function: infoblox:getARecord
      arguments:
        filters:
          name: very-interesting-host.example.com
          ipv4addr: 10.3.1.65
          view: nondefault_dnsview2
outputs:
  aRecRes: ${aRecTemp}
  # accessing individual field in results
  aRecName: ${aRecTemp.results[0].fqdn}
```
{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.infoblox.ARecord;
import com.pulumi.infoblox.ARecordArgs;
import com.pulumi.infoblox.InfobloxFunctions;
import com.pulumi.infoblox.inputs.GetARecordArgs;
import static com.pulumi.codegen.internal.Serialization.*;
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
        var vipHost = new ARecord("vipHost", ARecordArgs.builder()
            .fqdn("very-interesting-host.example.com")
            .ipAddr("10.3.1.65")
            .comment("special host")
            .dnsView("nondefault_dnsview2")
            .ttl(120)
            .extAttrs(serializeJson(
                jsonObject(
                    jsonProperty("Location", "65.8665701230204, -37.00791763398113")
                )))
            .build());

        final var aRecTemp = InfobloxFunctions.getARecord(GetARecordArgs.builder()
            .filters(Map.ofEntries(
                Map.entry("name", "very-interesting-host.example.com"),
                Map.entry("ipv4addr", "10.3.1.65"),
                Map.entry("view", "nondefault_dnsview2")
            ))
            .build());

        ctx.export("aRecRes", aRecTemp.applyValue(getARecordResult -> getARecordResult));
        ctx.export("aRecName", aRecTemp.applyValue(getARecordResult -> getARecordResult.results()[0].fqdn()));
    }
}
```
{{% /choosable %}}
{{< /chooser >}}

The list of matching objects as JSON format returned in output under results, with fields or arguments that are passed in the filters.

Filters will support `EA Search` i.e, fetches matching objects or records associated with the EAs' corresponding to provided function, if any.
### Example for using filters for EA Search:
{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as infoblox from "@pulumi/infoblox";

const aRec1 = infoblox.getARecord({
    filters: {
        "*TestEA": "Acceptance Test",
    },
});
```
{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import pulumi_infoblox as infoblox

a_rec1 = infoblox.get_a_record(filters={
    "*TestEA": "Acceptance Test",
})
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Infoblox = Pulumi.Infoblox;

return await Deployment.RunAsync(() =>
{
    var aRec1 = Infoblox.GetARecord.Invoke(new()
    {
        Filters =
        {
            { "*TestEA", "Acceptance Test" },
        },
    });

});

```
{{% /choosable %}}
{{% choosable language go %}}
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/infoblox/v2/infoblox"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := infoblox.LookupARecord(ctx, &infoblox.LookupARecordArgs{
			Filters: map[string]interface{}{
				"*TestEA": "Acceptance Test",
			},
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
  aRec1:
    fn::invoke:
      function: infoblox:getARecord
      arguments:
        filters:
          '*TestEA': Acceptance Test
```
{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.infoblox.InfobloxFunctions;
import com.pulumi.infoblox.inputs.GetARecordArgs;
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
        final var aRec1 = InfobloxFunctions.getARecord(GetARecordArgs.builder()
            .filters(Map.of("*TestEA", "Acceptance Test"))
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}

Filters will also support Multi Value EA Search, where if the EA has more than one value, to be passed as comma seperated
value as a string. In here EAs' can have multiple or multi values of types like 'string', 'integer', etc.
### Example for using Multi Value EA Search:
{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as infoblox from "@pulumi/infoblox";

const aRec2 = infoblox.getARecord({
    filters: {
        "*tf_multi_val": "test,test2,demo",
    },
});
// for negative condition, if there are common EA values associated with different objects, to fetch unique record or object
const aRec3 = infoblox.getARecord({
    filters: {
        "*tf_multi_val": "test",
        "*tf_multi_val!": "dummy",
    },
});
```
{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import pulumi_infoblox as infoblox

a_rec2 = infoblox.get_a_record(filters={
    "*tf_multi_val": "test,test2,demo",
})
# for negative condition, if there are common EA values associated with different objects, to fetch unique record or object
a_rec3 = infoblox.get_a_record(filters={
    "*tf_multi_val": "test",
    "*tf_multi_val!": "dummy",
})
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Infoblox = Pulumi.Infoblox;

return await Deployment.RunAsync(() =>
{
    var aRec2 = Infoblox.GetARecord.Invoke(new()
    {
        Filters =
        {
            { "*tf_multi_val", "test,test2,demo" },
        },
    });

    // for negative condition, if there are common EA values associated with different objects, to fetch unique record or object
    var aRec3 = Infoblox.GetARecord.Invoke(new()
    {
        Filters =
        {
            { "*tf_multi_val", "test" },
            { "*tf_multi_val!", "dummy" },
        },
    });

});

```
{{% /choosable %}}
{{% choosable language go %}}
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/infoblox/v2/infoblox"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := infoblox.LookupARecord(ctx, &infoblox.LookupARecordArgs{
			Filters: map[string]interface{}{
				"*tf_multi_val": "test,test2,demo",
			},
		}, nil)
		if err != nil {
			return err
		}
		// for negative condition, if there are common EA values associated with different objects, to fetch unique record or object
		_, err = infoblox.LookupARecord(ctx, &infoblox.LookupARecordArgs{
			Filters: map[string]interface{}{
				"*tf_multi_val":  "test",
				"*tf_multi_val!": "dummy",
			},
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
  aRec2:
    fn::invoke:
      function: infoblox:getARecord
      arguments:
        filters:
          '*tf_multi_val': test,test2,demo
  # for negative condition, if there are common EA values associated with different objects, to fetch unique record or object
  aRec3:
    fn::invoke:
      function: infoblox:getARecord
      arguments:
        filters:
          '*tf_multi_val': test
          '*tf_multi_val!': dummy
```
{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.infoblox.InfobloxFunctions;
import com.pulumi.infoblox.inputs.GetARecordArgs;
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
        final var aRec2 = InfobloxFunctions.getARecord(GetARecordArgs.builder()
            .filters(Map.of("*tf_multi_val", "test,test2,demo"))
            .build());

        // for negative condition, if there are common EA values associated with different objects, to fetch unique record or object
        final var aRec3 = InfobloxFunctions.getARecord(GetARecordArgs.builder()
            .filters(Map.ofEntries(
                Map.entry("*tf_multi_val", "test"),
                Map.entry("*tf_multi_val!", "dummy")
            ))
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Importing existing resources

There is a possibility to import existing resources, enabling them to be managed by Pulumi.
As of now, there is a limitation: you have to write full resource's definition yourself.

In general, the process of importing an existing resource looks like this:

- write a new Pulumi file (ex. a-record-imported.tf) with the content:
  {{< chooser language "typescript,python,go,csharp,java,yaml" >}}
  {{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as infoblox from "@pulumi/infoblox";

const aRec1Imported = new infoblox.ARecord("a_rec_1_imported", {
    fqdn: "rec-a-1.imported.test.com",
    ipAddr: "192.168.1.2",
    ttl: 10,
    comment: "A-record to be imported",
    extAttrs: JSON.stringify({
        Location: "New office",
    }),
});
```
{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import json
import pulumi_infoblox as infoblox

a_rec1_imported = infoblox.ARecord("a_rec_1_imported",
    fqdn="rec-a-1.imported.test.com",
    ip_addr="192.168.1.2",
    ttl=10,
    comment="A-record to be imported",
    ext_attrs=json.dumps({
        "Location": "New office",
    }))
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Infoblox = Pulumi.Infoblox;

return await Deployment.RunAsync(() =>
{
    var aRec1Imported = new Infoblox.ARecord("a_rec_1_imported", new()
    {
        Fqdn = "rec-a-1.imported.test.com",
        IpAddr = "192.168.1.2",
        Ttl = 10,
        Comment = "A-record to be imported",
        ExtAttrs = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["Location"] = "New office",
        }),
    });

});

```
{{% /choosable %}}
{{% choosable language go %}}
```go
package main

import (
	"encoding/json"

	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/infoblox/v2/infoblox"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"Location": "New office",
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		_, err = infoblox.NewARecord(ctx, "a_rec_1_imported", &infoblox.ARecordArgs{
			Fqdn:     pulumi.String("rec-a-1.imported.test.com"),
			IpAddr:   pulumi.String("192.168.1.2"),
			Ttl:      pulumi.Float64(10),
			Comment:  pulumi.String("A-record to be imported"),
			ExtAttrs: pulumi.String(json0),
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
  aRec1Imported:
    type: infoblox:ARecord
    name: a_rec_1_imported
    properties:
      fqdn: rec-a-1.imported.test.com
      ipAddr: 192.168.1.2
      ttl: 10
      comment: A-record to be imported
      extAttrs:
        fn::toJSON:
          Location: New office
```
{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.infoblox.ARecord;
import com.pulumi.infoblox.ARecordArgs;
import static com.pulumi.codegen.internal.Serialization.*;
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
        var aRec1Imported = new ARecord("aRec1Imported", ARecordArgs.builder()
            .fqdn("rec-a-1.imported.test.com")
            .ipAddr("192.168.1.2")
            .ttl(10)
            .comment("A-record to be imported")
            .extAttrs(serializeJson(
                jsonObject(
                    jsonProperty("Location", "New office")
                )))
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
``
- get a reference for the resource you want to import (ex. by using `curl` tool)
- issue a command of the form `pulumi import RESOURCE_TYPE.RESOURCE_NAME RESOURCE_REFERENCE`.
  Example: `pulumi import infoblox_a_record.a_rec_1_imported record:a/ZG5zLmJpbmRfYSQuX2RlZmF1bHQub3JnLmV4YW1wbGUsc3RhdGljMSwxLjIuMy40:rec-a-1.imported.test.com/default`

Please, note that if some of resource's properties (supported by the Infoblox provider plugin) is not defined or
is empty for the object on NIOS side, then appropriate resource's property must be empty or not defined.
Otherwise, you will get a difference in the resource's actual state and resource's description you specified,
and thus you will get a resource's update performed on the next `pulumi up` command invocation,
which will actually set the value of the property to the one which you defined (ex. empty value).

To import a host record (represented by the `infoblox.IpAllocation` and
`infoblox.IpAssociation` resources in Pulumi), add the `Pulumi Internal ID` extensible attribute
with a randomly generated value in the form of a UUID to the record.
- For steps to add the extensible attribute, refer to the [Infoblox NIOS Documentation](https://docs.infoblox.com).
- You may use the command-line tool `uuid` for Linux-based systems to generate a UUID.

> The `Pulumi Internal ID` extensible attribute is not shown in to pulumi.tfstate file. Use it to create
> or import the `infoblox.IpAllocation` and `infoblox.IpAssociation` resources.
> You must not add it in a resource block with other extensible attributes.

> You must not delete (ex. with 'pulumi destroy' command) an `infoblox.IpAssociation` resource right after importing, but you may do this after 'pulumi up'.
> The reason: after 'pulumi import' the dependency between `infoblox.IpAssociation` and respective `infoblox.IpAllocation` is not established by Pulumi.
### Utilizing the Import Block to Import Resources:

As a prerequisite, for the object that you need to import, obtain the reference ID assigned to the object in NIOS.

Define the import block in the Pulumi Configuration (.tf) file of a resource that must be imported. In the .tf file of the resource to import, include the following block:
#### Example for importing A-records from a zone

> **Note:**
> When using the Pulumi import block for a resource, a new Pulumi internal ID is assigned to the resource when the pulumi preview command is run for the first time. If a subsequent pulumi up is aborted, the record will still retain the Pulumi Internal ID though the resource is not managed by Pulumi.