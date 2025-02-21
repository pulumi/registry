---
# WARNING: this file was fetched from https://raw.githubusercontent.com/three141/pulumi-argocd/v1.0.1/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
title: ArgoCD
meta_desc: Provides an overview of the ArgoCD Provider for Pulumi.
layout: overview
---

The Argo CD Resource Provider lets you manage [Argo CD](https://argoproj.github.io/cd/) resources.
The ArgoCD provider must be configured with credentials to deploy and update resources in ArgoCD.

## Example

{{< chooser language "typescript,python,csharp,go" >}}
{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as argocd from "@three14/pulumi-argocd";

const publicNginxHelm = new argocd.Repository("public_nginx_helm", {
    repo: "https://helm.nginx.com/stable",
    name: "nginx-stable",
    type: "helm",
});

```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi
import pulumi_argocd as argocd

public_nginx_helm = argocd.Repository("public_nginx_helm",
    repo="https://helm.nginx.com/stable",
    name="nginx-stable",
    type="helm")

```

{{% /choosable %}}
{{% choosable language csharp %}}

```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Argocd = Three14.Argocd;

return await Deployment.RunAsync(() => 
{
    var publicNginxHelm = new Argocd.Repository("public_nginx_helm", new()
    {
        Repo = "https://helm.nginx.com/stable",
        Name = "nginx-stable",
        Type = "helm",
    });
});


```

{{% /choosable %}}
{{% choosable language go %}}

```go
package main

import (
	"github.com/Three141/pulumi-argocd/sdk/go/argocd"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Public Helm repository
		_, err := argocd.NewRepository(ctx, "public_nginx_helm", &argocd.RepositoryArgs{
			Repo: pulumi.String("https://helm.nginx.com/stable"),
			Name: pulumi.String("nginx-stable"),
			Type: pulumi.String("helm"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

```

{{% /choosable %}}
{{< /chooser >}}
