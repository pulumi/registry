---
title: Google Cloud (GCP) Classic
meta_desc: Learn how to use Pulumi's GCP Provider to reduce the complexity of managing and provisioning GCP resources.
layout: overview
---

The Google Cloud Platform (GCP) provider for Pulumi can provision many of the cloud resources available in [Google Cloud](https://cloud.google.com/).

The GCP provider must be configured with credentials to deploy and update resources in Google Cloud; see the [Installation & Configuration]({{<relref "./installation-configuration">}}) for instructions.

**New to Pulumi and Google Cloud?** [Get started with Google Cloud using our tutorial]({{<relref "/docs/get-started/gcp">}})

{{% notes %}}
Pulumi has a new Google Cloud provider: the [Pulumi Google Native Provider]({{<relref "/registry/packages/google-native">}}). Google Native gives you same-day access to all new Google Cloud resources.

Consider trying [Google Native]({{<relref "/registry/packages/google-native">}}) if you need Google Cloud resources that aren't available in this provider.
{{% /notes %}}

## Example

{{< chooser language "javascript,typescript,python,go,csharp" >}}

{{% choosable language javascript %}}

```javascript
const gcp = require("@pulumi/gcp")

const bucket = new gcp.storage.Bucket("my-bucket");
```

{{% /choosable %}}
{{% choosable language typescript %}}

```typescript
import * as gcp from "@pulumi/gcp";

const bucket = new gcp.storage.Bucket("my-bucket");
```

{{% /choosable %}}
{{% choosable language python %}}

```python
from pulumi_gcp import storage

bucket = storage.Bucket('my-bucket')
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
    "github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/storage"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		bucket, err := storage.NewBucket(ctx, "my-bucket", nil)
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
using System.Threading.Tasks;
using Pulumi;
using Pulumi.Gcp;

class Program
{
    static Task Main() =>
        Deployment.Run(() => {
            var bucket = new Gcp.Storage.Bucket("my-bucket");
        });
}
```

{{% /choosable %}}

{{< /chooser >}}

Visit the [How-to Guides]({{<relref "./how-to-guides">}}) to find step-by-step guides for specific scenarios like creating a serverless application using Google Cloud Functions or setting up a Google Kubernetes Engine (GKE) cluster.

## SDK packages

The Google Cloud (GCP) Classic provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/gcp`](https://www.npmjs.com/package/@pulumi/gcp)
* Python: [`pulumi-gcp`](https://pypi.org/project/pulumi-gcp/)
* Go: [`github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp`](https://github.com/pulumi/pulumi-gcp)
* .NET: [`Pulumi.Gcp`](https://www.nuget.org/packages/Pulumi.Gcp)
