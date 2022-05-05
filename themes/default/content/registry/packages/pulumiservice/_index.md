---
title: Pulumi Service
meta_desc: Provides an overview of the Pulumi Service provider for Pulumi.
layout: overview
---

The Pulumi Service provider for Pulumi can be used to provision resources available in the [Pulumi Service](https://www.scaleway.com/).
The Pulumi Service provider must be configured with credentials to deploy and update resources in Pulumi Service.

## Example

{{< chooser language "typescript,python,go,csharp" >}}
{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as pulumiservice from "@pulumi/pulumiservice";
const webhook = new pulumiservice.Webhook("example-webhook", {
    active: true,
    displayName: "webhook example",
    organizationName: "example",
    payloadUrl: "https://example.com/webhook",
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi
import pulumi_service
webhook = pulumi_service.Webhook("example-webhook",
    active: True,
    display_name: "webhook example",
    organization_name: "example",
    payload_url: "https://example.com/webhook",
)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"fmt"
	"github.com/pulumi/pulumi-pulumiservice/sdk/go/pulumiservice"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		webhook, err := pulumiservice.NewWebhook(ctx, "example-webhook", &pulumiservice.WebhookArgs{
			Active:           pulumi.Bool(true),
			DisplayName:      pulumi.String("example webhook"),
			OrganizationName: pulumi.String("example"),
			PayloadURL:       pulumi.String("https://example.com/webhook"),
		}, nil)
		if err != nil {
			return fmt.Errorf("error creating webhook: %v", err)
		}
		return nil
	})
}
```

{{% /choosable %}}
{{% choosable language csharp %}}

```csharp
using Pulumi;
using Pulumi.PulumiService;
class PulumiServiceWebhook: Stack
{
    public PulumiServiceWebhook()
    {
        var webhook = new Webhook("example-webhook", new WebhookArgs{
            Active = true,
            DisplayName = "example webhook",
            OrganizationName = "example",
            PayloadUrl = "https://example.com/webhook"
        })
    }
}
```

{{% /choosable %}}
{{< /chooser >}}
