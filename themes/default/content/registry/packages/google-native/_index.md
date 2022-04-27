---
title: Google Native
meta_desc: Learn how to use Pulumi's Google Native Provider to reduce the complexity of managing and provisioning GCP resources.
layout: overview
---

{{< notes >}}
Google Native is in public preview.

[Google Cloud Classic]({{<relref "/registry/packages/gcp">}}) remains fully supported and recommended for production use.
{{< /notes >}}

The Google Native provider for Pulumi can provision many of the cloud resources available in [Google Cloud](https://cloud.google.com/).

The Google Native provider must be configured with credentials to deploy and update resources in Google Cloud; see [Installation & Configuration]({{<relref "./installation-configuration">}}) for instructions.

## Example

{{< chooser language "typescript,python,go,csharp,yaml" >}}

{{% choosable language typescript %}}

```typescript
import * as storage from "@pulumi/google-native/storage/v1";
import * as pulumi from "@pulumi/pulumi";

const config = new pulumi.Config("google-native");
const project = config.require("project");
const bucketName = "pulumi-goog-native-ts-01";

// Create a Google Cloud resource (Storage Bucket)
const bucket = new storage.Bucket("my-bucket", {
    name:bucketName,
    bucket:bucketName,
    project: project,
});

// Export the bucket self-link
export const bucketName = bucket.selfLink;
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi
from pulumi_google_native.storage import v1 as storage

config = pulumi.Config()
project = config.require('project')
# Create a Google Cloud resource (Storage Bucket)
bucket_name = "pulumi-goog-native-bucket-py-01"
bucket = storage.Bucket('my-bucket', name=bucket_name, bucket=bucket_name, project=project)

# Export the bucket self-link
pulumi.export('bucket', bucket.self_link)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
package main

import (
	storage "github.com/pulumi/pulumi-google-native/sdk/go/google/storage/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	const bucketName = "pulumi-goog-native-bucket-go-01"
	pulumi.Run(func(ctx *pulumi.Context) error {
		conf := config.New(ctx, "google-native")
		project := conf.Require("project")
		// Create a Google Cloud resource (Storage Bucket)
		bucket, err := storage.NewBucket(ctx, "bucket", &storage.BucketArgs{
			Name:    pulumi.StringPtr(bucketName),
			Bucket:  pulumi.String(project),
			Project: project,
		})
		if err != nil {
			return err
		}
		// Export the bucket self-link
		ctx.Export("bucket", bucket.SelfLink)

		return nil
	})
}
```

{{% /choosable %}}
{{% choosable language csharp %}}

```csharp
using System.Threading.Tasks;
using Pulumi;
using Pulumi.GoogleNative.Storage.V1;

class Program
{
    static Task Main() =>
        Deployment.Run(() => {
            var config = new Config("google-native");
            var project = config.Require("project");
            var bucketName = "pulumi-goog-native-bucket-cs-01";
            // Create a Google Cloud resource (Storage Bucket)
            var bucket = new Bucket("my-bucket", new BucketArgs{
                Name = bucketName,
                Bucket = bucketName,
                Project = project,
            });
        });
}
```

{{% /choosable %}}
{{% choosable language yaml %}}

```yaml
configuration:
  project:
    type: String
variables:
  bucketName: pulumi-goog-native-bucket-go-01
resources:
  my-bucket:
    type: google-native:storage/v1:Bucket
    properties:
      name: ${bucketName}
      bucket: ${bucketName}
      project: ${project}
outputs:
  bucket: ${my-bucket.selfLink}
```

{{% /choosable %}}

{{< /chooser >}}

Visit the [How-to Guides]({{<relref "./how-to-guides">}}) to find step-by-step guides for specific scenarios like creating a serverless application using Google Cloud Functions or setting up a Google Kubernetes Engine (GKE) cluster.
