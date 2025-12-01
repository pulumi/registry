---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/redpanda-data/redpanda/1.4.0/index.md
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

const test = new redpanda.ResourceGroup("test", {name: resourceGroupName});
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
    allowDeletion: clusterAllowDeletion,
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
const testUser = new redpanda.User("test", {
    name: userName,
    password: userPw,
    mechanism: mechanism,
    clusterApiUrl: testCluster.clusterApiUrl,
    allowDeletion: userAllowDeletion,
});
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
    allowDeletion: aclAllowDeletion,
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
    allowDeletion: aclAllowDeletion,
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
    allowDeletion: aclAllowDeletion,
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
    allowDeletion: aclAllowDeletion,
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
const readRegistry = new redpanda.SchemaRegistryAcl("read_registry", {
    clusterId: testCluster.id,
    principal: pulumi.interpolate`User:${testUser.name}`,
    resourceType: "REGISTRY",
    resourceName: "*",
    patternType: "LITERAL",
    host: "*",
    operation: "READ",
    permission: "ALLOW",
    username: testUser.name,
    password: userPw,
    allowDeletion: true,
}, {
    dependsOn: [schemaRegistryAdmin],
});
const writeRegistry = new redpanda.SchemaRegistryAcl("write_registry", {
    clusterId: testCluster.id,
    principal: pulumi.interpolate`User:${testUser.name}`,
    resourceType: "REGISTRY",
    resourceName: "*",
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
        readRegistry,
        writeRegistry,
    ],
});
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
        readRegistry,
        writeRegistry,
    ],
});
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
        readRegistry,
        writeRegistry,
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
    allowDeletion: srAclAllowDeletion,
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
const developer = new redpanda.Role("developer", {
    name: roleName,
    clusterApiUrl: testCluster.clusterApiUrl,
    allowDeletion: roleAllowDeletion,
});
const developerAssignment = new redpanda.RoleAssignment("developer_assignment", {
    roleName: developer.name,
    principal: testUser.name,
    clusterApiUrl: testCluster.clusterApiUrl,
}, {
    dependsOn: [testUser],
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

test = redpanda.ResourceGroup("test", name=resource_group_name)
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
    allow_deletion=cluster_allow_deletion,
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
test_user = redpanda.User("test",
    name=user_name,
    password=user_pw,
    mechanism=mechanism,
    cluster_api_url=test_cluster.cluster_api_url,
    allow_deletion=user_allow_deletion)
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
    allow_deletion=acl_allow_deletion)
schema_registry_admin = redpanda.Acl("schema_registry_admin",
    resource_type="CLUSTER",
    resource_name_="kafka-cluster",
    resource_pattern_type="LITERAL",
    principal=test_user.name.apply(lambda name: f"User:{name}"),
    host="*",
    operation="ALTER",
    permission_type="ALLOW",
    cluster_api_url=test_cluster.cluster_api_url,
    allow_deletion=acl_allow_deletion)
cluster_action = redpanda.Acl("cluster_action",
    resource_type="CLUSTER",
    resource_name_="kafka-cluster",
    resource_pattern_type="LITERAL",
    principal=test_user.name.apply(lambda name: f"User:{name}"),
    host="*",
    operation="CLUSTER_ACTION",
    permission_type="ALLOW",
    cluster_api_url=test_cluster.cluster_api_url,
    allow_deletion=acl_allow_deletion)
topic_access = redpanda.Acl("topic_access",
    resource_type="TOPIC",
    resource_name_=topic_name,
    resource_pattern_type="LITERAL",
    principal=test_user.name.apply(lambda name: f"User:{name}"),
    host="*",
    operation="ALL",
    permission_type="ALLOW",
    cluster_api_url=test_cluster.cluster_api_url,
    allow_deletion=acl_allow_deletion)
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
read_registry = redpanda.SchemaRegistryAcl("read_registry",
    cluster_id=test_cluster.id,
    principal=test_user.name.apply(lambda name: f"User:{name}"),
    resource_type="REGISTRY",
    resource_name_="*",
    pattern_type="LITERAL",
    host="*",
    operation="READ",
    permission="ALLOW",
    username=test_user.name,
    password=user_pw,
    allow_deletion=True,
    opts = pulumi.ResourceOptions(depends_on=[schema_registry_admin]))
write_registry = redpanda.SchemaRegistryAcl("write_registry",
    cluster_id=test_cluster.id,
    principal=test_user.name.apply(lambda name: f"User:{name}"),
    resource_type="REGISTRY",
    resource_name_="*",
    pattern_type="LITERAL",
    host="*",
    operation="WRITE",
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
            cluster_admin,
            schema_registry_admin,
            cluster_action,
            topic_access,
            all_test_topic,
            describe_registry,
            alter_configs_registry,
            read_registry,
            write_registry,
        ]))
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
            read_registry,
            write_registry,
        ]))
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
            read_registry,
            write_registry,
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
    allow_deletion=sr_acl_allow_deletion,
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
developer = redpanda.Role("developer",
    name=role_name,
    cluster_api_url=test_cluster.cluster_api_url,
    allow_deletion=role_allow_deletion)
