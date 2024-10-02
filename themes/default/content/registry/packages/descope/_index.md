---
title: Descope
meta_desc: Provides an overview of the Descope Provider for Pulumi.
layout: package
---

The `Descope` provider for Pulumi can be used to provision any of the resources available with `Descope`.

## Example

{{< chooser language "javascript,typescript,python,go,csharp,yaml" >}}

{{% choosable language javascript %}}

```javascript
"use strict";
const descope = require("@descope/pulumi-descope");

export const project = new descope.Project("descope-project", {});
```

{{% /choosable %}}

{{% choosable language typescript %}}

```typescript
import * as descope from "@descope/pulumi-descope";

export const project = new descope.Project("descope-project", {});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi
import descope_pulumi

project = descope_pulumi.Project("pulumi-py-test", environment="production")

pulumi.export("project", project)
```

{{% /choosable %}}

{{% choosable language csharp %}}

```csharp
using Pulumi;
using Descope.Pulumi.Descope;

return await Deployment.RunAsync(() =>
{
	var project = new Project("descope-project");
});
```

{{% /choosable %}}

{{% choosable language go %}}

`````go
package main

import (
	"github.com/descope/pulumi-descope/sdk/go/descope"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		project, err := descope.NewProject(ctx, "pulumi-go-test", &descope.ProjectArgs{Environment: pulumi.String("production")})
		if err != nil {
			return err
		}
		ctx.Export("project", project)
		return nil
	})
}
```

{{% /choosable %}}

{{% choosable language yaml %}}

````yaml
resources:
  descopeProject:
    type: descope:project:Project

```

{{% /choosable %}}
{{< /chooser >}}
`````
