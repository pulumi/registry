---
# WARNING: this file was fetched from https://raw.githubusercontent.com/castai/pulumi-castai/v7.73.2/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/castai/pulumi-castai/blob/v7.73.2/docs/_index.md
title: CAST AI
meta_desc: Provides an overview of the CAST AI Provider for Pulumi.
layout: overview
---

The CAST AI Provider for Pulumi enables you to manage CAST AI resources in your cloud infrastructure using Pulumi. CAST AI is a Kubernetes cost optimization platform that helps you reduce cloud costs by automatically optimizing your Kubernetes clusters.

## Important: Installation Order

For CAST AI to work properly, you must install the CAST AI agent **before** creating the cluster connection. The correct order is:

1. **Install CAST AI agent** (using Helm or Kubernetes manifests)
2. **Install CAST AI cluster controller** (using Helm or Kubernetes manifests)
3. **Create CAST AI cluster connection** (using Pulumi CAST AI provider)

This ensures the cluster controller is running and can respond when CAST AI attempts to connect to your cluster.

## Example

{{< chooser language "typescript,python,go" >}}

{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as castai from "@castai/pulumi";
import * as gcp from "@pulumi/gcp";
import * as k8s from "@pulumi/kubernetes";

const castaiApiToken = process.env.CASTAI_API_TOKEN;
if (!castaiApiToken) {
    console.error("ERROR: CASTAI_API_TOKEN environment variable is required");
    process.exit(1);
}

const gcpProjectId = process.env.GCP_PROJECT_ID || "my-gcp-project-id";
const gkeClusterName = process.env.GKE_CLUSTER_NAME || "pulumi-test-cluster-" + Date.now();
const gkeLocation = process.env.GKE_LOCATION || "us-central1";

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
const roleBindings = requiredRoles.map((role, index) => {
    return new gcp.projects.IAMMember(`castai-role-${index}`, {
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
    apiToken: castaiApiToken,
    apiUrl: process.env.CASTAI_API_URL || "https://api.cast.ai",
});

// Create a GKE cluster for testing
const gkeClusterInfo = new gcp.container.Cluster("test-gke-cluster", {
    name: gkeClusterName,
    location: gkeLocation,
    project: gcpProjectId,
    initialNodeCount: 4, // Increased to 4 nodes to ensure enough capacity for CAST AI agents
    nodeConfig: {
        machineType: "e2-standard-2", // Increased from e2-medium to ensure enough resources for CAST AI agents
        oauthScopes: [
            "https://www.googleapis.com/auth/cloud-platform",
        ],
    },
    removeDefaultNodePool: false,
    deletionProtection: false,
});

// Create a Kubernetes provider to interact with the GKE cluster
const k8sProvider = new k8s.Provider("gke-k8s", {
    kubeconfig: gkeClusterInfo.endpoint.apply(endpoint => {
        return `apiVersion: v1
kind: Config
clusters:
- cluster:
    server: https://${endpoint}
    insecure-skip-tls-verify: true
  name: ${gkeClusterName}
contexts:
- context:
    cluster: ${gkeClusterName}
    user: ${gkeClusterName}
  name: ${gkeClusterName}
current-context: ${gkeClusterName}
users:
- name: ${gkeClusterName}
  user:
    exec:
      apiVersion: client.authentication.k8s.io/v1beta1
      command: gke-gcloud-auth-plugin
      installHint: Install gke-gcloud-auth-plugin for use with kubectl by following https://cloud.google.com/blog/products/containers-kubernetes/kubectl-auth-changes-in-gke
`;
    }),
});

// Install the CAST AI agent using Helm
const castaiAgent = new k8s.helm.v3.Release("castai-agent", {
    name: "castai-agent", // This will be the exact name used, without a suffix
    chart: "castai-agent",
    repositoryOpts: {
        repo: "https://castai.github.io/helm-charts",
    },
    namespace: "castai-agent",
    createNamespace: true,
    cleanupOnFail: true,
    timeout: 300,
    skipAwait: true,
    values: {
        replicaCount: 1,
        provider: "gke",
        additionalEnv: {
            STATIC_CLUSTER_ID: gkeCluster.id,
        },
        createNamespace: false,
        apiURL: process.env.CASTAI_API_URL || "https://api.cast.ai",
        apiKey: castaiApiToken,
        resources: {
            agent: {
                requests: {
                    memory: "512Mi",
                    cpu: "100m",
                },
                limits: {
                    memory: "1Gi",
                    cpu: "500m",
                },
            },
            monitor: {
                requests: {
                    memory: "64Mi",
                    cpu: "50m",
                },
            },
        },
    },
}, {
    provider: k8sProvider,
    dependsOn: [gkeCluster],
    customTimeouts: {
        create: "1m",
        update: "1m",
        delete: "5m",
    },
});

// Install the CAST AI cluster controller
const clusterController = new k8s.helm.v3.Release("cluster-controller", {
    name: "cluster-controller", // This will be the exact name used, without a suffix
    chart: "castai-cluster-controller",
    repositoryOpts: {
        repo: "https://castai.github.io/helm-charts",
    },
    namespace: "castai-agent",
    createNamespace: true,
    cleanupOnFail: true,
    timeout: 300,
    skipAwait: true,
    values: {
        castai: {
            clusterID: gkeCluster.id,
            apiURL: process.env.CASTAI_API_URL || "https://api.cast.ai",
            apiKey: castaiApiToken,
        },
        resources: {
            requests: {
                memory: "128Mi",
                cpu: "50m",
            },
            limits: {
                memory: "256Mi",
                cpu: "200m",
            },
        },
    },
}, {
    provider: k8sProvider,
    dependsOn: [castaiAgent, gkeCluster],
    customTimeouts: {
        create: "1m",
        update: "1m",
        delete: "5m",
    },
});

// Connect GKE cluster to CAST AI
const gkeCluster = new castai.GkeCluster("gke-cluster-connection", {
    projectId: gcpProjectId,
    location: gkeLocation,
    name: gkeClusterName,
    deleteNodesOnDisconnect: true,
    credentialsJson: serviceAccountKey.privateKey.apply(key =>
        Buffer.from(key, 'base64').toString('utf8')
    ),
}, {
    provider,
    dependsOn: [gkeClusterInfo, ...roleBindings],
    customTimeouts: {
        create: "2m",
        update: "5m",
        delete: "5m",
    },
});

// Install the CAST AI evictor
const castaiEvictor = new k8s.helm.v3.Release("castai-evictor", {
    name: "castai-evictor", // This will be the exact name used, without a suffix
    chart: "castai-evictor",
    repositoryOpts: {
        repo: "https://castai.github.io/helm-charts",
    },
    namespace: "castai-agent",
    createNamespace: true,
    cleanupOnFail: true,
    timeout: 300,
    skipAwait: true,
    values: {
        replicaCount: 1,
        managedByCASTAI: true,
    },
}, {
    provider: k8sProvider,
    dependsOn: [castaiAgent, clusterController],
    customTimeouts: {
        create: "1m",
        update: "1m",
        delete: "5m",
    },
});

// Install the CAST AI pod pinner
const castaiPodPinner = new k8s.helm.v3.Release("castai-pod-pinner", {
    name: "castai-pod-pinner", // This will be the exact name used, without a suffix
    chart: "castai-pod-pinner",
    repositoryOpts: {
        repo: "https://castai.github.io/helm-charts",
    },
    namespace: "castai-agent",
    createNamespace: true,
    cleanupOnFail: true,
    timeout: 300,
    skipAwait: true,
    values: {
        castai: {
            apiKey: castaiApiToken,
            clusterID: gkeCluster.id,
        },
        replicaCount: 0,
    },
}, {
    provider: k8sProvider,
    dependsOn: [castaiAgent, clusterController],
    customTimeouts: {
        create: "1m",
        update: "1m",
        delete: "5m",
    },
});

// Export useful information
export const clusterName = gkeClusterName;
export const clusterId = gkeCluster.id;
export const serviceAccountEmail = castaiServiceAccount.email;
export const serviceAccountName = castaiServiceAccount.name;
export const agentHelmRelease = castaiAgent.name;
export const controllerHelmRelease = clusterController.name;
export const evictorHelmRelease = castaiEvictor.name;
export const podPinnerHelmRelease = castaiPodPinner.name;
```

### Connecting to an Existing GKE Cluster

If you already have a GKE cluster and want to connect it to CAST AI, you can use this simplified example that reads cluster information from your current Kubernetes context:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as castai from "@castai/pulumi";
import * as gcp from "@pulumi/gcp";
import * as k8s from "@pulumi/kubernetes";

const castaiApiToken = process.env.CASTAI_API_TOKEN;
if (!castaiApiToken) {
    console.error("ERROR: CASTAI_API_TOKEN environment variable is required");
    process.exit(1);
}

// Get cluster information from environment or use defaults
const gcpProjectId = process.env.GCP_PROJECT_ID || "my-gcp-project-id";
const gkeClusterName = process.env.GKE_CLUSTER_NAME || "my-existing-cluster";
const gkeLocation = process.env.GKE_LOCATION || "us-central1";

// Generate a short unique suffix to avoid naming conflicts
const uniqueSuffix = Math.random().toString(36).substring(2, 8); // 6 character random string

// Create a service account for CAST AI
const castaiServiceAccount = new gcp.serviceaccount.Account("castai-sa", {
    accountId: `castai-gke-${uniqueSuffix}`, // Keep it short: castai-gke-abc123
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
const roleBindings = requiredRoles.map((role, index) => {
    return new gcp.projects.IAMMember(`castai-role-${index}-${uniqueSuffix}`, {
        project: gcpProjectId,
        role: role,
        member: castaiServiceAccount.email.apply(email => `serviceAccount:${email}`),
    });
});

// Create a service account key
const serviceAccountKey = new gcp.serviceaccount.Key(`castai-key-${uniqueSuffix}`, {
    serviceAccountId: castaiServiceAccount.name,
    publicKeyType: "TYPE_X509_PEM_FILE",
});

// Initialize the CAST AI provider
const provider = new castai.Provider(`castai-provider-${uniqueSuffix}`, {
    apiToken: castaiApiToken,
    apiUrl: process.env.CASTAI_API_URL || "https://api.cast.ai",
});

// Create a Kubernetes provider using current context (assumes kubectl is configured)
const k8sProvider = new k8s.Provider(`gke-k8s-${uniqueSuffix}`, {
    // Uses current kubeconfig context - make sure kubectl is pointing to your GKE cluster
});

// Connect existing GKE cluster to CAST AI
const gkeCluster = new castai.GkeCluster(`gke-cluster-${uniqueSuffix}`, {
    projectId: gcpProjectId,
    location: gkeLocation,
    name: gkeClusterName,
    deleteNodesOnDisconnect: true,
    credentialsJson: serviceAccountKey.privateKey.apply(key =>
        Buffer.from(key, 'base64').toString('utf8')
    ),
}, {
    provider,
    dependsOn: roleBindings,
    customTimeouts: {
        create: "2m",
        update: "5m",
        delete: "5m",
    },
});

// Install the CAST AI agent using Helm
const castaiAgent = new k8s.helm.v3.Release(`castai-agent-${uniqueSuffix}`, {
    name: `castai-agent-${uniqueSuffix}`, // Use unique name to avoid conflicts
    chart: "castai-agent",
    repositoryOpts: {
        repo: "https://castai.github.io/helm-charts",
    },
    namespace: "castai-agent",
    createNamespace: true,
    cleanupOnFail: true,
    timeout: 300,
    skipAwait: true,
    values: {
        replicaCount: 1,
        provider: "gke",
        additionalEnv: {
            STATIC_CLUSTER_ID: gkeCluster.id,
        },
        createNamespace: false,
        apiURL: process.env.CASTAI_API_URL || "https://api.cast.ai",
        apiKey: castaiApiToken,
        resources: {
            agent: {
                requests: {
                    memory: "512Mi",
                    cpu: "100m",
                },
                limits: {
                    memory: "1Gi",
                    cpu: "500m",
                },
            },
            monitor: {
                requests: {
                    memory: "64Mi",
                    cpu: "50m",
                },
            },
        },
    },
}, {
    provider: k8sProvider,
    dependsOn: [gkeCluster],
    customTimeouts: {
        create: "1m",
        update: "1m",
        delete: "5m",
    },
});

// Install the CAST AI cluster controller
const clusterController = new k8s.helm.v3.Release(`cluster-controller-${uniqueSuffix}`, {
    name: `cluster-controller-${uniqueSuffix}`, // Use unique name to avoid conflicts
    chart: "castai-cluster-controller",
    repositoryOpts: {
        repo: "https://castai.github.io/helm-charts",
    },
    namespace: "castai-agent",
    createNamespace: true,
    cleanupOnFail: true,
    timeout: 300,
    skipAwait: true,
    values: {
        castai: {
            clusterID: gkeCluster.id,
            apiURL: process.env.CASTAI_API_URL || "https://api.cast.ai",
            apiKey: castaiApiToken,
        },
        resources: {
            requests: {
                memory: "128Mi",
                cpu: "50m",
            },
            limits: {
                memory: "256Mi",
                cpu: "200m",
            },
        },
    },
}, {
    provider: k8sProvider,
    dependsOn: [castaiAgent, gkeCluster],
    customTimeouts: {
        create: "1m",
        update: "1m",
        delete: "5m",
    },
});

// Install the CAST AI evictor
const castaiEvictor = new k8s.helm.v3.Release(`castai-evictor-${uniqueSuffix}`, {
    name: `castai-evictor-${uniqueSuffix}`, // Use unique name to avoid conflicts
    chart: "castai-evictor",
    repositoryOpts: {
        repo: "https://castai.github.io/helm-charts",
    },
    namespace: "castai-agent",
    createNamespace: true,
    cleanupOnFail: true,
    timeout: 300,
    skipAwait: true,
    values: {
        replicaCount: 1,
        managedByCASTAI: true,
    },
}, {
    provider: k8sProvider,
    dependsOn: [castaiAgent, clusterController],
    customTimeouts: {
        create: "1m",
        update: "1m",
        delete: "5m",
    },
});

// Install the CAST AI pod pinner
const castaiPodPinner = new k8s.helm.v3.Release(`castai-pod-pinner-${uniqueSuffix}`, {
    name: `castai-pod-pinner-${uniqueSuffix}`, // Use unique name to avoid conflicts
    chart: "castai-pod-pinner",
    repositoryOpts: {
        repo: "https://castai.github.io/helm-charts",
    },
    namespace: "castai-agent",
    createNamespace: true,
    cleanupOnFail: true,
    timeout: 300,
    skipAwait: true,
    values: {
        castai: {
            apiKey: castaiApiToken,
            clusterID: gkeCluster.id,
        },
        replicaCount: 0,
    },
}, {
    provider: k8sProvider,
    dependsOn: [castaiAgent, clusterController],
    customTimeouts: {
        create: "1m",
        update: "1m",
        delete: "5m",
    },
});

// Export useful information
export const clusterName = gkeClusterName;
export const clusterId = gkeCluster.id;
export const serviceAccountEmail = castaiServiceAccount.email;
export const agentHelmRelease = castaiAgent.name;
export const controllerHelmRelease = clusterController.name;
export const evictorHelmRelease = castaiEvictor.name;
export const podPinnerHelmRelease = castaiPodPinner.name;
```

**Prerequisites for the existing cluster example:**
- Your GKE cluster must already exist
- `kubectl` must be configured to point to your GKE cluster
- You must have the necessary permissions to install Helm charts in the cluster
- Set the environment variables: `CASTAI_API_TOKEN`, `GCP_PROJECT_ID`, `GKE_CLUSTER_NAME`, and `GKE_LOCATION`

**To configure kubectl for your GKE cluster:**
```bash
gcloud container clusters get-credentials YOUR_CLUSTER_NAME --location=YOUR_LOCATION --project=YOUR_PROJECT_ID
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi
import os
import base64
import time
from pulumi_castai import Provider, GkeCluster
from pulumi_gcp import serviceaccount, projects, container
from pulumi_kubernetes import core, helm, Provider as K8sProvider

# Check for required environment variables
castai_api_token = os.environ.get("CASTAI_API_TOKEN")
if not castai_api_token:
    raise Exception("ERROR: CASTAI_API_TOKEN environment variable is required")

# Get configuration from environment variables
project_id = os.environ.get("GCP_PROJECT_ID", "my-gcp-project-id")
cluster_name = os.environ.get("GKE_CLUSTER_NAME", f"pulumi-test-cluster-{int(time.time())}")
gke_location = os.environ.get("GKE_LOCATION", "us-central1")

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
role_bindings = []
for i, role in enumerate(required_roles):
    role_binding = projects.IAMMember(
        f"castai-role-{i}",
        project=project_id,
        role=role,
        member=castai_service_account.email.apply(lambda email: f"serviceAccount:{email}")
    )
    role_bindings.append(role_binding)

# Create a service account key
service_account_key = serviceaccount.Key(
    "castai-service-account-key",
    service_account_id=castai_service_account.name,
    public_key_type="TYPE_X509_PEM_FILE"
)

# Initialize the CAST AI provider
castai_provider = Provider(
    "castai-provider",
    api_token=castai_api_token,
    api_url=os.environ.get("CASTAI_API_URL", "https://api.cast.ai")
)

# Create a GKE cluster for testing
gke_cluster_info = container.Cluster(
    "test-gke-cluster",
    name=cluster_name,
    location=gke_location,
    project=project_id,
    initial_node_count=4,  # Increased to 4 nodes to ensure enough capacity for CAST AI agents
    node_config={
        "machine_type": "e2-standard-2",  # Increased from e2-medium to ensure enough resources for CAST AI agents
        "oauth_scopes": [
            "https://www.googleapis.com/auth/cloud-platform",
        ],
    },
    remove_default_node_pool=False,
    deletion_protection=False
)

# Create a Kubernetes provider to interact with the GKE cluster
k8s_provider = K8sProvider(
    "gke-k8s",
    kubeconfig=gke_cluster_info.endpoint.apply(lambda endpoint: f"""apiVersion: v1
kind: Config
clusters:
- cluster:
    server: https://{endpoint}
    insecure-skip-tls-verify: true
  name: {cluster_name}
contexts:
- context:
    cluster: {cluster_name}
    user: {cluster_name}
  name: {cluster_name}
current-context: {cluster_name}
users:
- name: {cluster_name}
  user:
    exec:
      apiVersion: client.authentication.k8s.io/v1beta1
      command: gke-gcloud-auth-plugin
      installHint: Install gke-gcloud-auth-plugin for use with kubectl by following https://cloud.google.com/blog/products/containers-kubernetes/kubectl-auth-changes-in-gke
""")
)

# Install the CAST AI agent using Helm
castai_agent = helm.v3.Release(
    "castai-agent",
    name="castai-agent",  # This will be the exact name used, without a suffix
    chart="castai-agent",
    repository_opts={"repo": "https://castai.github.io/helm-charts"},
    namespace="castai-agent",
    create_namespace=True,
    cleanup_on_fail=True,
    timeout=300,
    skip_await=True,
    values={
        "replicaCount": 1,
        "provider": "gke",
        "additionalEnv": {
            "STATIC_CLUSTER_ID": gke_cluster.id,
        },
        "createNamespace": False,
        "apiURL": os.environ.get("CASTAI_API_URL", "https://api.cast.ai"),
        "apiKey": castai_api_token,
        "resources": {
            "agent": {
                "requests": {
                    "memory": "512Mi",
                    "cpu": "100m",
                },
                "limits": {
                    "memory": "1Gi",
                    "cpu": "500m",
                },
            },
            "monitor": {
                "requests": {
                    "memory": "64Mi",
                    "cpu": "50m",
                },
            },
        },
    },
    opts=pulumi.ResourceOptions(
        provider=k8s_provider,
        depends_on=[gke_cluster],
        custom_timeouts=pulumi.CustomTimeouts(
            create="1m",
            update="1m",
            delete="5m",
        ),
    )
)

# Install the CAST AI cluster controller
cluster_controller = helm.v3.Release(
    "cluster-controller",
    name="cluster-controller",  # This will be the exact name used, without a suffix
    chart="castai-cluster-controller",
    repository_opts={"repo": "https://castai.github.io/helm-charts"},
    namespace="castai-agent",
    create_namespace=True,
    cleanup_on_fail=True,
    timeout=300,
    skip_await=True,
    values={
        "castai": {
            "clusterID": gke_cluster.id,
            "apiURL": os.environ.get("CASTAI_API_URL", "https://api.cast.ai"),
            "apiKey": castai_api_token,
        },
        "resources": {
            "requests": {
                "memory": "128Mi",
                "cpu": "50m",
            },
            "limits": {
                "memory": "256Mi",
                "cpu": "200m",
            },
        },
    },
    opts=pulumi.ResourceOptions(
        provider=k8s_provider,
        depends_on=[castai_agent, gke_cluster],
        custom_timeouts=pulumi.CustomTimeouts(
            create="1m",
            update="1m",
            delete="5m",
        ),
    )
)

# Connect GKE cluster to CAST AI
gke_cluster = GkeCluster("gke-cluster-connection",
    project_id=project_id,
    location=gke_location,
    name=cluster_name,
    delete_nodes_on_disconnect=True,
    credentials_json=service_account_key.private_key.apply(
        lambda key: base64.b64decode(key).decode('utf-8')
    ),
    opts=pulumi.ResourceOptions(
        provider=castai_provider,
        depends_on=[gke_cluster_info] + role_bindings,
        custom_timeouts=pulumi.CustomTimeouts(
            create="2m",
            update="5m",
            delete="5m",
        ),
    )
)

# Install the CAST AI evictor
castai_evictor = helm.v3.Release(
    "castai-evictor",
    name="castai-evictor",  # This will be the exact name used, without a suffix
    chart="castai-evictor",
    repository_opts={"repo": "https://castai.github.io/helm-charts"},
    namespace="castai-agent",
    create_namespace=True,
    cleanup_on_fail=True,
    timeout=300,
    skip_await=True,
    values={
        "replicaCount": 1,
        "managedByCASTAI": True,
    },
    opts=pulumi.ResourceOptions(
        provider=k8s_provider,
        depends_on=[castai_agent, cluster_controller],
        custom_timeouts=pulumi.CustomTimeouts(
            create="1m",
            update="1m",
            delete="5m",
        ),
    )
)

# Install the CAST AI pod pinner
castai_pod_pinner = helm.v3.Release(
    "castai-pod-pinner",
    name="castai-pod-pinner",  # This will be the exact name used, without a suffix
    chart="castai-pod-pinner",
    repository_opts={"repo": "https://castai.github.io/helm-charts"},
    namespace="castai-agent",
    create_namespace=True,
    cleanup_on_fail=True,
    timeout=300,
    skip_await=True,
    values={
        "castai": {
            "apiKey": castai_api_token,
            "clusterID": gke_cluster.id,
        },
        "replicaCount": 0,
    },
    opts=pulumi.ResourceOptions(
        provider=k8s_provider,
        depends_on=[castai_agent, cluster_controller],
        custom_timeouts=pulumi.CustomTimeouts(
            create="1m",
            update="1m",
            delete="5m",
        ),
    )
)

# Export useful information
pulumi.export("cluster_name", cluster_name)
pulumi.export("cluster_id", gke_cluster.id)
pulumi.export("service_account_email", castai_service_account.email)
pulumi.export("service_account_name", castai_service_account.name)
pulumi.export("agent_helm_release", castai_agent.name)
pulumi.export("controller_helm_release", cluster_controller.name)
pulumi.export("evictor_helm_release", castai_evictor.name)
pulumi.export("pod_pinner_helm_release", castai_pod_pinner.name)
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
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	helmv3 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Get values from environment variables or use defaults
		projectID := os.Getenv("GCP_PROJECT_ID")
		if projectID == "" {
			projectID = "my-gcp-project-id"
		}
		clusterName := os.Getenv("GKE_CLUSTER_NAME")
		if clusterName == "" {
			clusterName = "my-gke-cluster"
		}
		apiToken := os.Getenv("CASTAI_API_TOKEN")

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

		// Define the required roles for CAST AI (using broader permissions)
		requiredRoles := []string{
			"roles/container.admin",
			"roles/compute.admin",
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
		})
		if err != nil {
			return err
		}

		// Initialize providers
		castaiProvider, err := castai.NewProvider(ctx, "castai-provider", &castai.ProviderArgs{
			ApiToken: pulumi.String(apiToken),
		})
		if err != nil {
			return err
		}

		k8sProvider, err := kubernetes.NewProvider(ctx, "gke-k8s", &kubernetes.ProviderArgs{})
		if err != nil {
			return err
		}

		// STEP 1: Create namespace with proper Helm labels
		castaiNamespace, err := corev1.NewNamespace(ctx, "castai-namespace", &corev1.NamespaceArgs{
			Metadata: &metav1.ObjectMetaArgs{
				Name: pulumi.String("castai-agent"),
				Labels: pulumi.StringMap{
					"app.kubernetes.io/managed-by": pulumi.String("Helm"),
				},
				Annotations: pulumi.StringMap{
					"meta.helm.sh/release-name":      pulumi.String("castai-agent"),
					"meta.helm.sh/release-namespace": pulumi.String("castai-agent"),
				},
			},
		}, pulumi.Provider(k8sProvider))
		if err != nil {
			return err
		}

		// STEP 2: Install CAST AI agent FIRST (before cluster connection)
		castaiAgent, err := helmv3.NewRelease(ctx, "castai-agent", &helmv3.ReleaseArgs{
			Name:  pulumi.String("castai-agent"),
			Chart: pulumi.String("castai-agent"),
			RepositoryOpts: &helmv3.RepositoryOptsArgs{
				Repo: pulumi.String("https://castai.github.io/helm-charts"),
			},
			Namespace:       pulumi.String("castai-agent"),
			CreateNamespace: pulumi.Bool(false),
			Values: pulumi.Map{
				"apiKey":  pulumi.String(apiToken),
				"provider": pulumi.String("gke"),
				"apiURL":  pulumi.String("https://api.cast.ai"),
			},
		}, pulumi.Provider(k8sProvider), pulumi.DependsOn([]pulumi.Resource{castaiNamespace}))
		if err != nil {
			return err
		}

		// STEP 3: Install CAST AI cluster controller
		clusterController, err := helmv3.NewRelease(ctx, "cluster-controller", &helmv3.ReleaseArgs{
			Name:  pulumi.String("cluster-controller"),
			Chart: pulumi.String("castai-cluster-controller"),
			RepositoryOpts: &helmv3.RepositoryOptsArgs{
				Repo: pulumi.String("https://castai.github.io/helm-charts"),
			},
			Namespace: pulumi.String("castai-agent"),
			Values: pulumi.Map{
				"castai": pulumi.Map{
					"apiKey": pulumi.String(apiToken),
					"apiURL": pulumi.String("https://api.cast.ai"),
				},
			},
		}, pulumi.Provider(k8sProvider), pulumi.DependsOn([]pulumi.Resource{castaiAgent}))
		if err != nil {
			return err
		}

		// STEP 4: Connect GKE cluster to CAST AI AFTER agent installation
		gkeCluster, err := castai.NewGkeCluster(ctx, "gke-cluster-connection", &castai.GkeClusterArgs{
			ProjectId:               pulumi.String(projectID),
			Location:                pulumi.String("us-central1"),
			Name:                    pulumi.String(clusterName),
			DeleteNodesOnDisconnect: pulumi.Bool(true),
			CredentialsJson:         serviceAccountKey.PrivateKey,
		}, pulumi.Provider(castaiProvider), pulumi.DependsOn([]pulumi.Resource{castaiAgent, clusterController}))
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
{{< /chooser >}}

## Alternative: Manual Agent Installation

If you prefer to install the CAST AI agent manually using Helm or kubectl, you can do so before running your Pulumi program:

```bash
# Add CAST AI Helm repository
helm repo add castai-helm https://castai.github.io/helm-charts
helm repo update

# Install CAST AI agent
helm install castai-agent castai-helm/castai-agent \
  --namespace castai-agent \
  --create-namespace \
  --set apiKey=$CASTAI_API_TOKEN \
  --set provider=gke \
  --set apiURL=https://api.cast.ai

# Install CAST AI cluster controller
helm install cluster-controller castai-helm/castai-cluster-controller \
  --namespace castai-agent \
  --set castai.apiKey=$CASTAI_API_TOKEN \
  --set castai.apiURL=https://api.cast.ai
```

After manual installation, you can use a simplified Pulumi program that only creates the cluster connection.

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
