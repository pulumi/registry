---
title: CoreDNS
meta_desc: Information on how to set up credentials to use the CoreDNS component.
layout: installation
---

To manage a CoreDNS cert-manager installation with this component, you need to have Kubernetes credentials. Use the instructions on the Kubernetes Provider's [Installation & Configuration](/registry/packages/kubernetes/installation-configuration) to get credentials if needed. Your Kubernetes credentials are never sent to pulumi.com. Pulumi uses the Kubernetes SDK and the credentials in your environment to authenticate requests from your computer to Kubernetes.

## Installation

The CoreDNS provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/kubernetes-coredns`](https://www.npmjs.com/package/@pulumi/kubernetes-coredns)
* Python: [`pulumi-kubernetes-coredns`](https://pypi.org/project/pulumi-kubernetes-coredns/)
* Go: [`github.com/pulumi/pulumi-kubernetes-coredns/sdk/go/kubernetes`](https://github.com/pulumi/pulumi-kubernetes-coredns)
* .NET: [`Pulumi.KubernetesCoreDNS`](https://www.nuget.org/packages/Pulumi.KubernetesCoreDNS)
