---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/redpanda-data/redpanda/2.2.0/index.md
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
- `awsAccessKeyId` (String, Sensitive) AWS access key ID for BYOC clusters. Can also be set via AWS_ACCESS_KEY_ID.
- `awsSecretAccessKey` (String, Sensitive) AWS secret access key for BYOC clusters. Can also be set via AWS_SECRET_ACCESS_KEY.
- `awsSessionToken` (String, Sensitive) AWS session token for BYOC clusters (for temporary credentials). Can also be set via AWS_SESSION_TOKEN.
- `azureClientId` (String) Used for creating and managing BYOC and BYOVPC clusters. Can also be specified in the environment as AZURE_CLIENT_ID or ARM_CLIENT_ID
- `azureClientSecret` (String, Sensitive) Used for creating and managing BYOC and BYOVPC clusters. Can also be specified in the environment as AZURE_CLIENT_SECRET or ARM_CLIENT_SECRET
- `azureSubscriptionId` (String) The default Azure Subscription ID which should be used for Redpanda BYOC clusters. If another subscription is specified on a resource, it will take precedence. This can also be sourced from the `ARM_SUBSCRIPTION_ID` environment variable.
- `azureTenantId` (String) Used for creating and managing BYOC and BYOVPC clusters. Can also be specified in the environment as AZURE_TENANT_ID or ARM_TENANT_ID
- `clientId` (String, Sensitive) The ID for the client. You need either `clientId` AND `clientSecret`, or `accessToken`, to use this provider. Can also be set with the `REDPANDA_CLIENT_ID` environment variable.
- `clientSecret` (String, Sensitive) Redpanda client secret. You need either `clientId` AND `clientSecret`, or `accessToken`, to use this provider. Can also be set with the `REDPANDA_CLIENT_SECRET` environment variable.
- `gcpProjectId` (String) The default Google Cloud Project ID to use for Redpanda BYOC clusters. If another project is specified on a resource, it will take precedence. This can also be sourced from the `GOOGLE_PROJECT` environment variable, or any of the following ordered by precedence: `GOOGLE_PROJECT`, `GOOGLE_CLOUD_PROJECT`, `GCLOUD_PROJECT`, or `CLOUDSDK_CORE_PROJECT`.
- `googleCredentials` (String, Sensitive) Used for creating and managing BYOC and BYOVPC clusters. Can also be specified in the environment as GOOGLE_CREDENTIALS
- `googleCredentialsBase64` (String, Sensitive) Used for creating and managing BYOC and BYOVPC clusters. Is a convenience passthrough for base64 encoded credentials intended for use in CI/CD. Can also be specified in the environment as GOOGLE_CREDENTIALS_BASE64
## Authentication with Redpanda Cloud

