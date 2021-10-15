---
title: Run My Darn Container Setup
meta_desc: Information on how to set up credentials to use the Run My Darn Container component.
layout: installation
---

To provision a container in one of the cloud providers with the this component, you need to have credentials.

For AWS, use the instructions from the AWS Classic Provider's [Installation & Configuration]({{< relref "/registry/packages/aws/installation-configuration" >}}) to get credentials if needed.

For Azure, use the instructions from the Azure Native Provider's [Installation & Configuration]({{< relref "/registry/packages/azure-native/installation-configuration" >}}) to get credentials if needed.

For Google, use the instructions from the GCP Classic Provider's [Installation & Configuration]({{< relref "/registry/packages/gcp/installation-configuration" >}}) to get credentials if needed.

**Your credentials are never sent to Pulumi.com.** Pulumi uses the respective cloud SDKs and the credentials in your environment to authenticate requests from your computer to the cloud providers.
