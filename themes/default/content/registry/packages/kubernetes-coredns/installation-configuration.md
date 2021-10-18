---
title: CoreDNS
meta_desc: Information on how to set up credentials to use the CoreDNS component.
layout: installation
---

To manage a CoreDNS cert-manager installation with this component, you need to have Kubernetes credentials. Use the instructions on the Kubernetes Provider's [Installation & Configuration]({{< relref "/registry/packages/kubernetes/installation-configuration" >}}) to get credentials if needed.

**Your Kubernetes credentials are never sent to Pulumi.com.** Pulumi uses the Kubernetes SDK and the credentials in your environment to authenticate requests from your computer to Kubernetes.