This provider requires a `clientId` and `clientSecret` for authentication with Redpanda Cloud services, enabling users to securely manage their Redpanda resources. You can get these by creating an account in [Redpanda Cloud](https://cloudv2.redpanda.com/home) and then [creating a client in the Redpanda Cloud UI](https://cloudv2.redpanda.com/clients).
### Token cache

The provider caches the OAuth2 access token to disk so that subsequent
`pulumi` invocations reuse it until expiry, instead of issuing a fresh
token on every plan or apply. This keeps usage well under the per-organization
daily token-issuance quota.

- **Location:** `<user-cache-dir>/redpanda/provider/creds-cache.json`. The
  user cache directory is `~/Library/Caches` on macOS, `$XDG_CACHE_HOME`
  (falling back to `~/.cache`) on Linux, and `%LocalAppData%` on Windows.
- **Permissions:** the file is written with mode `0600`; the
  `redpanda/provider/` directory is created with mode `0755`.
- **Cache key:** tokens are keyed by `<audience>:<client_id>`, so different
  Redpanda environments and clients each get their own entry.
- **Disable:** set `REDPANDA_TOKEN_CACHE_DISABLE=1` to skip the on-disk
  cache; a fresh token is fetched per provider invocation.
- **Clear:** delete the file (`rm <user-cache-dir>/redpanda/provider/creds-cache.json`).
  Useful if a server-side secret rotation invalidates the cached token before
  its local expiry.
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
            enable_shadow_linking: clusterEnableShadowLinking,
            schema_registry_enable_authorization: true,
        }),
    },
    maintenanceWindowConfig: {
        dayHour: {
            dayOfWeek: maintenanceDayOfWeek,
            hourOfDay: maintenanceHourOfDay,
        },
    },
    tags: clusterTags,
    timeouts: {
        create: "90m",
    },
});
const testTopic = new redpanda.Topic("test", {
    name: topicName,
    partitionCount: partitionCount,
    replicationFactor: replicationFactor,
    clusterApiUrl: testCluster.clusterApiUrl,
    allowDeletion: true,
    configuration: topicConfiguration != null ? topicConfiguration : {
        "cleanup.policy": "delete",
        "retention.ms": topicRetentionMs,
    },
});
const testUser = new redpanda.User("test", {
    name: userName,
    password: userPasswordWo != null ? null : userPw,
    passwordWo: userPasswordWo,
    passwordWoVersion: userPasswordWoVersion,
    mechanism: mechanism,
    clusterApiUrl: testCluster.clusterApiUrl,
    allowDeletion: userAllowDeletion,
}, {
    dependsOn: [testTopic],
});
// Console-endpoint canary. The topic canary proves the dataplane; the console API
// is a separate endpoint with separate readiness, so it needs its own. Role is the
// simplest console resource, and the console resources below wait on it.
const consoleCanary = new redpanda.Role("console_canary", {
    name: `${roleName}-console-canary`,
    clusterApiUrl: testCluster.clusterApiUrl,
    allowDeletion: true,
});
const developer = new redpanda.Role("developer", {
    name: roleName,
    clusterApiUrl: testCluster.clusterApiUrl,
    allowDeletion: roleAllowDeletion,
}, {
    dependsOn: [
        testTopic,
        consoleCanary,
    ],
});
// Bootstrap SR ACL grants for the provider's own Bearer-token principal.
// Two grants are required: the SUBJECT-scope grant authorizes
// POST /subjects/<subj>/versions; the REGISTRY-scope grant authorizes the
// follow-up GET /schemas/ids/<id>/versions that the SR client (franz-go) makes
// to fetch the full schema metadata after create. Without the REGISTRY grant,
// the schema is created but the follow-up GET 403s and the resource fails.
// User:* is broader than necessary; tighten once the exact principal is
// documented as discoverable.
const roleTopicRead = new redpanda.Acl("role_topic_read", {
    resourceType: "TOPIC",
    resourceName: testTopic.name,
    resourcePatternType: "LITERAL",
    principal: pulumi.interpolate`RedpandaRole:${developer.name}`,
    host: "*",
    operation: "READ",
    permissionType: "ALLOW",
    clusterApiUrl: testCluster.clusterApiUrl,
    allowDeletion: aclAllowDeletion,
});
const developerAssignment = new redpanda.RoleAssignment("developer_assignment", {
    roleName: developer.name,
    principal: pulumi.interpolate`User:${testUser.name}`,
    clusterApiUrl: testCluster.clusterApiUrl,
}, {
    dependsOn: [
        testUser,
        consoleCanary,
    ],
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
            "enable_shadow_linking": cluster_enable_shadow_linking,
            "schema_registry_enable_authorization": True,
        }),
    },
    maintenance_window_config={
        "day_hour": {
            "day_of_week": maintenance_day_of_week,
            "hour_of_day": maintenance_hour_of_day,
        },
    },
    tags=cluster_tags,
    timeouts={
        "create": "90m",
    })
test_topic = redpanda.Topic("test",
    name=topic_name,
    partition_count=partition_count,
    replication_factor=replication_factor,
    cluster_api_url=test_cluster.cluster_api_url,
    allow_deletion=True,
    configuration=topic_configuration if topic_configuration != None else {
        "cleanup.policy": "delete",
        "retention.ms": topic_retention_ms,
    })
test_user = redpanda.User("test",
    name=user_name,
    password=None if user_password_wo != None else user_pw,
    password_wo=user_password_wo,
    password_wo_version=user_password_wo_version,
    mechanism=mechanism,
    cluster_api_url=test_cluster.cluster_api_url,
    allow_deletion=user_allow_deletion,
    opts = pulumi.ResourceOptions(depends_on=[test_topic]))
# Console-endpoint canary. The topic canary proves the dataplane; the console API
# is a separate endpoint with separate readiness, so it needs its own. Role is the
# simplest console resource, and the console resources below wait on it.
console_canary = redpanda.Role("console_canary",
    name=f"{role_name}-console-canary",
    cluster_api_url=test_cluster.cluster_api_url,
    allow_deletion=True)
developer = redpanda.Role("developer",
    name=role_name,
    cluster_api_url=test_cluster.cluster_api_url,
    allow_deletion=role_allow_deletion,
    opts = pulumi.ResourceOptions(depends_on=[
            test_topic,
            console_canary,
        ]))
# Bootstrap SR ACL grants for the provider's own Bearer-token principal.
# Two grants are required: the SUBJECT-scope grant authorizes
# POST /subjects/<subj>/versions; the REGISTRY-scope grant authorizes the
# follow-up GET /schemas/ids/<id>/versions that the SR client (franz-go) makes
# to fetch the full schema metadata after create. Without the REGISTRY grant,
# the schema is created but the follow-up GET 403s and the resource fails.
# User:* is broader than necessary; tighten once the exact principal is
# documented as discoverable.
role_topic_read = redpanda.Acl("role_topic_read",
    resource_type="TOPIC",
    resource_name_=test_topic.name,
    resource_pattern_type="LITERAL",
    principal=developer.name.apply(lambda name: f"RedpandaRole:{name}"),
    host="*",
    operation="READ",
    permission_type="ALLOW",
    cluster_api_url=test_cluster.cluster_api_url,
    allow_deletion=acl_allow_deletion)
