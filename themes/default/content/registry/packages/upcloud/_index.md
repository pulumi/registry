---
# WARNING: this file was fetched from https://raw.githubusercontent.com/UpCloudLtd/pulumi-upcloud/v0.10.2/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/UpCloudLtd/pulumi-upcloud/blob/v0.10.2/docs/_index.md
title: UpCloud
meta_desc: Provides an overview of the UpCloud Pulumi Provider.
layout: package
---

Use the UpCloud Pulumi provider to deploy and manage resources in [UpCloud](https://upcloud.com/). Get started by installing the provider and configuring your credentials.

## Example usage

<!-- These are copied from examples/templates -->

{{< chooser language "go,typescript,python,csharp" >}}
{{% choosable language go %}}

```golang
package main

import (
	"github.com/UpCloudLtd/pulumi-upcloud/sdk/go/upcloud"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Load configuration from Pulumi config
		cfg := config.New(ctx, "")

		objectStorageName := cfg.Get("object_storage_name")
		if objectStorageName == "" {
			objectStorageName = "bucket-example-objstov2"
		}

		region := cfg.Get("region")
		if region == "" {
			region = "europe-1"
		}

		bucketName := cfg.Get("bucket_name")
		if bucketName == "" {
			bucketName = "bucket"
		}

		// Create an UpCloud Managed Object Storage
		objectStorage, err := upcloud.NewManagedObjectStorage(ctx, "objectStorage", &upcloud.ManagedObjectStorageArgs{
			Name:             pulumi.String(objectStorageName),
			Region:           pulumi.String(region),
			ConfiguredStatus: pulumi.String("started"),
		})
		if err != nil {
			return err
		}

		// Create a Bucket inside the Object Storage
		bucket, err := upcloud.NewManagedObjectStorageBucket(ctx, "storageBucket", &upcloud.ManagedObjectStorageBucketArgs{
			ServiceUuid: objectStorage.ID(),
			Name:        pulumi.String(bucketName),
		})
		if err != nil {
			return err
		}

		// Export outputs
		ctx.Export("object_storage_uuid", objectStorage.ID())
		ctx.Export("bucket_name", bucket.Name)

		return nil
	})
}
```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi
import pulumi_upcloud as upcloud

# Load user input from Pulumi config
config = pulumi.Config()
object_storage_name = config.require("object_storage_name")
region = config.require("region")
bucket_name = config.require("bucket_name")

# Create an UpCloud Managed Object Storage
object_storage = upcloud.ManagedObjectStorage(
    "objectStorage",
    name=object_storage_name,
    region=region,
    configured_status="started",
)

# Create a Bucket inside the Object Storage
bucket = upcloud.ManagedObjectStorageBucket(
    "storageBucket",
    service_uuid=object_storage.id,
    name=bucket_name,
)

# Export outputs
pulumi.export("object_storage_uuid", object_storage.id)
pulumi.export("bucket_name", bucket.name)
```

{{% /choosable %}}

{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as upcloud from "@upcloud/pulumi-upcloud";

// Load Pulumi config values
const config = new pulumi.Config();

const objectStorageName = config.require("object_storage_name");
const region = config.require("region");
const bucketName = config.require("bucket_name");

// Create an UpCloud Managed Object Storage
const objectStorage = new upcloud.ManagedObjectStorage("objectStorage", {
    name: objectStorageName,
    region: region,
    configuredStatus: "started"
});

// Create a Bucket inside the Object Storage
const bucket = new upcloud.ManagedObjectStorageBucket("storageBucket", {
    serviceUuid: objectStorage.id,
    name: bucketName
});

// Export outputs
export const objectStorageUuid = objectStorage.id;
export const bucketNameOutput = bucket.name;
```

{{% /choosable %}}

{{% choosable language csharp %}}

```csharp
using Pulumi;
using UpCloud.Pulumi.UpCloud;
using System.Collections.Generic;

return await Deployment.RunAsync(() =>
{
    var config = new Pulumi.Config();

    var objectStorageName = config.Require("object_storage_name");
    var region = config.Require("region");
    var bucketName = config.Require("bucket_name");

    var objectStorage = new ManagedObjectStorage("objectStorage", new()
    {
        Name = objectStorageName,
        Region = region,
        ConfiguredStatus = "started"
    });

    var bucket = new ManagedObjectStorageBucket("storageBucket", new()
    {
        ServiceUuid = objectStorage.Id,
        Name = bucketName
    });

    return new Dictionary<string, object?>
    {
        ["object_storage_uuid"] = objectStorage.Id,
        ["bucket_name"] = bucket.Name
    };
});
```

{{% /choosable %}}

{{< /chooser >}}
