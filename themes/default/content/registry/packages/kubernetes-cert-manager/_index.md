---
title: Jetstack Cert Manager
meta_desc: Use Pulumi's Component for offers a Pulumi-friendly and strongly-typed way to manage Cert Manager installations using infrastructure as code.
layout: overview
---

Easily manage Jetstack Cert Manager installations as a package available in all Pulumi languages.

Example:

{{< chooser language "typescript,python" >}}

{{% choosable language typescript %}}

```typescript
import * as k8s from "@pulumi/kubernetes";
import * as certmanager from "@pulumi/kubernetes-cert-manager";

// Create a sandbox namespace.
const nsName = "sandbox";
const ns = new k8s.core.v1.Namespace("sandbox-ns", {
    metadata: { name: nsName },
});

// Install a cert manager into our cluster.
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

# Install a cert manager into our cluster.
manager = CertManager('cert-manager',
                      install_crds=True,
                      helm_options=ReleaseArgs(
                          namespace=ns_name,
                      ))
```

{{% /choosable %}}

{{< /chooser >}}

## SDK packages

The Jetstack Cert Manager provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/kubernetes-cert-manager`](https://www.npmjs.com/package/@pulumi/kubernetes-cert-manager)
* Python: [`pulumi-kubernetes-cert-manager`](https://pypi.org/project/pulumi-kubernetes-cert-manager/)
* Go: [`github.com/pulumi/pulumi-kubernetes-cert-manager/sdk/go/kubernetes`](https://github.com/pulumi/pulumi-kubernetes-cert-manager)
* .NET: [`Pulumi.KubernetesCertManager`](https://www.nuget.org/packages/Pulumi.KubernetesCertManager)
