---
title: Synced Folder Installation
meta_desc: Instructions for installing and configuring the Synced Folder component.
layout: installation
---

## Installing the package

The Synced Folder component package is available in the following Pulumi languages:

* TypeScript/JavaScript: [`@pulumi/synced-folder`](https://www.npmjs.com/package/@pulumi/synced-folder)
* Python: [`pulumi_synced_folder`](https://pypi.org/project/pulumi-synced-folder/)
* Go: [`github.com/pulumi/pulumi-synced-folder/sdk/go/synced-folder`](https://github.com/pulumi/pulumi-synced-folder/)
* C#/.NET: [`Pulumi.SyncedFolder`](https://www.nuget.org/packages/Pulumi.SyncedFolder)
* Pulumi YAML

You can install the package with your package manager of choice:


```bash
$ npm install @pulumi/synced-folder
```

```bash
$ yarn add @pulumi/synced-folder
```

```bash
$ pip install pulumi_synced_folder
```

```bash
$ go get -u github.com/pulumi/pulumi-synced-folder/sdk/go/synced-folder
```

```bash
$ dotnet add package Pulumi.SyncedFolder
```

## Configuring

The package contains three components:

* `S3BucketFolder` syncs the contents of a local folder to a specified Amazon S3 bucket.
* `AzureBlobFolder` syncs to a specified Azure Blob Storage container.
* `GoogleCloudFolder` syncs to a specified Google Cloud Storage bucket.

Because each of these components is backed by a different cloud, they're each configured a little differently.

### Configuring cloud-provider settings

By default, the component inherits the cloud-provider settings of the project it belongs to, so if your project is configured to use the `us-east-1` region of AWS, for example, any `S3BucketFolder`s you declare will use that same region as well. You can adjust this behavior, however, on a case-by-case basis by [customizing the component's resource provider](/docs/intro/concepts/resources/components#inheriting-resource-providers).

For help configuring your cloud-provider credentials, see the Installation &amp; Configuration pages of the [AWS](/registry/packages/aws/installation-configuration), [Azure Native](/registry/packages/azure-native/installation-configuration), and [Google Cloud](/registry/packages/gcp/installation-configuration) provider packages.



### Configuring components

The following input properties are common to all three Synced Folder components:

| Property | Type | Description |
| -------- | ---- | ----------- |
| `path` | `string` | The path (relative or fully-qualified) to the folder containing the files to be synced. Required. |
| `managedObjects` | `boolean` | Whether to have Pulumi manage files as individual cloud resources. Defaults to `true`. See below for details. |

The sections below list component-specific input properties.

#### S3BucketFolder properties

| Property | Type | Description |
| -------- | ---- | ----------- |
| `bucketName` | `string` | The name of the S3 bucket to sync to (e.g., `my-bucket` in `s3://my-bucket`). Required. |
| `acl` | `string` | The AWS [Canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl-overview.html#canned-acl) to apply to each file (e.g., `public-read`). Required. |

#### AzureBlobFolder properties

| Property | Type | Description |
| -------- | ---- | ----------- |
| `containerName` | `string` | The name of the Azure storage container to sync to. Required. |
| `storageAccountName` | `string` | The name of the Azure storage account that the container belongs to. Required. |
| `resourceGroupName` | `string` | The name of the Azure resource group that the storage account belongs to. Required. |

#### GoogleCloudFolder properties

| Property | Type | Description |
| -------- | ---- | ----------- |
| `bucketName` | `string` | The name of the Google Cloud Storage bucket to sync to (e.g., `my-bucket` in `gs://my-bucket`). Required. |
