---
# WARNING: this file was fetched from https://raw.githubusercontent.com/chronosphereio/pulumi-chronosphere/v0.9.9/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Chronosphere
meta_desc: Provides an overview of the Chronosphere Provider for Pulumi.
layout: overview
---

The Chronosphere Pulumi Provider allows [Pulumi](https://www.pulumi.com/) to manage supported [Chronosphere resources](https://docs.chronosphere.io/administer).

- [Documentation](https://docs.chronosphere.io/administer)

## Example

{{< chooser language "typescript,python,go,csharp" >}}
{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as chronosphere from "@pulumi-chronosphere/pulumi-chronosphere";

let testTeam = new chronosphere.Team("testTeam", {
  name: "Pulumi Test Team",
  slug: "pulumi-test-team",
});

let testNotifier = new chronosphere.EmailAlertNotifier("testNotifier", {
  name: "Pulumi Test Notifier",
  slug: "pulumi-test-notifier",
  to: "blackhole@example.com",
});

let testNotificationPolicy = new chronosphere.NotificationPolicy(
  "testNotificationPolicy",
  {
    name: "Pulumi Test Notification Policy",
    slug: "pulumi-test-notification-policy",
    teamId: testTeam.id,
    routes: [
      {
        severity: "warn",
        notifiers: [testNotifier.slug],
      },
    ],
  }
);

let testCollection = new chronosphere.Collection("testCollection", {
  name: "Pulumi Test Collection",
  slug: "pulumi-test-collection",
  teamId: testTeam.id,
  notificationPolicyId: testNotificationPolicy.id,
});

let testMonitor = new chronosphere.Monitor("testMonitor", {
  name: "Pulumi Test Monitor",
  slug: "pulumi-test-monitor",
  collectionId: testCollection.id,
  query: {
    prometheusExpr: 'chrono_test_metric{job="chrono_test_job"}',
  },
  seriesConditions: {
    conditions: [
      {
        op: "GT",
        value: 1000,
        severity: "critical",
        sustain: "10h",
      },
    ],
  },
});

export const teamSlug = testTeam.slug;
export const monitorSlug = testMonitor.slug;
```

{{% /choosable %}}
{{% choosable language python %}}

```python
team = chronosphere.Team("testTeam",
    name="Pulumi Test Team",
    slug="pulumi-test-team",
)

notifier = chronosphere.EmailAlertNotifier("testNotifier",
    name="Pulumi Test Notifier",
    slug="pulumi-test-notifier",
    to="blackhole@example.com",
)

notification_policy = chronosphere.NotificationPolicy("testNotificationPolicy",
    name="Pulumi Test Notification Policy",
    slug="pulumi-test-notification-policy",
    team_id=team.id,
    routes=[{
        "severity": "warn",
        "notifiers": [notifier.slug],
    }],
)

collection = chronosphere.Collection("testCollection",
    name="Pulumi Test Collection",
    slug="pulumi-test-collection",
    team_id=team.id,
    notification_policy_id=notification_policy.id,
)

monitor = chronosphere.Monitor("testMonitor",
    name="Pulumi Test Monitor",
    slug="pulumi-test-monitor",
    collection_id=collection.id,
    query={
        "prometheus_expr": 'chrono_test_metric{job="chrono_test_job"}',
    },
    series_conditions={
        "conditions": [{
            "op": "GT",
            "value": 1000,
            "severity": "critical",
            "sustain": "10h",
        }],
    },
)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
package main

import (
	"log"

	"github.com/chronosphereio/pulumi-chronosphere/sdk/go/chronosphere"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		team, err := chronosphere.NewTeam(ctx, "example-team", &chronosphere.TeamArgs{
			Name: pulumi.String("Pulumi Test Team"),
			Slug: pulumi.String("pulumi-test-team"),
		})
		if err != nil {
			log.Fatal(err)
		}

		testNotifier, err := chronosphere.NewEmailAlertNotifier(ctx, "testNotifier", &chronosphere.EmailAlertNotifierArgs{
			Name: pulumi.String("Pulumi Test Notifier"),
			Slug: pulumi.String("pulumi-test-notifier"),
			To:   pulumi.String("blackhole@example.com"),
		})
		if err != nil {
			log.Fatal(err)
		}

		testNotificationPolicy, err := chronosphere.NewNotificationPolicy(ctx, "testNotificationPolicy", &chronosphere.NotificationPolicyArgs{
			Name:   pulumi.String("Pulumi Test Notification Policy"),
			Slug:   pulumi.String("pulumi-test-notification-policy"),
			TeamId: team.ID(),
			Routes: chronosphere.NotificationPolicyRouteArray{
				&chronosphere.NotificationPolicyRouteArgs{
					Severity:  pulumi.String("warn"),
					Notifiers: pulumi.StringArray{testNotifier.Slug},
				},
			},
		})
		if err != nil {
			log.Fatal(err)
		}

		testCollection, err := chronosphere.NewCollection(ctx, "testCollection", &chronosphere.CollectionArgs{
			Name:                 pulumi.String("Pulumi Test Collection"),
			Slug:                 pulumi.String("pulumi-test-collection"),
			TeamId:               team.ID(),
			NotificationPolicyId: testNotificationPolicy.ID(),
		})
		if err != nil {
			log.Fatal(err)
		}

		testMonitor, err := chronosphere.NewMonitor(ctx, "testMonitor", &chronosphere.MonitorArgs{
			Name:         pulumi.String("Pulumi Test Monitor"),
			Slug:         pulumi.String("pulumi-test-monitor"),
			CollectionId: testCollection.ID(),
			Query: &chronosphere.MonitorQueryArgs{
				PrometheusExpr: pulumi.String(`chrono_test_metric{job="chrono_test_job"}`),
			},
			SeriesConditions: &chronosphere.MonitorSeriesConditionsArgs{
				Conditions: chronosphere.MonitorSeriesConditionsConditionArray{
					&chronosphere.MonitorSeriesConditionsConditionArgs{
						Op:       pulumi.String("GT"),
						Value:    pulumi.Float64(1000),
						Severity: pulumi.String("critical"),
						Sustain:  pulumi.String("10h"),
					},
				},
			},
		})
		if err != nil {
			log.Fatal(err)
		}

		ctx.Export("teamSlug", team.Slug)
		ctx.Export("monitorSlug", testMonitor.Slug)

		return nil
	})
}
```

{{% /choosable %}}
{{% choosable language csharp %}}

```csharp
using System.Collections.Generic;
using System.Security.Cryptography.X509Certificates;
using Chronosphere.Pulumi.Chronosphere;
using Chronosphere.Pulumi.Chronosphere.Inputs;
using Pulumi;

return await Deployment.RunAsync(() =>
{
    var team = new Team("testTeam", new TeamArgs
    {
        Name = "Pulumi Test Team",
        Slug = "pulumi-test-team",
    });

    var notifier = new EmailAlertNotifier("testNotifier", new EmailAlertNotifierArgs
    {
        Name = "Pulumi Test Notifier",
        Slug = "pulumi-test-notifier",
        To = "blackhole@example.com",
    });

    var notificationPolicy = new NotificationPolicy("testNotificationPolicy", new NotificationPolicyArgs
    {
        Name = "Pulumi Test Notification Policy",
        Slug = "pulumi-test-notification-policy",
        TeamId = team.Id,
        Routes = new[]
        {
            new NotificationPolicyRouteArgs
            {
                Severity = "warn",
                Notifiers = new[] { notifier.Slug },
            },
        },
    });

    var collection = new Collection("testCollection", new CollectionArgs
    {
        Name = "Pulumi Test Collection",
        Slug = "pulumi-test-collection",
        TeamId = team.Id,
        NotificationPolicyId = notificationPolicy.Id,
    });

    var monitor = new Chronosphere.Pulumi.Chronosphere.Monitor("testMonitor", new MonitorArgs
    {
        Name = "Pulumi Test Monitor",
        Slug = "pulumi-test-monitor",
        CollectionId = collection.Id,
        Query = new MonitorQueryArgs
        {
            PrometheusExpr = "chrono_test_metric{job=\"chrono_test_job\"}",
        },
        SeriesConditions = new MonitorSeriesConditionsArgs
        {
            Conditions = new[]
            {
                new MonitorSeriesConditionsConditionArgs
                {
                    Op = "GT",
                    Value = 1000,
                    Severity = "critical",
                    Sustain = "10h",
                },
            },
        },
    });

    return new Dictionary<string, object?>
    {
        ["monitorSlug"] = monitor.Slug,
    };
});
```

{{% /choosable %}}

{{< /chooser >}}
