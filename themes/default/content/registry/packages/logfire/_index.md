---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pydantic/pulumi-logfire/v0.1.11/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/pydantic/pulumi-logfire/blob/v0.1.11/docs/_index.md
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

## Importing existing resources

Existing Logfire resources can be imported into Pulumi state with `pulumi import`. For resources that are easy to look up by name or slug, prefer the name-based import IDs:

| Resource | Recommended import ID | Example |
| --- | --- | --- |
| `logfire:Project` | `organization/project-name` | `pulumi import logfire:index/project:Project prod "acme/prod-logs"` |
| `logfire:Alert` | `project-name/alert-name` | `pulumi import logfire:index/alert:Alert errors "prod-logs/error-alert"` |
| `logfire:Dashboard` | `project-name/dashboard-slug` | `pulumi import logfire:index/dashboard:Dashboard overview "prod-logs/prod-overview"` |

UUID-based import IDs are also supported if you already have the backend IDs. The separators `/`, `,`, and `|` are accepted for multi-part import IDs.
