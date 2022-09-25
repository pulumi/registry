---
title: Synced Folder
meta_desc: A Pulumi component that synchronizes a local folder to Amazon S3, Azure Blob Storage, or Google Cloud Storage.
layout: overview
---

A Pulumi component that synchronizes the contents of a local folder to Amazon S3, Azure Blob Storage, or Google Cloud Storage.

## Examples

The following examples show you how to use the component, or something.

### Sync to Amazon S3

{{< chooser language "typescript,python,go,csharp,yaml" / >}}

{{% choosable language typescript %}}

```typescript
import * as aws from "@pulumi/aws";
import * as synced from "@pulumi/synced-folder";

const bucket = new aws.s3.Bucket("my-bucket", {
    acl: aws.s3.PublicReadAcl,
});

const folder = new synced.S3BucketFolder("synced-folder", {
    path: "./my-folder",
    bucketName: bucket.bucket,
    acl: aws.s3.PublicReadAcl,
});
```

{{% /choosable %}}

{{% choosable language python %}}

```python
from pulumi_aws import s3
import pulumi_synced_folder

bucket = s3.Bucket(
    "my-bucket",
    acl=s3.CannedAcl.PUBLIC_READ,
)

folder = pulumi_synced_folder.S3BucketFolder(
    "synced-folder",
    path="./my-folder",
    bucket_name=bucket.bucket,
    acl=s3.CannedAcl.PUBLIC_READ,
)
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

		bucket, err := s3.NewBucket(ctx, "my-bucket", &s3.BucketArgs{
			Acl: s3.CannedAclPublicRead,
		})
		if err != nil {
			return err
		}

		_, err = synced.NewS3BucketFolder(ctx, "synced-folder", &synced.S3BucketFolderArgs{
			Path:       pulumi.String("./my-folder"),
			BucketName: bucket.Bucket,
			Acl:        s3.CannedAclPublicRead,
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
using Pulumi;
using Pulumi.Aws.S3;
using Pulumi.SyncedFolder;

return await Deployment.RunAsync(() =>
{

    var bucket = new Bucket("my-bucket", new BucketArgs {
        Acl = CannedAcl.PublicRead,
    });

    var folder = new S3BucketFolder("synced-folder", new S3BucketFolderArgs {
        Path = "./my-folder",
        BucketName = bucket.BucketName,
        Acl = (string)CannedAcl.PublicRead,
    });
});
```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
name: synced-folder-aws-yaml
runtime: yaml

resources:

  my-bucket:
    type: aws:s3:Bucket
    properties:
      acl: public-read

  synced-folder:
    type: synced-folder:index:S3BucketFolder
    properties:
      path: ./my-folder
      bucketName: ${my-bucket.bucket}
      acl: public-read
```

{{% /choosable %}}

### Sync to Azure Blob Storage

{{< chooser language "typescript,python,go,csharp,yaml" / >}}

{{% choosable language typescript %}}

```typescript
import * as resources from "@pulumi/azure-native/resources";
import * as storage from "@pulumi/azure-native/storage";
import * as synced from "@pulumi/synced-folder";

const resourceGroup = new resources.ResourceGroup("resourceGroup");

const account = new storage.StorageAccount("account", {
    resourceGroupName: resourceGroup.name,
    kind: storage.Kind.StorageV2,
    sku: {
        name: storage.SkuName.Standard_LRS,
    },
});

const container = new storage.BlobContainer("container", {
    resourceGroupName: resourceGroup.name,
    accountName: account.name,
});

const folder = new synced.AzureBlobFolder("synced-folder", {
    resourceGroupName: resourceGroup.name,
    storageAccountName: account.name,
    containerName: container.name,
    path: "./my-folder",
});
```

{{% /choosable %}}

{{% choosable language python %}}

```python
import resource
from pulumi_azure_native import storage
from pulumi_azure_native import resources
import pulumi_synced_folder

resource_group = resources.ResourceGroup("resource_group")

account = storage.StorageAccount("account", storage.StorageAccountArgs(
    resource_group_name=resource_group.name,
    kind=storage.Kind.STORAGE_V2,
    sku=storage.SkuArgs(
        name=storage.SkuName.STANDARD_LRS,
    )
))

container = storage.BlobContainer("container", storage.BlobContainerArgs(
    resource_group_name=resource_group.name,
    account_name=account.name,
))

folder = pulumi_synced_folder.AzureBlobFolder("synced-folder", pulumi_synced_folder.AzureBlobFolderArgs(
    resource_group_name=resource_group.name,
    storage_account_name=account.name,
    container_name=container.name,
    path="./my-folder",
))
```

{{% /choosable %}}

{{% choosable language go %}}

```go
package main

