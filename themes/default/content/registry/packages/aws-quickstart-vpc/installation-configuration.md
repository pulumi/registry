---
title: AWS QuickStart VPC Installation & Configuration
meta_desc: Information on how to set up credentials to use the AWS QuickStart VPC component.
layout: package
---

{{< aws-resource-note >}}

To provision an AWS VPC based on the AWS QuickStart guide with this component, you need to have AWS credentials. Use the instructions on the AWS Provider's [Installation & Configuration](/registry/packages/aws/installation-configuration) to get credentials if needed. Your AWS credentials are never sent to Pulumi.com. Pulumi uses the AWS SDK and the credentials in your environment to authenticate requests from your computer to AWS.

## Installation

The AWS QuickStart VPC provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/aws-quickstart-vpc`](https://www.npmjs.com/package/@pulumi/aws-quickstart-vpc)
* Python: [`pulumi-aws-quickstart-vpc`](https://pypi.org/project/pulumi-aws-quickstart-vpc/)
* Go: [`github.com/pulumi/pulumi-aws-quickstart-vpc/sdk/go/aws`](https://github.com/pulumi/pulumi-aws-quickstart-vpc)
* .NET: [`Pulumi.AwsQuickStartVpc`](https://www.nuget.org/packages/Pulumi.AwsQuickStartVpc)
