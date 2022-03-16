---
title: Kafka
meta_desc: Provides an overview of the Kafka Provider for Pulumi.
layout: overview
---

{{% overview-description %}}
The Kafka provider for Pulumi can be used to provision any of the resources available for Kafka.
The Kafka provider must be configured with credentials to deploy and update resources in Kafka.
{{% /overview-description %}}

## Example

{{< chooser language "javascript,typescript,python,go,csharp" >}}

{{% choosable language javascript %}}

```javascript
const kafka = require("@pulumi/kafka")

const topic = new kafka.Topic("topic", {
  name: "sample-topic",
  replicationFactor: 1,
  partitions: 4,
});
```

{{% /choosable %}}
{{% choosable language typescript %}}

```typescript
import * as kafka from "@pulumi/kafka";

const topic = new kafka.Topic("topic", {
  name: "sample-topic",
  replicationFactor: 1,
  partitions: 4,
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_kafka as kafka

topic = kafka.Topic("topic",
  name="sample-topic",
  replication_factor=1,
  partitions=4,
)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	kafka "github.com/pulumi/pulumi-kafka/sdk/v3/go/kafka"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		topic, err := kafka.NewTopic(ctx, "topic", &kafka.TopicArgs{
			Name:              pulumi.String("sample-topic"),
			ReplicationFactor: pulumi.Int(1),
			Partitions:        pulumi.Int(4),
		})
		if err != nil {
			return err
		}

		return nil
	})
}

```

{{% /choosable %}}
{{% choosable language csharp %}}

```csharp
using System.Collections.Generic;
using System.Threading.Tasks;
using Pulumi;
using Pulumi.Kafka;

class Program
{
    static Task Main() =>
        Deployment.Run(() => {
            var topic = new Topic("topic", new TopicArgs
            {
                Name = "sample-topic",
                ReplicationFactor = 1,
                Partitions = 4,
            });
        });
}
```

{{% /choosable %}}

{{< /chooser >}}