developer_assignment = redpanda.RoleAssignment("developer_assignment",
    role_name=developer.name,
    principal=test_user.name.apply(lambda name: f"User:{name}"),
    cluster_api_url=test_cluster.cluster_api_url,
    opts = pulumi.ResourceOptions(depends_on=[
            test_user,
            console_canary,
        ]))
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
                ["enable_shadow_linking"] = clusterEnableShadowLinking,
                ["schema_registry_enable_authorization"] = true,
            }),
        },
        MaintenanceWindowConfig = new Redpanda.Inputs.ClusterMaintenanceWindowConfigArgs
        {
            DayHour = new Redpanda.Inputs.ClusterMaintenanceWindowConfigDayHourArgs
            {
                DayOfWeek = maintenanceDayOfWeek,
                HourOfDay = maintenanceHourOfDay,
            },
        },
        Tags = clusterTags,
        Timeouts = new Redpanda.Inputs.ClusterTimeoutsArgs
        {
            Create = "90m",
        },
    });

    var testTopic = new Redpanda.Topic("test", new()
    {
        Name = topicName,
        PartitionCount = partitionCount,
        ReplicationFactor = replicationFactor,
        ClusterApiUrl = testCluster.ClusterApiUrl,
        AllowDeletion = true,
        Configuration = topicConfiguration != null ? topicConfiguration :
        {
            { "cleanup.policy", "delete" },
            { "retention.ms", topicRetentionMs },
        },
    });

    var testUser = new Redpanda.User("test", new()
    {
        Name = userName,
        Password = userPasswordWo != null ? null : userPw,
        PasswordWo = userPasswordWo,
        PasswordWoVersion = userPasswordWoVersion,
        Mechanism = mechanism,
        ClusterApiUrl = testCluster.ClusterApiUrl,
        AllowDeletion = userAllowDeletion,
    }, new CustomResourceOptions
    {
        DependsOn =
        {
            testTopic,
        },
    });

    // Console-endpoint canary. The topic canary proves the dataplane; the console API
    // is a separate endpoint with separate readiness, so it needs its own. Role is the
    // simplest console resource, and the console resources below wait on it.
    var consoleCanary = new Redpanda.Role("console_canary", new()
    {
        Name = $"{roleName}-console-canary",
        ClusterApiUrl = testCluster.ClusterApiUrl,
        AllowDeletion = true,
    });

    var developer = new Redpanda.Role("developer", new()
    {
        Name = roleName,
        ClusterApiUrl = testCluster.ClusterApiUrl,
        AllowDeletion = roleAllowDeletion,
    }, new CustomResourceOptions
    {
        DependsOn =
        {
            testTopic,
            consoleCanary,
        },
    });

    // Bootstrap SR ACL grants for the provider's own Bearer-token principal.
    // Two grants are required: the SUBJECT-scope grant authorizes
    // POST /subjects/<subj>/versions; the REGISTRY-scope grant authorizes the
    // follow-up GET /schemas/ids/<id>/versions that the SR client (franz-go) makes
    // to fetch the full schema metadata after create. Without the REGISTRY grant,
    // the schema is created but the follow-up GET 403s and the resource fails.
    // User:* is broader than necessary; tighten once the exact principal is
    // documented as discoverable.
    var roleTopicRead = new Redpanda.Acl("role_topic_read", new()
    {
        ResourceType = "TOPIC",
        ResourceName = testTopic.Name,
        ResourcePatternType = "LITERAL",
        Principal = developer.Name.Apply(name => $"RedpandaRole:{name}"),
        Host = "*",
        Operation = "READ",
        PermissionType = "ALLOW",
        ClusterApiUrl = testCluster.ClusterApiUrl,
        AllowDeletion = aclAllowDeletion,
    });

    var developerAssignment = new Redpanda.RoleAssignment("developer_assignment", new()
    {
        RoleName = developer.Name,
        Principal = testUser.Name.Apply(name => $"User:{name}"),
        ClusterApiUrl = testCluster.ClusterApiUrl,
    }, new CustomResourceOptions
    {
        DependsOn =
        {
            testUser,
            consoleCanary,
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

```
```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/redpanda/v2/redpanda"
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
			"enable_shadow_linking":                clusterEnableShadowLinking,
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
			MaintenanceWindowConfig: &redpanda.ClusterMaintenanceWindowConfigArgs{
				DayHour: &redpanda.ClusterMaintenanceWindowConfigDayHourArgs{
					DayOfWeek: pulumi.Any(maintenanceDayOfWeek),
					HourOfDay: pulumi.Any(maintenanceHourOfDay),
				},
			},
			Tags: pulumi.Any(clusterTags),
			Timeouts: &redpanda.ClusterTimeoutsArgs{
				Create: pulumi.String("90m"),
			},
		})
		if err != nil {
			return err
		}
		var tmp0 pulumi.StringMap
		if topicConfiguration != nil {
			tmp0 = pulumi.Any(topicConfiguration)
		} else {
			tmp0 = pulumi.StringMap{
				"cleanup.policy": pulumi.String("delete"),
				"retention.ms":   pulumi.Any(topicRetentionMs),
			}
		}
		testTopic, err := redpanda.NewTopic(ctx, "test", &redpanda.TopicArgs{
			Name:              pulumi.Any(topicName),
			PartitionCount:    pulumi.Any(partitionCount),
			ReplicationFactor: pulumi.Any(replicationFactor),
			ClusterApiUrl:     testCluster.ClusterApiUrl,
			AllowDeletion:     pulumi.Bool(true),
			Configuration:     pulumi.StringMap(tmp0),
		})
		if err != nil {
			return err
		}
		var tmp1 pulumi.String
		if userPasswordWo != nil {
			tmp1 = nil
		} else {
			tmp1 = pulumi.Any(userPw)
		}
		testUser, err := redpanda.NewUser(ctx, "test", &redpanda.UserArgs{
			Name:              pulumi.Any(userName),
			Password:          pulumi.String(tmp1),
			PasswordWo:        pulumi.Any(userPasswordWo),
			PasswordWoVersion: pulumi.Any(userPasswordWoVersion),
			Mechanism:         pulumi.Any(mechanism),
			ClusterApiUrl:     testCluster.ClusterApiUrl,
			AllowDeletion:     pulumi.Any(userAllowDeletion),
		}, pulumi.DependsOn([]pulumi.Resource{
			testTopic,
		}))
		if err != nil {
			return err
		}
		// Console-endpoint canary. The topic canary proves the dataplane; the console API
		// is a separate endpoint with separate readiness, so it needs its own. Role is the
		// simplest console resource, and the console resources below wait on it.
		consoleCanary, err := redpanda.NewRole(ctx, "console_canary", &redpanda.RoleArgs{
			Name:          pulumi.Sprintf("%v-console-canary", roleName),
			ClusterApiUrl: testCluster.ClusterApiUrl,
			AllowDeletion: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		developer, err := redpanda.NewRole(ctx, "developer", &redpanda.RoleArgs{
			Name:          pulumi.Any(roleName),
			ClusterApiUrl: testCluster.ClusterApiUrl,
			AllowDeletion: pulumi.Any(roleAllowDeletion),
		}, pulumi.DependsOn([]pulumi.Resource{
			testTopic,
			consoleCanary,
		}))
		if err != nil {
			return err
		}
		// Bootstrap SR ACL grants for the provider's own Bearer-token principal.
		// Two grants are required: the SUBJECT-scope grant authorizes
		// POST /subjects/<subj>/versions; the REGISTRY-scope grant authorizes the
		// follow-up GET /schemas/ids/<id>/versions that the SR client (franz-go) makes
		// to fetch the full schema metadata after create. Without the REGISTRY grant,
		// the schema is created but the follow-up GET 403s and the resource fails.
		// User:* is broader than necessary; tighten once the exact principal is
		// documented as discoverable.
		_, err = redpanda.NewAcl(ctx, "role_topic_read", &redpanda.AclArgs{
			ResourceType:        pulumi.String("TOPIC"),
			ResourceName:        testTopic.Name,
			ResourcePatternType: pulumi.String("LITERAL"),
			Principal: developer.Name.ApplyT(func(name string) (string, error) {
				return fmt.Sprintf("RedpandaRole:%v", name), nil
			}).(pulumi.StringOutput),
			Host:           pulumi.String("*"),
			Operation:      pulumi.String("READ"),
			PermissionType: pulumi.String("ALLOW"),
			ClusterApiUrl:  testCluster.ClusterApiUrl,
			AllowDeletion:  pulumi.Any(aclAllowDeletion),
		})
		if err != nil {
			return err
		}
		_, err = redpanda.NewRoleAssignment(ctx, "developer_assignment", &redpanda.RoleAssignmentArgs{
			RoleName: developer.Name,
			Principal: testUser.Name.ApplyT(func(name string) (string, error) {
				return fmt.Sprintf("User:%v", name), nil
			}).(pulumi.StringOutput),
			ClusterApiUrl: testCluster.ClusterApiUrl,
		}, pulumi.DependsOn([]pulumi.Resource{
			testUser,
			consoleCanary,
		}))
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
Example currently unavailable in this language
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
import com.pulumi.redpanda.inputs.ClusterMaintenanceWindowConfigArgs;
import com.pulumi.redpanda.inputs.ClusterMaintenanceWindowConfigDayHourArgs;
import com.pulumi.redpanda.inputs.ClusterTimeoutsArgs;
import com.pulumi.redpanda.Topic;
import com.pulumi.redpanda.TopicArgs;
import com.pulumi.redpanda.User;
import com.pulumi.redpanda.UserArgs;
import com.pulumi.redpanda.Role;
import com.pulumi.redpanda.RoleArgs;
import com.pulumi.redpanda.Acl;
import com.pulumi.redpanda.AclArgs;
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
                        jsonProperty("enable_shadow_linking", clusterEnableShadowLinking),
                        jsonProperty("schema_registry_enable_authorization", true)
                    )))
                .build())
            .maintenanceWindowConfig(ClusterMaintenanceWindowConfigArgs.builder()
                .dayHour(ClusterMaintenanceWindowConfigDayHourArgs.builder()
                    .dayOfWeek(maintenanceDayOfWeek)
                    .hourOfDay(maintenanceHourOfDay)
                    .build())
                .build())
            .tags(clusterTags)
            .timeouts(ClusterTimeoutsArgs.builder()
                .create("90m")
                .build())
            .build());

        var testTopic = new Topic("testTopic", TopicArgs.builder()
            .name(topicName)
            .partitionCount(partitionCount)
            .replicationFactor(replicationFactor)
            .clusterApiUrl(testCluster.clusterApiUrl())
            .allowDeletion(true)
            .configuration(topicConfiguration != null ? topicConfiguration : Map.ofEntries(
                Map.entry("cleanup.policy", "delete"),
                Map.entry("retention.ms", topicRetentionMs)
            ))
            .build());

        var testUser = new User("testUser", UserArgs.builder()
            .name(userName)
            .password(userPasswordWo != null ? null : userPw)
            .passwordWo(userPasswordWo)
            .passwordWoVersion(userPasswordWoVersion)
            .mechanism(mechanism)
            .clusterApiUrl(testCluster.clusterApiUrl())
            .allowDeletion(userAllowDeletion)
            .build(), CustomResourceOptions.builder()
                .dependsOn(testTopic)
                .build());

        // Console-endpoint canary. The topic canary proves the dataplane; the console API
        // is a separate endpoint with separate readiness, so it needs its own. Role is the
        // simplest console resource, and the console resources below wait on it.
        var consoleCanary = new Role("consoleCanary", RoleArgs.builder()
            .name(String.format("%s-console-canary", roleName))
            .clusterApiUrl(testCluster.clusterApiUrl())
            .allowDeletion(true)
            .build());

        var developer = new Role("developer", RoleArgs.builder()
            .name(roleName)
            .clusterApiUrl(testCluster.clusterApiUrl())
            .allowDeletion(roleAllowDeletion)
            .build(), CustomResourceOptions.builder()
                .dependsOn(
                    testTopic,
                    consoleCanary)
                .build());

        // Bootstrap SR ACL grants for the provider's own Bearer-token principal.
        // Two grants are required: the SUBJECT-scope grant authorizes
        // POST /subjects/<subj>/versions; the REGISTRY-scope grant authorizes the
        // follow-up GET /schemas/ids/<id>/versions that the SR client (franz-go) makes
        // to fetch the full schema metadata after create. Without the REGISTRY grant,
        // the schema is created but the follow-up GET 403s and the resource fails.
        // User:* is broader than necessary; tighten once the exact principal is
        // documented as discoverable.
        var roleTopicRead = new Acl("roleTopicRead", AclArgs.builder()
            .resourceType("TOPIC")
            .resourceName(testTopic.name())
            .resourcePatternType("LITERAL")
            .principal(developer.name().applyValue(_name -> String.format("RedpandaRole:%s", _name)))
            .host("*")
            .operation("READ")
            .permissionType("ALLOW")
            .clusterApiUrl(testCluster.clusterApiUrl())
            .allowDeletion(aclAllowDeletion)
            .build());

        var developerAssignment = new RoleAssignment("developerAssignment", RoleAssignmentArgs.builder()
            .roleName(developer.name())
            .principal(testUser.name().applyValue(_name -> String.format("User:%s", _name)))
            .clusterApiUrl(testCluster.clusterApiUrl())
            .build(), CustomResourceOptions.builder()
                .dependsOn(
                    testUser,
                    consoleCanary)
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
            enable_shadow_linking: clusterEnableShadowLinking,
            schema_registry_enable_authorization: true,
        }),
    },
    maintenanceWindowConfig: {
        dayHour: {
            dayOfWeek: maintenanceDayOfWeek,
            hourOfDay: maintenanceHourOfDay,
        },
    },
    tags: clusterTags,
    timeouts: {
        create: "90m",
    },
});
const testTopic = new redpanda.Topic("test", {
    name: topicName,
    partitionCount: partitionCount,
    replicationFactor: replicationFactor,
    clusterApiUrl: testCluster.clusterApiUrl,
    allowDeletion: true,
    configuration: topicConfiguration != null ? topicConfiguration : {
        "cleanup.policy": "delete",
        "retention.ms": topicRetentionMs,
    },
});
const testUser = new redpanda.User("test", {
    name: userName,
    password: userPw,
    mechanism: mechanism,
    clusterApiUrl: testCluster.clusterApiUrl,
    allowDeletion: userAllowDeletion,
}, {
    dependsOn: [testTopic],
});
// Console-endpoint canary. The topic canary proves the dataplane; the console API
// is a separate endpoint with separate readiness, so it needs its own. Role is the
// simplest console resource, and the console resources below wait on it.
const consoleCanary = new redpanda.Role("console_canary", {
    name: `${roleName}-console-canary`,
    clusterApiUrl: testCluster.clusterApiUrl,
    allowDeletion: true,
});
const developer = new redpanda.Role("developer", {
    name: roleName,
    clusterApiUrl: testCluster.clusterApiUrl,
    allowDeletion: roleAllowDeletion,
}, {
    dependsOn: [
        testTopic,
        consoleCanary,
    ],
});
// Bootstrap SR ACL grants for the provider's own Bearer-token principal.
// SUBJECT grant authorizes POST /subjects/<subj>/versions; REGISTRY grant
// authorizes the franz-go follow-up GET /schemas/ids/<id>/versions.
const roleTopicRead = new redpanda.Acl("role_topic_read", {
    resourceType: "TOPIC",
    resourceName: testTopic.name,
    resourcePatternType: "LITERAL",
    principal: pulumi.interpolate`RedpandaRole:${developer.name}`,
    host: "*",
    operation: "READ",
    permissionType: "ALLOW",
    clusterApiUrl: testCluster.clusterApiUrl,
    allowDeletion: aclAllowDeletion,
});
const developerAssignment = new redpanda.RoleAssignment("developer_assignment", {
    roleName: developer.name,
    principal: pulumi.interpolate`User:${testUser.name}`,
    clusterApiUrl: testCluster.clusterApiUrl,
}, {
    dependsOn: [
        testUser,
        consoleCanary,
    ],
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
            "enable_shadow_linking": cluster_enable_shadow_linking,
            "schema_registry_enable_authorization": True,
        }),
    },
    maintenance_window_config={
        "day_hour": {
            "day_of_week": maintenance_day_of_week,
            "hour_of_day": maintenance_hour_of_day,
        },
    },
    tags=cluster_tags,
    timeouts={
        "create": "90m",
    })
