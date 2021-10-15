---
title: Pulumi GCP Global CloudRun
meta_desc: Information on how to set up credentials to use the AWS API Gateway component.
layout: installation
---

{{% configure-gcp %}}

To provision a GCP Global CloudRun Application with the this component, you need to have Google Cloud credentials. Use the instructions on the Google Classic Provider's [Installation & Configuration]({{< relref "/registry/packages/gcp/installation-configuration" >}}) to get credentials if needed.

**Your AWS credentials are never sent to Pulumi.com.** Pulumi uses the GCP SDK and the credentials in your environment to authenticate requests from your computer to Global Cloud.
