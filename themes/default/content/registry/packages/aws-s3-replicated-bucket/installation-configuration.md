---
title: AWS S3 Replicated Bucket Installation & Configuration
meta_desc: Information on how to set up credentials to use the AWS S3 Replicated Bucket component.
layout: package
---

{{< aws-resource-note >}}

To provision an AWS S3 Replicated Bucket with this component, you need to have AWS credentials. Use the instructions on the AWS Provider's [Installation & Configuration](/registry/packages/aws/installation-configuration) to get credentials if needed. Your AWS credentials are never sent to Pulumi.com. Pulumi uses the AWS SDK and the credentials in your environment to authenticate requests from your computer to AWS.

## Installation

The AWS S3 Replicated Bucket provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/aws-s3-replicated-bucket`](https://www.npmjs.com/package/@pulumi/aws-s3-replicated-bucket)
* Python: [`pulumi-aws-s3-replicated-bucket`](https://pypi.org/project/pulumi-aws-s3-replicated-bucket/)
* Go: [`github.com/pulumi/pulumi-aws-s3-replicated-bucket/sdk/go/aws`](https://github.com/pulumi/pulumi-aws-s3-replicated-bucket)
* .NET: [`Pulumi.AwsS3ReplicatedBucket`](https://www.nuget.org/packages/Pulumi.AwsS3ReplicatedBucket)
