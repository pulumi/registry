---
title: Metabase Installation
meta_desc: Information on how to set up credentials to use the Pulumi Metabase component.
layout: package
---

{{< aws-resource-note >}}

To provision a Metabase Fargate Service with this Component, you need to have AWS credentials configured. Use the instructions on the AWS Provider's [Installation & Configuration](/registry/packages/aws/installation-configuration) to configure your credentials. Your AWS credentials are never sent to Pulumi.com. Pulumi uses the AWS SDK and the credentials in your environment to authenticate requests from your computer to AWS.

## Installation

The AWS IAM Component is available as a package in:

* JavaScript/TypeScript: [`@pulumi/metabase`](https://www.npmjs.com/package/@pulumi/metabase)
* Python: [`pulumi-metabase`](https://pypi.org/project/pulumi-metabase/)
* Go: [`github.com/pulumi/pulumi-metabase/sdk/go/metabase`](https://github.com/pulumi/pulumi-metabase)
* .NET: [`Pulumi.Metabase`](https://www.nuget.org/packages/Pulumi.Metabase)
