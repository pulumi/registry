---
# WARNING: this file was fetched from https://raw.githubusercontent.com/castai/pulumi-castai/v0.1.45/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
title: CAST AI Installation & Configuration
meta_desc: Information on how to install the CAST AI provider for Pulumi.
layout: installation
---

## Installation

The Pulumi CAST AI provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@castai/pulumi`](https://www.npmjs.com/package/@castai/pulumi)
* Python: [`pulumi-castai`](https://pypi.org/project/pulumi-castai/)
* Go: [`github.com/castai/pulumi-castai/sdk/go/castai`](https://pkg.go.dev/github.com/castai/pulumi-castai/sdk/go/castai)
* .NET: [`Pulumi.CastAI`](https://www.nuget.org/packages/CastAI.Pulumi)

### Node.js (JavaScript/TypeScript)

```bash
npm install @castai/pulumi
```

### Python

```bash
pip install pulumi_castai
```

### Go

```bash
go get github.com/castai/pulumi-castai/sdk/go/castai
```

### .NET

```bash
dotnet add package Pulumi.CastAI
```

## Setup

To use the CAST AI provider, you need to have a CAST AI account and an API token. You can sign up for a CAST AI account at [https://cast.ai](https://cast.ai) and generate an API token from the CAST AI console.

### Authentication

The CAST AI provider requires an API token for authentication. You can provide this token in several ways:

1. Set the `CASTAI_API_TOKEN` environment variable:

```bash
export CASTAI_API_TOKEN=your-api-token-here
```

2. Provide the token directly in your Pulumi program:

{{< chooser language "typescript,python,go,csharp" >}}
{{% choosable language typescript %}}

```typescript
import * as castai from "@castai/pulumi";

