---
title: cert-manager
meta_desc: Information on how to set up credentials to use the cert-manager component.
layout: installation
---

To manage a cert-manager installation with this component, you need to have Kubernetes credentials. Use the instructions on the Kubernetes Provider's [Installation & Configuration](/registry/packages/kubernetes/installation-configuration) to get credentials if needed. Your Kubernetes credentials are never sent to pulumi.com. Pulumi uses the Kubernetes SDK and the credentials in your environment to authenticate requests from your computer to Kubernetes.

## Installation

The cert-manager provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/kubernetes-cert-manager`](https://www.npmjs.com/package/@pulumi/kubernetes-cert-manager)
* Python: [`pulumi-kubernetes-cert-manager`](https://pypi.org/project/pulumi-kubernetes-cert-manager/)
* Go: [`github.com/pulumi/pulumi-kubernetes-cert-manager/sdk/go/kubernetes`](https://github.com/pulumi/pulumi-kubernetes-cert-manager)
* .NET: [`Pulumi.KubernetesCertManager`](https://www.nuget.org/packages/Pulumi.KubernetesCertManager)
