---
title: AWS IAM Installation & Configuration
meta_desc: Information on how to set up credentials to use the Pulumi AWS IAM component.
layout: package
---

{{< aws-resource-note >}}

To provision an AWS IAM Roles with this Component, you need to have AWS credentials configured. Use the instructions on the [AWS provider's installation & configuration](/registry/packages/aws/installation-configuration) to configure your credentials. Your AWS credentials are never sent to Pulumi.com. Pulumi uses the AWS SDK and the credentials in your environment to authenticate requests from your computer to AWS.

## Installation

The AWS IAM Component is available as a package in:

* JavaScript/TypeScript: [`@pulumi/aws-iam`](https://www.npmjs.com/package/@pulumi/aws-iam)
* Python: [`pulumi-aws-iam`](https://pypi.org/project/pulumi-aws-iam/)
* Go: [`github.com/pulumi/pulumi-aws-iam/sdk/go/aws-iam`](https://github.com/pulumi/pulumi-aws-iam)
* .NET: [`Pulumi.AwsIam`](https://www.nuget.org/packages/Pulumi.AwsIam)
