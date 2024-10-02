---
title: AWS QuickStart Redshift Installtion & Configuration
meta_desc: Information on how to set up credentials to use the Pulumi AWS QuickStart Redshift component.
layout: package
---

{{< aws-resource-note >}}

To provision an AWS Redshift Cluster based on the AWS QuickStart guide with this component, you need to have AWS credentials. Use the instructions on the AWS Provider's [Installation & Configuration](/registry/packages/aws/installation-configuration) to get credentials if needed. Your AWS credentials are never sent to Pulumi.com. Pulumi uses the AWS SDK and the credentials in your environment to authenticate requests from your computer to AWS.

## Installation

The AWS QuickStart Redshift provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/aws-quickstart-redshift`](https://www.npmjs.com/package/@pulumi/aws-quickstart-redshift)
* Python: [`pulumi-aws-quickstart-redshift`](https://pypi.org/project/pulumi-aws-quickstart-redshift/)
* Go: [`github.com/pulumi/pulumi-aws-quickstart-redshift/sdk/go/aws`](https://github.com/pulumi/pulumi-aws-quickstart-redshift)
* .NET: [`Pulumi.AwsQuickStartRedshift`](https://www.nuget.org/packages/Pulumi.AwsQuickStartRedshift)
