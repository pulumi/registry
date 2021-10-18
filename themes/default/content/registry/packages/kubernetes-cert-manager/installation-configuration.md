---
title: Jetstack Cert Manager 
meta_desc: Information on how to set up credentials to use the Jetstack Cert Manager component.
layout: installation
---

To manage a Jetstack Cert Manager installation with this component, you need to have Kubernetes credentials. Use the instructions on the Kubernetes Provider's [Installation & Configuration]({{< relref "/registry/packages/kubernetes/installation-configuration" >}}) to get credentials if needed.

**Your Kubernetes credentials are never sent to Pulumi.com.** Pulumi uses the Kubernetes SDK and the credentials in your environment to authenticate requests from your computer to Kubernetes.
