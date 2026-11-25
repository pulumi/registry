---
title: Kubernetes Cert Manager
meta_desc: Use Pulumi's Component for cert-manager offers a Pulumi-friendly and strongly-typed way to manage cert-manager installations using infrastructure as code.
layout: package
---

cert-manager is a Kubernetes add-on to automate the management and issuance of TLS certificates from various issuing sources.
It will ensure certificates are valid and up to date periodically, and attempt to renew certificates at an appropriate time before expiry.
The Pulumi Component for cert-manager helps teams easily manage cert-manager installations as a package available in all Pulumi languages.
cert-manager was created by [Jetstack](https://jetstack.io) and is now a [CNCF Member Project](https://cert-manager.io/).

## Example

{{< chooser language "typescript,python,go,csharp,java,yaml,hcl" >}}

{{% choosable language typescript %}}

```typescript
import * as k8s from "@pulumi/kubernetes";
import * as certmanager from "@pulumi/kubernetes-cert-manager";

// Create a sandbox namespace.
const nsName = "sandbox";
const ns = new k8s.core.v1.Namespace("sandbox-ns", {
    metadata: { name: nsName },
});

// Install cert-manager into our cluster.
const manager = new certmanager.CertManager("cert-manager", {
    installCRDs: true,
    helmOptions: {
        namespace: nsName,
    },
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi
from pulumi_kubernetes.core.v1 import Namespace
from pulumi_kubernetes_cert_manager import CertManager, ReleaseArgs

# Create a sandbox namespace.
ns_name = 'sandbox'
ns = Namespace('sandbox-ns', metadata={ 'name': ns_name })

# Install cert-manager into our cluster.
manager = CertManager('cert-manager',
                      install_crds=True,
                      helm_options=ReleaseArgs(
                          namespace=ns_name,
                      ))
```

{{% /choosable %}}

{{% choosable language go %}}

```go
import (
	kubernetescertmanager "github.com/pulumi/pulumi-kubernetes-cert-manager/sdk/go/kubernetes-cert-manager"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a sandbox namespace.
		sandboxNs, err := corev1.NewNamespace(ctx, "sandbox-ns", &corev1.NamespaceArgs{
			Metadata: &metav1.ObjectMetaArgs{
				Name: pulumi.String("sandbox"),
			},
		})
		if err != nil {
			return err
		}
		// Install cert-manager into our cluster.
		_, err = kubernetescertmanager.NewCertManager(ctx, "cert-manager", &kubernetescertmanager.CertManagerArgs{
			InstallCRDs: pulumi.Bool(true),
			HelmOptions: &kubernetescertmanager.ReleaseArgs{
				Namespace: sandboxNs.Metadata.Name(),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
```

{{% /choosable %}}

{{% choosable language csharp %}}

```csharp
using Pulumi;
using Kubernetes = Pulumi.Kubernetes;
using KubernetesCertManager = Pulumi.KubernetesCertManager;

await Deployment.RunAsync(() =>
{
    // Create a sandbox namespace.
    var sandboxNs = new Kubernetes.Core.V1.Namespace("sandbox-ns", new()
    {
        Metadata = new Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs
        {
            Name = "sandbox",
        },
    });

    // Install cert-manager into our cluster.
    var certManager = new KubernetesCertManager.CertManager("cert-manager", new()
    {
        InstallCRDs = true,
        HelmOptions = new KubernetesCertManager.Inputs.ReleaseArgs
        {
            Namespace = sandboxNs.Metadata.Apply(metadata => metadata.Name),
        },
    });
});
```

{{% /choosable %}}

{{% choosable language java %}}

```java
import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.kubernetes.core.v1.Namespace;
import com.pulumi.kubernetes.core.v1.NamespaceArgs;
import com.pulumi.kubernetes.meta.v1.inputs.ObjectMetaArgs;
import com.pulumi.kubernetescertmanager.CertManager;
import com.pulumi.kubernetescertmanager.CertManagerArgs;
import com.pulumi.kubernetescertmanager.inputs.ReleaseArgs;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        // Create a sandbox namespace.
        var sandboxNs = new Namespace("sandbox-ns", NamespaceArgs.builder()
            .metadata(ObjectMetaArgs.builder()
                .name("sandbox")
                .build())
            .build());

        // Install cert-manager into our cluster.
        var certManager = new CertManager("cert-manager", CertManagerArgs.builder()
            .installCRDs(true)
            .helmOptions(ReleaseArgs.builder()
                .namespace(sandboxNs.metadata().applyValue(metadata -> metadata.name().get()))
                .build())
            .build());
    }
}
```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
resources:
  # Create a sandbox namespace.
  sandbox-ns:
    type: kubernetes:core/v1:Namespace
    properties:
      metadata:
        name: sandbox
  # Install cert-manager into our cluster.
  cert-manager:
    type: kubernetes-cert-manager:CertManager
    properties:
      installCRDs: true
      helmOptions:
        namespace: ${sandbox-ns.metadata.name}
```

{{% /choosable %}}

{{% choosable language hcl %}}

```hcl
terraform {
  required_providers {
    kubernetes = {
      source = "pulumi/kubernetes"
    }
    kubernetes-cert-manager = {
      source = "pulumi/kubernetes-cert-manager"
    }
  }
}

# Create a sandbox namespace.
resource "kubernetes_core_v1_namespace" "sandbox-ns" {
  metadata = {
    name = "sandbox"
  }
}

# Install cert-manager into our cluster.
resource "kubernetes-cert-manager_cert_manager" "cert-manager" {
  install_cr_ds = true
  helm_options = {
    namespace = kubernetes_core_v1_namespace.sandbox-ns.metadata.name
  }
}
```

{{% /choosable %}}

{{< /chooser >}}
