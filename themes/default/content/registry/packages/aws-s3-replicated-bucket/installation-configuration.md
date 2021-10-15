---
title: AWS S3 Replicated Bucket
meta_desc: Information on how to set up credentials to use the AWS S3 Replicated Bucket component.
layout: installation
---

{{< aws-resource-note >}}

To provision an AWS S3 Replicated Bucket with this component, you need to have AWS credentials. Use the instructions on the AWS Classic Provider's [Installation & Configuration]({{< relref "/registry/packages/aws/installation-configuration" >}}) to get credentials if needed.

**Your AWS credentials are never sent to Pulumi.com.** Pulumi uses the AWS SDK and the credentials in your environment to authenticate requests from your computer to AWS.