test_topic = redpanda.Topic("test",
    name=topic_name,
    partition_count=partition_count,
    replication_factor=replication_factor,
    cluster_api_url=test_cluster.cluster_api_url,
    allow_deletion=True,
    configuration=topic_configuration if topic_configuration != None else {
        "cleanup.policy": "delete",
        "retention.ms": topic_retention_ms,
    })
test_user = redpanda.User("test",
    name=user_name,
    password=user_pw,
    mechanism=mechanism,
    cluster_api_url=test_cluster.cluster_api_url,
    allow_deletion=user_allow_deletion,
    opts = pulumi.ResourceOptions(depends_on=[test_topic]))
# Console-endpoint canary. The topic canary proves the dataplane; the console API
# is a separate endpoint with separate readiness, so it needs its own. Role is the
# simplest console resource, and the console resources below wait on it.
console_canary = redpanda.Role("console_canary",
    name=f"{role_name}-console-canary",
    cluster_api_url=test_cluster.cluster_api_url,
    allow_deletion=True)
developer = redpanda.Role("developer",
    name=role_name,
    cluster_api_url=test_cluster.cluster_api_url,
    allow_deletion=role_allow_deletion,
    opts = pulumi.ResourceOptions(depends_on=[
            test_topic,
            console_canary,
        ]))
