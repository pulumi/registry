---
# WARNING: this file was fetched from https://raw.githubusercontent.com/castai/pulumi-castai/v7.73.2/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/castai/pulumi-castai/blob/v7.73.2/docs/installation-configuration.md
title: CAST AI Installation & Configuration
meta_desc: Information on how to install the CAST AI provider for Pulumi.
layout: installation
---

## Installation

The Pulumi CAST AI provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@castai/pulumi`](https://www.npmjs.com/package/@castai/pulumi)
* Python: [`pulumi-castai`](https://pypi.org/project/pulumi-castai/)
* Go: [`github.com/castai/pulumi-castai/sdk/go/castai`](https://pkg.go.dev/github.com/castai/pulumi-castai/sdk/go/castai)

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
go get github.com/castai/pulumi-castai/sdk/go/castai@latest
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

{{< chooser language "typescript,python,go" >}}
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
