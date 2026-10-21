---
title: Pulumi Cloud (Pulumi Service provider)
meta_desc: Provides an overview of the Pulumi Service provider for Pulumi.
layout: package
---

The Pulumi Service provider for Pulumi can be used to provision resources available in [Pulumi Cloud](https://app.pulumi.com/).

## Installation

{{< chooser language "typescript,python,go,csharp,java,yaml,hcl" >}}
{{% choosable language typescript %}}

```bash
npm install @pulumi/pulumiservice
```

{{% /choosable %}}
{{% choosable language python %}}

```bash
pip install pulumi-pulumiservice
```

{{% /choosable %}}
{{% choosable language go %}}

```bash
go get github.com/pulumi/pulumi-pulumiservice/sdk/go/pulumiservice
```

{{% /choosable %}}
{{% choosable language csharp %}}

```bash
dotnet add package Pulumi.PulumiService
```

{{% /choosable %}}
{{% choosable language java %}}

Maven:

```xml
<dependency>
    <groupId>com.pulumi</groupId>
    <artifactId>pulumiservice</artifactId>
    <version>1.5.0</version>
</dependency>
```

Gradle:

```groovy
implementation 'com.pulumi:pulumiservice:1.5.0'
```

{{% /choosable %}}
{{% choosable language yaml %}}

```bash
pulumi package add pulumiservice
```

{{% /choosable %}}
{{% choosable language hcl %}}

Declare the provider in your program, then run `pulumi install`:

```hcl
terraform {
  required_providers {
    pulumiservice = {
      source  = "pulumi/pulumiservice"
      version = "1.5.0"
    }
  }
}
```

```bash
pulumi install
```

{{% /choosable %}}
{{< /chooser >}}

## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml,hcl" >}}
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
import pulumi_pulumiservice
webhook = pulumi_pulumiservice.Webhook("example-webhook",
    active=True,
    display_name="webhook example",
    organization_name="example",
    payload_url="https://example.com/webhook",
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
			DisplayName:      pulumi.String("webhook example"),
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
            DisplayName = "webhook example",
            OrganizationName = "example",
            PayloadUrl = "https://example.com/webhook"
        });
    }
}
```

{{% /choosable %}}
{{% choosable language java %}}

```java
import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.pulumiservice.Webhook;
import com.pulumi.pulumiservice.WebhookArgs;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    private static void stack(Context ctx) {
        var webhook = new Webhook("example-webhook", WebhookArgs.builder()
            .active(true)
            .displayName("webhook example")
            .organizationName("example")
            .payloadUrl("https://example.com/webhook")
            .build());
    }
}
```

{{% /choosable %}}
{{% choosable language yaml %}}

```yaml
resources:
  webhook:
    type: pulumiservice:index:Webhook
    properties:
      active: true
      displayName: webhook example
      organizationName: example
      payloadUrl: https://example.com/webhook
```

{{% /choosable %}}
{{% choosable language hcl %}}

```hcl
terraform {
  required_providers {
    pulumiservice = {
      source = "pulumi/pulumiservice"
    }
  }
}

resource "pulumiservice_webhook" "webhook" {
  active            = true
  display_name      = "webhook example"
  organization_name = "example"
  payload_url       = "https://example.com/webhook"
}
```

{{% /choosable %}}

{{< /chooser >}}

## Configuration

To provision resources with the Pulumi Service provider, you need Pulumi Cloud credentials. The provider looks for an [access token](/docs/administration/concepts/access-tokens/) in this order:

1. The `pulumiservice:accessToken` configuration option.
2. The `PULUMI_ACCESS_TOKEN` environment variable.
3. The account you're logged in to with `pulumi login`.

If you've already run `pulumi login`, you don't need to configure anything else.

{{% notes "info" %}}
The provider gets your login's API URL from the Pulumi CLI, which passes it to providers starting with CLI v3.247.0. If you're logged in to a self-hosted Pulumi Cloud with an older CLI, the provider calls `https://api.pulumi.com` instead, so also set `PULUMI_BACKEND_URL` or the `apiUrl` option, or upgrade the CLI.
{{% /notes %}}

Use `pulumi config set pulumiservice:<option>` or pass options to the provider constructor.

| Name | Required | Secret | Description |
|---|---|---|---|
| `accessToken` | No | Yes | [Pulumi Cloud access token](/docs/administration/concepts/access-tokens/). Defaults to `PULUMI_ACCESS_TOKEN`, then to your `pulumi login` credentials. |
| `apiUrl` | No | No | Overrides the default [Pulumi Cloud API URL](/docs/pulumi-cloud/cloud-rest-api/) for [self-hosted customers](/docs/pulumi-cloud/self-hosted/). Defaults to `PULUMI_BACKEND_URL`, then `PULUMI_API`, then `https://api.pulumi.com`. |

### Example

```bash
# Not needed if you've run `pulumi login` or set PULUMI_ACCESS_TOKEN.
pulumi config set --secret pulumiservice:accessToken <PULUMI_ACCESS_TOKEN>

# Only needed for a self-hosted Pulumi Cloud, and only if you haven't set
# PULUMI_BACKEND_URL or logged in to it with Pulumi CLI v3.247.0 or later.
pulumi config set pulumiservice:apiUrl <PULUMI_API_URL>
```
