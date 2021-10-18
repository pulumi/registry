---
title: NGINX Ingress Controller Package
meta_desc: Information on how to set up credentials to use the NGINX Ingress Controller Package component.
layout: installation
---

To manage an NGINX Ingress Controller installation with this component, you need to have Kubernetes credentials. Use the instructions on the Kubernetes Provider's [Installation & Configuration]({{< relref "/registry/packages/kubernetes/installation-configuration" >}}) to get credentials if needed.

**Your Kubernetes credentials are never sent to Pulumi.com.** Pulumi uses the Kubernetes SDK and the credentials in your environment to authenticate requests from your computer to Kubernetes.
