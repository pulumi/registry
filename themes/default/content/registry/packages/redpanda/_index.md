---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/redpanda-data/redpanda/1.0.2/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Redpanda Provider
meta_desc: Provides an overview on how to configure the Pulumi Redpanda provider.
layout: package
---

## Generate Provider

The Redpanda provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider redpanda-data/redpanda
```
## Overview

The Redpanda provider is designed for managing Redpanda clusters and Kafka resources in Redpanda Dedicated and Cloud environments. It supports the provisioning, management, and configuration of clusters and Kafka resources, facilitating seamless integration into Pulumi workflows.
## Configuration Reference

- `accessToken` (String, Sensitive) Redpanda client token. You need either `accessToken`, or both `clientId` and `clientSecret` to use this provider. Can also be set with the `REDPANDA_ACCESS_TOKEN` environment variable.
- `azureClientId` (String) Used for creating and managing BYOC and BYOVPC clusters. Can also be specified in the environment as AZURE_CLIENT_ID or ARM_CLIENT_ID
- `azureClientSecret` (String) Used for creating and managing BYOC and BYOVPC clusters. Can also be specified in the environment as AZURE_CLIENT_SECRET or ARM_CLIENT_SECRET
- `azureSubscriptionId` (String) The default Azure Subscription ID which should be used for Redpanda BYOC clusters. If another subscription is specified on a resource, it will take precedence. This can also be sourced from the `ARM_SUBSCRIPTION_ID` environment variable.
- `azureTenantId` (String) Used for creating and managing BYOC and BYOVPC clusters. Can also be specified in the environment as AZURE_TENANT_ID or ARM_TENANT_ID
- `clientId` (String, Sensitive) The ID for the client. You need either `clientId` AND `clientSecret`, or `accessToken`, to use this provider. Can also be set with the `REDPANDA_CLIENT_ID` environment variable.
- `clientSecret` (String, Sensitive) Redpanda client secret. You need either `clientId` AND `clientSecret`, or `accessToken`, to use this provider. Can also be set with the `REDPANDA_CLIENT_SECRET` environment variable.
- `gcpProjectId` (String) The default Google Cloud Project ID to use for Redpanda BYOC clusters. If another project is specified on a resource, it will take precedence. This can also be sourced from the `GOOGLE_PROJECT` environment variable, or any of the following ordered by precedence: `GOOGLE_PROJECT`, `GOOGLE_CLOUD_PROJECT`, `GCLOUD_PROJECT`, or `CLOUDSDK_CORE_PROJECT`.
- `googleCredentials` (String) Used for creating and managing BYOC and BYOVPC clusters. Can also be specified in the environment as GOOGLE_CREDENTIALS
- `googleCredentialsBase64` (String) Used for creating and managing BYOC and BYOVPC clusters. Is a convenience passthrough for base64 encoded credentials intended for use in CI/CD. Can also be specified in the environment as GOOGLE_CREDENTIALS_BASE64
## Authentication with Redpanda Cloud

This provider requires a `clientId` and `clientSecret` for authentication with Redpanda Cloud services, enabling users to securely manage their Redpanda resources. You can get these by creating an account in [Redpanda Cloud](https://cloudv2.redpanda.com/home) and then [creating a client in the Redpanda Cloud UI](https://cloudv2.redpanda.com/clients).
## Example Provider Configuration

Pulumi 1.0 or later:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    redpanda:clientId:
        value: your_client_id
    redpanda:clientSecret:
        value: your_client_secret

```
```typescript
import * as pulumi from "@pulumi/pulumi";

```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    redpanda:clientId:
        value: your_client_id
    redpanda:clientSecret:
        value: your_client_secret

```
```python
import pulumi

```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    redpanda:clientId:
        value: your_client_id
    redpanda:clientSecret:
        value: your_client_secret

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;

return await Deployment.RunAsync(() =>
{
});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    redpanda:clientId:
        value: your_client_id
    redpanda:clientSecret:
        value: your_client_secret

```
```go
package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
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
    redpanda:clientId:
        value: your_client_id
    redpanda:clientSecret:
        value: your_client_secret

```
```yaml
{}
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    redpanda:clientId:
        value: your_client_id
    redpanda:clientSecret:
        value: your_client_secret

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
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
    }
}
```
{{% /choosable %}}
{{< /chooser >}}
### Example Usage for an AWS Dedicated Cluster

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as redpanda from "@pulumi/redpanda";

const config = new pulumi.Config();
const resourceGroupName = config.get("resourceGroupName") || "testname";
const test = new redpanda.ResourceGroup("test", {name: resourceGroupName});
const networkName = config.get("networkName") || "testname";
const region = config.get("region") || "us-east-2";
const cloudProvider = config.get("cloudProvider") || "aws";
const testNetwork = new redpanda.Network("test", {
    name: networkName,
    resourceGroupId: test.id,
    cloudProvider: cloudProvider,
    region: region,
    clusterType: "dedicated",
    cidrBlock: "10.0.0.0/20",
});
const clusterName = config.get("clusterName") || "testname";
const zones = config.getObject("zones") || [
    "use2-az1",
    "use2-az2",
    "use2-az3",
];
const throughputTier = config.get("throughputTier") || "tier-1-aws-v2-arm";
const testCluster = new redpanda.Cluster("test", {
    name: clusterName,
    resourceGroupId: test.id,
    networkId: testNetwork.id,
    cloudProvider: cloudProvider,
    region: region,
    clusterType: "dedicated",
    connectionType: "public",
    throughputTier: throughputTier,
    zones: zones,
    allowDeletion: true,
    tags: {
        key: "value",
    },
});
const userName = config.get("userName") || "test-username";
const userPw = config.get("userPw") || "password";
const mechanism = config.get("mechanism") || "scram-sha-256";
const testUser = new redpanda.User("test", {
    name: userName,
    password: userPw,
    mechanism: mechanism,
    clusterApiUrl: testCluster.clusterApiUrl,
});
const topicName = config.get("topicName") || "test-topic";
const partitionCount = config.getNumber("partitionCount") || 3;
const replicationFactor = config.getNumber("replicationFactor") || 3;
const testTopic = new redpanda.Topic("test", {
    name: topicName,
    partitionCount: partitionCount,
    replicationFactor: replicationFactor,
    clusterApiUrl: testCluster.clusterApiUrl,
    allowDeletion: true,
});
const testAcl = new redpanda.Acl("test", {
    resourceType: "TOPIC",
    resourceName: testTopic.name,
    resourcePatternType: "LITERAL",
    principal: pulumi.interpolate`User:${testUser.name}`,
    host: "*",
    operation: "READ",
    permissionType: "ALLOW",
    clusterApiUrl: testCluster.clusterApiUrl,
});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python

```
```python
import pulumi
import pulumi_redpanda as redpanda

config = pulumi.Config()
resource_group_name = config.get("resourceGroupName")
if resource_group_name is None:
    resource_group_name = "testname"
test = redpanda.ResourceGroup("test", name=resource_group_name)
network_name = config.get("networkName")
if network_name is None:
    network_name = "testname"
