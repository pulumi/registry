---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-newrelic/v5.56.0/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/pulumi/pulumi-newrelic/blob/v5.56.0/docs/_index.md
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Newrelic Provider
meta_desc: Provides an overview on how to configure the Pulumi Newrelic provider.
layout: package
---

## Installation

The Newrelic provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/newrelic`](https://www.npmjs.com/package/@pulumi/newrelic)
* Python: [`pulumi-newrelic`](https://pypi.org/project/pulumi-newrelic/)
* Go: [`github.com/pulumi/pulumi-newrelic/sdk/v5/go/newrelic`](https://github.com/pulumi/pulumi-newrelic)
* .NET: [`Pulumi.Newrelic`](https://www.nuget.org/packages/Pulumi.Newrelic)
* Java: [`com.pulumi/newrelic`](https://central.sonatype.com/artifact/com.pulumi/newrelic)

## Overview

[New Relic](https://newrelic.com/) offers tools that help you fix problems
quickly, maintain complex systems, improve your code, and accelerate your
digital transformation.

Use the navigation to the left to read about the available resources.
## Configuration Reference

The following configuration inputs are supported.

|        Argument        | Required? |                                                                                            Description                                                                                             |
|------------------------|-----------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `accountId`           | Required  | Your New Relic account ID. The `NEW_RELIC_ACCOUNT_ID` environment variable can also be used.                                                                                                       |
| `apiKey`              | Required  | Your New Relic Personal API key (usually prefixed with `NRAK`). The `NEW_RELIC_API_KEY` environment variable can also be used.                                                                     |
| `region`               | Optional  | The region for the data center for which your New Relic account is configured. The `NEW_RELIC_REGION` environment variable can also be used. Valid values are `US` or `EU`. Default value is `US`. |
| `insecureSkipVerify` | Optional  | Trust self-signed SSL certificates. If omitted, the `NEW_RELIC_API_SKIP_VERIFY` environment variable is used.                                                                                      |
| `insightsInsertKey`  | Optional  | Your Insights insert key used when inserting Insights events via the `newrelic.insights.Event` resource. Can also use `NEW_RELIC_INSIGHTS_INSERT_KEY` environment variable.                        |
| `cacertFile`          | Optional  | A path to a PEM-encoded certificate authority used to verify the remote agent's certificate. The `NEW_RELIC_API_CACERT` environment variable can also be used.                                     |
## Authentication Requirements

This provider is in the midst of migrating away from our older REST based APIs
to a newer GraphQL based API that we lovingly call NerdGraph. During this
transition, the provider will be using different endpoints depending on which
resource is in use. Below is a table that reflects the current state of the
resources compared to which endpoint is in use.
### Resources

|                      Resource                       |        Endpoint         |    Authentication     |
|-----------------------------------------------------|-------------------------|-----------------------|
| `newrelic.AccountManagement`                       | NerdGraph               | `apiKey`             |
| `newrelic.AlertChannel`                            | RESTv2                  | `apiKey`             |
| `newrelic.AlertCondition`                          | RESTv2                  | `apiKey`             |
| `newrelic.AlertMutingRule`                        | NerdGraph               | `apiKey`             |
| `newrelic.AlertPolicy`                             | NerdGraph               | `apiKey`             |
| `newrelic.AlertPolicyChannel`                     | RESTv2                  | `apiKey`             |
| `newrelic.ApiAccessKey`                           | NerdGraph               | `apiKey`             |
| `newrelic.plugins.ApplicationSettings`                     | RESTv2                  | `apiKey`             |
| `newrelic.BrowserApplication`                      | NerdGraph               | `apiKey`             |
| `newrelic.cloud.AwsGovcloudIntegrations`          | NerdGraph               | `apiKey`             |
| `newrelic.cloud.AwsGovcloudLinkAccount`          | NerdGraph               | `apiKey`             |
| `newrelic.cloud.AwsIntegrations`                   | NerdGraph               | `apiKey`             |
| `newrelic.cloud.AwsLinkAccount`                   | NerdGraph               | `apiKey`             |
| `newrelic.cloud.AzureIntegrations`                 | NerdGraph               | `apiKey`             |
| `newrelic.cloud.AzureLinkAccount`                 | NerdGraph               | `apiKey`             |
| `newrelic.cloud.GcpIntegrations`                   | NerdGraph               | `apiKey`             |
| `newrelic.cloud.GcpLinkAccount`                   | NerdGraph               | `apiKey`             |
| `newrelic.DataPartitionRule`                      | NerdGraph               | `apiKey`             |
| `newrelic.EntityTags`                              | NerdGraph               | `apiKey`             |
| `newrelic.EventsToMetricsRule`                   | NerdGraph               | `apiKey`             |
| `newrelic.Group`                                    | NerdGraph               | `apiKey`             |
| `newrelic.InfraAlertCondition`                    | Infrastructure REST API | `apiKey`             |
| `newrelic.insights.Event`                           | Insights API            | `insightsInsertKey` |
| `newrelic.KeyTransaction`                          | NerdGraph               | `apiKey`             |
| `newrelic.LogParsingRule`                         | NerdGraph               | `apiKey`             |
| `newrelic.NotificationChannel`                     | NerdGraph               | `apiKey`             |
| `newrelic.NotificationDestination`                 | NerdGraph               | `apiKey`             |
| `newrelic.NrqlAlertCondition`                     | NerdGraph               | `apiKey`             |
| `newrelic.NrqlDropRule`                           | NerdGraph               | `apiKey`             |
| `newrelic.ObfuscationExpression`                   | NerdGraph               | `apiKey`             |
| `newrelic.ObfuscationRule`                         | NerdGraph               | `apiKey`             |
| `newrelic.OneDashboard`                            | NerdGraph               | `apiKey`             |
| `newrelic.OneDashboardJson`                       | NerdGraph               | `apiKey`             |
| `newrelic.OneDashboardRaw`                        | NerdGraph               | `apiKey`             |
| `newrelic.ServiceLevel`                            | NerdGraph               | `apiKey`             |
| `newrelic.synthetics.AlertCondition`               | RESTv2                  | `apiKey`             |
| `newrelic.synthetics.BrokenLinksMonitor`          | NerdGraph               | `apiKey`             |
| `newrelic.synthetics.CertCheckMonitor`            | NerdGraph               | `apiKey`             |
| `newrelic.synthetics.Monitor`                       | NerdGraph               | `apiKey`             |
| `newrelic.synthetics.MultiLocationAlertCondition` | RESTv2                  | `apiKey`             |
| `newrelic.synthetics.PrivateLocation`              | NerdGraph               | `apiKey`             |
| `newrelic.synthetics.ScriptMonitor`                | NerdGraph               | `apiKey`             |
| `newrelic.synthetics.SecureCredential`             | NerdGraph               | `apiKey`             |
| `newrelic.synthetics.StepMonitor`                  | NerdGraph               | `apiKey`             |
| `newrelic.User`                                     | NerdGraph               | `apiKey`             |
| `newrelic.Workflow`                                 | NerdGraph               | `apiKey`             |
| `newrelic.plugins.Workload`                                 | NerdGraph               | `apiKey`             |
### Functions

|               Function               | Endpoint  | Authentication |
|-----------------------------------------|-----------|----------------|
| `newrelic.getAccount`                      | NerdGraph | `apiKey`      |
| `newrelic.AlertChannel`                | RESTv2    | `apiKey`      |
| `newrelic.AlertPolicy`                 | NerdGraph | `apiKey`      |
| `newrelic.getApplication`                  | RESTv2    | `apiKey`      |
| `newrelic.getCloudAccount`                | NerdGraph | `apiKey`      |
| `newrelic.getEntity`                       | NerdGraph | `apiKey`      |
| `newrelic.KeyTransaction`              | NerdGraph | `apiKey`      |
| `newrelic.NotificationDestination`     | NerdGraph | `apiKey`      |
| `newrelic.ObfuscationExpression`       | NerdGraph | `apiKey`      |
| `newrelic.synthetics.PrivateLocation`  | NerdGraph | `apiKey`      |
| `newrelic.synthetics.SecureCredential` | NerdGraph | `apiKey`      |
| `newrelic.getTestGrokPattern`            | NerdGraph | `apiKey`      |
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    newrelic:accountId:
        value: 12345
    newrelic:apiKey:
        value: 12345
    newrelic:region:
        value: US

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as newrelic from "@pulumi/newrelic";

// Read an APM application resource
const foo = newrelic.getEntity({
    name: "Your App Name",
    domain: "APM",
    type: "APPLICATION",
});
// Create an alert policy
const alert = new newrelic.AlertPolicy("alert", {name: "Your Concise Alert Name"});
// Add a condition
const fooNrqlAlertCondition = new newrelic.NrqlAlertCondition("foo", {
    policyId: alert.id,
    type: "static",
    name: "foo",
    description: "Alert when transactions are taking too long",
    runbookUrl: "https://www.example.com",
    enabled: true,
    violationTimeLimitSeconds: 3600,
    nrql: {
        query: foo.then(foo => `SELECT average(duration) FROM Transaction where appName = '${foo.name}'`),
    },
    critical: {
        operator: "above",
        threshold: 5.5,
        thresholdDuration: 300,
        thresholdOccurrences: "ALL",
    },
});
// Add a notification channel
const email = new newrelic.AlertChannel("email", {
    name: "email",
    type: "email",
    config: {
        recipients: "username@example.com",
        includeJsonAttachment: "1",
    },
});
// Link the channel to the policy
const alertEmail = new newrelic.AlertPolicyChannel("alert_email", {
    policyId: alert.id,
    channelIds: [email.id],
});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    newrelic:accountId:
        value: 12345
    newrelic:apiKey:
        value: 12345
    newrelic:region:
        value: US

```
```python
import pulumi
import pulumi_newrelic as newrelic

# Read an APM application resource
foo = newrelic.get_entity(name="Your App Name",
    domain="APM",
    type="APPLICATION")
# Create an alert policy
alert = newrelic.AlertPolicy("alert", name="Your Concise Alert Name")
# Add a condition
foo_nrql_alert_condition = newrelic.NrqlAlertCondition("foo",
    policy_id=alert.id,
    type="static",
    name="foo",
    description="Alert when transactions are taking too long",
    runbook_url="https://www.example.com",
    enabled=True,
    violation_time_limit_seconds=3600,
    nrql={
        "query": f"SELECT average(duration) FROM Transaction where appName = '{foo.name}'",
    },
    critical={
        "operator": "above",
        "threshold": 5.5,
        "threshold_duration": 300,
        "threshold_occurrences": "ALL",
    })
# Add a notification channel
email = newrelic.AlertChannel("email",
    name="email",
    type="email",
    config={
        "recipients": "username@example.com",
        "include_json_attachment": "1",
    })
# Link the channel to the policy
alert_email = newrelic.AlertPolicyChannel("alert_email",
    policy_id=alert.id,
    channel_ids=[email.id])
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    newrelic:accountId:
        value: 12345
    newrelic:apiKey:
        value: 12345
    newrelic:region:
        value: US

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using NewRelic = Pulumi.NewRelic;

return await Deployment.RunAsync(() =>
{
    // Read an APM application resource
    var foo = NewRelic.GetEntity.Invoke(new()
    {
        Name = "Your App Name",
        Domain = "APM",
        Type = "APPLICATION",
    });

    // Create an alert policy
    var alert = new NewRelic.AlertPolicy("alert", new()
    {
        Name = "Your Concise Alert Name",
    });

    // Add a condition
    var fooNrqlAlertCondition = new NewRelic.NrqlAlertCondition("foo", new()
    {
        PolicyId = alert.Id,
        Type = "static",
        Name = "foo",
        Description = "Alert when transactions are taking too long",
        RunbookUrl = "https://www.example.com",
        Enabled = true,
        ViolationTimeLimitSeconds = 3600,
        Nrql = new NewRelic.Inputs.NrqlAlertConditionNrqlArgs
        {
            Query = $"SELECT average(duration) FROM Transaction where appName = '{foo.Apply(getEntityResult => getEntityResult.Name)}'",
        },
        Critical = new NewRelic.Inputs.NrqlAlertConditionCriticalArgs
        {
            Operator = "above",
            Threshold = 5.5,
            ThresholdDuration = 300,
            ThresholdOccurrences = "ALL",
        },
    });

    // Add a notification channel
    var email = new NewRelic.AlertChannel("email", new()
    {
        Name = "email",
        Type = "email",
        Config = new NewRelic.Inputs.AlertChannelConfigArgs
        {
            Recipients = "username@example.com",
            IncludeJsonAttachment = "1",
        },
    });

    // Link the channel to the policy
    var alertEmail = new NewRelic.AlertPolicyChannel("alert_email", new()
    {
        PolicyId = alert.Id,
        ChannelIds = new[]
        {
            email.Id,
        },
    });

});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    newrelic:accountId:
        value: 12345
    newrelic:apiKey:
        value: 12345
    newrelic:region:
        value: US

```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-newrelic/sdk/v5/go/newrelic"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Read an APM application resource
		foo, err := newrelic.GetEntity(ctx, &newrelic.GetEntityArgs{
			Name:   "Your App Name",
			Domain: pulumi.StringRef("APM"),
			Type:   pulumi.StringRef("APPLICATION"),
		}, nil)
		if err != nil {
			return err
		}
		// Create an alert policy
		alert, err := newrelic.NewAlertPolicy(ctx, "alert", &newrelic.AlertPolicyArgs{
			Name: pulumi.String("Your Concise Alert Name"),
		})
		if err != nil {
			return err
		}
		// Add a condition
		_, err = newrelic.NewNrqlAlertCondition(ctx, "foo", &newrelic.NrqlAlertConditionArgs{
			PolicyId:                  alert.ID(),
			Type:                      pulumi.String("static"),
			Name:                      pulumi.String("foo"),
			Description:               pulumi.String("Alert when transactions are taking too long"),
			RunbookUrl:                pulumi.String("https://www.example.com"),
			Enabled:                   pulumi.Bool(true),
			ViolationTimeLimitSeconds: pulumi.Int(3600),
			Nrql: &newrelic.NrqlAlertConditionNrqlArgs{
				Query: pulumi.Sprintf("SELECT average(duration) FROM Transaction where appName = '%v'", foo.Name),
			},
			Critical: &newrelic.NrqlAlertConditionCriticalArgs{
				Operator:             pulumi.String("above"),
				Threshold:            pulumi.Float64(5.5),
				ThresholdDuration:    pulumi.Int(300),
				ThresholdOccurrences: pulumi.String("ALL"),
			},
		})
		if err != nil {
			return err
		}
		// Add a notification channel
		email, err := newrelic.NewAlertChannel(ctx, "email", &newrelic.AlertChannelArgs{
			Name: pulumi.String("email"),
			Type: pulumi.String("email"),
			Config: &newrelic.AlertChannelConfigArgs{
				Recipients:            pulumi.String("username@example.com"),
				IncludeJsonAttachment: pulumi.String("1"),
			},
		})
		if err != nil {
			return err
		}
		// Link the channel to the policy
		_, err = newrelic.NewAlertPolicyChannel(ctx, "alert_email", &newrelic.AlertPolicyChannelArgs{
			PolicyId: alert.ID(),
			ChannelIds: pulumi.StringArray{
				email.ID(),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
```
{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
config:
    newrelic:accountId:
        value: 12345
    newrelic:apiKey:
        value: 12345
    newrelic:region:
        value: US

```
```yaml
resources:
  # Create an alert policy
  alert:
    type: newrelic:AlertPolicy
    properties:
      name: Your Concise Alert Name
  # Add a condition
  fooNrqlAlertCondition:
    type: newrelic:NrqlAlertCondition
    name: foo
    properties:
      policyId: ${alert.id}
      type: static
      name: foo
      description: Alert when transactions are taking too long
      runbookUrl: https://www.example.com
      enabled: true
      violationTimeLimitSeconds: 3600
      nrql:
        query: SELECT average(duration) FROM Transaction where appName = '${foo.name}'
      critical:
        operator: above
        threshold: 5.5
        thresholdDuration: 300
        thresholdOccurrences: ALL
  # Add a notification channel
  email:
    type: newrelic:AlertChannel
    properties:
      name: email
      type: email
      config:
        recipients: username@example.com
        includeJsonAttachment: '1'
  # Link the channel to the policy
  alertEmail:
    type: newrelic:AlertPolicyChannel
    name: alert_email
    properties:
      policyId: ${alert.id}
      channelIds:
        - ${email.id}
variables:
  # Read an APM application resource
  foo:
    fn::invoke:
      function: newrelic:getEntity
      arguments:
        name: Your App Name
        domain: APM
        type: APPLICATION
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    newrelic:accountId:
        value: 12345
    newrelic:apiKey:
        value: 12345
    newrelic:region:
        value: US

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.newrelic.NewrelicFunctions;
import com.pulumi.newrelic.inputs.GetEntityArgs;
import com.pulumi.newrelic.AlertPolicy;
import com.pulumi.newrelic.AlertPolicyArgs;
import com.pulumi.newrelic.NrqlAlertCondition;
import com.pulumi.newrelic.NrqlAlertConditionArgs;
import com.pulumi.newrelic.inputs.NrqlAlertConditionNrqlArgs;
import com.pulumi.newrelic.inputs.NrqlAlertConditionCriticalArgs;
import com.pulumi.newrelic.AlertChannel;
import com.pulumi.newrelic.AlertChannelArgs;
import com.pulumi.newrelic.inputs.AlertChannelConfigArgs;
import com.pulumi.newrelic.AlertPolicyChannel;
import com.pulumi.newrelic.AlertPolicyChannelArgs;
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        // Read an APM application resource
        final var foo = NewrelicFunctions.getEntity(GetEntityArgs.builder()
            .name("Your App Name")
            .domain("APM")
            .type("APPLICATION")
            .build());

        // Create an alert policy
        var alert = new AlertPolicy("alert", AlertPolicyArgs.builder()
            .name("Your Concise Alert Name")
            .build());

        // Add a condition
        var fooNrqlAlertCondition = new NrqlAlertCondition("fooNrqlAlertCondition", NrqlAlertConditionArgs.builder()
            .policyId(alert.id())
            .type("static")
            .name("foo")
            .description("Alert when transactions are taking too long")
            .runbookUrl("https://www.example.com")
            .enabled(true)
            .violationTimeLimitSeconds(3600)
            .nrql(NrqlAlertConditionNrqlArgs.builder()
                .query(String.format("SELECT average(duration) FROM Transaction where appName = '%s'", foo.name()))
                .build())
            .critical(NrqlAlertConditionCriticalArgs.builder()
                .operator("above")
                .threshold(5.5)
                .thresholdDuration(300)
                .thresholdOccurrences("ALL")
                .build())
            .build());

        // Add a notification channel
        var email = new AlertChannel("email", AlertChannelArgs.builder()
            .name("email")
            .type("email")
            .config(AlertChannelConfigArgs.builder()
                .recipients("username@example.com")
                .includeJsonAttachment("1")
                .build())
            .build());

        // Link the channel to the policy
        var alertEmail = new AlertPolicyChannel("alertEmail", AlertPolicyChannelArgs.builder()
            .policyId(alert.id())
            .channelIds(email.id())
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}