---
title: AWS Classic
meta_desc: Learn how you can use Pulumi's AWS Classic Provider to reduce the complexity of provisioning and managing resources on AWS.
layout: overview
---

The Amazon Web Services (AWS) provider for Pulumi can provision many of the cloud resources available in [AWS](https://aws.amazon.com/). It uses the AWS SDK to manage and provision resources.

The AWS provider must be configured with credentials to deploy and update resources in AWS; see [Installation & Configuration]({{<relref "./installation-configuration">}}) for instructions.

**New to Pulumi and AWS?** [Get started with AWS using our tutorial]({{<relref "/docs/get-started/aws">}})

{{% notes %}}
Pulumi has a new AWS provider: the [Pulumi AWS Native Provider]({{<relref "/registry/packages/aws-native">}}). AWS Native gives you same-day access to all new AWS resources and includes coverage of all resources in the [AWS Cloud Control API](https://aws.amazon.com/blogs/aws/announcing-aws-cloud-control-api/).

Consider trying [AWS Native]({{<relref "/registry/packages/aws-native">}}) if you need AWS resources that aren't available in this provider.
{{% /notes %}}

## Example

```typescript
const aws = require("@pulumi/aws");

const bucket = new aws.s3.Bucket("mybucket");
```

Visit the [How-to Guides]({{<relref "./how-to-guides">}}) to find step-by-step guides for specific scenarios like creating a serverless application or setting up Athena search.

## Components

Pulumi offers Components that provide simpler interfaces and higher-productivity APIs for many areas of AWS:

* [Amazon EKS]({{<relref "/registry/packages/eks">}})
* [Crosswalk for AWS]({{<relref "/docs/guides/crosswalk/aws">}}), which includes API Gateway, CloudWatch, Elastic Container Registry, Elastic Container Service, Elastic Kubernetes Service, Elastic Load Balancing, Identity & Access Management, Lambda, Virtual Private Cloud, and more

## SDK packages

The AWS provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/aws`](https://www.npmjs.com/package/@pulumi/aws)
* Python: [`pulumi-aws`](https://pypi.org/project/pulumi-aws/)
* Go: [`github.com/pulumi/pulumi-aws/sdk/v4/go/aws`](https://github.com/pulumi/pulumi-aws)
* .NET: [`Pulumi.Aws`](https://www.nuget.org/packages/Pulumi.Aws)
