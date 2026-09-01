---
# WARNING: this file was fetched from https://raw.githubusercontent.com/matthiashamacher/pulumi-ohdear/v0.1.10/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/matthiashamacher/pulumi-ohdear/blob/v0.1.10/docs/_index.md
title: Oh Dear
meta_desc: Manage Oh Dear website monitoring — monitors, status pages, tags, and notification destinations — as Pulumi resources.
layout: package
---

The Oh Dear provider for Pulumi lets you manage [Oh Dear](https://ohdear.app)
website monitoring in code: monitors, status pages, tags and tag groups, and
notification destinations.

It talks to the [Oh Dear API](https://ohdear.app/docs/api) directly and is
authenticated with a single API token sent as a Bearer token on every request.

## Example

{{< chooser language "typescript" >}}

{{% choosable language typescript %}}

```typescript
import * as ohdear from "@matthiashamacher/ohdear";

// Monitor a website with the uptime and certificate checks.
const site = new ohdear.Monitor("docs", {
    url: "https://example.com",
    teamId: 1,
    type: "http",
    checks: ["uptime", "certificate_health"],
    tags: ["production"],
});

// Group related tags together.
const envs = new ohdear.TagGroup("environments", {
    teamId: 1,
    label: "Environments",
    tags: ["production", "staging"],
});

// A public status page that lists the monitor.
const status = new ohdear.StatusPage("public", {
    teamId: 1,
    title: "Example Status",
    monitors: [{ id: site.monitorId, clickable: true }],
});

// Email alerts for the monitor.
const alerts = new ohdear.NotificationDestination("email", {
    level: "monitor",
    ownerId: site.monitorId,
    channel: "mail",
    destination: { mail: "alerts@example.com" },
    notificationTypes: [
        "HttpUptimeCheckFailedNotification",
        "HttpUptimeCheckRecoveredNotification",
    ],
});
```

{{% /choosable %}}

{{< /chooser >}}

## Resources

| Resource | Purpose |
| --- | --- |
| `Monitor` | A monitored URL or host (`http`, `ping`, `tcp`, `ai`) and its checks. |
| `Tag` | A tag used to organize monitors. Create-only upstream. |
| `TagGroup` | A named group of tags, with wildcard matching. |
| `StatusPage` | A public status page and the monitors shown on it. |
| `StatusPageUpdateTemplate` | A reusable status-page update message. |
| `NotificationDestination` | Where alerts are delivered, scoped to a team, monitor, tag or tag group. |

See the **Installation & Configuration** tab for setup and the API Docs tab for
the full input and output reference of each resource.
