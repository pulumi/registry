---
title: Buildkite
meta_desc: Provides an overview of the Buildkite Provider for Pulumi.
layout: package
---

The Buildkite provider for Pulumi can be used to provision resources for [Buildkite](http://buildkite.com).

## Example

{{< chooser language "javascript,typescript,python,csharp" >}}

{{% choosable language javascript %}}

```javascript
const buildkite = require("@pulumiverse/buildkite")

const agentToken = new buildkite.agent.AgentToken('token', {
    description: 'agent token',
})
```

{{% /choosable %}} {{% choosable language typescript %}}

```typescript
import * as buildkite from "@pulumiverse/buildkite";

const agentToken = new buildkite.agent.AgentToken('token', {
    description: 'agent token',
})
```

{{% /choosable %}} {{% choosable language python %}}

```python
import pulumiverse_buildkite as buildkite

agent_token = buildkite.agent.AgentToken(
    description="agent token"
)
```

{{% /choosable %}} {{% choosable language csharp %}}

```csharp
using System.Collections.Generic;
using System.Threading.Tasks;
using Pulumi;
using Pulumiverse.Buildkite.Agent;
using Pulumiverse.Buildkite.Agent.Inputs;

class Program
{
    static Task Main() =>
        Deployment.Run(() => {
            var agentToken = new AgentToken("vm", new AgentTokenArgs
            {
                Description="agent token"
            });
        });
}
```

{{% /choosable %}}

{{< /chooser >}}
