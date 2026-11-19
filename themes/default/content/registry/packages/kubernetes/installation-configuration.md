---
title: Kubernetes Installation & Configuration
meta_desc: Install the Pulumi Kubernetes provider, connect it to a cluster, and configure how it manages resources.
layout: package
---

## Installation

{{< chooser language "typescript,python,go,csharp,java,yaml,hcl" >}}
{{% choosable language typescript %}}

npm:

```bash
npm install @pulumi/kubernetes
```

Yarn:

```bash
yarn add @pulumi/kubernetes
```

pnpm:

```bash
pnpm add @pulumi/kubernetes
```

Bun:

```bash
bun add @pulumi/kubernetes
```

{{% /choosable %}}
{{% choosable language python %}}

pip:

Add `pulumi-kubernetes` to `requirements.txt`, then run:

```bash
pip install -r requirements.txt
```

uv:

```bash
uv add pulumi-kubernetes
```

Poetry:

```bash
poetry add pulumi-kubernetes
```

{{% /choosable %}}
{{% choosable language go %}}

```bash
go get github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes
```

{{% /choosable %}}
{{% choosable language csharp %}}

```bash
dotnet add package Pulumi.Kubernetes
```

{{% /choosable %}}
{{% choosable language java %}}

Maven:

```xml
<dependency>
    <groupId>com.pulumi</groupId>
    <artifactId>kubernetes</artifactId>
    <version>4.34.3</version>
</dependency>
```

Gradle:

```groovy
implementation 'com.pulumi:kubernetes:4.34.3'
```

{{% /choosable %}}
{{% choosable language yaml %}}

```bash
pulumi package add kubernetes
```

{{% /choosable %}}
{{% choosable language hcl %}}

Declare the provider in your program, then run `pulumi install`:

```hcl
terraform {
  required_providers {
    kubernetes = {
      source  = "pulumi/kubernetes"
      version = "4.34.3"
    }
  }
}
```

```bash
pulumi install
```

{{% /choosable %}}
{{< /chooser >}}

## Connect to a cluster

