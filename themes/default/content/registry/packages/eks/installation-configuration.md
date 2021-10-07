---
title: Pulumi Amazon EKS Component Setup
meta_desc: This page provides an overview on how to configure credentials for the Pulumi component for Amazon EKS.
layout: installation
---

The Pulumi Amazon EKS provider uses the AWS SDK to manage and provision resources.

> Pulumi relies on the AWS SDK to authenticate requests from your computer to the resources. Your credentials are never sent
> to pulumi.com.

The Pulumi Amazon EKS Provider needs to be configured with AWS credentials
before it can be used to create resources.

## Prerequisites

Before getting started, you will need to install some pre-requisites:

* [AWS CLI version 1.16.156 or later](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-install.html):
  Amazon EKS uses IAM to provide secure authentication to your Kubernetes cluster.

These are not required but are recommended if you plan on interacting with your Kubernetes cluster:

* [`kubectl`](https://kubernetes.io/docs/tasks/tools/install-kubectl/): the standard Kubernetes command line interface.
* [`helm`](https://helm.sh/docs/using_helm/): if you plan on deploying Helm charts to your cluster.

## Getting Your AWS Credentials

See the AWS Provider's [Installation & Configuration]({{ relref "/registry/packages/aws/installation-configuration" }}) page for information on how to get AWS credentials.
