---
title: "Typed CustomResources with Provider Extensions"
meta_desc: How to generate typed Pulumi SDKs for Kubernetes CustomResourceDefinitions by adding them as an extension of the kubernetes provider.
layout: package
---

> **Note:** Requires pulumi-kubernetes v4.34.0 or later.

Pulumi can manage Kubernetes [`CustomResources`](/registry/packages/kubernetes/api-docs/apiextensions/customresource/) directly, but those resources are untyped — every CustomResourceDefinition (CRD) has its own schema, so the provider cannot know the shape of `spec` ahead of time.

Adding your CRDs as an *extension* of the `kubernetes` provider solves this. The provider reads the CRD's OpenAPI schema, generates a typed SDK for it, and serves the resulting resources itself. You get IDE autocompletion and compile-time type checking without installing a separate code generation tool, and the extension's resources are managed by the same provider instance as the rest of your Kubernetes resources.

This is the recommended approach for typed custom resources. If you are using [`crd2pulumi`](/docs/integrations/clouds/kubernetes/crd2pulumi/) today, see [Migrating from crd2pulumi](#migrating-from-crd2pulumi) below.

## Adding an extension package

Point `pulumi package add` at one or more CRD manifests. The extension's parameters are passed as a single quoted string of `key=value` pairs:

```bash
pulumi package add kubernetes --extension "name=gateway-networking crd-manifest=gateway-api-crds.yaml"
```

This generates an SDK under `sdks/gateway-networking` and records the package in `Pulumi.yaml`, so subsequent `pulumi up` runs reconstruct the same schema without re-reading the manifest.

The supported parameters are:

| Parameter | Required | Description |
|---|---|---|
| `crd-manifest` | Yes | Path to a YAML or JSON file containing one or more CustomResourceDefinitions. Repeat the key to supply several files. |
| `name` | No | Name of the generated package. Defaults to the first manifest's file name without its extension, falling back to `crds`. |
| `version` | No | Version of the generated package. A CRD bundle has no inherent version, so this only needs to be stable. Defaults to `1.0.0`. |

A single manifest may contain multiple CRDs, and documents that are not CustomResourceDefinitions are skipped. Only CRD versions marked `served: true` are included in the generated SDK.

## Using the generated SDK

Extension resources are tokenized under the extension's package name rather than `kubernetes`, in the form `<name>:<group>/<version>:<Kind>`.

{{< chooser language "typescript,yaml" >}}

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

Extension resources can be used alongside the base provider's resources in the same program — they are served by the same provider instance, so no additional provider configuration is required.

`pulumi package add` prints language-specific instructions for importing the generated SDK when it completes.

## Migrating from crd2pulumi

Extension-served resources alias their `kubernetes:`-namespaced token. Existing state written by a `crd2pulumi`-generated SDK is therefore *adopted* rather than replaced, so switching does not destroy and recreate your custom resources.

To migrate:

1. Run `pulumi package add` against the same CRD manifests you passed to `crd2pulumi`.
2. Replace imports of the `crd2pulumi`-generated SDK with the new extension package.
3. Run `pulumi preview` and confirm the plan reports no replacements.
4. Remove the checked-in `crd2pulumi` output and the tool from your build.

## Notes

The generated SDK depends on the matching version of the base `kubernetes` SDK — for example, an extension generated against pulumi-kubernetes v4.34.0 depends on `@pulumi/kubernetes` v4.34.0. Keep the base provider SDK installed and in step with the provider version you generate against.
