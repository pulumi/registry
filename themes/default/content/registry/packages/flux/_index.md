---
title: Flux
meta_desc: Provides an overview of the Flux provider for Pulumi.
layout: overview
---

The Pulumi Flux provider lets you manage [Flux](https://fluxcd.io/) resources.

## Example

{{< chooser language "typescript,python,go,csharp" >}}
{{% choosable language typescript %}}

```typescript
import * as flux from "@worawat/flux";
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_flux as flux
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
 "fmt"
 flux "github.com/oun/pulumi-flux/sdk/go/flux"
 "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
```

{{% /choosable %}}
{{% choosable language csharp %}}

```csharp
using Pulumi;
using Pulumi.Flux;

```

{{% /choosable %}}
{{< /chooser >}}
