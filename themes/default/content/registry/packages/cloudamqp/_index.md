---
# WARNING: this file was fetched from https://raw.githubusercontent.com/pulumi/pulumi-cloudamqp/v3.25.0/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/pulumi/pulumi-cloudamqp/blob/v3.25.0/docs/_index.md
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: CloudAMQP Provider
meta_desc: Provides an overview on how to configure the Pulumi CloudAMQP provider.
layout: package
---

## Installation

The CloudAMQP provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/cloudamqp`](https://www.npmjs.com/package/@pulumi/cloudamqp)
* Python: [`pulumi-cloudamqp`](https://pypi.org/project/pulumi-cloudamqp/)
* Go: [`github.com/pulumi/pulumi-cloudamqp/sdk/v3/go/cloudamqp`](https://github.com/pulumi/pulumi-cloudamqp)
* .NET: [`Pulumi.Cloudamqp`](https://www.nuget.org/packages/Pulumi.Cloudamqp)
* Java: [`com.pulumi/cloudamqp`](https://central.sonatype.com/artifact/com.pulumi/cloudamqp)

## Overview

The CloudAMQP provider is used to interact with CloudAMQP organization resources.

The provider allows you to manage your CloudAMQP instances and features. Create, configure and
deploy [**LavinMQ**](https://lavinmq.com/) or [**RabbitMQ**](https://www.rabbitmq.com/) to different cloud platforms. The provider needs to be
configured with the proper API key before it can be used.

Use the navigation to the left to read about the available resources.
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    cloudamqp:apikey:
        value: 'TODO: var.cloudamqp_customer_api_key'
    cloudamqp:enableFasterInstanceDestroy:
        value: true

```

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as cloudamqp from "@pulumi/cloudamqp";

// Create a new cloudamqp instance
const instance = new cloudamqp.Instance("instance", {
    name: "pulumi-cloudamqp-instance",
    plan: "penguin-1",
    region: "amazon-web-services::us-west-1",
    tags: ["pulumi"],
});
// New recipient to receieve notifications
const recipient01 = new cloudamqp.Notification("recipient_01", {
    instanceId: instance.id,
    type: "email",
    value: "alarm@example.com",
    name: "alarm",
});
// New cpu alarm
const cpuAlarm = new cloudamqp.Alarm("cpu_alarm", {
    instanceId: instance.id,
    type: "cpu",
    valueThreshold: 90,
    timeThreshold: 600,
    enabled: true,
    recipients: [recipient01.id],
});
// Configure firewall
const firewall = new cloudamqp.SecurityFirewall("firewall", {
    instanceId: instance.id,
    rules: [{
        ip: "10.54.72.0/0",
        ports: [4567],
        services: ["AMQPS"],
    }],
});
// Cloudwatch metrics integration
const cloudwatch = new cloudamqp.IntegrationMetric("cloudwatch", {
    instanceId: instance.id,
    name: "cloudwatch",
    accessKeyId: awsAccessKey,
    secretAccessKey: awsSecretKey,
    region: awsRegion,
});
```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    cloudamqp:apikey:
        value: 'TODO: var.cloudamqp_customer_api_key'
    cloudamqp:enableFasterInstanceDestroy:
        value: true

```

```python
import pulumi
import pulumi_cloudamqp as cloudamqp

# Create a new cloudamqp instance
instance = cloudamqp.Instance("instance",
    name="pulumi-cloudamqp-instance",
    plan="penguin-1",
    region="amazon-web-services::us-west-1",
    tags=["pulumi"])
# New recipient to receieve notifications
recipient01 = cloudamqp.Notification("recipient_01",
    instance_id=instance.id,
    type="email",
    value="alarm@example.com",
    name="alarm")
# New cpu alarm
cpu_alarm = cloudamqp.Alarm("cpu_alarm",
    instance_id=instance.id,
    type="cpu",
    value_threshold=90,
    time_threshold=600,
    enabled=True,
    recipients=[recipient01.id])
# Configure firewall
firewall = cloudamqp.SecurityFirewall("firewall",
    instance_id=instance.id,
    rules=[{
        "ip": "10.54.72.0/0",
        "ports": [4567],
        "services": ["AMQPS"],
    }])
# Cloudwatch metrics integration
cloudwatch = cloudamqp.IntegrationMetric("cloudwatch",
    instance_id=instance.id,
    name="cloudwatch",
    access_key_id=aws_access_key,
    secret_access_key=aws_secret_key,
    region=aws_region)
```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    cloudamqp:apikey:
        value: 'TODO: var.cloudamqp_customer_api_key'
    cloudamqp:enableFasterInstanceDestroy:
        value: true

```

```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using CloudAmqp = Pulumi.CloudAmqp;

