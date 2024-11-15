---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumiverse/pulumi-sentry/v0.0.9/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Sentry
meta_desc: Provides an overview of the Sentry Provider for Pulumi.
layout: overview
---

The Sentry provider for Pulumi can be used to provision Teams and Projects in [Sentry](https://sentry.io).

The Sentry provider must be configured with credentials to create and update resources in Sentry.

## Example

{{< chooser language "typescript,python,go,csharp" >}}
{{% choosable language typescript %}}

```typescript
import * as sentry from "@pulumiverse/sentry";
const project = new sentry.SentryProject("example", {
    name: "example-project",
    organization: "my-organization-id",
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumiverse_sentry as sentry

project = sentry.SentryProject("example",
    name="example-project"
    organization="my-organization-id",
)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"fmt"
	sentry "github.com/pulumiverse/pulumi-sentry/sdk/go/sentry"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		project, err := sentry.NewSentryProject(ctx, "example", &sentry.SentryProjectArgs{
            Name: pulumi.String("example-project"),
            Organization: pulumi.String("my-organization-id"),
		})

		if err != nil {
			return fmt.Errorf("error creating sentry project: %v", err)
		}

		return nil
	})
}
```

{{% /choosable %}}
{{% choosable language csharp %}}

```csharp
using Pulumi;
using Pulumiverse.sentry;

class sentryProject : Stack
{
    public sentryProject()
    {
        var project = new SentryProject("example", new SentryProjectArgs{
            Name: "example-project"
            Organization: "my-organization-id",
        });
    }
}
```

{{% /choosable %}}

{{< /chooser >}}
