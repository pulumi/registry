---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/iterative/iterative/0.11.20/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Iterative Provider
meta_desc: Provides an overview on how to configure the Pulumi Iterative provider.
layout: package
---

## Generate Provider

The Iterative provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider iterative/iterative
```
## Overview

![TPI](https://static.iterative.ai/img/tpi/banner.svg)

![tests](https://github.com/iterative/pulumi-provider-iterative/actions/workflows/test.yml)
[![Apache-2.0](https://img.shields.io/badge/licence-Apache%202.0-blue)](https://github.com/iterative/pulumi-provider-iterative/blob/main/LICENSE)

TPI is a Pulumi plugin built with machine learning in mind. This CLI tool offers full lifecycle management of computing resources (including GPUs and respawning spot instances) from several cloud vendors (AWS, Azure, GCP, K8s)... without needing to be a cloud expert.

- **Lower cost with spot recovery**: transparent data checkpoint/restore & auto-respawning of low-cost spot/preemptible instances
- **No cloud vendor lock-in**: switch between clouds with just one line thanks to unified abstraction
- **No waste**: auto-cleanup unused resources (terminate compute instances upon task completion/failure & remove storage upon download of results), pay only for what you use
- **Developer-first experience**: one-command data sync & code execution with no external server, making the cloud feel like a laptop

Supported cloud vendors [include](https://registry.pulumi.io/providers/iterative/iterative/latest/docs/guides/authentication):

|                                                                                                       ![Amazon Web Services                                                                                                        |                                                                                                            [![Microsoft                                                                                                            |                                                                                                         [![Google Cloud Platform                                                                                                          |                                                                                                   [![Kubernetes                                                                                                    |

| (AWS)](https://registry.pulumi.io/providers/iterative/iterative/latest/docs/guides/authentication#amazon-web-services>) | Azure](<https://img.shields.io/badge/Azure-Microsoft_Azure-black?colorA=white&logoColor=0078D4&logo=microsoftazure)](https://registry.pulumi.io/providers/iterative/iterative/latest/docs/guides/authentication#microsoft-azure>) | (GCP)](<https://img.shields.io/badge/GCP-Google_Cloud_Platform-black?colorA=white&logoColor=4285F4&logo=googlecloud)](https://registry.pulumi.io/providers/iterative/iterative/latest/docs/guides/authentication#google-cloud-platform>) | (K8s)](<https://img.shields.io/badge/K8s-Kubernetes-black?colorA=white&logoColor=326CE5&logo=kubernetes)](https://registry.pulumi.io/providers/iterative/iterative/latest/docs/guides/authentication#kubernetes>) |
|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|

![](https://static.iterative.ai/img/tpi/high-level-light.png)
## Why TPI?

There are a several reasons to use TPI instead of other related solutions (custom scripts and/or cloud orchestrators):

1. **Reduced management overhead and infrastructure cost**:
   TPI is a CLI tool, not a running service. It requires no additional orchestrating machine (control plane/head nodes) to schedule/recover/terminate instances. Instead, TPI runs (spot) instances via cloud-native scaling groups ([AWS Auto Scaling Groups](https://docs.aws.amazon.com/autoscaling/ec2/userguide/what-is-amazon-ec2-auto-scaling.html), [Azure VM Scale Sets](https://azure.microsoft.com/en-us/services/virtual-machine-scale-sets), [GCP managed instance groups](https://cloud.google.com/compute/docs/instance-groups#managed_instance_groups), and [Kubernetes Jobs](https://kubernetes.io/docs/concepts/workloads/controllers/job)), taking care of recovery and termination automatically on the cloud provider's side. This design reduces management overhead & infrastructure costs. You can close your laptop while cloud tasks are running -- auto-recovery happens even if you are offline.
2. **Unified tool for data science and software development teams**:
   TPI provides consistent tooling for both data scientists and DevOps engineers, improving cross-team collaboration. This simplifies compute management to a single config file, and reduces time to deliver ML models into production.
3. **Reproducible, codified environments**:
   Store hardware requirements in a single configuration file alongside the rest of your ML pipeline code.

<img width=24px src="https://static.iterative.ai/logo/cml.svg"/> TPI is used to power [CML](https://cml.dev), bringing cloud providers to existing GitHub, GitLab & Bitbucket CI/CD workflows ([repository](https://github.com/iterative/cml)).
## Links

- Getting Started
  - [Authentication](https://registry.pulumi.io/providers/iterative/iterative/latest/docs/guides/authentication)
- Full reference
- Example Projects
  - [Run Jupyter & TensorBoard in the cloud with one command](https://github.com/iterative/blog-tpi-jupyter)
  - [Move local ML experiments to the cloud](https://github.com/iterative/blog-tpi-bees)