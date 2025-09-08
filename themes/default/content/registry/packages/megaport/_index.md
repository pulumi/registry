---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/megaport/megaport/1.4.2/index.md
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
Virtual Cross Connects (VXCs), Megaport Cloud Routers (MCRs), and Partner VXCs.

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

        ctx.export("locationIdForNextdcB1", alternative.applyValue(getLocationResult -> getLocationResult.id()));
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

        ctx.export("locationId", lookup.applyValue(getLocationResult -> getLocationResult.id()));
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

        ctx.export("locationIdForNextdcB1", nameExample.applyValue(getLocationResult -> getLocationResult.id()));
    }
}
```
{{% /choosable %}}
{{< /chooser >}}

**Important:** Location IDs never change and provide the most reliable and deterministic behavior. Location names may be updated over time, which could cause Pulumi configurations to break unexpectedly.

The most up-to-date listing of Megaport Datacenter Locations can be accessed through the Megaport API at `GET /v3/locations`
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