---
title: ElasticCloud (EC)
meta_desc: Provides an overview of the ElasticCloud Provider for Pulumi.
layout: overview
---

{{% overview-description %}}
The ElasticCloud provider for Pulumi can be used to provision any of the resources available in [ElasticCloud](https://www.elastic.co/cloud/).
{{% /overview-description %}}

## Example

{{< chooser language "javascript,typescript,python,go,csharp" >}}

{{% choosable language javascript %}}

```javascript
const ec = require("@pulumi/ec")

const REGION = "us-east-1";

const latestVersion = elasticCloud.getStack({
    region: REGION,
    versionRegex: "latest"
});

new elasticCloud.Deployment('my-deployment', {
    region: REGION,
    version: latestVersion.then((x: { version: string; }) => x.version),
    deploymentTemplateId: "aws-io-optimized-v2",
    elasticsearch: {}
});
```

{{% /choosable %}}
{{% choosable language typescript %}}

```typescript
import * as elasticCloud from "@pulumi/ec";

const REGION = "us-east-1";

const latestVersion = elasticCloud.getStack({
    region: REGION,
    versionRegex: "latest"
});

new elasticCloud.Deployment('my-deployment', {
    region: REGION,
    version: latestVersion.then((x: { version: string; }) => x.version),
    deploymentTemplateId: "aws-io-optimized-v2",
    elasticsearch: {}
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_ec as ec

REGION = "us-east-1"

latest_version = ec.get_stack(
    region=REGION, version_regex="latest").version

deployment = ec.Deployment('my-deployment', region=REGION, version=latest_version,
                           deployment_template_id="aws-io-optimized-v2", elasticsearch={})
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	ec "github.com/pulumi/pulumi-ec/sdk/go/ec"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
        version, err := ec.GetStack(ctx, &GetStackArgs{
			Region:       pulumi.String("us-east-1"),
			VersionRegex: "latest",
        }
		if err != nil {
			return err
		}
		
		_, err := ec.NewDeployment(ctx, "my-deployment", &ec.DeploymentArgs{
            Region:               pulumi.String("us-east-1"),
			Version:              version.Version,
            DeploymentTemplateId: pulumi.String("aws-io-optimized-v2")
			ElasticSearch:        &ec.DeploymentElasticsearch{},
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
using ElasticCloud = Pulumi.ElasticCloud;

class Program
{
    static Task Main() =>
        Deployment.Run(() => {
            var latest = Output.Create(ElasticCloud.GetStack.InvokeAsync(new ElasticCloud.GetStackArgs
            {
                Region = "us-east-1",
                VersionRegex = "latest",
            });

            var deployment = ElasticCloud.Deployment("my-deployment", new ElasticCloud.DeploymentArgs
            {
                Region = "us-east-1",
                Version = latest.Version,
                DeploymentTemplateId = "aws-io-optimized-v2",
                ElasticSearch = {},
            });
        });
}
```

{{% /choosable %}}

{{< /chooser >}}
