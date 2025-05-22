---
# WARNING: this file was fetched from https://raw.githubusercontent.com/castai/pulumi-castai/v0.1.78/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
title: CAST AI
meta_desc: Provides an overview of the CAST AI Provider for Pulumi.
layout: overview
---

The CAST AI Provider for Pulumi enables you to manage CAST AI resources in your cloud infrastructure using Pulumi. CAST AI is a Kubernetes cost optimization platform that helps you reduce cloud costs by automatically optimizing your Kubernetes clusters.

## Example

{{< chooser language "typescript,python,go,csharp" >}}

{{% choosable language typescript %}}

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
    name: process.env.EKS_CLUSTER_NAME || "my-eks-cluster",
    deleteNodesOnDisconnect: true,
    overrideSecurityGroups: ["sg-12345678"],
    subnets: ["subnet-12345678", "subnet-87654321"],
}, { provider });

// Export the cluster ID
export const clusterId = eksCluster.id;
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi
import os
from pulumi_castai import Provider, EksCluster

# Initialize the CAST AI provider
api_token = os.environ.get("CASTAI_API_TOKEN", "your-api-token-here")
provider = Provider("castai-provider", api_token=api_token)

# Get AWS values from environment variables or use defaults
aws_region = os.environ.get("AWS_REGION", "us-west-2")
aws_account_id = os.environ.get("AWS_ACCOUNT_ID", "123456789012")
eks_cluster_name = os.environ.get("EKS_CLUSTER_NAME", "my-eks-cluster")

# Create a connection to an EKS cluster
eks_cluster = EksCluster("eks-cluster-connection",
    account_id=aws_account_id,
    region=aws_region,
    name=eks_cluster_name,
    delete_nodes_on_disconnect=True,
    override_security_groups=["sg-12345678"],
    subnets=["subnet-12345678", "subnet-87654321"],
    opts=pulumi.ResourceOptions(provider=provider)
)

# Export the cluster ID
pulumi.export("cluster_id", eks_cluster.id)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
package main

import (
	"os"

	"github.com/castai/pulumi-castai/sdk/go/castai"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Initialize the provider (API token will be read from environment variable CASTAI_API_TOKEN)
		provider, err := castai.NewProvider(ctx, "castai-provider", &castai.ProviderArgs{})
		if err != nil {
			return err
		}

		// Get AWS account ID from environment variable or use a default value
		accountID := os.Getenv("AWS_ACCOUNT_ID")
		if accountID == "" {
			accountID = "123456789012"
		}

		// Get AWS region from environment variable or use a default value
		region := os.Getenv("AWS_REGION")
		if region == "" {
			region = "us-west-2"
		}

		// Get EKS cluster name from environment variable or use a default value
		clusterName := os.Getenv("EKS_CLUSTER_NAME")
		if clusterName == "" {
			clusterName = "my-eks-cluster"
		}

		// Create a connection to an EKS cluster
		eksArgs := &castai.EksClusterArgs{
			AccountId:              pulumi.String(accountID),
			Region:                 pulumi.String(region),
			Name:         pulumi.String(clusterName),
			DeleteNodesOnDisconnect: pulumi.Bool(true),
			OverrideSecurityGroups:         pulumi.StringArray{pulumi.String("sg-12345678")},
			Subnets:              pulumi.StringArray{pulumi.String("subnet-12345678"), pulumi.String("subnet-87654321")},
		}

		eksCluster, err := castai.NewEksCluster(ctx, "eks-cluster-connection", eksArgs, pulumi.Provider(provider))
		if err != nil {
			return err
		}

		// Export the cluster ID
		ctx.Export("clusterId", eksCluster.ID())

		return nil
	})
}
```

{{% /choosable %}}
{{% choosable language csharp %}}

```csharp
using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using Pulumi;
using Pulumi.CastAI;

return await Deployment.RunAsync(() =>
{
    // Initialize the CAST AI provider
    var provider = new Provider("castai-provider", new ProviderArgs
    {
        ApiToken = Environment.GetEnvironmentVariable("CASTAI_API_TOKEN")
    });

    // Get AWS values from environment variables or use defaults
    var awsRegion = Environment.GetEnvironmentVariable("AWS_REGION") ?? "us-west-2";
    var awsAccountId = Environment.GetEnvironmentVariable("AWS_ACCOUNT_ID") ?? "123456789012";
    var eksClusterName = Environment.GetEnvironmentVariable("EKS_CLUSTER_NAME") ?? "my-eks-cluster";

    // Create a connection to an EKS cluster
    var eksCluster = new EksCluster("eks-cluster-connection", new EksClusterArgs
    {
        AccountId = awsAccountId,
        Region = awsRegion,
        Name = eksClusterName,
        DeleteNodesOnDisconnect = true,
        OverrideSecurityGroups = new[] {"sg-12345678"},
        Subnets = new[] { "subnet-12345678", "subnet-87654321" }
    }, new CustomResourceOptions
    {
        Provider = provider
    });

    // Export the cluster ID
    return new Dictionary<string, object?>
    {
        ["ClusterId"] = eksCluster.Id
    };
});
```

{{% /choosable %}}
{{< /chooser >}}

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
