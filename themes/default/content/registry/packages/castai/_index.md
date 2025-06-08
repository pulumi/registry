---
# WARNING: this file was fetched from https://raw.githubusercontent.com/castai/pulumi-castai/v0.1.87/docs/_index.md
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
import * as gcp from "@pulumi/gcp";

const gcpProjectId = process.env.GCP_PROJECT_ID || "my-gcp-project-id";
const gkeClusterName = process.env.GKE_CLUSTER_NAME || "my-gke-cluster";

// Create a service account for CAST AI
const castaiServiceAccount = new gcp.serviceaccount.Account("castai-service-account", {
    accountId: "castai-gke-access",
    displayName: "CAST AI GKE Access Service Account",
    description: "Service account for CAST AI to manage GKE cluster",
    project: gcpProjectId,
});

// Define the required roles for CAST AI
const requiredRoles = [
    "roles/container.clusterAdmin",
    "roles/compute.instanceAdmin.v1",
    "roles/iam.serviceAccountUser",
];

// Assign roles to the service account
requiredRoles.forEach((role, index) => {
    new gcp.projects.IAMMember(`castai-role-${index}`, {
        project: gcpProjectId,
        role: role,
        member: castaiServiceAccount.email.apply(email => `serviceAccount:${email}`),
    });
});

// Create a service account key
const serviceAccountKey = new gcp.serviceaccount.Key("castai-service-account-key", {
    serviceAccountId: castaiServiceAccount.name,
    publicKeyType: "TYPE_X509_PEM_FILE",
});

// Initialize the CAST AI provider
const provider = new castai.Provider("castai-provider", {
    apiToken: process.env.CASTAI_API_TOKEN,
});

// Connect a GKE cluster to CAST AI using the service account credentials
const gkeCluster = new castai.GkeCluster("gke-cluster-connection", {
    projectId: gcpProjectId,
    location: "us-central1",
    name: gkeClusterName,
    deleteNodesOnDisconnect: true,
    credentialsJson: serviceAccountKey.privateKey,
}, { provider });

// Export the cluster ID and service account information
export const clusterId = gkeCluster.id;
export const serviceAccountEmail = castaiServiceAccount.email;
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi
import os
from pulumi_castai import Provider, GkeCluster
from pulumi_gcp import serviceaccount, projects

# Get GCP project ID from environment variable or use a default value
project_id = os.environ.get("GCP_PROJECT_ID", "my-gcp-project-id")

# Create a service account for CAST AI
castai_service_account = serviceaccount.Account(
    "castai-service-account",
    account_id="castai-gke-access",
    display_name="CAST AI GKE Access Service Account",
    description="Service account for CAST AI to manage GKE cluster",
    project=project_id
)

# Define the required roles for CAST AI
required_roles = [
    "roles/container.clusterAdmin",
    "roles/compute.instanceAdmin.v1",
    "roles/iam.serviceAccountUser",
]

# Assign roles to the service account
for i, role in enumerate(required_roles):
    projects.IAMMember(
        f"castai-role-{i}",
        project=project_id,
        role=role,
        member=castai_service_account.email.apply(lambda email: f"serviceAccount:{email}")
    )

# Create a service account key
service_account_key = serviceaccount.Key(
    "castai-service-account-key",
    service_account_id=castai_service_account.name,
    public_key_type="TYPE_X509_PEM_FILE"
)

# Initialize the CAST AI provider
api_token = os.environ.get("CASTAI_API_TOKEN", "your-api-token-here")
provider = Provider("castai-provider", api_token=api_token)

# Get GKE cluster name from environment variable or use a default value
cluster_name = os.environ.get("GKE_CLUSTER_NAME", "my-gke-cluster")

# Create a connection to a GKE cluster using the service account credentials
gke_cluster = GkeCluster("gke-cluster-connection",
    project_id=project_id,
    location="us-central1",
    name=cluster_name,
    delete_nodes_on_disconnect=True,
    credentials_json=service_account_key.private_key,
    opts=pulumi.ResourceOptions(provider=provider)
)

# Export the cluster ID and service account information
pulumi.export("cluster_id", gke_cluster.id)
pulumi.export("service_account_email", castai_service_account.email)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
package main

import (
	"os"

	"github.com/castai/pulumi-castai/sdk/go/castai"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/projects"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/serviceaccount"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Get GCP project ID from environment variable or use a default value
		projectID := os.Getenv("GCP_PROJECT_ID")
		if projectID == "" {
			projectID = "my-gcp-project-id"
		}

		// Get GKE cluster name from environment variable or use a default value
		clusterName := os.Getenv("GKE_CLUSTER_NAME")
		if clusterName == "" {
			clusterName = "my-gke-cluster"
		}

		// Create a service account for CAST AI
		castaiServiceAccount, err := serviceaccount.NewAccount(ctx, "castai-service-account", &serviceaccount.AccountArgs{
			AccountId:   pulumi.String("castai-gke-access"),
			DisplayName: pulumi.String("CAST AI GKE Access Service Account"),
			Description: pulumi.String("Service account for CAST AI to manage GKE cluster"),
			Project:     pulumi.String(projectID),
		})
		if err != nil {
			return err
		}

		// Define the required roles for CAST AI
		requiredRoles := []string{
			"roles/container.clusterAdmin",
			"roles/compute.instanceAdmin.v1",
			"roles/iam.serviceAccountUser",
		}

		// Assign roles to the service account
		for i, role := range requiredRoles {
			_, err := projects.NewIAMMember(ctx, pulumi.Sprintf("castai-role-%d", i), &projects.IAMMemberArgs{
				Project: pulumi.String(projectID),
				Role:    pulumi.String(role),
				Member:  pulumi.Sprintf("serviceAccount:%s", castaiServiceAccount.Email),
			})
			if err != nil {
				return err
			}
		}

		// Create a service account key
		serviceAccountKey, err := serviceaccount.NewKey(ctx, "castai-service-account-key", &serviceaccount.KeyArgs{
			ServiceAccountId: castaiServiceAccount.Name,
			PublicKeyType:    pulumi.String("TYPE_X509_PEM_FILE"),
		})
		if err != nil {
			return err
		}

		// Initialize the CAST AI provider
		provider, err := castai.NewProvider(ctx, "castai-provider", &castai.ProviderArgs{
			ApiToken: pulumi.String(os.Getenv("CASTAI_API_TOKEN")),
		})
		if err != nil {
			return err
		}

		// Create a connection to a GKE cluster using the service account credentials
		gkeArgs := &castai.GkeClusterArgs{
			ProjectId:               pulumi.String(projectID),
			Location:                pulumi.String("us-central1"),
			Name:                    pulumi.String(clusterName),
			DeleteNodesOnDisconnect: pulumi.Bool(true),
			CredentialsJson:         serviceAccountKey.PrivateKey,
		}

		gkeCluster, err := castai.NewGkeCluster(ctx, "gke-cluster-connection", gkeArgs, pulumi.Provider(provider))
		if err != nil {
			return err
		}

		// Export useful information
		ctx.Export("clusterId", gkeCluster.ID())
		ctx.Export("serviceAccountEmail", castaiServiceAccount.Email)

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
