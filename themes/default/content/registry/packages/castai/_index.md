---
# WARNING: this file was fetched from https://raw.githubusercontent.com/castai/pulumi-castai/v0.1.45/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
title: CAST AI
meta_desc: Provides an overview of the CAST AI Provider for Pulumi.
layout: overview
---

# CAST AI Provider

The CAST AI Provider for Pulumi enables you to manage [CAST AI](https://cast.ai/) resources in your cloud infrastructure using Pulumi. CAST AI is a Kubernetes cost optimization platform that helps you reduce cloud costs by up to 50% through automated instance selection, scaling, and spot instance management.

## Example

{{< chooser language "typescript,python,go,csharp" >}}
{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as castai from "@castai/pulumi";

// Initialize the CAST AI provider
const provider = new castai.Provider("castai-provider", {
    apiToken: process.env.CASTAI_API_TOKEN,
    apiUrl: process.env.CASTAI_API_URL || "https://api.cast.ai",
});

// Connect a GKE cluster to CAST AI
const gkeCluster = new castai.GkeCluster("gke-cluster-connection", {
    projectId: process.env.GCP_PROJECT_ID || "my-gcp-project-id",
    location: "us-central1",
    name: process.env.GKE_CLUSTER_NAME || "cast_ai_test_cluster",
    deleteNodesOnDisconnect: true,
}, { provider });

// Export the cluster ID
export const clusterId = gkeCluster.id;
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi
import os
from pulumi_castai import Provider, GkeCluster

# Initialize the CAST AI provider
api_token = os.environ.get("CASTAI_API_TOKEN", "your-api-token-here")
provider = Provider("castai-provider", api_token=api_token)

# Connect a GKE cluster to CAST AI
gke_cluster = GkeCluster("gke-cluster-connection",
    project_id=os.environ.get("GCP_PROJECT_ID", "my-gcp-project-id"),
    location="us-central1",
    name=os.environ.get("GKE_CLUSTER_NAME", "cast_ai_test_cluster"),
    delete_nodes_on_disconnect=True,
    opts=pulumi.ResourceOptions(provider=provider)
)

# Export the cluster ID
pulumi.export("cluster_id", gke_cluster.id)
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
		// Initialize the provider
		provider, err := castai.NewProvider(ctx, "castai-provider", &castai.ProviderArgs{
			ApiToken: pulumi.String(os.Getenv("CASTAI_API_TOKEN")),
		})
		if err != nil {
			return err
		}

		// Get GCP project ID from environment variable or use a default value
		projectID := os.Getenv("GCP_PROJECT_ID")
		if projectID == "" {
			projectID = "my-gcp-project-id"
		}

		// Get GKE cluster name from environment variable or use a default value
		clusterName := os.Getenv("GKE_CLUSTER_NAME")
		if clusterName == "" {
			clusterName = "cast_ai_test_cluster"
		}

		// Connect a GKE cluster to CAST AI
		gkeCluster, err := castai.NewGkeCluster(ctx, "gke-cluster-connection", &castai.GkeClusterArgs{
			ProjectId:              pulumi.String(projectID),
			Location:               pulumi.String("us-central1"),
			Name:                   pulumi.String(clusterName),
			DeleteNodesOnDisconnect: pulumi.Bool(true),
		}, pulumi.Provider(provider))
		if err != nil {
			return err
		}

		// Export the cluster ID
		ctx.Export("clusterId", gkeCluster.ID())
		return nil
	})
}
```

{{% /choosable %}}
{{% choosable language csharp %}}

```csharp
using System;
using Pulumi;
using Pulumi.CastAI;

class MyStack : Stack
{
    public MyStack()
    {
        // Initialize the CAST AI provider
        var provider = new Provider("castai-provider", new ProviderArgs
        {
            ApiToken = Environment.GetEnvironmentVariable("CASTAI_API_TOKEN"),
            ApiUrl = Environment.GetEnvironmentVariable("CASTAI_API_URL") ?? "https://api.cast.ai"
        });

        // Connect a GKE cluster to CAST AI
        var gkeCluster = new GkeCluster("gke-cluster-connection", new GkeClusterArgs
        {
            ProjectId = Environment.GetEnvironmentVariable("GCP_PROJECT_ID") ?? "my-gcp-project-id",
            Location = "us-central1",
            Name = Environment.GetEnvironmentVariable("GKE_CLUSTER_NAME") ?? "cast_ai_test_cluster",
            DeleteNodesOnDisconnect = true
        }, new CustomResourceOptions
        {
            Provider = provider
        });

        // Export the cluster ID
        this.ClusterId = gkeCluster.Id;
    }

    [Output] public Output<string> ClusterId { get; set; }
}
```

{{% /choosable %}}
{{< /chooser >}}

## Features

The CAST AI Provider offers resources to:

* Connect your Kubernetes clusters (EKS, GKE, AKS) to CAST AI for cost optimization
* Configure autoscaling policies for your clusters
* Manage node configurations and templates
* Set up IAM roles and service accounts for CAST AI
* Configure cost optimization policies

## Available Resources

The CAST AI provider supports the following resources:

### Cloud Provider Resources
- AWS EKS clusters: `castai:aws:EksCluster`
- GCP GKE clusters: `castai:gcp:GkeCluster`
- Azure AKS clusters: `castai:azure:AksCluster`

### Core Resources
- Cluster: `castai:index:Cluster`
- Credentials: `castai:index:Credentials`
- Cluster Token: `castai:index:ClusterToken`

### Autoscaling Resources
- Autoscaler: `castai:autoscaling:Autoscaler`

### Organization Resources
- Service Account: `castai:organization:ServiceAccount`
- Service Account Key: `castai:organization:ServiceAccountKey`
