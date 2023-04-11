---
title: Kubernetes Installation & Configuration
meta_desc: Provides an overview on how to setup the Kubernetes Provider for Pulumi.
layout: installation
---

<!-- markdownlint-disable url -->
[pulumi-kubernetes-provider]: ./
[server-side-apply]: ./installation-configuration/#server-side-apply
[client-go]: https://github.com/kubernetes/client-go
[gke-tutorial]: /registry/packages/kubernetes/how-to-guides/gke
[eks-tutorial]: /registry/packages/kubernetes/how-to-guides/eks
[aks-tutorial]: /registry/packages/kubernetes/how-to-guides/aks
[Heptio AWS quickstart]: https://aws.amazon.com/quickstart/architecture/heptio-kubernetes/
[provider-args]: /registry/packages/kubernetes/api-docs/provider
[provider-kubeconfig]: /registry/packages/kubernetes/api-docs/provider#inputs
[kubeconfig]: https://kubernetes.io/docs/tasks/access-application-cluster/configure-access-multiple-clusters/
[install]: /docs/get-started/install
[nodejs]: https://nodejs.org/en/
[npm]: https://www.npmjs.com/get-npm
[yarn]: https://yarnpkg.com/en/docs/install
<!-- markdownlint-enable url -->

Your Pulumi program is required to import the [pulumi/kubernetes][pulumi-kubernetes-provider] provider package to allow the Pulumi CLI to authenticate and interact with a running Kubernetes cluster.

By default, Pulumi will use a local [kubeconfig] if available, or one can be passed as a [provider argument][provider-kubeconfig] in the request.

With the `kubeconfig` available, Pulumi communicates with the API Server using the official Kubernetes [client-go] library, just like `kubectl` does.

## Installation

The Kubernetes provider is available as a package in all Pulumi languages:

