---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi-labs/pulumi-openai/v0.3.1/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/pulumi-labs/pulumi-openai/blob/v0.3.1/docs/_index.md
title: OpenAI
meta_desc: Use the Pulumi OpenAI provider to manage OpenAI agents, automations, vaults and vault credentials.
layout: package
---

The OpenAI provider for Pulumi lets you manage [OpenAI](https://platform.openai.com) resources as part of your Pulumi programs: agents, the automations that run them on a schedule, and the vaults that hold the credentials they use.

This provider is experimental. Its resources use OpenAI's beta Agents API.

## Installation

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}

npm:

```bash
npm install @pulumi-labs/openai
```

Yarn:

```bash
yarn add @pulumi-labs/openai
```

pnpm:

```bash
pnpm add @pulumi-labs/openai
```

Bun:

```bash
bun add @pulumi-labs/openai
```

{{% /choosable %}}
{{% choosable language python %}}

pip:

Add `pulumi-labs-openai` to `requirements.txt`, then run:

```bash
pip install -r requirements.txt
```

uv:

```bash
uv add pulumi-labs-openai
```

Poetry:

```bash
poetry add pulumi-labs-openai
```

{{% /choosable %}}
{{% choosable language go %}}

```bash
go get github.com/pulumi-labs/pulumi-openai/sdk/go/pulumi-openai
```

{{% /choosable %}}
{{% choosable language csharp %}}

```bash
dotnet add package Pulumi.Labs.Openai
```

{{% /choosable %}}
{{% choosable language java %}}

Replace `VERSION` with the latest version on [Maven Central](https://central.sonatype.com/artifact/com.pulumi.labs/openai).

Maven:

```xml
<dependency>
    <groupId>com.pulumi.labs</groupId>
    <artifactId>openai</artifactId>
    <version>VERSION</version>
</dependency>
```

Gradle:

```groovy
implementation 'com.pulumi.labs:openai:VERSION'
```

{{% /choosable %}}
{{% choosable language yaml %}}

Install the latest provider plugin from its GitHub releases:

```bash
pulumi plugin install resource openai --server github://api.github.com/pulumi-labs/pulumi-openai
```

{{% /choosable %}}
{{< /chooser >}}

## Example Usage

This program creates an agent that can search the web.

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}

```typescript
import * as openai from "@pulumi-labs/openai";

const supportAgent = new openai.Agent("supportAgent", {
    model: "gpt-5.5",
    name: "support-agent",
    instructions: "You answer customer questions about orders.",
    reasoning: { effort: openai.ReasoningEffort.Low },
    tools: [{ type: openai.ToolType.WebSearch }],
});

export const agentId = supportAgent.id;
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi
import pulumi_labs_openai as openai

support_agent = openai.Agent(
    "supportAgent",
    model="gpt-5.5",
    name="support-agent",
    instructions="You answer customer questions about orders.",
    reasoning={"effort": openai.ReasoningEffort.LOW},
    tools=[{"type": openai.ToolType.WEB_SEARCH}],
)

pulumi.export("agentId", support_agent.id)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
package main

import (
	openai "github.com/pulumi-labs/pulumi-openai/sdk/go/pulumi-openai"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		supportAgent, err := openai.NewAgent(ctx, "supportAgent", &openai.AgentArgs{
			Model:        pulumi.String("gpt-5.5"),
			Name:         pulumi.String("support-agent"),
			Instructions: pulumi.String("You answer customer questions about orders."),
			Reasoning: &openai.AgentReasoningArgs{
				Effort: openai.ReasoningEffortLow,
			},
			Tools: openai.AgentToolArray{
				&openai.AgentToolArgs{Type: openai.ToolTypeWebSearch},
			},
		})
		if err != nil {
			return err
		}

		ctx.Export("agentId", supportAgent.ID())
		return nil
	})
}
```

{{% /choosable %}}
{{% choosable language csharp %}}

```csharp
using System.Collections.Generic;
using Pulumi;
using OpenAI = Pulumi.Labs.Openai;

return await Deployment.RunAsync(() =>
{
    var supportAgent = new OpenAI.Agent("supportAgent", new()
    {
        Model = "gpt-5.5",
        Name = "support-agent",
        Instructions = "You answer customer questions about orders.",
        Reasoning = new OpenAI.Inputs.AgentReasoningArgs
        {
            Effort = OpenAI.ReasoningEffort.Low,
        },
        Tools =
        {
            new OpenAI.Inputs.AgentToolArgs { Type = OpenAI.ToolType.WebSearch },
        },
    });

    return new Dictionary<string, object?>
    {
        ["agentId"] = supportAgent.Id,
    };
});
```

{{% /choosable %}}
{{% choosable language java %}}

```java
import com.pulumi.Pulumi;
import com.pulumi.labs.openai.Agent;
import com.pulumi.labs.openai.AgentArgs;
import com.pulumi.labs.openai.enums.ReasoningEffort;
import com.pulumi.labs.openai.enums.ToolType;
import com.pulumi.labs.openai.inputs.AgentReasoningArgs;
import com.pulumi.labs.openai.inputs.AgentToolArgs;

public class App {
    public static void main(String[] args) {
        Pulumi.run(ctx -> {
            var supportAgent = new Agent("supportAgent", AgentArgs.builder()
                .model("gpt-5.5")
                .name("support-agent")
                .instructions("You answer customer questions about orders.")
                .reasoning(AgentReasoningArgs.builder()
                    .effort(ReasoningEffort.Low)
                    .build())
                .tools(AgentToolArgs.builder()
                    .type(ToolType.WebSearch)
                    .build())
                .build());

            ctx.export("agentId", supportAgent.id());
        });
    }
}
```

{{% /choosable %}}
{{% choosable language yaml %}}

```yaml
name: openai-example
runtime: yaml

resources:
  supportAgent:
    type: openai:Agent
    properties:
      model: gpt-5.5
      name: support-agent
      instructions: You answer customer questions about orders.
      reasoning:
        effort: low
      tools:
        - type: web_search

outputs:
  agentId: ${supportAgent.id}
```

{{% /choosable %}}
{{< /chooser >}}

Set your OpenAI API key before running `pulumi up`:

```bash
pulumi config set --secret openai:apiToken <your-api-key>
```

More programs, including an automation that runs the agent every weekday, are in the [`examples`](https://github.com/pulumi-labs/pulumi-openai/tree/main/examples) directory of the provider repository.

## Configuration

- `apiToken` (Required, Secret): The OpenAI API key. May also be set with the `OPENAI_API_KEY` environment variable.
- `baseUrl` (Optional): The base URL of the OpenAI API. Defaults to `https://api.openai.com/v1`. May also be set with the `OPENAI_BASE_URL` environment variable.

## Resources

- `Agent`: A reusable agent: a model with its instructions, tools and settings.
- `PeriodicAutomation`: Runs an agent repeatedly, on a cron expression.
- `OneTimeAutomation`: Runs an agent once, at a given time.
- `Vault`: A store of credentials that sessions use without the secrets appearing in an agent's configuration.
- `VaultCredential`: A credential in a vault: a token for an MCP server, or a secret for OpenAI-hosted environments.
