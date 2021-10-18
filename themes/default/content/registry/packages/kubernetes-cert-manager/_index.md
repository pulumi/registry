---
title: cert-manager
meta_desc: Use Pulumi's Component for cert-manager offers a Pulumi-friendly and strongly-typed way to manage cert-manager installations using infrastructure as code.
layout: overview
---

cert-manager is a Kubernetes add-on to automate the management and issuance of TLS certificates from various issuing sources.

It will ensure certificates are valid and up to date periodically, and attempt to renew certificates at an appropriate time before expiry. 

The Pulumi Component for cert-manager helps teams easily manage cert-manager installations as a package available in all Pulumi languages.

cert-manager was created by [Jetstack](https://jetstack.io) and is now a [CNCF Member Project](https://cert-manager.io/). 

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

{{< /chooser >}}

## SDK packages

The cert-manager provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/kubernetes-cert-manager`](https://www.npmjs.com/package/@pulumi/kubernetes-cert-manager)
* Python: [`pulumi-kubernetes-cert-manager`](https://pypi.org/project/pulumi-kubernetes-cert-manager/)
* Go: [`github.com/pulumi/pulumi-kubernetes-cert-manager/sdk/go/kubernetes`](https://github.com/pulumi/pulumi-kubernetes-cert-manager)
* .NET: [`Pulumi.KubernetesCertManager`](https://www.nuget.org/packages/Pulumi.KubernetesCertManager)
