---
# WARNING: this file was fetched from https://raw.githubusercontent.com/incident-io/pulumi-incident/v0.1.2/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/incident-io/pulumi-incident/blob/v0.1.2/docs/_index.md
title: incident.io
meta_desc: Manage incident.io configuration as code with Pulumi, in TypeScript, Python or Go.
layout: package
---

The incident.io provider lets you manage your incident.io configuration from a Pulumi program. Escalation paths, schedules, alert routes, catalog, custom fields and workflows become code you review, version and roll back like the rest of your infrastructure.

It is generated from the [official incident.io Terraform provider](https://github.com/incident-io/terraform-provider-incident), so anything you can manage with Terraform you can manage here.

## Installation

The incident.io provider is available in three languages:

```bash
npm install @incident-io/pulumi
pip install pulumi-incident
go get github.com/incident-io/pulumi-incident/sdk/go/incident
```

## Example Usage

{{< chooser language "typescript,python,go" >}}
{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as incident from "@incident-io/pulumi";

const trivial = new incident.Severity("trivial", {
    name: "Trivial",
    description: "Issues causing no impact. No immediate response is required.",
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi
import pulumi_incident as incident

trivial = incident.Severity("trivial",
    name="Trivial",
    description="Issues causing no impact. No immediate response is required.")
```

{{% /choosable %}}
{{% choosable language go %}}

```go
package main

import (
	"github.com/incident-io/pulumi-incident/sdk/go/incident"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := incident.NewSeverity(ctx, "trivial", &incident.SeverityArgs{
			Name:        pulumi.String("Trivial"),
			Description: pulumi.String("Issues causing no impact. No immediate response is required."),
		})
		return err
	})
}
```

{{% /choosable %}}
{{< /chooser >}}

## Configuration

The provider needs an API key, created in [Settings → API keys](https://app.incident.io/settings/api-keys):

```bash
pulumi config set --secret incident:apiKey inc_...
```

`INCIDENT_API_KEY` works as an environment variable alternative. On a dedicated or self-hosted deployment, also set `incident:endpoint`. See [Installation & Configuration](/registry/packages/incident/installation-configuration/) for the full reference.

## What you can manage

- **On-call**: escalation paths, schedules, and the sync rules and targets that keep rotations in step with an external source of truth.
- **Alerts**: alert sources, alert routes, and the attributes you route on.
- **Catalog**: catalog types, their attributes, and entries, for modelling your services and teams and driving routing from that model.
- **Incident configuration**: severities, statuses, incident roles, custom fields and their options.
- **Automation**: workflows and maintenance windows.

Resources with `Beta` in their name track features still changing shape in the product, and their inputs may change in a minor release.

## Importing existing configuration

Every resource supports `pulumi import`, so you can bring an account you already configured by hand under management a piece at a time:

```bash
pulumi import incident:index/severity:Severity trivial 01ABCDEF...
```

By default, importing a resource also claims it as managed by code, which stops people editing it in the incident.io dashboard. That claim is a write to your account, so set `markImportedResourcesAsManaged` to `false` if you would rather imports left it untouched.
