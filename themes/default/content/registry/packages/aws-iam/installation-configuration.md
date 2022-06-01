---
title: AWS IAM
meta_desc: Information on how to set up credentials to use the Pulumi AWS IAM component.
layout: installation
---

{{< aws-resource-note >}}

To provision an AWS IAM Roles with this component, you need to have AWS credentials. Use the instructions on the AWS Classic Provider's [Installation & Configuration]({{< relref "/registry/packages/aws/installation-configuration" >}}) to get credentials if needed. Your AWS credentials are never sent to Pulumi.com. Pulumi uses the AWS SDK and the credentials in your environment to authenticate requests from your computer to AWS.

## Installation

The AWS QuickStart Redshift provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/aws-iam`](https://www.npmjs.com/package/@pulumi/aws-iam)
* Python: [`pulumi-aws-iam`](https://pypi.org/project/pulumi-aws-iam/)
* Go: [`github.com/pulumi/pulumi-aws-iam/sdk/go/aws-iam`](https://github.com/pulumi/pulumi-aws-iam)
* .NET: [`Pulumi.AwsIam`](https://www.nuget.org/packages/Pulumi.AwsIam)
