---
title: "Typed CustomResources with Provider Extensions"
meta_desc: How to generate typed Pulumi SDKs for Kubernetes CustomResourceDefinitions by adding them as an extension of the kubernetes provider.
layout: package
---

> **Note:** Requires pulumi-kubernetes v4.34.0 or later.

Pulumi can manage Kubernetes [`CustomResources`](/registry/packages/kubernetes/api-docs/apiextensions/customresource/) directly, but they are untyped. Every CustomResourceDefinition (CRD) has its own schema, so the provider cannot know the shape of `spec` in advance.

Adding your CRDs as an *extension* of the `kubernetes` provider solves this. The provider reads the CRD's OpenAPI schema, generates a typed SDK for it, and serves the resulting resources itself. You get IDE autocompletion and compile-time type checking without installing a separate code generation tool, and the extension's resources are managed by the same provider instance as the rest of your Kubernetes resources.

This is the recommended approach for typed custom resources. If you are using [`crd2pulumi`](/docs/integrations/clouds/kubernetes/crd2pulumi/) today, see [Migrating from crd2pulumi](#migrating-from-crd2pulumi) below.

## Usage

Place a CRD manifest in your program. The examples below use the [Gateway API](https://gateway-api.sigs.k8s.io/) CRDs.

```bash
pulumi package add kubernetes --extension "name=gateway-networking crd-manifest=gateway-api-crds.yaml"
```

This generates an SDK under `sdks/gateway-networking` and records the package in `Pulumi.yaml`, so subsequent `pulumi up` runs reconstruct the same schema without re-reading the manifest.

The generated SDK draws shared types such as `ObjectMeta` from the base `kubernetes` package, so your project also needs the `pulumi-kubernetes` SDK as a dependency.

The supported parameters are:

| Parameter | Required | Description |
|---|---|---|
| `crd-manifest` | Yes | Path to a YAML or JSON file containing one or more CustomResourceDefinitions. Repeat the key to supply several files. |
| `name` | No | Name of the generated package. Defaults to the first manifest's file name without its extension, falling back to `crds`. |
| `version` | No | Version of the generated package. A CRD bundle has no inherent version, so any value works as long as it stays stable across runs. Defaults to `1.0.0`. |

A single manifest may contain multiple CRDs, and documents that are not CustomResourceDefinitions are skipped. To supply several manifests, repeat the `crd-manifest` key inside the same quoted string:

```bash
pulumi package add kubernetes --extension "name=gateway-networking crd-manifest=gateway-api-crds.yaml crd-manifest=extra-crds.yaml"
```

## Using the generated SDK

Extension resources are tokenized under the extension's package name, in the form `<name>:<group>/<version>:<Kind>`. In the generated SDKs the module name is the group's first segment, so `gateway.networking.k8s.io/v1` becomes the `gateway.v1` module.

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}

{{% choosable language typescript %}}

```typescript
import * as gatewayNetworking from "@pulumi/gateway-networking";

const gatewayClass = new gatewayNetworking.gateway.v1.GatewayClass("example", {
    metadata: {
        name: "example-class",
    },
    spec: {
        controllerName: "example.com/gateway-controller",
    },
});
```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi_gateway_networking as gateway_networking

gateway_class = gateway_networking.gateway.v1.GatewayClass(
    "example",
    metadata={
        "name": "example-class",
    },
    spec={
        "controller_name": "example.com/gateway-controller",
    },
)
```

{{% /choosable %}}

{{% choosable language go %}}

The import path is rooted at the module name in the generated SDK's `go.mod`, which defaults to the extension name. Below that root, packages are laid out as `kubernetes/<group>/<version>`.

```go
package main

import (
	gatewayv1 "gateway-networking/kubernetes/gateway/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := gatewayv1.NewGatewayClass(ctx, "example", &gatewayv1.GatewayClassArgs{
			Metadata: &metav1.ObjectMetaArgs{
				Name: pulumi.String("example-class"),
			},
			Spec: &gatewayv1.GatewayClassSpecArgs{
				ControllerName: pulumi.String("example.com/gateway-controller"),
			},
		})
		return err
	})
}
```

{{% /choosable %}}

{{% choosable language csharp %}}

```csharp
using Pulumi;
using Pulumi.GatewayNetworking.Gateway.V1;
using Pulumi.GatewayNetworking.Types.Inputs.Gateway.V1;
using Pulumi.Kubernetes.Types.Inputs.Meta.V1;