# Bootstrap SR ACL grants for the provider's own Bearer-token principal.
# SUBJECT grant authorizes POST /subjects/<subj>/versions; REGISTRY grant
# authorizes the franz-go follow-up GET /schemas/ids/<id>/versions.
role_topic_read = redpanda.Acl("role_topic_read",
    resource_type="TOPIC",
    resource_name_=test_topic.name,
    resource_pattern_type="LITERAL",
    principal=developer.name.apply(lambda name: f"RedpandaRole:{name}"),
    host="*",
    operation="READ",
    permission_type="ALLOW",
    cluster_api_url=test_cluster.cluster_api_url,
    allow_deletion=acl_allow_deletion)
developer_assignment = redpanda.RoleAssignment("developer_assignment",
    role_name=developer.name,
    principal=test_user.name.apply(lambda name: f"User:{name}"),
    cluster_api_url=test_cluster.cluster_api_url,
    opts = pulumi.ResourceOptions(depends_on=[
            test_user,
            console_canary,
        ]))
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
                ["enable_shadow_linking"] = clusterEnableShadowLinking,
                ["schema_registry_enable_authorization"] = true,
            }),
        },
        MaintenanceWindowConfig = new Redpanda.Inputs.ClusterMaintenanceWindowConfigArgs
        {
            DayHour = new Redpanda.Inputs.ClusterMaintenanceWindowConfigDayHourArgs
            {
                DayOfWeek = maintenanceDayOfWeek,
                HourOfDay = maintenanceHourOfDay,
            },
        },
        Tags = clusterTags,
        Timeouts = new Redpanda.Inputs.ClusterTimeoutsArgs
        {
            Create = "90m",
        },
    });

    var testTopic = new Redpanda.Topic("test", new()
    {
        Name = topicName,
        PartitionCount = partitionCount,
        ReplicationFactor = replicationFactor,
        ClusterApiUrl = testCluster.ClusterApiUrl,
        AllowDeletion = true,
        Configuration = topicConfiguration != null ? topicConfiguration :
        {
            { "cleanup.policy", "delete" },
            { "retention.ms", topicRetentionMs },
        },
    });

    var testUser = new Redpanda.User("test", new()
    {
        Name = userName,
        Password = userPw,
        Mechanism = mechanism,
        ClusterApiUrl = testCluster.ClusterApiUrl,
        AllowDeletion = userAllowDeletion,
    }, new CustomResourceOptions
    {
        DependsOn =
        {
            testTopic,
        },
    });

    // Console-endpoint canary. The topic canary proves the dataplane; the console API
    // is a separate endpoint with separate readiness, so it needs its own. Role is the
    // simplest console resource, and the console resources below wait on it.
    var consoleCanary = new Redpanda.Role("console_canary", new()
    {
        Name = $"{roleName}-console-canary",
        ClusterApiUrl = testCluster.ClusterApiUrl,
        AllowDeletion = true,
    });

    var developer = new Redpanda.Role("developer", new()
    {
        Name = roleName,
        ClusterApiUrl = testCluster.ClusterApiUrl,
        AllowDeletion = roleAllowDeletion,
    }, new CustomResourceOptions
    {
        DependsOn =
        {
            testTopic,
            consoleCanary,
        },
    });

    // Bootstrap SR ACL grants for the provider's own Bearer-token principal.
    // SUBJECT grant authorizes POST /subjects/<subj>/versions; REGISTRY grant
    // authorizes the franz-go follow-up GET /schemas/ids/<id>/versions.
    var roleTopicRead = new Redpanda.Acl("role_topic_read", new()
    {
        ResourceType = "TOPIC",
        ResourceName = testTopic.Name,
        ResourcePatternType = "LITERAL",
        Principal = developer.Name.Apply(name => $"RedpandaRole:{name}"),
        Host = "*",
        Operation = "READ",
        PermissionType = "ALLOW",
        ClusterApiUrl = testCluster.ClusterApiUrl,
        AllowDeletion = aclAllowDeletion,
    });

    var developerAssignment = new Redpanda.RoleAssignment("developer_assignment", new()
    {
        RoleName = developer.Name,
        Principal = testUser.Name.Apply(name => $"User:{name}"),
        ClusterApiUrl = testCluster.ClusterApiUrl,
    }, new CustomResourceOptions
    {
        DependsOn =
        {
            testUser,
            consoleCanary,
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

```
```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/redpanda/v2/redpanda"
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
			"enable_shadow_linking":                clusterEnableShadowLinking,
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
			MaintenanceWindowConfig: &redpanda.ClusterMaintenanceWindowConfigArgs{
				DayHour: &redpanda.ClusterMaintenanceWindowConfigDayHourArgs{
					DayOfWeek: pulumi.Any(maintenanceDayOfWeek),
					HourOfDay: pulumi.Any(maintenanceHourOfDay),
				},
			},
			Tags: pulumi.Any(clusterTags),
			Timeouts: &redpanda.ClusterTimeoutsArgs{
				Create: pulumi.String("90m"),
			},
		})
		if err != nil {
			return err
		}
		var tmp0 pulumi.StringMap
		if topicConfiguration != nil {
			tmp0 = pulumi.Any(topicConfiguration)
		} else {
			tmp0 = pulumi.StringMap{
				"cleanup.policy": pulumi.String("delete"),
				"retention.ms":   pulumi.Any(topicRetentionMs),
			}
		}
		testTopic, err := redpanda.NewTopic(ctx, "test", &redpanda.TopicArgs{
			Name:              pulumi.Any(topicName),
			PartitionCount:    pulumi.Any(partitionCount),
			ReplicationFactor: pulumi.Any(replicationFactor),
			ClusterApiUrl:     testCluster.ClusterApiUrl,
			AllowDeletion:     pulumi.Bool(true),
			Configuration:     pulumi.StringMap(tmp0),
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
		}, pulumi.DependsOn([]pulumi.Resource{
			testTopic,
		}))
		if err != nil {
			return err
		}
		// Console-endpoint canary. The topic canary proves the dataplane; the console API
		// is a separate endpoint with separate readiness, so it needs its own. Role is the
		// simplest console resource, and the console resources below wait on it.
		consoleCanary, err := redpanda.NewRole(ctx, "console_canary", &redpanda.RoleArgs{
			Name:          pulumi.Sprintf("%v-console-canary", roleName),
			ClusterApiUrl: testCluster.ClusterApiUrl,
			AllowDeletion: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		developer, err := redpanda.NewRole(ctx, "developer", &redpanda.RoleArgs{
			Name:          pulumi.Any(roleName),
			ClusterApiUrl: testCluster.ClusterApiUrl,
			AllowDeletion: pulumi.Any(roleAllowDeletion),
		}, pulumi.DependsOn([]pulumi.Resource{
			testTopic,
			consoleCanary,
		}))
		if err != nil {
			return err
		}
		// Bootstrap SR ACL grants for the provider's own Bearer-token principal.
		// SUBJECT grant authorizes POST /subjects/<subj>/versions; REGISTRY grant
		// authorizes the franz-go follow-up GET /schemas/ids/<id>/versions.
		_, err = redpanda.NewAcl(ctx, "role_topic_read", &redpanda.AclArgs{
			ResourceType:        pulumi.String("TOPIC"),
			ResourceName:        testTopic.Name,
			ResourcePatternType: pulumi.String("LITERAL"),
			Principal: developer.Name.ApplyT(func(name string) (string, error) {
				return fmt.Sprintf("RedpandaRole:%v", name), nil
			}).(pulumi.StringOutput),
			Host:           pulumi.String("*"),
			Operation:      pulumi.String("READ"),
			PermissionType: pulumi.String("ALLOW"),
			ClusterApiUrl:  testCluster.ClusterApiUrl,
			AllowDeletion:  pulumi.Any(aclAllowDeletion),
		})
		if err != nil {
			return err
		}
		_, err = redpanda.NewRoleAssignment(ctx, "developer_assignment", &redpanda.RoleAssignmentArgs{
			RoleName: developer.Name,
			Principal: testUser.Name.ApplyT(func(name string) (string, error) {
				return fmt.Sprintf("User:%v", name), nil
			}).(pulumi.StringOutput),
			ClusterApiUrl: testCluster.ClusterApiUrl,
		}, pulumi.DependsOn([]pulumi.Resource{
			testUser,
			consoleCanary,
		}))
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
Example currently unavailable in this language
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
import com.pulumi.redpanda.inputs.ClusterMaintenanceWindowConfigArgs;
import com.pulumi.redpanda.inputs.ClusterMaintenanceWindowConfigDayHourArgs;
import com.pulumi.redpanda.inputs.ClusterTimeoutsArgs;
import com.pulumi.redpanda.Topic;
import com.pulumi.redpanda.TopicArgs;
import com.pulumi.redpanda.User;
import com.pulumi.redpanda.UserArgs;
import com.pulumi.redpanda.Role;
import com.pulumi.redpanda.RoleArgs;
import com.pulumi.redpanda.Acl;
import com.pulumi.redpanda.AclArgs;
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
                        jsonProperty("enable_shadow_linking", clusterEnableShadowLinking),
                        jsonProperty("schema_registry_enable_authorization", true)
                    )))
                .build())
            .maintenanceWindowConfig(ClusterMaintenanceWindowConfigArgs.builder()
                .dayHour(ClusterMaintenanceWindowConfigDayHourArgs.builder()
                    .dayOfWeek(maintenanceDayOfWeek)
                    .hourOfDay(maintenanceHourOfDay)
                    .build())
                .build())
            .tags(clusterTags)
            .timeouts(ClusterTimeoutsArgs.builder()
                .create("90m")
                .build())
            .build());

        var testTopic = new Topic("testTopic", TopicArgs.builder()
            .name(topicName)
            .partitionCount(partitionCount)
            .replicationFactor(replicationFactor)
            .clusterApiUrl(testCluster.clusterApiUrl())
            .allowDeletion(true)
            .configuration(topicConfiguration != null ? topicConfiguration : Map.ofEntries(
                Map.entry("cleanup.policy", "delete"),
                Map.entry("retention.ms", topicRetentionMs)
            ))
            .build());

        var testUser = new User("testUser", UserArgs.builder()
            .name(userName)
            .password(userPw)
            .mechanism(mechanism)
            .clusterApiUrl(testCluster.clusterApiUrl())
            .allowDeletion(userAllowDeletion)
            .build(), CustomResourceOptions.builder()
                .dependsOn(testTopic)
                .build());

        // Console-endpoint canary. The topic canary proves the dataplane; the console API
        // is a separate endpoint with separate readiness, so it needs its own. Role is the
        // simplest console resource, and the console resources below wait on it.
        var consoleCanary = new Role("consoleCanary", RoleArgs.builder()
            .name(String.format("%s-console-canary", roleName))
            .clusterApiUrl(testCluster.clusterApiUrl())
            .allowDeletion(true)
            .build());

        var developer = new Role("developer", RoleArgs.builder()
            .name(roleName)
            .clusterApiUrl(testCluster.clusterApiUrl())
            .allowDeletion(roleAllowDeletion)
            .build(), CustomResourceOptions.builder()
                .dependsOn(
                    testTopic,
                    consoleCanary)
                .build());

        // Bootstrap SR ACL grants for the provider's own Bearer-token principal.
        // SUBJECT grant authorizes POST /subjects/<subj>/versions; REGISTRY grant
        // authorizes the franz-go follow-up GET /schemas/ids/<id>/versions.
        var roleTopicRead = new Acl("roleTopicRead", AclArgs.builder()
            .resourceType("TOPIC")
            .resourceName(testTopic.name())
            .resourcePatternType("LITERAL")
            .principal(developer.name().applyValue(_name -> String.format("RedpandaRole:%s", _name)))
            .host("*")
            .operation("READ")
            .permissionType("ALLOW")
            .clusterApiUrl(testCluster.clusterApiUrl())
            .allowDeletion(aclAllowDeletion)
            .build());

        var developerAssignment = new RoleAssignment("developerAssignment", RoleAssignmentArgs.builder()
            .roleName(developer.name())
            .principal(testUser.name().applyValue(_name -> String.format("User:%s", _name)))
            .clusterApiUrl(testCluster.clusterApiUrl())
            .build(), CustomResourceOptions.builder()
                .dependsOn(
                    testUser,
                    consoleCanary)
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

	"github.com/pulumi/pulumi-pulumi-provider/sdks/go/redpanda/v2/redpanda"
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
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python

```
```python
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet

```
```csharp
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go

```
```go
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml

```
```yaml
Example currently unavailable in this language
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java

```
```java
Example currently unavailable in this language
```
{{% /choosable %}}
{{< /chooser >}}