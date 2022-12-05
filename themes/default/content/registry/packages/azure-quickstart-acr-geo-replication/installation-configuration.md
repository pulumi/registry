---
title: Azure Quickstart ACR Geo Replication Installation & Configuration
meta_desc: Information on how to set up credentials to use the Azure Quickstart ACR Geo Replicated Registry component.
layout: installation
---

To provision an Azure Quickstart ACR Geo Replication Registry with this component, you need to have Azure credentials. Use the instructions on the Azure Native Provider's [Installation & Configuration]({{< relref "/registry/packages/azure-native/installation-configuration" >}}) to get credentials if needed. Your Azure credentials are never sent to pulumi.com. Pulumi uses the Azure SDK and the credentials in your environment to authenticate requests from your computer to Azure.

## Installation

The Azure Quickstart ACR Geo Replication provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/aws-s3-replicated-bucket`](https://www.npmjs.com/package/@pulumi/aws-s3-replicated-bucket)
* Python: [`pulumi-aws-s3-replicated-bucket`](https://pypi.org/project/pulumi-aws-s3-replicated-bucket/)
* Go: [`github.com/pulumi/pulumi-aws-s3-replicated-bucket/sdk/go/aws`](https://github.com/pulumi/pulumi-aws-s3-replicated-bucket)
* .NET: [`Pulumi.AwsS3ReplicatedBucket`](https://www.nuget.org/packages/Pulumi.AwsS3ReplicatedBucket)