- JavaScript/TypeScript: [`@pulumi/kubernetes`](https://www.npmjs.com/package/@pulumi/kubernetes)
- Python: [`pulumi-kubernetes`](https://pypi.org/project/pulumi-kubernetes/)
- Go: [`github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes`](https://github.com/pulumi/pulumi-kubernetes)
- .NET: [`Pulumi.Kubernetes`](https://www.nuget.org/packages/Pulumi.Kubernetes)

## Pre-Requisites

If you do not have a cluster set up and running yet, you'll need to do the
following steps.

1. [Install the Pulumi CLI][install].
1. Install a package manager for your Pulumi program language runtime, such as [npm] or [Yarn] for [Node.js][nodejs], or PyPI for Python.
1. Provision a Kubernetes cluster. For a new managed Kubernetes cluster, check out the [guides.](/registry/packages/kubernetes/how-to-guides)
1. Download [`kubectl`](https://kubernetes.io/docs/tasks/tools/install-kubectl/).

## Setup

By default, Pulumi will look for a kubeconfig file in the following locations,
just like `kubectl`:

- The environment variable: `$KUBECONFIG`,
- Or in current user's default kubeconfig directory: `~/.kube/config`

If the kubeconfig file is not in either of these locations, Pulumi will **not** find it, and it will
fail to authenticate against the cluster. Set one of these locations to a valid kubeconfig file, if you have not done so
already.

Verify the cluster is configured and up by running `kubectl get pods`.

Then follow the [getting started guide](/docs/get-started/kubernetes) or explore some [code examples](https://github.com/pulumi/examples#kubernetes).

Note: Pulumi **never** sends **any** authentication [secrets](/docs/intro/concepts/secrets/) or credentials to the Pulumi Cloud.

### Contexts

The kubeconfig file defines some number of _contexts_. Each context is a name that is associated
with a _cluster_, _namespace_, and a "_user_" (a local-only name that's associated with a credential
that allows access to the cluster).

Your Kubernetes implementation may have already written out an appropriate kubeconfig. 
For example, `minikube start` does this, unless you specified `--keep-context=true`.
Moreover, multiple sources of kubeconfig are merged and the result may surprise you.
Therefore, check your current kubeconfig using:

```shell
$ kubectl \
    config \
        view
```

To create a context, for example, you can run the `kubectl set-context` command as follows:

```shell
$ kubectl config \
    set-context my-context \
    --cluster=my-cluster \
    --user=my-user
```

Mind that e.g., `minikube` uses the same name for the context and the cluster in it.

If you have completed your kubeconfig configuration, and are using the default kubeconfig file, you will be able to set the
configuration variable `kubernetes:context` in the Kubernetes provider to the given context name:

```shell
$ pulumi stack init new-kube-stack
$ pulumi config set kubernetes:context my-context
```

If you don't want to select a context, you can always make it the operating system user-wide default:

```shell
$ kubectl config \
    use-context my-context
```

> **Note:** Depending on a default context is a bad idea if you're going to share your stack with
> others; it makes your stack dependent on ambient information not known to Pulumi, an anti-pattern
> that leads to unrepeatable deployments.

Additionally, the Kubernetes provider accepts many [configuration settings][provider-args].

These can be provided to the default Kubernetes provider via `pulumi config set kubernetes:<option>`, or passed
to the constructor of a `new kubernetes.Provider` to construct a specific instance of the Kubernetes provider for your requests.

### kubectl proxy

Each Kubernetes resource managed by Pulumi will have a link in the [Pulumi Cloud](https://app.pulumi.com")
to view the resource in the cluster. These links are local, and require the client run `kubectl proxy` beforehand to access the resource.

To learn more about `kubectl proxy` check out the [reference docs](https://kubernetes.io/docs/concepts/cluster-administration/proxies/).

## Configuration

The Kubernetes provider accepts the following configuration settings. These can be provided to the default Kubernetes provider via `pulumi config set kubernetes:<option>`, or passed to the constructor of `new kubernetes.Provider` to construct a [specific instance](https://www.pulumi.com/docs/intro/concepts/resources/providers/#explicit-provider-configuration) of the Kubernetes provider.

See the [Provider configuration][provider-args] docs for a complete list of options.

## Server-Side Apply

[Server-Side Apply](https://kubernetes.io/docs/reference/using-api/server-side-apply/) (SSA) is a resource management strategy that was introduced in Kubernetes `v1.18`. Clients using SSA can safely share the management of Kubernetes resources by making the API Server responsible for computing diffs and resolving conflicts.

The [v3.20.0 release](https://github.com/pulumi/pulumi-kubernetes/releases/tag/v3.20.0) of the Pulumi Kubernetes provider added support for managing resources using SSA. This feature is currently opt-in using the `enableServerSideApply` [Provider option][provider-args], but will become the default in the next major release. Using SSA provides the following benefits:

1. Kubernetes resources may be safely managed by more than one controller.
2. It is now possible to "Upsert" resources; create the resource if it does not exist, or apply the configuration to an existing resource.
3. It is now possible to patch resources with the Patch resource types in the SDK. Each resource type in the SDK has a corresponding Patch resource.
4. The `last-applied-configuration` annotation is no longer used.

## Annotations

A few Pulumi-specific annotations can be applied to Kubernetes resources managed by Pulumi to control aspects of how Pulumi deploys and manages the Kubernetes resource:

- `pulumi.com/patchFieldManager`: Server-Side Apply option: Specify the `FieldManager` name to use for the Server-Side Apply operation.
- `pulumi.com/patchForce`: Server-Side Apply option: Force override any conflicts for the specified resource.
- `pulumi.com/replaceUnready`: If the resource failed to become ready in the previous Pulumi update, replace the resource rather than continuing to wait for it to become ready. Only `batch/v1/Job` currently supports this annotation.
- `pulumi.com/skipAwait`: Disables Pulumi's default await logic that waits for a Kubernetes resource to become "ready" before marking the resource as having created or updated succesfully.
- `pulumi.com/timeoutSeconds`: Specifies the number of seconds that the Pulumi Kubernetes provider will wait for the resource to become "ready".

In addition, the Pulumi provider may write the following annotations onto resources it manages:

- `app.kubernetes.io/managed-by`: Indicates the controller managing a Kubernetes resource. This annotation is not set when using the Server-Side Apply mode.
- `pulumi.com/autonamed`: Indicates that the Pulumi Kubernetes provider decided to autoname the resource (instead of using an explicitly provided `metadata.name`).
