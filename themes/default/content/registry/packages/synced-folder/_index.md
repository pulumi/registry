---
title: Synced Folder
meta_desc: An overview of the Synced Folder component.
layout: overview
---

A Pulumi component that synchronizes a local folder to Amazon S3, Azure Blob Storage, or Google Cloud Storage.

## Examples

{{< chooser language "typescript,python,go,csharp,yaml" / >}}

{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as synced from "@pulumi/synced-folder";

const bucket = new aws.s3.Bucket("s3-bucket", {
    acl: aws.s3.PublicReadAcl,
    website: {
        indexDocument: "index.html",
        errorDocument: "error.html",
    },
});

const folder = new synced.S3BucketFolder("synced-bucket-folder", {
    path: "./site",
    bucketName: bucket.bucket,
    acl: aws.s3.PublicReadAcl,
    managedObjects: false,
});

export const url = pulumi.interpolate`http://${bucket.websiteEndpoint}`;
```

{{% /choosable %}}

{{% choosable language python %}}

```python
"""An example of using the synced-folder component."""

import pulumi
from pulumi_aws import s3
import pulumi_synced_folder

bucket = s3.Bucket("s3-bucket", acl="public-read", 
    website=s3.BucketWebsiteArgs(
        index_document="index.html",
        error_document="error.html"
    )
)

folder = pulumi_synced_folder.S3BucketFolder("synced-bucket-folder", 
    path="./site",
    bucket_name=bucket.bucket,
    acl="public-read",
    managed_objects=False
)

pulumi.export("url", pulumi.Output.concat("http://", bucket.website_endpoint))
```

{{% /choosable %}}

{{% choosable language go %}}

```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/s3"
	synced "github.com/pulumi/pulumi-synced-folder/sdk/go/synced-folder"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		
		bucket, err := s3.NewBucket(ctx, "s3-bucket", &s3.BucketArgs{
			Acl: s3.CannedAclPublicRead,
			Website: s3.BucketWebsiteArgs{
				IndexDocument: pulumi.String("index.html"),
				ErrorDocument: pulumi.String("error.html"),
			},
		})
		if err != nil {
			return err
		}

		_, err = synced.NewS3BucketFolder(ctx, "synced-bucket-folder", &synced.S3BucketFolderArgs{
			Path: pulumi.String("./site"),
			Acl: s3.CannedAclPublicRead,
			BucketName: bucket.Bucket,
		})

		ctx.Export("url", pulumi.Sprintf("http://%s", bucket.WebsiteEndpoint))
		return nil
	})
}
```

{{% /choosable %}}

{{% choosable language csharp %}}

```csharp
using System.Collections.Generic;
using Pulumi;
using Pulumi.Aws.S3;
using Pulumi.Aws.S3.Inputs;
using Pulumi.SyncedFolder;

return await Deployment.RunAsync(() =>
{

    var bucket = new Bucket("s3-bucket", new BucketArgs {
        Acl = CannedAcl.PublicRead,
        Website = new BucketWebsiteArgs {
            IndexDocument = "index.html",
            ErrorDocument = "error.html",
        }
    });

    var folder = new S3BucketFolder("synced-bucket-folder", new S3BucketFolderArgs {
        Path = "./site",
        BucketName = bucket.BucketName,
        Acl = CannedAcl.PublicRead.ToString(),
        ManagedObjects = false,
    });

    return new Dictionary<string, object?>
    {
        ["url"] = Output.Format($"http://{bucket.WebsiteEndpoint}")
    };
});
```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
name: synced-folder-examples-aws-yaml
runtime: yaml
description: An example of using the synced-folder component.

resources:

  # Create an S3 bucket and configure it as a website.
  s3-bucket:
    type: aws:s3:Bucket
    properties:
      acl: public-read
      website:
        indexDocument: index.html
        errorDocument: error.html

  # Sync the contents of the ./site to the bucket
  synced-bucket-folder:
    type: synced-folder:index:S3BucketFolder
    properties:
      path: ./site
      bucketName: ${s3-bucket.bucket}
      acl: public-read

outputs:
  url: http://${s3-bucket.websiteEndpoint}
```

{{% /choosable %}}

## Notes

### Using the `managedObjects` property

By default, the component manages your files as individual Pulumi cloud resources, but you can opt out of this behavior by setting the component's `managedObjects` property to `false`. When you do this, the component assumes you've installed the appropriate CLI tool &mdash; [`aws`](https://aws.amazon.com/cli/), [`az`](https://docs.microsoft.com/en-us/cli/azure/), or [`gcloud`/`gsutil`](https://cloud.google.com/storage/docs/gsutil), depending on the cloud &mdash; and uses the [Command](https://www.pulumi.com/registry/packages/command/) provider to issue commands on that tool directly. Files are one-way synchronized only (local to remote), and files that exist remotely but not locally are deleted. All files are deleted from remote storage on `pulumi destroy`.

The component does not yet support switching seamlessly between `managedObjects: true` and `managedObjects: false`, however, so if you find after deploying a given folder with managed objects that you'd prefer to use unmanaged objects instead (or vice-versa), we recommend creating a second bucket/storage container and folder and removing the first. You can generally do this within the scope of a single program update. For example:

```yaml
# ...

resources:

  # The original bucket and synced-folder resources, using managed file objects.
  # 
  # my-first-bucket:
  #   type: aws:s3:Bucket
  #   properties:
  #     acl: public-read  
  #     website:
  #       indexDocument: index.html
  #       errorDocument: error.html
  #
  # my-first-synced-folder:
  #   type: synced-folder:index:S3BucketFolder
  #   properties:
  #     path: ./stuff
  #     bucketName: ${my-first-bucket.bucket}
  #     acl: public-read

  # A new bucket and synced-folder using unmanaged file objects.
  changed-my-mind-bucket:
    type: aws:s3:Bucket
    properties:
      acl: public-read
      website:
        indexDocument: index.html
        errorDocument: error.html

  changed-my-mind-synced-folder:
    type: synced-folder:index:S3BucketFolder
    properties:
      path: ./stuff
      bucketName: ${changed-my-mind-bucket.bucket}
      acl: public-read
      managedObjects: false 

outputs:

  # An updated program reference pointing to the new bucket.
  url: http://${changed-my-mind-bucket.websiteEndpoint}
```
