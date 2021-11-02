---
title: Pulumi AWS Miniflux Setup
meta_desc: Information on how to set up credentials to use the AWS Miniflux component.
layout: installation
---

{{< aws-resource-note >}}

To provision a Miniflux installation with the this component, you need to have AWS credentials. Use the instructions on the AWS Classic Provider's [Installation & Configuration]({{< relref "/registry/packages/aws/installation-configuration" >}}) to get credentials if needed.

Your AWS credentials are never sent to pulumi.com. Pulumi uses the AWS SDK and the credentials in your environment to authenticate requests from your computer to AWS.

## Installation

The AWS Miniflux component is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/aws-miniflux`](https://www.npmjs.com/package/@pulumi/aws-miniflux)
* Python: [`pulumi-aws-miniflux`](https://pypi.org/project/pulumi-aws-miniflux/)
* Go: [`github.com/pulumi/pulumi-aws-miniflux/sdk/go/miniflux`](https://github.com/pulumi/pulumi-aws-miniflux)
* .NET: [`Pulumi.AwsMiniflux`](https://www.nuget.org/packages/Pulumi.AwsMiniflux)