The provider talks to the Kubernetes API server with [client-go](https://github.com/kubernetes/client-go), the same library `kubectl` uses, and finds its credentials the same way. It uses the first of these that's available:

1. The `kubeconfig` setting, which takes either the contents of a kubeconfig file or a path to one.
1. The files listed in the `KUBECONFIG` environment variable.
1. `~/.kube/config`.
1. The pod's service account, when Pulumi runs inside a cluster.

To check what the provider will see, run `kubectl config view` and `kubectl get pods`.

### Choose a context

A kubeconfig can define several contexts, each pairing a cluster with a user and a namespace. Without a `context` setting, the provider uses the kubeconfig's current context, so the same stack can deploy to a different cluster depending on who runs it. Set the context for the stack instead:

```bash
pulumi config set kubernetes:context my-context
```

To list the available contexts, run `kubectl config get-contexts`. To create one, run `kubectl config set-context my-context --cluster=my-cluster --user=my-user`.

Setting `strictMode` to `true` requires every stack to name its `kubeconfig` and `context`, and disables the default provider so every resource must use an explicit one.

### Use more than one cluster

The `kubernetes:` stack settings configure the default provider. To deploy to several clusters from one program, or to a cluster the same program creates, construct an [explicit provider](/docs/iac/concepts/providers/) for each cluster and pass its kubeconfig as the `kubeconfig` input.

### Keep credentials secret

The provider reads the ambient kubeconfig on the machine running Pulumi and connects to the API server directly, so Pulumi Cloud never receives it. A kubeconfig you pass through the `kubeconfig` setting or an explicit provider's `kubeconfig` input is different: Pulumi stores provider inputs in the stack's configuration and state. If you pass a file path, only the path is stored. If you pass the file's contents, all of it is stored, and the provider doesn't mark it as a secret.

Whether the contents need to be a secret depends on how the kubeconfig authenticates:

- **Credentials embedded in the file** need to be secret. That covers a client certificate's private key (`client-key-data`), a bearer `token`, a `username` and `password`, and the access and refresh tokens a legacy `auth-provider` entry caches. Set the value with `pulumi config set --secret kubernetes:kubeconfig`, or wrap it as a [secret](/docs/iac/concepts/secrets/) before passing it to an explicit provider.
- **An `exec` credential plugin** fetches short-lived credentials when the provider connects, so the file holds no secrets. The kubeconfigs that EKS, GKE, and AKS with [kubelogin](https://azure.github.io/kubelogin/) produce work this way. The rest of a kubeconfig, including the server address and the cluster's CA certificate, isn't secret. Check the plugin's `env` entries, though, which can carry credentials of their own.

[Pulumi ESC](/docs/esc/guides/integrate-with/kubernetes-cluster-access/) can manage kubeconfigs and short-lived cluster credentials centrally, for Pulumi programs as well as `kubectl` and `helm`.

## Configuration

Set these options for the default provider with `pulumi config set kubernetes:<option>`, or pass them as inputs to an explicit provider. The [Provider resource docs](/registry/packages/kubernetes/api-docs/provider/) also list `kubeClientSettings` and `helmReleaseSettings`, which are available only on explicit providers.

- `alwaysRender` (Optional) — When `true`, rewrites every manifest in `renderYamlToDirectory` on each update, even if the resource hasn't changed. Useful for tools such as the Argo CD Config Management Plugin, which need every manifest regenerated on each run. Only valid when `renderYamlToDirectory` is set.
- `cluster` (Optional) — The name of the kubeconfig cluster to use.
- `clusterIdentifier` (Optional) — Controls when the provider is replaced. When set, the provider, and every resource it manages, is replaced only when this value changes, and every other configuration change is an update. Without it, Pulumi uses heuristics to decide, because Kubernetes has no API for identifying a cluster, and those heuristics can replace the provider when an update would do.
- `context` (Optional) — The name of the kubeconfig context to use. Required when `strictMode` is `true`.
- `deleteUnreachable` (Optional) — When `true`, removes resources from the stack's state if their cluster can't be reached, so you can clean up after a cluster is deleted. May also be set with `PULUMI_K8S_DELETE_UNREACHABLE`.
- `enableConfigMapMutable` (Optional) — When `true`, changes to a ConfigMap's data update it in place instead of replacing it. Defaults to `false` for backward compatibility. May also be set with `PULUMI_K8S_ENABLE_CONFIGMAP_MUTABLE`.
- `enablePatchForce` (Optional) — When `true`, Server-Side Apply operations override field conflicts with other field managers. The `pulumi.com/patchForce` annotation on a resource takes precedence over this setting. May also be set with `PULUMI_K8S_ENABLE_PATCH_FORCE`. See [pulumi-kubernetes#2280](https://github.com/pulumi/pulumi-kubernetes/issues/2280).
- `enableSecretMutable` (Optional) — When `true`, changes to a Secret's data update it in place instead of replacing it. Defaults to `false` for backward compatibility. May also be set with `PULUMI_K8S_ENABLE_SECRET_MUTABLE`.
- `enableServerSideApply` (Optional) — Set to `false` to stop using [Server-Side Apply](#server-side-apply). Defaults to `true`. May also be set with `PULUMI_K8S_ENABLE_SERVER_SIDE_APPLY`.
- `kubeconfig` (Optional) — The contents of a kubeconfig file, or a path to one. Overrides `KUBECONFIG`. Required when `strictMode` is `true`. Set it with `--secret` if the file embeds credentials; see [Keep credentials secret](#keep-credentials-secret).
- `namespace` (Optional) — The default namespace for namespaced resources. A resource's own `metadata.namespace` takes precedence, and this setting takes precedence over the namespace of the active kubeconfig context. Ignored for cluster-scoped resources.
- `renderYamlToDirectory` (Optional) — Writes resource manifests to this directory instead of applying them to a cluster, and keeps them in sync with the program. Secret values are written in plain text, and outputs the cluster would compute, such as `status`, stay empty. The provider still tries to reach the cluster to tell whether custom resources are namespaced; if it can't, it writes them without a namespace and warns. In beta.
- `skipUpdateUnreachable` (Optional) — When `true`, skips updates to resources whose cluster can't be reached instead of failing. May also be set with `PULUMI_K8S_SKIP_UPDATE_UNREACHABLE`.
- `strictMode` (Optional) — When `true`, disables the default provider and requires `kubeconfig` and `context` on every provider, so each resource is tied to a specific cluster. Recommended for production stacks.
- `suppressDeprecationWarnings` (Optional) — When `true`, hides warnings about deprecated `apiVersion`s. May also be set with `PULUMI_K8S_SUPPRESS_DEPRECATION_WARNINGS`.
- `suppressHelmHookWarnings` (Optional) — When `true`, hides warnings about unsupported Helm hooks. May also be set with `PULUMI_K8S_SUPPRESS_HELM_HOOK_WARNINGS`.
- `upsertExistingObjects` (Optional) — When `true`, creating a resource that already exists in the cluster takes it over instead of failing. Defaults to `false`, because with Server-Side Apply, renaming a resource without an alias would otherwise update the existing object and then delete it. May also be set with `PULUMI_K8S_UPSERT_EXISTING_OBJECTS`.

## Server-Side Apply

The provider manages resources with [Server-Side Apply](https://kubernetes.io/docs/reference/using-api/server-side-apply/) (SSA) by default. SSA has the API server compute diffs and track which client owns each field, so Pulumi can share a resource with other controllers and tools without overwriting their changes. It also lets you patch resources that Pulumi didn't create, using the `Patch` resource that each resource type has, such as [`NamespacePatch`](/registry/packages/kubernetes/api-docs/core/v1/namespacepatch/).

Creating a resource that already exists in the cluster fails with an `AlreadyExists` error. To adopt existing objects, set `upsertExistingObjects` to `true`, or use a `Patch` resource to manage some of an object's fields without owning the object.

See [Managing resources with Server-Side Apply](/dev/tutorials/managing-resources-with-server-side-apply/) for a walkthrough.

## Annotations

These annotations control how Pulumi deploys and manages a resource. Set them in the resource's `metadata.annotations`.

### pulumi.com/skipAwait

When set to `"true"`, Pulumi creates or updates the resource without waiting for it to become ready.

{{% notes type="warning" %}}
Deployments, DaemonSets, StatefulSets, Pods, and Namespaces also skip waiting for deletion when this annotation is set. With Server-Side Apply, that can cause race conditions during replacement, and this behavior may change in a future release. To delete resources quickly, use `pulumi.com/deletionPropagationPolicy` instead.
{{% /notes %}}

{{% notes type="info" %}}
Pulumi doesn't know how to check readiness for every resource type, and treats many resources as ready as soon as they're created. If something depends on a resource being ready, set `pulumi.com/waitFor` on it, or set the `PULUMI_K8S_AWAIT_ALL=true` environment variable to have Pulumi wait for every resource.
{{% /notes %}}

### pulumi.com/waitFor

Defines when Pulumi considers the resource ready. It takes one of three forms:

- A `kubectl` [JSONPath](https://kubernetes.io/docs/reference/kubectl/jsonpath/) expression, prefixed with `jsonpath=` and optionally followed by a value. With a value, the resource is ready when the expression evaluates to that value. Without one, it's ready when any value exists at the path.

  ```yaml
  pulumi.com/waitFor: "jsonpath={.status.phase}=Running"
  pulumi.com/waitFor: "jsonpath={.webhooks[*].clientConfig.caBundle}"
  ```

- A condition, prefixed with `condition=` and optionally followed by a status, which defaults to `True`. This matches `kubectl wait --for=condition=...`.

  ```yaml
  pulumi.com/waitFor: "condition=Synced"
  pulumi.com/waitFor: "condition=Reconciling=False"
  ```

- A JSON array of `jsonpath=` and `condition=` expressions. The resource is ready when all of them are met at once.

  ```yaml
  pulumi.com/waitFor: '["jsonpath={.foo}", "condition=Bar"]'
  ```

This annotation has no effect if `pulumi.com/skipAwait` is `"true"`.

### pulumi.com/timeoutSeconds

The number of seconds Pulumi waits for the resource to become ready. Prefer the [`customTimeouts`](/docs/iac/concepts/resources/options/customtimeouts/) resource option.

### pulumi.com/replaceUnready

If the resource didn't become ready during the previous update, replace it instead of waiting for it again. Only `batch/v1/Job` supports this annotation.

### pulumi.com/deletionPropagationPolicy

Sets how Kubernetes deletes the resource's dependents:

- `foreground` (default) — Deletes the resource and all of its dependents, and waits for them. Slower, but everything is cleaned up when it succeeds.
- `background` — Deletes the resource and lets Kubernetes garbage-collect its dependents asynchronously. Faster, but dependents that aren't finalized can remain.
- `orphan` — Deletes the resource and leaves its dependents in place, for example to keep them for migration or debugging.

See [Cascading deletion](https://kubernetes.io/docs/concepts/architecture/garbage-collection/#cascading-deletion) in the Kubernetes docs.

### pulumi.com/patchForce

When set to `"true"`, Server-Side Apply overrides field conflicts for this resource. The `enablePatchForce` setting does the same for every resource.

### pulumi.com/patchFieldManager

The field manager name Pulumi uses for Server-Side Apply operations on this resource.

### Annotations Pulumi sets

The provider writes these annotations to resources it manages:

- `pulumi.com/autonamed` — Pulumi generated the resource's name, because `metadata.name` wasn't set.
- `app.kubernetes.io/managed-by` — Pulumi manages the resource. Not set when Server-Side Apply is enabled.