region = config.get("region")
if region is None:
    region = "us-east-2"
cloud_provider = config.get("cloudProvider")
if cloud_provider is None:
    cloud_provider = "aws"
test_network = redpanda.Network("test",
    name=network_name,
    resource_group_id=test.id,
    cloud_provider=cloud_provider,
    region=region,
    cluster_type="dedicated",
    cidr_block="10.0.0.0/20")
cluster_name = config.get("clusterName")
if cluster_name is None:
    cluster_name = "testname"
zones = config.get_object("zones")
if zones is None:
    zones = [
        "use2-az1",
        "use2-az2",
        "use2-az3",
    ]
throughput_tier = config.get("throughputTier")
if throughput_tier is None:
    throughput_tier = "tier-1-aws-v2-arm"
test_cluster = redpanda.Cluster("test",
    name=cluster_name,
    resource_group_id=test.id,
    network_id=test_network.id,
    cloud_provider=cloud_provider,
    region=region,
    cluster_type="dedicated",
    connection_type="public",
    throughput_tier=throughput_tier,
    zones=zones,
    allow_deletion=True,
    tags={
        "key": "value",
    })
user_name = config.get("userName")
if user_name is None:
    user_name = "test-username"
user_pw = config.get("userPw")
if user_pw is None:
    user_pw = "password"
mechanism = config.get("mechanism")
if mechanism is None:
    mechanism = "scram-sha-256"
test_user = redpanda.User("test",
    name=user_name,
    password=user_pw,
    mechanism=mechanism,
    cluster_api_url=test_cluster.cluster_api_url)
topic_name = config.get("topicName")
if topic_name is None:
    topic_name = "test-topic"
partition_count = config.get_float("partitionCount")
if partition_count is None:
    partition_count = 3
replication_factor = config.get_float("replicationFactor")
if replication_factor is None:
    replication_factor = 3
test_topic = redpanda.Topic("test",
    name=topic_name,
    partition_count=partition_count,
    replication_factor=replication_factor,
    cluster_api_url=test_cluster.cluster_api_url,
    allow_deletion=True)
test_acl = redpanda.Acl("test",
    resource_type="TOPIC",
    resource_name_=test_topic.name,
    resource_pattern_type="LITERAL",
    principal=test_user.name.apply(lambda name: f"User:{name}"),
    host="*",
    operation="READ",
    permission_type="ALLOW",
    cluster_api_url=test_cluster.cluster_api_url)
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Redpanda = Pulumi.Redpanda;

return await Deployment.RunAsync(() =>
{
    var config = new Config();
    var resourceGroupName = config.Get("resourceGroupName") ?? "testname";
    var test = new Redpanda.ResourceGroup("test", new()
    {
        Name = resourceGroupName,
    });

    var networkName = config.Get("networkName") ?? "testname";
    var region = config.Get("region") ?? "us-east-2";
    var cloudProvider = config.Get("cloudProvider") ?? "aws";
    var testNetwork = new Redpanda.Network("test", new()
    {
        Name = networkName,
        ResourceGroupId = test.Id,
        CloudProvider = cloudProvider,
        Region = region,
        ClusterType = "dedicated",
        CidrBlock = "10.0.0.0/20",
    });

    var clusterName = config.Get("clusterName") ?? "testname";
    var zones = config.GetObject<dynamic>("zones") ?? new[]
    {
        "use2-az1",
        "use2-az2",
        "use2-az3",
    };
    var throughputTier = config.Get("throughputTier") ?? "tier-1-aws-v2-arm";
    var testCluster = new Redpanda.Cluster("test", new()
    {
        Name = clusterName,
        ResourceGroupId = test.Id,
        NetworkId = testNetwork.Id,
        CloudProvider = cloudProvider,
        Region = region,
        ClusterType = "dedicated",
        ConnectionType = "public",
        ThroughputTier = throughputTier,
        Zones = zones,
        AllowDeletion = true,
        Tags =
        {
            { "key", "value" },
        },
    });

    var userName = config.Get("userName") ?? "test-username";
    var userPw = config.Get("userPw") ?? "password";
    var mechanism = config.Get("mechanism") ?? "scram-sha-256";
    var testUser = new Redpanda.User("test", new()
    {
        Name = userName,
        Password = userPw,
        Mechanism = mechanism,
        ClusterApiUrl = testCluster.ClusterApiUrl,
    });

    var topicName = config.Get("topicName") ?? "test-topic";
    var partitionCount = config.GetDouble("partitionCount") ?? 3;
    var replicationFactor = config.GetDouble("replicationFactor") ?? 3;
    var testTopic = new Redpanda.Topic("test", new()
    {
        Name = topicName,
        PartitionCount = partitionCount,
        ReplicationFactor = replicationFactor,
        ClusterApiUrl = testCluster.ClusterApiUrl,
        AllowDeletion = true,
    });

    var testAcl = new Redpanda.Acl("test", new()
    {
        ResourceType = "TOPIC",
        ResourceName = testTopic.Name,
        ResourcePatternType = "LITERAL",
        Principal = testUser.Name.Apply(name => $"User:{name}"),
        Host = "*",
        Operation = "READ",
        PermissionType = "ALLOW",
        ClusterApiUrl = testCluster.ClusterApiUrl,
    });

});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go