const provider = new castai.Provider("castai-provider", {
    apiToken: "your-api-token-here",
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi_castai as castai

provider = castai.Provider("castai-provider", api_token="your-api-token-here")
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
    "github.com/castai/pulumi-castai/sdk/go/castai"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
    pulumi.Run(func(ctx *pulumi.Context) error {
        provider, err := castai.NewProvider(ctx, "castai-provider", &castai.ProviderArgs{
            ApiToken: pulumi.String("your-api-token-here"),
        })
        if err != nil {
            return err
        }

        // Use the provider to create resources
        return nil
    })
}
```

{{% /choosable %}}
{{% choosable language csharp %}}

```csharp
using Pulumi;
using Pulumi.CastAI;

class MyStack : Stack
{
    public MyStack()
    {
        var provider = new Provider("castai-provider", new ProviderArgs
        {
            ApiToken = "your-api-token-here",
        });
    }
}
```

{{% /choosable %}}
{{< /chooser >}}

### Configuration Options

The CAST AI provider accepts the following configuration options:

| Option | Description | Environment Variable | Default |
|--------|-------------|----------------------|---------|
| `apiToken` | CAST AI API token | `CASTAI_API_TOKEN` | - |
| `apiUrl` | CAST AI API URL | `CASTAI_API_URL` | `https://api.cast.ai` |

## Cloud Provider Credentials

To connect your Kubernetes clusters to CAST AI, you'll need to provide credentials for your cloud provider. The specific credentials required depend on the cloud provider:

### AWS (EKS)

For AWS EKS clusters, you'll need:

- AWS Access Key ID and Secret Access Key
- AWS Region
- EKS Cluster Name
- Security Group ID
- Subnet IDs

### GCP (GKE)

For GCP GKE clusters, you'll need:

- GCP Project ID
- GCP Location (region or zone)
- GKE Cluster Name
- GCP Credentials (service account key)

### Azure (AKS)

For Azure AKS clusters, you'll need:

- Azure Subscription ID
- Azure Tenant ID
- Azure Resource Group Name
- AKS Cluster Name

## Examples

Here are examples of connecting different types of Kubernetes clusters to CAST AI:

### AWS EKS

{{< chooser language "typescript,python,go" >}}
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
    eksClusterName: process.env.EKS_CLUSTER_NAME || "my-eks-cluster",
    deleteNodesOnDisconnect: true,
    securityGroupId: "sg-12345678",
    subnetIds: ["subnet-12345678", "subnet-87654321"],
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
provider = Provider("castai-provider", api_token=os.environ.get("CASTAI_API_TOKEN"))

# Connect an EKS cluster to CAST AI
eks_cluster = EksCluster("eks-cluster-connection",
    account_id=os.environ.get("AWS_ACCOUNT_ID", "123456789012"),
    region=os.environ.get("AWS_REGION", "us-west-2"),
    eks_cluster_name=os.environ.get("EKS_CLUSTER_NAME", "my-eks-cluster"),
    delete_nodes_on_disconnect=True,
    security_group_id="sg-12345678",
    subnet_ids=["subnet-12345678", "subnet-87654321"],
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
		// Initialize the provider
		provider, err := castai.NewProvider(ctx, "castai-provider", &castai.ProviderArgs{
			ApiToken: pulumi.String(os.Getenv("CASTAI_API_TOKEN")),
		})
		if err != nil {
			return err
		}

		// Connect an EKS cluster to CAST AI
		eksCluster, err := castai.NewEksCluster(ctx, "eks-cluster-connection", &castai.EksClusterArgs{
			AccountId:              pulumi.String(os.Getenv("AWS_ACCOUNT_ID")),
			Region:                 pulumi.String(os.Getenv("AWS_REGION")),
			EksClusterName:         pulumi.String(os.Getenv("EKS_CLUSTER_NAME")),
			DeleteNodesOnDisconnect: pulumi.Bool(true),
			SecurityGroupId:        pulumi.String("sg-12345678"),
			SubnetIds:              pulumi.StringArray{pulumi.String("subnet-12345678"), pulumi.String("subnet-87654321")},
		}, pulumi.Provider(provider))
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
{{< /chooser >}}

### GCP GKE

{{< chooser language "typescript,python,go" >}}
{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as castai from "@pulumi/castai";

// Initialize the CAST AI provider
const provider = new castai.Provider("castai-provider", {
    apiToken: process.env.CASTAI_API_TOKEN,
});

// Connect a GKE cluster to CAST AI
const gkeCluster = new castai.GkeCluster("gke-cluster-connection", {
    projectId: process.env.GCP_PROJECT_ID || "my-gcp-project-id",
    location: process.env.GKE_LOCATION || "us-central1",
    name: process.env.GKE_CLUSTER_NAME || "my-gke-cluster",
    deleteNodesOnDisconnect: true,
    credentialsJson: process.env.GOOGLE_CREDENTIALS,
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
provider = Provider("castai-provider", api_token=os.environ.get("CASTAI_API_TOKEN"))

# Connect a GKE cluster to CAST AI
gke_cluster = GkeCluster("gke-cluster-connection",
    project_id=os.environ.get("GCP_PROJECT_ID", "my-gcp-project-id"),
    location=os.environ.get("GKE_LOCATION", "us-central1"),
    name=os.environ.get("GKE_CLUSTER_NAME", "my-gke-cluster"),
    delete_nodes_on_disconnect=True,
    credentials_json=os.environ.get("GOOGLE_CREDENTIALS"),
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

		// Connect a GKE cluster to CAST AI
		gkeCluster, err := castai.NewGkeCluster(ctx, "gke-cluster-connection", &castai.GkeClusterArgs{
			ProjectId:              pulumi.String(os.Getenv("GCP_PROJECT_ID")),
			Location:               pulumi.String(os.Getenv("GKE_LOCATION")),
			Name:                   pulumi.String(os.Getenv("GKE_CLUSTER_NAME")),
			DeleteNodesOnDisconnect: pulumi.Bool(true),
			CredentialsJson:        pulumi.String(os.Getenv("GOOGLE_CREDENTIALS")),
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
{{< /chooser >}}

### Azure AKS

{{< chooser language "typescript,python,go" >}}
{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as castai from "@castai/pulumi";

// Initialize the CAST AI provider
const provider = new castai.Provider("castai-provider", {
    apiToken: process.env.CASTAI_API_TOKEN,
});

// Connect an AKS cluster to CAST AI
const aksCluster = new castai.AksCluster("aks-cluster-connection", {
    subscriptionId: process.env.AZURE_SUBSCRIPTION_ID || "00000000-0000-0000-0000-000000000000",
    tenantId: process.env.AZURE_TENANT_ID || "00000000-0000-0000-0000-000000000000",
    resourceGroupName: process.env.AZURE_RESOURCE_GROUP || "my-resource-group",
    aksClusterName: process.env.AKS_CLUSTER_NAME || "my-aks-cluster",
    deleteNodesOnDisconnect: true,
}, { provider });

// Export the cluster ID
export const clusterId = aksCluster.id;
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi
import os
from pulumi_castai import Provider, AksCluster

# Initialize the CAST AI provider
provider = Provider("castai-provider", api_token=os.environ.get("CASTAI_API_TOKEN"))

# Connect an AKS cluster to CAST AI
aks_cluster = AksCluster("aks-cluster-connection",
    subscription_id=os.environ.get("AZURE_SUBSCRIPTION_ID", "00000000-0000-0000-0000-000000000000"),
    tenant_id=os.environ.get("AZURE_TENANT_ID", "00000000-0000-0000-0000-000000000000"),
    resource_group_name=os.environ.get("AZURE_RESOURCE_GROUP", "my-resource-group"),
    aks_cluster_name=os.environ.get("AKS_CLUSTER_NAME", "my-aks-cluster"),
    delete_nodes_on_disconnect=True,
    opts=pulumi.ResourceOptions(provider=provider)
)

# Export the cluster ID
pulumi.export("cluster_id", aks_cluster.id)
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

		// Connect an AKS cluster to CAST AI
		aksCluster, err := castai.NewAksCluster(ctx, "aks-cluster-connection", &castai.AksClusterArgs{
			SubscriptionId:         pulumi.String(os.Getenv("AZURE_SUBSCRIPTION_ID")),
			TenantId:               pulumi.String(os.Getenv("AZURE_TENANT_ID")),
			ResourceGroupName:      pulumi.String(os.Getenv("AZURE_RESOURCE_GROUP")),
			AksClusterName:         pulumi.String(os.Getenv("AKS_CLUSTER_NAME")),
			DeleteNodesOnDisconnect: pulumi.Bool(true),
		}, pulumi.Provider(provider))
		if err != nil {
			return err
		}

		// Export the cluster ID
		ctx.Export("clusterId", aksCluster.ID())
		return nil
	})
}
```

{{% /choosable %}}
{{< /chooser >}}
