---
title: RedisCloud
meta_desc: Provides an overview of the RedisCloud Provider for Pulumi.
layout: overview
---

The [Redis Cloud](https://redis.com) Resource Provider lets you manage Redis Cloud resources. The provider is used to interact with the resources supported by Redis Enterprise Cloud. Please note that the provider is available only for Flexible or Annual subscriptions. It is not supported for Fixed or Free subscriptions.

## Example

{{< chooser language "typescript,python,go,yaml" >}}
{{% choosable language typescript %}}

```typescript
"use strict";
import rediscloud from "@rediscloud/pulumi-rediscloud";

const card = await rediscloud.getPaymentMethod({
  cardType: "Visa",
  lastFourNumbers: "1234",
});

const subscription = new rediscloud.Subscription("my-subscription", {
  name: "my-subscription",
  paymentMethod: "credit-card",
  paymentMethodId: card.id,
  cloudProvider: {
    regions: [
      {
        region: "us-east-1",
        multipleAvailabilityZones: false,
        networkingDeploymentCidr: "10.0.0.0/24",
        preferredAvailabilityZones: ["use1-az1", "use1-az2", "use1-az5"],
      },
    ],
  },

  creationPlan: {
    memoryLimitInGb: 10,
    quantity: 1,
    replication: true,
    supportOssClusterApi: false,
    throughputMeasurementBy: "operations-per-second",
    throughputMeasurementValue: 20000,
    modules: ["RedisJSON"],
  },
});

const database = new rediscloud.SubscriptionDatabase("my-db", {
  name: "my-db",
  subscriptionId: subscription.id,
  protocol: "redis",
  memoryLimitInGb: 10,
  dataPersistence: "aof-every-1-second",
  throughputMeasurementBy: "operations-per-second",
  throughputMeasurementValue: 20000,
  replication: true,
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_rediscloud

card = pulumi_rediscloud.get_payment_method(
    card_type="Visa",
    last_four_numbers="1234",
)

subscription = pulumi_rediscloud.Subscription(
    "my-subscription",
    name="my-subscription",
    payment_method="credit-card",
    payment_method_id=card.id,
    cloud_provider=pulumi_rediscloud.SubscriptionCloudProviderArgs(
        regions=[
            pulumi_rediscloud.SubscriptionCloudProviderRegionArgs(
                region="us-east-1",
                multiple_availability_zones=True,
                networking_deployment_cidr="10.0.0.0/24",
                preferred_availability_zones=["use1-az1", "use1-az2", "use1-az5"],
            )
        ]
    ),
    creation_plan=pulumi_rediscloud.SubscriptionCreationPlanArgs(
        memory_limit_in_gb=10,
        quantity=1,
        replication=True,
        support_oss_cluster_api=False,
        throughput_measurement_by="operations-per-second",
        throughput_measurement_value=20000,
        modules=["RedisJSON"],
    ),
)

database = pulumi_rediscloud.SubscriptionDatabase(
    "my-db",
    name="my-db",
    subscription_id=subscription.id,
    protocol="redis",
    memory_limit_in_gb=10,
    data_persistence="aof-every-1-second",
    throughput_measurement_by="operations-per-second",
    throughput_measurement_value=20000,
    replication=True,
)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
package main

import (
	"github.com/RedisLabs/pulumi-rediscloud/sdk/go/rediscloud"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		card, err := rediscloud.GetPaymentMethod(ctx, &rediscloud.GetPaymentMethodArgs{
			CardType:        pulumi.StringRef("Visa"),
			LastFourNumbers: pulumi.StringRef("1234"),
		}, nil)

		if err != nil {
			return err
		}

		cardId := card.Id

		subscription, err := rediscloud.NewSubscription(ctx, "subscription", &rediscloud.SubscriptionArgs{
			PaymentMethod:   pulumi.String("credit-card"),
			PaymentMethodId: pulumi.String(cardId),
			CloudProvider: &rediscloud.SubscriptionCloudProviderArgs{
				Regions: rediscloud.SubscriptionCloudProviderRegionArray{
					&rediscloud.SubscriptionCloudProviderRegionArgs{
						Region:                    pulumi.String("us-east-1"),
						MultipleAvailabilityZones: pulumi.Bool(false),
						NetworkingDeploymentCidr:  pulumi.String("10.0.0.0/24"),
						PreferredAvailabilityZones: pulumi.StringArray{
							pulumi.String("use1-az1"),
							pulumi.String("use1-az2"),
							pulumi.String("use1-az3"),
						},
					},
				},
			},
			CreationPlan: &rediscloud.SubscriptionCreationPlanArgs{
				MemoryLimitInGb:            pulumi.Float64(10),
				Quantity:                   pulumi.Int(1),
				Replication:                pulumi.Bool(true),
				SupportOssClusterApi:       pulumi.Bool(false),
				ThroughputMeasurementBy:    pulumi.String("operations-per-second"),
				ThroughputMeasurementValue: pulumi.Int(20000),
				Modules:                    pulumi.StringArray{"RedisJSON"},
			},
		})
		if err != nil {
			return err
		}
		_, err = rediscloud.NewSubscriptionDatabase(ctx, "database", &rediscloud.SubscriptionDatabaseArgs{
			SubscriptionId:             subscription.ID(),
			Protocol:                   pulumi.String("redis"),
			MemoryLimitInGb:            pulumi.Float64(10),
			DataPersistence:            pulumi.String("aof-every-1-second"),
			ThroughputMeasurementBy:    pulumi.String("operations-per-second"),
			ThroughputMeasurementValue: pulumi.Int(20000),
			SupportOssClusterApi:       pulumi.Bool(false),
			Replication:                pulumi.Bool(true),
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
name: rediscloud-yaml
runtime: yaml
description: A minimal Pulumi YAML program
variables:
  cardId:
    fn::invoke:
      function: rediscloud:getPaymentMethod
      arguments:
        cardType: Visa
        lastFourNumbers: "1234"
      return: id
outputs: {}
resources:
  subscription:
    type: rediscloud:Subscription
    properties:
      paymentMethod: credit-card
      paymentMethodId: ${cardId}
      cloudProvider:
        regions:
          - region: us-east-1
            multipleAvailabilityZones: true
            networkingDeploymentCidr: "10.0.0.0/24"
            preferredAvailabilityZones:
              - "use1-az1"
              - "use1-az2"
              - "use1-az5"
      creationPlan:
        memoryLimitInGb: 10
        quantity: 1
        replication: true
        supportOssClusterApi: false
        throughputMeasurementBy: operations-per-second
        throughputMeasurementValue: 20000
        modules:
          - RedisJSON
  database:
    type: rediscloud:SubscriptionDatabase
    properties:
      subscriptionId: ${subscription.id}
      protocol: redis
      memoryLimitInGb: 10
      dataPersistence: aof-every-1-second
      throughputMeasurementBy: "operations-per-second"
      throughputMeasurementValue: 20000
      replication: true
      modules:
        - name: RedisJSON
```

{{% /choosable %}}

{{< /chooser >}}