import (
	"github.com/pulumi/pulumi-azure-native/sdk/go/azure/resources"
	"github.com/pulumi/pulumi-azure-native/sdk/go/azure/storage"
	synced "github.com/pulumi/pulumi-synced-folder/sdk/go/synced-folder"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		resourceGroup, err := resources.NewResourceGroup(ctx, "resourceGroup", nil)
		if err != nil {
			return err
		}

		account, err := storage.NewStorageAccount(ctx, "account", &storage.StorageAccountArgs{
			ResourceGroupName: resourceGroup.Name,
			Kind:              pulumi.String("StorageV2"),
			Sku: &storage.SkuArgs{
				Name: pulumi.String("Standard_LRS"),
			},
		})
		if err != nil {
			return err
		}

		container, err := storage.NewBlobContainer(ctx, "container", &storage.BlobContainerArgs{
			ResourceGroupName: resourceGroup.Name,
			AccountName:       account.Name,
		})
		if err != nil {
			return err
		}

		_, err = synced.NewAzureBlobFolder(ctx, "folder", &synced.AzureBlobFolderArgs{
			ResourceGroupName:  resourceGroup.Name,
			StorageAccountName: account.Name,
			ContainerName:      container.Name,
			Path:               pulumi.String("./my-folder"),
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
using Pulumi.AzureNative.Resources;
using Pulumi.AzureNative.Storage;
using Pulumi.AzureNative.Storage.Inputs;
using Pulumi.SyncedFolder;

return await Pulumi.Deployment.RunAsync(() =>
{
    var resourceGroup = new ResourceGroup("resource-group");

    var storageAccount = new StorageAccount("account", new StorageAccountArgs
    {
        ResourceGroupName = resourceGroup.Name,
        Kind = Kind.StorageV2,
        Sku = new SkuArgs
        {
            Name = SkuName.Standard_LRS
        },
    });

    var container = new BlobContainer("container", new BlobContainerArgs
    {
        ResourceGroupName = resourceGroup.Name,
        AccountName = storageAccount.Name,
    });

    var folder = new AzureBlobFolder("synced-folder", new AzureBlobFolderArgs
    {
        ResourceGroupName = resourceGroup.Name,
        StorageAccountName = storageAccount.Name,
        ContainerName = container.Name,
        Path = "./my-folder",
    });
});
```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
name: synced-folder-azure-yaml
runtime: yaml

resources:

    resourceGroup:
      type: azure-native:resources:ResourceGroup

    account:
      type: azure-native:storage:StorageAccount
      properties:
        resourceGroupName: ${resourceGroup.name}
        kind: StorageV2
        sku:
          name: Standard_LRS

    container:
      type: azure-native:storage:BlobContainer
      properties:
        resourceGroupName: ${resourceGroup.name}
        accountName: ${account.name}

    folder:
      type: synced-folder:index:AzureBlobFolder
      properties:
        resourceGroupName: ${resourceGroup.name}
        storageAccountName: ${account.name}
        containerName: ${container.name}
        path: ./my-folder
```

{{% /choosable %}}
### Sync to Google Cloud Storage

{{< chooser language "typescript,python,go,csharp,yaml" / >}}

{{% choosable language typescript %}}

```typescript
import * as gcp from "@pulumi/gcp";
import * as synced from "@pulumi/synced-folder";

const bucket = new gcp.storage.Bucket("my-bucket", {
    location: "US"
});

const binding = new gcp.storage.BucketIAMBinding("binding", {
    bucket: bucket.name,
    role: "roles/storage.objectViewer",
    members: [
        "allUsers"
    ],
});

const folder = new synced.GoogleCloudFolder("folder", {
    bucketName: bucket.name,
    path: "./my-folder",
});
```

{{% /choosable %}}

{{% choosable language python %}}

```python
from pulumi_gcp import storage
import pulumi_synced_folder

bucket = storage.Bucket("my-bucket", location="US")

binding = storage.BucketIAMBinding(
    "binding",
    storage.BucketIAMBindingArgs(
        bucket=bucket.name,
        role="roles/storage.objectViewer",
        members=[
            "allUsers",
        ],
    )
)

folder = pulumi_synced_folder.GoogleCloudFolder(
    "folder",
    pulumi_synced_folder.GoogleCloudFolderArgs(
        bucket_name=bucket.name,
        path="./my-folder",
    ),
)
```

{{% /choosable %}}

{{% choosable language go %}}

```go
package main

import (
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/storage"
	synced "github.com/pulumi/pulumi-synced-folder/sdk/go/synced-folder"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		bucket, err := storage.NewBucket(ctx, "my-bucket", &storage.BucketArgs{
			Location: pulumi.String("US"),
		})
		if err != nil {
			return err
		}

		_, err = storage.NewBucketIAMBinding(ctx, "binding", &storage.BucketIAMBindingArgs{
			Bucket:  bucket.Name,
			Role:    pulumi.String("roles/storage.objectViewer"),
			Members: pulumi.ToStringArray([]string{"allUsers"}),
		})
		if err != nil {
			return err
		}

		_, err = synced.NewGoogleCloudFolder(ctx, "folder", &synced.GoogleCloudFolderArgs{
			BucketName: bucket.Name,
			Path:       pulumi.String("./my-folder"),
		})
		if err != nil {
			return err
		}

		return nil
	})
}
```

{{% /choosable %}}

{{% choosable %}}

```csharp
using Pulumi;
using Pulumi.Gcp.Storage;
using Pulumi.SyncedFolder;

return await Deployment.RunAsync(() =>
{
    var bucket = new Bucket("my-bucket", new BucketArgs
    {
        Location = "US"
    });

    var binding = new BucketIAMBinding("binding", new BucketIAMBindingArgs
    {
        Bucket = bucket.Name,
        Role = "roles/storage.objectViewer",
        Members = new[]
        {
            "allUsers",
        }
    });

    var folder = new GoogleCloudFolder("folder", new GoogleCloudFolderArgs
    {
        BucketName = bucket.Name,
        Path = "./my-folder",
    });
});
```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
name: synced-folder-gcp-yaml
runtime: yaml

resources:

  bucket:
    type: gcp:storage:Bucket
    properties:
      location: US

  binding:
    type: gcp:storage:BucketIAMBinding
    properties:
      bucket: ${bucket.name}
      role: roles/storage.objectViewer
      members:
        - allUsers

  folder:
    type: synced-folder:index:GoogleCloudFolder
    properties:
      bucketName: ${bucket.name}
      path: ./my-folder
```

{{% /choosable %}}

## Notes

### Managed and unmanaged file objects

By default, the Synced Folder component manages your files as individual Pulumi cloud resources (e.g., as `aws:S3:BucketObject`s), but you can opt out of this behavior by setting the component's `managedObjects` property to `false`:

{{< chooser language "typescript,python,go,csharp,yaml" / >}}

{{% choosable language typescript %}}

```typescript
const folder = new synced.S3BucketFolder("synced-folder", {
    path: "./my-folder",
    bucketName: bucket.bucket,
    acl: aws.s3.PublicReadAcl,

    // Set this property to false to manage files outside Pulumi.
    managedObjects: false,
});
```

{{% /choosable %}}

{{% choosable language python %}}

```python
folder = pulumi_synced_folder.S3BucketFolder(
    "synced-folder",
    path="./my-folder",
    bucket_name=bucket.bucket,
    acl=s3.CannedAcl.PUBLIC_READ,

    # Set this property to false to manage files outside Pulumi.
    managed_objects=False,
)
```

{{% /choosable %}}

{{% choosable language go %}}

```go
folder, err = synced.NewS3BucketFolder(ctx, "synced-folder", &synced.S3BucketFolderArgs{
    Path:           pulumi.String("./my-folder"),
    BucketName:     bucket.Bucket,
    Acl:            s3.CannedAclPublicRead,

    // Set this property to false to manage files outside Pulumi.
    ManagedObjects: pulumi.Bool(false),
})
```

{{% /choosable %}}

{{% choosable language csharp %}}

```csharp
var folder = new S3BucketFolder("synced-bucket-folder", new S3BucketFolderArgs {
    Path = "./my-folder",
    BucketName = bucket.BucketName,
    Acl = (string)CannedAcl.PublicRead,

    // Set this property to false to manage files outside Pulumi.
    ManagedObjects = false,
});
```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
folder:
  type: synced-folder:index:S3BucketFolder
  properties:
    path: ./my-folder
    bucketName: ${my-bucket.bucket}

    # Set this property to false to manage files outside Pulumi.
    managedObjects: false
```

{{% /choosable %}}

When you do this, the component assumes you've installed the appropriate CLI tool &mdash; [`aws`](https://aws.amazon.com/cli/), [`az`](https://docs.microsoft.com/en-us/cli/azure/), or [`gcloud`/`gsutil`](https://cloud.google.com/storage/docs/gsutil), depending on the cloud &mdash; and uses the [Pulumi Command](https://www.pulumi.com/registry/packages/command/) provider to issue commands the CLI tool directly.

| Resource | By default |  `managedObjects: false` |
| -- | -- | -- |
| `S3BucketFolder` | `aws:s3:BucketObject` | [`aws s3 sync`](https://docs.aws.amazon.com/cli/latest/reference/s3/sync.html) |
| `AzureBlobFolder` | `azure-native:storage:Blob` | [`az storage blob sync`](https://learn.microsoft.com/en-us/cli/azure/storage/blob?view=azure-cli-latest#az-storage-blob-sync) |
| `GoogleCloudFolder` | `gcp:storage:BucketObject` | [`gsutil rsync`](https://cloud.google.com/storage/docs/gsutil/commands/rsync) |


Files are one-way synchronized to the cloud, and any files that exist remotely but not locally are deleted. All files are deleted from remote storage on `pulumi destroy`. If the folder you're syncing contains thousands of files, you may want to consider using this option.

The component does not yet support switching seamlessly between `managedObjects: true` and `managedObjects: false`, however, so if you find after deploying a given folder with managed objects that you'd prefer to use unmanaged objects instead (or vice-versa), we recommend creating a second bucket/storage container and folder and removing the first. You can generally do this within the scope of a single program update. For example:

```yaml
name: synced-folder-aws-yaml
runtime: yaml

resources:

  # Original bucket and synced-folder resources, using managed file objects.
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

  changed-my-mind-synced-folder:
    type: synced-folder:index:S3BucketFolder
    properties:
      path: ./stuff
      bucketName: ${changed-my-mind-bucket.bucket}
      acl: public-read
      managedObjects: false
```
