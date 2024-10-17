---
title: CrateDB
meta_desc: Provides an overview of the CrateDB Provider for Pulumi.
layout: package
---

The CrateDB provider for Pulumi can be used to provision the resources available in [CrateDB](https://cratedb.com/database/).

The CrateDB provider must be configured with credentials to deploy and update resources in CrateDB; see [Installation & Configuration](./installation-configuration) for instructions.

## Example

{{< chooser language "typescript,python,go,csharp" >}}
{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as cratedb from "@komminarlabs/cratedb";

export const organization = new cratedb.Organization("default", {
    name: "default",
});

export const organizationName = organization.name;

console.log(`Organization Name: {organizationName}`);
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import komminarlabs_cratedb as cratedb

organization = cratedb.Organization(
    "default",
    name="default",
)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
package main

import (
	"github.com/komminarlabs/pulumi-cratedb/sdk/go/cratedb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		defaultOrg, err := cratedb.NewOrganization(ctx, "default", &cratedb.OrganizationArgs{
			Name: pulumi.String("default"),
		})
		if err != nil {
			return err
		}

		ctx.Export("defaultOrgName", defaultOrg.Name)
		return nil
	})
}
```

{{% /choosable %}}

{{% choosable language csharp %}}

```csharp
using Pulumi;
using KomminarLabs.CrateDB;

class CrateDB : Stack
{
    public CrateDB()
    {
        var db = new Organization("default", new OrganizationArgs{
            Name: "default"
        });
    }
}
```

{{% /choosable %}}

{{< /chooser >}}
