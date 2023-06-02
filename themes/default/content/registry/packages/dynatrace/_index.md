---
title: Dynatrace
meta_desc: Provides an overview of the Dynatrace Provider for Pulumi.
layout: package
---

The Dynatrace provider for Pulumi can be used to provision any of the cloud resources available in [Pulumi](https://www.dynatrace.com/).
The Dynatrace provider must be configured with credentials to deploy and update resources in Dynatrace.

## Example

{{< chooser language "typescript,python,go,csharp" >}}
{{% choosable language typescript %}}

```typescript
import * as dynatrace from "@lbrlabs/pulumi-dynatrace";
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import lbrlabs_pulumi_dynatrace as dynatrace
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"fmt"
	dynatrace "github.com/lbrlabs/pulumi-dynatrace/sdk/go/dynatrace"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
```

{{% /choosable %}}
{{% choosable language csharp %}}

```csharp
using Pulumi;
using Lbrlabs.PulumiPackage.Dynatrace;

```

{{% /choosable %}}

{{< /chooser >}}
