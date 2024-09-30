---
title: Pulumi AWS QuickStart Aurora Postgres
meta_desc: Information on how to set up credentials to use the AWS QuickStart Aurora Postgres component.
layout: package
---

{{< aws-resource-note >}}

To provision an AWS Aurora Postgres Cluster based on the AWS QuickStart guide with this component, you need to have AWS credentials. Use the instructions on the AWS Provider's [Installation & Configuration](/registry/packages/aws/installation-configuration) to get credentials if needed.

Your AWS credentials are never sent to pulumi.com. Pulumi uses the AWS SDK and the credentials in your environment to authenticate requests from your computer to AWS.

## Installation

The Pulumi AWS QuickStart Aurora Postgres provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/aws-quickstart-aurora-postgres`](https://www.npmjs.com/package/@pulumi/aws-quickstart-aurora-postgres)
* Python: [`pulumi-aws-quickstart-aurora-postgres`](https://pypi.org/project/pulumi-aws-quickstart-aurora-postgres/)
* Go: [`github.com/pulumi/pulumi-aws-quickstart-aurora-postgres/sdk/go/aws`](https://github.com/pulumi/pulumi-aws-quickstart-aurora-postgres)
* .NET: [`Pulumi.AwsQuickStartAuroraPostgres`](https://www.nuget.org/packages/Pulumi.AwsQuickStartAuroraPostgres)
