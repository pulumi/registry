---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumiverse/pulumi-harbor/v3.10.21/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Harbor
meta_desc: Provides an overview of the Harbor Provider for Pulumi.
layout: overview
---

The `Harbor` provider for Pulumi can be used to provision any of the resources available with `Harbor`.

## Example

{{< chooser language "javascript,typescript,python,go,csharp" >}}

{{% choosable language javascript %}}

```javascript
"use strict";
const harbor = require("@pulumiverse/harbor");

const registry = new harbor.Registry("registry", {
    providerName: "docker-hub",
    endpointUrl: "https://hub.docker.com",
    name: "pulumi-harbor"
})

const project = new harbor.Project("project", {
    name: "pulumi-harbor",
    registryId: registry.registryId,
    public: "true",
})
```

{{% /choosable %}}

{{% choosable language typescript %}}

```typescript
import * as harbor from '@pulumiverse/harbor';

let registry = new harbor.Registry('registry', {
    providerName: "docker-hub",
    endpointUrl: "https://hub.docker.com",
    name: "pulumi-harbor"
});

let project = new harbor.Project('project', {
    name: "pulumi-harbor",
    registryId: registry.registryId,
    public: "true",
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumiverse_harbor as harbor
import pulumi

registry = harbor.Registry("registry", name="pulumi-harbor",
                           endpoint_url="https://harbor.pulumi.com",
                           provider_name="docker-hub")

project = harbor.Project("project", name="pulumi-harbor",
                         registry_id=registry.registry_id,
                         public="true")
```

{{% /choosable %}}
{{% choosable language go %}}

```go
package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-harbor/sdk/v3/go/harbor"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		registry, err := harbor.NewRegistry(ctx, "registry", &harbor.RegistryArgs{
			ProviderName: pulumi.String("docker-hub"),
			EndpointUrl:  pulumi.String("https://hub.docker.com"),
			Name:         pulumi.String("pulumi-harbor"),
		})
		if err != nil {
			return err
		}

		_, err = harbor.NewProject(ctx, "project", &harbor.ProjectArgs{
			Name:       pulumi.String("pulumi-harbor"),
			Public:     pulumi.String("true"),
			RegistryId: registry.RegistryId,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
```

{{% /choosable %}}

{{% choosable language csharp %}}

```csharp
using System.Collections.Generic;
using Pulumi;
using Pulumiverse.Harbor;

return await Deployment.RunAsync(() =>
{
   var registry = new Registry("registry", new RegistryArgs
   {
      ProviderName= "docker-hub",
      EndpointUrl="https://hub.docker.com",
      Name= "pulumi-harbor",
   });
   var project = new Project("project", new ProjectArgs
   {
      RegistryId= registry.RegistryId,
      Name= "pulumi-harbor",
      Public= "true" 
   });
});
```

{{% /choosable %}}

{{< /chooser >}}
