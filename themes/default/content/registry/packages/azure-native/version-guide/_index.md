---
title: Azure Native version guide
meta_desc: How Azure API versions are represented in the native Azure provider for Pulumi.
layout: package
---

The native Azure provider SDKs follow the semantic versioning convention, similarly to other Pulumi providers. At the same time, Azure API is structured around date-based API versions like `2020-03-01` defined per Azure resource provider.

This guide describes how we combine both versioning approaches to achieve the following design goals:

- Provide access to the entire API surface of Azure resources.
- Ship new updates as soon as Microsoft publishes them.
- Provide stable SDKs without unexpected breaking changes.
- Keep the SDK size manageable.

## Azure Native Semantic Versioning

Releases of the Azure Native SDKs follow the semantic versioning convention, where each version is identified with three numbers: `major.minor.patch` (for example, `1.2.1`).

- Patch versions contain bug fixes only.
- Minor versions may ship fixes and new functionality that doesn't break existing code: new resources, functions, properties.
- Major versions include breaking changes, where existing user code may need to be modified.

Microsoft has defined [Breaking changes guidelines](https://github.com/Azure/azure-rest-api-specs/blob/master/documentation/Breaking%20changes%20guidelines.md) for all the Azure API specifications. The native Azure provider generates SDKs from the Open API specifications, so we rely on Microsoft to follow their breaking changes guideline.

Occasionally, Microsoft ships breaking changes. For example, they may rename a misspelled property in the spec to match it with the actual service behavior. We treat these occurrences as bug fixes and ship them in patch or minor provider versions without revving the major version.

## Azure Default API Versions

Azure's APIs are versioned. API versions are defined per Azure service (resource provider) and are published frequently. API versions are based on a date, for instance, `2020-03-01` or `2019-05-15-preview`. Breaking changes may occur between API versions.

The Pulumi provider chooses a set of default API versions for each release. For the resources in the Azure Native SDKs, these are the API versions that the resources are generated from and that the provider will use for its requests to Azure. The default API version is often, but not always, the same for all resources in a resource provider. The [API reference docs](/registry/packages/azure-native/api-docs/) describe these resources.

Following the semantic versioning outlined above, patch versions will not contain API version changes. Minor versions contain API version changes when we determine that they don't cause breaking changes, as far as we can tell from the Azure specification. Therefore, regular minor version upgrades allow users to stay current with Azure API changes while minimizing the risk of breaking changes.

TODO how we choose the default API versions

## Accessing any API Version via Local Packages

We expect most users to use the standard resources with the default API versions most of the time, without the need to specify an explicit API version. There are cases, however, where you may want to use a specific API version:
- you have a resource on an older API version and cannot migrate yet,
- the default version has an issue,
- you want to pin an API version for extra stability , or
- you need a new preview version for access to bleeding edge features.

For such cases, the Pulumi provider allows you to generate a local SDK for a specific API version and use it in your project. This is enabled by our new Local Packages feature, which lets you generate Pulumi Packages locally into your Pulumi project, instead of relying only on packages published to the Pulumi Registry.

A single command generates a local package for a specific API version of an Azure resource provider:

```bash
pulumi package add azure-native storage v20240101
```

This command generates a local package for the `storage` resource provider with the API version `2024-01-01` and adds it to your project.

You can then update your Pulumi program to use the new, local SDK:

{{< chooser language "typescript,python,csharp,go,java,yaml" >}}

{{% choosable language typescript %}}

```typescript
import * as storage_v20240101 from "@pulumi/azure-native_storage_v20240101";

const storageAccount = new storage_v20240101.storage.v20240101.StorageAccount(...)
```

{{% /choosable %}}

{{% choosable language python %}}

```python
TODO
```

{{% /choosable %}}

{{% choosable language csharp %}}

```csharp
TODO
```

{{% /choosable %}}

{{% choosable language go %}}

```go
TODO
```

{{% /choosable %}}

{{% choosable language java %}}

TODO Java is not supported.

{{% /choosable %}}


{{% choosable language yaml %}}

TODO
```yaml
resources:
  my-cluster:
   type: "azure-native:containerservice/v20220301:ManagedCluster"
   properties:
     # ...
```

{{% /choosable %}}

{{< /chooser >}}

## Using the generic resource

The generic resource [azure-native.resources.Resource](../api-docs/resources/resource/) is a special resource that allows you to access any Azure resource at any API version. The downside is that it is not explicitly typed and type-safe like the specific resources. It's useful when you need to access a resource that is not yet supported by the Pulumi provider or when an issue prevents you from using a specific resource.

The example below creates a storage account using the generic resource.

{{< chooser language "typescript,python,csharp,go,java,yaml" >}}

{{% choosable language typescript %}}

```typescript
TODO
```

{{% /choosable %}}

{{% choosable language python %}}

```python
TODO
```

{{% /choosable %}}

{{% choosable language csharp %}}

```csharp
TODO
```

{{% /choosable %}}

{{% choosable language go %}}

```go
TODO
```

{{% /choosable %}}

{{% choosable language java %}}

TODO

{{% /choosable %}}


{{% choosable language yaml %}}

TODO
```yaml
```

{{% /choosable %}}

{{< /chooser >}}