return await Deployment.RunAsync(() =>
{
    // Create a new cloudamqp instance
    var instance = new CloudAmqp.Instance("instance", new()
    {
        Name = "pulumi-cloudamqp-instance",
        Plan = "penguin-1",
        Region = "amazon-web-services::us-west-1",
        Tags = new[]
        {
            "pulumi",
        },
    });

    // New recipient to receieve notifications
    var recipient01 = new CloudAmqp.Notification("recipient_01", new()
    {
        InstanceId = instance.Id,
        Type = "email",
        Value = "alarm@example.com",
        Name = "alarm",
    });

    // New cpu alarm
    var cpuAlarm = new CloudAmqp.Alarm("cpu_alarm", new()
    {
        InstanceId = instance.Id,
        Type = "cpu",
        ValueThreshold = 90,
        TimeThreshold = 600,
        Enabled = true,
        Recipients = new[]
        {
            recipient01.Id,
        },
    });

    // Configure firewall
    var firewall = new CloudAmqp.SecurityFirewall("firewall", new()
    {
        InstanceId = instance.Id,
        Rules = new[]
        {
            new CloudAmqp.Inputs.SecurityFirewallRuleArgs
            {
                Ip = "10.54.72.0/0",
                Ports = new[]
                {
                    4567,
                },
                Services = new[]
                {
                    "AMQPS",
                },
            },
        },
    });

    // Cloudwatch metrics integration
    var cloudwatch = new CloudAmqp.IntegrationMetric("cloudwatch", new()
    {
        InstanceId = instance.Id,
        Name = "cloudwatch",
        AccessKeyId = awsAccessKey,
        SecretAccessKey = awsSecretKey,
        Region = awsRegion,
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
    cloudamqp:apikey:
        value: 'TODO: var.cloudamqp_customer_api_key'
    cloudamqp:enableFasterInstanceDestroy:
        value: true

```

```go
package main

import (
	"github.com/pulumi/pulumi-cloudamqp/sdk/v3/go/cloudamqp"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a new cloudamqp instance
		instance, err := cloudamqp.NewInstance(ctx, "instance", &cloudamqp.InstanceArgs{
			Name:   pulumi.String("pulumi-cloudamqp-instance"),
			Plan:   pulumi.String("penguin-1"),
			Region: pulumi.String("amazon-web-services::us-west-1"),
			Tags: pulumi.StringArray{
				pulumi.String("pulumi"),
			},
		})
		if err != nil {
			return err
		}
		// New recipient to receieve notifications
		recipient01, err := cloudamqp.NewNotification(ctx, "recipient_01", &cloudamqp.NotificationArgs{
			InstanceId: instance.ID(),
			Type:       pulumi.String("email"),
			Value:      pulumi.String("alarm@example.com"),
			Name:       pulumi.String("alarm"),
		})
		if err != nil {
			return err
		}
		// New cpu alarm
		_, err = cloudamqp.NewAlarm(ctx, "cpu_alarm", &cloudamqp.AlarmArgs{
			InstanceId:     instance.ID(),
			Type:           pulumi.String("cpu"),
			ValueThreshold: pulumi.Int(90),
			TimeThreshold:  pulumi.Int(600),
			Enabled:        pulumi.Bool(true),
			Recipients: pulumi.IntArray{
				recipient01.ID(),
			},
		})
		if err != nil {
			return err
		}
		// Configure firewall
		_, err = cloudamqp.NewSecurityFirewall(ctx, "firewall", &cloudamqp.SecurityFirewallArgs{
			InstanceId: instance.ID(),
			Rules: cloudamqp.SecurityFirewallRuleArray{
				&cloudamqp.SecurityFirewallRuleArgs{
					Ip: pulumi.String("10.54.72.0/0"),
					Ports: pulumi.IntArray{
						pulumi.Int(4567),
					},
					Services: pulumi.StringArray{
						pulumi.String("AMQPS"),
					},
				},
			},
		})
		if err != nil {
			return err
		}
		// Cloudwatch metrics integration
		_, err = cloudamqp.NewIntegrationMetric(ctx, "cloudwatch", &cloudamqp.IntegrationMetricArgs{
			InstanceId:      instance.ID(),
			Name:            pulumi.String("cloudwatch"),
			AccessKeyId:     pulumi.Any(awsAccessKey),
			SecretAccessKey: pulumi.Any(awsSecretKey),
			Region:          pulumi.Any(awsRegion),
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
    cloudamqp:apikey:
        value: 'TODO: var.cloudamqp_customer_api_key'
    cloudamqp:enableFasterInstanceDestroy:
        value: true

```

```yaml
resources:
  # Create a new cloudamqp instance
  instance:
    type: cloudamqp:Instance
    properties:
      name: pulumi-cloudamqp-instance
      plan: penguin-1
      region: amazon-web-services::us-west-1
      tags:
        - pulumi
  # New recipient to receieve notifications
  recipient01:
    type: cloudamqp:Notification
    name: recipient_01
    properties:
      instanceId: ${instance.id}
      type: email
      value: alarm@example.com
      name: alarm
  # New cpu alarm
  cpuAlarm:
    type: cloudamqp:Alarm
    name: cpu_alarm
    properties:
      instanceId: ${instance.id}
      type: cpu
      valueThreshold: 90
      timeThreshold: 600
      enabled: true
      recipients:
        - ${recipient01.id}
  # Configure firewall
  firewall:
    type: cloudamqp:SecurityFirewall
    properties:
      instanceId: ${instance.id}
      rules:
        - ip: 10.54.72.0/0
          ports:
            - 4567
          services:
            - AMQPS
  # Cloudwatch metrics integration
  cloudwatch:
    type: cloudamqp:IntegrationMetric
    properties:
      instanceId: ${instance.id}
      name: cloudwatch
      accessKeyId: ${awsAccessKey}
      secretAccessKey: ${awsSecretKey}
      region: ${awsRegion}
```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    cloudamqp:apikey:
        value: 'TODO: var.cloudamqp_customer_api_key'
    cloudamqp:enableFasterInstanceDestroy:
        value: true

```

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.cloudamqp.Instance;
import com.pulumi.cloudamqp.InstanceArgs;
import com.pulumi.cloudamqp.Notification;
import com.pulumi.cloudamqp.NotificationArgs;
import com.pulumi.cloudamqp.Alarm;
import com.pulumi.cloudamqp.AlarmArgs;
import com.pulumi.cloudamqp.SecurityFirewall;
import com.pulumi.cloudamqp.SecurityFirewallArgs;
import com.pulumi.cloudamqp.inputs.SecurityFirewallRuleArgs;
import com.pulumi.cloudamqp.IntegrationMetric;
import com.pulumi.cloudamqp.IntegrationMetricArgs;
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
        // Create a new cloudamqp instance
        var instance = new Instance("instance", InstanceArgs.builder()
            .name("pulumi-cloudamqp-instance")
            .plan("penguin-1")
            .region("amazon-web-services::us-west-1")
            .tags("pulumi")
            .build());

        // New recipient to receieve notifications
        var recipient01 = new Notification("recipient01", NotificationArgs.builder()
            .instanceId(instance.id())
            .type("email")
            .value("alarm@example.com")
            .name("alarm")
            .build());

        // New cpu alarm
        var cpuAlarm = new Alarm("cpuAlarm", AlarmArgs.builder()
            .instanceId(instance.id())
            .type("cpu")
            .valueThreshold(90)
            .timeThreshold(600)
            .enabled(true)
            .recipients(recipient01.id())
            .build());

        // Configure firewall
        var firewall = new SecurityFirewall("firewall", SecurityFirewallArgs.builder()
            .instanceId(instance.id())
            .rules(SecurityFirewallRuleArgs.builder()
                .ip("10.54.72.0/0")
                .ports(4567)
                .services("AMQPS")
                .build())
            .build());

        // Cloudwatch metrics integration
        var cloudwatch = new IntegrationMetric("cloudwatch", IntegrationMetricArgs.builder()
            .instanceId(instance.id())
            .name("cloudwatch")
            .accessKeyId(awsAccessKey)
            .secretAccessKey(awsSecretKey)
            .region(awsRegion)
            .build());

    }
}
```

{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

The following configuration inputs are supported:

* `apikey` - (Required) This is the CloudAMQP Customer API key needed to make calls to the customer
  API. It can be sourced from login in to your CloudAMQP account and go to API access or
  go directly to [API Keys](https://customer.cloudamqp.com/apikeys). The API key can also be read from the environment variable
  `CLOUDAMQP_APIKEY`.

* `enableFasterInstanceDestroy` - (Optional) This will speed up the destroy action for
  `cloudamqp.Instance` when running `pulumi destroy`. It's
  done by skipping delete behaviour for resources that don't
  need to be cleaned up when the servers are deleted. The
  argument can also be read from the environment variable
  `CLOUDAMQP_ENABLE_FASTER_INSTANCE_DESTROY`, default set to
  false.

  ***Note:*** Available from [v1.27.0](https://github.com/cloudamqp/pulumi-provider-cloudamqp/releases/tag/v1.27.0).

---

***List of resources affected by `enableFasterInstanceDestroy`:***

* cloudamqp.Plugin
* cloudamqp.PluginCommunity
* cloudamqp.SecurityFirewall

More information can be found under `Enable faster instance destroy` section on respective resource.