```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/redpanda/redpanda"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")
		resourceGroupName := "testname"
		if param := cfg.Get("resourceGroupName"); param != "" {
			resourceGroupName = param
		}
		test, err := redpanda.NewResourceGroup(ctx, "test", &redpanda.ResourceGroupArgs{
			Name: pulumi.String(resourceGroupName),
		})
		if err != nil {
			return err
		}
		networkName := "testname"
		if param := cfg.Get("networkName"); param != "" {
			networkName = param
		}
		region := "us-east-2"
		if param := cfg.Get("region"); param != "" {
			region = param
		}
		cloudProvider := "aws"
		if param := cfg.Get("cloudProvider"); param != "" {
			cloudProvider = param
		}
		testNetwork, err := redpanda.NewNetwork(ctx, "test", &redpanda.NetworkArgs{
			Name:            pulumi.String(networkName),
			ResourceGroupId: test.ID(),
			CloudProvider:   pulumi.String(cloudProvider),
			Region:          pulumi.String(region),
			ClusterType:     pulumi.String("dedicated"),
			CidrBlock:       pulumi.String("10.0.0.0/20"),
		})
		if err != nil {
			return err
		}
		clusterName := "testname"
		if param := cfg.Get("clusterName"); param != "" {
			clusterName = param
		}
		zones := []string{
			"use2-az1",
			"use2-az2",
			"use2-az3",
		}
		if param := cfg.GetObject("zones"); param != nil {
			zones = param
		}
		throughputTier := "tier-1-aws-v2-arm"
		if param := cfg.Get("throughputTier"); param != "" {
			throughputTier = param
		}
		testCluster, err := redpanda.NewCluster(ctx, "test", &redpanda.ClusterArgs{
			Name:            pulumi.String(clusterName),
			ResourceGroupId: test.ID(),
			NetworkId:       testNetwork.ID(),
			CloudProvider:   pulumi.String(cloudProvider),
			Region:          pulumi.String(region),
			ClusterType:     pulumi.String("dedicated"),
			ConnectionType:  pulumi.String("public"),
			ThroughputTier:  pulumi.String(throughputTier),
			Zones:           pulumi.Any(zones),
			AllowDeletion:   pulumi.Bool(true),
			Tags: pulumi.StringMap{
				"key": pulumi.String("value"),
			},
		})
		if err != nil {
			return err
		}
		userName := "test-username"
		if param := cfg.Get("userName"); param != "" {
			userName = param
		}
		userPw := "password"
		if param := cfg.Get("userPw"); param != "" {
			userPw = param
		}
		mechanism := "scram-sha-256"
		if param := cfg.Get("mechanism"); param != "" {
			mechanism = param
		}
		testUser, err := redpanda.NewUser(ctx, "test", &redpanda.UserArgs{
			Name:          pulumi.String(userName),
			Password:      pulumi.String(userPw),
			Mechanism:     pulumi.String(mechanism),
			ClusterApiUrl: testCluster.ClusterApiUrl,
		})
		if err != nil {
			return err
		}
		topicName := "test-topic"
		if param := cfg.Get("topicName"); param != "" {
			topicName = param
		}
		partitionCount := float64(3)
		if param := cfg.GetFloat64("partitionCount"); param != 0 {
			partitionCount = param
		}
		replicationFactor := float64(3)
		if param := cfg.GetFloat64("replicationFactor"); param != 0 {
			replicationFactor = param
		}
		testTopic, err := redpanda.NewTopic(ctx, "test", &redpanda.TopicArgs{
			Name:              pulumi.String(topicName),
			PartitionCount:    pulumi.Float64(partitionCount),
			ReplicationFactor: pulumi.Float64(replicationFactor),
			ClusterApiUrl:     testCluster.ClusterApiUrl,
			AllowDeletion:     pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		_, err = redpanda.NewAcl(ctx, "test", &redpanda.AclArgs{
			ResourceType:        pulumi.String("TOPIC"),
			ResourceName:        testTopic.Name,
			ResourcePatternType: pulumi.String("LITERAL"),
			Principal: testUser.Name.ApplyT(func(name string) (string, error) {
				return fmt.Sprintf("User:%v", name), nil
			}).(pulumi.StringOutput),
			Host:           pulumi.String("*"),
			Operation:      pulumi.String("READ"),
			PermissionType: pulumi.String("ALLOW"),
			ClusterApiUrl:  testCluster.ClusterApiUrl,
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

```
```yaml
configuration:
  resourceGroupName:
    type: string
    default: testname
  networkName:
    type: string
    default: testname
  clusterName:
    type: string
    default: testname
  region:
    type: string
    default: us-east-2
  zones:
    type: dynamic
    default:
      - use2-az1
      - use2-az2
      - use2-az3
  cloudProvider:
    type: string
    default: aws
  throughputTier:
    type: string
    default: tier-1-aws-v2-arm
  userName:
    type: string
    default: test-username
  userPw:
    type: string
    default: password
  mechanism:
    type: string
    default: scram-sha-256
  topicName:
    type: string
    default: test-topic
  partitionCount:
    type: number
    default: 3
  replicationFactor:
    type: number
    default: 3
resources:
  test:
    type: redpanda:ResourceGroup
    properties:
      name: ${resourceGroupName}
  testNetwork:
    type: redpanda:Network
    name: test
    properties:
      name: ${networkName}
      resourceGroupId: ${test.id}
      cloudProvider: ${cloudProvider}
      region: ${region}
      clusterType: dedicated
      cidrBlock: 10.0.0.0/20
  testCluster:
    type: redpanda:Cluster
    name: test
    properties:
      name: ${clusterName}
      resourceGroupId: ${test.id}
      networkId: ${testNetwork.id}
      cloudProvider: ${cloudProvider}
      region: ${region}
      clusterType: dedicated
      connectionType: public
      throughputTier: ${throughputTier}
      zones: ${zones}
      allowDeletion: true
      tags:
        key: value
  testUser:
    type: redpanda:User
    name: test
    properties:
      name: ${userName}
      password: ${userPw}
      mechanism: ${mechanism}
      clusterApiUrl: ${testCluster.clusterApiUrl}
  testTopic:
    type: redpanda:Topic
    name: test
    properties:
      name: ${topicName}
      partitionCount: ${partitionCount}
      replicationFactor: ${replicationFactor}
      clusterApiUrl: ${testCluster.clusterApiUrl}
      allowDeletion: true
  testAcl:
    type: redpanda:Acl
    name: test
    properties:
      resourceType: TOPIC
      resourceName: ${testTopic.name}
      resourcePatternType: LITERAL
      principal: User:${testUser.name}
      host: '*'
      operation: READ
      permissionType: ALLOW
      clusterApiUrl: ${testCluster.clusterApiUrl}
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.redpanda.ResourceGroup;
import com.pulumi.redpanda.ResourceGroupArgs;
import com.pulumi.redpanda.Network;
import com.pulumi.redpanda.NetworkArgs;
import com.pulumi.redpanda.Cluster;
import com.pulumi.redpanda.ClusterArgs;
import com.pulumi.redpanda.User;
import com.pulumi.redpanda.UserArgs;
import com.pulumi.redpanda.Topic;
import com.pulumi.redpanda.TopicArgs;
import com.pulumi.redpanda.Acl;
import com.pulumi.redpanda.AclArgs;
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
        final var config = ctx.config();
        final var resourceGroupName = config.get("resourceGroupName").orElse("testname");
        var test = new ResourceGroup("test", ResourceGroupArgs.builder()
            .name(resourceGroupName)
            .build());

        final var networkName = config.get("networkName").orElse("testname");
        final var region = config.get("region").orElse("us-east-2");
        final var cloudProvider = config.get("cloudProvider").orElse("aws");
        var testNetwork = new Network("testNetwork", NetworkArgs.builder()
            .name(networkName)
            .resourceGroupId(test.id())
            .cloudProvider(cloudProvider)
            .region(region)
            .clusterType("dedicated")
            .cidrBlock("10.0.0.0/20")
            .build());

        final var clusterName = config.get("clusterName").orElse("testname");
        final var zones = config.get("zones").orElse(
            "use2-az1",
            "use2-az2",
            "use2-az3");
        final var throughputTier = config.get("throughputTier").orElse("tier-1-aws-v2-arm");
        var testCluster = new Cluster("testCluster", ClusterArgs.builder()
            .name(clusterName)
            .resourceGroupId(test.id())
            .networkId(testNetwork.id())
            .cloudProvider(cloudProvider)
            .region(region)
            .clusterType("dedicated")
            .connectionType("public")
            .throughputTier(throughputTier)
            .zones(zones)
            .allowDeletion(true)
            .tags(Map.of("key", "value"))
            .build());

        final var userName = config.get("userName").orElse("test-username");
        final var userPw = config.get("userPw").orElse("password");
        final var mechanism = config.get("mechanism").orElse("scram-sha-256");
        var testUser = new User("testUser", UserArgs.builder()
            .name(userName)
            .password(userPw)
            .mechanism(mechanism)
            .clusterApiUrl(testCluster.clusterApiUrl())
            .build());

        final var topicName = config.get("topicName").orElse("test-topic");
        final var partitionCount = config.get("partitionCount").orElse(3);
        final var replicationFactor = config.get("replicationFactor").orElse(3);
        var testTopic = new Topic("testTopic", TopicArgs.builder()
            .name(topicName)
            .partitionCount(partitionCount)
            .replicationFactor(replicationFactor)
            .clusterApiUrl(testCluster.clusterApiUrl())
            .allowDeletion(true)
            .build());

        var testAcl = new Acl("testAcl", AclArgs.builder()
            .resourceType("TOPIC")
            .resourceName(testTopic.name())
            .resourcePatternType("LITERAL")
            .principal(testUser.name().applyValue(name -> String.format("User:%s", name)))
            .host("*")
            .operation("READ")
            .permissionType("ALLOW")
            .clusterApiUrl(testCluster.clusterApiUrl())
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
### Example Usage for a GCP Dedicated Cluster

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as redpanda from "@pulumi/redpanda";

const config = new pulumi.Config();
const resourceGroupName = config.get("resourceGroupName") || "";
const test = new redpanda.ResourceGroup("test", {name: resourceGroupName});
const networkName = config.get("networkName") || "";
const region = config.get("region") || "us-central1";
const cloudProvider = config.get("cloudProvider") || "gcp";
const testNetwork = new redpanda.Network("test", {
    name: networkName,
    resourceGroupId: test.id,
    cloudProvider: cloudProvider,
    region: region,
    clusterType: "dedicated",
    cidrBlock: "10.0.0.0/20",
});
const clusterName = config.get("clusterName") || "";
const zones = config.getObject("zones") || [
    "us-central1-a",
    "us-central1-b",
    "us-central1-c",
];
const throughputTier = config.get("throughputTier") || "tier-1-gcp-um4g";
const testCluster = new redpanda.Cluster("test", {
    name: clusterName,
    resourceGroupId: test.id,
    networkId: testNetwork.id,
    cloudProvider: cloudProvider,
    region: region,
    clusterType: "dedicated",
    connectionType: "public",
    throughputTier: throughputTier,
    zones: zones,
    allowDeletion: true,
});
const userName = config.get("userName") || "test-username";
const userPw = config.get("userPw") || "password";
const mechanism = config.get("mechanism") || "scram-sha-256";
const testUser = new redpanda.User("test", {
    name: userName,
    password: userPw,
    mechanism: mechanism,
    clusterApiUrl: testCluster.clusterApiUrl,
});
const topicName = config.get("topicName") || "test-topic";
const partitionCount = config.getNumber("partitionCount") || 3;
const replicationFactor = config.getNumber("replicationFactor") || 3;
const testTopic = new redpanda.Topic("test", {
    name: topicName,
    partitionCount: partitionCount,
    replicationFactor: replicationFactor,
    clusterApiUrl: testCluster.clusterApiUrl,
    allowDeletion: true,
});
const testAcl = new redpanda.Acl("test", {
    resourceType: "TOPIC",
    resourceName: testTopic.name,
    resourcePatternType: "LITERAL",
    principal: pulumi.interpolate`User:${testUser.name}`,
    host: "*",
    operation: "READ",
    permissionType: "ALLOW",
    clusterApiUrl: testCluster.clusterApiUrl,
});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python

```
```python
import pulumi
import pulumi_redpanda as redpanda

config = pulumi.Config()
resource_group_name = config.get("resourceGroupName")
if resource_group_name is None:
    resource_group_name = ""
test = redpanda.ResourceGroup("test", name=resource_group_name)
network_name = config.get("networkName")
if network_name is None:
    network_name = ""
region = config.get("region")
if region is None:
    region = "us-central1"
cloud_provider = config.get("cloudProvider")
if cloud_provider is None:
    cloud_provider = "gcp"
test_network = redpanda.Network("test",
    name=network_name,
    resource_group_id=test.id,
    cloud_provider=cloud_provider,
    region=region,
    cluster_type="dedicated",
    cidr_block="10.0.0.0/20")
cluster_name = config.get("clusterName")
if cluster_name is None:
    cluster_name = ""
zones = config.get_object("zones")
if zones is None:
    zones = [
        "us-central1-a",
        "us-central1-b",
        "us-central1-c",
    ]
throughput_tier = config.get("throughputTier")
if throughput_tier is None:
    throughput_tier = "tier-1-gcp-um4g"
test_cluster = redpanda.Cluster("test",
    name=cluster_name,
    resource_group_id=test.id,
    network_id=test_network.id,
    cloud_provider=cloud_provider,
    region=region,
    cluster_type="dedicated",
    connection_type="public",
    throughput_tier=throughput_tier,
    zones=zones,
    allow_deletion=True)
user_name = config.get("userName")
if user_name is None:
    user_name = "test-username"
user_pw = config.get("userPw")
if user_pw is None:
    user_pw = "password"
mechanism = config.get("mechanism")
if mechanism is None:
    mechanism = "scram-sha-256"
test_user = redpanda.User("test",
    name=user_name,
    password=user_pw,
    mechanism=mechanism,
    cluster_api_url=test_cluster.cluster_api_url)
topic_name = config.get("topicName")
if topic_name is None:
    topic_name = "test-topic"
partition_count = config.get_float("partitionCount")
if partition_count is None:
    partition_count = 3
replication_factor = config.get_float("replicationFactor")
if replication_factor is None:
    replication_factor = 3
test_topic = redpanda.Topic("test",
    name=topic_name,
    partition_count=partition_count,
    replication_factor=replication_factor,
    cluster_api_url=test_cluster.cluster_api_url,
    allow_deletion=True)
test_acl = redpanda.Acl("test",
    resource_type="TOPIC",
    resource_name_=test_topic.name,
    resource_pattern_type="LITERAL",
    principal=test_user.name.apply(lambda name: f"User:{name}"),
    host="*",
    operation="READ",
    permission_type="ALLOW",
    cluster_api_url=test_cluster.cluster_api_url)
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Redpanda = Pulumi.Redpanda;

return await Deployment.RunAsync(() =>
{
    var config = new Config();
    var resourceGroupName = config.Get("resourceGroupName") ?? "";
    var test = new Redpanda.ResourceGroup("test", new()
    {
        Name = resourceGroupName,
    });

    var networkName = config.Get("networkName") ?? "";
    var region = config.Get("region") ?? "us-central1";
    var cloudProvider = config.Get("cloudProvider") ?? "gcp";
    var testNetwork = new Redpanda.Network("test", new()
    {
        Name = networkName,
        ResourceGroupId = test.Id,
        CloudProvider = cloudProvider,
        Region = region,
        ClusterType = "dedicated",
        CidrBlock = "10.0.0.0/20",
    });

    var clusterName = config.Get("clusterName") ?? "";
    var zones = config.GetObject<dynamic>("zones") ?? new[]
    {
        "us-central1-a",
        "us-central1-b",
        "us-central1-c",
    };
    var throughputTier = config.Get("throughputTier") ?? "tier-1-gcp-um4g";
    var testCluster = new Redpanda.Cluster("test", new()
    {
        Name = clusterName,
        ResourceGroupId = test.Id,
        NetworkId = testNetwork.Id,
        CloudProvider = cloudProvider,
        Region = region,
        ClusterType = "dedicated",
        ConnectionType = "public",
        ThroughputTier = throughputTier,
        Zones = zones,
        AllowDeletion = true,
    });

    var userName = config.Get("userName") ?? "test-username";
    var userPw = config.Get("userPw") ?? "password";
    var mechanism = config.Get("mechanism") ?? "scram-sha-256";
    var testUser = new Redpanda.User("test", new()
    {
        Name = userName,
        Password = userPw,
        Mechanism = mechanism,
        ClusterApiUrl = testCluster.ClusterApiUrl,
    });

    var topicName = config.Get("topicName") ?? "test-topic";
    var partitionCount = config.GetDouble("partitionCount") ?? 3;
    var replicationFactor = config.GetDouble("replicationFactor") ?? 3;
    var testTopic = new Redpanda.Topic("test", new()
    {
        Name = topicName,
        PartitionCount = partitionCount,
        ReplicationFactor = replicationFactor,
        ClusterApiUrl = testCluster.ClusterApiUrl,
        AllowDeletion = true,
    });

    var testAcl = new Redpanda.Acl("test", new()
    {
        ResourceType = "TOPIC",
        ResourceName = testTopic.Name,
        ResourcePatternType = "LITERAL",
        Principal = testUser.Name.Apply(name => $"User:{name}"),
        Host = "*",
        Operation = "READ",
        PermissionType = "ALLOW",
        ClusterApiUrl = testCluster.ClusterApiUrl,
    });

});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go

```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/redpanda/redpanda"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")
		resourceGroupName := ""
		if param := cfg.Get("resourceGroupName"); param != "" {
			resourceGroupName = param
		}
		test, err := redpanda.NewResourceGroup(ctx, "test", &redpanda.ResourceGroupArgs{
			Name: pulumi.String(resourceGroupName),
		})
		if err != nil {
			return err
		}
		networkName := ""
		if param := cfg.Get("networkName"); param != "" {
			networkName = param
		}
		region := "us-central1"
		if param := cfg.Get("region"); param != "" {
			region = param
		}
		cloudProvider := "gcp"
		if param := cfg.Get("cloudProvider"); param != "" {
			cloudProvider = param
		}
		testNetwork, err := redpanda.NewNetwork(ctx, "test", &redpanda.NetworkArgs{
			Name:            pulumi.String(networkName),
			ResourceGroupId: test.ID(),
			CloudProvider:   pulumi.String(cloudProvider),
			Region:          pulumi.String(region),
			ClusterType:     pulumi.String("dedicated"),
			CidrBlock:       pulumi.String("10.0.0.0/20"),
		})
		if err != nil {
			return err
		}
		clusterName := ""
		if param := cfg.Get("clusterName"); param != "" {
			clusterName = param
		}
		zones := []string{
			"us-central1-a",
			"us-central1-b",
			"us-central1-c",
		}
		if param := cfg.GetObject("zones"); param != nil {
			zones = param
		}
		throughputTier := "tier-1-gcp-um4g"
		if param := cfg.Get("throughputTier"); param != "" {
			throughputTier = param
		}
		testCluster, err := redpanda.NewCluster(ctx, "test", &redpanda.ClusterArgs{
			Name:            pulumi.String(clusterName),
			ResourceGroupId: test.ID(),
			NetworkId:       testNetwork.ID(),
			CloudProvider:   pulumi.String(cloudProvider),
			Region:          pulumi.String(region),
			ClusterType:     pulumi.String("dedicated"),
			ConnectionType:  pulumi.String("public"),
			ThroughputTier:  pulumi.String(throughputTier),
			Zones:           pulumi.Any(zones),
			AllowDeletion:   pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		userName := "test-username"
		if param := cfg.Get("userName"); param != "" {
			userName = param
		}
		userPw := "password"
		if param := cfg.Get("userPw"); param != "" {
			userPw = param
		}
		mechanism := "scram-sha-256"
		if param := cfg.Get("mechanism"); param != "" {
			mechanism = param
		}
		testUser, err := redpanda.NewUser(ctx, "test", &redpanda.UserArgs{
			Name:          pulumi.String(userName),
			Password:      pulumi.String(userPw),
			Mechanism:     pulumi.String(mechanism),
			ClusterApiUrl: testCluster.ClusterApiUrl,
		})
		if err != nil {
			return err
		}
		topicName := "test-topic"
		if param := cfg.Get("topicName"); param != "" {
			topicName = param
		}
		partitionCount := float64(3)
		if param := cfg.GetFloat64("partitionCount"); param != 0 {
			partitionCount = param
		}
		replicationFactor := float64(3)
		if param := cfg.GetFloat64("replicationFactor"); param != 0 {
			replicationFactor = param
		}
		testTopic, err := redpanda.NewTopic(ctx, "test", &redpanda.TopicArgs{
			Name:              pulumi.String(topicName),
			PartitionCount:    pulumi.Float64(partitionCount),
			ReplicationFactor: pulumi.Float64(replicationFactor),
			ClusterApiUrl:     testCluster.ClusterApiUrl,
			AllowDeletion:     pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		_, err = redpanda.NewAcl(ctx, "test", &redpanda.AclArgs{
			ResourceType:        pulumi.String("TOPIC"),
			ResourceName:        testTopic.Name,
			ResourcePatternType: pulumi.String("LITERAL"),
			Principal: testUser.Name.ApplyT(func(name string) (string, error) {
				return fmt.Sprintf("User:%v", name), nil
			}).(pulumi.StringOutput),
			Host:           pulumi.String("*"),
			Operation:      pulumi.String("READ"),
			PermissionType: pulumi.String("ALLOW"),
			ClusterApiUrl:  testCluster.ClusterApiUrl,
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

```
```yaml
configuration:
  clusterName:
    type: string
    default: ""
  resourceGroupName:
    type: string
    default: ""
  networkName:
    type: string
    default: ""
  region:
    type: string
    default: us-central1
  zones:
    type: dynamic
    default:
      - us-central1-a
      - us-central1-b
      - us-central1-c
  cloudProvider:
    type: string
    default: gcp
  throughputTier:
    type: string
    default: tier-1-gcp-um4g
  userName:
    type: string
    default: test-username
  userPw:
    type: string
    default: password
  mechanism:
    type: string
    default: scram-sha-256
  topicName:
    type: string
    default: test-topic
  partitionCount:
    type: number
    default: 3
  replicationFactor:
    type: number
    default: 3
resources:
  test:
    type: redpanda:ResourceGroup
    properties:
      name: ${resourceGroupName}
  testNetwork:
    type: redpanda:Network
    name: test
    properties:
      name: ${networkName}
      resourceGroupId: ${test.id}
      cloudProvider: ${cloudProvider}
      region: ${region}
      clusterType: dedicated
      cidrBlock: 10.0.0.0/20
  testCluster:
    type: redpanda:Cluster
    name: test
    properties:
      name: ${clusterName}
      resourceGroupId: ${test.id}
      networkId: ${testNetwork.id}
      cloudProvider: ${cloudProvider}
      region: ${region}
      clusterType: dedicated
      connectionType: public
      throughputTier: ${throughputTier}
      zones: ${zones}
      allowDeletion: true ## This is a reference for GCP tags
      #   #   tags = {
      #   #     "key" = "value"
      #   #   }
      #   ## This is a reference for GCP Private Service Connect
      #   #   gcp_private_service_connect = {
      #   #     enabled               = true
      #   #     global_access_enabled = true
      #   #     consumer_accept_list = [
      #   #       {
      #   #         source = "projects/123456789012"
      #   #       }
      #   #     ]
      #   #   }
  testUser:
    type: redpanda:User
    name: test
    properties:
      name: ${userName}
      password: ${userPw}
      mechanism: ${mechanism}
      clusterApiUrl: ${testCluster.clusterApiUrl}
  testTopic:
    type: redpanda:Topic
    name: test
    properties:
      name: ${topicName}
      partitionCount: ${partitionCount}
      replicationFactor: ${replicationFactor}
      clusterApiUrl: ${testCluster.clusterApiUrl}
      allowDeletion: true
  testAcl:
    type: redpanda:Acl
    name: test
    properties:
      resourceType: TOPIC
      resourceName: ${testTopic.name}
      resourcePatternType: LITERAL
      principal: User:${testUser.name}
      host: '*'
      operation: READ
      permissionType: ALLOW
      clusterApiUrl: ${testCluster.clusterApiUrl}
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.redpanda.ResourceGroup;
import com.pulumi.redpanda.ResourceGroupArgs;
import com.pulumi.redpanda.Network;
import com.pulumi.redpanda.NetworkArgs;
import com.pulumi.redpanda.Cluster;
import com.pulumi.redpanda.ClusterArgs;
import com.pulumi.redpanda.User;
import com.pulumi.redpanda.UserArgs;
import com.pulumi.redpanda.Topic;
import com.pulumi.redpanda.TopicArgs;
import com.pulumi.redpanda.Acl;
import com.pulumi.redpanda.AclArgs;
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
        final var config = ctx.config();
        final var resourceGroupName = config.get("resourceGroupName").orElse("");
        var test = new ResourceGroup("test", ResourceGroupArgs.builder()
            .name(resourceGroupName)
            .build());

        final var networkName = config.get("networkName").orElse("");
        final var region = config.get("region").orElse("us-central1");
        final var cloudProvider = config.get("cloudProvider").orElse("gcp");
        var testNetwork = new Network("testNetwork", NetworkArgs.builder()
            .name(networkName)
            .resourceGroupId(test.id())
            .cloudProvider(cloudProvider)
            .region(region)
            .clusterType("dedicated")
            .cidrBlock("10.0.0.0/20")
            .build());

        final var clusterName = config.get("clusterName").orElse("");
        final var zones = config.get("zones").orElse(
            "us-central1-a",
            "us-central1-b",
            "us-central1-c");
        final var throughputTier = config.get("throughputTier").orElse("tier-1-gcp-um4g");
        var testCluster = new Cluster("testCluster", ClusterArgs.builder()
            .name(clusterName)
            .resourceGroupId(test.id())
            .networkId(testNetwork.id())
            .cloudProvider(cloudProvider)
            .region(region)
            .clusterType("dedicated")
            .connectionType("public")
            .throughputTier(throughputTier)
            .zones(zones)
            .allowDeletion(true)
            .build());

        final var userName = config.get("userName").orElse("test-username");
        final var userPw = config.get("userPw").orElse("password");
        final var mechanism = config.get("mechanism").orElse("scram-sha-256");
        var testUser = new User("testUser", UserArgs.builder()
            .name(userName)
            .password(userPw)
            .mechanism(mechanism)
            .clusterApiUrl(testCluster.clusterApiUrl())
            .build());

        final var topicName = config.get("topicName").orElse("test-topic");
        final var partitionCount = config.get("partitionCount").orElse(3);
        final var replicationFactor = config.get("replicationFactor").orElse(3);
        var testTopic = new Topic("testTopic", TopicArgs.builder()
            .name(topicName)
            .partitionCount(partitionCount)
            .replicationFactor(replicationFactor)
            .clusterApiUrl(testCluster.clusterApiUrl())
            .allowDeletion(true)
            .build());

        var testAcl = new Acl("testAcl", AclArgs.builder()
            .resourceType("TOPIC")
            .resourceName(testTopic.name())
            .resourcePatternType("LITERAL")
            .principal(testUser.name().applyValue(name -> String.format("User:%s", name)))
            .host("*")
            .operation("READ")
            .permissionType("ALLOW")
            .clusterApiUrl(testCluster.clusterApiUrl())
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
### Example Usage of a function BYOC to manage users and ACLs

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as redpanda from "@pulumi/redpanda";

const config = new pulumi.Config();
const clusterId = config.get("clusterId") || "";
const test = redpanda.getCluster({
    id: clusterId,
});
const topicConfig = config.getObject("topicConfig") || {
    "cleanup.policy": "compact",
    "compression.type": "snappy",
    "flush.ms": 100,
};
const topicName = config.get("topicName") || "data-test-topic";
const partitionCount = config.getNumber("partitionCount") || 3;
const replicationFactor = config.getNumber("replicationFactor") || 3;
const testTopic = new redpanda.Topic("test", {
    name: topicName,
    partitionCount: partitionCount,
    replicationFactor: replicationFactor,
    clusterApiUrl: test.then(test => test.clusterApiUrl),
    allowDeletion: true,
    configuration: topicConfig,
});
const userName = config.get("userName") || "data-test-username";
const userPw = config.get("userPw") || "password";
const mechanism = config.get("mechanism") || "scram-sha-256";
const testUser = new redpanda.User("test", {
    name: userName,
    password: userPw,
    mechanism: mechanism,
    clusterApiUrl: test.then(test => test.clusterApiUrl),
});
const testAcl = new redpanda.Acl("test", {
    resourceType: "CLUSTER",
    resourceName: "kafka-cluster",
    resourcePatternType: "LITERAL",
    principal: pulumi.interpolate`User:${testUser.name}`,
    host: "*",
    operation: "ALTER",
    permissionType: "ALLOW",
    clusterApiUrl: test.then(test => test.clusterApiUrl),
});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python

```
```python
import pulumi
import pulumi_redpanda as redpanda

config = pulumi.Config()
cluster_id = config.get("clusterId")
if cluster_id is None:
    cluster_id = ""
test = redpanda.get_cluster(id=cluster_id)
topic_config = config.get_object("topicConfig")
if topic_config is None:
    topic_config = {
        "cleanup.policy": "compact",
        "compression.type": "snappy",
        "flush.ms": 100,
    }
topic_name = config.get("topicName")
if topic_name is None:
    topic_name = "data-test-topic"
partition_count = config.get_float("partitionCount")
if partition_count is None:
    partition_count = 3
replication_factor = config.get_float("replicationFactor")
if replication_factor is None:
    replication_factor = 3
test_topic = redpanda.Topic("test",
    name=topic_name,
    partition_count=partition_count,
    replication_factor=replication_factor,
    cluster_api_url=test.cluster_api_url,
    allow_deletion=True,
    configuration=topic_config)
user_name = config.get("userName")
if user_name is None:
    user_name = "data-test-username"
user_pw = config.get("userPw")
if user_pw is None:
    user_pw = "password"
mechanism = config.get("mechanism")
if mechanism is None:
    mechanism = "scram-sha-256"
test_user = redpanda.User("test",
    name=user_name,
    password=user_pw,
    mechanism=mechanism,
    cluster_api_url=test.cluster_api_url)
test_acl = redpanda.Acl("test",
    resource_type="CLUSTER",
    resource_name_="kafka-cluster",
    resource_pattern_type="LITERAL",
    principal=test_user.name.apply(lambda name: f"User:{name}"),
    host="*",
    operation="ALTER",
    permission_type="ALLOW",
    cluster_api_url=test.cluster_api_url)
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Redpanda = Pulumi.Redpanda;

return await Deployment.RunAsync(() =>
{
    var config = new Config();
    var clusterId = config.Get("clusterId") ?? "";
    var test = Redpanda.GetCluster.Invoke(new()
    {
        Id = clusterId,
    });

    var topicConfig = config.GetObject<dynamic>("topicConfig") ??
    {
        { "cleanup.policy", "compact" },
        { "compression.type", "snappy" },
        { "flush.ms", 100 },
    };
    var topicName = config.Get("topicName") ?? "data-test-topic";
    var partitionCount = config.GetDouble("partitionCount") ?? 3;
    var replicationFactor = config.GetDouble("replicationFactor") ?? 3;
    var testTopic = new Redpanda.Topic("test", new()
    {
        Name = topicName,
        PartitionCount = partitionCount,
        ReplicationFactor = replicationFactor,
        ClusterApiUrl = test.Apply(getClusterResult => getClusterResult.ClusterApiUrl),
        AllowDeletion = true,
        Configuration = topicConfig,
    });

    var userName = config.Get("userName") ?? "data-test-username";
    var userPw = config.Get("userPw") ?? "password";
    var mechanism = config.Get("mechanism") ?? "scram-sha-256";
    var testUser = new Redpanda.User("test", new()
    {
        Name = userName,
        Password = userPw,
        Mechanism = mechanism,
        ClusterApiUrl = test.Apply(getClusterResult => getClusterResult.ClusterApiUrl),
    });

    var testAcl = new Redpanda.Acl("test", new()
    {
        ResourceType = "CLUSTER",
        ResourceName = "kafka-cluster",
        ResourcePatternType = "LITERAL",
        Principal = testUser.Name.Apply(name => $"User:{name}"),
        Host = "*",
        Operation = "ALTER",
        PermissionType = "ALLOW",
        ClusterApiUrl = test.Apply(getClusterResult => getClusterResult.ClusterApiUrl),
    });

});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go

```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/redpanda/redpanda"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")
		clusterId := ""
		if param := cfg.Get("clusterId"); param != "" {
			clusterId = param
		}
		test, err := redpanda.LookupCluster(ctx, &redpanda.LookupClusterArgs{
			Id: clusterId,
		}, nil)
		if err != nil {
			return err
		}
		topicConfig := map[string]interface{}{
			"cleanup.policy":   "compact",
			"compression.type": "snappy",
			"flush.ms":         100,
		}
		if param := cfg.GetObject("topicConfig"); param != nil {
			topicConfig = param
		}
		topicName := "data-test-topic"
		if param := cfg.Get("topicName"); param != "" {
			topicName = param
		}
		partitionCount := float64(3)
		if param := cfg.GetFloat64("partitionCount"); param != 0 {
			partitionCount = param
		}
		replicationFactor := float64(3)
		if param := cfg.GetFloat64("replicationFactor"); param != 0 {
			replicationFactor = param
		}
		_, err = redpanda.NewTopic(ctx, "test", &redpanda.TopicArgs{
			Name:              pulumi.String(topicName),
			PartitionCount:    pulumi.Float64(partitionCount),
			ReplicationFactor: pulumi.Float64(replicationFactor),
			ClusterApiUrl:     pulumi.String(test.ClusterApiUrl),
			AllowDeletion:     pulumi.Bool(true),
			Configuration:     pulumi.Any(topicConfig),
		})
		if err != nil {
			return err
		}
		userName := "data-test-username"
		if param := cfg.Get("userName"); param != "" {
			userName = param
		}
		userPw := "password"
		if param := cfg.Get("userPw"); param != "" {
			userPw = param
		}
		mechanism := "scram-sha-256"
		if param := cfg.Get("mechanism"); param != "" {
			mechanism = param
		}
		testUser, err := redpanda.NewUser(ctx, "test", &redpanda.UserArgs{
			Name:          pulumi.String(userName),
			Password:      pulumi.String(userPw),
			Mechanism:     pulumi.String(mechanism),
			ClusterApiUrl: pulumi.String(test.ClusterApiUrl),
		})
		if err != nil {
			return err
		}
		_, err = redpanda.NewAcl(ctx, "test", &redpanda.AclArgs{
			ResourceType:        pulumi.String("CLUSTER"),
			ResourceName:        pulumi.String("kafka-cluster"),
			ResourcePatternType: pulumi.String("LITERAL"),
			Principal: testUser.Name.ApplyT(func(name string) (string, error) {
				return fmt.Sprintf("User:%v", name), nil
			}).(pulumi.StringOutput),
			Host:           pulumi.String("*"),
			Operation:      pulumi.String("ALTER"),
			PermissionType: pulumi.String("ALLOW"),
			ClusterApiUrl:  pulumi.String(test.ClusterApiUrl),
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

```
```yaml
configuration:
  clusterId:
    type: string
    default: ""
  topicConfig:
    type: dynamic
    default:
      cleanup.policy: compact
      compression.type: snappy
      flush.ms: 100
  userName:
    type: string
    default: data-test-username
  userPw:
    type: string
    default: password
  mechanism:
    type: string
    default: scram-sha-256
  topicName:
    type: string
    default: data-test-topic
  partitionCount:
    type: number
    default: 3
  replicationFactor:
    type: number
    default: 3
resources:
  testTopic:
    type: redpanda:Topic
    name: test
    properties:
      name: ${topicName}
      partitionCount: ${partitionCount}
      replicationFactor: ${replicationFactor}
      clusterApiUrl: ${test.clusterApiUrl}
      allowDeletion: true
      configuration: ${topicConfig}
  testUser:
    type: redpanda:User
    name: test
    properties:
      name: ${userName}
      password: ${userPw}
      mechanism: ${mechanism}
      clusterApiUrl: ${test.clusterApiUrl}
  testAcl:
    type: redpanda:Acl
    name: test
    properties:
      resourceType: CLUSTER
      resourceName: kafka-cluster
      resourcePatternType: LITERAL
      principal: User:${testUser.name}
      host: '*'
      operation: ALTER
      permissionType: ALLOW
      clusterApiUrl: ${test.clusterApiUrl}
variables:
  test:
    fn::invoke:
      function: redpanda:getCluster
      arguments:
        id: ${clusterId}
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.redpanda.RedpandaFunctions;
import com.pulumi.redpanda.inputs.GetClusterArgs;
import com.pulumi.redpanda.Topic;
import com.pulumi.redpanda.TopicArgs;
import com.pulumi.redpanda.User;
import com.pulumi.redpanda.UserArgs;
import com.pulumi.redpanda.Acl;
import com.pulumi.redpanda.AclArgs;
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
        final var config = ctx.config();
        final var clusterId = config.get("clusterId").orElse("");
        final var test = RedpandaFunctions.getCluster(GetClusterArgs.builder()
            .id(clusterId)
            .build());

        final var topicConfig = config.get("topicConfig").orElse(%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference));
        final var topicName = config.get("topicName").orElse("data-test-topic");
        final var partitionCount = config.get("partitionCount").orElse(3);
        final var replicationFactor = config.get("replicationFactor").orElse(3);
        var testTopic = new Topic("testTopic", TopicArgs.builder()
            .name(topicName)
            .partitionCount(partitionCount)
            .replicationFactor(replicationFactor)
            .clusterApiUrl(test.applyValue(getClusterResult -> getClusterResult.clusterApiUrl()))
            .allowDeletion(true)
            .configuration(topicConfig)
            .build());

        final var userName = config.get("userName").orElse("data-test-username");
        final var userPw = config.get("userPw").orElse("password");
        final var mechanism = config.get("mechanism").orElse("scram-sha-256");
        var testUser = new User("testUser", UserArgs.builder()
            .name(userName)
            .password(userPw)
            .mechanism(mechanism)
            .clusterApiUrl(test.applyValue(getClusterResult -> getClusterResult.clusterApiUrl()))
            .build());

        var testAcl = new Acl("testAcl", AclArgs.builder()
            .resourceType("CLUSTER")
            .resourceName("kafka-cluster")
            .resourcePatternType("LITERAL")
            .principal(testUser.name().applyValue(name -> String.format("User:%s", name)))
            .host("*")
            .operation("ALTER")
            .permissionType("ALLOW")
            .clusterApiUrl(test.applyValue(getClusterResult -> getClusterResult.clusterApiUrl()))
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
### Example Usage to create a serverless cluster

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as redpanda from "@pulumi/redpanda";

const config = new pulumi.Config();
const resourceGroupName = config.get("resourceGroupName") || "testgroup";
const test = new redpanda.ResourceGroup("test", {name: resourceGroupName});
const clusterName = config.get("clusterName") || "testname";
const region = config.get("region") || "eu-west-1";
const testServerlessCluster = new redpanda.ServerlessCluster("test", {
    name: clusterName,
    resourceGroupId: test.id,
    serverlessRegion: region,
});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python

```
```python
import pulumi
import pulumi_redpanda as redpanda

config = pulumi.Config()
resource_group_name = config.get("resourceGroupName")
if resource_group_name is None:
    resource_group_name = "testgroup"
test = redpanda.ResourceGroup("test", name=resource_group_name)
cluster_name = config.get("clusterName")
if cluster_name is None:
    cluster_name = "testname"
region = config.get("region")
if region is None:
    region = "eu-west-1"
test_serverless_cluster = redpanda.ServerlessCluster("test",
    name=cluster_name,
    resource_group_id=test.id,
    serverless_region=region)
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Redpanda = Pulumi.Redpanda;

return await Deployment.RunAsync(() =>
{
    var config = new Config();
    var resourceGroupName = config.Get("resourceGroupName") ?? "testgroup";
    var test = new Redpanda.ResourceGroup("test", new()
    {
        Name = resourceGroupName,
    });

    var clusterName = config.Get("clusterName") ?? "testname";
    var region = config.Get("region") ?? "eu-west-1";
    var testServerlessCluster = new Redpanda.ServerlessCluster("test", new()
    {
        Name = clusterName,
        ResourceGroupId = test.Id,
        ServerlessRegion = region,
    });

});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go

```
```go
package main

import (
	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/redpanda/redpanda"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")
		resourceGroupName := "testgroup"
		if param := cfg.Get("resourceGroupName"); param != "" {
			resourceGroupName = param
		}
		test, err := redpanda.NewResourceGroup(ctx, "test", &redpanda.ResourceGroupArgs{
			Name: pulumi.String(resourceGroupName),
		})
		if err != nil {
			return err
		}
		clusterName := "testname"
		if param := cfg.Get("clusterName"); param != "" {
			clusterName = param
		}
		region := "eu-west-1"
		if param := cfg.Get("region"); param != "" {
			region = param
		}
		_, err = redpanda.NewServerlessCluster(ctx, "test", &redpanda.ServerlessClusterArgs{
			Name:             pulumi.String(clusterName),
			ResourceGroupId:  test.ID(),
			ServerlessRegion: pulumi.String(region),
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

```
```yaml
configuration:
  resourceGroupName:
    type: string
    default: testgroup
  clusterName:
    type: string
    default: testname
  region:
    type: string
    default: eu-west-1
resources:
  test:
    type: redpanda:ResourceGroup
    properties:
      name: ${resourceGroupName}
  testServerlessCluster:
    type: redpanda:ServerlessCluster
    name: test
    properties:
      name: ${clusterName}
      resourceGroupId: ${test.id}
      serverlessRegion: ${region}
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.redpanda.ResourceGroup;
import com.pulumi.redpanda.ResourceGroupArgs;
import com.pulumi.redpanda.ServerlessCluster;
import com.pulumi.redpanda.ServerlessClusterArgs;
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
        final var config = ctx.config();
        final var resourceGroupName = config.get("resourceGroupName").orElse("testgroup");
        var test = new ResourceGroup("test", ResourceGroupArgs.builder()
            .name(resourceGroupName)
            .build());

        final var clusterName = config.get("clusterName").orElse("testname");
        final var region = config.get("region").orElse("eu-west-1");
        var testServerlessCluster = new ServerlessCluster("testServerlessCluster", ServerlessClusterArgs.builder()
            .name(clusterName)
            .resourceGroupId(test.id())
            .serverlessRegion(region)
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}