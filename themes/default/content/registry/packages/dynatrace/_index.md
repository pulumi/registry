---
# WARNING: this file was fetched from https://raw.githubusercontent.com/lbrlabs/pulumi-dynatrace/v0.29.0/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
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
import * as dynatrace from "@pulumiverse/dynatrace";
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumiverse_dynatrace as dynatrace
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"fmt"
	dynatrace "github.com/pulumiverse/pulumi-dynatrace/sdk/go/dynatrace"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
```

{{% /choosable %}}
{{% choosable language csharp %}}

```csharp
using Pulumi;
using Pulumiverse.Dynatrace;

```

{{% /choosable %}}

{{< /chooser >}}
