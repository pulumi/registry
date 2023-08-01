---
title: Prodvana
meta_desc: Provides an overview of the Prodvana Provider for Pulumi.
layout: package
---

The [Prodvana](http://prodvana.io/) provider for Pulumi can be used to provision resources within your Prodvana organization. For example you can create and manage Runtimes, Applications, and Release Channels.

The Prodvana provider must be configured with credentials to manage the resources in your Prodvana organization.

## Example

{{< chooser language "typescript,python,go" />}}

{{% choosable language typescript %}}

```typescript
import * as prodvana from "@prodvana/pulumi-prodvana";

const app = new prodvana.Application("my-app", {
    name: "my-app",
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_prodvana as prodvana

app = prodvana.Application("my-app", name="my-app")
```

{{% /choosable %}}
{{% choosable language go %}}

```go
package main

import (
	"github.com/prodvana/pulumi-prodvana/sdk/go/prodvana"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		_, err := prodvana.NewApplication(ctx, "my-app", &prodvana.ApplicationArgs{
			Name: pulumi.String("my-app"),
		})
		if err != nil {
			return err
		}

		return nil
	})
}
```

{{% /choosable %}}
