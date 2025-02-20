---
# WARNING: this file was fetched from https://raw.githubusercontent.com/three141/pulumi-argocd/v1.0.0/docs/_index.md
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
import * as argocd from "@three14/pulumi-argocd";

```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_argocd as argocd


```

{{% /choosable %}}
{{% choosable language csharp %}}

```csharp
using Argocd = Three14.Argocd;


```

{{% /choosable %}}
{{% choosable language go %}}

```go
package main

import (
    "github.com/Three141/pulumi-argocd/sdk/go/argocd"
)

```

{{% /choosable %}}
{{< /chooser >}}
