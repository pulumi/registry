---
# WARNING: this file was fetched from https://raw.githubusercontent.com/azaurus1/pulumi-kafka-connect/v0.0.12/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
title: Kafka Connect
meta_desc: Provides an overview of the Kafka Connect Provider for Pulumi.
layout: package
---

This Pulumi provider enables you to manage you [Kafka Connect](https://www.redpanda.com/guides/kafka-tutorial-what-is-kafka-connect) connectors using any language of Pulumi IaC.

## Example
{{< chooser language "typescript,python,go,csharp,yaml" >}}

{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as kafkaconnect from "@azaurus/kafkaconnect";

const defaultProvider = new kafkaconnect.Provider("defaultProvider", {url: "http://localhost:8083"});
const mirrorHeartbeatConnectorYxdw = new kafkaconnect.connector.Connector("mirror-heartbeat-connector-yxdw", {config: {
    "connector.class": "org.apache.kafka.connect.mirror.MirrorHeartbeatConnector",
    "source.cluster.alias": "source",
    "heartbeats.topic.replication.factor": "-1",
    name: "mirror-heartbeat-connector-yxdw",
}}, {
    provider: defaultProvider,
});
export const output = {
    value: mirrorHeartbeatConnectorYxdw.result,
};
```
{{% /choosable %}}

{{% choosable language python %}}
```python
import pulumi
import pulumi-kafkaconnect as kafkaconnect

default_provider = kafkaconnect.Provider("defaultProvider", url="http://localhost:8083")
mirror_heartbeat_connector_yxdw = kafkaconnect.connector.Connector("mirror-heartbeat-connector-yxdw", config={
    "connector.class": "org.apache.kafka.connect.mirror.MirrorHeartbeatConnector",
    "source.cluster.alias": "source",
    "heartbeats.topic.replication.factor": "-1",
    "name": "mirror-heartbeat-connector-yxdw",
},
opts = pulumi.ResourceOptions(provider=default_provider))
pulumi.export("output", {
    "value": mirror_heartbeat_connector_yxdw.result,
})
```
{{% /choosable %}}

{{% choosable language go %}}
```go
package main

import (
	"github.com/azaurus1/pulumi-kafka-connect/sdk/go/kafkaconnect"
	"github.com/azaurus1/pulumi-kafka-connect/sdk/go/kafkaconnect/connector"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		defaultProvider, err := kafkaconnect.NewProvider(ctx, "defaultProvider", &kafkaconnect.ProviderArgs{
			Url: pulumi.String("http://localhost:8083"),
		})
		if err != nil {
			return err
		}
		mirrorHeartbeatConnectorYxdw, err := connector.NewConnector(ctx, "mirror-heartbeat-connector-yxdw", &connector.ConnectorArgs{
			Config: pulumi.Map{
				"connector.class":                     pulumi.Any("org.apache.kafka.connect.mirror.MirrorHeartbeatConnector"),
				"source.cluster.alias":                pulumi.Any("source"),
				"heartbeats.topic.replication.factor": pulumi.Any("-1"),
				"name":                                pulumi.Any("mirror-heartbeat-connector-yxdw"),
			},
		}, pulumi.Provider(defaultProvider))
		if err != nil {
			return err
		}
		ctx.Export("output", pulumi.StringMap{
			"value": mirrorHeartbeatConnectorYxdw.Result,
		})
		return nil
	})
}

```
{{% /choosable %}}

{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Kafkaconnect = Azaurus1.Kafkaconnect;

return await Deployment.RunAsync(() =>
{
    var defaultProvider = new Kafkaconnect.Provider("defaultProvider", new()
    {
        Url = "http://localhost:8083",
    });

    var mirrorHeartbeatConnectorYxdw = new Kafkaconnect.Connector.Connector("mirror-heartbeat-connector-yxdw", new()
    {
        Config =
        {
            { "connector.class", "org.apache.kafka.connect.mirror.MirrorHeartbeatConnector" },
            { "source.cluster.alias", "source" },
            { "heartbeats.topic.replication.factor", "-1" },
            { "name", "mirror-heartbeat-connector-yxdw" },
        },
    }, new CustomResourceOptions
    {
        Provider = defaultProvider,
    });

    return new Dictionary<string, object?>
    {
        ["output"] =
        {
            { "value", mirrorHeartbeatConnectorYxdw.Result },
        },
    };
});
```
{{% /choosable %}}

{{% choosable language yaml %}}
```yaml
name: provider-kafkaconnect-native
runtime: yaml

resources:
  # Explicitly configure the default provider with the required 'url' property
  defaultProvider:
    type: pulumi:providers:kafkaconnect
    properties:
      url: "http://localhost:8083" # Replace with your actual Kafka Connect URL

  mirror-heartbeat-connector-yxdw:
    type: kafkaconnect:connector:Connector
    options:
      provider: ${defaultProvider}
    properties:
      config:
        connector.class: "org.apache.kafka.connect.mirror.MirrorHeartbeatConnector"
        source.cluster.alias: "source"
        heartbeats.topic.replication.factor: "-1"
        name: "mirror-heartbeat-connector-yxdw"

outputs:
  output:
    value: ${mirror-heartbeat-connector-yxdw.result}
```
{{% /choosable %}}

{{< /chooser >}}
