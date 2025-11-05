---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/megaport/megaport/1.4.6/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Megaport Provider
meta_desc: Provides an overview on how to configure the Pulumi Megaport provider.
layout: package
---

## Generate Provider

The Megaport provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider megaport/megaport
```
## Overview

The `pulumi-provider-megaport` or Megaport Pulumi Provider lets you create and manage
Megaport's product and services using the [Megaport API](https://dev.megaport.com).

This provides an opportunity for true multi-cloud hybrid environments supported by Megaport's Software
Defined Network (SDN). Using the Pulumi provider, you can create and manage Ports,
Virtual Cross Connects (VXCs), Megaport Cloud Routers (MCRs), MCR Prefix Filter Lists, Megaport Virtual Edges (MVEs), and Partner VXCs.

This provider is compatible with HashiCorp Pulumi, and we have tested compatibility with OpenTofu and haven't seen issues.

The Megaport Pulumi Provider is released as a tool for use with the Megaport API.

**Important:** The usage of the Megaport Pulumi Provider constitutes your acceptance of the terms available
in the Megaport [Acceptable Use Policy](https://www.megaport.com/legal/acceptable-use-policy/) and
[Global Services Agreement](https://www.megaport.com/legal/global-services-agreement/).
## Documentation

Documentation is published on the Pulumi Provider Megaport registry and the [OpenTofu Provider Megaport](https://search.opentofu.org/provider/megaport/megaport/latest) registry.
## Installation
### Pulumi Installation

The preferred installation method is via the Pulumi Provider Megaport
registry.
### OpenTofu Installation

For OpenTofu users, the provider is available via the [OpenTofu Registry](https://search.opentofu.org/provider/megaport/megaport/latest). No configuration changes are needed - use the same provider source as you would with Pulumi.
## Configuration

The provider can be configured in the same way whether using HashiCorp Pulumi or OpenTofu:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    megaport:acceptPurchaseTerms:
        value: true
    megaport:accessKey:
        value: your-access-key
    megaport:environment:
        value: production
    megaport:secretKey:
        value: your-secret-key

```
```typescript
import * as pulumi from "@pulumi/pulumi";

```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    megaport:acceptPurchaseTerms:
        value: true
    megaport:accessKey:
        value: your-access-key
    megaport:environment:
        value: production
    megaport:secretKey:
        value: your-secret-key

```
```python
import pulumi

```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    megaport:acceptPurchaseTerms:
        value: true
    megaport:accessKey:
        value: your-access-key
    megaport:environment:
        value: production
    megaport:secretKey:
        value: your-secret-key

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;

return await Deployment.RunAsync(() =>
{
});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    megaport:acceptPurchaseTerms:
        value: true
    megaport:accessKey:
        value: your-access-key
    megaport:environment:
        value: production
    megaport:secretKey:
        value: your-secret-key

```
```go
package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
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
    megaport:acceptPurchaseTerms:
        value: true
    megaport:accessKey:
        value: your-access-key
    megaport:environment:
        value: production
    megaport:secretKey:
        value: your-secret-key

```
```yaml
{}
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    megaport:acceptPurchaseTerms:
        value: true
    megaport:accessKey:
        value: your-access-key
    megaport:environment:
        value: production
    megaport:secretKey:
        value: your-secret-key

```
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
    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## 🚨 NEW FEATURE: MCR Prefix Filter List Resources
### Enhanced MCR Management with Standalone Resources

The Megaport Pulumi Provider now supports managing MCR prefix filter lists as individual resources, providing better lifecycle management and improved state handling compared to the previous inline approach.
#### ✅ New Standalone Approach (Recommended)

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as megaport from "@pulumi/megaport";

