---
title: GCP Global CloudRun
meta_desc: Information on how to set up credentials to use the AWS API Gateway component.
layout: installation
---

## Installation

The GCP Global CloudRun provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/aws-gcp-global-cloudrun`](https://www.npmjs.com/package/@pulumi/gcp-global-cloudrun)
* Python: [`pulumi-gcp-global-cloudrun`](https://pypi.org/project/pulumi-gcp-global-cloudrun/)
* Go: [`github.com/pulumi/pulumi-gcp-global-cloudrun/sdk/go/gcp`](https://github.com/pulumi/pulumi-gcp-global-cloudrun)
* .NET: [`Pulumi.GcpGlobalCloudRun`](https://www.nuget.org/packages/Pulumi.GcpGlobalCloudRun)

## Configuration

{{% configure-gcp %}}

To provision a GCP Global CloudRun Application with the this component, you need to have Google Cloud credentials. Use the instructions on the Google Classic Provider's [Installation & Configuration]({{< relref "/registry/packages/gcp/installation-configuration" >}}) to get credentials if needed.

Your AWS credentials are never sent to pulumi.com. Pulumi uses the GCP SDK and the credentials in your environment to authenticate requests from your computer to Global Cloud.