return await Deployment.RunAsync(() =>
{
    var gatewayClass = new GatewayClass("example", new GatewayClassArgs
    {
        Metadata = new ObjectMetaArgs
        {
            Name = "example-class",
        },
        Spec = new GatewayClassSpecArgs
        {
            ControllerName = "example.com/gateway-controller",
        },
    });
});
```

{{% /choosable %}}

{{% choosable language java %}}

```java
import com.pulumi.Pulumi;
import com.pulumi.gatewaynetworking.gateway.v1.GatewayClass;
import com.pulumi.gatewaynetworking.gateway.v1.GatewayClassArgs;
import com.pulumi.gatewaynetworking.gateway.v1.inputs.GatewayClassSpecArgs;
import com.pulumi.kubernetes.meta.v1.inputs.ObjectMetaArgs;

public class App {
    public static void main(String[] args) {
        Pulumi.run(ctx -> {
            var gatewayClass = new GatewayClass("example", GatewayClassArgs.builder()
                .metadata(ObjectMetaArgs.builder()
                    .name("example-class")
                    .build())
                .spec(GatewayClassSpecArgs.builder()
                    .controllerName("example.com/gateway-controller")
                    .build())
                .build());
        });
    }
}
```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
resources:
  test-gateway-class:
    type: gateway-networking:gateway.networking.k8s.io/v1:GatewayClass
    properties:
      metadata:
        name: example-class
      spec:
        controllerName: example.com/gateway-controller
```

{{% /choosable %}}

{{< /chooser >}}

## Migrating from crd2pulumi

The provider registers aliases for the crd2pulumi-style `kubernetes:`-namespaced tokens, so existing resources are adopted rather than replaced.

Generate the extension package from the same CRD manifests you passed to `crd2pulumi`:

```bash
pulumi package add kubernetes --extension "name=gateway-networking crd-manifest=gateway-api-crds.yaml"
```

Then update your imports. Resource names, property names, and nesting are unchanged, so your call sites stay as they are.

{{< chooser language "typescript,python,go,csharp,java" >}}

{{% choosable language typescript %}}

```typescript
// Before
import * as crds from "./crds/nodejs";
const gatewayClass = new crds.gateway.v1.GatewayClass("example", { /* ... */ });

// After
import * as gatewayNetworking from "@pulumi/gateway-networking";
const gatewayClass = new gatewayNetworking.gateway.v1.GatewayClass("example", { /* ... */ });
```

{{% /choosable %}}

{{% choosable language python %}}

```python
# Before
import pulumi_crds as crds
gateway_class = crds.gateway.v1.GatewayClass("example", ...)

# After
import pulumi_gateway_networking as gateway_networking
gateway_class = gateway_networking.gateway.v1.GatewayClass("example", ...)
```

{{% /choosable %}}

{{% choosable language go %}}

```go
// Before
import gatewayv1 "example.com/myproject/crds/go/kubernetes/gateway/v1"

// After
import gatewayv1 "gateway-networking/kubernetes/gateway/v1"
```

{{% /choosable %}}

{{% choosable language csharp %}}

```csharp
// Before
using Pulumi.Crds.Gateway.V1;

// After
using Pulumi.GatewayNetworking.Gateway.V1;
```

{{% /choosable %}}

{{% choosable language java %}}

The two generators lay Java packages out differently, so the group and version segments change shape as well as the package root.

```java
// Before
import com.pulumi.crds.gateway.networking.k8s.io_v1.GatewayClass;

// After
import com.pulumi.gatewaynetworking.gateway.v1.GatewayClass;
```

{{% /choosable %}}

{{< /chooser >}}

Run `pulumi preview` and confirm the plan reports no replacements, then delete the checked-in `crd2pulumi` output and remove the tool from your build.
