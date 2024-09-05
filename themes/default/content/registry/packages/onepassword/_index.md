---
title: 1Password
meta_desc: Provides an overview of the 1Password Provider for Pulumi.
layout: package
---

The 1Password provider for Pulumi allows you to access and manage items in your [1Password](https://1password.com) vaults.
You'll need to configure the 1Password provider with credentials to access and manage items in 1Password.

## Example

{{< chooser language "javascript,typescript,python,go" >}}

{{% choosable language javascript %}}

```javascript
const pulumi = require("@pulumi/pulumi");
const onepassword = require("@1password/pulumi-onepassword");

const example = onepassword.getItem({
    vault: data.onepassword_vault.example.uuid,
    uuid: onepassword_item.demo_sections.uuid,
});
```

{{% /choosable %}}
{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as onepassword from "@1password/pulumi-onepassword";

const example = onepassword.getItem({
    vault: data.onepassword_vault.example.uuid,
    uuid: onepassword_item.demo_sections.uuid,
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi
import pulumi_onepassword as onepassword

example = onepassword.get_item(vault=data["onepassword_vault"]["example"]["uuid"],
    uuid=onepassword_item["demo_sections"]["uuid"])
```

{{% /choosable %}}
{{% choosable language go %}}

```go
package main

import (

	"github.com/1Password/pulumi-onepassword/sdk/go/onepassword"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := onepassword.LookupItem(ctx, &onepassword.LookupItemArgs{
			Vault: data.Onepassword_vault.Example.Uuid,
			Uuid:  pulumi.StringRef(onepassword_item.Demo_sections.Uuid),
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}
```

{{% /choosable %}}

<!--

{{% choosable language csharp %}}

```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Onepassword = Pulumi.Onepassword;

return await Deployment.RunAsync(() =&gt; 
{
    var example = Onepassword.GetItem.Invoke(new()
    {
        Vault = data.Onepassword_vault.Example.Uuid,
        Uuid = onepassword_item.Demo_sections.Uuid,
    });

});
```

{{% /choosable %}}

-->

{{< /chooser >}}
