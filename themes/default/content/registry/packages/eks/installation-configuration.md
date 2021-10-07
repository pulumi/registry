---
title: Pulumi Amazon EKS Setup
meta_desc: Information on how to set up credentials to use the Amazon EKS component.
layout: installation
---

<!-- markdownlint-disable url -->
[iam-user-console]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_users_create.html#id_users_create_console
[iam-manage-keys]: https://docs.aws.amazon.com/general/latest/gr/managing-aws-access-keys.html
[configure-aws-cli]: https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-getting-started.html
<!-- markdownlint-enable url -->

{{< aws-resource-note >}}

To provision a Kubernetes cluster with the Amazon EKS component, you need to have AWS credentials. Use the instructions on the AWS Classic Provider's [Installation & Configuration page]({{< relref "/registry/packages/aws/installation-configuration" >}}) to get credentials if needed.

**Your AWS credentials are never sent to Pulumi.com.** Pulumi uses the AWS SDK and the credentials in your environment to authenticate requests from your computer to AWS.

## Prerequisites

Before getting started, you will need to install some pre-requisites:

* [AWS CLI version 1.16.156 or later](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-install.html):
  Amazon EKS uses IAM to provide secure authentication to your Kubernetes cluster.

These are not required but are recommended if you plan on interacting with your Kubernetes cluster:

* [`kubectl`](https://kubernetes.io/docs/tasks/tools/install-kubectl/): the standard Kubernetes command line interface.
* [`helm`](https://helm.sh/docs/using_helm/): if you plan on deploying Helm charts to your cluster.
