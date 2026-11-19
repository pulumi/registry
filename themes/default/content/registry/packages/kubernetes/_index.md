---
title: Kubernetes
meta_desc: Use the Pulumi Kubernetes provider to manage Kubernetes resources, Helm charts, and YAML manifests in any cluster.
layout: package
---

The Kubernetes provider for Pulumi manages any resource in the [Kubernetes API](https://kubernetes.io/docs/reference/kubernetes-api/), including custom resources, in a cluster you already have. It also deploys [Helm charts](/registry/packages/kubernetes/api-docs/helm/v4/chart/) and [YAML manifests](/registry/packages/kubernetes/api-docs/yaml/v2/configfile/) as Pulumi resources.

The provider connects to your cluster with the same kubeconfig `kubectl` uses. See [Installation & Configuration](./installation-configuration/) for setup and provider options.

## Example

{{< chooser language "typescript,python,go,csharp,java,yaml,hcl" >}}

{{% choosable language typescript %}}

```typescript
import * as k8s from "@pulumi/kubernetes";

const appLabels = { app: "nginx" };
const deployment = new k8s.apps.v1.Deployment("nginx", {
    spec: {
        selector: { matchLabels: appLabels },
        replicas: 1,
        template: {
            metadata: { labels: appLabels },
            spec: { containers: [{ name: "nginx", image: "nginx" }] },
        },
    },
});

export const name = deployment.metadata.name;
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi
from pulumi_kubernetes.apps.v1 import Deployment

app_labels = {"app": "nginx"}

deployment = Deployment(
    "nginx",
    spec={
        "selector": {"match_labels": app_labels},
        "replicas": 1,
        "template": {
            "metadata": {"labels": app_labels},
            "spec": {"containers": [{"name": "nginx", "image": "nginx"}]},
        },
    },
)

pulumi.export("name", deployment.metadata["name"])
```

{{% /choosable %}}
{{% choosable language go %}}

```go
package main

import (
	appsv1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/apps/v1"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		appLabels := pulumi.StringMap{
			"app": pulumi.String("nginx"),
		}
		deployment, err := appsv1.NewDeployment(ctx, "nginx", &appsv1.DeploymentArgs{
			Spec: appsv1.DeploymentSpecArgs{
				Selector: &metav1.LabelSelectorArgs{
					MatchLabels: appLabels,
				},
				Replicas: pulumi.Int(1),
				Template: &corev1.PodTemplateSpecArgs{
					Metadata: &metav1.ObjectMetaArgs{
						Labels: appLabels,
					},
					Spec: &corev1.PodSpecArgs{
						Containers: corev1.ContainerArray{
							corev1.ContainerArgs{
								Name:  pulumi.String("nginx"),
								Image: pulumi.String("nginx"),
							},
						},
					},
				},
			},
		})
		if err != nil {
			return err
		}

		ctx.Export("name", deployment.Metadata.Name())
		return nil
	})
}
```

{{% /choosable %}}
{{% choosable language csharp %}}

```csharp
using System.Collections.Generic;
using Pulumi;
using Pulumi.Kubernetes.Types.Inputs.Apps.V1;
using Pulumi.Kubernetes.Types.Inputs.Core.V1;
using Pulumi.Kubernetes.Types.Inputs.Meta.V1;

return await Deployment.RunAsync(() =>
{
    var appLabels = new InputMap<string>
    {
        { "app", "nginx" },
    };

    var deployment = new Pulumi.Kubernetes.Apps.V1.Deployment("nginx", new DeploymentArgs
    {
        Spec = new DeploymentSpecArgs
        {
            Selector = new LabelSelectorArgs
            {
                MatchLabels = appLabels,
            },
            Replicas = 1,
            Template = new PodTemplateSpecArgs
            {
                Metadata = new ObjectMetaArgs
                {
                    Labels = appLabels,
                },
                Spec = new PodSpecArgs
                {
                    Containers =
                    {
                        new ContainerArgs
                        {
                            Name = "nginx",
                            Image = "nginx",
                        },
                    },
                },
            },
        },
    });

    return new Dictionary<string, object?>
    {
        ["name"] = deployment.Metadata.Apply(m => m.Name),
    };
});
```

{{% /choosable %}}
{{% choosable language java %}}

```java
import com.pulumi.Pulumi;
import com.pulumi.kubernetes.apps.v1.Deployment;
import com.pulumi.kubernetes.apps.v1.DeploymentArgs;
import com.pulumi.kubernetes.apps.v1.inputs.DeploymentSpecArgs;
import com.pulumi.kubernetes.core.v1.inputs.ContainerArgs;
import com.pulumi.kubernetes.core.v1.inputs.PodSpecArgs;
import com.pulumi.kubernetes.core.v1.inputs.PodTemplateSpecArgs;
import com.pulumi.kubernetes.meta.v1.inputs.LabelSelectorArgs;
import com.pulumi.kubernetes.meta.v1.inputs.ObjectMetaArgs;
import java.util.Map;

public class App {
    public static void main(String[] args) {
        Pulumi.run(ctx -> {
            var labels = Map.of("app", "nginx");

            var deployment = new Deployment("nginx", DeploymentArgs.builder()
                .spec(DeploymentSpecArgs.builder()
                    .selector(LabelSelectorArgs.builder()
                        .matchLabels(labels)
                        .build())
                    .replicas(1)
                    .template(PodTemplateSpecArgs.builder()
                        .metadata(ObjectMetaArgs.builder()
                            .labels(labels)
                            .build())
                        .spec(PodSpecArgs.builder()
                            .containers(ContainerArgs.builder()
                                .name("nginx")
                                .image("nginx")
                                .build())
                            .build())
                        .build())
                    .build())
                .build());

            ctx.export("name", deployment.metadata().applyValue(m -> m.name().orElse("")));
        });
    }
}
```

{{% /choosable %}}
{{% choosable language yaml %}}

```yaml
variables:
  appLabels:
    app: nginx

resources:
  nginx:
    type: kubernetes:apps/v1:Deployment
    properties:
      spec:
        selector:
          matchLabels: ${appLabels}
        replicas: 1
        template:
          metadata:
            labels: ${appLabels}
          spec:
            containers:
              - name: nginx
                image: nginx

outputs:
  name: ${nginx.metadata.name}
```

{{% /choosable %}}
{{% choosable language hcl %}}

```hcl
terraform {
  required_providers {
    kubernetes = {
      source  = "pulumi/kubernetes"
      version = "4.34.3"
    }
  }
}

locals {
  labels = { app = "nginx" }
}

resource "kubernetes_apps_v1_deployment" "nginx" {
  spec = {
    selector = { match_labels = local.labels }
    replicas = 1
    template = {
      metadata = { labels = local.labels }
      spec = {
        containers = [{
          name  = "nginx"
          image = "nginx"
        }]
      }
    }
  }
}

output "name" {
  value = kubernetes_apps_v1_deployment.nginx.metadata.name
}
```

{{% /choosable %}}

{{< /chooser >}}

## Create a cluster

The Kubernetes provider deploys to a running cluster. To create one with Pulumi, use your cloud's package:

- [Amazon EKS](/registry/packages/eks/) creates an EKS cluster and the AWS resources it needs in a single component.
- [AWS](/registry/packages/aws/), [Azure Native](/registry/packages/azure-native/), [Google Cloud](/registry/packages/gcp/), and [DigitalOcean](/registry/packages/digitalocean/) provide the lower-level resources for EKS, AKS, GKE, and DigitalOcean Kubernetes.

The [Kubernetes templates](/dev/templates/kubernetes/) start a new project that creates a managed cluster on AWS, Azure, or Google Cloud. To deploy workloads to a cluster from the program that creates it, pass the cluster's kubeconfig to an [explicit provider](/docs/iac/concepts/providers/).

## Learn more

- [Pulumi and Kubernetes](/docs/integrations/clouds/kubernetes/) lists everything Pulumi offers for Kubernetes, including ESC integrations and policy packs.
- [Dev Center tutorials](/dev/tutorials/) walk through complete scenarios, such as deploying Helm charts and managing resources with Server-Side Apply.
- The [Pulumi Kubernetes Operator](/docs/integrations/clouds/kubernetes/pulumi-kubernetes-operator/) runs Pulumi stacks from inside a cluster, driven by `Stack` custom resources.
- [Migrating from Kubernetes YAML](/docs/iac/guides/migration/migrating-to-pulumi/from-kubernetes/) shows how to bring existing manifests and Helm charts into Pulumi.