developer_assignment = redpanda.RoleAssignment("developer_assignment",
    role_name=developer.name,
    principal=test_user.name,
    cluster_api_url=test_cluster.cluster_api_url,
    opts = pulumi.ResourceOptions(depends_on=[test_user]))
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
    var test = new Redpanda.ResourceGroup("test", new()
    {
        Name = resourceGroupName,
    });

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
        AllowDeletion = clusterAllowDeletion,
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

    var testUser = new Redpanda.User("test", new()
    {
        Name = userName,
        Password = userPw,
        Mechanism = mechanism,
        ClusterApiUrl = testCluster.ClusterApiUrl,
        AllowDeletion = userAllowDeletion,
    });

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
        AllowDeletion = aclAllowDeletion,
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
        AllowDeletion = aclAllowDeletion,
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
        AllowDeletion = aclAllowDeletion,
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
        AllowDeletion = aclAllowDeletion,
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

    var readRegistry = new Redpanda.SchemaRegistryAcl("read_registry", new()
    {
        ClusterId = testCluster.Id,
        Principal = testUser.Name.Apply(name => $"User:{name}"),
        ResourceType = "REGISTRY",
        ResourceName = "*",
        PatternType = "LITERAL",
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

    var writeRegistry = new Redpanda.SchemaRegistryAcl("write_registry", new()
    {
        ClusterId = testCluster.Id,
        Principal = testUser.Name.Apply(name => $"User:{name}"),
        ResourceType = "REGISTRY",
        ResourceName = "*",
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
            readRegistry,
            writeRegistry,
        },
    });

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
            readRegistry,
            writeRegistry,
        },
    });

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
            readRegistry,
            writeRegistry,
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
        AllowDeletion = srAclAllowDeletion,
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

    var developer = new Redpanda.Role("developer", new()
    {
        Name = roleName,
        ClusterApiUrl = testCluster.ClusterApiUrl,
        AllowDeletion = roleAllowDeletion,
    });

    var developerAssignment = new Redpanda.RoleAssignment("developer_assignment", new()
    {
        RoleName = developer.Name,
        Principal = testUser.Name,
        ClusterApiUrl = testCluster.ClusterApiUrl,
    }, new CustomResourceOptions
    {
        DependsOn =
        {
            testUser,
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
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		test, err := redpanda.NewResourceGroup(ctx, "test", &redpanda.ResourceGroupArgs{
			Name: pulumi.Any(resourceGroupName),
		})
		if err != nil {
			return err
		}
		testNetwork, err := redpanda.NewNetwork(ctx, "test", &redpanda.NetworkArgs{
			Name:            pulumi.Any(networkName),
			ResourceGroupId: test.ID(),
			CloudProvider:   pulumi.Any(cloudProvider),
			Region:          pulumi.Any(region),
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
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"schema_registry_enable_authorization": true,
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		testCluster, err := redpanda.NewCluster(ctx, "test", &redpanda.ClusterArgs{
			Name:            pulumi.Any(clusterName),
			ResourceGroupId: test.ID(),
			NetworkId:       testNetwork.ID(),
			CloudProvider:   pulumi.Any(cloudProvider),
			Region:          pulumi.Any(region),
			ClusterType:     pulumi.String("dedicated"),
			ConnectionType:  pulumi.String("public"),
			ThroughputTier:  pulumi.Any(throughputTier),
			Zones:           pulumi.Any(zones),
			AllowDeletion:   pulumi.Any(clusterAllowDeletion),
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
		testUser, err := redpanda.NewUser(ctx, "test", &redpanda.UserArgs{
			Name:          pulumi.Any(userName),
			Password:      pulumi.Any(userPw),
			Mechanism:     pulumi.Any(mechanism),
			ClusterApiUrl: testCluster.ClusterApiUrl,
			AllowDeletion: pulumi.Any(userAllowDeletion),
		})
		if err != nil {
			return err
		}
		_, err = redpanda.NewTopic(ctx, "test", &redpanda.TopicArgs{
			Name:              pulumi.Any(topicName),
			PartitionCount:    pulumi.Any(partitionCount),
			ReplicationFactor: pulumi.Any(replicationFactor),
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
			AllowDeletion:  pulumi.Any(aclAllowDeletion),
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
			AllowDeletion:  pulumi.Any(aclAllowDeletion),
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
			AllowDeletion:  pulumi.Any(aclAllowDeletion),
		})
		if err != nil {
			return err
		}
		topicAccess, err := redpanda.NewAcl(ctx, "topic_access", &redpanda.AclArgs{
			ResourceType:        pulumi.String("TOPIC"),
			ResourceName:        pulumi.Any(topicName),
			ResourcePatternType: pulumi.String("LITERAL"),
			Principal: testUser.Name.ApplyT(func(name string) (string, error) {
				return fmt.Sprintf("User:%v", name), nil
			}).(pulumi.StringOutput),
			Host:           pulumi.String("*"),
			Operation:      pulumi.String("ALL"),
			PermissionType: pulumi.String("ALLOW"),
			ClusterApiUrl:  testCluster.ClusterApiUrl,
			AllowDeletion:  pulumi.Any(aclAllowDeletion),
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
			Password:      pulumi.Any(userPw),
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
			Password:      pulumi.Any(userPw),
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
			Password:      pulumi.Any(userPw),
			AllowDeletion: pulumi.Bool(true),
		}, pulumi.DependsOn([]pulumi.Resource{
			schemaRegistryAdmin,
		}))
		if err != nil {
			return err
		}
		readRegistry, err := redpanda.NewSchemaRegistryAcl(ctx, "read_registry", &redpanda.SchemaRegistryAclArgs{
			ClusterId: testCluster.ID(),
			Principal: testUser.Name.ApplyT(func(name string) (string, error) {
				return fmt.Sprintf("User:%v", name), nil
			}).(pulumi.StringOutput),
			ResourceType:  pulumi.String("REGISTRY"),
			ResourceName:  pulumi.String("*"),
			PatternType:   pulumi.String("LITERAL"),
			Host:          pulumi.String("*"),
			Operation:     pulumi.String("READ"),
			Permission:    pulumi.String("ALLOW"),
			Username:      testUser.Name,
			Password:      pulumi.Any(userPw),
			AllowDeletion: pulumi.Bool(true),
		}, pulumi.DependsOn([]pulumi.Resource{
			schemaRegistryAdmin,
		}))
		if err != nil {
			return err
		}
		writeRegistry, err := redpanda.NewSchemaRegistryAcl(ctx, "write_registry", &redpanda.SchemaRegistryAclArgs{
			ClusterId: testCluster.ID(),
			Principal: testUser.Name.ApplyT(func(name string) (string, error) {
				return fmt.Sprintf("User:%v", name), nil
			}).(pulumi.StringOutput),
			ResourceType:  pulumi.String("REGISTRY"),
			ResourceName:  pulumi.String("*"),
			PatternType:   pulumi.String("LITERAL"),
			Host:          pulumi.String("*"),
			Operation:     pulumi.String("WRITE"),
			Permission:    pulumi.String("ALLOW"),
			Username:      testUser.Name,
			Password:      pulumi.Any(userPw),
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
			SchemaType:    pulumi.Any(schemaType),
			Schema:        pulumi.Any(userSchemaDefinition),
			Username:      testUser.Name,
			Password:      pulumi.Any(userPw),
			AllowDeletion: pulumi.Bool(true),
		}, pulumi.DependsOn([]pulumi.Resource{
			clusterAdmin,
			schemaRegistryAdmin,
			clusterAction,
			topicAccess,
			allTestTopic,
			describeRegistry,
			alterConfigsRegistry,
			readRegistry,
			writeRegistry,
		}))
		if err != nil {
			return err
		}
		userEventSchema, err := redpanda.NewSchema(ctx, "user_event_schema", &redpanda.SchemaArgs{
			ClusterId:     testCluster.ID(),
			Subject:       pulumi.Sprintf("%v-events-value", topicName),
			SchemaType:    pulumi.Any(schemaType),
			Schema:        pulumi.Any(userEventSchemaDefinition),
			Username:      testUser.Name,
			Password:      pulumi.Any(userPw),
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
			readRegistry,
			writeRegistry,
		}))
		if err != nil {
			return err
		}
		productSchema, err := redpanda.NewSchema(ctx, "product_schema", &redpanda.SchemaArgs{
			ClusterId:     testCluster.ID(),
			Subject:       pulumi.Sprintf("%v-product-value", topicName),
			SchemaType:    pulumi.Any(schemaType),
			Schema:        pulumi.Any(productSchemaDefinition),
			Compatibility: pulumi.Any(compatibilityLevel),
			Username:      testUser.Name,
			Password:      pulumi.Any(userPw),
			AllowDeletion: pulumi.Bool(true),
		}, pulumi.DependsOn([]pulumi.Resource{
			clusterAdmin,
			schemaRegistryAdmin,
			clusterAction,
			topicAccess,
			allTestTopic,
			describeRegistry,
			alterConfigsRegistry,
			readRegistry,
			writeRegistry,
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
			Password:      pulumi.Any(userPw),
			AllowDeletion: pulumi.Any(srAclAllowDeletion),
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
			Password:      pulumi.Any(userPw),
			AllowDeletion: pulumi.Bool(true),
		}, pulumi.DependsOn([]pulumi.Resource{
			schemaRegistryAdmin,
		}))
		if err != nil {
			return err
		}
		developer, err := redpanda.NewRole(ctx, "developer", &redpanda.RoleArgs{
			Name:          pulumi.Any(roleName),
			ClusterApiUrl: testCluster.ClusterApiUrl,
			AllowDeletion: pulumi.Any(roleAllowDeletion),
		})
		if err != nil {
			return err
		}
		_, err = redpanda.NewRoleAssignment(ctx, "developer_assignment", &redpanda.RoleAssignmentArgs{
			RoleName:      developer.Name,
			Principal:     testUser.Name,
			ClusterApiUrl: testCluster.ClusterApiUrl,
		}, pulumi.DependsOn([]pulumi.Resource{
			testUser,
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
      allowDeletion: ${clusterAllowDeletion}
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
      allowDeletion: ${userAllowDeletion}
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
        - ${readRegistry}
        - ${writeRegistry}
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
        - ${readRegistry}
        - ${writeRegistry}
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
        - ${readRegistry}
        - ${writeRegistry}
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
      allowDeletion: ${aclAllowDeletion}
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
      allowDeletion: ${aclAllowDeletion}
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
      allowDeletion: ${aclAllowDeletion}
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
      allowDeletion: ${aclAllowDeletion}
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
      allowDeletion: ${srAclAllowDeletion}
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
  readRegistry:
    type: redpanda:SchemaRegistryAcl
    name: read_registry
    properties:
      clusterId: ${testCluster.id}
      principal: User:${testUser.name}
      resourceType: REGISTRY
      resourceName: '*'
      patternType: LITERAL
      host: '*'
      operation: READ
      permission: ALLOW
      username: ${testUser.name}
      password: ${userPw}
      allowDeletion: true
    options:
      dependsOn:
        - ${schemaRegistryAdmin}
  writeRegistry:
    type: redpanda:SchemaRegistryAcl
    name: write_registry
    properties:
      clusterId: ${testCluster.id}
      principal: User:${testUser.name}
      resourceType: REGISTRY
      resourceName: '*'
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
  developer:
    type: redpanda:Role
    properties:
      name: ${roleName}
      clusterApiUrl: ${testCluster.clusterApiUrl}
      allowDeletion: ${roleAllowDeletion}
  developerAssignment:
    type: redpanda:RoleAssignment
    name: developer_assignment
    properties:
      roleName: ${developer.name}
      principal: ${testUser.name}
      clusterApiUrl: ${testCluster.clusterApiUrl}
    options:
      dependsOn:
        - ${testUser}
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
import com.pulumi.redpanda.Role;
import com.pulumi.redpanda.RoleArgs;
import com.pulumi.redpanda.RoleAssignment;
import com.pulumi.redpanda.RoleAssignmentArgs;
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
        var test = new ResourceGroup("test", ResourceGroupArgs.builder()
            .name(resourceGroupName)
            .build());

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
            .allowDeletion(clusterAllowDeletion)
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

        var testUser = new User("testUser", UserArgs.builder()
            .name(userName)
            .password(userPw)
            .mechanism(mechanism)
            .clusterApiUrl(testCluster.clusterApiUrl())
            .allowDeletion(userAllowDeletion)
            .build());

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
            .principal(testUser.name().applyValue(_name -> String.format("User:%s", _name)))
            .host("*")
            .operation("ALL")
            .permissionType("ALLOW")
            .clusterApiUrl(testCluster.clusterApiUrl())
            .allowDeletion(aclAllowDeletion)
            .build());

        var schemaRegistryAdmin = new Acl("schemaRegistryAdmin", AclArgs.builder()
            .resourceType("CLUSTER")
            .resourceName("kafka-cluster")
            .resourcePatternType("LITERAL")
            .principal(testUser.name().applyValue(_name -> String.format("User:%s", _name)))
            .host("*")
            .operation("ALTER")
            .permissionType("ALLOW")
            .clusterApiUrl(testCluster.clusterApiUrl())
            .allowDeletion(aclAllowDeletion)
            .build());

        var clusterAction = new Acl("clusterAction", AclArgs.builder()
            .resourceType("CLUSTER")
            .resourceName("kafka-cluster")
            .resourcePatternType("LITERAL")
            .principal(testUser.name().applyValue(_name -> String.format("User:%s", _name)))
            .host("*")
            .operation("CLUSTER_ACTION")
            .permissionType("ALLOW")
            .clusterApiUrl(testCluster.clusterApiUrl())
            .allowDeletion(aclAllowDeletion)
            .build());

        var topicAccess = new Acl("topicAccess", AclArgs.builder()
            .resourceType("TOPIC")
            .resourceName(topicName)
            .resourcePatternType("LITERAL")
            .principal(testUser.name().applyValue(_name -> String.format("User:%s", _name)))
            .host("*")
            .operation("ALL")
            .permissionType("ALLOW")
            .clusterApiUrl(testCluster.clusterApiUrl())
            .allowDeletion(aclAllowDeletion)
            .build());

        var allTestTopic = new SchemaRegistryAcl("allTestTopic", SchemaRegistryAclArgs.builder()
            .clusterId(testCluster.id())
            .principal(testUser.name().applyValue(_name -> String.format("User:%s", _name)))
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
            .principal(testUser.name().applyValue(_name -> String.format("User:%s", _name)))
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
            .principal(testUser.name().applyValue(_name -> String.format("User:%s", _name)))
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

        var readRegistry = new SchemaRegistryAcl("readRegistry", SchemaRegistryAclArgs.builder()
            .clusterId(testCluster.id())
            .principal(testUser.name().applyValue(_name -> String.format("User:%s", _name)))
            .resourceType("REGISTRY")
            .resourceName("*")
            .patternType("LITERAL")
            .host("*")
            .operation("READ")
            .permission("ALLOW")
            .username(testUser.name())
            .password(userPw)
            .allowDeletion(true)
            .build(), CustomResourceOptions.builder()
                .dependsOn(schemaRegistryAdmin)
                .build());

        var writeRegistry = new SchemaRegistryAcl("writeRegistry", SchemaRegistryAclArgs.builder()
            .clusterId(testCluster.id())
            .principal(testUser.name().applyValue(_name -> String.format("User:%s", _name)))
            .resourceType("REGISTRY")
            .resourceName("*")
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
                    alterConfigsRegistry,
                    readRegistry,
                    writeRegistry)
                .build());

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
                    alterConfigsRegistry,
                    readRegistry,
                    writeRegistry)
                .build());

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
                    alterConfigsRegistry,
                    readRegistry,
                    writeRegistry)
                .build());

        var readProduct = new SchemaRegistryAcl("readProduct", SchemaRegistryAclArgs.builder()
            .clusterId(testCluster.id())
            .principal(testUser.name().applyValue(_name -> String.format("User:%s", _name)))
            .resourceType("SUBJECT")
            .resourceName("product-")
            .patternType("PREFIXED")
            .host("*")
            .operation("READ")
            .permission("ALLOW")
            .username(testUser.name())
            .password(userPw)
            .allowDeletion(srAclAllowDeletion)
            .build(), CustomResourceOptions.builder()
                .dependsOn(schemaRegistryAdmin)
                .build());

        var writeOrders = new SchemaRegistryAcl("writeOrders", SchemaRegistryAclArgs.builder()
            .clusterId(testCluster.id())
            .principal(testUser.name().applyValue(_name -> String.format("User:%s", _name)))
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

        var developer = new Role("developer", RoleArgs.builder()
            .name(roleName)
            .clusterApiUrl(testCluster.clusterApiUrl())
            .allowDeletion(roleAllowDeletion)
            .build());

        var developerAssignment = new RoleAssignment("developerAssignment", RoleAssignmentArgs.builder()
            .roleName(developer.name())
            .principal(testUser.name())
            .clusterApiUrl(testCluster.clusterApiUrl())
            .build(), CustomResourceOptions.builder()
                .dependsOn(testUser)
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

const test = new redpanda.ResourceGroup("test", {name: resourceGroupName});
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
    allowDeletion: clusterAllowDeletion,
    clusterConfiguration: {
        customPropertiesJson: JSON.stringify({
            schema_registry_enable_authorization: true,
        }),
    },
    timeouts: {
        create: "90m",
    },
});
const testUser = new redpanda.User("test", {
    name: userName,
    password: userPw,
    mechanism: mechanism,
    clusterApiUrl: testCluster.clusterApiUrl,
    allowDeletion: userAllowDeletion,
});
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
    allowDeletion: aclAllowDeletion,
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
    allowDeletion: aclAllowDeletion,
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
    allowDeletion: aclAllowDeletion,
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
    allowDeletion: aclAllowDeletion,
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
const readRegistry = new redpanda.SchemaRegistryAcl("read_registry", {
    clusterId: testCluster.id,
    principal: pulumi.interpolate`User:${testUser.name}`,
    resourceType: "REGISTRY",
    resourceName: "*",
    patternType: "LITERAL",
    host: "*",
    operation: "READ",
    permission: "ALLOW",
    username: testUser.name,
    password: userPw,
    allowDeletion: true,
}, {
    dependsOn: [schemaRegistryAdmin],
});
const writeRegistry = new redpanda.SchemaRegistryAcl("write_registry", {
    clusterId: testCluster.id,
    principal: pulumi.interpolate`User:${testUser.name}`,
    resourceType: "REGISTRY",
    resourceName: "*",
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
        readRegistry,
        writeRegistry,
    ],
});
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
        readRegistry,
        writeRegistry,
    ],
});
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
        readRegistry,
        writeRegistry,
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
    allowDeletion: srAclAllowDeletion,
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
const developer = new redpanda.Role("developer", {
    name: roleName,
    clusterApiUrl: testCluster.clusterApiUrl,
    allowDeletion: roleAllowDeletion,
});
const developerAssignment = new redpanda.RoleAssignment("developer_assignment", {
    roleName: developer.name,
    principal: testUser.name,
    clusterApiUrl: testCluster.clusterApiUrl,
}, {
    dependsOn: [testUser],
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

test = redpanda.ResourceGroup("test", name=resource_group_name)
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
    allow_deletion=cluster_allow_deletion,
    cluster_configuration={
        "custom_properties_json": json.dumps({
            "schema_registry_enable_authorization": True,
        }),
    },
    timeouts={
        "create": "90m",
    })
test_user = redpanda.User("test",
    name=user_name,
    password=user_pw,
    mechanism=mechanism,
    cluster_api_url=test_cluster.cluster_api_url,
    allow_deletion=user_allow_deletion)
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
    allow_deletion=acl_allow_deletion)
schema_registry_admin = redpanda.Acl("schema_registry_admin",
    resource_type="CLUSTER",
    resource_name_="kafka-cluster",
    resource_pattern_type="LITERAL",
    principal=test_user.name.apply(lambda name: f"User:{name}"),
    host="*",
    operation="ALTER",
    permission_type="ALLOW",
    cluster_api_url=test_cluster.cluster_api_url,
    allow_deletion=acl_allow_deletion)
cluster_action = redpanda.Acl("cluster_action",
    resource_type="CLUSTER",
    resource_name_="kafka-cluster",
    resource_pattern_type="LITERAL",
    principal=test_user.name.apply(lambda name: f"User:{name}"),
    host="*",
    operation="CLUSTER_ACTION",
    permission_type="ALLOW",
    cluster_api_url=test_cluster.cluster_api_url,
    allow_deletion=acl_allow_deletion)
topic_access = redpanda.Acl("topic_access",
    resource_type="TOPIC",
    resource_name_=topic_name,
    resource_pattern_type="LITERAL",
    principal=test_user.name.apply(lambda name: f"User:{name}"),
    host="*",
    operation="ALL",
    permission_type="ALLOW",
    cluster_api_url=test_cluster.cluster_api_url,
    allow_deletion=acl_allow_deletion)
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
read_registry = redpanda.SchemaRegistryAcl("read_registry",
    cluster_id=test_cluster.id,
    principal=test_user.name.apply(lambda name: f"User:{name}"),
    resource_type="REGISTRY",
    resource_name_="*",
    pattern_type="LITERAL",
    host="*",
    operation="READ",
    permission="ALLOW",
    username=test_user.name,
    password=user_pw,
    allow_deletion=True,
    opts = pulumi.ResourceOptions(depends_on=[schema_registry_admin]))
write_registry = redpanda.SchemaRegistryAcl("write_registry",
    cluster_id=test_cluster.id,
    principal=test_user.name.apply(lambda name: f"User:{name}"),
    resource_type="REGISTRY",
    resource_name_="*",
    pattern_type="LITERAL",
    host="*",
    operation="WRITE",
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
            cluster_admin,
            schema_registry_admin,
            cluster_action,
            topic_access,
            all_test_topic,
            describe_registry,
            alter_configs_registry,
            read_registry,
            write_registry,
        ]))
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
            read_registry,
            write_registry,
        ]))
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
            read_registry,
            write_registry,
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
    allow_deletion=sr_acl_allow_deletion,
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
developer = redpanda.Role("developer",
    name=role_name,
    cluster_api_url=test_cluster.cluster_api_url,
    allow_deletion=role_allow_deletion)
