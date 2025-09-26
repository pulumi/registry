---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/redpanda-data/redpanda/1.3.0/index.md
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
    timeouts: {
        create: "20m",
        "delete": "20m",
    },
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
    clusterConfiguration: {
        customPropertiesJson: JSON.stringify({
            schema_registry_enable_authorization: true,
        }),
    },
    tags: {
        key: "value",
    },
    timeouts: {
        create: "90m",
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
const clusterAdmin = new redpanda.Acl("cluster_admin", {
    resourceType: "CLUSTER",
    resourceName: "kafka-cluster",
    resourcePatternType: "LITERAL",
    principal: pulumi.interpolate`User:${testUser.name}`,
    host: "*",
    operation: "ALL",
    permissionType: "ALLOW",
    clusterApiUrl: testCluster.clusterApiUrl,
    allowDeletion: true,
});
const schemaRegistryAdmin = new redpanda.Acl("schema_registry_admin", {
    resourceType: "CLUSTER",
    resourceName: "kafka-cluster",
    resourcePatternType: "LITERAL",
    principal: pulumi.interpolate`User:${testUser.name}`,
    host: "*",
    operation: "ALTER",
    permissionType: "ALLOW",
    clusterApiUrl: testCluster.clusterApiUrl,
    allowDeletion: true,
});
const clusterAction = new redpanda.Acl("cluster_action", {
    resourceType: "CLUSTER",
    resourceName: "kafka-cluster",
    resourcePatternType: "LITERAL",
    principal: pulumi.interpolate`User:${testUser.name}`,
    host: "*",
    operation: "CLUSTER_ACTION",
    permissionType: "ALLOW",
    clusterApiUrl: testCluster.clusterApiUrl,
    allowDeletion: true,
});
const topicAccess = new redpanda.Acl("topic_access", {
    resourceType: "TOPIC",
    resourceName: topicName,
    resourcePatternType: "LITERAL",
    principal: pulumi.interpolate`User:${testUser.name}`,
    host: "*",
    operation: "ALL",
    permissionType: "ALLOW",
    clusterApiUrl: testCluster.clusterApiUrl,
    allowDeletion: true,
});
const allTestTopic = new redpanda.SchemaRegistryAcl("all_test_topic", {
    clusterId: testCluster.id,
    principal: pulumi.interpolate`User:${testUser.name}`,
    resourceType: "SUBJECT",
    resourceName: `${topicName}-`,
    patternType: "PREFIXED",
    host: "*",
    operation: "ALL",
    permission: "ALLOW",
    username: testUser.name,
    password: userPw,
    allowDeletion: true,
}, {
    dependsOn: [schemaRegistryAdmin],
});
const describeRegistry = new redpanda.SchemaRegistryAcl("describe_registry", {
    clusterId: testCluster.id,
    principal: pulumi.interpolate`User:${testUser.name}`,
    resourceType: "REGISTRY",
    resourceName: "*",
    patternType: "LITERAL",
    host: "*",
    operation: "DESCRIBE",
    permission: "ALLOW",
    username: testUser.name,
    password: userPw,
    allowDeletion: true,
}, {
    dependsOn: [schemaRegistryAdmin],
});
const alterConfigsRegistry = new redpanda.SchemaRegistryAcl("alter_configs_registry", {
    clusterId: testCluster.id,
    principal: pulumi.interpolate`User:${testUser.name}`,
    resourceType: "REGISTRY",
    resourceName: "*",
    patternType: "LITERAL",
    host: "*",
    operation: "ALTER_CONFIGS",
    permission: "ALLOW",
    username: testUser.name,
    password: userPw,
    allowDeletion: true,
}, {
    dependsOn: [schemaRegistryAdmin],
});
// The type of schema (AVRO, JSON, PROTOBUF)
const schemaType = config.get("schemaType") || "AVRO";
// The AVRO schema definition for user data
const userSchemaDefinition = config.get("userSchemaDefinition") || `{
  "type": "record",
  "name": "User",
  "fields": [
    {
      "name": "id",
      "type": "int"
    },
    {
      "name": "name",
      "type": "string"
    },
    {
      "name": "email",
      "type": "string"
    },
    {
      "name": "created_at",
      "type": "long",
      "logicalType": "timestamp-millis"
    }
  ]
}
`;
const userSchema = new redpanda.Schema("user_schema", {
    clusterId: testCluster.id,
    subject: `${topicName}-value`,
    schemaType: schemaType,
    schema: userSchemaDefinition,
    username: testUser.name,
    password: userPw,
    allowDeletion: true,
}, {
    dependsOn: [
        clusterAdmin,
        schemaRegistryAdmin,
        clusterAction,
        topicAccess,
        allTestTopic,
        describeRegistry,
        alterConfigsRegistry,
    ],
});
// The AVRO schema definition for user events that references the User schema
const userEventSchemaDefinition = config.get("userEventSchemaDefinition") || `{
  "type": "record",
  "name": "UserEvent",
  "fields": [
    {
      "name": "event_id",
      "type": "string"
    },
    {
      "name": "event_type",
      "type": {
        "type": "enum",
        "name": "EventType",
        "symbols": ["CREATED", "UPDATED", "DELETED"]
      }
    },
    {
      "name": "user",
      "type": "User"
    },
    {
      "name": "timestamp",
      "type": "long",
      "logicalType": "timestamp-millis"
    },
    {
      "name": "metadata",
      "type": ["null", {
        "type": "map",
        "values": "string"
      }],
      "default": null
    }
  ]
}
`;
const userEventSchema = new redpanda.Schema("user_event_schema", {
    clusterId: testCluster.id,
    subject: `${topicName}-events-value`,
    schemaType: schemaType,
    schema: userEventSchemaDefinition,
    username: testUser.name,
    password: userPw,
    allowDeletion: true,
    references: [{
        name: "User",
        subject: userSchema.subject,
        version: userSchema.version,
    }],
}, {
    dependsOn: [
        clusterAdmin,
        schemaRegistryAdmin,
        clusterAction,
        topicAccess,
        allTestTopic,
        describeRegistry,
        alterConfigsRegistry,
    ],
});
// The AVRO schema definition for product data with strict compatibility
const productSchemaDefinition = config.get("productSchemaDefinition") || `{
  "type": "record",
  "name": "Product",
  "fields": [
    {
      "name": "id",
      "type": "string"
    },
    {
      "name": "name",
      "type": "string"
    },
    {
      "name": "price",
      "type": {
        "type": "bytes",
        "logicalType": "decimal",
        "precision": 10,
        "scale": 2
      }
    },
    {
      "name": "category",
      "type": {
        "type": "enum",
        "name": "Category",
        "symbols": ["ELECTRONICS", "CLOTHING", "BOOKS", "HOME"]
      }
    },
    {
      "name": "description",
      "type": ["null", "string"],
      "default": null
    },
    {
      "name": "created_at",
      "type": "long",
      "logicalType": "timestamp-millis"
    }
  ]
}
`;
// The compatibility level for schema evolution (BACKWARD, BACKWARD_TRANSITIVE, FORWARD, FORWARD_TRANSITIVE, FULL, FULL_TRANSITIVE, NONE)
const compatibilityLevel = config.get("compatibilityLevel") || "FULL";
const productSchema = new redpanda.Schema("product_schema", {
    clusterId: testCluster.id,
    subject: `${topicName}-product-value`,
    schemaType: schemaType,
    schema: productSchemaDefinition,
    compatibility: compatibilityLevel,
    username: testUser.name,
    password: userPw,
    allowDeletion: true,
}, {
    dependsOn: [
        clusterAdmin,
        schemaRegistryAdmin,
        clusterAction,
        topicAccess,
        allTestTopic,
        describeRegistry,
        alterConfigsRegistry,
    ],
});
const readProduct = new redpanda.SchemaRegistryAcl("read_product", {
    clusterId: testCluster.id,
    principal: pulumi.interpolate`User:${testUser.name}`,
    resourceType: "SUBJECT",
    resourceName: "product-",
    patternType: "PREFIXED",
    host: "*",
    operation: "READ",
    permission: "ALLOW",
    username: testUser.name,
    password: userPw,
    allowDeletion: true,
}, {
    dependsOn: [schemaRegistryAdmin],
});
const writeOrders = new redpanda.SchemaRegistryAcl("write_orders", {
    clusterId: testCluster.id,
    principal: pulumi.interpolate`User:${testUser.name}`,
    resourceType: "SUBJECT",
    resourceName: "orders-value",
    patternType: "LITERAL",
    host: "*",
    operation: "WRITE",
    permission: "ALLOW",
    username: testUser.name,
    password: userPw,
    allowDeletion: true,
}, {
    dependsOn: [schemaRegistryAdmin],
});
export const userSchemaInfo = {
    id: userSchema.schemaId,
    subject: userSchema.subject,
    version: userSchema.version,
    type: userSchema.schemaType,
};
export const userEventSchemaInfo = {
    id: userEventSchema.schemaId,
    subject: userEventSchema.subject,
    version: userEventSchema.version,
    type: userEventSchema.schemaType,
    references: userEventSchema.references,
};
export const productSchemaInfo = {
    id: productSchema.schemaId,
    subject: productSchema.subject,
    version: productSchema.version,
    type: productSchema.schemaType,
    compatibility: productSchema.compatibility,
};
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
import json
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
    cidr_block="10.0.0.0/20",
    timeouts={
        "create": "20m",
        "delete": "20m",
    })
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
    cluster_configuration={
        "custom_properties_json": json.dumps({
            "schema_registry_enable_authorization": True,
        }),
    },
    tags={
        "key": "value",
    },
    timeouts={
        "create": "90m",
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
cluster_admin = redpanda.Acl("cluster_admin",
    resource_type="CLUSTER",
    resource_name_="kafka-cluster",
    resource_pattern_type="LITERAL",
    principal=test_user.name.apply(lambda name: f"User:{name}"),
    host="*",
    operation="ALL",
    permission_type="ALLOW",
    cluster_api_url=test_cluster.cluster_api_url,
    allow_deletion=True)
schema_registry_admin = redpanda.Acl("schema_registry_admin",
    resource_type="CLUSTER",
    resource_name_="kafka-cluster",
    resource_pattern_type="LITERAL",
    principal=test_user.name.apply(lambda name: f"User:{name}"),
    host="*",
    operation="ALTER",
    permission_type="ALLOW",
    cluster_api_url=test_cluster.cluster_api_url,
    allow_deletion=True)
cluster_action = redpanda.Acl("cluster_action",
    resource_type="CLUSTER",
    resource_name_="kafka-cluster",
    resource_pattern_type="LITERAL",
    principal=test_user.name.apply(lambda name: f"User:{name}"),
    host="*",
    operation="CLUSTER_ACTION",
    permission_type="ALLOW",
    cluster_api_url=test_cluster.cluster_api_url,
    allow_deletion=True)
topic_access = redpanda.Acl("topic_access",
    resource_type="TOPIC",
    resource_name_=topic_name,
    resource_pattern_type="LITERAL",
    principal=test_user.name.apply(lambda name: f"User:{name}"),
    host="*",
    operation="ALL",
    permission_type="ALLOW",
    cluster_api_url=test_cluster.cluster_api_url,
    allow_deletion=True)
all_test_topic = redpanda.SchemaRegistryAcl("all_test_topic",
    cluster_id=test_cluster.id,
    principal=test_user.name.apply(lambda name: f"User:{name}"),
    resource_type="SUBJECT",
    resource_name_=f"{topic_name}-",
    pattern_type="PREFIXED",
    host="*",
    operation="ALL",
    permission="ALLOW",
    username=test_user.name,
    password=user_pw,
    allow_deletion=True,
    opts = pulumi.ResourceOptions(depends_on=[schema_registry_admin]))
describe_registry = redpanda.SchemaRegistryAcl("describe_registry",
    cluster_id=test_cluster.id,
    principal=test_user.name.apply(lambda name: f"User:{name}"),
    resource_type="REGISTRY",
    resource_name_="*",
    pattern_type="LITERAL",
    host="*",
    operation="DESCRIBE",
    permission="ALLOW",
    username=test_user.name,
    password=user_pw,
    allow_deletion=True,
    opts = pulumi.ResourceOptions(depends_on=[schema_registry_admin]))
alter_configs_registry = redpanda.SchemaRegistryAcl("alter_configs_registry",
    cluster_id=test_cluster.id,
    principal=test_user.name.apply(lambda name: f"User:{name}"),
    resource_type="REGISTRY",
    resource_name_="*",
    pattern_type="LITERAL",
    host="*",
    operation="ALTER_CONFIGS",
    permission="ALLOW",
    username=test_user.name,
    password=user_pw,
    allow_deletion=True,
    opts = pulumi.ResourceOptions(depends_on=[schema_registry_admin]))
# The type of schema (AVRO, JSON, PROTOBUF)
schema_type = config.get("schemaType")
if schema_type is None:
    schema_type = "AVRO"
# The AVRO schema definition for user data
user_schema_definition = config.get("userSchemaDefinition")
if user_schema_definition is None:
    user_schema_definition = """{
  "type": "record",
  "name": "User",
  "fields": [
    {
      "name": "id",
      "type": "int"
    },
    {
      "name": "name",
      "type": "string"
    },
    {
      "name": "email",
      "type": "string"
    },
    {
      "name": "created_at",
      "type": "long",
      "logicalType": "timestamp-millis"
    }
  ]
}
"""
user_schema = redpanda.Schema("user_schema",
    cluster_id=test_cluster.id,
    subject=f"{topic_name}-value",
    schema_type=schema_type,
    schema=user_schema_definition,
    username=test_user.name,
    password=user_pw,
    allow_deletion=True,
    opts = pulumi.ResourceOptions(depends_on=[
            cluster_admin,
            schema_registry_admin,
            cluster_action,
            topic_access,
            all_test_topic,
            describe_registry,
            alter_configs_registry,
        ]))
# The AVRO schema definition for user events that references the User schema
user_event_schema_definition = config.get("userEventSchemaDefinition")
if user_event_schema_definition is None:
    user_event_schema_definition = """{
  "type": "record",
  "name": "UserEvent",
  "fields": [
    {
      "name": "event_id",
      "type": "string"
    },
    {
      "name": "event_type",
      "type": {
        "type": "enum",
        "name": "EventType",
        "symbols": ["CREATED", "UPDATED", "DELETED"]
      }
    },
    {
      "name": "user",
      "type": "User"
    },
    {
      "name": "timestamp",
      "type": "long",
      "logicalType": "timestamp-millis"
    },
    {
      "name": "metadata",
      "type": ["null", {
        "type": "map",
        "values": "string"
      }],
      "default": null
    }
  ]
}
"""
user_event_schema = redpanda.Schema("user_event_schema",
    cluster_id=test_cluster.id,
    subject=f"{topic_name}-events-value",
    schema_type=schema_type,
    schema=user_event_schema_definition,
    username=test_user.name,
    password=user_pw,
    allow_deletion=True,
    references=[{
        "name": "User",
        "subject": user_schema.subject,
        "version": user_schema.version,
    }],
    opts = pulumi.ResourceOptions(depends_on=[
            cluster_admin,
            schema_registry_admin,
            cluster_action,
            topic_access,
            all_test_topic,
            describe_registry,
            alter_configs_registry,
        ]))
# The AVRO schema definition for product data with strict compatibility
product_schema_definition = config.get("productSchemaDefinition")
if product_schema_definition is None:
    product_schema_definition = """{
  "type": "record",
  "name": "Product",
  "fields": [
    {
      "name": "id",
      "type": "string"
    },
    {
      "name": "name",
      "type": "string"
    },
    {
      "name": "price",
      "type": {
        "type": "bytes",
        "logicalType": "decimal",
        "precision": 10,
        "scale": 2
      }
    },
    {
      "name": "category",
      "type": {
        "type": "enum",
        "name": "Category",
        "symbols": ["ELECTRONICS", "CLOTHING", "BOOKS", "HOME"]
      }
    },
    {
      "name": "description",
      "type": ["null", "string"],
      "default": null
    },
    {
      "name": "created_at",
      "type": "long",
      "logicalType": "timestamp-millis"
    }
  ]
}
"""
# The compatibility level for schema evolution (BACKWARD, BACKWARD_TRANSITIVE, FORWARD, FORWARD_TRANSITIVE, FULL, FULL_TRANSITIVE, NONE)
compatibility_level = config.get("compatibilityLevel")
if compatibility_level is None:
    compatibility_level = "FULL"
product_schema = redpanda.Schema("product_schema",
    cluster_id=test_cluster.id,
    subject=f"{topic_name}-product-value",
    schema_type=schema_type,
    schema=product_schema_definition,
    compatibility=compatibility_level,
    username=test_user.name,
    password=user_pw,
    allow_deletion=True,
    opts = pulumi.ResourceOptions(depends_on=[
            cluster_admin,
            schema_registry_admin,
            cluster_action,
            topic_access,
            all_test_topic,
            describe_registry,
            alter_configs_registry,
        ]))
read_product = redpanda.SchemaRegistryAcl("read_product",
    cluster_id=test_cluster.id,
    principal=test_user.name.apply(lambda name: f"User:{name}"),
    resource_type="SUBJECT",
    resource_name_="product-",
    pattern_type="PREFIXED",
    host="*",
    operation="READ",
    permission="ALLOW",
    username=test_user.name,
    password=user_pw,
    allow_deletion=True,
    opts = pulumi.ResourceOptions(depends_on=[schema_registry_admin]))
write_orders = redpanda.SchemaRegistryAcl("write_orders",
    cluster_id=test_cluster.id,
    principal=test_user.name.apply(lambda name: f"User:{name}"),
    resource_type="SUBJECT",
    resource_name_="orders-value",
    pattern_type="LITERAL",
    host="*",
    operation="WRITE",
    permission="ALLOW",
    username=test_user.name,
    password=user_pw,
    allow_deletion=True,
    opts = pulumi.ResourceOptions(depends_on=[schema_registry_admin]))
pulumi.export("userSchemaInfo", {
    "id": user_schema.schema_id,
    "subject": user_schema.subject,
    "version": user_schema.version,
    "type": user_schema.schema_type,
})
pulumi.export("userEventSchemaInfo", {
    "id": user_event_schema.schema_id,
    "subject": user_event_schema.subject,
    "version": user_event_schema.version,
    "type": user_event_schema.schema_type,
    "references": user_event_schema.references,
})
pulumi.export("productSchemaInfo", {
    "id": product_schema.schema_id,
    "subject": product_schema.subject,
    "version": product_schema.version,
    "type": product_schema.schema_type,
    "compatibility": product_schema.compatibility,
})
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
using System.Text.Json;
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
        Timeouts = new Redpanda.Inputs.NetworkTimeoutsArgs
        {
            Create = "20m",
            Delete = "20m",
        },
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
        ClusterConfiguration = new Redpanda.Inputs.ClusterClusterConfigurationArgs
        {
            CustomPropertiesJson = JsonSerializer.Serialize(new Dictionary<string, object?>
            {
                ["schema_registry_enable_authorization"] = true,
            }),
        },
        Tags =
        {
            { "key", "value" },
        },
        Timeouts = new Redpanda.Inputs.ClusterTimeoutsArgs
        {
            Create = "90m",
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

    var clusterAdmin = new Redpanda.Acl("cluster_admin", new()
    {
        ResourceType = "CLUSTER",
        ResourceName = "kafka-cluster",
        ResourcePatternType = "LITERAL",
        Principal = testUser.Name.Apply(name => $"User:{name}"),
        Host = "*",
        Operation = "ALL",
        PermissionType = "ALLOW",
        ClusterApiUrl = testCluster.ClusterApiUrl,
        AllowDeletion = true,
    });

    var schemaRegistryAdmin = new Redpanda.Acl("schema_registry_admin", new()
    {
        ResourceType = "CLUSTER",
        ResourceName = "kafka-cluster",
        ResourcePatternType = "LITERAL",
        Principal = testUser.Name.Apply(name => $"User:{name}"),
        Host = "*",
        Operation = "ALTER",
        PermissionType = "ALLOW",
        ClusterApiUrl = testCluster.ClusterApiUrl,
        AllowDeletion = true,
    });

    var clusterAction = new Redpanda.Acl("cluster_action", new()
    {
        ResourceType = "CLUSTER",
        ResourceName = "kafka-cluster",
        ResourcePatternType = "LITERAL",
        Principal = testUser.Name.Apply(name => $"User:{name}"),
        Host = "*",
        Operation = "CLUSTER_ACTION",
        PermissionType = "ALLOW",
        ClusterApiUrl = testCluster.ClusterApiUrl,
        AllowDeletion = true,
    });

    var topicAccess = new Redpanda.Acl("topic_access", new()
    {
        ResourceType = "TOPIC",
        ResourceName = topicName,
        ResourcePatternType = "LITERAL",
        Principal = testUser.Name.Apply(name => $"User:{name}"),
        Host = "*",
        Operation = "ALL",
        PermissionType = "ALLOW",
        ClusterApiUrl = testCluster.ClusterApiUrl,
        AllowDeletion = true,
    });

    var allTestTopic = new Redpanda.SchemaRegistryAcl("all_test_topic", new()
    {
        ClusterId = testCluster.Id,
        Principal = testUser.Name.Apply(name => $"User:{name}"),
        ResourceType = "SUBJECT",
        ResourceName = $"{topicName}-",
        PatternType = "PREFIXED",
        Host = "*",
        Operation = "ALL",
        Permission = "ALLOW",
        Username = testUser.Name,
        Password = userPw,
        AllowDeletion = true,
    }, new CustomResourceOptions
    {
        DependsOn =
        {
            schemaRegistryAdmin,
        },
    });

    var describeRegistry = new Redpanda.SchemaRegistryAcl("describe_registry", new()
    {
        ClusterId = testCluster.Id,
        Principal = testUser.Name.Apply(name => $"User:{name}"),
        ResourceType = "REGISTRY",
        ResourceName = "*",
        PatternType = "LITERAL",
        Host = "*",
        Operation = "DESCRIBE",
        Permission = "ALLOW",
        Username = testUser.Name,
        Password = userPw,
        AllowDeletion = true,
    }, new CustomResourceOptions
    {
        DependsOn =
        {
            schemaRegistryAdmin,
        },
    });

    var alterConfigsRegistry = new Redpanda.SchemaRegistryAcl("alter_configs_registry", new()
    {
        ClusterId = testCluster.Id,
        Principal = testUser.Name.Apply(name => $"User:{name}"),
        ResourceType = "REGISTRY",
        ResourceName = "*",
        PatternType = "LITERAL",
        Host = "*",
        Operation = "ALTER_CONFIGS",
        Permission = "ALLOW",
        Username = testUser.Name,
        Password = userPw,
        AllowDeletion = true,
    }, new CustomResourceOptions
    {
        DependsOn =
        {
            schemaRegistryAdmin,
        },
    });

    // The type of schema (AVRO, JSON, PROTOBUF)
    var schemaType = config.Get("schemaType") ?? "AVRO";
    // The AVRO schema definition for user data
    var userSchemaDefinition = config.Get("userSchemaDefinition") ?? @"{
  ""type"": ""record"",
  ""name"": ""User"",
  ""fields"": [
    {
      ""name"": ""id"",
      ""type"": ""int""
    },
    {
      ""name"": ""name"",
      ""type"": ""string""
    },
    {
      ""name"": ""email"",
      ""type"": ""string""
    },
    {
      ""name"": ""created_at"",
      ""type"": ""long"",
      ""logicalType"": ""timestamp-millis""
    }
  ]
}
";
    var userSchema = new Redpanda.Schema("user_schema", new()
    {
        ClusterId = testCluster.Id,
        Subject = $"{topicName}-value",
        SchemaType = schemaType,
        Schema = userSchemaDefinition,
        Username = testUser.Name,
        Password = userPw,
        AllowDeletion = true,
    }, new CustomResourceOptions
    {
        DependsOn =
        {
            clusterAdmin,
            schemaRegistryAdmin,
            clusterAction,
            topicAccess,
            allTestTopic,
            describeRegistry,
            alterConfigsRegistry,
        },
    });

    // The AVRO schema definition for user events that references the User schema
    var userEventSchemaDefinition = config.Get("userEventSchemaDefinition") ?? @"{
  ""type"": ""record"",
  ""name"": ""UserEvent"",
  ""fields"": [
    {
      ""name"": ""event_id"",
      ""type"": ""string""
    },
    {
      ""name"": ""event_type"",
      ""type"": {
        ""type"": ""enum"",
        ""name"": ""EventType"",
        ""symbols"": [""CREATED"", ""UPDATED"", ""DELETED""]
      }
    },
    {
      ""name"": ""user"",
      ""type"": ""User""
    },
    {
      ""name"": ""timestamp"",
      ""type"": ""long"",
      ""logicalType"": ""timestamp-millis""
    },
    {
      ""name"": ""metadata"",
      ""type"": [""null"", {
        ""type"": ""map"",
        ""values"": ""string""
      }],
      ""default"": null
    }
  ]
}
";
    var userEventSchema = new Redpanda.Schema("user_event_schema", new()
    {
        ClusterId = testCluster.Id,
        Subject = $"{topicName}-events-value",
        SchemaType = schemaType,
        Schema = userEventSchemaDefinition,
        Username = testUser.Name,
        Password = userPw,
        AllowDeletion = true,
        References = new[]
        {
            new Redpanda.Inputs.SchemaReferenceArgs
            {
                Name = "User",
                Subject = userSchema.Subject,
                Version = userSchema.Version,
            },
        },
    }, new CustomResourceOptions
    {
        DependsOn =
        {
            clusterAdmin,
            schemaRegistryAdmin,
            clusterAction,
            topicAccess,
            allTestTopic,
            describeRegistry,
            alterConfigsRegistry,
        },
    });

    // The AVRO schema definition for product data with strict compatibility
    var productSchemaDefinition = config.Get("productSchemaDefinition") ?? @"{
  ""type"": ""record"",
  ""name"": ""Product"",
  ""fields"": [
    {
      ""name"": ""id"",
      ""type"": ""string""
    },
    {
      ""name"": ""name"",
      ""type"": ""string""
    },
    {
      ""name"": ""price"",
      ""type"": {
        ""type"": ""bytes"",
        ""logicalType"": ""decimal"",
        ""precision"": 10,
        ""scale"": 2
      }
    },
    {
      ""name"": ""category"",
      ""type"": {
        ""type"": ""enum"",
        ""name"": ""Category"",
        ""symbols"": [""ELECTRONICS"", ""CLOTHING"", ""BOOKS"", ""HOME""]
      }
    },
    {
      ""name"": ""description"",
      ""type"": [""null"", ""string""],
      ""default"": null
    },
    {
      ""name"": ""created_at"",
      ""type"": ""long"",
      ""logicalType"": ""timestamp-millis""
    }
  ]
}
";
    // The compatibility level for schema evolution (BACKWARD, BACKWARD_TRANSITIVE, FORWARD, FORWARD_TRANSITIVE, FULL, FULL_TRANSITIVE, NONE)
    var compatibilityLevel = config.Get("compatibilityLevel") ?? "FULL";
    var productSchema = new Redpanda.Schema("product_schema", new()
    {
        ClusterId = testCluster.Id,
        Subject = $"{topicName}-product-value",
        SchemaType = schemaType,
        Schema = productSchemaDefinition,
        Compatibility = compatibilityLevel,
        Username = testUser.Name,
        Password = userPw,
        AllowDeletion = true,
    }, new CustomResourceOptions
    {
        DependsOn =
        {
            clusterAdmin,
            schemaRegistryAdmin,
            clusterAction,
            topicAccess,
            allTestTopic,
            describeRegistry,
            alterConfigsRegistry,
        },
    });

    var readProduct = new Redpanda.SchemaRegistryAcl("read_product", new()
    {
        ClusterId = testCluster.Id,
        Principal = testUser.Name.Apply(name => $"User:{name}"),
        ResourceType = "SUBJECT",
        ResourceName = "product-",
        PatternType = "PREFIXED",
        Host = "*",
        Operation = "READ",
        Permission = "ALLOW",
        Username = testUser.Name,
        Password = userPw,
        AllowDeletion = true,
    }, new CustomResourceOptions
    {
        DependsOn =
        {
            schemaRegistryAdmin,
        },
    });

    var writeOrders = new Redpanda.SchemaRegistryAcl("write_orders", new()
    {
        ClusterId = testCluster.Id,
        Principal = testUser.Name.Apply(name => $"User:{name}"),
        ResourceType = "SUBJECT",
        ResourceName = "orders-value",
        PatternType = "LITERAL",
        Host = "*",
        Operation = "WRITE",
        Permission = "ALLOW",
        Username = testUser.Name,
        Password = userPw,
        AllowDeletion = true,
    }, new CustomResourceOptions
    {
        DependsOn =
        {
            schemaRegistryAdmin,
        },
    });

    return new Dictionary<string, object?>
    {
        ["userSchemaInfo"] =
        {
            { "id", userSchema.SchemaId },
            { "subject", userSchema.Subject },
            { "version", userSchema.Version },
            { "type", userSchema.SchemaType },
        },
        ["userEventSchemaInfo"] =
        {
            { "id", userEventSchema.SchemaId },
            { "subject", userEventSchema.Subject },
            { "version", userEventSchema.Version },
            { "type", userEventSchema.SchemaType },
            { "references", userEventSchema.References },
        },
        ["productSchemaInfo"] =
        {
            { "id", productSchema.SchemaId },
            { "subject", productSchema.Subject },
            { "version", productSchema.Version },
            { "type", productSchema.SchemaType },
            { "compatibility", productSchema.Compatibility },
        },
    };
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
	"encoding/json"
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
			Timeouts: &redpanda.NetworkTimeoutsArgs{
				Create: pulumi.String("20m"),
				Delete: pulumi.String("20m"),
			},
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
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"schema_registry_enable_authorization": true,
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
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
			ClusterConfiguration: &redpanda.ClusterClusterConfigurationArgs{
				CustomPropertiesJson: pulumi.String(json0),
			},
			Tags: pulumi.StringMap{
				"key": pulumi.String("value"),
			},
			Timeouts: &redpanda.ClusterTimeoutsArgs{
				Create: pulumi.String("90m"),
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
		_, err = redpanda.NewTopic(ctx, "test", &redpanda.TopicArgs{
			Name:              pulumi.String(topicName),
			PartitionCount:    pulumi.Float64(partitionCount),
			ReplicationFactor: pulumi.Float64(replicationFactor),
			ClusterApiUrl:     testCluster.ClusterApiUrl,
			AllowDeletion:     pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		clusterAdmin, err := redpanda.NewAcl(ctx, "cluster_admin", &redpanda.AclArgs{
			ResourceType:        pulumi.String("CLUSTER"),
			ResourceName:        pulumi.String("kafka-cluster"),
			ResourcePatternType: pulumi.String("LITERAL"),
			Principal: testUser.Name.ApplyT(func(name string) (string, error) {
				return fmt.Sprintf("User:%v", name), nil
			}).(pulumi.StringOutput),
			Host:           pulumi.String("*"),
			Operation:      pulumi.String("ALL"),
			PermissionType: pulumi.String("ALLOW"),
			ClusterApiUrl:  testCluster.ClusterApiUrl,
			AllowDeletion:  pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		schemaRegistryAdmin, err := redpanda.NewAcl(ctx, "schema_registry_admin", &redpanda.AclArgs{
			ResourceType:        pulumi.String("CLUSTER"),
			ResourceName:        pulumi.String("kafka-cluster"),
			ResourcePatternType: pulumi.String("LITERAL"),
			Principal: testUser.Name.ApplyT(func(name string) (string, error) {
				return fmt.Sprintf("User:%v", name), nil
			}).(pulumi.StringOutput),
			Host:           pulumi.String("*"),
			Operation:      pulumi.String("ALTER"),
			PermissionType: pulumi.String("ALLOW"),
			ClusterApiUrl:  testCluster.ClusterApiUrl,
			AllowDeletion:  pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		clusterAction, err := redpanda.NewAcl(ctx, "cluster_action", &redpanda.AclArgs{
			ResourceType:        pulumi.String("CLUSTER"),
			ResourceName:        pulumi.String("kafka-cluster"),
			ResourcePatternType: pulumi.String("LITERAL"),
			Principal: testUser.Name.ApplyT(func(name string) (string, error) {
				return fmt.Sprintf("User:%v", name), nil
			}).(pulumi.StringOutput),
			Host:           pulumi.String("*"),
			Operation:      pulumi.String("CLUSTER_ACTION"),
			PermissionType: pulumi.String("ALLOW"),
			ClusterApiUrl:  testCluster.ClusterApiUrl,
			AllowDeletion:  pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		topicAccess, err := redpanda.NewAcl(ctx, "topic_access", &redpanda.AclArgs{
			ResourceType:        pulumi.String("TOPIC"),
			ResourceName:        pulumi.String(topicName),
			ResourcePatternType: pulumi.String("LITERAL"),
			Principal: testUser.Name.ApplyT(func(name string) (string, error) {
				return fmt.Sprintf("User:%v", name), nil
			}).(pulumi.StringOutput),
			Host:           pulumi.String("*"),
			Operation:      pulumi.String("ALL"),
			PermissionType: pulumi.String("ALLOW"),
			ClusterApiUrl:  testCluster.ClusterApiUrl,
			AllowDeletion:  pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		allTestTopic, err := redpanda.NewSchemaRegistryAcl(ctx, "all_test_topic", &redpanda.SchemaRegistryAclArgs{
			ClusterId: testCluster.ID(),
			Principal: testUser.Name.ApplyT(func(name string) (string, error) {
				return fmt.Sprintf("User:%v", name), nil
			}).(pulumi.StringOutput),
			ResourceType:  pulumi.String("SUBJECT"),
			ResourceName:  pulumi.Sprintf("%v-", topicName),
			PatternType:   pulumi.String("PREFIXED"),
			Host:          pulumi.String("*"),
			Operation:     pulumi.String("ALL"),
			Permission:    pulumi.String("ALLOW"),
			Username:      testUser.Name,
			Password:      pulumi.String(userPw),
			AllowDeletion: pulumi.Bool(true),
		}, pulumi.DependsOn([]pulumi.Resource{
			schemaRegistryAdmin,
		}))
		if err != nil {
			return err
		}
		describeRegistry, err := redpanda.NewSchemaRegistryAcl(ctx, "describe_registry", &redpanda.SchemaRegistryAclArgs{
			ClusterId: testCluster.ID(),
			Principal: testUser.Name.ApplyT(func(name string) (string, error) {
				return fmt.Sprintf("User:%v", name), nil
			}).(pulumi.StringOutput),
			ResourceType:  pulumi.String("REGISTRY"),
			ResourceName:  pulumi.String("*"),
			PatternType:   pulumi.String("LITERAL"),
			Host:          pulumi.String("*"),
			Operation:     pulumi.String("DESCRIBE"),
			Permission:    pulumi.String("ALLOW"),
			Username:      testUser.Name,
			Password:      pulumi.String(userPw),
			AllowDeletion: pulumi.Bool(true),
		}, pulumi.DependsOn([]pulumi.Resource{
			schemaRegistryAdmin,
		}))
		if err != nil {
			return err
		}
		alterConfigsRegistry, err := redpanda.NewSchemaRegistryAcl(ctx, "alter_configs_registry", &redpanda.SchemaRegistryAclArgs{
			ClusterId: testCluster.ID(),
			Principal: testUser.Name.ApplyT(func(name string) (string, error) {
				return fmt.Sprintf("User:%v", name), nil
			}).(pulumi.StringOutput),
			ResourceType:  pulumi.String("REGISTRY"),
			ResourceName:  pulumi.String("*"),
			PatternType:   pulumi.String("LITERAL"),
			Host:          pulumi.String("*"),
			Operation:     pulumi.String("ALTER_CONFIGS"),
			Permission:    pulumi.String("ALLOW"),
			Username:      testUser.Name,
			Password:      pulumi.String(userPw),
			AllowDeletion: pulumi.Bool(true),
		}, pulumi.DependsOn([]pulumi.Resource{
			schemaRegistryAdmin,
		}))
		if err != nil {
			return err
		}
		// The type of schema (AVRO, JSON, PROTOBUF)
		schemaType := "AVRO"
		if param := cfg.Get("schemaType"); param != "" {
			schemaType = param
		}
		// The AVRO schema definition for user data
		userSchemaDefinition := `{
  "type": "record",
  "name": "User",
  "fields": [
    {
      "name": "id",
      "type": "int"
    },
    {
      "name": "name",
      "type": "string"
    },
    {
      "name": "email",
      "type": "string"
    },
    {
      "name": "created_at",
      "type": "long",
      "logicalType": "timestamp-millis"
    }
  ]
}
`
		if param := cfg.Get("userSchemaDefinition"); param != "" {
			userSchemaDefinition = param
		}
		userSchema, err := redpanda.NewSchema(ctx, "user_schema", &redpanda.SchemaArgs{
			ClusterId:     testCluster.ID(),
			Subject:       pulumi.Sprintf("%v-value", topicName),
			SchemaType:    pulumi.String(schemaType),
			Schema:        pulumi.String(userSchemaDefinition),
			Username:      testUser.Name,
			Password:      pulumi.String(userPw),
			AllowDeletion: pulumi.Bool(true),
		}, pulumi.DependsOn([]pulumi.Resource{
			clusterAdmin,
			schemaRegistryAdmin,
			clusterAction,
			topicAccess,
			allTestTopic,
			describeRegistry,
			alterConfigsRegistry,
		}))
		if err != nil {
			return err
		}
		// The AVRO schema definition for user events that references the User schema
		userEventSchemaDefinition := `{
  "type": "record",
  "name": "UserEvent",
  "fields": [
    {
      "name": "event_id",
      "type": "string"
    },
    {
      "name": "event_type",
      "type": {
        "type": "enum",
        "name": "EventType",
        "symbols": ["CREATED", "UPDATED", "DELETED"]
      }
    },
    {
      "name": "user",
      "type": "User"
    },
    {
      "name": "timestamp",
      "type": "long",
      "logicalType": "timestamp-millis"
    },
    {
      "name": "metadata",
      "type": ["null", {
        "type": "map",
        "values": "string"
      }],
      "default": null
    }
  ]
}
`
		if param := cfg.Get("userEventSchemaDefinition"); param != "" {
			userEventSchemaDefinition = param
		}
		userEventSchema, err := redpanda.NewSchema(ctx, "user_event_schema", &redpanda.SchemaArgs{
			ClusterId:     testCluster.ID(),
			Subject:       pulumi.Sprintf("%v-events-value", topicName),
			SchemaType:    pulumi.String(schemaType),
			Schema:        pulumi.String(userEventSchemaDefinition),
			Username:      testUser.Name,
			Password:      pulumi.String(userPw),
			AllowDeletion: pulumi.Bool(true),
			References: redpanda.SchemaReferenceArray{
				&redpanda.SchemaReferenceArgs{
					Name:    pulumi.String("User"),
					Subject: userSchema.Subject,
					Version: userSchema.Version,
				},
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			clusterAdmin,
			schemaRegistryAdmin,
			clusterAction,
			topicAccess,
			allTestTopic,
			describeRegistry,
			alterConfigsRegistry,
		}))
		if err != nil {
			return err
		}
		// The AVRO schema definition for product data with strict compatibility
		productSchemaDefinition := `{
  "type": "record",
  "name": "Product",
  "fields": [
    {
      "name": "id",
      "type": "string"
    },
    {
      "name": "name",
      "type": "string"
    },
    {
      "name": "price",
      "type": {
        "type": "bytes",
        "logicalType": "decimal",
        "precision": 10,
        "scale": 2
      }
    },
    {
      "name": "category",
      "type": {
        "type": "enum",
        "name": "Category",
        "symbols": ["ELECTRONICS", "CLOTHING", "BOOKS", "HOME"]
      }
    },
    {
      "name": "description",
      "type": ["null", "string"],
      "default": null
    },
    {
      "name": "created_at",
      "type": "long",
      "logicalType": "timestamp-millis"
    }
  ]
}
`
		if param := cfg.Get("productSchemaDefinition"); param != "" {
			productSchemaDefinition = param
		}
		// The compatibility level for schema evolution (BACKWARD, BACKWARD_TRANSITIVE, FORWARD, FORWARD_TRANSITIVE, FULL, FULL_TRANSITIVE, NONE)
		compatibilityLevel := "FULL"
		if param := cfg.Get("compatibilityLevel"); param != "" {
			compatibilityLevel = param
		}
		productSchema, err := redpanda.NewSchema(ctx, "product_schema", &redpanda.SchemaArgs{
			ClusterId:     testCluster.ID(),
			Subject:       pulumi.Sprintf("%v-product-value", topicName),
			SchemaType:    pulumi.String(schemaType),
			Schema:        pulumi.String(productSchemaDefinition),
			Compatibility: pulumi.String(compatibilityLevel),
			Username:      testUser.Name,
			Password:      pulumi.String(userPw),
			AllowDeletion: pulumi.Bool(true),
		}, pulumi.DependsOn([]pulumi.Resource{
			clusterAdmin,
			schemaRegistryAdmin,
			clusterAction,
			topicAccess,
			allTestTopic,
			describeRegistry,
			alterConfigsRegistry,
		}))
		if err != nil {
			return err
		}
		_, err = redpanda.NewSchemaRegistryAcl(ctx, "read_product", &redpanda.SchemaRegistryAclArgs{
			ClusterId: testCluster.ID(),
			Principal: testUser.Name.ApplyT(func(name string) (string, error) {
				return fmt.Sprintf("User:%v", name), nil
			}).(pulumi.StringOutput),
			ResourceType:  pulumi.String("SUBJECT"),
			ResourceName:  pulumi.String("product-"),
			PatternType:   pulumi.String("PREFIXED"),
			Host:          pulumi.String("*"),
			Operation:     pulumi.String("READ"),
			Permission:    pulumi.String("ALLOW"),
			Username:      testUser.Name,
			Password:      pulumi.String(userPw),
			AllowDeletion: pulumi.Bool(true),
		}, pulumi.DependsOn([]pulumi.Resource{
			schemaRegistryAdmin,
		}))
		if err != nil {
			return err
		}
		_, err = redpanda.NewSchemaRegistryAcl(ctx, "write_orders", &redpanda.SchemaRegistryAclArgs{
			ClusterId: testCluster.ID(),
			Principal: testUser.Name.ApplyT(func(name string) (string, error) {
				return fmt.Sprintf("User:%v", name), nil
			}).(pulumi.StringOutput),
			ResourceType:  pulumi.String("SUBJECT"),
			ResourceName:  pulumi.String("orders-value"),
			PatternType:   pulumi.String("LITERAL"),
			Host:          pulumi.String("*"),
			Operation:     pulumi.String("WRITE"),
			Permission:    pulumi.String("ALLOW"),
			Username:      testUser.Name,
			Password:      pulumi.String(userPw),
			AllowDeletion: pulumi.Bool(true),
		}, pulumi.DependsOn([]pulumi.Resource{
			schemaRegistryAdmin,
		}))
		if err != nil {
			return err
		}
		ctx.Export("userSchemaInfo", pulumi.Map{
			"id":      userSchema.SchemaId,
			"subject": userSchema.Subject,
			"version": userSchema.Version,
			"type":    userSchema.SchemaType,
		})
		ctx.Export("userEventSchemaInfo", pulumi.Map{
			"id":         userEventSchema.SchemaId,
			"subject":    userEventSchema.Subject,
			"version":    userEventSchema.Version,
			"type":       userEventSchema.SchemaType,
			"references": userEventSchema.References,
		})
		ctx.Export("productSchemaInfo", pulumi.Map{
			"id":            productSchema.SchemaId,
			"subject":       productSchema.Subject,
			"version":       productSchema.Version,
			"type":          productSchema.SchemaType,
			"compatibility": productSchema.Compatibility,
		})
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
  schemaType:
    type: string
    default: AVRO
  userSchemaDefinition:
    type: string
    default: |
      {
        "type": "record",
        "name": "User",
        "fields": [
          {
            "name": "id",
            "type": "int"
          },
          {
            "name": "name",
            "type": "string"
          },
          {
            "name": "email",
            "type": "string"
          },
          {
            "name": "created_at",
            "type": "long",
            "logicalType": "timestamp-millis"
          }
        ]
      }
  userEventSchemaDefinition:
    type: string
    default: |
      {
        "type": "record",
        "name": "UserEvent",
        "fields": [
          {
            "name": "event_id",
            "type": "string"
          },
          {
            "name": "event_type",
            "type": {
              "type": "enum",
              "name": "EventType",
              "symbols": ["CREATED", "UPDATED", "DELETED"]
            }
          },
          {
            "name": "user",
            "type": "User"
          },
          {
            "name": "timestamp",
            "type": "long",
            "logicalType": "timestamp-millis"
          },
          {
            "name": "metadata",
            "type": ["null", {
              "type": "map",
              "values": "string"
            }],
            "default": null
          }
        ]
      }
  productSchemaDefinition:
    type: string
    default: |
      {
        "type": "record",
        "name": "Product",
        "fields": [
          {
            "name": "id",
            "type": "string"
          },
          {
            "name": "name",
            "type": "string"
          },
          {
            "name": "price",
            "type": {
              "type": "bytes",
              "logicalType": "decimal",
              "precision": 10,
              "scale": 2
            }
          },
          {
            "name": "category",
            "type": {
              "type": "enum",
              "name": "Category",
              "symbols": ["ELECTRONICS", "CLOTHING", "BOOKS", "HOME"]
            }
          },
          {
            "name": "description",
            "type": ["null", "string"],
            "default": null
          },
          {
            "name": "created_at",
            "type": "long",
            "logicalType": "timestamp-millis"
          }
        ]
      }
  compatibilityLevel:
    type: string
    default: FULL
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
      timeouts:
        create: 20m
        delete: 20m
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
      clusterConfiguration:
        customPropertiesJson:
          fn::toJSON:
            schema_registry_enable_authorization: true
      tags:
        key: value
      timeouts:
        create: 90m
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
  userSchema:
    type: redpanda:Schema
    name: user_schema
    properties:
      clusterId: ${testCluster.id}
      subject: ${topicName}-value
      schemaType: ${schemaType}
      schema: ${userSchemaDefinition}
      username: ${testUser.name}
      password: ${userPw}
      allowDeletion: true
    options:
      dependsOn:
        - ${clusterAdmin}
        - ${schemaRegistryAdmin}
        - ${clusterAction}
        - ${topicAccess}
        - ${allTestTopic}
        - ${describeRegistry}
        - ${alterConfigsRegistry}
  userEventSchema:
    type: redpanda:Schema
    name: user_event_schema
    properties:
      clusterId: ${testCluster.id}
      subject: ${topicName}-events-value
      schemaType: ${schemaType}
      schema: ${userEventSchemaDefinition}
      username: ${testUser.name}
      password: ${userPw}
      allowDeletion: true
      references:
        - name: User
          subject: ${userSchema.subject}
          version: ${userSchema.version}
    options:
      dependsOn:
        - ${clusterAdmin}
        - ${schemaRegistryAdmin}
        - ${clusterAction}
        - ${topicAccess}
        - ${allTestTopic}
        - ${describeRegistry}
        - ${alterConfigsRegistry}
  productSchema:
    type: redpanda:Schema
    name: product_schema
    properties:
      clusterId: ${testCluster.id}
      subject: ${topicName}-product-value
      schemaType: ${schemaType}
      schema: ${productSchemaDefinition}
      compatibility: ${compatibilityLevel}
      username: ${testUser.name}
      password: ${userPw}
      allowDeletion: true
    options:
      dependsOn:
        - ${clusterAdmin}
        - ${schemaRegistryAdmin}
        - ${clusterAction}
        - ${topicAccess}
        - ${allTestTopic}
        - ${describeRegistry}
        - ${alterConfigsRegistry}
  clusterAdmin:
    type: redpanda:Acl
    name: cluster_admin
    properties:
      resourceType: CLUSTER
      resourceName: kafka-cluster
      resourcePatternType: LITERAL
      principal: User:${testUser.name}
      host: '*'
      operation: ALL
      permissionType: ALLOW
      clusterApiUrl: ${testCluster.clusterApiUrl}
      allowDeletion: true
  schemaRegistryAdmin:
    type: redpanda:Acl
    name: schema_registry_admin
    properties:
      resourceType: CLUSTER
      resourceName: kafka-cluster
      resourcePatternType: LITERAL
      principal: User:${testUser.name}
      host: '*'
      operation: ALTER
      permissionType: ALLOW
      clusterApiUrl: ${testCluster.clusterApiUrl}
      allowDeletion: true
  clusterAction:
    type: redpanda:Acl
    name: cluster_action
    properties:
      resourceType: CLUSTER
      resourceName: kafka-cluster
      resourcePatternType: LITERAL
      principal: User:${testUser.name}
      host: '*'
      operation: CLUSTER_ACTION
      permissionType: ALLOW
      clusterApiUrl: ${testCluster.clusterApiUrl}
      allowDeletion: true
  topicAccess:
    type: redpanda:Acl
    name: topic_access
    properties:
      resourceType: TOPIC
      resourceName: ${topicName}
      resourcePatternType: LITERAL
      principal: User:${testUser.name}
      host: '*'
      operation: ALL
      permissionType: ALLOW
      clusterApiUrl: ${testCluster.clusterApiUrl}
      allowDeletion: true
  readProduct:
    type: redpanda:SchemaRegistryAcl
    name: read_product
    properties:
      clusterId: ${testCluster.id}
      principal: User:${testUser.name}
      resourceType: SUBJECT
      resourceName: product-
      patternType: PREFIXED
      host: '*'
      operation: READ
      permission: ALLOW
      username: ${testUser.name}
      password: ${userPw}
      allowDeletion: true
    options:
      dependsOn:
        - ${schemaRegistryAdmin}
  writeOrders:
    type: redpanda:SchemaRegistryAcl
    name: write_orders
    properties:
      clusterId: ${testCluster.id}
      principal: User:${testUser.name}
      resourceType: SUBJECT
      resourceName: orders-value
      patternType: LITERAL
      host: '*'
      operation: WRITE
      permission: ALLOW
      username: ${testUser.name}
      password: ${userPw}
      allowDeletion: true
    options:
      dependsOn:
        - ${schemaRegistryAdmin}
  allTestTopic:
    type: redpanda:SchemaRegistryAcl
    name: all_test_topic
    properties:
      clusterId: ${testCluster.id}
      principal: User:${testUser.name}
      resourceType: SUBJECT
      resourceName: ${topicName}-
      patternType: PREFIXED
      host: '*'
      operation: ALL
      permission: ALLOW
      username: ${testUser.name}
      password: ${userPw}
      allowDeletion: true
    options:
      dependsOn:
        - ${schemaRegistryAdmin}
  describeRegistry:
    type: redpanda:SchemaRegistryAcl
    name: describe_registry
    properties:
      clusterId: ${testCluster.id}
      principal: User:${testUser.name}
      resourceType: REGISTRY
      resourceName: '*'
      patternType: LITERAL
      host: '*'
      operation: DESCRIBE
      permission: ALLOW
      username: ${testUser.name}
      password: ${userPw}
      allowDeletion: true
    options:
      dependsOn:
        - ${schemaRegistryAdmin}
  alterConfigsRegistry:
    type: redpanda:SchemaRegistryAcl
    name: alter_configs_registry
    properties:
      clusterId: ${testCluster.id}
      principal: User:${testUser.name}
      resourceType: REGISTRY
      resourceName: '*'
      patternType: LITERAL
      host: '*'
      operation: ALTER_CONFIGS
      permission: ALLOW
      username: ${testUser.name}
      password: ${userPw}
      allowDeletion: true
    options:
      dependsOn:
        - ${schemaRegistryAdmin}
outputs:
  userSchemaInfo:
    id: ${userSchema.schemaId}
    subject: ${userSchema.subject}
    version: ${userSchema.version}
    type: ${userSchema.schemaType}
  userEventSchemaInfo:
    id: ${userEventSchema.schemaId}
    subject: ${userEventSchema.subject}
    version: ${userEventSchema.version}
    type: ${userEventSchema.schemaType}
    references: ${userEventSchema.references}
  productSchemaInfo:
    id: ${productSchema.schemaId}
    subject: ${productSchema.subject}
    version: ${productSchema.version}
    type: ${productSchema.schemaType}
    compatibility: ${productSchema.compatibility}
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
import com.pulumi.redpanda.inputs.NetworkTimeoutsArgs;
import com.pulumi.redpanda.Cluster;
import com.pulumi.redpanda.ClusterArgs;
import com.pulumi.redpanda.inputs.ClusterClusterConfigurationArgs;
import com.pulumi.redpanda.inputs.ClusterTimeoutsArgs;
import com.pulumi.redpanda.User;
import com.pulumi.redpanda.UserArgs;
import com.pulumi.redpanda.Topic;
import com.pulumi.redpanda.TopicArgs;
import com.pulumi.redpanda.Acl;
import com.pulumi.redpanda.AclArgs;
import com.pulumi.redpanda.SchemaRegistryAcl;
import com.pulumi.redpanda.SchemaRegistryAclArgs;
import com.pulumi.redpanda.Schema;
import com.pulumi.redpanda.SchemaArgs;
import com.pulumi.redpanda.inputs.SchemaReferenceArgs;
import static com.pulumi.codegen.internal.Serialization.*;
import com.pulumi.resources.CustomResourceOptions;
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
            .timeouts(NetworkTimeoutsArgs.builder()
                .create("20m")
                .delete("20m")
                .build())
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
            .clusterConfiguration(ClusterClusterConfigurationArgs.builder()
                .customPropertiesJson(serializeJson(
                    jsonObject(
                        jsonProperty("schema_registry_enable_authorization", true)
                    )))
                .build())
            .tags(Map.of("key", "value"))
            .timeouts(ClusterTimeoutsArgs.builder()
                .create("90m")
                .build())
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

        var clusterAdmin = new Acl("clusterAdmin", AclArgs.builder()
            .resourceType("CLUSTER")
            .resourceName("kafka-cluster")
            .resourcePatternType("LITERAL")
            .principal(testUser.name().applyValue(name -> String.format("User:%s", name)))
            .host("*")
            .operation("ALL")
            .permissionType("ALLOW")
            .clusterApiUrl(testCluster.clusterApiUrl())
            .allowDeletion(true)
            .build());

        var schemaRegistryAdmin = new Acl("schemaRegistryAdmin", AclArgs.builder()
            .resourceType("CLUSTER")
            .resourceName("kafka-cluster")
            .resourcePatternType("LITERAL")
            .principal(testUser.name().applyValue(name -> String.format("User:%s", name)))
            .host("*")
            .operation("ALTER")
            .permissionType("ALLOW")
            .clusterApiUrl(testCluster.clusterApiUrl())
            .allowDeletion(true)
            .build());

        var clusterAction = new Acl("clusterAction", AclArgs.builder()
            .resourceType("CLUSTER")
            .resourceName("kafka-cluster")
            .resourcePatternType("LITERAL")
            .principal(testUser.name().applyValue(name -> String.format("User:%s", name)))
            .host("*")
            .operation("CLUSTER_ACTION")
            .permissionType("ALLOW")
            .clusterApiUrl(testCluster.clusterApiUrl())
            .allowDeletion(true)
            .build());

        var topicAccess = new Acl("topicAccess", AclArgs.builder()
            .resourceType("TOPIC")
            .resourceName(topicName)
            .resourcePatternType("LITERAL")
            .principal(testUser.name().applyValue(name -> String.format("User:%s", name)))
            .host("*")
            .operation("ALL")
            .permissionType("ALLOW")
            .clusterApiUrl(testCluster.clusterApiUrl())
            .allowDeletion(true)
            .build());

        var allTestTopic = new SchemaRegistryAcl("allTestTopic", SchemaRegistryAclArgs.builder()
            .clusterId(testCluster.id())
            .principal(testUser.name().applyValue(name -> String.format("User:%s", name)))
            .resourceType("SUBJECT")
            .resourceName(String.format("%s-", topicName))
            .patternType("PREFIXED")
            .host("*")
            .operation("ALL")
            .permission("ALLOW")
            .username(testUser.name())
            .password(userPw)
            .allowDeletion(true)
            .build(), CustomResourceOptions.builder()
                .dependsOn(schemaRegistryAdmin)
                .build());

        var describeRegistry = new SchemaRegistryAcl("describeRegistry", SchemaRegistryAclArgs.builder()
            .clusterId(testCluster.id())
            .principal(testUser.name().applyValue(name -> String.format("User:%s", name)))
            .resourceType("REGISTRY")
            .resourceName("*")
            .patternType("LITERAL")
            .host("*")
            .operation("DESCRIBE")
            .permission("ALLOW")
            .username(testUser.name())
            .password(userPw)
            .allowDeletion(true)
            .build(), CustomResourceOptions.builder()
                .dependsOn(schemaRegistryAdmin)
                .build());

        var alterConfigsRegistry = new SchemaRegistryAcl("alterConfigsRegistry", SchemaRegistryAclArgs.builder()
            .clusterId(testCluster.id())
            .principal(testUser.name().applyValue(name -> String.format("User:%s", name)))
            .resourceType("REGISTRY")
            .resourceName("*")
            .patternType("LITERAL")
            .host("*")
            .operation("ALTER_CONFIGS")
            .permission("ALLOW")
            .username(testUser.name())
            .password(userPw)
            .allowDeletion(true)
            .build(), CustomResourceOptions.builder()
                .dependsOn(schemaRegistryAdmin)
                .build());

        final var schemaType = config.get("schemaType").orElse("AVRO");
        final var userSchemaDefinition = config.get("userSchemaDefinition").orElse("""
{
  "type": "record",
  "name": "User",
  "fields": [
    {
      "name": "id",
      "type": "int"
    },
    {
      "name": "name",
      "type": "string"
    },
    {
      "name": "email",
      "type": "string"
    },
    {
      "name": "created_at",
      "type": "long",
      "logicalType": "timestamp-millis"
    }
  ]
}
        """);
        var userSchema = new Schema("userSchema", SchemaArgs.builder()
            .clusterId(testCluster.id())
            .subject(String.format("%s-value", topicName))
            .schemaType(schemaType)
            .schema(userSchemaDefinition)
            .username(testUser.name())
            .password(userPw)
            .allowDeletion(true)
            .build(), CustomResourceOptions.builder()
                .dependsOn(
                    clusterAdmin,
                    schemaRegistryAdmin,
                    clusterAction,
                    topicAccess,
                    allTestTopic,
                    describeRegistry,
                    alterConfigsRegistry)
                .build());

        final var userEventSchemaDefinition = config.get("userEventSchemaDefinition").orElse("""
{
  "type": "record",
  "name": "UserEvent",
  "fields": [
    {
      "name": "event_id",
      "type": "string"
    },
    {
      "name": "event_type",
      "type": {
        "type": "enum",
        "name": "EventType",
        "symbols": ["CREATED", "UPDATED", "DELETED"]
      }
    },
    {
      "name": "user",
      "type": "User"
    },
    {
      "name": "timestamp",
      "type": "long",
      "logicalType": "timestamp-millis"
    },
    {
      "name": "metadata",
      "type": ["null", {
        "type": "map",
        "values": "string"
      }],
      "default": null
    }
  ]
}
        """);
        var userEventSchema = new Schema("userEventSchema", SchemaArgs.builder()
            .clusterId(testCluster.id())
            .subject(String.format("%s-events-value", topicName))
            .schemaType(schemaType)
            .schema(userEventSchemaDefinition)
            .username(testUser.name())
            .password(userPw)
            .allowDeletion(true)
            .references(SchemaReferenceArgs.builder()
                .name("User")
                .subject(userSchema.subject())
                .version(userSchema.version())
                .build())
            .build(), CustomResourceOptions.builder()
                .dependsOn(
                    clusterAdmin,
                    schemaRegistryAdmin,
                    clusterAction,
                    topicAccess,
                    allTestTopic,
                    describeRegistry,
                    alterConfigsRegistry)
                .build());

        final var productSchemaDefinition = config.get("productSchemaDefinition").orElse("""
{
  "type": "record",
  "name": "Product",
  "fields": [
    {
      "name": "id",
      "type": "string"
    },
    {
      "name": "name",
      "type": "string"
    },
    {
      "name": "price",
      "type": {
        "type": "bytes",
        "logicalType": "decimal",
        "precision": 10,
        "scale": 2
      }
    },
    {
      "name": "category",
      "type": {
        "type": "enum",
        "name": "Category",
        "symbols": ["ELECTRONICS", "CLOTHING", "BOOKS", "HOME"]
      }
    },
    {
      "name": "description",
      "type": ["null", "string"],
      "default": null
    },
    {
      "name": "created_at",
      "type": "long",
      "logicalType": "timestamp-millis"
    }
  ]
}
        """);
        final var compatibilityLevel = config.get("compatibilityLevel").orElse("FULL");
        var productSchema = new Schema("productSchema", SchemaArgs.builder()
            .clusterId(testCluster.id())
            .subject(String.format("%s-product-value", topicName))
            .schemaType(schemaType)
            .schema(productSchemaDefinition)
            .compatibility(compatibilityLevel)
            .username(testUser.name())
            .password(userPw)
            .allowDeletion(true)
            .build(), CustomResourceOptions.builder()
                .dependsOn(
                    clusterAdmin,
                    schemaRegistryAdmin,
                    clusterAction,
                    topicAccess,
                    allTestTopic,
                    describeRegistry,
                    alterConfigsRegistry)
                .build());

        var readProduct = new SchemaRegistryAcl("readProduct", SchemaRegistryAclArgs.builder()
            .clusterId(testCluster.id())
            .principal(testUser.name().applyValue(name -> String.format("User:%s", name)))
            .resourceType("SUBJECT")
            .resourceName("product-")
            .patternType("PREFIXED")
            .host("*")
            .operation("READ")
            .permission("ALLOW")
            .username(testUser.name())
            .password(userPw)
            .allowDeletion(true)
            .build(), CustomResourceOptions.builder()
                .dependsOn(schemaRegistryAdmin)
                .build());

        var writeOrders = new SchemaRegistryAcl("writeOrders", SchemaRegistryAclArgs.builder()
            .clusterId(testCluster.id())
            .principal(testUser.name().applyValue(name -> String.format("User:%s", name)))
            .resourceType("SUBJECT")
            .resourceName("orders-value")
            .patternType("LITERAL")
            .host("*")
            .operation("WRITE")
            .permission("ALLOW")
            .username(testUser.name())
            .password(userPw)
            .allowDeletion(true)
            .build(), CustomResourceOptions.builder()
                .dependsOn(schemaRegistryAdmin)
                .build());

        ctx.export("userSchemaInfo", Map.ofEntries(
            Map.entry("id", userSchema.schemaId()),
            Map.entry("subject", userSchema.subject()),
            Map.entry("version", userSchema.version()),
            Map.entry("type", userSchema.schemaType())
        ));
        ctx.export("userEventSchemaInfo", Map.ofEntries(
            Map.entry("id", userEventSchema.schemaId()),
            Map.entry("subject", userEventSchema.subject()),
            Map.entry("version", userEventSchema.version()),
            Map.entry("type", userEventSchema.schemaType()),
            Map.entry("references", userEventSchema.references())
        ));
        ctx.export("productSchemaInfo", Map.ofEntries(
            Map.entry("id", productSchema.schemaId()),
            Map.entry("subject", productSchema.subject()),
            Map.entry("version", productSchema.version()),
            Map.entry("type", productSchema.schemaType()),
            Map.entry("compatibility", productSchema.compatibility())
        ));
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
const resourceGroupName = config.get("resourceGroupName") || "testname";
const test = new redpanda.ResourceGroup("test", {name: resourceGroupName});
const networkName = config.get("networkName") || "testname";
const region = config.get("region") || "us-central1";
const cloudProvider = config.get("cloudProvider") || "gcp";
const testNetwork = new redpanda.Network("test", {
    name: networkName,
    resourceGroupId: test.id,
    cloudProvider: cloudProvider,
    region: region,
    clusterType: "dedicated",
    cidrBlock: "10.0.0.0/20",
    timeouts: {
        create: "20m",
        "delete": "20m",
    },
});
const clusterName = config.get("clusterName") || "testname";
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
    clusterConfiguration: {
        customPropertiesJson: JSON.stringify({
            schema_registry_enable_authorization: true,
        }),
    },
    timeouts: {
        create: "90m",
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
const schemaRegistryAdmin = new redpanda.Acl("schema_registry_admin", {
    resourceType: "CLUSTER",
    resourceName: "kafka-cluster",
    resourcePatternType: "LITERAL",
    principal: pulumi.interpolate`User:${testUser.name}`,
    host: "*",
    operation: "ALTER",
    permissionType: "ALLOW",
    clusterApiUrl: testCluster.clusterApiUrl,
    allowDeletion: true,
});
// The type of schema (AVRO, JSON, PROTOBUF)
const schemaType = config.get("schemaType") || "AVRO";
// The AVRO schema definition for user data
const userSchemaDefinition = config.get("userSchemaDefinition") || `{
  "type": "record",
  "name": "User",
  "fields": [
    {
      "name": "id",
      "type": "int"
    },
    {
      "name": "name",
      "type": "string"
    },
    {
      "name": "email",
      "type": "string"
    },
    {
      "name": "created_at",
      "type": "long",
      "logicalType": "timestamp-millis"
    }
  ]
}
`;
const allTestTopic = new redpanda.SchemaRegistryAcl("all_test_topic", {
    clusterId: testCluster.id,
    principal: pulumi.interpolate`User:${testUser.name}`,
    resourceType: "SUBJECT",
    resourceName: `${topicName}-`,
    patternType: "PREFIXED",
    host: "*",
    operation: "ALL",
    permission: "ALLOW",
    username: testUser.name,
    password: userPw,
    allowDeletion: true,
}, {
    dependsOn: [schemaRegistryAdmin],
});
const describeRegistry = new redpanda.SchemaRegistryAcl("describe_registry", {
    clusterId: testCluster.id,
    principal: pulumi.interpolate`User:${testUser.name}`,
    resourceType: "REGISTRY",
    resourceName: "*",
    patternType: "LITERAL",
    host: "*",
    operation: "DESCRIBE",
    permission: "ALLOW",
    username: testUser.name,
    password: userPw,
    allowDeletion: true,
}, {
    dependsOn: [schemaRegistryAdmin],
});
const alterConfigsRegistry = new redpanda.SchemaRegistryAcl("alter_configs_registry", {
    clusterId: testCluster.id,
    principal: pulumi.interpolate`User:${testUser.name}`,
    resourceType: "REGISTRY",
    resourceName: "*",
    patternType: "LITERAL",
    host: "*",
    operation: "ALTER_CONFIGS",
    permission: "ALLOW",
    username: testUser.name,
    password: userPw,
    allowDeletion: true,
}, {
    dependsOn: [schemaRegistryAdmin],
});
const userSchema = new redpanda.Schema("user_schema", {
    clusterId: testCluster.id,
    subject: `${topicName}-value`,
    schemaType: schemaType,
    schema: userSchemaDefinition,
    username: testUser.name,
    password: userPw,
    allowDeletion: true,
}, {
    dependsOn: [
        schemaRegistryAdmin,
        allTestTopic,
        describeRegistry,
        alterConfigsRegistry,
    ],
});
// The AVRO schema definition for user events that references the User schema
const userEventSchemaDefinition = config.get("userEventSchemaDefinition") || `{
  "type": "record",
  "name": "UserEvent",
  "fields": [
    {
      "name": "event_id",
      "type": "string"
    },
    {
      "name": "event_type",
      "type": {
        "type": "enum",
        "name": "EventType",
        "symbols": ["CREATED", "UPDATED", "DELETED"]
      }
    },
    {
      "name": "user",
      "type": "User"
    },
    {
      "name": "timestamp",
      "type": "long",
      "logicalType": "timestamp-millis"
    },
    {
      "name": "metadata",
      "type": ["null", {
        "type": "map",
        "values": "string"
      }],
      "default": null
    }
  ]
}
`;
const userEventSchema = new redpanda.Schema("user_event_schema", {
    clusterId: testCluster.id,
    subject: `${topicName}-events-value`,
    schemaType: schemaType,
    schema: userEventSchemaDefinition,
    username: testUser.name,
    password: userPw,
    allowDeletion: true,
    references: [{
        name: "User",
        subject: userSchema.subject,
        version: userSchema.version,
    }],
}, {
    dependsOn: [
        schemaRegistryAdmin,
        allTestTopic,
        describeRegistry,
        alterConfigsRegistry,
    ],
});
// The AVRO schema definition for product data with strict compatibility
const productSchemaDefinition = config.get("productSchemaDefinition") || `{
  "type": "record",
  "name": "Product",
  "fields": [
    {
      "name": "id",
      "type": "string"
    },
    {
      "name": "name",
      "type": "string"
    },
    {
      "name": "price",
      "type": {
        "type": "bytes",
        "logicalType": "decimal",
        "precision": 10,
        "scale": 2
      }
    },
    {
      "name": "category",
      "type": {
        "type": "enum",
        "name": "Category",
        "symbols": ["ELECTRONICS", "CLOTHING", "BOOKS", "HOME"]
      }
    },
    {
      "name": "description",
      "type": ["null", "string"],
      "default": null
    },
    {
      "name": "created_at",
      "type": "long",
      "logicalType": "timestamp-millis"
    }
  ]
}
`;
// The compatibility level for schema evolution (BACKWARD, BACKWARD_TRANSITIVE, FORWARD, FORWARD_TRANSITIVE, FULL, FULL_TRANSITIVE, NONE)
const compatibilityLevel = config.get("compatibilityLevel") || "FULL";
const productSchema = new redpanda.Schema("product_schema", {
    clusterId: testCluster.id,
    subject: `${topicName}-product-value`,
    schemaType: schemaType,
    schema: productSchemaDefinition,
    compatibility: compatibilityLevel,
    username: testUser.name,
    password: userPw,
    allowDeletion: true,
}, {
    dependsOn: [
        schemaRegistryAdmin,
        allTestTopic,
        describeRegistry,
        alterConfigsRegistry,
    ],
});
const clusterAdmin = new redpanda.Acl("cluster_admin", {
    resourceType: "CLUSTER",
    resourceName: "kafka-cluster",
    resourcePatternType: "LITERAL",
    principal: pulumi.interpolate`User:${testUser.name}`,
    host: "*",
    operation: "ALL",
    permissionType: "ALLOW",
    clusterApiUrl: testCluster.clusterApiUrl,
    allowDeletion: true,
});
const readProduct = new redpanda.SchemaRegistryAcl("read_product", {
    clusterId: testCluster.id,
    principal: pulumi.interpolate`User:${testUser.name}`,
    resourceType: "SUBJECT",
    resourceName: "product-",
    patternType: "PREFIXED",
    host: "*",
    operation: "READ",
    permission: "ALLOW",
    username: testUser.name,
    password: userPw,
    allowDeletion: true,
}, {
    dependsOn: [schemaRegistryAdmin],
});
const writeOrders = new redpanda.SchemaRegistryAcl("write_orders", {
    clusterId: testCluster.id,
    principal: pulumi.interpolate`User:${testUser.name}`,
    resourceType: "SUBJECT",
    resourceName: "orders-value",
    patternType: "LITERAL",
    host: "*",
    operation: "WRITE",
    permission: "ALLOW",
    username: testUser.name,
    password: userPw,
    allowDeletion: true,
}, {
    dependsOn: [schemaRegistryAdmin],
});
const describeTestTopic = new redpanda.SchemaRegistryAcl("describe_test_topic", {
    clusterId: testCluster.id,
    principal: pulumi.interpolate`User:${testUser.name}`,
    resourceType: "SUBJECT",
    resourceName: `${topicName}-`,
    patternType: "PREFIXED",
    host: "*",
    operation: "DESCRIBE",
    permission: "ALLOW",
    username: testUser.name,
    password: userPw,
    allowDeletion: true,
}, {
    dependsOn: [schemaRegistryAdmin],
});
export const userSchemaInfo = {
    id: userSchema.schemaId,
    subject: userSchema.subject,
    version: userSchema.version,
    type: userSchema.schemaType,
};
export const userEventSchemaInfo = {
    id: userEventSchema.schemaId,
    subject: userEventSchema.subject,
    version: userEventSchema.version,
    type: userEventSchema.schemaType,
    references: userEventSchema.references,
};
export const productSchemaInfo = {
    id: productSchema.schemaId,
    subject: productSchema.subject,
    version: productSchema.version,
    type: productSchema.schemaType,
    compatibility: productSchema.compatibility,
};
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
import json
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
    cidr_block="10.0.0.0/20",
    timeouts={
        "create": "20m",
        "delete": "20m",
    })
cluster_name = config.get("clusterName")
if cluster_name is None:
    cluster_name = "testname"
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
    allow_deletion=True,
    cluster_configuration={
        "custom_properties_json": json.dumps({
            "schema_registry_enable_authorization": True,
        }),
    },
    timeouts={
        "create": "90m",
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
schema_registry_admin = redpanda.Acl("schema_registry_admin",
    resource_type="CLUSTER",
    resource_name_="kafka-cluster",
    resource_pattern_type="LITERAL",
    principal=test_user.name.apply(lambda name: f"User:{name}"),
    host="*",
    operation="ALTER",
    permission_type="ALLOW",
    cluster_api_url=test_cluster.cluster_api_url,
    allow_deletion=True)
# The type of schema (AVRO, JSON, PROTOBUF)
schema_type = config.get("schemaType")
if schema_type is None:
    schema_type = "AVRO"
# The AVRO schema definition for user data
user_schema_definition = config.get("userSchemaDefinition")
if user_schema_definition is None:
    user_schema_definition = """{
  "type": "record",
  "name": "User",
  "fields": [
    {
      "name": "id",
      "type": "int"
    },
    {
      "name": "name",
      "type": "string"
    },
    {
      "name": "email",
      "type": "string"
    },
    {
      "name": "created_at",
      "type": "long",
      "logicalType": "timestamp-millis"
    }
  ]
}
"""
all_test_topic = redpanda.SchemaRegistryAcl("all_test_topic",
    cluster_id=test_cluster.id,
    principal=test_user.name.apply(lambda name: f"User:{name}"),
    resource_type="SUBJECT",
    resource_name_=f"{topic_name}-",
    pattern_type="PREFIXED",
    host="*",
    operation="ALL",
    permission="ALLOW",
    username=test_user.name,
    password=user_pw,
    allow_deletion=True,
    opts = pulumi.ResourceOptions(depends_on=[schema_registry_admin]))
describe_registry = redpanda.SchemaRegistryAcl("describe_registry",
    cluster_id=test_cluster.id,
    principal=test_user.name.apply(lambda name: f"User:{name}"),
    resource_type="REGISTRY",
    resource_name_="*",
    pattern_type="LITERAL",
    host="*",
    operation="DESCRIBE",
    permission="ALLOW",
    username=test_user.name,
    password=user_pw,
    allow_deletion=True,
    opts = pulumi.ResourceOptions(depends_on=[schema_registry_admin]))
alter_configs_registry = redpanda.SchemaRegistryAcl("alter_configs_registry",
    cluster_id=test_cluster.id,
    principal=test_user.name.apply(lambda name: f"User:{name}"),
    resource_type="REGISTRY",
    resource_name_="*",
    pattern_type="LITERAL",
    host="*",
    operation="ALTER_CONFIGS",
    permission="ALLOW",
    username=test_user.name,
    password=user_pw,
    allow_deletion=True,
    opts = pulumi.ResourceOptions(depends_on=[schema_registry_admin]))
user_schema = redpanda.Schema("user_schema",
    cluster_id=test_cluster.id,
    subject=f"{topic_name}-value",
    schema_type=schema_type,
    schema=user_schema_definition,
    username=test_user.name,
    password=user_pw,
    allow_deletion=True,
    opts = pulumi.ResourceOptions(depends_on=[
            schema_registry_admin,
            all_test_topic,
            describe_registry,
            alter_configs_registry,
        ]))
# The AVRO schema definition for user events that references the User schema
user_event_schema_definition = config.get("userEventSchemaDefinition")
if user_event_schema_definition is None:
    user_event_schema_definition = """{
  "type": "record",
  "name": "UserEvent",
  "fields": [
    {
      "name": "event_id",
      "type": "string"
    },
    {
      "name": "event_type",
      "type": {
        "type": "enum",
        "name": "EventType",
        "symbols": ["CREATED", "UPDATED", "DELETED"]
      }
    },
    {
      "name": "user",
      "type": "User"
    },
    {
      "name": "timestamp",
      "type": "long",
      "logicalType": "timestamp-millis"
    },
    {
      "name": "metadata",
      "type": ["null", {
        "type": "map",
        "values": "string"
      }],
      "default": null
    }
  ]
}
"""
user_event_schema = redpanda.Schema("user_event_schema",
    cluster_id=test_cluster.id,
    subject=f"{topic_name}-events-value",
    schema_type=schema_type,
    schema=user_event_schema_definition,
    username=test_user.name,
    password=user_pw,
    allow_deletion=True,
    references=[{
        "name": "User",
        "subject": user_schema.subject,
        "version": user_schema.version,
    }],
    opts = pulumi.ResourceOptions(depends_on=[
            schema_registry_admin,
            all_test_topic,
            describe_registry,
            alter_configs_registry,
        ]))
# The AVRO schema definition for product data with strict compatibility
product_schema_definition = config.get("productSchemaDefinition")
if product_schema_definition is None:
    product_schema_definition = """{
  "type": "record",
  "name": "Product",
  "fields": [
    {
      "name": "id",
      "type": "string"
    },
    {
      "name": "name",
      "type": "string"
    },
    {
      "name": "price",
      "type": {
        "type": "bytes",
        "logicalType": "decimal",
        "precision": 10,
        "scale": 2
      }
    },
    {
      "name": "category",
      "type": {
        "type": "enum",
        "name": "Category",
        "symbols": ["ELECTRONICS", "CLOTHING", "BOOKS", "HOME"]
      }
    },
    {
      "name": "description",
      "type": ["null", "string"],
      "default": null
    },
    {
      "name": "created_at",
      "type": "long",
      "logicalType": "timestamp-millis"
    }
  ]
}
"""
# The compatibility level for schema evolution (BACKWARD, BACKWARD_TRANSITIVE, FORWARD, FORWARD_TRANSITIVE, FULL, FULL_TRANSITIVE, NONE)
compatibility_level = config.get("compatibilityLevel")
if compatibility_level is None:
    compatibility_level = "FULL"
product_schema = redpanda.Schema("product_schema",
    cluster_id=test_cluster.id,
    subject=f"{topic_name}-product-value",
    schema_type=schema_type,
    schema=product_schema_definition,
    compatibility=compatibility_level,
    username=test_user.name,
    password=user_pw,
    allow_deletion=True,
    opts = pulumi.ResourceOptions(depends_on=[
            schema_registry_admin,
            all_test_topic,
            describe_registry,
            alter_configs_registry,
        ]))
cluster_admin = redpanda.Acl("cluster_admin",
    resource_type="CLUSTER",
    resource_name_="kafka-cluster",
    resource_pattern_type="LITERAL",
    principal=test_user.name.apply(lambda name: f"User:{name}"),
    host="*",
    operation="ALL",
    permission_type="ALLOW",
    cluster_api_url=test_cluster.cluster_api_url,
    allow_deletion=True)
read_product = redpanda.SchemaRegistryAcl("read_product",
    cluster_id=test_cluster.id,
    principal=test_user.name.apply(lambda name: f"User:{name}"),
    resource_type="SUBJECT",
    resource_name_="product-",
    pattern_type="PREFIXED",
    host="*",
    operation="READ",
    permission="ALLOW",
    username=test_user.name,
    password=user_pw,
    allow_deletion=True,
    opts = pulumi.ResourceOptions(depends_on=[schema_registry_admin]))
write_orders = redpanda.SchemaRegistryAcl("write_orders",
    cluster_id=test_cluster.id,
    principal=test_user.name.apply(lambda name: f"User:{name}"),
    resource_type="SUBJECT",
    resource_name_="orders-value",
    pattern_type="LITERAL",
    host="*",
    operation="WRITE",
    permission="ALLOW",
    username=test_user.name,
    password=user_pw,
    allow_deletion=True,
    opts = pulumi.ResourceOptions(depends_on=[schema_registry_admin]))
describe_test_topic = redpanda.SchemaRegistryAcl("describe_test_topic",
    cluster_id=test_cluster.id,
    principal=test_user.name.apply(lambda name: f"User:{name}"),
    resource_type="SUBJECT",
    resource_name_=f"{topic_name}-",
    pattern_type="PREFIXED",
    host="*",
    operation="DESCRIBE",
    permission="ALLOW",
    username=test_user.name,
    password=user_pw,
    allow_deletion=True,
    opts = pulumi.ResourceOptions(depends_on=[schema_registry_admin]))
pulumi.export("userSchemaInfo", {
    "id": user_schema.schema_id,
    "subject": user_schema.subject,
    "version": user_schema.version,
    "type": user_schema.schema_type,
})
pulumi.export("userEventSchemaInfo", {
    "id": user_event_schema.schema_id,
    "subject": user_event_schema.subject,
    "version": user_event_schema.version,
    "type": user_event_schema.schema_type,
    "references": user_event_schema.references,
})
pulumi.export("productSchemaInfo", {
    "id": product_schema.schema_id,
    "subject": product_schema.subject,
    "version": product_schema.version,
    "type": product_schema.schema_type,
    "compatibility": product_schema.compatibility,
})
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
using System.Text.Json;
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
        Timeouts = new Redpanda.Inputs.NetworkTimeoutsArgs
        {
            Create = "20m",
            Delete = "20m",
        },
    });

    var clusterName = config.Get("clusterName") ?? "testname";
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
        ClusterConfiguration = new Redpanda.Inputs.ClusterClusterConfigurationArgs
        {
            CustomPropertiesJson = JsonSerializer.Serialize(new Dictionary<string, object?>
            {
                ["schema_registry_enable_authorization"] = true,
            }),
        },
        Timeouts = new Redpanda.Inputs.ClusterTimeoutsArgs
        {
            Create = "90m",
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

    var schemaRegistryAdmin = new Redpanda.Acl("schema_registry_admin", new()
    {
        ResourceType = "CLUSTER",
        ResourceName = "kafka-cluster",
        ResourcePatternType = "LITERAL",
        Principal = testUser.Name.Apply(name => $"User:{name}"),
        Host = "*",
        Operation = "ALTER",
        PermissionType = "ALLOW",
        ClusterApiUrl = testCluster.ClusterApiUrl,
        AllowDeletion = true,
    });

    // The type of schema (AVRO, JSON, PROTOBUF)
    var schemaType = config.Get("schemaType") ?? "AVRO";
    // The AVRO schema definition for user data
    var userSchemaDefinition = config.Get("userSchemaDefinition") ?? @"{
  ""type"": ""record"",
  ""name"": ""User"",
  ""fields"": [
    {
      ""name"": ""id"",
      ""type"": ""int""
    },
    {
      ""name"": ""name"",
      ""type"": ""string""
    },
    {
      ""name"": ""email"",
      ""type"": ""string""
    },
    {
      ""name"": ""created_at"",
      ""type"": ""long"",
      ""logicalType"": ""timestamp-millis""
    }
  ]
}
";
    var allTestTopic = new Redpanda.SchemaRegistryAcl("all_test_topic", new()
    {
        ClusterId = testCluster.Id,
        Principal = testUser.Name.Apply(name => $"User:{name}"),
        ResourceType = "SUBJECT",
        ResourceName = $"{topicName}-",
        PatternType = "PREFIXED",
        Host = "*",
        Operation = "ALL",
        Permission = "ALLOW",
        Username = testUser.Name,
        Password = userPw,
        AllowDeletion = true,
    }, new CustomResourceOptions
    {
        DependsOn =
        {
            schemaRegistryAdmin,
        },
    });

    var describeRegistry = new Redpanda.SchemaRegistryAcl("describe_registry", new()
    {
        ClusterId = testCluster.Id,
        Principal = testUser.Name.Apply(name => $"User:{name}"),
        ResourceType = "REGISTRY",
        ResourceName = "*",
        PatternType = "LITERAL",
        Host = "*",
        Operation = "DESCRIBE",
        Permission = "ALLOW",
        Username = testUser.Name,
        Password = userPw,
        AllowDeletion = true,
    }, new CustomResourceOptions
    {
        DependsOn =
        {
            schemaRegistryAdmin,
        },
    });

    var alterConfigsRegistry = new Redpanda.SchemaRegistryAcl("alter_configs_registry", new()
    {
        ClusterId = testCluster.Id,
        Principal = testUser.Name.Apply(name => $"User:{name}"),
        ResourceType = "REGISTRY",
        ResourceName = "*",
        PatternType = "LITERAL",
        Host = "*",
        Operation = "ALTER_CONFIGS",
        Permission = "ALLOW",
        Username = testUser.Name,
        Password = userPw,
        AllowDeletion = true,
    }, new CustomResourceOptions
    {
        DependsOn =
        {
            schemaRegistryAdmin,
        },
    });

    var userSchema = new Redpanda.Schema("user_schema", new()
    {
        ClusterId = testCluster.Id,
        Subject = $"{topicName}-value",
        SchemaType = schemaType,
        Schema = userSchemaDefinition,
        Username = testUser.Name,
        Password = userPw,
        AllowDeletion = true,
    }, new CustomResourceOptions
    {
        DependsOn =
        {
            schemaRegistryAdmin,
            allTestTopic,
            describeRegistry,
            alterConfigsRegistry,
        },
    });

    // The AVRO schema definition for user events that references the User schema
    var userEventSchemaDefinition = config.Get("userEventSchemaDefinition") ?? @"{
  ""type"": ""record"",
  ""name"": ""UserEvent"",
  ""fields"": [
    {
      ""name"": ""event_id"",
      ""type"": ""string""
    },
    {
      ""name"": ""event_type"",
      ""type"": {
        ""type"": ""enum"",
        ""name"": ""EventType"",
        ""symbols"": [""CREATED"", ""UPDATED"", ""DELETED""]
      }
    },
    {
      ""name"": ""user"",
      ""type"": ""User""
    },
    {
      ""name"": ""timestamp"",
      ""type"": ""long"",
      ""logicalType"": ""timestamp-millis""
    },
    {
      ""name"": ""metadata"",
      ""type"": [""null"", {
        ""type"": ""map"",
        ""values"": ""string""
      }],
      ""default"": null
    }
  ]
}
";
    var userEventSchema = new Redpanda.Schema("user_event_schema", new()
    {
        ClusterId = testCluster.Id,
        Subject = $"{topicName}-events-value",
        SchemaType = schemaType,
        Schema = userEventSchemaDefinition,
        Username = testUser.Name,
        Password = userPw,
        AllowDeletion = true,
        References = new[]
        {
            new Redpanda.Inputs.SchemaReferenceArgs
            {
                Name = "User",
                Subject = userSchema.Subject,
                Version = userSchema.Version,
            },
        },
    }, new CustomResourceOptions
    {
        DependsOn =
        {
            schemaRegistryAdmin,
            allTestTopic,
            describeRegistry,
            alterConfigsRegistry,
        },
    });

    // The AVRO schema definition for product data with strict compatibility
    var productSchemaDefinition = config.Get("productSchemaDefinition") ?? @"{
  ""type"": ""record"",
  ""name"": ""Product"",
  ""fields"": [
    {
      ""name"": ""id"",
      ""type"": ""string""
    },
    {
      ""name"": ""name"",
      ""type"": ""string""
    },
    {
      ""name"": ""price"",
      ""type"": {
        ""type"": ""bytes"",
        ""logicalType"": ""decimal"",
        ""precision"": 10,
        ""scale"": 2
      }
    },
    {
      ""name"": ""category"",
      ""type"": {
        ""type"": ""enum"",
        ""name"": ""Category"",
        ""symbols"": [""ELECTRONICS"", ""CLOTHING"", ""BOOKS"", ""HOME""]
      }
    },
    {
      ""name"": ""description"",
      ""type"": [""null"", ""string""],
      ""default"": null
    },
    {
      ""name"": ""created_at"",
      ""type"": ""long"",
      ""logicalType"": ""timestamp-millis""
    }
  ]
}
";
    // The compatibility level for schema evolution (BACKWARD, BACKWARD_TRANSITIVE, FORWARD, FORWARD_TRANSITIVE, FULL, FULL_TRANSITIVE, NONE)
    var compatibilityLevel = config.Get("compatibilityLevel") ?? "FULL";
    var productSchema = new Redpanda.Schema("product_schema", new()
    {
        ClusterId = testCluster.Id,
        Subject = $"{topicName}-product-value",
        SchemaType = schemaType,
        Schema = productSchemaDefinition,
        Compatibility = compatibilityLevel,
        Username = testUser.Name,
        Password = userPw,
        AllowDeletion = true,
    }, new CustomResourceOptions
    {
        DependsOn =
        {
            schemaRegistryAdmin,
            allTestTopic,
            describeRegistry,
            alterConfigsRegistry,
        },
    });

    var clusterAdmin = new Redpanda.Acl("cluster_admin", new()
    {
        ResourceType = "CLUSTER",
        ResourceName = "kafka-cluster",
        ResourcePatternType = "LITERAL",
        Principal = testUser.Name.Apply(name => $"User:{name}"),
        Host = "*",
        Operation = "ALL",
        PermissionType = "ALLOW",
        ClusterApiUrl = testCluster.ClusterApiUrl,
        AllowDeletion = true,
    });

    var readProduct = new Redpanda.SchemaRegistryAcl("read_product", new()
    {
        ClusterId = testCluster.Id,
        Principal = testUser.Name.Apply(name => $"User:{name}"),
        ResourceType = "SUBJECT",
        ResourceName = "product-",
        PatternType = "PREFIXED",
        Host = "*",
        Operation = "READ",
        Permission = "ALLOW",
        Username = testUser.Name,
        Password = userPw,
        AllowDeletion = true,
    }, new CustomResourceOptions
    {
        DependsOn =
        {
            schemaRegistryAdmin,
        },
    });

    var writeOrders = new Redpanda.SchemaRegistryAcl("write_orders", new()
    {
        ClusterId = testCluster.Id,
        Principal = testUser.Name.Apply(name => $"User:{name}"),
        ResourceType = "SUBJECT",
        ResourceName = "orders-value",
        PatternType = "LITERAL",
        Host = "*",
        Operation = "WRITE",
        Permission = "ALLOW",
        Username = testUser.Name,
        Password = userPw,
        AllowDeletion = true,
    }, new CustomResourceOptions
    {
        DependsOn =
        {
            schemaRegistryAdmin,
        },
    });

    var describeTestTopic = new Redpanda.SchemaRegistryAcl("describe_test_topic", new()
    {
        ClusterId = testCluster.Id,
        Principal = testUser.Name.Apply(name => $"User:{name}"),
        ResourceType = "SUBJECT",
        ResourceName = $"{topicName}-",
        PatternType = "PREFIXED",
        Host = "*",
        Operation = "DESCRIBE",
        Permission = "ALLOW",
        Username = testUser.Name,
        Password = userPw,
        AllowDeletion = true,
    }, new CustomResourceOptions
    {
        DependsOn =
        {
            schemaRegistryAdmin,
        },
    });

    return new Dictionary<string, object?>
    {
        ["userSchemaInfo"] =
        {
            { "id", userSchema.SchemaId },
            { "subject", userSchema.Subject },
            { "version", userSchema.Version },
            { "type", userSchema.SchemaType },
        },
        ["userEventSchemaInfo"] =
        {
            { "id", userEventSchema.SchemaId },
            { "subject", userEventSchema.Subject },
            { "version", userEventSchema.Version },
            { "type", userEventSchema.SchemaType },
            { "references", userEventSchema.References },
        },
        ["productSchemaInfo"] =
        {
            { "id", productSchema.SchemaId },
            { "subject", productSchema.Subject },
            { "version", productSchema.Version },
            { "type", productSchema.SchemaType },
            { "compatibility", productSchema.Compatibility },
        },
    };
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
	"encoding/json"
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
			Timeouts: &redpanda.NetworkTimeoutsArgs{
				Create: pulumi.String("20m"),
				Delete: pulumi.String("20m"),
			},
		})
		if err != nil {
			return err
		}
		clusterName := "testname"
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
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"schema_registry_enable_authorization": true,
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
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
			ClusterConfiguration: &redpanda.ClusterClusterConfigurationArgs{
				CustomPropertiesJson: pulumi.String(json0),
			},
			Timeouts: &redpanda.ClusterTimeoutsArgs{
				Create: pulumi.String("90m"),
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
		_, err = redpanda.NewTopic(ctx, "test", &redpanda.TopicArgs{
			Name:              pulumi.String(topicName),
			PartitionCount:    pulumi.Float64(partitionCount),
			ReplicationFactor: pulumi.Float64(replicationFactor),
			ClusterApiUrl:     testCluster.ClusterApiUrl,
			AllowDeletion:     pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		schemaRegistryAdmin, err := redpanda.NewAcl(ctx, "schema_registry_admin", &redpanda.AclArgs{
			ResourceType:        pulumi.String("CLUSTER"),
			ResourceName:        pulumi.String("kafka-cluster"),
			ResourcePatternType: pulumi.String("LITERAL"),
			Principal: testUser.Name.ApplyT(func(name string) (string, error) {
				return fmt.Sprintf("User:%v", name), nil
			}).(pulumi.StringOutput),
			Host:           pulumi.String("*"),
			Operation:      pulumi.String("ALTER"),
			PermissionType: pulumi.String("ALLOW"),
			ClusterApiUrl:  testCluster.ClusterApiUrl,
			AllowDeletion:  pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		// The type of schema (AVRO, JSON, PROTOBUF)
		schemaType := "AVRO"
		if param := cfg.Get("schemaType"); param != "" {
			schemaType = param
		}
		// The AVRO schema definition for user data
		userSchemaDefinition := `{
  "type": "record",
  "name": "User",
  "fields": [
    {
      "name": "id",
      "type": "int"
    },
    {
      "name": "name",
      "type": "string"
    },
    {
      "name": "email",
      "type": "string"
    },
    {
      "name": "created_at",
      "type": "long",
      "logicalType": "timestamp-millis"
    }
  ]
}
`
		if param := cfg.Get("userSchemaDefinition"); param != "" {
			userSchemaDefinition = param
		}
		allTestTopic, err := redpanda.NewSchemaRegistryAcl(ctx, "all_test_topic", &redpanda.SchemaRegistryAclArgs{
			ClusterId: testCluster.ID(),
			Principal: testUser.Name.ApplyT(func(name string) (string, error) {
				return fmt.Sprintf("User:%v", name), nil
			}).(pulumi.StringOutput),
			ResourceType:  pulumi.String("SUBJECT"),
			ResourceName:  pulumi.Sprintf("%v-", topicName),
			PatternType:   pulumi.String("PREFIXED"),
			Host:          pulumi.String("*"),
			Operation:     pulumi.String("ALL"),
			Permission:    pulumi.String("ALLOW"),
			Username:      testUser.Name,
			Password:      pulumi.String(userPw),
			AllowDeletion: pulumi.Bool(true),
		}, pulumi.DependsOn([]pulumi.Resource{
			schemaRegistryAdmin,
		}))
		if err != nil {
			return err
		}
		describeRegistry, err := redpanda.NewSchemaRegistryAcl(ctx, "describe_registry", &redpanda.SchemaRegistryAclArgs{
			ClusterId: testCluster.ID(),
			Principal: testUser.Name.ApplyT(func(name string) (string, error) {
				return fmt.Sprintf("User:%v", name), nil
			}).(pulumi.StringOutput),
			ResourceType:  pulumi.String("REGISTRY"),
			ResourceName:  pulumi.String("*"),
			PatternType:   pulumi.String("LITERAL"),
			Host:          pulumi.String("*"),
			Operation:     pulumi.String("DESCRIBE"),
			Permission:    pulumi.String("ALLOW"),
			Username:      testUser.Name,
			Password:      pulumi.String(userPw),
			AllowDeletion: pulumi.Bool(true),
		}, pulumi.DependsOn([]pulumi.Resource{
			schemaRegistryAdmin,
		}))
		if err != nil {
			return err
		}
		alterConfigsRegistry, err := redpanda.NewSchemaRegistryAcl(ctx, "alter_configs_registry", &redpanda.SchemaRegistryAclArgs{
			ClusterId: testCluster.ID(),
			Principal: testUser.Name.ApplyT(func(name string) (string, error) {
				return fmt.Sprintf("User:%v", name), nil
			}).(pulumi.StringOutput),
			ResourceType:  pulumi.String("REGISTRY"),
			ResourceName:  pulumi.String("*"),
			PatternType:   pulumi.String("LITERAL"),
			Host:          pulumi.String("*"),
			Operation:     pulumi.String("ALTER_CONFIGS"),
			Permission:    pulumi.String("ALLOW"),
			Username:      testUser.Name,
			Password:      pulumi.String(userPw),
			AllowDeletion: pulumi.Bool(true),
		}, pulumi.DependsOn([]pulumi.Resource{
			schemaRegistryAdmin,
		}))
		if err != nil {
			return err
		}
		userSchema, err := redpanda.NewSchema(ctx, "user_schema", &redpanda.SchemaArgs{
			ClusterId:     testCluster.ID(),
			Subject:       pulumi.Sprintf("%v-value", topicName),
			SchemaType:    pulumi.String(schemaType),
			Schema:        pulumi.String(userSchemaDefinition),
			Username:      testUser.Name,
			Password:      pulumi.String(userPw),
			AllowDeletion: pulumi.Bool(true),
		}, pulumi.DependsOn([]pulumi.Resource{
			schemaRegistryAdmin,
			allTestTopic,
			describeRegistry,
			alterConfigsRegistry,
		}))
		if err != nil {
			return err
		}
		// The AVRO schema definition for user events that references the User schema
		userEventSchemaDefinition := `{
  "type": "record",
  "name": "UserEvent",
  "fields": [
    {
      "name": "event_id",
      "type": "string"
    },
    {
      "name": "event_type",
      "type": {
        "type": "enum",
        "name": "EventType",
        "symbols": ["CREATED", "UPDATED", "DELETED"]
      }
    },
    {
      "name": "user",
      "type": "User"
    },
    {
      "name": "timestamp",
      "type": "long",
      "logicalType": "timestamp-millis"
    },
    {
      "name": "metadata",
      "type": ["null", {
        "type": "map",
        "values": "string"
      }],
      "default": null
    }
  ]
}
`
		if param := cfg.Get("userEventSchemaDefinition"); param != "" {
			userEventSchemaDefinition = param
		}
		userEventSchema, err := redpanda.NewSchema(ctx, "user_event_schema", &redpanda.SchemaArgs{
			ClusterId:     testCluster.ID(),
			Subject:       pulumi.Sprintf("%v-events-value", topicName),
			SchemaType:    pulumi.String(schemaType),
			Schema:        pulumi.String(userEventSchemaDefinition),
			Username:      testUser.Name,
			Password:      pulumi.String(userPw),
			AllowDeletion: pulumi.Bool(true),
			References: redpanda.SchemaReferenceArray{
				&redpanda.SchemaReferenceArgs{
					Name:    pulumi.String("User"),
					Subject: userSchema.Subject,
					Version: userSchema.Version,
				},
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			schemaRegistryAdmin,
			allTestTopic,
			describeRegistry,
			alterConfigsRegistry,
		}))
		if err != nil {
			return err
		}
		// The AVRO schema definition for product data with strict compatibility
		productSchemaDefinition := `{
  "type": "record",
  "name": "Product",
  "fields": [
    {
      "name": "id",
      "type": "string"
    },
    {
      "name": "name",
      "type": "string"
    },
    {
      "name": "price",
      "type": {
        "type": "bytes",
        "logicalType": "decimal",
        "precision": 10,
        "scale": 2
      }
    },
    {
      "name": "category",
      "type": {
        "type": "enum",
        "name": "Category",
        "symbols": ["ELECTRONICS", "CLOTHING", "BOOKS", "HOME"]
      }
    },
    {
      "name": "description",
      "type": ["null", "string"],
      "default": null
    },
    {
      "name": "created_at",
      "type": "long",
      "logicalType": "timestamp-millis"
    }
  ]
}
`
		if param := cfg.Get("productSchemaDefinition"); param != "" {
			productSchemaDefinition = param
		}
		// The compatibility level for schema evolution (BACKWARD, BACKWARD_TRANSITIVE, FORWARD, FORWARD_TRANSITIVE, FULL, FULL_TRANSITIVE, NONE)
		compatibilityLevel := "FULL"
		if param := cfg.Get("compatibilityLevel"); param != "" {
			compatibilityLevel = param
		}
		productSchema, err := redpanda.NewSchema(ctx, "product_schema", &redpanda.SchemaArgs{
			ClusterId:     testCluster.ID(),
			Subject:       pulumi.Sprintf("%v-product-value", topicName),
			SchemaType:    pulumi.String(schemaType),
			Schema:        pulumi.String(productSchemaDefinition),
			Compatibility: pulumi.String(compatibilityLevel),
			Username:      testUser.Name,
			Password:      pulumi.String(userPw),
			AllowDeletion: pulumi.Bool(true),
		}, pulumi.DependsOn([]pulumi.Resource{
			schemaRegistryAdmin,
			allTestTopic,
			describeRegistry,
			alterConfigsRegistry,
		}))
		if err != nil {
			return err
		}
		_, err = redpanda.NewAcl(ctx, "cluster_admin", &redpanda.AclArgs{
			ResourceType:        pulumi.String("CLUSTER"),
			ResourceName:        pulumi.String("kafka-cluster"),
			ResourcePatternType: pulumi.String("LITERAL"),
			Principal: testUser.Name.ApplyT(func(name string) (string, error) {
				return fmt.Sprintf("User:%v", name), nil
			}).(pulumi.StringOutput),
			Host:           pulumi.String("*"),
			Operation:      pulumi.String("ALL"),
			PermissionType: pulumi.String("ALLOW"),
			ClusterApiUrl:  testCluster.ClusterApiUrl,
			AllowDeletion:  pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		_, err = redpanda.NewSchemaRegistryAcl(ctx, "read_product", &redpanda.SchemaRegistryAclArgs{
			ClusterId: testCluster.ID(),
			Principal: testUser.Name.ApplyT(func(name string) (string, error) {
				return fmt.Sprintf("User:%v", name), nil
			}).(pulumi.StringOutput),
			ResourceType:  pulumi.String("SUBJECT"),
			ResourceName:  pulumi.String("product-"),
			PatternType:   pulumi.String("PREFIXED"),
			Host:          pulumi.String("*"),
			Operation:     pulumi.String("READ"),
			Permission:    pulumi.String("ALLOW"),
			Username:      testUser.Name,
			Password:      pulumi.String(userPw),
			AllowDeletion: pulumi.Bool(true),
		}, pulumi.DependsOn([]pulumi.Resource{
			schemaRegistryAdmin,
		}))
		if err != nil {
			return err
		}
		_, err = redpanda.NewSchemaRegistryAcl(ctx, "write_orders", &redpanda.SchemaRegistryAclArgs{
			ClusterId: testCluster.ID(),
			Principal: testUser.Name.ApplyT(func(name string) (string, error) {
				return fmt.Sprintf("User:%v", name), nil
			}).(pulumi.StringOutput),
			ResourceType:  pulumi.String("SUBJECT"),
			ResourceName:  pulumi.String("orders-value"),
			PatternType:   pulumi.String("LITERAL"),
			Host:          pulumi.String("*"),
			Operation:     pulumi.String("WRITE"),
			Permission:    pulumi.String("ALLOW"),
			Username:      testUser.Name,
			Password:      pulumi.String(userPw),
			AllowDeletion: pulumi.Bool(true),
		}, pulumi.DependsOn([]pulumi.Resource{
			schemaRegistryAdmin,
		}))
		if err != nil {
			return err
		}
		_, err = redpanda.NewSchemaRegistryAcl(ctx, "describe_test_topic", &redpanda.SchemaRegistryAclArgs{
			ClusterId: testCluster.ID(),
			Principal: testUser.Name.ApplyT(func(name string) (string, error) {
				return fmt.Sprintf("User:%v", name), nil
			}).(pulumi.StringOutput),
			ResourceType:  pulumi.String("SUBJECT"),
			ResourceName:  pulumi.Sprintf("%v-", topicName),
			PatternType:   pulumi.String("PREFIXED"),
			Host:          pulumi.String("*"),
			Operation:     pulumi.String("DESCRIBE"),
			Permission:    pulumi.String("ALLOW"),
			Username:      testUser.Name,
			Password:      pulumi.String(userPw),
			AllowDeletion: pulumi.Bool(true),
		}, pulumi.DependsOn([]pulumi.Resource{
			schemaRegistryAdmin,
		}))
		if err != nil {
			return err
		}
		ctx.Export("userSchemaInfo", pulumi.Map{
			"id":      userSchema.SchemaId,
			"subject": userSchema.Subject,
			"version": userSchema.Version,
			"type":    userSchema.SchemaType,
		})
		ctx.Export("userEventSchemaInfo", pulumi.Map{
			"id":         userEventSchema.SchemaId,
			"subject":    userEventSchema.Subject,
			"version":    userEventSchema.Version,
			"type":       userEventSchema.SchemaType,
			"references": userEventSchema.References,
		})
		ctx.Export("productSchemaInfo", pulumi.Map{
			"id":            productSchema.SchemaId,
			"subject":       productSchema.Subject,
			"version":       productSchema.Version,
			"type":          productSchema.SchemaType,
			"compatibility": productSchema.Compatibility,
		})
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
    default: testname
  resourceGroupName:
    type: string
    default: testname
  networkName:
    type: string
    default: testname
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
  schemaType:
    type: string
    default: AVRO
  userSchemaDefinition:
    type: string
    default: |
      {
        "type": "record",
        "name": "User",
        "fields": [
          {
            "name": "id",
            "type": "int"
          },
          {
            "name": "name",
            "type": "string"
          },
          {
            "name": "email",
            "type": "string"
          },
          {
            "name": "created_at",
            "type": "long",
            "logicalType": "timestamp-millis"
          }
        ]
      }
  userEventSchemaDefinition:
    type: string
    default: |
      {
        "type": "record",
        "name": "UserEvent",
        "fields": [
          {
            "name": "event_id",
            "type": "string"
          },
          {
            "name": "event_type",
            "type": {
              "type": "enum",
              "name": "EventType",
              "symbols": ["CREATED", "UPDATED", "DELETED"]
            }
          },
          {
            "name": "user",
            "type": "User"
          },
          {
            "name": "timestamp",
            "type": "long",
            "logicalType": "timestamp-millis"
          },
          {
            "name": "metadata",
            "type": ["null", {
              "type": "map",
              "values": "string"
            }],
            "default": null
          }
        ]
      }
  productSchemaDefinition:
    type: string
    default: |
      {
        "type": "record",
        "name": "Product",
        "fields": [
          {
            "name": "id",
            "type": "string"
          },
          {
            "name": "name",
            "type": "string"
          },
          {
            "name": "price",
            "type": {
              "type": "bytes",
              "logicalType": "decimal",
              "precision": 10,
              "scale": 2
            }
          },
          {
            "name": "category",
            "type": {
              "type": "enum",
              "name": "Category",
              "symbols": ["ELECTRONICS", "CLOTHING", "BOOKS", "HOME"]
            }
          },
          {
            "name": "description",
            "type": ["null", "string"],
            "default": null
          },
          {
            "name": "created_at",
            "type": "long",
            "logicalType": "timestamp-millis"
          }
        ]
      }
  compatibilityLevel:
    type: string
    default: FULL
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
      timeouts:
        create: 20m
        delete: 20m
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
      clusterConfiguration:
        customPropertiesJson:
          fn::toJSON:
            schema_registry_enable_authorization: true
      timeouts:
        create: 90m
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
  userSchema:
    type: redpanda:Schema
    name: user_schema
    properties:
      clusterId: ${testCluster.id}
      subject: ${topicName}-value
      schemaType: ${schemaType}
      schema: ${userSchemaDefinition}
      username: ${testUser.name}
      password: ${userPw}
      allowDeletion: true
    options:
      dependsOn:
        - ${schemaRegistryAdmin}
        - ${allTestTopic}
        - ${describeRegistry}
        - ${alterConfigsRegistry}
  userEventSchema:
    type: redpanda:Schema
    name: user_event_schema
    properties:
      clusterId: ${testCluster.id}
      subject: ${topicName}-events-value
      schemaType: ${schemaType}
      schema: ${userEventSchemaDefinition}
      username: ${testUser.name}
      password: ${userPw}
      allowDeletion: true
      references:
        - name: User
          subject: ${userSchema.subject}
          version: ${userSchema.version}
    options:
      dependsOn:
        - ${schemaRegistryAdmin}
        - ${allTestTopic}
        - ${describeRegistry}
        - ${alterConfigsRegistry}
  productSchema:
    type: redpanda:Schema
    name: product_schema
    properties:
      clusterId: ${testCluster.id}
      subject: ${topicName}-product-value
      schemaType: ${schemaType}
      schema: ${productSchemaDefinition}
      compatibility: ${compatibilityLevel}
      username: ${testUser.name}
      password: ${userPw}
      allowDeletion: true
    options:
      dependsOn:
        - ${schemaRegistryAdmin}
        - ${allTestTopic}
        - ${describeRegistry}
        - ${alterConfigsRegistry}
  clusterAdmin:
    type: redpanda:Acl
    name: cluster_admin
    properties:
      resourceType: CLUSTER
      resourceName: kafka-cluster
      resourcePatternType: LITERAL
      principal: User:${testUser.name}
      host: '*'
      operation: ALL
      permissionType: ALLOW
      clusterApiUrl: ${testCluster.clusterApiUrl}
      allowDeletion: true
  schemaRegistryAdmin:
    type: redpanda:Acl
    name: schema_registry_admin
    properties:
      resourceType: CLUSTER
      resourceName: kafka-cluster
      resourcePatternType: LITERAL
      principal: User:${testUser.name}
      host: '*'
      operation: ALTER
      permissionType: ALLOW
      clusterApiUrl: ${testCluster.clusterApiUrl}
      allowDeletion: true
  readProduct:
    type: redpanda:SchemaRegistryAcl
    name: read_product
    properties:
      clusterId: ${testCluster.id}
      principal: User:${testUser.name}
      resourceType: SUBJECT
      resourceName: product-
      patternType: PREFIXED
      host: '*'
      operation: READ
      permission: ALLOW
      username: ${testUser.name}
      password: ${userPw}
      allowDeletion: true
    options:
      dependsOn:
        - ${schemaRegistryAdmin}
  writeOrders:
    type: redpanda:SchemaRegistryAcl
    name: write_orders
    properties:
      clusterId: ${testCluster.id}
      principal: User:${testUser.name}
      resourceType: SUBJECT
      resourceName: orders-value
      patternType: LITERAL
      host: '*'
      operation: WRITE
      permission: ALLOW
      username: ${testUser.name}
      password: ${userPw}
      allowDeletion: true
    options:
      dependsOn:
        - ${schemaRegistryAdmin}
  allTestTopic:
    type: redpanda:SchemaRegistryAcl
    name: all_test_topic
    properties:
      clusterId: ${testCluster.id}
      principal: User:${testUser.name}
      resourceType: SUBJECT
      resourceName: ${topicName}-
      patternType: PREFIXED
      host: '*'
      operation: ALL
      permission: ALLOW
      username: ${testUser.name}
      password: ${userPw}
      allowDeletion: true
    options:
      dependsOn:
        - ${schemaRegistryAdmin}
  describeTestTopic:
    type: redpanda:SchemaRegistryAcl
    name: describe_test_topic
    properties:
      clusterId: ${testCluster.id}
      principal: User:${testUser.name}
      resourceType: SUBJECT
      resourceName: ${topicName}-
      patternType: PREFIXED
      host: '*'
      operation: DESCRIBE
      permission: ALLOW
      username: ${testUser.name}
      password: ${userPw}
      allowDeletion: true
    options:
      dependsOn:
        - ${schemaRegistryAdmin}
  describeRegistry:
    type: redpanda:SchemaRegistryAcl
    name: describe_registry
    properties:
      clusterId: ${testCluster.id}
      principal: User:${testUser.name}
      resourceType: REGISTRY
      resourceName: '*'
      patternType: LITERAL
      host: '*'
      operation: DESCRIBE
      permission: ALLOW
      username: ${testUser.name}
      password: ${userPw}
      allowDeletion: true
    options:
      dependsOn:
        - ${schemaRegistryAdmin}
  alterConfigsRegistry:
    type: redpanda:SchemaRegistryAcl
    name: alter_configs_registry
    properties:
      clusterId: ${testCluster.id}
      principal: User:${testUser.name}
      resourceType: REGISTRY
      resourceName: '*'
      patternType: LITERAL
      host: '*'
      operation: ALTER_CONFIGS
      permission: ALLOW
      username: ${testUser.name}
      password: ${userPw}
      allowDeletion: true
    options:
      dependsOn:
        - ${schemaRegistryAdmin}
outputs:
  userSchemaInfo:
    id: ${userSchema.schemaId}
    subject: ${userSchema.subject}
    version: ${userSchema.version}
    type: ${userSchema.schemaType}
  userEventSchemaInfo:
    id: ${userEventSchema.schemaId}
    subject: ${userEventSchema.subject}
    version: ${userEventSchema.version}
    type: ${userEventSchema.schemaType}
    references: ${userEventSchema.references}
  productSchemaInfo:
    id: ${productSchema.schemaId}
    subject: ${productSchema.subject}
    version: ${productSchema.version}
    type: ${productSchema.schemaType}
    compatibility: ${productSchema.compatibility}
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
import com.pulumi.redpanda.inputs.NetworkTimeoutsArgs;
import com.pulumi.redpanda.Cluster;
import com.pulumi.redpanda.ClusterArgs;
import com.pulumi.redpanda.inputs.ClusterClusterConfigurationArgs;
import com.pulumi.redpanda.inputs.ClusterTimeoutsArgs;
import com.pulumi.redpanda.User;
import com.pulumi.redpanda.UserArgs;
import com.pulumi.redpanda.Topic;
import com.pulumi.redpanda.TopicArgs;
import com.pulumi.redpanda.Acl;
import com.pulumi.redpanda.AclArgs;
import com.pulumi.redpanda.SchemaRegistryAcl;
import com.pulumi.redpanda.SchemaRegistryAclArgs;
import com.pulumi.redpanda.Schema;
import com.pulumi.redpanda.SchemaArgs;
import com.pulumi.redpanda.inputs.SchemaReferenceArgs;
import static com.pulumi.codegen.internal.Serialization.*;
import com.pulumi.resources.CustomResourceOptions;
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
        final var region = config.get("region").orElse("us-central1");
        final var cloudProvider = config.get("cloudProvider").orElse("gcp");
        var testNetwork = new Network("testNetwork", NetworkArgs.builder()
            .name(networkName)
            .resourceGroupId(test.id())
            .cloudProvider(cloudProvider)
            .region(region)
            .clusterType("dedicated")
            .cidrBlock("10.0.0.0/20")
            .timeouts(NetworkTimeoutsArgs.builder()
                .create("20m")
                .delete("20m")
                .build())
            .build());

        final var clusterName = config.get("clusterName").orElse("testname");
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
            .clusterConfiguration(ClusterClusterConfigurationArgs.builder()
                .customPropertiesJson(serializeJson(
                    jsonObject(
                        jsonProperty("schema_registry_enable_authorization", true)
                    )))
                .build())
            .timeouts(ClusterTimeoutsArgs.builder()
                .create("90m")
                .build())
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

        var schemaRegistryAdmin = new Acl("schemaRegistryAdmin", AclArgs.builder()
            .resourceType("CLUSTER")
            .resourceName("kafka-cluster")
            .resourcePatternType("LITERAL")
            .principal(testUser.name().applyValue(name -> String.format("User:%s", name)))
            .host("*")
            .operation("ALTER")
            .permissionType("ALLOW")
            .clusterApiUrl(testCluster.clusterApiUrl())
            .allowDeletion(true)
            .build());

        final var schemaType = config.get("schemaType").orElse("AVRO");
        final var userSchemaDefinition = config.get("userSchemaDefinition").orElse("""
{
  "type": "record",
  "name": "User",
  "fields": [
    {
      "name": "id",
      "type": "int"
    },
    {
      "name": "name",
      "type": "string"
    },
    {
      "name": "email",
      "type": "string"
    },
    {
      "name": "created_at",
      "type": "long",
      "logicalType": "timestamp-millis"
    }
  ]
}
        """);
        var allTestTopic = new SchemaRegistryAcl("allTestTopic", SchemaRegistryAclArgs.builder()
            .clusterId(testCluster.id())
            .principal(testUser.name().applyValue(name -> String.format("User:%s", name)))
            .resourceType("SUBJECT")
            .resourceName(String.format("%s-", topicName))
            .patternType("PREFIXED")
            .host("*")
            .operation("ALL")
            .permission("ALLOW")
            .username(testUser.name())
            .password(userPw)
            .allowDeletion(true)
            .build(), CustomResourceOptions.builder()
                .dependsOn(schemaRegistryAdmin)
                .build());

        var describeRegistry = new SchemaRegistryAcl("describeRegistry", SchemaRegistryAclArgs.builder()
            .clusterId(testCluster.id())
            .principal(testUser.name().applyValue(name -> String.format("User:%s", name)))
            .resourceType("REGISTRY")
            .resourceName("*")
            .patternType("LITERAL")
            .host("*")
            .operation("DESCRIBE")
            .permission("ALLOW")
            .username(testUser.name())
            .password(userPw)
            .allowDeletion(true)
            .build(), CustomResourceOptions.builder()
                .dependsOn(schemaRegistryAdmin)
                .build());

        var alterConfigsRegistry = new SchemaRegistryAcl("alterConfigsRegistry", SchemaRegistryAclArgs.builder()
            .clusterId(testCluster.id())
            .principal(testUser.name().applyValue(name -> String.format("User:%s", name)))
            .resourceType("REGISTRY")
            .resourceName("*")
            .patternType("LITERAL")
            .host("*")
            .operation("ALTER_CONFIGS")
            .permission("ALLOW")
            .username(testUser.name())
            .password(userPw)
            .allowDeletion(true)
            .build(), CustomResourceOptions.builder()
                .dependsOn(schemaRegistryAdmin)
                .build());

        var userSchema = new Schema("userSchema", SchemaArgs.builder()
            .clusterId(testCluster.id())
            .subject(String.format("%s-value", topicName))
            .schemaType(schemaType)
            .schema(userSchemaDefinition)
            .username(testUser.name())
            .password(userPw)
            .allowDeletion(true)
            .build(), CustomResourceOptions.builder()
                .dependsOn(
                    schemaRegistryAdmin,
                    allTestTopic,
                    describeRegistry,
                    alterConfigsRegistry)
                .build());

        final var userEventSchemaDefinition = config.get("userEventSchemaDefinition").orElse("""
{
  "type": "record",
  "name": "UserEvent",
  "fields": [
    {
      "name": "event_id",
      "type": "string"
    },
    {
      "name": "event_type",
      "type": {
        "type": "enum",
        "name": "EventType",
        "symbols": ["CREATED", "UPDATED", "DELETED"]
      }
    },
    {
      "name": "user",
      "type": "User"
    },
    {
      "name": "timestamp",
      "type": "long",
      "logicalType": "timestamp-millis"
    },
    {
      "name": "metadata",
      "type": ["null", {
        "type": "map",
        "values": "string"
      }],
      "default": null
    }
  ]
}
        """);
        var userEventSchema = new Schema("userEventSchema", SchemaArgs.builder()
            .clusterId(testCluster.id())
            .subject(String.format("%s-events-value", topicName))
            .schemaType(schemaType)
            .schema(userEventSchemaDefinition)
            .username(testUser.name())
            .password(userPw)
            .allowDeletion(true)
            .references(SchemaReferenceArgs.builder()
                .name("User")
                .subject(userSchema.subject())
                .version(userSchema.version())
                .build())
            .build(), CustomResourceOptions.builder()
                .dependsOn(
                    schemaRegistryAdmin,
                    allTestTopic,
                    describeRegistry,
                    alterConfigsRegistry)
                .build());

        final var productSchemaDefinition = config.get("productSchemaDefinition").orElse("""
{
  "type": "record",
  "name": "Product",
  "fields": [
    {
      "name": "id",
      "type": "string"
    },
    {
      "name": "name",
      "type": "string"
    },
    {
      "name": "price",
      "type": {
        "type": "bytes",
        "logicalType": "decimal",
        "precision": 10,
        "scale": 2
      }
    },
    {
      "name": "category",
      "type": {
        "type": "enum",
        "name": "Category",
        "symbols": ["ELECTRONICS", "CLOTHING", "BOOKS", "HOME"]
      }
    },
    {
      "name": "description",
      "type": ["null", "string"],
      "default": null
    },
    {
      "name": "created_at",
      "type": "long",
      "logicalType": "timestamp-millis"
    }
  ]
}
        """);
        final var compatibilityLevel = config.get("compatibilityLevel").orElse("FULL");
        var productSchema = new Schema("productSchema", SchemaArgs.builder()
            .clusterId(testCluster.id())
            .subject(String.format("%s-product-value", topicName))
            .schemaType(schemaType)
            .schema(productSchemaDefinition)
            .compatibility(compatibilityLevel)
            .username(testUser.name())
            .password(userPw)
            .allowDeletion(true)
            .build(), CustomResourceOptions.builder()
                .dependsOn(
                    schemaRegistryAdmin,
                    allTestTopic,
                    describeRegistry,
                    alterConfigsRegistry)
                .build());

        var clusterAdmin = new Acl("clusterAdmin", AclArgs.builder()
            .resourceType("CLUSTER")
            .resourceName("kafka-cluster")
            .resourcePatternType("LITERAL")
            .principal(testUser.name().applyValue(name -> String.format("User:%s", name)))
            .host("*")
            .operation("ALL")
            .permissionType("ALLOW")
            .clusterApiUrl(testCluster.clusterApiUrl())
            .allowDeletion(true)
            .build());

        var readProduct = new SchemaRegistryAcl("readProduct", SchemaRegistryAclArgs.builder()
            .clusterId(testCluster.id())
            .principal(testUser.name().applyValue(name -> String.format("User:%s", name)))
            .resourceType("SUBJECT")
            .resourceName("product-")
            .patternType("PREFIXED")
            .host("*")
            .operation("READ")
            .permission("ALLOW")
            .username(testUser.name())
            .password(userPw)
            .allowDeletion(true)
            .build(), CustomResourceOptions.builder()
                .dependsOn(schemaRegistryAdmin)
                .build());

        var writeOrders = new SchemaRegistryAcl("writeOrders", SchemaRegistryAclArgs.builder()
            .clusterId(testCluster.id())
            .principal(testUser.name().applyValue(name -> String.format("User:%s", name)))
            .resourceType("SUBJECT")
            .resourceName("orders-value")
            .patternType("LITERAL")
            .host("*")
            .operation("WRITE")
            .permission("ALLOW")
            .username(testUser.name())
            .password(userPw)
            .allowDeletion(true)
            .build(), CustomResourceOptions.builder()
                .dependsOn(schemaRegistryAdmin)
                .build());

        var describeTestTopic = new SchemaRegistryAcl("describeTestTopic", SchemaRegistryAclArgs.builder()
            .clusterId(testCluster.id())
            .principal(testUser.name().applyValue(name -> String.format("User:%s", name)))
            .resourceType("SUBJECT")
            .resourceName(String.format("%s-", topicName))
            .patternType("PREFIXED")
            .host("*")
            .operation("DESCRIBE")
            .permission("ALLOW")
            .username(testUser.name())
            .password(userPw)
            .allowDeletion(true)
            .build(), CustomResourceOptions.builder()
                .dependsOn(schemaRegistryAdmin)
                .build());

        ctx.export("userSchemaInfo", Map.ofEntries(
            Map.entry("id", userSchema.schemaId()),
            Map.entry("subject", userSchema.subject()),
            Map.entry("version", userSchema.version()),
            Map.entry("type", userSchema.schemaType())
        ));
        ctx.export("userEventSchemaInfo", Map.ofEntries(
            Map.entry("id", userEventSchema.schemaId()),
            Map.entry("subject", userEventSchema.subject()),
            Map.entry("version", userEventSchema.version()),
            Map.entry("type", userEventSchema.schemaType()),
            Map.entry("references", userEventSchema.references())
        ));
        ctx.export("productSchemaInfo", Map.ofEntries(
            Map.entry("id", productSchema.schemaId()),
            Map.entry("subject", productSchema.subject()),
            Map.entry("version", productSchema.version()),
            Map.entry("type", productSchema.schemaType()),
            Map.entry("compatibility", productSchema.compatibility())
        ));
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