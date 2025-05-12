---
# WARNING: this file was fetched from https://raw.githubusercontent.com/castai/pulumi-castai/v0.1.72/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
title: CAST AI
meta_desc: Provides an overview of the CAST AI Provider for Pulumi.
layout: overview
---

The CAST AI Provider for Pulumi enables you to manage CAST AI resources in your cloud infrastructure using Pulumi. CAST AI is a Kubernetes cost optimization platform that helps you reduce cloud costs by automatically optimizing your Kubernetes clusters.

## Example

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as castai from "@castai/pulumi";

// Initialize the CAST AI provider
const provider = new castai.Provider("castai-provider", {
    apiToken: process.env.CASTAI_API_TOKEN,
});

// Connect an EKS cluster to CAST AI
const eksCluster = new castai.EksCluster("eks-cluster-connection", {
    accountId: process.env.AWS_ACCOUNT_ID || "123456789012",
    region: process.env.AWS_REGION || "us-west-2",
    eksClusterName: process.env.EKS_CLUSTER_NAME || "my-eks-cluster",
    deleteNodesOnDisconnect: true,
    securityGroupId: "sg-12345678",
    subnetIds: ["subnet-12345678", "subnet-87654321"],
}, { provider });

// Export the cluster ID
export const clusterId = eksCluster.id;
```

## Features

The CAST AI Provider for Pulumi offers resources to:

* Connect your Kubernetes clusters (EKS, GKE, AKS) to CAST AI
* Configure autoscaling policies
* Manage node configurations
* Set up cost optimization policies
* Create and manage service accounts for CAST AI

## Supported Cloud Providers

CAST AI supports the following cloud providers:

* Amazon Web Services (AWS) - EKS clusters
* Google Cloud Platform (GCP) - GKE clusters
* Microsoft Azure - AKS clusters

## Authentication

To use the CAST AI provider, you need to have a CAST AI account and an API token. You can sign up for a CAST AI account at [https://cast.ai](https://cast.ai) and generate an API token from the CAST AI console.