// Create MCR without inline prefix filter lists
const example = new megaport.Mcr("example", {
    productName: "my-mcr",
    portSpeed: 1000,
    locationId: 5,
    contractTermMonths: 12,
});
// Manage prefix filter lists as individual resources
const allowPrivateIpv4 = new megaport.McrPrefixFilterList("allow_private_ipv4", {
    mcrId: example.productUid,
    description: "Allow private IPv4 networks",
    addressFamily: "IPv4",
    entries: [
        {
            action: "permit",
            prefix: "10.0.0.0/8",
            ge: 16,
            le: 24,
        },
        {
            action: "permit",
            prefix: "192.168.0.0/16",
            ge: 24,
            le: 32,
        },
    ],
});
const allowPrivateIpv6 = new megaport.McrPrefixFilterList("allow_private_ipv6", {
    mcrId: example.productUid,
    description: "Allow private IPv6 networks",
    addressFamily: "IPv6",
    entries: [{
        action: "permit",
        prefix: "fd00::/8",
        ge: 48,
        le: 64,
    }],
});
```
{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import pulumi_megaport as megaport

# Create MCR without inline prefix filter lists
example = megaport.Mcr("example",
    product_name="my-mcr",
    port_speed=1000,
    location_id=5,
    contract_term_months=12)
# Manage prefix filter lists as individual resources
allow_private_ipv4 = megaport.McrPrefixFilterList("allow_private_ipv4",
    mcr_id=example.product_uid,
    description="Allow private IPv4 networks",
    address_family="IPv4",
    entries=[
        {
            "action": "permit",
            "prefix": "10.0.0.0/8",
            "ge": 16,
            "le": 24,
        },
        {
            "action": "permit",
            "prefix": "192.168.0.0/16",
            "ge": 24,
            "le": 32,
        },
    ])
allow_private_ipv6 = megaport.McrPrefixFilterList("allow_private_ipv6",
    mcr_id=example.product_uid,
    description="Allow private IPv6 networks",
    address_family="IPv6",
    entries=[{
        "action": "permit",
        "prefix": "fd00::/8",
        "ge": 48,
        "le": 64,
    }])
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Megaport = Pulumi.Megaport;

return await Deployment.RunAsync(() =>
{
    // Create MCR without inline prefix filter lists
    var example = new Megaport.Mcr("example", new()
    {
        ProductName = "my-mcr",
        PortSpeed = 1000,
        LocationId = 5,
        ContractTermMonths = 12,
    });

    // Manage prefix filter lists as individual resources
    var allowPrivateIpv4 = new Megaport.McrPrefixFilterList("allow_private_ipv4", new()
    {
        McrId = example.ProductUid,
        Description = "Allow private IPv4 networks",
        AddressFamily = "IPv4",
        Entries = new[]
        {
            new Megaport.Inputs.McrPrefixFilterListEntryArgs
            {
                Action = "permit",
                Prefix = "10.0.0.0/8",
                Ge = 16,
                Le = 24,
            },
            new Megaport.Inputs.McrPrefixFilterListEntryArgs
            {
                Action = "permit",
                Prefix = "192.168.0.0/16",
                Ge = 24,
                Le = 32,
            },
        },
    });

    var allowPrivateIpv6 = new Megaport.McrPrefixFilterList("allow_private_ipv6", new()
    {
        McrId = example.ProductUid,
        Description = "Allow private IPv6 networks",
        AddressFamily = "IPv6",
        Entries = new[]
        {
            new Megaport.Inputs.McrPrefixFilterListEntryArgs
            {
                Action = "permit",
                Prefix = "fd00::/8",
                Ge = 48,
                Le = 64,
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
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/megaport/megaport"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create MCR without inline prefix filter lists
		example, err := megaport.NewMcr(ctx, "example", &megaport.McrArgs{
			ProductName:        pulumi.String("my-mcr"),
			PortSpeed:          pulumi.Float64(1000),
			LocationId:         pulumi.Float64(5),
			ContractTermMonths: pulumi.Float64(12),
		})
		if err != nil {
			return err
		}
		// Manage prefix filter lists as individual resources
		_, err = megaport.NewMcrPrefixFilterList(ctx, "allow_private_ipv4", &megaport.McrPrefixFilterListArgs{
			McrId:         example.ProductUid,
			Description:   pulumi.String("Allow private IPv4 networks"),
			AddressFamily: pulumi.String("IPv4"),
			Entries: megaport.McrPrefixFilterListEntryArray{
				&megaport.McrPrefixFilterListEntryArgs{
					Action: pulumi.String("permit"),
					Prefix: pulumi.String("10.0.0.0/8"),
					Ge:     pulumi.Float64(16),
					Le:     pulumi.Float64(24),
				},
				&megaport.McrPrefixFilterListEntryArgs{
					Action: pulumi.String("permit"),
					Prefix: pulumi.String("192.168.0.0/16"),
					Ge:     pulumi.Float64(24),
					Le:     pulumi.Float64(32),
				},
			},
		})
		if err != nil {
			return err
		}
		_, err = megaport.NewMcrPrefixFilterList(ctx, "allow_private_ipv6", &megaport.McrPrefixFilterListArgs{
			McrId:         example.ProductUid,
			Description:   pulumi.String("Allow private IPv6 networks"),
			AddressFamily: pulumi.String("IPv6"),
			Entries: megaport.McrPrefixFilterListEntryArray{
				&megaport.McrPrefixFilterListEntryArgs{
					Action: pulumi.String("permit"),
					Prefix: pulumi.String("fd00::/8"),
					Ge:     pulumi.Float64(48),
					Le:     pulumi.Float64(64),
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
  # Create MCR without inline prefix filter lists
  example:
    type: megaport:Mcr
    properties:
      productName: my-mcr
      portSpeed: 1000
      locationId: 5
      contractTermMonths: 12
  # Manage prefix filter lists as individual resources
  allowPrivateIpv4:
    type: megaport:McrPrefixFilterList
    name: allow_private_ipv4
    properties:
      mcrId: ${example.productUid}
      description: Allow private IPv4 networks
      addressFamily: IPv4
      entries:
        - action: permit
          prefix: 10.0.0.0/8
          ge: 16
          le: 24
        - action: permit
          prefix: 192.168.0.0/16
          ge: 24
          le: 32
  allowPrivateIpv6:
    type: megaport:McrPrefixFilterList
    name: allow_private_ipv6
    properties:
      mcrId: ${example.productUid}
      description: Allow private IPv6 networks
      addressFamily: IPv6
      entries:
        - action: permit
          prefix: fd00::/8
          ge: 48
          le: 64
```
{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.megaport.Mcr;
import com.pulumi.megaport.McrArgs;
import com.pulumi.megaport.McrPrefixFilterList;
import com.pulumi.megaport.McrPrefixFilterListArgs;
import com.pulumi.megaport.inputs.McrPrefixFilterListEntryArgs;
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
        // Create MCR without inline prefix filter lists
        var example = new Mcr("example", McrArgs.builder()
            .productName("my-mcr")
            .portSpeed(1000.0)
            .locationId(5.0)
            .contractTermMonths(12.0)
            .build());

        // Manage prefix filter lists as individual resources
        var allowPrivateIpv4 = new McrPrefixFilterList("allowPrivateIpv4", McrPrefixFilterListArgs.builder()
            .mcrId(example.productUid())
            .description("Allow private IPv4 networks")
            .addressFamily("IPv4")
            .entries(
                McrPrefixFilterListEntryArgs.builder()
                    .action("permit")
                    .prefix("10.0.0.0/8")
                    .ge(16.0)
                    .le(24.0)
                    .build(),
                McrPrefixFilterListEntryArgs.builder()
                    .action("permit")
                    .prefix("192.168.0.0/16")
                    .ge(24.0)
                    .le(32.0)
                    .build())
            .build());

        var allowPrivateIpv6 = new McrPrefixFilterList("allowPrivateIpv6", McrPrefixFilterListArgs.builder()
            .mcrId(example.productUid())
            .description("Allow private IPv6 networks")
            .addressFamily("IPv6")
            .entries(McrPrefixFilterListEntryArgs.builder()
                .action("permit")
                .prefix("fd00::/8")
                .ge(48.0)
                .le(64.0)
                .build())
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
#### ⚠️ Deprecated Inline Approach

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as megaport from "@pulumi/megaport";

// ❌ DEPRECATED: Inline prefix filter lists (will be removed in future version)
const deprecatedExample = new megaport.Mcr("deprecated_example", {
    productName: "my-mcr",
    portSpeed: 1000,
    locationId: 5,
    contractTermMonths: 12,
    prefixFilterLists: [{
        description: "Allow private networks",
        addressFamily: "IPv4",
        entries: [{
            action: "permit",
            prefix: "10.0.0.0/8",
            ge: 16,
            le: 24,
        }],
    }],
});
```
{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import pulumi_megaport as megaport

# ❌ DEPRECATED: Inline prefix filter lists (will be removed in future version)
deprecated_example = megaport.Mcr("deprecated_example",
    product_name="my-mcr",
    port_speed=1000,
    location_id=5,
    contract_term_months=12,
    prefix_filter_lists=[{
        "description": "Allow private networks",
        "address_family": "IPv4",
        "entries": [{
            "action": "permit",
            "prefix": "10.0.0.0/8",
            "ge": 16,
            "le": 24,
        }],
    }])
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Megaport = Pulumi.Megaport;

return await Deployment.RunAsync(() =>
{
    // ❌ DEPRECATED: Inline prefix filter lists (will be removed in future version)
    var deprecatedExample = new Megaport.Mcr("deprecated_example", new()
    {
        ProductName = "my-mcr",
        PortSpeed = 1000,
        LocationId = 5,
        ContractTermMonths = 12,
        PrefixFilterLists = new[]
        {
            new Megaport.Inputs.McrPrefixFilterListArgs
            {
                Description = "Allow private networks",
                AddressFamily = "IPv4",
                Entries = new[]
                {
                    new Megaport.Inputs.McrPrefixFilterListEntryArgs
                    {
                        Action = "permit",
                        Prefix = "10.0.0.0/8",
                        Ge = 16,
                        Le = 24,
                    },
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
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/megaport/megaport"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// ❌ DEPRECATED: Inline prefix filter lists (will be removed in future version)
		_, err := megaport.NewMcr(ctx, "deprecated_example", &megaport.McrArgs{
			ProductName:        pulumi.String("my-mcr"),
			PortSpeed:          pulumi.Float64(1000),
			LocationId:         pulumi.Float64(5),
			ContractTermMonths: pulumi.Float64(12),
			PrefixFilterLists: megaport.McrPrefixFilterListTypeArray{
				&megaport.McrPrefixFilterListTypeArgs{
					Description:   pulumi.String("Allow private networks"),
					AddressFamily: pulumi.String("IPv4"),
					Entries: megaport.McrPrefixFilterListEntryArray{
						&megaport.McrPrefixFilterListEntryArgs{
							Action: pulumi.String("permit"),
							Prefix: pulumi.String("10.0.0.0/8"),
							Ge:     pulumi.Float64(16),
							Le:     pulumi.Float64(24),
						},
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
  # ❌ DEPRECATED: Inline prefix filter lists (will be removed in future version)
  deprecatedExample:
    type: megaport:Mcr
    name: deprecated_example
    properties:
      productName: my-mcr
      portSpeed: 1000
      locationId: 5
      contractTermMonths: 12 # This approach is deprecated and will show warnings
      prefixFilterLists:
        - description: Allow private networks
          addressFamily: IPv4
          entries:
            - action: permit
              prefix: 10.0.0.0/8
              ge: 16
              le: 24
```
{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.megaport.Mcr;
import com.pulumi.megaport.McrArgs;
import com.pulumi.megaport.inputs.McrPrefixFilterListArgs;
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
        // ❌ DEPRECATED: Inline prefix filter lists (will be removed in future version)
        var deprecatedExample = new Mcr("deprecatedExample", McrArgs.builder()
            .productName("my-mcr")
            .portSpeed(1000.0)
            .locationId(5.0)
            .contractTermMonths(12.0)
            .prefixFilterLists(McrPrefixFilterListArgs.builder()
                .description("Allow private networks")
                .addressFamily("IPv4")
                .entries(McrPrefixFilterListEntryArgs.builder()
                    .action("permit")
                    .prefix("10.0.0.0/8")
                    .ge(16.0)
                    .le(24.0)
                    .build())
                .build())
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
### Benefits of Standalone Resources

- **Individual Lifecycle Management**: Each prefix filter list has its own Pulumi state and lifecycle
- **Better Error Handling**: Failures in one list don't affect others
- **Enhanced Reusability**: Lists can be referenced and managed independently
- **Cleaner State**: Avoid complex nested object handling in Pulumi state
- **Import Support**: Easy migration of existing lists using `pulumi import`
### Migration Guide
#### Step 1: Inventory Existing Lists

Use the function to see what prefix filter lists you currently have:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as megaport from "@pulumi/megaport";

const existing = megaport.getMcrPrefixFilterLists({
    mcrId: "your-mcr-uid-here",
});
export const currentLists = existing.then(existing => existing.prefixFilterLists);
```
{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import pulumi_megaport as megaport

existing = megaport.get_mcr_prefix_filter_lists(mcr_id="your-mcr-uid-here")
pulumi.export("currentLists", existing.prefix_filter_lists)
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Megaport = Pulumi.Megaport;

return await Deployment.RunAsync(() =>
{
    var existing = Megaport.GetMcrPrefixFilterLists.Invoke(new()
    {
        McrId = "your-mcr-uid-here",
    });

    return new Dictionary<string, object?>
    {
        ["currentLists"] = existing.Apply(getMcrPrefixFilterListsResult => getMcrPrefixFilterListsResult.PrefixFilterLists),
    };
});

```
{{% /choosable %}}
{{% choosable language go %}}
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/megaport/megaport"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		existing, err := megaport.GetMcrPrefixFilterLists(ctx, &megaport.GetMcrPrefixFilterListsArgs{
			McrId: "your-mcr-uid-here",
		}, nil)
		if err != nil {
			return err
		}
		ctx.Export("currentLists", existing.PrefixFilterLists)
		return nil
	})
}
```
{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
variables:
  existing:
    fn::invoke:
      function: megaport:getMcrPrefixFilterLists
      arguments:
        mcrId: your-mcr-uid-here
outputs:
  currentLists: ${existing.prefixFilterLists}
```
{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.megaport.MegaportFunctions;
import com.pulumi.megaport.inputs.GetMcrPrefixFilterListsArgs;
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
        final var existing = MegaportFunctions.getMcrPrefixFilterLists(GetMcrPrefixFilterListsArgs.builder()
            .mcrId("your-mcr-uid-here")
            .build());

        ctx.export("currentLists", existing.prefixFilterLists());
    }
}
```
{{% /choosable %}}
{{< /chooser >}}
#### Step 2: Create Standalone Resources

For each existing list, create a corresponding resource:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as megaport from "@pulumi/megaport";

const migratedList1 = new megaport.McrPrefixFilterList("migrated_list_1", {
    mcrId: "your-mcr-uid-here",
    description: "Copy description from existing list",
    addressFamily: "IPv4",
    entries: [],
});
```
{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import pulumi_megaport as megaport

migrated_list1 = megaport.McrPrefixFilterList("migrated_list_1",
    mcr_id="your-mcr-uid-here",
    description="Copy description from existing list",
    address_family="IPv4",
    entries=[])
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Megaport = Pulumi.Megaport;

return await Deployment.RunAsync(() =>
{
    var migratedList1 = new Megaport.McrPrefixFilterList("migrated_list_1", new()
    {
        McrId = "your-mcr-uid-here",
        Description = "Copy description from existing list",
        AddressFamily = "IPv4",
        Entries = new[] {},
    });

});

```
{{% /choosable %}}
{{% choosable language go %}}
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/megaport/megaport"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := megaport.NewMcrPrefixFilterList(ctx, "migrated_list_1", &megaport.McrPrefixFilterListArgs{
			McrId:         pulumi.String("your-mcr-uid-here"),
			Description:   pulumi.String("Copy description from existing list"),
			AddressFamily: pulumi.String("IPv4"),
			Entries:       megaport.McrPrefixFilterListEntryArray{},
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
  migratedList1:
    type: megaport:McrPrefixFilterList
    name: migrated_list_1
    properties:
      mcrId: your-mcr-uid-here
      description: Copy description from existing list
      addressFamily: IPv4
      entries: []
```
{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.megaport.McrPrefixFilterList;
import com.pulumi.megaport.McrPrefixFilterListArgs;
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
        var migratedList1 = new McrPrefixFilterList("migratedList1", McrPrefixFilterListArgs.builder()
            .mcrId("your-mcr-uid-here")
            .description("Copy description from existing list")
            .addressFamily("IPv4")
            .entries()
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
#### Step 3: Import Existing Lists

Import each existing list to avoid recreation:

```bash
pulumi import megaport_mcr_prefix_filter_list.migrated_list_1 mcr-uid:prefix-list-id
```
#### Step 4: Update MCR Resource

Remove the `prefixFilterLists` attribute from your MCR resource and add a lifecycle rule:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as megaport from "@pulumi/megaport";

const example = new megaport.Mcr("example", {
    productName: "my-mcr",
    portSpeed: 1000,
    locationId: 5,
    contractTermMonths: 12,
});
```
{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import pulumi_megaport as megaport

example = megaport.Mcr("example",
    product_name="my-mcr",
    port_speed=1000,
    location_id=5,
    contract_term_months=12)
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Megaport = Pulumi.Megaport;

return await Deployment.RunAsync(() =>
{
    var example = new Megaport.Mcr("example", new()
    {
        ProductName = "my-mcr",
        PortSpeed = 1000,
        LocationId = 5,
        ContractTermMonths = 12,
    });

});

```
{{% /choosable %}}
{{% choosable language go %}}
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/megaport/megaport"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := megaport.NewMcr(ctx, "example", &megaport.McrArgs{
			ProductName:        pulumi.String("my-mcr"),
			PortSpeed:          pulumi.Float64(1000),
			LocationId:         pulumi.Float64(5),
			ContractTermMonths: pulumi.Float64(12),
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
    type: megaport:Mcr
    properties:
      productName: my-mcr
      portSpeed: 1000
      locationId: 5
      contractTermMonths: 12 # Add lifecycle rule to prevent drift warnings
```
{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.megaport.Mcr;
import com.pulumi.megaport.McrArgs;
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
        var example = new Mcr("example", McrArgs.builder()
            .productName("my-mcr")
            .portSpeed(1000.0)
            .locationId(5.0)
            .contractTermMonths(12.0)
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
#### Step 5: Verify Migration

Run `pulumi preview` to ensure no unexpected changes are detected.
### Mixed Usage Prevention

The provider includes validation to prevent managing the same prefix filter lists through both methods simultaneously. If you attempt to use both inline and standalone management for the same MCR, you'll receive warnings about potential conflicts.
### Deprecation Notice

The inline `prefixFilterLists` attribute in the MCR resource is deprecated and will be removed in a future version. We recommend migrating to standalone `megaport.McrPrefixFilterList` resources for better lifecycle management and improved state handling.
### Troubleshooting and Best Practices
#### Common Issues

**MCR Resource Shows Drift with Standalone Resources**

When using standalone `megaport.McrPrefixFilterList` resources, you should add a lifecycle rule to your MCR resource to prevent Pulumi from detecting drift on the `prefixFilterLists` attribute. This is necessary because:

1. The standalone prefix filter list resources manage the lists independently
2. The MCR resource still reads the lists from the API, which can cause Pulumi to detect "changes" even though the lists are being managed by the standalone resources
3. This applies to both newly created MCRs and existing ones - the lifecycle rule tells Pulumi to ignore differences in this attribute since it's being managed elsewhere

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as megaport from "@pulumi/megaport";

const example = new megaport.Mcr("example", {});
```
{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import pulumi_megaport as megaport

example = megaport.Mcr("example")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Megaport = Pulumi.Megaport;

return await Deployment.RunAsync(() =>
{
    var example = new Megaport.Mcr("example");

});

```
{{% /choosable %}}
{{% choosable language go %}}
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/megaport/megaport"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := megaport.NewMcr(ctx, "example", nil)
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
    type: megaport:Mcr
```
{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.megaport.Mcr;
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
        var example = new Mcr("example");

    }
}
```
{{% /choosable %}}
{{< /chooser >}}

**Why is this needed for new resources?** Even when creating a new MCR alongside standalone prefix filter list resources, Pulumi's refresh cycle will detect that the MCR has prefix filter lists attached (via the standalone resources), and without the lifecycle rule, it may show these as unexpected changes on subsequent plan/apply operations.

**Mixed Usage Warning**

If you see warnings about mixed usage, ensure you're not managing the same prefix filter lists through both inline and standalone methods simultaneously.

**Import Format**

When importing existing prefix filter lists, use the format `mcr_uid:prefix_list_id`:

```bash
# Get the MCR UID and prefix list ID from the Megaport Portal or API
pulumi import megaport_mcr_prefix_filter_list.example a1b2c3d4-5678-90ef-ghij-klmnopqrstuv:1234
```
#### Best Practices

- **Use Location IDs**: Always use location IDs instead of names for MCR placement (more stable)
- **Validate Prefix Ranges**: Ensure ge (greater than or equal) and le (less than or equal) values make sense for your prefix lengths
- **Group Related Lists**: Create logically grouped prefix filter lists for easier management
- **Test Migrations**: Always test migrations in a non-production environment first
- **Document Purposes**: Use descriptive names and descriptions for prefix filter lists
#### Example Production Configuration

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as megaport from "@pulumi/megaport";

// Production MCR with standalone prefix filter lists
const production = new megaport.Mcr("production", {
    productName: "prod-mcr",
    portSpeed: 2500,
    locationId: 1,
    contractTermMonths: 12,
    resourceTags: {
        Environment: "production",
        Owner: "network-team",
        Purpose: "multi-cloud-connectivity",
    },
});
// Allow internal corporate networks
const corporateNetworks = new megaport.McrPrefixFilterList("corporate_networks", {
    mcrId: production.productUid,
    description: "Corporate internal networks",
    addressFamily: "IPv4",
    entries: [
        {
            action: "permit",
            prefix: "10.100.0.0/16",
            ge: 24,
            le: 28,
        },
        {
            action: "permit",
            prefix: "10.200.0.0/16",
            ge: 24,
            le: 28,
        },
    ],
});
// Allow cloud provider networks
const cloudNetworks = new megaport.McrPrefixFilterList("cloud_networks", {
    mcrId: production.productUid,
    description: "AWS and Azure networks",
    addressFamily: "IPv4",
    entries: [{
        action: "permit",
        prefix: "172.16.0.0/12",
        ge: 16,
        le: 24,
    }],
});
// IPv6 support for future expansion
const ipv6Networks = new megaport.McrPrefixFilterList("ipv6_networks", {
    mcrId: production.productUid,
    description: "IPv6 corporate networks",
    addressFamily: "IPv6",
    entries: [{
        action: "permit",
        prefix: "2001:db8:100::/48",
        ge: 56,
        le: 64,
    }],
});
```
{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import pulumi_megaport as megaport

# Production MCR with standalone prefix filter lists
production = megaport.Mcr("production",
    product_name="prod-mcr",
    port_speed=2500,
    location_id=1,
    contract_term_months=12,
    resource_tags={
        "Environment": "production",
        "Owner": "network-team",
        "Purpose": "multi-cloud-connectivity",
    })
# Allow internal corporate networks
corporate_networks = megaport.McrPrefixFilterList("corporate_networks",
    mcr_id=production.product_uid,
    description="Corporate internal networks",
    address_family="IPv4",
    entries=[
        {
            "action": "permit",
            "prefix": "10.100.0.0/16",
            "ge": 24,
            "le": 28,
        },
        {
            "action": "permit",
            "prefix": "10.200.0.0/16",
            "ge": 24,
            "le": 28,
        },
    ])
# Allow cloud provider networks
cloud_networks = megaport.McrPrefixFilterList("cloud_networks",
    mcr_id=production.product_uid,
    description="AWS and Azure networks",
    address_family="IPv4",
    entries=[{
        "action": "permit",
        "prefix": "172.16.0.0/12",
        "ge": 16,
        "le": 24,
    }])
# IPv6 support for future expansion
ipv6_networks = megaport.McrPrefixFilterList("ipv6_networks",
    mcr_id=production.product_uid,
    description="IPv6 corporate networks",
    address_family="IPv6",
    entries=[{
        "action": "permit",
        "prefix": "2001:db8:100::/48",
        "ge": 56,
        "le": 64,
    }])
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Megaport = Pulumi.Megaport;

return await Deployment.RunAsync(() =>
{
    // Production MCR with standalone prefix filter lists
    var production = new Megaport.Mcr("production", new()
    {
        ProductName = "prod-mcr",
        PortSpeed = 2500,
        LocationId = 1,
        ContractTermMonths = 12,
        ResourceTags =
        {
            { "Environment", "production" },
            { "Owner", "network-team" },
            { "Purpose", "multi-cloud-connectivity" },
        },
    });

    // Allow internal corporate networks
    var corporateNetworks = new Megaport.McrPrefixFilterList("corporate_networks", new()
    {
        McrId = production.ProductUid,
        Description = "Corporate internal networks",
        AddressFamily = "IPv4",
        Entries = new[]
        {
            new Megaport.Inputs.McrPrefixFilterListEntryArgs
            {
                Action = "permit",
                Prefix = "10.100.0.0/16",
                Ge = 24,
                Le = 28,
            },
            new Megaport.Inputs.McrPrefixFilterListEntryArgs
            {
                Action = "permit",
                Prefix = "10.200.0.0/16",
                Ge = 24,
                Le = 28,
            },
        },
    });

    // Allow cloud provider networks
    var cloudNetworks = new Megaport.McrPrefixFilterList("cloud_networks", new()
    {
        McrId = production.ProductUid,
        Description = "AWS and Azure networks",
        AddressFamily = "IPv4",
        Entries = new[]
        {
            new Megaport.Inputs.McrPrefixFilterListEntryArgs
            {
                Action = "permit",
                Prefix = "172.16.0.0/12",
                Ge = 16,
                Le = 24,
            },
        },
    });

    // IPv6 support for future expansion
    var ipv6Networks = new Megaport.McrPrefixFilterList("ipv6_networks", new()
    {
        McrId = production.ProductUid,
        Description = "IPv6 corporate networks",
        AddressFamily = "IPv6",
        Entries = new[]
        {
            new Megaport.Inputs.McrPrefixFilterListEntryArgs
            {
                Action = "permit",
                Prefix = "2001:db8:100::/48",
                Ge = 56,
                Le = 64,
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
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/megaport/megaport"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Production MCR with standalone prefix filter lists
		production, err := megaport.NewMcr(ctx, "production", &megaport.McrArgs{
			ProductName:        pulumi.String("prod-mcr"),
			PortSpeed:          pulumi.Float64(2500),
			LocationId:         pulumi.Float64(1),
			ContractTermMonths: pulumi.Float64(12),
			ResourceTags: pulumi.StringMap{
				"Environment": pulumi.String("production"),
				"Owner":       pulumi.String("network-team"),
				"Purpose":     pulumi.String("multi-cloud-connectivity"),
			},
		})
		if err != nil {
			return err
		}
		// Allow internal corporate networks
		_, err = megaport.NewMcrPrefixFilterList(ctx, "corporate_networks", &megaport.McrPrefixFilterListArgs{
			McrId:         production.ProductUid,
			Description:   pulumi.String("Corporate internal networks"),
			AddressFamily: pulumi.String("IPv4"),
			Entries: megaport.McrPrefixFilterListEntryArray{
				&megaport.McrPrefixFilterListEntryArgs{
					Action: pulumi.String("permit"),
					Prefix: pulumi.String("10.100.0.0/16"),
					Ge:     pulumi.Float64(24),
					Le:     pulumi.Float64(28),
				},
				&megaport.McrPrefixFilterListEntryArgs{
					Action: pulumi.String("permit"),
					Prefix: pulumi.String("10.200.0.0/16"),
					Ge:     pulumi.Float64(24),
					Le:     pulumi.Float64(28),
				},
			},
		})
		if err != nil {
			return err
		}
		// Allow cloud provider networks
		_, err = megaport.NewMcrPrefixFilterList(ctx, "cloud_networks", &megaport.McrPrefixFilterListArgs{
			McrId:         production.ProductUid,
			Description:   pulumi.String("AWS and Azure networks"),
			AddressFamily: pulumi.String("IPv4"),
			Entries: megaport.McrPrefixFilterListEntryArray{
				&megaport.McrPrefixFilterListEntryArgs{
					Action: pulumi.String("permit"),
					Prefix: pulumi.String("172.16.0.0/12"),
					Ge:     pulumi.Float64(16),
					Le:     pulumi.Float64(24),
				},
			},
		})
		if err != nil {
			return err
		}
		// IPv6 support for future expansion
		_, err = megaport.NewMcrPrefixFilterList(ctx, "ipv6_networks", &megaport.McrPrefixFilterListArgs{
			McrId:         production.ProductUid,
			Description:   pulumi.String("IPv6 corporate networks"),
			AddressFamily: pulumi.String("IPv6"),
			Entries: megaport.McrPrefixFilterListEntryArray{
				&megaport.McrPrefixFilterListEntryArgs{
					Action: pulumi.String("permit"),
					Prefix: pulumi.String("2001:db8:100::/48"),
					Ge:     pulumi.Float64(56),
					Le:     pulumi.Float64(64),
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
  # Production MCR with standalone prefix filter lists
  production:
    type: megaport:Mcr
    properties:
      productName: prod-mcr
      portSpeed: 2500
      locationId: 1 # Use stable location ID
      contractTermMonths: 12
      resourceTags:
        Environment: production
        Owner: network-team
        Purpose: multi-cloud-connectivity
  # Allow internal corporate networks
  corporateNetworks:
    type: megaport:McrPrefixFilterList
    name: corporate_networks
    properties:
      mcrId: ${production.productUid}
      description: Corporate internal networks
      addressFamily: IPv4
      entries:
        - action: permit
          prefix: 10.100.0.0/16
          ge: 24
          le: 28
        - action: permit
          prefix: 10.200.0.0/16
          ge: 24
          le: 28
  # Allow cloud provider networks
  cloudNetworks:
    type: megaport:McrPrefixFilterList
    name: cloud_networks
    properties:
      mcrId: ${production.productUid}
      description: AWS and Azure networks
      addressFamily: IPv4
      entries:
        - action: permit
          prefix: 172.16.0.0/12
          ge: 16
          le: 24
  # IPv6 support for future expansion
  ipv6Networks:
    type: megaport:McrPrefixFilterList
    name: ipv6_networks
    properties:
      mcrId: ${production.productUid}
      description: IPv6 corporate networks
      addressFamily: IPv6
      entries:
        - action: permit
          prefix: 2001:db8:100::/48
          ge: 56
          le: 64
```
{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.megaport.Mcr;
import com.pulumi.megaport.McrArgs;
import com.pulumi.megaport.McrPrefixFilterList;
import com.pulumi.megaport.McrPrefixFilterListArgs;
import com.pulumi.megaport.inputs.McrPrefixFilterListEntryArgs;
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
        // Production MCR with standalone prefix filter lists
        var production = new Mcr("production", McrArgs.builder()
            .productName("prod-mcr")
            .portSpeed(2500.0)
            .locationId(1.0)
            .contractTermMonths(12.0)
            .resourceTags(Map.ofEntries(
                Map.entry("Environment", "production"),
                Map.entry("Owner", "network-team"),
                Map.entry("Purpose", "multi-cloud-connectivity")
            ))
            .build());

        // Allow internal corporate networks
        var corporateNetworks = new McrPrefixFilterList("corporateNetworks", McrPrefixFilterListArgs.builder()
            .mcrId(production.productUid())
            .description("Corporate internal networks")
            .addressFamily("IPv4")
            .entries(
                McrPrefixFilterListEntryArgs.builder()
                    .action("permit")
                    .prefix("10.100.0.0/16")
                    .ge(24.0)
                    .le(28.0)
                    .build(),
                McrPrefixFilterListEntryArgs.builder()
                    .action("permit")
                    .prefix("10.200.0.0/16")
                    .ge(24.0)
                    .le(28.0)
                    .build())
            .build());

        // Allow cloud provider networks
        var cloudNetworks = new McrPrefixFilterList("cloudNetworks", McrPrefixFilterListArgs.builder()
            .mcrId(production.productUid())
            .description("AWS and Azure networks")
            .addressFamily("IPv4")
            .entries(McrPrefixFilterListEntryArgs.builder()
                .action("permit")
                .prefix("172.16.0.0/12")
                .ge(16.0)
                .le(24.0)
                .build())
            .build());

        // IPv6 support for future expansion
        var ipv6Networks = new McrPrefixFilterList("ipv6Networks", McrPrefixFilterListArgs.builder()
            .mcrId(production.productUid())
            .description("IPv6 corporate networks")
            .addressFamily("IPv6")
            .entries(McrPrefixFilterListEntryArgs.builder()
                .action("permit")
                .prefix("2001:db8:100::/48")
                .ge(56.0)
                .le(64.0)
                .build())
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Example Use Cases

A suite of tested examples is in the examples directory
## 🚨 BREAKING CHANGE: Location Function Migration
### ⚠️ URGENT: `siteCode` Filtering No Longer Supported

**If you are using `siteCode` to filter locations in your Pulumi configurations, you must update your code immediately or your configurations will fail.**

The Megaport Location API has been upgraded to v3, and several important changes affect how you interact with location data:
#### ❌ What No Longer Works

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as megaport from "@pulumi/megaport";

// THIS WILL FAIL - site_code filtering is no longer supported
const brokenExample = megaport.getLocation({
    siteCode: "NTT-TOK",
});
```
{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import pulumi_megaport as megaport

# THIS WILL FAIL - site_code filtering is no longer supported
broken_example = megaport.get_location(site_code="NTT-TOK")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Megaport = Pulumi.Megaport;

return await Deployment.RunAsync(() =>
{
    // THIS WILL FAIL - site_code filtering is no longer supported
    var brokenExample = Megaport.GetLocation.Invoke(new()
    {
        SiteCode = "NTT-TOK",
    });

});

```
{{% /choosable %}}
{{% choosable language go %}}
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/megaport/megaport"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// THIS WILL FAIL - site_code filtering is no longer supported
		_, err := megaport.GetLocation(ctx, &megaport.GetLocationArgs{
			SiteCode: pulumi.StringRef("NTT-TOK"),
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
  # THIS WILL FAIL - site_code filtering is no longer supported
  brokenExample:
    fn::invoke:
      function: megaport:getLocation
      arguments:
        siteCode: NTT-TOK
```
{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.megaport.MegaportFunctions;
import com.pulumi.megaport.inputs.GetLocationArgs;
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
        // THIS WILL FAIL - site_code filtering is no longer supported
        final var brokenExample = MegaportFunctions.getLocation(GetLocationArgs.builder()
            .siteCode("NTT-TOK")
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
#### ✅ What You Should Use Instead

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as megaport from "@pulumi/megaport";

// ✅ RECOMMENDED: Use location ID for most reliable results
const recommended = megaport.getLocation({
    id: 123,
});
// ✅ ALTERNATIVE: Use location name (may change over time)
const alternative = megaport.getLocation({
    name: "NextDC B1",
});
export const locationIdForNextdcB1 = alternative.then(alternative => alternative.id);
```
{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import pulumi_megaport as megaport

# ✅ RECOMMENDED: Use location ID for most reliable results
recommended = megaport.get_location(id=123)
# ✅ ALTERNATIVE: Use location name (may change over time)
alternative = megaport.get_location(name="NextDC B1")
pulumi.export("locationIdForNextdcB1", alternative.id)
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Megaport = Pulumi.Megaport;

return await Deployment.RunAsync(() =>
{
    // ✅ RECOMMENDED: Use location ID for most reliable results
    var recommended = Megaport.GetLocation.Invoke(new()
    {
        Id = 123,
    });

    // ✅ ALTERNATIVE: Use location name (may change over time)
    var alternative = Megaport.GetLocation.Invoke(new()
    {
        Name = "NextDC B1",
    });

    return new Dictionary<string, object?>
    {
        ["locationIdForNextdcB1"] = alternative.Apply(getLocationResult => getLocationResult.Id),
    };
});

```
{{% /choosable %}}
{{% choosable language go %}}
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/megaport/megaport"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// ✅ RECOMMENDED: Use location ID for most reliable results
		_, err := megaport.GetLocation(ctx, &megaport.GetLocationArgs{
			Id: pulumi.Float64Ref(123),
		}, nil)
		if err != nil {
			return err
		}
		// ✅ ALTERNATIVE: Use location name (may change over time)
		alternative, err := megaport.GetLocation(ctx, &megaport.GetLocationArgs{
			Name: pulumi.StringRef("NextDC B1"),
		}, nil)
		if err != nil {
			return err
		}
		ctx.Export("locationIdForNextdcB1", alternative.Id)
		return nil
	})
}
```
{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
variables:
  # ✅ RECOMMENDED: Use location ID for most reliable results
  recommended:
    fn::invoke:
      function: megaport:getLocation
      arguments:
        id: 123
  # ✅ ALTERNATIVE: Use location name (may change over time)
  alternative:
    fn::invoke:
      function: megaport:getLocation
      arguments:
        name: NextDC B1
outputs:
  # 💡 TIP: Save the location ID for future use
  locationIdForNextdcB1: ${alternative.id}
```
{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.megaport.MegaportFunctions;
import com.pulumi.megaport.inputs.GetLocationArgs;
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
        // ✅ RECOMMENDED: Use location ID for most reliable results
        final var recommended = MegaportFunctions.getLocation(GetLocationArgs.builder()
            .id(123)
            .build());

        // ✅ ALTERNATIVE: Use location name (may change over time)
        final var alternative = MegaportFunctions.getLocation(GetLocationArgs.builder()
            .name("NextDC B1")
            .build());

        ctx.export("locationIdForNextdcB1", alternative.id());
    }
}
```
{{% /choosable %}}
{{< /chooser >}}
### 🔧 Migration Guide

**Step 1: Identify affected configurations**
Search your Pulumi files for any usage of `siteCode`:

```bash
grep -r "site_code" *.tf
```

**Step 2: Replace with `id` or `name`**

- **Best option**: Replace with the location `id` (most stable)
- **Alternative**: Replace with the location `name` (may change over time)

**Step 3: Update deprecated field usage**
Several location fields are now deprecated and will show warnings:
### 📋 Complete Migration Checklist

- <input disabled="" type="checkbox"> Replace all `siteCode = "..."` filters with `id = ...` or `name = "..."`
- <input disabled="" type="checkbox"> Remove any code that depends on deprecated fields
- <input disabled="" type="checkbox"> Test your configurations thoroughly
- <input disabled="" type="checkbox"> Update any documentation or comments
### 🆘 Need Help?

If you need to find the location ID for a specific site code, you can:

1. **Use Pulumi function**: Query by name to get the ID:

   {{< chooser language "typescript,python,go,csharp,java,yaml" >}}
   {{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as megaport from "@pulumi/megaport";

const lookup = megaport.getLocation({
    name: "Your Location Name",
});
export const locationId = lookup.then(lookup => lookup.id);
```
{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import pulumi_megaport as megaport

lookup = megaport.get_location(name="Your Location Name")
pulumi.export("locationId", lookup.id)
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Megaport = Pulumi.Megaport;

return await Deployment.RunAsync(() =>
{
    var lookup = Megaport.GetLocation.Invoke(new()
    {
        Name = "Your Location Name",
    });

    return new Dictionary<string, object?>
    {
        ["locationId"] = lookup.Apply(getLocationResult => getLocationResult.Id),
    };
});

```
{{% /choosable %}}
{{% choosable language go %}}
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/megaport/megaport"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		lookup, err := megaport.GetLocation(ctx, &megaport.GetLocationArgs{
			Name: pulumi.StringRef("Your Location Name"),
		}, nil)
		if err != nil {
			return err
		}
		ctx.Export("locationId", lookup.Id)
		return nil
	})
}
```
{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
variables:
  lookup:
    fn::invoke:
      function: megaport:getLocation
      arguments:
        name: Your Location Name
outputs:
  locationId: ${lookup.id}
```
{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.megaport.MegaportFunctions;
import com.pulumi.megaport.inputs.GetLocationArgs;
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
        final var lookup = MegaportFunctions.getLocation(GetLocationArgs.builder()
            .name("Your Location Name")
            .build());

        ctx.export("locationId", lookup.id());
    }
}
```
{{% /choosable %}}
{{< /chooser >}}
```

2. **Use the API directly**: Call `GET /v3/locations` to see all available locations
3. **Contact Support**: Megaport support can help map site codes to location IDs

---
## Datacenter Location Function

Locations for Megaport Data Centers can be retrieved using the Locations Function in the Megaport Pulumi Provider.
### Using Location IDs for Reliable Configurations

In the Megaport API, data center location names and site codes might occasionally change as facilities are rebranded, merged, or updated. However, the location ID remains constant and is the most reliable way to identify a data center.

If your Pulumi configurations or API scripts rely on a location name or code that changes, those integrations will fail to execute until the references are manually updated. This can lead to automation failures, deployment delays, or service disruptions.

**To ensure consistency and reliability, we strongly recommend using the location ID instead of the name or site code when integrating with the Megaport API or defining resources in Pulumi.**

**Current supported search methods:**

- `id` - **RECOMMENDED** (most reliable and stable)
- `name` - Alternative option (may change over time)

Examples:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as megaport from "@pulumi/megaport";

// ✅ RECOMMENDED: Use ID for most reliable results
const stableExample = megaport.getLocation({
    id: 5,
});
// ✅ ALTERNATIVE: Use name (less stable, may change)
const nameExample = megaport.getLocation({
    name: "NextDC B1",
});
export const locationIdForNextdcB1 = nameExample.then(nameExample => nameExample.id);
```
{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import pulumi_megaport as megaport

# ✅ RECOMMENDED: Use ID for most reliable results
stable_example = megaport.get_location(id=5)
# ✅ ALTERNATIVE: Use name (less stable, may change)
name_example = megaport.get_location(name="NextDC B1")
pulumi.export("locationIdForNextdcB1", name_example.id)
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Megaport = Pulumi.Megaport;

return await Deployment.RunAsync(() =>
{
    // ✅ RECOMMENDED: Use ID for most reliable results
    var stableExample = Megaport.GetLocation.Invoke(new()
    {
        Id = 5,
    });

    // ✅ ALTERNATIVE: Use name (less stable, may change)
    var nameExample = Megaport.GetLocation.Invoke(new()
    {
        Name = "NextDC B1",
    });

    return new Dictionary<string, object?>
    {
        ["locationIdForNextdcB1"] = nameExample.Apply(getLocationResult => getLocationResult.Id),
    };
});

```
{{% /choosable %}}
{{% choosable language go %}}
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/megaport/megaport"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// ✅ RECOMMENDED: Use ID for most reliable results
		_, err := megaport.GetLocation(ctx, &megaport.GetLocationArgs{
			Id: pulumi.Float64Ref(5),
		}, nil)
		if err != nil {
			return err
		}
		// ✅ ALTERNATIVE: Use name (less stable, may change)
		nameExample, err := megaport.GetLocation(ctx, &megaport.GetLocationArgs{
			Name: pulumi.StringRef("NextDC B1"),
		}, nil)
		if err != nil {
			return err
		}
		ctx.Export("locationIdForNextdcB1", nameExample.Id)
		return nil
	})
}
```
{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
variables:
  # ✅ RECOMMENDED: Use ID for most reliable results
  stableExample:
    fn::invoke:
      function: megaport:getLocation
      arguments:
        id: 5
  # ✅ ALTERNATIVE: Use name (less stable, may change)
  nameExample:
    fn::invoke:
      function: megaport:getLocation
      arguments:
        name: NextDC B1
outputs:
  # 💡 TIP: Save the location ID for future use
  locationIdForNextdcB1: ${nameExample.id}
```
{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.megaport.MegaportFunctions;
import com.pulumi.megaport.inputs.GetLocationArgs;
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
        // ✅ RECOMMENDED: Use ID for most reliable results
        final var stableExample = MegaportFunctions.getLocation(GetLocationArgs.builder()
            .id(5)
            .build());

        // ✅ ALTERNATIVE: Use name (less stable, may change)
        final var nameExample = MegaportFunctions.getLocation(GetLocationArgs.builder()
            .name("NextDC B1")
            .build());

        ctx.export("locationIdForNextdcB1", nameExample.id());
    }
}
```
{{% /choosable %}}
{{< /chooser >}}
### Location Reference Resources

For the most up-to-date list of Megaport data center locations:

- **[Megaport Location IDs Documentation](https://docs.megaport.com/enabled-locations/location-ids/)** - Complete list of location IDs (dynamically updated with the API)
- **[Megaport API GET /v3/locations](https://dev.megaport.com/#7c8e0706-e138-4d9a-bc4f-0419d97604cf)** - Programmatic access to location data
## Partner Port Stability

When using filter criteria to select partner ports (used to connect to cloud service providers), the specific partner port (and therefore UID) that best matches your filters may change over time as Megaport manages capacity by rotating ports. This can lead to unexpected warning messages during Pulumi operations even when the VXCs themselves are not being modified:

```
Warning: VXC B-End product UID is from a partner port, therefore it will not be changed.
```

This warning appears because Pulumi detects a difference in the partner port UID even when applying changes unrelated to those specific VXCs.
### Workaround

To prevent these warnings and ensure configuration stability, we recommend explicitly specifying the `productUid` in your partner port function once your connections are established:
## End-of-Term Cancellation

By default, when Pulumi deletes resources, they are immediately cancelled in the Megaport portal. However, you may prefer to have resources marked for cancellation at the end of their current billing term instead of immediate cancellation.

The provider supports this with the `cancelAtEndOfTerm` configuration option:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    megaport:acceptPurchaseTerms:
        value: true
    megaport:accessKey:
        value: your-access-key
    megaport:cancelAtEndOfTerm:
        value: true
    megaport:environment:
        value: production
    megaport:secretKey:
        value: your-secret-key

```

**Important notes:**

- This feature is currently only supported for Single Ports and LAG Ports
- For other resource types, the option will be ignored and immediate cancellation will occur
- When `cancelAtEndOfTerm` is set to `true`, resources will show as "CANCELLING" in the Megaport portal until the end of their billing term
- Resources are removed from Pulumi state as soon as the API call returns successfully, regardless of whether immediate or end-of-term cancellation is used
- If you reapply your configuration after a resource has been deleted, Pulumi will create a new resource, even if the original resource is still visible in the Megaport portal with "CANCELLING" status