developer_assignment = redpanda.RoleAssignment("developer_assignment",
    role_name=developer.name,
    principal=test_user.name,
    cluster_api_url=test_cluster.cluster_api_url,
    opts = pulumi.ResourceOptions(depends_on=[test_user]))
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
    var test = new Redpanda.ResourceGroup("test", new()
    {
        Name = resourceGroupName,
    });

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
        AllowDeletion = clusterAllowDeletion,
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

    var testUser = new Redpanda.User("test", new()
    {
        Name = userName,
        Password = userPw,
        Mechanism = mechanism,
        ClusterApiUrl = testCluster.ClusterApiUrl,
        AllowDeletion = userAllowDeletion,
    });

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
        AllowDeletion = aclAllowDeletion,
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
        AllowDeletion = aclAllowDeletion,
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
        AllowDeletion = aclAllowDeletion,
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
        AllowDeletion = aclAllowDeletion,
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

    var readRegistry = new Redpanda.SchemaRegistryAcl("read_registry", new()
    {
        ClusterId = testCluster.Id,
        Principal = testUser.Name.Apply(name => $"User:{name}"),
        ResourceType = "REGISTRY",
        ResourceName = "*",
        PatternType = "LITERAL",
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

    var writeRegistry = new Redpanda.SchemaRegistryAcl("write_registry", new()
    {
        ClusterId = testCluster.Id,
        Principal = testUser.Name.Apply(name => $"User:{name}"),
        ResourceType = "REGISTRY",
        ResourceName = "*",
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
            readRegistry,
            writeRegistry,
        },
    });

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
            readRegistry,
            writeRegistry,
        },
    });

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
            readRegistry,
            writeRegistry,
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
        AllowDeletion = srAclAllowDeletion,
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

    var developer = new Redpanda.Role("developer", new()
    {
        Name = roleName,
        ClusterApiUrl = testCluster.ClusterApiUrl,
        AllowDeletion = roleAllowDeletion,
    });

    var developerAssignment = new Redpanda.RoleAssignment("developer_assignment", new()
    {
        RoleName = developer.Name,
        Principal = testUser.Name,
        ClusterApiUrl = testCluster.ClusterApiUrl,
    }, new CustomResourceOptions
    {
        DependsOn =
        {
            testUser,
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
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		test, err := redpanda.NewResourceGroup(ctx, "test", &redpanda.ResourceGroupArgs{
			Name: pulumi.Any(resourceGroupName),
		})
		if err != nil {
			return err
		}
		testNetwork, err := redpanda.NewNetwork(ctx, "test", &redpanda.NetworkArgs{
			Name:            pulumi.Any(networkName),
			ResourceGroupId: test.ID(),
			CloudProvider:   pulumi.Any(cloudProvider),
			Region:          pulumi.Any(region),
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
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"schema_registry_enable_authorization": true,
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		testCluster, err := redpanda.NewCluster(ctx, "test", &redpanda.ClusterArgs{
			Name:            pulumi.Any(clusterName),
			ResourceGroupId: test.ID(),
			NetworkId:       testNetwork.ID(),
			CloudProvider:   pulumi.Any(cloudProvider),
			Region:          pulumi.Any(region),
			ClusterType:     pulumi.String("dedicated"),
			ConnectionType:  pulumi.String("public"),
			ThroughputTier:  pulumi.Any(throughputTier),
			Zones:           pulumi.Any(zones),
			AllowDeletion:   pulumi.Any(clusterAllowDeletion),
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
		testUser, err := redpanda.NewUser(ctx, "test", &redpanda.UserArgs{
			Name:          pulumi.Any(userName),
			Password:      pulumi.Any(userPw),
			Mechanism:     pulumi.Any(mechanism),
			ClusterApiUrl: testCluster.ClusterApiUrl,
			AllowDeletion: pulumi.Any(userAllowDeletion),
		})
		if err != nil {
			return err
		}
		_, err = redpanda.NewTopic(ctx, "test", &redpanda.TopicArgs{
			Name:              pulumi.Any(topicName),
			PartitionCount:    pulumi.Any(partitionCount),
			ReplicationFactor: pulumi.Any(replicationFactor),
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
			AllowDeletion:  pulumi.Any(aclAllowDeletion),
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
			AllowDeletion:  pulumi.Any(aclAllowDeletion),
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
			AllowDeletion:  pulumi.Any(aclAllowDeletion),
		})
		if err != nil {
			return err
		}
		topicAccess, err := redpanda.NewAcl(ctx, "topic_access", &redpanda.AclArgs{
			ResourceType:        pulumi.String("TOPIC"),
			ResourceName:        pulumi.Any(topicName),
			ResourcePatternType: pulumi.String("LITERAL"),
			Principal: testUser.Name.ApplyT(func(name string) (string, error) {
				return fmt.Sprintf("User:%v", name), nil
			}).(pulumi.StringOutput),
			Host:           pulumi.String("*"),
			Operation:      pulumi.String("ALL"),
			PermissionType: pulumi.String("ALLOW"),
			ClusterApiUrl:  testCluster.ClusterApiUrl,
			AllowDeletion:  pulumi.Any(aclAllowDeletion),
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
			Password:      pulumi.Any(userPw),
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
			Password:      pulumi.Any(userPw),
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
			Password:      pulumi.Any(userPw),
			AllowDeletion: pulumi.Bool(true),
		}, pulumi.DependsOn([]pulumi.Resource{
			schemaRegistryAdmin,
		}))
		if err != nil {
			return err
		}
		readRegistry, err := redpanda.NewSchemaRegistryAcl(ctx, "read_registry", &redpanda.SchemaRegistryAclArgs{
			ClusterId: testCluster.ID(),
			Principal: testUser.Name.ApplyT(func(name string) (string, error) {
				return fmt.Sprintf("User:%v", name), nil
			}).(pulumi.StringOutput),
			ResourceType:  pulumi.String("REGISTRY"),
			ResourceName:  pulumi.String("*"),
			PatternType:   pulumi.String("LITERAL"),
			Host:          pulumi.String("*"),
			Operation:     pulumi.String("READ"),
			Permission:    pulumi.String("ALLOW"),
			Username:      testUser.Name,
			Password:      pulumi.Any(userPw),
			AllowDeletion: pulumi.Bool(true),
		}, pulumi.DependsOn([]pulumi.Resource{
			schemaRegistryAdmin,
		}))
		if err != nil {
			return err
		}
		writeRegistry, err := redpanda.NewSchemaRegistryAcl(ctx, "write_registry", &redpanda.SchemaRegistryAclArgs{
			ClusterId: testCluster.ID(),
			Principal: testUser.Name.ApplyT(func(name string) (string, error) {
				return fmt.Sprintf("User:%v", name), nil
			}).(pulumi.StringOutput),
			ResourceType:  pulumi.String("REGISTRY"),
			ResourceName:  pulumi.String("*"),
			PatternType:   pulumi.String("LITERAL"),
			Host:          pulumi.String("*"),
			Operation:     pulumi.String("WRITE"),
			Permission:    pulumi.String("ALLOW"),
			Username:      testUser.Name,
			Password:      pulumi.Any(userPw),
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
			SchemaType:    pulumi.Any(schemaType),
			Schema:        pulumi.Any(userSchemaDefinition),
			Username:      testUser.Name,
			Password:      pulumi.Any(userPw),
			AllowDeletion: pulumi.Bool(true),
		}, pulumi.DependsOn([]pulumi.Resource{
			clusterAdmin,
			schemaRegistryAdmin,
			clusterAction,
			topicAccess,
			allTestTopic,
			describeRegistry,
			alterConfigsRegistry,
			readRegistry,
			writeRegistry,
		}))
		if err != nil {
			return err
		}
		userEventSchema, err := redpanda.NewSchema(ctx, "user_event_schema", &redpanda.SchemaArgs{
			ClusterId:     testCluster.ID(),
			Subject:       pulumi.Sprintf("%v-events-value", topicName),
			SchemaType:    pulumi.Any(schemaType),
			Schema:        pulumi.Any(userEventSchemaDefinition),
			Username:      testUser.Name,
			Password:      pulumi.Any(userPw),
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
			readRegistry,
			writeRegistry,
		}))
		if err != nil {
			return err
		}
		productSchema, err := redpanda.NewSchema(ctx, "product_schema", &redpanda.SchemaArgs{
			ClusterId:     testCluster.ID(),
			Subject:       pulumi.Sprintf("%v-product-value", topicName),
			SchemaType:    pulumi.Any(schemaType),
			Schema:        pulumi.Any(productSchemaDefinition),
			Compatibility: pulumi.Any(compatibilityLevel),
			Username:      testUser.Name,
			Password:      pulumi.Any(userPw),
			AllowDeletion: pulumi.Bool(true),
		}, pulumi.DependsOn([]pulumi.Resource{
			clusterAdmin,
			schemaRegistryAdmin,
			clusterAction,
			topicAccess,
			allTestTopic,
			describeRegistry,
			alterConfigsRegistry,
			readRegistry,
			writeRegistry,
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
			Password:      pulumi.Any(userPw),
			AllowDeletion: pulumi.Any(srAclAllowDeletion),
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
			Password:      pulumi.Any(userPw),
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
			Password:      pulumi.Any(userPw),
			AllowDeletion: pulumi.Bool(true),
		}, pulumi.DependsOn([]pulumi.Resource{
			schemaRegistryAdmin,
		}))
		if err != nil {
			return err
		}
		developer, err := redpanda.NewRole(ctx, "developer", &redpanda.RoleArgs{
			Name:          pulumi.Any(roleName),
			ClusterApiUrl: testCluster.ClusterApiUrl,
			AllowDeletion: pulumi.Any(roleAllowDeletion),
		})
		if err != nil {
			return err
		}
		_, err = redpanda.NewRoleAssignment(ctx, "developer_assignment", &redpanda.RoleAssignmentArgs{
			RoleName:      developer.Name,
			Principal:     testUser.Name,
			ClusterApiUrl: testCluster.ClusterApiUrl,
		}, pulumi.DependsOn([]pulumi.Resource{
			testUser,
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
      allowDeletion: ${clusterAllowDeletion}
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
      allowDeletion: ${userAllowDeletion}
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
        - ${readRegistry}
        - ${writeRegistry}
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
        - ${readRegistry}
        - ${writeRegistry}
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
        - ${readRegistry}
        - ${writeRegistry}
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
      allowDeletion: ${aclAllowDeletion}
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
      allowDeletion: ${aclAllowDeletion}
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
      allowDeletion: ${aclAllowDeletion}
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
      allowDeletion: ${aclAllowDeletion}
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
      allowDeletion: ${srAclAllowDeletion}
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
  readRegistry:
    type: redpanda:SchemaRegistryAcl
    name: read_registry
    properties:
      clusterId: ${testCluster.id}
      principal: User:${testUser.name}
      resourceType: REGISTRY
      resourceName: '*'
      patternType: LITERAL
      host: '*'
      operation: READ
      permission: ALLOW
      username: ${testUser.name}
      password: ${userPw}
      allowDeletion: true
    options:
      dependsOn:
        - ${schemaRegistryAdmin}
  writeRegistry:
    type: redpanda:SchemaRegistryAcl
    name: write_registry
    properties:
      clusterId: ${testCluster.id}
      principal: User:${testUser.name}
      resourceType: REGISTRY
      resourceName: '*'
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
  developer:
    type: redpanda:Role
    properties:
      name: ${roleName}
      clusterApiUrl: ${testCluster.clusterApiUrl}
      allowDeletion: ${roleAllowDeletion}
  developerAssignment:
    type: redpanda:RoleAssignment
    name: developer_assignment
    properties:
      roleName: ${developer.name}
      principal: ${testUser.name}
      clusterApiUrl: ${testCluster.clusterApiUrl}
    options:
      dependsOn:
        - ${testUser}
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
import com.pulumi.redpanda.Role;
import com.pulumi.redpanda.RoleArgs;
import com.pulumi.redpanda.RoleAssignment;
import com.pulumi.redpanda.RoleAssignmentArgs;
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
        var test = new ResourceGroup("test", ResourceGroupArgs.builder()
            .name(resourceGroupName)
            .build());

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
            .allowDeletion(clusterAllowDeletion)
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

        var testUser = new User("testUser", UserArgs.builder()
            .name(userName)
            .password(userPw)
            .mechanism(mechanism)
            .clusterApiUrl(testCluster.clusterApiUrl())
            .allowDeletion(userAllowDeletion)
            .build());

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
            .principal(testUser.name().applyValue(_name -> String.format("User:%s", _name)))
            .host("*")
            .operation("ALL")
            .permissionType("ALLOW")
            .clusterApiUrl(testCluster.clusterApiUrl())
            .allowDeletion(aclAllowDeletion)
            .build());

        var schemaRegistryAdmin = new Acl("schemaRegistryAdmin", AclArgs.builder()
            .resourceType("CLUSTER")
            .resourceName("kafka-cluster")
            .resourcePatternType("LITERAL")
            .principal(testUser.name().applyValue(_name -> String.format("User:%s", _name)))
            .host("*")
            .operation("ALTER")
            .permissionType("ALLOW")
            .clusterApiUrl(testCluster.clusterApiUrl())
            .allowDeletion(aclAllowDeletion)
            .build());

        var clusterAction = new Acl("clusterAction", AclArgs.builder()
            .resourceType("CLUSTER")
            .resourceName("kafka-cluster")
            .resourcePatternType("LITERAL")
            .principal(testUser.name().applyValue(_name -> String.format("User:%s", _name)))
            .host("*")
            .operation("CLUSTER_ACTION")
            .permissionType("ALLOW")
            .clusterApiUrl(testCluster.clusterApiUrl())
            .allowDeletion(aclAllowDeletion)
            .build());

        var topicAccess = new Acl("topicAccess", AclArgs.builder()
            .resourceType("TOPIC")
            .resourceName(topicName)
            .resourcePatternType("LITERAL")
            .principal(testUser.name().applyValue(_name -> String.format("User:%s", _name)))
            .host("*")
            .operation("ALL")
            .permissionType("ALLOW")
            .clusterApiUrl(testCluster.clusterApiUrl())
            .allowDeletion(aclAllowDeletion)
            .build());

        var allTestTopic = new SchemaRegistryAcl("allTestTopic", SchemaRegistryAclArgs.builder()
            .clusterId(testCluster.id())
            .principal(testUser.name().applyValue(_name -> String.format("User:%s", _name)))
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
            .principal(testUser.name().applyValue(_name -> String.format("User:%s", _name)))
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
            .principal(testUser.name().applyValue(_name -> String.format("User:%s", _name)))
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

        var readRegistry = new SchemaRegistryAcl("readRegistry", SchemaRegistryAclArgs.builder()
            .clusterId(testCluster.id())
            .principal(testUser.name().applyValue(_name -> String.format("User:%s", _name)))
            .resourceType("REGISTRY")
            .resourceName("*")
            .patternType("LITERAL")
            .host("*")
            .operation("READ")
            .permission("ALLOW")
            .username(testUser.name())
            .password(userPw)
            .allowDeletion(true)
            .build(), CustomResourceOptions.builder()
                .dependsOn(schemaRegistryAdmin)
                .build());

        var writeRegistry = new SchemaRegistryAcl("writeRegistry", SchemaRegistryAclArgs.builder()
            .clusterId(testCluster.id())
            .principal(testUser.name().applyValue(_name -> String.format("User:%s", _name)))
            .resourceType("REGISTRY")
            .resourceName("*")
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
                    alterConfigsRegistry,
                    readRegistry,
                    writeRegistry)
                .build());

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
                    alterConfigsRegistry,
                    readRegistry,
                    writeRegistry)
                .build());

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
                    alterConfigsRegistry,
                    readRegistry,
                    writeRegistry)
                .build());

        var readProduct = new SchemaRegistryAcl("readProduct", SchemaRegistryAclArgs.builder()
            .clusterId(testCluster.id())
            .principal(testUser.name().applyValue(_name -> String.format("User:%s", _name)))
            .resourceType("SUBJECT")
            .resourceName("product-")
            .patternType("PREFIXED")
            .host("*")
            .operation("READ")
            .permission("ALLOW")
            .username(testUser.name())
            .password(userPw)
            .allowDeletion(srAclAllowDeletion)
            .build(), CustomResourceOptions.builder()
                .dependsOn(schemaRegistryAdmin)
                .build());

        var writeOrders = new SchemaRegistryAcl("writeOrders", SchemaRegistryAclArgs.builder()
            .clusterId(testCluster.id())
            .principal(testUser.name().applyValue(_name -> String.format("User:%s", _name)))
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
            .principal(testUser.name().applyValue(_name -> String.format("User:%s", _name)))
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

        var developer = new Role("developer", RoleArgs.builder()
            .name(roleName)
            .clusterApiUrl(testCluster.clusterApiUrl())
            .allowDeletion(roleAllowDeletion)
            .build());

        var developerAssignment = new RoleAssignment("developerAssignment", RoleAssignmentArgs.builder()
            .roleName(developer.name())
            .principal(testUser.name())
            .clusterApiUrl(testCluster.clusterApiUrl())
            .build(), CustomResourceOptions.builder()
                .dependsOn(testUser)
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

const test = redpanda.getCluster({
    id: clusterId,
});
const testTopic = new redpanda.Topic("test", {
    name: topicName,
    partitionCount: partitionCount,
    replicationFactor: replicationFactor,
    clusterApiUrl: test.then(test => test.clusterApiUrl),
    allowDeletion: true,
    configuration: topicConfig,
});
const testUser = new redpanda.User("test", {
    name: userName,
    password: userPw,
    mechanism: mechanism,
    clusterApiUrl: test.then(test => test.clusterApiUrl),
    allowDeletion: userAllowDeletion,
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
    allowDeletion: aclAllowDeletion,
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

test = redpanda.get_cluster(id=cluster_id)
test_topic = redpanda.Topic("test",
    name=topic_name,
    partition_count=partition_count,
    replication_factor=replication_factor,
    cluster_api_url=test.cluster_api_url,
    allow_deletion=True,
    configuration=topic_config)
test_user = redpanda.User("test",
    name=user_name,
    password=user_pw,
    mechanism=mechanism,
    cluster_api_url=test.cluster_api_url,
    allow_deletion=user_allow_deletion)
test_acl = redpanda.Acl("test",
    resource_type="CLUSTER",
    resource_name_="kafka-cluster",
    resource_pattern_type="LITERAL",
    principal=test_user.name.apply(lambda name: f"User:{name}"),
    host="*",
    operation="ALTER",
    permission_type="ALLOW",
    cluster_api_url=test.cluster_api_url,
    allow_deletion=acl_allow_deletion)
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
    var test = Redpanda.GetCluster.Invoke(new()
    {
        Id = clusterId,
    });

    var testTopic = new Redpanda.Topic("test", new()
    {
        Name = topicName,
        PartitionCount = partitionCount,
        ReplicationFactor = replicationFactor,
        ClusterApiUrl = test.Apply(getClusterResult => getClusterResult.ClusterApiUrl),
        AllowDeletion = true,
        Configuration = topicConfig,
    });

    var testUser = new Redpanda.User("test", new()
    {
        Name = userName,
        Password = userPw,
        Mechanism = mechanism,
        ClusterApiUrl = test.Apply(getClusterResult => getClusterResult.ClusterApiUrl),
        AllowDeletion = userAllowDeletion,
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
        AllowDeletion = aclAllowDeletion,
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
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		test, err := redpanda.LookupCluster(ctx, &redpanda.LookupClusterArgs{
			Id: clusterId,
		}, nil)
		if err != nil {
			return err
		}
		_, err = redpanda.NewTopic(ctx, "test", &redpanda.TopicArgs{
			Name:              pulumi.Any(topicName),
			PartitionCount:    pulumi.Any(partitionCount),
			ReplicationFactor: pulumi.Any(replicationFactor),
			ClusterApiUrl:     pulumi.String(test.ClusterApiUrl),
			AllowDeletion:     pulumi.Bool(true),
			Configuration:     pulumi.Any(topicConfig),
		})
		if err != nil {
			return err
		}
		testUser, err := redpanda.NewUser(ctx, "test", &redpanda.UserArgs{
			Name:          pulumi.Any(userName),
			Password:      pulumi.Any(userPw),
			Mechanism:     pulumi.Any(mechanism),
			ClusterApiUrl: pulumi.String(test.ClusterApiUrl),
			AllowDeletion: pulumi.Any(userAllowDeletion),
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
			AllowDeletion:  pulumi.Any(aclAllowDeletion),
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
      allowDeletion: ${userAllowDeletion}
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
      allowDeletion: ${aclAllowDeletion}
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
        final var test = RedpandaFunctions.getCluster(GetClusterArgs.builder()
            .id(clusterId)
            .build());

        var testTopic = new Topic("testTopic", TopicArgs.builder()
            .name(topicName)
            .partitionCount(partitionCount)
            .replicationFactor(replicationFactor)
            .clusterApiUrl(test.clusterApiUrl())
            .allowDeletion(true)
            .configuration(topicConfig)
            .build());

        var testUser = new User("testUser", UserArgs.builder()
            .name(userName)
            .password(userPw)
            .mechanism(mechanism)
            .clusterApiUrl(test.clusterApiUrl())
            .allowDeletion(userAllowDeletion)
            .build());

        var testAcl = new Acl("testAcl", AclArgs.builder()
            .resourceType("CLUSTER")
            .resourceName("kafka-cluster")
            .resourcePatternType("LITERAL")
            .principal(testUser.name().applyValue(_name -> String.format("User:%s", _name)))
            .host("*")
            .operation("ALTER")
            .permissionType("ALLOW")
            .clusterApiUrl(test.clusterApiUrl())
            .allowDeletion(aclAllowDeletion)
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

const test = new redpanda.ResourceGroup("test", {name: resourceGroupName});
const testServerlessCluster = new redpanda.ServerlessCluster("test", {
    name: clusterName,
    resourceGroupId: test.id,
    serverlessRegion: region,
});
const testTopic = new redpanda.Topic("test", {
    name: topicName,
    partitionCount: partitionCount,
    replicationFactor: replicationFactor,
    clusterApiUrl: testServerlessCluster.clusterApiUrl,
    allowDeletion: true,
});
const testUser = new redpanda.User("test", {
    name: userName,
    password: userPw,
    mechanism: mechanism,
    clusterApiUrl: testServerlessCluster.clusterApiUrl,
    allowDeletion: userAllowDeletion,
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

test = redpanda.ResourceGroup("test", name=resource_group_name)
test_serverless_cluster = redpanda.ServerlessCluster("test",
    name=cluster_name,
    resource_group_id=test.id,
    serverless_region=region)
test_topic = redpanda.Topic("test",
    name=topic_name,
    partition_count=partition_count,
    replication_factor=replication_factor,
    cluster_api_url=test_serverless_cluster.cluster_api_url,
    allow_deletion=True)
test_user = redpanda.User("test",
    name=user_name,
    password=user_pw,
    mechanism=mechanism,
    cluster_api_url=test_serverless_cluster.cluster_api_url,
    allow_deletion=user_allow_deletion)
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
    var test = new Redpanda.ResourceGroup("test", new()
    {
        Name = resourceGroupName,
    });

    var testServerlessCluster = new Redpanda.ServerlessCluster("test", new()
    {
        Name = clusterName,
        ResourceGroupId = test.Id,
        ServerlessRegion = region,
    });

    var testTopic = new Redpanda.Topic("test", new()
    {
        Name = topicName,
        PartitionCount = partitionCount,
        ReplicationFactor = replicationFactor,
        ClusterApiUrl = testServerlessCluster.ClusterApiUrl,
        AllowDeletion = true,
    });

    var testUser = new Redpanda.User("test", new()
    {
        Name = userName,
        Password = userPw,
        Mechanism = mechanism,
        ClusterApiUrl = testServerlessCluster.ClusterApiUrl,
        AllowDeletion = userAllowDeletion,
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
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		test, err := redpanda.NewResourceGroup(ctx, "test", &redpanda.ResourceGroupArgs{
			Name: pulumi.Any(resourceGroupName),
		})
		if err != nil {
			return err
		}
		testServerlessCluster, err := redpanda.NewServerlessCluster(ctx, "test", &redpanda.ServerlessClusterArgs{
			Name:             pulumi.Any(clusterName),
			ResourceGroupId:  test.ID(),
			ServerlessRegion: pulumi.Any(region),
		})
		if err != nil {
			return err
		}
		_, err = redpanda.NewTopic(ctx, "test", &redpanda.TopicArgs{
			Name:              pulumi.Any(topicName),
			PartitionCount:    pulumi.Any(partitionCount),
			ReplicationFactor: pulumi.Any(replicationFactor),
			ClusterApiUrl:     testServerlessCluster.ClusterApiUrl,
			AllowDeletion:     pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		_, err = redpanda.NewUser(ctx, "test", &redpanda.UserArgs{
			Name:          pulumi.Any(userName),
			Password:      pulumi.Any(userPw),
			Mechanism:     pulumi.Any(mechanism),
			ClusterApiUrl: testServerlessCluster.ClusterApiUrl,
			AllowDeletion: pulumi.Any(userAllowDeletion),
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
  testTopic:
    type: redpanda:Topic
    name: test
    properties:
      name: ${topicName}
      partitionCount: ${partitionCount}
      replicationFactor: ${replicationFactor}
      clusterApiUrl: ${testServerlessCluster.clusterApiUrl}
      allowDeletion: true
  testUser:
    type: redpanda:User
    name: test
    properties:
      name: ${userName}
      password: ${userPw}
      mechanism: ${mechanism}
      clusterApiUrl: ${testServerlessCluster.clusterApiUrl}
      allowDeletion: ${userAllowDeletion}
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
import com.pulumi.redpanda.Topic;
import com.pulumi.redpanda.TopicArgs;
import com.pulumi.redpanda.User;
import com.pulumi.redpanda.UserArgs;
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
        var test = new ResourceGroup("test", ResourceGroupArgs.builder()
            .name(resourceGroupName)
            .build());

        var testServerlessCluster = new ServerlessCluster("testServerlessCluster", ServerlessClusterArgs.builder()
            .name(clusterName)
            .resourceGroupId(test.id())
            .serverlessRegion(region)
            .build());

        var testTopic = new Topic("testTopic", TopicArgs.builder()
            .name(topicName)
            .partitionCount(partitionCount)
            .replicationFactor(replicationFactor)
            .clusterApiUrl(testServerlessCluster.clusterApiUrl())
            .allowDeletion(true)
            .build());

        var testUser = new User("testUser", UserArgs.builder()
            .name(userName)
            .password(userPw)
            .mechanism(mechanism)
            .clusterApiUrl(testServerlessCluster.clusterApiUrl())
            .allowDeletion(userAllowDeletion)
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}