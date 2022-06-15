---
title: Slack
meta_desc: Provides an overview of the Slack Provider for Pulumi.
layout: overview
---

The Slack provider for Pulumi can be used to provision resources available in [Slack](https://www.slack.com/).
The Slack provider must be configured with a token to manage resources in a Slack Workspace.

## Example

{{< chooser language "typescript,python,go,csharp" >}}
{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as slack from "@pulumi/slack";

// Create a new Slack channel
const channel = new slack.Conversation("acc-test", {
  topic: "my topic",
  isPrivate: false,
});

// Get the id of the new channel as an output
export const channelId = channel.id;
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_slack as slack

channel = slack.Conversation("test",
    topic: "test topic",
    isPrivate: false,
)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	slack "github.com/pulumi/pulumi-slack/sdk/go/slack"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		channel, err := slack.Conversation(ctx, "test", &slack.ConversationArgs{
            Topic: pulumi.String("my topic"),
            IsPrivate: pulumi.Bool(false),
        })
		if err != nil {
			return err
		}

		ctx.Export("channel", channel.Name)

		return nil
	})
}
```

{{% /choosable %}}
{{% choosable language csharp %}}

```csharp
using Pulumi;
using Pulumi.Slack;

class SlackChannel : Channel
{
    public SlackChannel()
    {
        var channel = new Conversation("test", new ConversationArgs{
            topic = "my topic",
            isPrivate = false,
        });
    }
}
```

{{% /choosable %}}

{{< /chooser >}}