---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pydantic/pulumi-logfire/v0.1.8/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/pydantic/pulumi-logfire/blob/v0.1.8/docs/_index.md
title: Logfire
meta_desc: Use the Pulumi Logfire provider to manage projects, alerts, channels, dashboards, and API tokens.
layout: package
---

# Logfire Provider

The Logfire provider for Pulumi lets you manage Logfire resources as code, including:

- Projects
- Channels
- Alerts
- Dashboards
- Read tokens
- Write tokens
- Organizations

## Example

### TypeScript

```typescript
import * as logfire from "@pydantic/pulumi-logfire";

const project = new logfire.Project("exampleProject");
```

### Python

```python
import pulumi_logfire as logfire

project = logfire.Project("exampleProject")
```

### Go

```go
package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pydantic/pulumi-logfire/sdk/go/logfire"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := logfire.NewProject(ctx, "exampleProject", nil)
		return err
	})
}
```
