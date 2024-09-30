---
title: Pulumi AWS API Gateway Installation & Configuration
meta_desc: Information on how to set up credentials to use the AWS API Gateway component.
layout: package
---

{{< aws-resource-note >}}

To provision an AWS API Gateway with this component, you need to have AWS credentials. Use the instructions on the [AWS provider's installation & configuration](/registry/packages/aws/installation-configuration) to get credentials if needed.

Your AWS credentials are never sent to pulumi.com. Pulumi uses the AWS SDK and the credentials in your environment to authenticate requests from your computer to AWS.

## Installation

The AWS API Gateway provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/aws-apigateway`](https://www.npmjs.com/package/@pulumi/aws-apigateway)
* Python: [`pulumi-aws-apigateway`](https://pypi.org/project/pulumi-aws-apigateway/)
* Go: [`github.com/pulumi/pulumi-aws-apigateway/sdk/go/apigateway`](https://github.com/pulumi/pulumi-aws-apigateway)
* .NET: [`Pulumi.AwsApiGateway`](https://www.nuget.org/packages/Pulumi.AwsApiGateway)
* Java: [`com.pulumi/aws-apigateway`](https://central.sonatype.com/artifact/com.pulumi/aws-apigateway)
